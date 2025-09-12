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

// ----------------------------- ON HOLD ----------------------------

// ------------------------------------------------------------------
// CONTRACT EXCHANGE RATE
// ------------------------------------------------------------------

const BM_TON_ADDRESS = "EQCSxGZPHqa3TtnODgMan8CEM0jf6HpY-uon_NMeFgjKqkEY"

type BMTonScraper struct {
	BaseScraper
	blockchain      string
	contractAddress string
}

func NewBMTonScraper(blockchain string, address string) *BMTonScraper {
	return &BMTonScraper{
		BaseScraper:     NewBaseScraper(),
		blockchain:      blockchain,
		contractAddress: address,
	}

}

func (scraper *BMTonScraper) TotalUnderlying() (totalTon *big.Int, totalTonValue *big.Int, err error) {
	result, err := getBmtonExecutionResult()
	if err != nil {
		return
	}
	totalTon, err = result.Int(1)
	return
}

func (scraper *BMTonScraper) TotalShares() (totalBmton *big.Int, err error) {
	result, err := getBmtonExecutionResult()
	if err != nil {
		return
	}
	totalBmton, err = result.Int(0)
	return
}

func (scraper *BMTonScraper) DataChannel() chan models.FairValueData {
	return scraper.dataChannel
}

// TO DO
func (scraper *BMTonScraper) Close() chan bool {
	return scraper.BaseScraper.Close()
}

func getBmtonExecutionResult() (result *ton.ExecutionResult, err error) {

	// Connect to TON mainnet
	client := liteclient.NewConnectionPool()
	err = client.AddConnectionsFromConfigUrl(context.Background(), "https://ton.org/global-config.json")
	if err != nil {
		return
	}

	api := ton.NewAPIClient(client)

	// Resolve contract address
	contractAddr := address.MustParseAddr(BM_TON_ADDRESS)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Get latest masterchain block
	var block *ton.BlockIDExt
	block, err = api.CurrentMasterchainInfo(ctx)
	if err != nil {
		return
	}

	return api.RunGetMethod(ctx, block, contractAddr, "get_full_data")

}
