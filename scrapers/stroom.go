package scrapers

import (
	"math/big"

	"github.com/diadata-org/fair-value/models"
	"github.com/diadata-org/fair-value/utils"
	"github.com/tidwall/gjson"
)

type stroomScraper struct {
	BaseScraper
	lightningRPC string
	macaroon     string
	config       models.FeedConfig
	chunkSize    uint64
}

func NewStroomScraper(config models.FeedConfig, metacontractData models.MetacontractData) *stroomScraper {

	scraper := stroomScraper{
		BaseScraper:  NewBaseScraper(metacontractData),
		lightningRPC: utils.Getenv("LIGHTNING_RPC_NODE_STROOM", ""),
		macaroon:     utils.Getenv("MACAROON_STROOM", ""),
		config:       config,
		chunkSize:    uint64(10000),
	}

	return &scraper

}

func (scraper *stroomScraper) TotalUnderlying() (totalUnderlying *big.Int, totalValueUnderlying *big.Int, err error) {

	// TO DO

	return
}

func (scraper *stroomScraper) TotalShares() (totalShares *big.Int, err error) {
	var channelBalance, walletBalance, pendingChannels float64

	channelBalance, err = scraper.getChannelBalance(scraper.macaroon)
	if err != nil {
		return
	}
	walletBalance, err = scraper.getWalletBalance(scraper.macaroon)
	if err != nil {
		return
	}
	pendingChannels, err = scraper.getPendingChannels(scraper.macaroon)
	if err != nil {
		return
	}
	log.Tracef("Stroom - channelBalance -- walletBalance -- pendingChannels: %v -- %v -- %v", channelBalance, walletBalance, pendingChannels)
	totalShares, _ = big.NewFloat(channelBalance + walletBalance + pendingChannels).Int(nil)

	return
}

func (scraper *stroomScraper) DataChannel() chan models.FairValueData {
	return scraper.dataChannel
}

func (scraper *stroomScraper) GetConfig() models.FeedConfig {
	return scraper.config
}

func (scraper *stroomScraper) Close() chan bool {
	return scraper.BaseScraper.Close()
}

func (scraper *stroomScraper) getChannelBalance(macaroon string) (channelBalance float64, err error) {

	url := scraper.lightningRPC + "/v1/balance/channels"
	header := make(map[string]string)
	header["Grpc-Metadata-macaroon"] = macaroon
	contents, _, err := utils.GetRequest(url, header)
	if err != nil {
		return
	}

	localBalance := gjson.Get(string(contents), "local_balance.sat").Float()
	unsettledLocalBalance := gjson.Get(string(contents), "unsettled_local_balance.sat").Float()
	pendingOpenLocalBalance := gjson.Get(string(contents), "pending_open_local_balance.sat").Float()
	channelBalance = localBalance + unsettledLocalBalance + pendingOpenLocalBalance

	return
}

func (scraper *stroomScraper) getWalletBalance(macaroon string) (walletBalance float64, err error) {

	url := scraper.lightningRPC + "/v1/balance/blockchain"
	header := make(map[string]string)
	header["Grpc-Metadata-macaroon"] = macaroon
	contents, _, err := utils.GetRequest(url, header)
	walletBalance = gjson.Get(string(contents), "total_balance").Float()

	return
}

func (scraper *stroomScraper) getPendingChannels(macaroon string) (pending float64, err error) {

	url := scraper.lightningRPC + "/v1/channels/pending"
	header := make(map[string]string)
	header["Grpc-Metadata-macaroon"] = macaroon
	contents, _, err := utils.GetRequest(url, header)
	pending = gjson.Get(string(contents), "total_limbo_balance").Float()

	return
}
