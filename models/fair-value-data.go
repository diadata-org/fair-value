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
	FeedType   string
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

// TO DO: How to address @Params? Can we assume Params[0] is always the feed differentiator?
func (fvd *FairValueData) FairValueDataIdentifier() string {
	return fvd.FeedType + FEED_CONFIG_SEPARATOR + fvd.Blockchain + FEED_CONFIG_SEPARATOR + fvd.Address
}
