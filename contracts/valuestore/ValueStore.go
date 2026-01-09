// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ValueStore

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

// ValueStoreMetaData contains all meta data concerning the ValueStore contract.
var ValueStoreMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getValue\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"fairValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"valueUsd\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"numerator\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denominator\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setMultipleValues\",\"inputs\":[{\"name\":\"keys\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"fairValues\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"valueUsds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"numerators\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"denominators\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setValue\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"fairValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"valueUsd\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"numerator\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denominator\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValueUpdated\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"fairValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"valueUsd\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"numerator\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"denominator\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6080604052348015600e575f5ffd5b50335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055503373ffffffffffffffffffffffffffffffffffffffff165f73ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36110bf806100b55f395ff3fe608060405234801561000f575f5ffd5b5060043610610055575f3560e01c80634a7edea8146100595780636080b54f146100755780638da5cb5b14610091578063960384a0146100af578063f2fde38b146100e3575b5f5ffd5b610073600480360381019061006e9190610951565b6100ff565b005b61008f600480360381019061008a9190610af0565b61047c565b005b61009961061c565b6040516100a69190610bc5565b60405180910390f35b6100c960048036038101906100c49190610bde565b610640565b6040516100da959493929190610c38565b60405180910390f35b6100fd60048036038101906100f89190610cb3565b6106db565b005b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461018d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161018490610d38565b60405180910390fd5b878790508a8a90501480156101a75750858590508a8a9050145b80156101b85750838390508a8a9050145b80156101c95750818190508a8a9050145b610208576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101ff90610da0565b60405180910390fd5b5f5f90505b8a8a905081101561046f575f85858381811061022c5761022b610dbe565b5b90506020020135148061025857505f83838381811061024e5761024d610dbe565b5b9050602002013514155b610297576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161028e90610e35565b60405180910390fd5b6040518060a001604052808a8a848181106102b5576102b4610dbe565b5b9050602002013581526020018888848181106102d4576102d3610dbe565b5b9050602002013581526020018686848181106102f3576102f2610dbe565b5b90506020020135815260200184848481811061031257610311610dbe565b5b9050602002013581526020014281525060018c8c8481811061033757610336610dbe565b5b90506020028101906103499190610e5f565b604051610357929190610efd565b90815260200160405180910390205f820151815f0155602082015181600101556040820151816002015560608201518160030155608082015181600401559050507f5f2a8c3a90ec95498a7028ec8d4e67159d8aa0e0bd28284ea430ba0f1da6877d8b8b838181106103cc576103cb610dbe565b5b90506020028101906103de9190610e5f565b8b8b858181106103f1576103f0610dbe565b5b905060200201358a8a8681811061040b5761040a610dbe565b5b9050602002013589898781811061042557610424610dbe565b5b9050602002013588888881811061043f5761043e610dbe565b5b905060200201354260405161045a9796959493929190610f51565b60405180910390a1808060010191505061020d565b5050505050505050505050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461050a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161050190610d38565b60405180910390fd5b5f82148061051857505f8114155b610557576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161054e90610e35565b60405180910390fd5b6040518060a001604052808581526020018481526020018381526020018281526020014281525060018787604051610590929190610efd565b90815260200160405180910390205f820151815f0155602082015181600101556040820151816002015560608201518160030155608082015181600401559050507f5f2a8c3a90ec95498a7028ec8d4e67159d8aa0e0bd28284ea430ba0f1da6877d8686868686864260405161060c9796959493929190610f51565b60405180910390a1505050505050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f5f5f5f5f5f60018888604051610658929190610efd565b908152602001604051809103902090505f8160040154036106ae576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106a590611003565b60405180910390fd5b805f0154816001015482600201548360030154846004015495509550955095509550509295509295909350565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610769576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161076090610d38565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036107d7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107ce9061106b565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff165f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3805f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f8401126108bc576108bb61089b565b5b8235905067ffffffffffffffff8111156108d9576108d861089f565b5b6020830191508360208202830111156108f5576108f46108a3565b5b9250929050565b5f5f83601f8401126109115761091061089b565b5b8235905067ffffffffffffffff81111561092e5761092d61089f565b5b60208301915083602082028301111561094a576109496108a3565b5b9250929050565b5f5f5f5f5f5f5f5f5f5f60a08b8d03121561096f5761096e610893565b5b5f8b013567ffffffffffffffff81111561098c5761098b610897565b5b6109988d828e016108a7565b9a509a505060208b013567ffffffffffffffff8111156109bb576109ba610897565b5b6109c78d828e016108fc565b985098505060408b013567ffffffffffffffff8111156109ea576109e9610897565b5b6109f68d828e016108fc565b965096505060608b013567ffffffffffffffff811115610a1957610a18610897565b5b610a258d828e016108fc565b945094505060808b013567ffffffffffffffff811115610a4857610a47610897565b5b610a548d828e016108fc565b92509250509295989b9194979a5092959850565b5f5f83601f840112610a7d57610a7c61089b565b5b8235905067ffffffffffffffff811115610a9a57610a9961089f565b5b602083019150836001820283011115610ab657610ab56108a3565b5b9250929050565b5f819050919050565b610acf81610abd565b8114610ad9575f5ffd5b50565b5f81359050610aea81610ac6565b92915050565b5f5f5f5f5f5f60a08789031215610b0a57610b09610893565b5b5f87013567ffffffffffffffff811115610b2757610b26610897565b5b610b3389828a01610a68565b96509650506020610b4689828a01610adc565b9450506040610b5789828a01610adc565b9350506060610b6889828a01610adc565b9250506080610b7989828a01610adc565b9150509295509295509295565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610baf82610b86565b9050919050565b610bbf81610ba5565b82525050565b5f602082019050610bd85f830184610bb6565b92915050565b5f5f60208385031215610bf457610bf3610893565b5b5f83013567ffffffffffffffff811115610c1157610c10610897565b5b610c1d85828601610a68565b92509250509250929050565b610c3281610abd565b82525050565b5f60a082019050610c4b5f830188610c29565b610c586020830187610c29565b610c656040830186610c29565b610c726060830185610c29565b610c7f6080830184610c29565b9695505050505050565b610c9281610ba5565b8114610c9c575f5ffd5b50565b5f81359050610cad81610c89565b92915050565b5f60208284031215610cc857610cc7610893565b5b5f610cd584828501610c9f565b91505092915050565b5f82825260208201905092915050565b7f43616c6c6572206973206e6f7420746865206f776e65720000000000000000005f82015250565b5f610d22601783610cde565b9150610d2d82610cee565b602082019050919050565b5f6020820190508181035f830152610d4f81610d16565b9050919050565b7f4172726179206c656e67746873206d757374206d6174636800000000000000005f82015250565b5f610d8a601883610cde565b9150610d9582610d56565b602082019050919050565b5f6020820190508181035f830152610db781610d7e565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f6469766973696f6e206279207a65726f000000000000000000000000000000005f82015250565b5f610e1f601083610cde565b9150610e2a82610deb565b602082019050919050565b5f6020820190508181035f830152610e4c81610e13565b9050919050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83356001602003843603038112610e7b57610e7a610e53565b5b80840192508235915067ffffffffffffffff821115610e9d57610e9c610e57565b5b602083019250600182023603831315610eb957610eb8610e5b565b5b509250929050565b5f81905092915050565b828183375f83830152505050565b5f610ee48385610ec1565b9350610ef1838584610ecb565b82840190509392505050565b5f610f09828486610ed9565b91508190509392505050565b5f601f19601f8301169050919050565b5f610f308385610cde565b9350610f3d838584610ecb565b610f4683610f15565b840190509392505050565b5f60c0820190508181035f830152610f6a81898b610f25565b9050610f796020830188610c29565b610f866040830187610c29565b610f936060830186610c29565b610fa06080830185610c29565b610fad60a0830184610c29565b98975050505050505050565b7f4e6f206461746120666f72206b657900000000000000000000000000000000005f82015250565b5f610fed600f83610cde565b9150610ff882610fb9565b602082019050919050565b5f6020820190508181035f83015261101a81610fe1565b9050919050565b7f4e6577206f776e6572206973207a65726f2061646472657373000000000000005f82015250565b5f611055601983610cde565b915061106082611021565b602082019050919050565b5f6020820190508181035f83015261108281611049565b905091905056fea264697066735822122046477468a2fe8c028640964929d9efa4e7b1c8ec32936a7b52a0b5c5c063756264736f6c634300081e0033",
}

// ValueStoreABI is the input ABI used to generate the binding from.
// Deprecated: Use ValueStoreMetaData.ABI instead.
var ValueStoreABI = ValueStoreMetaData.ABI

// ValueStoreBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ValueStoreMetaData.Bin instead.
var ValueStoreBin = ValueStoreMetaData.Bin

// DeployValueStore deploys a new Ethereum contract, binding an instance of ValueStore to it.
func DeployValueStore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ValueStore, error) {
	parsed, err := ValueStoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ValueStoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValueStore{ValueStoreCaller: ValueStoreCaller{contract: contract}, ValueStoreTransactor: ValueStoreTransactor{contract: contract}, ValueStoreFilterer: ValueStoreFilterer{contract: contract}}, nil
}

// ValueStore is an auto generated Go binding around an Ethereum contract.
type ValueStore struct {
	ValueStoreCaller     // Read-only binding to the contract
	ValueStoreTransactor // Write-only binding to the contract
	ValueStoreFilterer   // Log filterer for contract events
}

// ValueStoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValueStoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueStoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValueStoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueStoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValueStoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueStoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValueStoreSession struct {
	Contract     *ValueStore       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValueStoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValueStoreCallerSession struct {
	Contract *ValueStoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ValueStoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValueStoreTransactorSession struct {
	Contract     *ValueStoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ValueStoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValueStoreRaw struct {
	Contract *ValueStore // Generic contract binding to access the raw methods on
}

// ValueStoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValueStoreCallerRaw struct {
	Contract *ValueStoreCaller // Generic read-only contract binding to access the raw methods on
}

// ValueStoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValueStoreTransactorRaw struct {
	Contract *ValueStoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValueStore creates a new instance of ValueStore, bound to a specific deployed contract.
func NewValueStore(address common.Address, backend bind.ContractBackend) (*ValueStore, error) {
	contract, err := bindValueStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValueStore{ValueStoreCaller: ValueStoreCaller{contract: contract}, ValueStoreTransactor: ValueStoreTransactor{contract: contract}, ValueStoreFilterer: ValueStoreFilterer{contract: contract}}, nil
}

// NewValueStoreCaller creates a new read-only instance of ValueStore, bound to a specific deployed contract.
func NewValueStoreCaller(address common.Address, caller bind.ContractCaller) (*ValueStoreCaller, error) {
	contract, err := bindValueStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValueStoreCaller{contract: contract}, nil
}

// NewValueStoreTransactor creates a new write-only instance of ValueStore, bound to a specific deployed contract.
func NewValueStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*ValueStoreTransactor, error) {
	contract, err := bindValueStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValueStoreTransactor{contract: contract}, nil
}

// NewValueStoreFilterer creates a new log filterer instance of ValueStore, bound to a specific deployed contract.
func NewValueStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*ValueStoreFilterer, error) {
	contract, err := bindValueStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValueStoreFilterer{contract: contract}, nil
}

// bindValueStore binds a generic wrapper to an already deployed contract.
func bindValueStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ValueStoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValueStore *ValueStoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValueStore.Contract.ValueStoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValueStore *ValueStoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValueStore.Contract.ValueStoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValueStore *ValueStoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValueStore.Contract.ValueStoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValueStore *ValueStoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValueStore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValueStore *ValueStoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValueStore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValueStore *ValueStoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValueStore.Contract.contract.Transact(opts, method, params...)
}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint256 fairValue, uint256 valueUsd, uint256 numerator, uint256 denominator, uint256 timestamp)
func (_ValueStore *ValueStoreCaller) GetValue(opts *bind.CallOpts, key string) (struct {
	FairValue   *big.Int
	ValueUsd    *big.Int
	Numerator   *big.Int
	Denominator *big.Int
	Timestamp   *big.Int
}, error) {
	var out []interface{}
	err := _ValueStore.contract.Call(opts, &out, "getValue", key)

	outstruct := new(struct {
		FairValue   *big.Int
		ValueUsd    *big.Int
		Numerator   *big.Int
		Denominator *big.Int
		Timestamp   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FairValue = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ValueUsd = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Numerator = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Denominator = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint256 fairValue, uint256 valueUsd, uint256 numerator, uint256 denominator, uint256 timestamp)
func (_ValueStore *ValueStoreSession) GetValue(key string) (struct {
	FairValue   *big.Int
	ValueUsd    *big.Int
	Numerator   *big.Int
	Denominator *big.Int
	Timestamp   *big.Int
}, error) {
	return _ValueStore.Contract.GetValue(&_ValueStore.CallOpts, key)
}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint256 fairValue, uint256 valueUsd, uint256 numerator, uint256 denominator, uint256 timestamp)
func (_ValueStore *ValueStoreCallerSession) GetValue(key string) (struct {
	FairValue   *big.Int
	ValueUsd    *big.Int
	Numerator   *big.Int
	Denominator *big.Int
	Timestamp   *big.Int
}, error) {
	return _ValueStore.Contract.GetValue(&_ValueStore.CallOpts, key)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ValueStore *ValueStoreCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValueStore.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ValueStore *ValueStoreSession) Owner() (common.Address, error) {
	return _ValueStore.Contract.Owner(&_ValueStore.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ValueStore *ValueStoreCallerSession) Owner() (common.Address, error) {
	return _ValueStore.Contract.Owner(&_ValueStore.CallOpts)
}

// SetMultipleValues is a paid mutator transaction binding the contract method 0x4a7edea8.
//
// Solidity: function setMultipleValues(string[] keys, uint256[] fairValues, uint256[] valueUsds, uint256[] numerators, uint256[] denominators) returns()
func (_ValueStore *ValueStoreTransactor) SetMultipleValues(opts *bind.TransactOpts, keys []string, fairValues []*big.Int, valueUsds []*big.Int, numerators []*big.Int, denominators []*big.Int) (*types.Transaction, error) {
	return _ValueStore.contract.Transact(opts, "setMultipleValues", keys, fairValues, valueUsds, numerators, denominators)
}

// SetMultipleValues is a paid mutator transaction binding the contract method 0x4a7edea8.
//
// Solidity: function setMultipleValues(string[] keys, uint256[] fairValues, uint256[] valueUsds, uint256[] numerators, uint256[] denominators) returns()
func (_ValueStore *ValueStoreSession) SetMultipleValues(keys []string, fairValues []*big.Int, valueUsds []*big.Int, numerators []*big.Int, denominators []*big.Int) (*types.Transaction, error) {
	return _ValueStore.Contract.SetMultipleValues(&_ValueStore.TransactOpts, keys, fairValues, valueUsds, numerators, denominators)
}

// SetMultipleValues is a paid mutator transaction binding the contract method 0x4a7edea8.
//
// Solidity: function setMultipleValues(string[] keys, uint256[] fairValues, uint256[] valueUsds, uint256[] numerators, uint256[] denominators) returns()
func (_ValueStore *ValueStoreTransactorSession) SetMultipleValues(keys []string, fairValues []*big.Int, valueUsds []*big.Int, numerators []*big.Int, denominators []*big.Int) (*types.Transaction, error) {
	return _ValueStore.Contract.SetMultipleValues(&_ValueStore.TransactOpts, keys, fairValues, valueUsds, numerators, denominators)
}

// SetValue is a paid mutator transaction binding the contract method 0x6080b54f.
//
// Solidity: function setValue(string key, uint256 fairValue, uint256 valueUsd, uint256 numerator, uint256 denominator) returns()
func (_ValueStore *ValueStoreTransactor) SetValue(opts *bind.TransactOpts, key string, fairValue *big.Int, valueUsd *big.Int, numerator *big.Int, denominator *big.Int) (*types.Transaction, error) {
	return _ValueStore.contract.Transact(opts, "setValue", key, fairValue, valueUsd, numerator, denominator)
}

// SetValue is a paid mutator transaction binding the contract method 0x6080b54f.
//
// Solidity: function setValue(string key, uint256 fairValue, uint256 valueUsd, uint256 numerator, uint256 denominator) returns()
func (_ValueStore *ValueStoreSession) SetValue(key string, fairValue *big.Int, valueUsd *big.Int, numerator *big.Int, denominator *big.Int) (*types.Transaction, error) {
	return _ValueStore.Contract.SetValue(&_ValueStore.TransactOpts, key, fairValue, valueUsd, numerator, denominator)
}

// SetValue is a paid mutator transaction binding the contract method 0x6080b54f.
//
// Solidity: function setValue(string key, uint256 fairValue, uint256 valueUsd, uint256 numerator, uint256 denominator) returns()
func (_ValueStore *ValueStoreTransactorSession) SetValue(key string, fairValue *big.Int, valueUsd *big.Int, numerator *big.Int, denominator *big.Int) (*types.Transaction, error) {
	return _ValueStore.Contract.SetValue(&_ValueStore.TransactOpts, key, fairValue, valueUsd, numerator, denominator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ValueStore *ValueStoreTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ValueStore.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ValueStore *ValueStoreSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ValueStore.Contract.TransferOwnership(&_ValueStore.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ValueStore *ValueStoreTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ValueStore.Contract.TransferOwnership(&_ValueStore.TransactOpts, newOwner)
}

// ValueStoreOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ValueStore contract.
type ValueStoreOwnershipTransferredIterator struct {
	Event *ValueStoreOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ValueStoreOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValueStoreOwnershipTransferred)
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
		it.Event = new(ValueStoreOwnershipTransferred)
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
func (it *ValueStoreOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValueStoreOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValueStoreOwnershipTransferred represents a OwnershipTransferred event raised by the ValueStore contract.
type ValueStoreOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ValueStore *ValueStoreFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ValueStoreOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ValueStore.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ValueStoreOwnershipTransferredIterator{contract: _ValueStore.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ValueStore *ValueStoreFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ValueStoreOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ValueStore.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValueStoreOwnershipTransferred)
				if err := _ValueStore.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ValueStore *ValueStoreFilterer) ParseOwnershipTransferred(log types.Log) (*ValueStoreOwnershipTransferred, error) {
	event := new(ValueStoreOwnershipTransferred)
	if err := _ValueStore.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValueStoreValueUpdatedIterator is returned from FilterValueUpdated and is used to iterate over the raw logs and unpacked data for ValueUpdated events raised by the ValueStore contract.
type ValueStoreValueUpdatedIterator struct {
	Event *ValueStoreValueUpdated // Event containing the contract specifics and raw log

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
func (it *ValueStoreValueUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValueStoreValueUpdated)
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
		it.Event = new(ValueStoreValueUpdated)
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
func (it *ValueStoreValueUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValueStoreValueUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValueStoreValueUpdated represents a ValueUpdated event raised by the ValueStore contract.
type ValueStoreValueUpdated struct {
	Key         string
	FairValue   *big.Int
	ValueUsd    *big.Int
	Numerator   *big.Int
	Denominator *big.Int
	Timestamp   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterValueUpdated is a free log retrieval operation binding the contract event 0x5f2a8c3a90ec95498a7028ec8d4e67159d8aa0e0bd28284ea430ba0f1da6877d.
//
// Solidity: event ValueUpdated(string key, uint256 fairValue, uint256 valueUsd, uint256 numerator, uint256 denominator, uint256 timestamp)
func (_ValueStore *ValueStoreFilterer) FilterValueUpdated(opts *bind.FilterOpts) (*ValueStoreValueUpdatedIterator, error) {

	logs, sub, err := _ValueStore.contract.FilterLogs(opts, "ValueUpdated")
	if err != nil {
		return nil, err
	}
	return &ValueStoreValueUpdatedIterator{contract: _ValueStore.contract, event: "ValueUpdated", logs: logs, sub: sub}, nil
}

// WatchValueUpdated is a free log subscription operation binding the contract event 0x5f2a8c3a90ec95498a7028ec8d4e67159d8aa0e0bd28284ea430ba0f1da6877d.
//
// Solidity: event ValueUpdated(string key, uint256 fairValue, uint256 valueUsd, uint256 numerator, uint256 denominator, uint256 timestamp)
func (_ValueStore *ValueStoreFilterer) WatchValueUpdated(opts *bind.WatchOpts, sink chan<- *ValueStoreValueUpdated) (event.Subscription, error) {

	logs, sub, err := _ValueStore.contract.WatchLogs(opts, "ValueUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValueStoreValueUpdated)
				if err := _ValueStore.contract.UnpackLog(event, "ValueUpdated", log); err != nil {
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

// ParseValueUpdated is a log parse operation binding the contract event 0x5f2a8c3a90ec95498a7028ec8d4e67159d8aa0e0bd28284ea430ba0f1da6877d.
//
// Solidity: event ValueUpdated(string key, uint256 fairValue, uint256 valueUsd, uint256 numerator, uint256 denominator, uint256 timestamp)
func (_ValueStore *ValueStoreFilterer) ParseValueUpdated(log types.Log) (*ValueStoreValueUpdated, error) {
	event := new(ValueStoreValueUpdated)
	if err := _ValueStore.contract.UnpackLog(event, "ValueUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
