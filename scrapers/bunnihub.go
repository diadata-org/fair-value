package scrapers

import (
	"math"
	"math/big"

	"github.com/diadata-org/fair-value/models"
	"github.com/diadata-org/fair-value/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	erc20 "github.com/diadata-org/fair-value/contracts/erc20"
	univ4 "github.com/diadata-org/fair-value/contracts/univ4poolmanager"
)

// ------------------------------------------------------------------
// CONTRACT EXCHANGE RATE
// ------------------------------------------------------------------

const (
	// UniswapV4 pool manager.
	UNISWAPV4_POOL_MANAGER = "0x000000000049C7bcBCa294E63567b4D21EB765f1"
	// POOL_ID = "0x5e5bc151fdb581faf0c28ae30ceba9193da793ccd7c22a70d3feaf3408c07666"
	// LP_TOKEN_ADDRESS = "0x78fd58693ff7796fDF565bD744fdC21CB9B49C6c"
)

type BunnihubScraper struct {
	BaseScraper
	client          *ethclient.Client
	blockchain      string
	contractAddress common.Address
	poolID          common.Hash
	lpTokenAddress  common.Address
	config          models.FeedConfig
}

func NewBunnihubScraper(config models.FeedConfig) *BunnihubScraper {

	scraper := BunnihubScraper{
		BaseScraper:     NewBaseScraper(),
		blockchain:      config.Blockchain,
		contractAddress: common.HexToAddress(config.Address),
		config:          config,
	}

	if len(config.Params) > 0 {
		if p, ok := config.Params[0].(string); ok {
			scraper.poolID = common.HexToHash(p)
		}
	}

	scraper.lpTokenAddress = common.HexToAddress(config.Address)

	client, err := ethclient.Dial(utils.Getenv("RPC_NODE_BUNNIHUB", ""))
	if err != nil {
		log.Errorf("bunnihub -- make eth client for %s: %v", config.Symbol, err)
	}
	scraper.client = client

	return &scraper

}

func (scraper *BunnihubScraper) TotalUnderlying() (totalUnderlying *big.Int, totalValueUnderlying *big.Int, err error) {
	bunnihubCaller, err := univ4.NewUniv4poolmanagerCaller(common.HexToAddress(UNISWAPV4_POOL_MANAGER), scraper.client)
	if err != nil {
		return
	}
	poolBalances, err := bunnihubCaller.PoolBalances(&bind.CallOpts{}, scraper.poolID)
	if err != nil {
		return
	}

	ethBalance := new(big.Float).Quo(new(big.Float).SetInt(poolBalances.Balance0), big.NewFloat(1e18))
	usdcBalance := new(big.Float).Quo(new(big.Float).SetInt(poolBalances.Balance1), big.NewFloat(1e6))

	// TO DO: call vault assets from contract as follows.
	// poolState, err := bunnihubCaller.PoolState(&bind.CallOpts{}, poolID)
	// if err != nil {
	// 	log.Error("bunnihub -- PoolState: ", err)
	// }
	// poolState.Vault0 and poolState.Vault1 -> call "asset" on the respective vault.

	// DIA Prices
	ethPrice, err := utils.GetDiaQuotationPrice(models.ETHEREUM, "0x0000000000000000000000000000000000000000")
	if err != nil {
		log.Error("bunnihub -- GetDiaQuotationPrice for ETH: ", err)
		return
	}
	usdcPrice, err := utils.GetDiaQuotationPrice(models.ETHEREUM, "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
	if err != nil {
		log.Error("bunnihub -- GetDiaQuotationPrice for USDC: ", err)
		return
	}

	// Scaled sum of values.
	ethValue := new(big.Float).Mul(ethBalance, big.NewFloat(ethPrice))
	usdcValue := new(big.Float).Mul(usdcBalance, big.NewFloat(usdcPrice))
	totalValue := new(big.Float).Add(ethValue, usdcValue)
	totalValueUnderlying, _ = new(big.Float).Mul(totalValue, new(big.Float).SetFloat64(math.Pow10(int(DECIMALS)))).Int(nil)

	return
}

func (scraper *BunnihubScraper) TotalShares() (totalSupply *big.Int, err error) {
	erc20Caller, err := erc20.NewERC20Caller(scraper.lpTokenAddress, scraper.client)
	if err != nil {
		return
	}
	totalSupply, err = erc20Caller.TotalSupply(&bind.CallOpts{})
	return
}

func (scraper *BunnihubScraper) DataChannel() chan models.FairValueData {
	return scraper.dataChannel
}

func (scraper *BunnihubScraper) GetConfig() models.FeedConfig {
	return scraper.config
}

func (scraper *BunnihubScraper) Close() chan bool {
	return scraper.BaseScraper.Close()
}
