package scrapers

import (
	"math/big"

	"github.com/diadata-org/fair-value/models"
	"github.com/diadata-org/fair-value/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	spsei "github.com/diadata-org/fair-value/contracts/spSEI"
	spseistaking "github.com/diadata-org/fair-value/contracts/spSEIStaking"
)

// ------------------------------------------------------------------
// CONTRACT EXCHANGE RATE
// ------------------------------------------------------------------

type SpSEIScraper struct {
	BaseScraper
	client          *ethclient.Client
	blockchain      string
	contractAddress common.Address
	lpTokenAddress  common.Address
	config          models.FeedConfig
	spSeiContract   common.Address
}

func NewSpSEIScraper(config models.FeedConfig, metacontractData models.MetacontractData) *SpSEIScraper {

	scraper := SpSEIScraper{
		BaseScraper:     NewBaseScraper(metacontractData),
		blockchain:      config.Blockchain,
		contractAddress: common.HexToAddress(config.Address),
		config:          config,
	}

	scraper.lpTokenAddress = common.HexToAddress(config.Address)
	scraper.spSeiContract = common.HexToAddress(scraper.config.Params[1].([]any)[0].(string))

	client, err := ethclient.Dial(utils.Getenv("RPC_NODE_SPSEI", "https://sei-public.nodies.app"))
	if err != nil {
		log.Errorf("SpSEI -- make eth client for %s: %v", config.Symbol, err)
		return nil
	}
	scraper.client = client

	return &scraper

}

func (scraper *SpSEIScraper) TotalUnderlying() (totalUnderlying *big.Int, totalValueUnderlying *big.Int, err error) {
	spseistakingCaller, err := spseistaking.NewSpseistakingCaller(scraper.contractAddress, scraper.client)
	if err != nil {
		return
	}
	totalUnderlying, err = spseistakingCaller.GetTotalSei(&bind.CallOpts{})
	if err != nil {
		return
	}

	// DIA Prices
	SEI := models.Asset{Symbol: "SEI", Blockchain: models.SEICHAIN, Address: "0x0000000000000000000000000000000000000000"}
	satusdQuotation, err := SEI.GetPrice(scraper.metacontractData.Address, scraper.metacontractData.Precision, scraper.metacontractData.Client)
	if err != nil {
		log.Error("SpSEI -- GetPrice: ", err)
	}
	// Scaled sum of values.
	totalValueUnderlying, _ = new(big.Float).Mul(big.NewFloat(0).SetInt(totalUnderlying), big.NewFloat(satusdQuotation.Price)).Int(nil)
	log.Debug("spsei total value underlying: ", totalValueUnderlying)

	return
}

func (scraper *SpSEIScraper) TotalShares() (*big.Int, error) {
	spseiCaller, err := spsei.NewSpseiCaller(scraper.spSeiContract, scraper.client)
	if err != nil {
		return nil, err
	}
	return spseiCaller.TotalSupply(&bind.CallOpts{})
}

func (scraper *SpSEIScraper) DataChannel() chan models.FairValueData {
	return scraper.dataChannel
}

func (scraper *SpSEIScraper) GetConfig() models.FeedConfig {
	return scraper.config
}

func (scraper *SpSEIScraper) Close() chan bool {
	return scraper.BaseScraper.Close()
}
