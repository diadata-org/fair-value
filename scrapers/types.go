package scrapers

import (
	"math/big"
	"time"
)

// TO DO: This should be moved to models.

// NOTE: All big return values have 18 decimals.
// If the boolean native is set to true, values are considered to be supply.
// Otherwise, values are USD denominations.

const (
	DECIMALS = 18
)

type IScraper interface {
	DataChannel() chan FairValueData
	Close() chan bool
}

type IContractExchangeRate interface {
	IScraper
	TotalUnderlying() (*big.Int, bool, error)
	TotalShares() (*big.Int, error)
}

type IReserveBacking interface {
	IScraper
	Reserves() (*big.Int, bool, error)
	Supply() (*big.Int, bool, error)
}

// TO DO: Maybe better to add another interface method Native()?
// Because I think either both - Assets AND Liabilities are native or not. Total supply should always(?) be native.
type INetAssetValue interface {
	IScraper
	Assets() (*big.Int, bool, error)
	Liabilities() (*big.Int, bool, error)
	TotalSupply() (*big.Int, error)
}

// MAKECERData computes all return values for the IContractExchangeRate interface.
func MakeCERData(scraper IContractExchangeRate, address string, blockchain string) (data FairValueData) {

	underlying, _, err := scraper.TotalUnderlying()
	if err != nil {
		log.Error("TotalUnderlying: ", err)
	}

	totalShares, err := scraper.TotalShares()
	if err != nil {
		log.Error("TotalShares: ", err)
	}

	log.Infof("underlying -- totalShares: %s -- %s", underlying.String(), totalShares.String())

	numeratorFloat := big.NewFloat(0).SetInt(underlying)
	denominatorFloat := big.NewFloat(0).SetInt(totalShares)
	price := big.NewFloat(0).Quo(numeratorFloat, denominatorFloat)

	data.Address = address
	data.Blockchain = blockchain
	data.Numerator = underlying
	data.Denominator = totalShares
	data.PriceUSD, _ = price.Float64()
	data.Time = time.Now()
	log.Info("price: ", data.PriceUSD)
	log.Info("data: ", data)

	return
}

// TO DO: Is this the best way to produce data?
func MakeNAVData(scraper INetAssetValue, address string, blockchain string) (data FairValueData) {

	assets, _, err := scraper.Assets()
	if err != nil {
		log.Error("scraper.Assets(): ", err)
	}
	liabilities, _, err := scraper.Liabilities()
	if err != nil {
		log.Error("scraper.Liabilities(): ", err)
	}
	supply, err := scraper.TotalSupply()
	if err != nil {
		log.Error("TotalSupply: ", err)
	}

	// TO DO: Compute (assets-liabilities)/supply
	numerator := new(big.Int).Sub(assets, liabilities)
	numeratorFloat := big.NewFloat(0).SetInt(numerator)
	denominator := big.NewFloat(0).SetInt(supply)
	price := big.NewFloat(0).Quo(numeratorFloat, denominator)

	data.Address = address
	data.Blockchain = blockchain
	data.Numerator = numerator
	data.Denominator = supply
	data.PriceUSD, _ = price.Float64()
	data.Time = time.Now()

	return
}

type FairValueData struct {
	Address    string
	Blockchain string
	Symbol     string
	// Fair value price in USD.
	PriceUSD float64
	// Optional field if only 1 asset is involved.
	FairValueNative float64
	// Numerator
	Numerator *big.Int
	// Denominator
	Denominator *big.Int
	Time        time.Time
}
