package scrapers

type BaseScraper struct {
	dataChannel  chan FairValueData
	closeChannel chan bool
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
