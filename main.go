package main

import (
	"context"
	"os"
	"strconv"
	"sync"
	"time"

	ValueStore "github.com/diadata-org/fair-value/contracts/valuestore"
	"github.com/diadata-org/fair-value/metrics"
	"github.com/diadata-org/fair-value/models"
	"github.com/diadata-org/fair-value/onchain"
	"github.com/diadata-org/fair-value/scrapers"
	"github.com/diadata-org/fair-value/utils"
	log "github.com/sirupsen/logrus"
)

var (
	configUpdateSeconds         int
	writeTickerSeconds          int
	feedConfigs, feedConfigsNew []models.FeedConfig
)

func init() {
	var err error
	configUpdateSeconds, err = strconv.Atoi(utils.Getenv("CONFIG_UPDATE_SECONDS", "86400"))
	if err != nil {
		log.Errorf("parse CONFIG_UPDATE_SECONDS: %v", err)
		configUpdateSeconds = 86400
	}
	writeTickerSeconds, err = strconv.Atoi(utils.Getenv("WRITE_TICKER_SECONDS", "120"))
	if err != nil {
		log.Errorf("parse WRITE_TICKER_SECONDS: %v", err)
		writeTickerSeconds = 120
	}
}

func main() {

	// On-chain setup
	deployedContract := utils.Getenv("DEPLOYED_CONTRACT", "")
	chainId, err := strconv.ParseInt(utils.Getenv("CHAIN_ID", "100640"), 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse chainId: %v", err)
	}
	auth, conn, connBackup, privateKey, err := utils.SetupOnchain(
		utils.Getenv("BLOCKCHAIN_NODE", ""),
		utils.Getenv("BACKUP_NODE", ""),
		utils.Getenv("PRIVATE_KEY", ""),
		chainId,
	)
	if err != nil {
		log.Fatal("SetupOnchain: ", err)
	}
	var contract *ValueStore.ValueStore
	var contractBackup *ValueStore.ValueStore
	err = onchain.DeployOrBindContract(deployedContract, conn, connBackup, auth, &contract, &contractBackup)
	if err != nil {
		log.Fatalf("Failed to Deploy or Bind primary and backup contract: %v", err)
	}

	// Start collecting and pushing metrics.
	metrics.StartMetrics(
		conn,
		privateKey,
		deployedContract,
		chainId,
		os.Getenv("PUSHGATEWAY_URL"),
		os.Getenv("PUSHGATEWAY_USER"),
		os.Getenv("PUSHGATEWAY_PASSWORD"),
		utils.Getenv("ENABLE_METRICS_SERVER", "false"),
		utils.Getenv("NODE_OPERATOR_NAME", ""),
		utils.Getenv("METRICS_PORT", "9090"),
	)

	// Setting up feeders.
	feedConfigs, err = models.GetFeedsFromConfig("fair-value-feeds.json")
	if err != nil {
		log.Fatal("GetFeedsFromConfig: ", err)
	}
	for _, fc := range feedConfigs {
		log.Infof("loaded %s from config. %s -- %s", fc.Symbol, fc.Blockchain, fc.Address)
	}

	var wg sync.WaitGroup
	defer wg.Wait()
	collectorChannel := make(chan models.FairValueData)

	// Create scrapers for all assets available in config file.
	allIscrapers := make(map[string]scrapers.IScraper)
	for _, config := range feedConfigs {
		ctx, cancel := context.WithCancel(context.Background())
		allIscrapers[config.FeedConfigIdentifier()] = scrapers.NewIScraper(cancel, config)
		wg.Add(1)
		go handleData(ctx, allIscrapers[config.FeedConfigIdentifier()], &wg, collectorChannel)
	}

	// Routine that periodically fetches the configs and compares to deployed config.
	// Whenever something is added to or removed from the config, either deploy or close it.
	// For now, assume that configs are either added or removed, i.e. existing ones cannot be changed.
	go func() {
		configTicker := time.NewTicker(time.Duration(time.Duration(configUpdateSeconds) * time.Second))

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
				go handleData(ctx, allIscrapers[config.FeedConfigIdentifier()], &wg, collectorChannel)
			}
		}
	}()

	// Routine writing collected data to the oracle.
	collectedData := make(map[string]models.FairValueData)
	go func() {
		writeTicker := time.NewTicker(time.Duration(writeTickerSeconds) * time.Second)
		for range writeTicker.C {
			log.Info("collectedData:----------------------------------- ", collectedData)
			onchain.OracleUpdateExecutor(auth, contract, contractBackup, conn, connBackup, collectedData)
		}
	}()

	go func() {
		for d := range collectorChannel {
			// Only store the latest data point.
			collectedData[d.FairValueDataIdentifier()] = d
		}
	}()

}

// handleData handles data from dataChannel.
func handleData(ctx context.Context, scraper scrapers.IScraper, wg *sync.WaitGroup, collectorChannel chan models.FairValueData) {
	defer wg.Done()

	for {
		select {
		case d := <-scraper.DataChannel():
			log.Info("channel out: ", d)
			log.Infof("symbol -- fairValueNative: %s -- %v", d.Symbol, d.FairValueNative)
			collectorChannel <- d
		case <-ctx.Done():
			log.Warn("close data handler for scraper ", scraper.GetConfig().Symbol)
			return
		}
	}
}
