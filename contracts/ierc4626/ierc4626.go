// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ierc4626

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

// Ierc4626MetaData contains all meta data concerning the Ierc4626 contract.
var Ierc4626MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"asset\",\"stateMutability\":\"view\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"type\":\"function\",\"name\":\"totalAssets\",\"stateMutability\":\"view\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"type\":\"function\",\"name\":\"convertToShares\",\"stateMutability\":\"view\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\"}]},{\"type\":\"function\",\"name\":\"convertToAssets\",\"stateMutability\":\"view\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"assets\",\"type\":\"uint256\"}]},{\"type\":\"function\",\"name\":\"maxDeposit\",\"stateMutability\":\"view\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"maxAssets\",\"type\":\"uint256\"}]},{\"type\":\"function\",\"name\":\"previewDeposit\",\"stateMutability\":\"view\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\"}]},{\"type\":\"function\",\"name\":\"deposit\",\"stateMutability\":\"nonpayable\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\"}]},{\"type\":\"function\",\"name\":\"maxMint\",\"stateMutability\":\"view\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"maxShares\",\"type\":\"uint256\"}]},{\"type\":\"function\",\"name\":\"previewMint\",\"stateMutability\":\"view\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"assets\",\"type\":\"uint256\"}]},{\"type\":\"function\",\"name\":\"mint\",\"stateMutability\":\"nonpayable\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"assets\",\"type\":\"uint256\"}]},{\"type\":\"function\",\"name\":\"maxWithdraw\",\"stateMutability\":\"view\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"maxAssets\",\"type\":\"uint256\"}]},{\"type\":\"function\",\"name\":\"previewWithdraw\",\"stateMutability\":\"view\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\"}]},{\"type\":\"function\",\"name\":\"withdraw\",\"stateMutability\":\"nonpayable\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"owner\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\"}]},{\"type\":\"function\",\"name\":\"maxRedeem\",\"stateMutability\":\"view\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"maxShares\",\"type\":\"uint256\"}]},{\"type\":\"function\",\"name\":\"previewRedeem\",\"stateMutability\":\"view\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"assets\",\"type\":\"uint256\"}]},{\"type\":\"function\",\"name\":\"redeem\",\"stateMutability\":\"nonpayable\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"owner\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"assets\",\"type\":\"uint256\"}]},{\"type\":\"event\",\"name\":\"Deposit\",\"anonymous\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true},{\"name\":\"assets\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false}]},{\"type\":\"event\",\"name\":\"Withdraw\",\"anonymous\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true},{\"name\":\"assets\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false}]},{\"type\":\"function\",\"name\":\"balanceOf\",\"stateMutability\":\"view\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"type\":\"function\",\"name\":\"totalSupply\",\"stateMutability\":\"view\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"type\":\"function\",\"name\":\"decimals\",\"stateMutability\":\"view\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}]},{\"type\":\"function\",\"name\":\"symbol\",\"stateMutability\":\"view\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"type\":\"function\",\"name\":\"name\",\"stateMutability\":\"view\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]}]",
}

// Ierc4626ABI is the input ABI used to generate the binding from.
// Deprecated: Use Ierc4626MetaData.ABI instead.
var Ierc4626ABI = Ierc4626MetaData.ABI

// Ierc4626 is an auto generated Go binding around an Ethereum contract.
type Ierc4626 struct {
	Ierc4626Caller     // Read-only binding to the contract
	Ierc4626Transactor // Write-only binding to the contract
	Ierc4626Filterer   // Log filterer for contract events
}

// Ierc4626Caller is an auto generated read-only Go binding around an Ethereum contract.
type Ierc4626Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ierc4626Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Ierc4626Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ierc4626Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Ierc4626Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ierc4626Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Ierc4626Session struct {
	Contract     *Ierc4626         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Ierc4626CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Ierc4626CallerSession struct {
	Contract *Ierc4626Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// Ierc4626TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Ierc4626TransactorSession struct {
	Contract     *Ierc4626Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// Ierc4626Raw is an auto generated low-level Go binding around an Ethereum contract.
type Ierc4626Raw struct {
	Contract *Ierc4626 // Generic contract binding to access the raw methods on
}

// Ierc4626CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Ierc4626CallerRaw struct {
	Contract *Ierc4626Caller // Generic read-only contract binding to access the raw methods on
}

// Ierc4626TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Ierc4626TransactorRaw struct {
	Contract *Ierc4626Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIerc4626 creates a new instance of Ierc4626, bound to a specific deployed contract.
func NewIerc4626(address common.Address, backend bind.ContractBackend) (*Ierc4626, error) {
	contract, err := bindIerc4626(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ierc4626{Ierc4626Caller: Ierc4626Caller{contract: contract}, Ierc4626Transactor: Ierc4626Transactor{contract: contract}, Ierc4626Filterer: Ierc4626Filterer{contract: contract}}, nil
}

// NewIerc4626Caller creates a new read-only instance of Ierc4626, bound to a specific deployed contract.
func NewIerc4626Caller(address common.Address, caller bind.ContractCaller) (*Ierc4626Caller, error) {
	contract, err := bindIerc4626(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Ierc4626Caller{contract: contract}, nil
}

// NewIerc4626Transactor creates a new write-only instance of Ierc4626, bound to a specific deployed contract.
func NewIerc4626Transactor(address common.Address, transactor bind.ContractTransactor) (*Ierc4626Transactor, error) {
	contract, err := bindIerc4626(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Ierc4626Transactor{contract: contract}, nil
}

// NewIerc4626Filterer creates a new log filterer instance of Ierc4626, bound to a specific deployed contract.
func NewIerc4626Filterer(address common.Address, filterer bind.ContractFilterer) (*Ierc4626Filterer, error) {
	contract, err := bindIerc4626(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Ierc4626Filterer{contract: contract}, nil
}

// bindIerc4626 binds a generic wrapper to an already deployed contract.
func bindIerc4626(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Ierc4626MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ierc4626 *Ierc4626Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ierc4626.Contract.Ierc4626Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ierc4626 *Ierc4626Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ierc4626.Contract.Ierc4626Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ierc4626 *Ierc4626Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ierc4626.Contract.Ierc4626Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ierc4626 *Ierc4626CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ierc4626.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ierc4626 *Ierc4626TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ierc4626.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ierc4626 *Ierc4626TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ierc4626.Contract.contract.Transact(opts, method, params...)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_Ierc4626 *Ierc4626Caller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ierc4626.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_Ierc4626 *Ierc4626Session) Asset() (common.Address, error) {
	return _Ierc4626.Contract.Asset(&_Ierc4626.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_Ierc4626 *Ierc4626CallerSession) Asset() (common.Address, error) {
	return _Ierc4626.Contract.Asset(&_Ierc4626.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Ierc4626 *Ierc4626Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ierc4626.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Ierc4626 *Ierc4626Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _Ierc4626.Contract.BalanceOf(&_Ierc4626.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Ierc4626 *Ierc4626CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Ierc4626.Contract.BalanceOf(&_Ierc4626.CallOpts, account)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256 assets)
func (_Ierc4626 *Ierc4626Caller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Ierc4626.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256 assets)
func (_Ierc4626 *Ierc4626Session) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _Ierc4626.Contract.ConvertToAssets(&_Ierc4626.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256 assets)
func (_Ierc4626 *Ierc4626CallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _Ierc4626.Contract.ConvertToAssets(&_Ierc4626.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256 shares)
func (_Ierc4626 *Ierc4626Caller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Ierc4626.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256 shares)
func (_Ierc4626 *Ierc4626Session) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _Ierc4626.Contract.ConvertToShares(&_Ierc4626.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256 shares)
func (_Ierc4626 *Ierc4626CallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _Ierc4626.Contract.ConvertToShares(&_Ierc4626.CallOpts, assets)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Ierc4626 *Ierc4626Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Ierc4626.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Ierc4626 *Ierc4626Session) Decimals() (uint8, error) {
	return _Ierc4626.Contract.Decimals(&_Ierc4626.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Ierc4626 *Ierc4626CallerSession) Decimals() (uint8, error) {
	return _Ierc4626.Contract.Decimals(&_Ierc4626.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address receiver) view returns(uint256 maxAssets)
func (_Ierc4626 *Ierc4626Caller) MaxDeposit(opts *bind.CallOpts, receiver common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ierc4626.contract.Call(opts, &out, "maxDeposit", receiver)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address receiver) view returns(uint256 maxAssets)
func (_Ierc4626 *Ierc4626Session) MaxDeposit(receiver common.Address) (*big.Int, error) {
	return _Ierc4626.Contract.MaxDeposit(&_Ierc4626.CallOpts, receiver)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address receiver) view returns(uint256 maxAssets)
func (_Ierc4626 *Ierc4626CallerSession) MaxDeposit(receiver common.Address) (*big.Int, error) {
	return _Ierc4626.Contract.MaxDeposit(&_Ierc4626.CallOpts, receiver)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address receiver) view returns(uint256 maxShares)
func (_Ierc4626 *Ierc4626Caller) MaxMint(opts *bind.CallOpts, receiver common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ierc4626.contract.Call(opts, &out, "maxMint", receiver)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address receiver) view returns(uint256 maxShares)
func (_Ierc4626 *Ierc4626Session) MaxMint(receiver common.Address) (*big.Int, error) {
	return _Ierc4626.Contract.MaxMint(&_Ierc4626.CallOpts, receiver)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address receiver) view returns(uint256 maxShares)
func (_Ierc4626 *Ierc4626CallerSession) MaxMint(receiver common.Address) (*big.Int, error) {
	return _Ierc4626.Contract.MaxMint(&_Ierc4626.CallOpts, receiver)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256 maxShares)
func (_Ierc4626 *Ierc4626Caller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ierc4626.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256 maxShares)
func (_Ierc4626 *Ierc4626Session) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _Ierc4626.Contract.MaxRedeem(&_Ierc4626.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256 maxShares)
func (_Ierc4626 *Ierc4626CallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _Ierc4626.Contract.MaxRedeem(&_Ierc4626.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256 maxAssets)
func (_Ierc4626 *Ierc4626Caller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ierc4626.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256 maxAssets)
func (_Ierc4626 *Ierc4626Session) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _Ierc4626.Contract.MaxWithdraw(&_Ierc4626.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256 maxAssets)
func (_Ierc4626 *Ierc4626CallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _Ierc4626.Contract.MaxWithdraw(&_Ierc4626.CallOpts, owner)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ierc4626 *Ierc4626Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ierc4626.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ierc4626 *Ierc4626Session) Name() (string, error) {
	return _Ierc4626.Contract.Name(&_Ierc4626.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ierc4626 *Ierc4626CallerSession) Name() (string, error) {
	return _Ierc4626.Contract.Name(&_Ierc4626.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256 shares)
func (_Ierc4626 *Ierc4626Caller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Ierc4626.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256 shares)
func (_Ierc4626 *Ierc4626Session) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _Ierc4626.Contract.PreviewDeposit(&_Ierc4626.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256 shares)
func (_Ierc4626 *Ierc4626CallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _Ierc4626.Contract.PreviewDeposit(&_Ierc4626.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256 assets)
func (_Ierc4626 *Ierc4626Caller) PreviewMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Ierc4626.contract.Call(opts, &out, "previewMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256 assets)
func (_Ierc4626 *Ierc4626Session) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _Ierc4626.Contract.PreviewMint(&_Ierc4626.CallOpts, shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256 assets)
func (_Ierc4626 *Ierc4626CallerSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _Ierc4626.Contract.PreviewMint(&_Ierc4626.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256 assets)
func (_Ierc4626 *Ierc4626Caller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Ierc4626.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256 assets)
func (_Ierc4626 *Ierc4626Session) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _Ierc4626.Contract.PreviewRedeem(&_Ierc4626.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256 assets)
func (_Ierc4626 *Ierc4626CallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _Ierc4626.Contract.PreviewRedeem(&_Ierc4626.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256 shares)
func (_Ierc4626 *Ierc4626Caller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Ierc4626.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256 shares)
func (_Ierc4626 *Ierc4626Session) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _Ierc4626.Contract.PreviewWithdraw(&_Ierc4626.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256 shares)
func (_Ierc4626 *Ierc4626CallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _Ierc4626.Contract.PreviewWithdraw(&_Ierc4626.CallOpts, assets)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ierc4626 *Ierc4626Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ierc4626.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ierc4626 *Ierc4626Session) Symbol() (string, error) {
	return _Ierc4626.Contract.Symbol(&_Ierc4626.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ierc4626 *Ierc4626CallerSession) Symbol() (string, error) {
	return _Ierc4626.Contract.Symbol(&_Ierc4626.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Ierc4626 *Ierc4626Caller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ierc4626.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Ierc4626 *Ierc4626Session) TotalAssets() (*big.Int, error) {
	return _Ierc4626.Contract.TotalAssets(&_Ierc4626.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Ierc4626 *Ierc4626CallerSession) TotalAssets() (*big.Int, error) {
	return _Ierc4626.Contract.TotalAssets(&_Ierc4626.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ierc4626 *Ierc4626Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ierc4626.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ierc4626 *Ierc4626Session) TotalSupply() (*big.Int, error) {
	return _Ierc4626.Contract.TotalSupply(&_Ierc4626.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ierc4626 *Ierc4626CallerSession) TotalSupply() (*big.Int, error) {
	return _Ierc4626.Contract.TotalSupply(&_Ierc4626.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_Ierc4626 *Ierc4626Transactor) Deposit(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Ierc4626.contract.Transact(opts, "deposit", assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_Ierc4626 *Ierc4626Session) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Ierc4626.Contract.Deposit(&_Ierc4626.TransactOpts, assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_Ierc4626 *Ierc4626TransactorSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Ierc4626.Contract.Deposit(&_Ierc4626.TransactOpts, assets, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_Ierc4626 *Ierc4626Transactor) Mint(opts *bind.TransactOpts, shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Ierc4626.contract.Transact(opts, "mint", shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_Ierc4626 *Ierc4626Session) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Ierc4626.Contract.Mint(&_Ierc4626.TransactOpts, shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_Ierc4626 *Ierc4626TransactorSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Ierc4626.Contract.Mint(&_Ierc4626.TransactOpts, shares, receiver)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_Ierc4626 *Ierc4626Transactor) Redeem(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Ierc4626.contract.Transact(opts, "redeem", shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_Ierc4626 *Ierc4626Session) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Ierc4626.Contract.Redeem(&_Ierc4626.TransactOpts, shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_Ierc4626 *Ierc4626TransactorSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Ierc4626.Contract.Redeem(&_Ierc4626.TransactOpts, shares, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_Ierc4626 *Ierc4626Transactor) Withdraw(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Ierc4626.contract.Transact(opts, "withdraw", assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_Ierc4626 *Ierc4626Session) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Ierc4626.Contract.Withdraw(&_Ierc4626.TransactOpts, assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_Ierc4626 *Ierc4626TransactorSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Ierc4626.Contract.Withdraw(&_Ierc4626.TransactOpts, assets, receiver, owner)
}

// Ierc4626DepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Ierc4626 contract.
type Ierc4626DepositIterator struct {
	Event *Ierc4626Deposit // Event containing the contract specifics and raw log

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
func (it *Ierc4626DepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Ierc4626Deposit)
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
		it.Event = new(Ierc4626Deposit)
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
func (it *Ierc4626DepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Ierc4626DepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Ierc4626Deposit represents a Deposit event raised by the Ierc4626 contract.
type Ierc4626Deposit struct {
	Sender common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_Ierc4626 *Ierc4626Filterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*Ierc4626DepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ierc4626.contract.FilterLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &Ierc4626DepositIterator{contract: _Ierc4626.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_Ierc4626 *Ierc4626Filterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *Ierc4626Deposit, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ierc4626.contract.WatchLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Ierc4626Deposit)
				if err := _Ierc4626.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_Ierc4626 *Ierc4626Filterer) ParseDeposit(log types.Log) (*Ierc4626Deposit, error) {
	event := new(Ierc4626Deposit)
	if err := _Ierc4626.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Ierc4626WithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Ierc4626 contract.
type Ierc4626WithdrawIterator struct {
	Event *Ierc4626Withdraw // Event containing the contract specifics and raw log

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
func (it *Ierc4626WithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Ierc4626Withdraw)
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
		it.Event = new(Ierc4626Withdraw)
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
func (it *Ierc4626WithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Ierc4626WithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Ierc4626Withdraw represents a Withdraw event raised by the Ierc4626 contract.
type Ierc4626Withdraw struct {
	Sender   common.Address
	Receiver common.Address
	Owner    common.Address
	Assets   *big.Int
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_Ierc4626 *Ierc4626Filterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, owner []common.Address) (*Ierc4626WithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ierc4626.contract.FilterLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &Ierc4626WithdrawIterator{contract: _Ierc4626.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_Ierc4626 *Ierc4626Filterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *Ierc4626Withdraw, sender []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ierc4626.contract.WatchLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Ierc4626Withdraw)
				if err := _Ierc4626.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_Ierc4626 *Ierc4626Filterer) ParseWithdraw(log types.Log) (*Ierc4626Withdraw, error) {
	event := new(Ierc4626Withdraw)
	if err := _Ierc4626.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
