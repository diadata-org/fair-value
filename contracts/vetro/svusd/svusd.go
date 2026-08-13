// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package svusd

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

// IStakingVaultCooldownRequest is an auto generated low-level Go binding around an user-defined struct.
type IStakingVaultCooldownRequest struct {
	Owner       common.Address
	Assets      *big.Int
	ClaimableAt *big.Int
}

// SvusdMetaData contains all meta data concerning the Svusd contract.
var SvusdMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CooldownEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CooldownNotEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimableAt\",\"type\":\"uint256\"}],\"name\":\"CooldownNotMatured\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxDeposit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxRedeem\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxWithdraw\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDuration\",\"type\":\"uint256\"}],\"name\":\"InvalidCooldownDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"InvalidRequestId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotRequestOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDuration\",\"type\":\"uint256\"}],\"name\":\"CooldownDurationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"previousStatus\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"newStatus\",\"type\":\"bool\"}],\"name\":\"CooldownEnabledUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"InstantWithdrawWhitelistUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousVaultRewards\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newVaultRewards\",\"type\":\"address\"}],\"name\":\"VaultRewardsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"WithdrawCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"WithdrawClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimableAt\",\"type\":\"uint256\"}],\"name\":\"WithdrawRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousDistributor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newDistributor\",\"type\":\"address\"}],\"name\":\"YieldDistributorUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_COOLDOWN_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_COOLDOWN_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_COOLDOWN_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId_\",\"type\":\"uint256\"}],\"name\":\"cancelWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"}],\"name\":\"claimWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"requestIds_\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"}],\"name\":\"claimWithdrawBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAssets_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cooldownDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cooldownEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"getActiveRequestIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"getClaimableRequests\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"requestIds_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"assets_\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"getPendingRequests\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"requestIds_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"assets_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"claimableAt_\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId_\",\"type\":\"uint256\"}],\"name\":\"getRequestDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimableAt\",\"type\":\"uint256\"}],\"internalType\":\"structIStakingVault.CooldownRequest\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"instantWithdrawWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"requestRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"requestWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssetsInCooldown\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"duration_\",\"type\":\"uint256\"}],\"name\":\"updateCooldownDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled_\",\"type\":\"bool\"}],\"name\":\"updateCooldownEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"status_\",\"type\":\"bool\"}],\"name\":\"updateInstantWithdrawWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vaultRewards_\",\"type\":\"address\"}],\"name\":\"updateVaultRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"distributor_\",\"type\":\"address\"}],\"name\":\"updateYieldDistributor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultRewards\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"yieldDistributor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SvusdABI is the input ABI used to generate the binding from.
// Deprecated: Use SvusdMetaData.ABI instead.
var SvusdABI = SvusdMetaData.ABI

// Svusd is an auto generated Go binding around an Ethereum contract.
type Svusd struct {
	SvusdCaller     // Read-only binding to the contract
	SvusdTransactor // Write-only binding to the contract
	SvusdFilterer   // Log filterer for contract events
}

// SvusdCaller is an auto generated read-only Go binding around an Ethereum contract.
type SvusdCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SvusdTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SvusdTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SvusdFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SvusdFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SvusdSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SvusdSession struct {
	Contract     *Svusd            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SvusdCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SvusdCallerSession struct {
	Contract *SvusdCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SvusdTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SvusdTransactorSession struct {
	Contract     *SvusdTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SvusdRaw is an auto generated low-level Go binding around an Ethereum contract.
type SvusdRaw struct {
	Contract *Svusd // Generic contract binding to access the raw methods on
}

// SvusdCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SvusdCallerRaw struct {
	Contract *SvusdCaller // Generic read-only contract binding to access the raw methods on
}

// SvusdTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SvusdTransactorRaw struct {
	Contract *SvusdTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSvusd creates a new instance of Svusd, bound to a specific deployed contract.
func NewSvusd(address common.Address, backend bind.ContractBackend) (*Svusd, error) {
	contract, err := bindSvusd(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Svusd{SvusdCaller: SvusdCaller{contract: contract}, SvusdTransactor: SvusdTransactor{contract: contract}, SvusdFilterer: SvusdFilterer{contract: contract}}, nil
}

// NewSvusdCaller creates a new read-only instance of Svusd, bound to a specific deployed contract.
func NewSvusdCaller(address common.Address, caller bind.ContractCaller) (*SvusdCaller, error) {
	contract, err := bindSvusd(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SvusdCaller{contract: contract}, nil
}

// NewSvusdTransactor creates a new write-only instance of Svusd, bound to a specific deployed contract.
func NewSvusdTransactor(address common.Address, transactor bind.ContractTransactor) (*SvusdTransactor, error) {
	contract, err := bindSvusd(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SvusdTransactor{contract: contract}, nil
}

// NewSvusdFilterer creates a new log filterer instance of Svusd, bound to a specific deployed contract.
func NewSvusdFilterer(address common.Address, filterer bind.ContractFilterer) (*SvusdFilterer, error) {
	contract, err := bindSvusd(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SvusdFilterer{contract: contract}, nil
}

// bindSvusd binds a generic wrapper to an already deployed contract.
func bindSvusd(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SvusdMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Svusd *SvusdRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Svusd.Contract.SvusdCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Svusd *SvusdRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Svusd.Contract.SvusdTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Svusd *SvusdRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Svusd.Contract.SvusdTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Svusd *SvusdCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Svusd.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Svusd *SvusdTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Svusd.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Svusd *SvusdTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Svusd.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTCOOLDOWNDURATION is a free data retrieval call binding the contract method 0x6d1fee01.
//
// Solidity: function DEFAULT_COOLDOWN_DURATION() view returns(uint256)
func (_Svusd *SvusdCaller) DEFAULTCOOLDOWNDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Svusd.contract.Call(opts, &out, "DEFAULT_COOLDOWN_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTCOOLDOWNDURATION is a free data retrieval call binding the contract method 0x6d1fee01.
//
// Solidity: function DEFAULT_COOLDOWN_DURATION() view returns(uint256)
func (_Svusd *SvusdSession) DEFAULTCOOLDOWNDURATION() (*big.Int, error) {
	return _Svusd.Contract.DEFAULTCOOLDOWNDURATION(&_Svusd.CallOpts)
}

// DEFAULTCOOLDOWNDURATION is a free data retrieval call binding the contract method 0x6d1fee01.
//
// Solidity: function DEFAULT_COOLDOWN_DURATION() view returns(uint256)
func (_Svusd *SvusdCallerSession) DEFAULTCOOLDOWNDURATION() (*big.Int, error) {
	return _Svusd.Contract.DEFAULTCOOLDOWNDURATION(&_Svusd.CallOpts)
}

// MAXCOOLDOWNDURATION is a free data retrieval call binding the contract method 0x1e9049cf.
//
// Solidity: function MAX_COOLDOWN_DURATION() view returns(uint256)
func (_Svusd *SvusdCaller) MAXCOOLDOWNDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Svusd.contract.Call(opts, &out, "MAX_COOLDOWN_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXCOOLDOWNDURATION is a free data retrieval call binding the contract method 0x1e9049cf.
//
// Solidity: function MAX_COOLDOWN_DURATION() view returns(uint256)
func (_Svusd *SvusdSession) MAXCOOLDOWNDURATION() (*big.Int, error) {
	return _Svusd.Contract.MAXCOOLDOWNDURATION(&_Svusd.CallOpts)
}

// MAXCOOLDOWNDURATION is a free data retrieval call binding the contract method 0x1e9049cf.
//
// Solidity: function MAX_COOLDOWN_DURATION() view returns(uint256)
func (_Svusd *SvusdCallerSession) MAXCOOLDOWNDURATION() (*big.Int, error) {
	return _Svusd.Contract.MAXCOOLDOWNDURATION(&_Svusd.CallOpts)
}

// MINCOOLDOWNDURATION is a free data retrieval call binding the contract method 0x5faab043.
//
// Solidity: function MIN_COOLDOWN_DURATION() view returns(uint256)
func (_Svusd *SvusdCaller) MINCOOLDOWNDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Svusd.contract.Call(opts, &out, "MIN_COOLDOWN_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINCOOLDOWNDURATION is a free data retrieval call binding the contract method 0x5faab043.
//
// Solidity: function MIN_COOLDOWN_DURATION() view returns(uint256)
func (_Svusd *SvusdSession) MINCOOLDOWNDURATION() (*big.Int, error) {
	return _Svusd.Contract.MINCOOLDOWNDURATION(&_Svusd.CallOpts)
}

// MINCOOLDOWNDURATION is a free data retrieval call binding the contract method 0x5faab043.
//
// Solidity: function MIN_COOLDOWN_DURATION() view returns(uint256)
func (_Svusd *SvusdCallerSession) MINCOOLDOWNDURATION() (*big.Int, error) {
	return _Svusd.Contract.MINCOOLDOWNDURATION(&_Svusd.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Svusd *SvusdCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Svusd.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Svusd *SvusdSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Svusd.Contract.Allowance(&_Svusd.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Svusd *SvusdCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Svusd.Contract.Allowance(&_Svusd.CallOpts, owner, spender)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_Svusd *SvusdCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Svusd.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_Svusd *SvusdSession) Asset() (common.Address, error) {
	return _Svusd.Contract.Asset(&_Svusd.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_Svusd *SvusdCallerSession) Asset() (common.Address, error) {
	return _Svusd.Contract.Asset(&_Svusd.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Svusd *SvusdCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Svusd.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Svusd *SvusdSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Svusd.Contract.BalanceOf(&_Svusd.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Svusd *SvusdCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Svusd.Contract.BalanceOf(&_Svusd.CallOpts, account)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_Svusd *SvusdCaller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Svusd.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_Svusd *SvusdSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _Svusd.Contract.ConvertToAssets(&_Svusd.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_Svusd *SvusdCallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _Svusd.Contract.ConvertToAssets(&_Svusd.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_Svusd *SvusdCaller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Svusd.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_Svusd *SvusdSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _Svusd.Contract.ConvertToShares(&_Svusd.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_Svusd *SvusdCallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _Svusd.Contract.ConvertToShares(&_Svusd.CallOpts, assets)
}

// CooldownDuration is a free data retrieval call binding the contract method 0x35269315.
//
// Solidity: function cooldownDuration() view returns(uint256)
func (_Svusd *SvusdCaller) CooldownDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Svusd.contract.Call(opts, &out, "cooldownDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CooldownDuration is a free data retrieval call binding the contract method 0x35269315.
//
// Solidity: function cooldownDuration() view returns(uint256)
func (_Svusd *SvusdSession) CooldownDuration() (*big.Int, error) {
	return _Svusd.Contract.CooldownDuration(&_Svusd.CallOpts)
}

// CooldownDuration is a free data retrieval call binding the contract method 0x35269315.
//
// Solidity: function cooldownDuration() view returns(uint256)
func (_Svusd *SvusdCallerSession) CooldownDuration() (*big.Int, error) {
	return _Svusd.Contract.CooldownDuration(&_Svusd.CallOpts)
}

// CooldownEnabled is a free data retrieval call binding the contract method 0xa985ceef.
//
// Solidity: function cooldownEnabled() view returns(bool)
func (_Svusd *SvusdCaller) CooldownEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Svusd.contract.Call(opts, &out, "cooldownEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CooldownEnabled is a free data retrieval call binding the contract method 0xa985ceef.
//
// Solidity: function cooldownEnabled() view returns(bool)
func (_Svusd *SvusdSession) CooldownEnabled() (bool, error) {
	return _Svusd.Contract.CooldownEnabled(&_Svusd.CallOpts)
}

// CooldownEnabled is a free data retrieval call binding the contract method 0xa985ceef.
//
// Solidity: function cooldownEnabled() view returns(bool)
func (_Svusd *SvusdCallerSession) CooldownEnabled() (bool, error) {
	return _Svusd.Contract.CooldownEnabled(&_Svusd.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Svusd *SvusdCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Svusd.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Svusd *SvusdSession) Decimals() (uint8, error) {
	return _Svusd.Contract.Decimals(&_Svusd.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Svusd *SvusdCallerSession) Decimals() (uint8, error) {
	return _Svusd.Contract.Decimals(&_Svusd.CallOpts)
}

// GetActiveRequestIds is a free data retrieval call binding the contract method 0x214e38d6.
//
// Solidity: function getActiveRequestIds(address account_) view returns(uint256[])
func (_Svusd *SvusdCaller) GetActiveRequestIds(opts *bind.CallOpts, account_ common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Svusd.contract.Call(opts, &out, "getActiveRequestIds", account_)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetActiveRequestIds is a free data retrieval call binding the contract method 0x214e38d6.
//
// Solidity: function getActiveRequestIds(address account_) view returns(uint256[])
func (_Svusd *SvusdSession) GetActiveRequestIds(account_ common.Address) ([]*big.Int, error) {
	return _Svusd.Contract.GetActiveRequestIds(&_Svusd.CallOpts, account_)
}

// GetActiveRequestIds is a free data retrieval call binding the contract method 0x214e38d6.
//
// Solidity: function getActiveRequestIds(address account_) view returns(uint256[])
func (_Svusd *SvusdCallerSession) GetActiveRequestIds(account_ common.Address) ([]*big.Int, error) {
	return _Svusd.Contract.GetActiveRequestIds(&_Svusd.CallOpts, account_)
}

// GetClaimableRequests is a free data retrieval call binding the contract method 0xc7786a50.
//
// Solidity: function getClaimableRequests(address account_) view returns(uint256[] requestIds_, uint256[] assets_)
func (_Svusd *SvusdCaller) GetClaimableRequests(opts *bind.CallOpts, account_ common.Address) (struct {
	RequestIds []*big.Int
	Assets     []*big.Int
}, error) {
	var out []interface{}
	err := _Svusd.contract.Call(opts, &out, "getClaimableRequests", account_)

	outstruct := new(struct {
		RequestIds []*big.Int
		Assets     []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RequestIds = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Assets = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetClaimableRequests is a free data retrieval call binding the contract method 0xc7786a50.
//
// Solidity: function getClaimableRequests(address account_) view returns(uint256[] requestIds_, uint256[] assets_)
func (_Svusd *SvusdSession) GetClaimableRequests(account_ common.Address) (struct {
	RequestIds []*big.Int
	Assets     []*big.Int
}, error) {
	return _Svusd.Contract.GetClaimableRequests(&_Svusd.CallOpts, account_)
}

// GetClaimableRequests is a free data retrieval call binding the contract method 0xc7786a50.
//
// Solidity: function getClaimableRequests(address account_) view returns(uint256[] requestIds_, uint256[] assets_)
func (_Svusd *SvusdCallerSession) GetClaimableRequests(account_ common.Address) (struct {
	RequestIds []*big.Int
	Assets     []*big.Int
}, error) {
	return _Svusd.Contract.GetClaimableRequests(&_Svusd.CallOpts, account_)
}

// GetPendingRequests is a free data retrieval call binding the contract method 0xf05bfa7b.
//
// Solidity: function getPendingRequests(address account_) view returns(uint256[] requestIds_, uint256[] assets_, uint256[] claimableAt_)
func (_Svusd *SvusdCaller) GetPendingRequests(opts *bind.CallOpts, account_ common.Address) (struct {
	RequestIds  []*big.Int
	Assets      []*big.Int
	ClaimableAt []*big.Int
}, error) {
	var out []interface{}
	err := _Svusd.contract.Call(opts, &out, "getPendingRequests", account_)

	outstruct := new(struct {
		RequestIds  []*big.Int
		Assets      []*big.Int
		ClaimableAt []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RequestIds = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Assets = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.ClaimableAt = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetPendingRequests is a free data retrieval call binding the contract method 0xf05bfa7b.
//
// Solidity: function getPendingRequests(address account_) view returns(uint256[] requestIds_, uint256[] assets_, uint256[] claimableAt_)
func (_Svusd *SvusdSession) GetPendingRequests(account_ common.Address) (struct {
	RequestIds  []*big.Int
	Assets      []*big.Int
	ClaimableAt []*big.Int
}, error) {
	return _Svusd.Contract.GetPendingRequests(&_Svusd.CallOpts, account_)
}

// GetPendingRequests is a free data retrieval call binding the contract method 0xf05bfa7b.
//
// Solidity: function getPendingRequests(address account_) view returns(uint256[] requestIds_, uint256[] assets_, uint256[] claimableAt_)
func (_Svusd *SvusdCallerSession) GetPendingRequests(account_ common.Address) (struct {
	RequestIds  []*big.Int
	Assets      []*big.Int
	ClaimableAt []*big.Int
}, error) {
	return _Svusd.Contract.GetPendingRequests(&_Svusd.CallOpts, account_)
}

// GetRequestDetails is a free data retrieval call binding the contract method 0xf34d4c63.
//
// Solidity: function getRequestDetails(uint256 requestId_) view returns((address,uint256,uint256))
func (_Svusd *SvusdCaller) GetRequestDetails(opts *bind.CallOpts, requestId_ *big.Int) (IStakingVaultCooldownRequest, error) {
	var out []interface{}
	err := _Svusd.contract.Call(opts, &out, "getRequestDetails", requestId_)

	if err != nil {
		return *new(IStakingVaultCooldownRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(IStakingVaultCooldownRequest)).(*IStakingVaultCooldownRequest)

	return out0, err

}

// GetRequestDetails is a free data retrieval call binding the contract method 0xf34d4c63.
//
// Solidity: function getRequestDetails(uint256 requestId_) view returns((address,uint256,uint256))
func (_Svusd *SvusdSession) GetRequestDetails(requestId_ *big.Int) (IStakingVaultCooldownRequest, error) {
	return _Svusd.Contract.GetRequestDetails(&_Svusd.CallOpts, requestId_)
}

// GetRequestDetails is a free data retrieval call binding the contract method 0xf34d4c63.
//
// Solidity: function getRequestDetails(uint256 requestId_) view returns((address,uint256,uint256))
func (_Svusd *SvusdCallerSession) GetRequestDetails(requestId_ *big.Int) (IStakingVaultCooldownRequest, error) {
	return _Svusd.Contract.GetRequestDetails(&_Svusd.CallOpts, requestId_)
}

// InstantWithdrawWhitelist is a free data retrieval call binding the contract method 0xdf2f6bb1.
//
// Solidity: function instantWithdrawWhitelist(address account_) view returns(bool)
func (_Svusd *SvusdCaller) InstantWithdrawWhitelist(opts *bind.CallOpts, account_ common.Address) (bool, error) {
	var out []interface{}
	err := _Svusd.contract.Call(opts, &out, "instantWithdrawWhitelist", account_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InstantWithdrawWhitelist is a free data retrieval call binding the contract method 0xdf2f6bb1.
//
// Solidity: function instantWithdrawWhitelist(address account_) view returns(bool)
func (_Svusd *SvusdSession) InstantWithdrawWhitelist(account_ common.Address) (bool, error) {
	return _Svusd.Contract.InstantWithdrawWhitelist(&_Svusd.CallOpts, account_)
}

// InstantWithdrawWhitelist is a free data retrieval call binding the contract method 0xdf2f6bb1.
//
// Solidity: function instantWithdrawWhitelist(address account_) view returns(bool)
func (_Svusd *SvusdCallerSession) InstantWithdrawWhitelist(account_ common.Address) (bool, error) {
	return _Svusd.Contract.InstantWithdrawWhitelist(&_Svusd.CallOpts, account_)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_Svusd *SvusdCaller) MaxDeposit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Svusd.contract.Call(opts, &out, "maxDeposit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_Svusd *SvusdSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _Svusd.Contract.MaxDeposit(&_Svusd.CallOpts, arg0)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_Svusd *SvusdCallerSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _Svusd.Contract.MaxDeposit(&_Svusd.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_Svusd *SvusdCaller) MaxMint(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Svusd.contract.Call(opts, &out, "maxMint", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_Svusd *SvusdSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _Svusd.Contract.MaxMint(&_Svusd.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_Svusd *SvusdCallerSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _Svusd.Contract.MaxMint(&_Svusd.CallOpts, arg0)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_Svusd *SvusdCaller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Svusd.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_Svusd *SvusdSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _Svusd.Contract.MaxRedeem(&_Svusd.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_Svusd *SvusdCallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _Svusd.Contract.MaxRedeem(&_Svusd.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_Svusd *SvusdCaller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Svusd.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_Svusd *SvusdSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _Svusd.Contract.MaxWithdraw(&_Svusd.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_Svusd *SvusdCallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _Svusd.Contract.MaxWithdraw(&_Svusd.CallOpts, owner)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Svusd *SvusdCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Svusd.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Svusd *SvusdSession) Name() (string, error) {
	return _Svusd.Contract.Name(&_Svusd.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Svusd *SvusdCallerSession) Name() (string, error) {
	return _Svusd.Contract.Name(&_Svusd.CallOpts)
}

// NextRequestId is a free data retrieval call binding the contract method 0x6a84a985.
//
// Solidity: function nextRequestId() view returns(uint256)
func (_Svusd *SvusdCaller) NextRequestId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Svusd.contract.Call(opts, &out, "nextRequestId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextRequestId is a free data retrieval call binding the contract method 0x6a84a985.
//
// Solidity: function nextRequestId() view returns(uint256)
func (_Svusd *SvusdSession) NextRequestId() (*big.Int, error) {
	return _Svusd.Contract.NextRequestId(&_Svusd.CallOpts)
}

// NextRequestId is a free data retrieval call binding the contract method 0x6a84a985.
//
// Solidity: function nextRequestId() view returns(uint256)
func (_Svusd *SvusdCallerSession) NextRequestId() (*big.Int, error) {
	return _Svusd.Contract.NextRequestId(&_Svusd.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Svusd *SvusdCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Svusd.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Svusd *SvusdSession) Owner() (common.Address, error) {
	return _Svusd.Contract.Owner(&_Svusd.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Svusd *SvusdCallerSession) Owner() (common.Address, error) {
	return _Svusd.Contract.Owner(&_Svusd.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Svusd *SvusdCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Svusd.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Svusd *SvusdSession) PendingOwner() (common.Address, error) {
	return _Svusd.Contract.PendingOwner(&_Svusd.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Svusd *SvusdCallerSession) PendingOwner() (common.Address, error) {
	return _Svusd.Contract.PendingOwner(&_Svusd.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_Svusd *SvusdCaller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Svusd.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_Svusd *SvusdSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _Svusd.Contract.PreviewDeposit(&_Svusd.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_Svusd *SvusdCallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _Svusd.Contract.PreviewDeposit(&_Svusd.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_Svusd *SvusdCaller) PreviewMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Svusd.contract.Call(opts, &out, "previewMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_Svusd *SvusdSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _Svusd.Contract.PreviewMint(&_Svusd.CallOpts, shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_Svusd *SvusdCallerSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _Svusd.Contract.PreviewMint(&_Svusd.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_Svusd *SvusdCaller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Svusd.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_Svusd *SvusdSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _Svusd.Contract.PreviewRedeem(&_Svusd.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_Svusd *SvusdCallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _Svusd.Contract.PreviewRedeem(&_Svusd.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_Svusd *SvusdCaller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Svusd.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_Svusd *SvusdSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _Svusd.Contract.PreviewWithdraw(&_Svusd.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_Svusd *SvusdCallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _Svusd.Contract.PreviewWithdraw(&_Svusd.CallOpts, assets)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Svusd *SvusdCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Svusd.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Svusd *SvusdSession) Symbol() (string, error) {
	return _Svusd.Contract.Symbol(&_Svusd.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Svusd *SvusdCallerSession) Symbol() (string, error) {
	return _Svusd.Contract.Symbol(&_Svusd.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Svusd *SvusdCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Svusd.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Svusd *SvusdSession) TotalAssets() (*big.Int, error) {
	return _Svusd.Contract.TotalAssets(&_Svusd.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Svusd *SvusdCallerSession) TotalAssets() (*big.Int, error) {
	return _Svusd.Contract.TotalAssets(&_Svusd.CallOpts)
}

// TotalAssetsInCooldown is a free data retrieval call binding the contract method 0x87493f2f.
//
// Solidity: function totalAssetsInCooldown() view returns(uint256)
func (_Svusd *SvusdCaller) TotalAssetsInCooldown(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Svusd.contract.Call(opts, &out, "totalAssetsInCooldown")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssetsInCooldown is a free data retrieval call binding the contract method 0x87493f2f.
//
// Solidity: function totalAssetsInCooldown() view returns(uint256)
func (_Svusd *SvusdSession) TotalAssetsInCooldown() (*big.Int, error) {
	return _Svusd.Contract.TotalAssetsInCooldown(&_Svusd.CallOpts)
}

// TotalAssetsInCooldown is a free data retrieval call binding the contract method 0x87493f2f.
//
// Solidity: function totalAssetsInCooldown() view returns(uint256)
func (_Svusd *SvusdCallerSession) TotalAssetsInCooldown() (*big.Int, error) {
	return _Svusd.Contract.TotalAssetsInCooldown(&_Svusd.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Svusd *SvusdCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Svusd.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Svusd *SvusdSession) TotalSupply() (*big.Int, error) {
	return _Svusd.Contract.TotalSupply(&_Svusd.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Svusd *SvusdCallerSession) TotalSupply() (*big.Int, error) {
	return _Svusd.Contract.TotalSupply(&_Svusd.CallOpts)
}

// VaultRewards is a free data retrieval call binding the contract method 0x8429d51f.
//
// Solidity: function vaultRewards() view returns(address)
func (_Svusd *SvusdCaller) VaultRewards(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Svusd.contract.Call(opts, &out, "vaultRewards")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultRewards is a free data retrieval call binding the contract method 0x8429d51f.
//
// Solidity: function vaultRewards() view returns(address)
func (_Svusd *SvusdSession) VaultRewards() (common.Address, error) {
	return _Svusd.Contract.VaultRewards(&_Svusd.CallOpts)
}

// VaultRewards is a free data retrieval call binding the contract method 0x8429d51f.
//
// Solidity: function vaultRewards() view returns(address)
func (_Svusd *SvusdCallerSession) VaultRewards() (common.Address, error) {
	return _Svusd.Contract.VaultRewards(&_Svusd.CallOpts)
}

// YieldDistributor is a free data retrieval call binding the contract method 0xa38c2c3d.
//
// Solidity: function yieldDistributor() view returns(address)
func (_Svusd *SvusdCaller) YieldDistributor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Svusd.contract.Call(opts, &out, "yieldDistributor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// YieldDistributor is a free data retrieval call binding the contract method 0xa38c2c3d.
//
// Solidity: function yieldDistributor() view returns(address)
func (_Svusd *SvusdSession) YieldDistributor() (common.Address, error) {
	return _Svusd.Contract.YieldDistributor(&_Svusd.CallOpts)
}

// YieldDistributor is a free data retrieval call binding the contract method 0xa38c2c3d.
//
// Solidity: function yieldDistributor() view returns(address)
func (_Svusd *SvusdCallerSession) YieldDistributor() (common.Address, error) {
	return _Svusd.Contract.YieldDistributor(&_Svusd.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Svusd *SvusdTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Svusd.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Svusd *SvusdSession) AcceptOwnership() (*types.Transaction, error) {
	return _Svusd.Contract.AcceptOwnership(&_Svusd.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Svusd *SvusdTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Svusd.Contract.AcceptOwnership(&_Svusd.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Svusd *SvusdTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Svusd.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Svusd *SvusdSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Svusd.Contract.Approve(&_Svusd.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Svusd *SvusdTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Svusd.Contract.Approve(&_Svusd.TransactOpts, spender, value)
}

// CancelWithdraw is a paid mutator transaction binding the contract method 0x9f01f7ba.
//
// Solidity: function cancelWithdraw(uint256 requestId_) returns(uint256 shares_)
func (_Svusd *SvusdTransactor) CancelWithdraw(opts *bind.TransactOpts, requestId_ *big.Int) (*types.Transaction, error) {
	return _Svusd.contract.Transact(opts, "cancelWithdraw", requestId_)
}

// CancelWithdraw is a paid mutator transaction binding the contract method 0x9f01f7ba.
//
// Solidity: function cancelWithdraw(uint256 requestId_) returns(uint256 shares_)
func (_Svusd *SvusdSession) CancelWithdraw(requestId_ *big.Int) (*types.Transaction, error) {
	return _Svusd.Contract.CancelWithdraw(&_Svusd.TransactOpts, requestId_)
}

// CancelWithdraw is a paid mutator transaction binding the contract method 0x9f01f7ba.
//
// Solidity: function cancelWithdraw(uint256 requestId_) returns(uint256 shares_)
func (_Svusd *SvusdTransactorSession) CancelWithdraw(requestId_ *big.Int) (*types.Transaction, error) {
	return _Svusd.Contract.CancelWithdraw(&_Svusd.TransactOpts, requestId_)
}

// ClaimWithdraw is a paid mutator transaction binding the contract method 0xb708ae50.
//
// Solidity: function claimWithdraw(uint256 requestId_, address receiver_) returns(uint256 assets)
func (_Svusd *SvusdTransactor) ClaimWithdraw(opts *bind.TransactOpts, requestId_ *big.Int, receiver_ common.Address) (*types.Transaction, error) {
	return _Svusd.contract.Transact(opts, "claimWithdraw", requestId_, receiver_)
}

// ClaimWithdraw is a paid mutator transaction binding the contract method 0xb708ae50.
//
// Solidity: function claimWithdraw(uint256 requestId_, address receiver_) returns(uint256 assets)
func (_Svusd *SvusdSession) ClaimWithdraw(requestId_ *big.Int, receiver_ common.Address) (*types.Transaction, error) {
	return _Svusd.Contract.ClaimWithdraw(&_Svusd.TransactOpts, requestId_, receiver_)
}

// ClaimWithdraw is a paid mutator transaction binding the contract method 0xb708ae50.
//
// Solidity: function claimWithdraw(uint256 requestId_, address receiver_) returns(uint256 assets)
func (_Svusd *SvusdTransactorSession) ClaimWithdraw(requestId_ *big.Int, receiver_ common.Address) (*types.Transaction, error) {
	return _Svusd.Contract.ClaimWithdraw(&_Svusd.TransactOpts, requestId_, receiver_)
}

// ClaimWithdrawBatch is a paid mutator transaction binding the contract method 0x8be34038.
//
// Solidity: function claimWithdrawBatch(uint256[] requestIds_, address receiver_) returns(uint256 totalAssets_)
func (_Svusd *SvusdTransactor) ClaimWithdrawBatch(opts *bind.TransactOpts, requestIds_ []*big.Int, receiver_ common.Address) (*types.Transaction, error) {
	return _Svusd.contract.Transact(opts, "claimWithdrawBatch", requestIds_, receiver_)
}

// ClaimWithdrawBatch is a paid mutator transaction binding the contract method 0x8be34038.
//
// Solidity: function claimWithdrawBatch(uint256[] requestIds_, address receiver_) returns(uint256 totalAssets_)
func (_Svusd *SvusdSession) ClaimWithdrawBatch(requestIds_ []*big.Int, receiver_ common.Address) (*types.Transaction, error) {
	return _Svusd.Contract.ClaimWithdrawBatch(&_Svusd.TransactOpts, requestIds_, receiver_)
}

// ClaimWithdrawBatch is a paid mutator transaction binding the contract method 0x8be34038.
//
// Solidity: function claimWithdrawBatch(uint256[] requestIds_, address receiver_) returns(uint256 totalAssets_)
func (_Svusd *SvusdTransactorSession) ClaimWithdrawBatch(requestIds_ []*big.Int, receiver_ common.Address) (*types.Transaction, error) {
	return _Svusd.Contract.ClaimWithdrawBatch(&_Svusd.TransactOpts, requestIds_, receiver_)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets_, address receiver_) returns(uint256)
func (_Svusd *SvusdTransactor) Deposit(opts *bind.TransactOpts, assets_ *big.Int, receiver_ common.Address) (*types.Transaction, error) {
	return _Svusd.contract.Transact(opts, "deposit", assets_, receiver_)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets_, address receiver_) returns(uint256)
func (_Svusd *SvusdSession) Deposit(assets_ *big.Int, receiver_ common.Address) (*types.Transaction, error) {
	return _Svusd.Contract.Deposit(&_Svusd.TransactOpts, assets_, receiver_)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets_, address receiver_) returns(uint256)
func (_Svusd *SvusdTransactorSession) Deposit(assets_ *big.Int, receiver_ common.Address) (*types.Transaction, error) {
	return _Svusd.Contract.Deposit(&_Svusd.TransactOpts, assets_, receiver_)
}

// Initialize is a paid mutator transaction binding the contract method 0x613d25bb.
//
// Solidity: function initialize(address asset_, string name_, string symbol_, address owner_) returns()
func (_Svusd *SvusdTransactor) Initialize(opts *bind.TransactOpts, asset_ common.Address, name_ string, symbol_ string, owner_ common.Address) (*types.Transaction, error) {
	return _Svusd.contract.Transact(opts, "initialize", asset_, name_, symbol_, owner_)
}

// Initialize is a paid mutator transaction binding the contract method 0x613d25bb.
//
// Solidity: function initialize(address asset_, string name_, string symbol_, address owner_) returns()
func (_Svusd *SvusdSession) Initialize(asset_ common.Address, name_ string, symbol_ string, owner_ common.Address) (*types.Transaction, error) {
	return _Svusd.Contract.Initialize(&_Svusd.TransactOpts, asset_, name_, symbol_, owner_)
}

// Initialize is a paid mutator transaction binding the contract method 0x613d25bb.
//
// Solidity: function initialize(address asset_, string name_, string symbol_, address owner_) returns()
func (_Svusd *SvusdTransactorSession) Initialize(asset_ common.Address, name_ string, symbol_ string, owner_ common.Address) (*types.Transaction, error) {
	return _Svusd.Contract.Initialize(&_Svusd.TransactOpts, asset_, name_, symbol_, owner_)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares_, address receiver_) returns(uint256)
func (_Svusd *SvusdTransactor) Mint(opts *bind.TransactOpts, shares_ *big.Int, receiver_ common.Address) (*types.Transaction, error) {
	return _Svusd.contract.Transact(opts, "mint", shares_, receiver_)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares_, address receiver_) returns(uint256)
func (_Svusd *SvusdSession) Mint(shares_ *big.Int, receiver_ common.Address) (*types.Transaction, error) {
	return _Svusd.Contract.Mint(&_Svusd.TransactOpts, shares_, receiver_)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares_, address receiver_) returns(uint256)
func (_Svusd *SvusdTransactorSession) Mint(shares_ *big.Int, receiver_ common.Address) (*types.Transaction, error) {
	return _Svusd.Contract.Mint(&_Svusd.TransactOpts, shares_, receiver_)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares_, address receiver_, address owner_) returns(uint256)
func (_Svusd *SvusdTransactor) Redeem(opts *bind.TransactOpts, shares_ *big.Int, receiver_ common.Address, owner_ common.Address) (*types.Transaction, error) {
	return _Svusd.contract.Transact(opts, "redeem", shares_, receiver_, owner_)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares_, address receiver_, address owner_) returns(uint256)
func (_Svusd *SvusdSession) Redeem(shares_ *big.Int, receiver_ common.Address, owner_ common.Address) (*types.Transaction, error) {
	return _Svusd.Contract.Redeem(&_Svusd.TransactOpts, shares_, receiver_, owner_)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares_, address receiver_, address owner_) returns(uint256)
func (_Svusd *SvusdTransactorSession) Redeem(shares_ *big.Int, receiver_ common.Address, owner_ common.Address) (*types.Transaction, error) {
	return _Svusd.Contract.Redeem(&_Svusd.TransactOpts, shares_, receiver_, owner_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Svusd *SvusdTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Svusd.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Svusd *SvusdSession) RenounceOwnership() (*types.Transaction, error) {
	return _Svusd.Contract.RenounceOwnership(&_Svusd.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Svusd *SvusdTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Svusd.Contract.RenounceOwnership(&_Svusd.TransactOpts)
}

// RequestRedeem is a paid mutator transaction binding the contract method 0x107703ab.
//
// Solidity: function requestRedeem(uint256 shares_, address owner_) returns(uint256 requestId, uint256 assets)
func (_Svusd *SvusdTransactor) RequestRedeem(opts *bind.TransactOpts, shares_ *big.Int, owner_ common.Address) (*types.Transaction, error) {
	return _Svusd.contract.Transact(opts, "requestRedeem", shares_, owner_)
}

// RequestRedeem is a paid mutator transaction binding the contract method 0x107703ab.
//
// Solidity: function requestRedeem(uint256 shares_, address owner_) returns(uint256 requestId, uint256 assets)
func (_Svusd *SvusdSession) RequestRedeem(shares_ *big.Int, owner_ common.Address) (*types.Transaction, error) {
	return _Svusd.Contract.RequestRedeem(&_Svusd.TransactOpts, shares_, owner_)
}

// RequestRedeem is a paid mutator transaction binding the contract method 0x107703ab.
//
// Solidity: function requestRedeem(uint256 shares_, address owner_) returns(uint256 requestId, uint256 assets)
func (_Svusd *SvusdTransactorSession) RequestRedeem(shares_ *big.Int, owner_ common.Address) (*types.Transaction, error) {
	return _Svusd.Contract.RequestRedeem(&_Svusd.TransactOpts, shares_, owner_)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0xccc143b8.
//
// Solidity: function requestWithdraw(uint256 assets_, address owner_) returns(uint256 requestId, uint256 shares)
func (_Svusd *SvusdTransactor) RequestWithdraw(opts *bind.TransactOpts, assets_ *big.Int, owner_ common.Address) (*types.Transaction, error) {
	return _Svusd.contract.Transact(opts, "requestWithdraw", assets_, owner_)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0xccc143b8.
//
// Solidity: function requestWithdraw(uint256 assets_, address owner_) returns(uint256 requestId, uint256 shares)
func (_Svusd *SvusdSession) RequestWithdraw(assets_ *big.Int, owner_ common.Address) (*types.Transaction, error) {
	return _Svusd.Contract.RequestWithdraw(&_Svusd.TransactOpts, assets_, owner_)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0xccc143b8.
//
// Solidity: function requestWithdraw(uint256 assets_, address owner_) returns(uint256 requestId, uint256 shares)
func (_Svusd *SvusdTransactorSession) RequestWithdraw(assets_ *big.Int, owner_ common.Address) (*types.Transaction, error) {
	return _Svusd.Contract.RequestWithdraw(&_Svusd.TransactOpts, assets_, owner_)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Svusd *SvusdTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Svusd.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Svusd *SvusdSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Svusd.Contract.Transfer(&_Svusd.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Svusd *SvusdTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Svusd.Contract.Transfer(&_Svusd.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Svusd *SvusdTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Svusd.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Svusd *SvusdSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Svusd.Contract.TransferFrom(&_Svusd.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Svusd *SvusdTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Svusd.Contract.TransferFrom(&_Svusd.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Svusd *SvusdTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Svusd.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Svusd *SvusdSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Svusd.Contract.TransferOwnership(&_Svusd.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Svusd *SvusdTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Svusd.Contract.TransferOwnership(&_Svusd.TransactOpts, newOwner)
}

// UpdateCooldownDuration is a paid mutator transaction binding the contract method 0x07a2c653.
//
// Solidity: function updateCooldownDuration(uint256 duration_) returns()
func (_Svusd *SvusdTransactor) UpdateCooldownDuration(opts *bind.TransactOpts, duration_ *big.Int) (*types.Transaction, error) {
	return _Svusd.contract.Transact(opts, "updateCooldownDuration", duration_)
}

// UpdateCooldownDuration is a paid mutator transaction binding the contract method 0x07a2c653.
//
// Solidity: function updateCooldownDuration(uint256 duration_) returns()
func (_Svusd *SvusdSession) UpdateCooldownDuration(duration_ *big.Int) (*types.Transaction, error) {
	return _Svusd.Contract.UpdateCooldownDuration(&_Svusd.TransactOpts, duration_)
}

// UpdateCooldownDuration is a paid mutator transaction binding the contract method 0x07a2c653.
//
// Solidity: function updateCooldownDuration(uint256 duration_) returns()
func (_Svusd *SvusdTransactorSession) UpdateCooldownDuration(duration_ *big.Int) (*types.Transaction, error) {
	return _Svusd.Contract.UpdateCooldownDuration(&_Svusd.TransactOpts, duration_)
}

// UpdateCooldownEnabled is a paid mutator transaction binding the contract method 0xf05ca034.
//
// Solidity: function updateCooldownEnabled(bool enabled_) returns()
func (_Svusd *SvusdTransactor) UpdateCooldownEnabled(opts *bind.TransactOpts, enabled_ bool) (*types.Transaction, error) {
	return _Svusd.contract.Transact(opts, "updateCooldownEnabled", enabled_)
}

// UpdateCooldownEnabled is a paid mutator transaction binding the contract method 0xf05ca034.
//
// Solidity: function updateCooldownEnabled(bool enabled_) returns()
func (_Svusd *SvusdSession) UpdateCooldownEnabled(enabled_ bool) (*types.Transaction, error) {
	return _Svusd.Contract.UpdateCooldownEnabled(&_Svusd.TransactOpts, enabled_)
}

// UpdateCooldownEnabled is a paid mutator transaction binding the contract method 0xf05ca034.
//
// Solidity: function updateCooldownEnabled(bool enabled_) returns()
func (_Svusd *SvusdTransactorSession) UpdateCooldownEnabled(enabled_ bool) (*types.Transaction, error) {
	return _Svusd.Contract.UpdateCooldownEnabled(&_Svusd.TransactOpts, enabled_)
}

// UpdateInstantWithdrawWhitelist is a paid mutator transaction binding the contract method 0x1c37830c.
//
// Solidity: function updateInstantWithdrawWhitelist(address account_, bool status_) returns()
func (_Svusd *SvusdTransactor) UpdateInstantWithdrawWhitelist(opts *bind.TransactOpts, account_ common.Address, status_ bool) (*types.Transaction, error) {
	return _Svusd.contract.Transact(opts, "updateInstantWithdrawWhitelist", account_, status_)
}

// UpdateInstantWithdrawWhitelist is a paid mutator transaction binding the contract method 0x1c37830c.
//
// Solidity: function updateInstantWithdrawWhitelist(address account_, bool status_) returns()
func (_Svusd *SvusdSession) UpdateInstantWithdrawWhitelist(account_ common.Address, status_ bool) (*types.Transaction, error) {
	return _Svusd.Contract.UpdateInstantWithdrawWhitelist(&_Svusd.TransactOpts, account_, status_)
}

// UpdateInstantWithdrawWhitelist is a paid mutator transaction binding the contract method 0x1c37830c.
//
// Solidity: function updateInstantWithdrawWhitelist(address account_, bool status_) returns()
func (_Svusd *SvusdTransactorSession) UpdateInstantWithdrawWhitelist(account_ common.Address, status_ bool) (*types.Transaction, error) {
	return _Svusd.Contract.UpdateInstantWithdrawWhitelist(&_Svusd.TransactOpts, account_, status_)
}

// UpdateVaultRewards is a paid mutator transaction binding the contract method 0x8fa2d228.
//
// Solidity: function updateVaultRewards(address vaultRewards_) returns()
func (_Svusd *SvusdTransactor) UpdateVaultRewards(opts *bind.TransactOpts, vaultRewards_ common.Address) (*types.Transaction, error) {
	return _Svusd.contract.Transact(opts, "updateVaultRewards", vaultRewards_)
}

// UpdateVaultRewards is a paid mutator transaction binding the contract method 0x8fa2d228.
//
// Solidity: function updateVaultRewards(address vaultRewards_) returns()
func (_Svusd *SvusdSession) UpdateVaultRewards(vaultRewards_ common.Address) (*types.Transaction, error) {
	return _Svusd.Contract.UpdateVaultRewards(&_Svusd.TransactOpts, vaultRewards_)
}

// UpdateVaultRewards is a paid mutator transaction binding the contract method 0x8fa2d228.
//
// Solidity: function updateVaultRewards(address vaultRewards_) returns()
func (_Svusd *SvusdTransactorSession) UpdateVaultRewards(vaultRewards_ common.Address) (*types.Transaction, error) {
	return _Svusd.Contract.UpdateVaultRewards(&_Svusd.TransactOpts, vaultRewards_)
}

// UpdateYieldDistributor is a paid mutator transaction binding the contract method 0xd7b02c4b.
//
// Solidity: function updateYieldDistributor(address distributor_) returns()
func (_Svusd *SvusdTransactor) UpdateYieldDistributor(opts *bind.TransactOpts, distributor_ common.Address) (*types.Transaction, error) {
	return _Svusd.contract.Transact(opts, "updateYieldDistributor", distributor_)
}

// UpdateYieldDistributor is a paid mutator transaction binding the contract method 0xd7b02c4b.
//
// Solidity: function updateYieldDistributor(address distributor_) returns()
func (_Svusd *SvusdSession) UpdateYieldDistributor(distributor_ common.Address) (*types.Transaction, error) {
	return _Svusd.Contract.UpdateYieldDistributor(&_Svusd.TransactOpts, distributor_)
}

// UpdateYieldDistributor is a paid mutator transaction binding the contract method 0xd7b02c4b.
//
// Solidity: function updateYieldDistributor(address distributor_) returns()
func (_Svusd *SvusdTransactorSession) UpdateYieldDistributor(distributor_ common.Address) (*types.Transaction, error) {
	return _Svusd.Contract.UpdateYieldDistributor(&_Svusd.TransactOpts, distributor_)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets_, address receiver_, address owner_) returns(uint256)
func (_Svusd *SvusdTransactor) Withdraw(opts *bind.TransactOpts, assets_ *big.Int, receiver_ common.Address, owner_ common.Address) (*types.Transaction, error) {
	return _Svusd.contract.Transact(opts, "withdraw", assets_, receiver_, owner_)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets_, address receiver_, address owner_) returns(uint256)
func (_Svusd *SvusdSession) Withdraw(assets_ *big.Int, receiver_ common.Address, owner_ common.Address) (*types.Transaction, error) {
	return _Svusd.Contract.Withdraw(&_Svusd.TransactOpts, assets_, receiver_, owner_)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets_, address receiver_, address owner_) returns(uint256)
func (_Svusd *SvusdTransactorSession) Withdraw(assets_ *big.Int, receiver_ common.Address, owner_ common.Address) (*types.Transaction, error) {
	return _Svusd.Contract.Withdraw(&_Svusd.TransactOpts, assets_, receiver_, owner_)
}

// SvusdApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Svusd contract.
type SvusdApprovalIterator struct {
	Event *SvusdApproval // Event containing the contract specifics and raw log

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
func (it *SvusdApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SvusdApproval)
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
		it.Event = new(SvusdApproval)
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
func (it *SvusdApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SvusdApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SvusdApproval represents a Approval event raised by the Svusd contract.
type SvusdApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Svusd *SvusdFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SvusdApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Svusd.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SvusdApprovalIterator{contract: _Svusd.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Svusd *SvusdFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SvusdApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Svusd.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SvusdApproval)
				if err := _Svusd.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Svusd *SvusdFilterer) ParseApproval(log types.Log) (*SvusdApproval, error) {
	event := new(SvusdApproval)
	if err := _Svusd.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SvusdCooldownDurationUpdatedIterator is returned from FilterCooldownDurationUpdated and is used to iterate over the raw logs and unpacked data for CooldownDurationUpdated events raised by the Svusd contract.
type SvusdCooldownDurationUpdatedIterator struct {
	Event *SvusdCooldownDurationUpdated // Event containing the contract specifics and raw log

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
func (it *SvusdCooldownDurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SvusdCooldownDurationUpdated)
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
		it.Event = new(SvusdCooldownDurationUpdated)
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
func (it *SvusdCooldownDurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SvusdCooldownDurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SvusdCooldownDurationUpdated represents a CooldownDurationUpdated event raised by the Svusd contract.
type SvusdCooldownDurationUpdated struct {
	PreviousDuration *big.Int
	NewDuration      *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCooldownDurationUpdated is a free log retrieval operation binding the contract event 0x8f6ceadd6e9845ce194f7cd50e5a1ef4aec5b40c31652db44ea043bbbf27b92b.
//
// Solidity: event CooldownDurationUpdated(uint256 previousDuration, uint256 newDuration)
func (_Svusd *SvusdFilterer) FilterCooldownDurationUpdated(opts *bind.FilterOpts) (*SvusdCooldownDurationUpdatedIterator, error) {

	logs, sub, err := _Svusd.contract.FilterLogs(opts, "CooldownDurationUpdated")
	if err != nil {
		return nil, err
	}
	return &SvusdCooldownDurationUpdatedIterator{contract: _Svusd.contract, event: "CooldownDurationUpdated", logs: logs, sub: sub}, nil
}

// WatchCooldownDurationUpdated is a free log subscription operation binding the contract event 0x8f6ceadd6e9845ce194f7cd50e5a1ef4aec5b40c31652db44ea043bbbf27b92b.
//
// Solidity: event CooldownDurationUpdated(uint256 previousDuration, uint256 newDuration)
func (_Svusd *SvusdFilterer) WatchCooldownDurationUpdated(opts *bind.WatchOpts, sink chan<- *SvusdCooldownDurationUpdated) (event.Subscription, error) {

	logs, sub, err := _Svusd.contract.WatchLogs(opts, "CooldownDurationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SvusdCooldownDurationUpdated)
				if err := _Svusd.contract.UnpackLog(event, "CooldownDurationUpdated", log); err != nil {
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

// ParseCooldownDurationUpdated is a log parse operation binding the contract event 0x8f6ceadd6e9845ce194f7cd50e5a1ef4aec5b40c31652db44ea043bbbf27b92b.
//
// Solidity: event CooldownDurationUpdated(uint256 previousDuration, uint256 newDuration)
func (_Svusd *SvusdFilterer) ParseCooldownDurationUpdated(log types.Log) (*SvusdCooldownDurationUpdated, error) {
	event := new(SvusdCooldownDurationUpdated)
	if err := _Svusd.contract.UnpackLog(event, "CooldownDurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SvusdCooldownEnabledUpdatedIterator is returned from FilterCooldownEnabledUpdated and is used to iterate over the raw logs and unpacked data for CooldownEnabledUpdated events raised by the Svusd contract.
type SvusdCooldownEnabledUpdatedIterator struct {
	Event *SvusdCooldownEnabledUpdated // Event containing the contract specifics and raw log

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
func (it *SvusdCooldownEnabledUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SvusdCooldownEnabledUpdated)
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
		it.Event = new(SvusdCooldownEnabledUpdated)
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
func (it *SvusdCooldownEnabledUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SvusdCooldownEnabledUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SvusdCooldownEnabledUpdated represents a CooldownEnabledUpdated event raised by the Svusd contract.
type SvusdCooldownEnabledUpdated struct {
	PreviousStatus bool
	NewStatus      bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCooldownEnabledUpdated is a free log retrieval operation binding the contract event 0x166d9f924637aaef1818762071526936fd8768f14f982c72772024a986c44b49.
//
// Solidity: event CooldownEnabledUpdated(bool previousStatus, bool newStatus)
func (_Svusd *SvusdFilterer) FilterCooldownEnabledUpdated(opts *bind.FilterOpts) (*SvusdCooldownEnabledUpdatedIterator, error) {

	logs, sub, err := _Svusd.contract.FilterLogs(opts, "CooldownEnabledUpdated")
	if err != nil {
		return nil, err
	}
	return &SvusdCooldownEnabledUpdatedIterator{contract: _Svusd.contract, event: "CooldownEnabledUpdated", logs: logs, sub: sub}, nil
}

// WatchCooldownEnabledUpdated is a free log subscription operation binding the contract event 0x166d9f924637aaef1818762071526936fd8768f14f982c72772024a986c44b49.
//
// Solidity: event CooldownEnabledUpdated(bool previousStatus, bool newStatus)
func (_Svusd *SvusdFilterer) WatchCooldownEnabledUpdated(opts *bind.WatchOpts, sink chan<- *SvusdCooldownEnabledUpdated) (event.Subscription, error) {

	logs, sub, err := _Svusd.contract.WatchLogs(opts, "CooldownEnabledUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SvusdCooldownEnabledUpdated)
				if err := _Svusd.contract.UnpackLog(event, "CooldownEnabledUpdated", log); err != nil {
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

// ParseCooldownEnabledUpdated is a log parse operation binding the contract event 0x166d9f924637aaef1818762071526936fd8768f14f982c72772024a986c44b49.
//
// Solidity: event CooldownEnabledUpdated(bool previousStatus, bool newStatus)
func (_Svusd *SvusdFilterer) ParseCooldownEnabledUpdated(log types.Log) (*SvusdCooldownEnabledUpdated, error) {
	event := new(SvusdCooldownEnabledUpdated)
	if err := _Svusd.contract.UnpackLog(event, "CooldownEnabledUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SvusdDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Svusd contract.
type SvusdDepositIterator struct {
	Event *SvusdDeposit // Event containing the contract specifics and raw log

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
func (it *SvusdDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SvusdDeposit)
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
		it.Event = new(SvusdDeposit)
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
func (it *SvusdDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SvusdDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SvusdDeposit represents a Deposit event raised by the Svusd contract.
type SvusdDeposit struct {
	Sender common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_Svusd *SvusdFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*SvusdDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Svusd.contract.FilterLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &SvusdDepositIterator{contract: _Svusd.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_Svusd *SvusdFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *SvusdDeposit, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Svusd.contract.WatchLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SvusdDeposit)
				if err := _Svusd.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_Svusd *SvusdFilterer) ParseDeposit(log types.Log) (*SvusdDeposit, error) {
	event := new(SvusdDeposit)
	if err := _Svusd.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SvusdInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Svusd contract.
type SvusdInitializedIterator struct {
	Event *SvusdInitialized // Event containing the contract specifics and raw log

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
func (it *SvusdInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SvusdInitialized)
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
		it.Event = new(SvusdInitialized)
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
func (it *SvusdInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SvusdInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SvusdInitialized represents a Initialized event raised by the Svusd contract.
type SvusdInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Svusd *SvusdFilterer) FilterInitialized(opts *bind.FilterOpts) (*SvusdInitializedIterator, error) {

	logs, sub, err := _Svusd.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SvusdInitializedIterator{contract: _Svusd.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Svusd *SvusdFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SvusdInitialized) (event.Subscription, error) {

	logs, sub, err := _Svusd.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SvusdInitialized)
				if err := _Svusd.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Svusd *SvusdFilterer) ParseInitialized(log types.Log) (*SvusdInitialized, error) {
	event := new(SvusdInitialized)
	if err := _Svusd.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SvusdInstantWithdrawWhitelistUpdatedIterator is returned from FilterInstantWithdrawWhitelistUpdated and is used to iterate over the raw logs and unpacked data for InstantWithdrawWhitelistUpdated events raised by the Svusd contract.
type SvusdInstantWithdrawWhitelistUpdatedIterator struct {
	Event *SvusdInstantWithdrawWhitelistUpdated // Event containing the contract specifics and raw log

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
func (it *SvusdInstantWithdrawWhitelistUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SvusdInstantWithdrawWhitelistUpdated)
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
		it.Event = new(SvusdInstantWithdrawWhitelistUpdated)
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
func (it *SvusdInstantWithdrawWhitelistUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SvusdInstantWithdrawWhitelistUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SvusdInstantWithdrawWhitelistUpdated represents a InstantWithdrawWhitelistUpdated event raised by the Svusd contract.
type SvusdInstantWithdrawWhitelistUpdated struct {
	Account common.Address
	Status  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInstantWithdrawWhitelistUpdated is a free log retrieval operation binding the contract event 0xea453964d1ed5c30fcc4e0065c0703205b656d4ed8df6badadd5641c880765a5.
//
// Solidity: event InstantWithdrawWhitelistUpdated(address indexed account, bool status)
func (_Svusd *SvusdFilterer) FilterInstantWithdrawWhitelistUpdated(opts *bind.FilterOpts, account []common.Address) (*SvusdInstantWithdrawWhitelistUpdatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Svusd.contract.FilterLogs(opts, "InstantWithdrawWhitelistUpdated", accountRule)
	if err != nil {
		return nil, err
	}
	return &SvusdInstantWithdrawWhitelistUpdatedIterator{contract: _Svusd.contract, event: "InstantWithdrawWhitelistUpdated", logs: logs, sub: sub}, nil
}

// WatchInstantWithdrawWhitelistUpdated is a free log subscription operation binding the contract event 0xea453964d1ed5c30fcc4e0065c0703205b656d4ed8df6badadd5641c880765a5.
//
// Solidity: event InstantWithdrawWhitelistUpdated(address indexed account, bool status)
func (_Svusd *SvusdFilterer) WatchInstantWithdrawWhitelistUpdated(opts *bind.WatchOpts, sink chan<- *SvusdInstantWithdrawWhitelistUpdated, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Svusd.contract.WatchLogs(opts, "InstantWithdrawWhitelistUpdated", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SvusdInstantWithdrawWhitelistUpdated)
				if err := _Svusd.contract.UnpackLog(event, "InstantWithdrawWhitelistUpdated", log); err != nil {
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

// ParseInstantWithdrawWhitelistUpdated is a log parse operation binding the contract event 0xea453964d1ed5c30fcc4e0065c0703205b656d4ed8df6badadd5641c880765a5.
//
// Solidity: event InstantWithdrawWhitelistUpdated(address indexed account, bool status)
func (_Svusd *SvusdFilterer) ParseInstantWithdrawWhitelistUpdated(log types.Log) (*SvusdInstantWithdrawWhitelistUpdated, error) {
	event := new(SvusdInstantWithdrawWhitelistUpdated)
	if err := _Svusd.contract.UnpackLog(event, "InstantWithdrawWhitelistUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SvusdOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the Svusd contract.
type SvusdOwnershipTransferStartedIterator struct {
	Event *SvusdOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *SvusdOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SvusdOwnershipTransferStarted)
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
		it.Event = new(SvusdOwnershipTransferStarted)
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
func (it *SvusdOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SvusdOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SvusdOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the Svusd contract.
type SvusdOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Svusd *SvusdFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SvusdOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Svusd.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SvusdOwnershipTransferStartedIterator{contract: _Svusd.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Svusd *SvusdFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *SvusdOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Svusd.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SvusdOwnershipTransferStarted)
				if err := _Svusd.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Svusd *SvusdFilterer) ParseOwnershipTransferStarted(log types.Log) (*SvusdOwnershipTransferStarted, error) {
	event := new(SvusdOwnershipTransferStarted)
	if err := _Svusd.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SvusdOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Svusd contract.
type SvusdOwnershipTransferredIterator struct {
	Event *SvusdOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SvusdOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SvusdOwnershipTransferred)
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
		it.Event = new(SvusdOwnershipTransferred)
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
func (it *SvusdOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SvusdOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SvusdOwnershipTransferred represents a OwnershipTransferred event raised by the Svusd contract.
type SvusdOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Svusd *SvusdFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SvusdOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Svusd.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SvusdOwnershipTransferredIterator{contract: _Svusd.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Svusd *SvusdFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SvusdOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Svusd.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SvusdOwnershipTransferred)
				if err := _Svusd.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Svusd *SvusdFilterer) ParseOwnershipTransferred(log types.Log) (*SvusdOwnershipTransferred, error) {
	event := new(SvusdOwnershipTransferred)
	if err := _Svusd.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SvusdTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Svusd contract.
type SvusdTransferIterator struct {
	Event *SvusdTransfer // Event containing the contract specifics and raw log

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
func (it *SvusdTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SvusdTransfer)
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
		it.Event = new(SvusdTransfer)
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
func (it *SvusdTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SvusdTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SvusdTransfer represents a Transfer event raised by the Svusd contract.
type SvusdTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Svusd *SvusdFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SvusdTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Svusd.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SvusdTransferIterator{contract: _Svusd.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Svusd *SvusdFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SvusdTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Svusd.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SvusdTransfer)
				if err := _Svusd.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Svusd *SvusdFilterer) ParseTransfer(log types.Log) (*SvusdTransfer, error) {
	event := new(SvusdTransfer)
	if err := _Svusd.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SvusdVaultRewardsUpdatedIterator is returned from FilterVaultRewardsUpdated and is used to iterate over the raw logs and unpacked data for VaultRewardsUpdated events raised by the Svusd contract.
type SvusdVaultRewardsUpdatedIterator struct {
	Event *SvusdVaultRewardsUpdated // Event containing the contract specifics and raw log

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
func (it *SvusdVaultRewardsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SvusdVaultRewardsUpdated)
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
		it.Event = new(SvusdVaultRewardsUpdated)
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
func (it *SvusdVaultRewardsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SvusdVaultRewardsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SvusdVaultRewardsUpdated represents a VaultRewardsUpdated event raised by the Svusd contract.
type SvusdVaultRewardsUpdated struct {
	PreviousVaultRewards common.Address
	NewVaultRewards      common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterVaultRewardsUpdated is a free log retrieval operation binding the contract event 0xd17fd6c2dd730d188323e59f585fc05d728ef4a029db3cffede6779d2c6fb7be.
//
// Solidity: event VaultRewardsUpdated(address indexed previousVaultRewards, address indexed newVaultRewards)
func (_Svusd *SvusdFilterer) FilterVaultRewardsUpdated(opts *bind.FilterOpts, previousVaultRewards []common.Address, newVaultRewards []common.Address) (*SvusdVaultRewardsUpdatedIterator, error) {

	var previousVaultRewardsRule []interface{}
	for _, previousVaultRewardsItem := range previousVaultRewards {
		previousVaultRewardsRule = append(previousVaultRewardsRule, previousVaultRewardsItem)
	}
	var newVaultRewardsRule []interface{}
	for _, newVaultRewardsItem := range newVaultRewards {
		newVaultRewardsRule = append(newVaultRewardsRule, newVaultRewardsItem)
	}

	logs, sub, err := _Svusd.contract.FilterLogs(opts, "VaultRewardsUpdated", previousVaultRewardsRule, newVaultRewardsRule)
	if err != nil {
		return nil, err
	}
	return &SvusdVaultRewardsUpdatedIterator{contract: _Svusd.contract, event: "VaultRewardsUpdated", logs: logs, sub: sub}, nil
}

// WatchVaultRewardsUpdated is a free log subscription operation binding the contract event 0xd17fd6c2dd730d188323e59f585fc05d728ef4a029db3cffede6779d2c6fb7be.
//
// Solidity: event VaultRewardsUpdated(address indexed previousVaultRewards, address indexed newVaultRewards)
func (_Svusd *SvusdFilterer) WatchVaultRewardsUpdated(opts *bind.WatchOpts, sink chan<- *SvusdVaultRewardsUpdated, previousVaultRewards []common.Address, newVaultRewards []common.Address) (event.Subscription, error) {

	var previousVaultRewardsRule []interface{}
	for _, previousVaultRewardsItem := range previousVaultRewards {
		previousVaultRewardsRule = append(previousVaultRewardsRule, previousVaultRewardsItem)
	}
	var newVaultRewardsRule []interface{}
	for _, newVaultRewardsItem := range newVaultRewards {
		newVaultRewardsRule = append(newVaultRewardsRule, newVaultRewardsItem)
	}

	logs, sub, err := _Svusd.contract.WatchLogs(opts, "VaultRewardsUpdated", previousVaultRewardsRule, newVaultRewardsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SvusdVaultRewardsUpdated)
				if err := _Svusd.contract.UnpackLog(event, "VaultRewardsUpdated", log); err != nil {
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

// ParseVaultRewardsUpdated is a log parse operation binding the contract event 0xd17fd6c2dd730d188323e59f585fc05d728ef4a029db3cffede6779d2c6fb7be.
//
// Solidity: event VaultRewardsUpdated(address indexed previousVaultRewards, address indexed newVaultRewards)
func (_Svusd *SvusdFilterer) ParseVaultRewardsUpdated(log types.Log) (*SvusdVaultRewardsUpdated, error) {
	event := new(SvusdVaultRewardsUpdated)
	if err := _Svusd.contract.UnpackLog(event, "VaultRewardsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SvusdWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Svusd contract.
type SvusdWithdrawIterator struct {
	Event *SvusdWithdraw // Event containing the contract specifics and raw log

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
func (it *SvusdWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SvusdWithdraw)
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
		it.Event = new(SvusdWithdraw)
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
func (it *SvusdWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SvusdWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SvusdWithdraw represents a Withdraw event raised by the Svusd contract.
type SvusdWithdraw struct {
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
func (_Svusd *SvusdFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, owner []common.Address) (*SvusdWithdrawIterator, error) {

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

	logs, sub, err := _Svusd.contract.FilterLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &SvusdWithdrawIterator{contract: _Svusd.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_Svusd *SvusdFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *SvusdWithdraw, sender []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Svusd.contract.WatchLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SvusdWithdraw)
				if err := _Svusd.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_Svusd *SvusdFilterer) ParseWithdraw(log types.Log) (*SvusdWithdraw, error) {
	event := new(SvusdWithdraw)
	if err := _Svusd.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SvusdWithdrawCancelledIterator is returned from FilterWithdrawCancelled and is used to iterate over the raw logs and unpacked data for WithdrawCancelled events raised by the Svusd contract.
type SvusdWithdrawCancelledIterator struct {
	Event *SvusdWithdrawCancelled // Event containing the contract specifics and raw log

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
func (it *SvusdWithdrawCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SvusdWithdrawCancelled)
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
		it.Event = new(SvusdWithdrawCancelled)
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
func (it *SvusdWithdrawCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SvusdWithdrawCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SvusdWithdrawCancelled represents a WithdrawCancelled event raised by the Svusd contract.
type SvusdWithdrawCancelled struct {
	Owner     common.Address
	RequestId *big.Int
	Assets    *big.Int
	Shares    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawCancelled is a free log retrieval operation binding the contract event 0x8a6803caa183ccb5fc1996809a7790f6b60d61f2bc403317748b627fd9937dc7.
//
// Solidity: event WithdrawCancelled(address indexed owner, uint256 indexed requestId, uint256 assets, uint256 shares)
func (_Svusd *SvusdFilterer) FilterWithdrawCancelled(opts *bind.FilterOpts, owner []common.Address, requestId []*big.Int) (*SvusdWithdrawCancelledIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Svusd.contract.FilterLogs(opts, "WithdrawCancelled", ownerRule, requestIdRule)
	if err != nil {
		return nil, err
	}
	return &SvusdWithdrawCancelledIterator{contract: _Svusd.contract, event: "WithdrawCancelled", logs: logs, sub: sub}, nil
}

// WatchWithdrawCancelled is a free log subscription operation binding the contract event 0x8a6803caa183ccb5fc1996809a7790f6b60d61f2bc403317748b627fd9937dc7.
//
// Solidity: event WithdrawCancelled(address indexed owner, uint256 indexed requestId, uint256 assets, uint256 shares)
func (_Svusd *SvusdFilterer) WatchWithdrawCancelled(opts *bind.WatchOpts, sink chan<- *SvusdWithdrawCancelled, owner []common.Address, requestId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Svusd.contract.WatchLogs(opts, "WithdrawCancelled", ownerRule, requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SvusdWithdrawCancelled)
				if err := _Svusd.contract.UnpackLog(event, "WithdrawCancelled", log); err != nil {
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

// ParseWithdrawCancelled is a log parse operation binding the contract event 0x8a6803caa183ccb5fc1996809a7790f6b60d61f2bc403317748b627fd9937dc7.
//
// Solidity: event WithdrawCancelled(address indexed owner, uint256 indexed requestId, uint256 assets, uint256 shares)
func (_Svusd *SvusdFilterer) ParseWithdrawCancelled(log types.Log) (*SvusdWithdrawCancelled, error) {
	event := new(SvusdWithdrawCancelled)
	if err := _Svusd.contract.UnpackLog(event, "WithdrawCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SvusdWithdrawClaimedIterator is returned from FilterWithdrawClaimed and is used to iterate over the raw logs and unpacked data for WithdrawClaimed events raised by the Svusd contract.
type SvusdWithdrawClaimedIterator struct {
	Event *SvusdWithdrawClaimed // Event containing the contract specifics and raw log

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
func (it *SvusdWithdrawClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SvusdWithdrawClaimed)
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
		it.Event = new(SvusdWithdrawClaimed)
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
func (it *SvusdWithdrawClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SvusdWithdrawClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SvusdWithdrawClaimed represents a WithdrawClaimed event raised by the Svusd contract.
type SvusdWithdrawClaimed struct {
	Owner     common.Address
	Receiver  common.Address
	RequestId *big.Int
	Assets    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawClaimed is a free log retrieval operation binding the contract event 0xb4317ae2386e712514c7723055d29e3f5b96ebd3f972f1ef3bbcdcf4a8ce40c0.
//
// Solidity: event WithdrawClaimed(address indexed owner, address indexed receiver, uint256 indexed requestId, uint256 assets)
func (_Svusd *SvusdFilterer) FilterWithdrawClaimed(opts *bind.FilterOpts, owner []common.Address, receiver []common.Address, requestId []*big.Int) (*SvusdWithdrawClaimedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Svusd.contract.FilterLogs(opts, "WithdrawClaimed", ownerRule, receiverRule, requestIdRule)
	if err != nil {
		return nil, err
	}
	return &SvusdWithdrawClaimedIterator{contract: _Svusd.contract, event: "WithdrawClaimed", logs: logs, sub: sub}, nil
}

// WatchWithdrawClaimed is a free log subscription operation binding the contract event 0xb4317ae2386e712514c7723055d29e3f5b96ebd3f972f1ef3bbcdcf4a8ce40c0.
//
// Solidity: event WithdrawClaimed(address indexed owner, address indexed receiver, uint256 indexed requestId, uint256 assets)
func (_Svusd *SvusdFilterer) WatchWithdrawClaimed(opts *bind.WatchOpts, sink chan<- *SvusdWithdrawClaimed, owner []common.Address, receiver []common.Address, requestId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Svusd.contract.WatchLogs(opts, "WithdrawClaimed", ownerRule, receiverRule, requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SvusdWithdrawClaimed)
				if err := _Svusd.contract.UnpackLog(event, "WithdrawClaimed", log); err != nil {
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

// ParseWithdrawClaimed is a log parse operation binding the contract event 0xb4317ae2386e712514c7723055d29e3f5b96ebd3f972f1ef3bbcdcf4a8ce40c0.
//
// Solidity: event WithdrawClaimed(address indexed owner, address indexed receiver, uint256 indexed requestId, uint256 assets)
func (_Svusd *SvusdFilterer) ParseWithdrawClaimed(log types.Log) (*SvusdWithdrawClaimed, error) {
	event := new(SvusdWithdrawClaimed)
	if err := _Svusd.contract.UnpackLog(event, "WithdrawClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SvusdWithdrawRequestedIterator is returned from FilterWithdrawRequested and is used to iterate over the raw logs and unpacked data for WithdrawRequested events raised by the Svusd contract.
type SvusdWithdrawRequestedIterator struct {
	Event *SvusdWithdrawRequested // Event containing the contract specifics and raw log

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
func (it *SvusdWithdrawRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SvusdWithdrawRequested)
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
		it.Event = new(SvusdWithdrawRequested)
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
func (it *SvusdWithdrawRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SvusdWithdrawRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SvusdWithdrawRequested represents a WithdrawRequested event raised by the Svusd contract.
type SvusdWithdrawRequested struct {
	Owner       common.Address
	RequestId   *big.Int
	Shares      *big.Int
	Assets      *big.Int
	ClaimableAt *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawRequested is a free log retrieval operation binding the contract event 0x1f0d204a59639d21245dbaa239dfc9a6227aa86cb614bdda19843c52d155fcfa.
//
// Solidity: event WithdrawRequested(address indexed owner, uint256 indexed requestId, uint256 shares, uint256 assets, uint256 claimableAt)
func (_Svusd *SvusdFilterer) FilterWithdrawRequested(opts *bind.FilterOpts, owner []common.Address, requestId []*big.Int) (*SvusdWithdrawRequestedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Svusd.contract.FilterLogs(opts, "WithdrawRequested", ownerRule, requestIdRule)
	if err != nil {
		return nil, err
	}
	return &SvusdWithdrawRequestedIterator{contract: _Svusd.contract, event: "WithdrawRequested", logs: logs, sub: sub}, nil
}

// WatchWithdrawRequested is a free log subscription operation binding the contract event 0x1f0d204a59639d21245dbaa239dfc9a6227aa86cb614bdda19843c52d155fcfa.
//
// Solidity: event WithdrawRequested(address indexed owner, uint256 indexed requestId, uint256 shares, uint256 assets, uint256 claimableAt)
func (_Svusd *SvusdFilterer) WatchWithdrawRequested(opts *bind.WatchOpts, sink chan<- *SvusdWithdrawRequested, owner []common.Address, requestId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Svusd.contract.WatchLogs(opts, "WithdrawRequested", ownerRule, requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SvusdWithdrawRequested)
				if err := _Svusd.contract.UnpackLog(event, "WithdrawRequested", log); err != nil {
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

// ParseWithdrawRequested is a log parse operation binding the contract event 0x1f0d204a59639d21245dbaa239dfc9a6227aa86cb614bdda19843c52d155fcfa.
//
// Solidity: event WithdrawRequested(address indexed owner, uint256 indexed requestId, uint256 shares, uint256 assets, uint256 claimableAt)
func (_Svusd *SvusdFilterer) ParseWithdrawRequested(log types.Log) (*SvusdWithdrawRequested, error) {
	event := new(SvusdWithdrawRequested)
	if err := _Svusd.contract.UnpackLog(event, "WithdrawRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SvusdYieldDistributorUpdatedIterator is returned from FilterYieldDistributorUpdated and is used to iterate over the raw logs and unpacked data for YieldDistributorUpdated events raised by the Svusd contract.
type SvusdYieldDistributorUpdatedIterator struct {
	Event *SvusdYieldDistributorUpdated // Event containing the contract specifics and raw log

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
func (it *SvusdYieldDistributorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SvusdYieldDistributorUpdated)
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
		it.Event = new(SvusdYieldDistributorUpdated)
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
func (it *SvusdYieldDistributorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SvusdYieldDistributorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SvusdYieldDistributorUpdated represents a YieldDistributorUpdated event raised by the Svusd contract.
type SvusdYieldDistributorUpdated struct {
	PreviousDistributor common.Address
	NewDistributor      common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterYieldDistributorUpdated is a free log retrieval operation binding the contract event 0xeb35f0f0a95d2801f2d360efbbf9ee54c1eb3a1801fa152bc4d4cf9ee10a0890.
//
// Solidity: event YieldDistributorUpdated(address indexed previousDistributor, address indexed newDistributor)
func (_Svusd *SvusdFilterer) FilterYieldDistributorUpdated(opts *bind.FilterOpts, previousDistributor []common.Address, newDistributor []common.Address) (*SvusdYieldDistributorUpdatedIterator, error) {

	var previousDistributorRule []interface{}
	for _, previousDistributorItem := range previousDistributor {
		previousDistributorRule = append(previousDistributorRule, previousDistributorItem)
	}
	var newDistributorRule []interface{}
	for _, newDistributorItem := range newDistributor {
		newDistributorRule = append(newDistributorRule, newDistributorItem)
	}

	logs, sub, err := _Svusd.contract.FilterLogs(opts, "YieldDistributorUpdated", previousDistributorRule, newDistributorRule)
	if err != nil {
		return nil, err
	}
	return &SvusdYieldDistributorUpdatedIterator{contract: _Svusd.contract, event: "YieldDistributorUpdated", logs: logs, sub: sub}, nil
}

// WatchYieldDistributorUpdated is a free log subscription operation binding the contract event 0xeb35f0f0a95d2801f2d360efbbf9ee54c1eb3a1801fa152bc4d4cf9ee10a0890.
//
// Solidity: event YieldDistributorUpdated(address indexed previousDistributor, address indexed newDistributor)
func (_Svusd *SvusdFilterer) WatchYieldDistributorUpdated(opts *bind.WatchOpts, sink chan<- *SvusdYieldDistributorUpdated, previousDistributor []common.Address, newDistributor []common.Address) (event.Subscription, error) {

	var previousDistributorRule []interface{}
	for _, previousDistributorItem := range previousDistributor {
		previousDistributorRule = append(previousDistributorRule, previousDistributorItem)
	}
	var newDistributorRule []interface{}
	for _, newDistributorItem := range newDistributor {
		newDistributorRule = append(newDistributorRule, newDistributorItem)
	}

	logs, sub, err := _Svusd.contract.WatchLogs(opts, "YieldDistributorUpdated", previousDistributorRule, newDistributorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SvusdYieldDistributorUpdated)
				if err := _Svusd.contract.UnpackLog(event, "YieldDistributorUpdated", log); err != nil {
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

// ParseYieldDistributorUpdated is a log parse operation binding the contract event 0xeb35f0f0a95d2801f2d360efbbf9ee54c1eb3a1801fa152bc4d4cf9ee10a0890.
//
// Solidity: event YieldDistributorUpdated(address indexed previousDistributor, address indexed newDistributor)
func (_Svusd *SvusdFilterer) ParseYieldDistributorUpdated(log types.Log) (*SvusdYieldDistributorUpdated, error) {
	event := new(SvusdYieldDistributorUpdated)
	if err := _Svusd.contract.UnpackLog(event, "YieldDistributorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
