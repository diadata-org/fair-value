package main

import (
	"sync"

	"github.com/diadata-org/fair-value/models"
	"github.com/diadata-org/fair-value/scrapers"
	log "github.com/sirupsen/logrus"
)

// TO DO: Add config file (json) that determines ALL parameters needed for a scraper, i.e.
// feedType, blockchain, contractAddress, restClient, ...?
// main will go through the list of such entries and deploy the corresponding scraper.

func main() {
	// TEST

	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()

	// scraper := scrapers.NewIScraper("CONTRACT_EXCHANGE_RATE", "", "")

	// hOHM
	// scraper := scrapers.NewIScraper("NET_ASSET_VALUE", "Ethereum", "0x1db1591540d7a6062be0837ca3c808add28844f6")

	// bunnihub
	// poolID for UniV4 pool manager: 0x5e5bc151fdb581faf0c28ae30ceba9193da793ccd7c22a70d3feaf3408c07666
	scraper := scrapers.NewIScraper("CONTRACT_EXCHANGE_RATE", models.UNICHAIN, "0x78fd58693ff7796fDF565bD744fdC21CB9B49C6c", "0x5e5bc151fdb581faf0c28ae30ceba9193da793ccd7c22a70d3feaf3408c07666")

	go handleData(scraper.DataChannel(), &wg)

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
