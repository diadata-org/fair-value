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
func NewIScraper(cancel context.CancelFunc, config models.FeedConfig) IScraper {

	switch config.FeedType {

	case "CONTRACT_EXCHANGE_RATE":

		scraper := NewIContractExchangeRate(config)

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

		scraper := NewINetAssetValue(config)

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

func NewIContractExchangeRate(config models.FeedConfig) IContractExchangeRate {

	asset := models.Asset{Blockchain: config.Blockchain, Address: config.Address}
	log.Infof("start %s scraper.", config.Symbol)

	switch asset {

	// pBTC
	case models.Asset{Blockchain: models.ARBITRUM_SEPOLIA, Address: "0xF0e7E2c829e54a6052d80E2c2107a494b00BC359"}:
		cer := NewpBTCScraper(config)
		return cer

	// USDp
	case models.Asset{Blockchain: models.ETHEREUM, Address: "0x6efeDDF9269c3683Ba516cb0e2124FE335F262a2"}:
		cer := NewUSDPScraperScraper(config)
		return cer

	// bmTON
	case models.Asset{Blockchain: models.TONCHAIN, Address: "EQCSxGZPHqa3TtnODgMan8CEM0jf6HpY-uon_NMeFgjKqkEY"}:
		cer := NewBMTonScraper(config)
		return cer

	// satUSD+
	case models.Asset{Blockchain: models.BINANCESMARTCHAIN, Address: "0x03d9C4E4BC5D3678A9076caC50dB0251D8676872"}:
		cer := NewSatusdScraper(config)
		return cer

	// Bunnihub
	case models.Asset{Blockchain: models.UNICHAIN, Address: "0x78fd58693ff7796fDF565bD744fdC21CB9B49C6c"}:
		cer := NewBunnihubScraper(config)
		return cer
	}

	return nil
}

func NewINetAssetValue(config models.FeedConfig) INetAssetValue {
	asset := models.Asset{Blockchain: config.Blockchain, Address: config.Address}
	log.Infof("start %s scraper.", config.Symbol)

	switch asset {
	case models.Asset{Blockchain: models.ETHEREUM, Address: "0x1db1591540d7a6062be0837ca3c808add28844f6"}:
		nav := NewHohmScraper(config)
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
