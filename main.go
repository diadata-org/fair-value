package main

import (
	"sync"

	"github.com/diadata-org/fair-value/models"
	"github.com/diadata-org/fair-value/scrapers"
	log "github.com/sirupsen/logrus"
)

var feedConfigs, feedConfigsNew []models.FeedConfig

// TO DO: Add config file (json) that determines ALL parameters needed for a scraper, i.e.
// feedType, blockchain, contractAddress, restClient, ...?
// main will go through the list of such entries and deploy the corresponding scraper.

func main() {

	// TO DO: Add routine that periodically fetches the configs and compares to deployed config.
	// Whenever something is added to the config, also deploy it.
	// For now, assume that configs are either added or removed, i.e. existing ones are not changed.

	var err error
	feedConfigs, err = models.GetFeedsFromConfig("fair-value-feeds.json")
	if err != nil {
		log.Fatal("GetFeedsFromConfig: ", err)
	}

	var wg sync.WaitGroup
	defer wg.Wait()

	allIscrapers := make(map[*models.FeedConfig]scrapers.IScraper)
	for _, f := range feedConfigs {
		allIscrapers[&f] = scrapers.NewIScraper(f.FeedType, f.Blockchain, f.Address, f.UpdateSeconds, f.Params)
		wg.Add(1)
		go handleData(allIscrapers[&f].DataChannel(), &wg)
	}

	// go func() {
	// 	configTicker := time.NewTicker(time.Duration(1 * time.Minute))
	// 	for range configTicker.C {
	// 		feedConfigsNew, err = models.GetFeedsFromConfig("fair-value-feeds.json")
	// 		if err != nil {
	// 			log.Error("GetFeedsFromConfig: ", err)
	// 		}
	// 	}
	// 	plus, minus := models.GetDiffConfig(feedConfigs, feedConfigsNew)
	// 	log.Infof("plus -- minus: %v -- %v", plus, minus)
	// 	for _, config := range feedConfigsNew {
	// 		allIscrapers[&config].Close()
	// 	}
	// }()

	// hOHM
	// scraper := scrapers.NewIScraper("NET_ASSET_VALUE", "Ethereum", "0x1db1591540d7a6062be0837ca3c808add28844f6")

	// bunnihub
	// poolID for UniV4 pool manager: 0x5e5bc151fdb581faf0c28ae30ceba9193da793ccd7c22a70d3feaf3408c07666
	// scraper := scrapers.NewIScraper("CONTRACT_EXCHANGE_RATE", models.UNICHAIN, "0x78fd58693ff7796fDF565bD744fdC21CB9B49C6c", "0x5e5bc151fdb581faf0c28ae30ceba9193da793ccd7c22a70d3feaf3408c07666")

}

// TO DO: Forward price data to channel and write to contract here using handleData() function.
func handleData(dataChannel chan scrapers.FairValueData, wg *sync.WaitGroup) {
	defer wg.Done()
	// TO DO: make select statement with graceful close.
	log.Info("start handling data...")
	for d := range dataChannel {
		log.Info("channel out: ", d)
		// This should be the final line of main (blocking call)
		// onchain.OracleUpdateExecutor(auth, contract, contractBackup, conn, connBackup, chainID, filtersChannel)
	}
}
