// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package usdp

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

// FacetCut is an auto generated low-level Go binding around an user-defined struct.
type FacetCut struct {
	FacetAddress      common.Address
	Action            uint8
	FunctionSelectors [][4]byte
}

// UsdpMetaData contains all meta data concerning the Usdp contract.
var UsdpMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumFacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structFacetCut[]\",\"name\":\"_diamondCut\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"_init\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"CannotAddFunctionToDiamondThatAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4[]\",\"name\":\"_selectors\",\"type\":\"bytes4[]\"}],\"name\":\"CannotAddSelectorsToZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"CannotRemoveFunctionThatDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"CannotRemoveImmutableFunction\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"CannotReplaceFunctionThatDoesNotExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"CannotReplaceFunctionWithTheSameFunctionFromTheSameFacet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4[]\",\"name\":\"_selectors\",\"type\":\"bytes4[]\"}],\"name\":\"CannotReplaceFunctionsFromFacetWithZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"CannotReplaceImmutableFunction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ContractHasNoCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_functionSelector\",\"type\":\"bytes4\"}],\"name\":\"FunctionNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_initializationContractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"InitializationFunctionReverted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_facetAddress\",\"type\":\"address\"}],\"name\":\"NoSelectorsProvidedForFacetForCut\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_facetAddress\",\"type\":\"address\"}],\"name\":\"RemoveFacetAddressMustBeZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumFacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"indexed\":false,\"internalType\":\"structFacetCut[]\",\"name\":\"_diamondCut\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_init\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"DiamondCut\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"}]",
}

// UsdpABI is the input ABI used to generate the binding from.
// Deprecated: Use UsdpMetaData.ABI instead.
var UsdpABI = UsdpMetaData.ABI

// Usdp is an auto generated Go binding around an Ethereum contract.
type Usdp struct {
	UsdpCaller     // Read-only binding to the contract
	UsdpTransactor // Write-only binding to the contract
	UsdpFilterer   // Log filterer for contract events
}

// UsdpCaller is an auto generated read-only Go binding around an Ethereum contract.
type UsdpCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsdpTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UsdpTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsdpFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UsdpFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsdpSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UsdpSession struct {
	Contract     *Usdp             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UsdpCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UsdpCallerSession struct {
	Contract *UsdpCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UsdpTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UsdpTransactorSession struct {
	Contract     *UsdpTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UsdpRaw is an auto generated low-level Go binding around an Ethereum contract.
type UsdpRaw struct {
	Contract *Usdp // Generic contract binding to access the raw methods on
}

// UsdpCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UsdpCallerRaw struct {
	Contract *UsdpCaller // Generic read-only contract binding to access the raw methods on
}

// UsdpTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UsdpTransactorRaw struct {
	Contract *UsdpTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUsdp creates a new instance of Usdp, bound to a specific deployed contract.
func NewUsdp(address common.Address, backend bind.ContractBackend) (*Usdp, error) {
	contract, err := bindUsdp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Usdp{UsdpCaller: UsdpCaller{contract: contract}, UsdpTransactor: UsdpTransactor{contract: contract}, UsdpFilterer: UsdpFilterer{contract: contract}}, nil
}

// NewUsdpCaller creates a new read-only instance of Usdp, bound to a specific deployed contract.
func NewUsdpCaller(address common.Address, caller bind.ContractCaller) (*UsdpCaller, error) {
	contract, err := bindUsdp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UsdpCaller{contract: contract}, nil
}

// NewUsdpTransactor creates a new write-only instance of Usdp, bound to a specific deployed contract.
func NewUsdpTransactor(address common.Address, transactor bind.ContractTransactor) (*UsdpTransactor, error) {
	contract, err := bindUsdp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UsdpTransactor{contract: contract}, nil
}

// NewUsdpFilterer creates a new log filterer instance of Usdp, bound to a specific deployed contract.
func NewUsdpFilterer(address common.Address, filterer bind.ContractFilterer) (*UsdpFilterer, error) {
	contract, err := bindUsdp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UsdpFilterer{contract: contract}, nil
}

// bindUsdp binds a generic wrapper to an already deployed contract.
func bindUsdp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UsdpMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Usdp *UsdpRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Usdp.Contract.UsdpCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Usdp *UsdpRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Usdp.Contract.UsdpTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Usdp *UsdpRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Usdp.Contract.UsdpTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Usdp *UsdpCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Usdp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Usdp *UsdpTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Usdp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Usdp *UsdpTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Usdp.Contract.contract.Transact(opts, method, params...)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Usdp *UsdpTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Usdp.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Usdp *UsdpSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Usdp.Contract.Fallback(&_Usdp.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Usdp *UsdpTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Usdp.Contract.Fallback(&_Usdp.TransactOpts, calldata)
}

// UsdpDiamondCutIterator is returned from FilterDiamondCut and is used to iterate over the raw logs and unpacked data for DiamondCut events raised by the Usdp contract.
type UsdpDiamondCutIterator struct {
	Event *UsdpDiamondCut // Event containing the contract specifics and raw log

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
func (it *UsdpDiamondCutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsdpDiamondCut)
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
		it.Event = new(UsdpDiamondCut)
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
func (it *UsdpDiamondCutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsdpDiamondCutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsdpDiamondCut represents a DiamondCut event raised by the Usdp contract.
type UsdpDiamondCut struct {
	DiamondCut []FacetCut
	Init       common.Address
	Calldata   []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDiamondCut is a free log retrieval operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_Usdp *UsdpFilterer) FilterDiamondCut(opts *bind.FilterOpts) (*UsdpDiamondCutIterator, error) {

	logs, sub, err := _Usdp.contract.FilterLogs(opts, "DiamondCut")
	if err != nil {
		return nil, err
	}
	return &UsdpDiamondCutIterator{contract: _Usdp.contract, event: "DiamondCut", logs: logs, sub: sub}, nil
}

// WatchDiamondCut is a free log subscription operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_Usdp *UsdpFilterer) WatchDiamondCut(opts *bind.WatchOpts, sink chan<- *UsdpDiamondCut) (event.Subscription, error) {

	logs, sub, err := _Usdp.contract.WatchLogs(opts, "DiamondCut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsdpDiamondCut)
				if err := _Usdp.contract.UnpackLog(event, "DiamondCut", log); err != nil {
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

// ParseDiamondCut is a log parse operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_Usdp *UsdpFilterer) ParseDiamondCut(log types.Log) (*UsdpDiamondCut, error) {
	event := new(UsdpDiamondCut)
	if err := _Usdp.contract.UnpackLog(event, "DiamondCut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
