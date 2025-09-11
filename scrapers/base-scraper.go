package scrapers

import "github.com/diadata-org/fair-value/models"

type BaseScraper struct {
	dataChannel  chan FairValueData
	closeChannel chan bool
	config       models.FeedConfig
}

func NewBaseScraper() BaseScraper {
	return BaseScraper{
		dataChannel:  make(chan FairValueData),
		closeChannel: make(chan bool),
	}
}

func (b *BaseScraper) DataChannel() chan FairValueData {
	return b.dataChannel
}

func (b *BaseScraper) Close() chan bool {
	return b.closeChannel
}

// TO DO: Do we need a config getter? Then we also need a config setter.
func (b *BaseScraper) GetFeedConfig() models.FeedConfig {
	return b.config
}
