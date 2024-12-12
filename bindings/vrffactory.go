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

// VRFFactoryMetaData contains all meta data concerning the VRFFactory contract.
var VRFFactoryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createProxy\",\"inputs\":[{\"name\":\"imlementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createProxyByOZ\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"manager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ProxyCreated\",\"inputs\":[{\"name\":\"proxyAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"FailedDeployment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	Bin: "0x6080604052348015600f57600080fd5b5061035f8061001f6000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80636140c54c1461003b578063d1158f771461006a575b600080fd5b61004e6100493660046102d4565b61007d565b6040516001600160a01b03909116815260200160405180910390f35b61004e6100783660046102f6565b610190565b604051693d602d80600a3d3981f360b01b602082015269363d3d373d3d3d363d7360b01b602a8201526bffffffffffffffffffffffff19606083901b1660348201526e5af43d82803e903d91602b57fd5bf360881b60488201526000908190605701604051602081830303815290604052905060008151602083016000f090506001600160a01b03811661014e5760405162461bcd60e51b8152602060048201526014602482015273141c9bde1e4810dc99585d19590819985a5b195960621b604482015260640160405180910390fd5b6040516001600160a01b03821681527efffc2da0b561cae30d9826d37709e9421c4725faebc226cbbb7ef5fc5e73499060200160405180910390a19392505050565b60008061019c84610246565b60405163485cc95560e01b81523360048201526001600160a01b0385811660248301529192509082169063485cc95590604401600060405180830381600087803b1580156101e957600080fd5b505af11580156101fd573d6000803e3d6000fd5b50506040516001600160a01b03841681527efffc2da0b561cae30d9826d37709e9421c4725faebc226cbbb7ef5fc5e73499250602001905060405180910390a190505b92915050565b6000610240828280763d602d80600a3d3981f3363d3d373d3d3d363d730000008360601b60e81c176000526e5af43d82803e903d91602b57fd5bf38360781b176020526037600983f090506001600160a01b0381166102405760405163b06ebf3d60e01b815260040160405180910390fd5b80356001600160a01b03811681146102cf57600080fd5b919050565b6000602082840312156102e657600080fd5b6102ef826102b8565b9392505050565b6000806040838503121561030957600080fd5b610312836102b8565b9150610320602084016102b8565b9050925092905056fea264697066735822122026df86b371abaddfa2b12ccf22ecbee4e50c7f3e5e7bb74dfc2bd73c1bbeb23464736f6c634300081b0033",
}

// VRFFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use VRFFactoryMetaData.ABI instead.
var VRFFactoryABI = VRFFactoryMetaData.ABI

// VRFFactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VRFFactoryMetaData.Bin instead.
var VRFFactoryBin = VRFFactoryMetaData.Bin

// DeployVRFFactory deploys a new Ethereum contract, binding an instance of VRFFactory to it.
func DeployVRFFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VRFFactory, error) {
	parsed, err := VRFFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRFFactory{VRFFactoryCaller: VRFFactoryCaller{contract: contract}, VRFFactoryTransactor: VRFFactoryTransactor{contract: contract}, VRFFactoryFilterer: VRFFactoryFilterer{contract: contract}}, nil
}

// VRFFactory is an auto generated Go binding around an Ethereum contract.
type VRFFactory struct {
	VRFFactoryCaller     // Read-only binding to the contract
	VRFFactoryTransactor // Write-only binding to the contract
	VRFFactoryFilterer   // Log filterer for contract events
}

// VRFFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type VRFFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VRFFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VRFFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VRFFactorySession struct {
	Contract     *VRFFactory       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VRFFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VRFFactoryCallerSession struct {
	Contract *VRFFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// VRFFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VRFFactoryTransactorSession struct {
	Contract     *VRFFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// VRFFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type VRFFactoryRaw struct {
	Contract *VRFFactory // Generic contract binding to access the raw methods on
}

// VRFFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VRFFactoryCallerRaw struct {
	Contract *VRFFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// VRFFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VRFFactoryTransactorRaw struct {
	Contract *VRFFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVRFFactory creates a new instance of VRFFactory, bound to a specific deployed contract.
func NewVRFFactory(address common.Address, backend bind.ContractBackend) (*VRFFactory, error) {
	contract, err := bindVRFFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFFactory{VRFFactoryCaller: VRFFactoryCaller{contract: contract}, VRFFactoryTransactor: VRFFactoryTransactor{contract: contract}, VRFFactoryFilterer: VRFFactoryFilterer{contract: contract}}, nil
}

// NewVRFFactoryCaller creates a new read-only instance of VRFFactory, bound to a specific deployed contract.
func NewVRFFactoryCaller(address common.Address, caller bind.ContractCaller) (*VRFFactoryCaller, error) {
	contract, err := bindVRFFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFFactoryCaller{contract: contract}, nil
}

// NewVRFFactoryTransactor creates a new write-only instance of VRFFactory, bound to a specific deployed contract.
func NewVRFFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFFactoryTransactor, error) {
	contract, err := bindVRFFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFFactoryTransactor{contract: contract}, nil
}

// NewVRFFactoryFilterer creates a new log filterer instance of VRFFactory, bound to a specific deployed contract.
func NewVRFFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFFactoryFilterer, error) {
	contract, err := bindVRFFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFFactoryFilterer{contract: contract}, nil
}

// bindVRFFactory binds a generic wrapper to an already deployed contract.
func bindVRFFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VRFFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VRFFactory *VRFFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFFactory.Contract.VRFFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VRFFactory *VRFFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFFactory.Contract.VRFFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VRFFactory *VRFFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFFactory.Contract.VRFFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VRFFactory *VRFFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VRFFactory *VRFFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VRFFactory *VRFFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFFactory.Contract.contract.Transact(opts, method, params...)
}

// CreateProxy is a paid mutator transaction binding the contract method 0x6140c54c.
//
// Solidity: function createProxy(address imlementation) returns(address)
func (_VRFFactory *VRFFactoryTransactor) CreateProxy(opts *bind.TransactOpts, imlementation common.Address) (*types.Transaction, error) {
	return _VRFFactory.contract.Transact(opts, "createProxy", imlementation)
}

// CreateProxy is a paid mutator transaction binding the contract method 0x6140c54c.
//
// Solidity: function createProxy(address imlementation) returns(address)
func (_VRFFactory *VRFFactorySession) CreateProxy(imlementation common.Address) (*types.Transaction, error) {
	return _VRFFactory.Contract.CreateProxy(&_VRFFactory.TransactOpts, imlementation)
}

// CreateProxy is a paid mutator transaction binding the contract method 0x6140c54c.
//
// Solidity: function createProxy(address imlementation) returns(address)
func (_VRFFactory *VRFFactoryTransactorSession) CreateProxy(imlementation common.Address) (*types.Transaction, error) {
	return _VRFFactory.Contract.CreateProxy(&_VRFFactory.TransactOpts, imlementation)
}

// CreateProxyByOZ is a paid mutator transaction binding the contract method 0xd1158f77.
//
// Solidity: function createProxyByOZ(address implementation, address manager) returns(address)
func (_VRFFactory *VRFFactoryTransactor) CreateProxyByOZ(opts *bind.TransactOpts, implementation common.Address, manager common.Address) (*types.Transaction, error) {
	return _VRFFactory.contract.Transact(opts, "createProxyByOZ", implementation, manager)
}

// CreateProxyByOZ is a paid mutator transaction binding the contract method 0xd1158f77.
//
// Solidity: function createProxyByOZ(address implementation, address manager) returns(address)
func (_VRFFactory *VRFFactorySession) CreateProxyByOZ(implementation common.Address, manager common.Address) (*types.Transaction, error) {
	return _VRFFactory.Contract.CreateProxyByOZ(&_VRFFactory.TransactOpts, implementation, manager)
}

// CreateProxyByOZ is a paid mutator transaction binding the contract method 0xd1158f77.
//
// Solidity: function createProxyByOZ(address implementation, address manager) returns(address)
func (_VRFFactory *VRFFactoryTransactorSession) CreateProxyByOZ(implementation common.Address, manager common.Address) (*types.Transaction, error) {
	return _VRFFactory.Contract.CreateProxyByOZ(&_VRFFactory.TransactOpts, implementation, manager)
}

// VRFFactoryProxyCreatedIterator is returned from FilterProxyCreated and is used to iterate over the raw logs and unpacked data for ProxyCreated events raised by the VRFFactory contract.
type VRFFactoryProxyCreatedIterator struct {
	Event *VRFFactoryProxyCreated // Event containing the contract specifics and raw log

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
func (it *VRFFactoryProxyCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFFactoryProxyCreated)
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
		it.Event = new(VRFFactoryProxyCreated)
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
func (it *VRFFactoryProxyCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VRFFactoryProxyCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VRFFactoryProxyCreated represents a ProxyCreated event raised by the VRFFactory contract.
type VRFFactoryProxyCreated struct {
	ProxyAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProxyCreated is a free log retrieval operation binding the contract event 0x00fffc2da0b561cae30d9826d37709e9421c4725faebc226cbbb7ef5fc5e7349.
//
// Solidity: event ProxyCreated(address proxyAddress)
func (_VRFFactory *VRFFactoryFilterer) FilterProxyCreated(opts *bind.FilterOpts) (*VRFFactoryProxyCreatedIterator, error) {

	logs, sub, err := _VRFFactory.contract.FilterLogs(opts, "ProxyCreated")
	if err != nil {
		return nil, err
	}
	return &VRFFactoryProxyCreatedIterator{contract: _VRFFactory.contract, event: "ProxyCreated", logs: logs, sub: sub}, nil
}

// WatchProxyCreated is a free log subscription operation binding the contract event 0x00fffc2da0b561cae30d9826d37709e9421c4725faebc226cbbb7ef5fc5e7349.
//
// Solidity: event ProxyCreated(address proxyAddress)
func (_VRFFactory *VRFFactoryFilterer) WatchProxyCreated(opts *bind.WatchOpts, sink chan<- *VRFFactoryProxyCreated) (event.Subscription, error) {

	logs, sub, err := _VRFFactory.contract.WatchLogs(opts, "ProxyCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VRFFactoryProxyCreated)
				if err := _VRFFactory.contract.UnpackLog(event, "ProxyCreated", log); err != nil {
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

// ParseProxyCreated is a log parse operation binding the contract event 0x00fffc2da0b561cae30d9826d37709e9421c4725faebc226cbbb7ef5fc5e7349.
//
// Solidity: event ProxyCreated(address proxyAddress)
func (_VRFFactory *VRFFactoryFilterer) ParseProxyCreated(log types.Log) (*VRFFactoryProxyCreated, error) {
	event := new(VRFFactoryProxyCreated)
	if err := _VRFFactory.contract.UnpackLog(event, "ProxyCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
