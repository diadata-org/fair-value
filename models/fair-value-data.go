package models

import (
	"math/big"
	"time"
)

// TO DO: How to integrate symbol? Mabe through scraper.GetConfig()?

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
