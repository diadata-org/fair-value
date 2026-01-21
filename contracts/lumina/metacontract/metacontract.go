// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package metacontract

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

// MetacontractMetaData contains all meta data concerning the Metacontract contract.
var MetacontractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InvalidTimeOut\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OracleExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OracleNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validValues\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"ThresholdNotMet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TimeoutExceedsLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOracleAddress\",\"type\":\"address\"}],\"name\":\"OracleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"removedOracleAddress\",\"type\":\"address\"}],\"name\":\"OracleRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOracleAddress\",\"type\":\"address\"}],\"name\":\"addOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumOracles\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTimeoutSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getValue\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"oracles\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracleToRemove\",\"type\":\"address\"}],\"name\":\"removeOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"setThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTimeoutSeconds\",\"type\":\"uint256\"}],\"name\":\"setTimeoutSeconds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MetacontractABI is the input ABI used to generate the binding from.
// Deprecated: Use MetacontractMetaData.ABI instead.
var MetacontractABI = MetacontractMetaData.ABI

// Metacontract is an auto generated Go binding around an Ethereum contract.
type Metacontract struct {
	MetacontractCaller     // Read-only binding to the contract
	MetacontractTransactor // Write-only binding to the contract
	MetacontractFilterer   // Log filterer for contract events
}

// MetacontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type MetacontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetacontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MetacontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetacontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MetacontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetacontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MetacontractSession struct {
	Contract     *Metacontract     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MetacontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MetacontractCallerSession struct {
	Contract *MetacontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// MetacontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MetacontractTransactorSession struct {
	Contract     *MetacontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MetacontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type MetacontractRaw struct {
	Contract *Metacontract // Generic contract binding to access the raw methods on
}

// MetacontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MetacontractCallerRaw struct {
	Contract *MetacontractCaller // Generic read-only contract binding to access the raw methods on
}

// MetacontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MetacontractTransactorRaw struct {
	Contract *MetacontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMetacontract creates a new instance of Metacontract, bound to a specific deployed contract.
func NewMetacontract(address common.Address, backend bind.ContractBackend) (*Metacontract, error) {
	contract, err := bindMetacontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Metacontract{MetacontractCaller: MetacontractCaller{contract: contract}, MetacontractTransactor: MetacontractTransactor{contract: contract}, MetacontractFilterer: MetacontractFilterer{contract: contract}}, nil
}

// NewMetacontractCaller creates a new read-only instance of Metacontract, bound to a specific deployed contract.
func NewMetacontractCaller(address common.Address, caller bind.ContractCaller) (*MetacontractCaller, error) {
	contract, err := bindMetacontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MetacontractCaller{contract: contract}, nil
}

// NewMetacontractTransactor creates a new write-only instance of Metacontract, bound to a specific deployed contract.
func NewMetacontractTransactor(address common.Address, transactor bind.ContractTransactor) (*MetacontractTransactor, error) {
	contract, err := bindMetacontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MetacontractTransactor{contract: contract}, nil
}

// NewMetacontractFilterer creates a new log filterer instance of Metacontract, bound to a specific deployed contract.
func NewMetacontractFilterer(address common.Address, filterer bind.ContractFilterer) (*MetacontractFilterer, error) {
	contract, err := bindMetacontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MetacontractFilterer{contract: contract}, nil
}

// bindMetacontract binds a generic wrapper to an already deployed contract.
func bindMetacontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MetacontractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Metacontract *MetacontractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Metacontract.Contract.MetacontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Metacontract *MetacontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Metacontract.Contract.MetacontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Metacontract *MetacontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Metacontract.Contract.MetacontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Metacontract *MetacontractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Metacontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Metacontract *MetacontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Metacontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Metacontract *MetacontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Metacontract.Contract.contract.Transact(opts, method, params...)
}

// GetNumOracles is a free data retrieval call binding the contract method 0x30c344f7.
//
// Solidity: function getNumOracles() view returns(uint256)
func (_Metacontract *MetacontractCaller) GetNumOracles(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Metacontract.contract.Call(opts, &out, "getNumOracles")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumOracles is a free data retrieval call binding the contract method 0x30c344f7.
//
// Solidity: function getNumOracles() view returns(uint256)
func (_Metacontract *MetacontractSession) GetNumOracles() (*big.Int, error) {
	return _Metacontract.Contract.GetNumOracles(&_Metacontract.CallOpts)
}

// GetNumOracles is a free data retrieval call binding the contract method 0x30c344f7.
//
// Solidity: function getNumOracles() view returns(uint256)
func (_Metacontract *MetacontractCallerSession) GetNumOracles() (*big.Int, error) {
	return _Metacontract.Contract.GetNumOracles(&_Metacontract.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_Metacontract *MetacontractCaller) GetThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Metacontract.contract.Call(opts, &out, "getThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_Metacontract *MetacontractSession) GetThreshold() (*big.Int, error) {
	return _Metacontract.Contract.GetThreshold(&_Metacontract.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_Metacontract *MetacontractCallerSession) GetThreshold() (*big.Int, error) {
	return _Metacontract.Contract.GetThreshold(&_Metacontract.CallOpts)
}

// GetTimeoutSeconds is a free data retrieval call binding the contract method 0x8d4fd753.
//
// Solidity: function getTimeoutSeconds() view returns(uint256)
func (_Metacontract *MetacontractCaller) GetTimeoutSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Metacontract.contract.Call(opts, &out, "getTimeoutSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimeoutSeconds is a free data retrieval call binding the contract method 0x8d4fd753.
//
// Solidity: function getTimeoutSeconds() view returns(uint256)
func (_Metacontract *MetacontractSession) GetTimeoutSeconds() (*big.Int, error) {
	return _Metacontract.Contract.GetTimeoutSeconds(&_Metacontract.CallOpts)
}

// GetTimeoutSeconds is a free data retrieval call binding the contract method 0x8d4fd753.
//
// Solidity: function getTimeoutSeconds() view returns(uint256)
func (_Metacontract *MetacontractCallerSession) GetTimeoutSeconds() (*big.Int, error) {
	return _Metacontract.Contract.GetTimeoutSeconds(&_Metacontract.CallOpts)
}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint128, uint128)
func (_Metacontract *MetacontractCaller) GetValue(opts *bind.CallOpts, key string) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Metacontract.contract.Call(opts, &out, "getValue", key)

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
func (_Metacontract *MetacontractSession) GetValue(key string) (*big.Int, *big.Int, error) {
	return _Metacontract.Contract.GetValue(&_Metacontract.CallOpts, key)
}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint128, uint128)
func (_Metacontract *MetacontractCallerSession) GetValue(key string) (*big.Int, *big.Int, error) {
	return _Metacontract.Contract.GetValue(&_Metacontract.CallOpts, key)
}

// Oracles is a free data retrieval call binding the contract method 0x5b69a7d8.
//
// Solidity: function oracles(uint256 ) view returns(address)
func (_Metacontract *MetacontractCaller) Oracles(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Metacontract.contract.Call(opts, &out, "oracles", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracles is a free data retrieval call binding the contract method 0x5b69a7d8.
//
// Solidity: function oracles(uint256 ) view returns(address)
func (_Metacontract *MetacontractSession) Oracles(arg0 *big.Int) (common.Address, error) {
	return _Metacontract.Contract.Oracles(&_Metacontract.CallOpts, arg0)
}

// Oracles is a free data retrieval call binding the contract method 0x5b69a7d8.
//
// Solidity: function oracles(uint256 ) view returns(address)
func (_Metacontract *MetacontractCallerSession) Oracles(arg0 *big.Int) (common.Address, error) {
	return _Metacontract.Contract.Oracles(&_Metacontract.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Metacontract *MetacontractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Metacontract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Metacontract *MetacontractSession) Owner() (common.Address, error) {
	return _Metacontract.Contract.Owner(&_Metacontract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Metacontract *MetacontractCallerSession) Owner() (common.Address, error) {
	return _Metacontract.Contract.Owner(&_Metacontract.CallOpts)
}

// AddOracle is a paid mutator transaction binding the contract method 0xdf5dd1a5.
//
// Solidity: function addOracle(address newOracleAddress) returns()
func (_Metacontract *MetacontractTransactor) AddOracle(opts *bind.TransactOpts, newOracleAddress common.Address) (*types.Transaction, error) {
	return _Metacontract.contract.Transact(opts, "addOracle", newOracleAddress)
}

// AddOracle is a paid mutator transaction binding the contract method 0xdf5dd1a5.
//
// Solidity: function addOracle(address newOracleAddress) returns()
func (_Metacontract *MetacontractSession) AddOracle(newOracleAddress common.Address) (*types.Transaction, error) {
	return _Metacontract.Contract.AddOracle(&_Metacontract.TransactOpts, newOracleAddress)
}

// AddOracle is a paid mutator transaction binding the contract method 0xdf5dd1a5.
//
// Solidity: function addOracle(address newOracleAddress) returns()
func (_Metacontract *MetacontractTransactorSession) AddOracle(newOracleAddress common.Address) (*types.Transaction, error) {
	return _Metacontract.Contract.AddOracle(&_Metacontract.TransactOpts, newOracleAddress)
}

// RemoveOracle is a paid mutator transaction binding the contract method 0xfdc85fc4.
//
// Solidity: function removeOracle(address oracleToRemove) returns()
func (_Metacontract *MetacontractTransactor) RemoveOracle(opts *bind.TransactOpts, oracleToRemove common.Address) (*types.Transaction, error) {
	return _Metacontract.contract.Transact(opts, "removeOracle", oracleToRemove)
}

// RemoveOracle is a paid mutator transaction binding the contract method 0xfdc85fc4.
//
// Solidity: function removeOracle(address oracleToRemove) returns()
func (_Metacontract *MetacontractSession) RemoveOracle(oracleToRemove common.Address) (*types.Transaction, error) {
	return _Metacontract.Contract.RemoveOracle(&_Metacontract.TransactOpts, oracleToRemove)
}

// RemoveOracle is a paid mutator transaction binding the contract method 0xfdc85fc4.
//
// Solidity: function removeOracle(address oracleToRemove) returns()
func (_Metacontract *MetacontractTransactorSession) RemoveOracle(oracleToRemove common.Address) (*types.Transaction, error) {
	return _Metacontract.Contract.RemoveOracle(&_Metacontract.TransactOpts, oracleToRemove)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Metacontract *MetacontractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Metacontract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Metacontract *MetacontractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Metacontract.Contract.RenounceOwnership(&_Metacontract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Metacontract *MetacontractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Metacontract.Contract.RenounceOwnership(&_Metacontract.TransactOpts)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x960bfe04.
//
// Solidity: function setThreshold(uint256 newThreshold) returns()
func (_Metacontract *MetacontractTransactor) SetThreshold(opts *bind.TransactOpts, newThreshold *big.Int) (*types.Transaction, error) {
	return _Metacontract.contract.Transact(opts, "setThreshold", newThreshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x960bfe04.
//
// Solidity: function setThreshold(uint256 newThreshold) returns()
func (_Metacontract *MetacontractSession) SetThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _Metacontract.Contract.SetThreshold(&_Metacontract.TransactOpts, newThreshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x960bfe04.
//
// Solidity: function setThreshold(uint256 newThreshold) returns()
func (_Metacontract *MetacontractTransactorSession) SetThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _Metacontract.Contract.SetThreshold(&_Metacontract.TransactOpts, newThreshold)
}

// SetTimeoutSeconds is a paid mutator transaction binding the contract method 0x36f945cd.
//
// Solidity: function setTimeoutSeconds(uint256 newTimeoutSeconds) returns()
func (_Metacontract *MetacontractTransactor) SetTimeoutSeconds(opts *bind.TransactOpts, newTimeoutSeconds *big.Int) (*types.Transaction, error) {
	return _Metacontract.contract.Transact(opts, "setTimeoutSeconds", newTimeoutSeconds)
}

// SetTimeoutSeconds is a paid mutator transaction binding the contract method 0x36f945cd.
//
// Solidity: function setTimeoutSeconds(uint256 newTimeoutSeconds) returns()
func (_Metacontract *MetacontractSession) SetTimeoutSeconds(newTimeoutSeconds *big.Int) (*types.Transaction, error) {
	return _Metacontract.Contract.SetTimeoutSeconds(&_Metacontract.TransactOpts, newTimeoutSeconds)
}

// SetTimeoutSeconds is a paid mutator transaction binding the contract method 0x36f945cd.
//
// Solidity: function setTimeoutSeconds(uint256 newTimeoutSeconds) returns()
func (_Metacontract *MetacontractTransactorSession) SetTimeoutSeconds(newTimeoutSeconds *big.Int) (*types.Transaction, error) {
	return _Metacontract.Contract.SetTimeoutSeconds(&_Metacontract.TransactOpts, newTimeoutSeconds)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Metacontract *MetacontractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Metacontract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Metacontract *MetacontractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Metacontract.Contract.TransferOwnership(&_Metacontract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Metacontract *MetacontractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Metacontract.Contract.TransferOwnership(&_Metacontract.TransactOpts, newOwner)
}

// MetacontractOracleAddedIterator is returned from FilterOracleAdded and is used to iterate over the raw logs and unpacked data for OracleAdded events raised by the Metacontract contract.
type MetacontractOracleAddedIterator struct {
	Event *MetacontractOracleAdded // Event containing the contract specifics and raw log

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
func (it *MetacontractOracleAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetacontractOracleAdded)
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
		it.Event = new(MetacontractOracleAdded)
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
func (it *MetacontractOracleAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetacontractOracleAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetacontractOracleAdded represents a OracleAdded event raised by the Metacontract contract.
type MetacontractOracleAdded struct {
	NewOracleAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOracleAdded is a free log retrieval operation binding the contract event 0x0047706786c922d17b39285dc59d696bafea72c0b003d3841ae1202076f4c2e4.
//
// Solidity: event OracleAdded(address newOracleAddress)
func (_Metacontract *MetacontractFilterer) FilterOracleAdded(opts *bind.FilterOpts) (*MetacontractOracleAddedIterator, error) {

	logs, sub, err := _Metacontract.contract.FilterLogs(opts, "OracleAdded")
	if err != nil {
		return nil, err
	}
	return &MetacontractOracleAddedIterator{contract: _Metacontract.contract, event: "OracleAdded", logs: logs, sub: sub}, nil
}

// WatchOracleAdded is a free log subscription operation binding the contract event 0x0047706786c922d17b39285dc59d696bafea72c0b003d3841ae1202076f4c2e4.
//
// Solidity: event OracleAdded(address newOracleAddress)
func (_Metacontract *MetacontractFilterer) WatchOracleAdded(opts *bind.WatchOpts, sink chan<- *MetacontractOracleAdded) (event.Subscription, error) {

	logs, sub, err := _Metacontract.contract.WatchLogs(opts, "OracleAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetacontractOracleAdded)
				if err := _Metacontract.contract.UnpackLog(event, "OracleAdded", log); err != nil {
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

// ParseOracleAdded is a log parse operation binding the contract event 0x0047706786c922d17b39285dc59d696bafea72c0b003d3841ae1202076f4c2e4.
//
// Solidity: event OracleAdded(address newOracleAddress)
func (_Metacontract *MetacontractFilterer) ParseOracleAdded(log types.Log) (*MetacontractOracleAdded, error) {
	event := new(MetacontractOracleAdded)
	if err := _Metacontract.contract.UnpackLog(event, "OracleAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetacontractOracleRemovedIterator is returned from FilterOracleRemoved and is used to iterate over the raw logs and unpacked data for OracleRemoved events raised by the Metacontract contract.
type MetacontractOracleRemovedIterator struct {
	Event *MetacontractOracleRemoved // Event containing the contract specifics and raw log

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
func (it *MetacontractOracleRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetacontractOracleRemoved)
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
		it.Event = new(MetacontractOracleRemoved)
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
func (it *MetacontractOracleRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetacontractOracleRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetacontractOracleRemoved represents a OracleRemoved event raised by the Metacontract contract.
type MetacontractOracleRemoved struct {
	RemovedOracleAddress common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterOracleRemoved is a free log retrieval operation binding the contract event 0x9c8e7d83025bef8a04c664b2f753f64b8814bdb7e27291d7e50935f18cc3c712.
//
// Solidity: event OracleRemoved(address removedOracleAddress)
func (_Metacontract *MetacontractFilterer) FilterOracleRemoved(opts *bind.FilterOpts) (*MetacontractOracleRemovedIterator, error) {

	logs, sub, err := _Metacontract.contract.FilterLogs(opts, "OracleRemoved")
	if err != nil {
		return nil, err
	}
	return &MetacontractOracleRemovedIterator{contract: _Metacontract.contract, event: "OracleRemoved", logs: logs, sub: sub}, nil
}

// WatchOracleRemoved is a free log subscription operation binding the contract event 0x9c8e7d83025bef8a04c664b2f753f64b8814bdb7e27291d7e50935f18cc3c712.
//
// Solidity: event OracleRemoved(address removedOracleAddress)
func (_Metacontract *MetacontractFilterer) WatchOracleRemoved(opts *bind.WatchOpts, sink chan<- *MetacontractOracleRemoved) (event.Subscription, error) {

	logs, sub, err := _Metacontract.contract.WatchLogs(opts, "OracleRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetacontractOracleRemoved)
				if err := _Metacontract.contract.UnpackLog(event, "OracleRemoved", log); err != nil {
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

// ParseOracleRemoved is a log parse operation binding the contract event 0x9c8e7d83025bef8a04c664b2f753f64b8814bdb7e27291d7e50935f18cc3c712.
//
// Solidity: event OracleRemoved(address removedOracleAddress)
func (_Metacontract *MetacontractFilterer) ParseOracleRemoved(log types.Log) (*MetacontractOracleRemoved, error) {
	event := new(MetacontractOracleRemoved)
	if err := _Metacontract.contract.UnpackLog(event, "OracleRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetacontractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Metacontract contract.
type MetacontractOwnershipTransferredIterator struct {
	Event *MetacontractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MetacontractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetacontractOwnershipTransferred)
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
		it.Event = new(MetacontractOwnershipTransferred)
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
func (it *MetacontractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetacontractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetacontractOwnershipTransferred represents a OwnershipTransferred event raised by the Metacontract contract.
type MetacontractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Metacontract *MetacontractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MetacontractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Metacontract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MetacontractOwnershipTransferredIterator{contract: _Metacontract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Metacontract *MetacontractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MetacontractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Metacontract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetacontractOwnershipTransferred)
				if err := _Metacontract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Metacontract *MetacontractFilterer) ParseOwnershipTransferred(log types.Log) (*MetacontractOwnershipTransferred, error) {
	event := new(MetacontractOwnershipTransferred)
	if err := _Metacontract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
