package scrapers

import (
	"math/big"

	"github.com/diadata-org/fair-value/models"
	"github.com/diadata-org/fair-value/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/diadata-org/fair-value/contracts/erc20"
	getters "github.com/diadata-org/fair-value/contracts/usdp/getters"
	redeemer "github.com/diadata-org/fair-value/contracts/usdp/redeemer"
)

const (
	// TO DO: Determine total supply (shares) instead.
	SCALING = 1e18
)

// ------------------------------------------------------------------
// CONTRACT EXCHANGE RATE
// ------------------------------------------------------------------

type USDPScraper struct {
	BaseScraper
	client          *ethclient.Client
	blockchain      string
	contractAddress string
	config          models.FeedConfig
}

func NewUSDPScraperScraper(config models.FeedConfig, metacontractData models.MetacontractData) *USDPScraper {
	scraper := USDPScraper{
		BaseScraper:     NewBaseScraper(metacontractData),
		blockchain:      config.Blockchain,
		contractAddress: config.Address,
		config:          config,
	}
	client, err := ethclient.Dial(utils.Getenv("RPC_NODE_USDP", ""))
	if err != nil {
		log.Errorf("USDp -- get ETH client for %s: %v", scraper.config.Symbol, err)
		return nil
	}
	scraper.client = client
	return &scraper
}

func (scraper *USDPScraper) TotalUnderlying() (totalUnderlying *big.Int, totalValueUnderlying *big.Int, err error) {
	addr := common.HexToAddress(scraper.config.Address)

	redeemerCaller, err := redeemer.NewUsdpredeemerCaller(addr, scraper.client)
	if err != nil {
		return
	}

	totalShares, err := scraper.TotalShares()
	if err != nil {
		return
	}

	redemptionCurve, err := redeemerCaller.QuoteRedemptionCurve(&bind.CallOpts{}, totalShares)
	if err != nil {
		return
	}
	collateralTokens := redemptionCurve.Tokens

	gettersCaller, err := getters.NewUsdpgettersCaller(addr, scraper.client)
	if err != nil {
		return
	}

	// Iterate over collateral tokens
	totalValueUnderlying = big.NewInt(0)
	for i, token := range collateralTokens {

		oracleValues, errGetters := gettersCaller.GetOracleValues(&bind.CallOpts{}, token)
		if errGetters != nil {
			err = errGetters
			return
		}

		amount := new(big.Float).Quo(big.NewFloat(0).SetInt(redemptionCurve.Amounts[i]), big.NewFloat(SCALING))
		redemption := oracleValues.Redemption
		price, _ := new(big.Float).Mul(amount, big.NewFloat(0).SetInt(redemption)).Int(nil)
		log.Debug("USDp -- price: ", price.String())
		log.Debugf("USDp -- token --- amount -- price: %s -- %s -- %s", token.Hex(), amount.String(), redemption.String())
		totalValueUnderlying.Add(totalValueUnderlying, price)
	}
	return
}

func (scraper *USDPScraper) TotalShares() (*big.Int, error) {

	gettersCaller, err := getters.NewUsdpgettersCaller(common.HexToAddress(scraper.config.Address), scraper.client)
	if err != nil {
		return nil, err
	}

	tokenP, err := gettersCaller.TokenP(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	erc20Caller, err := erc20.NewERC20Caller(tokenP, scraper.client)
	if err != nil {
		return nil, err
	}

	return erc20Caller.TotalSupply(&bind.CallOpts{})

}

func (scraper *USDPScraper) DataChannel() chan models.FairValueData {
	return scraper.dataChannel
}

func (scraper *USDPScraper) Close() chan bool {
	return scraper.BaseScraper.Close()
}

func (scraper *USDPScraper) GetConfig() models.FeedConfig {
	return scraper.config
}
