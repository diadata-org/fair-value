package scrapers

import (
	"math"
	"math/big"

	"github.com/diadata-org/fair-value/models"
	"github.com/diadata-org/fair-value/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	hohm "github.com/diadata-org/fair-value/contracts/hohm"
)

// ------------------------------------------------------------------
// NET ASSET VALUE
// ------------------------------------------------------------------

// const (
// 	HOHM_ADDRESS = "0x1db1591540d7a6062be0837ca3c808add28844f6"
// )

type HohmScraper struct {
	BaseScraper
	blockchain      string
	contractAddress string
	client          *ethclient.Client
	hohmCaller      *hohm.HohmCaller
	config          models.FeedConfig
}

func NewHohmScraper(config models.FeedConfig) *HohmScraper {

	client, err := ethclient.Dial(utils.Getenv("RPC_NODE_HOHM", ""))
	if err != nil {
		log.Errorf("hOHM -- make eth client for %s: %v", config.Symbol, err)
		return nil
	}
	hohmCaller, err := hohm.NewHohmCaller(common.HexToAddress(config.Address), client)
	if err != nil {
		log.Errorf("hOHM -- NewHohmCaller: %v", err)
		return nil
	}

	scraper := HohmScraper{
		BaseScraper:     NewBaseScraper(),
		blockchain:      config.Blockchain,
		contractAddress: config.Address,
		client:          client,
		hohmCaller:      hohmCaller,
		config:          config,
	}

	return &scraper
}

func (scraper *HohmScraper) Assets() (assetValueUSD *big.Int, native bool, err error) {
	tokens, balances, err := scraper.getHohmTokensBalances()
	if err != nil {
		return
	}

	assetValueUSD = new(big.Int)
	for i, token := range tokens.Assets {
		// log.Infof("hOHM -- asset address: %s", token.Hex())
		supply := balances.TotalAssets[i]
		// TO DO: Get price from lumina metacontract.
		var price float64
		price, err = utils.GetDiaQuotationPrice("Ethereum", token.Hex())
		if err != nil {
			log.Errorf("hOHM -- failed to fetch price for %s: %v", token.Hex(), err)
			return
		}
		log.Debugf("hOHM -- assets supply -- price: %s -- %v", supply.String(), price)

		// Scale price and supply to big.Float.
		priceBig := big.NewFloat(0).Mul(big.NewFloat(price), new(big.Float).SetFloat64(math.Pow10(int(DECIMALS))))
		supplyFloat := big.NewFloat(0).SetInt(supply)

		// Decimals of token.
		decimals, err := utils.GetEthDecimals(token, scraper.client)
		if err != nil {
			log.Errorf("hOHM -- GetEthDecimals for %s: %v", token.Hex(), err)
			continue
		}

		// Scale price*supply to @DECIMALS decimals.
		resultBig := big.NewFloat(0).Quo(big.NewFloat(0).Mul(priceBig, supplyFloat), new(big.Float).SetFloat64(math.Pow10(int(decimals))))

		result, _ := resultBig.Int(nil)
		assetValueUSD = new(big.Int).Add(assetValueUSD, result)

	}
	return
}

func (scraper *HohmScraper) Liabilities() (liabilitiesValueUSD *big.Int, native bool, err error) {
	tokens, balances, err := scraper.getHohmTokensBalances()
	if err != nil {
		return
	}

	liabilitiesValueUSD = new(big.Int)
	for i, token := range tokens.Liabilities {
		// log.Infof("hOHM -- liability address: %s", token.Hex())
		supply := balances.TotalLiabilities[i]
		// TO DO: Get price from lumina metacontract.
		var price float64
		price, err = utils.GetDiaQuotationPrice("Ethereum", token.Hex())
		if err != nil {
			log.Errorf("hOHM -- failed to fetch price for %s: %v", token.Hex(), err)
			return
		}
		// log.Infof("hOHM -- assets supply -- price: %s -- %v", supply.String(), price)

		// Scale price and supply to big.Float.
		priceBig := big.NewFloat(0).Mul(big.NewFloat(price), new(big.Float).SetFloat64(math.Pow10(int(DECIMALS))))
		supplyFloat := big.NewFloat(0).SetInt(supply)

		// Decimals of token.
		decimals, err := utils.GetEthDecimals(token, scraper.client)
		if err != nil {
			log.Errorf("hOHM -- GetEthDecimals for %s: %v", token.Hex(), err)
			continue
		}
		// Scale price*supply to @DECIMALS decimals.
		resultBig := big.NewFloat(0).Quo(big.NewFloat(0).Mul(priceBig, supplyFloat), new(big.Float).SetFloat64(math.Pow10(int(decimals))))

		result, _ := resultBig.Int(nil)
		liabilitiesValueUSD = new(big.Int).Add(liabilitiesValueUSD, result)

	}
	return
}

func (scraper *HohmScraper) TotalSupply() (scaledSupply *big.Int, err error) {

	totalSupply, err := scraper.hohmCaller.TotalSupply(&bind.CallOpts{})
	if err != nil {
		return
	}

	// Scale total supply to @DECIMALS decimals.
	decimals, err := utils.GetEthDecimals(common.HexToAddress(scraper.contractAddress), scraper.client)
	if err != nil {
		return
	}

	scaledSupply, _ = big.NewFloat(0).Mul(new(big.Float).SetInt(totalSupply), new(big.Float).SetFloat64(math.Pow10(int(DECIMALS-decimals)))).Int(nil)
	return
}

func (scraper *HohmScraper) DataChannel() chan models.FairValueData {
	return scraper.dataChannel
}

func (scraper *HohmScraper) Close() chan bool {
	return scraper.BaseScraper.Close()
}

func (scraper *HohmScraper) GetConfig() models.FeedConfig {
	return scraper.config
}

// getOhmTokensBalances returns assets and liabilities for hohm contract.
func (scraper *HohmScraper) getHohmTokensBalances() (tokens struct {
	Assets      []common.Address
	Liabilities []common.Address
}, balances struct {
	TotalAssets      []*big.Int
	TotalLiabilities []*big.Int
}, err error) {

	tokens, err = scraper.hohmCaller.Tokens(&bind.CallOpts{})
	if err != nil {
		return
	}

	balances, err = scraper.hohmCaller.BalanceSheet(&bind.CallOpts{})
	if err != nil {
		return
	}

	return

}
