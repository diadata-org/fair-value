package scrapers

import "github.com/diadata-org/fair-value/models"

type BaseScraper struct {
	dataChannel  chan models.FairValueData
	closeChannel chan bool
}

func NewBaseScraper() BaseScraper {
	return BaseScraper{
		dataChannel:  make(chan models.FairValueData),
		closeChannel: make(chan bool),
	}
}

func (b *BaseScraper) DataChannel() chan models.FairValueData {
	return b.dataChannel
}

func (b *BaseScraper) Close() chan bool {
	return b.closeChannel
}
