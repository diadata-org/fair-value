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
func NewIScraper(feedType string, blockchain string, address string, updateSeconds int, params []any) IScraper {
	switch feedType {

	case "CONTRACT_EXCHANGE_RATE":

		scraper := NewIContractExchangeRate(blockchain, address, params)

		// TO DO: add select for scraper close
		ticker := time.NewTicker(time.Duration(updateSeconds) * time.Second)
		go func() {
			for {
				select {
				case <-ticker.C:
					scraper.DataChannel() <- MakeCERData(scraper, address, blockchain)
				case <-scraper.Close():
					log.Error("Close scraper!")
					return
				}
			}
		}()

		return scraper

	case "NET_ASSET_VALUE":

		scraper := NewINetAssetValue(blockchain, address, params)

		// Processing of nav.Methods to final fair value.
		ticker := time.NewTicker(time.Duration(updateSeconds) * time.Second)
		go func() {
			for {
				select {
				case <-ticker.C:
					scraper.DataChannel() <- MakeNAVData(scraper, address, blockchain)
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

func NewIContractExchangeRate(blockchain string, address string, params []any) IContractExchangeRate {

	asset := models.Asset{Blockchain: blockchain, Address: address}

	switch asset {
	// mbTON
	case models.Asset{Blockchain: "TON", Address: "EQCSxGZPHqa3TtnODgMan8CEM0jf6HpY-uon_NMeFgjKqkEY"}:

		log.Info("start mbTON scraper.")
		cer := NewBMTonScraper(blockchain, address)

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
		fairValueData := FairValueData{Address: "EQCSxGZPHqa3TtnODgMan8CEM0jf6HpY-uon_NMeFgjKqkEY", Blockchain: "TON", FairValueNative: price, Time: time.Now()}
		cer.DataChannel() <- fairValueData

		return cer
	// Bunnihub
	case models.Asset{Blockchain: models.UNICHAIN, Address: "0x78fd58693ff7796fDF565bD744fdC21CB9B49C6c"}:

		log.Info("start bunnihub scraper. ")
		cer := NewBunnihubScraper(blockchain, address, params)

		return cer
	}

	return nil
}

func NewINetAssetValue(blockchain string, address string, params []any) INetAssetValue {
	asset := models.Asset{Blockchain: blockchain, Address: address}

	switch asset {
	case models.Asset{Blockchain: models.ETHEREUM, Address: "0x1db1591540d7a6062be0837ca3c808add28844f6"}:
		log.Info("start hOHM scraper.")

		// TO DO: How do we deal with same asset/contract but different parameters?
		// Example: alphagrowth where asset is given by pool id but uni pool manager address is the same:
		// https://github.com/diadata-org/diadata-monitoring/blob/inspector-v2/experiments/funadmental_price_alphagrowth.py
		nav := NewHohmScraper(blockchain, address, params)

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
