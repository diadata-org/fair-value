package utils

import (
	"github.com/diadata-org/fair-value/contracts/erc20"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetEthDecimals(address common.Address, client *ethclient.Client) (decimals uint8, err error) {
	erc20Caller, err := erc20.NewERC20MetadataCaller(address, client)
	if err != nil {
		return
	}
	return erc20Caller.Decimals(&bind.CallOpts{})
}
