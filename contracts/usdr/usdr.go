// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package usdr

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// UsdrMetaData contains all meta data concerning the Usdr contract.
var UsdrMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_getterAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_redeemerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_susdpVaultAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_outputDecimals\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"}],\"name\":\"OracleUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"QUERY_DECIMALS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USDP_MARKET_QUERY_STRING\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USDP_QUERY_STRING\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainSpecificUsdpIssued\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainSpecificUsdpPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSusdpMarketBasedPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSusdpPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getValue\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getterAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"outputDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newGetterAddress\",\"type\":\"address\"}],\"name\":\"setGetterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"compressedValues\",\"type\":\"uint256[]\"}],\"name\":\"setMultipleValues\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newDecimals\",\"type\":\"uint256\"}],\"name\":\"setOutputDecimals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newRedeemerAddress\",\"type\":\"address\"}],\"name\":\"setRedeemerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newSusdpVaultAddress\",\"type\":\"address\"}],\"name\":\"setSusdpVaultAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newUsdpMarketQueryString\",\"type\":\"string\"}],\"name\":\"setUsdpMarketQueryString\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newUsdpQueryString\",\"type\":\"string\"}],\"name\":\"setUsdpQueryString\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"}],\"name\":\"setValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"susdpVaultAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"values\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// UsdrABI is the input ABI used to generate the binding from.
// Deprecated: Use UsdrMetaData.ABI instead.
var UsdrABI = UsdrMetaData.ABI

// Usdr is an auto generated Go binding around an Ethereum contract.
type Usdr struct {
	UsdrCaller     // Read-only binding to the contract
	UsdrTransactor // Write-only binding to the contract
	UsdrFilterer   // Log filterer for contract events
}

// UsdrCaller is an auto generated read-only Go binding around an Ethereum contract.
type UsdrCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsdrTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UsdrTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsdrFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UsdrFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsdrSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UsdrSession struct {
	Contract     *Usdr             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UsdrCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UsdrCallerSession struct {
	Contract *UsdrCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UsdrTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UsdrTransactorSession struct {
	Contract     *UsdrTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UsdrRaw is an auto generated low-level Go binding around an Ethereum contract.
type UsdrRaw struct {
	Contract *Usdr // Generic contract binding to access the raw methods on
}

// UsdrCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UsdrCallerRaw struct {
	Contract *UsdrCaller // Generic read-only contract binding to access the raw methods on
}

// UsdrTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UsdrTransactorRaw struct {
	Contract *UsdrTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUsdr creates a new instance of Usdr, bound to a specific deployed contract.
func NewUsdr(address common.Address, backend bind.ContractBackend) (*Usdr, error) {
	contract, err := bindUsdr(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Usdr{UsdrCaller: UsdrCaller{contract: contract}, UsdrTransactor: UsdrTransactor{contract: contract}, UsdrFilterer: UsdrFilterer{contract: contract}}, nil
}

// NewUsdrCaller creates a new read-only instance of Usdr, bound to a specific deployed contract.
func NewUsdrCaller(address common.Address, caller bind.ContractCaller) (*UsdrCaller, error) {
	contract, err := bindUsdr(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UsdrCaller{contract: contract}, nil
}

// NewUsdrTransactor creates a new write-only instance of Usdr, bound to a specific deployed contract.
func NewUsdrTransactor(address common.Address, transactor bind.ContractTransactor) (*UsdrTransactor, error) {
	contract, err := bindUsdr(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UsdrTransactor{contract: contract}, nil
}

// NewUsdrFilterer creates a new log filterer instance of Usdr, bound to a specific deployed contract.
func NewUsdrFilterer(address common.Address, filterer bind.ContractFilterer) (*UsdrFilterer, error) {
	contract, err := bindUsdr(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UsdrFilterer{contract: contract}, nil
}

// bindUsdr binds a generic wrapper to an already deployed contract.
func bindUsdr(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UsdrMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Usdr *UsdrRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Usdr.Contract.UsdrCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Usdr *UsdrRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Usdr.Contract.UsdrTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Usdr *UsdrRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Usdr.Contract.UsdrTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Usdr *UsdrCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Usdr.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Usdr *UsdrTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Usdr.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Usdr *UsdrTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Usdr.Contract.contract.Transact(opts, method, params...)
}

// QUERYDECIMALS is a free data retrieval call binding the contract method 0xf904dbcc.
//
// Solidity: function QUERY_DECIMALS() view returns(uint256)
func (_Usdr *UsdrCaller) QUERYDECIMALS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Usdr.contract.Call(opts, &out, "QUERY_DECIMALS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QUERYDECIMALS is a free data retrieval call binding the contract method 0xf904dbcc.
//
// Solidity: function QUERY_DECIMALS() view returns(uint256)
func (_Usdr *UsdrSession) QUERYDECIMALS() (*big.Int, error) {
	return _Usdr.Contract.QUERYDECIMALS(&_Usdr.CallOpts)
}

// QUERYDECIMALS is a free data retrieval call binding the contract method 0xf904dbcc.
//
// Solidity: function QUERY_DECIMALS() view returns(uint256)
func (_Usdr *UsdrCallerSession) QUERYDECIMALS() (*big.Int, error) {
	return _Usdr.Contract.QUERYDECIMALS(&_Usdr.CallOpts)
}

// USDPMARKETQUERYSTRING is a free data retrieval call binding the contract method 0x24d1c964.
//
// Solidity: function USDP_MARKET_QUERY_STRING() view returns(string)
func (_Usdr *UsdrCaller) USDPMARKETQUERYSTRING(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Usdr.contract.Call(opts, &out, "USDP_MARKET_QUERY_STRING")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// USDPMARKETQUERYSTRING is a free data retrieval call binding the contract method 0x24d1c964.
//
// Solidity: function USDP_MARKET_QUERY_STRING() view returns(string)
func (_Usdr *UsdrSession) USDPMARKETQUERYSTRING() (string, error) {
	return _Usdr.Contract.USDPMARKETQUERYSTRING(&_Usdr.CallOpts)
}

// USDPMARKETQUERYSTRING is a free data retrieval call binding the contract method 0x24d1c964.
//
// Solidity: function USDP_MARKET_QUERY_STRING() view returns(string)
func (_Usdr *UsdrCallerSession) USDPMARKETQUERYSTRING() (string, error) {
	return _Usdr.Contract.USDPMARKETQUERYSTRING(&_Usdr.CallOpts)
}

// USDPQUERYSTRING is a free data retrieval call binding the contract method 0x2b3eb84e.
//
// Solidity: function USDP_QUERY_STRING() view returns(string)
func (_Usdr *UsdrCaller) USDPQUERYSTRING(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Usdr.contract.Call(opts, &out, "USDP_QUERY_STRING")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// USDPQUERYSTRING is a free data retrieval call binding the contract method 0x2b3eb84e.
//
// Solidity: function USDP_QUERY_STRING() view returns(string)
func (_Usdr *UsdrSession) USDPQUERYSTRING() (string, error) {
	return _Usdr.Contract.USDPQUERYSTRING(&_Usdr.CallOpts)
}

// USDPQUERYSTRING is a free data retrieval call binding the contract method 0x2b3eb84e.
//
// Solidity: function USDP_QUERY_STRING() view returns(string)
func (_Usdr *UsdrCallerSession) USDPQUERYSTRING() (string, error) {
	return _Usdr.Contract.USDPQUERYSTRING(&_Usdr.CallOpts)
}

// GetChainSpecificUsdpIssued is a free data retrieval call binding the contract method 0xa44cee4f.
//
// Solidity: function getChainSpecificUsdpIssued() view returns(uint256)
func (_Usdr *UsdrCaller) GetChainSpecificUsdpIssued(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Usdr.contract.Call(opts, &out, "getChainSpecificUsdpIssued")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainSpecificUsdpIssued is a free data retrieval call binding the contract method 0xa44cee4f.
//
// Solidity: function getChainSpecificUsdpIssued() view returns(uint256)
func (_Usdr *UsdrSession) GetChainSpecificUsdpIssued() (*big.Int, error) {
	return _Usdr.Contract.GetChainSpecificUsdpIssued(&_Usdr.CallOpts)
}

// GetChainSpecificUsdpIssued is a free data retrieval call binding the contract method 0xa44cee4f.
//
// Solidity: function getChainSpecificUsdpIssued() view returns(uint256)
func (_Usdr *UsdrCallerSession) GetChainSpecificUsdpIssued() (*big.Int, error) {
	return _Usdr.Contract.GetChainSpecificUsdpIssued(&_Usdr.CallOpts)
}

// GetChainSpecificUsdpPrice is a free data retrieval call binding the contract method 0xf190fb89.
//
// Solidity: function getChainSpecificUsdpPrice() view returns(uint256)
func (_Usdr *UsdrCaller) GetChainSpecificUsdpPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Usdr.contract.Call(opts, &out, "getChainSpecificUsdpPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainSpecificUsdpPrice is a free data retrieval call binding the contract method 0xf190fb89.
//
// Solidity: function getChainSpecificUsdpPrice() view returns(uint256)
func (_Usdr *UsdrSession) GetChainSpecificUsdpPrice() (*big.Int, error) {
	return _Usdr.Contract.GetChainSpecificUsdpPrice(&_Usdr.CallOpts)
}

// GetChainSpecificUsdpPrice is a free data retrieval call binding the contract method 0xf190fb89.
//
// Solidity: function getChainSpecificUsdpPrice() view returns(uint256)
func (_Usdr *UsdrCallerSession) GetChainSpecificUsdpPrice() (*big.Int, error) {
	return _Usdr.Contract.GetChainSpecificUsdpPrice(&_Usdr.CallOpts)
}

// GetSusdpMarketBasedPrice is a free data retrieval call binding the contract method 0xfbd29009.
//
// Solidity: function getSusdpMarketBasedPrice() view returns(uint256)
func (_Usdr *UsdrCaller) GetSusdpMarketBasedPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Usdr.contract.Call(opts, &out, "getSusdpMarketBasedPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSusdpMarketBasedPrice is a free data retrieval call binding the contract method 0xfbd29009.
//
// Solidity: function getSusdpMarketBasedPrice() view returns(uint256)
func (_Usdr *UsdrSession) GetSusdpMarketBasedPrice() (*big.Int, error) {
	return _Usdr.Contract.GetSusdpMarketBasedPrice(&_Usdr.CallOpts)
}

// GetSusdpMarketBasedPrice is a free data retrieval call binding the contract method 0xfbd29009.
//
// Solidity: function getSusdpMarketBasedPrice() view returns(uint256)
func (_Usdr *UsdrCallerSession) GetSusdpMarketBasedPrice() (*big.Int, error) {
	return _Usdr.Contract.GetSusdpMarketBasedPrice(&_Usdr.CallOpts)
}

// GetSusdpPrice is a free data retrieval call binding the contract method 0xa1b6012f.
//
// Solidity: function getSusdpPrice() view returns(uint256)
func (_Usdr *UsdrCaller) GetSusdpPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Usdr.contract.Call(opts, &out, "getSusdpPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSusdpPrice is a free data retrieval call binding the contract method 0xa1b6012f.
//
// Solidity: function getSusdpPrice() view returns(uint256)
func (_Usdr *UsdrSession) GetSusdpPrice() (*big.Int, error) {
	return _Usdr.Contract.GetSusdpPrice(&_Usdr.CallOpts)
}

// GetSusdpPrice is a free data retrieval call binding the contract method 0xa1b6012f.
//
// Solidity: function getSusdpPrice() view returns(uint256)
func (_Usdr *UsdrCallerSession) GetSusdpPrice() (*big.Int, error) {
	return _Usdr.Contract.GetSusdpPrice(&_Usdr.CallOpts)
}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint128, uint128)
func (_Usdr *UsdrCaller) GetValue(opts *bind.CallOpts, key string) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Usdr.contract.Call(opts, &out, "getValue", key)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint128, uint128)
func (_Usdr *UsdrSession) GetValue(key string) (*big.Int, *big.Int, error) {
	return _Usdr.Contract.GetValue(&_Usdr.CallOpts, key)
}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint128, uint128)
func (_Usdr *UsdrCallerSession) GetValue(key string) (*big.Int, *big.Int, error) {
	return _Usdr.Contract.GetValue(&_Usdr.CallOpts, key)
}

// GetterAddress is a free data retrieval call binding the contract method 0x5daf9a97.
//
// Solidity: function getterAddress() view returns(address)
func (_Usdr *UsdrCaller) GetterAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Usdr.contract.Call(opts, &out, "getterAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetterAddress is a free data retrieval call binding the contract method 0x5daf9a97.
//
// Solidity: function getterAddress() view returns(address)
func (_Usdr *UsdrSession) GetterAddress() (common.Address, error) {
	return _Usdr.Contract.GetterAddress(&_Usdr.CallOpts)
}

// GetterAddress is a free data retrieval call binding the contract method 0x5daf9a97.
//
// Solidity: function getterAddress() view returns(address)
func (_Usdr *UsdrCallerSession) GetterAddress() (common.Address, error) {
	return _Usdr.Contract.GetterAddress(&_Usdr.CallOpts)
}

// OutputDecimals is a free data retrieval call binding the contract method 0xbf560ad8.
//
// Solidity: function outputDecimals() view returns(uint256)
func (_Usdr *UsdrCaller) OutputDecimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Usdr.contract.Call(opts, &out, "outputDecimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OutputDecimals is a free data retrieval call binding the contract method 0xbf560ad8.
//
// Solidity: function outputDecimals() view returns(uint256)
func (_Usdr *UsdrSession) OutputDecimals() (*big.Int, error) {
	return _Usdr.Contract.OutputDecimals(&_Usdr.CallOpts)
}

// OutputDecimals is a free data retrieval call binding the contract method 0xbf560ad8.
//
// Solidity: function outputDecimals() view returns(uint256)
func (_Usdr *UsdrCallerSession) OutputDecimals() (*big.Int, error) {
	return _Usdr.Contract.OutputDecimals(&_Usdr.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Usdr *UsdrCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Usdr.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Usdr *UsdrSession) Owner() (common.Address, error) {
	return _Usdr.Contract.Owner(&_Usdr.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Usdr *UsdrCallerSession) Owner() (common.Address, error) {
	return _Usdr.Contract.Owner(&_Usdr.CallOpts)
}

// RedeemerAddress is a free data retrieval call binding the contract method 0x1157b834.
//
// Solidity: function redeemerAddress() view returns(address)
func (_Usdr *UsdrCaller) RedeemerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Usdr.contract.Call(opts, &out, "redeemerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RedeemerAddress is a free data retrieval call binding the contract method 0x1157b834.
//
// Solidity: function redeemerAddress() view returns(address)
func (_Usdr *UsdrSession) RedeemerAddress() (common.Address, error) {
	return _Usdr.Contract.RedeemerAddress(&_Usdr.CallOpts)
}

// RedeemerAddress is a free data retrieval call binding the contract method 0x1157b834.
//
// Solidity: function redeemerAddress() view returns(address)
func (_Usdr *UsdrCallerSession) RedeemerAddress() (common.Address, error) {
	return _Usdr.Contract.RedeemerAddress(&_Usdr.CallOpts)
}

// SusdpVaultAddress is a free data retrieval call binding the contract method 0xe65f4a07.
//
// Solidity: function susdpVaultAddress() view returns(address)
func (_Usdr *UsdrCaller) SusdpVaultAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Usdr.contract.Call(opts, &out, "susdpVaultAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SusdpVaultAddress is a free data retrieval call binding the contract method 0xe65f4a07.
//
// Solidity: function susdpVaultAddress() view returns(address)
func (_Usdr *UsdrSession) SusdpVaultAddress() (common.Address, error) {
	return _Usdr.Contract.SusdpVaultAddress(&_Usdr.CallOpts)
}

// SusdpVaultAddress is a free data retrieval call binding the contract method 0xe65f4a07.
//
// Solidity: function susdpVaultAddress() view returns(address)
func (_Usdr *UsdrCallerSession) SusdpVaultAddress() (common.Address, error) {
	return _Usdr.Contract.SusdpVaultAddress(&_Usdr.CallOpts)
}

// Values is a free data retrieval call binding the contract method 0x5a9ade8b.
//
// Solidity: function values(string ) view returns(uint256)
func (_Usdr *UsdrCaller) Values(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _Usdr.contract.Call(opts, &out, "values", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Values is a free data retrieval call binding the contract method 0x5a9ade8b.
//
// Solidity: function values(string ) view returns(uint256)
func (_Usdr *UsdrSession) Values(arg0 string) (*big.Int, error) {
	return _Usdr.Contract.Values(&_Usdr.CallOpts, arg0)
}

// Values is a free data retrieval call binding the contract method 0x5a9ade8b.
//
// Solidity: function values(string ) view returns(uint256)
func (_Usdr *UsdrCallerSession) Values(arg0 string) (*big.Int, error) {
	return _Usdr.Contract.Values(&_Usdr.CallOpts, arg0)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Usdr *UsdrTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Usdr.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Usdr *UsdrSession) RenounceOwnership() (*types.Transaction, error) {
	return _Usdr.Contract.RenounceOwnership(&_Usdr.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Usdr *UsdrTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Usdr.Contract.RenounceOwnership(&_Usdr.TransactOpts)
}

// SetGetterAddress is a paid mutator transaction binding the contract method 0x45eb8b44.
//
// Solidity: function setGetterAddress(address _newGetterAddress) returns()
func (_Usdr *UsdrTransactor) SetGetterAddress(opts *bind.TransactOpts, _newGetterAddress common.Address) (*types.Transaction, error) {
	return _Usdr.contract.Transact(opts, "setGetterAddress", _newGetterAddress)
}

// SetGetterAddress is a paid mutator transaction binding the contract method 0x45eb8b44.
//
// Solidity: function setGetterAddress(address _newGetterAddress) returns()
func (_Usdr *UsdrSession) SetGetterAddress(_newGetterAddress common.Address) (*types.Transaction, error) {
	return _Usdr.Contract.SetGetterAddress(&_Usdr.TransactOpts, _newGetterAddress)
}

// SetGetterAddress is a paid mutator transaction binding the contract method 0x45eb8b44.
//
// Solidity: function setGetterAddress(address _newGetterAddress) returns()
func (_Usdr *UsdrTransactorSession) SetGetterAddress(_newGetterAddress common.Address) (*types.Transaction, error) {
	return _Usdr.Contract.SetGetterAddress(&_Usdr.TransactOpts, _newGetterAddress)
}

// SetMultipleValues is a paid mutator transaction binding the contract method 0x8d241526.
//
// Solidity: function setMultipleValues(string[] keys, uint256[] compressedValues) returns()
func (_Usdr *UsdrTransactor) SetMultipleValues(opts *bind.TransactOpts, keys []string, compressedValues []*big.Int) (*types.Transaction, error) {
	return _Usdr.contract.Transact(opts, "setMultipleValues", keys, compressedValues)
}

// SetMultipleValues is a paid mutator transaction binding the contract method 0x8d241526.
//
// Solidity: function setMultipleValues(string[] keys, uint256[] compressedValues) returns()
func (_Usdr *UsdrSession) SetMultipleValues(keys []string, compressedValues []*big.Int) (*types.Transaction, error) {
	return _Usdr.Contract.SetMultipleValues(&_Usdr.TransactOpts, keys, compressedValues)
}

// SetMultipleValues is a paid mutator transaction binding the contract method 0x8d241526.
//
// Solidity: function setMultipleValues(string[] keys, uint256[] compressedValues) returns()
func (_Usdr *UsdrTransactorSession) SetMultipleValues(keys []string, compressedValues []*big.Int) (*types.Transaction, error) {
	return _Usdr.Contract.SetMultipleValues(&_Usdr.TransactOpts, keys, compressedValues)
}

// SetOutputDecimals is a paid mutator transaction binding the contract method 0x6b59e8c3.
//
// Solidity: function setOutputDecimals(uint256 _newDecimals) returns()
func (_Usdr *UsdrTransactor) SetOutputDecimals(opts *bind.TransactOpts, _newDecimals *big.Int) (*types.Transaction, error) {
	return _Usdr.contract.Transact(opts, "setOutputDecimals", _newDecimals)
}

// SetOutputDecimals is a paid mutator transaction binding the contract method 0x6b59e8c3.
//
// Solidity: function setOutputDecimals(uint256 _newDecimals) returns()
func (_Usdr *UsdrSession) SetOutputDecimals(_newDecimals *big.Int) (*types.Transaction, error) {
	return _Usdr.Contract.SetOutputDecimals(&_Usdr.TransactOpts, _newDecimals)
}

// SetOutputDecimals is a paid mutator transaction binding the contract method 0x6b59e8c3.
//
// Solidity: function setOutputDecimals(uint256 _newDecimals) returns()
func (_Usdr *UsdrTransactorSession) SetOutputDecimals(_newDecimals *big.Int) (*types.Transaction, error) {
	return _Usdr.Contract.SetOutputDecimals(&_Usdr.TransactOpts, _newDecimals)
}

// SetRedeemerAddress is a paid mutator transaction binding the contract method 0xa2af0419.
//
// Solidity: function setRedeemerAddress(address _newRedeemerAddress) returns()
func (_Usdr *UsdrTransactor) SetRedeemerAddress(opts *bind.TransactOpts, _newRedeemerAddress common.Address) (*types.Transaction, error) {
	return _Usdr.contract.Transact(opts, "setRedeemerAddress", _newRedeemerAddress)
}

// SetRedeemerAddress is a paid mutator transaction binding the contract method 0xa2af0419.
//
// Solidity: function setRedeemerAddress(address _newRedeemerAddress) returns()
func (_Usdr *UsdrSession) SetRedeemerAddress(_newRedeemerAddress common.Address) (*types.Transaction, error) {
	return _Usdr.Contract.SetRedeemerAddress(&_Usdr.TransactOpts, _newRedeemerAddress)
}

// SetRedeemerAddress is a paid mutator transaction binding the contract method 0xa2af0419.
//
// Solidity: function setRedeemerAddress(address _newRedeemerAddress) returns()
func (_Usdr *UsdrTransactorSession) SetRedeemerAddress(_newRedeemerAddress common.Address) (*types.Transaction, error) {
	return _Usdr.Contract.SetRedeemerAddress(&_Usdr.TransactOpts, _newRedeemerAddress)
}

// SetSusdpVaultAddress is a paid mutator transaction binding the contract method 0x73893608.
//
// Solidity: function setSusdpVaultAddress(address _newSusdpVaultAddress) returns()
func (_Usdr *UsdrTransactor) SetSusdpVaultAddress(opts *bind.TransactOpts, _newSusdpVaultAddress common.Address) (*types.Transaction, error) {
	return _Usdr.contract.Transact(opts, "setSusdpVaultAddress", _newSusdpVaultAddress)
}

// SetSusdpVaultAddress is a paid mutator transaction binding the contract method 0x73893608.
//
// Solidity: function setSusdpVaultAddress(address _newSusdpVaultAddress) returns()
func (_Usdr *UsdrSession) SetSusdpVaultAddress(_newSusdpVaultAddress common.Address) (*types.Transaction, error) {
	return _Usdr.Contract.SetSusdpVaultAddress(&_Usdr.TransactOpts, _newSusdpVaultAddress)
}

// SetSusdpVaultAddress is a paid mutator transaction binding the contract method 0x73893608.
//
// Solidity: function setSusdpVaultAddress(address _newSusdpVaultAddress) returns()
func (_Usdr *UsdrTransactorSession) SetSusdpVaultAddress(_newSusdpVaultAddress common.Address) (*types.Transaction, error) {
	return _Usdr.Contract.SetSusdpVaultAddress(&_Usdr.TransactOpts, _newSusdpVaultAddress)
}

// SetUsdpMarketQueryString is a paid mutator transaction binding the contract method 0x4efcfd25.
//
// Solidity: function setUsdpMarketQueryString(string newUsdpMarketQueryString) returns()
func (_Usdr *UsdrTransactor) SetUsdpMarketQueryString(opts *bind.TransactOpts, newUsdpMarketQueryString string) (*types.Transaction, error) {
	return _Usdr.contract.Transact(opts, "setUsdpMarketQueryString", newUsdpMarketQueryString)
}

// SetUsdpMarketQueryString is a paid mutator transaction binding the contract method 0x4efcfd25.
//
// Solidity: function setUsdpMarketQueryString(string newUsdpMarketQueryString) returns()
func (_Usdr *UsdrSession) SetUsdpMarketQueryString(newUsdpMarketQueryString string) (*types.Transaction, error) {
	return _Usdr.Contract.SetUsdpMarketQueryString(&_Usdr.TransactOpts, newUsdpMarketQueryString)
}

// SetUsdpMarketQueryString is a paid mutator transaction binding the contract method 0x4efcfd25.
//
// Solidity: function setUsdpMarketQueryString(string newUsdpMarketQueryString) returns()
func (_Usdr *UsdrTransactorSession) SetUsdpMarketQueryString(newUsdpMarketQueryString string) (*types.Transaction, error) {
	return _Usdr.Contract.SetUsdpMarketQueryString(&_Usdr.TransactOpts, newUsdpMarketQueryString)
}

// SetUsdpQueryString is a paid mutator transaction binding the contract method 0x495ae3cf.
//
// Solidity: function setUsdpQueryString(string newUsdpQueryString) returns()
func (_Usdr *UsdrTransactor) SetUsdpQueryString(opts *bind.TransactOpts, newUsdpQueryString string) (*types.Transaction, error) {
	return _Usdr.contract.Transact(opts, "setUsdpQueryString", newUsdpQueryString)
}

// SetUsdpQueryString is a paid mutator transaction binding the contract method 0x495ae3cf.
//
// Solidity: function setUsdpQueryString(string newUsdpQueryString) returns()
func (_Usdr *UsdrSession) SetUsdpQueryString(newUsdpQueryString string) (*types.Transaction, error) {
	return _Usdr.Contract.SetUsdpQueryString(&_Usdr.TransactOpts, newUsdpQueryString)
}

// SetUsdpQueryString is a paid mutator transaction binding the contract method 0x495ae3cf.
//
// Solidity: function setUsdpQueryString(string newUsdpQueryString) returns()
func (_Usdr *UsdrTransactorSession) SetUsdpQueryString(newUsdpQueryString string) (*types.Transaction, error) {
	return _Usdr.Contract.SetUsdpQueryString(&_Usdr.TransactOpts, newUsdpQueryString)
}

// SetValue is a paid mutator transaction binding the contract method 0x7898e0c2.
//
// Solidity: function setValue(string key, uint128 value, uint128 timestamp) returns()
func (_Usdr *UsdrTransactor) SetValue(opts *bind.TransactOpts, key string, value *big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _Usdr.contract.Transact(opts, "setValue", key, value, timestamp)
}

// SetValue is a paid mutator transaction binding the contract method 0x7898e0c2.
//
// Solidity: function setValue(string key, uint128 value, uint128 timestamp) returns()
func (_Usdr *UsdrSession) SetValue(key string, value *big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _Usdr.Contract.SetValue(&_Usdr.TransactOpts, key, value, timestamp)
}

// SetValue is a paid mutator transaction binding the contract method 0x7898e0c2.
//
// Solidity: function setValue(string key, uint128 value, uint128 timestamp) returns()
func (_Usdr *UsdrTransactorSession) SetValue(key string, value *big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _Usdr.Contract.SetValue(&_Usdr.TransactOpts, key, value, timestamp)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Usdr *UsdrTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Usdr.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Usdr *UsdrSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Usdr.Contract.TransferOwnership(&_Usdr.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Usdr *UsdrTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Usdr.Contract.TransferOwnership(&_Usdr.TransactOpts, newOwner)
}

// UsdrOracleUpdateIterator is returned from FilterOracleUpdate and is used to iterate over the raw logs and unpacked data for OracleUpdate events raised by the Usdr contract.
type UsdrOracleUpdateIterator struct {
	Event *UsdrOracleUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UsdrOracleUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsdrOracleUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UsdrOracleUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UsdrOracleUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsdrOracleUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsdrOracleUpdate represents a OracleUpdate event raised by the Usdr contract.
type UsdrOracleUpdate struct {
	Key       string
	Value     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOracleUpdate is a free log retrieval operation binding the contract event 0xa7fc99ed7617309ee23f63ae90196a1e490d362e6f6a547a59bc809ee2291782.
//
// Solidity: event OracleUpdate(string key, uint128 value, uint128 timestamp)
func (_Usdr *UsdrFilterer) FilterOracleUpdate(opts *bind.FilterOpts) (*UsdrOracleUpdateIterator, error) {

	logs, sub, err := _Usdr.contract.FilterLogs(opts, "OracleUpdate")
	if err != nil {
		return nil, err
	}
	return &UsdrOracleUpdateIterator{contract: _Usdr.contract, event: "OracleUpdate", logs: logs, sub: sub}, nil
}

// WatchOracleUpdate is a free log subscription operation binding the contract event 0xa7fc99ed7617309ee23f63ae90196a1e490d362e6f6a547a59bc809ee2291782.
//
// Solidity: event OracleUpdate(string key, uint128 value, uint128 timestamp)
func (_Usdr *UsdrFilterer) WatchOracleUpdate(opts *bind.WatchOpts, sink chan<- *UsdrOracleUpdate) (event.Subscription, error) {

	logs, sub, err := _Usdr.contract.WatchLogs(opts, "OracleUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsdrOracleUpdate)
				if err := _Usdr.contract.UnpackLog(event, "OracleUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOracleUpdate is a log parse operation binding the contract event 0xa7fc99ed7617309ee23f63ae90196a1e490d362e6f6a547a59bc809ee2291782.
//
// Solidity: event OracleUpdate(string key, uint128 value, uint128 timestamp)
func (_Usdr *UsdrFilterer) ParseOracleUpdate(log types.Log) (*UsdrOracleUpdate, error) {
	event := new(UsdrOracleUpdate)
	if err := _Usdr.contract.UnpackLog(event, "OracleUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UsdrOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Usdr contract.
type UsdrOwnershipTransferredIterator struct {
	Event *UsdrOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UsdrOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsdrOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UsdrOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UsdrOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsdrOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsdrOwnershipTransferred represents a OwnershipTransferred event raised by the Usdr contract.
type UsdrOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Usdr *UsdrFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UsdrOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Usdr.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UsdrOwnershipTransferredIterator{contract: _Usdr.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Usdr *UsdrFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UsdrOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Usdr.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsdrOwnershipTransferred)
				if err := _Usdr.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Usdr *UsdrFilterer) ParseOwnershipTransferred(log types.Log) (*UsdrOwnershipTransferred, error) {
	event := new(UsdrOwnershipTransferred)
	if err := _Usdr.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
