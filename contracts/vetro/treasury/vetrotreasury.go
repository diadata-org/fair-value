// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vetrotreasury

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

// VetrotreasuryMetaData contains all meta data concerning the Vetrotreasury contract.
var VetrotreasuryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"peggedToken_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"admin_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"}],\"name\":\"AccessControlInvalidDefaultAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AddToListFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AddressIsZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AssetMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceShouldBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotAuthorized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"DepositIsPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOraclePrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPriceTolerance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStalePeriod\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"InvalidTokenDecimals\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxWhitelistedTokensReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PeggedTokenMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"latestPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceUpperBound\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceLowerBound\",\"type\":\"uint256\"}],\"name\":\"PriceExceedTolerance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RemoveFromListFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReservedToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"StalePrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapperNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TreasuryNotMigrated\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"WithdrawIsPaused\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"AddedToWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminDelayChangeCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"effectSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminDelayChangeScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminTransferCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"acceptSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminTransferScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ExcessWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newTreasury\",\"type\":\"address\"}],\"name\":\"Migrated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RemovedFromWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"newValue\",\"type\":\"bool\"}],\"name\":\"SetDepositActive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"newValue\",\"type\":\"bool\"}],\"name\":\"SetWithdrawActive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"Swapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"Swept\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stalePeriod\",\"type\":\"uint256\"}],\"name\":\"UpdatedOracle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousPriceTolerance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPriceTolerance\",\"type\":\"uint256\"}],\"name\":\"UpdatedPriceTolerance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousSwapper\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newSwapper\",\"type\":\"address\"}],\"name\":\"UpdatedSwapper\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"KEEPER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAINTAINER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_BPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_STALE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_WHITELISTED_TOKENS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PEGGED_TOKEN\",\"outputs\":[{\"internalType\":\"contractIPeggedToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UMM_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stalePeriod_\",\"type\":\"uint256\"}],\"name\":\"addToWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"beginDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"}],\"name\":\"changeDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelayIncreaseWait\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_latestPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_unitPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"}],\"name\":\"harvest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"}],\"name\":\"isWhitelistedToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTreasury_\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceTolerance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"pull\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"push\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"}],\"name\":\"removeFromWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_reserve\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollbackDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"active_\",\"type\":\"bool\"}],\"name\":\"setDepositActive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"active_\",\"type\":\"bool\"}],\"name\":\"setWithdrawActive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut_\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapper\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromToken_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"}],\"name\":\"sweep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"tokenConfig\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stalePeriod\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"depositActive\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"withdrawActive\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newStalePeriod_\",\"type\":\"uint256\"}],\"name\":\"updateOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPriceTolerance_\",\"type\":\"uint256\"}],\"name\":\"updatePriceTolerance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"swapper_\",\"type\":\"address\"}],\"name\":\"updateSwapper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenReceiver_\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"}],\"name\":\"withdrawable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VetrotreasuryABI is the input ABI used to generate the binding from.
// Deprecated: Use VetrotreasuryMetaData.ABI instead.
var VetrotreasuryABI = VetrotreasuryMetaData.ABI

// Vetrotreasury is an auto generated Go binding around an Ethereum contract.
type Vetrotreasury struct {
	VetrotreasuryCaller     // Read-only binding to the contract
	VetrotreasuryTransactor // Write-only binding to the contract
	VetrotreasuryFilterer   // Log filterer for contract events
}

// VetrotreasuryCaller is an auto generated read-only Go binding around an Ethereum contract.
type VetrotreasuryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VetrotreasuryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VetrotreasuryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VetrotreasuryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VetrotreasuryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VetrotreasurySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VetrotreasurySession struct {
	Contract     *Vetrotreasury    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VetrotreasuryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VetrotreasuryCallerSession struct {
	Contract *VetrotreasuryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// VetrotreasuryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VetrotreasuryTransactorSession struct {
	Contract     *VetrotreasuryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// VetrotreasuryRaw is an auto generated low-level Go binding around an Ethereum contract.
type VetrotreasuryRaw struct {
	Contract *Vetrotreasury // Generic contract binding to access the raw methods on
}

// VetrotreasuryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VetrotreasuryCallerRaw struct {
	Contract *VetrotreasuryCaller // Generic read-only contract binding to access the raw methods on
}

// VetrotreasuryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VetrotreasuryTransactorRaw struct {
	Contract *VetrotreasuryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVetrotreasury creates a new instance of Vetrotreasury, bound to a specific deployed contract.
func NewVetrotreasury(address common.Address, backend bind.ContractBackend) (*Vetrotreasury, error) {
	contract, err := bindVetrotreasury(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vetrotreasury{VetrotreasuryCaller: VetrotreasuryCaller{contract: contract}, VetrotreasuryTransactor: VetrotreasuryTransactor{contract: contract}, VetrotreasuryFilterer: VetrotreasuryFilterer{contract: contract}}, nil
}

// NewVetrotreasuryCaller creates a new read-only instance of Vetrotreasury, bound to a specific deployed contract.
func NewVetrotreasuryCaller(address common.Address, caller bind.ContractCaller) (*VetrotreasuryCaller, error) {
	contract, err := bindVetrotreasury(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VetrotreasuryCaller{contract: contract}, nil
}

// NewVetrotreasuryTransactor creates a new write-only instance of Vetrotreasury, bound to a specific deployed contract.
func NewVetrotreasuryTransactor(address common.Address, transactor bind.ContractTransactor) (*VetrotreasuryTransactor, error) {
	contract, err := bindVetrotreasury(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VetrotreasuryTransactor{contract: contract}, nil
}

// NewVetrotreasuryFilterer creates a new log filterer instance of Vetrotreasury, bound to a specific deployed contract.
func NewVetrotreasuryFilterer(address common.Address, filterer bind.ContractFilterer) (*VetrotreasuryFilterer, error) {
	contract, err := bindVetrotreasury(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VetrotreasuryFilterer{contract: contract}, nil
}

// bindVetrotreasury binds a generic wrapper to an already deployed contract.
func bindVetrotreasury(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VetrotreasuryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vetrotreasury *VetrotreasuryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vetrotreasury.Contract.VetrotreasuryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vetrotreasury *VetrotreasuryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.VetrotreasuryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vetrotreasury *VetrotreasuryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.VetrotreasuryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vetrotreasury *VetrotreasuryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vetrotreasury.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vetrotreasury *VetrotreasuryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vetrotreasury *VetrotreasuryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Vetrotreasury *VetrotreasuryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vetrotreasury.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Vetrotreasury *VetrotreasurySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Vetrotreasury.Contract.DEFAULTADMINROLE(&_Vetrotreasury.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Vetrotreasury *VetrotreasuryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Vetrotreasury.Contract.DEFAULTADMINROLE(&_Vetrotreasury.CallOpts)
}

// KEEPERROLE is a free data retrieval call binding the contract method 0x364bc15a.
//
// Solidity: function KEEPER_ROLE() view returns(bytes32)
func (_Vetrotreasury *VetrotreasuryCaller) KEEPERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vetrotreasury.contract.Call(opts, &out, "KEEPER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// KEEPERROLE is a free data retrieval call binding the contract method 0x364bc15a.
//
// Solidity: function KEEPER_ROLE() view returns(bytes32)
func (_Vetrotreasury *VetrotreasurySession) KEEPERROLE() ([32]byte, error) {
	return _Vetrotreasury.Contract.KEEPERROLE(&_Vetrotreasury.CallOpts)
}

// KEEPERROLE is a free data retrieval call binding the contract method 0x364bc15a.
//
// Solidity: function KEEPER_ROLE() view returns(bytes32)
func (_Vetrotreasury *VetrotreasuryCallerSession) KEEPERROLE() ([32]byte, error) {
	return _Vetrotreasury.Contract.KEEPERROLE(&_Vetrotreasury.CallOpts)
}

// MAINTAINERROLE is a free data retrieval call binding the contract method 0xf8742254.
//
// Solidity: function MAINTAINER_ROLE() view returns(bytes32)
func (_Vetrotreasury *VetrotreasuryCaller) MAINTAINERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vetrotreasury.contract.Call(opts, &out, "MAINTAINER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MAINTAINERROLE is a free data retrieval call binding the contract method 0xf8742254.
//
// Solidity: function MAINTAINER_ROLE() view returns(bytes32)
func (_Vetrotreasury *VetrotreasurySession) MAINTAINERROLE() ([32]byte, error) {
	return _Vetrotreasury.Contract.MAINTAINERROLE(&_Vetrotreasury.CallOpts)
}

// MAINTAINERROLE is a free data retrieval call binding the contract method 0xf8742254.
//
// Solidity: function MAINTAINER_ROLE() view returns(bytes32)
func (_Vetrotreasury *VetrotreasuryCallerSession) MAINTAINERROLE() ([32]byte, error) {
	return _Vetrotreasury.Contract.MAINTAINERROLE(&_Vetrotreasury.CallOpts)
}

// MAXBPS is a free data retrieval call binding the contract method 0xfd967f47.
//
// Solidity: function MAX_BPS() view returns(uint256)
func (_Vetrotreasury *VetrotreasuryCaller) MAXBPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vetrotreasury.contract.Call(opts, &out, "MAX_BPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBPS is a free data retrieval call binding the contract method 0xfd967f47.
//
// Solidity: function MAX_BPS() view returns(uint256)
func (_Vetrotreasury *VetrotreasurySession) MAXBPS() (*big.Int, error) {
	return _Vetrotreasury.Contract.MAXBPS(&_Vetrotreasury.CallOpts)
}

// MAXBPS is a free data retrieval call binding the contract method 0xfd967f47.
//
// Solidity: function MAX_BPS() view returns(uint256)
func (_Vetrotreasury *VetrotreasuryCallerSession) MAXBPS() (*big.Int, error) {
	return _Vetrotreasury.Contract.MAXBPS(&_Vetrotreasury.CallOpts)
}

// MAXSTALEPERIOD is a free data retrieval call binding the contract method 0x064b914f.
//
// Solidity: function MAX_STALE_PERIOD() view returns(uint256)
func (_Vetrotreasury *VetrotreasuryCaller) MAXSTALEPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vetrotreasury.contract.Call(opts, &out, "MAX_STALE_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSTALEPERIOD is a free data retrieval call binding the contract method 0x064b914f.
//
// Solidity: function MAX_STALE_PERIOD() view returns(uint256)
func (_Vetrotreasury *VetrotreasurySession) MAXSTALEPERIOD() (*big.Int, error) {
	return _Vetrotreasury.Contract.MAXSTALEPERIOD(&_Vetrotreasury.CallOpts)
}

// MAXSTALEPERIOD is a free data retrieval call binding the contract method 0x064b914f.
//
// Solidity: function MAX_STALE_PERIOD() view returns(uint256)
func (_Vetrotreasury *VetrotreasuryCallerSession) MAXSTALEPERIOD() (*big.Int, error) {
	return _Vetrotreasury.Contract.MAXSTALEPERIOD(&_Vetrotreasury.CallOpts)
}

// MAXWHITELISTEDTOKENS is a free data retrieval call binding the contract method 0x7c08fe02.
//
// Solidity: function MAX_WHITELISTED_TOKENS() view returns(uint256)
func (_Vetrotreasury *VetrotreasuryCaller) MAXWHITELISTEDTOKENS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vetrotreasury.contract.Call(opts, &out, "MAX_WHITELISTED_TOKENS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXWHITELISTEDTOKENS is a free data retrieval call binding the contract method 0x7c08fe02.
//
// Solidity: function MAX_WHITELISTED_TOKENS() view returns(uint256)
func (_Vetrotreasury *VetrotreasurySession) MAXWHITELISTEDTOKENS() (*big.Int, error) {
	return _Vetrotreasury.Contract.MAXWHITELISTEDTOKENS(&_Vetrotreasury.CallOpts)
}

// MAXWHITELISTEDTOKENS is a free data retrieval call binding the contract method 0x7c08fe02.
//
// Solidity: function MAX_WHITELISTED_TOKENS() view returns(uint256)
func (_Vetrotreasury *VetrotreasuryCallerSession) MAXWHITELISTEDTOKENS() (*big.Int, error) {
	return _Vetrotreasury.Contract.MAXWHITELISTEDTOKENS(&_Vetrotreasury.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_Vetrotreasury *VetrotreasuryCaller) NAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Vetrotreasury.contract.Call(opts, &out, "NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_Vetrotreasury *VetrotreasurySession) NAME() (string, error) {
	return _Vetrotreasury.Contract.NAME(&_Vetrotreasury.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_Vetrotreasury *VetrotreasuryCallerSession) NAME() (string, error) {
	return _Vetrotreasury.Contract.NAME(&_Vetrotreasury.CallOpts)
}

// PEGGEDTOKEN is a free data retrieval call binding the contract method 0x4e485cf7.
//
// Solidity: function PEGGED_TOKEN() view returns(address)
func (_Vetrotreasury *VetrotreasuryCaller) PEGGEDTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vetrotreasury.contract.Call(opts, &out, "PEGGED_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PEGGEDTOKEN is a free data retrieval call binding the contract method 0x4e485cf7.
//
// Solidity: function PEGGED_TOKEN() view returns(address)
func (_Vetrotreasury *VetrotreasurySession) PEGGEDTOKEN() (common.Address, error) {
	return _Vetrotreasury.Contract.PEGGEDTOKEN(&_Vetrotreasury.CallOpts)
}

// PEGGEDTOKEN is a free data retrieval call binding the contract method 0x4e485cf7.
//
// Solidity: function PEGGED_TOKEN() view returns(address)
func (_Vetrotreasury *VetrotreasuryCallerSession) PEGGEDTOKEN() (common.Address, error) {
	return _Vetrotreasury.Contract.PEGGEDTOKEN(&_Vetrotreasury.CallOpts)
}

// UMMROLE is a free data retrieval call binding the contract method 0xf69516bd.
//
// Solidity: function UMM_ROLE() view returns(bytes32)
func (_Vetrotreasury *VetrotreasuryCaller) UMMROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vetrotreasury.contract.Call(opts, &out, "UMM_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UMMROLE is a free data retrieval call binding the contract method 0xf69516bd.
//
// Solidity: function UMM_ROLE() view returns(bytes32)
func (_Vetrotreasury *VetrotreasurySession) UMMROLE() ([32]byte, error) {
	return _Vetrotreasury.Contract.UMMROLE(&_Vetrotreasury.CallOpts)
}

// UMMROLE is a free data retrieval call binding the contract method 0xf69516bd.
//
// Solidity: function UMM_ROLE() view returns(bytes32)
func (_Vetrotreasury *VetrotreasuryCallerSession) UMMROLE() ([32]byte, error) {
	return _Vetrotreasury.Contract.UMMROLE(&_Vetrotreasury.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Vetrotreasury *VetrotreasuryCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Vetrotreasury.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Vetrotreasury *VetrotreasurySession) VERSION() (string, error) {
	return _Vetrotreasury.Contract.VERSION(&_Vetrotreasury.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Vetrotreasury *VetrotreasuryCallerSession) VERSION() (string, error) {
	return _Vetrotreasury.Contract.VERSION(&_Vetrotreasury.CallOpts)
}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_Vetrotreasury *VetrotreasuryCaller) DefaultAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vetrotreasury.contract.Call(opts, &out, "defaultAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_Vetrotreasury *VetrotreasurySession) DefaultAdmin() (common.Address, error) {
	return _Vetrotreasury.Contract.DefaultAdmin(&_Vetrotreasury.CallOpts)
}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_Vetrotreasury *VetrotreasuryCallerSession) DefaultAdmin() (common.Address, error) {
	return _Vetrotreasury.Contract.DefaultAdmin(&_Vetrotreasury.CallOpts)
}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_Vetrotreasury *VetrotreasuryCaller) DefaultAdminDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vetrotreasury.contract.Call(opts, &out, "defaultAdminDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_Vetrotreasury *VetrotreasurySession) DefaultAdminDelay() (*big.Int, error) {
	return _Vetrotreasury.Contract.DefaultAdminDelay(&_Vetrotreasury.CallOpts)
}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_Vetrotreasury *VetrotreasuryCallerSession) DefaultAdminDelay() (*big.Int, error) {
	return _Vetrotreasury.Contract.DefaultAdminDelay(&_Vetrotreasury.CallOpts)
}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_Vetrotreasury *VetrotreasuryCaller) DefaultAdminDelayIncreaseWait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vetrotreasury.contract.Call(opts, &out, "defaultAdminDelayIncreaseWait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_Vetrotreasury *VetrotreasurySession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _Vetrotreasury.Contract.DefaultAdminDelayIncreaseWait(&_Vetrotreasury.CallOpts)
}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_Vetrotreasury *VetrotreasuryCallerSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _Vetrotreasury.Contract.DefaultAdminDelayIncreaseWait(&_Vetrotreasury.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_Vetrotreasury *VetrotreasuryCaller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vetrotreasury.contract.Call(opts, &out, "gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_Vetrotreasury *VetrotreasurySession) Gateway() (common.Address, error) {
	return _Vetrotreasury.Contract.Gateway(&_Vetrotreasury.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_Vetrotreasury *VetrotreasuryCallerSession) Gateway() (common.Address, error) {
	return _Vetrotreasury.Contract.Gateway(&_Vetrotreasury.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token_) view returns(uint256 _latestPrice, uint256 _unitPrice)
func (_Vetrotreasury *VetrotreasuryCaller) GetPrice(opts *bind.CallOpts, token_ common.Address) (struct {
	LatestPrice *big.Int
	UnitPrice   *big.Int
}, error) {
	var out []interface{}
	err := _Vetrotreasury.contract.Call(opts, &out, "getPrice", token_)

	outstruct := new(struct {
		LatestPrice *big.Int
		UnitPrice   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LatestPrice = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.UnitPrice = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token_) view returns(uint256 _latestPrice, uint256 _unitPrice)
func (_Vetrotreasury *VetrotreasurySession) GetPrice(token_ common.Address) (struct {
	LatestPrice *big.Int
	UnitPrice   *big.Int
}, error) {
	return _Vetrotreasury.Contract.GetPrice(&_Vetrotreasury.CallOpts, token_)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token_) view returns(uint256 _latestPrice, uint256 _unitPrice)
func (_Vetrotreasury *VetrotreasuryCallerSession) GetPrice(token_ common.Address) (struct {
	LatestPrice *big.Int
	UnitPrice   *big.Int
}, error) {
	return _Vetrotreasury.Contract.GetPrice(&_Vetrotreasury.CallOpts, token_)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Vetrotreasury *VetrotreasuryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Vetrotreasury.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Vetrotreasury *VetrotreasurySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Vetrotreasury.Contract.GetRoleAdmin(&_Vetrotreasury.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Vetrotreasury *VetrotreasuryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Vetrotreasury.Contract.GetRoleAdmin(&_Vetrotreasury.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Vetrotreasury *VetrotreasuryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Vetrotreasury.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Vetrotreasury *VetrotreasurySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Vetrotreasury.Contract.HasRole(&_Vetrotreasury.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Vetrotreasury *VetrotreasuryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Vetrotreasury.Contract.HasRole(&_Vetrotreasury.CallOpts, role, account)
}

// IsWhitelistedToken is a free data retrieval call binding the contract method 0xab37f486.
//
// Solidity: function isWhitelistedToken(address token_) view returns(bool)
func (_Vetrotreasury *VetrotreasuryCaller) IsWhitelistedToken(opts *bind.CallOpts, token_ common.Address) (bool, error) {
	var out []interface{}
	err := _Vetrotreasury.contract.Call(opts, &out, "isWhitelistedToken", token_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelistedToken is a free data retrieval call binding the contract method 0xab37f486.
//
// Solidity: function isWhitelistedToken(address token_) view returns(bool)
func (_Vetrotreasury *VetrotreasurySession) IsWhitelistedToken(token_ common.Address) (bool, error) {
	return _Vetrotreasury.Contract.IsWhitelistedToken(&_Vetrotreasury.CallOpts, token_)
}

// IsWhitelistedToken is a free data retrieval call binding the contract method 0xab37f486.
//
// Solidity: function isWhitelistedToken(address token_) view returns(bool)
func (_Vetrotreasury *VetrotreasuryCallerSession) IsWhitelistedToken(token_ common.Address) (bool, error) {
	return _Vetrotreasury.Contract.IsWhitelistedToken(&_Vetrotreasury.CallOpts, token_)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vetrotreasury *VetrotreasuryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vetrotreasury.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vetrotreasury *VetrotreasurySession) Owner() (common.Address, error) {
	return _Vetrotreasury.Contract.Owner(&_Vetrotreasury.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vetrotreasury *VetrotreasuryCallerSession) Owner() (common.Address, error) {
	return _Vetrotreasury.Contract.Owner(&_Vetrotreasury.CallOpts)
}

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_Vetrotreasury *VetrotreasuryCaller) PendingDefaultAdmin(opts *bind.CallOpts) (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	var out []interface{}
	err := _Vetrotreasury.contract.Call(opts, &out, "pendingDefaultAdmin")

	outstruct := new(struct {
		NewAdmin common.Address
		Schedule *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewAdmin = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_Vetrotreasury *VetrotreasurySession) PendingDefaultAdmin() (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	return _Vetrotreasury.Contract.PendingDefaultAdmin(&_Vetrotreasury.CallOpts)
}

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_Vetrotreasury *VetrotreasuryCallerSession) PendingDefaultAdmin() (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	return _Vetrotreasury.Contract.PendingDefaultAdmin(&_Vetrotreasury.CallOpts)
}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_Vetrotreasury *VetrotreasuryCaller) PendingDefaultAdminDelay(opts *bind.CallOpts) (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	var out []interface{}
	err := _Vetrotreasury.contract.Call(opts, &out, "pendingDefaultAdminDelay")

	outstruct := new(struct {
		NewDelay *big.Int
		Schedule *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewDelay = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_Vetrotreasury *VetrotreasurySession) PendingDefaultAdminDelay() (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	return _Vetrotreasury.Contract.PendingDefaultAdminDelay(&_Vetrotreasury.CallOpts)
}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_Vetrotreasury *VetrotreasuryCallerSession) PendingDefaultAdminDelay() (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	return _Vetrotreasury.Contract.PendingDefaultAdminDelay(&_Vetrotreasury.CallOpts)
}

// PriceTolerance is a free data retrieval call binding the contract method 0x59011cd1.
//
// Solidity: function priceTolerance() view returns(uint256)
func (_Vetrotreasury *VetrotreasuryCaller) PriceTolerance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vetrotreasury.contract.Call(opts, &out, "priceTolerance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceTolerance is a free data retrieval call binding the contract method 0x59011cd1.
//
// Solidity: function priceTolerance() view returns(uint256)
func (_Vetrotreasury *VetrotreasurySession) PriceTolerance() (*big.Int, error) {
	return _Vetrotreasury.Contract.PriceTolerance(&_Vetrotreasury.CallOpts)
}

// PriceTolerance is a free data retrieval call binding the contract method 0x59011cd1.
//
// Solidity: function priceTolerance() view returns(uint256)
func (_Vetrotreasury *VetrotreasuryCallerSession) PriceTolerance() (*big.Int, error) {
	return _Vetrotreasury.Contract.PriceTolerance(&_Vetrotreasury.CallOpts)
}

// Reserve is a free data retrieval call binding the contract method 0xcd3293de.
//
// Solidity: function reserve() view returns(uint256 _reserve)
func (_Vetrotreasury *VetrotreasuryCaller) Reserve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vetrotreasury.contract.Call(opts, &out, "reserve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Reserve is a free data retrieval call binding the contract method 0xcd3293de.
//
// Solidity: function reserve() view returns(uint256 _reserve)
func (_Vetrotreasury *VetrotreasurySession) Reserve() (*big.Int, error) {
	return _Vetrotreasury.Contract.Reserve(&_Vetrotreasury.CallOpts)
}

// Reserve is a free data retrieval call binding the contract method 0xcd3293de.
//
// Solidity: function reserve() view returns(uint256 _reserve)
func (_Vetrotreasury *VetrotreasuryCallerSession) Reserve() (*big.Int, error) {
	return _Vetrotreasury.Contract.Reserve(&_Vetrotreasury.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Vetrotreasury *VetrotreasuryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Vetrotreasury.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Vetrotreasury *VetrotreasurySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Vetrotreasury.Contract.SupportsInterface(&_Vetrotreasury.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Vetrotreasury *VetrotreasuryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Vetrotreasury.Contract.SupportsInterface(&_Vetrotreasury.CallOpts, interfaceId)
}

// Swapper is a free data retrieval call binding the contract method 0x2b3297f9.
//
// Solidity: function swapper() view returns(address)
func (_Vetrotreasury *VetrotreasuryCaller) Swapper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vetrotreasury.contract.Call(opts, &out, "swapper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Swapper is a free data retrieval call binding the contract method 0x2b3297f9.
//
// Solidity: function swapper() view returns(address)
func (_Vetrotreasury *VetrotreasurySession) Swapper() (common.Address, error) {
	return _Vetrotreasury.Contract.Swapper(&_Vetrotreasury.CallOpts)
}

// Swapper is a free data retrieval call binding the contract method 0x2b3297f9.
//
// Solidity: function swapper() view returns(address)
func (_Vetrotreasury *VetrotreasuryCallerSession) Swapper() (common.Address, error) {
	return _Vetrotreasury.Contract.Swapper(&_Vetrotreasury.CallOpts)
}

// TokenConfig is a free data retrieval call binding the contract method 0xfe136c4e.
//
// Solidity: function tokenConfig(address token) view returns(address vault, address oracle, uint256 stalePeriod, bool depositActive, bool withdrawActive, uint8 decimals)
func (_Vetrotreasury *VetrotreasuryCaller) TokenConfig(opts *bind.CallOpts, token common.Address) (struct {
	Vault          common.Address
	Oracle         common.Address
	StalePeriod    *big.Int
	DepositActive  bool
	WithdrawActive bool
	Decimals       uint8
}, error) {
	var out []interface{}
	err := _Vetrotreasury.contract.Call(opts, &out, "tokenConfig", token)

	outstruct := new(struct {
		Vault          common.Address
		Oracle         common.Address
		StalePeriod    *big.Int
		DepositActive  bool
		WithdrawActive bool
		Decimals       uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Vault = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Oracle = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.StalePeriod = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.DepositActive = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.WithdrawActive = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.Decimals = *abi.ConvertType(out[5], new(uint8)).(*uint8)

	return *outstruct, err

}

// TokenConfig is a free data retrieval call binding the contract method 0xfe136c4e.
//
// Solidity: function tokenConfig(address token) view returns(address vault, address oracle, uint256 stalePeriod, bool depositActive, bool withdrawActive, uint8 decimals)
func (_Vetrotreasury *VetrotreasurySession) TokenConfig(token common.Address) (struct {
	Vault          common.Address
	Oracle         common.Address
	StalePeriod    *big.Int
	DepositActive  bool
	WithdrawActive bool
	Decimals       uint8
}, error) {
	return _Vetrotreasury.Contract.TokenConfig(&_Vetrotreasury.CallOpts, token)
}

// TokenConfig is a free data retrieval call binding the contract method 0xfe136c4e.
//
// Solidity: function tokenConfig(address token) view returns(address vault, address oracle, uint256 stalePeriod, bool depositActive, bool withdrawActive, uint8 decimals)
func (_Vetrotreasury *VetrotreasuryCallerSession) TokenConfig(token common.Address) (struct {
	Vault          common.Address
	Oracle         common.Address
	StalePeriod    *big.Int
	DepositActive  bool
	WithdrawActive bool
	Decimals       uint8
}, error) {
	return _Vetrotreasury.Contract.TokenConfig(&_Vetrotreasury.CallOpts, token)
}

// WhitelistedTokens is a free data retrieval call binding the contract method 0x5e1762a0.
//
// Solidity: function whitelistedTokens() view returns(address[])
func (_Vetrotreasury *VetrotreasuryCaller) WhitelistedTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Vetrotreasury.contract.Call(opts, &out, "whitelistedTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// WhitelistedTokens is a free data retrieval call binding the contract method 0x5e1762a0.
//
// Solidity: function whitelistedTokens() view returns(address[])
func (_Vetrotreasury *VetrotreasurySession) WhitelistedTokens() ([]common.Address, error) {
	return _Vetrotreasury.Contract.WhitelistedTokens(&_Vetrotreasury.CallOpts)
}

// WhitelistedTokens is a free data retrieval call binding the contract method 0x5e1762a0.
//
// Solidity: function whitelistedTokens() view returns(address[])
func (_Vetrotreasury *VetrotreasuryCallerSession) WhitelistedTokens() ([]common.Address, error) {
	return _Vetrotreasury.Contract.WhitelistedTokens(&_Vetrotreasury.CallOpts)
}

// Withdrawable is a free data retrieval call binding the contract method 0xce513b6f.
//
// Solidity: function withdrawable(address token_) view returns(uint256)
func (_Vetrotreasury *VetrotreasuryCaller) Withdrawable(opts *bind.CallOpts, token_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vetrotreasury.contract.Call(opts, &out, "withdrawable", token_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Withdrawable is a free data retrieval call binding the contract method 0xce513b6f.
//
// Solidity: function withdrawable(address token_) view returns(uint256)
func (_Vetrotreasury *VetrotreasurySession) Withdrawable(token_ common.Address) (*big.Int, error) {
	return _Vetrotreasury.Contract.Withdrawable(&_Vetrotreasury.CallOpts, token_)
}

// Withdrawable is a free data retrieval call binding the contract method 0xce513b6f.
//
// Solidity: function withdrawable(address token_) view returns(uint256)
func (_Vetrotreasury *VetrotreasuryCallerSession) Withdrawable(token_ common.Address) (*big.Int, error) {
	return _Vetrotreasury.Contract.Withdrawable(&_Vetrotreasury.CallOpts, token_)
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_Vetrotreasury *VetrotreasuryTransactor) AcceptDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vetrotreasury.contract.Transact(opts, "acceptDefaultAdminTransfer")
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_Vetrotreasury *VetrotreasurySession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _Vetrotreasury.Contract.AcceptDefaultAdminTransfer(&_Vetrotreasury.TransactOpts)
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_Vetrotreasury *VetrotreasuryTransactorSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _Vetrotreasury.Contract.AcceptDefaultAdminTransfer(&_Vetrotreasury.TransactOpts)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xcd110f11.
//
// Solidity: function addToWhitelist(address token_, address vault_, address oracle_, uint256 stalePeriod_) returns()
func (_Vetrotreasury *VetrotreasuryTransactor) AddToWhitelist(opts *bind.TransactOpts, token_ common.Address, vault_ common.Address, oracle_ common.Address, stalePeriod_ *big.Int) (*types.Transaction, error) {
	return _Vetrotreasury.contract.Transact(opts, "addToWhitelist", token_, vault_, oracle_, stalePeriod_)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xcd110f11.
//
// Solidity: function addToWhitelist(address token_, address vault_, address oracle_, uint256 stalePeriod_) returns()
func (_Vetrotreasury *VetrotreasurySession) AddToWhitelist(token_ common.Address, vault_ common.Address, oracle_ common.Address, stalePeriod_ *big.Int) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.AddToWhitelist(&_Vetrotreasury.TransactOpts, token_, vault_, oracle_, stalePeriod_)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xcd110f11.
//
// Solidity: function addToWhitelist(address token_, address vault_, address oracle_, uint256 stalePeriod_) returns()
func (_Vetrotreasury *VetrotreasuryTransactorSession) AddToWhitelist(token_ common.Address, vault_ common.Address, oracle_ common.Address, stalePeriod_ *big.Int) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.AddToWhitelist(&_Vetrotreasury.TransactOpts, token_, vault_, oracle_, stalePeriod_)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_Vetrotreasury *VetrotreasuryTransactor) BeginDefaultAdminTransfer(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _Vetrotreasury.contract.Transact(opts, "beginDefaultAdminTransfer", newAdmin)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_Vetrotreasury *VetrotreasurySession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.BeginDefaultAdminTransfer(&_Vetrotreasury.TransactOpts, newAdmin)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_Vetrotreasury *VetrotreasuryTransactorSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.BeginDefaultAdminTransfer(&_Vetrotreasury.TransactOpts, newAdmin)
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_Vetrotreasury *VetrotreasuryTransactor) CancelDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vetrotreasury.contract.Transact(opts, "cancelDefaultAdminTransfer")
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_Vetrotreasury *VetrotreasurySession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _Vetrotreasury.Contract.CancelDefaultAdminTransfer(&_Vetrotreasury.TransactOpts)
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_Vetrotreasury *VetrotreasuryTransactorSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _Vetrotreasury.Contract.CancelDefaultAdminTransfer(&_Vetrotreasury.TransactOpts)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_Vetrotreasury *VetrotreasuryTransactor) ChangeDefaultAdminDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _Vetrotreasury.contract.Transact(opts, "changeDefaultAdminDelay", newDelay)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_Vetrotreasury *VetrotreasurySession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.ChangeDefaultAdminDelay(&_Vetrotreasury.TransactOpts, newDelay)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_Vetrotreasury *VetrotreasuryTransactorSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.ChangeDefaultAdminDelay(&_Vetrotreasury.TransactOpts, newDelay)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token_, uint256 amount_) returns()
func (_Vetrotreasury *VetrotreasuryTransactor) Deposit(opts *bind.TransactOpts, token_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Vetrotreasury.contract.Transact(opts, "deposit", token_, amount_)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token_, uint256 amount_) returns()
func (_Vetrotreasury *VetrotreasurySession) Deposit(token_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.Deposit(&_Vetrotreasury.TransactOpts, token_, amount_)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token_, uint256 amount_) returns()
func (_Vetrotreasury *VetrotreasuryTransactorSession) Deposit(token_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.Deposit(&_Vetrotreasury.TransactOpts, token_, amount_)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Vetrotreasury *VetrotreasuryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vetrotreasury.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Vetrotreasury *VetrotreasurySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.GrantRole(&_Vetrotreasury.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Vetrotreasury *VetrotreasuryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.GrantRole(&_Vetrotreasury.TransactOpts, role, account)
}

// Harvest is a paid mutator transaction binding the contract method 0x66cc1857.
//
// Solidity: function harvest(address token_, address receiver_) returns(uint256)
func (_Vetrotreasury *VetrotreasuryTransactor) Harvest(opts *bind.TransactOpts, token_ common.Address, receiver_ common.Address) (*types.Transaction, error) {
	return _Vetrotreasury.contract.Transact(opts, "harvest", token_, receiver_)
}

// Harvest is a paid mutator transaction binding the contract method 0x66cc1857.
//
// Solidity: function harvest(address token_, address receiver_) returns(uint256)
func (_Vetrotreasury *VetrotreasurySession) Harvest(token_ common.Address, receiver_ common.Address) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.Harvest(&_Vetrotreasury.TransactOpts, token_, receiver_)
}

// Harvest is a paid mutator transaction binding the contract method 0x66cc1857.
//
// Solidity: function harvest(address token_, address receiver_) returns(uint256)
func (_Vetrotreasury *VetrotreasuryTransactorSession) Harvest(token_ common.Address, receiver_ common.Address) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.Harvest(&_Vetrotreasury.TransactOpts, token_, receiver_)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address newTreasury_) returns()
func (_Vetrotreasury *VetrotreasuryTransactor) Migrate(opts *bind.TransactOpts, newTreasury_ common.Address) (*types.Transaction, error) {
	return _Vetrotreasury.contract.Transact(opts, "migrate", newTreasury_)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address newTreasury_) returns()
func (_Vetrotreasury *VetrotreasurySession) Migrate(newTreasury_ common.Address) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.Migrate(&_Vetrotreasury.TransactOpts, newTreasury_)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address newTreasury_) returns()
func (_Vetrotreasury *VetrotreasuryTransactorSession) Migrate(newTreasury_ common.Address) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.Migrate(&_Vetrotreasury.TransactOpts, newTreasury_)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address token_, uint256 amount_) returns()
func (_Vetrotreasury *VetrotreasuryTransactor) Pull(opts *bind.TransactOpts, token_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Vetrotreasury.contract.Transact(opts, "pull", token_, amount_)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address token_, uint256 amount_) returns()
func (_Vetrotreasury *VetrotreasurySession) Pull(token_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.Pull(&_Vetrotreasury.TransactOpts, token_, amount_)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address token_, uint256 amount_) returns()
func (_Vetrotreasury *VetrotreasuryTransactorSession) Pull(token_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.Pull(&_Vetrotreasury.TransactOpts, token_, amount_)
}

// Push is a paid mutator transaction binding the contract method 0xb753a98c.
//
// Solidity: function push(address token_, uint256 amount_) returns()
func (_Vetrotreasury *VetrotreasuryTransactor) Push(opts *bind.TransactOpts, token_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Vetrotreasury.contract.Transact(opts, "push", token_, amount_)
}

// Push is a paid mutator transaction binding the contract method 0xb753a98c.
//
// Solidity: function push(address token_, uint256 amount_) returns()
func (_Vetrotreasury *VetrotreasurySession) Push(token_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.Push(&_Vetrotreasury.TransactOpts, token_, amount_)
}

// Push is a paid mutator transaction binding the contract method 0xb753a98c.
//
// Solidity: function push(address token_, uint256 amount_) returns()
func (_Vetrotreasury *VetrotreasuryTransactorSession) Push(token_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.Push(&_Vetrotreasury.TransactOpts, token_, amount_)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(address token_) returns()
func (_Vetrotreasury *VetrotreasuryTransactor) RemoveFromWhitelist(opts *bind.TransactOpts, token_ common.Address) (*types.Transaction, error) {
	return _Vetrotreasury.contract.Transact(opts, "removeFromWhitelist", token_)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(address token_) returns()
func (_Vetrotreasury *VetrotreasurySession) RemoveFromWhitelist(token_ common.Address) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.RemoveFromWhitelist(&_Vetrotreasury.TransactOpts, token_)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(address token_) returns()
func (_Vetrotreasury *VetrotreasuryTransactorSession) RemoveFromWhitelist(token_ common.Address) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.RemoveFromWhitelist(&_Vetrotreasury.TransactOpts, token_)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Vetrotreasury *VetrotreasuryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vetrotreasury.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Vetrotreasury *VetrotreasurySession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.RenounceRole(&_Vetrotreasury.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Vetrotreasury *VetrotreasuryTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.RenounceRole(&_Vetrotreasury.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Vetrotreasury *VetrotreasuryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vetrotreasury.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Vetrotreasury *VetrotreasurySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.RevokeRole(&_Vetrotreasury.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Vetrotreasury *VetrotreasuryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.RevokeRole(&_Vetrotreasury.TransactOpts, role, account)
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_Vetrotreasury *VetrotreasuryTransactor) RollbackDefaultAdminDelay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vetrotreasury.contract.Transact(opts, "rollbackDefaultAdminDelay")
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_Vetrotreasury *VetrotreasurySession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _Vetrotreasury.Contract.RollbackDefaultAdminDelay(&_Vetrotreasury.TransactOpts)
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_Vetrotreasury *VetrotreasuryTransactorSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _Vetrotreasury.Contract.RollbackDefaultAdminDelay(&_Vetrotreasury.TransactOpts)
}

// SetDepositActive is a paid mutator transaction binding the contract method 0xb5034eb2.
//
// Solidity: function setDepositActive(address token_, bool active_) returns()
func (_Vetrotreasury *VetrotreasuryTransactor) SetDepositActive(opts *bind.TransactOpts, token_ common.Address, active_ bool) (*types.Transaction, error) {
	return _Vetrotreasury.contract.Transact(opts, "setDepositActive", token_, active_)
}

// SetDepositActive is a paid mutator transaction binding the contract method 0xb5034eb2.
//
// Solidity: function setDepositActive(address token_, bool active_) returns()
func (_Vetrotreasury *VetrotreasurySession) SetDepositActive(token_ common.Address, active_ bool) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.SetDepositActive(&_Vetrotreasury.TransactOpts, token_, active_)
}

// SetDepositActive is a paid mutator transaction binding the contract method 0xb5034eb2.
//
// Solidity: function setDepositActive(address token_, bool active_) returns()
func (_Vetrotreasury *VetrotreasuryTransactorSession) SetDepositActive(token_ common.Address, active_ bool) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.SetDepositActive(&_Vetrotreasury.TransactOpts, token_, active_)
}

// SetWithdrawActive is a paid mutator transaction binding the contract method 0x5bd34607.
//
// Solidity: function setWithdrawActive(address token_, bool active_) returns()
func (_Vetrotreasury *VetrotreasuryTransactor) SetWithdrawActive(opts *bind.TransactOpts, token_ common.Address, active_ bool) (*types.Transaction, error) {
	return _Vetrotreasury.contract.Transact(opts, "setWithdrawActive", token_, active_)
}

// SetWithdrawActive is a paid mutator transaction binding the contract method 0x5bd34607.
//
// Solidity: function setWithdrawActive(address token_, bool active_) returns()
func (_Vetrotreasury *VetrotreasurySession) SetWithdrawActive(token_ common.Address, active_ bool) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.SetWithdrawActive(&_Vetrotreasury.TransactOpts, token_, active_)
}

// SetWithdrawActive is a paid mutator transaction binding the contract method 0x5bd34607.
//
// Solidity: function setWithdrawActive(address token_, bool active_) returns()
func (_Vetrotreasury *VetrotreasuryTransactorSession) SetWithdrawActive(token_ common.Address, active_ bool) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.SetWithdrawActive(&_Vetrotreasury.TransactOpts, token_, active_)
}

// Swap is a paid mutator transaction binding the contract method 0xfe029156.
//
// Solidity: function swap(address tokenIn_, address tokenOut_, uint256 amountIn_, uint256 minAmountOut_) returns(uint256)
func (_Vetrotreasury *VetrotreasuryTransactor) Swap(opts *bind.TransactOpts, tokenIn_ common.Address, tokenOut_ common.Address, amountIn_ *big.Int, minAmountOut_ *big.Int) (*types.Transaction, error) {
	return _Vetrotreasury.contract.Transact(opts, "swap", tokenIn_, tokenOut_, amountIn_, minAmountOut_)
}

// Swap is a paid mutator transaction binding the contract method 0xfe029156.
//
// Solidity: function swap(address tokenIn_, address tokenOut_, uint256 amountIn_, uint256 minAmountOut_) returns(uint256)
func (_Vetrotreasury *VetrotreasurySession) Swap(tokenIn_ common.Address, tokenOut_ common.Address, amountIn_ *big.Int, minAmountOut_ *big.Int) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.Swap(&_Vetrotreasury.TransactOpts, tokenIn_, tokenOut_, amountIn_, minAmountOut_)
}

// Swap is a paid mutator transaction binding the contract method 0xfe029156.
//
// Solidity: function swap(address tokenIn_, address tokenOut_, uint256 amountIn_, uint256 minAmountOut_) returns(uint256)
func (_Vetrotreasury *VetrotreasuryTransactorSession) Swap(tokenIn_ common.Address, tokenOut_ common.Address, amountIn_ *big.Int, minAmountOut_ *big.Int) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.Swap(&_Vetrotreasury.TransactOpts, tokenIn_, tokenOut_, amountIn_, minAmountOut_)
}

// Sweep is a paid mutator transaction binding the contract method 0xb8dc491b.
//
// Solidity: function sweep(address fromToken_, address receiver_) returns()
func (_Vetrotreasury *VetrotreasuryTransactor) Sweep(opts *bind.TransactOpts, fromToken_ common.Address, receiver_ common.Address) (*types.Transaction, error) {
	return _Vetrotreasury.contract.Transact(opts, "sweep", fromToken_, receiver_)
}

// Sweep is a paid mutator transaction binding the contract method 0xb8dc491b.
//
// Solidity: function sweep(address fromToken_, address receiver_) returns()
func (_Vetrotreasury *VetrotreasurySession) Sweep(fromToken_ common.Address, receiver_ common.Address) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.Sweep(&_Vetrotreasury.TransactOpts, fromToken_, receiver_)
}

// Sweep is a paid mutator transaction binding the contract method 0xb8dc491b.
//
// Solidity: function sweep(address fromToken_, address receiver_) returns()
func (_Vetrotreasury *VetrotreasuryTransactorSession) Sweep(fromToken_ common.Address, receiver_ common.Address) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.Sweep(&_Vetrotreasury.TransactOpts, fromToken_, receiver_)
}

// UpdateOracle is a paid mutator transaction binding the contract method 0x49aafe82.
//
// Solidity: function updateOracle(address token_, address oracle_, uint256 newStalePeriod_) returns()
func (_Vetrotreasury *VetrotreasuryTransactor) UpdateOracle(opts *bind.TransactOpts, token_ common.Address, oracle_ common.Address, newStalePeriod_ *big.Int) (*types.Transaction, error) {
	return _Vetrotreasury.contract.Transact(opts, "updateOracle", token_, oracle_, newStalePeriod_)
}

// UpdateOracle is a paid mutator transaction binding the contract method 0x49aafe82.
//
// Solidity: function updateOracle(address token_, address oracle_, uint256 newStalePeriod_) returns()
func (_Vetrotreasury *VetrotreasurySession) UpdateOracle(token_ common.Address, oracle_ common.Address, newStalePeriod_ *big.Int) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.UpdateOracle(&_Vetrotreasury.TransactOpts, token_, oracle_, newStalePeriod_)
}

// UpdateOracle is a paid mutator transaction binding the contract method 0x49aafe82.
//
// Solidity: function updateOracle(address token_, address oracle_, uint256 newStalePeriod_) returns()
func (_Vetrotreasury *VetrotreasuryTransactorSession) UpdateOracle(token_ common.Address, oracle_ common.Address, newStalePeriod_ *big.Int) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.UpdateOracle(&_Vetrotreasury.TransactOpts, token_, oracle_, newStalePeriod_)
}

// UpdatePriceTolerance is a paid mutator transaction binding the contract method 0xc760af5c.
//
// Solidity: function updatePriceTolerance(uint256 newPriceTolerance_) returns()
func (_Vetrotreasury *VetrotreasuryTransactor) UpdatePriceTolerance(opts *bind.TransactOpts, newPriceTolerance_ *big.Int) (*types.Transaction, error) {
	return _Vetrotreasury.contract.Transact(opts, "updatePriceTolerance", newPriceTolerance_)
}

// UpdatePriceTolerance is a paid mutator transaction binding the contract method 0xc760af5c.
//
// Solidity: function updatePriceTolerance(uint256 newPriceTolerance_) returns()
func (_Vetrotreasury *VetrotreasurySession) UpdatePriceTolerance(newPriceTolerance_ *big.Int) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.UpdatePriceTolerance(&_Vetrotreasury.TransactOpts, newPriceTolerance_)
}

// UpdatePriceTolerance is a paid mutator transaction binding the contract method 0xc760af5c.
//
// Solidity: function updatePriceTolerance(uint256 newPriceTolerance_) returns()
func (_Vetrotreasury *VetrotreasuryTransactorSession) UpdatePriceTolerance(newPriceTolerance_ *big.Int) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.UpdatePriceTolerance(&_Vetrotreasury.TransactOpts, newPriceTolerance_)
}

// UpdateSwapper is a paid mutator transaction binding the contract method 0xd3033c39.
//
// Solidity: function updateSwapper(address swapper_) returns()
func (_Vetrotreasury *VetrotreasuryTransactor) UpdateSwapper(opts *bind.TransactOpts, swapper_ common.Address) (*types.Transaction, error) {
	return _Vetrotreasury.contract.Transact(opts, "updateSwapper", swapper_)
}

// UpdateSwapper is a paid mutator transaction binding the contract method 0xd3033c39.
//
// Solidity: function updateSwapper(address swapper_) returns()
func (_Vetrotreasury *VetrotreasurySession) UpdateSwapper(swapper_ common.Address) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.UpdateSwapper(&_Vetrotreasury.TransactOpts, swapper_)
}

// UpdateSwapper is a paid mutator transaction binding the contract method 0xd3033c39.
//
// Solidity: function updateSwapper(address swapper_) returns()
func (_Vetrotreasury *VetrotreasuryTransactorSession) UpdateSwapper(swapper_ common.Address) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.UpdateSwapper(&_Vetrotreasury.TransactOpts, swapper_)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address token_, uint256 amount_, address tokenReceiver_) returns()
func (_Vetrotreasury *VetrotreasuryTransactor) Withdraw(opts *bind.TransactOpts, token_ common.Address, amount_ *big.Int, tokenReceiver_ common.Address) (*types.Transaction, error) {
	return _Vetrotreasury.contract.Transact(opts, "withdraw", token_, amount_, tokenReceiver_)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address token_, uint256 amount_, address tokenReceiver_) returns()
func (_Vetrotreasury *VetrotreasurySession) Withdraw(token_ common.Address, amount_ *big.Int, tokenReceiver_ common.Address) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.Withdraw(&_Vetrotreasury.TransactOpts, token_, amount_, tokenReceiver_)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address token_, uint256 amount_, address tokenReceiver_) returns()
func (_Vetrotreasury *VetrotreasuryTransactorSession) Withdraw(token_ common.Address, amount_ *big.Int, tokenReceiver_ common.Address) (*types.Transaction, error) {
	return _Vetrotreasury.Contract.Withdraw(&_Vetrotreasury.TransactOpts, token_, amount_, tokenReceiver_)
}

// VetrotreasuryAddedToWhitelistIterator is returned from FilterAddedToWhitelist and is used to iterate over the raw logs and unpacked data for AddedToWhitelist events raised by the Vetrotreasury contract.
type VetrotreasuryAddedToWhitelistIterator struct {
	Event *VetrotreasuryAddedToWhitelist // Event containing the contract specifics and raw log

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
func (it *VetrotreasuryAddedToWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VetrotreasuryAddedToWhitelist)
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
		it.Event = new(VetrotreasuryAddedToWhitelist)
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
func (it *VetrotreasuryAddedToWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VetrotreasuryAddedToWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VetrotreasuryAddedToWhitelist represents a AddedToWhitelist event raised by the Vetrotreasury contract.
type VetrotreasuryAddedToWhitelist struct {
	Token  common.Address
	Vault  common.Address
	Oracle common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddedToWhitelist is a free log retrieval operation binding the contract event 0x4ef74c648ce2bd2158a6ad1e2a476c0ea9a36a554b8c31b8f2b80e18b658baf5.
//
// Solidity: event AddedToWhitelist(address indexed token, address indexed vault, address indexed oracle)
func (_Vetrotreasury *VetrotreasuryFilterer) FilterAddedToWhitelist(opts *bind.FilterOpts, token []common.Address, vault []common.Address, oracle []common.Address) (*VetrotreasuryAddedToWhitelistIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _Vetrotreasury.contract.FilterLogs(opts, "AddedToWhitelist", tokenRule, vaultRule, oracleRule)
	if err != nil {
		return nil, err
	}
	return &VetrotreasuryAddedToWhitelistIterator{contract: _Vetrotreasury.contract, event: "AddedToWhitelist", logs: logs, sub: sub}, nil
}

// WatchAddedToWhitelist is a free log subscription operation binding the contract event 0x4ef74c648ce2bd2158a6ad1e2a476c0ea9a36a554b8c31b8f2b80e18b658baf5.
//
// Solidity: event AddedToWhitelist(address indexed token, address indexed vault, address indexed oracle)
func (_Vetrotreasury *VetrotreasuryFilterer) WatchAddedToWhitelist(opts *bind.WatchOpts, sink chan<- *VetrotreasuryAddedToWhitelist, token []common.Address, vault []common.Address, oracle []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _Vetrotreasury.contract.WatchLogs(opts, "AddedToWhitelist", tokenRule, vaultRule, oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VetrotreasuryAddedToWhitelist)
				if err := _Vetrotreasury.contract.UnpackLog(event, "AddedToWhitelist", log); err != nil {
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

// ParseAddedToWhitelist is a log parse operation binding the contract event 0x4ef74c648ce2bd2158a6ad1e2a476c0ea9a36a554b8c31b8f2b80e18b658baf5.
//
// Solidity: event AddedToWhitelist(address indexed token, address indexed vault, address indexed oracle)
func (_Vetrotreasury *VetrotreasuryFilterer) ParseAddedToWhitelist(log types.Log) (*VetrotreasuryAddedToWhitelist, error) {
	event := new(VetrotreasuryAddedToWhitelist)
	if err := _Vetrotreasury.contract.UnpackLog(event, "AddedToWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VetrotreasuryDefaultAdminDelayChangeCanceledIterator is returned from FilterDefaultAdminDelayChangeCanceled and is used to iterate over the raw logs and unpacked data for DefaultAdminDelayChangeCanceled events raised by the Vetrotreasury contract.
type VetrotreasuryDefaultAdminDelayChangeCanceledIterator struct {
	Event *VetrotreasuryDefaultAdminDelayChangeCanceled // Event containing the contract specifics and raw log

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
func (it *VetrotreasuryDefaultAdminDelayChangeCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VetrotreasuryDefaultAdminDelayChangeCanceled)
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
		it.Event = new(VetrotreasuryDefaultAdminDelayChangeCanceled)
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
func (it *VetrotreasuryDefaultAdminDelayChangeCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VetrotreasuryDefaultAdminDelayChangeCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VetrotreasuryDefaultAdminDelayChangeCanceled represents a DefaultAdminDelayChangeCanceled event raised by the Vetrotreasury contract.
type VetrotreasuryDefaultAdminDelayChangeCanceled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminDelayChangeCanceled is a free log retrieval operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_Vetrotreasury *VetrotreasuryFilterer) FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*VetrotreasuryDefaultAdminDelayChangeCanceledIterator, error) {

	logs, sub, err := _Vetrotreasury.contract.FilterLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return &VetrotreasuryDefaultAdminDelayChangeCanceledIterator{contract: _Vetrotreasury.contract, event: "DefaultAdminDelayChangeCanceled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminDelayChangeCanceled is a free log subscription operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_Vetrotreasury *VetrotreasuryFilterer) WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *VetrotreasuryDefaultAdminDelayChangeCanceled) (event.Subscription, error) {

	logs, sub, err := _Vetrotreasury.contract.WatchLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VetrotreasuryDefaultAdminDelayChangeCanceled)
				if err := _Vetrotreasury.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
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

// ParseDefaultAdminDelayChangeCanceled is a log parse operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_Vetrotreasury *VetrotreasuryFilterer) ParseDefaultAdminDelayChangeCanceled(log types.Log) (*VetrotreasuryDefaultAdminDelayChangeCanceled, error) {
	event := new(VetrotreasuryDefaultAdminDelayChangeCanceled)
	if err := _Vetrotreasury.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VetrotreasuryDefaultAdminDelayChangeScheduledIterator is returned from FilterDefaultAdminDelayChangeScheduled and is used to iterate over the raw logs and unpacked data for DefaultAdminDelayChangeScheduled events raised by the Vetrotreasury contract.
type VetrotreasuryDefaultAdminDelayChangeScheduledIterator struct {
	Event *VetrotreasuryDefaultAdminDelayChangeScheduled // Event containing the contract specifics and raw log

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
func (it *VetrotreasuryDefaultAdminDelayChangeScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VetrotreasuryDefaultAdminDelayChangeScheduled)
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
		it.Event = new(VetrotreasuryDefaultAdminDelayChangeScheduled)
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
func (it *VetrotreasuryDefaultAdminDelayChangeScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VetrotreasuryDefaultAdminDelayChangeScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VetrotreasuryDefaultAdminDelayChangeScheduled represents a DefaultAdminDelayChangeScheduled event raised by the Vetrotreasury contract.
type VetrotreasuryDefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminDelayChangeScheduled is a free log retrieval operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_Vetrotreasury *VetrotreasuryFilterer) FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*VetrotreasuryDefaultAdminDelayChangeScheduledIterator, error) {

	logs, sub, err := _Vetrotreasury.contract.FilterLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return &VetrotreasuryDefaultAdminDelayChangeScheduledIterator{contract: _Vetrotreasury.contract, event: "DefaultAdminDelayChangeScheduled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminDelayChangeScheduled is a free log subscription operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_Vetrotreasury *VetrotreasuryFilterer) WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *VetrotreasuryDefaultAdminDelayChangeScheduled) (event.Subscription, error) {

	logs, sub, err := _Vetrotreasury.contract.WatchLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VetrotreasuryDefaultAdminDelayChangeScheduled)
				if err := _Vetrotreasury.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
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

// ParseDefaultAdminDelayChangeScheduled is a log parse operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_Vetrotreasury *VetrotreasuryFilterer) ParseDefaultAdminDelayChangeScheduled(log types.Log) (*VetrotreasuryDefaultAdminDelayChangeScheduled, error) {
	event := new(VetrotreasuryDefaultAdminDelayChangeScheduled)
	if err := _Vetrotreasury.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VetrotreasuryDefaultAdminTransferCanceledIterator is returned from FilterDefaultAdminTransferCanceled and is used to iterate over the raw logs and unpacked data for DefaultAdminTransferCanceled events raised by the Vetrotreasury contract.
type VetrotreasuryDefaultAdminTransferCanceledIterator struct {
	Event *VetrotreasuryDefaultAdminTransferCanceled // Event containing the contract specifics and raw log

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
func (it *VetrotreasuryDefaultAdminTransferCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VetrotreasuryDefaultAdminTransferCanceled)
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
		it.Event = new(VetrotreasuryDefaultAdminTransferCanceled)
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
func (it *VetrotreasuryDefaultAdminTransferCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VetrotreasuryDefaultAdminTransferCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VetrotreasuryDefaultAdminTransferCanceled represents a DefaultAdminTransferCanceled event raised by the Vetrotreasury contract.
type VetrotreasuryDefaultAdminTransferCanceled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminTransferCanceled is a free log retrieval operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_Vetrotreasury *VetrotreasuryFilterer) FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*VetrotreasuryDefaultAdminTransferCanceledIterator, error) {

	logs, sub, err := _Vetrotreasury.contract.FilterLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return &VetrotreasuryDefaultAdminTransferCanceledIterator{contract: _Vetrotreasury.contract, event: "DefaultAdminTransferCanceled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminTransferCanceled is a free log subscription operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_Vetrotreasury *VetrotreasuryFilterer) WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *VetrotreasuryDefaultAdminTransferCanceled) (event.Subscription, error) {

	logs, sub, err := _Vetrotreasury.contract.WatchLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VetrotreasuryDefaultAdminTransferCanceled)
				if err := _Vetrotreasury.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
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

// ParseDefaultAdminTransferCanceled is a log parse operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_Vetrotreasury *VetrotreasuryFilterer) ParseDefaultAdminTransferCanceled(log types.Log) (*VetrotreasuryDefaultAdminTransferCanceled, error) {
	event := new(VetrotreasuryDefaultAdminTransferCanceled)
	if err := _Vetrotreasury.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VetrotreasuryDefaultAdminTransferScheduledIterator is returned from FilterDefaultAdminTransferScheduled and is used to iterate over the raw logs and unpacked data for DefaultAdminTransferScheduled events raised by the Vetrotreasury contract.
type VetrotreasuryDefaultAdminTransferScheduledIterator struct {
	Event *VetrotreasuryDefaultAdminTransferScheduled // Event containing the contract specifics and raw log

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
func (it *VetrotreasuryDefaultAdminTransferScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VetrotreasuryDefaultAdminTransferScheduled)
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
		it.Event = new(VetrotreasuryDefaultAdminTransferScheduled)
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
func (it *VetrotreasuryDefaultAdminTransferScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VetrotreasuryDefaultAdminTransferScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VetrotreasuryDefaultAdminTransferScheduled represents a DefaultAdminTransferScheduled event raised by the Vetrotreasury contract.
type VetrotreasuryDefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminTransferScheduled is a free log retrieval operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_Vetrotreasury *VetrotreasuryFilterer) FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*VetrotreasuryDefaultAdminTransferScheduledIterator, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _Vetrotreasury.contract.FilterLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return &VetrotreasuryDefaultAdminTransferScheduledIterator{contract: _Vetrotreasury.contract, event: "DefaultAdminTransferScheduled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminTransferScheduled is a free log subscription operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_Vetrotreasury *VetrotreasuryFilterer) WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *VetrotreasuryDefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _Vetrotreasury.contract.WatchLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VetrotreasuryDefaultAdminTransferScheduled)
				if err := _Vetrotreasury.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
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

// ParseDefaultAdminTransferScheduled is a log parse operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_Vetrotreasury *VetrotreasuryFilterer) ParseDefaultAdminTransferScheduled(log types.Log) (*VetrotreasuryDefaultAdminTransferScheduled, error) {
	event := new(VetrotreasuryDefaultAdminTransferScheduled)
	if err := _Vetrotreasury.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VetrotreasuryExcessWithdrawnIterator is returned from FilterExcessWithdrawn and is used to iterate over the raw logs and unpacked data for ExcessWithdrawn events raised by the Vetrotreasury contract.
type VetrotreasuryExcessWithdrawnIterator struct {
	Event *VetrotreasuryExcessWithdrawn // Event containing the contract specifics and raw log

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
func (it *VetrotreasuryExcessWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VetrotreasuryExcessWithdrawn)
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
		it.Event = new(VetrotreasuryExcessWithdrawn)
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
func (it *VetrotreasuryExcessWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VetrotreasuryExcessWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VetrotreasuryExcessWithdrawn represents a ExcessWithdrawn event raised by the Vetrotreasury contract.
type VetrotreasuryExcessWithdrawn struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExcessWithdrawn is a free log retrieval operation binding the contract event 0x9bdedcc968f2c0682145c20289dfbb3081a39684571116ab45eabfef512e9473.
//
// Solidity: event ExcessWithdrawn(address indexed token, uint256 amount)
func (_Vetrotreasury *VetrotreasuryFilterer) FilterExcessWithdrawn(opts *bind.FilterOpts, token []common.Address) (*VetrotreasuryExcessWithdrawnIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Vetrotreasury.contract.FilterLogs(opts, "ExcessWithdrawn", tokenRule)
	if err != nil {
		return nil, err
	}
	return &VetrotreasuryExcessWithdrawnIterator{contract: _Vetrotreasury.contract, event: "ExcessWithdrawn", logs: logs, sub: sub}, nil
}

// WatchExcessWithdrawn is a free log subscription operation binding the contract event 0x9bdedcc968f2c0682145c20289dfbb3081a39684571116ab45eabfef512e9473.
//
// Solidity: event ExcessWithdrawn(address indexed token, uint256 amount)
func (_Vetrotreasury *VetrotreasuryFilterer) WatchExcessWithdrawn(opts *bind.WatchOpts, sink chan<- *VetrotreasuryExcessWithdrawn, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Vetrotreasury.contract.WatchLogs(opts, "ExcessWithdrawn", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VetrotreasuryExcessWithdrawn)
				if err := _Vetrotreasury.contract.UnpackLog(event, "ExcessWithdrawn", log); err != nil {
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

// ParseExcessWithdrawn is a log parse operation binding the contract event 0x9bdedcc968f2c0682145c20289dfbb3081a39684571116ab45eabfef512e9473.
//
// Solidity: event ExcessWithdrawn(address indexed token, uint256 amount)
func (_Vetrotreasury *VetrotreasuryFilterer) ParseExcessWithdrawn(log types.Log) (*VetrotreasuryExcessWithdrawn, error) {
	event := new(VetrotreasuryExcessWithdrawn)
	if err := _Vetrotreasury.contract.UnpackLog(event, "ExcessWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VetrotreasuryMigratedIterator is returned from FilterMigrated and is used to iterate over the raw logs and unpacked data for Migrated events raised by the Vetrotreasury contract.
type VetrotreasuryMigratedIterator struct {
	Event *VetrotreasuryMigrated // Event containing the contract specifics and raw log

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
func (it *VetrotreasuryMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VetrotreasuryMigrated)
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
		it.Event = new(VetrotreasuryMigrated)
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
func (it *VetrotreasuryMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VetrotreasuryMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VetrotreasuryMigrated represents a Migrated event raised by the Vetrotreasury contract.
type VetrotreasuryMigrated struct {
	NewTreasury common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMigrated is a free log retrieval operation binding the contract event 0xa2e7361c23d7820040603b83c0cd3f494d377bac69736377d75bb56c651a5098.
//
// Solidity: event Migrated(address indexed newTreasury)
func (_Vetrotreasury *VetrotreasuryFilterer) FilterMigrated(opts *bind.FilterOpts, newTreasury []common.Address) (*VetrotreasuryMigratedIterator, error) {

	var newTreasuryRule []interface{}
	for _, newTreasuryItem := range newTreasury {
		newTreasuryRule = append(newTreasuryRule, newTreasuryItem)
	}

	logs, sub, err := _Vetrotreasury.contract.FilterLogs(opts, "Migrated", newTreasuryRule)
	if err != nil {
		return nil, err
	}
	return &VetrotreasuryMigratedIterator{contract: _Vetrotreasury.contract, event: "Migrated", logs: logs, sub: sub}, nil
}

// WatchMigrated is a free log subscription operation binding the contract event 0xa2e7361c23d7820040603b83c0cd3f494d377bac69736377d75bb56c651a5098.
//
// Solidity: event Migrated(address indexed newTreasury)
func (_Vetrotreasury *VetrotreasuryFilterer) WatchMigrated(opts *bind.WatchOpts, sink chan<- *VetrotreasuryMigrated, newTreasury []common.Address) (event.Subscription, error) {

	var newTreasuryRule []interface{}
	for _, newTreasuryItem := range newTreasury {
		newTreasuryRule = append(newTreasuryRule, newTreasuryItem)
	}

	logs, sub, err := _Vetrotreasury.contract.WatchLogs(opts, "Migrated", newTreasuryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VetrotreasuryMigrated)
				if err := _Vetrotreasury.contract.UnpackLog(event, "Migrated", log); err != nil {
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

// ParseMigrated is a log parse operation binding the contract event 0xa2e7361c23d7820040603b83c0cd3f494d377bac69736377d75bb56c651a5098.
//
// Solidity: event Migrated(address indexed newTreasury)
func (_Vetrotreasury *VetrotreasuryFilterer) ParseMigrated(log types.Log) (*VetrotreasuryMigrated, error) {
	event := new(VetrotreasuryMigrated)
	if err := _Vetrotreasury.contract.UnpackLog(event, "Migrated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VetrotreasuryRemovedFromWhitelistIterator is returned from FilterRemovedFromWhitelist and is used to iterate over the raw logs and unpacked data for RemovedFromWhitelist events raised by the Vetrotreasury contract.
type VetrotreasuryRemovedFromWhitelistIterator struct {
	Event *VetrotreasuryRemovedFromWhitelist // Event containing the contract specifics and raw log

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
func (it *VetrotreasuryRemovedFromWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VetrotreasuryRemovedFromWhitelist)
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
		it.Event = new(VetrotreasuryRemovedFromWhitelist)
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
func (it *VetrotreasuryRemovedFromWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VetrotreasuryRemovedFromWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VetrotreasuryRemovedFromWhitelist represents a RemovedFromWhitelist event raised by the Vetrotreasury contract.
type VetrotreasuryRemovedFromWhitelist struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRemovedFromWhitelist is a free log retrieval operation binding the contract event 0xcdd2e9b91a56913d370075169cefa1602ba36be5301664f752192bb1709df757.
//
// Solidity: event RemovedFromWhitelist(address indexed token)
func (_Vetrotreasury *VetrotreasuryFilterer) FilterRemovedFromWhitelist(opts *bind.FilterOpts, token []common.Address) (*VetrotreasuryRemovedFromWhitelistIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Vetrotreasury.contract.FilterLogs(opts, "RemovedFromWhitelist", tokenRule)
	if err != nil {
		return nil, err
	}
	return &VetrotreasuryRemovedFromWhitelistIterator{contract: _Vetrotreasury.contract, event: "RemovedFromWhitelist", logs: logs, sub: sub}, nil
}

// WatchRemovedFromWhitelist is a free log subscription operation binding the contract event 0xcdd2e9b91a56913d370075169cefa1602ba36be5301664f752192bb1709df757.
//
// Solidity: event RemovedFromWhitelist(address indexed token)
func (_Vetrotreasury *VetrotreasuryFilterer) WatchRemovedFromWhitelist(opts *bind.WatchOpts, sink chan<- *VetrotreasuryRemovedFromWhitelist, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Vetrotreasury.contract.WatchLogs(opts, "RemovedFromWhitelist", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VetrotreasuryRemovedFromWhitelist)
				if err := _Vetrotreasury.contract.UnpackLog(event, "RemovedFromWhitelist", log); err != nil {
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

// ParseRemovedFromWhitelist is a log parse operation binding the contract event 0xcdd2e9b91a56913d370075169cefa1602ba36be5301664f752192bb1709df757.
//
// Solidity: event RemovedFromWhitelist(address indexed token)
func (_Vetrotreasury *VetrotreasuryFilterer) ParseRemovedFromWhitelist(log types.Log) (*VetrotreasuryRemovedFromWhitelist, error) {
	event := new(VetrotreasuryRemovedFromWhitelist)
	if err := _Vetrotreasury.contract.UnpackLog(event, "RemovedFromWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VetrotreasuryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Vetrotreasury contract.
type VetrotreasuryRoleAdminChangedIterator struct {
	Event *VetrotreasuryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *VetrotreasuryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VetrotreasuryRoleAdminChanged)
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
		it.Event = new(VetrotreasuryRoleAdminChanged)
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
func (it *VetrotreasuryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VetrotreasuryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VetrotreasuryRoleAdminChanged represents a RoleAdminChanged event raised by the Vetrotreasury contract.
type VetrotreasuryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Vetrotreasury *VetrotreasuryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*VetrotreasuryRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Vetrotreasury.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &VetrotreasuryRoleAdminChangedIterator{contract: _Vetrotreasury.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Vetrotreasury *VetrotreasuryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *VetrotreasuryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Vetrotreasury.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VetrotreasuryRoleAdminChanged)
				if err := _Vetrotreasury.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Vetrotreasury *VetrotreasuryFilterer) ParseRoleAdminChanged(log types.Log) (*VetrotreasuryRoleAdminChanged, error) {
	event := new(VetrotreasuryRoleAdminChanged)
	if err := _Vetrotreasury.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VetrotreasuryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Vetrotreasury contract.
type VetrotreasuryRoleGrantedIterator struct {
	Event *VetrotreasuryRoleGranted // Event containing the contract specifics and raw log

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
func (it *VetrotreasuryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VetrotreasuryRoleGranted)
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
		it.Event = new(VetrotreasuryRoleGranted)
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
func (it *VetrotreasuryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VetrotreasuryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VetrotreasuryRoleGranted represents a RoleGranted event raised by the Vetrotreasury contract.
type VetrotreasuryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Vetrotreasury *VetrotreasuryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*VetrotreasuryRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Vetrotreasury.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VetrotreasuryRoleGrantedIterator{contract: _Vetrotreasury.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Vetrotreasury *VetrotreasuryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *VetrotreasuryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Vetrotreasury.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VetrotreasuryRoleGranted)
				if err := _Vetrotreasury.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Vetrotreasury *VetrotreasuryFilterer) ParseRoleGranted(log types.Log) (*VetrotreasuryRoleGranted, error) {
	event := new(VetrotreasuryRoleGranted)
	if err := _Vetrotreasury.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VetrotreasuryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Vetrotreasury contract.
type VetrotreasuryRoleRevokedIterator struct {
	Event *VetrotreasuryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *VetrotreasuryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VetrotreasuryRoleRevoked)
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
		it.Event = new(VetrotreasuryRoleRevoked)
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
func (it *VetrotreasuryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VetrotreasuryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VetrotreasuryRoleRevoked represents a RoleRevoked event raised by the Vetrotreasury contract.
type VetrotreasuryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Vetrotreasury *VetrotreasuryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*VetrotreasuryRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Vetrotreasury.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VetrotreasuryRoleRevokedIterator{contract: _Vetrotreasury.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Vetrotreasury *VetrotreasuryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *VetrotreasuryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Vetrotreasury.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VetrotreasuryRoleRevoked)
				if err := _Vetrotreasury.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Vetrotreasury *VetrotreasuryFilterer) ParseRoleRevoked(log types.Log) (*VetrotreasuryRoleRevoked, error) {
	event := new(VetrotreasuryRoleRevoked)
	if err := _Vetrotreasury.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VetrotreasurySetDepositActiveIterator is returned from FilterSetDepositActive and is used to iterate over the raw logs and unpacked data for SetDepositActive events raised by the Vetrotreasury contract.
type VetrotreasurySetDepositActiveIterator struct {
	Event *VetrotreasurySetDepositActive // Event containing the contract specifics and raw log

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
func (it *VetrotreasurySetDepositActiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VetrotreasurySetDepositActive)
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
		it.Event = new(VetrotreasurySetDepositActive)
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
func (it *VetrotreasurySetDepositActiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VetrotreasurySetDepositActiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VetrotreasurySetDepositActive represents a SetDepositActive event raised by the Vetrotreasury contract.
type VetrotreasurySetDepositActive struct {
	Token    common.Address
	NewValue bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetDepositActive is a free log retrieval operation binding the contract event 0xec8e92846f090e284fe3af69214673b93034178308baf850014c7b178521c63d.
//
// Solidity: event SetDepositActive(address indexed token, bool newValue)
func (_Vetrotreasury *VetrotreasuryFilterer) FilterSetDepositActive(opts *bind.FilterOpts, token []common.Address) (*VetrotreasurySetDepositActiveIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Vetrotreasury.contract.FilterLogs(opts, "SetDepositActive", tokenRule)
	if err != nil {
		return nil, err
	}
	return &VetrotreasurySetDepositActiveIterator{contract: _Vetrotreasury.contract, event: "SetDepositActive", logs: logs, sub: sub}, nil
}

// WatchSetDepositActive is a free log subscription operation binding the contract event 0xec8e92846f090e284fe3af69214673b93034178308baf850014c7b178521c63d.
//
// Solidity: event SetDepositActive(address indexed token, bool newValue)
func (_Vetrotreasury *VetrotreasuryFilterer) WatchSetDepositActive(opts *bind.WatchOpts, sink chan<- *VetrotreasurySetDepositActive, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Vetrotreasury.contract.WatchLogs(opts, "SetDepositActive", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VetrotreasurySetDepositActive)
				if err := _Vetrotreasury.contract.UnpackLog(event, "SetDepositActive", log); err != nil {
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

// ParseSetDepositActive is a log parse operation binding the contract event 0xec8e92846f090e284fe3af69214673b93034178308baf850014c7b178521c63d.
//
// Solidity: event SetDepositActive(address indexed token, bool newValue)
func (_Vetrotreasury *VetrotreasuryFilterer) ParseSetDepositActive(log types.Log) (*VetrotreasurySetDepositActive, error) {
	event := new(VetrotreasurySetDepositActive)
	if err := _Vetrotreasury.contract.UnpackLog(event, "SetDepositActive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VetrotreasurySetWithdrawActiveIterator is returned from FilterSetWithdrawActive and is used to iterate over the raw logs and unpacked data for SetWithdrawActive events raised by the Vetrotreasury contract.
type VetrotreasurySetWithdrawActiveIterator struct {
	Event *VetrotreasurySetWithdrawActive // Event containing the contract specifics and raw log

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
func (it *VetrotreasurySetWithdrawActiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VetrotreasurySetWithdrawActive)
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
		it.Event = new(VetrotreasurySetWithdrawActive)
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
func (it *VetrotreasurySetWithdrawActiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VetrotreasurySetWithdrawActiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VetrotreasurySetWithdrawActive represents a SetWithdrawActive event raised by the Vetrotreasury contract.
type VetrotreasurySetWithdrawActive struct {
	Token    common.Address
	NewValue bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetWithdrawActive is a free log retrieval operation binding the contract event 0xe42f32c8178bf44f474966892eb868bc036112f750b5aba609e616be843948bc.
//
// Solidity: event SetWithdrawActive(address indexed token, bool newValue)
func (_Vetrotreasury *VetrotreasuryFilterer) FilterSetWithdrawActive(opts *bind.FilterOpts, token []common.Address) (*VetrotreasurySetWithdrawActiveIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Vetrotreasury.contract.FilterLogs(opts, "SetWithdrawActive", tokenRule)
	if err != nil {
		return nil, err
	}
	return &VetrotreasurySetWithdrawActiveIterator{contract: _Vetrotreasury.contract, event: "SetWithdrawActive", logs: logs, sub: sub}, nil
}

// WatchSetWithdrawActive is a free log subscription operation binding the contract event 0xe42f32c8178bf44f474966892eb868bc036112f750b5aba609e616be843948bc.
//
// Solidity: event SetWithdrawActive(address indexed token, bool newValue)
func (_Vetrotreasury *VetrotreasuryFilterer) WatchSetWithdrawActive(opts *bind.WatchOpts, sink chan<- *VetrotreasurySetWithdrawActive, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Vetrotreasury.contract.WatchLogs(opts, "SetWithdrawActive", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VetrotreasurySetWithdrawActive)
				if err := _Vetrotreasury.contract.UnpackLog(event, "SetWithdrawActive", log); err != nil {
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

// ParseSetWithdrawActive is a log parse operation binding the contract event 0xe42f32c8178bf44f474966892eb868bc036112f750b5aba609e616be843948bc.
//
// Solidity: event SetWithdrawActive(address indexed token, bool newValue)
func (_Vetrotreasury *VetrotreasuryFilterer) ParseSetWithdrawActive(log types.Log) (*VetrotreasurySetWithdrawActive, error) {
	event := new(VetrotreasurySetWithdrawActive)
	if err := _Vetrotreasury.contract.UnpackLog(event, "SetWithdrawActive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VetrotreasurySwappedIterator is returned from FilterSwapped and is used to iterate over the raw logs and unpacked data for Swapped events raised by the Vetrotreasury contract.
type VetrotreasurySwappedIterator struct {
	Event *VetrotreasurySwapped // Event containing the contract specifics and raw log

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
func (it *VetrotreasurySwappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VetrotreasurySwapped)
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
		it.Event = new(VetrotreasurySwapped)
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
func (it *VetrotreasurySwappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VetrotreasurySwappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VetrotreasurySwapped represents a Swapped event raised by the Vetrotreasury contract.
type VetrotreasurySwapped struct {
	TokenIn  common.Address
	TokenOut common.Address
	AmountIn *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSwapped is a free log retrieval operation binding the contract event 0x2e7f8a64aa3240292c0adfa332e1e8945dd31589fcb0bce2721fa21c69b1390f.
//
// Solidity: event Swapped(address indexed tokenIn, address indexed tokenOut, uint256 amountIn)
func (_Vetrotreasury *VetrotreasuryFilterer) FilterSwapped(opts *bind.FilterOpts, tokenIn []common.Address, tokenOut []common.Address) (*VetrotreasurySwappedIterator, error) {

	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _Vetrotreasury.contract.FilterLogs(opts, "Swapped", tokenInRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return &VetrotreasurySwappedIterator{contract: _Vetrotreasury.contract, event: "Swapped", logs: logs, sub: sub}, nil
}

// WatchSwapped is a free log subscription operation binding the contract event 0x2e7f8a64aa3240292c0adfa332e1e8945dd31589fcb0bce2721fa21c69b1390f.
//
// Solidity: event Swapped(address indexed tokenIn, address indexed tokenOut, uint256 amountIn)
func (_Vetrotreasury *VetrotreasuryFilterer) WatchSwapped(opts *bind.WatchOpts, sink chan<- *VetrotreasurySwapped, tokenIn []common.Address, tokenOut []common.Address) (event.Subscription, error) {

	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _Vetrotreasury.contract.WatchLogs(opts, "Swapped", tokenInRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VetrotreasurySwapped)
				if err := _Vetrotreasury.contract.UnpackLog(event, "Swapped", log); err != nil {
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

// ParseSwapped is a log parse operation binding the contract event 0x2e7f8a64aa3240292c0adfa332e1e8945dd31589fcb0bce2721fa21c69b1390f.
//
// Solidity: event Swapped(address indexed tokenIn, address indexed tokenOut, uint256 amountIn)
func (_Vetrotreasury *VetrotreasuryFilterer) ParseSwapped(log types.Log) (*VetrotreasurySwapped, error) {
	event := new(VetrotreasurySwapped)
	if err := _Vetrotreasury.contract.UnpackLog(event, "Swapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VetrotreasurySweptIterator is returned from FilterSwept and is used to iterate over the raw logs and unpacked data for Swept events raised by the Vetrotreasury contract.
type VetrotreasurySweptIterator struct {
	Event *VetrotreasurySwept // Event containing the contract specifics and raw log

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
func (it *VetrotreasurySweptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VetrotreasurySwept)
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
		it.Event = new(VetrotreasurySwept)
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
func (it *VetrotreasurySweptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VetrotreasurySweptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VetrotreasurySwept represents a Swept event raised by the Vetrotreasury contract.
type VetrotreasurySwept struct {
	Token    common.Address
	Amount   *big.Int
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSwept is a free log retrieval operation binding the contract event 0xbb3f74f3539ea7725781ff6810125a75c183f5c944318fc94873d1324f0482ae.
//
// Solidity: event Swept(address indexed token, uint256 amount, address indexed receiver)
func (_Vetrotreasury *VetrotreasuryFilterer) FilterSwept(opts *bind.FilterOpts, token []common.Address, receiver []common.Address) (*VetrotreasurySweptIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Vetrotreasury.contract.FilterLogs(opts, "Swept", tokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &VetrotreasurySweptIterator{contract: _Vetrotreasury.contract, event: "Swept", logs: logs, sub: sub}, nil
}

// WatchSwept is a free log subscription operation binding the contract event 0xbb3f74f3539ea7725781ff6810125a75c183f5c944318fc94873d1324f0482ae.
//
// Solidity: event Swept(address indexed token, uint256 amount, address indexed receiver)
func (_Vetrotreasury *VetrotreasuryFilterer) WatchSwept(opts *bind.WatchOpts, sink chan<- *VetrotreasurySwept, token []common.Address, receiver []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Vetrotreasury.contract.WatchLogs(opts, "Swept", tokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VetrotreasurySwept)
				if err := _Vetrotreasury.contract.UnpackLog(event, "Swept", log); err != nil {
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

// ParseSwept is a log parse operation binding the contract event 0xbb3f74f3539ea7725781ff6810125a75c183f5c944318fc94873d1324f0482ae.
//
// Solidity: event Swept(address indexed token, uint256 amount, address indexed receiver)
func (_Vetrotreasury *VetrotreasuryFilterer) ParseSwept(log types.Log) (*VetrotreasurySwept, error) {
	event := new(VetrotreasurySwept)
	if err := _Vetrotreasury.contract.UnpackLog(event, "Swept", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VetrotreasuryUpdatedOracleIterator is returned from FilterUpdatedOracle and is used to iterate over the raw logs and unpacked data for UpdatedOracle events raised by the Vetrotreasury contract.
type VetrotreasuryUpdatedOracleIterator struct {
	Event *VetrotreasuryUpdatedOracle // Event containing the contract specifics and raw log

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
func (it *VetrotreasuryUpdatedOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VetrotreasuryUpdatedOracle)
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
		it.Event = new(VetrotreasuryUpdatedOracle)
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
func (it *VetrotreasuryUpdatedOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VetrotreasuryUpdatedOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VetrotreasuryUpdatedOracle represents a UpdatedOracle event raised by the Vetrotreasury contract.
type VetrotreasuryUpdatedOracle struct {
	Token       common.Address
	Oracle      common.Address
	StalePeriod *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatedOracle is a free log retrieval operation binding the contract event 0x97617a5525c215c6ce8be06021f0cb489578937e2ce04b9d8d1209d0527db7c5.
//
// Solidity: event UpdatedOracle(address indexed token, address indexed oracle, uint256 stalePeriod)
func (_Vetrotreasury *VetrotreasuryFilterer) FilterUpdatedOracle(opts *bind.FilterOpts, token []common.Address, oracle []common.Address) (*VetrotreasuryUpdatedOracleIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _Vetrotreasury.contract.FilterLogs(opts, "UpdatedOracle", tokenRule, oracleRule)
	if err != nil {
		return nil, err
	}
	return &VetrotreasuryUpdatedOracleIterator{contract: _Vetrotreasury.contract, event: "UpdatedOracle", logs: logs, sub: sub}, nil
}

// WatchUpdatedOracle is a free log subscription operation binding the contract event 0x97617a5525c215c6ce8be06021f0cb489578937e2ce04b9d8d1209d0527db7c5.
//
// Solidity: event UpdatedOracle(address indexed token, address indexed oracle, uint256 stalePeriod)
func (_Vetrotreasury *VetrotreasuryFilterer) WatchUpdatedOracle(opts *bind.WatchOpts, sink chan<- *VetrotreasuryUpdatedOracle, token []common.Address, oracle []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _Vetrotreasury.contract.WatchLogs(opts, "UpdatedOracle", tokenRule, oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VetrotreasuryUpdatedOracle)
				if err := _Vetrotreasury.contract.UnpackLog(event, "UpdatedOracle", log); err != nil {
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

// ParseUpdatedOracle is a log parse operation binding the contract event 0x97617a5525c215c6ce8be06021f0cb489578937e2ce04b9d8d1209d0527db7c5.
//
// Solidity: event UpdatedOracle(address indexed token, address indexed oracle, uint256 stalePeriod)
func (_Vetrotreasury *VetrotreasuryFilterer) ParseUpdatedOracle(log types.Log) (*VetrotreasuryUpdatedOracle, error) {
	event := new(VetrotreasuryUpdatedOracle)
	if err := _Vetrotreasury.contract.UnpackLog(event, "UpdatedOracle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VetrotreasuryUpdatedPriceToleranceIterator is returned from FilterUpdatedPriceTolerance and is used to iterate over the raw logs and unpacked data for UpdatedPriceTolerance events raised by the Vetrotreasury contract.
type VetrotreasuryUpdatedPriceToleranceIterator struct {
	Event *VetrotreasuryUpdatedPriceTolerance // Event containing the contract specifics and raw log

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
func (it *VetrotreasuryUpdatedPriceToleranceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VetrotreasuryUpdatedPriceTolerance)
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
		it.Event = new(VetrotreasuryUpdatedPriceTolerance)
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
func (it *VetrotreasuryUpdatedPriceToleranceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VetrotreasuryUpdatedPriceToleranceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VetrotreasuryUpdatedPriceTolerance represents a UpdatedPriceTolerance event raised by the Vetrotreasury contract.
type VetrotreasuryUpdatedPriceTolerance struct {
	PreviousPriceTolerance *big.Int
	NewPriceTolerance      *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterUpdatedPriceTolerance is a free log retrieval operation binding the contract event 0x3cb23ae97dcd603618548f333b96dbb9e41059b6fd0eb7face59058c5ac3b11b.
//
// Solidity: event UpdatedPriceTolerance(uint256 previousPriceTolerance, uint256 newPriceTolerance)
func (_Vetrotreasury *VetrotreasuryFilterer) FilterUpdatedPriceTolerance(opts *bind.FilterOpts) (*VetrotreasuryUpdatedPriceToleranceIterator, error) {

	logs, sub, err := _Vetrotreasury.contract.FilterLogs(opts, "UpdatedPriceTolerance")
	if err != nil {
		return nil, err
	}
	return &VetrotreasuryUpdatedPriceToleranceIterator{contract: _Vetrotreasury.contract, event: "UpdatedPriceTolerance", logs: logs, sub: sub}, nil
}

// WatchUpdatedPriceTolerance is a free log subscription operation binding the contract event 0x3cb23ae97dcd603618548f333b96dbb9e41059b6fd0eb7face59058c5ac3b11b.
//
// Solidity: event UpdatedPriceTolerance(uint256 previousPriceTolerance, uint256 newPriceTolerance)
func (_Vetrotreasury *VetrotreasuryFilterer) WatchUpdatedPriceTolerance(opts *bind.WatchOpts, sink chan<- *VetrotreasuryUpdatedPriceTolerance) (event.Subscription, error) {

	logs, sub, err := _Vetrotreasury.contract.WatchLogs(opts, "UpdatedPriceTolerance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VetrotreasuryUpdatedPriceTolerance)
				if err := _Vetrotreasury.contract.UnpackLog(event, "UpdatedPriceTolerance", log); err != nil {
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

// ParseUpdatedPriceTolerance is a log parse operation binding the contract event 0x3cb23ae97dcd603618548f333b96dbb9e41059b6fd0eb7face59058c5ac3b11b.
//
// Solidity: event UpdatedPriceTolerance(uint256 previousPriceTolerance, uint256 newPriceTolerance)
func (_Vetrotreasury *VetrotreasuryFilterer) ParseUpdatedPriceTolerance(log types.Log) (*VetrotreasuryUpdatedPriceTolerance, error) {
	event := new(VetrotreasuryUpdatedPriceTolerance)
	if err := _Vetrotreasury.contract.UnpackLog(event, "UpdatedPriceTolerance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VetrotreasuryUpdatedSwapperIterator is returned from FilterUpdatedSwapper and is used to iterate over the raw logs and unpacked data for UpdatedSwapper events raised by the Vetrotreasury contract.
type VetrotreasuryUpdatedSwapperIterator struct {
	Event *VetrotreasuryUpdatedSwapper // Event containing the contract specifics and raw log

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
func (it *VetrotreasuryUpdatedSwapperIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VetrotreasuryUpdatedSwapper)
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
		it.Event = new(VetrotreasuryUpdatedSwapper)
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
func (it *VetrotreasuryUpdatedSwapperIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VetrotreasuryUpdatedSwapperIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VetrotreasuryUpdatedSwapper represents a UpdatedSwapper event raised by the Vetrotreasury contract.
type VetrotreasuryUpdatedSwapper struct {
	PreviousSwapper common.Address
	NewSwapper      common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdatedSwapper is a free log retrieval operation binding the contract event 0x6c953b7ec311055c20b96a42cea31e89528e375b1bf953a503db40854b3188fe.
//
// Solidity: event UpdatedSwapper(address indexed previousSwapper, address indexed newSwapper)
func (_Vetrotreasury *VetrotreasuryFilterer) FilterUpdatedSwapper(opts *bind.FilterOpts, previousSwapper []common.Address, newSwapper []common.Address) (*VetrotreasuryUpdatedSwapperIterator, error) {

	var previousSwapperRule []interface{}
	for _, previousSwapperItem := range previousSwapper {
		previousSwapperRule = append(previousSwapperRule, previousSwapperItem)
	}
	var newSwapperRule []interface{}
	for _, newSwapperItem := range newSwapper {
		newSwapperRule = append(newSwapperRule, newSwapperItem)
	}

	logs, sub, err := _Vetrotreasury.contract.FilterLogs(opts, "UpdatedSwapper", previousSwapperRule, newSwapperRule)
	if err != nil {
		return nil, err
	}
	return &VetrotreasuryUpdatedSwapperIterator{contract: _Vetrotreasury.contract, event: "UpdatedSwapper", logs: logs, sub: sub}, nil
}

// WatchUpdatedSwapper is a free log subscription operation binding the contract event 0x6c953b7ec311055c20b96a42cea31e89528e375b1bf953a503db40854b3188fe.
//
// Solidity: event UpdatedSwapper(address indexed previousSwapper, address indexed newSwapper)
func (_Vetrotreasury *VetrotreasuryFilterer) WatchUpdatedSwapper(opts *bind.WatchOpts, sink chan<- *VetrotreasuryUpdatedSwapper, previousSwapper []common.Address, newSwapper []common.Address) (event.Subscription, error) {

	var previousSwapperRule []interface{}
	for _, previousSwapperItem := range previousSwapper {
		previousSwapperRule = append(previousSwapperRule, previousSwapperItem)
	}
	var newSwapperRule []interface{}
	for _, newSwapperItem := range newSwapper {
		newSwapperRule = append(newSwapperRule, newSwapperItem)
	}

	logs, sub, err := _Vetrotreasury.contract.WatchLogs(opts, "UpdatedSwapper", previousSwapperRule, newSwapperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VetrotreasuryUpdatedSwapper)
				if err := _Vetrotreasury.contract.UnpackLog(event, "UpdatedSwapper", log); err != nil {
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

// ParseUpdatedSwapper is a log parse operation binding the contract event 0x6c953b7ec311055c20b96a42cea31e89528e375b1bf953a503db40854b3188fe.
//
// Solidity: event UpdatedSwapper(address indexed previousSwapper, address indexed newSwapper)
func (_Vetrotreasury *VetrotreasuryFilterer) ParseUpdatedSwapper(log types.Log) (*VetrotreasuryUpdatedSwapper, error) {
	event := new(VetrotreasuryUpdatedSwapper)
	if err := _Vetrotreasury.contract.UnpackLog(event, "UpdatedSwapper", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
