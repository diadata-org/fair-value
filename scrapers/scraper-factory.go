package scrapers

import (
	"context"
	"errors"
	"time"

	"github.com/diadata-org/fair-value/models"
)

// ---------------------------------------------------------------------------------
// SCRAPER FACTORY
// ---------------------------------------------------------------------------------

// TO DO: Can we set params to string?

// NewIScraper is the factory function for the basic Scraper interface.
// @params is a set of optional parameters such as poolID for bunnihub UniV4 pools.
func NewIScraper(
	cancel context.CancelFunc,
	config models.FeedConfig,
	metacontractData models.MetacontractData,
	metacontractFV models.MetacontractData,
) (IScraper, error) {

	switch config.FeedType {

	case "CONTRACT_EXCHANGE_RATE":

		scraper, err := NewIContractExchangeRate(config, metacontractData, metacontractFV)
		if err != nil {
			return nil, err
		}

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

		return scraper, nil

	case "NET_ASSET_VALUE":

		scraper, err := NewINetAssetValue(config, metacontractData)
		if err != nil {
			return nil, err
		}

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

		return scraper, nil

	}
	return nil, errors.New("no scraper created")
}

// ---------------------------------------------------------------------------------
// SPECIFIC FAIR VALUE SCRAPER FACTORIES
// ---------------------------------------------------------------------------------

func NewIContractExchangeRate(
	config models.FeedConfig,
	metacontractData models.MetacontractData,
	metacontractFV models.MetacontractData,
) (IContractExchangeRate, error) {

	symbol := config.Symbol
	log.Infof("start %s scraper.", symbol)

	switch symbol {

	case "sVUSD":
		return NewSVUSDScraper(config, metacontractData, metacontractFV)

	case "VUSD":
		return NewVUSDScraper(config, metacontractData)

	case "spSEI":
		return NewSpSEIScraper(config, metacontractData)

	case "pBTC":
		return NewpBTCScraper(config, metacontractData)

	case "hemiBTC":
		return NewhemiBTCScraper(config, metacontractData)

	case "USDp":
		return NewUSDPScraperScraper(config, metacontractData)

	case "bmTON":
		return NewBMTonScraper(config, metacontractData)

	case "satUSD+":
		return NewSatusdScraper(config, metacontractData)

	// case "hOHM":
	// cer := NewBunnihubScraper(config, metacontractData)
	// return cer

	case "stroom":
		return NewStroomScraper(config, metacontractData)

	case "USDr":
		return NewUSDRScraperScraper(config, metacontractData)
	}

	return nil, errors.New("symbol not available")
}

func NewINetAssetValue(config models.FeedConfig, metacontractData models.MetacontractData) (INetAssetValue, error) {
	symbol := config.Symbol
	log.Infof("start %s scraper.", symbol)

	switch symbol {
	case "hOHM":
		return NewHohmScraper(config, metacontractData)
	}

	return nil, errors.New("symbol not available")
}

func NewIReserveBacking(blockchain string, address string) IReserveBacking {

	asset := models.Asset{Blockchain: blockchain, Address: address}

	switch asset {
	// TO DO
	}
	return nil
}
