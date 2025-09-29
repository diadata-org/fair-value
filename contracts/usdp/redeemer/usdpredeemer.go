// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package usdpredeemer

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

// UsdpredeemerMetaData contains all meta data concerning the Usdpredeemer contract.
var UsdpredeemerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidChainlinkRate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLengths\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotTrusted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Paused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooBigAmountIn\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooLate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooSmallAmountOut\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNormalizerValue\",\"type\":\"uint256\"}],\"name\":\"NormalizerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"forfeitTokens\",\"type\":\"address[]\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Redeemed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"quoteRedemptionCurve\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"minAmountOuts\",\"type\":\"uint256[]\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"minAmountOuts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"forfeitTokens\",\"type\":\"address[]\"}],\"name\":\"redeemWithForfeit\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"increase\",\"type\":\"bool\"}],\"name\":\"updateNormalizer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// UsdpredeemerABI is the input ABI used to generate the binding from.
// Deprecated: Use UsdpredeemerMetaData.ABI instead.
var UsdpredeemerABI = UsdpredeemerMetaData.ABI

// Usdpredeemer is an auto generated Go binding around an Ethereum contract.
type Usdpredeemer struct {
	UsdpredeemerCaller     // Read-only binding to the contract
	UsdpredeemerTransactor // Write-only binding to the contract
	UsdpredeemerFilterer   // Log filterer for contract events
}

// UsdpredeemerCaller is an auto generated read-only Go binding around an Ethereum contract.
type UsdpredeemerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsdpredeemerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UsdpredeemerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsdpredeemerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UsdpredeemerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsdpredeemerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UsdpredeemerSession struct {
	Contract     *Usdpredeemer     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UsdpredeemerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UsdpredeemerCallerSession struct {
	Contract *UsdpredeemerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// UsdpredeemerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UsdpredeemerTransactorSession struct {
	Contract     *UsdpredeemerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// UsdpredeemerRaw is an auto generated low-level Go binding around an Ethereum contract.
type UsdpredeemerRaw struct {
	Contract *Usdpredeemer // Generic contract binding to access the raw methods on
}

// UsdpredeemerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UsdpredeemerCallerRaw struct {
	Contract *UsdpredeemerCaller // Generic read-only contract binding to access the raw methods on
}

// UsdpredeemerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UsdpredeemerTransactorRaw struct {
	Contract *UsdpredeemerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUsdpredeemer creates a new instance of Usdpredeemer, bound to a specific deployed contract.
func NewUsdpredeemer(address common.Address, backend bind.ContractBackend) (*Usdpredeemer, error) {
	contract, err := bindUsdpredeemer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Usdpredeemer{UsdpredeemerCaller: UsdpredeemerCaller{contract: contract}, UsdpredeemerTransactor: UsdpredeemerTransactor{contract: contract}, UsdpredeemerFilterer: UsdpredeemerFilterer{contract: contract}}, nil
}

// NewUsdpredeemerCaller creates a new read-only instance of Usdpredeemer, bound to a specific deployed contract.
func NewUsdpredeemerCaller(address common.Address, caller bind.ContractCaller) (*UsdpredeemerCaller, error) {
	contract, err := bindUsdpredeemer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UsdpredeemerCaller{contract: contract}, nil
}

// NewUsdpredeemerTransactor creates a new write-only instance of Usdpredeemer, bound to a specific deployed contract.
func NewUsdpredeemerTransactor(address common.Address, transactor bind.ContractTransactor) (*UsdpredeemerTransactor, error) {
	contract, err := bindUsdpredeemer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UsdpredeemerTransactor{contract: contract}, nil
}

// NewUsdpredeemerFilterer creates a new log filterer instance of Usdpredeemer, bound to a specific deployed contract.
func NewUsdpredeemerFilterer(address common.Address, filterer bind.ContractFilterer) (*UsdpredeemerFilterer, error) {
	contract, err := bindUsdpredeemer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UsdpredeemerFilterer{contract: contract}, nil
}

// bindUsdpredeemer binds a generic wrapper to an already deployed contract.
func bindUsdpredeemer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UsdpredeemerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Usdpredeemer *UsdpredeemerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Usdpredeemer.Contract.UsdpredeemerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Usdpredeemer *UsdpredeemerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Usdpredeemer.Contract.UsdpredeemerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Usdpredeemer *UsdpredeemerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Usdpredeemer.Contract.UsdpredeemerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Usdpredeemer *UsdpredeemerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Usdpredeemer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Usdpredeemer *UsdpredeemerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Usdpredeemer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Usdpredeemer *UsdpredeemerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Usdpredeemer.Contract.contract.Transact(opts, method, params...)
}

// QuoteRedemptionCurve is a free data retrieval call binding the contract method 0xd703a0cd.
//
// Solidity: function quoteRedemptionCurve(uint256 amount) view returns(address[] tokens, uint256[] amounts)
func (_Usdpredeemer *UsdpredeemerCaller) QuoteRedemptionCurve(opts *bind.CallOpts, amount *big.Int) (struct {
	Tokens  []common.Address
	Amounts []*big.Int
}, error) {
	var out []interface{}
	err := _Usdpredeemer.contract.Call(opts, &out, "quoteRedemptionCurve", amount)

	outstruct := new(struct {
		Tokens  []common.Address
		Amounts []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Tokens = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Amounts = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// QuoteRedemptionCurve is a free data retrieval call binding the contract method 0xd703a0cd.
//
// Solidity: function quoteRedemptionCurve(uint256 amount) view returns(address[] tokens, uint256[] amounts)
func (_Usdpredeemer *UsdpredeemerSession) QuoteRedemptionCurve(amount *big.Int) (struct {
	Tokens  []common.Address
	Amounts []*big.Int
}, error) {
	return _Usdpredeemer.Contract.QuoteRedemptionCurve(&_Usdpredeemer.CallOpts, amount)
}

// QuoteRedemptionCurve is a free data retrieval call binding the contract method 0xd703a0cd.
//
// Solidity: function quoteRedemptionCurve(uint256 amount) view returns(address[] tokens, uint256[] amounts)
func (_Usdpredeemer *UsdpredeemerCallerSession) QuoteRedemptionCurve(amount *big.Int) (struct {
	Tokens  []common.Address
	Amounts []*big.Int
}, error) {
	return _Usdpredeemer.Contract.QuoteRedemptionCurve(&_Usdpredeemer.CallOpts, amount)
}

// Redeem is a paid mutator transaction binding the contract method 0x815822c1.
//
// Solidity: function redeem(uint256 amount, address receiver, uint256 deadline, uint256[] minAmountOuts) returns(address[] tokens, uint256[] amounts)
func (_Usdpredeemer *UsdpredeemerTransactor) Redeem(opts *bind.TransactOpts, amount *big.Int, receiver common.Address, deadline *big.Int, minAmountOuts []*big.Int) (*types.Transaction, error) {
	return _Usdpredeemer.contract.Transact(opts, "redeem", amount, receiver, deadline, minAmountOuts)
}

// Redeem is a paid mutator transaction binding the contract method 0x815822c1.
//
// Solidity: function redeem(uint256 amount, address receiver, uint256 deadline, uint256[] minAmountOuts) returns(address[] tokens, uint256[] amounts)
func (_Usdpredeemer *UsdpredeemerSession) Redeem(amount *big.Int, receiver common.Address, deadline *big.Int, minAmountOuts []*big.Int) (*types.Transaction, error) {
	return _Usdpredeemer.Contract.Redeem(&_Usdpredeemer.TransactOpts, amount, receiver, deadline, minAmountOuts)
}

// Redeem is a paid mutator transaction binding the contract method 0x815822c1.
//
// Solidity: function redeem(uint256 amount, address receiver, uint256 deadline, uint256[] minAmountOuts) returns(address[] tokens, uint256[] amounts)
func (_Usdpredeemer *UsdpredeemerTransactorSession) Redeem(amount *big.Int, receiver common.Address, deadline *big.Int, minAmountOuts []*big.Int) (*types.Transaction, error) {
	return _Usdpredeemer.Contract.Redeem(&_Usdpredeemer.TransactOpts, amount, receiver, deadline, minAmountOuts)
}

// RedeemWithForfeit is a paid mutator transaction binding the contract method 0x2e7639bc.
//
// Solidity: function redeemWithForfeit(uint256 amount, address receiver, uint256 deadline, uint256[] minAmountOuts, address[] forfeitTokens) returns(address[] tokens, uint256[] amounts)
func (_Usdpredeemer *UsdpredeemerTransactor) RedeemWithForfeit(opts *bind.TransactOpts, amount *big.Int, receiver common.Address, deadline *big.Int, minAmountOuts []*big.Int, forfeitTokens []common.Address) (*types.Transaction, error) {
	return _Usdpredeemer.contract.Transact(opts, "redeemWithForfeit", amount, receiver, deadline, minAmountOuts, forfeitTokens)
}

// RedeemWithForfeit is a paid mutator transaction binding the contract method 0x2e7639bc.
//
// Solidity: function redeemWithForfeit(uint256 amount, address receiver, uint256 deadline, uint256[] minAmountOuts, address[] forfeitTokens) returns(address[] tokens, uint256[] amounts)
func (_Usdpredeemer *UsdpredeemerSession) RedeemWithForfeit(amount *big.Int, receiver common.Address, deadline *big.Int, minAmountOuts []*big.Int, forfeitTokens []common.Address) (*types.Transaction, error) {
	return _Usdpredeemer.Contract.RedeemWithForfeit(&_Usdpredeemer.TransactOpts, amount, receiver, deadline, minAmountOuts, forfeitTokens)
}

// RedeemWithForfeit is a paid mutator transaction binding the contract method 0x2e7639bc.
//
// Solidity: function redeemWithForfeit(uint256 amount, address receiver, uint256 deadline, uint256[] minAmountOuts, address[] forfeitTokens) returns(address[] tokens, uint256[] amounts)
func (_Usdpredeemer *UsdpredeemerTransactorSession) RedeemWithForfeit(amount *big.Int, receiver common.Address, deadline *big.Int, minAmountOuts []*big.Int, forfeitTokens []common.Address) (*types.Transaction, error) {
	return _Usdpredeemer.Contract.RedeemWithForfeit(&_Usdpredeemer.TransactOpts, amount, receiver, deadline, minAmountOuts, forfeitTokens)
}

// UpdateNormalizer is a paid mutator transaction binding the contract method 0xfd7daaf8.
//
// Solidity: function updateNormalizer(uint256 amount, bool increase) returns(uint256)
func (_Usdpredeemer *UsdpredeemerTransactor) UpdateNormalizer(opts *bind.TransactOpts, amount *big.Int, increase bool) (*types.Transaction, error) {
	return _Usdpredeemer.contract.Transact(opts, "updateNormalizer", amount, increase)
}

// UpdateNormalizer is a paid mutator transaction binding the contract method 0xfd7daaf8.
//
// Solidity: function updateNormalizer(uint256 amount, bool increase) returns(uint256)
func (_Usdpredeemer *UsdpredeemerSession) UpdateNormalizer(amount *big.Int, increase bool) (*types.Transaction, error) {
	return _Usdpredeemer.Contract.UpdateNormalizer(&_Usdpredeemer.TransactOpts, amount, increase)
}

// UpdateNormalizer is a paid mutator transaction binding the contract method 0xfd7daaf8.
//
// Solidity: function updateNormalizer(uint256 amount, bool increase) returns(uint256)
func (_Usdpredeemer *UsdpredeemerTransactorSession) UpdateNormalizer(amount *big.Int, increase bool) (*types.Transaction, error) {
	return _Usdpredeemer.Contract.UpdateNormalizer(&_Usdpredeemer.TransactOpts, amount, increase)
}

// UsdpredeemerNormalizerUpdatedIterator is returned from FilterNormalizerUpdated and is used to iterate over the raw logs and unpacked data for NormalizerUpdated events raised by the Usdpredeemer contract.
type UsdpredeemerNormalizerUpdatedIterator struct {
	Event *UsdpredeemerNormalizerUpdated // Event containing the contract specifics and raw log

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
func (it *UsdpredeemerNormalizerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsdpredeemerNormalizerUpdated)
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
		it.Event = new(UsdpredeemerNormalizerUpdated)
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
func (it *UsdpredeemerNormalizerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsdpredeemerNormalizerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsdpredeemerNormalizerUpdated represents a NormalizerUpdated event raised by the Usdpredeemer contract.
type UsdpredeemerNormalizerUpdated struct {
	NewNormalizerValue *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNormalizerUpdated is a free log retrieval operation binding the contract event 0xb4fde7403c001c7cdeafc0c3bb46e972111b7efac4e832398cbf9fabe1694ed9.
//
// Solidity: event NormalizerUpdated(uint256 newNormalizerValue)
func (_Usdpredeemer *UsdpredeemerFilterer) FilterNormalizerUpdated(opts *bind.FilterOpts) (*UsdpredeemerNormalizerUpdatedIterator, error) {

	logs, sub, err := _Usdpredeemer.contract.FilterLogs(opts, "NormalizerUpdated")
	if err != nil {
		return nil, err
	}
	return &UsdpredeemerNormalizerUpdatedIterator{contract: _Usdpredeemer.contract, event: "NormalizerUpdated", logs: logs, sub: sub}, nil
}

// WatchNormalizerUpdated is a free log subscription operation binding the contract event 0xb4fde7403c001c7cdeafc0c3bb46e972111b7efac4e832398cbf9fabe1694ed9.
//
// Solidity: event NormalizerUpdated(uint256 newNormalizerValue)
func (_Usdpredeemer *UsdpredeemerFilterer) WatchNormalizerUpdated(opts *bind.WatchOpts, sink chan<- *UsdpredeemerNormalizerUpdated) (event.Subscription, error) {

	logs, sub, err := _Usdpredeemer.contract.WatchLogs(opts, "NormalizerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsdpredeemerNormalizerUpdated)
				if err := _Usdpredeemer.contract.UnpackLog(event, "NormalizerUpdated", log); err != nil {
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

// ParseNormalizerUpdated is a log parse operation binding the contract event 0xb4fde7403c001c7cdeafc0c3bb46e972111b7efac4e832398cbf9fabe1694ed9.
//
// Solidity: event NormalizerUpdated(uint256 newNormalizerValue)
func (_Usdpredeemer *UsdpredeemerFilterer) ParseNormalizerUpdated(log types.Log) (*UsdpredeemerNormalizerUpdated, error) {
	event := new(UsdpredeemerNormalizerUpdated)
	if err := _Usdpredeemer.contract.UnpackLog(event, "NormalizerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UsdpredeemerRedeemedIterator is returned from FilterRedeemed and is used to iterate over the raw logs and unpacked data for Redeemed events raised by the Usdpredeemer contract.
type UsdpredeemerRedeemedIterator struct {
	Event *UsdpredeemerRedeemed // Event containing the contract specifics and raw log

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
func (it *UsdpredeemerRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsdpredeemerRedeemed)
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
		it.Event = new(UsdpredeemerRedeemed)
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
func (it *UsdpredeemerRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsdpredeemerRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsdpredeemerRedeemed represents a Redeemed event raised by the Usdpredeemer contract.
type UsdpredeemerRedeemed struct {
	Amount        *big.Int
	Tokens        []common.Address
	Amounts       []*big.Int
	ForfeitTokens []common.Address
	From          common.Address
	To            common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRedeemed is a free log retrieval operation binding the contract event 0x3998bcf315dae4953d294b3ca6e1d0074c249adc136b50d1ff99b04753c73e7f.
//
// Solidity: event Redeemed(uint256 amount, address[] tokens, uint256[] amounts, address[] forfeitTokens, address indexed from, address indexed to)
func (_Usdpredeemer *UsdpredeemerFilterer) FilterRedeemed(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*UsdpredeemerRedeemedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Usdpredeemer.contract.FilterLogs(opts, "Redeemed", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &UsdpredeemerRedeemedIterator{contract: _Usdpredeemer.contract, event: "Redeemed", logs: logs, sub: sub}, nil
}

// WatchRedeemed is a free log subscription operation binding the contract event 0x3998bcf315dae4953d294b3ca6e1d0074c249adc136b50d1ff99b04753c73e7f.
//
// Solidity: event Redeemed(uint256 amount, address[] tokens, uint256[] amounts, address[] forfeitTokens, address indexed from, address indexed to)
func (_Usdpredeemer *UsdpredeemerFilterer) WatchRedeemed(opts *bind.WatchOpts, sink chan<- *UsdpredeemerRedeemed, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Usdpredeemer.contract.WatchLogs(opts, "Redeemed", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsdpredeemerRedeemed)
				if err := _Usdpredeemer.contract.UnpackLog(event, "Redeemed", log); err != nil {
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

// ParseRedeemed is a log parse operation binding the contract event 0x3998bcf315dae4953d294b3ca6e1d0074c249adc136b50d1ff99b04753c73e7f.
//
// Solidity: event Redeemed(uint256 amount, address[] tokens, uint256[] amounts, address[] forfeitTokens, address indexed from, address indexed to)
func (_Usdpredeemer *UsdpredeemerFilterer) ParseRedeemed(log types.Log) (*UsdpredeemerRedeemed, error) {
	event := new(UsdpredeemerRedeemed)
	if err := _Usdpredeemer.contract.UnpackLog(event, "Redeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
