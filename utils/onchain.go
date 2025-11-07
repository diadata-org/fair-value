package utils

import (
	"crypto/ecdsa"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func SetupOnchain(
	blockchainNode string,
	backupNode string,
	privateKeyHex string,
	chainId int64,
) (
	auth *bind.TransactOpts,
	conn *ethclient.Client,
	connBackup *ethclient.Client,
	privateKey *ecdsa.PrivateKey,
	err error,
) {
	conn, err = MakeEthClient(blockchainNode, backupNode)
	if err != nil {
		return
	}
	connBackup, err = MakeEthClient(backupNode, blockchainNode)
	if err != nil {
		return
	}

	privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x")
	privateKey, err = crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return
	}

	auth, err = bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainId))
	if err != nil {
		return
	}

	return auth, conn, connBackup, privateKey, nil
}
