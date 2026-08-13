package scrapers

import "github.com/diadata-org/fair-value/models"

type BaseScraper struct {
	dataChannel      chan models.FairValueData
	closeChannel     chan bool
	metacontractData models.MetacontractData
	capFairValueUSD  *float64
}

func NewBaseScraper(metacontractData models.MetacontractData) BaseScraper {
	return BaseScraper{
		dataChannel:      make(chan models.FairValueData),
		closeChannel:     make(chan bool),
		metacontractData: metacontractData,
	}
}

func (b *BaseScraper) DataChannel() chan models.FairValueData {
	return b.dataChannel
}

func (b *BaseScraper) Close() chan bool {
	return b.closeChannel
}

func (b *BaseScraper) CapFairValueUSD() *float64 {
	if b.capFairValueUSD == nil {
		return nil
	}
	v := *b.capFairValueUSD
	return &v
}

func (b *BaseScraper) SetCapFairValueUSD(v *float64) {
	if v == nil {
		b.capFairValueUSD = nil
		return
	}
	x := *v
	b.capFairValueUSD = &x
}
