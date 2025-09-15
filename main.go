package main

import (
	"context"
	"sync"
	"time"

	"github.com/diadata-org/fair-value/models"
	"github.com/diadata-org/fair-value/scrapers"
	log "github.com/sirupsen/logrus"
)

const (
	CONFIG_UPDATE_SECONDS = 60
)

var (
	feedConfigs, feedConfigsNew []models.FeedConfig
)

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
		ctx, cancel := context.WithCancel(context.Background())
		allIscrapers[config.FeedConfigIdentifier()] = scrapers.NewIScraper(cancel, config)
		wg.Add(1)
		go handleData(ctx, allIscrapers[config.FeedConfigIdentifier()], &wg)
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
			log.Info("REMOVING the following feeds...")
			for _, m := range minus {
				log.Infof("Symbol -- Address: %s -- %s", m.Symbol, m.Address)
			}
			log.Info("ADDING the following feeds...")
			for _, m := range plus {
				log.Infof("Symbol -- Address: %s -- %s", m.Symbol, m.Address)
			}

			// TEST removing a feed
			minus = []models.FeedConfig{
				{
					Symbol:     "satUSD+",
					FeedType:   "CONTRACT_EXCHANGE_RATE",
					Address:    "0x03d9C4E4BC5D3678A9076caC50dB0251D8676872",
					Blockchain: "BinanceSmartChain",
				},
			}

			// Close scrapers for removed configs.
			for _, config := range minus {
				if _, ok := allIscrapers[config.FeedConfigIdentifier()]; ok {
					allIscrapers[config.FeedConfigIdentifier()].Close() <- true
					delete(allIscrapers, config.FeedConfigIdentifier())
				}
			}

			// Add scraper for added configs.
			for _, config := range plus {
				ctx, cancel := context.WithCancel(context.Background())
				allIscrapers[config.FeedConfigIdentifier()] = scrapers.NewIScraper(cancel, config)
				wg.Add(1)
				go handleData(ctx, allIscrapers[config.FeedConfigIdentifier()], &wg)
			}
		}
	}()

}

// TO DO: Add onchain functionality.
// handleData handles data from dataChannel.
func handleData(ctx context.Context, scraper scrapers.IScraper, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case d := <-scraper.DataChannel():
			log.Info("channel out: ", d)
			// This should be the final line of main (blocking call)
			// onchain.OracleUpdateExecutor(auth, contract, contractBackup, conn, connBackup, chainID, filtersChannel)
		case <-ctx.Done():
			log.Warn("close data handler for scraper ", scraper.GetConfig().Symbol)
			return
		}
	}
}
