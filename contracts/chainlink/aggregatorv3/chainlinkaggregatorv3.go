// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package chainlinkaggregatorv3

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

// Chainlinkaggregatorv3MetaData contains all meta data concerning the Chainlinkaggregatorv3 contract.
var Chainlinkaggregatorv3MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"decimals\",\"stateMutability\":\"view\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}]},{\"type\":\"function\",\"name\":\"latestRoundData\",\"stateMutability\":\"view\",\"inputs\":[],\"outputs\":[{\"name\":\"roundId\",\"type\":\"uint80\"},{\"name\":\"answer\",\"type\":\"int256\"},{\"name\":\"startedAt\",\"type\":\"uint256\"},{\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"name\":\"answeredInRound\",\"type\":\"uint80\"}]}]",
}

// Chainlinkaggregatorv3ABI is the input ABI used to generate the binding from.
// Deprecated: Use Chainlinkaggregatorv3MetaData.ABI instead.
var Chainlinkaggregatorv3ABI = Chainlinkaggregatorv3MetaData.ABI

// Chainlinkaggregatorv3 is an auto generated Go binding around an Ethereum contract.
type Chainlinkaggregatorv3 struct {
	Chainlinkaggregatorv3Caller     // Read-only binding to the contract
	Chainlinkaggregatorv3Transactor // Write-only binding to the contract
	Chainlinkaggregatorv3Filterer   // Log filterer for contract events
}

// Chainlinkaggregatorv3Caller is an auto generated read-only Go binding around an Ethereum contract.
type Chainlinkaggregatorv3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Chainlinkaggregatorv3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Chainlinkaggregatorv3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Chainlinkaggregatorv3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Chainlinkaggregatorv3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Chainlinkaggregatorv3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Chainlinkaggregatorv3Session struct {
	Contract     *Chainlinkaggregatorv3 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// Chainlinkaggregatorv3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Chainlinkaggregatorv3CallerSession struct {
	Contract *Chainlinkaggregatorv3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// Chainlinkaggregatorv3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Chainlinkaggregatorv3TransactorSession struct {
	Contract     *Chainlinkaggregatorv3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// Chainlinkaggregatorv3Raw is an auto generated low-level Go binding around an Ethereum contract.
type Chainlinkaggregatorv3Raw struct {
	Contract *Chainlinkaggregatorv3 // Generic contract binding to access the raw methods on
}

// Chainlinkaggregatorv3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Chainlinkaggregatorv3CallerRaw struct {
	Contract *Chainlinkaggregatorv3Caller // Generic read-only contract binding to access the raw methods on
}

// Chainlinkaggregatorv3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Chainlinkaggregatorv3TransactorRaw struct {
	Contract *Chainlinkaggregatorv3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewChainlinkaggregatorv3 creates a new instance of Chainlinkaggregatorv3, bound to a specific deployed contract.
func NewChainlinkaggregatorv3(address common.Address, backend bind.ContractBackend) (*Chainlinkaggregatorv3, error) {
	contract, err := bindChainlinkaggregatorv3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Chainlinkaggregatorv3{Chainlinkaggregatorv3Caller: Chainlinkaggregatorv3Caller{contract: contract}, Chainlinkaggregatorv3Transactor: Chainlinkaggregatorv3Transactor{contract: contract}, Chainlinkaggregatorv3Filterer: Chainlinkaggregatorv3Filterer{contract: contract}}, nil
}

// NewChainlinkaggregatorv3Caller creates a new read-only instance of Chainlinkaggregatorv3, bound to a specific deployed contract.
func NewChainlinkaggregatorv3Caller(address common.Address, caller bind.ContractCaller) (*Chainlinkaggregatorv3Caller, error) {
	contract, err := bindChainlinkaggregatorv3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Chainlinkaggregatorv3Caller{contract: contract}, nil
}

// NewChainlinkaggregatorv3Transactor creates a new write-only instance of Chainlinkaggregatorv3, bound to a specific deployed contract.
func NewChainlinkaggregatorv3Transactor(address common.Address, transactor bind.ContractTransactor) (*Chainlinkaggregatorv3Transactor, error) {
	contract, err := bindChainlinkaggregatorv3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Chainlinkaggregatorv3Transactor{contract: contract}, nil
}

// NewChainlinkaggregatorv3Filterer creates a new log filterer instance of Chainlinkaggregatorv3, bound to a specific deployed contract.
func NewChainlinkaggregatorv3Filterer(address common.Address, filterer bind.ContractFilterer) (*Chainlinkaggregatorv3Filterer, error) {
	contract, err := bindChainlinkaggregatorv3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Chainlinkaggregatorv3Filterer{contract: contract}, nil
}

// bindChainlinkaggregatorv3 binds a generic wrapper to an already deployed contract.
func bindChainlinkaggregatorv3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Chainlinkaggregatorv3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Chainlinkaggregatorv3 *Chainlinkaggregatorv3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Chainlinkaggregatorv3.Contract.Chainlinkaggregatorv3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Chainlinkaggregatorv3 *Chainlinkaggregatorv3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Chainlinkaggregatorv3.Contract.Chainlinkaggregatorv3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Chainlinkaggregatorv3 *Chainlinkaggregatorv3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Chainlinkaggregatorv3.Contract.Chainlinkaggregatorv3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Chainlinkaggregatorv3 *Chainlinkaggregatorv3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Chainlinkaggregatorv3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Chainlinkaggregatorv3 *Chainlinkaggregatorv3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Chainlinkaggregatorv3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Chainlinkaggregatorv3 *Chainlinkaggregatorv3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Chainlinkaggregatorv3.Contract.contract.Transact(opts, method, params...)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Chainlinkaggregatorv3 *Chainlinkaggregatorv3Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Chainlinkaggregatorv3.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Chainlinkaggregatorv3 *Chainlinkaggregatorv3Session) Decimals() (uint8, error) {
	return _Chainlinkaggregatorv3.Contract.Decimals(&_Chainlinkaggregatorv3.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Chainlinkaggregatorv3 *Chainlinkaggregatorv3CallerSession) Decimals() (uint8, error) {
	return _Chainlinkaggregatorv3.Contract.Decimals(&_Chainlinkaggregatorv3.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_Chainlinkaggregatorv3 *Chainlinkaggregatorv3Caller) LatestRoundData(opts *bind.CallOpts) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _Chainlinkaggregatorv3.contract.Call(opts, &out, "latestRoundData")

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_Chainlinkaggregatorv3 *Chainlinkaggregatorv3Session) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _Chainlinkaggregatorv3.Contract.LatestRoundData(&_Chainlinkaggregatorv3.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_Chainlinkaggregatorv3 *Chainlinkaggregatorv3CallerSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _Chainlinkaggregatorv3.Contract.LatestRoundData(&_Chainlinkaggregatorv3.CallOpts)
}
