// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package satusd

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

// SatusdMetaData contains all meta data concerning the Satusd contract.
var SatusdMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxDeposit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxRedeem\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxWithdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"AssetAddedToEmissionList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"AssetRemovedFromEmissionList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pendingReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardIntegral\",\"type\":\"uint256\"}],\"name\":\"RewardIntegralForAccountUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardIntegral\",\"type\":\"uint256\"}],\"name\":\"RewardIntegralUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardIntegral\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"claimStartTime\",\"type\":\"uint32\"}],\"name\":\"RewardParamsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardVault\",\"type\":\"address\"}],\"name\":\"RewardVaultSetting\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"INITIAL_DEPOSIT_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SHARE_NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SHARE_SYMBOL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNDERLYING\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset_\",\"type\":\"address\"}],\"name\":\"addAssetToEmissionList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"claimableReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"emissionList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emissionListLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_underlying\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardVault\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isClaimStart\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdate\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset_\",\"type\":\"address\"}],\"name\":\"removeAssetFromEmissionList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resume\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"rewardIntegralFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"rewardParamsMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardIntegral\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"claimStartTime\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardVault\",\"outputs\":[{\"internalType\":\"contractIRewardVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"claimStartTime\",\"type\":\"uint32\"}],\"name\":\"setRewardParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardVault\",\"type\":\"address\"}],\"name\":\"setRewardVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"simulateDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"simulateMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"simulateRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"simulateWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"storedPendingReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SatusdABI is the input ABI used to generate the binding from.
// Deprecated: Use SatusdMetaData.ABI instead.
var SatusdABI = SatusdMetaData.ABI

// Satusd is an auto generated Go binding around an Ethereum contract.
type Satusd struct {
	SatusdCaller     // Read-only binding to the contract
	SatusdTransactor // Write-only binding to the contract
	SatusdFilterer   // Log filterer for contract events
}

// SatusdCaller is an auto generated read-only Go binding around an Ethereum contract.
type SatusdCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SatusdTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SatusdTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SatusdFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SatusdFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SatusdSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SatusdSession struct {
	Contract     *Satusd           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SatusdCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SatusdCallerSession struct {
	Contract *SatusdCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SatusdTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SatusdTransactorSession struct {
	Contract     *SatusdTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SatusdRaw is an auto generated low-level Go binding around an Ethereum contract.
type SatusdRaw struct {
	Contract *Satusd // Generic contract binding to access the raw methods on
}

// SatusdCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SatusdCallerRaw struct {
	Contract *SatusdCaller // Generic read-only contract binding to access the raw methods on
}

// SatusdTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SatusdTransactorRaw struct {
	Contract *SatusdTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSatusd creates a new instance of Satusd, bound to a specific deployed contract.
func NewSatusd(address common.Address, backend bind.ContractBackend) (*Satusd, error) {
	contract, err := bindSatusd(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Satusd{SatusdCaller: SatusdCaller{contract: contract}, SatusdTransactor: SatusdTransactor{contract: contract}, SatusdFilterer: SatusdFilterer{contract: contract}}, nil
}

// NewSatusdCaller creates a new read-only instance of Satusd, bound to a specific deployed contract.
func NewSatusdCaller(address common.Address, caller bind.ContractCaller) (*SatusdCaller, error) {
	contract, err := bindSatusd(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SatusdCaller{contract: contract}, nil
}

// NewSatusdTransactor creates a new write-only instance of Satusd, bound to a specific deployed contract.
func NewSatusdTransactor(address common.Address, transactor bind.ContractTransactor) (*SatusdTransactor, error) {
	contract, err := bindSatusd(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SatusdTransactor{contract: contract}, nil
}

// NewSatusdFilterer creates a new log filterer instance of Satusd, bound to a specific deployed contract.
func NewSatusdFilterer(address common.Address, filterer bind.ContractFilterer) (*SatusdFilterer, error) {
	contract, err := bindSatusd(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SatusdFilterer{contract: contract}, nil
}

// bindSatusd binds a generic wrapper to an already deployed contract.
func bindSatusd(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SatusdMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Satusd *SatusdRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Satusd.Contract.SatusdCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Satusd *SatusdRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Satusd.Contract.SatusdTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Satusd *SatusdRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Satusd.Contract.SatusdTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Satusd *SatusdCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Satusd.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Satusd *SatusdTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Satusd.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Satusd *SatusdTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Satusd.Contract.contract.Transact(opts, method, params...)
}

// INITIALDEPOSITLIMIT is a free data retrieval call binding the contract method 0xd16002da.
//
// Solidity: function INITIAL_DEPOSIT_LIMIT() view returns(uint256)
func (_Satusd *SatusdCaller) INITIALDEPOSITLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Satusd.contract.Call(opts, &out, "INITIAL_DEPOSIT_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITIALDEPOSITLIMIT is a free data retrieval call binding the contract method 0xd16002da.
//
// Solidity: function INITIAL_DEPOSIT_LIMIT() view returns(uint256)
func (_Satusd *SatusdSession) INITIALDEPOSITLIMIT() (*big.Int, error) {
	return _Satusd.Contract.INITIALDEPOSITLIMIT(&_Satusd.CallOpts)
}

// INITIALDEPOSITLIMIT is a free data retrieval call binding the contract method 0xd16002da.
//
// Solidity: function INITIAL_DEPOSIT_LIMIT() view returns(uint256)
func (_Satusd *SatusdCallerSession) INITIALDEPOSITLIMIT() (*big.Int, error) {
	return _Satusd.Contract.INITIALDEPOSITLIMIT(&_Satusd.CallOpts)
}

// SHARENAME is a free data retrieval call binding the contract method 0xc2aa8164.
//
// Solidity: function SHARE_NAME() view returns(string)
func (_Satusd *SatusdCaller) SHARENAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Satusd.contract.Call(opts, &out, "SHARE_NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SHARENAME is a free data retrieval call binding the contract method 0xc2aa8164.
//
// Solidity: function SHARE_NAME() view returns(string)
func (_Satusd *SatusdSession) SHARENAME() (string, error) {
	return _Satusd.Contract.SHARENAME(&_Satusd.CallOpts)
}

// SHARENAME is a free data retrieval call binding the contract method 0xc2aa8164.
//
// Solidity: function SHARE_NAME() view returns(string)
func (_Satusd *SatusdCallerSession) SHARENAME() (string, error) {
	return _Satusd.Contract.SHARENAME(&_Satusd.CallOpts)
}

// SHARESYMBOL is a free data retrieval call binding the contract method 0xe6f33f40.
//
// Solidity: function SHARE_SYMBOL() view returns(string)
func (_Satusd *SatusdCaller) SHARESYMBOL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Satusd.contract.Call(opts, &out, "SHARE_SYMBOL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SHARESYMBOL is a free data retrieval call binding the contract method 0xe6f33f40.
//
// Solidity: function SHARE_SYMBOL() view returns(string)
func (_Satusd *SatusdSession) SHARESYMBOL() (string, error) {
	return _Satusd.Contract.SHARESYMBOL(&_Satusd.CallOpts)
}

// SHARESYMBOL is a free data retrieval call binding the contract method 0xe6f33f40.
//
// Solidity: function SHARE_SYMBOL() view returns(string)
func (_Satusd *SatusdCallerSession) SHARESYMBOL() (string, error) {
	return _Satusd.Contract.SHARESYMBOL(&_Satusd.CallOpts)
}

// UNDERLYING is a free data retrieval call binding the contract method 0xc5d664c6.
//
// Solidity: function UNDERLYING() view returns(address)
func (_Satusd *SatusdCaller) UNDERLYING(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Satusd.contract.Call(opts, &out, "UNDERLYING")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UNDERLYING is a free data retrieval call binding the contract method 0xc5d664c6.
//
// Solidity: function UNDERLYING() view returns(address)
func (_Satusd *SatusdSession) UNDERLYING() (common.Address, error) {
	return _Satusd.Contract.UNDERLYING(&_Satusd.CallOpts)
}

// UNDERLYING is a free data retrieval call binding the contract method 0xc5d664c6.
//
// Solidity: function UNDERLYING() view returns(address)
func (_Satusd *SatusdCallerSession) UNDERLYING() (common.Address, error) {
	return _Satusd.Contract.UNDERLYING(&_Satusd.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Satusd *SatusdCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Satusd.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Satusd *SatusdSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Satusd.Contract.UPGRADEINTERFACEVERSION(&_Satusd.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Satusd *SatusdCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Satusd.Contract.UPGRADEINTERFACEVERSION(&_Satusd.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Satusd *SatusdCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Satusd.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Satusd *SatusdSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Satusd.Contract.Allowance(&_Satusd.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Satusd *SatusdCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Satusd.Contract.Allowance(&_Satusd.CallOpts, owner, spender)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_Satusd *SatusdCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Satusd.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_Satusd *SatusdSession) Asset() (common.Address, error) {
	return _Satusd.Contract.Asset(&_Satusd.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_Satusd *SatusdCallerSession) Asset() (common.Address, error) {
	return _Satusd.Contract.Asset(&_Satusd.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Satusd *SatusdCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Satusd.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Satusd *SatusdSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Satusd.Contract.BalanceOf(&_Satusd.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Satusd *SatusdCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Satusd.Contract.BalanceOf(&_Satusd.CallOpts, account)
}

// ClaimableReward is a free data retrieval call binding the contract method 0xd26abffa.
//
// Solidity: function claimableReward(address token, address account) view returns(uint256)
func (_Satusd *SatusdCaller) ClaimableReward(opts *bind.CallOpts, token common.Address, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Satusd.contract.Call(opts, &out, "claimableReward", token, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimableReward is a free data retrieval call binding the contract method 0xd26abffa.
//
// Solidity: function claimableReward(address token, address account) view returns(uint256)
func (_Satusd *SatusdSession) ClaimableReward(token common.Address, account common.Address) (*big.Int, error) {
	return _Satusd.Contract.ClaimableReward(&_Satusd.CallOpts, token, account)
}

// ClaimableReward is a free data retrieval call binding the contract method 0xd26abffa.
//
// Solidity: function claimableReward(address token, address account) view returns(uint256)
func (_Satusd *SatusdCallerSession) ClaimableReward(token common.Address, account common.Address) (*big.Int, error) {
	return _Satusd.Contract.ClaimableReward(&_Satusd.CallOpts, token, account)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_Satusd *SatusdCaller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Satusd.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_Satusd *SatusdSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _Satusd.Contract.ConvertToAssets(&_Satusd.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_Satusd *SatusdCallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _Satusd.Contract.ConvertToAssets(&_Satusd.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_Satusd *SatusdCaller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Satusd.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_Satusd *SatusdSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _Satusd.Contract.ConvertToShares(&_Satusd.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_Satusd *SatusdCallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _Satusd.Contract.ConvertToShares(&_Satusd.CallOpts, assets)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Satusd *SatusdCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Satusd.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Satusd *SatusdSession) Decimals() (uint8, error) {
	return _Satusd.Contract.Decimals(&_Satusd.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Satusd *SatusdCallerSession) Decimals() (uint8, error) {
	return _Satusd.Contract.Decimals(&_Satusd.CallOpts)
}

// EmissionList is a free data retrieval call binding the contract method 0x7db808e9.
//
// Solidity: function emissionList(uint256 ) view returns(address)
func (_Satusd *SatusdCaller) EmissionList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Satusd.contract.Call(opts, &out, "emissionList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EmissionList is a free data retrieval call binding the contract method 0x7db808e9.
//
// Solidity: function emissionList(uint256 ) view returns(address)
func (_Satusd *SatusdSession) EmissionList(arg0 *big.Int) (common.Address, error) {
	return _Satusd.Contract.EmissionList(&_Satusd.CallOpts, arg0)
}

// EmissionList is a free data retrieval call binding the contract method 0x7db808e9.
//
// Solidity: function emissionList(uint256 ) view returns(address)
func (_Satusd *SatusdCallerSession) EmissionList(arg0 *big.Int) (common.Address, error) {
	return _Satusd.Contract.EmissionList(&_Satusd.CallOpts, arg0)
}

// EmissionListLength is a free data retrieval call binding the contract method 0x1036f087.
//
// Solidity: function emissionListLength() view returns(uint256)
func (_Satusd *SatusdCaller) EmissionListLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Satusd.contract.Call(opts, &out, "emissionListLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EmissionListLength is a free data retrieval call binding the contract method 0x1036f087.
//
// Solidity: function emissionListLength() view returns(uint256)
func (_Satusd *SatusdSession) EmissionListLength() (*big.Int, error) {
	return _Satusd.Contract.EmissionListLength(&_Satusd.CallOpts)
}

// EmissionListLength is a free data retrieval call binding the contract method 0x1036f087.
//
// Solidity: function emissionListLength() view returns(uint256)
func (_Satusd *SatusdCallerSession) EmissionListLength() (*big.Int, error) {
	return _Satusd.Contract.EmissionListLength(&_Satusd.CallOpts)
}

// IsClaimStart is a free data retrieval call binding the contract method 0xe0427c51.
//
// Solidity: function isClaimStart(address token) view returns(bool)
func (_Satusd *SatusdCaller) IsClaimStart(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _Satusd.contract.Call(opts, &out, "isClaimStart", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsClaimStart is a free data retrieval call binding the contract method 0xe0427c51.
//
// Solidity: function isClaimStart(address token) view returns(bool)
func (_Satusd *SatusdSession) IsClaimStart(token common.Address) (bool, error) {
	return _Satusd.Contract.IsClaimStart(&_Satusd.CallOpts, token)
}

// IsClaimStart is a free data retrieval call binding the contract method 0xe0427c51.
//
// Solidity: function isClaimStart(address token) view returns(bool)
func (_Satusd *SatusdCallerSession) IsClaimStart(token common.Address) (bool, error) {
	return _Satusd.Contract.IsClaimStart(&_Satusd.CallOpts, token)
}

// LastUpdate is a free data retrieval call binding the contract method 0xc0463711.
//
// Solidity: function lastUpdate() view returns(uint32)
func (_Satusd *SatusdCaller) LastUpdate(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Satusd.contract.Call(opts, &out, "lastUpdate")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LastUpdate is a free data retrieval call binding the contract method 0xc0463711.
//
// Solidity: function lastUpdate() view returns(uint32)
func (_Satusd *SatusdSession) LastUpdate() (uint32, error) {
	return _Satusd.Contract.LastUpdate(&_Satusd.CallOpts)
}

// LastUpdate is a free data retrieval call binding the contract method 0xc0463711.
//
// Solidity: function lastUpdate() view returns(uint32)
func (_Satusd *SatusdCallerSession) LastUpdate() (uint32, error) {
	return _Satusd.Contract.LastUpdate(&_Satusd.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_Satusd *SatusdCaller) MaxDeposit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Satusd.contract.Call(opts, &out, "maxDeposit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_Satusd *SatusdSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _Satusd.Contract.MaxDeposit(&_Satusd.CallOpts, arg0)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_Satusd *SatusdCallerSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _Satusd.Contract.MaxDeposit(&_Satusd.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_Satusd *SatusdCaller) MaxMint(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Satusd.contract.Call(opts, &out, "maxMint", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_Satusd *SatusdSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _Satusd.Contract.MaxMint(&_Satusd.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_Satusd *SatusdCallerSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _Satusd.Contract.MaxMint(&_Satusd.CallOpts, arg0)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_Satusd *SatusdCaller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Satusd.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_Satusd *SatusdSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _Satusd.Contract.MaxRedeem(&_Satusd.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_Satusd *SatusdCallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _Satusd.Contract.MaxRedeem(&_Satusd.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_Satusd *SatusdCaller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Satusd.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_Satusd *SatusdSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _Satusd.Contract.MaxWithdraw(&_Satusd.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_Satusd *SatusdCallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _Satusd.Contract.MaxWithdraw(&_Satusd.CallOpts, owner)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Satusd *SatusdCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Satusd.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Satusd *SatusdSession) Name() (string, error) {
	return _Satusd.Contract.Name(&_Satusd.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Satusd *SatusdCallerSession) Name() (string, error) {
	return _Satusd.Contract.Name(&_Satusd.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Satusd *SatusdCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Satusd.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Satusd *SatusdSession) Owner() (common.Address, error) {
	return _Satusd.Contract.Owner(&_Satusd.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Satusd *SatusdCallerSession) Owner() (common.Address, error) {
	return _Satusd.Contract.Owner(&_Satusd.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Satusd *SatusdCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Satusd.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Satusd *SatusdSession) Paused() (bool, error) {
	return _Satusd.Contract.Paused(&_Satusd.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Satusd *SatusdCallerSession) Paused() (bool, error) {
	return _Satusd.Contract.Paused(&_Satusd.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_Satusd *SatusdCaller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Satusd.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_Satusd *SatusdSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _Satusd.Contract.PreviewDeposit(&_Satusd.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_Satusd *SatusdCallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _Satusd.Contract.PreviewDeposit(&_Satusd.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_Satusd *SatusdCaller) PreviewMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Satusd.contract.Call(opts, &out, "previewMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_Satusd *SatusdSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _Satusd.Contract.PreviewMint(&_Satusd.CallOpts, shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_Satusd *SatusdCallerSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _Satusd.Contract.PreviewMint(&_Satusd.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_Satusd *SatusdCaller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Satusd.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_Satusd *SatusdSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _Satusd.Contract.PreviewRedeem(&_Satusd.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_Satusd *SatusdCallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _Satusd.Contract.PreviewRedeem(&_Satusd.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_Satusd *SatusdCaller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Satusd.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_Satusd *SatusdSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _Satusd.Contract.PreviewWithdraw(&_Satusd.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_Satusd *SatusdCallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _Satusd.Contract.PreviewWithdraw(&_Satusd.CallOpts, assets)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Satusd *SatusdCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Satusd.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Satusd *SatusdSession) ProxiableUUID() ([32]byte, error) {
	return _Satusd.Contract.ProxiableUUID(&_Satusd.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Satusd *SatusdCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Satusd.Contract.ProxiableUUID(&_Satusd.CallOpts)
}

// RewardIntegralFor is a free data retrieval call binding the contract method 0x63f57777.
//
// Solidity: function rewardIntegralFor(address token, address account) view returns(uint256)
func (_Satusd *SatusdCaller) RewardIntegralFor(opts *bind.CallOpts, token common.Address, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Satusd.contract.Call(opts, &out, "rewardIntegralFor", token, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardIntegralFor is a free data retrieval call binding the contract method 0x63f57777.
//
// Solidity: function rewardIntegralFor(address token, address account) view returns(uint256)
func (_Satusd *SatusdSession) RewardIntegralFor(token common.Address, account common.Address) (*big.Int, error) {
	return _Satusd.Contract.RewardIntegralFor(&_Satusd.CallOpts, token, account)
}

// RewardIntegralFor is a free data retrieval call binding the contract method 0x63f57777.
//
// Solidity: function rewardIntegralFor(address token, address account) view returns(uint256)
func (_Satusd *SatusdCallerSession) RewardIntegralFor(token common.Address, account common.Address) (*big.Int, error) {
	return _Satusd.Contract.RewardIntegralFor(&_Satusd.CallOpts, token, account)
}

// RewardParamsMap is a free data retrieval call binding the contract method 0x227ef81f.
//
// Solidity: function rewardParamsMap(address token) view returns(uint256 rewardRate, uint256 rewardIntegral, uint32 claimStartTime)
func (_Satusd *SatusdCaller) RewardParamsMap(opts *bind.CallOpts, token common.Address) (struct {
	RewardRate     *big.Int
	RewardIntegral *big.Int
	ClaimStartTime uint32
}, error) {
	var out []interface{}
	err := _Satusd.contract.Call(opts, &out, "rewardParamsMap", token)

	outstruct := new(struct {
		RewardRate     *big.Int
		RewardIntegral *big.Int
		ClaimStartTime uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RewardRate = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardIntegral = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ClaimStartTime = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// RewardParamsMap is a free data retrieval call binding the contract method 0x227ef81f.
//
// Solidity: function rewardParamsMap(address token) view returns(uint256 rewardRate, uint256 rewardIntegral, uint32 claimStartTime)
func (_Satusd *SatusdSession) RewardParamsMap(token common.Address) (struct {
	RewardRate     *big.Int
	RewardIntegral *big.Int
	ClaimStartTime uint32
}, error) {
	return _Satusd.Contract.RewardParamsMap(&_Satusd.CallOpts, token)
}

// RewardParamsMap is a free data retrieval call binding the contract method 0x227ef81f.
//
// Solidity: function rewardParamsMap(address token) view returns(uint256 rewardRate, uint256 rewardIntegral, uint32 claimStartTime)
func (_Satusd *SatusdCallerSession) RewardParamsMap(token common.Address) (struct {
	RewardRate     *big.Int
	RewardIntegral *big.Int
	ClaimStartTime uint32
}, error) {
	return _Satusd.Contract.RewardParamsMap(&_Satusd.CallOpts, token)
}

// RewardVault is a free data retrieval call binding the contract method 0x3a2c6777.
//
// Solidity: function rewardVault() view returns(address)
func (_Satusd *SatusdCaller) RewardVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Satusd.contract.Call(opts, &out, "rewardVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardVault is a free data retrieval call binding the contract method 0x3a2c6777.
//
// Solidity: function rewardVault() view returns(address)
func (_Satusd *SatusdSession) RewardVault() (common.Address, error) {
	return _Satusd.Contract.RewardVault(&_Satusd.CallOpts)
}

// RewardVault is a free data retrieval call binding the contract method 0x3a2c6777.
//
// Solidity: function rewardVault() view returns(address)
func (_Satusd *SatusdCallerSession) RewardVault() (common.Address, error) {
	return _Satusd.Contract.RewardVault(&_Satusd.CallOpts)
}

// SimulateDeposit is a free data retrieval call binding the contract method 0x923c86e6.
//
// Solidity: function simulateDeposit(uint256 assets) view returns(uint256)
func (_Satusd *SatusdCaller) SimulateDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Satusd.contract.Call(opts, &out, "simulateDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SimulateDeposit is a free data retrieval call binding the contract method 0x923c86e6.
//
// Solidity: function simulateDeposit(uint256 assets) view returns(uint256)
func (_Satusd *SatusdSession) SimulateDeposit(assets *big.Int) (*big.Int, error) {
	return _Satusd.Contract.SimulateDeposit(&_Satusd.CallOpts, assets)
}

// SimulateDeposit is a free data retrieval call binding the contract method 0x923c86e6.
//
// Solidity: function simulateDeposit(uint256 assets) view returns(uint256)
func (_Satusd *SatusdCallerSession) SimulateDeposit(assets *big.Int) (*big.Int, error) {
	return _Satusd.Contract.SimulateDeposit(&_Satusd.CallOpts, assets)
}

// SimulateMint is a free data retrieval call binding the contract method 0xf9f18df7.
//
// Solidity: function simulateMint(uint256 shares) view returns(uint256)
func (_Satusd *SatusdCaller) SimulateMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Satusd.contract.Call(opts, &out, "simulateMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SimulateMint is a free data retrieval call binding the contract method 0xf9f18df7.
//
// Solidity: function simulateMint(uint256 shares) view returns(uint256)
func (_Satusd *SatusdSession) SimulateMint(shares *big.Int) (*big.Int, error) {
	return _Satusd.Contract.SimulateMint(&_Satusd.CallOpts, shares)
}

// SimulateMint is a free data retrieval call binding the contract method 0xf9f18df7.
//
// Solidity: function simulateMint(uint256 shares) view returns(uint256)
func (_Satusd *SatusdCallerSession) SimulateMint(shares *big.Int) (*big.Int, error) {
	return _Satusd.Contract.SimulateMint(&_Satusd.CallOpts, shares)
}

// SimulateRedeem is a free data retrieval call binding the contract method 0xec1ac75b.
//
// Solidity: function simulateRedeem(uint256 shares) view returns(uint256)
func (_Satusd *SatusdCaller) SimulateRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Satusd.contract.Call(opts, &out, "simulateRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SimulateRedeem is a free data retrieval call binding the contract method 0xec1ac75b.
//
// Solidity: function simulateRedeem(uint256 shares) view returns(uint256)
func (_Satusd *SatusdSession) SimulateRedeem(shares *big.Int) (*big.Int, error) {
	return _Satusd.Contract.SimulateRedeem(&_Satusd.CallOpts, shares)
}

// SimulateRedeem is a free data retrieval call binding the contract method 0xec1ac75b.
//
// Solidity: function simulateRedeem(uint256 shares) view returns(uint256)
func (_Satusd *SatusdCallerSession) SimulateRedeem(shares *big.Int) (*big.Int, error) {
	return _Satusd.Contract.SimulateRedeem(&_Satusd.CallOpts, shares)
}

// SimulateWithdraw is a free data retrieval call binding the contract method 0xbb31203d.
//
// Solidity: function simulateWithdraw(uint256 assets) view returns(uint256)
func (_Satusd *SatusdCaller) SimulateWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Satusd.contract.Call(opts, &out, "simulateWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SimulateWithdraw is a free data retrieval call binding the contract method 0xbb31203d.
//
// Solidity: function simulateWithdraw(uint256 assets) view returns(uint256)
func (_Satusd *SatusdSession) SimulateWithdraw(assets *big.Int) (*big.Int, error) {
	return _Satusd.Contract.SimulateWithdraw(&_Satusd.CallOpts, assets)
}

// SimulateWithdraw is a free data retrieval call binding the contract method 0xbb31203d.
//
// Solidity: function simulateWithdraw(uint256 assets) view returns(uint256)
func (_Satusd *SatusdCallerSession) SimulateWithdraw(assets *big.Int) (*big.Int, error) {
	return _Satusd.Contract.SimulateWithdraw(&_Satusd.CallOpts, assets)
}

// StoredPendingReward is a free data retrieval call binding the contract method 0x51210430.
//
// Solidity: function storedPendingReward(address token, address account) view returns(uint256)
func (_Satusd *SatusdCaller) StoredPendingReward(opts *bind.CallOpts, token common.Address, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Satusd.contract.Call(opts, &out, "storedPendingReward", token, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StoredPendingReward is a free data retrieval call binding the contract method 0x51210430.
//
// Solidity: function storedPendingReward(address token, address account) view returns(uint256)
func (_Satusd *SatusdSession) StoredPendingReward(token common.Address, account common.Address) (*big.Int, error) {
	return _Satusd.Contract.StoredPendingReward(&_Satusd.CallOpts, token, account)
}

// StoredPendingReward is a free data retrieval call binding the contract method 0x51210430.
//
// Solidity: function storedPendingReward(address token, address account) view returns(uint256)
func (_Satusd *SatusdCallerSession) StoredPendingReward(token common.Address, account common.Address) (*big.Int, error) {
	return _Satusd.Contract.StoredPendingReward(&_Satusd.CallOpts, token, account)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Satusd *SatusdCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Satusd.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Satusd *SatusdSession) Symbol() (string, error) {
	return _Satusd.Contract.Symbol(&_Satusd.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Satusd *SatusdCallerSession) Symbol() (string, error) {
	return _Satusd.Contract.Symbol(&_Satusd.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Satusd *SatusdCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Satusd.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Satusd *SatusdSession) TotalAssets() (*big.Int, error) {
	return _Satusd.Contract.TotalAssets(&_Satusd.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Satusd *SatusdCallerSession) TotalAssets() (*big.Int, error) {
	return _Satusd.Contract.TotalAssets(&_Satusd.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Satusd *SatusdCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Satusd.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Satusd *SatusdSession) TotalSupply() (*big.Int, error) {
	return _Satusd.Contract.TotalSupply(&_Satusd.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Satusd *SatusdCallerSession) TotalSupply() (*big.Int, error) {
	return _Satusd.Contract.TotalSupply(&_Satusd.CallOpts)
}

// AddAssetToEmissionList is a paid mutator transaction binding the contract method 0xf23331c6.
//
// Solidity: function addAssetToEmissionList(address asset_) returns()
func (_Satusd *SatusdTransactor) AddAssetToEmissionList(opts *bind.TransactOpts, asset_ common.Address) (*types.Transaction, error) {
	return _Satusd.contract.Transact(opts, "addAssetToEmissionList", asset_)
}

// AddAssetToEmissionList is a paid mutator transaction binding the contract method 0xf23331c6.
//
// Solidity: function addAssetToEmissionList(address asset_) returns()
func (_Satusd *SatusdSession) AddAssetToEmissionList(asset_ common.Address) (*types.Transaction, error) {
	return _Satusd.Contract.AddAssetToEmissionList(&_Satusd.TransactOpts, asset_)
}

// AddAssetToEmissionList is a paid mutator transaction binding the contract method 0xf23331c6.
//
// Solidity: function addAssetToEmissionList(address asset_) returns()
func (_Satusd *SatusdTransactorSession) AddAssetToEmissionList(asset_ common.Address) (*types.Transaction, error) {
	return _Satusd.Contract.AddAssetToEmissionList(&_Satusd.TransactOpts, asset_)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Satusd *SatusdTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Satusd.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Satusd *SatusdSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Satusd.Contract.Approve(&_Satusd.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Satusd *SatusdTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Satusd.Contract.Approve(&_Satusd.TransactOpts, spender, value)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xb88a802f.
//
// Solidity: function claimReward() returns()
func (_Satusd *SatusdTransactor) ClaimReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Satusd.contract.Transact(opts, "claimReward")
}

// ClaimReward is a paid mutator transaction binding the contract method 0xb88a802f.
//
// Solidity: function claimReward() returns()
func (_Satusd *SatusdSession) ClaimReward() (*types.Transaction, error) {
	return _Satusd.Contract.ClaimReward(&_Satusd.TransactOpts)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xb88a802f.
//
// Solidity: function claimReward() returns()
func (_Satusd *SatusdTransactorSession) ClaimReward() (*types.Transaction, error) {
	return _Satusd.Contract.ClaimReward(&_Satusd.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_Satusd *SatusdTransactor) Deposit(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Satusd.contract.Transact(opts, "deposit", assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_Satusd *SatusdSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Satusd.Contract.Deposit(&_Satusd.TransactOpts, assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_Satusd *SatusdTransactorSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Satusd.Contract.Deposit(&_Satusd.TransactOpts, assets, receiver)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _underlying, address _rewardVault) returns()
func (_Satusd *SatusdTransactor) Initialize(opts *bind.TransactOpts, _underlying common.Address, _rewardVault common.Address) (*types.Transaction, error) {
	return _Satusd.contract.Transact(opts, "initialize", _underlying, _rewardVault)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _underlying, address _rewardVault) returns()
func (_Satusd *SatusdSession) Initialize(_underlying common.Address, _rewardVault common.Address) (*types.Transaction, error) {
	return _Satusd.Contract.Initialize(&_Satusd.TransactOpts, _underlying, _rewardVault)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _underlying, address _rewardVault) returns()
func (_Satusd *SatusdTransactorSession) Initialize(_underlying common.Address, _rewardVault common.Address) (*types.Transaction, error) {
	return _Satusd.Contract.Initialize(&_Satusd.TransactOpts, _underlying, _rewardVault)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_Satusd *SatusdTransactor) Mint(opts *bind.TransactOpts, shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Satusd.contract.Transact(opts, "mint", shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_Satusd *SatusdSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Satusd.Contract.Mint(&_Satusd.TransactOpts, shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_Satusd *SatusdTransactorSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Satusd.Contract.Mint(&_Satusd.TransactOpts, shares, receiver)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Satusd *SatusdTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Satusd.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Satusd *SatusdSession) Pause() (*types.Transaction, error) {
	return _Satusd.Contract.Pause(&_Satusd.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Satusd *SatusdTransactorSession) Pause() (*types.Transaction, error) {
	return _Satusd.Contract.Pause(&_Satusd.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_Satusd *SatusdTransactor) Redeem(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Satusd.contract.Transact(opts, "redeem", shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_Satusd *SatusdSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Satusd.Contract.Redeem(&_Satusd.TransactOpts, shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_Satusd *SatusdTransactorSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Satusd.Contract.Redeem(&_Satusd.TransactOpts, shares, receiver, owner)
}

// RemoveAssetFromEmissionList is a paid mutator transaction binding the contract method 0xab571c60.
//
// Solidity: function removeAssetFromEmissionList(address asset_) returns()
func (_Satusd *SatusdTransactor) RemoveAssetFromEmissionList(opts *bind.TransactOpts, asset_ common.Address) (*types.Transaction, error) {
	return _Satusd.contract.Transact(opts, "removeAssetFromEmissionList", asset_)
}

// RemoveAssetFromEmissionList is a paid mutator transaction binding the contract method 0xab571c60.
//
// Solidity: function removeAssetFromEmissionList(address asset_) returns()
func (_Satusd *SatusdSession) RemoveAssetFromEmissionList(asset_ common.Address) (*types.Transaction, error) {
	return _Satusd.Contract.RemoveAssetFromEmissionList(&_Satusd.TransactOpts, asset_)
}

// RemoveAssetFromEmissionList is a paid mutator transaction binding the contract method 0xab571c60.
//
// Solidity: function removeAssetFromEmissionList(address asset_) returns()
func (_Satusd *SatusdTransactorSession) RemoveAssetFromEmissionList(asset_ common.Address) (*types.Transaction, error) {
	return _Satusd.Contract.RemoveAssetFromEmissionList(&_Satusd.TransactOpts, asset_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Satusd *SatusdTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Satusd.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Satusd *SatusdSession) RenounceOwnership() (*types.Transaction, error) {
	return _Satusd.Contract.RenounceOwnership(&_Satusd.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Satusd *SatusdTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Satusd.Contract.RenounceOwnership(&_Satusd.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Satusd *SatusdTransactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Satusd.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Satusd *SatusdSession) Resume() (*types.Transaction, error) {
	return _Satusd.Contract.Resume(&_Satusd.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Satusd *SatusdTransactorSession) Resume() (*types.Transaction, error) {
	return _Satusd.Contract.Resume(&_Satusd.TransactOpts)
}

// SetRewardParams is a paid mutator transaction binding the contract method 0xbd3e7d97.
//
// Solidity: function setRewardParams(address token, uint256 rewardRate, uint32 claimStartTime) returns()
func (_Satusd *SatusdTransactor) SetRewardParams(opts *bind.TransactOpts, token common.Address, rewardRate *big.Int, claimStartTime uint32) (*types.Transaction, error) {
	return _Satusd.contract.Transact(opts, "setRewardParams", token, rewardRate, claimStartTime)
}

// SetRewardParams is a paid mutator transaction binding the contract method 0xbd3e7d97.
//
// Solidity: function setRewardParams(address token, uint256 rewardRate, uint32 claimStartTime) returns()
func (_Satusd *SatusdSession) SetRewardParams(token common.Address, rewardRate *big.Int, claimStartTime uint32) (*types.Transaction, error) {
	return _Satusd.Contract.SetRewardParams(&_Satusd.TransactOpts, token, rewardRate, claimStartTime)
}

// SetRewardParams is a paid mutator transaction binding the contract method 0xbd3e7d97.
//
// Solidity: function setRewardParams(address token, uint256 rewardRate, uint32 claimStartTime) returns()
func (_Satusd *SatusdTransactorSession) SetRewardParams(token common.Address, rewardRate *big.Int, claimStartTime uint32) (*types.Transaction, error) {
	return _Satusd.Contract.SetRewardParams(&_Satusd.TransactOpts, token, rewardRate, claimStartTime)
}

// SetRewardVault is a paid mutator transaction binding the contract method 0x8125dd10.
//
// Solidity: function setRewardVault(address _rewardVault) returns()
func (_Satusd *SatusdTransactor) SetRewardVault(opts *bind.TransactOpts, _rewardVault common.Address) (*types.Transaction, error) {
	return _Satusd.contract.Transact(opts, "setRewardVault", _rewardVault)
}

// SetRewardVault is a paid mutator transaction binding the contract method 0x8125dd10.
//
// Solidity: function setRewardVault(address _rewardVault) returns()
func (_Satusd *SatusdSession) SetRewardVault(_rewardVault common.Address) (*types.Transaction, error) {
	return _Satusd.Contract.SetRewardVault(&_Satusd.TransactOpts, _rewardVault)
}

// SetRewardVault is a paid mutator transaction binding the contract method 0x8125dd10.
//
// Solidity: function setRewardVault(address _rewardVault) returns()
func (_Satusd *SatusdTransactorSession) SetRewardVault(_rewardVault common.Address) (*types.Transaction, error) {
	return _Satusd.Contract.SetRewardVault(&_Satusd.TransactOpts, _rewardVault)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Satusd *SatusdTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Satusd.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Satusd *SatusdSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Satusd.Contract.Transfer(&_Satusd.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Satusd *SatusdTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Satusd.Contract.Transfer(&_Satusd.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Satusd *SatusdTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Satusd.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Satusd *SatusdSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Satusd.Contract.TransferFrom(&_Satusd.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Satusd *SatusdTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Satusd.Contract.TransferFrom(&_Satusd.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Satusd *SatusdTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Satusd.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Satusd *SatusdSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Satusd.Contract.TransferOwnership(&_Satusd.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Satusd *SatusdTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Satusd.Contract.TransferOwnership(&_Satusd.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Satusd *SatusdTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Satusd.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Satusd *SatusdSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Satusd.Contract.UpgradeToAndCall(&_Satusd.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Satusd *SatusdTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Satusd.Contract.UpgradeToAndCall(&_Satusd.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_Satusd *SatusdTransactor) Withdraw(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Satusd.contract.Transact(opts, "withdraw", assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_Satusd *SatusdSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Satusd.Contract.Withdraw(&_Satusd.TransactOpts, assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_Satusd *SatusdTransactorSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Satusd.Contract.Withdraw(&_Satusd.TransactOpts, assets, receiver, owner)
}

// SatusdApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Satusd contract.
type SatusdApprovalIterator struct {
	Event *SatusdApproval // Event containing the contract specifics and raw log

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
func (it *SatusdApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SatusdApproval)
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
		it.Event = new(SatusdApproval)
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
func (it *SatusdApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SatusdApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SatusdApproval represents a Approval event raised by the Satusd contract.
type SatusdApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Satusd *SatusdFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SatusdApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Satusd.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SatusdApprovalIterator{contract: _Satusd.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Satusd *SatusdFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SatusdApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Satusd.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SatusdApproval)
				if err := _Satusd.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Satusd *SatusdFilterer) ParseApproval(log types.Log) (*SatusdApproval, error) {
	event := new(SatusdApproval)
	if err := _Satusd.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SatusdAssetAddedToEmissionListIterator is returned from FilterAssetAddedToEmissionList and is used to iterate over the raw logs and unpacked data for AssetAddedToEmissionList events raised by the Satusd contract.
type SatusdAssetAddedToEmissionListIterator struct {
	Event *SatusdAssetAddedToEmissionList // Event containing the contract specifics and raw log

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
func (it *SatusdAssetAddedToEmissionListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SatusdAssetAddedToEmissionList)
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
		it.Event = new(SatusdAssetAddedToEmissionList)
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
func (it *SatusdAssetAddedToEmissionListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SatusdAssetAddedToEmissionListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SatusdAssetAddedToEmissionList represents a AssetAddedToEmissionList event raised by the Satusd contract.
type SatusdAssetAddedToEmissionList struct {
	Asset common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAssetAddedToEmissionList is a free log retrieval operation binding the contract event 0xfe55b62d54eea531cb2ab0224ef4aa29eb159c1303f08e0b4e5811cabe810a30.
//
// Solidity: event AssetAddedToEmissionList(address indexed asset)
func (_Satusd *SatusdFilterer) FilterAssetAddedToEmissionList(opts *bind.FilterOpts, asset []common.Address) (*SatusdAssetAddedToEmissionListIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _Satusd.contract.FilterLogs(opts, "AssetAddedToEmissionList", assetRule)
	if err != nil {
		return nil, err
	}
	return &SatusdAssetAddedToEmissionListIterator{contract: _Satusd.contract, event: "AssetAddedToEmissionList", logs: logs, sub: sub}, nil
}

// WatchAssetAddedToEmissionList is a free log subscription operation binding the contract event 0xfe55b62d54eea531cb2ab0224ef4aa29eb159c1303f08e0b4e5811cabe810a30.
//
// Solidity: event AssetAddedToEmissionList(address indexed asset)
func (_Satusd *SatusdFilterer) WatchAssetAddedToEmissionList(opts *bind.WatchOpts, sink chan<- *SatusdAssetAddedToEmissionList, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _Satusd.contract.WatchLogs(opts, "AssetAddedToEmissionList", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SatusdAssetAddedToEmissionList)
				if err := _Satusd.contract.UnpackLog(event, "AssetAddedToEmissionList", log); err != nil {
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

// ParseAssetAddedToEmissionList is a log parse operation binding the contract event 0xfe55b62d54eea531cb2ab0224ef4aa29eb159c1303f08e0b4e5811cabe810a30.
//
// Solidity: event AssetAddedToEmissionList(address indexed asset)
func (_Satusd *SatusdFilterer) ParseAssetAddedToEmissionList(log types.Log) (*SatusdAssetAddedToEmissionList, error) {
	event := new(SatusdAssetAddedToEmissionList)
	if err := _Satusd.contract.UnpackLog(event, "AssetAddedToEmissionList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SatusdAssetRemovedFromEmissionListIterator is returned from FilterAssetRemovedFromEmissionList and is used to iterate over the raw logs and unpacked data for AssetRemovedFromEmissionList events raised by the Satusd contract.
type SatusdAssetRemovedFromEmissionListIterator struct {
	Event *SatusdAssetRemovedFromEmissionList // Event containing the contract specifics and raw log

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
func (it *SatusdAssetRemovedFromEmissionListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SatusdAssetRemovedFromEmissionList)
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
		it.Event = new(SatusdAssetRemovedFromEmissionList)
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
func (it *SatusdAssetRemovedFromEmissionListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SatusdAssetRemovedFromEmissionListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SatusdAssetRemovedFromEmissionList represents a AssetRemovedFromEmissionList event raised by the Satusd contract.
type SatusdAssetRemovedFromEmissionList struct {
	Asset common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAssetRemovedFromEmissionList is a free log retrieval operation binding the contract event 0xf317de7b5d0ba757a221e919e6249a2e16d542c2f90b59d64d50484be4be3bd7.
//
// Solidity: event AssetRemovedFromEmissionList(address indexed asset)
func (_Satusd *SatusdFilterer) FilterAssetRemovedFromEmissionList(opts *bind.FilterOpts, asset []common.Address) (*SatusdAssetRemovedFromEmissionListIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _Satusd.contract.FilterLogs(opts, "AssetRemovedFromEmissionList", assetRule)
	if err != nil {
		return nil, err
	}
	return &SatusdAssetRemovedFromEmissionListIterator{contract: _Satusd.contract, event: "AssetRemovedFromEmissionList", logs: logs, sub: sub}, nil
}

// WatchAssetRemovedFromEmissionList is a free log subscription operation binding the contract event 0xf317de7b5d0ba757a221e919e6249a2e16d542c2f90b59d64d50484be4be3bd7.
//
// Solidity: event AssetRemovedFromEmissionList(address indexed asset)
func (_Satusd *SatusdFilterer) WatchAssetRemovedFromEmissionList(opts *bind.WatchOpts, sink chan<- *SatusdAssetRemovedFromEmissionList, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _Satusd.contract.WatchLogs(opts, "AssetRemovedFromEmissionList", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SatusdAssetRemovedFromEmissionList)
				if err := _Satusd.contract.UnpackLog(event, "AssetRemovedFromEmissionList", log); err != nil {
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

// ParseAssetRemovedFromEmissionList is a log parse operation binding the contract event 0xf317de7b5d0ba757a221e919e6249a2e16d542c2f90b59d64d50484be4be3bd7.
//
// Solidity: event AssetRemovedFromEmissionList(address indexed asset)
func (_Satusd *SatusdFilterer) ParseAssetRemovedFromEmissionList(log types.Log) (*SatusdAssetRemovedFromEmissionList, error) {
	event := new(SatusdAssetRemovedFromEmissionList)
	if err := _Satusd.contract.UnpackLog(event, "AssetRemovedFromEmissionList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SatusdDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Satusd contract.
type SatusdDepositIterator struct {
	Event *SatusdDeposit // Event containing the contract specifics and raw log

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
func (it *SatusdDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SatusdDeposit)
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
		it.Event = new(SatusdDeposit)
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
func (it *SatusdDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SatusdDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SatusdDeposit represents a Deposit event raised by the Satusd contract.
type SatusdDeposit struct {
	Sender common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_Satusd *SatusdFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*SatusdDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Satusd.contract.FilterLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &SatusdDepositIterator{contract: _Satusd.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_Satusd *SatusdFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *SatusdDeposit, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Satusd.contract.WatchLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SatusdDeposit)
				if err := _Satusd.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_Satusd *SatusdFilterer) ParseDeposit(log types.Log) (*SatusdDeposit, error) {
	event := new(SatusdDeposit)
	if err := _Satusd.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SatusdInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Satusd contract.
type SatusdInitializedIterator struct {
	Event *SatusdInitialized // Event containing the contract specifics and raw log

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
func (it *SatusdInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SatusdInitialized)
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
		it.Event = new(SatusdInitialized)
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
func (it *SatusdInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SatusdInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SatusdInitialized represents a Initialized event raised by the Satusd contract.
type SatusdInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Satusd *SatusdFilterer) FilterInitialized(opts *bind.FilterOpts) (*SatusdInitializedIterator, error) {

	logs, sub, err := _Satusd.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SatusdInitializedIterator{contract: _Satusd.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Satusd *SatusdFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SatusdInitialized) (event.Subscription, error) {

	logs, sub, err := _Satusd.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SatusdInitialized)
				if err := _Satusd.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Satusd *SatusdFilterer) ParseInitialized(log types.Log) (*SatusdInitialized, error) {
	event := new(SatusdInitialized)
	if err := _Satusd.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SatusdOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Satusd contract.
type SatusdOwnershipTransferredIterator struct {
	Event *SatusdOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SatusdOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SatusdOwnershipTransferred)
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
		it.Event = new(SatusdOwnershipTransferred)
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
func (it *SatusdOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SatusdOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SatusdOwnershipTransferred represents a OwnershipTransferred event raised by the Satusd contract.
type SatusdOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Satusd *SatusdFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SatusdOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Satusd.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SatusdOwnershipTransferredIterator{contract: _Satusd.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Satusd *SatusdFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SatusdOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Satusd.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SatusdOwnershipTransferred)
				if err := _Satusd.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Satusd *SatusdFilterer) ParseOwnershipTransferred(log types.Log) (*SatusdOwnershipTransferred, error) {
	event := new(SatusdOwnershipTransferred)
	if err := _Satusd.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SatusdPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Satusd contract.
type SatusdPausedIterator struct {
	Event *SatusdPaused // Event containing the contract specifics and raw log

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
func (it *SatusdPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SatusdPaused)
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
		it.Event = new(SatusdPaused)
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
func (it *SatusdPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SatusdPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SatusdPaused represents a Paused event raised by the Satusd contract.
type SatusdPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Satusd *SatusdFilterer) FilterPaused(opts *bind.FilterOpts) (*SatusdPausedIterator, error) {

	logs, sub, err := _Satusd.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &SatusdPausedIterator{contract: _Satusd.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Satusd *SatusdFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *SatusdPaused) (event.Subscription, error) {

	logs, sub, err := _Satusd.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SatusdPaused)
				if err := _Satusd.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Satusd *SatusdFilterer) ParsePaused(log types.Log) (*SatusdPaused, error) {
	event := new(SatusdPaused)
	if err := _Satusd.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SatusdRewardClaimedIterator is returned from FilterRewardClaimed and is used to iterate over the raw logs and unpacked data for RewardClaimed events raised by the Satusd contract.
type SatusdRewardClaimedIterator struct {
	Event *SatusdRewardClaimed // Event containing the contract specifics and raw log

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
func (it *SatusdRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SatusdRewardClaimed)
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
		it.Event = new(SatusdRewardClaimed)
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
func (it *SatusdRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SatusdRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SatusdRewardClaimed represents a RewardClaimed event raised by the Satusd contract.
type SatusdRewardClaimed struct {
	Token    common.Address
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRewardClaimed is a free log retrieval operation binding the contract event 0x0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b7.
//
// Solidity: event RewardClaimed(address indexed token, address indexed receiver, uint256 amount)
func (_Satusd *SatusdFilterer) FilterRewardClaimed(opts *bind.FilterOpts, token []common.Address, receiver []common.Address) (*SatusdRewardClaimedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Satusd.contract.FilterLogs(opts, "RewardClaimed", tokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &SatusdRewardClaimedIterator{contract: _Satusd.contract, event: "RewardClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardClaimed is a free log subscription operation binding the contract event 0x0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b7.
//
// Solidity: event RewardClaimed(address indexed token, address indexed receiver, uint256 amount)
func (_Satusd *SatusdFilterer) WatchRewardClaimed(opts *bind.WatchOpts, sink chan<- *SatusdRewardClaimed, token []common.Address, receiver []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Satusd.contract.WatchLogs(opts, "RewardClaimed", tokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SatusdRewardClaimed)
				if err := _Satusd.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
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

// ParseRewardClaimed is a log parse operation binding the contract event 0x0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b7.
//
// Solidity: event RewardClaimed(address indexed token, address indexed receiver, uint256 amount)
func (_Satusd *SatusdFilterer) ParseRewardClaimed(log types.Log) (*SatusdRewardClaimed, error) {
	event := new(SatusdRewardClaimed)
	if err := _Satusd.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SatusdRewardIntegralForAccountUpdatedIterator is returned from FilterRewardIntegralForAccountUpdated and is used to iterate over the raw logs and unpacked data for RewardIntegralForAccountUpdated events raised by the Satusd contract.
type SatusdRewardIntegralForAccountUpdatedIterator struct {
	Event *SatusdRewardIntegralForAccountUpdated // Event containing the contract specifics and raw log

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
func (it *SatusdRewardIntegralForAccountUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SatusdRewardIntegralForAccountUpdated)
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
		it.Event = new(SatusdRewardIntegralForAccountUpdated)
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
func (it *SatusdRewardIntegralForAccountUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SatusdRewardIntegralForAccountUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SatusdRewardIntegralForAccountUpdated represents a RewardIntegralForAccountUpdated event raised by the Satusd contract.
type SatusdRewardIntegralForAccountUpdated struct {
	Account        common.Address
	Token          common.Address
	PendingReward  *big.Int
	RewardIntegral *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRewardIntegralForAccountUpdated is a free log retrieval operation binding the contract event 0xdb1b81b8ffba907fa76754265baca950292a08fdd6584527e35204dea596dcc1.
//
// Solidity: event RewardIntegralForAccountUpdated(address indexed account, address indexed token, uint256 pendingReward, uint256 rewardIntegral)
func (_Satusd *SatusdFilterer) FilterRewardIntegralForAccountUpdated(opts *bind.FilterOpts, account []common.Address, token []common.Address) (*SatusdRewardIntegralForAccountUpdatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Satusd.contract.FilterLogs(opts, "RewardIntegralForAccountUpdated", accountRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &SatusdRewardIntegralForAccountUpdatedIterator{contract: _Satusd.contract, event: "RewardIntegralForAccountUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardIntegralForAccountUpdated is a free log subscription operation binding the contract event 0xdb1b81b8ffba907fa76754265baca950292a08fdd6584527e35204dea596dcc1.
//
// Solidity: event RewardIntegralForAccountUpdated(address indexed account, address indexed token, uint256 pendingReward, uint256 rewardIntegral)
func (_Satusd *SatusdFilterer) WatchRewardIntegralForAccountUpdated(opts *bind.WatchOpts, sink chan<- *SatusdRewardIntegralForAccountUpdated, account []common.Address, token []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Satusd.contract.WatchLogs(opts, "RewardIntegralForAccountUpdated", accountRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SatusdRewardIntegralForAccountUpdated)
				if err := _Satusd.contract.UnpackLog(event, "RewardIntegralForAccountUpdated", log); err != nil {
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

// ParseRewardIntegralForAccountUpdated is a log parse operation binding the contract event 0xdb1b81b8ffba907fa76754265baca950292a08fdd6584527e35204dea596dcc1.
//
// Solidity: event RewardIntegralForAccountUpdated(address indexed account, address indexed token, uint256 pendingReward, uint256 rewardIntegral)
func (_Satusd *SatusdFilterer) ParseRewardIntegralForAccountUpdated(log types.Log) (*SatusdRewardIntegralForAccountUpdated, error) {
	event := new(SatusdRewardIntegralForAccountUpdated)
	if err := _Satusd.contract.UnpackLog(event, "RewardIntegralForAccountUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SatusdRewardIntegralUpdatedIterator is returned from FilterRewardIntegralUpdated and is used to iterate over the raw logs and unpacked data for RewardIntegralUpdated events raised by the Satusd contract.
type SatusdRewardIntegralUpdatedIterator struct {
	Event *SatusdRewardIntegralUpdated // Event containing the contract specifics and raw log

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
func (it *SatusdRewardIntegralUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SatusdRewardIntegralUpdated)
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
		it.Event = new(SatusdRewardIntegralUpdated)
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
func (it *SatusdRewardIntegralUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SatusdRewardIntegralUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SatusdRewardIntegralUpdated represents a RewardIntegralUpdated event raised by the Satusd contract.
type SatusdRewardIntegralUpdated struct {
	Token          common.Address
	RewardIntegral *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRewardIntegralUpdated is a free log retrieval operation binding the contract event 0x61d927f91635642c4c52b78d3f339c15ff959118da79ceb10122a91902182b04.
//
// Solidity: event RewardIntegralUpdated(address indexed token, uint256 rewardIntegral)
func (_Satusd *SatusdFilterer) FilterRewardIntegralUpdated(opts *bind.FilterOpts, token []common.Address) (*SatusdRewardIntegralUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Satusd.contract.FilterLogs(opts, "RewardIntegralUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &SatusdRewardIntegralUpdatedIterator{contract: _Satusd.contract, event: "RewardIntegralUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardIntegralUpdated is a free log subscription operation binding the contract event 0x61d927f91635642c4c52b78d3f339c15ff959118da79ceb10122a91902182b04.
//
// Solidity: event RewardIntegralUpdated(address indexed token, uint256 rewardIntegral)
func (_Satusd *SatusdFilterer) WatchRewardIntegralUpdated(opts *bind.WatchOpts, sink chan<- *SatusdRewardIntegralUpdated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Satusd.contract.WatchLogs(opts, "RewardIntegralUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SatusdRewardIntegralUpdated)
				if err := _Satusd.contract.UnpackLog(event, "RewardIntegralUpdated", log); err != nil {
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

// ParseRewardIntegralUpdated is a log parse operation binding the contract event 0x61d927f91635642c4c52b78d3f339c15ff959118da79ceb10122a91902182b04.
//
// Solidity: event RewardIntegralUpdated(address indexed token, uint256 rewardIntegral)
func (_Satusd *SatusdFilterer) ParseRewardIntegralUpdated(log types.Log) (*SatusdRewardIntegralUpdated, error) {
	event := new(SatusdRewardIntegralUpdated)
	if err := _Satusd.contract.UnpackLog(event, "RewardIntegralUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SatusdRewardParamsSetIterator is returned from FilterRewardParamsSet and is used to iterate over the raw logs and unpacked data for RewardParamsSet events raised by the Satusd contract.
type SatusdRewardParamsSetIterator struct {
	Event *SatusdRewardParamsSet // Event containing the contract specifics and raw log

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
func (it *SatusdRewardParamsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SatusdRewardParamsSet)
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
		it.Event = new(SatusdRewardParamsSet)
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
func (it *SatusdRewardParamsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SatusdRewardParamsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SatusdRewardParamsSet represents a RewardParamsSet event raised by the Satusd contract.
type SatusdRewardParamsSet struct {
	Token          common.Address
	RewardRate     *big.Int
	RewardIntegral *big.Int
	ClaimStartTime uint32
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRewardParamsSet is a free log retrieval operation binding the contract event 0xa82b7ef40535dbf27232e95dedda07f0ccf028e6f5432c9f79b843d5ed11d0f5.
//
// Solidity: event RewardParamsSet(address indexed token, uint256 rewardRate, uint256 rewardIntegral, uint32 claimStartTime)
func (_Satusd *SatusdFilterer) FilterRewardParamsSet(opts *bind.FilterOpts, token []common.Address) (*SatusdRewardParamsSetIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Satusd.contract.FilterLogs(opts, "RewardParamsSet", tokenRule)
	if err != nil {
		return nil, err
	}
	return &SatusdRewardParamsSetIterator{contract: _Satusd.contract, event: "RewardParamsSet", logs: logs, sub: sub}, nil
}

// WatchRewardParamsSet is a free log subscription operation binding the contract event 0xa82b7ef40535dbf27232e95dedda07f0ccf028e6f5432c9f79b843d5ed11d0f5.
//
// Solidity: event RewardParamsSet(address indexed token, uint256 rewardRate, uint256 rewardIntegral, uint32 claimStartTime)
func (_Satusd *SatusdFilterer) WatchRewardParamsSet(opts *bind.WatchOpts, sink chan<- *SatusdRewardParamsSet, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Satusd.contract.WatchLogs(opts, "RewardParamsSet", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SatusdRewardParamsSet)
				if err := _Satusd.contract.UnpackLog(event, "RewardParamsSet", log); err != nil {
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

// ParseRewardParamsSet is a log parse operation binding the contract event 0xa82b7ef40535dbf27232e95dedda07f0ccf028e6f5432c9f79b843d5ed11d0f5.
//
// Solidity: event RewardParamsSet(address indexed token, uint256 rewardRate, uint256 rewardIntegral, uint32 claimStartTime)
func (_Satusd *SatusdFilterer) ParseRewardParamsSet(log types.Log) (*SatusdRewardParamsSet, error) {
	event := new(SatusdRewardParamsSet)
	if err := _Satusd.contract.UnpackLog(event, "RewardParamsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SatusdRewardVaultSettingIterator is returned from FilterRewardVaultSetting and is used to iterate over the raw logs and unpacked data for RewardVaultSetting events raised by the Satusd contract.
type SatusdRewardVaultSettingIterator struct {
	Event *SatusdRewardVaultSetting // Event containing the contract specifics and raw log

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
func (it *SatusdRewardVaultSettingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SatusdRewardVaultSetting)
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
		it.Event = new(SatusdRewardVaultSetting)
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
func (it *SatusdRewardVaultSettingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SatusdRewardVaultSettingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SatusdRewardVaultSetting represents a RewardVaultSetting event raised by the Satusd contract.
type SatusdRewardVaultSetting struct {
	RewardVault common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardVaultSetting is a free log retrieval operation binding the contract event 0x7ccb2cd577263a3595bff3e18ab7719c1f427c7fd48cdd043eff96fd3c583753.
//
// Solidity: event RewardVaultSetting(address indexed rewardVault)
func (_Satusd *SatusdFilterer) FilterRewardVaultSetting(opts *bind.FilterOpts, rewardVault []common.Address) (*SatusdRewardVaultSettingIterator, error) {

	var rewardVaultRule []interface{}
	for _, rewardVaultItem := range rewardVault {
		rewardVaultRule = append(rewardVaultRule, rewardVaultItem)
	}

	logs, sub, err := _Satusd.contract.FilterLogs(opts, "RewardVaultSetting", rewardVaultRule)
	if err != nil {
		return nil, err
	}
	return &SatusdRewardVaultSettingIterator{contract: _Satusd.contract, event: "RewardVaultSetting", logs: logs, sub: sub}, nil
}

// WatchRewardVaultSetting is a free log subscription operation binding the contract event 0x7ccb2cd577263a3595bff3e18ab7719c1f427c7fd48cdd043eff96fd3c583753.
//
// Solidity: event RewardVaultSetting(address indexed rewardVault)
func (_Satusd *SatusdFilterer) WatchRewardVaultSetting(opts *bind.WatchOpts, sink chan<- *SatusdRewardVaultSetting, rewardVault []common.Address) (event.Subscription, error) {

	var rewardVaultRule []interface{}
	for _, rewardVaultItem := range rewardVault {
		rewardVaultRule = append(rewardVaultRule, rewardVaultItem)
	}

	logs, sub, err := _Satusd.contract.WatchLogs(opts, "RewardVaultSetting", rewardVaultRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SatusdRewardVaultSetting)
				if err := _Satusd.contract.UnpackLog(event, "RewardVaultSetting", log); err != nil {
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

// ParseRewardVaultSetting is a log parse operation binding the contract event 0x7ccb2cd577263a3595bff3e18ab7719c1f427c7fd48cdd043eff96fd3c583753.
//
// Solidity: event RewardVaultSetting(address indexed rewardVault)
func (_Satusd *SatusdFilterer) ParseRewardVaultSetting(log types.Log) (*SatusdRewardVaultSetting, error) {
	event := new(SatusdRewardVaultSetting)
	if err := _Satusd.contract.UnpackLog(event, "RewardVaultSetting", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SatusdTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Satusd contract.
type SatusdTransferIterator struct {
	Event *SatusdTransfer // Event containing the contract specifics and raw log

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
func (it *SatusdTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SatusdTransfer)
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
		it.Event = new(SatusdTransfer)
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
func (it *SatusdTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SatusdTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SatusdTransfer represents a Transfer event raised by the Satusd contract.
type SatusdTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Satusd *SatusdFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SatusdTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Satusd.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SatusdTransferIterator{contract: _Satusd.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Satusd *SatusdFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SatusdTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Satusd.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SatusdTransfer)
				if err := _Satusd.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Satusd *SatusdFilterer) ParseTransfer(log types.Log) (*SatusdTransfer, error) {
	event := new(SatusdTransfer)
	if err := _Satusd.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SatusdUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Satusd contract.
type SatusdUnpausedIterator struct {
	Event *SatusdUnpaused // Event containing the contract specifics and raw log

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
func (it *SatusdUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SatusdUnpaused)
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
		it.Event = new(SatusdUnpaused)
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
func (it *SatusdUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SatusdUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SatusdUnpaused represents a Unpaused event raised by the Satusd contract.
type SatusdUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Satusd *SatusdFilterer) FilterUnpaused(opts *bind.FilterOpts) (*SatusdUnpausedIterator, error) {

	logs, sub, err := _Satusd.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &SatusdUnpausedIterator{contract: _Satusd.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Satusd *SatusdFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *SatusdUnpaused) (event.Subscription, error) {

	logs, sub, err := _Satusd.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SatusdUnpaused)
				if err := _Satusd.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Satusd *SatusdFilterer) ParseUnpaused(log types.Log) (*SatusdUnpaused, error) {
	event := new(SatusdUnpaused)
	if err := _Satusd.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SatusdUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Satusd contract.
type SatusdUpgradedIterator struct {
	Event *SatusdUpgraded // Event containing the contract specifics and raw log

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
func (it *SatusdUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SatusdUpgraded)
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
		it.Event = new(SatusdUpgraded)
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
func (it *SatusdUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SatusdUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SatusdUpgraded represents a Upgraded event raised by the Satusd contract.
type SatusdUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Satusd *SatusdFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*SatusdUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Satusd.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &SatusdUpgradedIterator{contract: _Satusd.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Satusd *SatusdFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *SatusdUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Satusd.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SatusdUpgraded)
				if err := _Satusd.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Satusd *SatusdFilterer) ParseUpgraded(log types.Log) (*SatusdUpgraded, error) {
	event := new(SatusdUpgraded)
	if err := _Satusd.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SatusdWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Satusd contract.
type SatusdWithdrawIterator struct {
	Event *SatusdWithdraw // Event containing the contract specifics and raw log

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
func (it *SatusdWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SatusdWithdraw)
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
		it.Event = new(SatusdWithdraw)
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
func (it *SatusdWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SatusdWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SatusdWithdraw represents a Withdraw event raised by the Satusd contract.
type SatusdWithdraw struct {
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
func (_Satusd *SatusdFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, owner []common.Address) (*SatusdWithdrawIterator, error) {

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

	logs, sub, err := _Satusd.contract.FilterLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &SatusdWithdrawIterator{contract: _Satusd.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_Satusd *SatusdFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *SatusdWithdraw, sender []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Satusd.contract.WatchLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SatusdWithdraw)
				if err := _Satusd.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_Satusd *SatusdFilterer) ParseWithdraw(log types.Log) (*SatusdWithdraw, error) {
	event := new(SatusdWithdraw)
	if err := _Satusd.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
