// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package spseistaking

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

// SpseistakingMetaData contains all meta data concerning the Spseistaking contract.
var SpseistakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unstakeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bufferBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"InstantUnstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"param\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ParamsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"from\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"PendingUnstakeMigrated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ProtocolFeesClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"from\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Redistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Refreshed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Reinvested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lstAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bufferBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentExchangeRatio\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tiers\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tierFeeRatios\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TiersSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"unstakeValidators\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"UnstakeValidatorsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unstakeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentExchangeRatio\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"validator\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"targetAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ValidatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"priorityList\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isInPriority\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ValidatorPrioritySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"validator\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ValidatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"validators\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"targetAmounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ValidatorTargetAmountSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Withdrawed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"HOUR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_FEE_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_FILL_LIMIT\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_TIERS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_CLAIM_INTERVAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_STAKE_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_UNSTAKE_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NORMALIZE_FACTOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RATIO_DENOM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAW_DAYS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accumulatedFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accumulatedVolume\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"validator\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"targetAmount\",\"type\":\"uint256\"}],\"name\":\"addValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bufferBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bufferFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bufferRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimProtocolFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimableBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fillInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExchangeRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"getFeeWithBuffer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getNormalizeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTiers\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalSei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingStorage\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_orderQueue\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"_validators\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_targetAmounts\",\"type\":\"uint256[]\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"instantUnstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastClaimTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastFillTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRefreshTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"from\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"migratePendingUnstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"orderQueue\",\"outputs\":[{\"internalType\":\"contractIOrderQueue\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"from\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"redistribute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refreshInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"validator\",\"type\":\"string\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardStorage\",\"outputs\":[{\"internalType\":\"contractRewardStorage\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bufferRatio\",\"type\":\"uint256\"}],\"name\":\"setBufferRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_refreshInterval\",\"type\":\"uint256\"}],\"name\":\"setRefreshInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_spreadFeeRate\",\"type\":\"uint256\"}],\"name\":\"setSpreadFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tiers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_tierFeeRatios\",\"type\":\"uint256[]\"}],\"name\":\"setTiers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_unstakeFeeRate\",\"type\":\"uint256\"}],\"name\":\"setUnstakeFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_unstakeValidators\",\"type\":\"string[]\"}],\"name\":\"setUnstakeValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_priorityList\",\"type\":\"string[]\"},{\"internalType\":\"bool\",\"name\":\"_isInPriority\",\"type\":\"bool\"}],\"name\":\"setValidatorPriority\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_validators\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_targetAmount\",\"type\":\"uint256[]\"}],\"name\":\"setValidatorTargetAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spSei\",\"outputs\":[{\"internalType\":\"contractSpSei\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spreadFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isDonation\",\"type\":\"bool\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetBufferBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unstakeFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// SpseistakingABI is the input ABI used to generate the binding from.
// Deprecated: Use SpseistakingMetaData.ABI instead.
var SpseistakingABI = SpseistakingMetaData.ABI

// Spseistaking is an auto generated Go binding around an Ethereum contract.
type Spseistaking struct {
	SpseistakingCaller     // Read-only binding to the contract
	SpseistakingTransactor // Write-only binding to the contract
	SpseistakingFilterer   // Log filterer for contract events
}

// SpseistakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type SpseistakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpseistakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SpseistakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpseistakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SpseistakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpseistakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SpseistakingSession struct {
	Contract     *Spseistaking     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SpseistakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SpseistakingCallerSession struct {
	Contract *SpseistakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// SpseistakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SpseistakingTransactorSession struct {
	Contract     *SpseistakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SpseistakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type SpseistakingRaw struct {
	Contract *Spseistaking // Generic contract binding to access the raw methods on
}

// SpseistakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SpseistakingCallerRaw struct {
	Contract *SpseistakingCaller // Generic read-only contract binding to access the raw methods on
}

// SpseistakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SpseistakingTransactorRaw struct {
	Contract *SpseistakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSpseistaking creates a new instance of Spseistaking, bound to a specific deployed contract.
func NewSpseistaking(address common.Address, backend bind.ContractBackend) (*Spseistaking, error) {
	contract, err := bindSpseistaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Spseistaking{SpseistakingCaller: SpseistakingCaller{contract: contract}, SpseistakingTransactor: SpseistakingTransactor{contract: contract}, SpseistakingFilterer: SpseistakingFilterer{contract: contract}}, nil
}

// NewSpseistakingCaller creates a new read-only instance of Spseistaking, bound to a specific deployed contract.
func NewSpseistakingCaller(address common.Address, caller bind.ContractCaller) (*SpseistakingCaller, error) {
	contract, err := bindSpseistaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SpseistakingCaller{contract: contract}, nil
}

// NewSpseistakingTransactor creates a new write-only instance of Spseistaking, bound to a specific deployed contract.
func NewSpseistakingTransactor(address common.Address, transactor bind.ContractTransactor) (*SpseistakingTransactor, error) {
	contract, err := bindSpseistaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SpseistakingTransactor{contract: contract}, nil
}

// NewSpseistakingFilterer creates a new log filterer instance of Spseistaking, bound to a specific deployed contract.
func NewSpseistakingFilterer(address common.Address, filterer bind.ContractFilterer) (*SpseistakingFilterer, error) {
	contract, err := bindSpseistaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SpseistakingFilterer{contract: contract}, nil
}

// bindSpseistaking binds a generic wrapper to an already deployed contract.
func bindSpseistaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SpseistakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Spseistaking *SpseistakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Spseistaking.Contract.SpseistakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Spseistaking *SpseistakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Spseistaking.Contract.SpseistakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Spseistaking *SpseistakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Spseistaking.Contract.SpseistakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Spseistaking *SpseistakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Spseistaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Spseistaking *SpseistakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Spseistaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Spseistaking *SpseistakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Spseistaking.Contract.contract.Transact(opts, method, params...)
}

// HOUR is a free data retrieval call binding the contract method 0xa39dc9be.
//
// Solidity: function HOUR() view returns(uint256)
func (_Spseistaking *SpseistakingCaller) HOUR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Spseistaking.contract.Call(opts, &out, "HOUR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HOUR is a free data retrieval call binding the contract method 0xa39dc9be.
//
// Solidity: function HOUR() view returns(uint256)
func (_Spseistaking *SpseistakingSession) HOUR() (*big.Int, error) {
	return _Spseistaking.Contract.HOUR(&_Spseistaking.CallOpts)
}

// HOUR is a free data retrieval call binding the contract method 0xa39dc9be.
//
// Solidity: function HOUR() view returns(uint256)
func (_Spseistaking *SpseistakingCallerSession) HOUR() (*big.Int, error) {
	return _Spseistaking.Contract.HOUR(&_Spseistaking.CallOpts)
}

// MAXFEERATIO is a free data retrieval call binding the contract method 0x29c69f4b.
//
// Solidity: function MAX_FEE_RATIO() view returns(uint256)
func (_Spseistaking *SpseistakingCaller) MAXFEERATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Spseistaking.contract.Call(opts, &out, "MAX_FEE_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFEERATIO is a free data retrieval call binding the contract method 0x29c69f4b.
//
// Solidity: function MAX_FEE_RATIO() view returns(uint256)
func (_Spseistaking *SpseistakingSession) MAXFEERATIO() (*big.Int, error) {
	return _Spseistaking.Contract.MAXFEERATIO(&_Spseistaking.CallOpts)
}

// MAXFEERATIO is a free data retrieval call binding the contract method 0x29c69f4b.
//
// Solidity: function MAX_FEE_RATIO() view returns(uint256)
func (_Spseistaking *SpseistakingCallerSession) MAXFEERATIO() (*big.Int, error) {
	return _Spseistaking.Contract.MAXFEERATIO(&_Spseistaking.CallOpts)
}

// MAXFILLLIMIT is a free data retrieval call binding the contract method 0xc3c0f1ba.
//
// Solidity: function MAX_FILL_LIMIT() view returns(uint8)
func (_Spseistaking *SpseistakingCaller) MAXFILLLIMIT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Spseistaking.contract.Call(opts, &out, "MAX_FILL_LIMIT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAXFILLLIMIT is a free data retrieval call binding the contract method 0xc3c0f1ba.
//
// Solidity: function MAX_FILL_LIMIT() view returns(uint8)
func (_Spseistaking *SpseistakingSession) MAXFILLLIMIT() (uint8, error) {
	return _Spseistaking.Contract.MAXFILLLIMIT(&_Spseistaking.CallOpts)
}

// MAXFILLLIMIT is a free data retrieval call binding the contract method 0xc3c0f1ba.
//
// Solidity: function MAX_FILL_LIMIT() view returns(uint8)
func (_Spseistaking *SpseistakingCallerSession) MAXFILLLIMIT() (uint8, error) {
	return _Spseistaking.Contract.MAXFILLLIMIT(&_Spseistaking.CallOpts)
}

// MAXTIERS is a free data retrieval call binding the contract method 0x07c3d4af.
//
// Solidity: function MAX_TIERS() view returns(uint256)
func (_Spseistaking *SpseistakingCaller) MAXTIERS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Spseistaking.contract.Call(opts, &out, "MAX_TIERS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTIERS is a free data retrieval call binding the contract method 0x07c3d4af.
//
// Solidity: function MAX_TIERS() view returns(uint256)
func (_Spseistaking *SpseistakingSession) MAXTIERS() (*big.Int, error) {
	return _Spseistaking.Contract.MAXTIERS(&_Spseistaking.CallOpts)
}

// MAXTIERS is a free data retrieval call binding the contract method 0x07c3d4af.
//
// Solidity: function MAX_TIERS() view returns(uint256)
func (_Spseistaking *SpseistakingCallerSession) MAXTIERS() (*big.Int, error) {
	return _Spseistaking.Contract.MAXTIERS(&_Spseistaking.CallOpts)
}

// MINCLAIMINTERVAL is a free data retrieval call binding the contract method 0xcd6606e2.
//
// Solidity: function MIN_CLAIM_INTERVAL() view returns(uint256)
func (_Spseistaking *SpseistakingCaller) MINCLAIMINTERVAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Spseistaking.contract.Call(opts, &out, "MIN_CLAIM_INTERVAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINCLAIMINTERVAL is a free data retrieval call binding the contract method 0xcd6606e2.
//
// Solidity: function MIN_CLAIM_INTERVAL() view returns(uint256)
func (_Spseistaking *SpseistakingSession) MINCLAIMINTERVAL() (*big.Int, error) {
	return _Spseistaking.Contract.MINCLAIMINTERVAL(&_Spseistaking.CallOpts)
}

// MINCLAIMINTERVAL is a free data retrieval call binding the contract method 0xcd6606e2.
//
// Solidity: function MIN_CLAIM_INTERVAL() view returns(uint256)
func (_Spseistaking *SpseistakingCallerSession) MINCLAIMINTERVAL() (*big.Int, error) {
	return _Spseistaking.Contract.MINCLAIMINTERVAL(&_Spseistaking.CallOpts)
}

// MINSTAKEAMOUNT is a free data retrieval call binding the contract method 0x27ed7188.
//
// Solidity: function MIN_STAKE_AMOUNT() view returns(uint256)
func (_Spseistaking *SpseistakingCaller) MINSTAKEAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Spseistaking.contract.Call(opts, &out, "MIN_STAKE_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINSTAKEAMOUNT is a free data retrieval call binding the contract method 0x27ed7188.
//
// Solidity: function MIN_STAKE_AMOUNT() view returns(uint256)
func (_Spseistaking *SpseistakingSession) MINSTAKEAMOUNT() (*big.Int, error) {
	return _Spseistaking.Contract.MINSTAKEAMOUNT(&_Spseistaking.CallOpts)
}

// MINSTAKEAMOUNT is a free data retrieval call binding the contract method 0x27ed7188.
//
// Solidity: function MIN_STAKE_AMOUNT() view returns(uint256)
func (_Spseistaking *SpseistakingCallerSession) MINSTAKEAMOUNT() (*big.Int, error) {
	return _Spseistaking.Contract.MINSTAKEAMOUNT(&_Spseistaking.CallOpts)
}

// MINUNSTAKEAMOUNT is a free data retrieval call binding the contract method 0xbc6738cb.
//
// Solidity: function MIN_UNSTAKE_AMOUNT() view returns(uint256)
func (_Spseistaking *SpseistakingCaller) MINUNSTAKEAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Spseistaking.contract.Call(opts, &out, "MIN_UNSTAKE_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINUNSTAKEAMOUNT is a free data retrieval call binding the contract method 0xbc6738cb.
//
// Solidity: function MIN_UNSTAKE_AMOUNT() view returns(uint256)
func (_Spseistaking *SpseistakingSession) MINUNSTAKEAMOUNT() (*big.Int, error) {
	return _Spseistaking.Contract.MINUNSTAKEAMOUNT(&_Spseistaking.CallOpts)
}

// MINUNSTAKEAMOUNT is a free data retrieval call binding the contract method 0xbc6738cb.
//
// Solidity: function MIN_UNSTAKE_AMOUNT() view returns(uint256)
func (_Spseistaking *SpseistakingCallerSession) MINUNSTAKEAMOUNT() (*big.Int, error) {
	return _Spseistaking.Contract.MINUNSTAKEAMOUNT(&_Spseistaking.CallOpts)
}

// NORMALIZEFACTOR is a free data retrieval call binding the contract method 0x545c2030.
//
// Solidity: function NORMALIZE_FACTOR() view returns(uint256)
func (_Spseistaking *SpseistakingCaller) NORMALIZEFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Spseistaking.contract.Call(opts, &out, "NORMALIZE_FACTOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NORMALIZEFACTOR is a free data retrieval call binding the contract method 0x545c2030.
//
// Solidity: function NORMALIZE_FACTOR() view returns(uint256)
func (_Spseistaking *SpseistakingSession) NORMALIZEFACTOR() (*big.Int, error) {
	return _Spseistaking.Contract.NORMALIZEFACTOR(&_Spseistaking.CallOpts)
}

// NORMALIZEFACTOR is a free data retrieval call binding the contract method 0x545c2030.
//
// Solidity: function NORMALIZE_FACTOR() view returns(uint256)
func (_Spseistaking *SpseistakingCallerSession) NORMALIZEFACTOR() (*big.Int, error) {
	return _Spseistaking.Contract.NORMALIZEFACTOR(&_Spseistaking.CallOpts)
}

// RATIODENOM is a free data retrieval call binding the contract method 0x7c86be0d.
//
// Solidity: function RATIO_DENOM() view returns(uint256)
func (_Spseistaking *SpseistakingCaller) RATIODENOM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Spseistaking.contract.Call(opts, &out, "RATIO_DENOM")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RATIODENOM is a free data retrieval call binding the contract method 0x7c86be0d.
//
// Solidity: function RATIO_DENOM() view returns(uint256)
func (_Spseistaking *SpseistakingSession) RATIODENOM() (*big.Int, error) {
	return _Spseistaking.Contract.RATIODENOM(&_Spseistaking.CallOpts)
}

// RATIODENOM is a free data retrieval call binding the contract method 0x7c86be0d.
//
// Solidity: function RATIO_DENOM() view returns(uint256)
func (_Spseistaking *SpseistakingCallerSession) RATIODENOM() (*big.Int, error) {
	return _Spseistaking.Contract.RATIODENOM(&_Spseistaking.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Spseistaking *SpseistakingCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Spseistaking.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Spseistaking *SpseistakingSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Spseistaking.Contract.UPGRADEINTERFACEVERSION(&_Spseistaking.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Spseistaking *SpseistakingCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Spseistaking.Contract.UPGRADEINTERFACEVERSION(&_Spseistaking.CallOpts)
}

// WITHDRAWDAYS is a free data retrieval call binding the contract method 0x371bf6ae.
//
// Solidity: function WITHDRAW_DAYS() view returns(uint256)
func (_Spseistaking *SpseistakingCaller) WITHDRAWDAYS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Spseistaking.contract.Call(opts, &out, "WITHDRAW_DAYS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WITHDRAWDAYS is a free data retrieval call binding the contract method 0x371bf6ae.
//
// Solidity: function WITHDRAW_DAYS() view returns(uint256)
func (_Spseistaking *SpseistakingSession) WITHDRAWDAYS() (*big.Int, error) {
	return _Spseistaking.Contract.WITHDRAWDAYS(&_Spseistaking.CallOpts)
}

// WITHDRAWDAYS is a free data retrieval call binding the contract method 0x371bf6ae.
//
// Solidity: function WITHDRAW_DAYS() view returns(uint256)
func (_Spseistaking *SpseistakingCallerSession) WITHDRAWDAYS() (*big.Int, error) {
	return _Spseistaking.Contract.WITHDRAWDAYS(&_Spseistaking.CallOpts)
}

// AccumulatedFee is a free data retrieval call binding the contract method 0xf8ce3164.
//
// Solidity: function accumulatedFee() view returns(uint256)
func (_Spseistaking *SpseistakingCaller) AccumulatedFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Spseistaking.contract.Call(opts, &out, "accumulatedFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccumulatedFee is a free data retrieval call binding the contract method 0xf8ce3164.
//
// Solidity: function accumulatedFee() view returns(uint256)
func (_Spseistaking *SpseistakingSession) AccumulatedFee() (*big.Int, error) {
	return _Spseistaking.Contract.AccumulatedFee(&_Spseistaking.CallOpts)
}

// AccumulatedFee is a free data retrieval call binding the contract method 0xf8ce3164.
//
// Solidity: function accumulatedFee() view returns(uint256)
func (_Spseistaking *SpseistakingCallerSession) AccumulatedFee() (*big.Int, error) {
	return _Spseistaking.Contract.AccumulatedFee(&_Spseistaking.CallOpts)
}

// AccumulatedVolume is a free data retrieval call binding the contract method 0x4fb5ea18.
//
// Solidity: function accumulatedVolume() view returns(uint256)
func (_Spseistaking *SpseistakingCaller) AccumulatedVolume(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Spseistaking.contract.Call(opts, &out, "accumulatedVolume")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccumulatedVolume is a free data retrieval call binding the contract method 0x4fb5ea18.
//
// Solidity: function accumulatedVolume() view returns(uint256)
func (_Spseistaking *SpseistakingSession) AccumulatedVolume() (*big.Int, error) {
	return _Spseistaking.Contract.AccumulatedVolume(&_Spseistaking.CallOpts)
}

// AccumulatedVolume is a free data retrieval call binding the contract method 0x4fb5ea18.
//
// Solidity: function accumulatedVolume() view returns(uint256)
func (_Spseistaking *SpseistakingCallerSession) AccumulatedVolume() (*big.Int, error) {
	return _Spseistaking.Contract.AccumulatedVolume(&_Spseistaking.CallOpts)
}

// BufferBalance is a free data retrieval call binding the contract method 0x1298f13c.
//
// Solidity: function bufferBalance() view returns(uint256)
func (_Spseistaking *SpseistakingCaller) BufferBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Spseistaking.contract.Call(opts, &out, "bufferBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BufferBalance is a free data retrieval call binding the contract method 0x1298f13c.
//
// Solidity: function bufferBalance() view returns(uint256)
func (_Spseistaking *SpseistakingSession) BufferBalance() (*big.Int, error) {
	return _Spseistaking.Contract.BufferBalance(&_Spseistaking.CallOpts)
}

// BufferBalance is a free data retrieval call binding the contract method 0x1298f13c.
//
// Solidity: function bufferBalance() view returns(uint256)
func (_Spseistaking *SpseistakingCallerSession) BufferBalance() (*big.Int, error) {
	return _Spseistaking.Contract.BufferBalance(&_Spseistaking.CallOpts)
}

// BufferFeeRate is a free data retrieval call binding the contract method 0x504adcfb.
//
// Solidity: function bufferFeeRate() view returns(uint256)
func (_Spseistaking *SpseistakingCaller) BufferFeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Spseistaking.contract.Call(opts, &out, "bufferFeeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BufferFeeRate is a free data retrieval call binding the contract method 0x504adcfb.
//
// Solidity: function bufferFeeRate() view returns(uint256)
func (_Spseistaking *SpseistakingSession) BufferFeeRate() (*big.Int, error) {
	return _Spseistaking.Contract.BufferFeeRate(&_Spseistaking.CallOpts)
}

// BufferFeeRate is a free data retrieval call binding the contract method 0x504adcfb.
//
// Solidity: function bufferFeeRate() view returns(uint256)
func (_Spseistaking *SpseistakingCallerSession) BufferFeeRate() (*big.Int, error) {
	return _Spseistaking.Contract.BufferFeeRate(&_Spseistaking.CallOpts)
}

// BufferRatio is a free data retrieval call binding the contract method 0xa37f6911.
//
// Solidity: function bufferRatio() view returns(uint256)
func (_Spseistaking *SpseistakingCaller) BufferRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Spseistaking.contract.Call(opts, &out, "bufferRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BufferRatio is a free data retrieval call binding the contract method 0xa37f6911.
//
// Solidity: function bufferRatio() view returns(uint256)
func (_Spseistaking *SpseistakingSession) BufferRatio() (*big.Int, error) {
	return _Spseistaking.Contract.BufferRatio(&_Spseistaking.CallOpts)
}

// BufferRatio is a free data retrieval call binding the contract method 0xa37f6911.
//
// Solidity: function bufferRatio() view returns(uint256)
func (_Spseistaking *SpseistakingCallerSession) BufferRatio() (*big.Int, error) {
	return _Spseistaking.Contract.BufferRatio(&_Spseistaking.CallOpts)
}

// ClaimableBalance is a free data retrieval call binding the contract method 0x60f3309b.
//
// Solidity: function claimableBalance(address ) view returns(uint256)
func (_Spseistaking *SpseistakingCaller) ClaimableBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Spseistaking.contract.Call(opts, &out, "claimableBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimableBalance is a free data retrieval call binding the contract method 0x60f3309b.
//
// Solidity: function claimableBalance(address ) view returns(uint256)
func (_Spseistaking *SpseistakingSession) ClaimableBalance(arg0 common.Address) (*big.Int, error) {
	return _Spseistaking.Contract.ClaimableBalance(&_Spseistaking.CallOpts, arg0)
}

// ClaimableBalance is a free data retrieval call binding the contract method 0x60f3309b.
//
// Solidity: function claimableBalance(address ) view returns(uint256)
func (_Spseistaking *SpseistakingCallerSession) ClaimableBalance(arg0 common.Address) (*big.Int, error) {
	return _Spseistaking.Contract.ClaimableBalance(&_Spseistaking.CallOpts, arg0)
}

// FillInterval is a free data retrieval call binding the contract method 0x0c9eb8c5.
//
// Solidity: function fillInterval() view returns(uint256)
func (_Spseistaking *SpseistakingCaller) FillInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Spseistaking.contract.Call(opts, &out, "fillInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FillInterval is a free data retrieval call binding the contract method 0x0c9eb8c5.
//
// Solidity: function fillInterval() view returns(uint256)
func (_Spseistaking *SpseistakingSession) FillInterval() (*big.Int, error) {
	return _Spseistaking.Contract.FillInterval(&_Spseistaking.CallOpts)
}

// FillInterval is a free data retrieval call binding the contract method 0x0c9eb8c5.
//
// Solidity: function fillInterval() view returns(uint256)
func (_Spseistaking *SpseistakingCallerSession) FillInterval() (*big.Int, error) {
	return _Spseistaking.Contract.FillInterval(&_Spseistaking.CallOpts)
}

// GetExchangeRatio is a free data retrieval call binding the contract method 0xe1f156a3.
//
// Solidity: function getExchangeRatio() view returns(uint256)
func (_Spseistaking *SpseistakingCaller) GetExchangeRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Spseistaking.contract.Call(opts, &out, "getExchangeRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExchangeRatio is a free data retrieval call binding the contract method 0xe1f156a3.
//
// Solidity: function getExchangeRatio() view returns(uint256)
func (_Spseistaking *SpseistakingSession) GetExchangeRatio() (*big.Int, error) {
	return _Spseistaking.Contract.GetExchangeRatio(&_Spseistaking.CallOpts)
}

// GetExchangeRatio is a free data retrieval call binding the contract method 0xe1f156a3.
//
// Solidity: function getExchangeRatio() view returns(uint256)
func (_Spseistaking *SpseistakingCallerSession) GetExchangeRatio() (*big.Int, error) {
	return _Spseistaking.Contract.GetExchangeRatio(&_Spseistaking.CallOpts)
}

// GetFeeWithBuffer is a free data retrieval call binding the contract method 0x6ba885d7.
//
// Solidity: function getFeeWithBuffer(uint256 withdrawAmount) view returns(bool, uint256)
func (_Spseistaking *SpseistakingCaller) GetFeeWithBuffer(opts *bind.CallOpts, withdrawAmount *big.Int) (bool, *big.Int, error) {
	var out []interface{}
	err := _Spseistaking.contract.Call(opts, &out, "getFeeWithBuffer", withdrawAmount)

	if err != nil {
		return *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetFeeWithBuffer is a free data retrieval call binding the contract method 0x6ba885d7.
//
// Solidity: function getFeeWithBuffer(uint256 withdrawAmount) view returns(bool, uint256)
func (_Spseistaking *SpseistakingSession) GetFeeWithBuffer(withdrawAmount *big.Int) (bool, *big.Int, error) {
	return _Spseistaking.Contract.GetFeeWithBuffer(&_Spseistaking.CallOpts, withdrawAmount)
}

// GetFeeWithBuffer is a free data retrieval call binding the contract method 0x6ba885d7.
//
// Solidity: function getFeeWithBuffer(uint256 withdrawAmount) view returns(bool, uint256)
func (_Spseistaking *SpseistakingCallerSession) GetFeeWithBuffer(withdrawAmount *big.Int) (bool, *big.Int, error) {
	return _Spseistaking.Contract.GetFeeWithBuffer(&_Spseistaking.CallOpts, withdrawAmount)
}

// GetNormalizeAmount is a free data retrieval call binding the contract method 0xb425eada.
//
// Solidity: function getNormalizeAmount(uint256 amount) pure returns(uint256, uint256)
func (_Spseistaking *SpseistakingCaller) GetNormalizeAmount(opts *bind.CallOpts, amount *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Spseistaking.contract.Call(opts, &out, "getNormalizeAmount", amount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetNormalizeAmount is a free data retrieval call binding the contract method 0xb425eada.
//
// Solidity: function getNormalizeAmount(uint256 amount) pure returns(uint256, uint256)
func (_Spseistaking *SpseistakingSession) GetNormalizeAmount(amount *big.Int) (*big.Int, *big.Int, error) {
	return _Spseistaking.Contract.GetNormalizeAmount(&_Spseistaking.CallOpts, amount)
}

// GetNormalizeAmount is a free data retrieval call binding the contract method 0xb425eada.
//
// Solidity: function getNormalizeAmount(uint256 amount) pure returns(uint256, uint256)
func (_Spseistaking *SpseistakingCallerSession) GetNormalizeAmount(amount *big.Int) (*big.Int, *big.Int, error) {
	return _Spseistaking.Contract.GetNormalizeAmount(&_Spseistaking.CallOpts, amount)
}

// GetTiers is a free data retrieval call binding the contract method 0xde170570.
//
// Solidity: function getTiers() view returns(uint256[], uint256[])
func (_Spseistaking *SpseistakingCaller) GetTiers(opts *bind.CallOpts) ([]*big.Int, []*big.Int, error) {
	var out []interface{}
	err := _Spseistaking.contract.Call(opts, &out, "getTiers")

	if err != nil {
		return *new([]*big.Int), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetTiers is a free data retrieval call binding the contract method 0xde170570.
//
// Solidity: function getTiers() view returns(uint256[], uint256[])
func (_Spseistaking *SpseistakingSession) GetTiers() ([]*big.Int, []*big.Int, error) {
	return _Spseistaking.Contract.GetTiers(&_Spseistaking.CallOpts)
}

// GetTiers is a free data retrieval call binding the contract method 0xde170570.
//
// Solidity: function getTiers() view returns(uint256[], uint256[])
func (_Spseistaking *SpseistakingCallerSession) GetTiers() ([]*big.Int, []*big.Int, error) {
	return _Spseistaking.Contract.GetTiers(&_Spseistaking.CallOpts)
}

// GetTotalSei is a free data retrieval call binding the contract method 0x1afc9b7d.
//
// Solidity: function getTotalSei() view returns(uint256)
func (_Spseistaking *SpseistakingCaller) GetTotalSei(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Spseistaking.contract.Call(opts, &out, "getTotalSei")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalSei is a free data retrieval call binding the contract method 0x1afc9b7d.
//
// Solidity: function getTotalSei() view returns(uint256)
func (_Spseistaking *SpseistakingSession) GetTotalSei() (*big.Int, error) {
	return _Spseistaking.Contract.GetTotalSei(&_Spseistaking.CallOpts)
}

// GetTotalSei is a free data retrieval call binding the contract method 0x1afc9b7d.
//
// Solidity: function getTotalSei() view returns(uint256)
func (_Spseistaking *SpseistakingCallerSession) GetTotalSei() (*big.Int, error) {
	return _Spseistaking.Contract.GetTotalSei(&_Spseistaking.CallOpts)
}

// LastClaimTime is a free data retrieval call binding the contract method 0xc21bae0c.
//
// Solidity: function lastClaimTime() view returns(uint256)
func (_Spseistaking *SpseistakingCaller) LastClaimTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Spseistaking.contract.Call(opts, &out, "lastClaimTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastClaimTime is a free data retrieval call binding the contract method 0xc21bae0c.
//
// Solidity: function lastClaimTime() view returns(uint256)
func (_Spseistaking *SpseistakingSession) LastClaimTime() (*big.Int, error) {
	return _Spseistaking.Contract.LastClaimTime(&_Spseistaking.CallOpts)
}

// LastClaimTime is a free data retrieval call binding the contract method 0xc21bae0c.
//
// Solidity: function lastClaimTime() view returns(uint256)
func (_Spseistaking *SpseistakingCallerSession) LastClaimTime() (*big.Int, error) {
	return _Spseistaking.Contract.LastClaimTime(&_Spseistaking.CallOpts)
}

// LastFillTime is a free data retrieval call binding the contract method 0x0892622f.
//
// Solidity: function lastFillTime() view returns(uint256)
func (_Spseistaking *SpseistakingCaller) LastFillTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Spseistaking.contract.Call(opts, &out, "lastFillTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastFillTime is a free data retrieval call binding the contract method 0x0892622f.
//
// Solidity: function lastFillTime() view returns(uint256)
func (_Spseistaking *SpseistakingSession) LastFillTime() (*big.Int, error) {
	return _Spseistaking.Contract.LastFillTime(&_Spseistaking.CallOpts)
}

// LastFillTime is a free data retrieval call binding the contract method 0x0892622f.
//
// Solidity: function lastFillTime() view returns(uint256)
func (_Spseistaking *SpseistakingCallerSession) LastFillTime() (*big.Int, error) {
	return _Spseistaking.Contract.LastFillTime(&_Spseistaking.CallOpts)
}

// LastRefreshTime is a free data retrieval call binding the contract method 0x6d5f0c66.
//
// Solidity: function lastRefreshTime() view returns(uint256)
func (_Spseistaking *SpseistakingCaller) LastRefreshTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Spseistaking.contract.Call(opts, &out, "lastRefreshTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRefreshTime is a free data retrieval call binding the contract method 0x6d5f0c66.
//
// Solidity: function lastRefreshTime() view returns(uint256)
func (_Spseistaking *SpseistakingSession) LastRefreshTime() (*big.Int, error) {
	return _Spseistaking.Contract.LastRefreshTime(&_Spseistaking.CallOpts)
}

// LastRefreshTime is a free data retrieval call binding the contract method 0x6d5f0c66.
//
// Solidity: function lastRefreshTime() view returns(uint256)
func (_Spseistaking *SpseistakingCallerSession) LastRefreshTime() (*big.Int, error) {
	return _Spseistaking.Contract.LastRefreshTime(&_Spseistaking.CallOpts)
}

// OrderQueue is a free data retrieval call binding the contract method 0xa4b64265.
//
// Solidity: function orderQueue() view returns(address)
func (_Spseistaking *SpseistakingCaller) OrderQueue(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Spseistaking.contract.Call(opts, &out, "orderQueue")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OrderQueue is a free data retrieval call binding the contract method 0xa4b64265.
//
// Solidity: function orderQueue() view returns(address)
func (_Spseistaking *SpseistakingSession) OrderQueue() (common.Address, error) {
	return _Spseistaking.Contract.OrderQueue(&_Spseistaking.CallOpts)
}

// OrderQueue is a free data retrieval call binding the contract method 0xa4b64265.
//
// Solidity: function orderQueue() view returns(address)
func (_Spseistaking *SpseistakingCallerSession) OrderQueue() (common.Address, error) {
	return _Spseistaking.Contract.OrderQueue(&_Spseistaking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Spseistaking *SpseistakingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Spseistaking.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Spseistaking *SpseistakingSession) Owner() (common.Address, error) {
	return _Spseistaking.Contract.Owner(&_Spseistaking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Spseistaking *SpseistakingCallerSession) Owner() (common.Address, error) {
	return _Spseistaking.Contract.Owner(&_Spseistaking.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Spseistaking *SpseistakingCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Spseistaking.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Spseistaking *SpseistakingSession) Paused() (bool, error) {
	return _Spseistaking.Contract.Paused(&_Spseistaking.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Spseistaking *SpseistakingCallerSession) Paused() (bool, error) {
	return _Spseistaking.Contract.Paused(&_Spseistaking.CallOpts)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint256)
func (_Spseistaking *SpseistakingCaller) ProtocolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Spseistaking.contract.Call(opts, &out, "protocolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint256)
func (_Spseistaking *SpseistakingSession) ProtocolFee() (*big.Int, error) {
	return _Spseistaking.Contract.ProtocolFee(&_Spseistaking.CallOpts)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint256)
func (_Spseistaking *SpseistakingCallerSession) ProtocolFee() (*big.Int, error) {
	return _Spseistaking.Contract.ProtocolFee(&_Spseistaking.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Spseistaking *SpseistakingCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Spseistaking.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Spseistaking *SpseistakingSession) ProxiableUUID() ([32]byte, error) {
	return _Spseistaking.Contract.ProxiableUUID(&_Spseistaking.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Spseistaking *SpseistakingCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Spseistaking.Contract.ProxiableUUID(&_Spseistaking.CallOpts)
}

// RefreshInterval is a free data retrieval call binding the contract method 0xdcfebe51.
//
// Solidity: function refreshInterval() view returns(uint256)
func (_Spseistaking *SpseistakingCaller) RefreshInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Spseistaking.contract.Call(opts, &out, "refreshInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RefreshInterval is a free data retrieval call binding the contract method 0xdcfebe51.
//
// Solidity: function refreshInterval() view returns(uint256)
func (_Spseistaking *SpseistakingSession) RefreshInterval() (*big.Int, error) {
	return _Spseistaking.Contract.RefreshInterval(&_Spseistaking.CallOpts)
}

// RefreshInterval is a free data retrieval call binding the contract method 0xdcfebe51.
//
// Solidity: function refreshInterval() view returns(uint256)
func (_Spseistaking *SpseistakingCallerSession) RefreshInterval() (*big.Int, error) {
	return _Spseistaking.Contract.RefreshInterval(&_Spseistaking.CallOpts)
}

// RewardStorage is a free data retrieval call binding the contract method 0xfde717f6.
//
// Solidity: function rewardStorage() view returns(address)
func (_Spseistaking *SpseistakingCaller) RewardStorage(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Spseistaking.contract.Call(opts, &out, "rewardStorage")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardStorage is a free data retrieval call binding the contract method 0xfde717f6.
//
// Solidity: function rewardStorage() view returns(address)
func (_Spseistaking *SpseistakingSession) RewardStorage() (common.Address, error) {
	return _Spseistaking.Contract.RewardStorage(&_Spseistaking.CallOpts)
}

// RewardStorage is a free data retrieval call binding the contract method 0xfde717f6.
//
// Solidity: function rewardStorage() view returns(address)
func (_Spseistaking *SpseistakingCallerSession) RewardStorage() (common.Address, error) {
	return _Spseistaking.Contract.RewardStorage(&_Spseistaking.CallOpts)
}

// SpSei is a free data retrieval call binding the contract method 0xc6554d39.
//
// Solidity: function spSei() view returns(address)
func (_Spseistaking *SpseistakingCaller) SpSei(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Spseistaking.contract.Call(opts, &out, "spSei")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SpSei is a free data retrieval call binding the contract method 0xc6554d39.
//
// Solidity: function spSei() view returns(address)
func (_Spseistaking *SpseistakingSession) SpSei() (common.Address, error) {
	return _Spseistaking.Contract.SpSei(&_Spseistaking.CallOpts)
}

// SpSei is a free data retrieval call binding the contract method 0xc6554d39.
//
// Solidity: function spSei() view returns(address)
func (_Spseistaking *SpseistakingCallerSession) SpSei() (common.Address, error) {
	return _Spseistaking.Contract.SpSei(&_Spseistaking.CallOpts)
}

// SpreadFeeRate is a free data retrieval call binding the contract method 0x448ed042.
//
// Solidity: function spreadFeeRate() view returns(uint256)
func (_Spseistaking *SpseistakingCaller) SpreadFeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Spseistaking.contract.Call(opts, &out, "spreadFeeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpreadFeeRate is a free data retrieval call binding the contract method 0x448ed042.
//
// Solidity: function spreadFeeRate() view returns(uint256)
func (_Spseistaking *SpseistakingSession) SpreadFeeRate() (*big.Int, error) {
	return _Spseistaking.Contract.SpreadFeeRate(&_Spseistaking.CallOpts)
}

// SpreadFeeRate is a free data retrieval call binding the contract method 0x448ed042.
//
// Solidity: function spreadFeeRate() view returns(uint256)
func (_Spseistaking *SpseistakingCallerSession) SpreadFeeRate() (*big.Int, error) {
	return _Spseistaking.Contract.SpreadFeeRate(&_Spseistaking.CallOpts)
}

// TargetBufferBalance is a free data retrieval call binding the contract method 0x3dbc9dc0.
//
// Solidity: function targetBufferBalance() view returns(uint256)
func (_Spseistaking *SpseistakingCaller) TargetBufferBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Spseistaking.contract.Call(opts, &out, "targetBufferBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TargetBufferBalance is a free data retrieval call binding the contract method 0x3dbc9dc0.
//
// Solidity: function targetBufferBalance() view returns(uint256)
func (_Spseistaking *SpseistakingSession) TargetBufferBalance() (*big.Int, error) {
	return _Spseistaking.Contract.TargetBufferBalance(&_Spseistaking.CallOpts)
}

// TargetBufferBalance is a free data retrieval call binding the contract method 0x3dbc9dc0.
//
// Solidity: function targetBufferBalance() view returns(uint256)
func (_Spseistaking *SpseistakingCallerSession) TargetBufferBalance() (*big.Int, error) {
	return _Spseistaking.Contract.TargetBufferBalance(&_Spseistaking.CallOpts)
}

// UnstakeFeeRate is a free data retrieval call binding the contract method 0x6ae28833.
//
// Solidity: function unstakeFeeRate() view returns(uint256)
func (_Spseistaking *SpseistakingCaller) UnstakeFeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Spseistaking.contract.Call(opts, &out, "unstakeFeeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnstakeFeeRate is a free data retrieval call binding the contract method 0x6ae28833.
//
// Solidity: function unstakeFeeRate() view returns(uint256)
func (_Spseistaking *SpseistakingSession) UnstakeFeeRate() (*big.Int, error) {
	return _Spseistaking.Contract.UnstakeFeeRate(&_Spseistaking.CallOpts)
}

// UnstakeFeeRate is a free data retrieval call binding the contract method 0x6ae28833.
//
// Solidity: function unstakeFeeRate() view returns(uint256)
func (_Spseistaking *SpseistakingCallerSession) UnstakeFeeRate() (*big.Int, error) {
	return _Spseistaking.Contract.UnstakeFeeRate(&_Spseistaking.CallOpts)
}

// AddValidator is a paid mutator transaction binding the contract method 0x156be05a.
//
// Solidity: function addValidator(string validator, uint256 targetAmount) returns()
func (_Spseistaking *SpseistakingTransactor) AddValidator(opts *bind.TransactOpts, validator string, targetAmount *big.Int) (*types.Transaction, error) {
	return _Spseistaking.contract.Transact(opts, "addValidator", validator, targetAmount)
}

// AddValidator is a paid mutator transaction binding the contract method 0x156be05a.
//
// Solidity: function addValidator(string validator, uint256 targetAmount) returns()
func (_Spseistaking *SpseistakingSession) AddValidator(validator string, targetAmount *big.Int) (*types.Transaction, error) {
	return _Spseistaking.Contract.AddValidator(&_Spseistaking.TransactOpts, validator, targetAmount)
}

// AddValidator is a paid mutator transaction binding the contract method 0x156be05a.
//
// Solidity: function addValidator(string validator, uint256 targetAmount) returns()
func (_Spseistaking *SpseistakingTransactorSession) AddValidator(validator string, targetAmount *big.Int) (*types.Transaction, error) {
	return _Spseistaking.Contract.AddValidator(&_Spseistaking.TransactOpts, validator, targetAmount)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Spseistaking *SpseistakingTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Spseistaking.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Spseistaking *SpseistakingSession) Claim() (*types.Transaction, error) {
	return _Spseistaking.Contract.Claim(&_Spseistaking.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Spseistaking *SpseistakingTransactorSession) Claim() (*types.Transaction, error) {
	return _Spseistaking.Contract.Claim(&_Spseistaking.TransactOpts)
}

// ClaimProtocolFees is a paid mutator transaction binding the contract method 0x4a7d0369.
//
// Solidity: function claimProtocolFees() returns()
func (_Spseistaking *SpseistakingTransactor) ClaimProtocolFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Spseistaking.contract.Transact(opts, "claimProtocolFees")
}

// ClaimProtocolFees is a paid mutator transaction binding the contract method 0x4a7d0369.
//
// Solidity: function claimProtocolFees() returns()
func (_Spseistaking *SpseistakingSession) ClaimProtocolFees() (*types.Transaction, error) {
	return _Spseistaking.Contract.ClaimProtocolFees(&_Spseistaking.TransactOpts)
}

// ClaimProtocolFees is a paid mutator transaction binding the contract method 0x4a7d0369.
//
// Solidity: function claimProtocolFees() returns()
func (_Spseistaking *SpseistakingTransactorSession) ClaimProtocolFees() (*types.Transaction, error) {
	return _Spseistaking.Contract.ClaimProtocolFees(&_Spseistaking.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0x617847eb.
//
// Solidity: function init(address _stakingStorage, address _orderQueue, string[] _validators, uint256[] _targetAmounts) payable returns()
func (_Spseistaking *SpseistakingTransactor) Init(opts *bind.TransactOpts, _stakingStorage common.Address, _orderQueue common.Address, _validators []string, _targetAmounts []*big.Int) (*types.Transaction, error) {
	return _Spseistaking.contract.Transact(opts, "init", _stakingStorage, _orderQueue, _validators, _targetAmounts)
}

// Init is a paid mutator transaction binding the contract method 0x617847eb.
//
// Solidity: function init(address _stakingStorage, address _orderQueue, string[] _validators, uint256[] _targetAmounts) payable returns()
func (_Spseistaking *SpseistakingSession) Init(_stakingStorage common.Address, _orderQueue common.Address, _validators []string, _targetAmounts []*big.Int) (*types.Transaction, error) {
	return _Spseistaking.Contract.Init(&_Spseistaking.TransactOpts, _stakingStorage, _orderQueue, _validators, _targetAmounts)
}

// Init is a paid mutator transaction binding the contract method 0x617847eb.
//
// Solidity: function init(address _stakingStorage, address _orderQueue, string[] _validators, uint256[] _targetAmounts) payable returns()
func (_Spseistaking *SpseistakingTransactorSession) Init(_stakingStorage common.Address, _orderQueue common.Address, _validators []string, _targetAmounts []*big.Int) (*types.Transaction, error) {
	return _Spseistaking.Contract.Init(&_Spseistaking.TransactOpts, _stakingStorage, _orderQueue, _validators, _targetAmounts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Spseistaking *SpseistakingTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Spseistaking.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Spseistaking *SpseistakingSession) Initialize() (*types.Transaction, error) {
	return _Spseistaking.Contract.Initialize(&_Spseistaking.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Spseistaking *SpseistakingTransactorSession) Initialize() (*types.Transaction, error) {
	return _Spseistaking.Contract.Initialize(&_Spseistaking.TransactOpts)
}

// InstantUnstake is a paid mutator transaction binding the contract method 0x7c0696e8.
//
// Solidity: function instantUnstake(uint256 amount) returns()
func (_Spseistaking *SpseistakingTransactor) InstantUnstake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Spseistaking.contract.Transact(opts, "instantUnstake", amount)
}

// InstantUnstake is a paid mutator transaction binding the contract method 0x7c0696e8.
//
// Solidity: function instantUnstake(uint256 amount) returns()
func (_Spseistaking *SpseistakingSession) InstantUnstake(amount *big.Int) (*types.Transaction, error) {
	return _Spseistaking.Contract.InstantUnstake(&_Spseistaking.TransactOpts, amount)
}

// InstantUnstake is a paid mutator transaction binding the contract method 0x7c0696e8.
//
// Solidity: function instantUnstake(uint256 amount) returns()
func (_Spseistaking *SpseistakingTransactorSession) InstantUnstake(amount *big.Int) (*types.Transaction, error) {
	return _Spseistaking.Contract.InstantUnstake(&_Spseistaking.TransactOpts, amount)
}

// MigratePendingUnstake is a paid mutator transaction binding the contract method 0x69e2d5ea.
//
// Solidity: function migratePendingUnstake(string from, string to, uint256 amount) returns()
func (_Spseistaking *SpseistakingTransactor) MigratePendingUnstake(opts *bind.TransactOpts, from string, to string, amount *big.Int) (*types.Transaction, error) {
	return _Spseistaking.contract.Transact(opts, "migratePendingUnstake", from, to, amount)
}

// MigratePendingUnstake is a paid mutator transaction binding the contract method 0x69e2d5ea.
//
// Solidity: function migratePendingUnstake(string from, string to, uint256 amount) returns()
func (_Spseistaking *SpseistakingSession) MigratePendingUnstake(from string, to string, amount *big.Int) (*types.Transaction, error) {
	return _Spseistaking.Contract.MigratePendingUnstake(&_Spseistaking.TransactOpts, from, to, amount)
}

// MigratePendingUnstake is a paid mutator transaction binding the contract method 0x69e2d5ea.
//
// Solidity: function migratePendingUnstake(string from, string to, uint256 amount) returns()
func (_Spseistaking *SpseistakingTransactorSession) MigratePendingUnstake(from string, to string, amount *big.Int) (*types.Transaction, error) {
	return _Spseistaking.Contract.MigratePendingUnstake(&_Spseistaking.TransactOpts, from, to, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Spseistaking *SpseistakingTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Spseistaking.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Spseistaking *SpseistakingSession) Pause() (*types.Transaction, error) {
	return _Spseistaking.Contract.Pause(&_Spseistaking.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Spseistaking *SpseistakingTransactorSession) Pause() (*types.Transaction, error) {
	return _Spseistaking.Contract.Pause(&_Spseistaking.TransactOpts)
}

// Redistribute is a paid mutator transaction binding the contract method 0xa226b0f8.
//
// Solidity: function redistribute(string from, string to, uint256 amount) returns()
func (_Spseistaking *SpseistakingTransactor) Redistribute(opts *bind.TransactOpts, from string, to string, amount *big.Int) (*types.Transaction, error) {
	return _Spseistaking.contract.Transact(opts, "redistribute", from, to, amount)
}

// Redistribute is a paid mutator transaction binding the contract method 0xa226b0f8.
//
// Solidity: function redistribute(string from, string to, uint256 amount) returns()
func (_Spseistaking *SpseistakingSession) Redistribute(from string, to string, amount *big.Int) (*types.Transaction, error) {
	return _Spseistaking.Contract.Redistribute(&_Spseistaking.TransactOpts, from, to, amount)
}

// Redistribute is a paid mutator transaction binding the contract method 0xa226b0f8.
//
// Solidity: function redistribute(string from, string to, uint256 amount) returns()
func (_Spseistaking *SpseistakingTransactorSession) Redistribute(from string, to string, amount *big.Int) (*types.Transaction, error) {
	return _Spseistaking.Contract.Redistribute(&_Spseistaking.TransactOpts, from, to, amount)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x3f52cb34.
//
// Solidity: function removeValidator(string validator) returns()
func (_Spseistaking *SpseistakingTransactor) RemoveValidator(opts *bind.TransactOpts, validator string) (*types.Transaction, error) {
	return _Spseistaking.contract.Transact(opts, "removeValidator", validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x3f52cb34.
//
// Solidity: function removeValidator(string validator) returns()
func (_Spseistaking *SpseistakingSession) RemoveValidator(validator string) (*types.Transaction, error) {
	return _Spseistaking.Contract.RemoveValidator(&_Spseistaking.TransactOpts, validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x3f52cb34.
//
// Solidity: function removeValidator(string validator) returns()
func (_Spseistaking *SpseistakingTransactorSession) RemoveValidator(validator string) (*types.Transaction, error) {
	return _Spseistaking.Contract.RemoveValidator(&_Spseistaking.TransactOpts, validator)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Spseistaking *SpseistakingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Spseistaking.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Spseistaking *SpseistakingSession) RenounceOwnership() (*types.Transaction, error) {
	return _Spseistaking.Contract.RenounceOwnership(&_Spseistaking.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Spseistaking *SpseistakingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Spseistaking.Contract.RenounceOwnership(&_Spseistaking.TransactOpts)
}

// SetBufferRatio is a paid mutator transaction binding the contract method 0x35cd299e.
//
// Solidity: function setBufferRatio(uint256 _bufferRatio) returns()
func (_Spseistaking *SpseistakingTransactor) SetBufferRatio(opts *bind.TransactOpts, _bufferRatio *big.Int) (*types.Transaction, error) {
	return _Spseistaking.contract.Transact(opts, "setBufferRatio", _bufferRatio)
}

// SetBufferRatio is a paid mutator transaction binding the contract method 0x35cd299e.
//
// Solidity: function setBufferRatio(uint256 _bufferRatio) returns()
func (_Spseistaking *SpseistakingSession) SetBufferRatio(_bufferRatio *big.Int) (*types.Transaction, error) {
	return _Spseistaking.Contract.SetBufferRatio(&_Spseistaking.TransactOpts, _bufferRatio)
}

// SetBufferRatio is a paid mutator transaction binding the contract method 0x35cd299e.
//
// Solidity: function setBufferRatio(uint256 _bufferRatio) returns()
func (_Spseistaking *SpseistakingTransactorSession) SetBufferRatio(_bufferRatio *big.Int) (*types.Transaction, error) {
	return _Spseistaking.Contract.SetBufferRatio(&_Spseistaking.TransactOpts, _bufferRatio)
}

// SetRefreshInterval is a paid mutator transaction binding the contract method 0xad4c0259.
//
// Solidity: function setRefreshInterval(uint256 _refreshInterval) returns()
func (_Spseistaking *SpseistakingTransactor) SetRefreshInterval(opts *bind.TransactOpts, _refreshInterval *big.Int) (*types.Transaction, error) {
	return _Spseistaking.contract.Transact(opts, "setRefreshInterval", _refreshInterval)
}

// SetRefreshInterval is a paid mutator transaction binding the contract method 0xad4c0259.
//
// Solidity: function setRefreshInterval(uint256 _refreshInterval) returns()
func (_Spseistaking *SpseistakingSession) SetRefreshInterval(_refreshInterval *big.Int) (*types.Transaction, error) {
	return _Spseistaking.Contract.SetRefreshInterval(&_Spseistaking.TransactOpts, _refreshInterval)
}

// SetRefreshInterval is a paid mutator transaction binding the contract method 0xad4c0259.
//
// Solidity: function setRefreshInterval(uint256 _refreshInterval) returns()
func (_Spseistaking *SpseistakingTransactorSession) SetRefreshInterval(_refreshInterval *big.Int) (*types.Transaction, error) {
	return _Spseistaking.Contract.SetRefreshInterval(&_Spseistaking.TransactOpts, _refreshInterval)
}

// SetSpreadFeeRate is a paid mutator transaction binding the contract method 0xa4f82c73.
//
// Solidity: function setSpreadFeeRate(uint256 _spreadFeeRate) returns()
func (_Spseistaking *SpseistakingTransactor) SetSpreadFeeRate(opts *bind.TransactOpts, _spreadFeeRate *big.Int) (*types.Transaction, error) {
	return _Spseistaking.contract.Transact(opts, "setSpreadFeeRate", _spreadFeeRate)
}

// SetSpreadFeeRate is a paid mutator transaction binding the contract method 0xa4f82c73.
//
// Solidity: function setSpreadFeeRate(uint256 _spreadFeeRate) returns()
func (_Spseistaking *SpseistakingSession) SetSpreadFeeRate(_spreadFeeRate *big.Int) (*types.Transaction, error) {
	return _Spseistaking.Contract.SetSpreadFeeRate(&_Spseistaking.TransactOpts, _spreadFeeRate)
}

// SetSpreadFeeRate is a paid mutator transaction binding the contract method 0xa4f82c73.
//
// Solidity: function setSpreadFeeRate(uint256 _spreadFeeRate) returns()
func (_Spseistaking *SpseistakingTransactorSession) SetSpreadFeeRate(_spreadFeeRate *big.Int) (*types.Transaction, error) {
	return _Spseistaking.Contract.SetSpreadFeeRate(&_Spseistaking.TransactOpts, _spreadFeeRate)
}

// SetTiers is a paid mutator transaction binding the contract method 0x1b8ebc3c.
//
// Solidity: function setTiers(uint256[] _tiers, uint256[] _tierFeeRatios) returns()
func (_Spseistaking *SpseistakingTransactor) SetTiers(opts *bind.TransactOpts, _tiers []*big.Int, _tierFeeRatios []*big.Int) (*types.Transaction, error) {
	return _Spseistaking.contract.Transact(opts, "setTiers", _tiers, _tierFeeRatios)
}

// SetTiers is a paid mutator transaction binding the contract method 0x1b8ebc3c.
//
// Solidity: function setTiers(uint256[] _tiers, uint256[] _tierFeeRatios) returns()
func (_Spseistaking *SpseistakingSession) SetTiers(_tiers []*big.Int, _tierFeeRatios []*big.Int) (*types.Transaction, error) {
	return _Spseistaking.Contract.SetTiers(&_Spseistaking.TransactOpts, _tiers, _tierFeeRatios)
}

// SetTiers is a paid mutator transaction binding the contract method 0x1b8ebc3c.
//
// Solidity: function setTiers(uint256[] _tiers, uint256[] _tierFeeRatios) returns()
func (_Spseistaking *SpseistakingTransactorSession) SetTiers(_tiers []*big.Int, _tierFeeRatios []*big.Int) (*types.Transaction, error) {
	return _Spseistaking.Contract.SetTiers(&_Spseistaking.TransactOpts, _tiers, _tierFeeRatios)
}

// SetUnstakeFeeRate is a paid mutator transaction binding the contract method 0x774e27c3.
//
// Solidity: function setUnstakeFeeRate(uint256 _unstakeFeeRate) returns()
func (_Spseistaking *SpseistakingTransactor) SetUnstakeFeeRate(opts *bind.TransactOpts, _unstakeFeeRate *big.Int) (*types.Transaction, error) {
	return _Spseistaking.contract.Transact(opts, "setUnstakeFeeRate", _unstakeFeeRate)
}

// SetUnstakeFeeRate is a paid mutator transaction binding the contract method 0x774e27c3.
//
// Solidity: function setUnstakeFeeRate(uint256 _unstakeFeeRate) returns()
func (_Spseistaking *SpseistakingSession) SetUnstakeFeeRate(_unstakeFeeRate *big.Int) (*types.Transaction, error) {
	return _Spseistaking.Contract.SetUnstakeFeeRate(&_Spseistaking.TransactOpts, _unstakeFeeRate)
}

// SetUnstakeFeeRate is a paid mutator transaction binding the contract method 0x774e27c3.
//
// Solidity: function setUnstakeFeeRate(uint256 _unstakeFeeRate) returns()
func (_Spseistaking *SpseistakingTransactorSession) SetUnstakeFeeRate(_unstakeFeeRate *big.Int) (*types.Transaction, error) {
	return _Spseistaking.Contract.SetUnstakeFeeRate(&_Spseistaking.TransactOpts, _unstakeFeeRate)
}

// SetUnstakeValidators is a paid mutator transaction binding the contract method 0xcb4202fe.
//
// Solidity: function setUnstakeValidators(string[] _unstakeValidators) returns()
func (_Spseistaking *SpseistakingTransactor) SetUnstakeValidators(opts *bind.TransactOpts, _unstakeValidators []string) (*types.Transaction, error) {
	return _Spseistaking.contract.Transact(opts, "setUnstakeValidators", _unstakeValidators)
}

// SetUnstakeValidators is a paid mutator transaction binding the contract method 0xcb4202fe.
//
// Solidity: function setUnstakeValidators(string[] _unstakeValidators) returns()
func (_Spseistaking *SpseistakingSession) SetUnstakeValidators(_unstakeValidators []string) (*types.Transaction, error) {
	return _Spseistaking.Contract.SetUnstakeValidators(&_Spseistaking.TransactOpts, _unstakeValidators)
}

// SetUnstakeValidators is a paid mutator transaction binding the contract method 0xcb4202fe.
//
// Solidity: function setUnstakeValidators(string[] _unstakeValidators) returns()
func (_Spseistaking *SpseistakingTransactorSession) SetUnstakeValidators(_unstakeValidators []string) (*types.Transaction, error) {
	return _Spseistaking.Contract.SetUnstakeValidators(&_Spseistaking.TransactOpts, _unstakeValidators)
}

// SetValidatorPriority is a paid mutator transaction binding the contract method 0x708c12a5.
//
// Solidity: function setValidatorPriority(string[] _priorityList, bool _isInPriority) returns()
func (_Spseistaking *SpseistakingTransactor) SetValidatorPriority(opts *bind.TransactOpts, _priorityList []string, _isInPriority bool) (*types.Transaction, error) {
	return _Spseistaking.contract.Transact(opts, "setValidatorPriority", _priorityList, _isInPriority)
}

// SetValidatorPriority is a paid mutator transaction binding the contract method 0x708c12a5.
//
// Solidity: function setValidatorPriority(string[] _priorityList, bool _isInPriority) returns()
func (_Spseistaking *SpseistakingSession) SetValidatorPriority(_priorityList []string, _isInPriority bool) (*types.Transaction, error) {
	return _Spseistaking.Contract.SetValidatorPriority(&_Spseistaking.TransactOpts, _priorityList, _isInPriority)
}

// SetValidatorPriority is a paid mutator transaction binding the contract method 0x708c12a5.
//
// Solidity: function setValidatorPriority(string[] _priorityList, bool _isInPriority) returns()
func (_Spseistaking *SpseistakingTransactorSession) SetValidatorPriority(_priorityList []string, _isInPriority bool) (*types.Transaction, error) {
	return _Spseistaking.Contract.SetValidatorPriority(&_Spseistaking.TransactOpts, _priorityList, _isInPriority)
}

// SetValidatorTargetAmount is a paid mutator transaction binding the contract method 0xe30963c9.
//
// Solidity: function setValidatorTargetAmount(string[] _validators, uint256[] _targetAmount) returns()
func (_Spseistaking *SpseistakingTransactor) SetValidatorTargetAmount(opts *bind.TransactOpts, _validators []string, _targetAmount []*big.Int) (*types.Transaction, error) {
	return _Spseistaking.contract.Transact(opts, "setValidatorTargetAmount", _validators, _targetAmount)
}

// SetValidatorTargetAmount is a paid mutator transaction binding the contract method 0xe30963c9.
//
// Solidity: function setValidatorTargetAmount(string[] _validators, uint256[] _targetAmount) returns()
func (_Spseistaking *SpseistakingSession) SetValidatorTargetAmount(_validators []string, _targetAmount []*big.Int) (*types.Transaction, error) {
	return _Spseistaking.Contract.SetValidatorTargetAmount(&_Spseistaking.TransactOpts, _validators, _targetAmount)
}

// SetValidatorTargetAmount is a paid mutator transaction binding the contract method 0xe30963c9.
//
// Solidity: function setValidatorTargetAmount(string[] _validators, uint256[] _targetAmount) returns()
func (_Spseistaking *SpseistakingTransactorSession) SetValidatorTargetAmount(_validators []string, _targetAmount []*big.Int) (*types.Transaction, error) {
	return _Spseistaking.Contract.SetValidatorTargetAmount(&_Spseistaking.TransactOpts, _validators, _targetAmount)
}

// Stake is a paid mutator transaction binding the contract method 0x7fab9e46.
//
// Solidity: function stake(bool isDonation) payable returns()
func (_Spseistaking *SpseistakingTransactor) Stake(opts *bind.TransactOpts, isDonation bool) (*types.Transaction, error) {
	return _Spseistaking.contract.Transact(opts, "stake", isDonation)
}

// Stake is a paid mutator transaction binding the contract method 0x7fab9e46.
//
// Solidity: function stake(bool isDonation) payable returns()
func (_Spseistaking *SpseistakingSession) Stake(isDonation bool) (*types.Transaction, error) {
	return _Spseistaking.Contract.Stake(&_Spseistaking.TransactOpts, isDonation)
}

// Stake is a paid mutator transaction binding the contract method 0x7fab9e46.
//
// Solidity: function stake(bool isDonation) payable returns()
func (_Spseistaking *SpseistakingTransactorSession) Stake(isDonation bool) (*types.Transaction, error) {
	return _Spseistaking.Contract.Stake(&_Spseistaking.TransactOpts, isDonation)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Spseistaking *SpseistakingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Spseistaking.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Spseistaking *SpseistakingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Spseistaking.Contract.TransferOwnership(&_Spseistaking.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Spseistaking *SpseistakingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Spseistaking.Contract.TransferOwnership(&_Spseistaking.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Spseistaking *SpseistakingTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Spseistaking.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Spseistaking *SpseistakingSession) Unpause() (*types.Transaction, error) {
	return _Spseistaking.Contract.Unpause(&_Spseistaking.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Spseistaking *SpseistakingTransactorSession) Unpause() (*types.Transaction, error) {
	return _Spseistaking.Contract.Unpause(&_Spseistaking.TransactOpts)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 amount) returns()
func (_Spseistaking *SpseistakingTransactor) Unstake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Spseistaking.contract.Transact(opts, "unstake", amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 amount) returns()
func (_Spseistaking *SpseistakingSession) Unstake(amount *big.Int) (*types.Transaction, error) {
	return _Spseistaking.Contract.Unstake(&_Spseistaking.TransactOpts, amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 amount) returns()
func (_Spseistaking *SpseistakingTransactorSession) Unstake(amount *big.Int) (*types.Transaction, error) {
	return _Spseistaking.Contract.Unstake(&_Spseistaking.TransactOpts, amount)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Spseistaking *SpseistakingTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Spseistaking.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Spseistaking *SpseistakingSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Spseistaking.Contract.UpgradeToAndCall(&_Spseistaking.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Spseistaking *SpseistakingTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Spseistaking.Contract.UpgradeToAndCall(&_Spseistaking.TransactOpts, newImplementation, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Spseistaking *SpseistakingTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Spseistaking.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Spseistaking *SpseistakingSession) Receive() (*types.Transaction, error) {
	return _Spseistaking.Contract.Receive(&_Spseistaking.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Spseistaking *SpseistakingTransactorSession) Receive() (*types.Transaction, error) {
	return _Spseistaking.Contract.Receive(&_Spseistaking.TransactOpts)
}

// SpseistakingClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the Spseistaking contract.
type SpseistakingClaimedIterator struct {
	Event *SpseistakingClaimed // Event containing the contract specifics and raw log

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
func (it *SpseistakingClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpseistakingClaimed)
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
		it.Event = new(SpseistakingClaimed)
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
func (it *SpseistakingClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpseistakingClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpseistakingClaimed represents a Claimed event raised by the Spseistaking contract.
type SpseistakingClaimed struct {
	User      common.Address
	Amount    *big.Int
	Fee       *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0x9cdcf2f7714cca3508c7f0110b04a90a80a3a8dd0e35de99689db74d28c5383e.
//
// Solidity: event Claimed(address indexed user, uint256 amount, uint256 fee, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) FilterClaimed(opts *bind.FilterOpts, user []common.Address) (*SpseistakingClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Spseistaking.contract.FilterLogs(opts, "Claimed", userRule)
	if err != nil {
		return nil, err
	}
	return &SpseistakingClaimedIterator{contract: _Spseistaking.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0x9cdcf2f7714cca3508c7f0110b04a90a80a3a8dd0e35de99689db74d28c5383e.
//
// Solidity: event Claimed(address indexed user, uint256 amount, uint256 fee, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *SpseistakingClaimed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Spseistaking.contract.WatchLogs(opts, "Claimed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpseistakingClaimed)
				if err := _Spseistaking.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0x9cdcf2f7714cca3508c7f0110b04a90a80a3a8dd0e35de99689db74d28c5383e.
//
// Solidity: event Claimed(address indexed user, uint256 amount, uint256 fee, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) ParseClaimed(log types.Log) (*SpseistakingClaimed, error) {
	event := new(SpseistakingClaimed)
	if err := _Spseistaking.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpseistakingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Spseistaking contract.
type SpseistakingInitializedIterator struct {
	Event *SpseistakingInitialized // Event containing the contract specifics and raw log

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
func (it *SpseistakingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpseistakingInitialized)
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
		it.Event = new(SpseistakingInitialized)
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
func (it *SpseistakingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpseistakingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpseistakingInitialized represents a Initialized event raised by the Spseistaking contract.
type SpseistakingInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Spseistaking *SpseistakingFilterer) FilterInitialized(opts *bind.FilterOpts) (*SpseistakingInitializedIterator, error) {

	logs, sub, err := _Spseistaking.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SpseistakingInitializedIterator{contract: _Spseistaking.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Spseistaking *SpseistakingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SpseistakingInitialized) (event.Subscription, error) {

	logs, sub, err := _Spseistaking.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpseistakingInitialized)
				if err := _Spseistaking.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Spseistaking *SpseistakingFilterer) ParseInitialized(log types.Log) (*SpseistakingInitialized, error) {
	event := new(SpseistakingInitialized)
	if err := _Spseistaking.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpseistakingInstantUnstakedIterator is returned from FilterInstantUnstaked and is used to iterate over the raw logs and unpacked data for InstantUnstaked events raised by the Spseistaking contract.
type SpseistakingInstantUnstakedIterator struct {
	Event *SpseistakingInstantUnstaked // Event containing the contract specifics and raw log

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
func (it *SpseistakingInstantUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpseistakingInstantUnstaked)
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
		it.Event = new(SpseistakingInstantUnstaked)
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
func (it *SpseistakingInstantUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpseistakingInstantUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpseistakingInstantUnstaked represents a InstantUnstaked event raised by the Spseistaking contract.
type SpseistakingInstantUnstaked struct {
	User          common.Address
	Amount        *big.Int
	UnstakeAmount *big.Int
	BufferBalance *big.Int
	Fee           *big.Int
	Timestamp     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInstantUnstaked is a free log retrieval operation binding the contract event 0xaabf6b8b0ff9cc2ebf84269b4b54c159afd61041501b5f9e21fad57abda46777.
//
// Solidity: event InstantUnstaked(address indexed user, uint256 amount, uint256 unstakeAmount, uint256 bufferBalance, uint256 fee, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) FilterInstantUnstaked(opts *bind.FilterOpts, user []common.Address) (*SpseistakingInstantUnstakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Spseistaking.contract.FilterLogs(opts, "InstantUnstaked", userRule)
	if err != nil {
		return nil, err
	}
	return &SpseistakingInstantUnstakedIterator{contract: _Spseistaking.contract, event: "InstantUnstaked", logs: logs, sub: sub}, nil
}

// WatchInstantUnstaked is a free log subscription operation binding the contract event 0xaabf6b8b0ff9cc2ebf84269b4b54c159afd61041501b5f9e21fad57abda46777.
//
// Solidity: event InstantUnstaked(address indexed user, uint256 amount, uint256 unstakeAmount, uint256 bufferBalance, uint256 fee, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) WatchInstantUnstaked(opts *bind.WatchOpts, sink chan<- *SpseistakingInstantUnstaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Spseistaking.contract.WatchLogs(opts, "InstantUnstaked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpseistakingInstantUnstaked)
				if err := _Spseistaking.contract.UnpackLog(event, "InstantUnstaked", log); err != nil {
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

// ParseInstantUnstaked is a log parse operation binding the contract event 0xaabf6b8b0ff9cc2ebf84269b4b54c159afd61041501b5f9e21fad57abda46777.
//
// Solidity: event InstantUnstaked(address indexed user, uint256 amount, uint256 unstakeAmount, uint256 bufferBalance, uint256 fee, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) ParseInstantUnstaked(log types.Log) (*SpseistakingInstantUnstaked, error) {
	event := new(SpseistakingInstantUnstaked)
	if err := _Spseistaking.contract.UnpackLog(event, "InstantUnstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpseistakingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Spseistaking contract.
type SpseistakingOwnershipTransferredIterator struct {
	Event *SpseistakingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SpseistakingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpseistakingOwnershipTransferred)
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
		it.Event = new(SpseistakingOwnershipTransferred)
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
func (it *SpseistakingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpseistakingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpseistakingOwnershipTransferred represents a OwnershipTransferred event raised by the Spseistaking contract.
type SpseistakingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Spseistaking *SpseistakingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SpseistakingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Spseistaking.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SpseistakingOwnershipTransferredIterator{contract: _Spseistaking.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Spseistaking *SpseistakingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SpseistakingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Spseistaking.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpseistakingOwnershipTransferred)
				if err := _Spseistaking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Spseistaking *SpseistakingFilterer) ParseOwnershipTransferred(log types.Log) (*SpseistakingOwnershipTransferred, error) {
	event := new(SpseistakingOwnershipTransferred)
	if err := _Spseistaking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpseistakingParamsUpdatedIterator is returned from FilterParamsUpdated and is used to iterate over the raw logs and unpacked data for ParamsUpdated events raised by the Spseistaking contract.
type SpseistakingParamsUpdatedIterator struct {
	Event *SpseistakingParamsUpdated // Event containing the contract specifics and raw log

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
func (it *SpseistakingParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpseistakingParamsUpdated)
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
		it.Event = new(SpseistakingParamsUpdated)
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
func (it *SpseistakingParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpseistakingParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpseistakingParamsUpdated represents a ParamsUpdated event raised by the Spseistaking contract.
type SpseistakingParamsUpdated struct {
	Param     string
	Value     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterParamsUpdated is a free log retrieval operation binding the contract event 0xc194909a05d344d541a4418f8a1afcb8c4a13f35e87d3faf2b4d1f736a2f00b6.
//
// Solidity: event ParamsUpdated(string param, uint256 value, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) FilterParamsUpdated(opts *bind.FilterOpts) (*SpseistakingParamsUpdatedIterator, error) {

	logs, sub, err := _Spseistaking.contract.FilterLogs(opts, "ParamsUpdated")
	if err != nil {
		return nil, err
	}
	return &SpseistakingParamsUpdatedIterator{contract: _Spseistaking.contract, event: "ParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchParamsUpdated is a free log subscription operation binding the contract event 0xc194909a05d344d541a4418f8a1afcb8c4a13f35e87d3faf2b4d1f736a2f00b6.
//
// Solidity: event ParamsUpdated(string param, uint256 value, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) WatchParamsUpdated(opts *bind.WatchOpts, sink chan<- *SpseistakingParamsUpdated) (event.Subscription, error) {

	logs, sub, err := _Spseistaking.contract.WatchLogs(opts, "ParamsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpseistakingParamsUpdated)
				if err := _Spseistaking.contract.UnpackLog(event, "ParamsUpdated", log); err != nil {
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

// ParseParamsUpdated is a log parse operation binding the contract event 0xc194909a05d344d541a4418f8a1afcb8c4a13f35e87d3faf2b4d1f736a2f00b6.
//
// Solidity: event ParamsUpdated(string param, uint256 value, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) ParseParamsUpdated(log types.Log) (*SpseistakingParamsUpdated, error) {
	event := new(SpseistakingParamsUpdated)
	if err := _Spseistaking.contract.UnpackLog(event, "ParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpseistakingPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Spseistaking contract.
type SpseistakingPausedIterator struct {
	Event *SpseistakingPaused // Event containing the contract specifics and raw log

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
func (it *SpseistakingPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpseistakingPaused)
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
		it.Event = new(SpseistakingPaused)
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
func (it *SpseistakingPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpseistakingPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpseistakingPaused represents a Paused event raised by the Spseistaking contract.
type SpseistakingPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Spseistaking *SpseistakingFilterer) FilterPaused(opts *bind.FilterOpts) (*SpseistakingPausedIterator, error) {

	logs, sub, err := _Spseistaking.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &SpseistakingPausedIterator{contract: _Spseistaking.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Spseistaking *SpseistakingFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *SpseistakingPaused) (event.Subscription, error) {

	logs, sub, err := _Spseistaking.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpseistakingPaused)
				if err := _Spseistaking.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Spseistaking *SpseistakingFilterer) ParsePaused(log types.Log) (*SpseistakingPaused, error) {
	event := new(SpseistakingPaused)
	if err := _Spseistaking.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpseistakingPendingUnstakeMigratedIterator is returned from FilterPendingUnstakeMigrated and is used to iterate over the raw logs and unpacked data for PendingUnstakeMigrated events raised by the Spseistaking contract.
type SpseistakingPendingUnstakeMigratedIterator struct {
	Event *SpseistakingPendingUnstakeMigrated // Event containing the contract specifics and raw log

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
func (it *SpseistakingPendingUnstakeMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpseistakingPendingUnstakeMigrated)
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
		it.Event = new(SpseistakingPendingUnstakeMigrated)
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
func (it *SpseistakingPendingUnstakeMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpseistakingPendingUnstakeMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpseistakingPendingUnstakeMigrated represents a PendingUnstakeMigrated event raised by the Spseistaking contract.
type SpseistakingPendingUnstakeMigrated struct {
	From      string
	To        string
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPendingUnstakeMigrated is a free log retrieval operation binding the contract event 0xca8e8158b56ce26e206590403f11ee7d59927b8217700331c156352fd5d11078.
//
// Solidity: event PendingUnstakeMigrated(string from, string to, uint256 amount, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) FilterPendingUnstakeMigrated(opts *bind.FilterOpts) (*SpseistakingPendingUnstakeMigratedIterator, error) {

	logs, sub, err := _Spseistaking.contract.FilterLogs(opts, "PendingUnstakeMigrated")
	if err != nil {
		return nil, err
	}
	return &SpseistakingPendingUnstakeMigratedIterator{contract: _Spseistaking.contract, event: "PendingUnstakeMigrated", logs: logs, sub: sub}, nil
}

// WatchPendingUnstakeMigrated is a free log subscription operation binding the contract event 0xca8e8158b56ce26e206590403f11ee7d59927b8217700331c156352fd5d11078.
//
// Solidity: event PendingUnstakeMigrated(string from, string to, uint256 amount, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) WatchPendingUnstakeMigrated(opts *bind.WatchOpts, sink chan<- *SpseistakingPendingUnstakeMigrated) (event.Subscription, error) {

	logs, sub, err := _Spseistaking.contract.WatchLogs(opts, "PendingUnstakeMigrated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpseistakingPendingUnstakeMigrated)
				if err := _Spseistaking.contract.UnpackLog(event, "PendingUnstakeMigrated", log); err != nil {
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

// ParsePendingUnstakeMigrated is a log parse operation binding the contract event 0xca8e8158b56ce26e206590403f11ee7d59927b8217700331c156352fd5d11078.
//
// Solidity: event PendingUnstakeMigrated(string from, string to, uint256 amount, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) ParsePendingUnstakeMigrated(log types.Log) (*SpseistakingPendingUnstakeMigrated, error) {
	event := new(SpseistakingPendingUnstakeMigrated)
	if err := _Spseistaking.contract.UnpackLog(event, "PendingUnstakeMigrated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpseistakingProtocolFeesClaimedIterator is returned from FilterProtocolFeesClaimed and is used to iterate over the raw logs and unpacked data for ProtocolFeesClaimed events raised by the Spseistaking contract.
type SpseistakingProtocolFeesClaimedIterator struct {
	Event *SpseistakingProtocolFeesClaimed // Event containing the contract specifics and raw log

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
func (it *SpseistakingProtocolFeesClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpseistakingProtocolFeesClaimed)
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
		it.Event = new(SpseistakingProtocolFeesClaimed)
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
func (it *SpseistakingProtocolFeesClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpseistakingProtocolFeesClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpseistakingProtocolFeesClaimed represents a ProtocolFeesClaimed event raised by the Spseistaking contract.
type SpseistakingProtocolFeesClaimed struct {
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeesClaimed is a free log retrieval operation binding the contract event 0xa183fe3fab3451e4394f6ec7bc6c71c823d891df2df809a3c0e87d104408fa25.
//
// Solidity: event ProtocolFeesClaimed(uint256 amount, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) FilterProtocolFeesClaimed(opts *bind.FilterOpts) (*SpseistakingProtocolFeesClaimedIterator, error) {

	logs, sub, err := _Spseistaking.contract.FilterLogs(opts, "ProtocolFeesClaimed")
	if err != nil {
		return nil, err
	}
	return &SpseistakingProtocolFeesClaimedIterator{contract: _Spseistaking.contract, event: "ProtocolFeesClaimed", logs: logs, sub: sub}, nil
}

// WatchProtocolFeesClaimed is a free log subscription operation binding the contract event 0xa183fe3fab3451e4394f6ec7bc6c71c823d891df2df809a3c0e87d104408fa25.
//
// Solidity: event ProtocolFeesClaimed(uint256 amount, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) WatchProtocolFeesClaimed(opts *bind.WatchOpts, sink chan<- *SpseistakingProtocolFeesClaimed) (event.Subscription, error) {

	logs, sub, err := _Spseistaking.contract.WatchLogs(opts, "ProtocolFeesClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpseistakingProtocolFeesClaimed)
				if err := _Spseistaking.contract.UnpackLog(event, "ProtocolFeesClaimed", log); err != nil {
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

// ParseProtocolFeesClaimed is a log parse operation binding the contract event 0xa183fe3fab3451e4394f6ec7bc6c71c823d891df2df809a3c0e87d104408fa25.
//
// Solidity: event ProtocolFeesClaimed(uint256 amount, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) ParseProtocolFeesClaimed(log types.Log) (*SpseistakingProtocolFeesClaimed, error) {
	event := new(SpseistakingProtocolFeesClaimed)
	if err := _Spseistaking.contract.UnpackLog(event, "ProtocolFeesClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpseistakingRedistributedIterator is returned from FilterRedistributed and is used to iterate over the raw logs and unpacked data for Redistributed events raised by the Spseistaking contract.
type SpseistakingRedistributedIterator struct {
	Event *SpseistakingRedistributed // Event containing the contract specifics and raw log

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
func (it *SpseistakingRedistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpseistakingRedistributed)
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
		it.Event = new(SpseistakingRedistributed)
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
func (it *SpseistakingRedistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpseistakingRedistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpseistakingRedistributed represents a Redistributed event raised by the Spseistaking contract.
type SpseistakingRedistributed struct {
	From      string
	To        string
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRedistributed is a free log retrieval operation binding the contract event 0x2ada4ee9ac5765ba568c5973da7dbedbbc1a149cf01864a1d9f577ba93308136.
//
// Solidity: event Redistributed(string from, string to, uint256 amount, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) FilterRedistributed(opts *bind.FilterOpts) (*SpseistakingRedistributedIterator, error) {

	logs, sub, err := _Spseistaking.contract.FilterLogs(opts, "Redistributed")
	if err != nil {
		return nil, err
	}
	return &SpseistakingRedistributedIterator{contract: _Spseistaking.contract, event: "Redistributed", logs: logs, sub: sub}, nil
}

// WatchRedistributed is a free log subscription operation binding the contract event 0x2ada4ee9ac5765ba568c5973da7dbedbbc1a149cf01864a1d9f577ba93308136.
//
// Solidity: event Redistributed(string from, string to, uint256 amount, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) WatchRedistributed(opts *bind.WatchOpts, sink chan<- *SpseistakingRedistributed) (event.Subscription, error) {

	logs, sub, err := _Spseistaking.contract.WatchLogs(opts, "Redistributed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpseistakingRedistributed)
				if err := _Spseistaking.contract.UnpackLog(event, "Redistributed", log); err != nil {
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

// ParseRedistributed is a log parse operation binding the contract event 0x2ada4ee9ac5765ba568c5973da7dbedbbc1a149cf01864a1d9f577ba93308136.
//
// Solidity: event Redistributed(string from, string to, uint256 amount, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) ParseRedistributed(log types.Log) (*SpseistakingRedistributed, error) {
	event := new(SpseistakingRedistributed)
	if err := _Spseistaking.contract.UnpackLog(event, "Redistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpseistakingRefreshedIterator is returned from FilterRefreshed and is used to iterate over the raw logs and unpacked data for Refreshed events raised by the Spseistaking contract.
type SpseistakingRefreshedIterator struct {
	Event *SpseistakingRefreshed // Event containing the contract specifics and raw log

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
func (it *SpseistakingRefreshedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpseistakingRefreshed)
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
		it.Event = new(SpseistakingRefreshed)
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
func (it *SpseistakingRefreshedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpseistakingRefreshedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpseistakingRefreshed represents a Refreshed event raised by the Spseistaking contract.
type SpseistakingRefreshed struct {
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRefreshed is a free log retrieval operation binding the contract event 0x2cfeea4f295d5c39c4838499d1635915f420852e4bc3ea07ced148e33b57f286.
//
// Solidity: event Refreshed(uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) FilterRefreshed(opts *bind.FilterOpts) (*SpseistakingRefreshedIterator, error) {

	logs, sub, err := _Spseistaking.contract.FilterLogs(opts, "Refreshed")
	if err != nil {
		return nil, err
	}
	return &SpseistakingRefreshedIterator{contract: _Spseistaking.contract, event: "Refreshed", logs: logs, sub: sub}, nil
}

// WatchRefreshed is a free log subscription operation binding the contract event 0x2cfeea4f295d5c39c4838499d1635915f420852e4bc3ea07ced148e33b57f286.
//
// Solidity: event Refreshed(uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) WatchRefreshed(opts *bind.WatchOpts, sink chan<- *SpseistakingRefreshed) (event.Subscription, error) {

	logs, sub, err := _Spseistaking.contract.WatchLogs(opts, "Refreshed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpseistakingRefreshed)
				if err := _Spseistaking.contract.UnpackLog(event, "Refreshed", log); err != nil {
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

// ParseRefreshed is a log parse operation binding the contract event 0x2cfeea4f295d5c39c4838499d1635915f420852e4bc3ea07ced148e33b57f286.
//
// Solidity: event Refreshed(uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) ParseRefreshed(log types.Log) (*SpseistakingRefreshed, error) {
	event := new(SpseistakingRefreshed)
	if err := _Spseistaking.contract.UnpackLog(event, "Refreshed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpseistakingReinvestedIterator is returned from FilterReinvested and is used to iterate over the raw logs and unpacked data for Reinvested events raised by the Spseistaking contract.
type SpseistakingReinvestedIterator struct {
	Event *SpseistakingReinvested // Event containing the contract specifics and raw log

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
func (it *SpseistakingReinvestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpseistakingReinvested)
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
		it.Event = new(SpseistakingReinvested)
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
func (it *SpseistakingReinvestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpseistakingReinvestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpseistakingReinvested represents a Reinvested event raised by the Spseistaking contract.
type SpseistakingReinvested struct {
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReinvested is a free log retrieval operation binding the contract event 0x61d66d8767ff702762e19c18b3cf4a23e2eef70b62fde2720244871ba6d796c7.
//
// Solidity: event Reinvested(uint256 amount, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) FilterReinvested(opts *bind.FilterOpts) (*SpseistakingReinvestedIterator, error) {

	logs, sub, err := _Spseistaking.contract.FilterLogs(opts, "Reinvested")
	if err != nil {
		return nil, err
	}
	return &SpseistakingReinvestedIterator{contract: _Spseistaking.contract, event: "Reinvested", logs: logs, sub: sub}, nil
}

// WatchReinvested is a free log subscription operation binding the contract event 0x61d66d8767ff702762e19c18b3cf4a23e2eef70b62fde2720244871ba6d796c7.
//
// Solidity: event Reinvested(uint256 amount, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) WatchReinvested(opts *bind.WatchOpts, sink chan<- *SpseistakingReinvested) (event.Subscription, error) {

	logs, sub, err := _Spseistaking.contract.WatchLogs(opts, "Reinvested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpseistakingReinvested)
				if err := _Spseistaking.contract.UnpackLog(event, "Reinvested", log); err != nil {
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

// ParseReinvested is a log parse operation binding the contract event 0x61d66d8767ff702762e19c18b3cf4a23e2eef70b62fde2720244871ba6d796c7.
//
// Solidity: event Reinvested(uint256 amount, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) ParseReinvested(log types.Log) (*SpseistakingReinvested, error) {
	event := new(SpseistakingReinvested)
	if err := _Spseistaking.contract.UnpackLog(event, "Reinvested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpseistakingStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Spseistaking contract.
type SpseistakingStakedIterator struct {
	Event *SpseistakingStaked // Event containing the contract specifics and raw log

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
func (it *SpseistakingStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpseistakingStaked)
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
		it.Event = new(SpseistakingStaked)
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
func (it *SpseistakingStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpseistakingStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpseistakingStaked represents a Staked event raised by the Spseistaking contract.
type SpseistakingStaked struct {
	User                 common.Address
	InAmount             *big.Int
	LstAmount            *big.Int
	StakeAmount          *big.Int
	BufferBalance        *big.Int
	CurrentExchangeRatio *big.Int
	Timestamp            *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0xc16be9a586414a157dd46b4d023aa9997a025dd1cbbaa67ac0c1b8273a5eaf55.
//
// Solidity: event Staked(address indexed user, uint256 inAmount, uint256 lstAmount, uint256 stakeAmount, uint256 bufferBalance, uint256 currentExchangeRatio, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*SpseistakingStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Spseistaking.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &SpseistakingStakedIterator{contract: _Spseistaking.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0xc16be9a586414a157dd46b4d023aa9997a025dd1cbbaa67ac0c1b8273a5eaf55.
//
// Solidity: event Staked(address indexed user, uint256 inAmount, uint256 lstAmount, uint256 stakeAmount, uint256 bufferBalance, uint256 currentExchangeRatio, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *SpseistakingStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Spseistaking.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpseistakingStaked)
				if err := _Spseistaking.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0xc16be9a586414a157dd46b4d023aa9997a025dd1cbbaa67ac0c1b8273a5eaf55.
//
// Solidity: event Staked(address indexed user, uint256 inAmount, uint256 lstAmount, uint256 stakeAmount, uint256 bufferBalance, uint256 currentExchangeRatio, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) ParseStaked(log types.Log) (*SpseistakingStaked, error) {
	event := new(SpseistakingStaked)
	if err := _Spseistaking.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpseistakingTiersSetIterator is returned from FilterTiersSet and is used to iterate over the raw logs and unpacked data for TiersSet events raised by the Spseistaking contract.
type SpseistakingTiersSetIterator struct {
	Event *SpseistakingTiersSet // Event containing the contract specifics and raw log

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
func (it *SpseistakingTiersSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpseistakingTiersSet)
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
		it.Event = new(SpseistakingTiersSet)
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
func (it *SpseistakingTiersSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpseistakingTiersSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpseistakingTiersSet represents a TiersSet event raised by the Spseistaking contract.
type SpseistakingTiersSet struct {
	Tiers         []*big.Int
	TierFeeRatios []*big.Int
	Timestamp     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTiersSet is a free log retrieval operation binding the contract event 0x18136faa580e170ea2de80374b9e4ebb0010f1f8d372246dfe30a4df3d2e0bf1.
//
// Solidity: event TiersSet(uint256[] tiers, uint256[] tierFeeRatios, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) FilterTiersSet(opts *bind.FilterOpts) (*SpseistakingTiersSetIterator, error) {

	logs, sub, err := _Spseistaking.contract.FilterLogs(opts, "TiersSet")
	if err != nil {
		return nil, err
	}
	return &SpseistakingTiersSetIterator{contract: _Spseistaking.contract, event: "TiersSet", logs: logs, sub: sub}, nil
}

// WatchTiersSet is a free log subscription operation binding the contract event 0x18136faa580e170ea2de80374b9e4ebb0010f1f8d372246dfe30a4df3d2e0bf1.
//
// Solidity: event TiersSet(uint256[] tiers, uint256[] tierFeeRatios, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) WatchTiersSet(opts *bind.WatchOpts, sink chan<- *SpseistakingTiersSet) (event.Subscription, error) {

	logs, sub, err := _Spseistaking.contract.WatchLogs(opts, "TiersSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpseistakingTiersSet)
				if err := _Spseistaking.contract.UnpackLog(event, "TiersSet", log); err != nil {
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

// ParseTiersSet is a log parse operation binding the contract event 0x18136faa580e170ea2de80374b9e4ebb0010f1f8d372246dfe30a4df3d2e0bf1.
//
// Solidity: event TiersSet(uint256[] tiers, uint256[] tierFeeRatios, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) ParseTiersSet(log types.Log) (*SpseistakingTiersSet, error) {
	event := new(SpseistakingTiersSet)
	if err := _Spseistaking.contract.UnpackLog(event, "TiersSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpseistakingUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Spseistaking contract.
type SpseistakingUnpausedIterator struct {
	Event *SpseistakingUnpaused // Event containing the contract specifics and raw log

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
func (it *SpseistakingUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpseistakingUnpaused)
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
		it.Event = new(SpseistakingUnpaused)
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
func (it *SpseistakingUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpseistakingUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpseistakingUnpaused represents a Unpaused event raised by the Spseistaking contract.
type SpseistakingUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Spseistaking *SpseistakingFilterer) FilterUnpaused(opts *bind.FilterOpts) (*SpseistakingUnpausedIterator, error) {

	logs, sub, err := _Spseistaking.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &SpseistakingUnpausedIterator{contract: _Spseistaking.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Spseistaking *SpseistakingFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *SpseistakingUnpaused) (event.Subscription, error) {

	logs, sub, err := _Spseistaking.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpseistakingUnpaused)
				if err := _Spseistaking.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Spseistaking *SpseistakingFilterer) ParseUnpaused(log types.Log) (*SpseistakingUnpaused, error) {
	event := new(SpseistakingUnpaused)
	if err := _Spseistaking.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpseistakingUnstakeValidatorsSetIterator is returned from FilterUnstakeValidatorsSet and is used to iterate over the raw logs and unpacked data for UnstakeValidatorsSet events raised by the Spseistaking contract.
type SpseistakingUnstakeValidatorsSetIterator struct {
	Event *SpseistakingUnstakeValidatorsSet // Event containing the contract specifics and raw log

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
func (it *SpseistakingUnstakeValidatorsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpseistakingUnstakeValidatorsSet)
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
		it.Event = new(SpseistakingUnstakeValidatorsSet)
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
func (it *SpseistakingUnstakeValidatorsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpseistakingUnstakeValidatorsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpseistakingUnstakeValidatorsSet represents a UnstakeValidatorsSet event raised by the Spseistaking contract.
type SpseistakingUnstakeValidatorsSet struct {
	UnstakeValidators []string
	Timestamp         *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUnstakeValidatorsSet is a free log retrieval operation binding the contract event 0xd751653b41e615bb0b2459b08f5cd8e0ace710db9ded60bc1304d7f61ecc8c8e.
//
// Solidity: event UnstakeValidatorsSet(string[] unstakeValidators, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) FilterUnstakeValidatorsSet(opts *bind.FilterOpts) (*SpseistakingUnstakeValidatorsSetIterator, error) {

	logs, sub, err := _Spseistaking.contract.FilterLogs(opts, "UnstakeValidatorsSet")
	if err != nil {
		return nil, err
	}
	return &SpseistakingUnstakeValidatorsSetIterator{contract: _Spseistaking.contract, event: "UnstakeValidatorsSet", logs: logs, sub: sub}, nil
}

// WatchUnstakeValidatorsSet is a free log subscription operation binding the contract event 0xd751653b41e615bb0b2459b08f5cd8e0ace710db9ded60bc1304d7f61ecc8c8e.
//
// Solidity: event UnstakeValidatorsSet(string[] unstakeValidators, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) WatchUnstakeValidatorsSet(opts *bind.WatchOpts, sink chan<- *SpseistakingUnstakeValidatorsSet) (event.Subscription, error) {

	logs, sub, err := _Spseistaking.contract.WatchLogs(opts, "UnstakeValidatorsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpseistakingUnstakeValidatorsSet)
				if err := _Spseistaking.contract.UnpackLog(event, "UnstakeValidatorsSet", log); err != nil {
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

// ParseUnstakeValidatorsSet is a log parse operation binding the contract event 0xd751653b41e615bb0b2459b08f5cd8e0ace710db9ded60bc1304d7f61ecc8c8e.
//
// Solidity: event UnstakeValidatorsSet(string[] unstakeValidators, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) ParseUnstakeValidatorsSet(log types.Log) (*SpseistakingUnstakeValidatorsSet, error) {
	event := new(SpseistakingUnstakeValidatorsSet)
	if err := _Spseistaking.contract.UnpackLog(event, "UnstakeValidatorsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpseistakingUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the Spseistaking contract.
type SpseistakingUnstakedIterator struct {
	Event *SpseistakingUnstaked // Event containing the contract specifics and raw log

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
func (it *SpseistakingUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpseistakingUnstaked)
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
		it.Event = new(SpseistakingUnstaked)
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
func (it *SpseistakingUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpseistakingUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpseistakingUnstaked represents a Unstaked event raised by the Spseistaking contract.
type SpseistakingUnstaked struct {
	User                 common.Address
	Amount               *big.Int
	UnstakeAmount        *big.Int
	CurrentExchangeRatio *big.Int
	Timestamp            *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0xdcfd2b4017d03f7e541021db793b2f9b31e4acdee005f789e52853c390e3e962.
//
// Solidity: event Unstaked(address indexed user, uint256 amount, uint256 unstakeAmount, uint256 currentExchangeRatio, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) FilterUnstaked(opts *bind.FilterOpts, user []common.Address) (*SpseistakingUnstakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Spseistaking.contract.FilterLogs(opts, "Unstaked", userRule)
	if err != nil {
		return nil, err
	}
	return &SpseistakingUnstakedIterator{contract: _Spseistaking.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0xdcfd2b4017d03f7e541021db793b2f9b31e4acdee005f789e52853c390e3e962.
//
// Solidity: event Unstaked(address indexed user, uint256 amount, uint256 unstakeAmount, uint256 currentExchangeRatio, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *SpseistakingUnstaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Spseistaking.contract.WatchLogs(opts, "Unstaked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpseistakingUnstaked)
				if err := _Spseistaking.contract.UnpackLog(event, "Unstaked", log); err != nil {
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

// ParseUnstaked is a log parse operation binding the contract event 0xdcfd2b4017d03f7e541021db793b2f9b31e4acdee005f789e52853c390e3e962.
//
// Solidity: event Unstaked(address indexed user, uint256 amount, uint256 unstakeAmount, uint256 currentExchangeRatio, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) ParseUnstaked(log types.Log) (*SpseistakingUnstaked, error) {
	event := new(SpseistakingUnstaked)
	if err := _Spseistaking.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpseistakingUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Spseistaking contract.
type SpseistakingUpgradedIterator struct {
	Event *SpseistakingUpgraded // Event containing the contract specifics and raw log

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
func (it *SpseistakingUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpseistakingUpgraded)
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
		it.Event = new(SpseistakingUpgraded)
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
func (it *SpseistakingUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpseistakingUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpseistakingUpgraded represents a Upgraded event raised by the Spseistaking contract.
type SpseistakingUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Spseistaking *SpseistakingFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*SpseistakingUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Spseistaking.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &SpseistakingUpgradedIterator{contract: _Spseistaking.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Spseistaking *SpseistakingFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *SpseistakingUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Spseistaking.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpseistakingUpgraded)
				if err := _Spseistaking.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Spseistaking *SpseistakingFilterer) ParseUpgraded(log types.Log) (*SpseistakingUpgraded, error) {
	event := new(SpseistakingUpgraded)
	if err := _Spseistaking.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpseistakingValidatorAddedIterator is returned from FilterValidatorAdded and is used to iterate over the raw logs and unpacked data for ValidatorAdded events raised by the Spseistaking contract.
type SpseistakingValidatorAddedIterator struct {
	Event *SpseistakingValidatorAdded // Event containing the contract specifics and raw log

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
func (it *SpseistakingValidatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpseistakingValidatorAdded)
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
		it.Event = new(SpseistakingValidatorAdded)
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
func (it *SpseistakingValidatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpseistakingValidatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpseistakingValidatorAdded represents a ValidatorAdded event raised by the Spseistaking contract.
type SpseistakingValidatorAdded struct {
	Validator    string
	TargetAmount *big.Int
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidatorAdded is a free log retrieval operation binding the contract event 0xd821982dda62ce3db9a315253d7be0e05252af13522413d6880d84365491a68b.
//
// Solidity: event ValidatorAdded(string validator, uint256 targetAmount, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) FilterValidatorAdded(opts *bind.FilterOpts) (*SpseistakingValidatorAddedIterator, error) {

	logs, sub, err := _Spseistaking.contract.FilterLogs(opts, "ValidatorAdded")
	if err != nil {
		return nil, err
	}
	return &SpseistakingValidatorAddedIterator{contract: _Spseistaking.contract, event: "ValidatorAdded", logs: logs, sub: sub}, nil
}

// WatchValidatorAdded is a free log subscription operation binding the contract event 0xd821982dda62ce3db9a315253d7be0e05252af13522413d6880d84365491a68b.
//
// Solidity: event ValidatorAdded(string validator, uint256 targetAmount, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) WatchValidatorAdded(opts *bind.WatchOpts, sink chan<- *SpseistakingValidatorAdded) (event.Subscription, error) {

	logs, sub, err := _Spseistaking.contract.WatchLogs(opts, "ValidatorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpseistakingValidatorAdded)
				if err := _Spseistaking.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
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

// ParseValidatorAdded is a log parse operation binding the contract event 0xd821982dda62ce3db9a315253d7be0e05252af13522413d6880d84365491a68b.
//
// Solidity: event ValidatorAdded(string validator, uint256 targetAmount, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) ParseValidatorAdded(log types.Log) (*SpseistakingValidatorAdded, error) {
	event := new(SpseistakingValidatorAdded)
	if err := _Spseistaking.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpseistakingValidatorPrioritySetIterator is returned from FilterValidatorPrioritySet and is used to iterate over the raw logs and unpacked data for ValidatorPrioritySet events raised by the Spseistaking contract.
type SpseistakingValidatorPrioritySetIterator struct {
	Event *SpseistakingValidatorPrioritySet // Event containing the contract specifics and raw log

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
func (it *SpseistakingValidatorPrioritySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpseistakingValidatorPrioritySet)
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
		it.Event = new(SpseistakingValidatorPrioritySet)
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
func (it *SpseistakingValidatorPrioritySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpseistakingValidatorPrioritySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpseistakingValidatorPrioritySet represents a ValidatorPrioritySet event raised by the Spseistaking contract.
type SpseistakingValidatorPrioritySet struct {
	PriorityList []string
	IsInPriority bool
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidatorPrioritySet is a free log retrieval operation binding the contract event 0xcc62ec8622ba05b7d5a28f983b670f90e3271ccf38941ec69bfc4030ac37a9c7.
//
// Solidity: event ValidatorPrioritySet(string[] priorityList, bool isInPriority, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) FilterValidatorPrioritySet(opts *bind.FilterOpts) (*SpseistakingValidatorPrioritySetIterator, error) {

	logs, sub, err := _Spseistaking.contract.FilterLogs(opts, "ValidatorPrioritySet")
	if err != nil {
		return nil, err
	}
	return &SpseistakingValidatorPrioritySetIterator{contract: _Spseistaking.contract, event: "ValidatorPrioritySet", logs: logs, sub: sub}, nil
}

// WatchValidatorPrioritySet is a free log subscription operation binding the contract event 0xcc62ec8622ba05b7d5a28f983b670f90e3271ccf38941ec69bfc4030ac37a9c7.
//
// Solidity: event ValidatorPrioritySet(string[] priorityList, bool isInPriority, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) WatchValidatorPrioritySet(opts *bind.WatchOpts, sink chan<- *SpseistakingValidatorPrioritySet) (event.Subscription, error) {

	logs, sub, err := _Spseistaking.contract.WatchLogs(opts, "ValidatorPrioritySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpseistakingValidatorPrioritySet)
				if err := _Spseistaking.contract.UnpackLog(event, "ValidatorPrioritySet", log); err != nil {
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

// ParseValidatorPrioritySet is a log parse operation binding the contract event 0xcc62ec8622ba05b7d5a28f983b670f90e3271ccf38941ec69bfc4030ac37a9c7.
//
// Solidity: event ValidatorPrioritySet(string[] priorityList, bool isInPriority, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) ParseValidatorPrioritySet(log types.Log) (*SpseistakingValidatorPrioritySet, error) {
	event := new(SpseistakingValidatorPrioritySet)
	if err := _Spseistaking.contract.UnpackLog(event, "ValidatorPrioritySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpseistakingValidatorRemovedIterator is returned from FilterValidatorRemoved and is used to iterate over the raw logs and unpacked data for ValidatorRemoved events raised by the Spseistaking contract.
type SpseistakingValidatorRemovedIterator struct {
	Event *SpseistakingValidatorRemoved // Event containing the contract specifics and raw log

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
func (it *SpseistakingValidatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpseistakingValidatorRemoved)
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
		it.Event = new(SpseistakingValidatorRemoved)
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
func (it *SpseistakingValidatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpseistakingValidatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpseistakingValidatorRemoved represents a ValidatorRemoved event raised by the Spseistaking contract.
type SpseistakingValidatorRemoved struct {
	Validator string
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorRemoved is a free log retrieval operation binding the contract event 0x162b44184e6b5c95dfb83f541f52017151025d91613378f94997681063319580.
//
// Solidity: event ValidatorRemoved(string validator, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) FilterValidatorRemoved(opts *bind.FilterOpts) (*SpseistakingValidatorRemovedIterator, error) {

	logs, sub, err := _Spseistaking.contract.FilterLogs(opts, "ValidatorRemoved")
	if err != nil {
		return nil, err
	}
	return &SpseistakingValidatorRemovedIterator{contract: _Spseistaking.contract, event: "ValidatorRemoved", logs: logs, sub: sub}, nil
}

// WatchValidatorRemoved is a free log subscription operation binding the contract event 0x162b44184e6b5c95dfb83f541f52017151025d91613378f94997681063319580.
//
// Solidity: event ValidatorRemoved(string validator, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) WatchValidatorRemoved(opts *bind.WatchOpts, sink chan<- *SpseistakingValidatorRemoved) (event.Subscription, error) {

	logs, sub, err := _Spseistaking.contract.WatchLogs(opts, "ValidatorRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpseistakingValidatorRemoved)
				if err := _Spseistaking.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
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

// ParseValidatorRemoved is a log parse operation binding the contract event 0x162b44184e6b5c95dfb83f541f52017151025d91613378f94997681063319580.
//
// Solidity: event ValidatorRemoved(string validator, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) ParseValidatorRemoved(log types.Log) (*SpseistakingValidatorRemoved, error) {
	event := new(SpseistakingValidatorRemoved)
	if err := _Spseistaking.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpseistakingValidatorTargetAmountSetIterator is returned from FilterValidatorTargetAmountSet and is used to iterate over the raw logs and unpacked data for ValidatorTargetAmountSet events raised by the Spseistaking contract.
type SpseistakingValidatorTargetAmountSetIterator struct {
	Event *SpseistakingValidatorTargetAmountSet // Event containing the contract specifics and raw log

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
func (it *SpseistakingValidatorTargetAmountSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpseistakingValidatorTargetAmountSet)
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
		it.Event = new(SpseistakingValidatorTargetAmountSet)
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
func (it *SpseistakingValidatorTargetAmountSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpseistakingValidatorTargetAmountSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpseistakingValidatorTargetAmountSet represents a ValidatorTargetAmountSet event raised by the Spseistaking contract.
type SpseistakingValidatorTargetAmountSet struct {
	Validators    []string
	TargetAmounts []*big.Int
	Timestamp     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterValidatorTargetAmountSet is a free log retrieval operation binding the contract event 0x6dd3afe483c7cafca94f4c794c6f9f4573febb358baf8ac1f00db2a561146029.
//
// Solidity: event ValidatorTargetAmountSet(string[] validators, uint256[] targetAmounts, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) FilterValidatorTargetAmountSet(opts *bind.FilterOpts) (*SpseistakingValidatorTargetAmountSetIterator, error) {

	logs, sub, err := _Spseistaking.contract.FilterLogs(opts, "ValidatorTargetAmountSet")
	if err != nil {
		return nil, err
	}
	return &SpseistakingValidatorTargetAmountSetIterator{contract: _Spseistaking.contract, event: "ValidatorTargetAmountSet", logs: logs, sub: sub}, nil
}

// WatchValidatorTargetAmountSet is a free log subscription operation binding the contract event 0x6dd3afe483c7cafca94f4c794c6f9f4573febb358baf8ac1f00db2a561146029.
//
// Solidity: event ValidatorTargetAmountSet(string[] validators, uint256[] targetAmounts, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) WatchValidatorTargetAmountSet(opts *bind.WatchOpts, sink chan<- *SpseistakingValidatorTargetAmountSet) (event.Subscription, error) {

	logs, sub, err := _Spseistaking.contract.WatchLogs(opts, "ValidatorTargetAmountSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpseistakingValidatorTargetAmountSet)
				if err := _Spseistaking.contract.UnpackLog(event, "ValidatorTargetAmountSet", log); err != nil {
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

// ParseValidatorTargetAmountSet is a log parse operation binding the contract event 0x6dd3afe483c7cafca94f4c794c6f9f4573febb358baf8ac1f00db2a561146029.
//
// Solidity: event ValidatorTargetAmountSet(string[] validators, uint256[] targetAmounts, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) ParseValidatorTargetAmountSet(log types.Log) (*SpseistakingValidatorTargetAmountSet, error) {
	event := new(SpseistakingValidatorTargetAmountSet)
	if err := _Spseistaking.contract.UnpackLog(event, "ValidatorTargetAmountSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpseistakingWithdrawedIterator is returned from FilterWithdrawed and is used to iterate over the raw logs and unpacked data for Withdrawed events raised by the Spseistaking contract.
type SpseistakingWithdrawedIterator struct {
	Event *SpseistakingWithdrawed // Event containing the contract specifics and raw log

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
func (it *SpseistakingWithdrawedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpseistakingWithdrawed)
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
		it.Event = new(SpseistakingWithdrawed)
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
func (it *SpseistakingWithdrawedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpseistakingWithdrawedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpseistakingWithdrawed represents a Withdrawed event raised by the Spseistaking contract.
type SpseistakingWithdrawed struct {
	User      common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawed is a free log retrieval operation binding the contract event 0x4cdcd27ae88503b2d4d3034a348b78aec00eca6369f48e5002ca3df8686b9b3e.
//
// Solidity: event Withdrawed(address indexed user, uint256 amount, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) FilterWithdrawed(opts *bind.FilterOpts, user []common.Address) (*SpseistakingWithdrawedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Spseistaking.contract.FilterLogs(opts, "Withdrawed", userRule)
	if err != nil {
		return nil, err
	}
	return &SpseistakingWithdrawedIterator{contract: _Spseistaking.contract, event: "Withdrawed", logs: logs, sub: sub}, nil
}

// WatchWithdrawed is a free log subscription operation binding the contract event 0x4cdcd27ae88503b2d4d3034a348b78aec00eca6369f48e5002ca3df8686b9b3e.
//
// Solidity: event Withdrawed(address indexed user, uint256 amount, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) WatchWithdrawed(opts *bind.WatchOpts, sink chan<- *SpseistakingWithdrawed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Spseistaking.contract.WatchLogs(opts, "Withdrawed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpseistakingWithdrawed)
				if err := _Spseistaking.contract.UnpackLog(event, "Withdrawed", log); err != nil {
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

// ParseWithdrawed is a log parse operation binding the contract event 0x4cdcd27ae88503b2d4d3034a348b78aec00eca6369f48e5002ca3df8686b9b3e.
//
// Solidity: event Withdrawed(address indexed user, uint256 amount, uint256 timestamp)
func (_Spseistaking *SpseistakingFilterer) ParseWithdrawed(log types.Log) (*SpseistakingWithdrawed, error) {
	event := new(SpseistakingWithdrawed)
	if err := _Spseistaking.contract.UnpackLog(event, "Withdrawed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
