package scrapers

import (
	"context"
	"math/big"
	"time"

	"github.com/diadata-org/fair-value/models"
	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/ton"
)

// ------------------------------------------------------------------
// CONTRACT EXCHANGE RATE
// ------------------------------------------------------------------

const BM_TON_ADDRESS = "EQCSxGZPHqa3TtnODgMan8CEM0jf6HpY-uon_NMeFgjKqkEY"

type BMTonScraper struct {
	BaseScraper
	blockchain      string
	contractAddress string
	config          models.FeedConfig
	api             *ton.APIClient
}

func NewBMTonScraper(config models.FeedConfig) *BMTonScraper {
	scraper := BMTonScraper{
		BaseScraper:     NewBaseScraper(),
		blockchain:      config.Blockchain,
		contractAddress: config.Address,
		config:          config,
	}

	// Connect to TON mainnet
	client := liteclient.NewConnectionPool()
	err := client.AddConnectionsFromConfigUrl(context.Background(), "https://ton.org/global-config.json")
	if err != nil {
		return nil
	}
	api := ton.NewAPIClient(client)
	scraper.api = api

	return &scraper
}

func (scraper *BMTonScraper) TotalUnderlying() (totalTon *big.Int, totalTonValue *big.Int, err error) {
	result, err := scraper.getBmtonExecutionResult()
	if err != nil {
		return
	}
	totalTon, err = result.Int(1)
	return
}

func (scraper *BMTonScraper) TotalShares() (totalBmton *big.Int, err error) {
	result, err := scraper.getBmtonExecutionResult()
	if err != nil {
		return
	}
	totalBmton, err = result.Int(0)
	return
}

func (scraper *BMTonScraper) DataChannel() chan models.FairValueData {
	return scraper.dataChannel
}

func (scraper *BMTonScraper) Close() chan bool {
	return scraper.BaseScraper.Close()
}

func (scraper *BMTonScraper) GetConfig() models.FeedConfig {
	return scraper.config
}

func (scraper *BMTonScraper) getBmtonExecutionResult() (result *ton.ExecutionResult, err error) {

	// Resolve contract address
	contractAddr := address.MustParseAddr(scraper.config.Address)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Get latest masterchain block
	var block *ton.BlockIDExt
	block, err = scraper.api.CurrentMasterchainInfo(ctx)
	if err != nil {
		return
	}

	return scraper.api.RunGetMethod(ctx, block, contractAddr, "get_full_data")

}
