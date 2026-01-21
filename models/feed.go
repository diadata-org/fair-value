package models

import (
	"encoding/json"
	"reflect"

	"github.com/diadata-org/fair-value/utils"
)

const FEED_CONFIG_SEPARATOR = "-"

// We have the following convention for the @Params slice:
// Params[0]: poolID 				string
// Params[1]: wallet addresses 		[]string
type FeedConfig struct {
	Symbol        string `json:"Symbol"`
	FeedType      string `json:"Feed_Type"`
	Address       string `json:"Address"`
	Blockchain    string `json:"Blockchain"`
	UpdateSeconds int    `json:"Update_Seconds"`
	Params        []any  `json:"Params"`
}

// TO DO: How to address @Params? Can we assume Params[0] is always the feed differentiator?
func (fc *FeedConfig) FeedConfigIdentifier() string {
	return fc.FeedType + FEED_CONFIG_SEPARATOR + fc.Blockchain + FEED_CONFIG_SEPARATOR + fc.Address
}

// Contained returns true when @fc is contained in @fcSlice.
func (fc *FeedConfig) Contained(fcSlice []FeedConfig) bool {
	for _, item := range fcSlice {
		if fc.IsEqual(item) {
			return true
		}
	}
	return false
}

func (fc *FeedConfig) IsEqual(fc1 FeedConfig) bool {
	return reflect.DeepEqual(*fc, fc1)
}

func GetDiffConfig(fcOld, fcNew []FeedConfig) (plus []FeedConfig, minus []FeedConfig) {

	for _, fc := range fcOld {
		if !fc.Contained(fcNew) {
			minus = append(minus, fc)
		}
	}

	for _, fc := range fcNew {
		if !fc.Contained(fcOld) {
			plus = append(plus, fc)
		}
	}
	return
}

func GetFeedsFromConfig(filename string, remoteConfig bool, branchConfig string) ([]FeedConfig, error) {

	data, err := utils.ReadFile(filename, remoteConfig, branchConfig)
	if err != nil {
		return []FeedConfig{}, err
	}

	type feedConfigs struct {
		Feeds []FeedConfig `json:"Feeds"`
	}
	var fc feedConfigs
	err = json.Unmarshal(data, &fc)
	if err != nil {
		return []FeedConfig{}, err
	}
	return fc.Feeds, nil

}
