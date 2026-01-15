// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hemibtc

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

// HemibtcMetaData contains all meta data concerning the Hemibtc contract.
var HemibtcMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authorizedMinter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"invalidMinter\",\"type\":\"address\"}],\"name\":\"InvalidAuthorizedMinter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"invalidMinter\",\"type\":\"address\"}],\"name\":\"MinterAddressNotAuthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"burnFrom\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnBTC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mintBTC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// HemibtcABI is the input ABI used to generate the binding from.
// Deprecated: Use HemibtcMetaData.ABI instead.
var HemibtcABI = HemibtcMetaData.ABI

// Hemibtc is an auto generated Go binding around an Ethereum contract.
type Hemibtc struct {
	HemibtcCaller     // Read-only binding to the contract
	HemibtcTransactor // Write-only binding to the contract
	HemibtcFilterer   // Log filterer for contract events
}

// HemibtcCaller is an auto generated read-only Go binding around an Ethereum contract.
type HemibtcCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HemibtcTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HemibtcTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HemibtcFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HemibtcFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HemibtcSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HemibtcSession struct {
	Contract     *Hemibtc          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HemibtcCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HemibtcCallerSession struct {
	Contract *HemibtcCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// HemibtcTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HemibtcTransactorSession struct {
	Contract     *HemibtcTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// HemibtcRaw is an auto generated low-level Go binding around an Ethereum contract.
type HemibtcRaw struct {
	Contract *Hemibtc // Generic contract binding to access the raw methods on
}

// HemibtcCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HemibtcCallerRaw struct {
	Contract *HemibtcCaller // Generic read-only contract binding to access the raw methods on
}

// HemibtcTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HemibtcTransactorRaw struct {
	Contract *HemibtcTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHemibtc creates a new instance of Hemibtc, bound to a specific deployed contract.
func NewHemibtc(address common.Address, backend bind.ContractBackend) (*Hemibtc, error) {
	contract, err := bindHemibtc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hemibtc{HemibtcCaller: HemibtcCaller{contract: contract}, HemibtcTransactor: HemibtcTransactor{contract: contract}, HemibtcFilterer: HemibtcFilterer{contract: contract}}, nil
}

// NewHemibtcCaller creates a new read-only instance of Hemibtc, bound to a specific deployed contract.
func NewHemibtcCaller(address common.Address, caller bind.ContractCaller) (*HemibtcCaller, error) {
	contract, err := bindHemibtc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HemibtcCaller{contract: contract}, nil
}

// NewHemibtcTransactor creates a new write-only instance of Hemibtc, bound to a specific deployed contract.
func NewHemibtcTransactor(address common.Address, transactor bind.ContractTransactor) (*HemibtcTransactor, error) {
	contract, err := bindHemibtc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HemibtcTransactor{contract: contract}, nil
}

// NewHemibtcFilterer creates a new log filterer instance of Hemibtc, bound to a specific deployed contract.
func NewHemibtcFilterer(address common.Address, filterer bind.ContractFilterer) (*HemibtcFilterer, error) {
	contract, err := bindHemibtc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HemibtcFilterer{contract: contract}, nil
}

// bindHemibtc binds a generic wrapper to an already deployed contract.
func bindHemibtc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HemibtcMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hemibtc *HemibtcRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hemibtc.Contract.HemibtcCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hemibtc *HemibtcRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hemibtc.Contract.HemibtcTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hemibtc *HemibtcRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hemibtc.Contract.HemibtcTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hemibtc *HemibtcCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hemibtc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hemibtc *HemibtcTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hemibtc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hemibtc *HemibtcTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hemibtc.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Hemibtc *HemibtcCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Hemibtc.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Hemibtc *HemibtcSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Hemibtc.Contract.Allowance(&_Hemibtc.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Hemibtc *HemibtcCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Hemibtc.Contract.Allowance(&_Hemibtc.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Hemibtc *HemibtcCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Hemibtc.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Hemibtc *HemibtcSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Hemibtc.Contract.BalanceOf(&_Hemibtc.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Hemibtc *HemibtcCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Hemibtc.Contract.BalanceOf(&_Hemibtc.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Hemibtc *HemibtcCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Hemibtc.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Hemibtc *HemibtcSession) Decimals() (uint8, error) {
	return _Hemibtc.Contract.Decimals(&_Hemibtc.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Hemibtc *HemibtcCallerSession) Decimals() (uint8, error) {
	return _Hemibtc.Contract.Decimals(&_Hemibtc.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Hemibtc *HemibtcCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hemibtc.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Hemibtc *HemibtcSession) Minter() (common.Address, error) {
	return _Hemibtc.Contract.Minter(&_Hemibtc.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Hemibtc *HemibtcCallerSession) Minter() (common.Address, error) {
	return _Hemibtc.Contract.Minter(&_Hemibtc.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Hemibtc *HemibtcCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Hemibtc.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Hemibtc *HemibtcSession) Name() (string, error) {
	return _Hemibtc.Contract.Name(&_Hemibtc.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Hemibtc *HemibtcCallerSession) Name() (string, error) {
	return _Hemibtc.Contract.Name(&_Hemibtc.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Hemibtc *HemibtcCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Hemibtc.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Hemibtc *HemibtcSession) Symbol() (string, error) {
	return _Hemibtc.Contract.Symbol(&_Hemibtc.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Hemibtc *HemibtcCallerSession) Symbol() (string, error) {
	return _Hemibtc.Contract.Symbol(&_Hemibtc.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Hemibtc *HemibtcCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hemibtc.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Hemibtc *HemibtcSession) TotalSupply() (*big.Int, error) {
	return _Hemibtc.Contract.TotalSupply(&_Hemibtc.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Hemibtc *HemibtcCallerSession) TotalSupply() (*big.Int, error) {
	return _Hemibtc.Contract.TotalSupply(&_Hemibtc.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Hemibtc *HemibtcTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Hemibtc.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Hemibtc *HemibtcSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Hemibtc.Contract.Approve(&_Hemibtc.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Hemibtc *HemibtcTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Hemibtc.Contract.Approve(&_Hemibtc.TransactOpts, spender, value)
}

// BurnBTC is a paid mutator transaction binding the contract method 0xfb06c472.
//
// Solidity: function burnBTC(address burnFrom, uint256 amount) returns()
func (_Hemibtc *HemibtcTransactor) BurnBTC(opts *bind.TransactOpts, burnFrom common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Hemibtc.contract.Transact(opts, "burnBTC", burnFrom, amount)
}

// BurnBTC is a paid mutator transaction binding the contract method 0xfb06c472.
//
// Solidity: function burnBTC(address burnFrom, uint256 amount) returns()
func (_Hemibtc *HemibtcSession) BurnBTC(burnFrom common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Hemibtc.Contract.BurnBTC(&_Hemibtc.TransactOpts, burnFrom, amount)
}

// BurnBTC is a paid mutator transaction binding the contract method 0xfb06c472.
//
// Solidity: function burnBTC(address burnFrom, uint256 amount) returns()
func (_Hemibtc *HemibtcTransactorSession) BurnBTC(burnFrom common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Hemibtc.Contract.BurnBTC(&_Hemibtc.TransactOpts, burnFrom, amount)
}

// MintBTC is a paid mutator transaction binding the contract method 0x8394c837.
//
// Solidity: function mintBTC(address recipient, uint256 amount) returns()
func (_Hemibtc *HemibtcTransactor) MintBTC(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Hemibtc.contract.Transact(opts, "mintBTC", recipient, amount)
}

// MintBTC is a paid mutator transaction binding the contract method 0x8394c837.
//
// Solidity: function mintBTC(address recipient, uint256 amount) returns()
func (_Hemibtc *HemibtcSession) MintBTC(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Hemibtc.Contract.MintBTC(&_Hemibtc.TransactOpts, recipient, amount)
}

// MintBTC is a paid mutator transaction binding the contract method 0x8394c837.
//
// Solidity: function mintBTC(address recipient, uint256 amount) returns()
func (_Hemibtc *HemibtcTransactorSession) MintBTC(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Hemibtc.Contract.MintBTC(&_Hemibtc.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Hemibtc *HemibtcTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Hemibtc.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Hemibtc *HemibtcSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Hemibtc.Contract.Transfer(&_Hemibtc.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Hemibtc *HemibtcTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Hemibtc.Contract.Transfer(&_Hemibtc.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Hemibtc *HemibtcTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Hemibtc.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Hemibtc *HemibtcSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Hemibtc.Contract.TransferFrom(&_Hemibtc.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Hemibtc *HemibtcTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Hemibtc.Contract.TransferFrom(&_Hemibtc.TransactOpts, from, to, value)
}

// HemibtcApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Hemibtc contract.
type HemibtcApprovalIterator struct {
	Event *HemibtcApproval // Event containing the contract specifics and raw log

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
func (it *HemibtcApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HemibtcApproval)
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
		it.Event = new(HemibtcApproval)
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
func (it *HemibtcApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HemibtcApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HemibtcApproval represents a Approval event raised by the Hemibtc contract.
type HemibtcApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Hemibtc *HemibtcFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*HemibtcApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Hemibtc.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &HemibtcApprovalIterator{contract: _Hemibtc.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Hemibtc *HemibtcFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *HemibtcApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Hemibtc.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HemibtcApproval)
				if err := _Hemibtc.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Hemibtc *HemibtcFilterer) ParseApproval(log types.Log) (*HemibtcApproval, error) {
	event := new(HemibtcApproval)
	if err := _Hemibtc.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HemibtcTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Hemibtc contract.
type HemibtcTransferIterator struct {
	Event *HemibtcTransfer // Event containing the contract specifics and raw log

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
func (it *HemibtcTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HemibtcTransfer)
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
		it.Event = new(HemibtcTransfer)
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
func (it *HemibtcTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HemibtcTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HemibtcTransfer represents a Transfer event raised by the Hemibtc contract.
type HemibtcTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Hemibtc *HemibtcFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*HemibtcTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Hemibtc.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &HemibtcTransferIterator{contract: _Hemibtc.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Hemibtc *HemibtcFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *HemibtcTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Hemibtc.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HemibtcTransfer)
				if err := _Hemibtc.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Hemibtc *HemibtcFilterer) ParseTransfer(log types.Log) (*HemibtcTransfer, error) {
	event := new(HemibtcTransfer)
	if err := _Hemibtc.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
