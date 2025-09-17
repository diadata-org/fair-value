package scrapers

import (
	"math/big"
	"time"

	"github.com/diadata-org/fair-value/models"
)

// NOTE: All big return values have 18 decimals.
// If the boolean native is set to true, values are considered to be supply.
// Otherwise, values are USD denominations.

const (
	DECIMALS = 18
)

type IScraper interface {
	DataChannel() chan models.FairValueData
	Close() chan bool
	GetConfig() models.FeedConfig
}

// // TO DO: Better to make an overall data factory and switch inside?
// func MakeDataMain(config models.FeedConfig) models.FairValueData {

// 	switch config.FeedType {
// 	case "CONTRACT_EXCHANGE_RATE":
// 		scraper := NewIContractExchangeRate(config)
// 		return MakeCERData(scraper)

// 	case "NET_ASSET_VALUE":
// 		scraper := NewINetAssetValue(config)
// 		return MakeNAVData(scraper)
// 	}
// 	return models.FairValueData{}
// }

type IContractExchangeRate interface {
	IScraper
	//  The total number of underlying assets as a *big.Int.
	// 2. The USD value of item 1.
	// In some cases, for instance when there is several underlying assets,
	// totalUnderlying will be nil as it is not possible to assign it a meaningful value.
	TotalUnderlying() (*big.Int, *big.Int, error)
	TotalShares() (*big.Int, error)
}

type IReserveBacking interface {
	IScraper
	Reserves() (*big.Int, bool, error)
	Supply() (*big.Int, bool, error)
}

// TO DO: Maybe better to add another interface method Native() bool?
// Because I think either both - Assets AND Liabilities are native or not. Total supply should always(?) be native.
type INetAssetValue interface {
	IScraper
	Assets() (*big.Int, bool, error)
	Liabilities() (*big.Int, bool, error)
	TotalSupply() (*big.Int, error)
}

// MAKECERData computes all return values for the IContractExchangeRate interface.
func MakeCERData(scraper IContractExchangeRate) (data models.FairValueData) {

	underlying, underlyingValue, err := scraper.TotalUnderlying()
	if err != nil {
		log.Error("TotalUnderlying: ", err)
	}

	totalShares, err := scraper.TotalShares()
	if err != nil {
		log.Error("TotalShares: ", err)
	}

	config := scraper.GetConfig()
	log.Debugf("underlying -- totalShares for address %s: %s -- %s", config.Address, underlying.String(), totalShares.String())

	// Compute USD price
	var priceUSD float64
	if underlyingValue != nil {
		numeratorFloat := big.NewFloat(0).SetInt(underlyingValue)
		denominatorFloat := big.NewFloat(0).SetInt(totalShares)
		priceUSD, _ = big.NewFloat(0).Quo(numeratorFloat, denominatorFloat).Float64()
	}

	// Compute fair value native
	var fairValueNative float64
	if underlying != nil {
		numeratorFloat := big.NewFloat(0).SetInt(underlying)
		denominatorFloat := big.NewFloat(0).SetInt(totalShares)
		fairValueNative, _ = big.NewFloat(0).Quo(numeratorFloat, denominatorFloat).Float64()
	}

	data.Symbol = config.Symbol
	data.Address = config.Address
	data.Blockchain = config.Blockchain
	data.Numerator = underlying
	data.Denominator = totalShares
	if underlying == nil {
		data.Numerator = big.NewInt(0)
	}
	if totalShares == nil {
		data.Denominator = big.NewInt(0)
	}
	data.PriceUSD = priceUSD
	data.FairValueNative = fairValueNative
	data.Time = time.Now()
	log.Debugf("price for address %s: %v", config.Address, data.PriceUSD)

	return
}

// TO DO: Is this the best way to produce data?
// MakeNAVData computes all return values for the INetAssetValue interface.
func MakeNAVData(scraper INetAssetValue) (data models.FairValueData) {
	config := scraper.GetConfig()

	assets, _, err := scraper.Assets()
	if err != nil {
		log.Errorf("scraper.Assets() for %s: %v", config.Symbol, err)
		return
	}
	liabilities, _, err := scraper.Liabilities()
	if err != nil {
		log.Errorf("scraper.Liabilities() for %s: %v", config.Symbol, err)
		return
	}
	supply, err := scraper.TotalSupply()
	if err != nil {
		log.Errorf("TotalSupply for %s: %v", config.Symbol, err)
		return
	}

	// Compute (assets-liabilities)/supply
	numerator := new(big.Int).Sub(assets, liabilities)
	numeratorFloat := big.NewFloat(0).SetInt(numerator)
	denominator := big.NewFloat(0).SetInt(supply)
	price := big.NewFloat(0).Quo(numeratorFloat, denominator)

	data.Symbol = config.Symbol
	data.Address = config.Address
	data.Blockchain = config.Blockchain
	data.Numerator = numerator
	data.Denominator = supply
	if numerator == nil {
		data.Numerator = big.NewInt(0)
	}
	if supply == nil {
		data.Denominator = big.NewInt(0)
	}
	data.PriceUSD, _ = price.Float64()
	data.Time = time.Now()

	return
}
