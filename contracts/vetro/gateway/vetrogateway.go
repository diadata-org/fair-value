// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vetrogateway

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

// VetrogatewayMetaData contains all meta data concerning the Vetrogateway contract.
var VetrogatewayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AccountAlreadyWhitelisted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AccountNotWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AddressIsZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"AmoBurnExceedsSupply\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountIsZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerNotWhitelisted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"ExceededMaxMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"ExceededMaxWithdraw\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"FeeOnTransferToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"constraint\",\"type\":\"uint256\"}],\"name\":\"InvalidAmoMintLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"InvalidMintFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pegBand\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPegBandBps\",\"type\":\"uint256\"}],\"name\":\"InvalidPegBand\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"InvalidRedeemFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWithdrawalDelay\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"peggedTokenOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPeggedTokenOut\",\"type\":\"uint256\"}],\"name\":\"MintableIsLessThanMinimum\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoActiveWithdrawalRequest\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"peggedTokenIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPeggedTokenIn\",\"type\":\"uint256\"}],\"name\":\"PeggedTokenToBurnIsHigherThanMax\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTokenOut\",\"type\":\"uint256\"}],\"name\":\"RedeemableIsLessThanMinimum\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokenIn\",\"type\":\"uint256\"}],\"name\":\"TokenAmountIsHigherThanMax\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenNotWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawalDelayFeatureNotEnabled\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddedToInstantRedeemWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountBurned\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAmoSupply\",\"type\":\"uint256\"}],\"name\":\"BurnedFromAMO\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"peggedTokenAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousMintFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMintFee\",\"type\":\"uint256\"}],\"name\":\"MintFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousMintLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMintLimit\",\"type\":\"uint256\"}],\"name\":\"MintLimitUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountMinted\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAmoSupply\",\"type\":\"uint256\"}],\"name\":\"MintedToAMO\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousPegBandBps\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPegBandBps\",\"type\":\"uint256\"}],\"name\":\"PegBandUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousRedeemFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRedeemFee\",\"type\":\"uint256\"}],\"name\":\"RedeemFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RedeemRequestCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimableAt\",\"type\":\"uint256\"}],\"name\":\"RedeemRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RemovedFromInstantRedeemWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"UpdatedAmoMintLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"peggedTokenAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"WithdrawalDelayEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousDelay\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDelay\",\"type\":\"uint256\"}],\"name\":\"WithdrawalDelayUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAINTAINER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_BPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_FEE_BPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_WITHDRAWAL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PEGGED_TOKEN\",\"outputs\":[{\"internalType\":\"contractIPeggedToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UMM_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"addToInstantRedeemWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"amoMintLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"amoSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"burnFromAMO\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelRedeemRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPeggedTokenOut_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInstantRedeemWhitelist\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_\",\"type\":\"address\"}],\"name\":\"getRedeemRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountLocked\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_claimableAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"peggedToken_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintLimit_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialWithdrawalDelay_\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"isInstantRedeemWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxAmoMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOut_\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"peggedTokenOut_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAmountIn_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"}],\"name\":\"mintFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"}],\"name\":\"mintToAMO\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"}],\"name\":\"pegBand\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn_\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"peggedTokenOut_\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOut_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"peggedTokenIn_\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOut_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut_\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOut_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"peggedTokenIn_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"}],\"name\":\"redeemFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"removeFromInstantRedeemWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"peggedTokenAmount_\",\"type\":\"uint256\"}],\"name\":\"requestRedeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled_\",\"type\":\"bool\"}],\"name\":\"setWithdrawalDelayEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newAmoMintLimit_\",\"type\":\"uint256\"}],\"name\":\"updateAmoMintLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newMintFee_\",\"type\":\"uint256\"}],\"name\":\"updateMintFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMintLimit_\",\"type\":\"uint256\"}],\"name\":\"updateMintLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newPegBandBps_\",\"type\":\"uint256\"}],\"name\":\"updatePegBand\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newRedeemFee_\",\"type\":\"uint256\"}],\"name\":\"updateRedeemFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDelay_\",\"type\":\"uint256\"}],\"name\":\"updateWithdrawalDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOut_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPeggedTokenIn_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalDelayEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VetrogatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use VetrogatewayMetaData.ABI instead.
var VetrogatewayABI = VetrogatewayMetaData.ABI

// Vetrogateway is an auto generated Go binding around an Ethereum contract.
type Vetrogateway struct {
	VetrogatewayCaller     // Read-only binding to the contract
	VetrogatewayTransactor // Write-only binding to the contract
	VetrogatewayFilterer   // Log filterer for contract events
}

// VetrogatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type VetrogatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VetrogatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VetrogatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VetrogatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VetrogatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VetrogatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VetrogatewaySession struct {
	Contract     *Vetrogateway     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VetrogatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VetrogatewayCallerSession struct {
	Contract *VetrogatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// VetrogatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VetrogatewayTransactorSession struct {
	Contract     *VetrogatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// VetrogatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type VetrogatewayRaw struct {
	Contract *Vetrogateway // Generic contract binding to access the raw methods on
}

// VetrogatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VetrogatewayCallerRaw struct {
	Contract *VetrogatewayCaller // Generic read-only contract binding to access the raw methods on
}

// VetrogatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VetrogatewayTransactorRaw struct {
	Contract *VetrogatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVetrogateway creates a new instance of Vetrogateway, bound to a specific deployed contract.
func NewVetrogateway(address common.Address, backend bind.ContractBackend) (*Vetrogateway, error) {
	contract, err := bindVetrogateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vetrogateway{VetrogatewayCaller: VetrogatewayCaller{contract: contract}, VetrogatewayTransactor: VetrogatewayTransactor{contract: contract}, VetrogatewayFilterer: VetrogatewayFilterer{contract: contract}}, nil
}

// NewVetrogatewayCaller creates a new read-only instance of Vetrogateway, bound to a specific deployed contract.
func NewVetrogatewayCaller(address common.Address, caller bind.ContractCaller) (*VetrogatewayCaller, error) {
	contract, err := bindVetrogateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VetrogatewayCaller{contract: contract}, nil
}

// NewVetrogatewayTransactor creates a new write-only instance of Vetrogateway, bound to a specific deployed contract.
func NewVetrogatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*VetrogatewayTransactor, error) {
	contract, err := bindVetrogateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VetrogatewayTransactor{contract: contract}, nil
}

// NewVetrogatewayFilterer creates a new log filterer instance of Vetrogateway, bound to a specific deployed contract.
func NewVetrogatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*VetrogatewayFilterer, error) {
	contract, err := bindVetrogateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VetrogatewayFilterer{contract: contract}, nil
}

// bindVetrogateway binds a generic wrapper to an already deployed contract.
func bindVetrogateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VetrogatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vetrogateway *VetrogatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vetrogateway.Contract.VetrogatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vetrogateway *VetrogatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vetrogateway.Contract.VetrogatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vetrogateway *VetrogatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vetrogateway.Contract.VetrogatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vetrogateway *VetrogatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vetrogateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vetrogateway *VetrogatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vetrogateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vetrogateway *VetrogatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vetrogateway.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Vetrogateway *VetrogatewayCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vetrogateway.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Vetrogateway *VetrogatewaySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Vetrogateway.Contract.DEFAULTADMINROLE(&_Vetrogateway.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Vetrogateway *VetrogatewayCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Vetrogateway.Contract.DEFAULTADMINROLE(&_Vetrogateway.CallOpts)
}

// MAINTAINERROLE is a free data retrieval call binding the contract method 0xf8742254.
//
// Solidity: function MAINTAINER_ROLE() view returns(bytes32)
func (_Vetrogateway *VetrogatewayCaller) MAINTAINERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vetrogateway.contract.Call(opts, &out, "MAINTAINER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MAINTAINERROLE is a free data retrieval call binding the contract method 0xf8742254.
//
// Solidity: function MAINTAINER_ROLE() view returns(bytes32)
func (_Vetrogateway *VetrogatewaySession) MAINTAINERROLE() ([32]byte, error) {
	return _Vetrogateway.Contract.MAINTAINERROLE(&_Vetrogateway.CallOpts)
}

// MAINTAINERROLE is a free data retrieval call binding the contract method 0xf8742254.
//
// Solidity: function MAINTAINER_ROLE() view returns(bytes32)
func (_Vetrogateway *VetrogatewayCallerSession) MAINTAINERROLE() ([32]byte, error) {
	return _Vetrogateway.Contract.MAINTAINERROLE(&_Vetrogateway.CallOpts)
}

// MAXBPS is a free data retrieval call binding the contract method 0xfd967f47.
//
// Solidity: function MAX_BPS() view returns(uint256)
func (_Vetrogateway *VetrogatewayCaller) MAXBPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vetrogateway.contract.Call(opts, &out, "MAX_BPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBPS is a free data retrieval call binding the contract method 0xfd967f47.
//
// Solidity: function MAX_BPS() view returns(uint256)
func (_Vetrogateway *VetrogatewaySession) MAXBPS() (*big.Int, error) {
	return _Vetrogateway.Contract.MAXBPS(&_Vetrogateway.CallOpts)
}

// MAXBPS is a free data retrieval call binding the contract method 0xfd967f47.
//
// Solidity: function MAX_BPS() view returns(uint256)
func (_Vetrogateway *VetrogatewayCallerSession) MAXBPS() (*big.Int, error) {
	return _Vetrogateway.Contract.MAXBPS(&_Vetrogateway.CallOpts)
}

// MAXFEEBPS is a free data retrieval call binding the contract method 0xd55be8c6.
//
// Solidity: function MAX_FEE_BPS() view returns(uint256)
func (_Vetrogateway *VetrogatewayCaller) MAXFEEBPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vetrogateway.contract.Call(opts, &out, "MAX_FEE_BPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFEEBPS is a free data retrieval call binding the contract method 0xd55be8c6.
//
// Solidity: function MAX_FEE_BPS() view returns(uint256)
func (_Vetrogateway *VetrogatewaySession) MAXFEEBPS() (*big.Int, error) {
	return _Vetrogateway.Contract.MAXFEEBPS(&_Vetrogateway.CallOpts)
}

// MAXFEEBPS is a free data retrieval call binding the contract method 0xd55be8c6.
//
// Solidity: function MAX_FEE_BPS() view returns(uint256)
func (_Vetrogateway *VetrogatewayCallerSession) MAXFEEBPS() (*big.Int, error) {
	return _Vetrogateway.Contract.MAXFEEBPS(&_Vetrogateway.CallOpts)
}

// MAXWITHDRAWALDELAY is a free data retrieval call binding the contract method 0xa238f9df.
//
// Solidity: function MAX_WITHDRAWAL_DELAY() view returns(uint256)
func (_Vetrogateway *VetrogatewayCaller) MAXWITHDRAWALDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vetrogateway.contract.Call(opts, &out, "MAX_WITHDRAWAL_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXWITHDRAWALDELAY is a free data retrieval call binding the contract method 0xa238f9df.
//
// Solidity: function MAX_WITHDRAWAL_DELAY() view returns(uint256)
func (_Vetrogateway *VetrogatewaySession) MAXWITHDRAWALDELAY() (*big.Int, error) {
	return _Vetrogateway.Contract.MAXWITHDRAWALDELAY(&_Vetrogateway.CallOpts)
}

// MAXWITHDRAWALDELAY is a free data retrieval call binding the contract method 0xa238f9df.
//
// Solidity: function MAX_WITHDRAWAL_DELAY() view returns(uint256)
func (_Vetrogateway *VetrogatewayCallerSession) MAXWITHDRAWALDELAY() (*big.Int, error) {
	return _Vetrogateway.Contract.MAXWITHDRAWALDELAY(&_Vetrogateway.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_Vetrogateway *VetrogatewayCaller) NAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Vetrogateway.contract.Call(opts, &out, "NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_Vetrogateway *VetrogatewaySession) NAME() (string, error) {
	return _Vetrogateway.Contract.NAME(&_Vetrogateway.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_Vetrogateway *VetrogatewayCallerSession) NAME() (string, error) {
	return _Vetrogateway.Contract.NAME(&_Vetrogateway.CallOpts)
}

// PEGGEDTOKEN is a free data retrieval call binding the contract method 0x4e485cf7.
//
// Solidity: function PEGGED_TOKEN() view returns(address)
func (_Vetrogateway *VetrogatewayCaller) PEGGEDTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vetrogateway.contract.Call(opts, &out, "PEGGED_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PEGGEDTOKEN is a free data retrieval call binding the contract method 0x4e485cf7.
//
// Solidity: function PEGGED_TOKEN() view returns(address)
func (_Vetrogateway *VetrogatewaySession) PEGGEDTOKEN() (common.Address, error) {
	return _Vetrogateway.Contract.PEGGEDTOKEN(&_Vetrogateway.CallOpts)
}

// PEGGEDTOKEN is a free data retrieval call binding the contract method 0x4e485cf7.
//
// Solidity: function PEGGED_TOKEN() view returns(address)
func (_Vetrogateway *VetrogatewayCallerSession) PEGGEDTOKEN() (common.Address, error) {
	return _Vetrogateway.Contract.PEGGEDTOKEN(&_Vetrogateway.CallOpts)
}

// UMMROLE is a free data retrieval call binding the contract method 0xf69516bd.
//
// Solidity: function UMM_ROLE() view returns(bytes32)
func (_Vetrogateway *VetrogatewayCaller) UMMROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vetrogateway.contract.Call(opts, &out, "UMM_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UMMROLE is a free data retrieval call binding the contract method 0xf69516bd.
//
// Solidity: function UMM_ROLE() view returns(bytes32)
func (_Vetrogateway *VetrogatewaySession) UMMROLE() ([32]byte, error) {
	return _Vetrogateway.Contract.UMMROLE(&_Vetrogateway.CallOpts)
}

// UMMROLE is a free data retrieval call binding the contract method 0xf69516bd.
//
// Solidity: function UMM_ROLE() view returns(bytes32)
func (_Vetrogateway *VetrogatewayCallerSession) UMMROLE() ([32]byte, error) {
	return _Vetrogateway.Contract.UMMROLE(&_Vetrogateway.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Vetrogateway *VetrogatewayCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Vetrogateway.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Vetrogateway *VetrogatewaySession) VERSION() (string, error) {
	return _Vetrogateway.Contract.VERSION(&_Vetrogateway.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Vetrogateway *VetrogatewayCallerSession) VERSION() (string, error) {
	return _Vetrogateway.Contract.VERSION(&_Vetrogateway.CallOpts)
}

// AmoMintLimit is a free data retrieval call binding the contract method 0x1d0f9254.
//
// Solidity: function amoMintLimit() view returns(uint256)
func (_Vetrogateway *VetrogatewayCaller) AmoMintLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vetrogateway.contract.Call(opts, &out, "amoMintLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmoMintLimit is a free data retrieval call binding the contract method 0x1d0f9254.
//
// Solidity: function amoMintLimit() view returns(uint256)
func (_Vetrogateway *VetrogatewaySession) AmoMintLimit() (*big.Int, error) {
	return _Vetrogateway.Contract.AmoMintLimit(&_Vetrogateway.CallOpts)
}

// AmoMintLimit is a free data retrieval call binding the contract method 0x1d0f9254.
//
// Solidity: function amoMintLimit() view returns(uint256)
func (_Vetrogateway *VetrogatewayCallerSession) AmoMintLimit() (*big.Int, error) {
	return _Vetrogateway.Contract.AmoMintLimit(&_Vetrogateway.CallOpts)
}

// AmoSupply is a free data retrieval call binding the contract method 0x1543b996.
//
// Solidity: function amoSupply() view returns(uint256)
func (_Vetrogateway *VetrogatewayCaller) AmoSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vetrogateway.contract.Call(opts, &out, "amoSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmoSupply is a free data retrieval call binding the contract method 0x1543b996.
//
// Solidity: function amoSupply() view returns(uint256)
func (_Vetrogateway *VetrogatewaySession) AmoSupply() (*big.Int, error) {
	return _Vetrogateway.Contract.AmoSupply(&_Vetrogateway.CallOpts)
}

// AmoSupply is a free data retrieval call binding the contract method 0x1543b996.
//
// Solidity: function amoSupply() view returns(uint256)
func (_Vetrogateway *VetrogatewayCallerSession) AmoSupply() (*big.Int, error) {
	return _Vetrogateway.Contract.AmoSupply(&_Vetrogateway.CallOpts)
}

// GetInstantRedeemWhitelist is a free data retrieval call binding the contract method 0x11dbcb6a.
//
// Solidity: function getInstantRedeemWhitelist() view returns(address[])
func (_Vetrogateway *VetrogatewayCaller) GetInstantRedeemWhitelist(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Vetrogateway.contract.Call(opts, &out, "getInstantRedeemWhitelist")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetInstantRedeemWhitelist is a free data retrieval call binding the contract method 0x11dbcb6a.
//
// Solidity: function getInstantRedeemWhitelist() view returns(address[])
func (_Vetrogateway *VetrogatewaySession) GetInstantRedeemWhitelist() ([]common.Address, error) {
	return _Vetrogateway.Contract.GetInstantRedeemWhitelist(&_Vetrogateway.CallOpts)
}

// GetInstantRedeemWhitelist is a free data retrieval call binding the contract method 0x11dbcb6a.
//
// Solidity: function getInstantRedeemWhitelist() view returns(address[])
func (_Vetrogateway *VetrogatewayCallerSession) GetInstantRedeemWhitelist() ([]common.Address, error) {
	return _Vetrogateway.Contract.GetInstantRedeemWhitelist(&_Vetrogateway.CallOpts)
}

// GetRedeemRequest is a free data retrieval call binding the contract method 0x3c05e18f.
//
// Solidity: function getRedeemRequest(address user_) view returns(uint256 _amountLocked, uint256 _claimableAt)
func (_Vetrogateway *VetrogatewayCaller) GetRedeemRequest(opts *bind.CallOpts, user_ common.Address) (struct {
	AmountLocked *big.Int
	ClaimableAt  *big.Int
}, error) {
	var out []interface{}
	err := _Vetrogateway.contract.Call(opts, &out, "getRedeemRequest", user_)

	outstruct := new(struct {
		AmountLocked *big.Int
		ClaimableAt  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountLocked = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ClaimableAt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetRedeemRequest is a free data retrieval call binding the contract method 0x3c05e18f.
//
// Solidity: function getRedeemRequest(address user_) view returns(uint256 _amountLocked, uint256 _claimableAt)
func (_Vetrogateway *VetrogatewaySession) GetRedeemRequest(user_ common.Address) (struct {
	AmountLocked *big.Int
	ClaimableAt  *big.Int
}, error) {
	return _Vetrogateway.Contract.GetRedeemRequest(&_Vetrogateway.CallOpts, user_)
}

// GetRedeemRequest is a free data retrieval call binding the contract method 0x3c05e18f.
//
// Solidity: function getRedeemRequest(address user_) view returns(uint256 _amountLocked, uint256 _claimableAt)
func (_Vetrogateway *VetrogatewayCallerSession) GetRedeemRequest(user_ common.Address) (struct {
	AmountLocked *big.Int
	ClaimableAt  *big.Int
}, error) {
	return _Vetrogateway.Contract.GetRedeemRequest(&_Vetrogateway.CallOpts, user_)
}

// IsInstantRedeemWhitelisted is a free data retrieval call binding the contract method 0x569e7665.
//
// Solidity: function isInstantRedeemWhitelisted(address account_) view returns(bool)
func (_Vetrogateway *VetrogatewayCaller) IsInstantRedeemWhitelisted(opts *bind.CallOpts, account_ common.Address) (bool, error) {
	var out []interface{}
	err := _Vetrogateway.contract.Call(opts, &out, "isInstantRedeemWhitelisted", account_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInstantRedeemWhitelisted is a free data retrieval call binding the contract method 0x569e7665.
//
// Solidity: function isInstantRedeemWhitelisted(address account_) view returns(bool)
func (_Vetrogateway *VetrogatewaySession) IsInstantRedeemWhitelisted(account_ common.Address) (bool, error) {
	return _Vetrogateway.Contract.IsInstantRedeemWhitelisted(&_Vetrogateway.CallOpts, account_)
}

// IsInstantRedeemWhitelisted is a free data retrieval call binding the contract method 0x569e7665.
//
// Solidity: function isInstantRedeemWhitelisted(address account_) view returns(bool)
func (_Vetrogateway *VetrogatewayCallerSession) IsInstantRedeemWhitelisted(account_ common.Address) (bool, error) {
	return _Vetrogateway.Contract.IsInstantRedeemWhitelisted(&_Vetrogateway.CallOpts, account_)
}

// MaxAmoMint is a free data retrieval call binding the contract method 0xe2ad4430.
//
// Solidity: function maxAmoMint() view returns(uint256)
func (_Vetrogateway *VetrogatewayCaller) MaxAmoMint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vetrogateway.contract.Call(opts, &out, "maxAmoMint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAmoMint is a free data retrieval call binding the contract method 0xe2ad4430.
//
// Solidity: function maxAmoMint() view returns(uint256)
func (_Vetrogateway *VetrogatewaySession) MaxAmoMint() (*big.Int, error) {
	return _Vetrogateway.Contract.MaxAmoMint(&_Vetrogateway.CallOpts)
}

// MaxAmoMint is a free data retrieval call binding the contract method 0xe2ad4430.
//
// Solidity: function maxAmoMint() view returns(uint256)
func (_Vetrogateway *VetrogatewayCallerSession) MaxAmoMint() (*big.Int, error) {
	return _Vetrogateway.Contract.MaxAmoMint(&_Vetrogateway.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x6083e59a.
//
// Solidity: function maxDeposit() pure returns(uint256)
func (_Vetrogateway *VetrogatewayCaller) MaxDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vetrogateway.contract.Call(opts, &out, "maxDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x6083e59a.
//
// Solidity: function maxDeposit() pure returns(uint256)
func (_Vetrogateway *VetrogatewaySession) MaxDeposit() (*big.Int, error) {
	return _Vetrogateway.Contract.MaxDeposit(&_Vetrogateway.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x6083e59a.
//
// Solidity: function maxDeposit() pure returns(uint256)
func (_Vetrogateway *VetrogatewayCallerSession) MaxDeposit() (*big.Int, error) {
	return _Vetrogateway.Contract.MaxDeposit(&_Vetrogateway.CallOpts)
}

// MaxMint is a free data retrieval call binding the contract method 0x7501f741.
//
// Solidity: function maxMint() view returns(uint256)
func (_Vetrogateway *VetrogatewayCaller) MaxMint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vetrogateway.contract.Call(opts, &out, "maxMint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0x7501f741.
//
// Solidity: function maxMint() view returns(uint256)
func (_Vetrogateway *VetrogatewaySession) MaxMint() (*big.Int, error) {
	return _Vetrogateway.Contract.MaxMint(&_Vetrogateway.CallOpts)
}

// MaxMint is a free data retrieval call binding the contract method 0x7501f741.
//
// Solidity: function maxMint() view returns(uint256)
func (_Vetrogateway *VetrogatewayCallerSession) MaxMint() (*big.Int, error) {
	return _Vetrogateway.Contract.MaxMint(&_Vetrogateway.CallOpts)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner_) view returns(uint256)
func (_Vetrogateway *VetrogatewayCaller) MaxRedeem(opts *bind.CallOpts, owner_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vetrogateway.contract.Call(opts, &out, "maxRedeem", owner_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner_) view returns(uint256)
func (_Vetrogateway *VetrogatewaySession) MaxRedeem(owner_ common.Address) (*big.Int, error) {
	return _Vetrogateway.Contract.MaxRedeem(&_Vetrogateway.CallOpts, owner_)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner_) view returns(uint256)
func (_Vetrogateway *VetrogatewayCallerSession) MaxRedeem(owner_ common.Address) (*big.Int, error) {
	return _Vetrogateway.Contract.MaxRedeem(&_Vetrogateway.CallOpts, owner_)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address tokenOut_) view returns(uint256)
func (_Vetrogateway *VetrogatewayCaller) MaxWithdraw(opts *bind.CallOpts, tokenOut_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vetrogateway.contract.Call(opts, &out, "maxWithdraw", tokenOut_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address tokenOut_) view returns(uint256)
func (_Vetrogateway *VetrogatewaySession) MaxWithdraw(tokenOut_ common.Address) (*big.Int, error) {
	return _Vetrogateway.Contract.MaxWithdraw(&_Vetrogateway.CallOpts, tokenOut_)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address tokenOut_) view returns(uint256)
func (_Vetrogateway *VetrogatewayCallerSession) MaxWithdraw(tokenOut_ common.Address) (*big.Int, error) {
	return _Vetrogateway.Contract.MaxWithdraw(&_Vetrogateway.CallOpts, tokenOut_)
}

// MintFee is a free data retrieval call binding the contract method 0xf8b49e72.
//
// Solidity: function mintFee(address token_) view returns(uint256)
func (_Vetrogateway *VetrogatewayCaller) MintFee(opts *bind.CallOpts, token_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vetrogateway.contract.Call(opts, &out, "mintFee", token_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintFee is a free data retrieval call binding the contract method 0xf8b49e72.
//
// Solidity: function mintFee(address token_) view returns(uint256)
func (_Vetrogateway *VetrogatewaySession) MintFee(token_ common.Address) (*big.Int, error) {
	return _Vetrogateway.Contract.MintFee(&_Vetrogateway.CallOpts, token_)
}

// MintFee is a free data retrieval call binding the contract method 0xf8b49e72.
//
// Solidity: function mintFee(address token_) view returns(uint256)
func (_Vetrogateway *VetrogatewayCallerSession) MintFee(token_ common.Address) (*big.Int, error) {
	return _Vetrogateway.Contract.MintFee(&_Vetrogateway.CallOpts, token_)
}

// MintLimit is a free data retrieval call binding the contract method 0x996517cf.
//
// Solidity: function mintLimit() view returns(uint256)
func (_Vetrogateway *VetrogatewayCaller) MintLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vetrogateway.contract.Call(opts, &out, "mintLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintLimit is a free data retrieval call binding the contract method 0x996517cf.
//
// Solidity: function mintLimit() view returns(uint256)
func (_Vetrogateway *VetrogatewaySession) MintLimit() (*big.Int, error) {
	return _Vetrogateway.Contract.MintLimit(&_Vetrogateway.CallOpts)
}

// MintLimit is a free data retrieval call binding the contract method 0x996517cf.
//
// Solidity: function mintLimit() view returns(uint256)
func (_Vetrogateway *VetrogatewayCallerSession) MintLimit() (*big.Int, error) {
	return _Vetrogateway.Contract.MintLimit(&_Vetrogateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vetrogateway *VetrogatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vetrogateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vetrogateway *VetrogatewaySession) Owner() (common.Address, error) {
	return _Vetrogateway.Contract.Owner(&_Vetrogateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vetrogateway *VetrogatewayCallerSession) Owner() (common.Address, error) {
	return _Vetrogateway.Contract.Owner(&_Vetrogateway.CallOpts)
}

// PegBand is a free data retrieval call binding the contract method 0x726c19b6.
//
// Solidity: function pegBand(address token_) view returns(uint256)
func (_Vetrogateway *VetrogatewayCaller) PegBand(opts *bind.CallOpts, token_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vetrogateway.contract.Call(opts, &out, "pegBand", token_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PegBand is a free data retrieval call binding the contract method 0x726c19b6.
//
// Solidity: function pegBand(address token_) view returns(uint256)
func (_Vetrogateway *VetrogatewaySession) PegBand(token_ common.Address) (*big.Int, error) {
	return _Vetrogateway.Contract.PegBand(&_Vetrogateway.CallOpts, token_)
}

// PegBand is a free data retrieval call binding the contract method 0x726c19b6.
//
// Solidity: function pegBand(address token_) view returns(uint256)
func (_Vetrogateway *VetrogatewayCallerSession) PegBand(token_ common.Address) (*big.Int, error) {
	return _Vetrogateway.Contract.PegBand(&_Vetrogateway.CallOpts, token_)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xb8f82b26.
//
// Solidity: function previewDeposit(address tokenIn_, uint256 amountIn_) view returns(uint256)
func (_Vetrogateway *VetrogatewayCaller) PreviewDeposit(opts *bind.CallOpts, tokenIn_ common.Address, amountIn_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vetrogateway.contract.Call(opts, &out, "previewDeposit", tokenIn_, amountIn_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xb8f82b26.
//
// Solidity: function previewDeposit(address tokenIn_, uint256 amountIn_) view returns(uint256)
func (_Vetrogateway *VetrogatewaySession) PreviewDeposit(tokenIn_ common.Address, amountIn_ *big.Int) (*big.Int, error) {
	return _Vetrogateway.Contract.PreviewDeposit(&_Vetrogateway.CallOpts, tokenIn_, amountIn_)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xb8f82b26.
//
// Solidity: function previewDeposit(address tokenIn_, uint256 amountIn_) view returns(uint256)
func (_Vetrogateway *VetrogatewayCallerSession) PreviewDeposit(tokenIn_ common.Address, amountIn_ *big.Int) (*big.Int, error) {
	return _Vetrogateway.Contract.PreviewDeposit(&_Vetrogateway.CallOpts, tokenIn_, amountIn_)
}

// PreviewMint is a free data retrieval call binding the contract method 0xd1f810a5.
//
// Solidity: function previewMint(address tokenIn_, uint256 peggedTokenOut_) view returns(uint256)
func (_Vetrogateway *VetrogatewayCaller) PreviewMint(opts *bind.CallOpts, tokenIn_ common.Address, peggedTokenOut_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vetrogateway.contract.Call(opts, &out, "previewMint", tokenIn_, peggedTokenOut_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xd1f810a5.
//
// Solidity: function previewMint(address tokenIn_, uint256 peggedTokenOut_) view returns(uint256)
func (_Vetrogateway *VetrogatewaySession) PreviewMint(tokenIn_ common.Address, peggedTokenOut_ *big.Int) (*big.Int, error) {
	return _Vetrogateway.Contract.PreviewMint(&_Vetrogateway.CallOpts, tokenIn_, peggedTokenOut_)
}

// PreviewMint is a free data retrieval call binding the contract method 0xd1f810a5.
//
// Solidity: function previewMint(address tokenIn_, uint256 peggedTokenOut_) view returns(uint256)
func (_Vetrogateway *VetrogatewayCallerSession) PreviewMint(tokenIn_ common.Address, peggedTokenOut_ *big.Int) (*big.Int, error) {
	return _Vetrogateway.Contract.PreviewMint(&_Vetrogateway.CallOpts, tokenIn_, peggedTokenOut_)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0xcbe52ae3.
//
// Solidity: function previewRedeem(address tokenOut_, uint256 peggedTokenIn_) view returns(uint256)
func (_Vetrogateway *VetrogatewayCaller) PreviewRedeem(opts *bind.CallOpts, tokenOut_ common.Address, peggedTokenIn_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vetrogateway.contract.Call(opts, &out, "previewRedeem", tokenOut_, peggedTokenIn_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0xcbe52ae3.
//
// Solidity: function previewRedeem(address tokenOut_, uint256 peggedTokenIn_) view returns(uint256)
func (_Vetrogateway *VetrogatewaySession) PreviewRedeem(tokenOut_ common.Address, peggedTokenIn_ *big.Int) (*big.Int, error) {
	return _Vetrogateway.Contract.PreviewRedeem(&_Vetrogateway.CallOpts, tokenOut_, peggedTokenIn_)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0xcbe52ae3.
//
// Solidity: function previewRedeem(address tokenOut_, uint256 peggedTokenIn_) view returns(uint256)
func (_Vetrogateway *VetrogatewayCallerSession) PreviewRedeem(tokenOut_ common.Address, peggedTokenIn_ *big.Int) (*big.Int, error) {
	return _Vetrogateway.Contract.PreviewRedeem(&_Vetrogateway.CallOpts, tokenOut_, peggedTokenIn_)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0xbbc6f1dc.
//
// Solidity: function previewWithdraw(address tokenOut_, uint256 amountOut_) view returns(uint256)
func (_Vetrogateway *VetrogatewayCaller) PreviewWithdraw(opts *bind.CallOpts, tokenOut_ common.Address, amountOut_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vetrogateway.contract.Call(opts, &out, "previewWithdraw", tokenOut_, amountOut_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0xbbc6f1dc.
//
// Solidity: function previewWithdraw(address tokenOut_, uint256 amountOut_) view returns(uint256)
func (_Vetrogateway *VetrogatewaySession) PreviewWithdraw(tokenOut_ common.Address, amountOut_ *big.Int) (*big.Int, error) {
	return _Vetrogateway.Contract.PreviewWithdraw(&_Vetrogateway.CallOpts, tokenOut_, amountOut_)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0xbbc6f1dc.
//
// Solidity: function previewWithdraw(address tokenOut_, uint256 amountOut_) view returns(uint256)
func (_Vetrogateway *VetrogatewayCallerSession) PreviewWithdraw(tokenOut_ common.Address, amountOut_ *big.Int) (*big.Int, error) {
	return _Vetrogateway.Contract.PreviewWithdraw(&_Vetrogateway.CallOpts, tokenOut_, amountOut_)
}

// RedeemFee is a free data retrieval call binding the contract method 0xde7a8d11.
//
// Solidity: function redeemFee(address token_) view returns(uint256)
func (_Vetrogateway *VetrogatewayCaller) RedeemFee(opts *bind.CallOpts, token_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vetrogateway.contract.Call(opts, &out, "redeemFee", token_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RedeemFee is a free data retrieval call binding the contract method 0xde7a8d11.
//
// Solidity: function redeemFee(address token_) view returns(uint256)
func (_Vetrogateway *VetrogatewaySession) RedeemFee(token_ common.Address) (*big.Int, error) {
	return _Vetrogateway.Contract.RedeemFee(&_Vetrogateway.CallOpts, token_)
}

// RedeemFee is a free data retrieval call binding the contract method 0xde7a8d11.
//
// Solidity: function redeemFee(address token_) view returns(uint256)
func (_Vetrogateway *VetrogatewayCallerSession) RedeemFee(token_ common.Address) (*big.Int, error) {
	return _Vetrogateway.Contract.RedeemFee(&_Vetrogateway.CallOpts, token_)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Vetrogateway *VetrogatewayCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vetrogateway.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Vetrogateway *VetrogatewaySession) Treasury() (common.Address, error) {
	return _Vetrogateway.Contract.Treasury(&_Vetrogateway.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Vetrogateway *VetrogatewayCallerSession) Treasury() (common.Address, error) {
	return _Vetrogateway.Contract.Treasury(&_Vetrogateway.CallOpts)
}

// WithdrawalDelay is a free data retrieval call binding the contract method 0xa7ab6961.
//
// Solidity: function withdrawalDelay() view returns(uint256)
func (_Vetrogateway *VetrogatewayCaller) WithdrawalDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vetrogateway.contract.Call(opts, &out, "withdrawalDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalDelay is a free data retrieval call binding the contract method 0xa7ab6961.
//
// Solidity: function withdrawalDelay() view returns(uint256)
func (_Vetrogateway *VetrogatewaySession) WithdrawalDelay() (*big.Int, error) {
	return _Vetrogateway.Contract.WithdrawalDelay(&_Vetrogateway.CallOpts)
}

// WithdrawalDelay is a free data retrieval call binding the contract method 0xa7ab6961.
//
// Solidity: function withdrawalDelay() view returns(uint256)
func (_Vetrogateway *VetrogatewayCallerSession) WithdrawalDelay() (*big.Int, error) {
	return _Vetrogateway.Contract.WithdrawalDelay(&_Vetrogateway.CallOpts)
}

// WithdrawalDelayEnabled is a free data retrieval call binding the contract method 0x4b5f70e9.
//
// Solidity: function withdrawalDelayEnabled() view returns(bool)
func (_Vetrogateway *VetrogatewayCaller) WithdrawalDelayEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vetrogateway.contract.Call(opts, &out, "withdrawalDelayEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WithdrawalDelayEnabled is a free data retrieval call binding the contract method 0x4b5f70e9.
//
// Solidity: function withdrawalDelayEnabled() view returns(bool)
func (_Vetrogateway *VetrogatewaySession) WithdrawalDelayEnabled() (bool, error) {
	return _Vetrogateway.Contract.WithdrawalDelayEnabled(&_Vetrogateway.CallOpts)
}

// WithdrawalDelayEnabled is a free data retrieval call binding the contract method 0x4b5f70e9.
//
// Solidity: function withdrawalDelayEnabled() view returns(bool)
func (_Vetrogateway *VetrogatewayCallerSession) WithdrawalDelayEnabled() (bool, error) {
	return _Vetrogateway.Contract.WithdrawalDelayEnabled(&_Vetrogateway.CallOpts)
}

// AddToInstantRedeemWhitelist is a paid mutator transaction binding the contract method 0xd7811ca8.
//
// Solidity: function addToInstantRedeemWhitelist(address account_) returns()
func (_Vetrogateway *VetrogatewayTransactor) AddToInstantRedeemWhitelist(opts *bind.TransactOpts, account_ common.Address) (*types.Transaction, error) {
	return _Vetrogateway.contract.Transact(opts, "addToInstantRedeemWhitelist", account_)
}

// AddToInstantRedeemWhitelist is a paid mutator transaction binding the contract method 0xd7811ca8.
//
// Solidity: function addToInstantRedeemWhitelist(address account_) returns()
func (_Vetrogateway *VetrogatewaySession) AddToInstantRedeemWhitelist(account_ common.Address) (*types.Transaction, error) {
	return _Vetrogateway.Contract.AddToInstantRedeemWhitelist(&_Vetrogateway.TransactOpts, account_)
}

// AddToInstantRedeemWhitelist is a paid mutator transaction binding the contract method 0xd7811ca8.
//
// Solidity: function addToInstantRedeemWhitelist(address account_) returns()
func (_Vetrogateway *VetrogatewayTransactorSession) AddToInstantRedeemWhitelist(account_ common.Address) (*types.Transaction, error) {
	return _Vetrogateway.Contract.AddToInstantRedeemWhitelist(&_Vetrogateway.TransactOpts, account_)
}

// BurnFromAMO is a paid mutator transaction binding the contract method 0x85a8715c.
//
// Solidity: function burnFromAMO(uint256 amount_) returns()
func (_Vetrogateway *VetrogatewayTransactor) BurnFromAMO(opts *bind.TransactOpts, amount_ *big.Int) (*types.Transaction, error) {
	return _Vetrogateway.contract.Transact(opts, "burnFromAMO", amount_)
}

// BurnFromAMO is a paid mutator transaction binding the contract method 0x85a8715c.
//
// Solidity: function burnFromAMO(uint256 amount_) returns()
func (_Vetrogateway *VetrogatewaySession) BurnFromAMO(amount_ *big.Int) (*types.Transaction, error) {
	return _Vetrogateway.Contract.BurnFromAMO(&_Vetrogateway.TransactOpts, amount_)
}

// BurnFromAMO is a paid mutator transaction binding the contract method 0x85a8715c.
//
// Solidity: function burnFromAMO(uint256 amount_) returns()
func (_Vetrogateway *VetrogatewayTransactorSession) BurnFromAMO(amount_ *big.Int) (*types.Transaction, error) {
	return _Vetrogateway.Contract.BurnFromAMO(&_Vetrogateway.TransactOpts, amount_)
}

// CancelRedeemRequest is a paid mutator transaction binding the contract method 0xd392ff84.
//
// Solidity: function cancelRedeemRequest() returns()
func (_Vetrogateway *VetrogatewayTransactor) CancelRedeemRequest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vetrogateway.contract.Transact(opts, "cancelRedeemRequest")
}

// CancelRedeemRequest is a paid mutator transaction binding the contract method 0xd392ff84.
//
// Solidity: function cancelRedeemRequest() returns()
func (_Vetrogateway *VetrogatewaySession) CancelRedeemRequest() (*types.Transaction, error) {
	return _Vetrogateway.Contract.CancelRedeemRequest(&_Vetrogateway.TransactOpts)
}

// CancelRedeemRequest is a paid mutator transaction binding the contract method 0xd392ff84.
//
// Solidity: function cancelRedeemRequest() returns()
func (_Vetrogateway *VetrogatewayTransactorSession) CancelRedeemRequest() (*types.Transaction, error) {
	return _Vetrogateway.Contract.CancelRedeemRequest(&_Vetrogateway.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x8b6099db.
//
// Solidity: function deposit(address tokenIn_, uint256 amountIn_, uint256 minPeggedTokenOut_, address receiver_) returns(uint256)
func (_Vetrogateway *VetrogatewayTransactor) Deposit(opts *bind.TransactOpts, tokenIn_ common.Address, amountIn_ *big.Int, minPeggedTokenOut_ *big.Int, receiver_ common.Address) (*types.Transaction, error) {
	return _Vetrogateway.contract.Transact(opts, "deposit", tokenIn_, amountIn_, minPeggedTokenOut_, receiver_)
}

// Deposit is a paid mutator transaction binding the contract method 0x8b6099db.
//
// Solidity: function deposit(address tokenIn_, uint256 amountIn_, uint256 minPeggedTokenOut_, address receiver_) returns(uint256)
func (_Vetrogateway *VetrogatewaySession) Deposit(tokenIn_ common.Address, amountIn_ *big.Int, minPeggedTokenOut_ *big.Int, receiver_ common.Address) (*types.Transaction, error) {
	return _Vetrogateway.Contract.Deposit(&_Vetrogateway.TransactOpts, tokenIn_, amountIn_, minPeggedTokenOut_, receiver_)
}

// Deposit is a paid mutator transaction binding the contract method 0x8b6099db.
//
// Solidity: function deposit(address tokenIn_, uint256 amountIn_, uint256 minPeggedTokenOut_, address receiver_) returns(uint256)
func (_Vetrogateway *VetrogatewayTransactorSession) Deposit(tokenIn_ common.Address, amountIn_ *big.Int, minPeggedTokenOut_ *big.Int, receiver_ common.Address) (*types.Transaction, error) {
	return _Vetrogateway.Contract.Deposit(&_Vetrogateway.TransactOpts, tokenIn_, amountIn_, minPeggedTokenOut_, receiver_)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address peggedToken_, uint256 mintLimit_, uint256 initialWithdrawalDelay_) returns()
func (_Vetrogateway *VetrogatewayTransactor) Initialize(opts *bind.TransactOpts, peggedToken_ common.Address, mintLimit_ *big.Int, initialWithdrawalDelay_ *big.Int) (*types.Transaction, error) {
	return _Vetrogateway.contract.Transact(opts, "initialize", peggedToken_, mintLimit_, initialWithdrawalDelay_)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address peggedToken_, uint256 mintLimit_, uint256 initialWithdrawalDelay_) returns()
func (_Vetrogateway *VetrogatewaySession) Initialize(peggedToken_ common.Address, mintLimit_ *big.Int, initialWithdrawalDelay_ *big.Int) (*types.Transaction, error) {
	return _Vetrogateway.Contract.Initialize(&_Vetrogateway.TransactOpts, peggedToken_, mintLimit_, initialWithdrawalDelay_)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address peggedToken_, uint256 mintLimit_, uint256 initialWithdrawalDelay_) returns()
func (_Vetrogateway *VetrogatewayTransactorSession) Initialize(peggedToken_ common.Address, mintLimit_ *big.Int, initialWithdrawalDelay_ *big.Int) (*types.Transaction, error) {
	return _Vetrogateway.Contract.Initialize(&_Vetrogateway.TransactOpts, peggedToken_, mintLimit_, initialWithdrawalDelay_)
}

// Mint is a paid mutator transaction binding the contract method 0xf74bfe8e.
//
// Solidity: function mint(address tokenIn_, uint256 peggedTokenOut_, uint256 maxAmountIn_, address receiver_) returns(uint256)
func (_Vetrogateway *VetrogatewayTransactor) Mint(opts *bind.TransactOpts, tokenIn_ common.Address, peggedTokenOut_ *big.Int, maxAmountIn_ *big.Int, receiver_ common.Address) (*types.Transaction, error) {
	return _Vetrogateway.contract.Transact(opts, "mint", tokenIn_, peggedTokenOut_, maxAmountIn_, receiver_)
}

// Mint is a paid mutator transaction binding the contract method 0xf74bfe8e.
//
// Solidity: function mint(address tokenIn_, uint256 peggedTokenOut_, uint256 maxAmountIn_, address receiver_) returns(uint256)
func (_Vetrogateway *VetrogatewaySession) Mint(tokenIn_ common.Address, peggedTokenOut_ *big.Int, maxAmountIn_ *big.Int, receiver_ common.Address) (*types.Transaction, error) {
	return _Vetrogateway.Contract.Mint(&_Vetrogateway.TransactOpts, tokenIn_, peggedTokenOut_, maxAmountIn_, receiver_)
}

// Mint is a paid mutator transaction binding the contract method 0xf74bfe8e.
//
// Solidity: function mint(address tokenIn_, uint256 peggedTokenOut_, uint256 maxAmountIn_, address receiver_) returns(uint256)
func (_Vetrogateway *VetrogatewayTransactorSession) Mint(tokenIn_ common.Address, peggedTokenOut_ *big.Int, maxAmountIn_ *big.Int, receiver_ common.Address) (*types.Transaction, error) {
	return _Vetrogateway.Contract.Mint(&_Vetrogateway.TransactOpts, tokenIn_, peggedTokenOut_, maxAmountIn_, receiver_)
}

// MintToAMO is a paid mutator transaction binding the contract method 0xbba6b6d9.
//
// Solidity: function mintToAMO(uint256 amount_, address receiver_) returns()
func (_Vetrogateway *VetrogatewayTransactor) MintToAMO(opts *bind.TransactOpts, amount_ *big.Int, receiver_ common.Address) (*types.Transaction, error) {
	return _Vetrogateway.contract.Transact(opts, "mintToAMO", amount_, receiver_)
}

// MintToAMO is a paid mutator transaction binding the contract method 0xbba6b6d9.
//
// Solidity: function mintToAMO(uint256 amount_, address receiver_) returns()
func (_Vetrogateway *VetrogatewaySession) MintToAMO(amount_ *big.Int, receiver_ common.Address) (*types.Transaction, error) {
	return _Vetrogateway.Contract.MintToAMO(&_Vetrogateway.TransactOpts, amount_, receiver_)
}

// MintToAMO is a paid mutator transaction binding the contract method 0xbba6b6d9.
//
// Solidity: function mintToAMO(uint256 amount_, address receiver_) returns()
func (_Vetrogateway *VetrogatewayTransactorSession) MintToAMO(amount_ *big.Int, receiver_ common.Address) (*types.Transaction, error) {
	return _Vetrogateway.Contract.MintToAMO(&_Vetrogateway.TransactOpts, amount_, receiver_)
}

// Redeem is a paid mutator transaction binding the contract method 0x43bcfab6.
//
// Solidity: function redeem(address tokenOut_, uint256 peggedTokenIn_, uint256 minAmountOut_, address receiver_) returns(uint256)
func (_Vetrogateway *VetrogatewayTransactor) Redeem(opts *bind.TransactOpts, tokenOut_ common.Address, peggedTokenIn_ *big.Int, minAmountOut_ *big.Int, receiver_ common.Address) (*types.Transaction, error) {
	return _Vetrogateway.contract.Transact(opts, "redeem", tokenOut_, peggedTokenIn_, minAmountOut_, receiver_)
}

// Redeem is a paid mutator transaction binding the contract method 0x43bcfab6.
//
// Solidity: function redeem(address tokenOut_, uint256 peggedTokenIn_, uint256 minAmountOut_, address receiver_) returns(uint256)
func (_Vetrogateway *VetrogatewaySession) Redeem(tokenOut_ common.Address, peggedTokenIn_ *big.Int, minAmountOut_ *big.Int, receiver_ common.Address) (*types.Transaction, error) {
	return _Vetrogateway.Contract.Redeem(&_Vetrogateway.TransactOpts, tokenOut_, peggedTokenIn_, minAmountOut_, receiver_)
}

// Redeem is a paid mutator transaction binding the contract method 0x43bcfab6.
//
// Solidity: function redeem(address tokenOut_, uint256 peggedTokenIn_, uint256 minAmountOut_, address receiver_) returns(uint256)
func (_Vetrogateway *VetrogatewayTransactorSession) Redeem(tokenOut_ common.Address, peggedTokenIn_ *big.Int, minAmountOut_ *big.Int, receiver_ common.Address) (*types.Transaction, error) {
	return _Vetrogateway.Contract.Redeem(&_Vetrogateway.TransactOpts, tokenOut_, peggedTokenIn_, minAmountOut_, receiver_)
}

// RemoveFromInstantRedeemWhitelist is a paid mutator transaction binding the contract method 0xf4a8bfce.
//
// Solidity: function removeFromInstantRedeemWhitelist(address account_) returns()
func (_Vetrogateway *VetrogatewayTransactor) RemoveFromInstantRedeemWhitelist(opts *bind.TransactOpts, account_ common.Address) (*types.Transaction, error) {
	return _Vetrogateway.contract.Transact(opts, "removeFromInstantRedeemWhitelist", account_)
}

// RemoveFromInstantRedeemWhitelist is a paid mutator transaction binding the contract method 0xf4a8bfce.
//
// Solidity: function removeFromInstantRedeemWhitelist(address account_) returns()
func (_Vetrogateway *VetrogatewaySession) RemoveFromInstantRedeemWhitelist(account_ common.Address) (*types.Transaction, error) {
	return _Vetrogateway.Contract.RemoveFromInstantRedeemWhitelist(&_Vetrogateway.TransactOpts, account_)
}

// RemoveFromInstantRedeemWhitelist is a paid mutator transaction binding the contract method 0xf4a8bfce.
//
// Solidity: function removeFromInstantRedeemWhitelist(address account_) returns()
func (_Vetrogateway *VetrogatewayTransactorSession) RemoveFromInstantRedeemWhitelist(account_ common.Address) (*types.Transaction, error) {
	return _Vetrogateway.Contract.RemoveFromInstantRedeemWhitelist(&_Vetrogateway.TransactOpts, account_)
}

// RequestRedeem is a paid mutator transaction binding the contract method 0xaa2f892d.
//
// Solidity: function requestRedeem(uint256 peggedTokenAmount_) returns()
func (_Vetrogateway *VetrogatewayTransactor) RequestRedeem(opts *bind.TransactOpts, peggedTokenAmount_ *big.Int) (*types.Transaction, error) {
	return _Vetrogateway.contract.Transact(opts, "requestRedeem", peggedTokenAmount_)
}

// RequestRedeem is a paid mutator transaction binding the contract method 0xaa2f892d.
//
// Solidity: function requestRedeem(uint256 peggedTokenAmount_) returns()
func (_Vetrogateway *VetrogatewaySession) RequestRedeem(peggedTokenAmount_ *big.Int) (*types.Transaction, error) {
	return _Vetrogateway.Contract.RequestRedeem(&_Vetrogateway.TransactOpts, peggedTokenAmount_)
}

// RequestRedeem is a paid mutator transaction binding the contract method 0xaa2f892d.
//
// Solidity: function requestRedeem(uint256 peggedTokenAmount_) returns()
func (_Vetrogateway *VetrogatewayTransactorSession) RequestRedeem(peggedTokenAmount_ *big.Int) (*types.Transaction, error) {
	return _Vetrogateway.Contract.RequestRedeem(&_Vetrogateway.TransactOpts, peggedTokenAmount_)
}

// SetWithdrawalDelayEnabled is a paid mutator transaction binding the contract method 0x0ca4bea6.
//
// Solidity: function setWithdrawalDelayEnabled(bool enabled_) returns()
func (_Vetrogateway *VetrogatewayTransactor) SetWithdrawalDelayEnabled(opts *bind.TransactOpts, enabled_ bool) (*types.Transaction, error) {
	return _Vetrogateway.contract.Transact(opts, "setWithdrawalDelayEnabled", enabled_)
}

// SetWithdrawalDelayEnabled is a paid mutator transaction binding the contract method 0x0ca4bea6.
//
// Solidity: function setWithdrawalDelayEnabled(bool enabled_) returns()
func (_Vetrogateway *VetrogatewaySession) SetWithdrawalDelayEnabled(enabled_ bool) (*types.Transaction, error) {
	return _Vetrogateway.Contract.SetWithdrawalDelayEnabled(&_Vetrogateway.TransactOpts, enabled_)
}

// SetWithdrawalDelayEnabled is a paid mutator transaction binding the contract method 0x0ca4bea6.
//
// Solidity: function setWithdrawalDelayEnabled(bool enabled_) returns()
func (_Vetrogateway *VetrogatewayTransactorSession) SetWithdrawalDelayEnabled(enabled_ bool) (*types.Transaction, error) {
	return _Vetrogateway.Contract.SetWithdrawalDelayEnabled(&_Vetrogateway.TransactOpts, enabled_)
}

// UpdateAmoMintLimit is a paid mutator transaction binding the contract method 0xdc4aca0b.
//
// Solidity: function updateAmoMintLimit(uint256 newAmoMintLimit_) returns()
func (_Vetrogateway *VetrogatewayTransactor) UpdateAmoMintLimit(opts *bind.TransactOpts, newAmoMintLimit_ *big.Int) (*types.Transaction, error) {
	return _Vetrogateway.contract.Transact(opts, "updateAmoMintLimit", newAmoMintLimit_)
}

// UpdateAmoMintLimit is a paid mutator transaction binding the contract method 0xdc4aca0b.
//
// Solidity: function updateAmoMintLimit(uint256 newAmoMintLimit_) returns()
func (_Vetrogateway *VetrogatewaySession) UpdateAmoMintLimit(newAmoMintLimit_ *big.Int) (*types.Transaction, error) {
	return _Vetrogateway.Contract.UpdateAmoMintLimit(&_Vetrogateway.TransactOpts, newAmoMintLimit_)
}

// UpdateAmoMintLimit is a paid mutator transaction binding the contract method 0xdc4aca0b.
//
// Solidity: function updateAmoMintLimit(uint256 newAmoMintLimit_) returns()
func (_Vetrogateway *VetrogatewayTransactorSession) UpdateAmoMintLimit(newAmoMintLimit_ *big.Int) (*types.Transaction, error) {
	return _Vetrogateway.Contract.UpdateAmoMintLimit(&_Vetrogateway.TransactOpts, newAmoMintLimit_)
}

// UpdateMintFee is a paid mutator transaction binding the contract method 0x37c1c973.
//
// Solidity: function updateMintFee(address token_, uint256 newMintFee_) returns()
func (_Vetrogateway *VetrogatewayTransactor) UpdateMintFee(opts *bind.TransactOpts, token_ common.Address, newMintFee_ *big.Int) (*types.Transaction, error) {
	return _Vetrogateway.contract.Transact(opts, "updateMintFee", token_, newMintFee_)
}

// UpdateMintFee is a paid mutator transaction binding the contract method 0x37c1c973.
//
// Solidity: function updateMintFee(address token_, uint256 newMintFee_) returns()
func (_Vetrogateway *VetrogatewaySession) UpdateMintFee(token_ common.Address, newMintFee_ *big.Int) (*types.Transaction, error) {
	return _Vetrogateway.Contract.UpdateMintFee(&_Vetrogateway.TransactOpts, token_, newMintFee_)
}

// UpdateMintFee is a paid mutator transaction binding the contract method 0x37c1c973.
//
// Solidity: function updateMintFee(address token_, uint256 newMintFee_) returns()
func (_Vetrogateway *VetrogatewayTransactorSession) UpdateMintFee(token_ common.Address, newMintFee_ *big.Int) (*types.Transaction, error) {
	return _Vetrogateway.Contract.UpdateMintFee(&_Vetrogateway.TransactOpts, token_, newMintFee_)
}

// UpdateMintLimit is a paid mutator transaction binding the contract method 0xe01d55c5.
//
// Solidity: function updateMintLimit(uint256 newMintLimit_) returns()
func (_Vetrogateway *VetrogatewayTransactor) UpdateMintLimit(opts *bind.TransactOpts, newMintLimit_ *big.Int) (*types.Transaction, error) {
	return _Vetrogateway.contract.Transact(opts, "updateMintLimit", newMintLimit_)
}

// UpdateMintLimit is a paid mutator transaction binding the contract method 0xe01d55c5.
//
// Solidity: function updateMintLimit(uint256 newMintLimit_) returns()
func (_Vetrogateway *VetrogatewaySession) UpdateMintLimit(newMintLimit_ *big.Int) (*types.Transaction, error) {
	return _Vetrogateway.Contract.UpdateMintLimit(&_Vetrogateway.TransactOpts, newMintLimit_)
}

// UpdateMintLimit is a paid mutator transaction binding the contract method 0xe01d55c5.
//
// Solidity: function updateMintLimit(uint256 newMintLimit_) returns()
func (_Vetrogateway *VetrogatewayTransactorSession) UpdateMintLimit(newMintLimit_ *big.Int) (*types.Transaction, error) {
	return _Vetrogateway.Contract.UpdateMintLimit(&_Vetrogateway.TransactOpts, newMintLimit_)
}

// UpdatePegBand is a paid mutator transaction binding the contract method 0xef34e1d3.
//
// Solidity: function updatePegBand(address token_, uint256 newPegBandBps_) returns()
func (_Vetrogateway *VetrogatewayTransactor) UpdatePegBand(opts *bind.TransactOpts, token_ common.Address, newPegBandBps_ *big.Int) (*types.Transaction, error) {
	return _Vetrogateway.contract.Transact(opts, "updatePegBand", token_, newPegBandBps_)
}

// UpdatePegBand is a paid mutator transaction binding the contract method 0xef34e1d3.
//
// Solidity: function updatePegBand(address token_, uint256 newPegBandBps_) returns()
func (_Vetrogateway *VetrogatewaySession) UpdatePegBand(token_ common.Address, newPegBandBps_ *big.Int) (*types.Transaction, error) {
	return _Vetrogateway.Contract.UpdatePegBand(&_Vetrogateway.TransactOpts, token_, newPegBandBps_)
}

// UpdatePegBand is a paid mutator transaction binding the contract method 0xef34e1d3.
//
// Solidity: function updatePegBand(address token_, uint256 newPegBandBps_) returns()
func (_Vetrogateway *VetrogatewayTransactorSession) UpdatePegBand(token_ common.Address, newPegBandBps_ *big.Int) (*types.Transaction, error) {
	return _Vetrogateway.Contract.UpdatePegBand(&_Vetrogateway.TransactOpts, token_, newPegBandBps_)
}

// UpdateRedeemFee is a paid mutator transaction binding the contract method 0x6dd86aa9.
//
// Solidity: function updateRedeemFee(address token_, uint256 newRedeemFee_) returns()
func (_Vetrogateway *VetrogatewayTransactor) UpdateRedeemFee(opts *bind.TransactOpts, token_ common.Address, newRedeemFee_ *big.Int) (*types.Transaction, error) {
	return _Vetrogateway.contract.Transact(opts, "updateRedeemFee", token_, newRedeemFee_)
}

// UpdateRedeemFee is a paid mutator transaction binding the contract method 0x6dd86aa9.
//
// Solidity: function updateRedeemFee(address token_, uint256 newRedeemFee_) returns()
func (_Vetrogateway *VetrogatewaySession) UpdateRedeemFee(token_ common.Address, newRedeemFee_ *big.Int) (*types.Transaction, error) {
	return _Vetrogateway.Contract.UpdateRedeemFee(&_Vetrogateway.TransactOpts, token_, newRedeemFee_)
}

// UpdateRedeemFee is a paid mutator transaction binding the contract method 0x6dd86aa9.
//
// Solidity: function updateRedeemFee(address token_, uint256 newRedeemFee_) returns()
func (_Vetrogateway *VetrogatewayTransactorSession) UpdateRedeemFee(token_ common.Address, newRedeemFee_ *big.Int) (*types.Transaction, error) {
	return _Vetrogateway.Contract.UpdateRedeemFee(&_Vetrogateway.TransactOpts, token_, newRedeemFee_)
}

// UpdateWithdrawalDelay is a paid mutator transaction binding the contract method 0x52cd67e9.
//
// Solidity: function updateWithdrawalDelay(uint256 newDelay_) returns()
func (_Vetrogateway *VetrogatewayTransactor) UpdateWithdrawalDelay(opts *bind.TransactOpts, newDelay_ *big.Int) (*types.Transaction, error) {
	return _Vetrogateway.contract.Transact(opts, "updateWithdrawalDelay", newDelay_)
}

// UpdateWithdrawalDelay is a paid mutator transaction binding the contract method 0x52cd67e9.
//
// Solidity: function updateWithdrawalDelay(uint256 newDelay_) returns()
func (_Vetrogateway *VetrogatewaySession) UpdateWithdrawalDelay(newDelay_ *big.Int) (*types.Transaction, error) {
	return _Vetrogateway.Contract.UpdateWithdrawalDelay(&_Vetrogateway.TransactOpts, newDelay_)
}

// UpdateWithdrawalDelay is a paid mutator transaction binding the contract method 0x52cd67e9.
//
// Solidity: function updateWithdrawalDelay(uint256 newDelay_) returns()
func (_Vetrogateway *VetrogatewayTransactorSession) UpdateWithdrawalDelay(newDelay_ *big.Int) (*types.Transaction, error) {
	return _Vetrogateway.Contract.UpdateWithdrawalDelay(&_Vetrogateway.TransactOpts, newDelay_)
}

// Withdraw is a paid mutator transaction binding the contract method 0x16762eed.
//
// Solidity: function withdraw(address tokenOut_, uint256 amountOut_, uint256 maxPeggedTokenIn_, address receiver_) returns(uint256)
func (_Vetrogateway *VetrogatewayTransactor) Withdraw(opts *bind.TransactOpts, tokenOut_ common.Address, amountOut_ *big.Int, maxPeggedTokenIn_ *big.Int, receiver_ common.Address) (*types.Transaction, error) {
	return _Vetrogateway.contract.Transact(opts, "withdraw", tokenOut_, amountOut_, maxPeggedTokenIn_, receiver_)
}

// Withdraw is a paid mutator transaction binding the contract method 0x16762eed.
//
// Solidity: function withdraw(address tokenOut_, uint256 amountOut_, uint256 maxPeggedTokenIn_, address receiver_) returns(uint256)
func (_Vetrogateway *VetrogatewaySession) Withdraw(tokenOut_ common.Address, amountOut_ *big.Int, maxPeggedTokenIn_ *big.Int, receiver_ common.Address) (*types.Transaction, error) {
	return _Vetrogateway.Contract.Withdraw(&_Vetrogateway.TransactOpts, tokenOut_, amountOut_, maxPeggedTokenIn_, receiver_)
}

// Withdraw is a paid mutator transaction binding the contract method 0x16762eed.
//
// Solidity: function withdraw(address tokenOut_, uint256 amountOut_, uint256 maxPeggedTokenIn_, address receiver_) returns(uint256)
func (_Vetrogateway *VetrogatewayTransactorSession) Withdraw(tokenOut_ common.Address, amountOut_ *big.Int, maxPeggedTokenIn_ *big.Int, receiver_ common.Address) (*types.Transaction, error) {
	return _Vetrogateway.Contract.Withdraw(&_Vetrogateway.TransactOpts, tokenOut_, amountOut_, maxPeggedTokenIn_, receiver_)
}

// VetrogatewayAddedToInstantRedeemWhitelistIterator is returned from FilterAddedToInstantRedeemWhitelist and is used to iterate over the raw logs and unpacked data for AddedToInstantRedeemWhitelist events raised by the Vetrogateway contract.
type VetrogatewayAddedToInstantRedeemWhitelistIterator struct {
	Event *VetrogatewayAddedToInstantRedeemWhitelist // Event containing the contract specifics and raw log

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
func (it *VetrogatewayAddedToInstantRedeemWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VetrogatewayAddedToInstantRedeemWhitelist)
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
		it.Event = new(VetrogatewayAddedToInstantRedeemWhitelist)
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
func (it *VetrogatewayAddedToInstantRedeemWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VetrogatewayAddedToInstantRedeemWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VetrogatewayAddedToInstantRedeemWhitelist represents a AddedToInstantRedeemWhitelist event raised by the Vetrogateway contract.
type VetrogatewayAddedToInstantRedeemWhitelist struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddedToInstantRedeemWhitelist is a free log retrieval operation binding the contract event 0x797517369e23ffd3af45a8a42a13b6db0a4826ad2ef98121c434e14f150e6240.
//
// Solidity: event AddedToInstantRedeemWhitelist(address indexed account)
func (_Vetrogateway *VetrogatewayFilterer) FilterAddedToInstantRedeemWhitelist(opts *bind.FilterOpts, account []common.Address) (*VetrogatewayAddedToInstantRedeemWhitelistIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Vetrogateway.contract.FilterLogs(opts, "AddedToInstantRedeemWhitelist", accountRule)
	if err != nil {
		return nil, err
	}
	return &VetrogatewayAddedToInstantRedeemWhitelistIterator{contract: _Vetrogateway.contract, event: "AddedToInstantRedeemWhitelist", logs: logs, sub: sub}, nil
}

// WatchAddedToInstantRedeemWhitelist is a free log subscription operation binding the contract event 0x797517369e23ffd3af45a8a42a13b6db0a4826ad2ef98121c434e14f150e6240.
//
// Solidity: event AddedToInstantRedeemWhitelist(address indexed account)
func (_Vetrogateway *VetrogatewayFilterer) WatchAddedToInstantRedeemWhitelist(opts *bind.WatchOpts, sink chan<- *VetrogatewayAddedToInstantRedeemWhitelist, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Vetrogateway.contract.WatchLogs(opts, "AddedToInstantRedeemWhitelist", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VetrogatewayAddedToInstantRedeemWhitelist)
				if err := _Vetrogateway.contract.UnpackLog(event, "AddedToInstantRedeemWhitelist", log); err != nil {
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

// ParseAddedToInstantRedeemWhitelist is a log parse operation binding the contract event 0x797517369e23ffd3af45a8a42a13b6db0a4826ad2ef98121c434e14f150e6240.
//
// Solidity: event AddedToInstantRedeemWhitelist(address indexed account)
func (_Vetrogateway *VetrogatewayFilterer) ParseAddedToInstantRedeemWhitelist(log types.Log) (*VetrogatewayAddedToInstantRedeemWhitelist, error) {
	event := new(VetrogatewayAddedToInstantRedeemWhitelist)
	if err := _Vetrogateway.contract.UnpackLog(event, "AddedToInstantRedeemWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VetrogatewayBurnedFromAMOIterator is returned from FilterBurnedFromAMO and is used to iterate over the raw logs and unpacked data for BurnedFromAMO events raised by the Vetrogateway contract.
type VetrogatewayBurnedFromAMOIterator struct {
	Event *VetrogatewayBurnedFromAMO // Event containing the contract specifics and raw log

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
func (it *VetrogatewayBurnedFromAMOIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VetrogatewayBurnedFromAMO)
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
		it.Event = new(VetrogatewayBurnedFromAMO)
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
func (it *VetrogatewayBurnedFromAMOIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VetrogatewayBurnedFromAMOIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VetrogatewayBurnedFromAMO represents a BurnedFromAMO event raised by the Vetrogateway contract.
type VetrogatewayBurnedFromAMO struct {
	AmountBurned *big.Int
	NewAmoSupply *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBurnedFromAMO is a free log retrieval operation binding the contract event 0xa840d857be4fa627a8b83209753d23db02a27dc32782579e9332a69d5fc881dd.
//
// Solidity: event BurnedFromAMO(uint256 amountBurned, uint256 newAmoSupply)
func (_Vetrogateway *VetrogatewayFilterer) FilterBurnedFromAMO(opts *bind.FilterOpts) (*VetrogatewayBurnedFromAMOIterator, error) {

	logs, sub, err := _Vetrogateway.contract.FilterLogs(opts, "BurnedFromAMO")
	if err != nil {
		return nil, err
	}
	return &VetrogatewayBurnedFromAMOIterator{contract: _Vetrogateway.contract, event: "BurnedFromAMO", logs: logs, sub: sub}, nil
}

// WatchBurnedFromAMO is a free log subscription operation binding the contract event 0xa840d857be4fa627a8b83209753d23db02a27dc32782579e9332a69d5fc881dd.
//
// Solidity: event BurnedFromAMO(uint256 amountBurned, uint256 newAmoSupply)
func (_Vetrogateway *VetrogatewayFilterer) WatchBurnedFromAMO(opts *bind.WatchOpts, sink chan<- *VetrogatewayBurnedFromAMO) (event.Subscription, error) {

	logs, sub, err := _Vetrogateway.contract.WatchLogs(opts, "BurnedFromAMO")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VetrogatewayBurnedFromAMO)
				if err := _Vetrogateway.contract.UnpackLog(event, "BurnedFromAMO", log); err != nil {
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

// ParseBurnedFromAMO is a log parse operation binding the contract event 0xa840d857be4fa627a8b83209753d23db02a27dc32782579e9332a69d5fc881dd.
//
// Solidity: event BurnedFromAMO(uint256 amountBurned, uint256 newAmoSupply)
func (_Vetrogateway *VetrogatewayFilterer) ParseBurnedFromAMO(log types.Log) (*VetrogatewayBurnedFromAMO, error) {
	event := new(VetrogatewayBurnedFromAMO)
	if err := _Vetrogateway.contract.UnpackLog(event, "BurnedFromAMO", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VetrogatewayDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Vetrogateway contract.
type VetrogatewayDepositIterator struct {
	Event *VetrogatewayDeposit // Event containing the contract specifics and raw log

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
func (it *VetrogatewayDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VetrogatewayDeposit)
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
		it.Event = new(VetrogatewayDeposit)
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
func (it *VetrogatewayDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VetrogatewayDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VetrogatewayDeposit represents a Deposit event raised by the Vetrogateway contract.
type VetrogatewayDeposit struct {
	Token             common.Address
	TokenAmount       *big.Int
	PeggedTokenAmount *big.Int
	Receiver          common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x02d7e648dd130fc184d383e55bb126ac4c9c60e8f94bf05acdf557ba2d540b47.
//
// Solidity: event Deposit(address indexed token, uint256 tokenAmount, uint256 peggedTokenAmount, address indexed receiver)
func (_Vetrogateway *VetrogatewayFilterer) FilterDeposit(opts *bind.FilterOpts, token []common.Address, receiver []common.Address) (*VetrogatewayDepositIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Vetrogateway.contract.FilterLogs(opts, "Deposit", tokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &VetrogatewayDepositIterator{contract: _Vetrogateway.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x02d7e648dd130fc184d383e55bb126ac4c9c60e8f94bf05acdf557ba2d540b47.
//
// Solidity: event Deposit(address indexed token, uint256 tokenAmount, uint256 peggedTokenAmount, address indexed receiver)
func (_Vetrogateway *VetrogatewayFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *VetrogatewayDeposit, token []common.Address, receiver []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Vetrogateway.contract.WatchLogs(opts, "Deposit", tokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VetrogatewayDeposit)
				if err := _Vetrogateway.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x02d7e648dd130fc184d383e55bb126ac4c9c60e8f94bf05acdf557ba2d540b47.
//
// Solidity: event Deposit(address indexed token, uint256 tokenAmount, uint256 peggedTokenAmount, address indexed receiver)
func (_Vetrogateway *VetrogatewayFilterer) ParseDeposit(log types.Log) (*VetrogatewayDeposit, error) {
	event := new(VetrogatewayDeposit)
	if err := _Vetrogateway.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VetrogatewayInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Vetrogateway contract.
type VetrogatewayInitializedIterator struct {
	Event *VetrogatewayInitialized // Event containing the contract specifics and raw log

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
func (it *VetrogatewayInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VetrogatewayInitialized)
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
		it.Event = new(VetrogatewayInitialized)
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
func (it *VetrogatewayInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VetrogatewayInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VetrogatewayInitialized represents a Initialized event raised by the Vetrogateway contract.
type VetrogatewayInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Vetrogateway *VetrogatewayFilterer) FilterInitialized(opts *bind.FilterOpts) (*VetrogatewayInitializedIterator, error) {

	logs, sub, err := _Vetrogateway.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &VetrogatewayInitializedIterator{contract: _Vetrogateway.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Vetrogateway *VetrogatewayFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *VetrogatewayInitialized) (event.Subscription, error) {

	logs, sub, err := _Vetrogateway.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VetrogatewayInitialized)
				if err := _Vetrogateway.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Vetrogateway *VetrogatewayFilterer) ParseInitialized(log types.Log) (*VetrogatewayInitialized, error) {
	event := new(VetrogatewayInitialized)
	if err := _Vetrogateway.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VetrogatewayMintFeeUpdatedIterator is returned from FilterMintFeeUpdated and is used to iterate over the raw logs and unpacked data for MintFeeUpdated events raised by the Vetrogateway contract.
type VetrogatewayMintFeeUpdatedIterator struct {
	Event *VetrogatewayMintFeeUpdated // Event containing the contract specifics and raw log

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
func (it *VetrogatewayMintFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VetrogatewayMintFeeUpdated)
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
		it.Event = new(VetrogatewayMintFeeUpdated)
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
func (it *VetrogatewayMintFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VetrogatewayMintFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VetrogatewayMintFeeUpdated represents a MintFeeUpdated event raised by the Vetrogateway contract.
type VetrogatewayMintFeeUpdated struct {
	Token           common.Address
	PreviousMintFee *big.Int
	NewMintFee      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMintFeeUpdated is a free log retrieval operation binding the contract event 0x1a15b544093bd3556343e4087772291340e180214a1c35469a6dd8aa1a4d3c68.
//
// Solidity: event MintFeeUpdated(address indexed token, uint256 previousMintFee, uint256 newMintFee)
func (_Vetrogateway *VetrogatewayFilterer) FilterMintFeeUpdated(opts *bind.FilterOpts, token []common.Address) (*VetrogatewayMintFeeUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Vetrogateway.contract.FilterLogs(opts, "MintFeeUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &VetrogatewayMintFeeUpdatedIterator{contract: _Vetrogateway.contract, event: "MintFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchMintFeeUpdated is a free log subscription operation binding the contract event 0x1a15b544093bd3556343e4087772291340e180214a1c35469a6dd8aa1a4d3c68.
//
// Solidity: event MintFeeUpdated(address indexed token, uint256 previousMintFee, uint256 newMintFee)
func (_Vetrogateway *VetrogatewayFilterer) WatchMintFeeUpdated(opts *bind.WatchOpts, sink chan<- *VetrogatewayMintFeeUpdated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Vetrogateway.contract.WatchLogs(opts, "MintFeeUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VetrogatewayMintFeeUpdated)
				if err := _Vetrogateway.contract.UnpackLog(event, "MintFeeUpdated", log); err != nil {
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

// ParseMintFeeUpdated is a log parse operation binding the contract event 0x1a15b544093bd3556343e4087772291340e180214a1c35469a6dd8aa1a4d3c68.
//
// Solidity: event MintFeeUpdated(address indexed token, uint256 previousMintFee, uint256 newMintFee)
func (_Vetrogateway *VetrogatewayFilterer) ParseMintFeeUpdated(log types.Log) (*VetrogatewayMintFeeUpdated, error) {
	event := new(VetrogatewayMintFeeUpdated)
	if err := _Vetrogateway.contract.UnpackLog(event, "MintFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VetrogatewayMintLimitUpdatedIterator is returned from FilterMintLimitUpdated and is used to iterate over the raw logs and unpacked data for MintLimitUpdated events raised by the Vetrogateway contract.
type VetrogatewayMintLimitUpdatedIterator struct {
	Event *VetrogatewayMintLimitUpdated // Event containing the contract specifics and raw log

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
func (it *VetrogatewayMintLimitUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VetrogatewayMintLimitUpdated)
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
		it.Event = new(VetrogatewayMintLimitUpdated)
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
func (it *VetrogatewayMintLimitUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VetrogatewayMintLimitUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VetrogatewayMintLimitUpdated represents a MintLimitUpdated event raised by the Vetrogateway contract.
type VetrogatewayMintLimitUpdated struct {
	PreviousMintLimit *big.Int
	NewMintLimit      *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterMintLimitUpdated is a free log retrieval operation binding the contract event 0x864790bdf9878a0378c6fc2b0ce53bf74ca13b901bc97a1cb94aa88f1600e482.
//
// Solidity: event MintLimitUpdated(uint256 previousMintLimit, uint256 newMintLimit)
func (_Vetrogateway *VetrogatewayFilterer) FilterMintLimitUpdated(opts *bind.FilterOpts) (*VetrogatewayMintLimitUpdatedIterator, error) {

	logs, sub, err := _Vetrogateway.contract.FilterLogs(opts, "MintLimitUpdated")
	if err != nil {
		return nil, err
	}
	return &VetrogatewayMintLimitUpdatedIterator{contract: _Vetrogateway.contract, event: "MintLimitUpdated", logs: logs, sub: sub}, nil
}

// WatchMintLimitUpdated is a free log subscription operation binding the contract event 0x864790bdf9878a0378c6fc2b0ce53bf74ca13b901bc97a1cb94aa88f1600e482.
//
// Solidity: event MintLimitUpdated(uint256 previousMintLimit, uint256 newMintLimit)
func (_Vetrogateway *VetrogatewayFilterer) WatchMintLimitUpdated(opts *bind.WatchOpts, sink chan<- *VetrogatewayMintLimitUpdated) (event.Subscription, error) {

	logs, sub, err := _Vetrogateway.contract.WatchLogs(opts, "MintLimitUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VetrogatewayMintLimitUpdated)
				if err := _Vetrogateway.contract.UnpackLog(event, "MintLimitUpdated", log); err != nil {
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

// ParseMintLimitUpdated is a log parse operation binding the contract event 0x864790bdf9878a0378c6fc2b0ce53bf74ca13b901bc97a1cb94aa88f1600e482.
//
// Solidity: event MintLimitUpdated(uint256 previousMintLimit, uint256 newMintLimit)
func (_Vetrogateway *VetrogatewayFilterer) ParseMintLimitUpdated(log types.Log) (*VetrogatewayMintLimitUpdated, error) {
	event := new(VetrogatewayMintLimitUpdated)
	if err := _Vetrogateway.contract.UnpackLog(event, "MintLimitUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VetrogatewayMintedToAMOIterator is returned from FilterMintedToAMO and is used to iterate over the raw logs and unpacked data for MintedToAMO events raised by the Vetrogateway contract.
type VetrogatewayMintedToAMOIterator struct {
	Event *VetrogatewayMintedToAMO // Event containing the contract specifics and raw log

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
func (it *VetrogatewayMintedToAMOIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VetrogatewayMintedToAMO)
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
		it.Event = new(VetrogatewayMintedToAMO)
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
func (it *VetrogatewayMintedToAMOIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VetrogatewayMintedToAMOIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VetrogatewayMintedToAMO represents a MintedToAMO event raised by the Vetrogateway contract.
type VetrogatewayMintedToAMO struct {
	Receiver     common.Address
	AmountMinted *big.Int
	NewAmoSupply *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMintedToAMO is a free log retrieval operation binding the contract event 0x5b12134695cb4f1d04e9c6f5747e4f28d084940f9ac912162a12ea4c50790c7b.
//
// Solidity: event MintedToAMO(address indexed receiver, uint256 amountMinted, uint256 newAmoSupply)
func (_Vetrogateway *VetrogatewayFilterer) FilterMintedToAMO(opts *bind.FilterOpts, receiver []common.Address) (*VetrogatewayMintedToAMOIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Vetrogateway.contract.FilterLogs(opts, "MintedToAMO", receiverRule)
	if err != nil {
		return nil, err
	}
	return &VetrogatewayMintedToAMOIterator{contract: _Vetrogateway.contract, event: "MintedToAMO", logs: logs, sub: sub}, nil
}

// WatchMintedToAMO is a free log subscription operation binding the contract event 0x5b12134695cb4f1d04e9c6f5747e4f28d084940f9ac912162a12ea4c50790c7b.
//
// Solidity: event MintedToAMO(address indexed receiver, uint256 amountMinted, uint256 newAmoSupply)
func (_Vetrogateway *VetrogatewayFilterer) WatchMintedToAMO(opts *bind.WatchOpts, sink chan<- *VetrogatewayMintedToAMO, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Vetrogateway.contract.WatchLogs(opts, "MintedToAMO", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VetrogatewayMintedToAMO)
				if err := _Vetrogateway.contract.UnpackLog(event, "MintedToAMO", log); err != nil {
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

// ParseMintedToAMO is a log parse operation binding the contract event 0x5b12134695cb4f1d04e9c6f5747e4f28d084940f9ac912162a12ea4c50790c7b.
//
// Solidity: event MintedToAMO(address indexed receiver, uint256 amountMinted, uint256 newAmoSupply)
func (_Vetrogateway *VetrogatewayFilterer) ParseMintedToAMO(log types.Log) (*VetrogatewayMintedToAMO, error) {
	event := new(VetrogatewayMintedToAMO)
	if err := _Vetrogateway.contract.UnpackLog(event, "MintedToAMO", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VetrogatewayPegBandUpdatedIterator is returned from FilterPegBandUpdated and is used to iterate over the raw logs and unpacked data for PegBandUpdated events raised by the Vetrogateway contract.
type VetrogatewayPegBandUpdatedIterator struct {
	Event *VetrogatewayPegBandUpdated // Event containing the contract specifics and raw log

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
func (it *VetrogatewayPegBandUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VetrogatewayPegBandUpdated)
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
		it.Event = new(VetrogatewayPegBandUpdated)
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
func (it *VetrogatewayPegBandUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VetrogatewayPegBandUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VetrogatewayPegBandUpdated represents a PegBandUpdated event raised by the Vetrogateway contract.
type VetrogatewayPegBandUpdated struct {
	Token              common.Address
	PreviousPegBandBps *big.Int
	NewPegBandBps      *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPegBandUpdated is a free log retrieval operation binding the contract event 0x91b5c669b95c944e391cafcd29858c0b8d0f840d1efececa05babf51b3d4684a.
//
// Solidity: event PegBandUpdated(address indexed token, uint256 previousPegBandBps, uint256 newPegBandBps)
func (_Vetrogateway *VetrogatewayFilterer) FilterPegBandUpdated(opts *bind.FilterOpts, token []common.Address) (*VetrogatewayPegBandUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Vetrogateway.contract.FilterLogs(opts, "PegBandUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &VetrogatewayPegBandUpdatedIterator{contract: _Vetrogateway.contract, event: "PegBandUpdated", logs: logs, sub: sub}, nil
}

// WatchPegBandUpdated is a free log subscription operation binding the contract event 0x91b5c669b95c944e391cafcd29858c0b8d0f840d1efececa05babf51b3d4684a.
//
// Solidity: event PegBandUpdated(address indexed token, uint256 previousPegBandBps, uint256 newPegBandBps)
func (_Vetrogateway *VetrogatewayFilterer) WatchPegBandUpdated(opts *bind.WatchOpts, sink chan<- *VetrogatewayPegBandUpdated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Vetrogateway.contract.WatchLogs(opts, "PegBandUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VetrogatewayPegBandUpdated)
				if err := _Vetrogateway.contract.UnpackLog(event, "PegBandUpdated", log); err != nil {
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

// ParsePegBandUpdated is a log parse operation binding the contract event 0x91b5c669b95c944e391cafcd29858c0b8d0f840d1efececa05babf51b3d4684a.
//
// Solidity: event PegBandUpdated(address indexed token, uint256 previousPegBandBps, uint256 newPegBandBps)
func (_Vetrogateway *VetrogatewayFilterer) ParsePegBandUpdated(log types.Log) (*VetrogatewayPegBandUpdated, error) {
	event := new(VetrogatewayPegBandUpdated)
	if err := _Vetrogateway.contract.UnpackLog(event, "PegBandUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VetrogatewayRedeemFeeUpdatedIterator is returned from FilterRedeemFeeUpdated and is used to iterate over the raw logs and unpacked data for RedeemFeeUpdated events raised by the Vetrogateway contract.
type VetrogatewayRedeemFeeUpdatedIterator struct {
	Event *VetrogatewayRedeemFeeUpdated // Event containing the contract specifics and raw log

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
func (it *VetrogatewayRedeemFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VetrogatewayRedeemFeeUpdated)
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
		it.Event = new(VetrogatewayRedeemFeeUpdated)
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
func (it *VetrogatewayRedeemFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VetrogatewayRedeemFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VetrogatewayRedeemFeeUpdated represents a RedeemFeeUpdated event raised by the Vetrogateway contract.
type VetrogatewayRedeemFeeUpdated struct {
	Token             common.Address
	PreviousRedeemFee *big.Int
	NewRedeemFee      *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRedeemFeeUpdated is a free log retrieval operation binding the contract event 0x08cf3e85f82bb8308a1225e82694466368f02a7803dbfac0b5b7a5db38d52ca2.
//
// Solidity: event RedeemFeeUpdated(address indexed token, uint256 previousRedeemFee, uint256 newRedeemFee)
func (_Vetrogateway *VetrogatewayFilterer) FilterRedeemFeeUpdated(opts *bind.FilterOpts, token []common.Address) (*VetrogatewayRedeemFeeUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Vetrogateway.contract.FilterLogs(opts, "RedeemFeeUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &VetrogatewayRedeemFeeUpdatedIterator{contract: _Vetrogateway.contract, event: "RedeemFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchRedeemFeeUpdated is a free log subscription operation binding the contract event 0x08cf3e85f82bb8308a1225e82694466368f02a7803dbfac0b5b7a5db38d52ca2.
//
// Solidity: event RedeemFeeUpdated(address indexed token, uint256 previousRedeemFee, uint256 newRedeemFee)
func (_Vetrogateway *VetrogatewayFilterer) WatchRedeemFeeUpdated(opts *bind.WatchOpts, sink chan<- *VetrogatewayRedeemFeeUpdated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Vetrogateway.contract.WatchLogs(opts, "RedeemFeeUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VetrogatewayRedeemFeeUpdated)
				if err := _Vetrogateway.contract.UnpackLog(event, "RedeemFeeUpdated", log); err != nil {
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

// ParseRedeemFeeUpdated is a log parse operation binding the contract event 0x08cf3e85f82bb8308a1225e82694466368f02a7803dbfac0b5b7a5db38d52ca2.
//
// Solidity: event RedeemFeeUpdated(address indexed token, uint256 previousRedeemFee, uint256 newRedeemFee)
func (_Vetrogateway *VetrogatewayFilterer) ParseRedeemFeeUpdated(log types.Log) (*VetrogatewayRedeemFeeUpdated, error) {
	event := new(VetrogatewayRedeemFeeUpdated)
	if err := _Vetrogateway.contract.UnpackLog(event, "RedeemFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VetrogatewayRedeemRequestCancelledIterator is returned from FilterRedeemRequestCancelled and is used to iterate over the raw logs and unpacked data for RedeemRequestCancelled events raised by the Vetrogateway contract.
type VetrogatewayRedeemRequestCancelledIterator struct {
	Event *VetrogatewayRedeemRequestCancelled // Event containing the contract specifics and raw log

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
func (it *VetrogatewayRedeemRequestCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VetrogatewayRedeemRequestCancelled)
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
		it.Event = new(VetrogatewayRedeemRequestCancelled)
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
func (it *VetrogatewayRedeemRequestCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VetrogatewayRedeemRequestCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VetrogatewayRedeemRequestCancelled represents a RedeemRequestCancelled event raised by the Vetrogateway contract.
type VetrogatewayRedeemRequestCancelled struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRedeemRequestCancelled is a free log retrieval operation binding the contract event 0x046d3e17087d323298991516e3fdec36fe1204d6bff7fb6e2b8eca77b80a22ad.
//
// Solidity: event RedeemRequestCancelled(address indexed user, uint256 amount)
func (_Vetrogateway *VetrogatewayFilterer) FilterRedeemRequestCancelled(opts *bind.FilterOpts, user []common.Address) (*VetrogatewayRedeemRequestCancelledIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Vetrogateway.contract.FilterLogs(opts, "RedeemRequestCancelled", userRule)
	if err != nil {
		return nil, err
	}
	return &VetrogatewayRedeemRequestCancelledIterator{contract: _Vetrogateway.contract, event: "RedeemRequestCancelled", logs: logs, sub: sub}, nil
}

// WatchRedeemRequestCancelled is a free log subscription operation binding the contract event 0x046d3e17087d323298991516e3fdec36fe1204d6bff7fb6e2b8eca77b80a22ad.
//
// Solidity: event RedeemRequestCancelled(address indexed user, uint256 amount)
func (_Vetrogateway *VetrogatewayFilterer) WatchRedeemRequestCancelled(opts *bind.WatchOpts, sink chan<- *VetrogatewayRedeemRequestCancelled, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Vetrogateway.contract.WatchLogs(opts, "RedeemRequestCancelled", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VetrogatewayRedeemRequestCancelled)
				if err := _Vetrogateway.contract.UnpackLog(event, "RedeemRequestCancelled", log); err != nil {
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

// ParseRedeemRequestCancelled is a log parse operation binding the contract event 0x046d3e17087d323298991516e3fdec36fe1204d6bff7fb6e2b8eca77b80a22ad.
//
// Solidity: event RedeemRequestCancelled(address indexed user, uint256 amount)
func (_Vetrogateway *VetrogatewayFilterer) ParseRedeemRequestCancelled(log types.Log) (*VetrogatewayRedeemRequestCancelled, error) {
	event := new(VetrogatewayRedeemRequestCancelled)
	if err := _Vetrogateway.contract.UnpackLog(event, "RedeemRequestCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VetrogatewayRedeemRequestedIterator is returned from FilterRedeemRequested and is used to iterate over the raw logs and unpacked data for RedeemRequested events raised by the Vetrogateway contract.
type VetrogatewayRedeemRequestedIterator struct {
	Event *VetrogatewayRedeemRequested // Event containing the contract specifics and raw log

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
func (it *VetrogatewayRedeemRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VetrogatewayRedeemRequested)
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
		it.Event = new(VetrogatewayRedeemRequested)
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
func (it *VetrogatewayRedeemRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VetrogatewayRedeemRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VetrogatewayRedeemRequested represents a RedeemRequested event raised by the Vetrogateway contract.
type VetrogatewayRedeemRequested struct {
	User        common.Address
	Amount      *big.Int
	ClaimableAt *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRedeemRequested is a free log retrieval operation binding the contract event 0x58fe322fc5911ed072ec92f570e517b9793e350eb1ff7be0019fd9f3fade87bc.
//
// Solidity: event RedeemRequested(address indexed user, uint256 amount, uint256 claimableAt)
func (_Vetrogateway *VetrogatewayFilterer) FilterRedeemRequested(opts *bind.FilterOpts, user []common.Address) (*VetrogatewayRedeemRequestedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Vetrogateway.contract.FilterLogs(opts, "RedeemRequested", userRule)
	if err != nil {
		return nil, err
	}
	return &VetrogatewayRedeemRequestedIterator{contract: _Vetrogateway.contract, event: "RedeemRequested", logs: logs, sub: sub}, nil
}

// WatchRedeemRequested is a free log subscription operation binding the contract event 0x58fe322fc5911ed072ec92f570e517b9793e350eb1ff7be0019fd9f3fade87bc.
//
// Solidity: event RedeemRequested(address indexed user, uint256 amount, uint256 claimableAt)
func (_Vetrogateway *VetrogatewayFilterer) WatchRedeemRequested(opts *bind.WatchOpts, sink chan<- *VetrogatewayRedeemRequested, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Vetrogateway.contract.WatchLogs(opts, "RedeemRequested", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VetrogatewayRedeemRequested)
				if err := _Vetrogateway.contract.UnpackLog(event, "RedeemRequested", log); err != nil {
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

// ParseRedeemRequested is a log parse operation binding the contract event 0x58fe322fc5911ed072ec92f570e517b9793e350eb1ff7be0019fd9f3fade87bc.
//
// Solidity: event RedeemRequested(address indexed user, uint256 amount, uint256 claimableAt)
func (_Vetrogateway *VetrogatewayFilterer) ParseRedeemRequested(log types.Log) (*VetrogatewayRedeemRequested, error) {
	event := new(VetrogatewayRedeemRequested)
	if err := _Vetrogateway.contract.UnpackLog(event, "RedeemRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VetrogatewayRemovedFromInstantRedeemWhitelistIterator is returned from FilterRemovedFromInstantRedeemWhitelist and is used to iterate over the raw logs and unpacked data for RemovedFromInstantRedeemWhitelist events raised by the Vetrogateway contract.
type VetrogatewayRemovedFromInstantRedeemWhitelistIterator struct {
	Event *VetrogatewayRemovedFromInstantRedeemWhitelist // Event containing the contract specifics and raw log

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
func (it *VetrogatewayRemovedFromInstantRedeemWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VetrogatewayRemovedFromInstantRedeemWhitelist)
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
		it.Event = new(VetrogatewayRemovedFromInstantRedeemWhitelist)
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
func (it *VetrogatewayRemovedFromInstantRedeemWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VetrogatewayRemovedFromInstantRedeemWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VetrogatewayRemovedFromInstantRedeemWhitelist represents a RemovedFromInstantRedeemWhitelist event raised by the Vetrogateway contract.
type VetrogatewayRemovedFromInstantRedeemWhitelist struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRemovedFromInstantRedeemWhitelist is a free log retrieval operation binding the contract event 0x33ab47afb8a85cbb067a9be15fcd6a7d60556a446b47f204e65a37cf7ad19c38.
//
// Solidity: event RemovedFromInstantRedeemWhitelist(address indexed account)
func (_Vetrogateway *VetrogatewayFilterer) FilterRemovedFromInstantRedeemWhitelist(opts *bind.FilterOpts, account []common.Address) (*VetrogatewayRemovedFromInstantRedeemWhitelistIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Vetrogateway.contract.FilterLogs(opts, "RemovedFromInstantRedeemWhitelist", accountRule)
	if err != nil {
		return nil, err
	}
	return &VetrogatewayRemovedFromInstantRedeemWhitelistIterator{contract: _Vetrogateway.contract, event: "RemovedFromInstantRedeemWhitelist", logs: logs, sub: sub}, nil
}

// WatchRemovedFromInstantRedeemWhitelist is a free log subscription operation binding the contract event 0x33ab47afb8a85cbb067a9be15fcd6a7d60556a446b47f204e65a37cf7ad19c38.
//
// Solidity: event RemovedFromInstantRedeemWhitelist(address indexed account)
func (_Vetrogateway *VetrogatewayFilterer) WatchRemovedFromInstantRedeemWhitelist(opts *bind.WatchOpts, sink chan<- *VetrogatewayRemovedFromInstantRedeemWhitelist, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Vetrogateway.contract.WatchLogs(opts, "RemovedFromInstantRedeemWhitelist", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VetrogatewayRemovedFromInstantRedeemWhitelist)
				if err := _Vetrogateway.contract.UnpackLog(event, "RemovedFromInstantRedeemWhitelist", log); err != nil {
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

// ParseRemovedFromInstantRedeemWhitelist is a log parse operation binding the contract event 0x33ab47afb8a85cbb067a9be15fcd6a7d60556a446b47f204e65a37cf7ad19c38.
//
// Solidity: event RemovedFromInstantRedeemWhitelist(address indexed account)
func (_Vetrogateway *VetrogatewayFilterer) ParseRemovedFromInstantRedeemWhitelist(log types.Log) (*VetrogatewayRemovedFromInstantRedeemWhitelist, error) {
	event := new(VetrogatewayRemovedFromInstantRedeemWhitelist)
	if err := _Vetrogateway.contract.UnpackLog(event, "RemovedFromInstantRedeemWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VetrogatewayUpdatedAmoMintLimitIterator is returned from FilterUpdatedAmoMintLimit and is used to iterate over the raw logs and unpacked data for UpdatedAmoMintLimit events raised by the Vetrogateway contract.
type VetrogatewayUpdatedAmoMintLimitIterator struct {
	Event *VetrogatewayUpdatedAmoMintLimit // Event containing the contract specifics and raw log

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
func (it *VetrogatewayUpdatedAmoMintLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VetrogatewayUpdatedAmoMintLimit)
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
		it.Event = new(VetrogatewayUpdatedAmoMintLimit)
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
func (it *VetrogatewayUpdatedAmoMintLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VetrogatewayUpdatedAmoMintLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VetrogatewayUpdatedAmoMintLimit represents a UpdatedAmoMintLimit event raised by the Vetrogateway contract.
type VetrogatewayUpdatedAmoMintLimit struct {
	PreviousLimit *big.Int
	NewLimit      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedAmoMintLimit is a free log retrieval operation binding the contract event 0xa15e75b23c4dcef1647713af3c7fdf585008390108be4b7936c438297047726d.
//
// Solidity: event UpdatedAmoMintLimit(uint256 previousLimit, uint256 newLimit)
func (_Vetrogateway *VetrogatewayFilterer) FilterUpdatedAmoMintLimit(opts *bind.FilterOpts) (*VetrogatewayUpdatedAmoMintLimitIterator, error) {

	logs, sub, err := _Vetrogateway.contract.FilterLogs(opts, "UpdatedAmoMintLimit")
	if err != nil {
		return nil, err
	}
	return &VetrogatewayUpdatedAmoMintLimitIterator{contract: _Vetrogateway.contract, event: "UpdatedAmoMintLimit", logs: logs, sub: sub}, nil
}

// WatchUpdatedAmoMintLimit is a free log subscription operation binding the contract event 0xa15e75b23c4dcef1647713af3c7fdf585008390108be4b7936c438297047726d.
//
// Solidity: event UpdatedAmoMintLimit(uint256 previousLimit, uint256 newLimit)
func (_Vetrogateway *VetrogatewayFilterer) WatchUpdatedAmoMintLimit(opts *bind.WatchOpts, sink chan<- *VetrogatewayUpdatedAmoMintLimit) (event.Subscription, error) {

	logs, sub, err := _Vetrogateway.contract.WatchLogs(opts, "UpdatedAmoMintLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VetrogatewayUpdatedAmoMintLimit)
				if err := _Vetrogateway.contract.UnpackLog(event, "UpdatedAmoMintLimit", log); err != nil {
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

// ParseUpdatedAmoMintLimit is a log parse operation binding the contract event 0xa15e75b23c4dcef1647713af3c7fdf585008390108be4b7936c438297047726d.
//
// Solidity: event UpdatedAmoMintLimit(uint256 previousLimit, uint256 newLimit)
func (_Vetrogateway *VetrogatewayFilterer) ParseUpdatedAmoMintLimit(log types.Log) (*VetrogatewayUpdatedAmoMintLimit, error) {
	event := new(VetrogatewayUpdatedAmoMintLimit)
	if err := _Vetrogateway.contract.UnpackLog(event, "UpdatedAmoMintLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VetrogatewayWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Vetrogateway contract.
type VetrogatewayWithdrawIterator struct {
	Event *VetrogatewayWithdraw // Event containing the contract specifics and raw log

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
func (it *VetrogatewayWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VetrogatewayWithdraw)
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
		it.Event = new(VetrogatewayWithdraw)
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
func (it *VetrogatewayWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VetrogatewayWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VetrogatewayWithdraw represents a Withdraw event raised by the Vetrogateway contract.
type VetrogatewayWithdraw struct {
	Token             common.Address
	TokenAmount       *big.Int
	PeggedTokenAmount *big.Int
	Receiver          common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x8166bf25f8a2b7ed3c85049207da4358d16edbed977d23fa2ee6f0dde3ec2132.
//
// Solidity: event Withdraw(address indexed token, uint256 tokenAmount, uint256 peggedTokenAmount, address indexed receiver)
func (_Vetrogateway *VetrogatewayFilterer) FilterWithdraw(opts *bind.FilterOpts, token []common.Address, receiver []common.Address) (*VetrogatewayWithdrawIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Vetrogateway.contract.FilterLogs(opts, "Withdraw", tokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &VetrogatewayWithdrawIterator{contract: _Vetrogateway.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x8166bf25f8a2b7ed3c85049207da4358d16edbed977d23fa2ee6f0dde3ec2132.
//
// Solidity: event Withdraw(address indexed token, uint256 tokenAmount, uint256 peggedTokenAmount, address indexed receiver)
func (_Vetrogateway *VetrogatewayFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *VetrogatewayWithdraw, token []common.Address, receiver []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Vetrogateway.contract.WatchLogs(opts, "Withdraw", tokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VetrogatewayWithdraw)
				if err := _Vetrogateway.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x8166bf25f8a2b7ed3c85049207da4358d16edbed977d23fa2ee6f0dde3ec2132.
//
// Solidity: event Withdraw(address indexed token, uint256 tokenAmount, uint256 peggedTokenAmount, address indexed receiver)
func (_Vetrogateway *VetrogatewayFilterer) ParseWithdraw(log types.Log) (*VetrogatewayWithdraw, error) {
	event := new(VetrogatewayWithdraw)
	if err := _Vetrogateway.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VetrogatewayWithdrawalDelayEnabledIterator is returned from FilterWithdrawalDelayEnabled and is used to iterate over the raw logs and unpacked data for WithdrawalDelayEnabled events raised by the Vetrogateway contract.
type VetrogatewayWithdrawalDelayEnabledIterator struct {
	Event *VetrogatewayWithdrawalDelayEnabled // Event containing the contract specifics and raw log

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
func (it *VetrogatewayWithdrawalDelayEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VetrogatewayWithdrawalDelayEnabled)
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
		it.Event = new(VetrogatewayWithdrawalDelayEnabled)
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
func (it *VetrogatewayWithdrawalDelayEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VetrogatewayWithdrawalDelayEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VetrogatewayWithdrawalDelayEnabled represents a WithdrawalDelayEnabled event raised by the Vetrogateway contract.
type VetrogatewayWithdrawalDelayEnabled struct {
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalDelayEnabled is a free log retrieval operation binding the contract event 0x1ab3cf66155e779afea5e22ccc9d9309ab0b8ac9f8bb280be843272f54202ead.
//
// Solidity: event WithdrawalDelayEnabled(bool enabled)
func (_Vetrogateway *VetrogatewayFilterer) FilterWithdrawalDelayEnabled(opts *bind.FilterOpts) (*VetrogatewayWithdrawalDelayEnabledIterator, error) {

	logs, sub, err := _Vetrogateway.contract.FilterLogs(opts, "WithdrawalDelayEnabled")
	if err != nil {
		return nil, err
	}
	return &VetrogatewayWithdrawalDelayEnabledIterator{contract: _Vetrogateway.contract, event: "WithdrawalDelayEnabled", logs: logs, sub: sub}, nil
}

// WatchWithdrawalDelayEnabled is a free log subscription operation binding the contract event 0x1ab3cf66155e779afea5e22ccc9d9309ab0b8ac9f8bb280be843272f54202ead.
//
// Solidity: event WithdrawalDelayEnabled(bool enabled)
func (_Vetrogateway *VetrogatewayFilterer) WatchWithdrawalDelayEnabled(opts *bind.WatchOpts, sink chan<- *VetrogatewayWithdrawalDelayEnabled) (event.Subscription, error) {

	logs, sub, err := _Vetrogateway.contract.WatchLogs(opts, "WithdrawalDelayEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VetrogatewayWithdrawalDelayEnabled)
				if err := _Vetrogateway.contract.UnpackLog(event, "WithdrawalDelayEnabled", log); err != nil {
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

// ParseWithdrawalDelayEnabled is a log parse operation binding the contract event 0x1ab3cf66155e779afea5e22ccc9d9309ab0b8ac9f8bb280be843272f54202ead.
//
// Solidity: event WithdrawalDelayEnabled(bool enabled)
func (_Vetrogateway *VetrogatewayFilterer) ParseWithdrawalDelayEnabled(log types.Log) (*VetrogatewayWithdrawalDelayEnabled, error) {
	event := new(VetrogatewayWithdrawalDelayEnabled)
	if err := _Vetrogateway.contract.UnpackLog(event, "WithdrawalDelayEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VetrogatewayWithdrawalDelayUpdatedIterator is returned from FilterWithdrawalDelayUpdated and is used to iterate over the raw logs and unpacked data for WithdrawalDelayUpdated events raised by the Vetrogateway contract.
type VetrogatewayWithdrawalDelayUpdatedIterator struct {
	Event *VetrogatewayWithdrawalDelayUpdated // Event containing the contract specifics and raw log

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
func (it *VetrogatewayWithdrawalDelayUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VetrogatewayWithdrawalDelayUpdated)
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
		it.Event = new(VetrogatewayWithdrawalDelayUpdated)
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
func (it *VetrogatewayWithdrawalDelayUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VetrogatewayWithdrawalDelayUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VetrogatewayWithdrawalDelayUpdated represents a WithdrawalDelayUpdated event raised by the Vetrogateway contract.
type VetrogatewayWithdrawalDelayUpdated struct {
	PreviousDelay *big.Int
	NewDelay      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalDelayUpdated is a free log retrieval operation binding the contract event 0x9c3f1b54b1487e018f1d0593ff5cf7fb625b2df6332c974a6cc56bb358879841.
//
// Solidity: event WithdrawalDelayUpdated(uint256 previousDelay, uint256 newDelay)
func (_Vetrogateway *VetrogatewayFilterer) FilterWithdrawalDelayUpdated(opts *bind.FilterOpts) (*VetrogatewayWithdrawalDelayUpdatedIterator, error) {

	logs, sub, err := _Vetrogateway.contract.FilterLogs(opts, "WithdrawalDelayUpdated")
	if err != nil {
		return nil, err
	}
	return &VetrogatewayWithdrawalDelayUpdatedIterator{contract: _Vetrogateway.contract, event: "WithdrawalDelayUpdated", logs: logs, sub: sub}, nil
}

// WatchWithdrawalDelayUpdated is a free log subscription operation binding the contract event 0x9c3f1b54b1487e018f1d0593ff5cf7fb625b2df6332c974a6cc56bb358879841.
//
// Solidity: event WithdrawalDelayUpdated(uint256 previousDelay, uint256 newDelay)
func (_Vetrogateway *VetrogatewayFilterer) WatchWithdrawalDelayUpdated(opts *bind.WatchOpts, sink chan<- *VetrogatewayWithdrawalDelayUpdated) (event.Subscription, error) {

	logs, sub, err := _Vetrogateway.contract.WatchLogs(opts, "WithdrawalDelayUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VetrogatewayWithdrawalDelayUpdated)
				if err := _Vetrogateway.contract.UnpackLog(event, "WithdrawalDelayUpdated", log); err != nil {
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

// ParseWithdrawalDelayUpdated is a log parse operation binding the contract event 0x9c3f1b54b1487e018f1d0593ff5cf7fb625b2df6332c974a6cc56bb358879841.
//
// Solidity: event WithdrawalDelayUpdated(uint256 previousDelay, uint256 newDelay)
func (_Vetrogateway *VetrogatewayFilterer) ParseWithdrawalDelayUpdated(log types.Log) (*VetrogatewayWithdrawalDelayUpdated, error) {
	event := new(VetrogatewayWithdrawalDelayUpdated)
	if err := _Vetrogateway.contract.UnpackLog(event, "WithdrawalDelayUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
