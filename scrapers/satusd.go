package scrapers

import (
	"math/big"

	"github.com/diadata-org/fair-value/models"
	"github.com/diadata-org/fair-value/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	satusd "github.com/diadata-org/fair-value/contracts/satusd"
)

// ------------------------------------------------------------------
// CONTRACT EXCHANGE RATE
// ------------------------------------------------------------------

type SatusdScraper struct {
	BaseScraper
	client          *ethclient.Client
	blockchain      string
	contractAddress common.Address
	poolID          common.Hash
	lpTokenAddress  common.Address
	config          models.FeedConfig
}

func NewSatusdScraper(config models.FeedConfig) *SatusdScraper {

	scraper := SatusdScraper{
		BaseScraper:     NewBaseScraper(),
		blockchain:      config.Blockchain,
		contractAddress: common.HexToAddress(config.Address),
		config:          config,
	}

	if len(config.Params) > 0 {
		scraper.poolID = common.HexToHash(config.Params[0].(string))
	}

	scraper.lpTokenAddress = common.HexToAddress(config.Address)

	client, err := ethclient.Dial(utils.Getenv("RPC_NODE_SATUSD", ""))
	if err != nil {
		log.Fatal(err)
	}
	scraper.client = client

	return &scraper

}

func (scraper *SatusdScraper) TotalUnderlying() (totalUnderlying *big.Int, totalValueUnderlying *big.Int, err error) {
	satusdCaller, err := satusd.NewSatusdCaller(scraper.contractAddress, scraper.client)
	if err != nil {
		return
	}
	totalUnderlying, err = satusdCaller.TotalAssets(&bind.CallOpts{})
	if err != nil {
		return
	}

	// TO DO: Get underlying asset from contract.

	// DIA Prices
	satusdPrice, err := utils.GetDiaQuotationPrice(models.BINANCESMARTCHAIN, "0xb4818BB69478730EF4e33Cc068dD94278e2766cB")
	if err != nil {
		log.Fatal(err)
	}
	// Scaled sum of values.
	totalValueUnderlying, _ = new(big.Float).Mul(big.NewFloat(0).SetInt(totalUnderlying), big.NewFloat(satusdPrice)).Int(nil)
	log.Debug("satUSD total value underlying: ", totalValueUnderlying)

	return
}

func (scraper *SatusdScraper) TotalShares() (*big.Int, error) {
	satusdCaller, err := satusd.NewSatusdCaller(scraper.contractAddress, scraper.client)
	if err != nil {
		return nil, err
	}
	return satusdCaller.TotalSupply(&bind.CallOpts{})
}

func (scraper *SatusdScraper) DataChannel() chan models.FairValueData {
	return scraper.dataChannel
}

func (scraper *SatusdScraper) GetConfig() models.FeedConfig {
	return scraper.config
}

func (scraper *SatusdScraper) Close() chan bool {
	return scraper.BaseScraper.Close()
}
