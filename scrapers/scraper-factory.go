package scrapers

import (
	"context"
	"time"

	"github.com/diadata-org/fair-value/models"
)

// ---------------------------------------------------------------------------------
// SCRAPER FACTORY
// ---------------------------------------------------------------------------------

// TO DO: Can we set params to string?

// NewIScraper is the factory function for the basic Scraper interface.
// @params is a set of optional parameters such as poolID for bunnihub UniV4 pools.
func NewIScraper(cancel context.CancelFunc, config models.FeedConfig, metacontractData models.MetacontractData) IScraper {

	switch config.FeedType {

	case "CONTRACT_EXCHANGE_RATE":

		scraper := NewIContractExchangeRate(config, metacontractData)

		// Processing of CER data for final value.
		ticker := time.NewTicker(time.Duration(config.UpdateSeconds) * time.Second)
		go func() {
			for {
				select {
				case <-ticker.C:
					if data, err := MakeCERData(scraper); err == nil {
						scraper.DataChannel() <- data
					} else {
						log.Errorf("MakeCERData for %s: %v", config.Symbol, err)
					}
				case <-scraper.Close():
					log.Warnf("Close %s scraper!", config.Symbol)
					cancel()
					return
				}
			}
		}()

		return scraper

	case "NET_ASSET_VALUE":

		scraper := NewINetAssetValue(config, metacontractData)

		// Processing of nav.Methods to final fair value.
		ticker := time.NewTicker(time.Duration(config.UpdateSeconds) * time.Second)
		go func() {
			for {
				select {
				case <-ticker.C:
					if data, err := MakeNAVData(scraper); err == nil {
						scraper.DataChannel() <- data
					} else {
						log.Errorf("MakeNAVData for %s: %v", config.Symbol, err)
					}
				case <-scraper.Close():
					log.Warnf("Close %s scraper!", config.Symbol)
					cancel()
					return
				}
			}
		}()

		return scraper

	}
	return nil
}

// ---------------------------------------------------------------------------------
// SPECIFIC FAIR VALUE SCRAPER FACTORIES
// ---------------------------------------------------------------------------------

func NewIContractExchangeRate(config models.FeedConfig, metacontractData models.MetacontractData) IContractExchangeRate {

	symbol := config.Symbol
	log.Infof("start %s scraper.", symbol)

	switch symbol {

	case "pBTC":
		cer := NewpBTCScraper(config, metacontractData)
		return cer

	case "hemiBTC":
		cer := NewhemiBTCScraper(config, metacontractData)
		return cer

	case "USDp":
		cer := NewUSDPScraperScraper(config, metacontractData)
		return cer

	case "bmTON":
		cer := NewBMTonScraper(config, metacontractData)
		return cer

	case "satUSD+":
		cer := NewSatusdScraper(config, metacontractData)
		return cer

	// case "hOHM":
	// cer := NewBunnihubScraper(config, metacontractData)
	// return cer

	case "stroom":
		cer := NewStroomScraper(config, metacontractData)
		return cer

	case "USDr":
		cer := NewUSDRScraperScraper(config, metacontractData)
		return cer
	}

	return nil
}

func NewINetAssetValue(config models.FeedConfig, metacontractData models.MetacontractData) INetAssetValue {
	symbol := config.Symbol
	log.Infof("start %s scraper.", symbol)

	switch symbol {
	case "hOHM":
		nav := NewHohmScraper(config, metacontractData)
		return nav
	}

	return nil
}

func NewIReserveBacking(blockchain string, address string) IReserveBacking {

	asset := models.Asset{Blockchain: blockchain, Address: address}

	switch asset {
	// TO DO
	}
	return nil
}
