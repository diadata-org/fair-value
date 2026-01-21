package models

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	BINANCESMARTCHAIN = "BinanceSmartChain"
	BITCOIN           = "Bitcoin"
	ETHEREUM          = "Ethereum"
	ARBITRUM_SEPOLIA  = "Arbitrum-Sepolia"
	TONCHAIN          = "Tonchain"
	UNICHAIN          = "Unichain"
)

type MetacontractData struct {
	Address   common.Address
	Precision int
	Client    *ethclient.Client
}
