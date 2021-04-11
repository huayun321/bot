// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package router

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// RouterABI is the input ABI used to generate the binding from.
const RouterABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_toAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"hua\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RouterBin is the compiled bytecode used for deploying new contracts.
var RouterBin = "0x60c060405234801561001057600080fd5b50604051610e8e380380610e8e8339818101604052604081101561003357600080fd5b5080516020909101516001600160601b0319606092831b8116608052911b1660a05260805160601c60a05160601c610df76100976000398061014052806103445250806101bf5280610296528061037952806107f9528061083c5250610df76000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806319c87f1f1461004657806386917a251461006a578063c45a015514610136575b600080fd5b61004e61013e565b604080516001600160a01b039092168252519081900360200190f35b6100e66004803603608081101561008057600080fd5b8135916020810135918101906060810160408201356401000000008111156100a757600080fd5b8201836020820111156100b957600080fd5b803590602001918460208302840111640100000000831117156100db57600080fd5b919350915035610162565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561012257818101518382015260200161010a565b505050509050019250505060405180910390f35b61004e610377565b7f000000000000000000000000000000000000000000000000000000000000000081565b606081428110156101ba576040805162461bcd60e51b815260206004820152601660248201527f50616e63616b65526f757465723a204558504952454400000000000000000000604482015290519081900360640190fd5b6102187f00000000000000000000000000000000000000000000000000000000000000008887878080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525061039b92505050565b9150858260018451038151811061022b57fe5b602002602001015110156102705760405162461bcd60e51b8152600401808060200182810382526029815260200180610d2c6029913960400191505060405180910390fd5b61030e8585600081811061028057fe5b905060200201356001600160a01b0316336102f47f0000000000000000000000000000000000000000000000000000000000000000898960008181106102c257fe5b905060200201356001600160a01b03168a8a60018181106102df57fe5b905060200201356001600160a01b03166104e7565b8560008151811061030157fe5b60200260200101516105bf565b61036d828686808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152507f0000000000000000000000000000000000000000000000000000000000000000925061074a915050565b5095945050505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60606002825110156103f4576040805162461bcd60e51b815260206004820152601c60248201527f50616e63616b654c6962726172793a20494e56414c49445f5041544800000000604482015290519081900360640190fd5b815167ffffffffffffffff8111801561040c57600080fd5b50604051908082528060200260200182016040528015610436578160200160208202803683370190505b509050828160008151811061044757fe5b60200260200101818152505060005b60018351038110156104df576000806104998786858151811061047557fe5b602002602001015187866001018151811061048c57fe5b6020026020010151610990565b915091506104bb8484815181106104ac57fe5b60200260200101518383610a6a565b8484600101815181106104ca57fe5b60209081029190910101525050600101610456565b509392505050565b60008060006104f68585610b5a565b604080516bffffffffffffffffffffffff19606094851b811660208084019190915293851b81166034830152825160288184030181526048830184528051908501207fff0000000000000000000000000000000000000000000000000000000000000060688401529a90941b9093166069840152607d8301989098527fd0d4c4cd0848c93cb4fd1f498d7013ee6bfb25783ea21593d5834f5d250ece66609d808401919091528851808403909101815260bd909201909752805196019590952095945050505050565b604080516001600160a01b0385811660248301528481166044830152606480830185905283518084039091018152608490920183526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd0000000000000000000000000000000000000000000000000000000017815292518251600094606094938a169392918291908083835b602083106106725780518252601f199092019160209182019101610653565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d80600081146106d4576040519150601f19603f3d011682016040523d82523d6000602084013e6106d9565b606091505b5091509150818015610707575080511580610707575080806020019051602081101561070457600080fd5b50515b6107425760405162461bcd60e51b8152600401808060200182810382526024815260200180610d9e6024913960400191505060405180910390fd5b505050505050565b60005b600183510381101561098a5760008084838151811061076857fe5b602002602001015185846001018151811061077f57fe5b60200260200101519150915060006107978383610b5a565b50905060008785600101815181106107ab57fe5b60200260200101519050600080836001600160a01b0316866001600160a01b0316146107d9578260006107dd565b6000835b91509150600060028a510388106107f45788610835565b6108357f0000000000000000000000000000000000000000000000000000000000000000878c8b6002018151811061082857fe5b60200260200101516104e7565b90506108627f000000000000000000000000000000000000000000000000000000000000000088886104e7565b6001600160a01b031663022c0d9f84848460006040519080825280601f01601f19166020018201604052801561089f576020820181803683370190505b506040518563ffffffff1660e01b815260040180858152602001848152602001836001600160a01b03166001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b838110156109105781810151838201526020016108f8565b50505050905090810190601f16801561093d5780820380516001836020036101000a031916815260200191505b5095505050505050600060405180830381600087803b15801561095f57600080fd5b505af1158015610973573d6000803e3d6000fd5b50506001909901985061074d975050505050505050565b50505050565b600080600061099f8585610b5a565b5090506109ad8686866104e7565b506000806109bc8888886104e7565b6001600160a01b0316630902f1ac6040518163ffffffff1660e01b815260040160606040518083038186803b1580156109f457600080fd5b505afa158015610a08573d6000803e3d6000fd5b505050506040513d6060811015610a1e57600080fd5b5080516020909101516dffffffffffffffffffffffffffff91821693501690506001600160a01b0387811690841614610a58578082610a5b565b81815b90999098509650505050505050565b6000808411610aaa5760405162461bcd60e51b8152600401808060200182810382526029815260200180610d036029913960400191505060405180910390fd5b600083118015610aba5750600082115b610af55760405162461bcd60e51b8152600401808060200182810382526026815260200180610d786026913960400191505060405180910390fd5b6000610b09856103e663ffffffff610c3816565b90506000610b1d828563ffffffff610c3816565b90506000610b4383610b37886103e863ffffffff610c3816565b9063ffffffff610caa16565b9050808281610b4e57fe5b04979650505050505050565b600080826001600160a01b0316846001600160a01b03161415610bae5760405162461bcd60e51b8152600401808060200182810382526023815260200180610d556023913960400191505060405180910390fd5b826001600160a01b0316846001600160a01b031610610bce578284610bd1565b83835b90925090506001600160a01b038216610c31576040805162461bcd60e51b815260206004820152601c60248201527f50616e63616b654c6962726172793a205a45524f5f4144445245535300000000604482015290519081900360640190fd5b9250929050565b6000811580610c5357505080820282828281610c5057fe5b04145b610ca4576040805162461bcd60e51b815260206004820152601460248201527f64732d6d6174682d6d756c2d6f766572666c6f77000000000000000000000000604482015290519081900360640190fd5b92915050565b80820182811015610ca4576040805162461bcd60e51b815260206004820152601460248201527f64732d6d6174682d6164642d6f766572666c6f77000000000000000000000000604482015290519081900360640190fdfe50616e63616b654c6962726172793a20494e53554646494349454e545f494e5055545f414d4f554e5450616e63616b65526f757465723a20494e53554646494349454e545f4f55545055545f414d4f554e5450616e63616b654c6962726172793a204944454e544943414c5f41444452455353455350616e63616b654c6962726172793a20494e53554646494349454e545f4c49515549444954595472616e7366657248656c7065723a205452414e534645525f46524f4d5f4641494c4544a264697066735822122064a4e6eaf46dff0487395dfc6c6df3c53aaadf07c3b35ff1c8766a5c86d30a3864736f6c63430006060033"

// DeployRouter deploys a new Ethereum contract, binding an instance of Router to it.
func DeployRouter(auth *bind.TransactOpts, backend bind.ContractBackend, _factory common.Address, _toAddress common.Address) (common.Address, *types.Transaction, *Router, error) {
	parsed, err := abi.JSON(strings.NewReader(RouterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RouterBin), backend, _factory, _toAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Router{RouterCaller: RouterCaller{contract: contract}, RouterTransactor: RouterTransactor{contract: contract}, RouterFilterer: RouterFilterer{contract: contract}}, nil
}

// Router is an auto generated Go binding around an Ethereum contract.
type Router struct {
	RouterCaller     // Read-only binding to the contract
	RouterTransactor // Write-only binding to the contract
	RouterFilterer   // Log filterer for contract events
}

// RouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type RouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RouterSession struct {
	Contract     *Router           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RouterCallerSession struct {
	Contract *RouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RouterTransactorSession struct {
	Contract     *RouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type RouterRaw struct {
	Contract *Router // Generic contract binding to access the raw methods on
}

// RouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RouterCallerRaw struct {
	Contract *RouterCaller // Generic read-only contract binding to access the raw methods on
}

// RouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RouterTransactorRaw struct {
	Contract *RouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRouter creates a new instance of Router, bound to a specific deployed contract.
func NewRouter(address common.Address, backend bind.ContractBackend) (*Router, error) {
	contract, err := bindRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Router{RouterCaller: RouterCaller{contract: contract}, RouterTransactor: RouterTransactor{contract: contract}, RouterFilterer: RouterFilterer{contract: contract}}, nil
}

// NewRouterCaller creates a new read-only instance of Router, bound to a specific deployed contract.
func NewRouterCaller(address common.Address, caller bind.ContractCaller) (*RouterCaller, error) {
	contract, err := bindRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RouterCaller{contract: contract}, nil
}

// NewRouterTransactor creates a new write-only instance of Router, bound to a specific deployed contract.
func NewRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*RouterTransactor, error) {
	contract, err := bindRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RouterTransactor{contract: contract}, nil
}

// NewRouterFilterer creates a new log filterer instance of Router, bound to a specific deployed contract.
func NewRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*RouterFilterer, error) {
	contract, err := bindRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RouterFilterer{contract: contract}, nil
}

// bindRouter binds a generic wrapper to an already deployed contract.
func bindRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.RouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.contract.Transact(opts, method, params...)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Router *RouterCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Router *RouterSession) Factory() (common.Address, error) {
	return _Router.Contract.Factory(&_Router.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Router *RouterCallerSession) Factory() (common.Address, error) {
	return _Router.Contract.Factory(&_Router.CallOpts)
}

// ToAddress is a free data retrieval call binding the contract method 0x19c87f1f.
//
// Solidity: function toAddress() view returns(address)
func (_Router *RouterCaller) ToAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "toAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ToAddress is a free data retrieval call binding the contract method 0x19c87f1f.
//
// Solidity: function toAddress() view returns(address)
func (_Router *RouterSession) ToAddress() (common.Address, error) {
	return _Router.Contract.ToAddress(&_Router.CallOpts)
}

// ToAddress is a free data retrieval call binding the contract method 0x19c87f1f.
//
// Solidity: function toAddress() view returns(address)
func (_Router *RouterCallerSession) ToAddress() (common.Address, error) {
	return _Router.Contract.ToAddress(&_Router.CallOpts)
}

// Hua is a paid mutator transaction binding the contract method 0x86917a25.
//
// Solidity: function hua(uint256 amountIn, uint256 amountOutMin, address[] path, uint256 deadline) returns(uint256[] amounts)
func (_Router *RouterTransactor) Hua(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "hua", amountIn, amountOutMin, path, deadline)
}

// Hua is a paid mutator transaction binding the contract method 0x86917a25.
//
// Solidity: function hua(uint256 amountIn, uint256 amountOutMin, address[] path, uint256 deadline) returns(uint256[] amounts)
func (_Router *RouterSession) Hua(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.Hua(&_Router.TransactOpts, amountIn, amountOutMin, path, deadline)
}

// Hua is a paid mutator transaction binding the contract method 0x86917a25.
//
// Solidity: function hua(uint256 amountIn, uint256 amountOutMin, address[] path, uint256 deadline) returns(uint256[] amounts)
func (_Router *RouterTransactorSession) Hua(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.Hua(&_Router.TransactOpts, amountIn, amountOutMin, path, deadline)
}
