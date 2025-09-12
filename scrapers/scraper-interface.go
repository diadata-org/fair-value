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
	// TO DO: Should we make this Close() error and implement Close() on scraper level as closeChannel <- true?
	Close() chan bool
}

// TO DO: Better to make an overall data factory and switch inside?
// func MakeData(feedType string, address string, blockchain string, params []any) FairValueData {

// 	switch feedType {
// 	case "CONTRACT_EXCHANGE_RATE":
// 		scraper := NewIContractExchangeRate(blockchain, address, params)
// 		return MakeCERData(scraper, address, blockchain)

// 	case "NET_ASSET_VALUE":
// 		scraper := NewINetAssetValue(blockchain, address, params)
// 		return MakeNAVData(scraper, address, blockchain)
// 	}
// 	return FairValueData{}
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
func MakeCERData(scraper IContractExchangeRate, address string, blockchain string) (data models.FairValueData) {

	underlying, underlyingValue, err := scraper.TotalUnderlying()
	if err != nil {
		log.Error("TotalUnderlying: ", err)
	}

	totalShares, err := scraper.TotalShares()
	if err != nil {
		log.Error("TotalShares: ", err)
	}

	log.Debugf("underlying -- totalShares for address %s: %s -- %s", address, underlying.String(), totalShares.String())

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

	data.Address = address
	data.Blockchain = blockchain
	data.Numerator = underlying
	data.Denominator = totalShares
	data.PriceUSD = priceUSD
	data.FairValueNative = fairValueNative
	data.Time = time.Now()
	log.Debugf("price for address %s: %v", address, data.PriceUSD)

	return
}

// TO DO: Is this the best way to produce data?
// MakeNAVData computes all return values for the INetAssetValue interface.
func MakeNAVData(scraper INetAssetValue, address string, blockchain string) (data models.FairValueData) {

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

	// Compute (assets-liabilities)/supply
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
