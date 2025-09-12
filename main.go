package main

import (
	"sync"
	"time"

	"github.com/diadata-org/fair-value/models"
	"github.com/diadata-org/fair-value/scrapers"
	log "github.com/sirupsen/logrus"
)

const (
	CONFIG_UPDATE_SECONDS = 60
)

var feedConfigs, feedConfigsNew []models.FeedConfig

func main() {

	var err error
	feedConfigs, err = models.GetFeedsFromConfig("fair-value-feeds.json")
	if err != nil {
		log.Fatal("GetFeedsFromConfig: ", err)
	}

	var wg sync.WaitGroup
	defer wg.Wait()

	allIscrapers := make(map[string]scrapers.IScraper)
	for _, config := range feedConfigs {
		allIscrapers[config.FeedConfigIdentifier()] = scrapers.NewIScraper(config)
		wg.Add(1)
		go handleData(allIscrapers[config.FeedConfigIdentifier()].DataChannel(), &wg)
	}

	// Routine that periodically fetches the configs and compares to deployed config.
	// Whenever something is added to or removed from the config, either deploy or close it.
	// For now, assume that configs are either added or removed, i.e. existing ones cannot be changed.
	go func() {
		configTicker := time.NewTicker(time.Duration(time.Duration(CONFIG_UPDATE_SECONDS) * time.Second))

		for range configTicker.C {

			feedConfigsNew, err = models.GetFeedsFromConfig("fair-value-feeds.json")
			if err != nil {
				log.Error("GetFeedsFromConfig: ", err)
			}

			plus, minus := models.GetDiffConfig(feedConfigs, feedConfigsNew)
			log.Warnf("plus -- minus: %v -- %v", plus, minus)

			// Close scrapers for removed configs.
			for _, config := range minus {
				if _, ok := allIscrapers[config.FeedConfigIdentifier()]; ok {
					allIscrapers[config.FeedConfigIdentifier()].Close() <- true
					delete(allIscrapers, config.FeedConfigIdentifier())
				}
			}

			// Add scraper for added configs.
			for _, config := range plus {
				allIscrapers[config.FeedConfigIdentifier()] = scrapers.NewIScraper(config)
				wg.Add(1)
				go handleData(allIscrapers[config.FeedConfigIdentifier()].DataChannel(), &wg)
			}
		}
	}()

}

// TO DO: Add onchain functionality.
// Handles data from dataChannel.
func handleData(dataChannel chan models.FairValueData, wg *sync.WaitGroup) {
	defer wg.Done()
	// TO DO: make select statement with graceful close.
	for d := range dataChannel {
		log.Info("channel out: ", d)
		// This should be the final line of main (blocking call)
		// onchain.OracleUpdateExecutor(auth, contract, contractBackup, conn, connBackup, chainID, filtersChannel)
	}
}
