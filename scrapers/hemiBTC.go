package scrapers

import (
	"math/big"

	hemibtc "github.com/diadata-org/fair-value/contracts/hemibtc"
	"github.com/diadata-org/fair-value/models"
	"github.com/diadata-org/fair-value/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type hemiBTCScraper struct {
	BaseScraper
	client          *ethclient.Client
	bitcoinRPC      string
	blockchain      string
	contractAddress common.Address
	config          models.FeedConfig
	bitcoinAPI      string
	chunkSize       uint64
}

func NewhemiBTCScraper(config models.FeedConfig, metacontractData models.MetacontractData) *hemiBTCScraper {

	scraper := hemiBTCScraper{
		BaseScraper:     NewBaseScraper(metacontractData),
		blockchain:      config.Blockchain,
		contractAddress: common.HexToAddress(config.Address),
		bitcoinRPC:      utils.Getenv("BITCOIN_RPC_NODE_HEMIBTC", ""),
		config:          config,
		bitcoinAPI:      utils.Getenv("BITCOIN_API_SOURCE_HEMIBTC", "MEMPOOL"),
		chunkSize:       uint64(10000),
	}

	client, err := ethclient.Dial(utils.Getenv("RPC_NODE_EVM_HEMIBTC", ""))
	if err != nil {
		log.Errorf("hemiBTC -- make eth client for %s: %v", config.Symbol, err)
		return nil
	}
	scraper.client = client

	return &scraper

}

func (scraper *hemiBTCScraper) TotalUnderlying() (totalUnderlying *big.Int, totalValueUnderlying *big.Int, err error) {
	var totalBalance float64

	// Compute balances of reserve wallets from config.
	if len(scraper.config.Params) > 0 {
		for _, wallet := range scraper.config.Params[1].([]any) {
			var balance float64

			switch scraper.bitcoinAPI {
			case "CHAINSTACK":
				balance, err = utils.GetBitcoinWalletBalance(wallet.(string), scraper.bitcoinRPC)
			case "ANKR":
				balance, err = utils.GetBitcoinWalletBalanceAnkr(wallet.(string))
			case "MEMPOOL":
				balance, err = utils.GetBitcoinWalletBalanceMempool(wallet.(string))
				if err != nil {
					balance, err = utils.GetBitcoinWalletBalanceAnkr(wallet.(string))
				}
			}
			if err != nil {
				log.Errorf("hemiBTC -- getBitcoinWalletBalance for address %s: %v", wallet, err)
				return
			}

			log.Debugf("hemiBTC -- balance for %s: %v", wallet, balance)
			totalBalance += balance
		}
	}
	log.Debugf("hemiBTC -- total balance for config wallets: %v", totalBalance)

	// Convert balance to big.Int with 18 decimals.
	totalUnderlying, _ = big.NewFloat(0).Mul(big.NewFloat(totalBalance), big.NewFloat(1e18)).Int(nil)

	// Compute totalValueUnderlying.
	btc := models.Asset{Symbol: "BTC", Blockchain: models.BITCOIN, Address: "0x0000000000000000000000000000000000000000"}
	quotationBitcoin, err := btc.GetPrice(scraper.metacontractData.Address, scraper.metacontractData.Precision, scraper.metacontractData.Client)
	if err != nil {
		log.Error("hemiBTC -- GetDiaQuotationPrice: ", err)
	}
	totalValueUnderlying, _ = big.NewFloat(0).Mul(big.NewFloat(totalBalance*quotationBitcoin.Price), big.NewFloat(1e18)).Int(nil)

	return
}

func (scraper *hemiBTCScraper) TotalShares() (totalShares *big.Int, err error) {
	hemibtcCaller, err := hemibtc.NewHemibtcCaller(scraper.contractAddress, scraper.client)
	if err != nil {
		return
	}
	// Total shares are with 8 decimals and have to be normalized to 18.
	totalShares, err = hemibtcCaller.TotalSupply(&bind.CallOpts{})
	totalShares, _ = big.NewFloat(0).Mul(big.NewFloat(0).SetInt(totalShares), big.NewFloat(1e10)).Int(nil)
	return
}

func (scraper *hemiBTCScraper) DataChannel() chan models.FairValueData {
	return scraper.dataChannel
}

func (scraper *hemiBTCScraper) GetConfig() models.FeedConfig {
	return scraper.config
}

func (scraper *hemiBTCScraper) Close() chan bool {
	return scraper.BaseScraper.Close()
}
