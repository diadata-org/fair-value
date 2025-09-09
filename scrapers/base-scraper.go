package scrapers

type BaseScraper struct {
	dataChannel chan FairValueData
}

func NewBaseScraper() BaseScraper {
	return BaseScraper{
		dataChannel: make(chan FairValueData),
	}
}

func (b *BaseScraper) DataChannel() chan FairValueData {
	return b.dataChannel
}
