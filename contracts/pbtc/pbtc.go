// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pbtc

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

// EnforcedOptionParam is an auto generated low-level Go binding around an user-defined struct.
type EnforcedOptionParam struct {
	Eid     uint32
	MsgType uint16
	Options []byte
}

// InboundPacket is an auto generated low-level Go binding around an user-defined struct.
type InboundPacket struct {
	Origin    Origin
	DstEid    uint32
	Receiver  common.Address
	Guid      [32]byte
	Value     *big.Int
	Executor  common.Address
	Message   []byte
	ExtraData []byte
}

// IpBTCRequest is an auto generated low-level Go binding around an user-defined struct.
type IpBTCRequest struct {
	Requester         common.Address
	Amount            *big.Int
	BtcDepositAddress string
	BtcTxid           string
	Nonce             *big.Int
	Timestamp         *big.Int
	Status            uint8
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

// OFTFeeDetail is an auto generated low-level Go binding around an user-defined struct.
type OFTFeeDetail struct {
	FeeAmountLD *big.Int
	Description string
}

// OFTLimit is an auto generated low-level Go binding around an user-defined struct.
type OFTLimit struct {
	MinAmountLD *big.Int
	MaxAmountLD *big.Int
}

// OFTReceipt is an auto generated low-level Go binding around an user-defined struct.
type OFTReceipt struct {
	AmountSentLD     *big.Int
	AmountReceivedLD *big.Int
}

// Origin is an auto generated low-level Go binding around an user-defined struct.
type Origin struct {
	SrcEid uint32
	Sender [32]byte
	Nonce  uint64
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

// PbtcMetaData contains all meta data concerning the Pbtc contract.
var PbtcMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lzEndpoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountSD\",\"type\":\"uint256\"}],\"name\":\"AmountSDOverflowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDelegate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEndpointCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLocalDecimals\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"name\":\"InvalidOptions\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LzTokenUnavailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"}],\"name\":\"NoPeer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"}],\"name\":\"NotEnoughNative\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"OnlyEndpoint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"}],\"name\":\"OnlyPeer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlySelf\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"name\":\"SimulationResult\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountLD\",\"type\":\"uint256\"}],\"name\":\"SlippageExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"pBTC__InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"pBTC__InvalidAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"pBTC__InvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"pBTC__InvalidBtcDepositAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"pBTC__InvalidBtcTxid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"pBTC__InvalidPor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"pBTC__InvalidRequest\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"pBTC__InvalidRequestHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"pBTC__InvalidRequestInitiator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"pBTC__InvalidRequestStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"pBTC__InvalidUserAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"pBTC__NotCustodian\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"pBTC__NotWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"pBTC__UnderCollateralized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"btcDepositAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"btcTxid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestHash\",\"type\":\"bytes32\"}],\"name\":\"BurnConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"btcDepositAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestHash\",\"type\":\"bytes32\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"btcDepositAddress\",\"type\":\"string\"}],\"name\":\"CustodianBtcDepositAddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"msgType\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structEnforcedOptionParam[]\",\"name\":\"_enforcedOptions\",\"type\":\"tuple[]\"}],\"name\":\"EnforcedOptionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"btcDepositAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"btcTxid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestHash\",\"type\":\"bytes32\"}],\"name\":\"MintConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"btcDepositAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"btcTxid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestHash\",\"type\":\"bytes32\"}],\"name\":\"MintRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"btcDepositAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"btcTxid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestHash\",\"type\":\"bytes32\"}],\"name\":\"MintRequestAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestHash\",\"type\":\"bytes32\"}],\"name\":\"MintRequestCancel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"inspector\",\"type\":\"address\"}],\"name\":\"MsgInspectorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountReceivedLD\",\"type\":\"uint256\"}],\"name\":\"OFTReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountSentLD\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountReceivedLD\",\"type\":\"uint256\"}],\"name\":\"OFTSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"peer\",\"type\":\"bytes32\"}],\"name\":\"PeerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldPoR\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPoR\",\"type\":\"address\"}],\"name\":\"PoRChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"preCrimeAddress\",\"type\":\"address\"}],\"name\":\"PreCrimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"btcDepositAddress\",\"type\":\"string\"}],\"name\":\"UserBtcDepositAddressSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CUSTODIAN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SEND\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SEND_AND_CALL\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WHITELIST_MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addToWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"origin\",\"type\":\"tuple\"}],\"name\":\"allowInitializePath\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"approvalRequired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestHash\",\"type\":\"bytes32\"}],\"name\":\"cancelMintRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"_msgType\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_extraOptions\",\"type\":\"bytes\"}],\"name\":\"combineOptions\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"createBurnRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"btcTxid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"btcDepositAddress\",\"type\":\"string\"}],\"name\":\"createMintRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimalConversionRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endpoint\",\"outputs\":[{\"internalType\":\"contractILayerZeroEndpointV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"_msgType\",\"type\":\"uint16\"}],\"name\":\"enforcedOptions\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"btcTxid\",\"type\":\"string\"}],\"name\":\"executeBurnRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestHash\",\"type\":\"bytes32\"}],\"name\":\"executeMintRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"getBurnRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requestNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"btcDepositAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"btcTxid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"requestHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBurnRequestsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCirculationSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getCustodianBtcDepositAddressForUser\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"getMintRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requestNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"btcDepositAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"btcTxid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"requestHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestHash\",\"type\":\"bytes32\"}],\"name\":\"getMintRequestNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMintRequestsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestHash\",\"type\":\"bytes32\"}],\"name\":\"getPendingBurnRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"btcDepositAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"btcTxid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumIpBTC.RequestStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structIpBTC.Request\",\"name\":\"request\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestHash\",\"type\":\"bytes32\"}],\"name\":\"getPendingMintRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"btcDepositAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"btcTxid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumIpBTC.RequestStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structIpBTC.Request\",\"name\":\"request\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoR\",\"outputs\":[{\"internalType\":\"contractIPoR\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWhitelist\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"contractIPoR\",\"name\":\"_newPor\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"isComposeMsgSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_peer\",\"type\":\"bytes32\"}],\"name\":\"isPeer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_guid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"lzReceive\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"origin\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structInboundPacket[]\",\"name\":\"_packets\",\"type\":\"tuple[]\"}],\"name\":\"lzReceiveAndRevert\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_guid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"lzReceiveSimulate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"msgInspector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nextNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oApp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oAppVersion\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"senderVersion\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"receiverVersion\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oftVersion\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"}],\"name\":\"peers\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"preCrime\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountLD\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraOptions\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"composeMsg\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"oftCmd\",\"type\":\"bytes\"}],\"internalType\":\"structSendParam\",\"name\":\"_sendParam\",\"type\":\"tuple\"}],\"name\":\"quoteOFT\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minAmountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAmountLD\",\"type\":\"uint256\"}],\"internalType\":\"structOFTLimit\",\"name\":\"oftLimit\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"feeAmountLD\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structOFTFeeDetail[]\",\"name\":\"oftFeeDetails\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountSentLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountReceivedLD\",\"type\":\"uint256\"}],\"internalType\":\"structOFTReceipt\",\"name\":\"oftReceipt\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountLD\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraOptions\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"composeMsg\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"oftCmd\",\"type\":\"bytes\"}],\"internalType\":\"structSendParam\",\"name\":\"_sendParam\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"_payInLzToken\",\"type\":\"bool\"}],\"name\":\"quoteSend\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lzTokenFee\",\"type\":\"uint256\"}],\"internalType\":\"structMessagingFee\",\"name\":\"msgFee\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestHash\",\"type\":\"bytes32\"}],\"name\":\"rejectMintRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeFromWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountLD\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraOptions\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"composeMsg\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"oftCmd\",\"type\":\"bytes\"}],\"internalType\":\"structSendParam\",\"name\":\"_sendParam\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lzTokenFee\",\"type\":\"uint256\"}],\"internalType\":\"structMessagingFee\",\"name\":\"_fee\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_refundAddress\",\"type\":\"address\"}],\"name\":\"send\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lzTokenFee\",\"type\":\"uint256\"}],\"internalType\":\"structMessagingFee\",\"name\":\"fee\",\"type\":\"tuple\"}],\"internalType\":\"structMessagingReceipt\",\"name\":\"msgReceipt\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountSentLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountReceivedLD\",\"type\":\"uint256\"}],\"internalType\":\"structOFTReceipt\",\"name\":\"oftReceipt\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"btcDestinationAddress\",\"type\":\"string\"}],\"name\":\"setCustodianMintDestinationAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegate\",\"type\":\"address\"}],\"name\":\"setDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"msgType\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"internalType\":\"structEnforcedOptionParam[]\",\"name\":\"_enforcedOptions\",\"type\":\"tuple[]\"}],\"name\":\"setEnforcedOptions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_msgInspector\",\"type\":\"address\"}],\"name\":\"setMsgInspector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_peer\",\"type\":\"bytes32\"}],\"name\":\"setPeer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPoR\",\"name\":\"_newPoR\",\"type\":\"address\"}],\"name\":\"setPoR\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_preCrime\",\"type\":\"address\"}],\"name\":\"setPreCrime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"btcDestinationAddress\",\"type\":\"string\"}],\"name\":\"setUserBurnDestinationAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sharedDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PbtcABI is the input ABI used to generate the binding from.
// Deprecated: Use PbtcMetaData.ABI instead.
var PbtcABI = PbtcMetaData.ABI

// Pbtc is an auto generated Go binding around an Ethereum contract.
type Pbtc struct {
	PbtcCaller     // Read-only binding to the contract
	PbtcTransactor // Write-only binding to the contract
	PbtcFilterer   // Log filterer for contract events
}

// PbtcCaller is an auto generated read-only Go binding around an Ethereum contract.
type PbtcCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbtcTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PbtcTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbtcFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PbtcFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbtcSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PbtcSession struct {
	Contract     *Pbtc             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PbtcCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PbtcCallerSession struct {
	Contract *PbtcCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PbtcTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PbtcTransactorSession struct {
	Contract     *PbtcTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PbtcRaw is an auto generated low-level Go binding around an Ethereum contract.
type PbtcRaw struct {
	Contract *Pbtc // Generic contract binding to access the raw methods on
}

// PbtcCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PbtcCallerRaw struct {
	Contract *PbtcCaller // Generic read-only contract binding to access the raw methods on
}

// PbtcTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PbtcTransactorRaw struct {
	Contract *PbtcTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPbtc creates a new instance of Pbtc, bound to a specific deployed contract.
func NewPbtc(address common.Address, backend bind.ContractBackend) (*Pbtc, error) {
	contract, err := bindPbtc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pbtc{PbtcCaller: PbtcCaller{contract: contract}, PbtcTransactor: PbtcTransactor{contract: contract}, PbtcFilterer: PbtcFilterer{contract: contract}}, nil
}

// NewPbtcCaller creates a new read-only instance of Pbtc, bound to a specific deployed contract.
func NewPbtcCaller(address common.Address, caller bind.ContractCaller) (*PbtcCaller, error) {
	contract, err := bindPbtc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PbtcCaller{contract: contract}, nil
}

// NewPbtcTransactor creates a new write-only instance of Pbtc, bound to a specific deployed contract.
func NewPbtcTransactor(address common.Address, transactor bind.ContractTransactor) (*PbtcTransactor, error) {
	contract, err := bindPbtc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PbtcTransactor{contract: contract}, nil
}

// NewPbtcFilterer creates a new log filterer instance of Pbtc, bound to a specific deployed contract.
func NewPbtcFilterer(address common.Address, filterer bind.ContractFilterer) (*PbtcFilterer, error) {
	contract, err := bindPbtc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PbtcFilterer{contract: contract}, nil
}

// bindPbtc binds a generic wrapper to an already deployed contract.
func bindPbtc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PbtcMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pbtc *PbtcRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pbtc.Contract.PbtcCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pbtc *PbtcRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pbtc.Contract.PbtcTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pbtc *PbtcRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pbtc.Contract.PbtcTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pbtc *PbtcCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pbtc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pbtc *PbtcTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pbtc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pbtc *PbtcTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pbtc.Contract.contract.Transact(opts, method, params...)
}

// CUSTODIANROLE is a free data retrieval call binding the contract method 0xc79445d0.
//
// Solidity: function CUSTODIAN_ROLE() view returns(bytes32)
func (_Pbtc *PbtcCaller) CUSTODIANROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "CUSTODIAN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CUSTODIANROLE is a free data retrieval call binding the contract method 0xc79445d0.
//
// Solidity: function CUSTODIAN_ROLE() view returns(bytes32)
func (_Pbtc *PbtcSession) CUSTODIANROLE() ([32]byte, error) {
	return _Pbtc.Contract.CUSTODIANROLE(&_Pbtc.CallOpts)
}

// CUSTODIANROLE is a free data retrieval call binding the contract method 0xc79445d0.
//
// Solidity: function CUSTODIAN_ROLE() view returns(bytes32)
func (_Pbtc *PbtcCallerSession) CUSTODIANROLE() ([32]byte, error) {
	return _Pbtc.Contract.CUSTODIANROLE(&_Pbtc.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Pbtc *PbtcCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Pbtc *PbtcSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Pbtc.Contract.DEFAULTADMINROLE(&_Pbtc.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Pbtc *PbtcCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Pbtc.Contract.DEFAULTADMINROLE(&_Pbtc.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_Pbtc *PbtcCaller) MAXFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "MAX_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_Pbtc *PbtcSession) MAXFEE() (*big.Int, error) {
	return _Pbtc.Contract.MAXFEE(&_Pbtc.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_Pbtc *PbtcCallerSession) MAXFEE() (*big.Int, error) {
	return _Pbtc.Contract.MAXFEE(&_Pbtc.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Pbtc *PbtcCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Pbtc *PbtcSession) MINTERROLE() ([32]byte, error) {
	return _Pbtc.Contract.MINTERROLE(&_Pbtc.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Pbtc *PbtcCallerSession) MINTERROLE() ([32]byte, error) {
	return _Pbtc.Contract.MINTERROLE(&_Pbtc.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Pbtc *PbtcCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Pbtc *PbtcSession) PAUSERROLE() ([32]byte, error) {
	return _Pbtc.Contract.PAUSERROLE(&_Pbtc.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Pbtc *PbtcCallerSession) PAUSERROLE() ([32]byte, error) {
	return _Pbtc.Contract.PAUSERROLE(&_Pbtc.CallOpts)
}

// SEND is a free data retrieval call binding the contract method 0x1f5e1334.
//
// Solidity: function SEND() view returns(uint16)
func (_Pbtc *PbtcCaller) SEND(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "SEND")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// SEND is a free data retrieval call binding the contract method 0x1f5e1334.
//
// Solidity: function SEND() view returns(uint16)
func (_Pbtc *PbtcSession) SEND() (uint16, error) {
	return _Pbtc.Contract.SEND(&_Pbtc.CallOpts)
}

// SEND is a free data retrieval call binding the contract method 0x1f5e1334.
//
// Solidity: function SEND() view returns(uint16)
func (_Pbtc *PbtcCallerSession) SEND() (uint16, error) {
	return _Pbtc.Contract.SEND(&_Pbtc.CallOpts)
}

// SENDANDCALL is a free data retrieval call binding the contract method 0x134d4f25.
//
// Solidity: function SEND_AND_CALL() view returns(uint16)
func (_Pbtc *PbtcCaller) SENDANDCALL(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "SEND_AND_CALL")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// SENDANDCALL is a free data retrieval call binding the contract method 0x134d4f25.
//
// Solidity: function SEND_AND_CALL() view returns(uint16)
func (_Pbtc *PbtcSession) SENDANDCALL() (uint16, error) {
	return _Pbtc.Contract.SENDANDCALL(&_Pbtc.CallOpts)
}

// SENDANDCALL is a free data retrieval call binding the contract method 0x134d4f25.
//
// Solidity: function SEND_AND_CALL() view returns(uint16)
func (_Pbtc *PbtcCallerSession) SENDANDCALL() (uint16, error) {
	return _Pbtc.Contract.SENDANDCALL(&_Pbtc.CallOpts)
}

// USERROLE is a free data retrieval call binding the contract method 0x13119161.
//
// Solidity: function USER_ROLE() view returns(bytes32)
func (_Pbtc *PbtcCaller) USERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "USER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// USERROLE is a free data retrieval call binding the contract method 0x13119161.
//
// Solidity: function USER_ROLE() view returns(bytes32)
func (_Pbtc *PbtcSession) USERROLE() ([32]byte, error) {
	return _Pbtc.Contract.USERROLE(&_Pbtc.CallOpts)
}

// USERROLE is a free data retrieval call binding the contract method 0x13119161.
//
// Solidity: function USER_ROLE() view returns(bytes32)
func (_Pbtc *PbtcCallerSession) USERROLE() ([32]byte, error) {
	return _Pbtc.Contract.USERROLE(&_Pbtc.CallOpts)
}

// WHITELISTMANAGERROLE is a free data retrieval call binding the contract method 0x7295ed93.
//
// Solidity: function WHITELIST_MANAGER_ROLE() view returns(bytes32)
func (_Pbtc *PbtcCaller) WHITELISTMANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "WHITELIST_MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WHITELISTMANAGERROLE is a free data retrieval call binding the contract method 0x7295ed93.
//
// Solidity: function WHITELIST_MANAGER_ROLE() view returns(bytes32)
func (_Pbtc *PbtcSession) WHITELISTMANAGERROLE() ([32]byte, error) {
	return _Pbtc.Contract.WHITELISTMANAGERROLE(&_Pbtc.CallOpts)
}

// WHITELISTMANAGERROLE is a free data retrieval call binding the contract method 0x7295ed93.
//
// Solidity: function WHITELIST_MANAGER_ROLE() view returns(bytes32)
func (_Pbtc *PbtcCallerSession) WHITELISTMANAGERROLE() ([32]byte, error) {
	return _Pbtc.Contract.WHITELISTMANAGERROLE(&_Pbtc.CallOpts)
}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_Pbtc *PbtcCaller) AllowInitializePath(opts *bind.CallOpts, origin Origin) (bool, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "allowInitializePath", origin)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_Pbtc *PbtcSession) AllowInitializePath(origin Origin) (bool, error) {
	return _Pbtc.Contract.AllowInitializePath(&_Pbtc.CallOpts, origin)
}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_Pbtc *PbtcCallerSession) AllowInitializePath(origin Origin) (bool, error) {
	return _Pbtc.Contract.AllowInitializePath(&_Pbtc.CallOpts, origin)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Pbtc *PbtcCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Pbtc *PbtcSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Pbtc.Contract.Allowance(&_Pbtc.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Pbtc *PbtcCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Pbtc.Contract.Allowance(&_Pbtc.CallOpts, owner, spender)
}

// ApprovalRequired is a free data retrieval call binding the contract method 0x9f68b964.
//
// Solidity: function approvalRequired() pure returns(bool)
func (_Pbtc *PbtcCaller) ApprovalRequired(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "approvalRequired")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApprovalRequired is a free data retrieval call binding the contract method 0x9f68b964.
//
// Solidity: function approvalRequired() pure returns(bool)
func (_Pbtc *PbtcSession) ApprovalRequired() (bool, error) {
	return _Pbtc.Contract.ApprovalRequired(&_Pbtc.CallOpts)
}

// ApprovalRequired is a free data retrieval call binding the contract method 0x9f68b964.
//
// Solidity: function approvalRequired() pure returns(bool)
func (_Pbtc *PbtcCallerSession) ApprovalRequired() (bool, error) {
	return _Pbtc.Contract.ApprovalRequired(&_Pbtc.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Pbtc *PbtcCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Pbtc *PbtcSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Pbtc.Contract.BalanceOf(&_Pbtc.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Pbtc *PbtcCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Pbtc.Contract.BalanceOf(&_Pbtc.CallOpts, account)
}

// CombineOptions is a free data retrieval call binding the contract method 0xbc70b354.
//
// Solidity: function combineOptions(uint32 _eid, uint16 _msgType, bytes _extraOptions) view returns(bytes)
func (_Pbtc *PbtcCaller) CombineOptions(opts *bind.CallOpts, _eid uint32, _msgType uint16, _extraOptions []byte) ([]byte, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "combineOptions", _eid, _msgType, _extraOptions)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CombineOptions is a free data retrieval call binding the contract method 0xbc70b354.
//
// Solidity: function combineOptions(uint32 _eid, uint16 _msgType, bytes _extraOptions) view returns(bytes)
func (_Pbtc *PbtcSession) CombineOptions(_eid uint32, _msgType uint16, _extraOptions []byte) ([]byte, error) {
	return _Pbtc.Contract.CombineOptions(&_Pbtc.CallOpts, _eid, _msgType, _extraOptions)
}

// CombineOptions is a free data retrieval call binding the contract method 0xbc70b354.
//
// Solidity: function combineOptions(uint32 _eid, uint16 _msgType, bytes _extraOptions) view returns(bytes)
func (_Pbtc *PbtcCallerSession) CombineOptions(_eid uint32, _msgType uint16, _extraOptions []byte) ([]byte, error) {
	return _Pbtc.Contract.CombineOptions(&_Pbtc.CallOpts, _eid, _msgType, _extraOptions)
}

// DecimalConversionRate is a free data retrieval call binding the contract method 0x963efcaa.
//
// Solidity: function decimalConversionRate() view returns(uint256)
func (_Pbtc *PbtcCaller) DecimalConversionRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "decimalConversionRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DecimalConversionRate is a free data retrieval call binding the contract method 0x963efcaa.
//
// Solidity: function decimalConversionRate() view returns(uint256)
func (_Pbtc *PbtcSession) DecimalConversionRate() (*big.Int, error) {
	return _Pbtc.Contract.DecimalConversionRate(&_Pbtc.CallOpts)
}

// DecimalConversionRate is a free data retrieval call binding the contract method 0x963efcaa.
//
// Solidity: function decimalConversionRate() view returns(uint256)
func (_Pbtc *PbtcCallerSession) DecimalConversionRate() (*big.Int, error) {
	return _Pbtc.Contract.DecimalConversionRate(&_Pbtc.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Pbtc *PbtcCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Pbtc *PbtcSession) Decimals() (uint8, error) {
	return _Pbtc.Contract.Decimals(&_Pbtc.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Pbtc *PbtcCallerSession) Decimals() (uint8, error) {
	return _Pbtc.Contract.Decimals(&_Pbtc.CallOpts)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_Pbtc *PbtcCaller) Endpoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "endpoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_Pbtc *PbtcSession) Endpoint() (common.Address, error) {
	return _Pbtc.Contract.Endpoint(&_Pbtc.CallOpts)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_Pbtc *PbtcCallerSession) Endpoint() (common.Address, error) {
	return _Pbtc.Contract.Endpoint(&_Pbtc.CallOpts)
}

// EnforcedOptions is a free data retrieval call binding the contract method 0x5535d461.
//
// Solidity: function enforcedOptions(uint32 _eid, uint16 _msgType) view returns(bytes)
func (_Pbtc *PbtcCaller) EnforcedOptions(opts *bind.CallOpts, _eid uint32, _msgType uint16) ([]byte, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "enforcedOptions", _eid, _msgType)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EnforcedOptions is a free data retrieval call binding the contract method 0x5535d461.
//
// Solidity: function enforcedOptions(uint32 _eid, uint16 _msgType) view returns(bytes)
func (_Pbtc *PbtcSession) EnforcedOptions(_eid uint32, _msgType uint16) ([]byte, error) {
	return _Pbtc.Contract.EnforcedOptions(&_Pbtc.CallOpts, _eid, _msgType)
}

// EnforcedOptions is a free data retrieval call binding the contract method 0x5535d461.
//
// Solidity: function enforcedOptions(uint32 _eid, uint16 _msgType) view returns(bytes)
func (_Pbtc *PbtcCallerSession) EnforcedOptions(_eid uint32, _msgType uint16) ([]byte, error) {
	return _Pbtc.Contract.EnforcedOptions(&_Pbtc.CallOpts, _eid, _msgType)
}

// GetBurnRequest is a free data retrieval call binding the contract method 0xe21c40c0.
//
// Solidity: function getBurnRequest(uint256 nonce) view returns(uint256 requestNonce, address requester, uint256 amount, string btcDepositAddress, string btcTxid, uint256 timestamp, string status, bytes32 requestHash)
func (_Pbtc *PbtcCaller) GetBurnRequest(opts *bind.CallOpts, nonce *big.Int) (struct {
	RequestNonce      *big.Int
	Requester         common.Address
	Amount            *big.Int
	BtcDepositAddress string
	BtcTxid           string
	Timestamp         *big.Int
	Status            string
	RequestHash       [32]byte
}, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "getBurnRequest", nonce)

	outstruct := new(struct {
		RequestNonce      *big.Int
		Requester         common.Address
		Amount            *big.Int
		BtcDepositAddress string
		BtcTxid           string
		Timestamp         *big.Int
		Status            string
		RequestHash       [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RequestNonce = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Requester = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.BtcDepositAddress = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.BtcTxid = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[6], new(string)).(*string)
	outstruct.RequestHash = *abi.ConvertType(out[7], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// GetBurnRequest is a free data retrieval call binding the contract method 0xe21c40c0.
//
// Solidity: function getBurnRequest(uint256 nonce) view returns(uint256 requestNonce, address requester, uint256 amount, string btcDepositAddress, string btcTxid, uint256 timestamp, string status, bytes32 requestHash)
func (_Pbtc *PbtcSession) GetBurnRequest(nonce *big.Int) (struct {
	RequestNonce      *big.Int
	Requester         common.Address
	Amount            *big.Int
	BtcDepositAddress string
	BtcTxid           string
	Timestamp         *big.Int
	Status            string
	RequestHash       [32]byte
}, error) {
	return _Pbtc.Contract.GetBurnRequest(&_Pbtc.CallOpts, nonce)
}

// GetBurnRequest is a free data retrieval call binding the contract method 0xe21c40c0.
//
// Solidity: function getBurnRequest(uint256 nonce) view returns(uint256 requestNonce, address requester, uint256 amount, string btcDepositAddress, string btcTxid, uint256 timestamp, string status, bytes32 requestHash)
func (_Pbtc *PbtcCallerSession) GetBurnRequest(nonce *big.Int) (struct {
	RequestNonce      *big.Int
	Requester         common.Address
	Amount            *big.Int
	BtcDepositAddress string
	BtcTxid           string
	Timestamp         *big.Int
	Status            string
	RequestHash       [32]byte
}, error) {
	return _Pbtc.Contract.GetBurnRequest(&_Pbtc.CallOpts, nonce)
}

// GetBurnRequestsLength is a free data retrieval call binding the contract method 0x72f69a72.
//
// Solidity: function getBurnRequestsLength() view returns(uint256 length)
func (_Pbtc *PbtcCaller) GetBurnRequestsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "getBurnRequestsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBurnRequestsLength is a free data retrieval call binding the contract method 0x72f69a72.
//
// Solidity: function getBurnRequestsLength() view returns(uint256 length)
func (_Pbtc *PbtcSession) GetBurnRequestsLength() (*big.Int, error) {
	return _Pbtc.Contract.GetBurnRequestsLength(&_Pbtc.CallOpts)
}

// GetBurnRequestsLength is a free data retrieval call binding the contract method 0x72f69a72.
//
// Solidity: function getBurnRequestsLength() view returns(uint256 length)
func (_Pbtc *PbtcCallerSession) GetBurnRequestsLength() (*big.Int, error) {
	return _Pbtc.Contract.GetBurnRequestsLength(&_Pbtc.CallOpts)
}

// GetCirculationSupply is a free data retrieval call binding the contract method 0xe7a4581a.
//
// Solidity: function getCirculationSupply() view returns(uint256)
func (_Pbtc *PbtcCaller) GetCirculationSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "getCirculationSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCirculationSupply is a free data retrieval call binding the contract method 0xe7a4581a.
//
// Solidity: function getCirculationSupply() view returns(uint256)
func (_Pbtc *PbtcSession) GetCirculationSupply() (*big.Int, error) {
	return _Pbtc.Contract.GetCirculationSupply(&_Pbtc.CallOpts)
}

// GetCirculationSupply is a free data retrieval call binding the contract method 0xe7a4581a.
//
// Solidity: function getCirculationSupply() view returns(uint256)
func (_Pbtc *PbtcCallerSession) GetCirculationSupply() (*big.Int, error) {
	return _Pbtc.Contract.GetCirculationSupply(&_Pbtc.CallOpts)
}

// GetCustodianBtcDepositAddressForUser is a free data retrieval call binding the contract method 0x3981933e.
//
// Solidity: function getCustodianBtcDepositAddressForUser(address user) view returns(string)
func (_Pbtc *PbtcCaller) GetCustodianBtcDepositAddressForUser(opts *bind.CallOpts, user common.Address) (string, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "getCustodianBtcDepositAddressForUser", user)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetCustodianBtcDepositAddressForUser is a free data retrieval call binding the contract method 0x3981933e.
//
// Solidity: function getCustodianBtcDepositAddressForUser(address user) view returns(string)
func (_Pbtc *PbtcSession) GetCustodianBtcDepositAddressForUser(user common.Address) (string, error) {
	return _Pbtc.Contract.GetCustodianBtcDepositAddressForUser(&_Pbtc.CallOpts, user)
}

// GetCustodianBtcDepositAddressForUser is a free data retrieval call binding the contract method 0x3981933e.
//
// Solidity: function getCustodianBtcDepositAddressForUser(address user) view returns(string)
func (_Pbtc *PbtcCallerSession) GetCustodianBtcDepositAddressForUser(user common.Address) (string, error) {
	return _Pbtc.Contract.GetCustodianBtcDepositAddressForUser(&_Pbtc.CallOpts, user)
}

// GetMintRequest is a free data retrieval call binding the contract method 0xb69b22d8.
//
// Solidity: function getMintRequest(uint256 nonce) view returns(uint256 requestNonce, address requester, uint256 amount, string btcDepositAddress, string btcTxid, uint256 timestamp, string status, bytes32 requestHash)
func (_Pbtc *PbtcCaller) GetMintRequest(opts *bind.CallOpts, nonce *big.Int) (struct {
	RequestNonce      *big.Int
	Requester         common.Address
	Amount            *big.Int
	BtcDepositAddress string
	BtcTxid           string
	Timestamp         *big.Int
	Status            string
	RequestHash       [32]byte
}, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "getMintRequest", nonce)

	outstruct := new(struct {
		RequestNonce      *big.Int
		Requester         common.Address
		Amount            *big.Int
		BtcDepositAddress string
		BtcTxid           string
		Timestamp         *big.Int
		Status            string
		RequestHash       [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RequestNonce = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Requester = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.BtcDepositAddress = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.BtcTxid = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[6], new(string)).(*string)
	outstruct.RequestHash = *abi.ConvertType(out[7], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// GetMintRequest is a free data retrieval call binding the contract method 0xb69b22d8.
//
// Solidity: function getMintRequest(uint256 nonce) view returns(uint256 requestNonce, address requester, uint256 amount, string btcDepositAddress, string btcTxid, uint256 timestamp, string status, bytes32 requestHash)
func (_Pbtc *PbtcSession) GetMintRequest(nonce *big.Int) (struct {
	RequestNonce      *big.Int
	Requester         common.Address
	Amount            *big.Int
	BtcDepositAddress string
	BtcTxid           string
	Timestamp         *big.Int
	Status            string
	RequestHash       [32]byte
}, error) {
	return _Pbtc.Contract.GetMintRequest(&_Pbtc.CallOpts, nonce)
}

// GetMintRequest is a free data retrieval call binding the contract method 0xb69b22d8.
//
// Solidity: function getMintRequest(uint256 nonce) view returns(uint256 requestNonce, address requester, uint256 amount, string btcDepositAddress, string btcTxid, uint256 timestamp, string status, bytes32 requestHash)
func (_Pbtc *PbtcCallerSession) GetMintRequest(nonce *big.Int) (struct {
	RequestNonce      *big.Int
	Requester         common.Address
	Amount            *big.Int
	BtcDepositAddress string
	BtcTxid           string
	Timestamp         *big.Int
	Status            string
	RequestHash       [32]byte
}, error) {
	return _Pbtc.Contract.GetMintRequest(&_Pbtc.CallOpts, nonce)
}

// GetMintRequestNonce is a free data retrieval call binding the contract method 0x618caab0.
//
// Solidity: function getMintRequestNonce(bytes32 requestHash) view returns(uint256)
func (_Pbtc *PbtcCaller) GetMintRequestNonce(opts *bind.CallOpts, requestHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "getMintRequestNonce", requestHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMintRequestNonce is a free data retrieval call binding the contract method 0x618caab0.
//
// Solidity: function getMintRequestNonce(bytes32 requestHash) view returns(uint256)
func (_Pbtc *PbtcSession) GetMintRequestNonce(requestHash [32]byte) (*big.Int, error) {
	return _Pbtc.Contract.GetMintRequestNonce(&_Pbtc.CallOpts, requestHash)
}

// GetMintRequestNonce is a free data retrieval call binding the contract method 0x618caab0.
//
// Solidity: function getMintRequestNonce(bytes32 requestHash) view returns(uint256)
func (_Pbtc *PbtcCallerSession) GetMintRequestNonce(requestHash [32]byte) (*big.Int, error) {
	return _Pbtc.Contract.GetMintRequestNonce(&_Pbtc.CallOpts, requestHash)
}

// GetMintRequestsLength is a free data retrieval call binding the contract method 0x311104f3.
//
// Solidity: function getMintRequestsLength() view returns(uint256 length)
func (_Pbtc *PbtcCaller) GetMintRequestsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "getMintRequestsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMintRequestsLength is a free data retrieval call binding the contract method 0x311104f3.
//
// Solidity: function getMintRequestsLength() view returns(uint256 length)
func (_Pbtc *PbtcSession) GetMintRequestsLength() (*big.Int, error) {
	return _Pbtc.Contract.GetMintRequestsLength(&_Pbtc.CallOpts)
}

// GetMintRequestsLength is a free data retrieval call binding the contract method 0x311104f3.
//
// Solidity: function getMintRequestsLength() view returns(uint256 length)
func (_Pbtc *PbtcCallerSession) GetMintRequestsLength() (*big.Int, error) {
	return _Pbtc.Contract.GetMintRequestsLength(&_Pbtc.CallOpts)
}

// GetPendingBurnRequest is a free data retrieval call binding the contract method 0xfd3cc314.
//
// Solidity: function getPendingBurnRequest(bytes32 requestHash) view returns(uint256 nonce, (address,uint256,string,string,uint256,uint256,uint8) request)
func (_Pbtc *PbtcCaller) GetPendingBurnRequest(opts *bind.CallOpts, requestHash [32]byte) (struct {
	Nonce   *big.Int
	Request IpBTCRequest
}, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "getPendingBurnRequest", requestHash)

	outstruct := new(struct {
		Nonce   *big.Int
		Request IpBTCRequest
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Nonce = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Request = *abi.ConvertType(out[1], new(IpBTCRequest)).(*IpBTCRequest)

	return *outstruct, err

}

// GetPendingBurnRequest is a free data retrieval call binding the contract method 0xfd3cc314.
//
// Solidity: function getPendingBurnRequest(bytes32 requestHash) view returns(uint256 nonce, (address,uint256,string,string,uint256,uint256,uint8) request)
func (_Pbtc *PbtcSession) GetPendingBurnRequest(requestHash [32]byte) (struct {
	Nonce   *big.Int
	Request IpBTCRequest
}, error) {
	return _Pbtc.Contract.GetPendingBurnRequest(&_Pbtc.CallOpts, requestHash)
}

// GetPendingBurnRequest is a free data retrieval call binding the contract method 0xfd3cc314.
//
// Solidity: function getPendingBurnRequest(bytes32 requestHash) view returns(uint256 nonce, (address,uint256,string,string,uint256,uint256,uint8) request)
func (_Pbtc *PbtcCallerSession) GetPendingBurnRequest(requestHash [32]byte) (struct {
	Nonce   *big.Int
	Request IpBTCRequest
}, error) {
	return _Pbtc.Contract.GetPendingBurnRequest(&_Pbtc.CallOpts, requestHash)
}

// GetPendingMintRequest is a free data retrieval call binding the contract method 0x939e4dd6.
//
// Solidity: function getPendingMintRequest(bytes32 requestHash) view returns(uint256 nonce, (address,uint256,string,string,uint256,uint256,uint8) request)
func (_Pbtc *PbtcCaller) GetPendingMintRequest(opts *bind.CallOpts, requestHash [32]byte) (struct {
	Nonce   *big.Int
	Request IpBTCRequest
}, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "getPendingMintRequest", requestHash)

	outstruct := new(struct {
		Nonce   *big.Int
		Request IpBTCRequest
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Nonce = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Request = *abi.ConvertType(out[1], new(IpBTCRequest)).(*IpBTCRequest)

	return *outstruct, err

}

// GetPendingMintRequest is a free data retrieval call binding the contract method 0x939e4dd6.
//
// Solidity: function getPendingMintRequest(bytes32 requestHash) view returns(uint256 nonce, (address,uint256,string,string,uint256,uint256,uint8) request)
func (_Pbtc *PbtcSession) GetPendingMintRequest(requestHash [32]byte) (struct {
	Nonce   *big.Int
	Request IpBTCRequest
}, error) {
	return _Pbtc.Contract.GetPendingMintRequest(&_Pbtc.CallOpts, requestHash)
}

// GetPendingMintRequest is a free data retrieval call binding the contract method 0x939e4dd6.
//
// Solidity: function getPendingMintRequest(bytes32 requestHash) view returns(uint256 nonce, (address,uint256,string,string,uint256,uint256,uint8) request)
func (_Pbtc *PbtcCallerSession) GetPendingMintRequest(requestHash [32]byte) (struct {
	Nonce   *big.Int
	Request IpBTCRequest
}, error) {
	return _Pbtc.Contract.GetPendingMintRequest(&_Pbtc.CallOpts, requestHash)
}

// GetPoR is a free data retrieval call binding the contract method 0xf7982ee5.
//
// Solidity: function getPoR() view returns(address)
func (_Pbtc *PbtcCaller) GetPoR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "getPoR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPoR is a free data retrieval call binding the contract method 0xf7982ee5.
//
// Solidity: function getPoR() view returns(address)
func (_Pbtc *PbtcSession) GetPoR() (common.Address, error) {
	return _Pbtc.Contract.GetPoR(&_Pbtc.CallOpts)
}

// GetPoR is a free data retrieval call binding the contract method 0xf7982ee5.
//
// Solidity: function getPoR() view returns(address)
func (_Pbtc *PbtcCallerSession) GetPoR() (common.Address, error) {
	return _Pbtc.Contract.GetPoR(&_Pbtc.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Pbtc *PbtcCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Pbtc *PbtcSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Pbtc.Contract.GetRoleAdmin(&_Pbtc.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Pbtc *PbtcCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Pbtc.Contract.GetRoleAdmin(&_Pbtc.CallOpts, role)
}

// GetWhitelist is a free data retrieval call binding the contract method 0xd01f63f5.
//
// Solidity: function getWhitelist() view returns(address[])
func (_Pbtc *PbtcCaller) GetWhitelist(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "getWhitelist")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetWhitelist is a free data retrieval call binding the contract method 0xd01f63f5.
//
// Solidity: function getWhitelist() view returns(address[])
func (_Pbtc *PbtcSession) GetWhitelist() ([]common.Address, error) {
	return _Pbtc.Contract.GetWhitelist(&_Pbtc.CallOpts)
}

// GetWhitelist is a free data retrieval call binding the contract method 0xd01f63f5.
//
// Solidity: function getWhitelist() view returns(address[])
func (_Pbtc *PbtcCallerSession) GetWhitelist() ([]common.Address, error) {
	return _Pbtc.Contract.GetWhitelist(&_Pbtc.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Pbtc *PbtcCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Pbtc *PbtcSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Pbtc.Contract.HasRole(&_Pbtc.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Pbtc *PbtcCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Pbtc.Contract.HasRole(&_Pbtc.CallOpts, role, account)
}

// IsComposeMsgSender is a free data retrieval call binding the contract method 0x82413eac.
//
// Solidity: function isComposeMsgSender((uint32,bytes32,uint64) , bytes , address _sender) view returns(bool)
func (_Pbtc *PbtcCaller) IsComposeMsgSender(opts *bind.CallOpts, arg0 Origin, arg1 []byte, _sender common.Address) (bool, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "isComposeMsgSender", arg0, arg1, _sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsComposeMsgSender is a free data retrieval call binding the contract method 0x82413eac.
//
// Solidity: function isComposeMsgSender((uint32,bytes32,uint64) , bytes , address _sender) view returns(bool)
func (_Pbtc *PbtcSession) IsComposeMsgSender(arg0 Origin, arg1 []byte, _sender common.Address) (bool, error) {
	return _Pbtc.Contract.IsComposeMsgSender(&_Pbtc.CallOpts, arg0, arg1, _sender)
}

// IsComposeMsgSender is a free data retrieval call binding the contract method 0x82413eac.
//
// Solidity: function isComposeMsgSender((uint32,bytes32,uint64) , bytes , address _sender) view returns(bool)
func (_Pbtc *PbtcCallerSession) IsComposeMsgSender(arg0 Origin, arg1 []byte, _sender common.Address) (bool, error) {
	return _Pbtc.Contract.IsComposeMsgSender(&_Pbtc.CallOpts, arg0, arg1, _sender)
}

// IsPeer is a free data retrieval call binding the contract method 0x5a0dfe4d.
//
// Solidity: function isPeer(uint32 _eid, bytes32 _peer) view returns(bool)
func (_Pbtc *PbtcCaller) IsPeer(opts *bind.CallOpts, _eid uint32, _peer [32]byte) (bool, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "isPeer", _eid, _peer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPeer is a free data retrieval call binding the contract method 0x5a0dfe4d.
//
// Solidity: function isPeer(uint32 _eid, bytes32 _peer) view returns(bool)
func (_Pbtc *PbtcSession) IsPeer(_eid uint32, _peer [32]byte) (bool, error) {
	return _Pbtc.Contract.IsPeer(&_Pbtc.CallOpts, _eid, _peer)
}

// IsPeer is a free data retrieval call binding the contract method 0x5a0dfe4d.
//
// Solidity: function isPeer(uint32 _eid, bytes32 _peer) view returns(bool)
func (_Pbtc *PbtcCallerSession) IsPeer(_eid uint32, _peer [32]byte) (bool, error) {
	return _Pbtc.Contract.IsPeer(&_Pbtc.CallOpts, _eid, _peer)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address account) view returns(bool)
func (_Pbtc *PbtcCaller) IsWhitelisted(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "isWhitelisted", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address account) view returns(bool)
func (_Pbtc *PbtcSession) IsWhitelisted(account common.Address) (bool, error) {
	return _Pbtc.Contract.IsWhitelisted(&_Pbtc.CallOpts, account)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address account) view returns(bool)
func (_Pbtc *PbtcCallerSession) IsWhitelisted(account common.Address) (bool, error) {
	return _Pbtc.Contract.IsWhitelisted(&_Pbtc.CallOpts, account)
}

// MsgInspector is a free data retrieval call binding the contract method 0x111ecdad.
//
// Solidity: function msgInspector() view returns(address)
func (_Pbtc *PbtcCaller) MsgInspector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "msgInspector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MsgInspector is a free data retrieval call binding the contract method 0x111ecdad.
//
// Solidity: function msgInspector() view returns(address)
func (_Pbtc *PbtcSession) MsgInspector() (common.Address, error) {
	return _Pbtc.Contract.MsgInspector(&_Pbtc.CallOpts)
}

// MsgInspector is a free data retrieval call binding the contract method 0x111ecdad.
//
// Solidity: function msgInspector() view returns(address)
func (_Pbtc *PbtcCallerSession) MsgInspector() (common.Address, error) {
	return _Pbtc.Contract.MsgInspector(&_Pbtc.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Pbtc *PbtcCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Pbtc *PbtcSession) Name() (string, error) {
	return _Pbtc.Contract.Name(&_Pbtc.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Pbtc *PbtcCallerSession) Name() (string, error) {
	return _Pbtc.Contract.Name(&_Pbtc.CallOpts)
}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 , bytes32 ) view returns(uint64 nonce)
func (_Pbtc *PbtcCaller) NextNonce(opts *bind.CallOpts, arg0 uint32, arg1 [32]byte) (uint64, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "nextNonce", arg0, arg1)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 , bytes32 ) view returns(uint64 nonce)
func (_Pbtc *PbtcSession) NextNonce(arg0 uint32, arg1 [32]byte) (uint64, error) {
	return _Pbtc.Contract.NextNonce(&_Pbtc.CallOpts, arg0, arg1)
}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 , bytes32 ) view returns(uint64 nonce)
func (_Pbtc *PbtcCallerSession) NextNonce(arg0 uint32, arg1 [32]byte) (uint64, error) {
	return _Pbtc.Contract.NextNonce(&_Pbtc.CallOpts, arg0, arg1)
}

// OApp is a free data retrieval call binding the contract method 0x52ae2879.
//
// Solidity: function oApp() view returns(address)
func (_Pbtc *PbtcCaller) OApp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "oApp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OApp is a free data retrieval call binding the contract method 0x52ae2879.
//
// Solidity: function oApp() view returns(address)
func (_Pbtc *PbtcSession) OApp() (common.Address, error) {
	return _Pbtc.Contract.OApp(&_Pbtc.CallOpts)
}

// OApp is a free data retrieval call binding the contract method 0x52ae2879.
//
// Solidity: function oApp() view returns(address)
func (_Pbtc *PbtcCallerSession) OApp() (common.Address, error) {
	return _Pbtc.Contract.OApp(&_Pbtc.CallOpts)
}

// OAppVersion is a free data retrieval call binding the contract method 0x17442b70.
//
// Solidity: function oAppVersion() pure returns(uint64 senderVersion, uint64 receiverVersion)
func (_Pbtc *PbtcCaller) OAppVersion(opts *bind.CallOpts) (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "oAppVersion")

	outstruct := new(struct {
		SenderVersion   uint64
		ReceiverVersion uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SenderVersion = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.ReceiverVersion = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// OAppVersion is a free data retrieval call binding the contract method 0x17442b70.
//
// Solidity: function oAppVersion() pure returns(uint64 senderVersion, uint64 receiverVersion)
func (_Pbtc *PbtcSession) OAppVersion() (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	return _Pbtc.Contract.OAppVersion(&_Pbtc.CallOpts)
}

// OAppVersion is a free data retrieval call binding the contract method 0x17442b70.
//
// Solidity: function oAppVersion() pure returns(uint64 senderVersion, uint64 receiverVersion)
func (_Pbtc *PbtcCallerSession) OAppVersion() (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	return _Pbtc.Contract.OAppVersion(&_Pbtc.CallOpts)
}

// OftVersion is a free data retrieval call binding the contract method 0x156a0d0f.
//
// Solidity: function oftVersion() pure returns(bytes4 interfaceId, uint64 version)
func (_Pbtc *PbtcCaller) OftVersion(opts *bind.CallOpts) (struct {
	InterfaceId [4]byte
	Version     uint64
}, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "oftVersion")

	outstruct := new(struct {
		InterfaceId [4]byte
		Version     uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.InterfaceId = *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	outstruct.Version = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// OftVersion is a free data retrieval call binding the contract method 0x156a0d0f.
//
// Solidity: function oftVersion() pure returns(bytes4 interfaceId, uint64 version)
func (_Pbtc *PbtcSession) OftVersion() (struct {
	InterfaceId [4]byte
	Version     uint64
}, error) {
	return _Pbtc.Contract.OftVersion(&_Pbtc.CallOpts)
}

// OftVersion is a free data retrieval call binding the contract method 0x156a0d0f.
//
// Solidity: function oftVersion() pure returns(bytes4 interfaceId, uint64 version)
func (_Pbtc *PbtcCallerSession) OftVersion() (struct {
	InterfaceId [4]byte
	Version     uint64
}, error) {
	return _Pbtc.Contract.OftVersion(&_Pbtc.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pbtc *PbtcCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pbtc *PbtcSession) Owner() (common.Address, error) {
	return _Pbtc.Contract.Owner(&_Pbtc.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pbtc *PbtcCallerSession) Owner() (common.Address, error) {
	return _Pbtc.Contract.Owner(&_Pbtc.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pbtc *PbtcCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pbtc *PbtcSession) Paused() (bool, error) {
	return _Pbtc.Contract.Paused(&_Pbtc.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pbtc *PbtcCallerSession) Paused() (bool, error) {
	return _Pbtc.Contract.Paused(&_Pbtc.CallOpts)
}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 _eid) view returns(bytes32)
func (_Pbtc *PbtcCaller) Peers(opts *bind.CallOpts, _eid uint32) ([32]byte, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "peers", _eid)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 _eid) view returns(bytes32)
func (_Pbtc *PbtcSession) Peers(_eid uint32) ([32]byte, error) {
	return _Pbtc.Contract.Peers(&_Pbtc.CallOpts, _eid)
}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 _eid) view returns(bytes32)
func (_Pbtc *PbtcCallerSession) Peers(_eid uint32) ([32]byte, error) {
	return _Pbtc.Contract.Peers(&_Pbtc.CallOpts, _eid)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Pbtc *PbtcCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Pbtc *PbtcSession) PendingOwner() (common.Address, error) {
	return _Pbtc.Contract.PendingOwner(&_Pbtc.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Pbtc *PbtcCallerSession) PendingOwner() (common.Address, error) {
	return _Pbtc.Contract.PendingOwner(&_Pbtc.CallOpts)
}

// PreCrime is a free data retrieval call binding the contract method 0xb731ea0a.
//
// Solidity: function preCrime() view returns(address)
func (_Pbtc *PbtcCaller) PreCrime(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "preCrime")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PreCrime is a free data retrieval call binding the contract method 0xb731ea0a.
//
// Solidity: function preCrime() view returns(address)
func (_Pbtc *PbtcSession) PreCrime() (common.Address, error) {
	return _Pbtc.Contract.PreCrime(&_Pbtc.CallOpts)
}

// PreCrime is a free data retrieval call binding the contract method 0xb731ea0a.
//
// Solidity: function preCrime() view returns(address)
func (_Pbtc *PbtcCallerSession) PreCrime() (common.Address, error) {
	return _Pbtc.Contract.PreCrime(&_Pbtc.CallOpts)
}

// QuoteOFT is a free data retrieval call binding the contract method 0x0d35b415.
//
// Solidity: function quoteOFT((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam) view returns((uint256,uint256) oftLimit, (int256,string)[] oftFeeDetails, (uint256,uint256) oftReceipt)
func (_Pbtc *PbtcCaller) QuoteOFT(opts *bind.CallOpts, _sendParam SendParam) (struct {
	OftLimit      OFTLimit
	OftFeeDetails []OFTFeeDetail
	OftReceipt    OFTReceipt
}, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "quoteOFT", _sendParam)

	outstruct := new(struct {
		OftLimit      OFTLimit
		OftFeeDetails []OFTFeeDetail
		OftReceipt    OFTReceipt
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OftLimit = *abi.ConvertType(out[0], new(OFTLimit)).(*OFTLimit)
	outstruct.OftFeeDetails = *abi.ConvertType(out[1], new([]OFTFeeDetail)).(*[]OFTFeeDetail)
	outstruct.OftReceipt = *abi.ConvertType(out[2], new(OFTReceipt)).(*OFTReceipt)

	return *outstruct, err

}

// QuoteOFT is a free data retrieval call binding the contract method 0x0d35b415.
//
// Solidity: function quoteOFT((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam) view returns((uint256,uint256) oftLimit, (int256,string)[] oftFeeDetails, (uint256,uint256) oftReceipt)
func (_Pbtc *PbtcSession) QuoteOFT(_sendParam SendParam) (struct {
	OftLimit      OFTLimit
	OftFeeDetails []OFTFeeDetail
	OftReceipt    OFTReceipt
}, error) {
	return _Pbtc.Contract.QuoteOFT(&_Pbtc.CallOpts, _sendParam)
}

// QuoteOFT is a free data retrieval call binding the contract method 0x0d35b415.
//
// Solidity: function quoteOFT((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam) view returns((uint256,uint256) oftLimit, (int256,string)[] oftFeeDetails, (uint256,uint256) oftReceipt)
func (_Pbtc *PbtcCallerSession) QuoteOFT(_sendParam SendParam) (struct {
	OftLimit      OFTLimit
	OftFeeDetails []OFTFeeDetail
	OftReceipt    OFTReceipt
}, error) {
	return _Pbtc.Contract.QuoteOFT(&_Pbtc.CallOpts, _sendParam)
}

// QuoteSend is a free data retrieval call binding the contract method 0x3b6f743b.
//
// Solidity: function quoteSend((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, bool _payInLzToken) view returns((uint256,uint256) msgFee)
func (_Pbtc *PbtcCaller) QuoteSend(opts *bind.CallOpts, _sendParam SendParam, _payInLzToken bool) (MessagingFee, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "quoteSend", _sendParam, _payInLzToken)

	if err != nil {
		return *new(MessagingFee), err
	}

	out0 := *abi.ConvertType(out[0], new(MessagingFee)).(*MessagingFee)

	return out0, err

}

// QuoteSend is a free data retrieval call binding the contract method 0x3b6f743b.
//
// Solidity: function quoteSend((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, bool _payInLzToken) view returns((uint256,uint256) msgFee)
func (_Pbtc *PbtcSession) QuoteSend(_sendParam SendParam, _payInLzToken bool) (MessagingFee, error) {
	return _Pbtc.Contract.QuoteSend(&_Pbtc.CallOpts, _sendParam, _payInLzToken)
}

// QuoteSend is a free data retrieval call binding the contract method 0x3b6f743b.
//
// Solidity: function quoteSend((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, bool _payInLzToken) view returns((uint256,uint256) msgFee)
func (_Pbtc *PbtcCallerSession) QuoteSend(_sendParam SendParam, _payInLzToken bool) (MessagingFee, error) {
	return _Pbtc.Contract.QuoteSend(&_Pbtc.CallOpts, _sendParam, _payInLzToken)
}

// SharedDecimals is a free data retrieval call binding the contract method 0x857749b0.
//
// Solidity: function sharedDecimals() pure returns(uint8)
func (_Pbtc *PbtcCaller) SharedDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "sharedDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SharedDecimals is a free data retrieval call binding the contract method 0x857749b0.
//
// Solidity: function sharedDecimals() pure returns(uint8)
func (_Pbtc *PbtcSession) SharedDecimals() (uint8, error) {
	return _Pbtc.Contract.SharedDecimals(&_Pbtc.CallOpts)
}

// SharedDecimals is a free data retrieval call binding the contract method 0x857749b0.
//
// Solidity: function sharedDecimals() pure returns(uint8)
func (_Pbtc *PbtcCallerSession) SharedDecimals() (uint8, error) {
	return _Pbtc.Contract.SharedDecimals(&_Pbtc.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Pbtc *PbtcCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Pbtc *PbtcSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Pbtc.Contract.SupportsInterface(&_Pbtc.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Pbtc *PbtcCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Pbtc.Contract.SupportsInterface(&_Pbtc.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Pbtc *PbtcCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Pbtc *PbtcSession) Symbol() (string, error) {
	return _Pbtc.Contract.Symbol(&_Pbtc.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Pbtc *PbtcCallerSession) Symbol() (string, error) {
	return _Pbtc.Contract.Symbol(&_Pbtc.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Pbtc *PbtcCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Pbtc *PbtcSession) Token() (common.Address, error) {
	return _Pbtc.Contract.Token(&_Pbtc.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Pbtc *PbtcCallerSession) Token() (common.Address, error) {
	return _Pbtc.Contract.Token(&_Pbtc.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Pbtc *PbtcCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pbtc.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Pbtc *PbtcSession) TotalSupply() (*big.Int, error) {
	return _Pbtc.Contract.TotalSupply(&_Pbtc.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Pbtc *PbtcCallerSession) TotalSupply() (*big.Int, error) {
	return _Pbtc.Contract.TotalSupply(&_Pbtc.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Pbtc *PbtcTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pbtc.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Pbtc *PbtcSession) AcceptOwnership() (*types.Transaction, error) {
	return _Pbtc.Contract.AcceptOwnership(&_Pbtc.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Pbtc *PbtcTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Pbtc.Contract.AcceptOwnership(&_Pbtc.TransactOpts)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xe43252d7.
//
// Solidity: function addToWhitelist(address account) returns()
func (_Pbtc *PbtcTransactor) AddToWhitelist(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Pbtc.contract.Transact(opts, "addToWhitelist", account)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xe43252d7.
//
// Solidity: function addToWhitelist(address account) returns()
func (_Pbtc *PbtcSession) AddToWhitelist(account common.Address) (*types.Transaction, error) {
	return _Pbtc.Contract.AddToWhitelist(&_Pbtc.TransactOpts, account)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xe43252d7.
//
// Solidity: function addToWhitelist(address account) returns()
func (_Pbtc *PbtcTransactorSession) AddToWhitelist(account common.Address) (*types.Transaction, error) {
	return _Pbtc.Contract.AddToWhitelist(&_Pbtc.TransactOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Pbtc *PbtcTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pbtc.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Pbtc *PbtcSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pbtc.Contract.Approve(&_Pbtc.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Pbtc *PbtcTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pbtc.Contract.Approve(&_Pbtc.TransactOpts, spender, value)
}

// CancelMintRequest is a paid mutator transaction binding the contract method 0xc06e2d24.
//
// Solidity: function cancelMintRequest(bytes32 requestHash) returns()
func (_Pbtc *PbtcTransactor) CancelMintRequest(opts *bind.TransactOpts, requestHash [32]byte) (*types.Transaction, error) {
	return _Pbtc.contract.Transact(opts, "cancelMintRequest", requestHash)
}

// CancelMintRequest is a paid mutator transaction binding the contract method 0xc06e2d24.
//
// Solidity: function cancelMintRequest(bytes32 requestHash) returns()
func (_Pbtc *PbtcSession) CancelMintRequest(requestHash [32]byte) (*types.Transaction, error) {
	return _Pbtc.Contract.CancelMintRequest(&_Pbtc.TransactOpts, requestHash)
}

// CancelMintRequest is a paid mutator transaction binding the contract method 0xc06e2d24.
//
// Solidity: function cancelMintRequest(bytes32 requestHash) returns()
func (_Pbtc *PbtcTransactorSession) CancelMintRequest(requestHash [32]byte) (*types.Transaction, error) {
	return _Pbtc.Contract.CancelMintRequest(&_Pbtc.TransactOpts, requestHash)
}

// CreateBurnRequest is a paid mutator transaction binding the contract method 0xab20390a.
//
// Solidity: function createBurnRequest(uint256 amount) returns()
func (_Pbtc *PbtcTransactor) CreateBurnRequest(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Pbtc.contract.Transact(opts, "createBurnRequest", amount)
}

// CreateBurnRequest is a paid mutator transaction binding the contract method 0xab20390a.
//
// Solidity: function createBurnRequest(uint256 amount) returns()
func (_Pbtc *PbtcSession) CreateBurnRequest(amount *big.Int) (*types.Transaction, error) {
	return _Pbtc.Contract.CreateBurnRequest(&_Pbtc.TransactOpts, amount)
}

// CreateBurnRequest is a paid mutator transaction binding the contract method 0xab20390a.
//
// Solidity: function createBurnRequest(uint256 amount) returns()
func (_Pbtc *PbtcTransactorSession) CreateBurnRequest(amount *big.Int) (*types.Transaction, error) {
	return _Pbtc.Contract.CreateBurnRequest(&_Pbtc.TransactOpts, amount)
}

// CreateMintRequest is a paid mutator transaction binding the contract method 0x9be224ba.
//
// Solidity: function createMintRequest(uint256 amount, string btcTxid, string btcDepositAddress) returns()
func (_Pbtc *PbtcTransactor) CreateMintRequest(opts *bind.TransactOpts, amount *big.Int, btcTxid string, btcDepositAddress string) (*types.Transaction, error) {
	return _Pbtc.contract.Transact(opts, "createMintRequest", amount, btcTxid, btcDepositAddress)
}

// CreateMintRequest is a paid mutator transaction binding the contract method 0x9be224ba.
//
// Solidity: function createMintRequest(uint256 amount, string btcTxid, string btcDepositAddress) returns()
func (_Pbtc *PbtcSession) CreateMintRequest(amount *big.Int, btcTxid string, btcDepositAddress string) (*types.Transaction, error) {
	return _Pbtc.Contract.CreateMintRequest(&_Pbtc.TransactOpts, amount, btcTxid, btcDepositAddress)
}

// CreateMintRequest is a paid mutator transaction binding the contract method 0x9be224ba.
//
// Solidity: function createMintRequest(uint256 amount, string btcTxid, string btcDepositAddress) returns()
func (_Pbtc *PbtcTransactorSession) CreateMintRequest(amount *big.Int, btcTxid string, btcDepositAddress string) (*types.Transaction, error) {
	return _Pbtc.Contract.CreateMintRequest(&_Pbtc.TransactOpts, amount, btcTxid, btcDepositAddress)
}

// ExecuteBurnRequest is a paid mutator transaction binding the contract method 0x6e6d8c45.
//
// Solidity: function executeBurnRequest(bytes32 requestHash, string btcTxid) returns()
func (_Pbtc *PbtcTransactor) ExecuteBurnRequest(opts *bind.TransactOpts, requestHash [32]byte, btcTxid string) (*types.Transaction, error) {
	return _Pbtc.contract.Transact(opts, "executeBurnRequest", requestHash, btcTxid)
}

// ExecuteBurnRequest is a paid mutator transaction binding the contract method 0x6e6d8c45.
//
// Solidity: function executeBurnRequest(bytes32 requestHash, string btcTxid) returns()
func (_Pbtc *PbtcSession) ExecuteBurnRequest(requestHash [32]byte, btcTxid string) (*types.Transaction, error) {
	return _Pbtc.Contract.ExecuteBurnRequest(&_Pbtc.TransactOpts, requestHash, btcTxid)
}

// ExecuteBurnRequest is a paid mutator transaction binding the contract method 0x6e6d8c45.
//
// Solidity: function executeBurnRequest(bytes32 requestHash, string btcTxid) returns()
func (_Pbtc *PbtcTransactorSession) ExecuteBurnRequest(requestHash [32]byte, btcTxid string) (*types.Transaction, error) {
	return _Pbtc.Contract.ExecuteBurnRequest(&_Pbtc.TransactOpts, requestHash, btcTxid)
}

// ExecuteMintRequest is a paid mutator transaction binding the contract method 0x18618324.
//
// Solidity: function executeMintRequest(bytes32 requestHash) returns()
func (_Pbtc *PbtcTransactor) ExecuteMintRequest(opts *bind.TransactOpts, requestHash [32]byte) (*types.Transaction, error) {
	return _Pbtc.contract.Transact(opts, "executeMintRequest", requestHash)
}

// ExecuteMintRequest is a paid mutator transaction binding the contract method 0x18618324.
//
// Solidity: function executeMintRequest(bytes32 requestHash) returns()
func (_Pbtc *PbtcSession) ExecuteMintRequest(requestHash [32]byte) (*types.Transaction, error) {
	return _Pbtc.Contract.ExecuteMintRequest(&_Pbtc.TransactOpts, requestHash)
}

// ExecuteMintRequest is a paid mutator transaction binding the contract method 0x18618324.
//
// Solidity: function executeMintRequest(bytes32 requestHash) returns()
func (_Pbtc *PbtcTransactorSession) ExecuteMintRequest(requestHash [32]byte) (*types.Transaction, error) {
	return _Pbtc.Contract.ExecuteMintRequest(&_Pbtc.TransactOpts, requestHash)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Pbtc *PbtcTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pbtc.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Pbtc *PbtcSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pbtc.Contract.GrantRole(&_Pbtc.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Pbtc *PbtcTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pbtc.Contract.GrantRole(&_Pbtc.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string _name, string _symbol, address _newPor) returns()
func (_Pbtc *PbtcTransactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string, _newPor common.Address) (*types.Transaction, error) {
	return _Pbtc.contract.Transact(opts, "initialize", _name, _symbol, _newPor)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string _name, string _symbol, address _newPor) returns()
func (_Pbtc *PbtcSession) Initialize(_name string, _symbol string, _newPor common.Address) (*types.Transaction, error) {
	return _Pbtc.Contract.Initialize(&_Pbtc.TransactOpts, _name, _symbol, _newPor)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string _name, string _symbol, address _newPor) returns()
func (_Pbtc *PbtcTransactorSession) Initialize(_name string, _symbol string, _newPor common.Address) (*types.Transaction, error) {
	return _Pbtc.Contract.Initialize(&_Pbtc.TransactOpts, _name, _symbol, _newPor)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_Pbtc *PbtcTransactor) LzReceive(opts *bind.TransactOpts, _origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Pbtc.contract.Transact(opts, "lzReceive", _origin, _guid, _message, _executor, _extraData)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_Pbtc *PbtcSession) LzReceive(_origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Pbtc.Contract.LzReceive(&_Pbtc.TransactOpts, _origin, _guid, _message, _executor, _extraData)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_Pbtc *PbtcTransactorSession) LzReceive(_origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Pbtc.Contract.LzReceive(&_Pbtc.TransactOpts, _origin, _guid, _message, _executor, _extraData)
}

// LzReceiveAndRevert is a paid mutator transaction binding the contract method 0xbd815db0.
//
// Solidity: function lzReceiveAndRevert(((uint32,bytes32,uint64),uint32,address,bytes32,uint256,address,bytes,bytes)[] _packets) payable returns()
func (_Pbtc *PbtcTransactor) LzReceiveAndRevert(opts *bind.TransactOpts, _packets []InboundPacket) (*types.Transaction, error) {
	return _Pbtc.contract.Transact(opts, "lzReceiveAndRevert", _packets)
}

// LzReceiveAndRevert is a paid mutator transaction binding the contract method 0xbd815db0.
//
// Solidity: function lzReceiveAndRevert(((uint32,bytes32,uint64),uint32,address,bytes32,uint256,address,bytes,bytes)[] _packets) payable returns()
func (_Pbtc *PbtcSession) LzReceiveAndRevert(_packets []InboundPacket) (*types.Transaction, error) {
	return _Pbtc.Contract.LzReceiveAndRevert(&_Pbtc.TransactOpts, _packets)
}

// LzReceiveAndRevert is a paid mutator transaction binding the contract method 0xbd815db0.
//
// Solidity: function lzReceiveAndRevert(((uint32,bytes32,uint64),uint32,address,bytes32,uint256,address,bytes,bytes)[] _packets) payable returns()
func (_Pbtc *PbtcTransactorSession) LzReceiveAndRevert(_packets []InboundPacket) (*types.Transaction, error) {
	return _Pbtc.Contract.LzReceiveAndRevert(&_Pbtc.TransactOpts, _packets)
}

// LzReceiveSimulate is a paid mutator transaction binding the contract method 0xd045a0dc.
//
// Solidity: function lzReceiveSimulate((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_Pbtc *PbtcTransactor) LzReceiveSimulate(opts *bind.TransactOpts, _origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Pbtc.contract.Transact(opts, "lzReceiveSimulate", _origin, _guid, _message, _executor, _extraData)
}

// LzReceiveSimulate is a paid mutator transaction binding the contract method 0xd045a0dc.
//
// Solidity: function lzReceiveSimulate((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_Pbtc *PbtcSession) LzReceiveSimulate(_origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Pbtc.Contract.LzReceiveSimulate(&_Pbtc.TransactOpts, _origin, _guid, _message, _executor, _extraData)
}

// LzReceiveSimulate is a paid mutator transaction binding the contract method 0xd045a0dc.
//
// Solidity: function lzReceiveSimulate((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_Pbtc *PbtcTransactorSession) LzReceiveSimulate(_origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Pbtc.Contract.LzReceiveSimulate(&_Pbtc.TransactOpts, _origin, _guid, _message, _executor, _extraData)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pbtc *PbtcTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pbtc.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pbtc *PbtcSession) Pause() (*types.Transaction, error) {
	return _Pbtc.Contract.Pause(&_Pbtc.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pbtc *PbtcTransactorSession) Pause() (*types.Transaction, error) {
	return _Pbtc.Contract.Pause(&_Pbtc.TransactOpts)
}

// RejectMintRequest is a paid mutator transaction binding the contract method 0x861f92a8.
//
// Solidity: function rejectMintRequest(bytes32 requestHash) returns()
func (_Pbtc *PbtcTransactor) RejectMintRequest(opts *bind.TransactOpts, requestHash [32]byte) (*types.Transaction, error) {
	return _Pbtc.contract.Transact(opts, "rejectMintRequest", requestHash)
}

// RejectMintRequest is a paid mutator transaction binding the contract method 0x861f92a8.
//
// Solidity: function rejectMintRequest(bytes32 requestHash) returns()
func (_Pbtc *PbtcSession) RejectMintRequest(requestHash [32]byte) (*types.Transaction, error) {
	return _Pbtc.Contract.RejectMintRequest(&_Pbtc.TransactOpts, requestHash)
}

// RejectMintRequest is a paid mutator transaction binding the contract method 0x861f92a8.
//
// Solidity: function rejectMintRequest(bytes32 requestHash) returns()
func (_Pbtc *PbtcTransactorSession) RejectMintRequest(requestHash [32]byte) (*types.Transaction, error) {
	return _Pbtc.Contract.RejectMintRequest(&_Pbtc.TransactOpts, requestHash)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(address account) returns()
func (_Pbtc *PbtcTransactor) RemoveFromWhitelist(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Pbtc.contract.Transact(opts, "removeFromWhitelist", account)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(address account) returns()
func (_Pbtc *PbtcSession) RemoveFromWhitelist(account common.Address) (*types.Transaction, error) {
	return _Pbtc.Contract.RemoveFromWhitelist(&_Pbtc.TransactOpts, account)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(address account) returns()
func (_Pbtc *PbtcTransactorSession) RemoveFromWhitelist(account common.Address) (*types.Transaction, error) {
	return _Pbtc.Contract.RemoveFromWhitelist(&_Pbtc.TransactOpts, account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pbtc *PbtcTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pbtc.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pbtc *PbtcSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pbtc.Contract.RenounceOwnership(&_Pbtc.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pbtc *PbtcTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pbtc.Contract.RenounceOwnership(&_Pbtc.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Pbtc *PbtcTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Pbtc.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Pbtc *PbtcSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Pbtc.Contract.RenounceRole(&_Pbtc.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Pbtc *PbtcTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Pbtc.Contract.RenounceRole(&_Pbtc.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Pbtc *PbtcTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pbtc.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Pbtc *PbtcSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pbtc.Contract.RevokeRole(&_Pbtc.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Pbtc *PbtcTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pbtc.Contract.RevokeRole(&_Pbtc.TransactOpts, role, account)
}

// Send is a paid mutator transaction binding the contract method 0xc7c7f5b3.
//
// Solidity: function send((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, (uint256,uint256) _fee, address _refundAddress) payable returns((bytes32,uint64,(uint256,uint256)) msgReceipt, (uint256,uint256) oftReceipt)
func (_Pbtc *PbtcTransactor) Send(opts *bind.TransactOpts, _sendParam SendParam, _fee MessagingFee, _refundAddress common.Address) (*types.Transaction, error) {
	return _Pbtc.contract.Transact(opts, "send", _sendParam, _fee, _refundAddress)
}

// Send is a paid mutator transaction binding the contract method 0xc7c7f5b3.
//
// Solidity: function send((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, (uint256,uint256) _fee, address _refundAddress) payable returns((bytes32,uint64,(uint256,uint256)) msgReceipt, (uint256,uint256) oftReceipt)
func (_Pbtc *PbtcSession) Send(_sendParam SendParam, _fee MessagingFee, _refundAddress common.Address) (*types.Transaction, error) {
	return _Pbtc.Contract.Send(&_Pbtc.TransactOpts, _sendParam, _fee, _refundAddress)
}

// Send is a paid mutator transaction binding the contract method 0xc7c7f5b3.
//
// Solidity: function send((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, (uint256,uint256) _fee, address _refundAddress) payable returns((bytes32,uint64,(uint256,uint256)) msgReceipt, (uint256,uint256) oftReceipt)
func (_Pbtc *PbtcTransactorSession) Send(_sendParam SendParam, _fee MessagingFee, _refundAddress common.Address) (*types.Transaction, error) {
	return _Pbtc.Contract.Send(&_Pbtc.TransactOpts, _sendParam, _fee, _refundAddress)
}

// SetCustodianMintDestinationAddress is a paid mutator transaction binding the contract method 0xd251075c.
//
// Solidity: function setCustodianMintDestinationAddress(address user, string btcDestinationAddress) returns()
func (_Pbtc *PbtcTransactor) SetCustodianMintDestinationAddress(opts *bind.TransactOpts, user common.Address, btcDestinationAddress string) (*types.Transaction, error) {
	return _Pbtc.contract.Transact(opts, "setCustodianMintDestinationAddress", user, btcDestinationAddress)
}

// SetCustodianMintDestinationAddress is a paid mutator transaction binding the contract method 0xd251075c.
//
// Solidity: function setCustodianMintDestinationAddress(address user, string btcDestinationAddress) returns()
func (_Pbtc *PbtcSession) SetCustodianMintDestinationAddress(user common.Address, btcDestinationAddress string) (*types.Transaction, error) {
	return _Pbtc.Contract.SetCustodianMintDestinationAddress(&_Pbtc.TransactOpts, user, btcDestinationAddress)
}

// SetCustodianMintDestinationAddress is a paid mutator transaction binding the contract method 0xd251075c.
//
// Solidity: function setCustodianMintDestinationAddress(address user, string btcDestinationAddress) returns()
func (_Pbtc *PbtcTransactorSession) SetCustodianMintDestinationAddress(user common.Address, btcDestinationAddress string) (*types.Transaction, error) {
	return _Pbtc.Contract.SetCustodianMintDestinationAddress(&_Pbtc.TransactOpts, user, btcDestinationAddress)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_Pbtc *PbtcTransactor) SetDelegate(opts *bind.TransactOpts, _delegate common.Address) (*types.Transaction, error) {
	return _Pbtc.contract.Transact(opts, "setDelegate", _delegate)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_Pbtc *PbtcSession) SetDelegate(_delegate common.Address) (*types.Transaction, error) {
	return _Pbtc.Contract.SetDelegate(&_Pbtc.TransactOpts, _delegate)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_Pbtc *PbtcTransactorSession) SetDelegate(_delegate common.Address) (*types.Transaction, error) {
	return _Pbtc.Contract.SetDelegate(&_Pbtc.TransactOpts, _delegate)
}

// SetEnforcedOptions is a paid mutator transaction binding the contract method 0xb98bd070.
//
// Solidity: function setEnforcedOptions((uint32,uint16,bytes)[] _enforcedOptions) returns()
func (_Pbtc *PbtcTransactor) SetEnforcedOptions(opts *bind.TransactOpts, _enforcedOptions []EnforcedOptionParam) (*types.Transaction, error) {
	return _Pbtc.contract.Transact(opts, "setEnforcedOptions", _enforcedOptions)
}

// SetEnforcedOptions is a paid mutator transaction binding the contract method 0xb98bd070.
//
// Solidity: function setEnforcedOptions((uint32,uint16,bytes)[] _enforcedOptions) returns()
func (_Pbtc *PbtcSession) SetEnforcedOptions(_enforcedOptions []EnforcedOptionParam) (*types.Transaction, error) {
	return _Pbtc.Contract.SetEnforcedOptions(&_Pbtc.TransactOpts, _enforcedOptions)
}

// SetEnforcedOptions is a paid mutator transaction binding the contract method 0xb98bd070.
//
// Solidity: function setEnforcedOptions((uint32,uint16,bytes)[] _enforcedOptions) returns()
func (_Pbtc *PbtcTransactorSession) SetEnforcedOptions(_enforcedOptions []EnforcedOptionParam) (*types.Transaction, error) {
	return _Pbtc.Contract.SetEnforcedOptions(&_Pbtc.TransactOpts, _enforcedOptions)
}

// SetMsgInspector is a paid mutator transaction binding the contract method 0x6fc1b31e.
//
// Solidity: function setMsgInspector(address _msgInspector) returns()
func (_Pbtc *PbtcTransactor) SetMsgInspector(opts *bind.TransactOpts, _msgInspector common.Address) (*types.Transaction, error) {
	return _Pbtc.contract.Transact(opts, "setMsgInspector", _msgInspector)
}

// SetMsgInspector is a paid mutator transaction binding the contract method 0x6fc1b31e.
//
// Solidity: function setMsgInspector(address _msgInspector) returns()
func (_Pbtc *PbtcSession) SetMsgInspector(_msgInspector common.Address) (*types.Transaction, error) {
	return _Pbtc.Contract.SetMsgInspector(&_Pbtc.TransactOpts, _msgInspector)
}

// SetMsgInspector is a paid mutator transaction binding the contract method 0x6fc1b31e.
//
// Solidity: function setMsgInspector(address _msgInspector) returns()
func (_Pbtc *PbtcTransactorSession) SetMsgInspector(_msgInspector common.Address) (*types.Transaction, error) {
	return _Pbtc.Contract.SetMsgInspector(&_Pbtc.TransactOpts, _msgInspector)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_Pbtc *PbtcTransactor) SetPeer(opts *bind.TransactOpts, _eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _Pbtc.contract.Transact(opts, "setPeer", _eid, _peer)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_Pbtc *PbtcSession) SetPeer(_eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _Pbtc.Contract.SetPeer(&_Pbtc.TransactOpts, _eid, _peer)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_Pbtc *PbtcTransactorSession) SetPeer(_eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _Pbtc.Contract.SetPeer(&_Pbtc.TransactOpts, _eid, _peer)
}

// SetPoR is a paid mutator transaction binding the contract method 0xcdac99b0.
//
// Solidity: function setPoR(address _newPoR) returns()
func (_Pbtc *PbtcTransactor) SetPoR(opts *bind.TransactOpts, _newPoR common.Address) (*types.Transaction, error) {
	return _Pbtc.contract.Transact(opts, "setPoR", _newPoR)
}

// SetPoR is a paid mutator transaction binding the contract method 0xcdac99b0.
//
// Solidity: function setPoR(address _newPoR) returns()
func (_Pbtc *PbtcSession) SetPoR(_newPoR common.Address) (*types.Transaction, error) {
	return _Pbtc.Contract.SetPoR(&_Pbtc.TransactOpts, _newPoR)
}

// SetPoR is a paid mutator transaction binding the contract method 0xcdac99b0.
//
// Solidity: function setPoR(address _newPoR) returns()
func (_Pbtc *PbtcTransactorSession) SetPoR(_newPoR common.Address) (*types.Transaction, error) {
	return _Pbtc.Contract.SetPoR(&_Pbtc.TransactOpts, _newPoR)
}

// SetPreCrime is a paid mutator transaction binding the contract method 0xd4243885.
//
// Solidity: function setPreCrime(address _preCrime) returns()
func (_Pbtc *PbtcTransactor) SetPreCrime(opts *bind.TransactOpts, _preCrime common.Address) (*types.Transaction, error) {
	return _Pbtc.contract.Transact(opts, "setPreCrime", _preCrime)
}

// SetPreCrime is a paid mutator transaction binding the contract method 0xd4243885.
//
// Solidity: function setPreCrime(address _preCrime) returns()
func (_Pbtc *PbtcSession) SetPreCrime(_preCrime common.Address) (*types.Transaction, error) {
	return _Pbtc.Contract.SetPreCrime(&_Pbtc.TransactOpts, _preCrime)
}

// SetPreCrime is a paid mutator transaction binding the contract method 0xd4243885.
//
// Solidity: function setPreCrime(address _preCrime) returns()
func (_Pbtc *PbtcTransactorSession) SetPreCrime(_preCrime common.Address) (*types.Transaction, error) {
	return _Pbtc.Contract.SetPreCrime(&_Pbtc.TransactOpts, _preCrime)
}

// SetUserBurnDestinationAddress is a paid mutator transaction binding the contract method 0xa80bb388.
//
// Solidity: function setUserBurnDestinationAddress(address user, string btcDestinationAddress) returns()
func (_Pbtc *PbtcTransactor) SetUserBurnDestinationAddress(opts *bind.TransactOpts, user common.Address, btcDestinationAddress string) (*types.Transaction, error) {
	return _Pbtc.contract.Transact(opts, "setUserBurnDestinationAddress", user, btcDestinationAddress)
}

// SetUserBurnDestinationAddress is a paid mutator transaction binding the contract method 0xa80bb388.
//
// Solidity: function setUserBurnDestinationAddress(address user, string btcDestinationAddress) returns()
func (_Pbtc *PbtcSession) SetUserBurnDestinationAddress(user common.Address, btcDestinationAddress string) (*types.Transaction, error) {
	return _Pbtc.Contract.SetUserBurnDestinationAddress(&_Pbtc.TransactOpts, user, btcDestinationAddress)
}

// SetUserBurnDestinationAddress is a paid mutator transaction binding the contract method 0xa80bb388.
//
// Solidity: function setUserBurnDestinationAddress(address user, string btcDestinationAddress) returns()
func (_Pbtc *PbtcTransactorSession) SetUserBurnDestinationAddress(user common.Address, btcDestinationAddress string) (*types.Transaction, error) {
	return _Pbtc.Contract.SetUserBurnDestinationAddress(&_Pbtc.TransactOpts, user, btcDestinationAddress)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Pbtc *PbtcTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pbtc.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Pbtc *PbtcSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pbtc.Contract.Transfer(&_Pbtc.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Pbtc *PbtcTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pbtc.Contract.Transfer(&_Pbtc.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Pbtc *PbtcTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pbtc.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Pbtc *PbtcSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pbtc.Contract.TransferFrom(&_Pbtc.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Pbtc *PbtcTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pbtc.Contract.TransferFrom(&_Pbtc.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pbtc *PbtcTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Pbtc.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pbtc *PbtcSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pbtc.Contract.TransferOwnership(&_Pbtc.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pbtc *PbtcTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pbtc.Contract.TransferOwnership(&_Pbtc.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pbtc *PbtcTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pbtc.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pbtc *PbtcSession) Unpause() (*types.Transaction, error) {
	return _Pbtc.Contract.Unpause(&_Pbtc.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pbtc *PbtcTransactorSession) Unpause() (*types.Transaction, error) {
	return _Pbtc.Contract.Unpause(&_Pbtc.TransactOpts)
}

// PbtcApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Pbtc contract.
type PbtcApprovalIterator struct {
	Event *PbtcApproval // Event containing the contract specifics and raw log

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
func (it *PbtcApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PbtcApproval)
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
		it.Event = new(PbtcApproval)
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
func (it *PbtcApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PbtcApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PbtcApproval represents a Approval event raised by the Pbtc contract.
type PbtcApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Pbtc *PbtcFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PbtcApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Pbtc.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PbtcApprovalIterator{contract: _Pbtc.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Pbtc *PbtcFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PbtcApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Pbtc.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PbtcApproval)
				if err := _Pbtc.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Pbtc *PbtcFilterer) ParseApproval(log types.Log) (*PbtcApproval, error) {
	event := new(PbtcApproval)
	if err := _Pbtc.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PbtcBurnConfirmedIterator is returned from FilterBurnConfirmed and is used to iterate over the raw logs and unpacked data for BurnConfirmed events raised by the Pbtc contract.
type PbtcBurnConfirmedIterator struct {
	Event *PbtcBurnConfirmed // Event containing the contract specifics and raw log

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
func (it *PbtcBurnConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PbtcBurnConfirmed)
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
		it.Event = new(PbtcBurnConfirmed)
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
func (it *PbtcBurnConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PbtcBurnConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PbtcBurnConfirmed represents a BurnConfirmed event raised by the Pbtc contract.
type PbtcBurnConfirmed struct {
	Nonce             *big.Int
	Requester         common.Address
	Amount            *big.Int
	BtcDepositAddress string
	BtcTxid           string
	Timestamp         *big.Int
	RequestHash       [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterBurnConfirmed is a free log retrieval operation binding the contract event 0x1949e77206780c38f7c6487c926f8a51280fcdbf63397a01a3428dbfccd2b09f.
//
// Solidity: event BurnConfirmed(uint256 indexed nonce, address indexed requester, uint256 amount, string btcDepositAddress, string btcTxid, uint256 timestamp, bytes32 requestHash)
func (_Pbtc *PbtcFilterer) FilterBurnConfirmed(opts *bind.FilterOpts, nonce []*big.Int, requester []common.Address) (*PbtcBurnConfirmedIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _Pbtc.contract.FilterLogs(opts, "BurnConfirmed", nonceRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &PbtcBurnConfirmedIterator{contract: _Pbtc.contract, event: "BurnConfirmed", logs: logs, sub: sub}, nil
}

// WatchBurnConfirmed is a free log subscription operation binding the contract event 0x1949e77206780c38f7c6487c926f8a51280fcdbf63397a01a3428dbfccd2b09f.
//
// Solidity: event BurnConfirmed(uint256 indexed nonce, address indexed requester, uint256 amount, string btcDepositAddress, string btcTxid, uint256 timestamp, bytes32 requestHash)
func (_Pbtc *PbtcFilterer) WatchBurnConfirmed(opts *bind.WatchOpts, sink chan<- *PbtcBurnConfirmed, nonce []*big.Int, requester []common.Address) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _Pbtc.contract.WatchLogs(opts, "BurnConfirmed", nonceRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PbtcBurnConfirmed)
				if err := _Pbtc.contract.UnpackLog(event, "BurnConfirmed", log); err != nil {
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

// ParseBurnConfirmed is a log parse operation binding the contract event 0x1949e77206780c38f7c6487c926f8a51280fcdbf63397a01a3428dbfccd2b09f.
//
// Solidity: event BurnConfirmed(uint256 indexed nonce, address indexed requester, uint256 amount, string btcDepositAddress, string btcTxid, uint256 timestamp, bytes32 requestHash)
func (_Pbtc *PbtcFilterer) ParseBurnConfirmed(log types.Log) (*PbtcBurnConfirmed, error) {
	event := new(PbtcBurnConfirmed)
	if err := _Pbtc.contract.UnpackLog(event, "BurnConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PbtcBurnedIterator is returned from FilterBurned and is used to iterate over the raw logs and unpacked data for Burned events raised by the Pbtc contract.
type PbtcBurnedIterator struct {
	Event *PbtcBurned // Event containing the contract specifics and raw log

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
func (it *PbtcBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PbtcBurned)
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
		it.Event = new(PbtcBurned)
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
func (it *PbtcBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PbtcBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PbtcBurned represents a Burned event raised by the Pbtc contract.
type PbtcBurned struct {
	Nonce             *big.Int
	Requester         common.Address
	Amount            *big.Int
	BtcDepositAddress string
	Timestamp         *big.Int
	RequestHash       [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterBurned is a free log retrieval operation binding the contract event 0x865e64c3fa22a0daee479fc02875d3e97d581930b9679232344d4d5dcce6a7b2.
//
// Solidity: event Burned(uint256 indexed nonce, address indexed requester, uint256 amount, string btcDepositAddress, uint256 timestamp, bytes32 requestHash)
func (_Pbtc *PbtcFilterer) FilterBurned(opts *bind.FilterOpts, nonce []*big.Int, requester []common.Address) (*PbtcBurnedIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _Pbtc.contract.FilterLogs(opts, "Burned", nonceRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &PbtcBurnedIterator{contract: _Pbtc.contract, event: "Burned", logs: logs, sub: sub}, nil
}

// WatchBurned is a free log subscription operation binding the contract event 0x865e64c3fa22a0daee479fc02875d3e97d581930b9679232344d4d5dcce6a7b2.
//
// Solidity: event Burned(uint256 indexed nonce, address indexed requester, uint256 amount, string btcDepositAddress, uint256 timestamp, bytes32 requestHash)
func (_Pbtc *PbtcFilterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *PbtcBurned, nonce []*big.Int, requester []common.Address) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _Pbtc.contract.WatchLogs(opts, "Burned", nonceRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PbtcBurned)
				if err := _Pbtc.contract.UnpackLog(event, "Burned", log); err != nil {
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

// ParseBurned is a log parse operation binding the contract event 0x865e64c3fa22a0daee479fc02875d3e97d581930b9679232344d4d5dcce6a7b2.
//
// Solidity: event Burned(uint256 indexed nonce, address indexed requester, uint256 amount, string btcDepositAddress, uint256 timestamp, bytes32 requestHash)
func (_Pbtc *PbtcFilterer) ParseBurned(log types.Log) (*PbtcBurned, error) {
	event := new(PbtcBurned)
	if err := _Pbtc.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PbtcCustodianBtcDepositAddressSetIterator is returned from FilterCustodianBtcDepositAddressSet and is used to iterate over the raw logs and unpacked data for CustodianBtcDepositAddressSet events raised by the Pbtc contract.
type PbtcCustodianBtcDepositAddressSetIterator struct {
	Event *PbtcCustodianBtcDepositAddressSet // Event containing the contract specifics and raw log

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
func (it *PbtcCustodianBtcDepositAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PbtcCustodianBtcDepositAddressSet)
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
		it.Event = new(PbtcCustodianBtcDepositAddressSet)
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
func (it *PbtcCustodianBtcDepositAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PbtcCustodianBtcDepositAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PbtcCustodianBtcDepositAddressSet represents a CustodianBtcDepositAddressSet event raised by the Pbtc contract.
type PbtcCustodianBtcDepositAddressSet struct {
	User              common.Address
	Sender            common.Address
	BtcDepositAddress string
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCustodianBtcDepositAddressSet is a free log retrieval operation binding the contract event 0xa1cea79438a06b74491693be087a2035e62acbe738749fc0ba7fc87df2eed939.
//
// Solidity: event CustodianBtcDepositAddressSet(address indexed user, address indexed sender, string btcDepositAddress)
func (_Pbtc *PbtcFilterer) FilterCustodianBtcDepositAddressSet(opts *bind.FilterOpts, user []common.Address, sender []common.Address) (*PbtcCustodianBtcDepositAddressSetIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Pbtc.contract.FilterLogs(opts, "CustodianBtcDepositAddressSet", userRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PbtcCustodianBtcDepositAddressSetIterator{contract: _Pbtc.contract, event: "CustodianBtcDepositAddressSet", logs: logs, sub: sub}, nil
}

// WatchCustodianBtcDepositAddressSet is a free log subscription operation binding the contract event 0xa1cea79438a06b74491693be087a2035e62acbe738749fc0ba7fc87df2eed939.
//
// Solidity: event CustodianBtcDepositAddressSet(address indexed user, address indexed sender, string btcDepositAddress)
func (_Pbtc *PbtcFilterer) WatchCustodianBtcDepositAddressSet(opts *bind.WatchOpts, sink chan<- *PbtcCustodianBtcDepositAddressSet, user []common.Address, sender []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Pbtc.contract.WatchLogs(opts, "CustodianBtcDepositAddressSet", userRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PbtcCustodianBtcDepositAddressSet)
				if err := _Pbtc.contract.UnpackLog(event, "CustodianBtcDepositAddressSet", log); err != nil {
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

// ParseCustodianBtcDepositAddressSet is a log parse operation binding the contract event 0xa1cea79438a06b74491693be087a2035e62acbe738749fc0ba7fc87df2eed939.
//
// Solidity: event CustodianBtcDepositAddressSet(address indexed user, address indexed sender, string btcDepositAddress)
func (_Pbtc *PbtcFilterer) ParseCustodianBtcDepositAddressSet(log types.Log) (*PbtcCustodianBtcDepositAddressSet, error) {
	event := new(PbtcCustodianBtcDepositAddressSet)
	if err := _Pbtc.contract.UnpackLog(event, "CustodianBtcDepositAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PbtcEnforcedOptionSetIterator is returned from FilterEnforcedOptionSet and is used to iterate over the raw logs and unpacked data for EnforcedOptionSet events raised by the Pbtc contract.
type PbtcEnforcedOptionSetIterator struct {
	Event *PbtcEnforcedOptionSet // Event containing the contract specifics and raw log

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
func (it *PbtcEnforcedOptionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PbtcEnforcedOptionSet)
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
		it.Event = new(PbtcEnforcedOptionSet)
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
func (it *PbtcEnforcedOptionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PbtcEnforcedOptionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PbtcEnforcedOptionSet represents a EnforcedOptionSet event raised by the Pbtc contract.
type PbtcEnforcedOptionSet struct {
	EnforcedOptions []EnforcedOptionParam
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterEnforcedOptionSet is a free log retrieval operation binding the contract event 0xbe4864a8e820971c0247f5992e2da559595f7bf076a21cb5928d443d2a13b674.
//
// Solidity: event EnforcedOptionSet((uint32,uint16,bytes)[] _enforcedOptions)
func (_Pbtc *PbtcFilterer) FilterEnforcedOptionSet(opts *bind.FilterOpts) (*PbtcEnforcedOptionSetIterator, error) {

	logs, sub, err := _Pbtc.contract.FilterLogs(opts, "EnforcedOptionSet")
	if err != nil {
		return nil, err
	}
	return &PbtcEnforcedOptionSetIterator{contract: _Pbtc.contract, event: "EnforcedOptionSet", logs: logs, sub: sub}, nil
}

// WatchEnforcedOptionSet is a free log subscription operation binding the contract event 0xbe4864a8e820971c0247f5992e2da559595f7bf076a21cb5928d443d2a13b674.
//
// Solidity: event EnforcedOptionSet((uint32,uint16,bytes)[] _enforcedOptions)
func (_Pbtc *PbtcFilterer) WatchEnforcedOptionSet(opts *bind.WatchOpts, sink chan<- *PbtcEnforcedOptionSet) (event.Subscription, error) {

	logs, sub, err := _Pbtc.contract.WatchLogs(opts, "EnforcedOptionSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PbtcEnforcedOptionSet)
				if err := _Pbtc.contract.UnpackLog(event, "EnforcedOptionSet", log); err != nil {
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

// ParseEnforcedOptionSet is a log parse operation binding the contract event 0xbe4864a8e820971c0247f5992e2da559595f7bf076a21cb5928d443d2a13b674.
//
// Solidity: event EnforcedOptionSet((uint32,uint16,bytes)[] _enforcedOptions)
func (_Pbtc *PbtcFilterer) ParseEnforcedOptionSet(log types.Log) (*PbtcEnforcedOptionSet, error) {
	event := new(PbtcEnforcedOptionSet)
	if err := _Pbtc.contract.UnpackLog(event, "EnforcedOptionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PbtcInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Pbtc contract.
type PbtcInitializedIterator struct {
	Event *PbtcInitialized // Event containing the contract specifics and raw log

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
func (it *PbtcInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PbtcInitialized)
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
		it.Event = new(PbtcInitialized)
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
func (it *PbtcInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PbtcInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PbtcInitialized represents a Initialized event raised by the Pbtc contract.
type PbtcInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Pbtc *PbtcFilterer) FilterInitialized(opts *bind.FilterOpts) (*PbtcInitializedIterator, error) {

	logs, sub, err := _Pbtc.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PbtcInitializedIterator{contract: _Pbtc.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Pbtc *PbtcFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PbtcInitialized) (event.Subscription, error) {

	logs, sub, err := _Pbtc.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PbtcInitialized)
				if err := _Pbtc.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Pbtc *PbtcFilterer) ParseInitialized(log types.Log) (*PbtcInitialized, error) {
	event := new(PbtcInitialized)
	if err := _Pbtc.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PbtcMintConfirmedIterator is returned from FilterMintConfirmed and is used to iterate over the raw logs and unpacked data for MintConfirmed events raised by the Pbtc contract.
type PbtcMintConfirmedIterator struct {
	Event *PbtcMintConfirmed // Event containing the contract specifics and raw log

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
func (it *PbtcMintConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PbtcMintConfirmed)
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
		it.Event = new(PbtcMintConfirmed)
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
func (it *PbtcMintConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PbtcMintConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PbtcMintConfirmed represents a MintConfirmed event raised by the Pbtc contract.
type PbtcMintConfirmed struct {
	Nonce             *big.Int
	Requester         common.Address
	Amount            *big.Int
	BtcDepositAddress string
	BtcTxid           string
	Timestamp         *big.Int
	RequestHash       [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterMintConfirmed is a free log retrieval operation binding the contract event 0x051f4ba27061b0e6dc829669a7baa8bba9cf7f6cd2f95e1f0bdd9c22126d8b21.
//
// Solidity: event MintConfirmed(uint256 indexed nonce, address indexed requester, uint256 amount, string btcDepositAddress, string btcTxid, uint256 timestamp, bytes32 requestHash)
func (_Pbtc *PbtcFilterer) FilterMintConfirmed(opts *bind.FilterOpts, nonce []*big.Int, requester []common.Address) (*PbtcMintConfirmedIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _Pbtc.contract.FilterLogs(opts, "MintConfirmed", nonceRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &PbtcMintConfirmedIterator{contract: _Pbtc.contract, event: "MintConfirmed", logs: logs, sub: sub}, nil
}

// WatchMintConfirmed is a free log subscription operation binding the contract event 0x051f4ba27061b0e6dc829669a7baa8bba9cf7f6cd2f95e1f0bdd9c22126d8b21.
//
// Solidity: event MintConfirmed(uint256 indexed nonce, address indexed requester, uint256 amount, string btcDepositAddress, string btcTxid, uint256 timestamp, bytes32 requestHash)
func (_Pbtc *PbtcFilterer) WatchMintConfirmed(opts *bind.WatchOpts, sink chan<- *PbtcMintConfirmed, nonce []*big.Int, requester []common.Address) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _Pbtc.contract.WatchLogs(opts, "MintConfirmed", nonceRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PbtcMintConfirmed)
				if err := _Pbtc.contract.UnpackLog(event, "MintConfirmed", log); err != nil {
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

// ParseMintConfirmed is a log parse operation binding the contract event 0x051f4ba27061b0e6dc829669a7baa8bba9cf7f6cd2f95e1f0bdd9c22126d8b21.
//
// Solidity: event MintConfirmed(uint256 indexed nonce, address indexed requester, uint256 amount, string btcDepositAddress, string btcTxid, uint256 timestamp, bytes32 requestHash)
func (_Pbtc *PbtcFilterer) ParseMintConfirmed(log types.Log) (*PbtcMintConfirmed, error) {
	event := new(PbtcMintConfirmed)
	if err := _Pbtc.contract.UnpackLog(event, "MintConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PbtcMintRejectedIterator is returned from FilterMintRejected and is used to iterate over the raw logs and unpacked data for MintRejected events raised by the Pbtc contract.
type PbtcMintRejectedIterator struct {
	Event *PbtcMintRejected // Event containing the contract specifics and raw log

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
func (it *PbtcMintRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PbtcMintRejected)
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
		it.Event = new(PbtcMintRejected)
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
func (it *PbtcMintRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PbtcMintRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PbtcMintRejected represents a MintRejected event raised by the Pbtc contract.
type PbtcMintRejected struct {
	Nonce             *big.Int
	Requester         common.Address
	Amount            *big.Int
	BtcDepositAddress string
	BtcTxid           string
	Timestamp         *big.Int
	RequestHash       [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterMintRejected is a free log retrieval operation binding the contract event 0xdadc06f5b98131083e96b856c044184efd23ae2e797a876fd80aa5dae4f72455.
//
// Solidity: event MintRejected(uint256 indexed nonce, address indexed requester, uint256 amount, string btcDepositAddress, string btcTxid, uint256 timestamp, bytes32 requestHash)
func (_Pbtc *PbtcFilterer) FilterMintRejected(opts *bind.FilterOpts, nonce []*big.Int, requester []common.Address) (*PbtcMintRejectedIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _Pbtc.contract.FilterLogs(opts, "MintRejected", nonceRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &PbtcMintRejectedIterator{contract: _Pbtc.contract, event: "MintRejected", logs: logs, sub: sub}, nil
}

// WatchMintRejected is a free log subscription operation binding the contract event 0xdadc06f5b98131083e96b856c044184efd23ae2e797a876fd80aa5dae4f72455.
//
// Solidity: event MintRejected(uint256 indexed nonce, address indexed requester, uint256 amount, string btcDepositAddress, string btcTxid, uint256 timestamp, bytes32 requestHash)
func (_Pbtc *PbtcFilterer) WatchMintRejected(opts *bind.WatchOpts, sink chan<- *PbtcMintRejected, nonce []*big.Int, requester []common.Address) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _Pbtc.contract.WatchLogs(opts, "MintRejected", nonceRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PbtcMintRejected)
				if err := _Pbtc.contract.UnpackLog(event, "MintRejected", log); err != nil {
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

// ParseMintRejected is a log parse operation binding the contract event 0xdadc06f5b98131083e96b856c044184efd23ae2e797a876fd80aa5dae4f72455.
//
// Solidity: event MintRejected(uint256 indexed nonce, address indexed requester, uint256 amount, string btcDepositAddress, string btcTxid, uint256 timestamp, bytes32 requestHash)
func (_Pbtc *PbtcFilterer) ParseMintRejected(log types.Log) (*PbtcMintRejected, error) {
	event := new(PbtcMintRejected)
	if err := _Pbtc.contract.UnpackLog(event, "MintRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PbtcMintRequestAddIterator is returned from FilterMintRequestAdd and is used to iterate over the raw logs and unpacked data for MintRequestAdd events raised by the Pbtc contract.
type PbtcMintRequestAddIterator struct {
	Event *PbtcMintRequestAdd // Event containing the contract specifics and raw log

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
func (it *PbtcMintRequestAddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PbtcMintRequestAdd)
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
		it.Event = new(PbtcMintRequestAdd)
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
func (it *PbtcMintRequestAddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PbtcMintRequestAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PbtcMintRequestAdd represents a MintRequestAdd event raised by the Pbtc contract.
type PbtcMintRequestAdd struct {
	Nonce             *big.Int
	Requester         common.Address
	Amount            *big.Int
	BtcDepositAddress string
	BtcTxid           string
	Timestamp         *big.Int
	RequestHash       [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterMintRequestAdd is a free log retrieval operation binding the contract event 0x09e00024b3e14e42d4e78c05bf370a34c2e4ce4027dad38abafdb1bf49da432f.
//
// Solidity: event MintRequestAdd(uint256 indexed nonce, address indexed requester, uint256 amount, string btcDepositAddress, string btcTxid, uint256 timestamp, bytes32 requestHash)
func (_Pbtc *PbtcFilterer) FilterMintRequestAdd(opts *bind.FilterOpts, nonce []*big.Int, requester []common.Address) (*PbtcMintRequestAddIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _Pbtc.contract.FilterLogs(opts, "MintRequestAdd", nonceRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &PbtcMintRequestAddIterator{contract: _Pbtc.contract, event: "MintRequestAdd", logs: logs, sub: sub}, nil
}

// WatchMintRequestAdd is a free log subscription operation binding the contract event 0x09e00024b3e14e42d4e78c05bf370a34c2e4ce4027dad38abafdb1bf49da432f.
//
// Solidity: event MintRequestAdd(uint256 indexed nonce, address indexed requester, uint256 amount, string btcDepositAddress, string btcTxid, uint256 timestamp, bytes32 requestHash)
func (_Pbtc *PbtcFilterer) WatchMintRequestAdd(opts *bind.WatchOpts, sink chan<- *PbtcMintRequestAdd, nonce []*big.Int, requester []common.Address) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _Pbtc.contract.WatchLogs(opts, "MintRequestAdd", nonceRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PbtcMintRequestAdd)
				if err := _Pbtc.contract.UnpackLog(event, "MintRequestAdd", log); err != nil {
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

// ParseMintRequestAdd is a log parse operation binding the contract event 0x09e00024b3e14e42d4e78c05bf370a34c2e4ce4027dad38abafdb1bf49da432f.
//
// Solidity: event MintRequestAdd(uint256 indexed nonce, address indexed requester, uint256 amount, string btcDepositAddress, string btcTxid, uint256 timestamp, bytes32 requestHash)
func (_Pbtc *PbtcFilterer) ParseMintRequestAdd(log types.Log) (*PbtcMintRequestAdd, error) {
	event := new(PbtcMintRequestAdd)
	if err := _Pbtc.contract.UnpackLog(event, "MintRequestAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PbtcMintRequestCancelIterator is returned from FilterMintRequestCancel and is used to iterate over the raw logs and unpacked data for MintRequestCancel events raised by the Pbtc contract.
type PbtcMintRequestCancelIterator struct {
	Event *PbtcMintRequestCancel // Event containing the contract specifics and raw log

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
func (it *PbtcMintRequestCancelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PbtcMintRequestCancel)
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
		it.Event = new(PbtcMintRequestCancel)
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
func (it *PbtcMintRequestCancelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PbtcMintRequestCancelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PbtcMintRequestCancel represents a MintRequestCancel event raised by the Pbtc contract.
type PbtcMintRequestCancel struct {
	Nonce       *big.Int
	Requester   common.Address
	RequestHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMintRequestCancel is a free log retrieval operation binding the contract event 0xb419f275eebfa354bbab2709955ee0c0e25ca95fae50a8e3672c5e3d9c931f58.
//
// Solidity: event MintRequestCancel(uint256 indexed nonce, address indexed requester, bytes32 requestHash)
func (_Pbtc *PbtcFilterer) FilterMintRequestCancel(opts *bind.FilterOpts, nonce []*big.Int, requester []common.Address) (*PbtcMintRequestCancelIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _Pbtc.contract.FilterLogs(opts, "MintRequestCancel", nonceRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &PbtcMintRequestCancelIterator{contract: _Pbtc.contract, event: "MintRequestCancel", logs: logs, sub: sub}, nil
}

// WatchMintRequestCancel is a free log subscription operation binding the contract event 0xb419f275eebfa354bbab2709955ee0c0e25ca95fae50a8e3672c5e3d9c931f58.
//
// Solidity: event MintRequestCancel(uint256 indexed nonce, address indexed requester, bytes32 requestHash)
func (_Pbtc *PbtcFilterer) WatchMintRequestCancel(opts *bind.WatchOpts, sink chan<- *PbtcMintRequestCancel, nonce []*big.Int, requester []common.Address) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _Pbtc.contract.WatchLogs(opts, "MintRequestCancel", nonceRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PbtcMintRequestCancel)
				if err := _Pbtc.contract.UnpackLog(event, "MintRequestCancel", log); err != nil {
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

// ParseMintRequestCancel is a log parse operation binding the contract event 0xb419f275eebfa354bbab2709955ee0c0e25ca95fae50a8e3672c5e3d9c931f58.
//
// Solidity: event MintRequestCancel(uint256 indexed nonce, address indexed requester, bytes32 requestHash)
func (_Pbtc *PbtcFilterer) ParseMintRequestCancel(log types.Log) (*PbtcMintRequestCancel, error) {
	event := new(PbtcMintRequestCancel)
	if err := _Pbtc.contract.UnpackLog(event, "MintRequestCancel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PbtcMsgInspectorSetIterator is returned from FilterMsgInspectorSet and is used to iterate over the raw logs and unpacked data for MsgInspectorSet events raised by the Pbtc contract.
type PbtcMsgInspectorSetIterator struct {
	Event *PbtcMsgInspectorSet // Event containing the contract specifics and raw log

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
func (it *PbtcMsgInspectorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PbtcMsgInspectorSet)
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
		it.Event = new(PbtcMsgInspectorSet)
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
func (it *PbtcMsgInspectorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PbtcMsgInspectorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PbtcMsgInspectorSet represents a MsgInspectorSet event raised by the Pbtc contract.
type PbtcMsgInspectorSet struct {
	Inspector common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMsgInspectorSet is a free log retrieval operation binding the contract event 0xf0be4f1e87349231d80c36b33f9e8639658eeaf474014dee15a3e6a4d4414197.
//
// Solidity: event MsgInspectorSet(address inspector)
func (_Pbtc *PbtcFilterer) FilterMsgInspectorSet(opts *bind.FilterOpts) (*PbtcMsgInspectorSetIterator, error) {

	logs, sub, err := _Pbtc.contract.FilterLogs(opts, "MsgInspectorSet")
	if err != nil {
		return nil, err
	}
	return &PbtcMsgInspectorSetIterator{contract: _Pbtc.contract, event: "MsgInspectorSet", logs: logs, sub: sub}, nil
}

// WatchMsgInspectorSet is a free log subscription operation binding the contract event 0xf0be4f1e87349231d80c36b33f9e8639658eeaf474014dee15a3e6a4d4414197.
//
// Solidity: event MsgInspectorSet(address inspector)
func (_Pbtc *PbtcFilterer) WatchMsgInspectorSet(opts *bind.WatchOpts, sink chan<- *PbtcMsgInspectorSet) (event.Subscription, error) {

	logs, sub, err := _Pbtc.contract.WatchLogs(opts, "MsgInspectorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PbtcMsgInspectorSet)
				if err := _Pbtc.contract.UnpackLog(event, "MsgInspectorSet", log); err != nil {
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

// ParseMsgInspectorSet is a log parse operation binding the contract event 0xf0be4f1e87349231d80c36b33f9e8639658eeaf474014dee15a3e6a4d4414197.
//
// Solidity: event MsgInspectorSet(address inspector)
func (_Pbtc *PbtcFilterer) ParseMsgInspectorSet(log types.Log) (*PbtcMsgInspectorSet, error) {
	event := new(PbtcMsgInspectorSet)
	if err := _Pbtc.contract.UnpackLog(event, "MsgInspectorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PbtcOFTReceivedIterator is returned from FilterOFTReceived and is used to iterate over the raw logs and unpacked data for OFTReceived events raised by the Pbtc contract.
type PbtcOFTReceivedIterator struct {
	Event *PbtcOFTReceived // Event containing the contract specifics and raw log

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
func (it *PbtcOFTReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PbtcOFTReceived)
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
		it.Event = new(PbtcOFTReceived)
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
func (it *PbtcOFTReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PbtcOFTReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PbtcOFTReceived represents a OFTReceived event raised by the Pbtc contract.
type PbtcOFTReceived struct {
	Guid             [32]byte
	SrcEid           uint32
	ToAddress        common.Address
	AmountReceivedLD *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOFTReceived is a free log retrieval operation binding the contract event 0xefed6d3500546b29533b128a29e3a94d70788727f0507505ac12eaf2e578fd9c.
//
// Solidity: event OFTReceived(bytes32 indexed guid, uint32 srcEid, address indexed toAddress, uint256 amountReceivedLD)
func (_Pbtc *PbtcFilterer) FilterOFTReceived(opts *bind.FilterOpts, guid [][32]byte, toAddress []common.Address) (*PbtcOFTReceivedIterator, error) {

	var guidRule []interface{}
	for _, guidItem := range guid {
		guidRule = append(guidRule, guidItem)
	}

	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _Pbtc.contract.FilterLogs(opts, "OFTReceived", guidRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return &PbtcOFTReceivedIterator{contract: _Pbtc.contract, event: "OFTReceived", logs: logs, sub: sub}, nil
}

// WatchOFTReceived is a free log subscription operation binding the contract event 0xefed6d3500546b29533b128a29e3a94d70788727f0507505ac12eaf2e578fd9c.
//
// Solidity: event OFTReceived(bytes32 indexed guid, uint32 srcEid, address indexed toAddress, uint256 amountReceivedLD)
func (_Pbtc *PbtcFilterer) WatchOFTReceived(opts *bind.WatchOpts, sink chan<- *PbtcOFTReceived, guid [][32]byte, toAddress []common.Address) (event.Subscription, error) {

	var guidRule []interface{}
	for _, guidItem := range guid {
		guidRule = append(guidRule, guidItem)
	}

	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _Pbtc.contract.WatchLogs(opts, "OFTReceived", guidRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PbtcOFTReceived)
				if err := _Pbtc.contract.UnpackLog(event, "OFTReceived", log); err != nil {
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

// ParseOFTReceived is a log parse operation binding the contract event 0xefed6d3500546b29533b128a29e3a94d70788727f0507505ac12eaf2e578fd9c.
//
// Solidity: event OFTReceived(bytes32 indexed guid, uint32 srcEid, address indexed toAddress, uint256 amountReceivedLD)
func (_Pbtc *PbtcFilterer) ParseOFTReceived(log types.Log) (*PbtcOFTReceived, error) {
	event := new(PbtcOFTReceived)
	if err := _Pbtc.contract.UnpackLog(event, "OFTReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PbtcOFTSentIterator is returned from FilterOFTSent and is used to iterate over the raw logs and unpacked data for OFTSent events raised by the Pbtc contract.
type PbtcOFTSentIterator struct {
	Event *PbtcOFTSent // Event containing the contract specifics and raw log

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
func (it *PbtcOFTSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PbtcOFTSent)
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
		it.Event = new(PbtcOFTSent)
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
func (it *PbtcOFTSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PbtcOFTSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PbtcOFTSent represents a OFTSent event raised by the Pbtc contract.
type PbtcOFTSent struct {
	Guid             [32]byte
	DstEid           uint32
	FromAddress      common.Address
	AmountSentLD     *big.Int
	AmountReceivedLD *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOFTSent is a free log retrieval operation binding the contract event 0x85496b760a4b7f8d66384b9df21b381f5d1b1e79f229a47aaf4c232edc2fe59a.
//
// Solidity: event OFTSent(bytes32 indexed guid, uint32 dstEid, address indexed fromAddress, uint256 amountSentLD, uint256 amountReceivedLD)
func (_Pbtc *PbtcFilterer) FilterOFTSent(opts *bind.FilterOpts, guid [][32]byte, fromAddress []common.Address) (*PbtcOFTSentIterator, error) {

	var guidRule []interface{}
	for _, guidItem := range guid {
		guidRule = append(guidRule, guidItem)
	}

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}

	logs, sub, err := _Pbtc.contract.FilterLogs(opts, "OFTSent", guidRule, fromAddressRule)
	if err != nil {
		return nil, err
	}
	return &PbtcOFTSentIterator{contract: _Pbtc.contract, event: "OFTSent", logs: logs, sub: sub}, nil
}

// WatchOFTSent is a free log subscription operation binding the contract event 0x85496b760a4b7f8d66384b9df21b381f5d1b1e79f229a47aaf4c232edc2fe59a.
//
// Solidity: event OFTSent(bytes32 indexed guid, uint32 dstEid, address indexed fromAddress, uint256 amountSentLD, uint256 amountReceivedLD)
func (_Pbtc *PbtcFilterer) WatchOFTSent(opts *bind.WatchOpts, sink chan<- *PbtcOFTSent, guid [][32]byte, fromAddress []common.Address) (event.Subscription, error) {

	var guidRule []interface{}
	for _, guidItem := range guid {
		guidRule = append(guidRule, guidItem)
	}

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}

	logs, sub, err := _Pbtc.contract.WatchLogs(opts, "OFTSent", guidRule, fromAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PbtcOFTSent)
				if err := _Pbtc.contract.UnpackLog(event, "OFTSent", log); err != nil {
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

// ParseOFTSent is a log parse operation binding the contract event 0x85496b760a4b7f8d66384b9df21b381f5d1b1e79f229a47aaf4c232edc2fe59a.
//
// Solidity: event OFTSent(bytes32 indexed guid, uint32 dstEid, address indexed fromAddress, uint256 amountSentLD, uint256 amountReceivedLD)
func (_Pbtc *PbtcFilterer) ParseOFTSent(log types.Log) (*PbtcOFTSent, error) {
	event := new(PbtcOFTSent)
	if err := _Pbtc.contract.UnpackLog(event, "OFTSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PbtcOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the Pbtc contract.
type PbtcOwnershipTransferStartedIterator struct {
	Event *PbtcOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *PbtcOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PbtcOwnershipTransferStarted)
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
		it.Event = new(PbtcOwnershipTransferStarted)
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
func (it *PbtcOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PbtcOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PbtcOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the Pbtc contract.
type PbtcOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Pbtc *PbtcFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PbtcOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pbtc.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PbtcOwnershipTransferStartedIterator{contract: _Pbtc.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Pbtc *PbtcFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *PbtcOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pbtc.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PbtcOwnershipTransferStarted)
				if err := _Pbtc.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_Pbtc *PbtcFilterer) ParseOwnershipTransferStarted(log types.Log) (*PbtcOwnershipTransferStarted, error) {
	event := new(PbtcOwnershipTransferStarted)
	if err := _Pbtc.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PbtcOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Pbtc contract.
type PbtcOwnershipTransferredIterator struct {
	Event *PbtcOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PbtcOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PbtcOwnershipTransferred)
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
		it.Event = new(PbtcOwnershipTransferred)
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
func (it *PbtcOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PbtcOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PbtcOwnershipTransferred represents a OwnershipTransferred event raised by the Pbtc contract.
type PbtcOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pbtc *PbtcFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PbtcOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pbtc.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PbtcOwnershipTransferredIterator{contract: _Pbtc.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pbtc *PbtcFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PbtcOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pbtc.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PbtcOwnershipTransferred)
				if err := _Pbtc.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Pbtc *PbtcFilterer) ParseOwnershipTransferred(log types.Log) (*PbtcOwnershipTransferred, error) {
	event := new(PbtcOwnershipTransferred)
	if err := _Pbtc.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PbtcPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Pbtc contract.
type PbtcPausedIterator struct {
	Event *PbtcPaused // Event containing the contract specifics and raw log

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
func (it *PbtcPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PbtcPaused)
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
		it.Event = new(PbtcPaused)
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
func (it *PbtcPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PbtcPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PbtcPaused represents a Paused event raised by the Pbtc contract.
type PbtcPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pbtc *PbtcFilterer) FilterPaused(opts *bind.FilterOpts) (*PbtcPausedIterator, error) {

	logs, sub, err := _Pbtc.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PbtcPausedIterator{contract: _Pbtc.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pbtc *PbtcFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PbtcPaused) (event.Subscription, error) {

	logs, sub, err := _Pbtc.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PbtcPaused)
				if err := _Pbtc.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Pbtc *PbtcFilterer) ParsePaused(log types.Log) (*PbtcPaused, error) {
	event := new(PbtcPaused)
	if err := _Pbtc.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PbtcPeerSetIterator is returned from FilterPeerSet and is used to iterate over the raw logs and unpacked data for PeerSet events raised by the Pbtc contract.
type PbtcPeerSetIterator struct {
	Event *PbtcPeerSet // Event containing the contract specifics and raw log

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
func (it *PbtcPeerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PbtcPeerSet)
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
		it.Event = new(PbtcPeerSet)
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
func (it *PbtcPeerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PbtcPeerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PbtcPeerSet represents a PeerSet event raised by the Pbtc contract.
type PbtcPeerSet struct {
	Eid  uint32
	Peer [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPeerSet is a free log retrieval operation binding the contract event 0x238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b.
//
// Solidity: event PeerSet(uint32 eid, bytes32 peer)
func (_Pbtc *PbtcFilterer) FilterPeerSet(opts *bind.FilterOpts) (*PbtcPeerSetIterator, error) {

	logs, sub, err := _Pbtc.contract.FilterLogs(opts, "PeerSet")
	if err != nil {
		return nil, err
	}
	return &PbtcPeerSetIterator{contract: _Pbtc.contract, event: "PeerSet", logs: logs, sub: sub}, nil
}

// WatchPeerSet is a free log subscription operation binding the contract event 0x238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b.
//
// Solidity: event PeerSet(uint32 eid, bytes32 peer)
func (_Pbtc *PbtcFilterer) WatchPeerSet(opts *bind.WatchOpts, sink chan<- *PbtcPeerSet) (event.Subscription, error) {

	logs, sub, err := _Pbtc.contract.WatchLogs(opts, "PeerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PbtcPeerSet)
				if err := _Pbtc.contract.UnpackLog(event, "PeerSet", log); err != nil {
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

// ParsePeerSet is a log parse operation binding the contract event 0x238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b.
//
// Solidity: event PeerSet(uint32 eid, bytes32 peer)
func (_Pbtc *PbtcFilterer) ParsePeerSet(log types.Log) (*PbtcPeerSet, error) {
	event := new(PbtcPeerSet)
	if err := _Pbtc.contract.UnpackLog(event, "PeerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PbtcPoRChangedIterator is returned from FilterPoRChanged and is used to iterate over the raw logs and unpacked data for PoRChanged events raised by the Pbtc contract.
type PbtcPoRChangedIterator struct {
	Event *PbtcPoRChanged // Event containing the contract specifics and raw log

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
func (it *PbtcPoRChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PbtcPoRChanged)
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
		it.Event = new(PbtcPoRChanged)
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
func (it *PbtcPoRChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PbtcPoRChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PbtcPoRChanged represents a PoRChanged event raised by the Pbtc contract.
type PbtcPoRChanged struct {
	OldPoR common.Address
	NewPoR common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoRChanged is a free log retrieval operation binding the contract event 0x7ddf314244433e132607dbe5882c97a2b22af5263e52fd435e1d1f01e6c12301.
//
// Solidity: event PoRChanged(address indexed oldPoR, address indexed newPoR)
func (_Pbtc *PbtcFilterer) FilterPoRChanged(opts *bind.FilterOpts, oldPoR []common.Address, newPoR []common.Address) (*PbtcPoRChangedIterator, error) {

	var oldPoRRule []interface{}
	for _, oldPoRItem := range oldPoR {
		oldPoRRule = append(oldPoRRule, oldPoRItem)
	}
	var newPoRRule []interface{}
	for _, newPoRItem := range newPoR {
		newPoRRule = append(newPoRRule, newPoRItem)
	}

	logs, sub, err := _Pbtc.contract.FilterLogs(opts, "PoRChanged", oldPoRRule, newPoRRule)
	if err != nil {
		return nil, err
	}
	return &PbtcPoRChangedIterator{contract: _Pbtc.contract, event: "PoRChanged", logs: logs, sub: sub}, nil
}

// WatchPoRChanged is a free log subscription operation binding the contract event 0x7ddf314244433e132607dbe5882c97a2b22af5263e52fd435e1d1f01e6c12301.
//
// Solidity: event PoRChanged(address indexed oldPoR, address indexed newPoR)
func (_Pbtc *PbtcFilterer) WatchPoRChanged(opts *bind.WatchOpts, sink chan<- *PbtcPoRChanged, oldPoR []common.Address, newPoR []common.Address) (event.Subscription, error) {

	var oldPoRRule []interface{}
	for _, oldPoRItem := range oldPoR {
		oldPoRRule = append(oldPoRRule, oldPoRItem)
	}
	var newPoRRule []interface{}
	for _, newPoRItem := range newPoR {
		newPoRRule = append(newPoRRule, newPoRItem)
	}

	logs, sub, err := _Pbtc.contract.WatchLogs(opts, "PoRChanged", oldPoRRule, newPoRRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PbtcPoRChanged)
				if err := _Pbtc.contract.UnpackLog(event, "PoRChanged", log); err != nil {
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

// ParsePoRChanged is a log parse operation binding the contract event 0x7ddf314244433e132607dbe5882c97a2b22af5263e52fd435e1d1f01e6c12301.
//
// Solidity: event PoRChanged(address indexed oldPoR, address indexed newPoR)
func (_Pbtc *PbtcFilterer) ParsePoRChanged(log types.Log) (*PbtcPoRChanged, error) {
	event := new(PbtcPoRChanged)
	if err := _Pbtc.contract.UnpackLog(event, "PoRChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PbtcPreCrimeSetIterator is returned from FilterPreCrimeSet and is used to iterate over the raw logs and unpacked data for PreCrimeSet events raised by the Pbtc contract.
type PbtcPreCrimeSetIterator struct {
	Event *PbtcPreCrimeSet // Event containing the contract specifics and raw log

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
func (it *PbtcPreCrimeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PbtcPreCrimeSet)
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
		it.Event = new(PbtcPreCrimeSet)
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
func (it *PbtcPreCrimeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PbtcPreCrimeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PbtcPreCrimeSet represents a PreCrimeSet event raised by the Pbtc contract.
type PbtcPreCrimeSet struct {
	PreCrimeAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPreCrimeSet is a free log retrieval operation binding the contract event 0xd48d879cef83a1c0bdda516f27b13ddb1b3f8bbac1c9e1511bb2a659c2427760.
//
// Solidity: event PreCrimeSet(address preCrimeAddress)
func (_Pbtc *PbtcFilterer) FilterPreCrimeSet(opts *bind.FilterOpts) (*PbtcPreCrimeSetIterator, error) {

	logs, sub, err := _Pbtc.contract.FilterLogs(opts, "PreCrimeSet")
	if err != nil {
		return nil, err
	}
	return &PbtcPreCrimeSetIterator{contract: _Pbtc.contract, event: "PreCrimeSet", logs: logs, sub: sub}, nil
}

// WatchPreCrimeSet is a free log subscription operation binding the contract event 0xd48d879cef83a1c0bdda516f27b13ddb1b3f8bbac1c9e1511bb2a659c2427760.
//
// Solidity: event PreCrimeSet(address preCrimeAddress)
func (_Pbtc *PbtcFilterer) WatchPreCrimeSet(opts *bind.WatchOpts, sink chan<- *PbtcPreCrimeSet) (event.Subscription, error) {

	logs, sub, err := _Pbtc.contract.WatchLogs(opts, "PreCrimeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PbtcPreCrimeSet)
				if err := _Pbtc.contract.UnpackLog(event, "PreCrimeSet", log); err != nil {
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

// ParsePreCrimeSet is a log parse operation binding the contract event 0xd48d879cef83a1c0bdda516f27b13ddb1b3f8bbac1c9e1511bb2a659c2427760.
//
// Solidity: event PreCrimeSet(address preCrimeAddress)
func (_Pbtc *PbtcFilterer) ParsePreCrimeSet(log types.Log) (*PbtcPreCrimeSet, error) {
	event := new(PbtcPreCrimeSet)
	if err := _Pbtc.contract.UnpackLog(event, "PreCrimeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PbtcRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Pbtc contract.
type PbtcRoleAdminChangedIterator struct {
	Event *PbtcRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PbtcRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PbtcRoleAdminChanged)
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
		it.Event = new(PbtcRoleAdminChanged)
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
func (it *PbtcRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PbtcRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PbtcRoleAdminChanged represents a RoleAdminChanged event raised by the Pbtc contract.
type PbtcRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Pbtc *PbtcFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PbtcRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Pbtc.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PbtcRoleAdminChangedIterator{contract: _Pbtc.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Pbtc *PbtcFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PbtcRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Pbtc.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PbtcRoleAdminChanged)
				if err := _Pbtc.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Pbtc *PbtcFilterer) ParseRoleAdminChanged(log types.Log) (*PbtcRoleAdminChanged, error) {
	event := new(PbtcRoleAdminChanged)
	if err := _Pbtc.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PbtcRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Pbtc contract.
type PbtcRoleGrantedIterator struct {
	Event *PbtcRoleGranted // Event containing the contract specifics and raw log

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
func (it *PbtcRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PbtcRoleGranted)
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
		it.Event = new(PbtcRoleGranted)
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
func (it *PbtcRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PbtcRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PbtcRoleGranted represents a RoleGranted event raised by the Pbtc contract.
type PbtcRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Pbtc *PbtcFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PbtcRoleGrantedIterator, error) {

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

	logs, sub, err := _Pbtc.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PbtcRoleGrantedIterator{contract: _Pbtc.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Pbtc *PbtcFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PbtcRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Pbtc.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PbtcRoleGranted)
				if err := _Pbtc.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Pbtc *PbtcFilterer) ParseRoleGranted(log types.Log) (*PbtcRoleGranted, error) {
	event := new(PbtcRoleGranted)
	if err := _Pbtc.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PbtcRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Pbtc contract.
type PbtcRoleRevokedIterator struct {
	Event *PbtcRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PbtcRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PbtcRoleRevoked)
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
		it.Event = new(PbtcRoleRevoked)
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
func (it *PbtcRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PbtcRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PbtcRoleRevoked represents a RoleRevoked event raised by the Pbtc contract.
type PbtcRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Pbtc *PbtcFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PbtcRoleRevokedIterator, error) {

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

	logs, sub, err := _Pbtc.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PbtcRoleRevokedIterator{contract: _Pbtc.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Pbtc *PbtcFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PbtcRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Pbtc.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PbtcRoleRevoked)
				if err := _Pbtc.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Pbtc *PbtcFilterer) ParseRoleRevoked(log types.Log) (*PbtcRoleRevoked, error) {
	event := new(PbtcRoleRevoked)
	if err := _Pbtc.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PbtcTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Pbtc contract.
type PbtcTransferIterator struct {
	Event *PbtcTransfer // Event containing the contract specifics and raw log

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
func (it *PbtcTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PbtcTransfer)
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
		it.Event = new(PbtcTransfer)
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
func (it *PbtcTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PbtcTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PbtcTransfer represents a Transfer event raised by the Pbtc contract.
type PbtcTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Pbtc *PbtcFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PbtcTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pbtc.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PbtcTransferIterator{contract: _Pbtc.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Pbtc *PbtcFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PbtcTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pbtc.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PbtcTransfer)
				if err := _Pbtc.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Pbtc *PbtcFilterer) ParseTransfer(log types.Log) (*PbtcTransfer, error) {
	event := new(PbtcTransfer)
	if err := _Pbtc.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PbtcUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Pbtc contract.
type PbtcUnpausedIterator struct {
	Event *PbtcUnpaused // Event containing the contract specifics and raw log

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
func (it *PbtcUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PbtcUnpaused)
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
		it.Event = new(PbtcUnpaused)
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
func (it *PbtcUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PbtcUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PbtcUnpaused represents a Unpaused event raised by the Pbtc contract.
type PbtcUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pbtc *PbtcFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PbtcUnpausedIterator, error) {

	logs, sub, err := _Pbtc.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PbtcUnpausedIterator{contract: _Pbtc.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pbtc *PbtcFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PbtcUnpaused) (event.Subscription, error) {

	logs, sub, err := _Pbtc.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PbtcUnpaused)
				if err := _Pbtc.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Pbtc *PbtcFilterer) ParseUnpaused(log types.Log) (*PbtcUnpaused, error) {
	event := new(PbtcUnpaused)
	if err := _Pbtc.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PbtcUserBtcDepositAddressSetIterator is returned from FilterUserBtcDepositAddressSet and is used to iterate over the raw logs and unpacked data for UserBtcDepositAddressSet events raised by the Pbtc contract.
type PbtcUserBtcDepositAddressSetIterator struct {
	Event *PbtcUserBtcDepositAddressSet // Event containing the contract specifics and raw log

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
func (it *PbtcUserBtcDepositAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PbtcUserBtcDepositAddressSet)
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
		it.Event = new(PbtcUserBtcDepositAddressSet)
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
func (it *PbtcUserBtcDepositAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PbtcUserBtcDepositAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PbtcUserBtcDepositAddressSet represents a UserBtcDepositAddressSet event raised by the Pbtc contract.
type PbtcUserBtcDepositAddressSet struct {
	User              common.Address
	BtcDepositAddress string
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUserBtcDepositAddressSet is a free log retrieval operation binding the contract event 0x6a0f2361ff648a6c43cb6324ae5066ff0479212bc2a5741505bab935953911c2.
//
// Solidity: event UserBtcDepositAddressSet(address indexed user, string btcDepositAddress)
func (_Pbtc *PbtcFilterer) FilterUserBtcDepositAddressSet(opts *bind.FilterOpts, user []common.Address) (*PbtcUserBtcDepositAddressSetIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Pbtc.contract.FilterLogs(opts, "UserBtcDepositAddressSet", userRule)
	if err != nil {
		return nil, err
	}
	return &PbtcUserBtcDepositAddressSetIterator{contract: _Pbtc.contract, event: "UserBtcDepositAddressSet", logs: logs, sub: sub}, nil
}

// WatchUserBtcDepositAddressSet is a free log subscription operation binding the contract event 0x6a0f2361ff648a6c43cb6324ae5066ff0479212bc2a5741505bab935953911c2.
//
// Solidity: event UserBtcDepositAddressSet(address indexed user, string btcDepositAddress)
func (_Pbtc *PbtcFilterer) WatchUserBtcDepositAddressSet(opts *bind.WatchOpts, sink chan<- *PbtcUserBtcDepositAddressSet, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Pbtc.contract.WatchLogs(opts, "UserBtcDepositAddressSet", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PbtcUserBtcDepositAddressSet)
				if err := _Pbtc.contract.UnpackLog(event, "UserBtcDepositAddressSet", log); err != nil {
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

// ParseUserBtcDepositAddressSet is a log parse operation binding the contract event 0x6a0f2361ff648a6c43cb6324ae5066ff0479212bc2a5741505bab935953911c2.
//
// Solidity: event UserBtcDepositAddressSet(address indexed user, string btcDepositAddress)
func (_Pbtc *PbtcFilterer) ParseUserBtcDepositAddressSet(log types.Log) (*PbtcUserBtcDepositAddressSet, error) {
	event := new(PbtcUserBtcDepositAddressSet)
	if err := _Pbtc.contract.UnpackLog(event, "UserBtcDepositAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
