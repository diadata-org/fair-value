package scrapers

import (
	"math/big"

	"github.com/diadata-org/fair-value/models"
	"github.com/diadata-org/fair-value/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

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

func NewUSDPScraperScraper(config models.FeedConfig) *USDPScraper {
	scraper := USDPScraper{
		BaseScraper:     NewBaseScraper(),
		blockchain:      config.Blockchain,
		contractAddress: config.Address,
		config:          config,
	}
	client, err := ethclient.Dial(utils.Getenv("RPC_NODE_USDP", ""))
	if err != nil {
		log.Fatalf("get ETH client for %s: %v", scraper.config.Symbol, err)
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

	redemptionCurve, err := redeemerCaller.QuoteRedemptionCurve(&bind.CallOpts{}, big.NewInt(SCALING))
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
		log.Debug("price: ", price.String())
		log.Debugf("token --- value -- redemption: %s -- %s -- %s", token.Hex(), amount.String(), redemption.String())
		totalValueUnderlying.Add(totalValueUnderlying, price)
	}
	return
}

func (scraper *USDPScraper) TotalShares() (*big.Int, error) {
	return big.NewInt(SCALING), nil
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

func (scraper *USDPScraper) getUSDPExecutionResult() (result *big.Int, err error) {

	addr := common.HexToAddress(scraper.config.Address)

	redeemerCaller, err := redeemer.NewUsdpredeemerCaller(addr, scraper.client)
	if err != nil {
		return
	}

	redemptionCurve, err := redeemerCaller.QuoteRedemptionCurve(&bind.CallOpts{}, big.NewInt(SCALING))
	if err != nil {
		return
	}
	collateralTokens := redemptionCurve.Tokens

	gettersCaller, err := getters.NewUsdpgettersCaller(addr, scraper.client)
	if err != nil {
		return
	}

	// Iterate over collateral tokens
	result = big.NewInt(0)
	for i, token := range collateralTokens {

		oracleValues, errGetters := gettersCaller.GetOracleValues(&bind.CallOpts{}, token)
		if errGetters != nil {
			err = errGetters
			return
		}

		amount := new(big.Float).Quo(big.NewFloat(0).SetInt(redemptionCurve.Amounts[i]), big.NewFloat(SCALING))
		redemption := oracleValues.Redemption
		price, _ := new(big.Float).Mul(amount, big.NewFloat(0).SetInt(redemption)).Int(nil)
		log.Info("price: ", price.String())
		log.Infof("token --- value -- redemption: %s -- %s -- %s", token.Hex(), amount.String(), redemption.String())
		result.Add(result, price)
	}

	// // Prepare readonly call options
	// callOpts := &bind.CallOpts{Context: context.Background()}

	// // --- 1. Call quoteRedemptionCurve(1e8) ---
	// amount := big.NewInt(1e8)
	// tokens, amounts, err := redeemerContract.QuoteRedemptionCurve(callOpts, amount)
	// if err != nil {
	// 	log.Fatal("QuoteRedemptionCurve failed:", err)
	// }

	// // --- 2. Loop over collateral tokens ---
	// usdpPrice := big.NewFloat(0)

	// for i, token := range tokens {
	// 	value := new(big.Float).Quo(new(big.Float).SetInt(amounts[i]), big.NewFloat(1e8))

	// 	// Get oracle values for this collateral
	// 	oracleVals, err := gettersContract.GetOracleValues(callOpts, token)
	// 	if err != nil {
	// 		log.Fatalf("GetOracleValues failed for token %s: %v", token.Hex(), err)
	// 	}

	// 	// Redemption value is the 5th element
	// 	collateralPrice := new(big.Float).Quo(new(big.Float).SetInt(oracleVals.Redemption), big.NewFloat(1e18))

	// 	redemptionUsdValue := new(big.Float).Mul(value, collateralPrice)
	// 	usdpPrice.Add(usdpPrice, redemptionUsdValue)
	// }

	// fmt.Printf("USDp Price: %f USD\n", usdpPrice)
	return

}
