package scrapers

import (
	"time"

	"github.com/diadata-org/fair-value/models"
)

// ---------------------------------------------------------------------------------
// SCRAPER FACTORY
// ---------------------------------------------------------------------------------

// TO DO: Can we set params to string?

// NewIScraper is the factory function for the basic Scraper interface.
// @params is a set of optional parameters such as poolID for bunnihub UniV4 pools.
func NewIScraper(config models.FeedConfig) IScraper {
	switch config.FeedType {

	case "CONTRACT_EXCHANGE_RATE":

		scraper := NewIContractExchangeRate(config)

		// Processing of CER data for final value.
		ticker := time.NewTicker(time.Duration(config.UpdateSeconds) * time.Second)
		go func() {
			for {
				select {
				case <-ticker.C:
					scraper.DataChannel() <- MakeCERData(scraper)
				case <-scraper.Close():
					log.Error("Close scraper!")
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
					scraper.DataChannel() <- MakeNAVData(scraper)
				case <-scraper.Close():
					log.Error("Close scraper!")
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

	switch asset {
	// mbTON
	case models.Asset{Blockchain: "TON", Address: "EQCSxGZPHqa3TtnODgMan8CEM0jf6HpY-uon_NMeFgjKqkEY"}:

		log.Info("start mbTON scraper.")
		cer := NewBMTonScraper(config)

		totalUnderlying, _, err := cer.TotalUnderlying()
		if err != nil {
			log.Fatal("totalUnderlying: ", err)
		}
		totalShares, err := cer.TotalShares()
		if err != nil {
			log.Fatal("totalShares: ", err)
		}
		log.Infof("totalUnderlying -- totalShares: %s -- %s", totalUnderlying.String(), totalShares.String())
		price, err := models.ComputeContractExchangeRatePrice(totalUnderlying, totalShares)
		if err != nil {
			log.Error("ComputeContractExchangeRatePrice: ", err)
		}
		fairValueData := models.FairValueData{Address: "EQCSxGZPHqa3TtnODgMan8CEM0jf6HpY-uon_NMeFgjKqkEY", Blockchain: "TON", FairValueNative: price, Time: time.Now()}
		cer.DataChannel() <- fairValueData

		return cer

	// satUSD+
	case models.Asset{Blockchain: models.BINANCESMARTCHAIN, Address: "0x03d9C4E4BC5D3678A9076caC50dB0251D8676872"}:
		log.Info("start satUSD+ scraper.")
		cer := NewSatusdScraper(config)
		return cer

	// Bunnihub
	case models.Asset{Blockchain: models.UNICHAIN, Address: "0x78fd58693ff7796fDF565bD744fdC21CB9B49C6c"}:

		log.Info("start bunnihub scraper.")
		cer := NewBunnihubScraper(config)

		return cer
	}

	return nil
}

func NewINetAssetValue(config models.FeedConfig) INetAssetValue {
	asset := models.Asset{Blockchain: config.Blockchain, Address: config.Address}

	switch asset {
	case models.Asset{Blockchain: models.ETHEREUM, Address: "0x1db1591540d7a6062be0837ca3c808add28844f6"}:
		log.Info("start hOHM scraper.")

		// TO DO: How do we deal with same asset/contract but different parameters?
		// Example: alphagrowth where asset is given by pool id but uni pool manager address is the same:
		// https://github.com/diadata-org/diadata-monitoring/blob/inspector-v2/experiments/funadmental_price_alphagrowth.py
		// Answer: use params.
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
