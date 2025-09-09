// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package univ4poolmanager

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

// IBunniHubDeployBunniTokenParams is an auto generated low-level Go binding around an user-defined struct.
type IBunniHubDeployBunniTokenParams struct {
	Currency0                common.Address
	Currency1                common.Address
	TickSpacing              *big.Int
	TwapSecondsAgo           *big.Int
	LiquidityDensityFunction common.Address
	Hooklet                  common.Address
	LdfType                  uint8
	LdfParams                [32]byte
	Hooks                    common.Address
	HookParams               []byte
	Vault0                   common.Address
	Vault1                   common.Address
	MinRawTokenRatio0        *big.Int
	TargetRawTokenRatio0     *big.Int
	MaxRawTokenRatio0        *big.Int
	MinRawTokenRatio1        *big.Int
	TargetRawTokenRatio1     *big.Int
	MaxRawTokenRatio1        *big.Int
	SqrtPriceX96             *big.Int
	Name                     [32]byte
	Symbol                   [32]byte
	Owner                    common.Address
	MetadataURI              string
	Salt                     [32]byte
}

// IBunniHubDepositParams is an auto generated low-level Go binding around an user-defined struct.
type IBunniHubDepositParams struct {
	PoolKey         PoolKey
	Recipient       common.Address
	RefundRecipient common.Address
	Amount0Desired  *big.Int
	Amount1Desired  *big.Int
	Amount0Min      *big.Int
	Amount1Min      *big.Int
	VaultFee0       *big.Int
	VaultFee1       *big.Int
	Deadline        *big.Int
}

// IBunniHubQueueWithdrawParams is an auto generated low-level Go binding around an user-defined struct.
type IBunniHubQueueWithdrawParams struct {
	PoolKey PoolKey
	Shares  *big.Int
}

// IBunniHubWithdrawParams is an auto generated low-level Go binding around an user-defined struct.
type IBunniHubWithdrawParams struct {
	PoolKey             PoolKey
	Recipient           common.Address
	Shares              *big.Int
	Amount0Min          *big.Int
	Amount1Min          *big.Int
	Deadline            *big.Int
	UseQueuedWithdrawal bool
}

// PoolKey is an auto generated low-level Go binding around an user-defined struct.
type PoolKey struct {
	Currency0   common.Address
	Currency1   common.Address
	Fee         *big.Int
	TickSpacing *big.Int
	Hooks       common.Address
}

// PoolState is an auto generated low-level Go binding around an user-defined struct.
type PoolState struct {
	LiquidityDensityFunction common.Address
	BunniToken               common.Address
	Hooklet                  common.Address
	TwapSecondsAgo           *big.Int
	LdfParams                [32]byte
	HookParams               []byte
	Vault0                   common.Address
	Vault1                   common.Address
	LdfType                  uint8
	MinRawTokenRatio0        *big.Int
	TargetRawTokenRatio0     *big.Int
	MaxRawTokenRatio0        *big.Int
	MinRawTokenRatio1        *big.Int
	TargetRawTokenRatio1     *big.Int
	MaxRawTokenRatio1        *big.Int
	Currency0Decimals        uint8
	Currency1Decimals        uint8
	Vault0Decimals           uint8
	Vault1Decimals           uint8
	RawBalance0              *big.Int
	RawBalance1              *big.Int
	Reserve0                 *big.Int
	Reserve1                 *big.Int
	IdleBalance              [32]byte
}

// Univ4poolmanagerMetaData contains all meta data concerning the Univ4poolmanager contract.
var Univ4poolmanagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPoolManager\",\"name\":\"poolManager_\",\"type\":\"address\"},{\"internalType\":\"contractWETH\",\"name\":\"weth_\",\"type\":\"address\"},{\"internalType\":\"contractIPermit2\",\"name\":\"permit2_\",\"type\":\"address\"},{\"internalType\":\"contractIBunniToken\",\"name\":\"bunniTokenImplementation_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"contractIBunniHook[]\",\"name\":\"initialHookWhitelist\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BunniHub__BunniTokenNotInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BunniHub__MsgValueInsufficient\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BunniHub__NoExpiredWithdrawal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BunniHub__PastDeadline\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BunniHub__Paused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BunniHub__Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BunniHub__VaultTookMoreThanRequested\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BunniHub__ZeroInput\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewOwnerIsZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoHandoverRequest\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuard__ReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"BurnPauseFuse\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"PoolId\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIBunniToken\",\"name\":\"bunniToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"PoolId\",\"name\":\"poolId\",\"type\":\"bytes32\"}],\"name\":\"NewBunni\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipHandoverCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipHandoverRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"PoolId\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"QueueWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIBunniHook\",\"name\":\"hook\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"whitelisted\",\"type\":\"bool\"}],\"name\":\"SetHookWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"pauseFlags\",\"type\":\"uint8\"}],\"name\":\"SetPauseFlags\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"isPauser\",\"type\":\"bool\"}],\"name\":\"SetPauser\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"PoolId\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"PoolId\",\"name\":\"poolId\",\"type\":\"bytes32\"}],\"name\":\"bunniTokenOfPool\",\"outputs\":[{\"internalType\":\"contractIBunniToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnPauseFuse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"completeOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"Currency\",\"name\":\"currency0\",\"type\":\"address\"},{\"internalType\":\"Currency\",\"name\":\"currency1\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"internalType\":\"uint24\",\"name\":\"twapSecondsAgo\",\"type\":\"uint24\"},{\"internalType\":\"contractILiquidityDensityFunction\",\"name\":\"liquidityDensityFunction\",\"type\":\"address\"},{\"internalType\":\"contractIHooklet\",\"name\":\"hooklet\",\"type\":\"address\"},{\"internalType\":\"enumLDFType\",\"name\":\"ldfType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"ldfParams\",\"type\":\"bytes32\"},{\"internalType\":\"contractIBunniHook\",\"name\":\"hooks\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"hookParams\",\"type\":\"bytes\"},{\"internalType\":\"contractERC4626\",\"name\":\"vault0\",\"type\":\"address\"},{\"internalType\":\"contractERC4626\",\"name\":\"vault1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"minRawTokenRatio0\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"targetRawTokenRatio0\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"maxRawTokenRatio0\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"minRawTokenRatio1\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"targetRawTokenRatio1\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"maxRawTokenRatio1\",\"type\":\"uint24\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"symbol\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"internalType\":\"structIBunniHub.DeployBunniTokenParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"deployBunniToken\",\"outputs\":[{\"internalType\":\"contractIBunniToken\",\"name\":\"token\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"Currency\",\"name\":\"currency0\",\"type\":\"address\"},{\"internalType\":\"Currency\",\"name\":\"currency1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"internalType\":\"contractIHooks\",\"name\":\"hooks\",\"type\":\"address\"}],\"internalType\":\"structPoolKey\",\"name\":\"key\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"Currency\",\"name\":\"currency0\",\"type\":\"address\"},{\"internalType\":\"Currency\",\"name\":\"currency1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"internalType\":\"contractIHooks\",\"name\":\"hooks\",\"type\":\"address\"}],\"internalType\":\"structPoolKey\",\"name\":\"poolKey\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0Desired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Desired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount0Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultFee0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultFee1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structIBunniHub.DepositParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPauseStatus\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"pauseFlags\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"unpauseFuse\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"Currency\",\"name\":\"currency0\",\"type\":\"address\"},{\"internalType\":\"Currency\",\"name\":\"currency1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"internalType\":\"contractIHooks\",\"name\":\"hooks\",\"type\":\"address\"}],\"internalType\":\"structPoolKey\",\"name\":\"key\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"isCurrency0\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"hookGive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"Currency\",\"name\":\"currency0\",\"type\":\"address\"},{\"internalType\":\"Currency\",\"name\":\"currency1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"internalType\":\"contractIHooks\",\"name\":\"hooks\",\"type\":\"address\"}],\"internalType\":\"structPoolKey\",\"name\":\"key\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"zeroForOne\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outputAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"shouldSurge\",\"type\":\"bool\"}],\"name\":\"hookHandleSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBunniHook\",\"name\":\"hook\",\"type\":\"address\"}],\"name\":\"hookIsWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"PoolId\",\"name\":\"poolId\",\"type\":\"bytes32\"}],\"name\":\"hookParams\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"Currency\",\"name\":\"currency0\",\"type\":\"address\"},{\"internalType\":\"Currency\",\"name\":\"currency1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"internalType\":\"contractIHooks\",\"name\":\"hooks\",\"type\":\"address\"}],\"internalType\":\"structPoolKey\",\"name\":\"key\",\"type\":\"tuple\"},{\"internalType\":\"IdleBalance\",\"name\":\"newIdleBalance\",\"type\":\"bytes32\"}],\"name\":\"hookSetIdleBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"PoolId\",\"name\":\"poolId\",\"type\":\"bytes32\"}],\"name\":\"hookletOfPool\",\"outputs\":[{\"internalType\":\"contractIHooklet\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"PoolId\",\"name\":\"poolId\",\"type\":\"bytes32\"}],\"name\":\"idleBalance\",\"outputs\":[{\"internalType\":\"IdleBalance\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"Currency\",\"name\":\"currency0\",\"type\":\"address\"},{\"internalType\":\"Currency\",\"name\":\"currency1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"internalType\":\"contractIHooks\",\"name\":\"hooks\",\"type\":\"address\"}],\"internalType\":\"structPoolKey\",\"name\":\"key\",\"type\":\"tuple\"}],\"name\":\"lockForRebalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bunniSubspace\",\"type\":\"bytes32\"}],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"result\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"ownershipHandoverExpiresAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"PoolId\",\"name\":\"poolId\",\"type\":\"bytes32\"}],\"name\":\"poolBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBunniToken\",\"name\":\"bunniToken\",\"type\":\"address\"}],\"name\":\"poolIdOfBunniToken\",\"outputs\":[{\"internalType\":\"PoolId\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolInitData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"PoolId\",\"name\":\"poolId\",\"type\":\"bytes32\"}],\"name\":\"poolParams\",\"outputs\":[{\"components\":[{\"internalType\":\"contractILiquidityDensityFunction\",\"name\":\"liquidityDensityFunction\",\"type\":\"address\"},{\"internalType\":\"contractIBunniToken\",\"name\":\"bunniToken\",\"type\":\"address\"},{\"internalType\":\"contractIHooklet\",\"name\":\"hooklet\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"twapSecondsAgo\",\"type\":\"uint24\"},{\"internalType\":\"bytes32\",\"name\":\"ldfParams\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"hookParams\",\"type\":\"bytes\"},{\"internalType\":\"contractERC4626\",\"name\":\"vault0\",\"type\":\"address\"},{\"internalType\":\"contractERC4626\",\"name\":\"vault1\",\"type\":\"address\"},{\"internalType\":\"enumLDFType\",\"name\":\"ldfType\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"minRawTokenRatio0\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"targetRawTokenRatio0\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"maxRawTokenRatio0\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"minRawTokenRatio1\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"targetRawTokenRatio1\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"maxRawTokenRatio1\",\"type\":\"uint24\"},{\"internalType\":\"uint8\",\"name\":\"currency0Decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"currency1Decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"vault0Decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"vault1Decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"rawBalance0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rawBalance1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserve0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserve1\",\"type\":\"uint256\"},{\"internalType\":\"IdleBalance\",\"name\":\"idleBalance\",\"type\":\"bytes32\"}],\"internalType\":\"structPoolState\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"PoolId\",\"name\":\"poolId\",\"type\":\"bytes32\"}],\"name\":\"poolState\",\"outputs\":[{\"components\":[{\"internalType\":\"contractILiquidityDensityFunction\",\"name\":\"liquidityDensityFunction\",\"type\":\"address\"},{\"internalType\":\"contractIBunniToken\",\"name\":\"bunniToken\",\"type\":\"address\"},{\"internalType\":\"contractIHooklet\",\"name\":\"hooklet\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"twapSecondsAgo\",\"type\":\"uint24\"},{\"internalType\":\"bytes32\",\"name\":\"ldfParams\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"hookParams\",\"type\":\"bytes\"},{\"internalType\":\"contractERC4626\",\"name\":\"vault0\",\"type\":\"address\"},{\"internalType\":\"contractERC4626\",\"name\":\"vault1\",\"type\":\"address\"},{\"internalType\":\"enumLDFType\",\"name\":\"ldfType\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"minRawTokenRatio0\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"targetRawTokenRatio0\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"maxRawTokenRatio0\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"minRawTokenRatio1\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"targetRawTokenRatio1\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"maxRawTokenRatio1\",\"type\":\"uint24\"},{\"internalType\":\"uint8\",\"name\":\"currency0Decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"currency1Decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"vault0Decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"vault1Decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"rawBalance0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rawBalance1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserve0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserve1\",\"type\":\"uint256\"},{\"internalType\":\"IdleBalance\",\"name\":\"idleBalance\",\"type\":\"bytes32\"}],\"internalType\":\"structPoolState\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"Currency\",\"name\":\"currency0\",\"type\":\"address\"},{\"internalType\":\"Currency\",\"name\":\"currency1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"internalType\":\"contractIHooks\",\"name\":\"hooks\",\"type\":\"address\"}],\"internalType\":\"structPoolKey\",\"name\":\"poolKey\",\"type\":\"tuple\"},{\"internalType\":\"uint200\",\"name\":\"shares\",\"type\":\"uint200\"}],\"internalType\":\"structIBunniHub.QueueWithdrawParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"queueWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBunniHook\",\"name\":\"hook\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"whitelisted\",\"type\":\"bool\"}],\"name\":\"setHookWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"pauseFlags\",\"type\":\"uint8\"}],\"name\":\"setPauseFlags\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"setPauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"unlockCallback\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"Currency\",\"name\":\"currency0\",\"type\":\"address\"},{\"internalType\":\"Currency\",\"name\":\"currency1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"internalType\":\"contractIHooks\",\"name\":\"hooks\",\"type\":\"address\"}],\"internalType\":\"structPoolKey\",\"name\":\"poolKey\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount0Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"useQueuedWithdrawal\",\"type\":\"bool\"}],\"internalType\":\"structIBunniHub.WithdrawParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// Univ4poolmanagerABI is the input ABI used to generate the binding from.
// Deprecated: Use Univ4poolmanagerMetaData.ABI instead.
var Univ4poolmanagerABI = Univ4poolmanagerMetaData.ABI

// Univ4poolmanager is an auto generated Go binding around an Ethereum contract.
type Univ4poolmanager struct {
	Univ4poolmanagerCaller     // Read-only binding to the contract
	Univ4poolmanagerTransactor // Write-only binding to the contract
	Univ4poolmanagerFilterer   // Log filterer for contract events
}

// Univ4poolmanagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type Univ4poolmanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Univ4poolmanagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Univ4poolmanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Univ4poolmanagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Univ4poolmanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Univ4poolmanagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Univ4poolmanagerSession struct {
	Contract     *Univ4poolmanager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Univ4poolmanagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Univ4poolmanagerCallerSession struct {
	Contract *Univ4poolmanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// Univ4poolmanagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Univ4poolmanagerTransactorSession struct {
	Contract     *Univ4poolmanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// Univ4poolmanagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type Univ4poolmanagerRaw struct {
	Contract *Univ4poolmanager // Generic contract binding to access the raw methods on
}

// Univ4poolmanagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Univ4poolmanagerCallerRaw struct {
	Contract *Univ4poolmanagerCaller // Generic read-only contract binding to access the raw methods on
}

// Univ4poolmanagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Univ4poolmanagerTransactorRaw struct {
	Contract *Univ4poolmanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniv4poolmanager creates a new instance of Univ4poolmanager, bound to a specific deployed contract.
func NewUniv4poolmanager(address common.Address, backend bind.ContractBackend) (*Univ4poolmanager, error) {
	contract, err := bindUniv4poolmanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Univ4poolmanager{Univ4poolmanagerCaller: Univ4poolmanagerCaller{contract: contract}, Univ4poolmanagerTransactor: Univ4poolmanagerTransactor{contract: contract}, Univ4poolmanagerFilterer: Univ4poolmanagerFilterer{contract: contract}}, nil
}

// NewUniv4poolmanagerCaller creates a new read-only instance of Univ4poolmanager, bound to a specific deployed contract.
func NewUniv4poolmanagerCaller(address common.Address, caller bind.ContractCaller) (*Univ4poolmanagerCaller, error) {
	contract, err := bindUniv4poolmanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Univ4poolmanagerCaller{contract: contract}, nil
}

// NewUniv4poolmanagerTransactor creates a new write-only instance of Univ4poolmanager, bound to a specific deployed contract.
func NewUniv4poolmanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*Univ4poolmanagerTransactor, error) {
	contract, err := bindUniv4poolmanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Univ4poolmanagerTransactor{contract: contract}, nil
}

// NewUniv4poolmanagerFilterer creates a new log filterer instance of Univ4poolmanager, bound to a specific deployed contract.
func NewUniv4poolmanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*Univ4poolmanagerFilterer, error) {
	contract, err := bindUniv4poolmanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Univ4poolmanagerFilterer{contract: contract}, nil
}

// bindUniv4poolmanager binds a generic wrapper to an already deployed contract.
func bindUniv4poolmanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Univ4poolmanagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Univ4poolmanager *Univ4poolmanagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Univ4poolmanager.Contract.Univ4poolmanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Univ4poolmanager *Univ4poolmanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.Univ4poolmanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Univ4poolmanager *Univ4poolmanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.Univ4poolmanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Univ4poolmanager *Univ4poolmanagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Univ4poolmanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Univ4poolmanager *Univ4poolmanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Univ4poolmanager *Univ4poolmanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.contract.Transact(opts, method, params...)
}

// BunniTokenOfPool is a free data retrieval call binding the contract method 0xa2a56697.
//
// Solidity: function bunniTokenOfPool(bytes32 poolId) view returns(address)
func (_Univ4poolmanager *Univ4poolmanagerCaller) BunniTokenOfPool(opts *bind.CallOpts, poolId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Univ4poolmanager.contract.Call(opts, &out, "bunniTokenOfPool", poolId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BunniTokenOfPool is a free data retrieval call binding the contract method 0xa2a56697.
//
// Solidity: function bunniTokenOfPool(bytes32 poolId) view returns(address)
func (_Univ4poolmanager *Univ4poolmanagerSession) BunniTokenOfPool(poolId [32]byte) (common.Address, error) {
	return _Univ4poolmanager.Contract.BunniTokenOfPool(&_Univ4poolmanager.CallOpts, poolId)
}

// BunniTokenOfPool is a free data retrieval call binding the contract method 0xa2a56697.
//
// Solidity: function bunniTokenOfPool(bytes32 poolId) view returns(address)
func (_Univ4poolmanager *Univ4poolmanagerCallerSession) BunniTokenOfPool(poolId [32]byte) (common.Address, error) {
	return _Univ4poolmanager.Contract.BunniTokenOfPool(&_Univ4poolmanager.CallOpts, poolId)
}

// GetPauseStatus is a free data retrieval call binding the contract method 0x1d9023cb.
//
// Solidity: function getPauseStatus() view returns(uint8 pauseFlags, bool unpauseFuse)
func (_Univ4poolmanager *Univ4poolmanagerCaller) GetPauseStatus(opts *bind.CallOpts) (struct {
	PauseFlags  uint8
	UnpauseFuse bool
}, error) {
	var out []interface{}
	err := _Univ4poolmanager.contract.Call(opts, &out, "getPauseStatus")

	outstruct := new(struct {
		PauseFlags  uint8
		UnpauseFuse bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PauseFlags = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.UnpauseFuse = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetPauseStatus is a free data retrieval call binding the contract method 0x1d9023cb.
//
// Solidity: function getPauseStatus() view returns(uint8 pauseFlags, bool unpauseFuse)
func (_Univ4poolmanager *Univ4poolmanagerSession) GetPauseStatus() (struct {
	PauseFlags  uint8
	UnpauseFuse bool
}, error) {
	return _Univ4poolmanager.Contract.GetPauseStatus(&_Univ4poolmanager.CallOpts)
}

// GetPauseStatus is a free data retrieval call binding the contract method 0x1d9023cb.
//
// Solidity: function getPauseStatus() view returns(uint8 pauseFlags, bool unpauseFuse)
func (_Univ4poolmanager *Univ4poolmanagerCallerSession) GetPauseStatus() (struct {
	PauseFlags  uint8
	UnpauseFuse bool
}, error) {
	return _Univ4poolmanager.Contract.GetPauseStatus(&_Univ4poolmanager.CallOpts)
}

// HookIsWhitelisted is a free data retrieval call binding the contract method 0x8980a32d.
//
// Solidity: function hookIsWhitelisted(address hook) view returns(bool)
func (_Univ4poolmanager *Univ4poolmanagerCaller) HookIsWhitelisted(opts *bind.CallOpts, hook common.Address) (bool, error) {
	var out []interface{}
	err := _Univ4poolmanager.contract.Call(opts, &out, "hookIsWhitelisted", hook)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HookIsWhitelisted is a free data retrieval call binding the contract method 0x8980a32d.
//
// Solidity: function hookIsWhitelisted(address hook) view returns(bool)
func (_Univ4poolmanager *Univ4poolmanagerSession) HookIsWhitelisted(hook common.Address) (bool, error) {
	return _Univ4poolmanager.Contract.HookIsWhitelisted(&_Univ4poolmanager.CallOpts, hook)
}

// HookIsWhitelisted is a free data retrieval call binding the contract method 0x8980a32d.
//
// Solidity: function hookIsWhitelisted(address hook) view returns(bool)
func (_Univ4poolmanager *Univ4poolmanagerCallerSession) HookIsWhitelisted(hook common.Address) (bool, error) {
	return _Univ4poolmanager.Contract.HookIsWhitelisted(&_Univ4poolmanager.CallOpts, hook)
}

// HookParams is a free data retrieval call binding the contract method 0x129f38ea.
//
// Solidity: function hookParams(bytes32 poolId) view returns(bytes)
func (_Univ4poolmanager *Univ4poolmanagerCaller) HookParams(opts *bind.CallOpts, poolId [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Univ4poolmanager.contract.Call(opts, &out, "hookParams", poolId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// HookParams is a free data retrieval call binding the contract method 0x129f38ea.
//
// Solidity: function hookParams(bytes32 poolId) view returns(bytes)
func (_Univ4poolmanager *Univ4poolmanagerSession) HookParams(poolId [32]byte) ([]byte, error) {
	return _Univ4poolmanager.Contract.HookParams(&_Univ4poolmanager.CallOpts, poolId)
}

// HookParams is a free data retrieval call binding the contract method 0x129f38ea.
//
// Solidity: function hookParams(bytes32 poolId) view returns(bytes)
func (_Univ4poolmanager *Univ4poolmanagerCallerSession) HookParams(poolId [32]byte) ([]byte, error) {
	return _Univ4poolmanager.Contract.HookParams(&_Univ4poolmanager.CallOpts, poolId)
}

// HookletOfPool is a free data retrieval call binding the contract method 0x14ff92a6.
//
// Solidity: function hookletOfPool(bytes32 poolId) view returns(address)
func (_Univ4poolmanager *Univ4poolmanagerCaller) HookletOfPool(opts *bind.CallOpts, poolId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Univ4poolmanager.contract.Call(opts, &out, "hookletOfPool", poolId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HookletOfPool is a free data retrieval call binding the contract method 0x14ff92a6.
//
// Solidity: function hookletOfPool(bytes32 poolId) view returns(address)
func (_Univ4poolmanager *Univ4poolmanagerSession) HookletOfPool(poolId [32]byte) (common.Address, error) {
	return _Univ4poolmanager.Contract.HookletOfPool(&_Univ4poolmanager.CallOpts, poolId)
}

// HookletOfPool is a free data retrieval call binding the contract method 0x14ff92a6.
//
// Solidity: function hookletOfPool(bytes32 poolId) view returns(address)
func (_Univ4poolmanager *Univ4poolmanagerCallerSession) HookletOfPool(poolId [32]byte) (common.Address, error) {
	return _Univ4poolmanager.Contract.HookletOfPool(&_Univ4poolmanager.CallOpts, poolId)
}

// IdleBalance is a free data retrieval call binding the contract method 0x88dd6e53.
//
// Solidity: function idleBalance(bytes32 poolId) view returns(bytes32)
func (_Univ4poolmanager *Univ4poolmanagerCaller) IdleBalance(opts *bind.CallOpts, poolId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Univ4poolmanager.contract.Call(opts, &out, "idleBalance", poolId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// IdleBalance is a free data retrieval call binding the contract method 0x88dd6e53.
//
// Solidity: function idleBalance(bytes32 poolId) view returns(bytes32)
func (_Univ4poolmanager *Univ4poolmanagerSession) IdleBalance(poolId [32]byte) ([32]byte, error) {
	return _Univ4poolmanager.Contract.IdleBalance(&_Univ4poolmanager.CallOpts, poolId)
}

// IdleBalance is a free data retrieval call binding the contract method 0x88dd6e53.
//
// Solidity: function idleBalance(bytes32 poolId) view returns(bytes32)
func (_Univ4poolmanager *Univ4poolmanagerCallerSession) IdleBalance(poolId [32]byte) ([32]byte, error) {
	return _Univ4poolmanager.Contract.IdleBalance(&_Univ4poolmanager.CallOpts, poolId)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address guy) view returns(bool)
func (_Univ4poolmanager *Univ4poolmanagerCaller) IsPauser(opts *bind.CallOpts, guy common.Address) (bool, error) {
	var out []interface{}
	err := _Univ4poolmanager.contract.Call(opts, &out, "isPauser", guy)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address guy) view returns(bool)
func (_Univ4poolmanager *Univ4poolmanagerSession) IsPauser(guy common.Address) (bool, error) {
	return _Univ4poolmanager.Contract.IsPauser(&_Univ4poolmanager.CallOpts, guy)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address guy) view returns(bool)
func (_Univ4poolmanager *Univ4poolmanagerCallerSession) IsPauser(guy common.Address) (bool, error) {
	return _Univ4poolmanager.Contract.IsPauser(&_Univ4poolmanager.CallOpts, guy)
}

// Nonce is a free data retrieval call binding the contract method 0x905da30f.
//
// Solidity: function nonce(bytes32 bunniSubspace) view returns(uint24)
func (_Univ4poolmanager *Univ4poolmanagerCaller) Nonce(opts *bind.CallOpts, bunniSubspace [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Univ4poolmanager.contract.Call(opts, &out, "nonce", bunniSubspace)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0x905da30f.
//
// Solidity: function nonce(bytes32 bunniSubspace) view returns(uint24)
func (_Univ4poolmanager *Univ4poolmanagerSession) Nonce(bunniSubspace [32]byte) (*big.Int, error) {
	return _Univ4poolmanager.Contract.Nonce(&_Univ4poolmanager.CallOpts, bunniSubspace)
}

// Nonce is a free data retrieval call binding the contract method 0x905da30f.
//
// Solidity: function nonce(bytes32 bunniSubspace) view returns(uint24)
func (_Univ4poolmanager *Univ4poolmanagerCallerSession) Nonce(bunniSubspace [32]byte) (*big.Int, error) {
	return _Univ4poolmanager.Contract.Nonce(&_Univ4poolmanager.CallOpts, bunniSubspace)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_Univ4poolmanager *Univ4poolmanagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Univ4poolmanager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_Univ4poolmanager *Univ4poolmanagerSession) Owner() (common.Address, error) {
	return _Univ4poolmanager.Contract.Owner(&_Univ4poolmanager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_Univ4poolmanager *Univ4poolmanagerCallerSession) Owner() (common.Address, error) {
	return _Univ4poolmanager.Contract.Owner(&_Univ4poolmanager.CallOpts)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_Univ4poolmanager *Univ4poolmanagerCaller) OwnershipHandoverExpiresAt(opts *bind.CallOpts, pendingOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Univ4poolmanager.contract.Call(opts, &out, "ownershipHandoverExpiresAt", pendingOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_Univ4poolmanager *Univ4poolmanagerSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _Univ4poolmanager.Contract.OwnershipHandoverExpiresAt(&_Univ4poolmanager.CallOpts, pendingOwner)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_Univ4poolmanager *Univ4poolmanagerCallerSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _Univ4poolmanager.Contract.OwnershipHandoverExpiresAt(&_Univ4poolmanager.CallOpts, pendingOwner)
}

// PoolBalances is a free data retrieval call binding the contract method 0x809b1f38.
//
// Solidity: function poolBalances(bytes32 poolId) view returns(uint256 balance0, uint256 balance1)
func (_Univ4poolmanager *Univ4poolmanagerCaller) PoolBalances(opts *bind.CallOpts, poolId [32]byte) (struct {
	Balance0 *big.Int
	Balance1 *big.Int
}, error) {
	var out []interface{}
	err := _Univ4poolmanager.contract.Call(opts, &out, "poolBalances", poolId)

	outstruct := new(struct {
		Balance0 *big.Int
		Balance1 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Balance0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Balance1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PoolBalances is a free data retrieval call binding the contract method 0x809b1f38.
//
// Solidity: function poolBalances(bytes32 poolId) view returns(uint256 balance0, uint256 balance1)
func (_Univ4poolmanager *Univ4poolmanagerSession) PoolBalances(poolId [32]byte) (struct {
	Balance0 *big.Int
	Balance1 *big.Int
}, error) {
	return _Univ4poolmanager.Contract.PoolBalances(&_Univ4poolmanager.CallOpts, poolId)
}

// PoolBalances is a free data retrieval call binding the contract method 0x809b1f38.
//
// Solidity: function poolBalances(bytes32 poolId) view returns(uint256 balance0, uint256 balance1)
func (_Univ4poolmanager *Univ4poolmanagerCallerSession) PoolBalances(poolId [32]byte) (struct {
	Balance0 *big.Int
	Balance1 *big.Int
}, error) {
	return _Univ4poolmanager.Contract.PoolBalances(&_Univ4poolmanager.CallOpts, poolId)
}

// PoolIdOfBunniToken is a free data retrieval call binding the contract method 0x7676cce0.
//
// Solidity: function poolIdOfBunniToken(address bunniToken) view returns(bytes32)
func (_Univ4poolmanager *Univ4poolmanagerCaller) PoolIdOfBunniToken(opts *bind.CallOpts, bunniToken common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Univ4poolmanager.contract.Call(opts, &out, "poolIdOfBunniToken", bunniToken)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PoolIdOfBunniToken is a free data retrieval call binding the contract method 0x7676cce0.
//
// Solidity: function poolIdOfBunniToken(address bunniToken) view returns(bytes32)
func (_Univ4poolmanager *Univ4poolmanagerSession) PoolIdOfBunniToken(bunniToken common.Address) ([32]byte, error) {
	return _Univ4poolmanager.Contract.PoolIdOfBunniToken(&_Univ4poolmanager.CallOpts, bunniToken)
}

// PoolIdOfBunniToken is a free data retrieval call binding the contract method 0x7676cce0.
//
// Solidity: function poolIdOfBunniToken(address bunniToken) view returns(bytes32)
func (_Univ4poolmanager *Univ4poolmanagerCallerSession) PoolIdOfBunniToken(bunniToken common.Address) ([32]byte, error) {
	return _Univ4poolmanager.Contract.PoolIdOfBunniToken(&_Univ4poolmanager.CallOpts, bunniToken)
}

// PoolInitData is a free data retrieval call binding the contract method 0xf0960848.
//
// Solidity: function poolInitData() view returns(bytes)
func (_Univ4poolmanager *Univ4poolmanagerCaller) PoolInitData(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Univ4poolmanager.contract.Call(opts, &out, "poolInitData")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PoolInitData is a free data retrieval call binding the contract method 0xf0960848.
//
// Solidity: function poolInitData() view returns(bytes)
func (_Univ4poolmanager *Univ4poolmanagerSession) PoolInitData() ([]byte, error) {
	return _Univ4poolmanager.Contract.PoolInitData(&_Univ4poolmanager.CallOpts)
}

// PoolInitData is a free data retrieval call binding the contract method 0xf0960848.
//
// Solidity: function poolInitData() view returns(bytes)
func (_Univ4poolmanager *Univ4poolmanagerCallerSession) PoolInitData() ([]byte, error) {
	return _Univ4poolmanager.Contract.PoolInitData(&_Univ4poolmanager.CallOpts)
}

// PoolParams is a free data retrieval call binding the contract method 0xa0fd3f7e.
//
// Solidity: function poolParams(bytes32 poolId) view returns((address,address,address,uint24,bytes32,bytes,address,address,uint8,uint24,uint24,uint24,uint24,uint24,uint24,uint8,uint8,uint8,uint8,uint256,uint256,uint256,uint256,bytes32))
func (_Univ4poolmanager *Univ4poolmanagerCaller) PoolParams(opts *bind.CallOpts, poolId [32]byte) (PoolState, error) {
	var out []interface{}
	err := _Univ4poolmanager.contract.Call(opts, &out, "poolParams", poolId)

	if err != nil {
		return *new(PoolState), err
	}

	out0 := *abi.ConvertType(out[0], new(PoolState)).(*PoolState)

	return out0, err

}

// PoolParams is a free data retrieval call binding the contract method 0xa0fd3f7e.
//
// Solidity: function poolParams(bytes32 poolId) view returns((address,address,address,uint24,bytes32,bytes,address,address,uint8,uint24,uint24,uint24,uint24,uint24,uint24,uint8,uint8,uint8,uint8,uint256,uint256,uint256,uint256,bytes32))
func (_Univ4poolmanager *Univ4poolmanagerSession) PoolParams(poolId [32]byte) (PoolState, error) {
	return _Univ4poolmanager.Contract.PoolParams(&_Univ4poolmanager.CallOpts, poolId)
}

// PoolParams is a free data retrieval call binding the contract method 0xa0fd3f7e.
//
// Solidity: function poolParams(bytes32 poolId) view returns((address,address,address,uint24,bytes32,bytes,address,address,uint8,uint24,uint24,uint24,uint24,uint24,uint24,uint8,uint8,uint8,uint8,uint256,uint256,uint256,uint256,bytes32))
func (_Univ4poolmanager *Univ4poolmanagerCallerSession) PoolParams(poolId [32]byte) (PoolState, error) {
	return _Univ4poolmanager.Contract.PoolParams(&_Univ4poolmanager.CallOpts, poolId)
}

// PoolState is a free data retrieval call binding the contract method 0xe0b01bac.
//
// Solidity: function poolState(bytes32 poolId) view returns((address,address,address,uint24,bytes32,bytes,address,address,uint8,uint24,uint24,uint24,uint24,uint24,uint24,uint8,uint8,uint8,uint8,uint256,uint256,uint256,uint256,bytes32))
func (_Univ4poolmanager *Univ4poolmanagerCaller) PoolState(opts *bind.CallOpts, poolId [32]byte) (PoolState, error) {
	var out []interface{}
	err := _Univ4poolmanager.contract.Call(opts, &out, "poolState", poolId)

	if err != nil {
		return *new(PoolState), err
	}

	out0 := *abi.ConvertType(out[0], new(PoolState)).(*PoolState)

	return out0, err

}

// PoolState is a free data retrieval call binding the contract method 0xe0b01bac.
//
// Solidity: function poolState(bytes32 poolId) view returns((address,address,address,uint24,bytes32,bytes,address,address,uint8,uint24,uint24,uint24,uint24,uint24,uint24,uint8,uint8,uint8,uint8,uint256,uint256,uint256,uint256,bytes32))
func (_Univ4poolmanager *Univ4poolmanagerSession) PoolState(poolId [32]byte) (PoolState, error) {
	return _Univ4poolmanager.Contract.PoolState(&_Univ4poolmanager.CallOpts, poolId)
}

// PoolState is a free data retrieval call binding the contract method 0xe0b01bac.
//
// Solidity: function poolState(bytes32 poolId) view returns((address,address,address,uint24,bytes32,bytes,address,address,uint8,uint24,uint24,uint24,uint24,uint24,uint24,uint8,uint8,uint8,uint8,uint256,uint256,uint256,uint256,bytes32))
func (_Univ4poolmanager *Univ4poolmanagerCallerSession) PoolState(poolId [32]byte) (PoolState, error) {
	return _Univ4poolmanager.Contract.PoolState(&_Univ4poolmanager.CallOpts, poolId)
}

// BurnPauseFuse is a paid mutator transaction binding the contract method 0x1ed08cb9.
//
// Solidity: function burnPauseFuse() returns()
func (_Univ4poolmanager *Univ4poolmanagerTransactor) BurnPauseFuse(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Univ4poolmanager.contract.Transact(opts, "burnPauseFuse")
}

// BurnPauseFuse is a paid mutator transaction binding the contract method 0x1ed08cb9.
//
// Solidity: function burnPauseFuse() returns()
func (_Univ4poolmanager *Univ4poolmanagerSession) BurnPauseFuse() (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.BurnPauseFuse(&_Univ4poolmanager.TransactOpts)
}

// BurnPauseFuse is a paid mutator transaction binding the contract method 0x1ed08cb9.
//
// Solidity: function burnPauseFuse() returns()
func (_Univ4poolmanager *Univ4poolmanagerTransactorSession) BurnPauseFuse() (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.BurnPauseFuse(&_Univ4poolmanager.TransactOpts)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_Univ4poolmanager *Univ4poolmanagerTransactor) CancelOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Univ4poolmanager.contract.Transact(opts, "cancelOwnershipHandover")
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_Univ4poolmanager *Univ4poolmanagerSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.CancelOwnershipHandover(&_Univ4poolmanager.TransactOpts)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_Univ4poolmanager *Univ4poolmanagerTransactorSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.CancelOwnershipHandover(&_Univ4poolmanager.TransactOpts)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_Univ4poolmanager *Univ4poolmanagerTransactor) CompleteOwnershipHandover(opts *bind.TransactOpts, pendingOwner common.Address) (*types.Transaction, error) {
	return _Univ4poolmanager.contract.Transact(opts, "completeOwnershipHandover", pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_Univ4poolmanager *Univ4poolmanagerSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.CompleteOwnershipHandover(&_Univ4poolmanager.TransactOpts, pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_Univ4poolmanager *Univ4poolmanagerTransactorSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.CompleteOwnershipHandover(&_Univ4poolmanager.TransactOpts, pendingOwner)
}

// DeployBunniToken is a paid mutator transaction binding the contract method 0xe56ba808.
//
// Solidity: function deployBunniToken((address,address,int24,uint24,address,address,uint8,bytes32,address,bytes,address,address,uint24,uint24,uint24,uint24,uint24,uint24,uint160,bytes32,bytes32,address,string,bytes32) params) returns(address token, (address,address,uint24,int24,address) key)
func (_Univ4poolmanager *Univ4poolmanagerTransactor) DeployBunniToken(opts *bind.TransactOpts, params IBunniHubDeployBunniTokenParams) (*types.Transaction, error) {
	return _Univ4poolmanager.contract.Transact(opts, "deployBunniToken", params)
}

// DeployBunniToken is a paid mutator transaction binding the contract method 0xe56ba808.
//
// Solidity: function deployBunniToken((address,address,int24,uint24,address,address,uint8,bytes32,address,bytes,address,address,uint24,uint24,uint24,uint24,uint24,uint24,uint160,bytes32,bytes32,address,string,bytes32) params) returns(address token, (address,address,uint24,int24,address) key)
func (_Univ4poolmanager *Univ4poolmanagerSession) DeployBunniToken(params IBunniHubDeployBunniTokenParams) (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.DeployBunniToken(&_Univ4poolmanager.TransactOpts, params)
}

// DeployBunniToken is a paid mutator transaction binding the contract method 0xe56ba808.
//
// Solidity: function deployBunniToken((address,address,int24,uint24,address,address,uint8,bytes32,address,bytes,address,address,uint24,uint24,uint24,uint24,uint24,uint24,uint160,bytes32,bytes32,address,string,bytes32) params) returns(address token, (address,address,uint24,int24,address) key)
func (_Univ4poolmanager *Univ4poolmanagerTransactorSession) DeployBunniToken(params IBunniHubDeployBunniTokenParams) (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.DeployBunniToken(&_Univ4poolmanager.TransactOpts, params)
}

// Deposit is a paid mutator transaction binding the contract method 0xedb2b1bf.
//
// Solidity: function deposit(((address,address,uint24,int24,address),address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256) params) payable returns(uint256 shares, uint256 amount0, uint256 amount1)
func (_Univ4poolmanager *Univ4poolmanagerTransactor) Deposit(opts *bind.TransactOpts, params IBunniHubDepositParams) (*types.Transaction, error) {
	return _Univ4poolmanager.contract.Transact(opts, "deposit", params)
}

// Deposit is a paid mutator transaction binding the contract method 0xedb2b1bf.
//
// Solidity: function deposit(((address,address,uint24,int24,address),address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256) params) payable returns(uint256 shares, uint256 amount0, uint256 amount1)
func (_Univ4poolmanager *Univ4poolmanagerSession) Deposit(params IBunniHubDepositParams) (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.Deposit(&_Univ4poolmanager.TransactOpts, params)
}

// Deposit is a paid mutator transaction binding the contract method 0xedb2b1bf.
//
// Solidity: function deposit(((address,address,uint24,int24,address),address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256) params) payable returns(uint256 shares, uint256 amount0, uint256 amount1)
func (_Univ4poolmanager *Univ4poolmanagerTransactorSession) Deposit(params IBunniHubDepositParams) (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.Deposit(&_Univ4poolmanager.TransactOpts, params)
}

// HookGive is a paid mutator transaction binding the contract method 0x63f7de6f.
//
// Solidity: function hookGive((address,address,uint24,int24,address) key, bool isCurrency0, uint256 amount) returns()
func (_Univ4poolmanager *Univ4poolmanagerTransactor) HookGive(opts *bind.TransactOpts, key PoolKey, isCurrency0 bool, amount *big.Int) (*types.Transaction, error) {
	return _Univ4poolmanager.contract.Transact(opts, "hookGive", key, isCurrency0, amount)
}

// HookGive is a paid mutator transaction binding the contract method 0x63f7de6f.
//
// Solidity: function hookGive((address,address,uint24,int24,address) key, bool isCurrency0, uint256 amount) returns()
func (_Univ4poolmanager *Univ4poolmanagerSession) HookGive(key PoolKey, isCurrency0 bool, amount *big.Int) (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.HookGive(&_Univ4poolmanager.TransactOpts, key, isCurrency0, amount)
}

// HookGive is a paid mutator transaction binding the contract method 0x63f7de6f.
//
// Solidity: function hookGive((address,address,uint24,int24,address) key, bool isCurrency0, uint256 amount) returns()
func (_Univ4poolmanager *Univ4poolmanagerTransactorSession) HookGive(key PoolKey, isCurrency0 bool, amount *big.Int) (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.HookGive(&_Univ4poolmanager.TransactOpts, key, isCurrency0, amount)
}

// HookHandleSwap is a paid mutator transaction binding the contract method 0x23407c16.
//
// Solidity: function hookHandleSwap((address,address,uint24,int24,address) key, bool zeroForOne, uint256 inputAmount, uint256 outputAmount, bool shouldSurge) returns()
func (_Univ4poolmanager *Univ4poolmanagerTransactor) HookHandleSwap(opts *bind.TransactOpts, key PoolKey, zeroForOne bool, inputAmount *big.Int, outputAmount *big.Int, shouldSurge bool) (*types.Transaction, error) {
	return _Univ4poolmanager.contract.Transact(opts, "hookHandleSwap", key, zeroForOne, inputAmount, outputAmount, shouldSurge)
}

// HookHandleSwap is a paid mutator transaction binding the contract method 0x23407c16.
//
// Solidity: function hookHandleSwap((address,address,uint24,int24,address) key, bool zeroForOne, uint256 inputAmount, uint256 outputAmount, bool shouldSurge) returns()
func (_Univ4poolmanager *Univ4poolmanagerSession) HookHandleSwap(key PoolKey, zeroForOne bool, inputAmount *big.Int, outputAmount *big.Int, shouldSurge bool) (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.HookHandleSwap(&_Univ4poolmanager.TransactOpts, key, zeroForOne, inputAmount, outputAmount, shouldSurge)
}

// HookHandleSwap is a paid mutator transaction binding the contract method 0x23407c16.
//
// Solidity: function hookHandleSwap((address,address,uint24,int24,address) key, bool zeroForOne, uint256 inputAmount, uint256 outputAmount, bool shouldSurge) returns()
func (_Univ4poolmanager *Univ4poolmanagerTransactorSession) HookHandleSwap(key PoolKey, zeroForOne bool, inputAmount *big.Int, outputAmount *big.Int, shouldSurge bool) (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.HookHandleSwap(&_Univ4poolmanager.TransactOpts, key, zeroForOne, inputAmount, outputAmount, shouldSurge)
}

// HookSetIdleBalance is a paid mutator transaction binding the contract method 0xef760335.
//
// Solidity: function hookSetIdleBalance((address,address,uint24,int24,address) key, bytes32 newIdleBalance) returns()
func (_Univ4poolmanager *Univ4poolmanagerTransactor) HookSetIdleBalance(opts *bind.TransactOpts, key PoolKey, newIdleBalance [32]byte) (*types.Transaction, error) {
	return _Univ4poolmanager.contract.Transact(opts, "hookSetIdleBalance", key, newIdleBalance)
}

// HookSetIdleBalance is a paid mutator transaction binding the contract method 0xef760335.
//
// Solidity: function hookSetIdleBalance((address,address,uint24,int24,address) key, bytes32 newIdleBalance) returns()
func (_Univ4poolmanager *Univ4poolmanagerSession) HookSetIdleBalance(key PoolKey, newIdleBalance [32]byte) (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.HookSetIdleBalance(&_Univ4poolmanager.TransactOpts, key, newIdleBalance)
}

// HookSetIdleBalance is a paid mutator transaction binding the contract method 0xef760335.
//
// Solidity: function hookSetIdleBalance((address,address,uint24,int24,address) key, bytes32 newIdleBalance) returns()
func (_Univ4poolmanager *Univ4poolmanagerTransactorSession) HookSetIdleBalance(key PoolKey, newIdleBalance [32]byte) (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.HookSetIdleBalance(&_Univ4poolmanager.TransactOpts, key, newIdleBalance)
}

// LockForRebalance is a paid mutator transaction binding the contract method 0x3fac6506.
//
// Solidity: function lockForRebalance((address,address,uint24,int24,address) key) returns()
func (_Univ4poolmanager *Univ4poolmanagerTransactor) LockForRebalance(opts *bind.TransactOpts, key PoolKey) (*types.Transaction, error) {
	return _Univ4poolmanager.contract.Transact(opts, "lockForRebalance", key)
}

// LockForRebalance is a paid mutator transaction binding the contract method 0x3fac6506.
//
// Solidity: function lockForRebalance((address,address,uint24,int24,address) key) returns()
func (_Univ4poolmanager *Univ4poolmanagerSession) LockForRebalance(key PoolKey) (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.LockForRebalance(&_Univ4poolmanager.TransactOpts, key)
}

// LockForRebalance is a paid mutator transaction binding the contract method 0x3fac6506.
//
// Solidity: function lockForRebalance((address,address,uint24,int24,address) key) returns()
func (_Univ4poolmanager *Univ4poolmanagerTransactorSession) LockForRebalance(key PoolKey) (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.LockForRebalance(&_Univ4poolmanager.TransactOpts, key)
}

// QueueWithdraw is a paid mutator transaction binding the contract method 0x5658d0b4.
//
// Solidity: function queueWithdraw(((address,address,uint24,int24,address),uint200) params) returns()
func (_Univ4poolmanager *Univ4poolmanagerTransactor) QueueWithdraw(opts *bind.TransactOpts, params IBunniHubQueueWithdrawParams) (*types.Transaction, error) {
	return _Univ4poolmanager.contract.Transact(opts, "queueWithdraw", params)
}

// QueueWithdraw is a paid mutator transaction binding the contract method 0x5658d0b4.
//
// Solidity: function queueWithdraw(((address,address,uint24,int24,address),uint200) params) returns()
func (_Univ4poolmanager *Univ4poolmanagerSession) QueueWithdraw(params IBunniHubQueueWithdrawParams) (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.QueueWithdraw(&_Univ4poolmanager.TransactOpts, params)
}

// QueueWithdraw is a paid mutator transaction binding the contract method 0x5658d0b4.
//
// Solidity: function queueWithdraw(((address,address,uint24,int24,address),uint200) params) returns()
func (_Univ4poolmanager *Univ4poolmanagerTransactorSession) QueueWithdraw(params IBunniHubQueueWithdrawParams) (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.QueueWithdraw(&_Univ4poolmanager.TransactOpts, params)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_Univ4poolmanager *Univ4poolmanagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Univ4poolmanager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_Univ4poolmanager *Univ4poolmanagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.RenounceOwnership(&_Univ4poolmanager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_Univ4poolmanager *Univ4poolmanagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.RenounceOwnership(&_Univ4poolmanager.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_Univ4poolmanager *Univ4poolmanagerTransactor) RequestOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Univ4poolmanager.contract.Transact(opts, "requestOwnershipHandover")
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_Univ4poolmanager *Univ4poolmanagerSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.RequestOwnershipHandover(&_Univ4poolmanager.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_Univ4poolmanager *Univ4poolmanagerTransactorSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.RequestOwnershipHandover(&_Univ4poolmanager.TransactOpts)
}

// SetHookWhitelist is a paid mutator transaction binding the contract method 0xce79eb60.
//
// Solidity: function setHookWhitelist(address hook, bool whitelisted) returns()
func (_Univ4poolmanager *Univ4poolmanagerTransactor) SetHookWhitelist(opts *bind.TransactOpts, hook common.Address, whitelisted bool) (*types.Transaction, error) {
	return _Univ4poolmanager.contract.Transact(opts, "setHookWhitelist", hook, whitelisted)
}

// SetHookWhitelist is a paid mutator transaction binding the contract method 0xce79eb60.
//
// Solidity: function setHookWhitelist(address hook, bool whitelisted) returns()
func (_Univ4poolmanager *Univ4poolmanagerSession) SetHookWhitelist(hook common.Address, whitelisted bool) (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.SetHookWhitelist(&_Univ4poolmanager.TransactOpts, hook, whitelisted)
}

// SetHookWhitelist is a paid mutator transaction binding the contract method 0xce79eb60.
//
// Solidity: function setHookWhitelist(address hook, bool whitelisted) returns()
func (_Univ4poolmanager *Univ4poolmanagerTransactorSession) SetHookWhitelist(hook common.Address, whitelisted bool) (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.SetHookWhitelist(&_Univ4poolmanager.TransactOpts, hook, whitelisted)
}

// SetPauseFlags is a paid mutator transaction binding the contract method 0xa56dd053.
//
// Solidity: function setPauseFlags(uint8 pauseFlags) returns()
func (_Univ4poolmanager *Univ4poolmanagerTransactor) SetPauseFlags(opts *bind.TransactOpts, pauseFlags uint8) (*types.Transaction, error) {
	return _Univ4poolmanager.contract.Transact(opts, "setPauseFlags", pauseFlags)
}

// SetPauseFlags is a paid mutator transaction binding the contract method 0xa56dd053.
//
// Solidity: function setPauseFlags(uint8 pauseFlags) returns()
func (_Univ4poolmanager *Univ4poolmanagerSession) SetPauseFlags(pauseFlags uint8) (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.SetPauseFlags(&_Univ4poolmanager.TransactOpts, pauseFlags)
}

// SetPauseFlags is a paid mutator transaction binding the contract method 0xa56dd053.
//
// Solidity: function setPauseFlags(uint8 pauseFlags) returns()
func (_Univ4poolmanager *Univ4poolmanagerTransactorSession) SetPauseFlags(pauseFlags uint8) (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.SetPauseFlags(&_Univ4poolmanager.TransactOpts, pauseFlags)
}

// SetPauser is a paid mutator transaction binding the contract method 0x7180c8ca.
//
// Solidity: function setPauser(address guy, bool status) returns()
func (_Univ4poolmanager *Univ4poolmanagerTransactor) SetPauser(opts *bind.TransactOpts, guy common.Address, status bool) (*types.Transaction, error) {
	return _Univ4poolmanager.contract.Transact(opts, "setPauser", guy, status)
}

// SetPauser is a paid mutator transaction binding the contract method 0x7180c8ca.
//
// Solidity: function setPauser(address guy, bool status) returns()
func (_Univ4poolmanager *Univ4poolmanagerSession) SetPauser(guy common.Address, status bool) (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.SetPauser(&_Univ4poolmanager.TransactOpts, guy, status)
}

// SetPauser is a paid mutator transaction binding the contract method 0x7180c8ca.
//
// Solidity: function setPauser(address guy, bool status) returns()
func (_Univ4poolmanager *Univ4poolmanagerTransactorSession) SetPauser(guy common.Address, status bool) (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.SetPauser(&_Univ4poolmanager.TransactOpts, guy, status)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_Univ4poolmanager *Univ4poolmanagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Univ4poolmanager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_Univ4poolmanager *Univ4poolmanagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.TransferOwnership(&_Univ4poolmanager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_Univ4poolmanager *Univ4poolmanagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.TransferOwnership(&_Univ4poolmanager.TransactOpts, newOwner)
}

// UnlockCallback is a paid mutator transaction binding the contract method 0x91dd7346.
//
// Solidity: function unlockCallback(bytes data) returns(bytes)
func (_Univ4poolmanager *Univ4poolmanagerTransactor) UnlockCallback(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _Univ4poolmanager.contract.Transact(opts, "unlockCallback", data)
}

// UnlockCallback is a paid mutator transaction binding the contract method 0x91dd7346.
//
// Solidity: function unlockCallback(bytes data) returns(bytes)
func (_Univ4poolmanager *Univ4poolmanagerSession) UnlockCallback(data []byte) (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.UnlockCallback(&_Univ4poolmanager.TransactOpts, data)
}

// UnlockCallback is a paid mutator transaction binding the contract method 0x91dd7346.
//
// Solidity: function unlockCallback(bytes data) returns(bytes)
func (_Univ4poolmanager *Univ4poolmanagerTransactorSession) UnlockCallback(data []byte) (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.UnlockCallback(&_Univ4poolmanager.TransactOpts, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5d4a505e.
//
// Solidity: function withdraw(((address,address,uint24,int24,address),address,uint256,uint256,uint256,uint256,bool) params) returns(uint256 amount0, uint256 amount1)
func (_Univ4poolmanager *Univ4poolmanagerTransactor) Withdraw(opts *bind.TransactOpts, params IBunniHubWithdrawParams) (*types.Transaction, error) {
	return _Univ4poolmanager.contract.Transact(opts, "withdraw", params)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5d4a505e.
//
// Solidity: function withdraw(((address,address,uint24,int24,address),address,uint256,uint256,uint256,uint256,bool) params) returns(uint256 amount0, uint256 amount1)
func (_Univ4poolmanager *Univ4poolmanagerSession) Withdraw(params IBunniHubWithdrawParams) (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.Withdraw(&_Univ4poolmanager.TransactOpts, params)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5d4a505e.
//
// Solidity: function withdraw(((address,address,uint24,int24,address),address,uint256,uint256,uint256,uint256,bool) params) returns(uint256 amount0, uint256 amount1)
func (_Univ4poolmanager *Univ4poolmanagerTransactorSession) Withdraw(params IBunniHubWithdrawParams) (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.Withdraw(&_Univ4poolmanager.TransactOpts, params)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Univ4poolmanager *Univ4poolmanagerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Univ4poolmanager.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Univ4poolmanager *Univ4poolmanagerSession) Receive() (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.Receive(&_Univ4poolmanager.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Univ4poolmanager *Univ4poolmanagerTransactorSession) Receive() (*types.Transaction, error) {
	return _Univ4poolmanager.Contract.Receive(&_Univ4poolmanager.TransactOpts)
}

// Univ4poolmanagerBurnPauseFuseIterator is returned from FilterBurnPauseFuse and is used to iterate over the raw logs and unpacked data for BurnPauseFuse events raised by the Univ4poolmanager contract.
type Univ4poolmanagerBurnPauseFuseIterator struct {
	Event *Univ4poolmanagerBurnPauseFuse // Event containing the contract specifics and raw log

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
func (it *Univ4poolmanagerBurnPauseFuseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Univ4poolmanagerBurnPauseFuse)
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
		it.Event = new(Univ4poolmanagerBurnPauseFuse)
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
func (it *Univ4poolmanagerBurnPauseFuseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Univ4poolmanagerBurnPauseFuseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Univ4poolmanagerBurnPauseFuse represents a BurnPauseFuse event raised by the Univ4poolmanager contract.
type Univ4poolmanagerBurnPauseFuse struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBurnPauseFuse is a free log retrieval operation binding the contract event 0xa4058ba547bb832da5ae671cc4d748c09c98c85226ec320325d641a1a3d64adf.
//
// Solidity: event BurnPauseFuse()
func (_Univ4poolmanager *Univ4poolmanagerFilterer) FilterBurnPauseFuse(opts *bind.FilterOpts) (*Univ4poolmanagerBurnPauseFuseIterator, error) {

	logs, sub, err := _Univ4poolmanager.contract.FilterLogs(opts, "BurnPauseFuse")
	if err != nil {
		return nil, err
	}
	return &Univ4poolmanagerBurnPauseFuseIterator{contract: _Univ4poolmanager.contract, event: "BurnPauseFuse", logs: logs, sub: sub}, nil
}

// WatchBurnPauseFuse is a free log subscription operation binding the contract event 0xa4058ba547bb832da5ae671cc4d748c09c98c85226ec320325d641a1a3d64adf.
//
// Solidity: event BurnPauseFuse()
func (_Univ4poolmanager *Univ4poolmanagerFilterer) WatchBurnPauseFuse(opts *bind.WatchOpts, sink chan<- *Univ4poolmanagerBurnPauseFuse) (event.Subscription, error) {

	logs, sub, err := _Univ4poolmanager.contract.WatchLogs(opts, "BurnPauseFuse")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Univ4poolmanagerBurnPauseFuse)
				if err := _Univ4poolmanager.contract.UnpackLog(event, "BurnPauseFuse", log); err != nil {
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

// ParseBurnPauseFuse is a log parse operation binding the contract event 0xa4058ba547bb832da5ae671cc4d748c09c98c85226ec320325d641a1a3d64adf.
//
// Solidity: event BurnPauseFuse()
func (_Univ4poolmanager *Univ4poolmanagerFilterer) ParseBurnPauseFuse(log types.Log) (*Univ4poolmanagerBurnPauseFuse, error) {
	event := new(Univ4poolmanagerBurnPauseFuse)
	if err := _Univ4poolmanager.contract.UnpackLog(event, "BurnPauseFuse", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Univ4poolmanagerDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Univ4poolmanager contract.
type Univ4poolmanagerDepositIterator struct {
	Event *Univ4poolmanagerDeposit // Event containing the contract specifics and raw log

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
func (it *Univ4poolmanagerDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Univ4poolmanagerDeposit)
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
		it.Event = new(Univ4poolmanagerDeposit)
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
func (it *Univ4poolmanagerDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Univ4poolmanagerDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Univ4poolmanagerDeposit represents a Deposit event raised by the Univ4poolmanager contract.
type Univ4poolmanagerDeposit struct {
	Sender    common.Address
	Recipient common.Address
	PoolId    [32]byte
	Amount0   *big.Int
	Amount1   *big.Int
	Shares    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xb18066d48ef2004e3dcc96ec09f8e738f9e8692565ae7108c2b593f8199af466.
//
// Solidity: event Deposit(address indexed sender, address indexed recipient, bytes32 indexed poolId, uint256 amount0, uint256 amount1, uint256 shares)
func (_Univ4poolmanager *Univ4poolmanagerFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address, poolId [][32]byte) (*Univ4poolmanagerDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Univ4poolmanager.contract.FilterLogs(opts, "Deposit", senderRule, recipientRule, poolIdRule)
	if err != nil {
		return nil, err
	}
	return &Univ4poolmanagerDepositIterator{contract: _Univ4poolmanager.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xb18066d48ef2004e3dcc96ec09f8e738f9e8692565ae7108c2b593f8199af466.
//
// Solidity: event Deposit(address indexed sender, address indexed recipient, bytes32 indexed poolId, uint256 amount0, uint256 amount1, uint256 shares)
func (_Univ4poolmanager *Univ4poolmanagerFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *Univ4poolmanagerDeposit, sender []common.Address, recipient []common.Address, poolId [][32]byte) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Univ4poolmanager.contract.WatchLogs(opts, "Deposit", senderRule, recipientRule, poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Univ4poolmanagerDeposit)
				if err := _Univ4poolmanager.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xb18066d48ef2004e3dcc96ec09f8e738f9e8692565ae7108c2b593f8199af466.
//
// Solidity: event Deposit(address indexed sender, address indexed recipient, bytes32 indexed poolId, uint256 amount0, uint256 amount1, uint256 shares)
func (_Univ4poolmanager *Univ4poolmanagerFilterer) ParseDeposit(log types.Log) (*Univ4poolmanagerDeposit, error) {
	event := new(Univ4poolmanagerDeposit)
	if err := _Univ4poolmanager.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Univ4poolmanagerNewBunniIterator is returned from FilterNewBunni and is used to iterate over the raw logs and unpacked data for NewBunni events raised by the Univ4poolmanager contract.
type Univ4poolmanagerNewBunniIterator struct {
	Event *Univ4poolmanagerNewBunni // Event containing the contract specifics and raw log

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
func (it *Univ4poolmanagerNewBunniIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Univ4poolmanagerNewBunni)
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
		it.Event = new(Univ4poolmanagerNewBunni)
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
func (it *Univ4poolmanagerNewBunniIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Univ4poolmanagerNewBunniIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Univ4poolmanagerNewBunni represents a NewBunni event raised by the Univ4poolmanager contract.
type Univ4poolmanagerNewBunni struct {
	BunniToken common.Address
	PoolId     [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewBunni is a free log retrieval operation binding the contract event 0x3ba5df143a8e4c83b7cb37037a87ae0cfb0f8c3a784d4a10e0b329d5706dce1a.
//
// Solidity: event NewBunni(address indexed bunniToken, bytes32 indexed poolId)
func (_Univ4poolmanager *Univ4poolmanagerFilterer) FilterNewBunni(opts *bind.FilterOpts, bunniToken []common.Address, poolId [][32]byte) (*Univ4poolmanagerNewBunniIterator, error) {

	var bunniTokenRule []interface{}
	for _, bunniTokenItem := range bunniToken {
		bunniTokenRule = append(bunniTokenRule, bunniTokenItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Univ4poolmanager.contract.FilterLogs(opts, "NewBunni", bunniTokenRule, poolIdRule)
	if err != nil {
		return nil, err
	}
	return &Univ4poolmanagerNewBunniIterator{contract: _Univ4poolmanager.contract, event: "NewBunni", logs: logs, sub: sub}, nil
}

// WatchNewBunni is a free log subscription operation binding the contract event 0x3ba5df143a8e4c83b7cb37037a87ae0cfb0f8c3a784d4a10e0b329d5706dce1a.
//
// Solidity: event NewBunni(address indexed bunniToken, bytes32 indexed poolId)
func (_Univ4poolmanager *Univ4poolmanagerFilterer) WatchNewBunni(opts *bind.WatchOpts, sink chan<- *Univ4poolmanagerNewBunni, bunniToken []common.Address, poolId [][32]byte) (event.Subscription, error) {

	var bunniTokenRule []interface{}
	for _, bunniTokenItem := range bunniToken {
		bunniTokenRule = append(bunniTokenRule, bunniTokenItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Univ4poolmanager.contract.WatchLogs(opts, "NewBunni", bunniTokenRule, poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Univ4poolmanagerNewBunni)
				if err := _Univ4poolmanager.contract.UnpackLog(event, "NewBunni", log); err != nil {
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

// ParseNewBunni is a log parse operation binding the contract event 0x3ba5df143a8e4c83b7cb37037a87ae0cfb0f8c3a784d4a10e0b329d5706dce1a.
//
// Solidity: event NewBunni(address indexed bunniToken, bytes32 indexed poolId)
func (_Univ4poolmanager *Univ4poolmanagerFilterer) ParseNewBunni(log types.Log) (*Univ4poolmanagerNewBunni, error) {
	event := new(Univ4poolmanagerNewBunni)
	if err := _Univ4poolmanager.contract.UnpackLog(event, "NewBunni", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Univ4poolmanagerOwnershipHandoverCanceledIterator is returned from FilterOwnershipHandoverCanceled and is used to iterate over the raw logs and unpacked data for OwnershipHandoverCanceled events raised by the Univ4poolmanager contract.
type Univ4poolmanagerOwnershipHandoverCanceledIterator struct {
	Event *Univ4poolmanagerOwnershipHandoverCanceled // Event containing the contract specifics and raw log

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
func (it *Univ4poolmanagerOwnershipHandoverCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Univ4poolmanagerOwnershipHandoverCanceled)
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
		it.Event = new(Univ4poolmanagerOwnershipHandoverCanceled)
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
func (it *Univ4poolmanagerOwnershipHandoverCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Univ4poolmanagerOwnershipHandoverCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Univ4poolmanagerOwnershipHandoverCanceled represents a OwnershipHandoverCanceled event raised by the Univ4poolmanager contract.
type Univ4poolmanagerOwnershipHandoverCanceled struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverCanceled is a free log retrieval operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_Univ4poolmanager *Univ4poolmanagerFilterer) FilterOwnershipHandoverCanceled(opts *bind.FilterOpts, pendingOwner []common.Address) (*Univ4poolmanagerOwnershipHandoverCanceledIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Univ4poolmanager.contract.FilterLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Univ4poolmanagerOwnershipHandoverCanceledIterator{contract: _Univ4poolmanager.contract, event: "OwnershipHandoverCanceled", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverCanceled is a free log subscription operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_Univ4poolmanager *Univ4poolmanagerFilterer) WatchOwnershipHandoverCanceled(opts *bind.WatchOpts, sink chan<- *Univ4poolmanagerOwnershipHandoverCanceled, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Univ4poolmanager.contract.WatchLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Univ4poolmanagerOwnershipHandoverCanceled)
				if err := _Univ4poolmanager.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
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

// ParseOwnershipHandoverCanceled is a log parse operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_Univ4poolmanager *Univ4poolmanagerFilterer) ParseOwnershipHandoverCanceled(log types.Log) (*Univ4poolmanagerOwnershipHandoverCanceled, error) {
	event := new(Univ4poolmanagerOwnershipHandoverCanceled)
	if err := _Univ4poolmanager.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Univ4poolmanagerOwnershipHandoverRequestedIterator is returned from FilterOwnershipHandoverRequested and is used to iterate over the raw logs and unpacked data for OwnershipHandoverRequested events raised by the Univ4poolmanager contract.
type Univ4poolmanagerOwnershipHandoverRequestedIterator struct {
	Event *Univ4poolmanagerOwnershipHandoverRequested // Event containing the contract specifics and raw log

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
func (it *Univ4poolmanagerOwnershipHandoverRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Univ4poolmanagerOwnershipHandoverRequested)
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
		it.Event = new(Univ4poolmanagerOwnershipHandoverRequested)
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
func (it *Univ4poolmanagerOwnershipHandoverRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Univ4poolmanagerOwnershipHandoverRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Univ4poolmanagerOwnershipHandoverRequested represents a OwnershipHandoverRequested event raised by the Univ4poolmanager contract.
type Univ4poolmanagerOwnershipHandoverRequested struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverRequested is a free log retrieval operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_Univ4poolmanager *Univ4poolmanagerFilterer) FilterOwnershipHandoverRequested(opts *bind.FilterOpts, pendingOwner []common.Address) (*Univ4poolmanagerOwnershipHandoverRequestedIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Univ4poolmanager.contract.FilterLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Univ4poolmanagerOwnershipHandoverRequestedIterator{contract: _Univ4poolmanager.contract, event: "OwnershipHandoverRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverRequested is a free log subscription operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_Univ4poolmanager *Univ4poolmanagerFilterer) WatchOwnershipHandoverRequested(opts *bind.WatchOpts, sink chan<- *Univ4poolmanagerOwnershipHandoverRequested, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Univ4poolmanager.contract.WatchLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Univ4poolmanagerOwnershipHandoverRequested)
				if err := _Univ4poolmanager.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
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

// ParseOwnershipHandoverRequested is a log parse operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_Univ4poolmanager *Univ4poolmanagerFilterer) ParseOwnershipHandoverRequested(log types.Log) (*Univ4poolmanagerOwnershipHandoverRequested, error) {
	event := new(Univ4poolmanagerOwnershipHandoverRequested)
	if err := _Univ4poolmanager.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Univ4poolmanagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Univ4poolmanager contract.
type Univ4poolmanagerOwnershipTransferredIterator struct {
	Event *Univ4poolmanagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Univ4poolmanagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Univ4poolmanagerOwnershipTransferred)
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
		it.Event = new(Univ4poolmanagerOwnershipTransferred)
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
func (it *Univ4poolmanagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Univ4poolmanagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Univ4poolmanagerOwnershipTransferred represents a OwnershipTransferred event raised by the Univ4poolmanager contract.
type Univ4poolmanagerOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_Univ4poolmanager *Univ4poolmanagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*Univ4poolmanagerOwnershipTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Univ4poolmanager.contract.FilterLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Univ4poolmanagerOwnershipTransferredIterator{contract: _Univ4poolmanager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_Univ4poolmanager *Univ4poolmanagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Univ4poolmanagerOwnershipTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Univ4poolmanager.contract.WatchLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Univ4poolmanagerOwnershipTransferred)
				if err := _Univ4poolmanager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_Univ4poolmanager *Univ4poolmanagerFilterer) ParseOwnershipTransferred(log types.Log) (*Univ4poolmanagerOwnershipTransferred, error) {
	event := new(Univ4poolmanagerOwnershipTransferred)
	if err := _Univ4poolmanager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Univ4poolmanagerQueueWithdrawIterator is returned from FilterQueueWithdraw and is used to iterate over the raw logs and unpacked data for QueueWithdraw events raised by the Univ4poolmanager contract.
type Univ4poolmanagerQueueWithdrawIterator struct {
	Event *Univ4poolmanagerQueueWithdraw // Event containing the contract specifics and raw log

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
func (it *Univ4poolmanagerQueueWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Univ4poolmanagerQueueWithdraw)
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
		it.Event = new(Univ4poolmanagerQueueWithdraw)
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
func (it *Univ4poolmanagerQueueWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Univ4poolmanagerQueueWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Univ4poolmanagerQueueWithdraw represents a QueueWithdraw event raised by the Univ4poolmanager contract.
type Univ4poolmanagerQueueWithdraw struct {
	Sender common.Address
	PoolId [32]byte
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterQueueWithdraw is a free log retrieval operation binding the contract event 0x0ee885e060478e5bf89befb890ae82fdcc47aa2a9c8e4d668fcce310318d28a1.
//
// Solidity: event QueueWithdraw(address indexed sender, bytes32 indexed poolId, uint256 shares)
func (_Univ4poolmanager *Univ4poolmanagerFilterer) FilterQueueWithdraw(opts *bind.FilterOpts, sender []common.Address, poolId [][32]byte) (*Univ4poolmanagerQueueWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Univ4poolmanager.contract.FilterLogs(opts, "QueueWithdraw", senderRule, poolIdRule)
	if err != nil {
		return nil, err
	}
	return &Univ4poolmanagerQueueWithdrawIterator{contract: _Univ4poolmanager.contract, event: "QueueWithdraw", logs: logs, sub: sub}, nil
}

// WatchQueueWithdraw is a free log subscription operation binding the contract event 0x0ee885e060478e5bf89befb890ae82fdcc47aa2a9c8e4d668fcce310318d28a1.
//
// Solidity: event QueueWithdraw(address indexed sender, bytes32 indexed poolId, uint256 shares)
func (_Univ4poolmanager *Univ4poolmanagerFilterer) WatchQueueWithdraw(opts *bind.WatchOpts, sink chan<- *Univ4poolmanagerQueueWithdraw, sender []common.Address, poolId [][32]byte) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Univ4poolmanager.contract.WatchLogs(opts, "QueueWithdraw", senderRule, poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Univ4poolmanagerQueueWithdraw)
				if err := _Univ4poolmanager.contract.UnpackLog(event, "QueueWithdraw", log); err != nil {
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

// ParseQueueWithdraw is a log parse operation binding the contract event 0x0ee885e060478e5bf89befb890ae82fdcc47aa2a9c8e4d668fcce310318d28a1.
//
// Solidity: event QueueWithdraw(address indexed sender, bytes32 indexed poolId, uint256 shares)
func (_Univ4poolmanager *Univ4poolmanagerFilterer) ParseQueueWithdraw(log types.Log) (*Univ4poolmanagerQueueWithdraw, error) {
	event := new(Univ4poolmanagerQueueWithdraw)
	if err := _Univ4poolmanager.contract.UnpackLog(event, "QueueWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Univ4poolmanagerSetHookWhitelistIterator is returned from FilterSetHookWhitelist and is used to iterate over the raw logs and unpacked data for SetHookWhitelist events raised by the Univ4poolmanager contract.
type Univ4poolmanagerSetHookWhitelistIterator struct {
	Event *Univ4poolmanagerSetHookWhitelist // Event containing the contract specifics and raw log

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
func (it *Univ4poolmanagerSetHookWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Univ4poolmanagerSetHookWhitelist)
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
		it.Event = new(Univ4poolmanagerSetHookWhitelist)
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
func (it *Univ4poolmanagerSetHookWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Univ4poolmanagerSetHookWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Univ4poolmanagerSetHookWhitelist represents a SetHookWhitelist event raised by the Univ4poolmanager contract.
type Univ4poolmanagerSetHookWhitelist struct {
	Hook        common.Address
	Whitelisted bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetHookWhitelist is a free log retrieval operation binding the contract event 0x2b15baa03b2135b36006038a39516fa43ec5f0062ca42d22e3a29e9ff1c051ef.
//
// Solidity: event SetHookWhitelist(address indexed hook, bool indexed whitelisted)
func (_Univ4poolmanager *Univ4poolmanagerFilterer) FilterSetHookWhitelist(opts *bind.FilterOpts, hook []common.Address, whitelisted []bool) (*Univ4poolmanagerSetHookWhitelistIterator, error) {

	var hookRule []interface{}
	for _, hookItem := range hook {
		hookRule = append(hookRule, hookItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _Univ4poolmanager.contract.FilterLogs(opts, "SetHookWhitelist", hookRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return &Univ4poolmanagerSetHookWhitelistIterator{contract: _Univ4poolmanager.contract, event: "SetHookWhitelist", logs: logs, sub: sub}, nil
}

// WatchSetHookWhitelist is a free log subscription operation binding the contract event 0x2b15baa03b2135b36006038a39516fa43ec5f0062ca42d22e3a29e9ff1c051ef.
//
// Solidity: event SetHookWhitelist(address indexed hook, bool indexed whitelisted)
func (_Univ4poolmanager *Univ4poolmanagerFilterer) WatchSetHookWhitelist(opts *bind.WatchOpts, sink chan<- *Univ4poolmanagerSetHookWhitelist, hook []common.Address, whitelisted []bool) (event.Subscription, error) {

	var hookRule []interface{}
	for _, hookItem := range hook {
		hookRule = append(hookRule, hookItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _Univ4poolmanager.contract.WatchLogs(opts, "SetHookWhitelist", hookRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Univ4poolmanagerSetHookWhitelist)
				if err := _Univ4poolmanager.contract.UnpackLog(event, "SetHookWhitelist", log); err != nil {
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

// ParseSetHookWhitelist is a log parse operation binding the contract event 0x2b15baa03b2135b36006038a39516fa43ec5f0062ca42d22e3a29e9ff1c051ef.
//
// Solidity: event SetHookWhitelist(address indexed hook, bool indexed whitelisted)
func (_Univ4poolmanager *Univ4poolmanagerFilterer) ParseSetHookWhitelist(log types.Log) (*Univ4poolmanagerSetHookWhitelist, error) {
	event := new(Univ4poolmanagerSetHookWhitelist)
	if err := _Univ4poolmanager.contract.UnpackLog(event, "SetHookWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Univ4poolmanagerSetPauseFlagsIterator is returned from FilterSetPauseFlags and is used to iterate over the raw logs and unpacked data for SetPauseFlags events raised by the Univ4poolmanager contract.
type Univ4poolmanagerSetPauseFlagsIterator struct {
	Event *Univ4poolmanagerSetPauseFlags // Event containing the contract specifics and raw log

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
func (it *Univ4poolmanagerSetPauseFlagsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Univ4poolmanagerSetPauseFlags)
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
		it.Event = new(Univ4poolmanagerSetPauseFlags)
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
func (it *Univ4poolmanagerSetPauseFlagsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Univ4poolmanagerSetPauseFlagsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Univ4poolmanagerSetPauseFlags represents a SetPauseFlags event raised by the Univ4poolmanager contract.
type Univ4poolmanagerSetPauseFlags struct {
	PauseFlags uint8
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetPauseFlags is a free log retrieval operation binding the contract event 0x3021cc5514f1ea312648df4d3e6c9cf9c5bd12c429f0849d4c903af7010c6afa.
//
// Solidity: event SetPauseFlags(uint8 indexed pauseFlags)
func (_Univ4poolmanager *Univ4poolmanagerFilterer) FilterSetPauseFlags(opts *bind.FilterOpts, pauseFlags []uint8) (*Univ4poolmanagerSetPauseFlagsIterator, error) {

	var pauseFlagsRule []interface{}
	for _, pauseFlagsItem := range pauseFlags {
		pauseFlagsRule = append(pauseFlagsRule, pauseFlagsItem)
	}

	logs, sub, err := _Univ4poolmanager.contract.FilterLogs(opts, "SetPauseFlags", pauseFlagsRule)
	if err != nil {
		return nil, err
	}
	return &Univ4poolmanagerSetPauseFlagsIterator{contract: _Univ4poolmanager.contract, event: "SetPauseFlags", logs: logs, sub: sub}, nil
}

// WatchSetPauseFlags is a free log subscription operation binding the contract event 0x3021cc5514f1ea312648df4d3e6c9cf9c5bd12c429f0849d4c903af7010c6afa.
//
// Solidity: event SetPauseFlags(uint8 indexed pauseFlags)
func (_Univ4poolmanager *Univ4poolmanagerFilterer) WatchSetPauseFlags(opts *bind.WatchOpts, sink chan<- *Univ4poolmanagerSetPauseFlags, pauseFlags []uint8) (event.Subscription, error) {

	var pauseFlagsRule []interface{}
	for _, pauseFlagsItem := range pauseFlags {
		pauseFlagsRule = append(pauseFlagsRule, pauseFlagsItem)
	}

	logs, sub, err := _Univ4poolmanager.contract.WatchLogs(opts, "SetPauseFlags", pauseFlagsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Univ4poolmanagerSetPauseFlags)
				if err := _Univ4poolmanager.contract.UnpackLog(event, "SetPauseFlags", log); err != nil {
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

// ParseSetPauseFlags is a log parse operation binding the contract event 0x3021cc5514f1ea312648df4d3e6c9cf9c5bd12c429f0849d4c903af7010c6afa.
//
// Solidity: event SetPauseFlags(uint8 indexed pauseFlags)
func (_Univ4poolmanager *Univ4poolmanagerFilterer) ParseSetPauseFlags(log types.Log) (*Univ4poolmanagerSetPauseFlags, error) {
	event := new(Univ4poolmanagerSetPauseFlags)
	if err := _Univ4poolmanager.contract.UnpackLog(event, "SetPauseFlags", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Univ4poolmanagerSetPauserIterator is returned from FilterSetPauser and is used to iterate over the raw logs and unpacked data for SetPauser events raised by the Univ4poolmanager contract.
type Univ4poolmanagerSetPauserIterator struct {
	Event *Univ4poolmanagerSetPauser // Event containing the contract specifics and raw log

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
func (it *Univ4poolmanagerSetPauserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Univ4poolmanagerSetPauser)
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
		it.Event = new(Univ4poolmanagerSetPauser)
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
func (it *Univ4poolmanagerSetPauserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Univ4poolmanagerSetPauserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Univ4poolmanagerSetPauser represents a SetPauser event raised by the Univ4poolmanager contract.
type Univ4poolmanagerSetPauser struct {
	Guy      common.Address
	IsPauser bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetPauser is a free log retrieval operation binding the contract event 0xd34f4aa5f94a385f2fa0ca25e5f01c6f331018f35c3d43a7b8057a86704de3df.
//
// Solidity: event SetPauser(address indexed guy, bool indexed isPauser)
func (_Univ4poolmanager *Univ4poolmanagerFilterer) FilterSetPauser(opts *bind.FilterOpts, guy []common.Address, isPauser []bool) (*Univ4poolmanagerSetPauserIterator, error) {

	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}
	var isPauserRule []interface{}
	for _, isPauserItem := range isPauser {
		isPauserRule = append(isPauserRule, isPauserItem)
	}

	logs, sub, err := _Univ4poolmanager.contract.FilterLogs(opts, "SetPauser", guyRule, isPauserRule)
	if err != nil {
		return nil, err
	}
	return &Univ4poolmanagerSetPauserIterator{contract: _Univ4poolmanager.contract, event: "SetPauser", logs: logs, sub: sub}, nil
}

// WatchSetPauser is a free log subscription operation binding the contract event 0xd34f4aa5f94a385f2fa0ca25e5f01c6f331018f35c3d43a7b8057a86704de3df.
//
// Solidity: event SetPauser(address indexed guy, bool indexed isPauser)
func (_Univ4poolmanager *Univ4poolmanagerFilterer) WatchSetPauser(opts *bind.WatchOpts, sink chan<- *Univ4poolmanagerSetPauser, guy []common.Address, isPauser []bool) (event.Subscription, error) {

	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}
	var isPauserRule []interface{}
	for _, isPauserItem := range isPauser {
		isPauserRule = append(isPauserRule, isPauserItem)
	}

	logs, sub, err := _Univ4poolmanager.contract.WatchLogs(opts, "SetPauser", guyRule, isPauserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Univ4poolmanagerSetPauser)
				if err := _Univ4poolmanager.contract.UnpackLog(event, "SetPauser", log); err != nil {
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

// ParseSetPauser is a log parse operation binding the contract event 0xd34f4aa5f94a385f2fa0ca25e5f01c6f331018f35c3d43a7b8057a86704de3df.
//
// Solidity: event SetPauser(address indexed guy, bool indexed isPauser)
func (_Univ4poolmanager *Univ4poolmanagerFilterer) ParseSetPauser(log types.Log) (*Univ4poolmanagerSetPauser, error) {
	event := new(Univ4poolmanagerSetPauser)
	if err := _Univ4poolmanager.contract.UnpackLog(event, "SetPauser", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Univ4poolmanagerWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Univ4poolmanager contract.
type Univ4poolmanagerWithdrawIterator struct {
	Event *Univ4poolmanagerWithdraw // Event containing the contract specifics and raw log

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
func (it *Univ4poolmanagerWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Univ4poolmanagerWithdraw)
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
		it.Event = new(Univ4poolmanagerWithdraw)
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
func (it *Univ4poolmanagerWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Univ4poolmanagerWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Univ4poolmanagerWithdraw represents a Withdraw event raised by the Univ4poolmanager contract.
type Univ4poolmanagerWithdraw struct {
	Sender    common.Address
	Recipient common.Address
	PoolId    [32]byte
	Amount0   *big.Int
	Amount1   *big.Int
	Shares    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xbc70c2ef3795ca1df41695488a7ff6060de75f86dd892696bbfd76bdd123270f.
//
// Solidity: event Withdraw(address indexed sender, address indexed recipient, bytes32 indexed poolId, uint256 amount0, uint256 amount1, uint256 shares)
func (_Univ4poolmanager *Univ4poolmanagerFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address, poolId [][32]byte) (*Univ4poolmanagerWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Univ4poolmanager.contract.FilterLogs(opts, "Withdraw", senderRule, recipientRule, poolIdRule)
	if err != nil {
		return nil, err
	}
	return &Univ4poolmanagerWithdrawIterator{contract: _Univ4poolmanager.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xbc70c2ef3795ca1df41695488a7ff6060de75f86dd892696bbfd76bdd123270f.
//
// Solidity: event Withdraw(address indexed sender, address indexed recipient, bytes32 indexed poolId, uint256 amount0, uint256 amount1, uint256 shares)
func (_Univ4poolmanager *Univ4poolmanagerFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *Univ4poolmanagerWithdraw, sender []common.Address, recipient []common.Address, poolId [][32]byte) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Univ4poolmanager.contract.WatchLogs(opts, "Withdraw", senderRule, recipientRule, poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Univ4poolmanagerWithdraw)
				if err := _Univ4poolmanager.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xbc70c2ef3795ca1df41695488a7ff6060de75f86dd892696bbfd76bdd123270f.
//
// Solidity: event Withdraw(address indexed sender, address indexed recipient, bytes32 indexed poolId, uint256 amount0, uint256 amount1, uint256 shares)
func (_Univ4poolmanager *Univ4poolmanagerFilterer) ParseWithdraw(log types.Log) (*Univ4poolmanagerWithdraw, error) {
	event := new(Univ4poolmanagerWithdraw)
	if err := _Univ4poolmanager.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
