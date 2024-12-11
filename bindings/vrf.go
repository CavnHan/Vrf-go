// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// VRFMetaData contains all meta data concerning the VRF contract.
var VRFMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"VrfManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fulfillRandomWords\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"randomWords\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRequestStatus\",\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"fulfilled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"randomWords\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"Manager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lastRequestId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestIds\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requestMappings\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"fulfilled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requestRandWords\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"numwords\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setManager\",\"inputs\":[{\"name\":\"newManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"FillRandomWords\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"randomWords\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RequestSend\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"numwords\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"current\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x6080604052348015600f57600080fd5b506016601a565b60ca565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161560695760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b039081161460c75780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6109db806100d96000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c8063d0ebdbe711610071578063d0ebdbe714610186578063d8a4676f14610199578063e5e1b972146101ba578063f2fde38b146101cd578063f68b2da8146101e0578063fc2a88c3146101f357600080fd5b806338ba4614146100b9578063485cc955146100ce578063715018a6146100e1578063873767b4146100e95780638796ba8c146101215780638da5cb5b14610142575b600080fd5b6100cc6100c73660046107a8565b6101fc565b005b6100cc6100dc36600461089d565b6102f0565b6100cc61041c565b61010c6100f73660046108d0565b60016020526000908152604090205460ff1681565b60405190151581526020015b60405180910390f35b61013461012f3660046108d0565b610430565b604051908152602001610118565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b03165b6040516001600160a01b039091168152602001610118565b6100cc6101943660046108e9565b610451565b6101ac6101a73660046108d0565b61047b565b604051610118929190610947565b60005461016e906001600160a01b031681565b6100cc6101db3660046108e9565b6104ed565b6100cc6101ee36600461096a565b61052b565b61013460035481565b6000546001600160a01b031633146102675760405162461bcd60e51b815260206004820152602360248201527f4f6e6c79206d616e616765722063616e2063616c6c20746869732066756e637460448201526234b7b760e91b60648201526084015b60405180910390fd5b604080518082018252600180825260208083018581526000878152838352949094208351815460ff1916901515178155935180519394936102af938501929190910190610732565b509050507ff3cb4deb0441dd096356debf166f879d78cadc19e4b94053c8bea6d3940de93a82826040516102e492919061098c565b60405180910390a15050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff166000811580156103365750825b905060008267ffffffffffffffff1660011480156103535750303b155b905081158015610361575080155b1561037f5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156103a957845460ff60401b1916600160401b1785555b6103b287610604565b600080546001600160a01b0319166001600160a01b038816179055831561041357845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050565b610424610615565b61042e6000610670565b565b6002818154811061044057600080fd5b600091825260209091200154905081565b610459610615565b600080546001600160a01b0319166001600160a01b0392909216919091179055565b600081815260016020818152604080842080549301805482518185028101850190935280835260609460ff16939192918391908301828280156104dd57602002820191906000526020600020905b8154815260200190600101908083116104c9575b5050505050905091509150915091565b6104f5610615565b6001600160a01b03811661051f57604051631e4fbdf760e01b81526000600482015260240161025e565b61052881610670565b50565b610533610615565b60408051808201825260008082528251818152602080820185528084019182528683526001808252949092208351815460ff1916901515178155905180519394919361058793928501929190910190610732565b5050600280546001810182556000919091527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace01839055506003829055604080518381526020810183905230918101919091527fb0089ea4415f0160195bcb1ffe462e932fb0012a47725589204b54981a6306f7906060016102e4565b61060c6106e1565b6105288161072a565b336106477f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b03161461042e5760405163118cdaa760e01b815233600482015260240161025e565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661042e57604051631afcd79f60e31b815260040160405180910390fd5b6104f56106e1565b82805482825590600052602060002090810192821561076d579160200282015b8281111561076d578251825591602001919060010190610752565b5061077992915061077d565b5090565b5b80821115610779576000815560010161077e565b634e487b7160e01b600052604160045260246000fd5b600080604083850312156107bb57600080fd5b82359150602083013567ffffffffffffffff8111156107d957600080fd5b8301601f810185136107ea57600080fd5b803567ffffffffffffffff81111561080457610804610792565b8060051b604051601f19603f830116810181811067ffffffffffffffff8211171561083157610831610792565b60405291825260208184018101929081018884111561084f57600080fd5b6020850194505b8385101561087257843580825260209586019590935001610856565b50809450505050509250929050565b80356001600160a01b038116811461089857600080fd5b919050565b600080604083850312156108b057600080fd5b6108b983610881565b91506108c760208401610881565b90509250929050565b6000602082840312156108e257600080fd5b5035919050565b6000602082840312156108fb57600080fd5b61090482610881565b9392505050565b600081518084526020840193506020830160005b8281101561093d57815186526020958601959091019060010161091f565b5093949350505050565b8215158152604060208201526000610962604083018461090b565b949350505050565b6000806040838503121561097d57600080fd5b50508035926020909101359150565b828152604060208201526000610962604083018461090b56fea26469706673582212208b9e6f937e4a5f213ca9525a86493666ec17a72c49703a4f08e733debae59e2664736f6c634300081b0033",
}

// VRFABI is the input ABI used to generate the binding from.
// Deprecated: Use VRFMetaData.ABI instead.
var VRFABI = VRFMetaData.ABI

// VRFBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VRFMetaData.Bin instead.
var VRFBin = VRFMetaData.Bin

// DeployVRF deploys a new Ethereum contract, binding an instance of VRF to it.
func DeployVRF(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VRF, error) {
	parsed, err := VRFMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRF{VRFCaller: VRFCaller{contract: contract}, VRFTransactor: VRFTransactor{contract: contract}, VRFFilterer: VRFFilterer{contract: contract}}, nil
}

// VRF is an auto generated Go binding around an Ethereum contract.
type VRF struct {
	VRFCaller     // Read-only binding to the contract
	VRFTransactor // Write-only binding to the contract
	VRFFilterer   // Log filterer for contract events
}

// VRFCaller is an auto generated read-only Go binding around an Ethereum contract.
type VRFCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VRFTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VRFFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VRFSession struct {
	Contract     *VRF              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VRFCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VRFCallerSession struct {
	Contract *VRFCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VRFTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VRFTransactorSession struct {
	Contract     *VRFTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VRFRaw is an auto generated low-level Go binding around an Ethereum contract.
type VRFRaw struct {
	Contract *VRF // Generic contract binding to access the raw methods on
}

// VRFCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VRFCallerRaw struct {
	Contract *VRFCaller // Generic read-only contract binding to access the raw methods on
}

// VRFTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VRFTransactorRaw struct {
	Contract *VRFTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVRF creates a new instance of VRF, bound to a specific deployed contract.
func NewVRF(address common.Address, backend bind.ContractBackend) (*VRF, error) {
	contract, err := bindVRF(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRF{VRFCaller: VRFCaller{contract: contract}, VRFTransactor: VRFTransactor{contract: contract}, VRFFilterer: VRFFilterer{contract: contract}}, nil
}

// NewVRFCaller creates a new read-only instance of VRF, bound to a specific deployed contract.
func NewVRFCaller(address common.Address, caller bind.ContractCaller) (*VRFCaller, error) {
	contract, err := bindVRF(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFCaller{contract: contract}, nil
}

// NewVRFTransactor creates a new write-only instance of VRF, bound to a specific deployed contract.
func NewVRFTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFTransactor, error) {
	contract, err := bindVRF(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFTransactor{contract: contract}, nil
}

// NewVRFFilterer creates a new log filterer instance of VRF, bound to a specific deployed contract.
func NewVRFFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFFilterer, error) {
	contract, err := bindVRF(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFFilterer{contract: contract}, nil
}

// bindVRF binds a generic wrapper to an already deployed contract.
func bindVRF(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VRFMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VRF *VRFRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRF.Contract.VRFCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VRF *VRFRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRF.Contract.VRFTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VRF *VRFRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRF.Contract.VRFTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VRF *VRFCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRF.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VRF *VRFTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRF.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VRF *VRFTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRF.Contract.contract.Transact(opts, method, params...)
}

// VrfManager is a free data retrieval call binding the contract method 0xe5e1b972.
//
// Solidity: function VrfManager() view returns(address)
func (_VRF *VRFCaller) VrfManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRF.contract.Call(opts, &out, "VrfManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VrfManager is a free data retrieval call binding the contract method 0xe5e1b972.
//
// Solidity: function VrfManager() view returns(address)
func (_VRF *VRFSession) VrfManager() (common.Address, error) {
	return _VRF.Contract.VrfManager(&_VRF.CallOpts)
}

// VrfManager is a free data retrieval call binding the contract method 0xe5e1b972.
//
// Solidity: function VrfManager() view returns(address)
func (_VRF *VRFCallerSession) VrfManager() (common.Address, error) {
	return _VRF.Contract.VrfManager(&_VRF.CallOpts)
}

// GetRequestStatus is a free data retrieval call binding the contract method 0xd8a4676f.
//
// Solidity: function getRequestStatus(uint256 _requestId) view returns(bool fulfilled, uint256[] randomWords)
func (_VRF *VRFCaller) GetRequestStatus(opts *bind.CallOpts, _requestId *big.Int) (struct {
	Fulfilled   bool
	RandomWords []*big.Int
}, error) {
	var out []interface{}
	err := _VRF.contract.Call(opts, &out, "getRequestStatus", _requestId)

	outstruct := new(struct {
		Fulfilled   bool
		RandomWords []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fulfilled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.RandomWords = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetRequestStatus is a free data retrieval call binding the contract method 0xd8a4676f.
//
// Solidity: function getRequestStatus(uint256 _requestId) view returns(bool fulfilled, uint256[] randomWords)
func (_VRF *VRFSession) GetRequestStatus(_requestId *big.Int) (struct {
	Fulfilled   bool
	RandomWords []*big.Int
}, error) {
	return _VRF.Contract.GetRequestStatus(&_VRF.CallOpts, _requestId)
}

// GetRequestStatus is a free data retrieval call binding the contract method 0xd8a4676f.
//
// Solidity: function getRequestStatus(uint256 _requestId) view returns(bool fulfilled, uint256[] randomWords)
func (_VRF *VRFCallerSession) GetRequestStatus(_requestId *big.Int) (struct {
	Fulfilled   bool
	RandomWords []*big.Int
}, error) {
	return _VRF.Contract.GetRequestStatus(&_VRF.CallOpts, _requestId)
}

// LastRequestId is a free data retrieval call binding the contract method 0xfc2a88c3.
//
// Solidity: function lastRequestId() view returns(uint256)
func (_VRF *VRFCaller) LastRequestId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRF.contract.Call(opts, &out, "lastRequestId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRequestId is a free data retrieval call binding the contract method 0xfc2a88c3.
//
// Solidity: function lastRequestId() view returns(uint256)
func (_VRF *VRFSession) LastRequestId() (*big.Int, error) {
	return _VRF.Contract.LastRequestId(&_VRF.CallOpts)
}

// LastRequestId is a free data retrieval call binding the contract method 0xfc2a88c3.
//
// Solidity: function lastRequestId() view returns(uint256)
func (_VRF *VRFCallerSession) LastRequestId() (*big.Int, error) {
	return _VRF.Contract.LastRequestId(&_VRF.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VRF *VRFCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRF.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VRF *VRFSession) Owner() (common.Address, error) {
	return _VRF.Contract.Owner(&_VRF.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VRF *VRFCallerSession) Owner() (common.Address, error) {
	return _VRF.Contract.Owner(&_VRF.CallOpts)
}

// RequestIds is a free data retrieval call binding the contract method 0x8796ba8c.
//
// Solidity: function requestIds(uint256 ) view returns(uint256)
func (_VRF *VRFCaller) RequestIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VRF.contract.Call(opts, &out, "requestIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestIds is a free data retrieval call binding the contract method 0x8796ba8c.
//
// Solidity: function requestIds(uint256 ) view returns(uint256)
func (_VRF *VRFSession) RequestIds(arg0 *big.Int) (*big.Int, error) {
	return _VRF.Contract.RequestIds(&_VRF.CallOpts, arg0)
}

// RequestIds is a free data retrieval call binding the contract method 0x8796ba8c.
//
// Solidity: function requestIds(uint256 ) view returns(uint256)
func (_VRF *VRFCallerSession) RequestIds(arg0 *big.Int) (*big.Int, error) {
	return _VRF.Contract.RequestIds(&_VRF.CallOpts, arg0)
}

// RequestMappings is a free data retrieval call binding the contract method 0x873767b4.
//
// Solidity: function requestMappings(uint256 ) view returns(bool fulfilled)
func (_VRF *VRFCaller) RequestMappings(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _VRF.contract.Call(opts, &out, "requestMappings", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RequestMappings is a free data retrieval call binding the contract method 0x873767b4.
//
// Solidity: function requestMappings(uint256 ) view returns(bool fulfilled)
func (_VRF *VRFSession) RequestMappings(arg0 *big.Int) (bool, error) {
	return _VRF.Contract.RequestMappings(&_VRF.CallOpts, arg0)
}

// RequestMappings is a free data retrieval call binding the contract method 0x873767b4.
//
// Solidity: function requestMappings(uint256 ) view returns(bool fulfilled)
func (_VRF *VRFCallerSession) RequestMappings(arg0 *big.Int) (bool, error) {
	return _VRF.Contract.RequestMappings(&_VRF.CallOpts, arg0)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0x38ba4614.
//
// Solidity: function fulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_VRF *VRFTransactor) FulfillRandomWords(opts *bind.TransactOpts, requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _VRF.contract.Transact(opts, "fulfillRandomWords", requestId, randomWords)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0x38ba4614.
//
// Solidity: function fulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_VRF *VRFSession) FulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _VRF.Contract.FulfillRandomWords(&_VRF.TransactOpts, requestId, randomWords)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0x38ba4614.
//
// Solidity: function fulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_VRF *VRFTransactorSession) FulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _VRF.Contract.FulfillRandomWords(&_VRF.TransactOpts, requestId, randomWords)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initOwner, address Manager) returns()
func (_VRF *VRFTransactor) Initialize(opts *bind.TransactOpts, initOwner common.Address, Manager common.Address) (*types.Transaction, error) {
	return _VRF.contract.Transact(opts, "initialize", initOwner, Manager)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initOwner, address Manager) returns()
func (_VRF *VRFSession) Initialize(initOwner common.Address, Manager common.Address) (*types.Transaction, error) {
	return _VRF.Contract.Initialize(&_VRF.TransactOpts, initOwner, Manager)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initOwner, address Manager) returns()
func (_VRF *VRFTransactorSession) Initialize(initOwner common.Address, Manager common.Address) (*types.Transaction, error) {
	return _VRF.Contract.Initialize(&_VRF.TransactOpts, initOwner, Manager)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VRF *VRFTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRF.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VRF *VRFSession) RenounceOwnership() (*types.Transaction, error) {
	return _VRF.Contract.RenounceOwnership(&_VRF.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VRF *VRFTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VRF.Contract.RenounceOwnership(&_VRF.TransactOpts)
}

// RequestRandWords is a paid mutator transaction binding the contract method 0xf68b2da8.
//
// Solidity: function requestRandWords(uint256 requestId, uint256 numwords) returns()
func (_VRF *VRFTransactor) RequestRandWords(opts *bind.TransactOpts, requestId *big.Int, numwords *big.Int) (*types.Transaction, error) {
	return _VRF.contract.Transact(opts, "requestRandWords", requestId, numwords)
}

// RequestRandWords is a paid mutator transaction binding the contract method 0xf68b2da8.
//
// Solidity: function requestRandWords(uint256 requestId, uint256 numwords) returns()
func (_VRF *VRFSession) RequestRandWords(requestId *big.Int, numwords *big.Int) (*types.Transaction, error) {
	return _VRF.Contract.RequestRandWords(&_VRF.TransactOpts, requestId, numwords)
}

// RequestRandWords is a paid mutator transaction binding the contract method 0xf68b2da8.
//
// Solidity: function requestRandWords(uint256 requestId, uint256 numwords) returns()
func (_VRF *VRFTransactorSession) RequestRandWords(requestId *big.Int, numwords *big.Int) (*types.Transaction, error) {
	return _VRF.Contract.RequestRandWords(&_VRF.TransactOpts, requestId, numwords)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address newManager) returns()
func (_VRF *VRFTransactor) SetManager(opts *bind.TransactOpts, newManager common.Address) (*types.Transaction, error) {
	return _VRF.contract.Transact(opts, "setManager", newManager)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address newManager) returns()
func (_VRF *VRFSession) SetManager(newManager common.Address) (*types.Transaction, error) {
	return _VRF.Contract.SetManager(&_VRF.TransactOpts, newManager)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address newManager) returns()
func (_VRF *VRFTransactorSession) SetManager(newManager common.Address) (*types.Transaction, error) {
	return _VRF.Contract.SetManager(&_VRF.TransactOpts, newManager)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VRF *VRFTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VRF.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VRF *VRFSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VRF.Contract.TransferOwnership(&_VRF.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VRF *VRFTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VRF.Contract.TransferOwnership(&_VRF.TransactOpts, newOwner)
}

// VRFFillRandomWordsIterator is returned from FilterFillRandomWords and is used to iterate over the raw logs and unpacked data for FillRandomWords events raised by the VRF contract.
type VRFFillRandomWordsIterator struct {
	Event *VRFFillRandomWords // Event containing the contract specifics and raw log

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
func (it *VRFFillRandomWordsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFFillRandomWords)
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
		it.Event = new(VRFFillRandomWords)
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
func (it *VRFFillRandomWordsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VRFFillRandomWordsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VRFFillRandomWords represents a FillRandomWords event raised by the VRF contract.
type VRFFillRandomWords struct {
	RequestId   *big.Int
	RandomWords []*big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFillRandomWords is a free log retrieval operation binding the contract event 0xf3cb4deb0441dd096356debf166f879d78cadc19e4b94053c8bea6d3940de93a.
//
// Solidity: event FillRandomWords(uint256 requestId, uint256[] randomWords)
func (_VRF *VRFFilterer) FilterFillRandomWords(opts *bind.FilterOpts) (*VRFFillRandomWordsIterator, error) {

	logs, sub, err := _VRF.contract.FilterLogs(opts, "FillRandomWords")
	if err != nil {
		return nil, err
	}
	return &VRFFillRandomWordsIterator{contract: _VRF.contract, event: "FillRandomWords", logs: logs, sub: sub}, nil
}

// WatchFillRandomWords is a free log subscription operation binding the contract event 0xf3cb4deb0441dd096356debf166f879d78cadc19e4b94053c8bea6d3940de93a.
//
// Solidity: event FillRandomWords(uint256 requestId, uint256[] randomWords)
func (_VRF *VRFFilterer) WatchFillRandomWords(opts *bind.WatchOpts, sink chan<- *VRFFillRandomWords) (event.Subscription, error) {

	logs, sub, err := _VRF.contract.WatchLogs(opts, "FillRandomWords")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VRFFillRandomWords)
				if err := _VRF.contract.UnpackLog(event, "FillRandomWords", log); err != nil {
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

// ParseFillRandomWords is a log parse operation binding the contract event 0xf3cb4deb0441dd096356debf166f879d78cadc19e4b94053c8bea6d3940de93a.
//
// Solidity: event FillRandomWords(uint256 requestId, uint256[] randomWords)
func (_VRF *VRFFilterer) ParseFillRandomWords(log types.Log) (*VRFFillRandomWords, error) {
	event := new(VRFFillRandomWords)
	if err := _VRF.contract.UnpackLog(event, "FillRandomWords", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VRFInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the VRF contract.
type VRFInitializedIterator struct {
	Event *VRFInitialized // Event containing the contract specifics and raw log

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
func (it *VRFInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFInitialized)
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
		it.Event = new(VRFInitialized)
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
func (it *VRFInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VRFInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VRFInitialized represents a Initialized event raised by the VRF contract.
type VRFInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_VRF *VRFFilterer) FilterInitialized(opts *bind.FilterOpts) (*VRFInitializedIterator, error) {

	logs, sub, err := _VRF.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &VRFInitializedIterator{contract: _VRF.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_VRF *VRFFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *VRFInitialized) (event.Subscription, error) {

	logs, sub, err := _VRF.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VRFInitialized)
				if err := _VRF.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_VRF *VRFFilterer) ParseInitialized(log types.Log) (*VRFInitialized, error) {
	event := new(VRFInitialized)
	if err := _VRF.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VRFOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VRF contract.
type VRFOwnershipTransferredIterator struct {
	Event *VRFOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VRFOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFOwnershipTransferred)
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
		it.Event = new(VRFOwnershipTransferred)
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
func (it *VRFOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VRFOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VRFOwnershipTransferred represents a OwnershipTransferred event raised by the VRF contract.
type VRFOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VRF *VRFFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VRFOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VRF.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VRFOwnershipTransferredIterator{contract: _VRF.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VRF *VRFFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VRF.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VRFOwnershipTransferred)
				if err := _VRF.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VRF *VRFFilterer) ParseOwnershipTransferred(log types.Log) (*VRFOwnershipTransferred, error) {
	event := new(VRFOwnershipTransferred)
	if err := _VRF.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VRFRequestSendIterator is returned from FilterRequestSend and is used to iterate over the raw logs and unpacked data for RequestSend events raised by the VRF contract.
type VRFRequestSendIterator struct {
	Event *VRFRequestSend // Event containing the contract specifics and raw log

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
func (it *VRFRequestSendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFRequestSend)
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
		it.Event = new(VRFRequestSend)
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
func (it *VRFRequestSendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VRFRequestSendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VRFRequestSend represents a RequestSend event raised by the VRF contract.
type VRFRequestSend struct {
	RequestId *big.Int
	Numwords  *big.Int
	Current   common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequestSend is a free log retrieval operation binding the contract event 0xb0089ea4415f0160195bcb1ffe462e932fb0012a47725589204b54981a6306f7.
//
// Solidity: event RequestSend(uint256 requestId, uint256 numwords, address current)
func (_VRF *VRFFilterer) FilterRequestSend(opts *bind.FilterOpts) (*VRFRequestSendIterator, error) {

	logs, sub, err := _VRF.contract.FilterLogs(opts, "RequestSend")
	if err != nil {
		return nil, err
	}
	return &VRFRequestSendIterator{contract: _VRF.contract, event: "RequestSend", logs: logs, sub: sub}, nil
}

// WatchRequestSend is a free log subscription operation binding the contract event 0xb0089ea4415f0160195bcb1ffe462e932fb0012a47725589204b54981a6306f7.
//
// Solidity: event RequestSend(uint256 requestId, uint256 numwords, address current)
func (_VRF *VRFFilterer) WatchRequestSend(opts *bind.WatchOpts, sink chan<- *VRFRequestSend) (event.Subscription, error) {

	logs, sub, err := _VRF.contract.WatchLogs(opts, "RequestSend")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VRFRequestSend)
				if err := _VRF.contract.UnpackLog(event, "RequestSend", log); err != nil {
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

// ParseRequestSend is a log parse operation binding the contract event 0xb0089ea4415f0160195bcb1ffe462e932fb0012a47725589204b54981a6306f7.
//
// Solidity: event RequestSend(uint256 requestId, uint256 numwords, address current)
func (_VRF *VRFFilterer) ParseRequestSend(log types.Log) (*VRFRequestSend, error) {
	event := new(VRFRequestSend)
	if err := _VRF.contract.UnpackLog(event, "RequestSend", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
