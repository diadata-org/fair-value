// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package usdpgetters

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

// Collateral is an auto generated low-level Go binding around an user-defined struct.
type Collateral struct {
	IsManaged         uint8
	IsMintLive        uint8
	IsBurnLive        uint8
	Decimals          uint8
	OnlyWhitelisted   uint8
	NormalizedStables *big.Int
	XFeeMint          []uint64
	YFeeMint          []int64
	XFeeBurn          []uint64
	YFeeBurn          []int64
	OracleConfig      []byte
	WhitelistData     []byte
	ManagerData       ManagerStorage
	StablecoinCap     *big.Int
}

// ManagerStorage is an auto generated low-level Go binding around an user-defined struct.
type ManagerStorage struct {
	SubCollaterals []common.Address
	Config         []byte
}

// UsdpgettersMetaData contains all meta data concerning the Usdpgetters contract.
var UsdpgettersMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidChainlinkRate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotCollateral\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"accessManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"getCollateralBurnFees\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"xFeeBurn\",\"type\":\"uint64[]\"},{\"internalType\":\"int64[]\",\"name\":\"yFeeBurn\",\"type\":\"int64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"getCollateralDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"getCollateralInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"isManaged\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"isMintLive\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"isBurnLive\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"onlyWhitelisted\",\"type\":\"uint8\"},{\"internalType\":\"uint216\",\"name\":\"normalizedStables\",\"type\":\"uint216\"},{\"internalType\":\"uint64[]\",\"name\":\"xFeeMint\",\"type\":\"uint64[]\"},{\"internalType\":\"int64[]\",\"name\":\"yFeeMint\",\"type\":\"int64[]\"},{\"internalType\":\"uint64[]\",\"name\":\"xFeeBurn\",\"type\":\"uint64[]\"},{\"internalType\":\"int64[]\",\"name\":\"yFeeBurn\",\"type\":\"int64[]\"},{\"internalType\":\"bytes\",\"name\":\"oracleConfig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"whitelistData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"subCollaterals\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"config\",\"type\":\"bytes\"}],\"internalType\":\"structManagerStorage\",\"name\":\"managerData\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"stablecoinCap\",\"type\":\"uint256\"}],\"internalType\":\"structCollateral\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCollateralList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"getCollateralMintFees\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"xFeeMint\",\"type\":\"uint64[]\"},{\"internalType\":\"int64[]\",\"name\":\"yFeeMint\",\"type\":\"int64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCollateralRatio\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"collatRatio\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"stablecoinsIssued\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"getCollateralWhitelistData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"getIssuedByCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stablecoinsFromCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stablecoinsIssued\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"getManagerData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"getOracle\",\"outputs\":[{\"internalType\":\"enumOracleReadType\",\"name\":\"oracleType\",\"type\":\"uint8\"},{\"internalType\":\"enumOracleReadType\",\"name\":\"targetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"oracleData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"targetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"hyperparameters\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"getOracleValues\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"mint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ratio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redemption\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRedemptionFees\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"xRedemptionCurve\",\"type\":\"uint64[]\"},{\"internalType\":\"int64[]\",\"name\":\"yRedemptionCurve\",\"type\":\"int64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"getStablecoinCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalIssued\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isConsumingScheduledOp\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"enumActionType\",\"name\":\"action\",\"type\":\"uint8\"}],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isTrusted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isTrustedSeller\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"isValidSelector\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"isWhitelistedCollateral\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isWhitelistedForCollateral\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumWhitelistType\",\"name\":\"whitelistType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isWhitelistedForType\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenP\",\"outputs\":[{\"internalType\":\"contractITokenP\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// UsdpgettersABI is the input ABI used to generate the binding from.
// Deprecated: Use UsdpgettersMetaData.ABI instead.
var UsdpgettersABI = UsdpgettersMetaData.ABI

// Usdpgetters is an auto generated Go binding around an Ethereum contract.
type Usdpgetters struct {
	UsdpgettersCaller     // Read-only binding to the contract
	UsdpgettersTransactor // Write-only binding to the contract
	UsdpgettersFilterer   // Log filterer for contract events
}

// UsdpgettersCaller is an auto generated read-only Go binding around an Ethereum contract.
type UsdpgettersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsdpgettersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UsdpgettersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsdpgettersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UsdpgettersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsdpgettersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UsdpgettersSession struct {
	Contract     *Usdpgetters      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UsdpgettersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UsdpgettersCallerSession struct {
	Contract *UsdpgettersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// UsdpgettersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UsdpgettersTransactorSession struct {
	Contract     *UsdpgettersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// UsdpgettersRaw is an auto generated low-level Go binding around an Ethereum contract.
type UsdpgettersRaw struct {
	Contract *Usdpgetters // Generic contract binding to access the raw methods on
}

// UsdpgettersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UsdpgettersCallerRaw struct {
	Contract *UsdpgettersCaller // Generic read-only contract binding to access the raw methods on
}

// UsdpgettersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UsdpgettersTransactorRaw struct {
	Contract *UsdpgettersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUsdpgetters creates a new instance of Usdpgetters, bound to a specific deployed contract.
func NewUsdpgetters(address common.Address, backend bind.ContractBackend) (*Usdpgetters, error) {
	contract, err := bindUsdpgetters(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Usdpgetters{UsdpgettersCaller: UsdpgettersCaller{contract: contract}, UsdpgettersTransactor: UsdpgettersTransactor{contract: contract}, UsdpgettersFilterer: UsdpgettersFilterer{contract: contract}}, nil
}

// NewUsdpgettersCaller creates a new read-only instance of Usdpgetters, bound to a specific deployed contract.
func NewUsdpgettersCaller(address common.Address, caller bind.ContractCaller) (*UsdpgettersCaller, error) {
	contract, err := bindUsdpgetters(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UsdpgettersCaller{contract: contract}, nil
}

// NewUsdpgettersTransactor creates a new write-only instance of Usdpgetters, bound to a specific deployed contract.
func NewUsdpgettersTransactor(address common.Address, transactor bind.ContractTransactor) (*UsdpgettersTransactor, error) {
	contract, err := bindUsdpgetters(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UsdpgettersTransactor{contract: contract}, nil
}

// NewUsdpgettersFilterer creates a new log filterer instance of Usdpgetters, bound to a specific deployed contract.
func NewUsdpgettersFilterer(address common.Address, filterer bind.ContractFilterer) (*UsdpgettersFilterer, error) {
	contract, err := bindUsdpgetters(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UsdpgettersFilterer{contract: contract}, nil
}

// bindUsdpgetters binds a generic wrapper to an already deployed contract.
func bindUsdpgetters(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UsdpgettersMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Usdpgetters *UsdpgettersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Usdpgetters.Contract.UsdpgettersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Usdpgetters *UsdpgettersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Usdpgetters.Contract.UsdpgettersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Usdpgetters *UsdpgettersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Usdpgetters.Contract.UsdpgettersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Usdpgetters *UsdpgettersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Usdpgetters.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Usdpgetters *UsdpgettersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Usdpgetters.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Usdpgetters *UsdpgettersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Usdpgetters.Contract.contract.Transact(opts, method, params...)
}

// AccessManager is a free data retrieval call binding the contract method 0xfdcb6068.
//
// Solidity: function accessManager() view returns(address)
func (_Usdpgetters *UsdpgettersCaller) AccessManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Usdpgetters.contract.Call(opts, &out, "accessManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccessManager is a free data retrieval call binding the contract method 0xfdcb6068.
//
// Solidity: function accessManager() view returns(address)
func (_Usdpgetters *UsdpgettersSession) AccessManager() (common.Address, error) {
	return _Usdpgetters.Contract.AccessManager(&_Usdpgetters.CallOpts)
}

// AccessManager is a free data retrieval call binding the contract method 0xfdcb6068.
//
// Solidity: function accessManager() view returns(address)
func (_Usdpgetters *UsdpgettersCallerSession) AccessManager() (common.Address, error) {
	return _Usdpgetters.Contract.AccessManager(&_Usdpgetters.CallOpts)
}

// GetCollateralBurnFees is a free data retrieval call binding the contract method 0x847da7be.
//
// Solidity: function getCollateralBurnFees(address collateral) view returns(uint64[] xFeeBurn, int64[] yFeeBurn)
func (_Usdpgetters *UsdpgettersCaller) GetCollateralBurnFees(opts *bind.CallOpts, collateral common.Address) (struct {
	XFeeBurn []uint64
	YFeeBurn []int64
}, error) {
	var out []interface{}
	err := _Usdpgetters.contract.Call(opts, &out, "getCollateralBurnFees", collateral)

	outstruct := new(struct {
		XFeeBurn []uint64
		YFeeBurn []int64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.XFeeBurn = *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)
	outstruct.YFeeBurn = *abi.ConvertType(out[1], new([]int64)).(*[]int64)

	return *outstruct, err

}

// GetCollateralBurnFees is a free data retrieval call binding the contract method 0x847da7be.
//
// Solidity: function getCollateralBurnFees(address collateral) view returns(uint64[] xFeeBurn, int64[] yFeeBurn)
func (_Usdpgetters *UsdpgettersSession) GetCollateralBurnFees(collateral common.Address) (struct {
	XFeeBurn []uint64
	YFeeBurn []int64
}, error) {
	return _Usdpgetters.Contract.GetCollateralBurnFees(&_Usdpgetters.CallOpts, collateral)
}

// GetCollateralBurnFees is a free data retrieval call binding the contract method 0x847da7be.
//
// Solidity: function getCollateralBurnFees(address collateral) view returns(uint64[] xFeeBurn, int64[] yFeeBurn)
func (_Usdpgetters *UsdpgettersCallerSession) GetCollateralBurnFees(collateral common.Address) (struct {
	XFeeBurn []uint64
	YFeeBurn []int64
}, error) {
	return _Usdpgetters.Contract.GetCollateralBurnFees(&_Usdpgetters.CallOpts, collateral)
}

// GetCollateralDecimals is a free data retrieval call binding the contract method 0xeb7aac5f.
//
// Solidity: function getCollateralDecimals(address collateral) view returns(uint8)
func (_Usdpgetters *UsdpgettersCaller) GetCollateralDecimals(opts *bind.CallOpts, collateral common.Address) (uint8, error) {
	var out []interface{}
	err := _Usdpgetters.contract.Call(opts, &out, "getCollateralDecimals", collateral)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetCollateralDecimals is a free data retrieval call binding the contract method 0xeb7aac5f.
//
// Solidity: function getCollateralDecimals(address collateral) view returns(uint8)
func (_Usdpgetters *UsdpgettersSession) GetCollateralDecimals(collateral common.Address) (uint8, error) {
	return _Usdpgetters.Contract.GetCollateralDecimals(&_Usdpgetters.CallOpts, collateral)
}

// GetCollateralDecimals is a free data retrieval call binding the contract method 0xeb7aac5f.
//
// Solidity: function getCollateralDecimals(address collateral) view returns(uint8)
func (_Usdpgetters *UsdpgettersCallerSession) GetCollateralDecimals(collateral common.Address) (uint8, error) {
	return _Usdpgetters.Contract.GetCollateralDecimals(&_Usdpgetters.CallOpts, collateral)
}

// GetCollateralInfo is a free data retrieval call binding the contract method 0x33352210.
//
// Solidity: function getCollateralInfo(address collateral) view returns((uint8,uint8,uint8,uint8,uint8,uint216,uint64[],int64[],uint64[],int64[],bytes,bytes,(address[],bytes),uint256))
func (_Usdpgetters *UsdpgettersCaller) GetCollateralInfo(opts *bind.CallOpts, collateral common.Address) (Collateral, error) {
	var out []interface{}
	err := _Usdpgetters.contract.Call(opts, &out, "getCollateralInfo", collateral)

	if err != nil {
		return *new(Collateral), err
	}

	out0 := *abi.ConvertType(out[0], new(Collateral)).(*Collateral)

	return out0, err

}

// GetCollateralInfo is a free data retrieval call binding the contract method 0x33352210.
//
// Solidity: function getCollateralInfo(address collateral) view returns((uint8,uint8,uint8,uint8,uint8,uint216,uint64[],int64[],uint64[],int64[],bytes,bytes,(address[],bytes),uint256))
func (_Usdpgetters *UsdpgettersSession) GetCollateralInfo(collateral common.Address) (Collateral, error) {
	return _Usdpgetters.Contract.GetCollateralInfo(&_Usdpgetters.CallOpts, collateral)
}

// GetCollateralInfo is a free data retrieval call binding the contract method 0x33352210.
//
// Solidity: function getCollateralInfo(address collateral) view returns((uint8,uint8,uint8,uint8,uint8,uint216,uint64[],int64[],uint64[],int64[],bytes,bytes,(address[],bytes),uint256))
func (_Usdpgetters *UsdpgettersCallerSession) GetCollateralInfo(collateral common.Address) (Collateral, error) {
	return _Usdpgetters.Contract.GetCollateralInfo(&_Usdpgetters.CallOpts, collateral)
}

// GetCollateralList is a free data retrieval call binding the contract method 0xb7181361.
//
// Solidity: function getCollateralList() view returns(address[])
func (_Usdpgetters *UsdpgettersCaller) GetCollateralList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Usdpgetters.contract.Call(opts, &out, "getCollateralList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCollateralList is a free data retrieval call binding the contract method 0xb7181361.
//
// Solidity: function getCollateralList() view returns(address[])
func (_Usdpgetters *UsdpgettersSession) GetCollateralList() ([]common.Address, error) {
	return _Usdpgetters.Contract.GetCollateralList(&_Usdpgetters.CallOpts)
}

// GetCollateralList is a free data retrieval call binding the contract method 0xb7181361.
//
// Solidity: function getCollateralList() view returns(address[])
func (_Usdpgetters *UsdpgettersCallerSession) GetCollateralList() ([]common.Address, error) {
	return _Usdpgetters.Contract.GetCollateralList(&_Usdpgetters.CallOpts)
}

// GetCollateralMintFees is a free data retrieval call binding the contract method 0xb85780bc.
//
// Solidity: function getCollateralMintFees(address collateral) view returns(uint64[] xFeeMint, int64[] yFeeMint)
func (_Usdpgetters *UsdpgettersCaller) GetCollateralMintFees(opts *bind.CallOpts, collateral common.Address) (struct {
	XFeeMint []uint64
	YFeeMint []int64
}, error) {
	var out []interface{}
	err := _Usdpgetters.contract.Call(opts, &out, "getCollateralMintFees", collateral)

	outstruct := new(struct {
		XFeeMint []uint64
		YFeeMint []int64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.XFeeMint = *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)
	outstruct.YFeeMint = *abi.ConvertType(out[1], new([]int64)).(*[]int64)

	return *outstruct, err

}

// GetCollateralMintFees is a free data retrieval call binding the contract method 0xb85780bc.
//
// Solidity: function getCollateralMintFees(address collateral) view returns(uint64[] xFeeMint, int64[] yFeeMint)
func (_Usdpgetters *UsdpgettersSession) GetCollateralMintFees(collateral common.Address) (struct {
	XFeeMint []uint64
	YFeeMint []int64
}, error) {
	return _Usdpgetters.Contract.GetCollateralMintFees(&_Usdpgetters.CallOpts, collateral)
}

// GetCollateralMintFees is a free data retrieval call binding the contract method 0xb85780bc.
//
// Solidity: function getCollateralMintFees(address collateral) view returns(uint64[] xFeeMint, int64[] yFeeMint)
func (_Usdpgetters *UsdpgettersCallerSession) GetCollateralMintFees(collateral common.Address) (struct {
	XFeeMint []uint64
	YFeeMint []int64
}, error) {
	return _Usdpgetters.Contract.GetCollateralMintFees(&_Usdpgetters.CallOpts, collateral)
}

// GetCollateralRatio is a free data retrieval call binding the contract method 0xcd377c53.
//
// Solidity: function getCollateralRatio() view returns(uint64 collatRatio, uint256 stablecoinsIssued)
func (_Usdpgetters *UsdpgettersCaller) GetCollateralRatio(opts *bind.CallOpts) (struct {
	CollatRatio       uint64
	StablecoinsIssued *big.Int
}, error) {
	var out []interface{}
	err := _Usdpgetters.contract.Call(opts, &out, "getCollateralRatio")

	outstruct := new(struct {
		CollatRatio       uint64
		StablecoinsIssued *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CollatRatio = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.StablecoinsIssued = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetCollateralRatio is a free data retrieval call binding the contract method 0xcd377c53.
//
// Solidity: function getCollateralRatio() view returns(uint64 collatRatio, uint256 stablecoinsIssued)
func (_Usdpgetters *UsdpgettersSession) GetCollateralRatio() (struct {
	CollatRatio       uint64
	StablecoinsIssued *big.Int
}, error) {
	return _Usdpgetters.Contract.GetCollateralRatio(&_Usdpgetters.CallOpts)
}

// GetCollateralRatio is a free data retrieval call binding the contract method 0xcd377c53.
//
// Solidity: function getCollateralRatio() view returns(uint64 collatRatio, uint256 stablecoinsIssued)
func (_Usdpgetters *UsdpgettersCallerSession) GetCollateralRatio() (struct {
	CollatRatio       uint64
	StablecoinsIssued *big.Int
}, error) {
	return _Usdpgetters.Contract.GetCollateralRatio(&_Usdpgetters.CallOpts)
}

// GetCollateralWhitelistData is a free data retrieval call binding the contract method 0x782513bd.
//
// Solidity: function getCollateralWhitelistData(address collateral) view returns(bytes)
func (_Usdpgetters *UsdpgettersCaller) GetCollateralWhitelistData(opts *bind.CallOpts, collateral common.Address) ([]byte, error) {
	var out []interface{}
	err := _Usdpgetters.contract.Call(opts, &out, "getCollateralWhitelistData", collateral)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetCollateralWhitelistData is a free data retrieval call binding the contract method 0x782513bd.
//
// Solidity: function getCollateralWhitelistData(address collateral) view returns(bytes)
func (_Usdpgetters *UsdpgettersSession) GetCollateralWhitelistData(collateral common.Address) ([]byte, error) {
	return _Usdpgetters.Contract.GetCollateralWhitelistData(&_Usdpgetters.CallOpts, collateral)
}

// GetCollateralWhitelistData is a free data retrieval call binding the contract method 0x782513bd.
//
// Solidity: function getCollateralWhitelistData(address collateral) view returns(bytes)
func (_Usdpgetters *UsdpgettersCallerSession) GetCollateralWhitelistData(collateral common.Address) ([]byte, error) {
	return _Usdpgetters.Contract.GetCollateralWhitelistData(&_Usdpgetters.CallOpts, collateral)
}

// GetIssuedByCollateral is a free data retrieval call binding the contract method 0x94e35d9e.
//
// Solidity: function getIssuedByCollateral(address collateral) view returns(uint256 stablecoinsFromCollateral, uint256 stablecoinsIssued)
func (_Usdpgetters *UsdpgettersCaller) GetIssuedByCollateral(opts *bind.CallOpts, collateral common.Address) (struct {
	StablecoinsFromCollateral *big.Int
	StablecoinsIssued         *big.Int
}, error) {
	var out []interface{}
	err := _Usdpgetters.contract.Call(opts, &out, "getIssuedByCollateral", collateral)

	outstruct := new(struct {
		StablecoinsFromCollateral *big.Int
		StablecoinsIssued         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StablecoinsFromCollateral = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StablecoinsIssued = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetIssuedByCollateral is a free data retrieval call binding the contract method 0x94e35d9e.
//
// Solidity: function getIssuedByCollateral(address collateral) view returns(uint256 stablecoinsFromCollateral, uint256 stablecoinsIssued)
func (_Usdpgetters *UsdpgettersSession) GetIssuedByCollateral(collateral common.Address) (struct {
	StablecoinsFromCollateral *big.Int
	StablecoinsIssued         *big.Int
}, error) {
	return _Usdpgetters.Contract.GetIssuedByCollateral(&_Usdpgetters.CallOpts, collateral)
}

// GetIssuedByCollateral is a free data retrieval call binding the contract method 0x94e35d9e.
//
// Solidity: function getIssuedByCollateral(address collateral) view returns(uint256 stablecoinsFromCollateral, uint256 stablecoinsIssued)
func (_Usdpgetters *UsdpgettersCallerSession) GetIssuedByCollateral(collateral common.Address) (struct {
	StablecoinsFromCollateral *big.Int
	StablecoinsIssued         *big.Int
}, error) {
	return _Usdpgetters.Contract.GetIssuedByCollateral(&_Usdpgetters.CallOpts, collateral)
}

// GetManagerData is a free data retrieval call binding the contract method 0x4ea3e343.
//
// Solidity: function getManagerData(address collateral) view returns(bool, address[], bytes)
func (_Usdpgetters *UsdpgettersCaller) GetManagerData(opts *bind.CallOpts, collateral common.Address) (bool, []common.Address, []byte, error) {
	var out []interface{}
	err := _Usdpgetters.contract.Call(opts, &out, "getManagerData", collateral)

	if err != nil {
		return *new(bool), *new([]common.Address), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	out2 := *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return out0, out1, out2, err

}

// GetManagerData is a free data retrieval call binding the contract method 0x4ea3e343.
//
// Solidity: function getManagerData(address collateral) view returns(bool, address[], bytes)
func (_Usdpgetters *UsdpgettersSession) GetManagerData(collateral common.Address) (bool, []common.Address, []byte, error) {
	return _Usdpgetters.Contract.GetManagerData(&_Usdpgetters.CallOpts, collateral)
}

// GetManagerData is a free data retrieval call binding the contract method 0x4ea3e343.
//
// Solidity: function getManagerData(address collateral) view returns(bool, address[], bytes)
func (_Usdpgetters *UsdpgettersCallerSession) GetManagerData(collateral common.Address) (bool, []common.Address, []byte, error) {
	return _Usdpgetters.Contract.GetManagerData(&_Usdpgetters.CallOpts, collateral)
}

// GetOracle is a free data retrieval call binding the contract method 0x10d3d22e.
//
// Solidity: function getOracle(address collateral) view returns(uint8 oracleType, uint8 targetType, bytes oracleData, bytes targetData, bytes hyperparameters)
func (_Usdpgetters *UsdpgettersCaller) GetOracle(opts *bind.CallOpts, collateral common.Address) (struct {
	OracleType      uint8
	TargetType      uint8
	OracleData      []byte
	TargetData      []byte
	Hyperparameters []byte
}, error) {
	var out []interface{}
	err := _Usdpgetters.contract.Call(opts, &out, "getOracle", collateral)

	outstruct := new(struct {
		OracleType      uint8
		TargetType      uint8
		OracleData      []byte
		TargetData      []byte
		Hyperparameters []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OracleType = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.TargetType = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.OracleData = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.TargetData = *abi.ConvertType(out[3], new([]byte)).(*[]byte)
	outstruct.Hyperparameters = *abi.ConvertType(out[4], new([]byte)).(*[]byte)

	return *outstruct, err

}

// GetOracle is a free data retrieval call binding the contract method 0x10d3d22e.
//
// Solidity: function getOracle(address collateral) view returns(uint8 oracleType, uint8 targetType, bytes oracleData, bytes targetData, bytes hyperparameters)
func (_Usdpgetters *UsdpgettersSession) GetOracle(collateral common.Address) (struct {
	OracleType      uint8
	TargetType      uint8
	OracleData      []byte
	TargetData      []byte
	Hyperparameters []byte
}, error) {
	return _Usdpgetters.Contract.GetOracle(&_Usdpgetters.CallOpts, collateral)
}

// GetOracle is a free data retrieval call binding the contract method 0x10d3d22e.
//
// Solidity: function getOracle(address collateral) view returns(uint8 oracleType, uint8 targetType, bytes oracleData, bytes targetData, bytes hyperparameters)
func (_Usdpgetters *UsdpgettersCallerSession) GetOracle(collateral common.Address) (struct {
	OracleType      uint8
	TargetType      uint8
	OracleData      []byte
	TargetData      []byte
	Hyperparameters []byte
}, error) {
	return _Usdpgetters.Contract.GetOracle(&_Usdpgetters.CallOpts, collateral)
}

// GetOracleValues is a free data retrieval call binding the contract method 0x38c269eb.
//
// Solidity: function getOracleValues(address collateral) view returns(uint256 mint, uint256 burn, uint256 ratio, uint256 minRatio, uint256 redemption)
func (_Usdpgetters *UsdpgettersCaller) GetOracleValues(opts *bind.CallOpts, collateral common.Address) (struct {
	Mint       *big.Int
	Burn       *big.Int
	Ratio      *big.Int
	MinRatio   *big.Int
	Redemption *big.Int
}, error) {
	var out []interface{}
	err := _Usdpgetters.contract.Call(opts, &out, "getOracleValues", collateral)

	outstruct := new(struct {
		Mint       *big.Int
		Burn       *big.Int
		Ratio      *big.Int
		MinRatio   *big.Int
		Redemption *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Mint = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Burn = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Ratio = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MinRatio = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Redemption = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetOracleValues is a free data retrieval call binding the contract method 0x38c269eb.
//
// Solidity: function getOracleValues(address collateral) view returns(uint256 mint, uint256 burn, uint256 ratio, uint256 minRatio, uint256 redemption)
func (_Usdpgetters *UsdpgettersSession) GetOracleValues(collateral common.Address) (struct {
	Mint       *big.Int
	Burn       *big.Int
	Ratio      *big.Int
	MinRatio   *big.Int
	Redemption *big.Int
}, error) {
	return _Usdpgetters.Contract.GetOracleValues(&_Usdpgetters.CallOpts, collateral)
}

// GetOracleValues is a free data retrieval call binding the contract method 0x38c269eb.
//
// Solidity: function getOracleValues(address collateral) view returns(uint256 mint, uint256 burn, uint256 ratio, uint256 minRatio, uint256 redemption)
func (_Usdpgetters *UsdpgettersCallerSession) GetOracleValues(collateral common.Address) (struct {
	Mint       *big.Int
	Burn       *big.Int
	Ratio      *big.Int
	MinRatio   *big.Int
	Redemption *big.Int
}, error) {
	return _Usdpgetters.Contract.GetOracleValues(&_Usdpgetters.CallOpts, collateral)
}

// GetRedemptionFees is a free data retrieval call binding the contract method 0xadc9d1f7.
//
// Solidity: function getRedemptionFees() view returns(uint64[] xRedemptionCurve, int64[] yRedemptionCurve)
func (_Usdpgetters *UsdpgettersCaller) GetRedemptionFees(opts *bind.CallOpts) (struct {
	XRedemptionCurve []uint64
	YRedemptionCurve []int64
}, error) {
	var out []interface{}
	err := _Usdpgetters.contract.Call(opts, &out, "getRedemptionFees")

	outstruct := new(struct {
		XRedemptionCurve []uint64
		YRedemptionCurve []int64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.XRedemptionCurve = *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)
	outstruct.YRedemptionCurve = *abi.ConvertType(out[1], new([]int64)).(*[]int64)

	return *outstruct, err

}

// GetRedemptionFees is a free data retrieval call binding the contract method 0xadc9d1f7.
//
// Solidity: function getRedemptionFees() view returns(uint64[] xRedemptionCurve, int64[] yRedemptionCurve)
func (_Usdpgetters *UsdpgettersSession) GetRedemptionFees() (struct {
	XRedemptionCurve []uint64
	YRedemptionCurve []int64
}, error) {
	return _Usdpgetters.Contract.GetRedemptionFees(&_Usdpgetters.CallOpts)
}

// GetRedemptionFees is a free data retrieval call binding the contract method 0xadc9d1f7.
//
// Solidity: function getRedemptionFees() view returns(uint64[] xRedemptionCurve, int64[] yRedemptionCurve)
func (_Usdpgetters *UsdpgettersCallerSession) GetRedemptionFees() (struct {
	XRedemptionCurve []uint64
	YRedemptionCurve []int64
}, error) {
	return _Usdpgetters.Contract.GetRedemptionFees(&_Usdpgetters.CallOpts)
}

// GetStablecoinCap is a free data retrieval call binding the contract method 0x31da6b13.
//
// Solidity: function getStablecoinCap(address collateral) view returns(uint256)
func (_Usdpgetters *UsdpgettersCaller) GetStablecoinCap(opts *bind.CallOpts, collateral common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Usdpgetters.contract.Call(opts, &out, "getStablecoinCap", collateral)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStablecoinCap is a free data retrieval call binding the contract method 0x31da6b13.
//
// Solidity: function getStablecoinCap(address collateral) view returns(uint256)
func (_Usdpgetters *UsdpgettersSession) GetStablecoinCap(collateral common.Address) (*big.Int, error) {
	return _Usdpgetters.Contract.GetStablecoinCap(&_Usdpgetters.CallOpts, collateral)
}

// GetStablecoinCap is a free data retrieval call binding the contract method 0x31da6b13.
//
// Solidity: function getStablecoinCap(address collateral) view returns(uint256)
func (_Usdpgetters *UsdpgettersCallerSession) GetStablecoinCap(collateral common.Address) (*big.Int, error) {
	return _Usdpgetters.Contract.GetStablecoinCap(&_Usdpgetters.CallOpts, collateral)
}

// GetTotalIssued is a free data retrieval call binding the contract method 0x8db9653f.
//
// Solidity: function getTotalIssued() view returns(uint256)
func (_Usdpgetters *UsdpgettersCaller) GetTotalIssued(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Usdpgetters.contract.Call(opts, &out, "getTotalIssued")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalIssued is a free data retrieval call binding the contract method 0x8db9653f.
//
// Solidity: function getTotalIssued() view returns(uint256)
func (_Usdpgetters *UsdpgettersSession) GetTotalIssued() (*big.Int, error) {
	return _Usdpgetters.Contract.GetTotalIssued(&_Usdpgetters.CallOpts)
}

// GetTotalIssued is a free data retrieval call binding the contract method 0x8db9653f.
//
// Solidity: function getTotalIssued() view returns(uint256)
func (_Usdpgetters *UsdpgettersCallerSession) GetTotalIssued() (*big.Int, error) {
	return _Usdpgetters.Contract.GetTotalIssued(&_Usdpgetters.CallOpts)
}

// IsConsumingScheduledOp is a free data retrieval call binding the contract method 0x8fb36037.
//
// Solidity: function isConsumingScheduledOp() view returns(bytes4)
func (_Usdpgetters *UsdpgettersCaller) IsConsumingScheduledOp(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _Usdpgetters.contract.Call(opts, &out, "isConsumingScheduledOp")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsConsumingScheduledOp is a free data retrieval call binding the contract method 0x8fb36037.
//
// Solidity: function isConsumingScheduledOp() view returns(bytes4)
func (_Usdpgetters *UsdpgettersSession) IsConsumingScheduledOp() ([4]byte, error) {
	return _Usdpgetters.Contract.IsConsumingScheduledOp(&_Usdpgetters.CallOpts)
}

// IsConsumingScheduledOp is a free data retrieval call binding the contract method 0x8fb36037.
//
// Solidity: function isConsumingScheduledOp() view returns(bytes4)
func (_Usdpgetters *UsdpgettersCallerSession) IsConsumingScheduledOp() ([4]byte, error) {
	return _Usdpgetters.Contract.IsConsumingScheduledOp(&_Usdpgetters.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0x0d126627.
//
// Solidity: function isPaused(address collateral, uint8 action) view returns(bool)
func (_Usdpgetters *UsdpgettersCaller) IsPaused(opts *bind.CallOpts, collateral common.Address, action uint8) (bool, error) {
	var out []interface{}
	err := _Usdpgetters.contract.Call(opts, &out, "isPaused", collateral, action)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0x0d126627.
//
// Solidity: function isPaused(address collateral, uint8 action) view returns(bool)
func (_Usdpgetters *UsdpgettersSession) IsPaused(collateral common.Address, action uint8) (bool, error) {
	return _Usdpgetters.Contract.IsPaused(&_Usdpgetters.CallOpts, collateral, action)
}

// IsPaused is a free data retrieval call binding the contract method 0x0d126627.
//
// Solidity: function isPaused(address collateral, uint8 action) view returns(bool)
func (_Usdpgetters *UsdpgettersCallerSession) IsPaused(collateral common.Address, action uint8) (bool, error) {
	return _Usdpgetters.Contract.IsPaused(&_Usdpgetters.CallOpts, collateral, action)
}

// IsTrusted is a free data retrieval call binding the contract method 0x96d64879.
//
// Solidity: function isTrusted(address sender) view returns(bool)
func (_Usdpgetters *UsdpgettersCaller) IsTrusted(opts *bind.CallOpts, sender common.Address) (bool, error) {
	var out []interface{}
	err := _Usdpgetters.contract.Call(opts, &out, "isTrusted", sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrusted is a free data retrieval call binding the contract method 0x96d64879.
//
// Solidity: function isTrusted(address sender) view returns(bool)
func (_Usdpgetters *UsdpgettersSession) IsTrusted(sender common.Address) (bool, error) {
	return _Usdpgetters.Contract.IsTrusted(&_Usdpgetters.CallOpts, sender)
}

// IsTrusted is a free data retrieval call binding the contract method 0x96d64879.
//
// Solidity: function isTrusted(address sender) view returns(bool)
func (_Usdpgetters *UsdpgettersCallerSession) IsTrusted(sender common.Address) (bool, error) {
	return _Usdpgetters.Contract.IsTrusted(&_Usdpgetters.CallOpts, sender)
}

// IsTrustedSeller is a free data retrieval call binding the contract method 0xfe7d0c54.
//
// Solidity: function isTrustedSeller(address sender) view returns(bool)
func (_Usdpgetters *UsdpgettersCaller) IsTrustedSeller(opts *bind.CallOpts, sender common.Address) (bool, error) {
	var out []interface{}
	err := _Usdpgetters.contract.Call(opts, &out, "isTrustedSeller", sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedSeller is a free data retrieval call binding the contract method 0xfe7d0c54.
//
// Solidity: function isTrustedSeller(address sender) view returns(bool)
func (_Usdpgetters *UsdpgettersSession) IsTrustedSeller(sender common.Address) (bool, error) {
	return _Usdpgetters.Contract.IsTrustedSeller(&_Usdpgetters.CallOpts, sender)
}

// IsTrustedSeller is a free data retrieval call binding the contract method 0xfe7d0c54.
//
// Solidity: function isTrustedSeller(address sender) view returns(bool)
func (_Usdpgetters *UsdpgettersCallerSession) IsTrustedSeller(sender common.Address) (bool, error) {
	return _Usdpgetters.Contract.IsTrustedSeller(&_Usdpgetters.CallOpts, sender)
}

// IsValidSelector is a free data retrieval call binding the contract method 0x77dc3429.
//
// Solidity: function isValidSelector(bytes4 selector) view returns(bool)
func (_Usdpgetters *UsdpgettersCaller) IsValidSelector(opts *bind.CallOpts, selector [4]byte) (bool, error) {
	var out []interface{}
	err := _Usdpgetters.contract.Call(opts, &out, "isValidSelector", selector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidSelector is a free data retrieval call binding the contract method 0x77dc3429.
//
// Solidity: function isValidSelector(bytes4 selector) view returns(bool)
func (_Usdpgetters *UsdpgettersSession) IsValidSelector(selector [4]byte) (bool, error) {
	return _Usdpgetters.Contract.IsValidSelector(&_Usdpgetters.CallOpts, selector)
}

// IsValidSelector is a free data retrieval call binding the contract method 0x77dc3429.
//
// Solidity: function isValidSelector(bytes4 selector) view returns(bool)
func (_Usdpgetters *UsdpgettersCallerSession) IsValidSelector(selector [4]byte) (bool, error) {
	return _Usdpgetters.Contract.IsValidSelector(&_Usdpgetters.CallOpts, selector)
}

// IsWhitelistedCollateral is a free data retrieval call binding the contract method 0xf9839d89.
//
// Solidity: function isWhitelistedCollateral(address collateral) view returns(bool)
func (_Usdpgetters *UsdpgettersCaller) IsWhitelistedCollateral(opts *bind.CallOpts, collateral common.Address) (bool, error) {
	var out []interface{}
	err := _Usdpgetters.contract.Call(opts, &out, "isWhitelistedCollateral", collateral)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelistedCollateral is a free data retrieval call binding the contract method 0xf9839d89.
//
// Solidity: function isWhitelistedCollateral(address collateral) view returns(bool)
func (_Usdpgetters *UsdpgettersSession) IsWhitelistedCollateral(collateral common.Address) (bool, error) {
	return _Usdpgetters.Contract.IsWhitelistedCollateral(&_Usdpgetters.CallOpts, collateral)
}

// IsWhitelistedCollateral is a free data retrieval call binding the contract method 0xf9839d89.
//
// Solidity: function isWhitelistedCollateral(address collateral) view returns(bool)
func (_Usdpgetters *UsdpgettersCallerSession) IsWhitelistedCollateral(collateral common.Address) (bool, error) {
	return _Usdpgetters.Contract.IsWhitelistedCollateral(&_Usdpgetters.CallOpts, collateral)
}

// IsWhitelistedForType is a free data retrieval call binding the contract method 0x99eeca49.
//
// Solidity: function isWhitelistedForType(uint8 whitelistType, address sender) view returns(bool)
func (_Usdpgetters *UsdpgettersCaller) IsWhitelistedForType(opts *bind.CallOpts, whitelistType uint8, sender common.Address) (bool, error) {
	var out []interface{}
	err := _Usdpgetters.contract.Call(opts, &out, "isWhitelistedForType", whitelistType, sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelistedForType is a free data retrieval call binding the contract method 0x99eeca49.
//
// Solidity: function isWhitelistedForType(uint8 whitelistType, address sender) view returns(bool)
func (_Usdpgetters *UsdpgettersSession) IsWhitelistedForType(whitelistType uint8, sender common.Address) (bool, error) {
	return _Usdpgetters.Contract.IsWhitelistedForType(&_Usdpgetters.CallOpts, whitelistType, sender)
}

// IsWhitelistedForType is a free data retrieval call binding the contract method 0x99eeca49.
//
// Solidity: function isWhitelistedForType(uint8 whitelistType, address sender) view returns(bool)
func (_Usdpgetters *UsdpgettersCallerSession) IsWhitelistedForType(whitelistType uint8, sender common.Address) (bool, error) {
	return _Usdpgetters.Contract.IsWhitelistedForType(&_Usdpgetters.CallOpts, whitelistType, sender)
}

// TokenP is a free data retrieval call binding the contract method 0x1978a5ed.
//
// Solidity: function tokenP() view returns(address)
func (_Usdpgetters *UsdpgettersCaller) TokenP(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Usdpgetters.contract.Call(opts, &out, "tokenP")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenP is a free data retrieval call binding the contract method 0x1978a5ed.
//
// Solidity: function tokenP() view returns(address)
func (_Usdpgetters *UsdpgettersSession) TokenP() (common.Address, error) {
	return _Usdpgetters.Contract.TokenP(&_Usdpgetters.CallOpts)
}

// TokenP is a free data retrieval call binding the contract method 0x1978a5ed.
//
// Solidity: function tokenP() view returns(address)
func (_Usdpgetters *UsdpgettersCallerSession) TokenP() (common.Address, error) {
	return _Usdpgetters.Contract.TokenP(&_Usdpgetters.CallOpts)
}

// IsWhitelistedForCollateral is a paid mutator transaction binding the contract method 0xa52aefd4.
//
// Solidity: function isWhitelistedForCollateral(address collateral, address sender) returns(bool)
func (_Usdpgetters *UsdpgettersTransactor) IsWhitelistedForCollateral(opts *bind.TransactOpts, collateral common.Address, sender common.Address) (*types.Transaction, error) {
	return _Usdpgetters.contract.Transact(opts, "isWhitelistedForCollateral", collateral, sender)
}

// IsWhitelistedForCollateral is a paid mutator transaction binding the contract method 0xa52aefd4.
//
// Solidity: function isWhitelistedForCollateral(address collateral, address sender) returns(bool)
func (_Usdpgetters *UsdpgettersSession) IsWhitelistedForCollateral(collateral common.Address, sender common.Address) (*types.Transaction, error) {
	return _Usdpgetters.Contract.IsWhitelistedForCollateral(&_Usdpgetters.TransactOpts, collateral, sender)
}

// IsWhitelistedForCollateral is a paid mutator transaction binding the contract method 0xa52aefd4.
//
// Solidity: function isWhitelistedForCollateral(address collateral, address sender) returns(bool)
func (_Usdpgetters *UsdpgettersTransactorSession) IsWhitelistedForCollateral(collateral common.Address, sender common.Address) (*types.Transaction, error) {
	return _Usdpgetters.Contract.IsWhitelistedForCollateral(&_Usdpgetters.TransactOpts, collateral, sender)
}
