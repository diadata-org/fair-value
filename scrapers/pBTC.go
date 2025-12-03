package scrapers

import (
	"context"
	"math/big"
	"strconv"

	pbtc "github.com/diadata-org/fair-value/contracts/pbtc"
	"github.com/diadata-org/fair-value/models"
	"github.com/diadata-org/fair-value/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type pBTCScraper struct {
	BaseScraper
	client          *ethclient.Client
	bitcoinRPC      string
	blockchain      string
	contractAddress common.Address
	poolID          common.Hash
	lpTokenAddress  common.Address
	config          models.FeedConfig
	bitcoinAPI      string
	lastBlock       uint64
	reserveWallets  map[string]struct{}
	chunkSize       uint64
}

func NewpBTCScraper(config models.FeedConfig) *pBTCScraper {

	contractCreation, err := strconv.Atoi(utils.Getenv("CONTRACT_CREATION_PBTC", "216834000"))
	if err != nil {
		log.Error("parse CONTRACT_CREATION_PBTC: ", err)
		contractCreation = 216834000
	}
	scraper := pBTCScraper{
		BaseScraper:     NewBaseScraper(),
		blockchain:      config.Blockchain,
		contractAddress: common.HexToAddress(config.Address),
		bitcoinRPC:      utils.Getenv("BITCOIN_RPC_NODE_PBTC", ""),
		config:          config,
		bitcoinAPI:      utils.Getenv("BITCOIN_API_SOURCE_PBTC", "ANKR"),
		lastBlock:       uint64(contractCreation),
		reserveWallets:  make(map[string]struct{}),
		chunkSize:       uint64(10000),
	}

	if len(config.Params) > 0 {
		scraper.poolID = common.HexToHash(config.Params[0].(string))
	}

	scraper.lpTokenAddress = common.HexToAddress(config.Address)

	client, err := ethclient.Dial(utils.Getenv("RPC_NODE_EVM_PBTC", ""))
	if err != nil {
		log.Errorf("pBTC -- make eth client for %s: %v", config.Symbol, err)
		return nil
	}
	scraper.client = client

	return &scraper

}

func (scraper *pBTCScraper) TotalUnderlying() (totalUnderlying *big.Int, totalValueUnderlying *big.Int, err error) {
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
			}
			if err != nil {
				log.Errorf("pBTC -- getBitcoinWalletBalance for address %s: %v", wallet, err)
				return
			}

			log.Debugf("pBTC -- balance for %s: %v", wallet, balance)
			totalBalance += balance
		}
	}
	log.Debugf("pBTC -- total balance for config wallets: %v", totalBalance)

	err = scraper.getReserveWallets()
	if err != nil {
		return
	}
	log.Debug("pBTC -- reserveWallets: ", scraper.reserveWallets)

	// Add balances from custodian deposit wallet addresses.
	for wallet := range scraper.reserveWallets {
		var balance float64

		switch scraper.bitcoinAPI {
		case "CHAINSTACK":
			balance, err = utils.GetBitcoinWalletBalance(wallet, scraper.bitcoinRPC)
		case "ANKR":
			balance, err = utils.GetBitcoinWalletBalanceAnkr(wallet)
		}
		if err != nil {
			log.Errorf("pBTC -- getBitcoinWalletBalance for address %s: %v", wallet, err)
			return
		}

		totalBalance += balance
		log.Debugf("pBTC -- balance for %s: %v", wallet, balance)

	}
	log.Debugf("pBTC -- total balance for all wallets: %v", totalBalance)

	// Convert balance to big.Int with 18 decimals.
	totalUnderlying, _ = big.NewFloat(0).Mul(big.NewFloat(totalBalance), big.NewFloat(1e18)).Int(nil)

	// Compute totalValueUnderlying.
	priceBitcoin, err := utils.GetDiaQuotationPrice(models.BITCOIN, "0x0000000000000000000000000000000000000000")
	if err != nil {
		log.Error("pBTC -- GetDiaQuotationPrice: ", err)
	}
	totalValueUnderlying, _ = big.NewFloat(0).Mul(big.NewFloat(totalBalance*priceBitcoin), big.NewFloat(1e18)).Int(nil)

	return
}

func (scraper *pBTCScraper) TotalShares() (totalShares *big.Int, err error) {
	pbtcCaller, err := pbtc.NewPbtcCaller(scraper.contractAddress, scraper.client)
	if err != nil {
		return
	}
	totalShares, err = pbtcCaller.TotalSupply(&bind.CallOpts{})
	return
}

func (scraper *pBTCScraper) DataChannel() chan models.FairValueData {
	return scraper.dataChannel
}

func (scraper *pBTCScraper) GetConfig() models.FeedConfig {
	return scraper.config
}

func (scraper *pBTCScraper) Close() chan bool {
	return scraper.BaseScraper.Close()
}

func (scraper *pBTCScraper) getReserveWallets() (err error) {
	pbtcFilterer, err := pbtc.NewPbtcFilterer(scraper.contractAddress, scraper.client)
	if err != nil {
		return
	}

	// Iterate over chunks of blocks.
	currentBlockNumber, err := scraper.client.BlockNumber(context.Background())
	if err != nil {
		return
	}
	log.Debugf("pBTC -- startBlock -- currentBlock: %v -- %v ", scraper.lastBlock, currentBlockNumber)
	startblock := scraper.lastBlock
	endblock := scraper.lastBlock + scraper.chunkSize

	for endblock < currentBlockNumber+scraper.chunkSize {

		if endblock > currentBlockNumber {
			endblock = currentBlockNumber
		}
		if currentBlockNumber > endblock {
			log.Tracef("pBTC -- blocks left: %v", currentBlockNumber-endblock)
		}
		log.Debugf("pBTC -- filter custodian wallets from block %v to %v", startblock, endblock)
		setDepositAddress, err := pbtcFilterer.FilterCustodianBtcDepositAddressSet(
			&bind.FilterOpts{
				Start: startblock,
				End:   &endblock,
			},
			[]common.Address{},
			[]common.Address{},
		)
		if err != nil {
			log.Error("pBTC -- filterCustodianBtcDepositAddressSet : ", err)
			startblock = endblock
			endblock = startblock + scraper.chunkSize
			continue
		}

		for setDepositAddress.Next() {
			if _, ok := scraper.reserveWallets[setDepositAddress.Event.BtcDepositAddress]; !ok {
				scraper.reserveWallets[setDepositAddress.Event.BtcDepositAddress] = struct{}{}
			}
		}

		// increment block range for the next iteration.
		startblock = endblock
		endblock = startblock + scraper.chunkSize
	}

	scraper.lastBlock = startblock

	return
}
