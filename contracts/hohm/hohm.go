// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hohm

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

// IOrigamiElevatedAccessExplicitAccess is an auto generated low-level Go binding around an user-defined struct.
type IOrigamiElevatedAccessExplicitAccess struct {
	FnSelector [4]byte
	Allowed    bool
}

// MessagingFee is an auto generated low-level Go binding around an user-defined struct.
type MessagingFee struct {
	NativeFee  *big.Int
	LzTokenFee *big.Int
}

// MessagingReceipt is an auto generated low-level Go binding around an user-defined struct.
type MessagingReceipt struct {
	Guid  [32]byte
	Nonce uint64
	Fee   MessagingFee
}

// OFTReceipt is an auto generated low-level Go binding around an user-defined struct.
type OFTReceipt struct {
	AmountSentLD     *big.Int
	AmountReceivedLD *big.Int
}

// SendParam is an auto generated low-level Go binding around an user-defined struct.
type SendParam struct {
	DstEid       uint32
	To           [32]byte
	AmountLD     *big.Int
	MinAmountLD  *big.Int
	ExtraOptions []byte
	ComposeMsg   []byte
	OftCmd       []byte
}

// HohmMetaData contains all meta data concerning the Hohm contract.
var HohmMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"collateralToken_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenPrices_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ExceededMaxExitWithShares\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ExceededMaxExitWithToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ExceededMaxJoinWithShares\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ExceededMaxJoinWithToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedNonZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAccess\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidParam\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denominator\",\"type\":\"uint256\"}],\"name\":\"PRBMath_MulDiv_Overflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"debtToken\",\"type\":\"address\"}],\"name\":\"DebtTokenSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"assets\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"liabilities\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Exit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"fnSelector\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"ExplicitAccessSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumIOrigamiTokenizedBalanceSheetVault.FeeType\",\"name\":\"feeType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"InKindFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"assets\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"liabilities\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Join\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"ManagerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxTotalSupply\",\"type\":\"uint256\"}],\"name\":\"MaxTotalSupplySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwnerAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldProposedOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newProposedOwner\",\"type\":\"address\"}],\"name\":\"NewOwnerProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporter\",\"type\":\"address\"}],\"name\":\"TeleporterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenPrices\",\"type\":\"address\"}],\"name\":\"TokenPricesSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"accountDelegationBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalCollateral\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"delegateAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"delegatedCollateral\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"areExitsPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"areJoinsPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"assetTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"availableSharesCapacity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balanceSheet\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"totalAssets\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"totalLiabilities\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateralToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"convertFromShares\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"assets\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"liabilities\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"convertFromToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"assets\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"liabilities\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"debtToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"delegateVotingPower\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exitFeeBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sharesOwner\",\"type\":\"address\"}],\"name\":\"exitWithShares\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"assets\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"liabilities\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sharesOwner\",\"type\":\"address\"}],\"name\":\"exitWithToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"assets\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"liabilities\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"name\":\"explicitFunctionAccess\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"isBalanceSheetToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isAsset\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isLiability\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"joinFeeBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"joinWithShares\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"assets\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"liabilities\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"joinWithToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"assets\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"liabilities\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liabilityTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"liabilities\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sharesOwner\",\"type\":\"address\"}],\"name\":\"maxExitWithShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxShares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sharesOwner\",\"type\":\"address\"}],\"name\":\"maxExitWithToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxToken\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxJoinWithShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxShares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxJoinWithToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxToken\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewExitWithShares\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"assets\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"liabilities\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"previewExitWithToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"assets\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"liabilities\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewJoinWithShares\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"assets\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"liabilities\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"previewJoinWithToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"assets\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"liabilities\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"proposeNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountLD\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraOptions\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"composeMsg\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"oftCmd\",\"type\":\"bytes\"}],\"internalType\":\"structSendParam\",\"name\":\"sendParam\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"payInLzToken\",\"type\":\"bool\"}],\"name\":\"quoteSend\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lzTokenFee\",\"type\":\"uint256\"}],\"internalType\":\"structMessagingFee\",\"name\":\"fee\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"recoverToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokeOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"assetAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"liabilityAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"sharesToMint\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newMaxTotalSupply\",\"type\":\"uint256\"}],\"name\":\"seed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountLD\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraOptions\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"composeMsg\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"oftCmd\",\"type\":\"bytes\"}],\"internalType\":\"structSendParam\",\"name\":\"sendParam\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lzTokenFee\",\"type\":\"uint256\"}],\"internalType\":\"structMessagingFee\",\"name\":\"fee\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"refundAddress\",\"type\":\"address\"}],\"name\":\"send\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lzTokenFee\",\"type\":\"uint256\"}],\"internalType\":\"structMessagingFee\",\"name\":\"fee\",\"type\":\"tuple\"}],\"internalType\":\"structMessagingReceipt\",\"name\":\"msgReceipt\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountSentLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountReceivedLD\",\"type\":\"uint256\"}],\"internalType\":\"structOFTReceipt\",\"name\":\"oftReceipt\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"allowedCaller\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes4\",\"name\":\"fnSelector\",\"type\":\"bytes4\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"internalType\":\"structIOrigamiElevatedAccess.ExplicitAccess[]\",\"name\":\"access\",\"type\":\"tuple[]\"}],\"name\":\"setExplicitAccess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newManager\",\"type\":\"address\"}],\"name\":\"setManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxTotalSupply_\",\"type\":\"uint256\"}],\"name\":\"setMaxTotalSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTeleporter\",\"type\":\"address\"}],\"name\":\"setTeleporter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenPrices\",\"type\":\"address\"}],\"name\":\"setTokenPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"syncDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporter\",\"outputs\":[{\"internalType\":\"contractIOFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenPrices\",\"outputs\":[{\"internalType\":\"contractITokenPrices\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"liabilities\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// HohmABI is the input ABI used to generate the binding from.
// Deprecated: Use HohmMetaData.ABI instead.
var HohmABI = HohmMetaData.ABI

// Hohm is an auto generated Go binding around an Ethereum contract.
type Hohm struct {
	HohmCaller     // Read-only binding to the contract
	HohmTransactor // Write-only binding to the contract
	HohmFilterer   // Log filterer for contract events
}

// HohmCaller is an auto generated read-only Go binding around an Ethereum contract.
type HohmCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HohmTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HohmTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HohmFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HohmFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HohmSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HohmSession struct {
	Contract     *Hohm             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HohmCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HohmCallerSession struct {
	Contract *HohmCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HohmTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HohmTransactorSession struct {
	Contract     *HohmTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HohmRaw is an auto generated low-level Go binding around an Ethereum contract.
type HohmRaw struct {
	Contract *Hohm // Generic contract binding to access the raw methods on
}

// HohmCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HohmCallerRaw struct {
	Contract *HohmCaller // Generic read-only contract binding to access the raw methods on
}

// HohmTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HohmTransactorRaw struct {
	Contract *HohmTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHohm creates a new instance of Hohm, bound to a specific deployed contract.
func NewHohm(address common.Address, backend bind.ContractBackend) (*Hohm, error) {
	contract, err := bindHohm(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hohm{HohmCaller: HohmCaller{contract: contract}, HohmTransactor: HohmTransactor{contract: contract}, HohmFilterer: HohmFilterer{contract: contract}}, nil
}

// NewHohmCaller creates a new read-only instance of Hohm, bound to a specific deployed contract.
func NewHohmCaller(address common.Address, caller bind.ContractCaller) (*HohmCaller, error) {
	contract, err := bindHohm(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HohmCaller{contract: contract}, nil
}

// NewHohmTransactor creates a new write-only instance of Hohm, bound to a specific deployed contract.
func NewHohmTransactor(address common.Address, transactor bind.ContractTransactor) (*HohmTransactor, error) {
	contract, err := bindHohm(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HohmTransactor{contract: contract}, nil
}

// NewHohmFilterer creates a new log filterer instance of Hohm, bound to a specific deployed contract.
func NewHohmFilterer(address common.Address, filterer bind.ContractFilterer) (*HohmFilterer, error) {
	contract, err := bindHohm(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HohmFilterer{contract: contract}, nil
}

// bindHohm binds a generic wrapper to an already deployed contract.
func bindHohm(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HohmMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hohm *HohmRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hohm.Contract.HohmCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hohm *HohmRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hohm.Contract.HohmTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hohm *HohmRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hohm.Contract.HohmTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hohm *HohmCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hohm.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hohm *HohmTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hohm.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hohm *HohmTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hohm.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Hohm *HohmCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Hohm.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Hohm *HohmSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Hohm.Contract.DOMAINSEPARATOR(&_Hohm.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Hohm *HohmCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Hohm.Contract.DOMAINSEPARATOR(&_Hohm.CallOpts)
}

// AccountDelegationBalances is a free data retrieval call binding the contract method 0x61b3009a.
//
// Solidity: function accountDelegationBalances(address account) view returns(uint256 totalCollateral, address delegateAddress, uint256 delegatedCollateral)
func (_Hohm *HohmCaller) AccountDelegationBalances(opts *bind.CallOpts, account common.Address) (struct {
	TotalCollateral     *big.Int
	DelegateAddress     common.Address
	DelegatedCollateral *big.Int
}, error) {
	var out []interface{}
	err := _Hohm.contract.Call(opts, &out, "accountDelegationBalances", account)

	outstruct := new(struct {
		TotalCollateral     *big.Int
		DelegateAddress     common.Address
		DelegatedCollateral *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalCollateral = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DelegateAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.DelegatedCollateral = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AccountDelegationBalances is a free data retrieval call binding the contract method 0x61b3009a.
//
// Solidity: function accountDelegationBalances(address account) view returns(uint256 totalCollateral, address delegateAddress, uint256 delegatedCollateral)
func (_Hohm *HohmSession) AccountDelegationBalances(account common.Address) (struct {
	TotalCollateral     *big.Int
	DelegateAddress     common.Address
	DelegatedCollateral *big.Int
}, error) {
	return _Hohm.Contract.AccountDelegationBalances(&_Hohm.CallOpts, account)
}

// AccountDelegationBalances is a free data retrieval call binding the contract method 0x61b3009a.
//
// Solidity: function accountDelegationBalances(address account) view returns(uint256 totalCollateral, address delegateAddress, uint256 delegatedCollateral)
func (_Hohm *HohmCallerSession) AccountDelegationBalances(account common.Address) (struct {
	TotalCollateral     *big.Int
	DelegateAddress     common.Address
	DelegatedCollateral *big.Int
}, error) {
	return _Hohm.Contract.AccountDelegationBalances(&_Hohm.CallOpts, account)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address tokenOwner, address spender) view returns(uint256)
func (_Hohm *HohmCaller) Allowance(opts *bind.CallOpts, tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Hohm.contract.Call(opts, &out, "allowance", tokenOwner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address tokenOwner, address spender) view returns(uint256)
func (_Hohm *HohmSession) Allowance(tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	return _Hohm.Contract.Allowance(&_Hohm.CallOpts, tokenOwner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address tokenOwner, address spender) view returns(uint256)
func (_Hohm *HohmCallerSession) Allowance(tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	return _Hohm.Contract.Allowance(&_Hohm.CallOpts, tokenOwner, spender)
}

// AreExitsPaused is a free data retrieval call binding the contract method 0x9a6b27cf.
//
// Solidity: function areExitsPaused() view returns(bool)
func (_Hohm *HohmCaller) AreExitsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Hohm.contract.Call(opts, &out, "areExitsPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AreExitsPaused is a free data retrieval call binding the contract method 0x9a6b27cf.
//
// Solidity: function areExitsPaused() view returns(bool)
func (_Hohm *HohmSession) AreExitsPaused() (bool, error) {
	return _Hohm.Contract.AreExitsPaused(&_Hohm.CallOpts)
}

// AreExitsPaused is a free data retrieval call binding the contract method 0x9a6b27cf.
//
// Solidity: function areExitsPaused() view returns(bool)
func (_Hohm *HohmCallerSession) AreExitsPaused() (bool, error) {
	return _Hohm.Contract.AreExitsPaused(&_Hohm.CallOpts)
}

// AreJoinsPaused is a free data retrieval call binding the contract method 0xaf1454d3.
//
// Solidity: function areJoinsPaused() view returns(bool)
func (_Hohm *HohmCaller) AreJoinsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Hohm.contract.Call(opts, &out, "areJoinsPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AreJoinsPaused is a free data retrieval call binding the contract method 0xaf1454d3.
//
// Solidity: function areJoinsPaused() view returns(bool)
func (_Hohm *HohmSession) AreJoinsPaused() (bool, error) {
	return _Hohm.Contract.AreJoinsPaused(&_Hohm.CallOpts)
}

// AreJoinsPaused is a free data retrieval call binding the contract method 0xaf1454d3.
//
// Solidity: function areJoinsPaused() view returns(bool)
func (_Hohm *HohmCallerSession) AreJoinsPaused() (bool, error) {
	return _Hohm.Contract.AreJoinsPaused(&_Hohm.CallOpts)
}

// AssetTokens is a free data retrieval call binding the contract method 0xe9607c01.
//
// Solidity: function assetTokens() view returns(address[] assets)
func (_Hohm *HohmCaller) AssetTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Hohm.contract.Call(opts, &out, "assetTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AssetTokens is a free data retrieval call binding the contract method 0xe9607c01.
//
// Solidity: function assetTokens() view returns(address[] assets)
func (_Hohm *HohmSession) AssetTokens() ([]common.Address, error) {
	return _Hohm.Contract.AssetTokens(&_Hohm.CallOpts)
}

// AssetTokens is a free data retrieval call binding the contract method 0xe9607c01.
//
// Solidity: function assetTokens() view returns(address[] assets)
func (_Hohm *HohmCallerSession) AssetTokens() ([]common.Address, error) {
	return _Hohm.Contract.AssetTokens(&_Hohm.CallOpts)
}

// AvailableSharesCapacity is a free data retrieval call binding the contract method 0xe9cad766.
//
// Solidity: function availableSharesCapacity() view returns(uint256 shares)
func (_Hohm *HohmCaller) AvailableSharesCapacity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hohm.contract.Call(opts, &out, "availableSharesCapacity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableSharesCapacity is a free data retrieval call binding the contract method 0xe9cad766.
//
// Solidity: function availableSharesCapacity() view returns(uint256 shares)
func (_Hohm *HohmSession) AvailableSharesCapacity() (*big.Int, error) {
	return _Hohm.Contract.AvailableSharesCapacity(&_Hohm.CallOpts)
}

// AvailableSharesCapacity is a free data retrieval call binding the contract method 0xe9cad766.
//
// Solidity: function availableSharesCapacity() view returns(uint256 shares)
func (_Hohm *HohmCallerSession) AvailableSharesCapacity() (*big.Int, error) {
	return _Hohm.Contract.AvailableSharesCapacity(&_Hohm.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Hohm *HohmCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Hohm.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Hohm *HohmSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Hohm.Contract.BalanceOf(&_Hohm.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Hohm *HohmCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Hohm.Contract.BalanceOf(&_Hohm.CallOpts, account)
}

// BalanceSheet is a free data retrieval call binding the contract method 0x22285cf6.
//
// Solidity: function balanceSheet() view returns(uint256[] totalAssets, uint256[] totalLiabilities)
func (_Hohm *HohmCaller) BalanceSheet(opts *bind.CallOpts) (struct {
	TotalAssets      []*big.Int
	TotalLiabilities []*big.Int
}, error) {
	var out []interface{}
	err := _Hohm.contract.Call(opts, &out, "balanceSheet")

	outstruct := new(struct {
		TotalAssets      []*big.Int
		TotalLiabilities []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalAssets = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.TotalLiabilities = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// BalanceSheet is a free data retrieval call binding the contract method 0x22285cf6.
//
// Solidity: function balanceSheet() view returns(uint256[] totalAssets, uint256[] totalLiabilities)
func (_Hohm *HohmSession) BalanceSheet() (struct {
	TotalAssets      []*big.Int
	TotalLiabilities []*big.Int
}, error) {
	return _Hohm.Contract.BalanceSheet(&_Hohm.CallOpts)
}

// BalanceSheet is a free data retrieval call binding the contract method 0x22285cf6.
//
// Solidity: function balanceSheet() view returns(uint256[] totalAssets, uint256[] totalLiabilities)
func (_Hohm *HohmCallerSession) BalanceSheet() (struct {
	TotalAssets      []*big.Int
	TotalLiabilities []*big.Int
}, error) {
	return _Hohm.Contract.BalanceSheet(&_Hohm.CallOpts)
}

// CollateralToken is a free data retrieval call binding the contract method 0xb2016bd4.
//
// Solidity: function collateralToken() view returns(address)
func (_Hohm *HohmCaller) CollateralToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hohm.contract.Call(opts, &out, "collateralToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CollateralToken is a free data retrieval call binding the contract method 0xb2016bd4.
//
// Solidity: function collateralToken() view returns(address)
func (_Hohm *HohmSession) CollateralToken() (common.Address, error) {
	return _Hohm.Contract.CollateralToken(&_Hohm.CallOpts)
}

// CollateralToken is a free data retrieval call binding the contract method 0xb2016bd4.
//
// Solidity: function collateralToken() view returns(address)
func (_Hohm *HohmCallerSession) CollateralToken() (common.Address, error) {
	return _Hohm.Contract.CollateralToken(&_Hohm.CallOpts)
}

// ConvertFromShares is a free data retrieval call binding the contract method 0x8f39a80b.
//
// Solidity: function convertFromShares(uint256 shares) view returns(uint256[] assets, uint256[] liabilities)
func (_Hohm *HohmCaller) ConvertFromShares(opts *bind.CallOpts, shares *big.Int) (struct {
	Assets      []*big.Int
	Liabilities []*big.Int
}, error) {
	var out []interface{}
	err := _Hohm.contract.Call(opts, &out, "convertFromShares", shares)

	outstruct := new(struct {
		Assets      []*big.Int
		Liabilities []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Assets = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Liabilities = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// ConvertFromShares is a free data retrieval call binding the contract method 0x8f39a80b.
//
// Solidity: function convertFromShares(uint256 shares) view returns(uint256[] assets, uint256[] liabilities)
func (_Hohm *HohmSession) ConvertFromShares(shares *big.Int) (struct {
	Assets      []*big.Int
	Liabilities []*big.Int
}, error) {
	return _Hohm.Contract.ConvertFromShares(&_Hohm.CallOpts, shares)
}

// ConvertFromShares is a free data retrieval call binding the contract method 0x8f39a80b.
//
// Solidity: function convertFromShares(uint256 shares) view returns(uint256[] assets, uint256[] liabilities)
func (_Hohm *HohmCallerSession) ConvertFromShares(shares *big.Int) (struct {
	Assets      []*big.Int
	Liabilities []*big.Int
}, error) {
	return _Hohm.Contract.ConvertFromShares(&_Hohm.CallOpts, shares)
}

// ConvertFromToken is a free data retrieval call binding the contract method 0x2113f244.
//
// Solidity: function convertFromToken(address tokenAddress, uint256 tokenAmount) view returns(uint256 shares, uint256[] assets, uint256[] liabilities)
func (_Hohm *HohmCaller) ConvertFromToken(opts *bind.CallOpts, tokenAddress common.Address, tokenAmount *big.Int) (struct {
	Shares      *big.Int
	Assets      []*big.Int
	Liabilities []*big.Int
}, error) {
	var out []interface{}
	err := _Hohm.contract.Call(opts, &out, "convertFromToken", tokenAddress, tokenAmount)

	outstruct := new(struct {
		Shares      *big.Int
		Assets      []*big.Int
		Liabilities []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Shares = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Assets = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.Liabilities = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// ConvertFromToken is a free data retrieval call binding the contract method 0x2113f244.
//
// Solidity: function convertFromToken(address tokenAddress, uint256 tokenAmount) view returns(uint256 shares, uint256[] assets, uint256[] liabilities)
func (_Hohm *HohmSession) ConvertFromToken(tokenAddress common.Address, tokenAmount *big.Int) (struct {
	Shares      *big.Int
	Assets      []*big.Int
	Liabilities []*big.Int
}, error) {
	return _Hohm.Contract.ConvertFromToken(&_Hohm.CallOpts, tokenAddress, tokenAmount)
}

// ConvertFromToken is a free data retrieval call binding the contract method 0x2113f244.
//
// Solidity: function convertFromToken(address tokenAddress, uint256 tokenAmount) view returns(uint256 shares, uint256[] assets, uint256[] liabilities)
func (_Hohm *HohmCallerSession) ConvertFromToken(tokenAddress common.Address, tokenAmount *big.Int) (struct {
	Shares      *big.Int
	Assets      []*big.Int
	Liabilities []*big.Int
}, error) {
	return _Hohm.Contract.ConvertFromToken(&_Hohm.CallOpts, tokenAddress, tokenAmount)
}

// DebtToken is a free data retrieval call binding the contract method 0xf8d89898.
//
// Solidity: function debtToken() view returns(address)
func (_Hohm *HohmCaller) DebtToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hohm.contract.Call(opts, &out, "debtToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DebtToken is a free data retrieval call binding the contract method 0xf8d89898.
//
// Solidity: function debtToken() view returns(address)
func (_Hohm *HohmSession) DebtToken() (common.Address, error) {
	return _Hohm.Contract.DebtToken(&_Hohm.CallOpts)
}

// DebtToken is a free data retrieval call binding the contract method 0xf8d89898.
//
// Solidity: function debtToken() view returns(address)
func (_Hohm *HohmCallerSession) DebtToken() (common.Address, error) {
	return _Hohm.Contract.DebtToken(&_Hohm.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Hohm *HohmCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Hohm.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Hohm *HohmSession) Decimals() (uint8, error) {
	return _Hohm.Contract.Decimals(&_Hohm.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Hohm *HohmCallerSession) Decimals() (uint8, error) {
	return _Hohm.Contract.Decimals(&_Hohm.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Hohm *HohmCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _Hohm.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Hohm *HohmSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Hohm.Contract.Eip712Domain(&_Hohm.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Hohm *HohmCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Hohm.Contract.Eip712Domain(&_Hohm.CallOpts)
}

// ExitFeeBps is a free data retrieval call binding the contract method 0x57b17a52.
//
// Solidity: function exitFeeBps() view returns(uint256)
func (_Hohm *HohmCaller) ExitFeeBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hohm.contract.Call(opts, &out, "exitFeeBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExitFeeBps is a free data retrieval call binding the contract method 0x57b17a52.
//
// Solidity: function exitFeeBps() view returns(uint256)
func (_Hohm *HohmSession) ExitFeeBps() (*big.Int, error) {
	return _Hohm.Contract.ExitFeeBps(&_Hohm.CallOpts)
}

// ExitFeeBps is a free data retrieval call binding the contract method 0x57b17a52.
//
// Solidity: function exitFeeBps() view returns(uint256)
func (_Hohm *HohmCallerSession) ExitFeeBps() (*big.Int, error) {
	return _Hohm.Contract.ExitFeeBps(&_Hohm.CallOpts)
}

// ExplicitFunctionAccess is a free data retrieval call binding the contract method 0xdaeccc79.
//
// Solidity: function explicitFunctionAccess(address , bytes4 ) view returns(bool)
func (_Hohm *HohmCaller) ExplicitFunctionAccess(opts *bind.CallOpts, arg0 common.Address, arg1 [4]byte) (bool, error) {
	var out []interface{}
	err := _Hohm.contract.Call(opts, &out, "explicitFunctionAccess", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExplicitFunctionAccess is a free data retrieval call binding the contract method 0xdaeccc79.
//
// Solidity: function explicitFunctionAccess(address , bytes4 ) view returns(bool)
func (_Hohm *HohmSession) ExplicitFunctionAccess(arg0 common.Address, arg1 [4]byte) (bool, error) {
	return _Hohm.Contract.ExplicitFunctionAccess(&_Hohm.CallOpts, arg0, arg1)
}

// ExplicitFunctionAccess is a free data retrieval call binding the contract method 0xdaeccc79.
//
// Solidity: function explicitFunctionAccess(address , bytes4 ) view returns(bool)
func (_Hohm *HohmCallerSession) ExplicitFunctionAccess(arg0 common.Address, arg1 [4]byte) (bool, error) {
	return _Hohm.Contract.ExplicitFunctionAccess(&_Hohm.CallOpts, arg0, arg1)
}

// IsBalanceSheetToken is a free data retrieval call binding the contract method 0xff249518.
//
// Solidity: function isBalanceSheetToken(address tokenAddress) view returns(bool isAsset, bool isLiability)
func (_Hohm *HohmCaller) IsBalanceSheetToken(opts *bind.CallOpts, tokenAddress common.Address) (struct {
	IsAsset     bool
	IsLiability bool
}, error) {
	var out []interface{}
	err := _Hohm.contract.Call(opts, &out, "isBalanceSheetToken", tokenAddress)

	outstruct := new(struct {
		IsAsset     bool
		IsLiability bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsAsset = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.IsLiability = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// IsBalanceSheetToken is a free data retrieval call binding the contract method 0xff249518.
//
// Solidity: function isBalanceSheetToken(address tokenAddress) view returns(bool isAsset, bool isLiability)
func (_Hohm *HohmSession) IsBalanceSheetToken(tokenAddress common.Address) (struct {
	IsAsset     bool
	IsLiability bool
}, error) {
	return _Hohm.Contract.IsBalanceSheetToken(&_Hohm.CallOpts, tokenAddress)
}

// IsBalanceSheetToken is a free data retrieval call binding the contract method 0xff249518.
//
// Solidity: function isBalanceSheetToken(address tokenAddress) view returns(bool isAsset, bool isLiability)
func (_Hohm *HohmCallerSession) IsBalanceSheetToken(tokenAddress common.Address) (struct {
	IsAsset     bool
	IsLiability bool
}, error) {
	return _Hohm.Contract.IsBalanceSheetToken(&_Hohm.CallOpts, tokenAddress)
}

// JoinFeeBps is a free data retrieval call binding the contract method 0x4e7e240d.
//
// Solidity: function joinFeeBps() view returns(uint256)
func (_Hohm *HohmCaller) JoinFeeBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hohm.contract.Call(opts, &out, "joinFeeBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// JoinFeeBps is a free data retrieval call binding the contract method 0x4e7e240d.
//
// Solidity: function joinFeeBps() view returns(uint256)
func (_Hohm *HohmSession) JoinFeeBps() (*big.Int, error) {
	return _Hohm.Contract.JoinFeeBps(&_Hohm.CallOpts)
}

// JoinFeeBps is a free data retrieval call binding the contract method 0x4e7e240d.
//
// Solidity: function joinFeeBps() view returns(uint256)
func (_Hohm *HohmCallerSession) JoinFeeBps() (*big.Int, error) {
	return _Hohm.Contract.JoinFeeBps(&_Hohm.CallOpts)
}

// LiabilityTokens is a free data retrieval call binding the contract method 0x00274eb3.
//
// Solidity: function liabilityTokens() view returns(address[] liabilities)
func (_Hohm *HohmCaller) LiabilityTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Hohm.contract.Call(opts, &out, "liabilityTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// LiabilityTokens is a free data retrieval call binding the contract method 0x00274eb3.
//
// Solidity: function liabilityTokens() view returns(address[] liabilities)
func (_Hohm *HohmSession) LiabilityTokens() ([]common.Address, error) {
	return _Hohm.Contract.LiabilityTokens(&_Hohm.CallOpts)
}

// LiabilityTokens is a free data retrieval call binding the contract method 0x00274eb3.
//
// Solidity: function liabilityTokens() view returns(address[] liabilities)
func (_Hohm *HohmCallerSession) LiabilityTokens() ([]common.Address, error) {
	return _Hohm.Contract.LiabilityTokens(&_Hohm.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Hohm *HohmCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hohm.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Hohm *HohmSession) Manager() (common.Address, error) {
	return _Hohm.Contract.Manager(&_Hohm.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Hohm *HohmCallerSession) Manager() (common.Address, error) {
	return _Hohm.Contract.Manager(&_Hohm.CallOpts)
}

// MaxExitWithShares is a free data retrieval call binding the contract method 0xc28ab675.
//
// Solidity: function maxExitWithShares(address sharesOwner) view returns(uint256 maxShares)
func (_Hohm *HohmCaller) MaxExitWithShares(opts *bind.CallOpts, sharesOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Hohm.contract.Call(opts, &out, "maxExitWithShares", sharesOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxExitWithShares is a free data retrieval call binding the contract method 0xc28ab675.
//
// Solidity: function maxExitWithShares(address sharesOwner) view returns(uint256 maxShares)
func (_Hohm *HohmSession) MaxExitWithShares(sharesOwner common.Address) (*big.Int, error) {
	return _Hohm.Contract.MaxExitWithShares(&_Hohm.CallOpts, sharesOwner)
}

// MaxExitWithShares is a free data retrieval call binding the contract method 0xc28ab675.
//
// Solidity: function maxExitWithShares(address sharesOwner) view returns(uint256 maxShares)
func (_Hohm *HohmCallerSession) MaxExitWithShares(sharesOwner common.Address) (*big.Int, error) {
	return _Hohm.Contract.MaxExitWithShares(&_Hohm.CallOpts, sharesOwner)
}

// MaxExitWithToken is a free data retrieval call binding the contract method 0xb2317e6b.
//
// Solidity: function maxExitWithToken(address tokenAddress, address sharesOwner) view returns(uint256 maxToken)
func (_Hohm *HohmCaller) MaxExitWithToken(opts *bind.CallOpts, tokenAddress common.Address, sharesOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Hohm.contract.Call(opts, &out, "maxExitWithToken", tokenAddress, sharesOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxExitWithToken is a free data retrieval call binding the contract method 0xb2317e6b.
//
// Solidity: function maxExitWithToken(address tokenAddress, address sharesOwner) view returns(uint256 maxToken)
func (_Hohm *HohmSession) MaxExitWithToken(tokenAddress common.Address, sharesOwner common.Address) (*big.Int, error) {
	return _Hohm.Contract.MaxExitWithToken(&_Hohm.CallOpts, tokenAddress, sharesOwner)
}

// MaxExitWithToken is a free data retrieval call binding the contract method 0xb2317e6b.
//
// Solidity: function maxExitWithToken(address tokenAddress, address sharesOwner) view returns(uint256 maxToken)
func (_Hohm *HohmCallerSession) MaxExitWithToken(tokenAddress common.Address, sharesOwner common.Address) (*big.Int, error) {
	return _Hohm.Contract.MaxExitWithToken(&_Hohm.CallOpts, tokenAddress, sharesOwner)
}

// MaxJoinWithShares is a free data retrieval call binding the contract method 0x359d6ea9.
//
// Solidity: function maxJoinWithShares(address ) view returns(uint256 maxShares)
func (_Hohm *HohmCaller) MaxJoinWithShares(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Hohm.contract.Call(opts, &out, "maxJoinWithShares", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxJoinWithShares is a free data retrieval call binding the contract method 0x359d6ea9.
//
// Solidity: function maxJoinWithShares(address ) view returns(uint256 maxShares)
func (_Hohm *HohmSession) MaxJoinWithShares(arg0 common.Address) (*big.Int, error) {
	return _Hohm.Contract.MaxJoinWithShares(&_Hohm.CallOpts, arg0)
}

// MaxJoinWithShares is a free data retrieval call binding the contract method 0x359d6ea9.
//
// Solidity: function maxJoinWithShares(address ) view returns(uint256 maxShares)
func (_Hohm *HohmCallerSession) MaxJoinWithShares(arg0 common.Address) (*big.Int, error) {
	return _Hohm.Contract.MaxJoinWithShares(&_Hohm.CallOpts, arg0)
}

// MaxJoinWithToken is a free data retrieval call binding the contract method 0xc1a97590.
//
// Solidity: function maxJoinWithToken(address tokenAddress, address ) view returns(uint256 maxToken)
func (_Hohm *HohmCaller) MaxJoinWithToken(opts *bind.CallOpts, tokenAddress common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Hohm.contract.Call(opts, &out, "maxJoinWithToken", tokenAddress, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxJoinWithToken is a free data retrieval call binding the contract method 0xc1a97590.
//
// Solidity: function maxJoinWithToken(address tokenAddress, address ) view returns(uint256 maxToken)
func (_Hohm *HohmSession) MaxJoinWithToken(tokenAddress common.Address, arg1 common.Address) (*big.Int, error) {
	return _Hohm.Contract.MaxJoinWithToken(&_Hohm.CallOpts, tokenAddress, arg1)
}

// MaxJoinWithToken is a free data retrieval call binding the contract method 0xc1a97590.
//
// Solidity: function maxJoinWithToken(address tokenAddress, address ) view returns(uint256 maxToken)
func (_Hohm *HohmCallerSession) MaxJoinWithToken(tokenAddress common.Address, arg1 common.Address) (*big.Int, error) {
	return _Hohm.Contract.MaxJoinWithToken(&_Hohm.CallOpts, tokenAddress, arg1)
}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x2ab4d052.
//
// Solidity: function maxTotalSupply() view returns(uint256)
func (_Hohm *HohmCaller) MaxTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hohm.contract.Call(opts, &out, "maxTotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x2ab4d052.
//
// Solidity: function maxTotalSupply() view returns(uint256)
func (_Hohm *HohmSession) MaxTotalSupply() (*big.Int, error) {
	return _Hohm.Contract.MaxTotalSupply(&_Hohm.CallOpts)
}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x2ab4d052.
//
// Solidity: function maxTotalSupply() view returns(uint256)
func (_Hohm *HohmCallerSession) MaxTotalSupply() (*big.Int, error) {
	return _Hohm.Contract.MaxTotalSupply(&_Hohm.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Hohm *HohmCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Hohm.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Hohm *HohmSession) Name() (string, error) {
	return _Hohm.Contract.Name(&_Hohm.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Hohm *HohmCallerSession) Name() (string, error) {
	return _Hohm.Contract.Name(&_Hohm.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Hohm *HohmCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Hohm.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Hohm *HohmSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Hohm.Contract.Nonces(&_Hohm.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Hohm *HohmCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Hohm.Contract.Nonces(&_Hohm.CallOpts, owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Hohm *HohmCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hohm.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Hohm *HohmSession) Owner() (common.Address, error) {
	return _Hohm.Contract.Owner(&_Hohm.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Hohm *HohmCallerSession) Owner() (common.Address, error) {
	return _Hohm.Contract.Owner(&_Hohm.CallOpts)
}

// PreviewExitWithShares is a free data retrieval call binding the contract method 0x2ab41950.
//
// Solidity: function previewExitWithShares(uint256 shares) view returns(uint256[] assets, uint256[] liabilities)
func (_Hohm *HohmCaller) PreviewExitWithShares(opts *bind.CallOpts, shares *big.Int) (struct {
	Assets      []*big.Int
	Liabilities []*big.Int
}, error) {
	var out []interface{}
	err := _Hohm.contract.Call(opts, &out, "previewExitWithShares", shares)

	outstruct := new(struct {
		Assets      []*big.Int
		Liabilities []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Assets = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Liabilities = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// PreviewExitWithShares is a free data retrieval call binding the contract method 0x2ab41950.
//
// Solidity: function previewExitWithShares(uint256 shares) view returns(uint256[] assets, uint256[] liabilities)
func (_Hohm *HohmSession) PreviewExitWithShares(shares *big.Int) (struct {
	Assets      []*big.Int
	Liabilities []*big.Int
}, error) {
	return _Hohm.Contract.PreviewExitWithShares(&_Hohm.CallOpts, shares)
}

// PreviewExitWithShares is a free data retrieval call binding the contract method 0x2ab41950.
//
// Solidity: function previewExitWithShares(uint256 shares) view returns(uint256[] assets, uint256[] liabilities)
func (_Hohm *HohmCallerSession) PreviewExitWithShares(shares *big.Int) (struct {
	Assets      []*big.Int
	Liabilities []*big.Int
}, error) {
	return _Hohm.Contract.PreviewExitWithShares(&_Hohm.CallOpts, shares)
}

// PreviewExitWithToken is a free data retrieval call binding the contract method 0x20c3627f.
//
// Solidity: function previewExitWithToken(address tokenAddress, uint256 tokenAmount) view returns(uint256 shares, uint256[] assets, uint256[] liabilities)
func (_Hohm *HohmCaller) PreviewExitWithToken(opts *bind.CallOpts, tokenAddress common.Address, tokenAmount *big.Int) (struct {
	Shares      *big.Int
	Assets      []*big.Int
	Liabilities []*big.Int
}, error) {
	var out []interface{}
	err := _Hohm.contract.Call(opts, &out, "previewExitWithToken", tokenAddress, tokenAmount)

	outstruct := new(struct {
		Shares      *big.Int
		Assets      []*big.Int
		Liabilities []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Shares = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Assets = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.Liabilities = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// PreviewExitWithToken is a free data retrieval call binding the contract method 0x20c3627f.
//
// Solidity: function previewExitWithToken(address tokenAddress, uint256 tokenAmount) view returns(uint256 shares, uint256[] assets, uint256[] liabilities)
func (_Hohm *HohmSession) PreviewExitWithToken(tokenAddress common.Address, tokenAmount *big.Int) (struct {
	Shares      *big.Int
	Assets      []*big.Int
	Liabilities []*big.Int
}, error) {
	return _Hohm.Contract.PreviewExitWithToken(&_Hohm.CallOpts, tokenAddress, tokenAmount)
}

// PreviewExitWithToken is a free data retrieval call binding the contract method 0x20c3627f.
//
// Solidity: function previewExitWithToken(address tokenAddress, uint256 tokenAmount) view returns(uint256 shares, uint256[] assets, uint256[] liabilities)
func (_Hohm *HohmCallerSession) PreviewExitWithToken(tokenAddress common.Address, tokenAmount *big.Int) (struct {
	Shares      *big.Int
	Assets      []*big.Int
	Liabilities []*big.Int
}, error) {
	return _Hohm.Contract.PreviewExitWithToken(&_Hohm.CallOpts, tokenAddress, tokenAmount)
}

// PreviewJoinWithShares is a free data retrieval call binding the contract method 0xd9b47d1c.
//
// Solidity: function previewJoinWithShares(uint256 shares) view returns(uint256[] assets, uint256[] liabilities)
func (_Hohm *HohmCaller) PreviewJoinWithShares(opts *bind.CallOpts, shares *big.Int) (struct {
	Assets      []*big.Int
	Liabilities []*big.Int
}, error) {
	var out []interface{}
	err := _Hohm.contract.Call(opts, &out, "previewJoinWithShares", shares)

	outstruct := new(struct {
		Assets      []*big.Int
		Liabilities []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Assets = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Liabilities = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// PreviewJoinWithShares is a free data retrieval call binding the contract method 0xd9b47d1c.
//
// Solidity: function previewJoinWithShares(uint256 shares) view returns(uint256[] assets, uint256[] liabilities)
func (_Hohm *HohmSession) PreviewJoinWithShares(shares *big.Int) (struct {
	Assets      []*big.Int
	Liabilities []*big.Int
}, error) {
	return _Hohm.Contract.PreviewJoinWithShares(&_Hohm.CallOpts, shares)
}

// PreviewJoinWithShares is a free data retrieval call binding the contract method 0xd9b47d1c.
//
// Solidity: function previewJoinWithShares(uint256 shares) view returns(uint256[] assets, uint256[] liabilities)
func (_Hohm *HohmCallerSession) PreviewJoinWithShares(shares *big.Int) (struct {
	Assets      []*big.Int
	Liabilities []*big.Int
}, error) {
	return _Hohm.Contract.PreviewJoinWithShares(&_Hohm.CallOpts, shares)
}

// PreviewJoinWithToken is a free data retrieval call binding the contract method 0x907ad64a.
//
// Solidity: function previewJoinWithToken(address tokenAddress, uint256 tokenAmount) view returns(uint256 shares, uint256[] assets, uint256[] liabilities)
func (_Hohm *HohmCaller) PreviewJoinWithToken(opts *bind.CallOpts, tokenAddress common.Address, tokenAmount *big.Int) (struct {
	Shares      *big.Int
	Assets      []*big.Int
	Liabilities []*big.Int
}, error) {
	var out []interface{}
	err := _Hohm.contract.Call(opts, &out, "previewJoinWithToken", tokenAddress, tokenAmount)

	outstruct := new(struct {
		Shares      *big.Int
		Assets      []*big.Int
		Liabilities []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Shares = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Assets = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.Liabilities = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// PreviewJoinWithToken is a free data retrieval call binding the contract method 0x907ad64a.
//
// Solidity: function previewJoinWithToken(address tokenAddress, uint256 tokenAmount) view returns(uint256 shares, uint256[] assets, uint256[] liabilities)
func (_Hohm *HohmSession) PreviewJoinWithToken(tokenAddress common.Address, tokenAmount *big.Int) (struct {
	Shares      *big.Int
	Assets      []*big.Int
	Liabilities []*big.Int
}, error) {
	return _Hohm.Contract.PreviewJoinWithToken(&_Hohm.CallOpts, tokenAddress, tokenAmount)
}

// PreviewJoinWithToken is a free data retrieval call binding the contract method 0x907ad64a.
//
// Solidity: function previewJoinWithToken(address tokenAddress, uint256 tokenAmount) view returns(uint256 shares, uint256[] assets, uint256[] liabilities)
func (_Hohm *HohmCallerSession) PreviewJoinWithToken(tokenAddress common.Address, tokenAmount *big.Int) (struct {
	Shares      *big.Int
	Assets      []*big.Int
	Liabilities []*big.Int
}, error) {
	return _Hohm.Contract.PreviewJoinWithToken(&_Hohm.CallOpts, tokenAddress, tokenAmount)
}

// QuoteSend is a free data retrieval call binding the contract method 0x3b6f743b.
//
// Solidity: function quoteSend((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) sendParam, bool payInLzToken) view returns((uint256,uint256) fee)
func (_Hohm *HohmCaller) QuoteSend(opts *bind.CallOpts, sendParam SendParam, payInLzToken bool) (MessagingFee, error) {
	var out []interface{}
	err := _Hohm.contract.Call(opts, &out, "quoteSend", sendParam, payInLzToken)

	if err != nil {
		return *new(MessagingFee), err
	}

	out0 := *abi.ConvertType(out[0], new(MessagingFee)).(*MessagingFee)

	return out0, err

}

// QuoteSend is a free data retrieval call binding the contract method 0x3b6f743b.
//
// Solidity: function quoteSend((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) sendParam, bool payInLzToken) view returns((uint256,uint256) fee)
func (_Hohm *HohmSession) QuoteSend(sendParam SendParam, payInLzToken bool) (MessagingFee, error) {
	return _Hohm.Contract.QuoteSend(&_Hohm.CallOpts, sendParam, payInLzToken)
}

// QuoteSend is a free data retrieval call binding the contract method 0x3b6f743b.
//
// Solidity: function quoteSend((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) sendParam, bool payInLzToken) view returns((uint256,uint256) fee)
func (_Hohm *HohmCallerSession) QuoteSend(sendParam SendParam, payInLzToken bool) (MessagingFee, error) {
	return _Hohm.Contract.QuoteSend(&_Hohm.CallOpts, sendParam, payInLzToken)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Hohm *HohmCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Hohm.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Hohm *HohmSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Hohm.Contract.SupportsInterface(&_Hohm.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Hohm *HohmCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Hohm.Contract.SupportsInterface(&_Hohm.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Hohm *HohmCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Hohm.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Hohm *HohmSession) Symbol() (string, error) {
	return _Hohm.Contract.Symbol(&_Hohm.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Hohm *HohmCallerSession) Symbol() (string, error) {
	return _Hohm.Contract.Symbol(&_Hohm.CallOpts)
}

// Teleporter is a free data retrieval call binding the contract method 0xb3ff9f49.
//
// Solidity: function teleporter() view returns(address)
func (_Hohm *HohmCaller) Teleporter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hohm.contract.Call(opts, &out, "teleporter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Teleporter is a free data retrieval call binding the contract method 0xb3ff9f49.
//
// Solidity: function teleporter() view returns(address)
func (_Hohm *HohmSession) Teleporter() (common.Address, error) {
	return _Hohm.Contract.Teleporter(&_Hohm.CallOpts)
}

// Teleporter is a free data retrieval call binding the contract method 0xb3ff9f49.
//
// Solidity: function teleporter() view returns(address)
func (_Hohm *HohmCallerSession) Teleporter() (common.Address, error) {
	return _Hohm.Contract.Teleporter(&_Hohm.CallOpts)
}

// TokenPrices is a free data retrieval call binding the contract method 0xb1e1fca4.
//
// Solidity: function tokenPrices() view returns(address)
func (_Hohm *HohmCaller) TokenPrices(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hohm.contract.Call(opts, &out, "tokenPrices")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenPrices is a free data retrieval call binding the contract method 0xb1e1fca4.
//
// Solidity: function tokenPrices() view returns(address)
func (_Hohm *HohmSession) TokenPrices() (common.Address, error) {
	return _Hohm.Contract.TokenPrices(&_Hohm.CallOpts)
}

// TokenPrices is a free data retrieval call binding the contract method 0xb1e1fca4.
//
// Solidity: function tokenPrices() view returns(address)
func (_Hohm *HohmCallerSession) TokenPrices() (common.Address, error) {
	return _Hohm.Contract.TokenPrices(&_Hohm.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() view returns(address[] assets, address[] liabilities)
func (_Hohm *HohmCaller) Tokens(opts *bind.CallOpts) (struct {
	Assets      []common.Address
	Liabilities []common.Address
}, error) {
	var out []interface{}
	err := _Hohm.contract.Call(opts, &out, "tokens")

	outstruct := new(struct {
		Assets      []common.Address
		Liabilities []common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Assets = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Liabilities = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() view returns(address[] assets, address[] liabilities)
func (_Hohm *HohmSession) Tokens() (struct {
	Assets      []common.Address
	Liabilities []common.Address
}, error) {
	return _Hohm.Contract.Tokens(&_Hohm.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() view returns(address[] assets, address[] liabilities)
func (_Hohm *HohmCallerSession) Tokens() (struct {
	Assets      []common.Address
	Liabilities []common.Address
}, error) {
	return _Hohm.Contract.Tokens(&_Hohm.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Hohm *HohmCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hohm.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Hohm *HohmSession) TotalSupply() (*big.Int, error) {
	return _Hohm.Contract.TotalSupply(&_Hohm.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Hohm *HohmCallerSession) TotalSupply() (*big.Int, error) {
	return _Hohm.Contract.TotalSupply(&_Hohm.CallOpts)
}

// AcceptOwner is a paid mutator transaction binding the contract method 0xebbc4965.
//
// Solidity: function acceptOwner() returns()
func (_Hohm *HohmTransactor) AcceptOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hohm.contract.Transact(opts, "acceptOwner")
}

// AcceptOwner is a paid mutator transaction binding the contract method 0xebbc4965.
//
// Solidity: function acceptOwner() returns()
func (_Hohm *HohmSession) AcceptOwner() (*types.Transaction, error) {
	return _Hohm.Contract.AcceptOwner(&_Hohm.TransactOpts)
}

// AcceptOwner is a paid mutator transaction binding the contract method 0xebbc4965.
//
// Solidity: function acceptOwner() returns()
func (_Hohm *HohmTransactorSession) AcceptOwner() (*types.Transaction, error) {
	return _Hohm.Contract.AcceptOwner(&_Hohm.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Hohm *HohmTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Hohm.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Hohm *HohmSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Hohm.Contract.Approve(&_Hohm.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Hohm *HohmTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Hohm.Contract.Approve(&_Hohm.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Hohm *HohmTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Hohm.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Hohm *HohmSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Hohm.Contract.Burn(&_Hohm.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Hohm *HohmTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Hohm.Contract.Burn(&_Hohm.TransactOpts, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Hohm *HohmTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Hohm.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Hohm *HohmSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Hohm.Contract.DecreaseAllowance(&_Hohm.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Hohm *HohmTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Hohm.Contract.DecreaseAllowance(&_Hohm.TransactOpts, spender, subtractedValue)
}

// DelegateVotingPower is a paid mutator transaction binding the contract method 0xf3ff955a.
//
// Solidity: function delegateVotingPower(address delegate) returns()
func (_Hohm *HohmTransactor) DelegateVotingPower(opts *bind.TransactOpts, delegate common.Address) (*types.Transaction, error) {
	return _Hohm.contract.Transact(opts, "delegateVotingPower", delegate)
}

// DelegateVotingPower is a paid mutator transaction binding the contract method 0xf3ff955a.
//
// Solidity: function delegateVotingPower(address delegate) returns()
func (_Hohm *HohmSession) DelegateVotingPower(delegate common.Address) (*types.Transaction, error) {
	return _Hohm.Contract.DelegateVotingPower(&_Hohm.TransactOpts, delegate)
}

// DelegateVotingPower is a paid mutator transaction binding the contract method 0xf3ff955a.
//
// Solidity: function delegateVotingPower(address delegate) returns()
func (_Hohm *HohmTransactorSession) DelegateVotingPower(delegate common.Address) (*types.Transaction, error) {
	return _Hohm.Contract.DelegateVotingPower(&_Hohm.TransactOpts, delegate)
}

// ExitWithShares is a paid mutator transaction binding the contract method 0x98d7e295.
//
// Solidity: function exitWithShares(uint256 shares, address receiver, address sharesOwner) returns(uint256[] assets, uint256[] liabilities)
func (_Hohm *HohmTransactor) ExitWithShares(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, sharesOwner common.Address) (*types.Transaction, error) {
	return _Hohm.contract.Transact(opts, "exitWithShares", shares, receiver, sharesOwner)
}

// ExitWithShares is a paid mutator transaction binding the contract method 0x98d7e295.
//
// Solidity: function exitWithShares(uint256 shares, address receiver, address sharesOwner) returns(uint256[] assets, uint256[] liabilities)
func (_Hohm *HohmSession) ExitWithShares(shares *big.Int, receiver common.Address, sharesOwner common.Address) (*types.Transaction, error) {
	return _Hohm.Contract.ExitWithShares(&_Hohm.TransactOpts, shares, receiver, sharesOwner)
}

// ExitWithShares is a paid mutator transaction binding the contract method 0x98d7e295.
//
// Solidity: function exitWithShares(uint256 shares, address receiver, address sharesOwner) returns(uint256[] assets, uint256[] liabilities)
func (_Hohm *HohmTransactorSession) ExitWithShares(shares *big.Int, receiver common.Address, sharesOwner common.Address) (*types.Transaction, error) {
	return _Hohm.Contract.ExitWithShares(&_Hohm.TransactOpts, shares, receiver, sharesOwner)
}

// ExitWithToken is a paid mutator transaction binding the contract method 0x2c66d203.
//
// Solidity: function exitWithToken(address tokenAddress, uint256 tokenAmount, address receiver, address sharesOwner) returns(uint256 shares, uint256[] assets, uint256[] liabilities)
func (_Hohm *HohmTransactor) ExitWithToken(opts *bind.TransactOpts, tokenAddress common.Address, tokenAmount *big.Int, receiver common.Address, sharesOwner common.Address) (*types.Transaction, error) {
	return _Hohm.contract.Transact(opts, "exitWithToken", tokenAddress, tokenAmount, receiver, sharesOwner)
}

// ExitWithToken is a paid mutator transaction binding the contract method 0x2c66d203.
//
// Solidity: function exitWithToken(address tokenAddress, uint256 tokenAmount, address receiver, address sharesOwner) returns(uint256 shares, uint256[] assets, uint256[] liabilities)
func (_Hohm *HohmSession) ExitWithToken(tokenAddress common.Address, tokenAmount *big.Int, receiver common.Address, sharesOwner common.Address) (*types.Transaction, error) {
	return _Hohm.Contract.ExitWithToken(&_Hohm.TransactOpts, tokenAddress, tokenAmount, receiver, sharesOwner)
}

// ExitWithToken is a paid mutator transaction binding the contract method 0x2c66d203.
//
// Solidity: function exitWithToken(address tokenAddress, uint256 tokenAmount, address receiver, address sharesOwner) returns(uint256 shares, uint256[] assets, uint256[] liabilities)
func (_Hohm *HohmTransactorSession) ExitWithToken(tokenAddress common.Address, tokenAmount *big.Int, receiver common.Address, sharesOwner common.Address) (*types.Transaction, error) {
	return _Hohm.Contract.ExitWithToken(&_Hohm.TransactOpts, tokenAddress, tokenAmount, receiver, sharesOwner)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Hohm *HohmTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Hohm.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Hohm *HohmSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Hohm.Contract.IncreaseAllowance(&_Hohm.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Hohm *HohmTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Hohm.Contract.IncreaseAllowance(&_Hohm.TransactOpts, spender, addedValue)
}

// JoinWithShares is a paid mutator transaction binding the contract method 0xcfa8d692.
//
// Solidity: function joinWithShares(uint256 shares, address receiver) returns(uint256[] assets, uint256[] liabilities)
func (_Hohm *HohmTransactor) JoinWithShares(opts *bind.TransactOpts, shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Hohm.contract.Transact(opts, "joinWithShares", shares, receiver)
}

// JoinWithShares is a paid mutator transaction binding the contract method 0xcfa8d692.
//
// Solidity: function joinWithShares(uint256 shares, address receiver) returns(uint256[] assets, uint256[] liabilities)
func (_Hohm *HohmSession) JoinWithShares(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Hohm.Contract.JoinWithShares(&_Hohm.TransactOpts, shares, receiver)
}

// JoinWithShares is a paid mutator transaction binding the contract method 0xcfa8d692.
//
// Solidity: function joinWithShares(uint256 shares, address receiver) returns(uint256[] assets, uint256[] liabilities)
func (_Hohm *HohmTransactorSession) JoinWithShares(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Hohm.Contract.JoinWithShares(&_Hohm.TransactOpts, shares, receiver)
}

// JoinWithToken is a paid mutator transaction binding the contract method 0x2fa7f924.
//
// Solidity: function joinWithToken(address tokenAddress, uint256 tokenAmount, address receiver) returns(uint256 shares, uint256[] assets, uint256[] liabilities)
func (_Hohm *HohmTransactor) JoinWithToken(opts *bind.TransactOpts, tokenAddress common.Address, tokenAmount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Hohm.contract.Transact(opts, "joinWithToken", tokenAddress, tokenAmount, receiver)
}

// JoinWithToken is a paid mutator transaction binding the contract method 0x2fa7f924.
//
// Solidity: function joinWithToken(address tokenAddress, uint256 tokenAmount, address receiver) returns(uint256 shares, uint256[] assets, uint256[] liabilities)
func (_Hohm *HohmSession) JoinWithToken(tokenAddress common.Address, tokenAmount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Hohm.Contract.JoinWithToken(&_Hohm.TransactOpts, tokenAddress, tokenAmount, receiver)
}

// JoinWithToken is a paid mutator transaction binding the contract method 0x2fa7f924.
//
// Solidity: function joinWithToken(address tokenAddress, uint256 tokenAmount, address receiver) returns(uint256 shares, uint256[] assets, uint256[] liabilities)
func (_Hohm *HohmTransactorSession) JoinWithToken(tokenAddress common.Address, tokenAmount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Hohm.Contract.JoinWithToken(&_Hohm.TransactOpts, tokenAddress, tokenAmount, receiver)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Hohm *HohmTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Hohm.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Hohm *HohmSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Hohm.Contract.Multicall(&_Hohm.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Hohm *HohmTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Hohm.Contract.Multicall(&_Hohm.TransactOpts, data)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Hohm *HohmTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Hohm.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Hohm *HohmSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Hohm.Contract.Permit(&_Hohm.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Hohm *HohmTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Hohm.Contract.Permit(&_Hohm.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// ProposeNewOwner is a paid mutator transaction binding the contract method 0xb1f8100d.
//
// Solidity: function proposeNewOwner(address account) returns()
func (_Hohm *HohmTransactor) ProposeNewOwner(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Hohm.contract.Transact(opts, "proposeNewOwner", account)
}

// ProposeNewOwner is a paid mutator transaction binding the contract method 0xb1f8100d.
//
// Solidity: function proposeNewOwner(address account) returns()
func (_Hohm *HohmSession) ProposeNewOwner(account common.Address) (*types.Transaction, error) {
	return _Hohm.Contract.ProposeNewOwner(&_Hohm.TransactOpts, account)
}

// ProposeNewOwner is a paid mutator transaction binding the contract method 0xb1f8100d.
//
// Solidity: function proposeNewOwner(address account) returns()
func (_Hohm *HohmTransactorSession) ProposeNewOwner(account common.Address) (*types.Transaction, error) {
	return _Hohm.Contract.ProposeNewOwner(&_Hohm.TransactOpts, account)
}

// RecoverToken is a paid mutator transaction binding the contract method 0xa7229fd9.
//
// Solidity: function recoverToken(address token, address to, uint256 amount) returns()
func (_Hohm *HohmTransactor) RecoverToken(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Hohm.contract.Transact(opts, "recoverToken", token, to, amount)
}

// RecoverToken is a paid mutator transaction binding the contract method 0xa7229fd9.
//
// Solidity: function recoverToken(address token, address to, uint256 amount) returns()
func (_Hohm *HohmSession) RecoverToken(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Hohm.Contract.RecoverToken(&_Hohm.TransactOpts, token, to, amount)
}

// RecoverToken is a paid mutator transaction binding the contract method 0xa7229fd9.
//
// Solidity: function recoverToken(address token, address to, uint256 amount) returns()
func (_Hohm *HohmTransactorSession) RecoverToken(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Hohm.Contract.RecoverToken(&_Hohm.TransactOpts, token, to, amount)
}

// RevokeOwnership is a paid mutator transaction binding the contract method 0x2b968958.
//
// Solidity: function revokeOwnership() returns()
func (_Hohm *HohmTransactor) RevokeOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hohm.contract.Transact(opts, "revokeOwnership")
}

// RevokeOwnership is a paid mutator transaction binding the contract method 0x2b968958.
//
// Solidity: function revokeOwnership() returns()
func (_Hohm *HohmSession) RevokeOwnership() (*types.Transaction, error) {
	return _Hohm.Contract.RevokeOwnership(&_Hohm.TransactOpts)
}

// RevokeOwnership is a paid mutator transaction binding the contract method 0x2b968958.
//
// Solidity: function revokeOwnership() returns()
func (_Hohm *HohmTransactorSession) RevokeOwnership() (*types.Transaction, error) {
	return _Hohm.Contract.RevokeOwnership(&_Hohm.TransactOpts)
}

// Seed is a paid mutator transaction binding the contract method 0x65310a58.
//
// Solidity: function seed(uint256[] assetAmounts, uint256[] liabilityAmounts, uint256 sharesToMint, address receiver, uint256 newMaxTotalSupply) returns()
func (_Hohm *HohmTransactor) Seed(opts *bind.TransactOpts, assetAmounts []*big.Int, liabilityAmounts []*big.Int, sharesToMint *big.Int, receiver common.Address, newMaxTotalSupply *big.Int) (*types.Transaction, error) {
	return _Hohm.contract.Transact(opts, "seed", assetAmounts, liabilityAmounts, sharesToMint, receiver, newMaxTotalSupply)
}

// Seed is a paid mutator transaction binding the contract method 0x65310a58.
//
// Solidity: function seed(uint256[] assetAmounts, uint256[] liabilityAmounts, uint256 sharesToMint, address receiver, uint256 newMaxTotalSupply) returns()
func (_Hohm *HohmSession) Seed(assetAmounts []*big.Int, liabilityAmounts []*big.Int, sharesToMint *big.Int, receiver common.Address, newMaxTotalSupply *big.Int) (*types.Transaction, error) {
	return _Hohm.Contract.Seed(&_Hohm.TransactOpts, assetAmounts, liabilityAmounts, sharesToMint, receiver, newMaxTotalSupply)
}

// Seed is a paid mutator transaction binding the contract method 0x65310a58.
//
// Solidity: function seed(uint256[] assetAmounts, uint256[] liabilityAmounts, uint256 sharesToMint, address receiver, uint256 newMaxTotalSupply) returns()
func (_Hohm *HohmTransactorSession) Seed(assetAmounts []*big.Int, liabilityAmounts []*big.Int, sharesToMint *big.Int, receiver common.Address, newMaxTotalSupply *big.Int) (*types.Transaction, error) {
	return _Hohm.Contract.Seed(&_Hohm.TransactOpts, assetAmounts, liabilityAmounts, sharesToMint, receiver, newMaxTotalSupply)
}

// Send is a paid mutator transaction binding the contract method 0xc7c7f5b3.
//
// Solidity: function send((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) sendParam, (uint256,uint256) fee, address refundAddress) payable returns((bytes32,uint64,(uint256,uint256)) msgReceipt, (uint256,uint256) oftReceipt)
func (_Hohm *HohmTransactor) Send(opts *bind.TransactOpts, sendParam SendParam, fee MessagingFee, refundAddress common.Address) (*types.Transaction, error) {
	return _Hohm.contract.Transact(opts, "send", sendParam, fee, refundAddress)
}

// Send is a paid mutator transaction binding the contract method 0xc7c7f5b3.
//
// Solidity: function send((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) sendParam, (uint256,uint256) fee, address refundAddress) payable returns((bytes32,uint64,(uint256,uint256)) msgReceipt, (uint256,uint256) oftReceipt)
func (_Hohm *HohmSession) Send(sendParam SendParam, fee MessagingFee, refundAddress common.Address) (*types.Transaction, error) {
	return _Hohm.Contract.Send(&_Hohm.TransactOpts, sendParam, fee, refundAddress)
}

// Send is a paid mutator transaction binding the contract method 0xc7c7f5b3.
//
// Solidity: function send((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) sendParam, (uint256,uint256) fee, address refundAddress) payable returns((bytes32,uint64,(uint256,uint256)) msgReceipt, (uint256,uint256) oftReceipt)
func (_Hohm *HohmTransactorSession) Send(sendParam SendParam, fee MessagingFee, refundAddress common.Address) (*types.Transaction, error) {
	return _Hohm.Contract.Send(&_Hohm.TransactOpts, sendParam, fee, refundAddress)
}

// SetExplicitAccess is a paid mutator transaction binding the contract method 0xbfccf0ec.
//
// Solidity: function setExplicitAccess(address allowedCaller, (bytes4,bool)[] access) returns()
func (_Hohm *HohmTransactor) SetExplicitAccess(opts *bind.TransactOpts, allowedCaller common.Address, access []IOrigamiElevatedAccessExplicitAccess) (*types.Transaction, error) {
	return _Hohm.contract.Transact(opts, "setExplicitAccess", allowedCaller, access)
}

// SetExplicitAccess is a paid mutator transaction binding the contract method 0xbfccf0ec.
//
// Solidity: function setExplicitAccess(address allowedCaller, (bytes4,bool)[] access) returns()
func (_Hohm *HohmSession) SetExplicitAccess(allowedCaller common.Address, access []IOrigamiElevatedAccessExplicitAccess) (*types.Transaction, error) {
	return _Hohm.Contract.SetExplicitAccess(&_Hohm.TransactOpts, allowedCaller, access)
}

// SetExplicitAccess is a paid mutator transaction binding the contract method 0xbfccf0ec.
//
// Solidity: function setExplicitAccess(address allowedCaller, (bytes4,bool)[] access) returns()
func (_Hohm *HohmTransactorSession) SetExplicitAccess(allowedCaller common.Address, access []IOrigamiElevatedAccessExplicitAccess) (*types.Transaction, error) {
	return _Hohm.Contract.SetExplicitAccess(&_Hohm.TransactOpts, allowedCaller, access)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address newManager) returns()
func (_Hohm *HohmTransactor) SetManager(opts *bind.TransactOpts, newManager common.Address) (*types.Transaction, error) {
	return _Hohm.contract.Transact(opts, "setManager", newManager)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address newManager) returns()
func (_Hohm *HohmSession) SetManager(newManager common.Address) (*types.Transaction, error) {
	return _Hohm.Contract.SetManager(&_Hohm.TransactOpts, newManager)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address newManager) returns()
func (_Hohm *HohmTransactorSession) SetManager(newManager common.Address) (*types.Transaction, error) {
	return _Hohm.Contract.SetManager(&_Hohm.TransactOpts, newManager)
}

// SetMaxTotalSupply is a paid mutator transaction binding the contract method 0x3f3e4c11.
//
// Solidity: function setMaxTotalSupply(uint256 maxTotalSupply_) returns()
func (_Hohm *HohmTransactor) SetMaxTotalSupply(opts *bind.TransactOpts, maxTotalSupply_ *big.Int) (*types.Transaction, error) {
	return _Hohm.contract.Transact(opts, "setMaxTotalSupply", maxTotalSupply_)
}

// SetMaxTotalSupply is a paid mutator transaction binding the contract method 0x3f3e4c11.
//
// Solidity: function setMaxTotalSupply(uint256 maxTotalSupply_) returns()
func (_Hohm *HohmSession) SetMaxTotalSupply(maxTotalSupply_ *big.Int) (*types.Transaction, error) {
	return _Hohm.Contract.SetMaxTotalSupply(&_Hohm.TransactOpts, maxTotalSupply_)
}

// SetMaxTotalSupply is a paid mutator transaction binding the contract method 0x3f3e4c11.
//
// Solidity: function setMaxTotalSupply(uint256 maxTotalSupply_) returns()
func (_Hohm *HohmTransactorSession) SetMaxTotalSupply(maxTotalSupply_ *big.Int) (*types.Transaction, error) {
	return _Hohm.Contract.SetMaxTotalSupply(&_Hohm.TransactOpts, maxTotalSupply_)
}

// SetTeleporter is a paid mutator transaction binding the contract method 0xa3696a09.
//
// Solidity: function setTeleporter(address newTeleporter) returns()
func (_Hohm *HohmTransactor) SetTeleporter(opts *bind.TransactOpts, newTeleporter common.Address) (*types.Transaction, error) {
	return _Hohm.contract.Transact(opts, "setTeleporter", newTeleporter)
}

// SetTeleporter is a paid mutator transaction binding the contract method 0xa3696a09.
//
// Solidity: function setTeleporter(address newTeleporter) returns()
func (_Hohm *HohmSession) SetTeleporter(newTeleporter common.Address) (*types.Transaction, error) {
	return _Hohm.Contract.SetTeleporter(&_Hohm.TransactOpts, newTeleporter)
}

// SetTeleporter is a paid mutator transaction binding the contract method 0xa3696a09.
//
// Solidity: function setTeleporter(address newTeleporter) returns()
func (_Hohm *HohmTransactorSession) SetTeleporter(newTeleporter common.Address) (*types.Transaction, error) {
	return _Hohm.Contract.SetTeleporter(&_Hohm.TransactOpts, newTeleporter)
}

// SetTokenPrices is a paid mutator transaction binding the contract method 0xbe2f5039.
//
// Solidity: function setTokenPrices(address _tokenPrices) returns()
func (_Hohm *HohmTransactor) SetTokenPrices(opts *bind.TransactOpts, _tokenPrices common.Address) (*types.Transaction, error) {
	return _Hohm.contract.Transact(opts, "setTokenPrices", _tokenPrices)
}

// SetTokenPrices is a paid mutator transaction binding the contract method 0xbe2f5039.
//
// Solidity: function setTokenPrices(address _tokenPrices) returns()
func (_Hohm *HohmSession) SetTokenPrices(_tokenPrices common.Address) (*types.Transaction, error) {
	return _Hohm.Contract.SetTokenPrices(&_Hohm.TransactOpts, _tokenPrices)
}

// SetTokenPrices is a paid mutator transaction binding the contract method 0xbe2f5039.
//
// Solidity: function setTokenPrices(address _tokenPrices) returns()
func (_Hohm *HohmTransactorSession) SetTokenPrices(_tokenPrices common.Address) (*types.Transaction, error) {
	return _Hohm.Contract.SetTokenPrices(&_Hohm.TransactOpts, _tokenPrices)
}

// SyncDelegation is a paid mutator transaction binding the contract method 0x0c5bf964.
//
// Solidity: function syncDelegation(address account) returns()
func (_Hohm *HohmTransactor) SyncDelegation(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Hohm.contract.Transact(opts, "syncDelegation", account)
}

// SyncDelegation is a paid mutator transaction binding the contract method 0x0c5bf964.
//
// Solidity: function syncDelegation(address account) returns()
func (_Hohm *HohmSession) SyncDelegation(account common.Address) (*types.Transaction, error) {
	return _Hohm.Contract.SyncDelegation(&_Hohm.TransactOpts, account)
}

// SyncDelegation is a paid mutator transaction binding the contract method 0x0c5bf964.
//
// Solidity: function syncDelegation(address account) returns()
func (_Hohm *HohmTransactorSession) SyncDelegation(account common.Address) (*types.Transaction, error) {
	return _Hohm.Contract.SyncDelegation(&_Hohm.TransactOpts, account)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Hohm *HohmTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Hohm.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Hohm *HohmSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Hohm.Contract.Transfer(&_Hohm.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Hohm *HohmTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Hohm.Contract.Transfer(&_Hohm.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Hohm *HohmTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Hohm.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Hohm *HohmSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Hohm.Contract.TransferFrom(&_Hohm.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Hohm *HohmTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Hohm.Contract.TransferFrom(&_Hohm.TransactOpts, from, to, amount)
}

// HohmApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Hohm contract.
type HohmApprovalIterator struct {
	Event *HohmApproval // Event containing the contract specifics and raw log

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
func (it *HohmApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HohmApproval)
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
		it.Event = new(HohmApproval)
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
func (it *HohmApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HohmApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HohmApproval represents a Approval event raised by the Hohm contract.
type HohmApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Hohm *HohmFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*HohmApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Hohm.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &HohmApprovalIterator{contract: _Hohm.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Hohm *HohmFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *HohmApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Hohm.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HohmApproval)
				if err := _Hohm.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Hohm *HohmFilterer) ParseApproval(log types.Log) (*HohmApproval, error) {
	event := new(HohmApproval)
	if err := _Hohm.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HohmDebtTokenSetIterator is returned from FilterDebtTokenSet and is used to iterate over the raw logs and unpacked data for DebtTokenSet events raised by the Hohm contract.
type HohmDebtTokenSetIterator struct {
	Event *HohmDebtTokenSet // Event containing the contract specifics and raw log

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
func (it *HohmDebtTokenSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HohmDebtTokenSet)
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
		it.Event = new(HohmDebtTokenSet)
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
func (it *HohmDebtTokenSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HohmDebtTokenSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HohmDebtTokenSet represents a DebtTokenSet event raised by the Hohm contract.
type HohmDebtTokenSet struct {
	DebtToken common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDebtTokenSet is a free log retrieval operation binding the contract event 0x066f10e42df49bba948e2e987181156b9451f4d4db55b013c5ccc92e6ca80c1f.
//
// Solidity: event DebtTokenSet(address indexed debtToken)
func (_Hohm *HohmFilterer) FilterDebtTokenSet(opts *bind.FilterOpts, debtToken []common.Address) (*HohmDebtTokenSetIterator, error) {

	var debtTokenRule []interface{}
	for _, debtTokenItem := range debtToken {
		debtTokenRule = append(debtTokenRule, debtTokenItem)
	}

	logs, sub, err := _Hohm.contract.FilterLogs(opts, "DebtTokenSet", debtTokenRule)
	if err != nil {
		return nil, err
	}
	return &HohmDebtTokenSetIterator{contract: _Hohm.contract, event: "DebtTokenSet", logs: logs, sub: sub}, nil
}

// WatchDebtTokenSet is a free log subscription operation binding the contract event 0x066f10e42df49bba948e2e987181156b9451f4d4db55b013c5ccc92e6ca80c1f.
//
// Solidity: event DebtTokenSet(address indexed debtToken)
func (_Hohm *HohmFilterer) WatchDebtTokenSet(opts *bind.WatchOpts, sink chan<- *HohmDebtTokenSet, debtToken []common.Address) (event.Subscription, error) {

	var debtTokenRule []interface{}
	for _, debtTokenItem := range debtToken {
		debtTokenRule = append(debtTokenRule, debtTokenItem)
	}

	logs, sub, err := _Hohm.contract.WatchLogs(opts, "DebtTokenSet", debtTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HohmDebtTokenSet)
				if err := _Hohm.contract.UnpackLog(event, "DebtTokenSet", log); err != nil {
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

// ParseDebtTokenSet is a log parse operation binding the contract event 0x066f10e42df49bba948e2e987181156b9451f4d4db55b013c5ccc92e6ca80c1f.
//
// Solidity: event DebtTokenSet(address indexed debtToken)
func (_Hohm *HohmFilterer) ParseDebtTokenSet(log types.Log) (*HohmDebtTokenSet, error) {
	event := new(HohmDebtTokenSet)
	if err := _Hohm.contract.UnpackLog(event, "DebtTokenSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HohmEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the Hohm contract.
type HohmEIP712DomainChangedIterator struct {
	Event *HohmEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *HohmEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HohmEIP712DomainChanged)
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
		it.Event = new(HohmEIP712DomainChanged)
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
func (it *HohmEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HohmEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HohmEIP712DomainChanged represents a EIP712DomainChanged event raised by the Hohm contract.
type HohmEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Hohm *HohmFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*HohmEIP712DomainChangedIterator, error) {

	logs, sub, err := _Hohm.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &HohmEIP712DomainChangedIterator{contract: _Hohm.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Hohm *HohmFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *HohmEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _Hohm.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HohmEIP712DomainChanged)
				if err := _Hohm.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Hohm *HohmFilterer) ParseEIP712DomainChanged(log types.Log) (*HohmEIP712DomainChanged, error) {
	event := new(HohmEIP712DomainChanged)
	if err := _Hohm.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HohmExitIterator is returned from FilterExit and is used to iterate over the raw logs and unpacked data for Exit events raised by the Hohm contract.
type HohmExitIterator struct {
	Event *HohmExit // Event containing the contract specifics and raw log

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
func (it *HohmExitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HohmExit)
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
		it.Event = new(HohmExit)
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
func (it *HohmExitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HohmExitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HohmExit represents a Exit event raised by the Hohm contract.
type HohmExit struct {
	Sender      common.Address
	Receiver    common.Address
	Owner       common.Address
	Assets      []*big.Int
	Liabilities []*big.Int
	Shares      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExit is a free log retrieval operation binding the contract event 0x728a5684b1f77b03e07bd227ddcef96ef31c5268eb61122aabcd651e08f752a1.
//
// Solidity: event Exit(address indexed sender, address indexed receiver, address indexed owner, uint256[] assets, uint256[] liabilities, uint256 shares)
func (_Hohm *HohmFilterer) FilterExit(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, owner []common.Address) (*HohmExitIterator, error) {

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

	logs, sub, err := _Hohm.contract.FilterLogs(opts, "Exit", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &HohmExitIterator{contract: _Hohm.contract, event: "Exit", logs: logs, sub: sub}, nil
}

// WatchExit is a free log subscription operation binding the contract event 0x728a5684b1f77b03e07bd227ddcef96ef31c5268eb61122aabcd651e08f752a1.
//
// Solidity: event Exit(address indexed sender, address indexed receiver, address indexed owner, uint256[] assets, uint256[] liabilities, uint256 shares)
func (_Hohm *HohmFilterer) WatchExit(opts *bind.WatchOpts, sink chan<- *HohmExit, sender []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Hohm.contract.WatchLogs(opts, "Exit", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HohmExit)
				if err := _Hohm.contract.UnpackLog(event, "Exit", log); err != nil {
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

// ParseExit is a log parse operation binding the contract event 0x728a5684b1f77b03e07bd227ddcef96ef31c5268eb61122aabcd651e08f752a1.
//
// Solidity: event Exit(address indexed sender, address indexed receiver, address indexed owner, uint256[] assets, uint256[] liabilities, uint256 shares)
func (_Hohm *HohmFilterer) ParseExit(log types.Log) (*HohmExit, error) {
	event := new(HohmExit)
	if err := _Hohm.contract.UnpackLog(event, "Exit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HohmExplicitAccessSetIterator is returned from FilterExplicitAccessSet and is used to iterate over the raw logs and unpacked data for ExplicitAccessSet events raised by the Hohm contract.
type HohmExplicitAccessSetIterator struct {
	Event *HohmExplicitAccessSet // Event containing the contract specifics and raw log

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
func (it *HohmExplicitAccessSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HohmExplicitAccessSet)
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
		it.Event = new(HohmExplicitAccessSet)
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
func (it *HohmExplicitAccessSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HohmExplicitAccessSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HohmExplicitAccessSet represents a ExplicitAccessSet event raised by the Hohm contract.
type HohmExplicitAccessSet struct {
	Account    common.Address
	FnSelector [4]byte
	Value      bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterExplicitAccessSet is a free log retrieval operation binding the contract event 0xf5736e75de2c751f775d4c5ed517289f77074f8c337f451ba4c0c3ed1dd7f9ad.
//
// Solidity: event ExplicitAccessSet(address indexed account, bytes4 indexed fnSelector, bool indexed value)
func (_Hohm *HohmFilterer) FilterExplicitAccessSet(opts *bind.FilterOpts, account []common.Address, fnSelector [][4]byte, value []bool) (*HohmExplicitAccessSetIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var fnSelectorRule []interface{}
	for _, fnSelectorItem := range fnSelector {
		fnSelectorRule = append(fnSelectorRule, fnSelectorItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Hohm.contract.FilterLogs(opts, "ExplicitAccessSet", accountRule, fnSelectorRule, valueRule)
	if err != nil {
		return nil, err
	}
	return &HohmExplicitAccessSetIterator{contract: _Hohm.contract, event: "ExplicitAccessSet", logs: logs, sub: sub}, nil
}

// WatchExplicitAccessSet is a free log subscription operation binding the contract event 0xf5736e75de2c751f775d4c5ed517289f77074f8c337f451ba4c0c3ed1dd7f9ad.
//
// Solidity: event ExplicitAccessSet(address indexed account, bytes4 indexed fnSelector, bool indexed value)
func (_Hohm *HohmFilterer) WatchExplicitAccessSet(opts *bind.WatchOpts, sink chan<- *HohmExplicitAccessSet, account []common.Address, fnSelector [][4]byte, value []bool) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var fnSelectorRule []interface{}
	for _, fnSelectorItem := range fnSelector {
		fnSelectorRule = append(fnSelectorRule, fnSelectorItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Hohm.contract.WatchLogs(opts, "ExplicitAccessSet", accountRule, fnSelectorRule, valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HohmExplicitAccessSet)
				if err := _Hohm.contract.UnpackLog(event, "ExplicitAccessSet", log); err != nil {
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

// ParseExplicitAccessSet is a log parse operation binding the contract event 0xf5736e75de2c751f775d4c5ed517289f77074f8c337f451ba4c0c3ed1dd7f9ad.
//
// Solidity: event ExplicitAccessSet(address indexed account, bytes4 indexed fnSelector, bool indexed value)
func (_Hohm *HohmFilterer) ParseExplicitAccessSet(log types.Log) (*HohmExplicitAccessSet, error) {
	event := new(HohmExplicitAccessSet)
	if err := _Hohm.contract.UnpackLog(event, "ExplicitAccessSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HohmInKindFeesIterator is returned from FilterInKindFees and is used to iterate over the raw logs and unpacked data for InKindFees events raised by the Hohm contract.
type HohmInKindFeesIterator struct {
	Event *HohmInKindFees // Event containing the contract specifics and raw log

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
func (it *HohmInKindFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HohmInKindFees)
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
		it.Event = new(HohmInKindFees)
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
func (it *HohmInKindFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HohmInKindFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HohmInKindFees represents a InKindFees event raised by the Hohm contract.
type HohmInKindFees struct {
	FeeType   uint8
	FeeBps    *big.Int
	FeeAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterInKindFees is a free log retrieval operation binding the contract event 0x7e81c1439e2f6851efe3288a5d0ae235c1729a6272f98ed0a4b5eb7809140427.
//
// Solidity: event InKindFees(uint8 feeType, uint256 feeBps, uint256 feeAmount)
func (_Hohm *HohmFilterer) FilterInKindFees(opts *bind.FilterOpts) (*HohmInKindFeesIterator, error) {

	logs, sub, err := _Hohm.contract.FilterLogs(opts, "InKindFees")
	if err != nil {
		return nil, err
	}
	return &HohmInKindFeesIterator{contract: _Hohm.contract, event: "InKindFees", logs: logs, sub: sub}, nil
}

// WatchInKindFees is a free log subscription operation binding the contract event 0x7e81c1439e2f6851efe3288a5d0ae235c1729a6272f98ed0a4b5eb7809140427.
//
// Solidity: event InKindFees(uint8 feeType, uint256 feeBps, uint256 feeAmount)
func (_Hohm *HohmFilterer) WatchInKindFees(opts *bind.WatchOpts, sink chan<- *HohmInKindFees) (event.Subscription, error) {

	logs, sub, err := _Hohm.contract.WatchLogs(opts, "InKindFees")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HohmInKindFees)
				if err := _Hohm.contract.UnpackLog(event, "InKindFees", log); err != nil {
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

// ParseInKindFees is a log parse operation binding the contract event 0x7e81c1439e2f6851efe3288a5d0ae235c1729a6272f98ed0a4b5eb7809140427.
//
// Solidity: event InKindFees(uint8 feeType, uint256 feeBps, uint256 feeAmount)
func (_Hohm *HohmFilterer) ParseInKindFees(log types.Log) (*HohmInKindFees, error) {
	event := new(HohmInKindFees)
	if err := _Hohm.contract.UnpackLog(event, "InKindFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HohmJoinIterator is returned from FilterJoin and is used to iterate over the raw logs and unpacked data for Join events raised by the Hohm contract.
type HohmJoinIterator struct {
	Event *HohmJoin // Event containing the contract specifics and raw log

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
func (it *HohmJoinIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HohmJoin)
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
		it.Event = new(HohmJoin)
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
func (it *HohmJoinIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HohmJoinIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HohmJoin represents a Join event raised by the Hohm contract.
type HohmJoin struct {
	Sender      common.Address
	Owner       common.Address
	Assets      []*big.Int
	Liabilities []*big.Int
	Shares      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterJoin is a free log retrieval operation binding the contract event 0x8bf99c37e0a09491ed7cb29409f9fcd5a61a01752ab57bbf20d72ce465330dab.
//
// Solidity: event Join(address indexed sender, address indexed owner, uint256[] assets, uint256[] liabilities, uint256 shares)
func (_Hohm *HohmFilterer) FilterJoin(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*HohmJoinIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Hohm.contract.FilterLogs(opts, "Join", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &HohmJoinIterator{contract: _Hohm.contract, event: "Join", logs: logs, sub: sub}, nil
}

// WatchJoin is a free log subscription operation binding the contract event 0x8bf99c37e0a09491ed7cb29409f9fcd5a61a01752ab57bbf20d72ce465330dab.
//
// Solidity: event Join(address indexed sender, address indexed owner, uint256[] assets, uint256[] liabilities, uint256 shares)
func (_Hohm *HohmFilterer) WatchJoin(opts *bind.WatchOpts, sink chan<- *HohmJoin, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Hohm.contract.WatchLogs(opts, "Join", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HohmJoin)
				if err := _Hohm.contract.UnpackLog(event, "Join", log); err != nil {
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

// ParseJoin is a log parse operation binding the contract event 0x8bf99c37e0a09491ed7cb29409f9fcd5a61a01752ab57bbf20d72ce465330dab.
//
// Solidity: event Join(address indexed sender, address indexed owner, uint256[] assets, uint256[] liabilities, uint256 shares)
func (_Hohm *HohmFilterer) ParseJoin(log types.Log) (*HohmJoin, error) {
	event := new(HohmJoin)
	if err := _Hohm.contract.UnpackLog(event, "Join", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HohmManagerSetIterator is returned from FilterManagerSet and is used to iterate over the raw logs and unpacked data for ManagerSet events raised by the Hohm contract.
type HohmManagerSetIterator struct {
	Event *HohmManagerSet // Event containing the contract specifics and raw log

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
func (it *HohmManagerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HohmManagerSet)
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
		it.Event = new(HohmManagerSet)
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
func (it *HohmManagerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HohmManagerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HohmManagerSet represents a ManagerSet event raised by the Hohm contract.
type HohmManagerSet struct {
	Manager common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterManagerSet is a free log retrieval operation binding the contract event 0x60a0f5b9f9e81e98216071b85826681c796256fe3d1354ecb675580fba64fa69.
//
// Solidity: event ManagerSet(address indexed manager)
func (_Hohm *HohmFilterer) FilterManagerSet(opts *bind.FilterOpts, manager []common.Address) (*HohmManagerSetIterator, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _Hohm.contract.FilterLogs(opts, "ManagerSet", managerRule)
	if err != nil {
		return nil, err
	}
	return &HohmManagerSetIterator{contract: _Hohm.contract, event: "ManagerSet", logs: logs, sub: sub}, nil
}

// WatchManagerSet is a free log subscription operation binding the contract event 0x60a0f5b9f9e81e98216071b85826681c796256fe3d1354ecb675580fba64fa69.
//
// Solidity: event ManagerSet(address indexed manager)
func (_Hohm *HohmFilterer) WatchManagerSet(opts *bind.WatchOpts, sink chan<- *HohmManagerSet, manager []common.Address) (event.Subscription, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _Hohm.contract.WatchLogs(opts, "ManagerSet", managerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HohmManagerSet)
				if err := _Hohm.contract.UnpackLog(event, "ManagerSet", log); err != nil {
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

// ParseManagerSet is a log parse operation binding the contract event 0x60a0f5b9f9e81e98216071b85826681c796256fe3d1354ecb675580fba64fa69.
//
// Solidity: event ManagerSet(address indexed manager)
func (_Hohm *HohmFilterer) ParseManagerSet(log types.Log) (*HohmManagerSet, error) {
	event := new(HohmManagerSet)
	if err := _Hohm.contract.UnpackLog(event, "ManagerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HohmMaxTotalSupplySetIterator is returned from FilterMaxTotalSupplySet and is used to iterate over the raw logs and unpacked data for MaxTotalSupplySet events raised by the Hohm contract.
type HohmMaxTotalSupplySetIterator struct {
	Event *HohmMaxTotalSupplySet // Event containing the contract specifics and raw log

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
func (it *HohmMaxTotalSupplySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HohmMaxTotalSupplySet)
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
		it.Event = new(HohmMaxTotalSupplySet)
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
func (it *HohmMaxTotalSupplySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HohmMaxTotalSupplySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HohmMaxTotalSupplySet represents a MaxTotalSupplySet event raised by the Hohm contract.
type HohmMaxTotalSupplySet struct {
	MaxTotalSupply *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMaxTotalSupplySet is a free log retrieval operation binding the contract event 0x0120f799fc820eabb910038e9cce6e8024add369b4d780181846e300df284484.
//
// Solidity: event MaxTotalSupplySet(uint256 maxTotalSupply)
func (_Hohm *HohmFilterer) FilterMaxTotalSupplySet(opts *bind.FilterOpts) (*HohmMaxTotalSupplySetIterator, error) {

	logs, sub, err := _Hohm.contract.FilterLogs(opts, "MaxTotalSupplySet")
	if err != nil {
		return nil, err
	}
	return &HohmMaxTotalSupplySetIterator{contract: _Hohm.contract, event: "MaxTotalSupplySet", logs: logs, sub: sub}, nil
}

// WatchMaxTotalSupplySet is a free log subscription operation binding the contract event 0x0120f799fc820eabb910038e9cce6e8024add369b4d780181846e300df284484.
//
// Solidity: event MaxTotalSupplySet(uint256 maxTotalSupply)
func (_Hohm *HohmFilterer) WatchMaxTotalSupplySet(opts *bind.WatchOpts, sink chan<- *HohmMaxTotalSupplySet) (event.Subscription, error) {

	logs, sub, err := _Hohm.contract.WatchLogs(opts, "MaxTotalSupplySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HohmMaxTotalSupplySet)
				if err := _Hohm.contract.UnpackLog(event, "MaxTotalSupplySet", log); err != nil {
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

// ParseMaxTotalSupplySet is a log parse operation binding the contract event 0x0120f799fc820eabb910038e9cce6e8024add369b4d780181846e300df284484.
//
// Solidity: event MaxTotalSupplySet(uint256 maxTotalSupply)
func (_Hohm *HohmFilterer) ParseMaxTotalSupplySet(log types.Log) (*HohmMaxTotalSupplySet, error) {
	event := new(HohmMaxTotalSupplySet)
	if err := _Hohm.contract.UnpackLog(event, "MaxTotalSupplySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HohmNewOwnerAcceptedIterator is returned from FilterNewOwnerAccepted and is used to iterate over the raw logs and unpacked data for NewOwnerAccepted events raised by the Hohm contract.
type HohmNewOwnerAcceptedIterator struct {
	Event *HohmNewOwnerAccepted // Event containing the contract specifics and raw log

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
func (it *HohmNewOwnerAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HohmNewOwnerAccepted)
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
		it.Event = new(HohmNewOwnerAccepted)
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
func (it *HohmNewOwnerAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HohmNewOwnerAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HohmNewOwnerAccepted represents a NewOwnerAccepted event raised by the Hohm contract.
type HohmNewOwnerAccepted struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewOwnerAccepted is a free log retrieval operation binding the contract event 0x5cd6b24c0149d980c82592262b3a81294b39f8f6e3c004126aaf0828c787d554.
//
// Solidity: event NewOwnerAccepted(address indexed oldOwner, address indexed newOwner)
func (_Hohm *HohmFilterer) FilterNewOwnerAccepted(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*HohmNewOwnerAcceptedIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Hohm.contract.FilterLogs(opts, "NewOwnerAccepted", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HohmNewOwnerAcceptedIterator{contract: _Hohm.contract, event: "NewOwnerAccepted", logs: logs, sub: sub}, nil
}

// WatchNewOwnerAccepted is a free log subscription operation binding the contract event 0x5cd6b24c0149d980c82592262b3a81294b39f8f6e3c004126aaf0828c787d554.
//
// Solidity: event NewOwnerAccepted(address indexed oldOwner, address indexed newOwner)
func (_Hohm *HohmFilterer) WatchNewOwnerAccepted(opts *bind.WatchOpts, sink chan<- *HohmNewOwnerAccepted, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Hohm.contract.WatchLogs(opts, "NewOwnerAccepted", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HohmNewOwnerAccepted)
				if err := _Hohm.contract.UnpackLog(event, "NewOwnerAccepted", log); err != nil {
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

// ParseNewOwnerAccepted is a log parse operation binding the contract event 0x5cd6b24c0149d980c82592262b3a81294b39f8f6e3c004126aaf0828c787d554.
//
// Solidity: event NewOwnerAccepted(address indexed oldOwner, address indexed newOwner)
func (_Hohm *HohmFilterer) ParseNewOwnerAccepted(log types.Log) (*HohmNewOwnerAccepted, error) {
	event := new(HohmNewOwnerAccepted)
	if err := _Hohm.contract.UnpackLog(event, "NewOwnerAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HohmNewOwnerProposedIterator is returned from FilterNewOwnerProposed and is used to iterate over the raw logs and unpacked data for NewOwnerProposed events raised by the Hohm contract.
type HohmNewOwnerProposedIterator struct {
	Event *HohmNewOwnerProposed // Event containing the contract specifics and raw log

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
func (it *HohmNewOwnerProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HohmNewOwnerProposed)
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
		it.Event = new(HohmNewOwnerProposed)
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
func (it *HohmNewOwnerProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HohmNewOwnerProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HohmNewOwnerProposed represents a NewOwnerProposed event raised by the Hohm contract.
type HohmNewOwnerProposed struct {
	OldOwner         common.Address
	OldProposedOwner common.Address
	NewProposedOwner common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewOwnerProposed is a free log retrieval operation binding the contract event 0x64420d4a41c6ed4de2bccbf33192eea18e576c5b23c79c3a722d4e9534c2e8d8.
//
// Solidity: event NewOwnerProposed(address indexed oldOwner, address indexed oldProposedOwner, address indexed newProposedOwner)
func (_Hohm *HohmFilterer) FilterNewOwnerProposed(opts *bind.FilterOpts, oldOwner []common.Address, oldProposedOwner []common.Address, newProposedOwner []common.Address) (*HohmNewOwnerProposedIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var oldProposedOwnerRule []interface{}
	for _, oldProposedOwnerItem := range oldProposedOwner {
		oldProposedOwnerRule = append(oldProposedOwnerRule, oldProposedOwnerItem)
	}
	var newProposedOwnerRule []interface{}
	for _, newProposedOwnerItem := range newProposedOwner {
		newProposedOwnerRule = append(newProposedOwnerRule, newProposedOwnerItem)
	}

	logs, sub, err := _Hohm.contract.FilterLogs(opts, "NewOwnerProposed", oldOwnerRule, oldProposedOwnerRule, newProposedOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HohmNewOwnerProposedIterator{contract: _Hohm.contract, event: "NewOwnerProposed", logs: logs, sub: sub}, nil
}

// WatchNewOwnerProposed is a free log subscription operation binding the contract event 0x64420d4a41c6ed4de2bccbf33192eea18e576c5b23c79c3a722d4e9534c2e8d8.
//
// Solidity: event NewOwnerProposed(address indexed oldOwner, address indexed oldProposedOwner, address indexed newProposedOwner)
func (_Hohm *HohmFilterer) WatchNewOwnerProposed(opts *bind.WatchOpts, sink chan<- *HohmNewOwnerProposed, oldOwner []common.Address, oldProposedOwner []common.Address, newProposedOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var oldProposedOwnerRule []interface{}
	for _, oldProposedOwnerItem := range oldProposedOwner {
		oldProposedOwnerRule = append(oldProposedOwnerRule, oldProposedOwnerItem)
	}
	var newProposedOwnerRule []interface{}
	for _, newProposedOwnerItem := range newProposedOwner {
		newProposedOwnerRule = append(newProposedOwnerRule, newProposedOwnerItem)
	}

	logs, sub, err := _Hohm.contract.WatchLogs(opts, "NewOwnerProposed", oldOwnerRule, oldProposedOwnerRule, newProposedOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HohmNewOwnerProposed)
				if err := _Hohm.contract.UnpackLog(event, "NewOwnerProposed", log); err != nil {
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

// ParseNewOwnerProposed is a log parse operation binding the contract event 0x64420d4a41c6ed4de2bccbf33192eea18e576c5b23c79c3a722d4e9534c2e8d8.
//
// Solidity: event NewOwnerProposed(address indexed oldOwner, address indexed oldProposedOwner, address indexed newProposedOwner)
func (_Hohm *HohmFilterer) ParseNewOwnerProposed(log types.Log) (*HohmNewOwnerProposed, error) {
	event := new(HohmNewOwnerProposed)
	if err := _Hohm.contract.UnpackLog(event, "NewOwnerProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HohmTeleporterSetIterator is returned from FilterTeleporterSet and is used to iterate over the raw logs and unpacked data for TeleporterSet events raised by the Hohm contract.
type HohmTeleporterSetIterator struct {
	Event *HohmTeleporterSet // Event containing the contract specifics and raw log

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
func (it *HohmTeleporterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HohmTeleporterSet)
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
		it.Event = new(HohmTeleporterSet)
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
func (it *HohmTeleporterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HohmTeleporterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HohmTeleporterSet represents a TeleporterSet event raised by the Hohm contract.
type HohmTeleporterSet struct {
	Teleporter common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTeleporterSet is a free log retrieval operation binding the contract event 0x7ab5072a2f334db02e20cd6aa963f87434f813ef38fba53eea68704c7445ddb9.
//
// Solidity: event TeleporterSet(address indexed teleporter)
func (_Hohm *HohmFilterer) FilterTeleporterSet(opts *bind.FilterOpts, teleporter []common.Address) (*HohmTeleporterSetIterator, error) {

	var teleporterRule []interface{}
	for _, teleporterItem := range teleporter {
		teleporterRule = append(teleporterRule, teleporterItem)
	}

	logs, sub, err := _Hohm.contract.FilterLogs(opts, "TeleporterSet", teleporterRule)
	if err != nil {
		return nil, err
	}
	return &HohmTeleporterSetIterator{contract: _Hohm.contract, event: "TeleporterSet", logs: logs, sub: sub}, nil
}

// WatchTeleporterSet is a free log subscription operation binding the contract event 0x7ab5072a2f334db02e20cd6aa963f87434f813ef38fba53eea68704c7445ddb9.
//
// Solidity: event TeleporterSet(address indexed teleporter)
func (_Hohm *HohmFilterer) WatchTeleporterSet(opts *bind.WatchOpts, sink chan<- *HohmTeleporterSet, teleporter []common.Address) (event.Subscription, error) {

	var teleporterRule []interface{}
	for _, teleporterItem := range teleporter {
		teleporterRule = append(teleporterRule, teleporterItem)
	}

	logs, sub, err := _Hohm.contract.WatchLogs(opts, "TeleporterSet", teleporterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HohmTeleporterSet)
				if err := _Hohm.contract.UnpackLog(event, "TeleporterSet", log); err != nil {
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

// ParseTeleporterSet is a log parse operation binding the contract event 0x7ab5072a2f334db02e20cd6aa963f87434f813ef38fba53eea68704c7445ddb9.
//
// Solidity: event TeleporterSet(address indexed teleporter)
func (_Hohm *HohmFilterer) ParseTeleporterSet(log types.Log) (*HohmTeleporterSet, error) {
	event := new(HohmTeleporterSet)
	if err := _Hohm.contract.UnpackLog(event, "TeleporterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HohmTokenPricesSetIterator is returned from FilterTokenPricesSet and is used to iterate over the raw logs and unpacked data for TokenPricesSet events raised by the Hohm contract.
type HohmTokenPricesSetIterator struct {
	Event *HohmTokenPricesSet // Event containing the contract specifics and raw log

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
func (it *HohmTokenPricesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HohmTokenPricesSet)
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
		it.Event = new(HohmTokenPricesSet)
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
func (it *HohmTokenPricesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HohmTokenPricesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HohmTokenPricesSet represents a TokenPricesSet event raised by the Hohm contract.
type HohmTokenPricesSet struct {
	TokenPrices common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenPricesSet is a free log retrieval operation binding the contract event 0x2781e03d8cf8be1845f40e150af1187b0cdb48dccd761a708f5e5b612a865d1d.
//
// Solidity: event TokenPricesSet(address indexed tokenPrices)
func (_Hohm *HohmFilterer) FilterTokenPricesSet(opts *bind.FilterOpts, tokenPrices []common.Address) (*HohmTokenPricesSetIterator, error) {

	var tokenPricesRule []interface{}
	for _, tokenPricesItem := range tokenPrices {
		tokenPricesRule = append(tokenPricesRule, tokenPricesItem)
	}

	logs, sub, err := _Hohm.contract.FilterLogs(opts, "TokenPricesSet", tokenPricesRule)
	if err != nil {
		return nil, err
	}
	return &HohmTokenPricesSetIterator{contract: _Hohm.contract, event: "TokenPricesSet", logs: logs, sub: sub}, nil
}

// WatchTokenPricesSet is a free log subscription operation binding the contract event 0x2781e03d8cf8be1845f40e150af1187b0cdb48dccd761a708f5e5b612a865d1d.
//
// Solidity: event TokenPricesSet(address indexed tokenPrices)
func (_Hohm *HohmFilterer) WatchTokenPricesSet(opts *bind.WatchOpts, sink chan<- *HohmTokenPricesSet, tokenPrices []common.Address) (event.Subscription, error) {

	var tokenPricesRule []interface{}
	for _, tokenPricesItem := range tokenPrices {
		tokenPricesRule = append(tokenPricesRule, tokenPricesItem)
	}

	logs, sub, err := _Hohm.contract.WatchLogs(opts, "TokenPricesSet", tokenPricesRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HohmTokenPricesSet)
				if err := _Hohm.contract.UnpackLog(event, "TokenPricesSet", log); err != nil {
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

// ParseTokenPricesSet is a log parse operation binding the contract event 0x2781e03d8cf8be1845f40e150af1187b0cdb48dccd761a708f5e5b612a865d1d.
//
// Solidity: event TokenPricesSet(address indexed tokenPrices)
func (_Hohm *HohmFilterer) ParseTokenPricesSet(log types.Log) (*HohmTokenPricesSet, error) {
	event := new(HohmTokenPricesSet)
	if err := _Hohm.contract.UnpackLog(event, "TokenPricesSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HohmTokenRecoveredIterator is returned from FilterTokenRecovered and is used to iterate over the raw logs and unpacked data for TokenRecovered events raised by the Hohm contract.
type HohmTokenRecoveredIterator struct {
	Event *HohmTokenRecovered // Event containing the contract specifics and raw log

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
func (it *HohmTokenRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HohmTokenRecovered)
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
		it.Event = new(HohmTokenRecovered)
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
func (it *HohmTokenRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HohmTokenRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HohmTokenRecovered represents a TokenRecovered event raised by the Hohm contract.
type HohmTokenRecovered struct {
	To     common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenRecovered is a free log retrieval operation binding the contract event 0x879f92dded0f26b83c3e00b12e0395dc72cfc3077343d1854ed6988edd1f9096.
//
// Solidity: event TokenRecovered(address indexed to, address indexed token, uint256 amount)
func (_Hohm *HohmFilterer) FilterTokenRecovered(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*HohmTokenRecoveredIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Hohm.contract.FilterLogs(opts, "TokenRecovered", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &HohmTokenRecoveredIterator{contract: _Hohm.contract, event: "TokenRecovered", logs: logs, sub: sub}, nil
}

// WatchTokenRecovered is a free log subscription operation binding the contract event 0x879f92dded0f26b83c3e00b12e0395dc72cfc3077343d1854ed6988edd1f9096.
//
// Solidity: event TokenRecovered(address indexed to, address indexed token, uint256 amount)
func (_Hohm *HohmFilterer) WatchTokenRecovered(opts *bind.WatchOpts, sink chan<- *HohmTokenRecovered, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Hohm.contract.WatchLogs(opts, "TokenRecovered", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HohmTokenRecovered)
				if err := _Hohm.contract.UnpackLog(event, "TokenRecovered", log); err != nil {
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

// ParseTokenRecovered is a log parse operation binding the contract event 0x879f92dded0f26b83c3e00b12e0395dc72cfc3077343d1854ed6988edd1f9096.
//
// Solidity: event TokenRecovered(address indexed to, address indexed token, uint256 amount)
func (_Hohm *HohmFilterer) ParseTokenRecovered(log types.Log) (*HohmTokenRecovered, error) {
	event := new(HohmTokenRecovered)
	if err := _Hohm.contract.UnpackLog(event, "TokenRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HohmTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Hohm contract.
type HohmTransferIterator struct {
	Event *HohmTransfer // Event containing the contract specifics and raw log

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
func (it *HohmTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HohmTransfer)
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
		it.Event = new(HohmTransfer)
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
func (it *HohmTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HohmTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HohmTransfer represents a Transfer event raised by the Hohm contract.
type HohmTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Hohm *HohmFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*HohmTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Hohm.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &HohmTransferIterator{contract: _Hohm.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Hohm *HohmFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *HohmTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Hohm.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HohmTransfer)
				if err := _Hohm.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Hohm *HohmFilterer) ParseTransfer(log types.Log) (*HohmTransfer, error) {
	event := new(HohmTransfer)
	if err := _Hohm.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
