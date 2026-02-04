package scrapers

import (
	"fmt"
	"math"
	"math/big"
	"strings"

	"github.com/diadata-org/fair-value/models"
	"github.com/diadata-org/fair-value/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	usdr "github.com/diadata-org/fair-value/contracts/usdr"
)

// ------------------------------------------------------------------
// CONTRACT EXCHANGE RATE
// ------------------------------------------------------------------

type USDRScraper struct {
	BaseScraper
	blockchain      string
	contractAddress string
	dataAddresses   map[string]string
	config          models.FeedConfig
}

func NewUSDRScraperScraper(config models.FeedConfig, metacontractData models.MetacontractData) *USDRScraper {
	scraper := USDRScraper{
		BaseScraper:     NewBaseScraper(metacontractData),
		blockchain:      config.Blockchain,
		contractAddress: config.Address,
		dataAddresses:   make(map[string]string),
		config:          config,
	}

	scraper.getDataAddresses()

	return &scraper
}

func (scraper *USDRScraper) TotalUnderlying() (totalUnderlying *big.Int, totalValueUnderlying *big.Int, err error) {

	rpcNodes := makeDataRPCMap()
	var totalValue float64

	// Iterate over blockchains and sum up.
	for blockchain, address := range scraper.dataAddresses {
		if _, ok := rpcNodes[blockchain]; !ok {
			log.Warnf("USDR - no rpc node available for blockchain %s", blockchain)
			continue
		}
		var usdpIssued float64
		usdpIssued, err = getChainSpecificUsdpIssued(rpcNodes[blockchain], address)
		if err != nil {
			return
		}
		var usdpPrice float64
		usdpPrice, err = getChainSpecificUsdpPrice(rpcNodes[blockchain], address)
		if err != nil {
			return
		}
		totalValue += usdpIssued * usdpPrice
	}

	// Scale to 18 decimals.
	totalValueUnderlying, _ = big.NewFloat(0).Mul(big.NewFloat(totalValue), big.NewFloat(math.Pow10(DECIMALS))).Int(nil)

	return
}

func (scraper *USDRScraper) TotalShares() (*big.Int, error) {
	var total float64
	rpcNodes := makeDataRPCMap()

	for blockchain, address := range scraper.dataAddresses {
		if _, ok := rpcNodes[blockchain]; !ok {
			return nil, fmt.Errorf("USDR - missing node for blockchain %s", blockchain)
		}
		usdpIssued, err := getChainSpecificUsdpIssued(rpcNodes[blockchain], address)
		if err != nil {
			return nil, err
		}
		total += usdpIssued
	}

	// Normalize to 18 decimals
	totalShares, _ := big.NewFloat(0).Mul(big.NewFloat(total), big.NewFloat(math.Pow10(DECIMALS))).Int(nil)
	return totalShares, nil

}

func (scraper *USDRScraper) DataChannel() chan models.FairValueData {
	return scraper.dataChannel
}

func (scraper *USDRScraper) Close() chan bool {
	return scraper.BaseScraper.Close()
}

func (scraper *USDRScraper) GetConfig() models.FeedConfig {
	return scraper.config
}

func getChainSpecificUsdpPrice(dataNode string, diaOracleAddress string) (usdpPrice float64, err error) {
	dataConn, err := ethclient.Dial(dataNode)
	if err != nil {
		return
	}

	diaOracleConn, err := usdr.NewUsdrCaller(common.HexToAddress(diaOracleAddress), dataConn)
	if err != nil {
		return
	}
	chainSpecificUsdpPrice, err := diaOracleConn.GetChainSpecificUsdpPrice(&bind.CallOpts{})
	if err != nil {
		return
	}
	outputDecimalsBI, err := diaOracleConn.OutputDecimals(&bind.CallOpts{})
	if err != nil {
		return
	}

	outputDecimals := outputDecimalsBI.Int64()
	usdpPrice, _ = chainSpecificUsdpPrice.Float64()
	usdpPrice /= math.Pow(10, float64(outputDecimals))

	return
}

func getChainSpecificUsdpIssued(dataNode string, diaOracleAddress string) (usdpIssued float64, err error) {
	dataConn, err := ethclient.Dial(dataNode)
	if err != nil {
		return
	}

	diaOracleConn, err := usdr.NewUsdrCaller(common.HexToAddress(diaOracleAddress), dataConn)
	if err != nil {
		return
	}

	chainSpecificUsdpIssued, err := diaOracleConn.GetChainSpecificUsdpIssued(&bind.CallOpts{})
	if err != nil {
		return
	}
	usdpIssued, _ = chainSpecificUsdpIssued.Float64()
	usdpIssued /= math.Pow10(DECIMALS)
	return
}

func (scraper *USDRScraper) getDataAddresses() {

	if len(scraper.config.Params) > 0 {
		dataAddresses := make(map[string]string)
		for _, chainData := range scraper.config.Params[1].([]any) {
			if s := strings.Split(chainData.(string), "-"); len(s) == 2 {
				dataAddresses[s[0]] = s[1]
			}
		}
		scraper.dataAddresses = dataAddresses
	}

}

func makeDataRPCMap() map[string]string {
	m := make(map[string]string)
	rpcNodes := utils.Getenv("DATA_RPC_NODES_USDR", "")
	for _, node := range strings.Split(rpcNodes, ",") {
		if s := strings.Split(node, "§"); len(s) == 2 {
			m[s[0]] = s[1]
		}
	}
	return m
}
