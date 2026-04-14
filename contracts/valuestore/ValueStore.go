// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package valuestore

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

// ValueStoreMetaData contains all meta data concerning the ValueStore contract.
var ValueStoreMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MAX_TIMESTAMP_GAP\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValue\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"fairValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"valueUsd\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"numerator\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denominator\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMultipleValues\",\"inputs\":[{\"name\":\"keys\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"fairValues\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"valueUsds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"numerators\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"denominators\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"timestamps\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setValue\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"fairValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"valueUsd\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"numerator\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denominator\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValueUpdated\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"fairValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"valueUsd\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"numerator\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"denominator\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"DivisionByZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DivisionByZeroInBatch\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidKeyInBatch\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NoDataForKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TimestampNotIncreasing\",\"inputs\":[{\"name\":\"newTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"existingTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"TimestampTooFarInFuture\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"TimestampTooFarInPast\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x60a080604052346100c257306080525f516020610f985f395f51905f525460ff8160401c166100b3576002600160401b03196001600160401b03821601610060575b604051610ed190816100c7823960805181818161084001526108e30152f35b6001600160401b0319166001600160401b039081175f516020610f985f395f51905f525581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80610041565b63f92ee8a960e01b5f5260045ffd5b5f80fdfe60a0806040526004361015610012575f80fd5b5f3560e01c90816301ffc9a714610a5f575080634f1ef2861461089457806352d1902d1461082e578063715018a6146107c757806372f755b2146104725780638da5cb5b1461043e578063960384a0146103aa578063ad3cb1cc1461034c578063be2e017914610330578063c4d66de814610224578063e329c28c146100cf5763f2fde38b146100a0575f80fd5b346100cb5760203660031901126100cb576100c96100bc610ab2565b6100c4610c90565b610c1f565b005b5f80fd5b346100cb5760c03660031901126100cb576004356001600160401b0381116100cb576100ff903690600401610b8d565b60243590604435606435926084359460a4359561011a610c90565b8415610215578515158061020d575b6101fe577f5f2a8c3a90ec95498a7028ec8d4e67159d8aa0e0bd28284ea430ba0f1da6877d946101f99261016789610162368585610b18565b610cc3565b60405161017381610ac8565b85815260048a60208301898152604084018c8152606085019188835260808601938452604051888882376020818a81015f815203019020955186555160018601555160028501555160038401555191015581604051928392833781015f815203902095604051948594859094939260609260808301968352602083015260408201520152565b0390a3005b6323d359a360e01b5f5260045ffd5b508015610129565b630eda9c3d60e31b5f5260045ffd5b346100cb5760203660031901126100cb5761023d610ab2565b5f516020610e7c5f395f51905f525460ff8160401c16801561031c575b61030d5768ffffffffffffffffff191668010000000000000001175f516020610e7c5f395f51905f52556001600160a01b038116156102fe576102a79061029f610d84565b6100c4610d84565b68ff0000000000000000195f516020610e7c5f395f51905f5254165f516020610e7c5f395f51905f52557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a1005b63d92e233d60e01b5f5260045ffd5b63f92ee8a960e01b5f5260045ffd5b5060016001600160401b038216101561025a565b346100cb575f3660031901126100cb576020604051610e108152f35b346100cb575f3660031901126100cb57604080519061036b8183610af7565b600582526020820191640352e302e360dc1b83528151928391602083525180918160208501528484015e5f828201840152601f01601f19168101030190f35b346100cb5760203660031901126100cb576004356001600160401b0381116100cb576103dc6020913690600401610b8d565b919082604051938492833781015f815203019020600481015490811561042f5760a09181549160018101549060036002820154910154916040519485526020850152604084015260608301526080820152f35b6339a07f5560e11b5f5260045ffd5b346100cb575f3660031901126100cb575f516020610e3c5f395f51905f52546040516001600160a01b039091168152602090f35b346100cb5760c03660031901126100cb576004356001600160401b0381116100cb576104a2903690600401610b5d565b906024356001600160401b0381116100cb576104c2903690600401610b5d565b90608052906044356001600160401b0381116100cb576104e6903690600401610b5d565b6064929192356001600160401b0381116100cb57610508903690600401610b5d565b9190956084356001600160401b0381116100cb5761052a903690600401610b5d565b94909560a4356001600160401b0381116100cb5761054c903690600401610b5d565b979092610557610c90565b8985148015906107bd575b80156107b3575b80156107a9575b801561079f575b610790575f5b85811061058657005b610591818784610bba565b90501561077e576105a381898e610c0f565b3515158061076b575b610759578a81808a8f828f8f838f948f8f8f918f908f928b8f898d938d809a97818f9599818d81809d8f8b8f846105e4918194610bba565b926105f091928d610c0f565b359136906105fd92610b18565b9061060791610cc3565b816080519161061592610c0f565b359b61062092610c0f565b359861062b92610c0f565b359461063692610c0f565b359361064192610c0f565b35926040519461065086610ac8565b85526020850190815260408501918252606085019283526080850193845261067987878a610bba565b9081604051928392833781015f8152036020019020945185555160018501555160028401555160038301555190600401556106b392610bba565b96909b81608051916106c492610c0f565b35976106cf92610c0f565b35966106da92610c0f565b35956106e6918c610c0f565b35956106f2918d610c0f565b359581604051928392833781015f81520390209360405193849361072d93859094939260609260808301968352602083015260408201520152565b037f5f2a8c3a90ec95498a7028ec8d4e67159d8aa0e0bd28284ea430ba0f1da6877d91a360010161057d565b6379ae650760e01b5f5260045260245ffd5b50610777818a85610c0f565b35156105ac565b630b414c8760e21b5f5260045260245ffd5b63a9854bc960e01b5f5260045ffd5b5088851415610577565b5087851415610570565b5086851415610569565b5085851415610562565b346100cb575f3660031901126100cb576107df610c90565b5f516020610e3c5f395f51905f5280546001600160a01b031981169091555f906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b346100cb575f3660031901126100cb577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031630036108855760206040515f516020610e5c5f395f51905f528152f35b63703e46dd60e11b5f5260045ffd5b60403660031901126100cb576108a8610ab2565b6024356001600160401b0381116100cb57366023820112156100cb576108d8903690602481600401359101610b18565b906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016308114908115610a3d575b506108855761091b610c90565b6040516352d1902d60e01b81526001600160a01b0382169290602081600481875afa5f9181610a09575b5061095d5783634c9c8ce360e01b5f5260045260245ffd5b805f516020610e5c5f395f51905f528592036109f75750823b156109e5575f516020610e5c5f395f51905f5280546001600160a01b031916821790557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a28051156109cd576100c991610daf565b5050346109d657005b63b398979f60e01b5f5260045ffd5b634c9c8ce360e01b5f5260045260245ffd5b632a87526960e21b5f5260045260245ffd5b9091506020813d602011610a35575b81610a2560209383610af7565b810103126100cb57519085610945565b3d9150610a18565b5f516020610e5c5f395f51905f52546001600160a01b0316141590508361090e565b346100cb5760203660031901126100cb576004359063ffffffff60e01b82168092036100cb576020916301ffc9a760e01b8114908115610aa1575b5015158152f35b6304b01c2560e51b14905083610a9a565b600435906001600160a01b03821682036100cb57565b60a081019081106001600160401b03821117610ae357604052565b634e487b7160e01b5f52604160045260245ffd5b90601f801991011681019081106001600160401b03821117610ae357604052565b9291926001600160401b038211610ae35760405191610b41601f8201601f191660200184610af7565b8294818452818301116100cb578281602093845f960137010152565b9181601f840112156100cb578235916001600160401b0383116100cb576020808501948460051b0101116100cb57565b9181601f840112156100cb578235916001600160401b0383116100cb57602083818601950101116100cb57565b9190811015610bfb5760051b81013590601e19813603018212156100cb5701908135916001600160401b0383116100cb5760200182360381136100cb579190565b634e487b7160e01b5f52603260045260245ffd5b9190811015610bfb5760051b0190565b6001600160a01b03168015610c7d575f516020610e3c5f395f51905f5280546001600160a01b0319811683179091556001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3565b631e4fbdf760e01b5f525f60045260245ffd5b5f516020610e3c5f395f51905f52546001600160a01b03163303610cb057565b63118cdaa760e01b5f523360045260245ffd5b610e104201804211610d59578211610d6d57610e10421180610d44575b610d2d576040518151600492602091839181908401835e81015f81520301902001549081610d0c575050565b81811115610d18575050565b631a6e512d60e31b5f5260045260245260445ffd5b5063272c289560e11b5f526004524260245260445ffd5b50610e0f194201428111610d59578210610ce0565b634e487b7160e01b5f52601160045260245ffd5b50632161dd6360e21b5f526004524260245260445ffd5b60ff5f516020610e7c5f395f51905f525460401c1615610da057565b631afcd79f60e31b5f5260045ffd5b905f8091602081519101845af48080610e28575b15610de35750506040513d81523d5f602083013e60203d82010160405290565b15610e0857639996b31560e01b5f9081526001600160a01b0391909116600452602490fd5b3d15610e19576040513d5f823e3d90fd5b63d6bda27560e01b5f5260045ffd5b503d151580610dc35750813b1515610dc356fe9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbcf0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a26469706673582212206299334f0a82cdac751ef8aa1ea547ebf15ae6b8cb9615bb6fb3d8e01b4e13b464736f6c63430008220033f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
}

// ValueStoreABI is the input ABI used to generate the binding from.
// Deprecated: Use ValueStoreMetaData.ABI instead.
var ValueStoreABI = ValueStoreMetaData.ABI

// ValueStoreBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ValueStoreMetaData.Bin instead.
var ValueStoreBin = ValueStoreMetaData.Bin

// DeployValueStore deploys a new Ethereum contract, binding an instance of ValueStore to it.
func DeployValueStore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ValueStore, error) {
	parsed, err := ValueStoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ValueStoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValueStore{ValueStoreCaller: ValueStoreCaller{contract: contract}, ValueStoreTransactor: ValueStoreTransactor{contract: contract}, ValueStoreFilterer: ValueStoreFilterer{contract: contract}}, nil
}

// ValueStore is an auto generated Go binding around an Ethereum contract.
type ValueStore struct {
	ValueStoreCaller     // Read-only binding to the contract
	ValueStoreTransactor // Write-only binding to the contract
	ValueStoreFilterer   // Log filterer for contract events
}

// ValueStoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValueStoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueStoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValueStoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueStoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValueStoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueStoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValueStoreSession struct {
	Contract     *ValueStore       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValueStoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValueStoreCallerSession struct {
	Contract *ValueStoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ValueStoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValueStoreTransactorSession struct {
	Contract     *ValueStoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ValueStoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValueStoreRaw struct {
	Contract *ValueStore // Generic contract binding to access the raw methods on
}

// ValueStoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValueStoreCallerRaw struct {
	Contract *ValueStoreCaller // Generic read-only contract binding to access the raw methods on
}

// ValueStoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValueStoreTransactorRaw struct {
	Contract *ValueStoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValueStore creates a new instance of ValueStore, bound to a specific deployed contract.
func NewValueStore(address common.Address, backend bind.ContractBackend) (*ValueStore, error) {
	contract, err := bindValueStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValueStore{ValueStoreCaller: ValueStoreCaller{contract: contract}, ValueStoreTransactor: ValueStoreTransactor{contract: contract}, ValueStoreFilterer: ValueStoreFilterer{contract: contract}}, nil
}

// NewValueStoreCaller creates a new read-only instance of ValueStore, bound to a specific deployed contract.
func NewValueStoreCaller(address common.Address, caller bind.ContractCaller) (*ValueStoreCaller, error) {
	contract, err := bindValueStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValueStoreCaller{contract: contract}, nil
}

// NewValueStoreTransactor creates a new write-only instance of ValueStore, bound to a specific deployed contract.
func NewValueStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*ValueStoreTransactor, error) {
	contract, err := bindValueStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValueStoreTransactor{contract: contract}, nil
}

// NewValueStoreFilterer creates a new log filterer instance of ValueStore, bound to a specific deployed contract.
func NewValueStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*ValueStoreFilterer, error) {
	contract, err := bindValueStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValueStoreFilterer{contract: contract}, nil
}

// bindValueStore binds a generic wrapper to an already deployed contract.
func bindValueStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ValueStoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValueStore *ValueStoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValueStore.Contract.ValueStoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValueStore *ValueStoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValueStore.Contract.ValueStoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValueStore *ValueStoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValueStore.Contract.ValueStoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValueStore *ValueStoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValueStore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValueStore *ValueStoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValueStore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValueStore *ValueStoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValueStore.Contract.contract.Transact(opts, method, params...)
}

// MAXTIMESTAMPGAP is a free data retrieval call binding the contract method 0xbe2e0179.
//
// Solidity: function MAX_TIMESTAMP_GAP() view returns(uint256)
func (_ValueStore *ValueStoreCaller) MAXTIMESTAMPGAP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValueStore.contract.Call(opts, &out, "MAX_TIMESTAMP_GAP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTIMESTAMPGAP is a free data retrieval call binding the contract method 0xbe2e0179.
//
// Solidity: function MAX_TIMESTAMP_GAP() view returns(uint256)
func (_ValueStore *ValueStoreSession) MAXTIMESTAMPGAP() (*big.Int, error) {
	return _ValueStore.Contract.MAXTIMESTAMPGAP(&_ValueStore.CallOpts)
}

// MAXTIMESTAMPGAP is a free data retrieval call binding the contract method 0xbe2e0179.
//
// Solidity: function MAX_TIMESTAMP_GAP() view returns(uint256)
func (_ValueStore *ValueStoreCallerSession) MAXTIMESTAMPGAP() (*big.Int, error) {
	return _ValueStore.Contract.MAXTIMESTAMPGAP(&_ValueStore.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ValueStore *ValueStoreCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ValueStore.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ValueStore *ValueStoreSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ValueStore.Contract.UPGRADEINTERFACEVERSION(&_ValueStore.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ValueStore *ValueStoreCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ValueStore.Contract.UPGRADEINTERFACEVERSION(&_ValueStore.CallOpts)
}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint256 fairValue, uint256 valueUsd, uint256 numerator, uint256 denominator, uint256 timestamp)
func (_ValueStore *ValueStoreCaller) GetValue(opts *bind.CallOpts, key string) (struct {
	FairValue   *big.Int
	ValueUsd    *big.Int
	Numerator   *big.Int
	Denominator *big.Int
	Timestamp   *big.Int
}, error) {
	var out []interface{}
	err := _ValueStore.contract.Call(opts, &out, "getValue", key)

	outstruct := new(struct {
		FairValue   *big.Int
		ValueUsd    *big.Int
		Numerator   *big.Int
		Denominator *big.Int
		Timestamp   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FairValue = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ValueUsd = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Numerator = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Denominator = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint256 fairValue, uint256 valueUsd, uint256 numerator, uint256 denominator, uint256 timestamp)
func (_ValueStore *ValueStoreSession) GetValue(key string) (struct {
	FairValue   *big.Int
	ValueUsd    *big.Int
	Numerator   *big.Int
	Denominator *big.Int
	Timestamp   *big.Int
}, error) {
	return _ValueStore.Contract.GetValue(&_ValueStore.CallOpts, key)
}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint256 fairValue, uint256 valueUsd, uint256 numerator, uint256 denominator, uint256 timestamp)
func (_ValueStore *ValueStoreCallerSession) GetValue(key string) (struct {
	FairValue   *big.Int
	ValueUsd    *big.Int
	Numerator   *big.Int
	Denominator *big.Int
	Timestamp   *big.Int
}, error) {
	return _ValueStore.Contract.GetValue(&_ValueStore.CallOpts, key)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ValueStore *ValueStoreCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValueStore.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ValueStore *ValueStoreSession) Owner() (common.Address, error) {
	return _ValueStore.Contract.Owner(&_ValueStore.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ValueStore *ValueStoreCallerSession) Owner() (common.Address, error) {
	return _ValueStore.Contract.Owner(&_ValueStore.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ValueStore *ValueStoreCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ValueStore.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ValueStore *ValueStoreSession) ProxiableUUID() ([32]byte, error) {
	return _ValueStore.Contract.ProxiableUUID(&_ValueStore.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ValueStore *ValueStoreCallerSession) ProxiableUUID() ([32]byte, error) {
	return _ValueStore.Contract.ProxiableUUID(&_ValueStore.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ValueStore *ValueStoreCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ValueStore.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ValueStore *ValueStoreSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ValueStore.Contract.SupportsInterface(&_ValueStore.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ValueStore *ValueStoreCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ValueStore.Contract.SupportsInterface(&_ValueStore.CallOpts, interfaceId)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_ValueStore *ValueStoreTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address) (*types.Transaction, error) {
	return _ValueStore.contract.Transact(opts, "initialize", initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_ValueStore *ValueStoreSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _ValueStore.Contract.Initialize(&_ValueStore.TransactOpts, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_ValueStore *ValueStoreTransactorSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _ValueStore.Contract.Initialize(&_ValueStore.TransactOpts, initialOwner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ValueStore *ValueStoreTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValueStore.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ValueStore *ValueStoreSession) RenounceOwnership() (*types.Transaction, error) {
	return _ValueStore.Contract.RenounceOwnership(&_ValueStore.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ValueStore *ValueStoreTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ValueStore.Contract.RenounceOwnership(&_ValueStore.TransactOpts)
}

// SetMultipleValues is a paid mutator transaction binding the contract method 0x72f755b2.
//
// Solidity: function setMultipleValues(string[] keys, uint256[] fairValues, uint256[] valueUsds, uint256[] numerators, uint256[] denominators, uint256[] timestamps) returns()
func (_ValueStore *ValueStoreTransactor) SetMultipleValues(opts *bind.TransactOpts, keys []string, fairValues []*big.Int, valueUsds []*big.Int, numerators []*big.Int, denominators []*big.Int, timestamps []*big.Int) (*types.Transaction, error) {
	return _ValueStore.contract.Transact(opts, "setMultipleValues", keys, fairValues, valueUsds, numerators, denominators, timestamps)
}

// SetMultipleValues is a paid mutator transaction binding the contract method 0x72f755b2.
//
// Solidity: function setMultipleValues(string[] keys, uint256[] fairValues, uint256[] valueUsds, uint256[] numerators, uint256[] denominators, uint256[] timestamps) returns()
func (_ValueStore *ValueStoreSession) SetMultipleValues(keys []string, fairValues []*big.Int, valueUsds []*big.Int, numerators []*big.Int, denominators []*big.Int, timestamps []*big.Int) (*types.Transaction, error) {
	return _ValueStore.Contract.SetMultipleValues(&_ValueStore.TransactOpts, keys, fairValues, valueUsds, numerators, denominators, timestamps)
}

// SetMultipleValues is a paid mutator transaction binding the contract method 0x72f755b2.
//
// Solidity: function setMultipleValues(string[] keys, uint256[] fairValues, uint256[] valueUsds, uint256[] numerators, uint256[] denominators, uint256[] timestamps) returns()
func (_ValueStore *ValueStoreTransactorSession) SetMultipleValues(keys []string, fairValues []*big.Int, valueUsds []*big.Int, numerators []*big.Int, denominators []*big.Int, timestamps []*big.Int) (*types.Transaction, error) {
	return _ValueStore.Contract.SetMultipleValues(&_ValueStore.TransactOpts, keys, fairValues, valueUsds, numerators, denominators, timestamps)
}

// SetValue is a paid mutator transaction binding the contract method 0xe329c28c.
//
// Solidity: function setValue(string key, uint256 fairValue, uint256 valueUsd, uint256 numerator, uint256 denominator, uint256 timestamp) returns()
func (_ValueStore *ValueStoreTransactor) SetValue(opts *bind.TransactOpts, key string, fairValue *big.Int, valueUsd *big.Int, numerator *big.Int, denominator *big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _ValueStore.contract.Transact(opts, "setValue", key, fairValue, valueUsd, numerator, denominator, timestamp)
}

// SetValue is a paid mutator transaction binding the contract method 0xe329c28c.
//
// Solidity: function setValue(string key, uint256 fairValue, uint256 valueUsd, uint256 numerator, uint256 denominator, uint256 timestamp) returns()
func (_ValueStore *ValueStoreSession) SetValue(key string, fairValue *big.Int, valueUsd *big.Int, numerator *big.Int, denominator *big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _ValueStore.Contract.SetValue(&_ValueStore.TransactOpts, key, fairValue, valueUsd, numerator, denominator, timestamp)
}

// SetValue is a paid mutator transaction binding the contract method 0xe329c28c.
//
// Solidity: function setValue(string key, uint256 fairValue, uint256 valueUsd, uint256 numerator, uint256 denominator, uint256 timestamp) returns()
func (_ValueStore *ValueStoreTransactorSession) SetValue(key string, fairValue *big.Int, valueUsd *big.Int, numerator *big.Int, denominator *big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _ValueStore.Contract.SetValue(&_ValueStore.TransactOpts, key, fairValue, valueUsd, numerator, denominator, timestamp)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ValueStore *ValueStoreTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ValueStore.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ValueStore *ValueStoreSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ValueStore.Contract.TransferOwnership(&_ValueStore.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ValueStore *ValueStoreTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ValueStore.Contract.TransferOwnership(&_ValueStore.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ValueStore *ValueStoreTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ValueStore.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ValueStore *ValueStoreSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ValueStore.Contract.UpgradeToAndCall(&_ValueStore.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ValueStore *ValueStoreTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ValueStore.Contract.UpgradeToAndCall(&_ValueStore.TransactOpts, newImplementation, data)
}

// ValueStoreInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ValueStore contract.
type ValueStoreInitializedIterator struct {
	Event *ValueStoreInitialized // Event containing the contract specifics and raw log

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
func (it *ValueStoreInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValueStoreInitialized)
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
		it.Event = new(ValueStoreInitialized)
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
func (it *ValueStoreInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValueStoreInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValueStoreInitialized represents a Initialized event raised by the ValueStore contract.
type ValueStoreInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ValueStore *ValueStoreFilterer) FilterInitialized(opts *bind.FilterOpts) (*ValueStoreInitializedIterator, error) {

	logs, sub, err := _ValueStore.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ValueStoreInitializedIterator{contract: _ValueStore.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ValueStore *ValueStoreFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ValueStoreInitialized) (event.Subscription, error) {

	logs, sub, err := _ValueStore.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValueStoreInitialized)
				if err := _ValueStore.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ValueStore *ValueStoreFilterer) ParseInitialized(log types.Log) (*ValueStoreInitialized, error) {
	event := new(ValueStoreInitialized)
	if err := _ValueStore.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValueStoreOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ValueStore contract.
type ValueStoreOwnershipTransferredIterator struct {
	Event *ValueStoreOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ValueStoreOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValueStoreOwnershipTransferred)
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
		it.Event = new(ValueStoreOwnershipTransferred)
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
func (it *ValueStoreOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValueStoreOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValueStoreOwnershipTransferred represents a OwnershipTransferred event raised by the ValueStore contract.
type ValueStoreOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ValueStore *ValueStoreFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ValueStoreOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ValueStore.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ValueStoreOwnershipTransferredIterator{contract: _ValueStore.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ValueStore *ValueStoreFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ValueStoreOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ValueStore.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValueStoreOwnershipTransferred)
				if err := _ValueStore.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ValueStore *ValueStoreFilterer) ParseOwnershipTransferred(log types.Log) (*ValueStoreOwnershipTransferred, error) {
	event := new(ValueStoreOwnershipTransferred)
	if err := _ValueStore.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValueStoreUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ValueStore contract.
type ValueStoreUpgradedIterator struct {
	Event *ValueStoreUpgraded // Event containing the contract specifics and raw log

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
func (it *ValueStoreUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValueStoreUpgraded)
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
		it.Event = new(ValueStoreUpgraded)
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
func (it *ValueStoreUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValueStoreUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValueStoreUpgraded represents a Upgraded event raised by the ValueStore contract.
type ValueStoreUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ValueStore *ValueStoreFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ValueStoreUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ValueStore.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ValueStoreUpgradedIterator{contract: _ValueStore.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ValueStore *ValueStoreFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ValueStoreUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ValueStore.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValueStoreUpgraded)
				if err := _ValueStore.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_ValueStore *ValueStoreFilterer) ParseUpgraded(log types.Log) (*ValueStoreUpgraded, error) {
	event := new(ValueStoreUpgraded)
	if err := _ValueStore.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValueStoreValueUpdatedIterator is returned from FilterValueUpdated and is used to iterate over the raw logs and unpacked data for ValueUpdated events raised by the ValueStore contract.
type ValueStoreValueUpdatedIterator struct {
	Event *ValueStoreValueUpdated // Event containing the contract specifics and raw log

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
func (it *ValueStoreValueUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValueStoreValueUpdated)
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
		it.Event = new(ValueStoreValueUpdated)
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
func (it *ValueStoreValueUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValueStoreValueUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValueStoreValueUpdated represents a ValueUpdated event raised by the ValueStore contract.
type ValueStoreValueUpdated struct {
	Key         common.Hash
	FairValue   *big.Int
	ValueUsd    *big.Int
	Numerator   *big.Int
	Denominator *big.Int
	Timestamp   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterValueUpdated is a free log retrieval operation binding the contract event 0x5f2a8c3a90ec95498a7028ec8d4e67159d8aa0e0bd28284ea430ba0f1da6877d.
//
// Solidity: event ValueUpdated(string indexed key, uint256 fairValue, uint256 valueUsd, uint256 numerator, uint256 denominator, uint256 indexed timestamp)
func (_ValueStore *ValueStoreFilterer) FilterValueUpdated(opts *bind.FilterOpts, key []string, timestamp []*big.Int) (*ValueStoreValueUpdatedIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}

	logs, sub, err := _ValueStore.contract.FilterLogs(opts, "ValueUpdated", keyRule, timestampRule)
	if err != nil {
		return nil, err
	}
	return &ValueStoreValueUpdatedIterator{contract: _ValueStore.contract, event: "ValueUpdated", logs: logs, sub: sub}, nil
}

// WatchValueUpdated is a free log subscription operation binding the contract event 0x5f2a8c3a90ec95498a7028ec8d4e67159d8aa0e0bd28284ea430ba0f1da6877d.
//
// Solidity: event ValueUpdated(string indexed key, uint256 fairValue, uint256 valueUsd, uint256 numerator, uint256 denominator, uint256 indexed timestamp)
func (_ValueStore *ValueStoreFilterer) WatchValueUpdated(opts *bind.WatchOpts, sink chan<- *ValueStoreValueUpdated, key []string, timestamp []*big.Int) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}

	logs, sub, err := _ValueStore.contract.WatchLogs(opts, "ValueUpdated", keyRule, timestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValueStoreValueUpdated)
				if err := _ValueStore.contract.UnpackLog(event, "ValueUpdated", log); err != nil {
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

// ParseValueUpdated is a log parse operation binding the contract event 0x5f2a8c3a90ec95498a7028ec8d4e67159d8aa0e0bd28284ea430ba0f1da6877d.
//
// Solidity: event ValueUpdated(string indexed key, uint256 fairValue, uint256 valueUsd, uint256 numerator, uint256 denominator, uint256 indexed timestamp)
func (_ValueStore *ValueStoreFilterer) ParseValueUpdated(log types.Log) (*ValueStoreValueUpdated, error) {
	event := new(ValueStoreValueUpdated)
	if err := _ValueStore.contract.UnpackLog(event, "ValueUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
