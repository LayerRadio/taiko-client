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

// AddressManagerMetaData contains all meta data concerning the AddressManager contract.
var AddressManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AM_INVALID_ADDRESS\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"domain\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"}],\"name\":\"AddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"domain\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"domain\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"setAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AddressManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressManagerMetaData.ABI instead.
var AddressManagerABI = AddressManagerMetaData.ABI

// AddressManager is an auto generated Go binding around an Ethereum contract.
type AddressManager struct {
	AddressManagerCaller     // Read-only binding to the contract
	AddressManagerTransactor // Write-only binding to the contract
	AddressManagerFilterer   // Log filterer for contract events
}

// AddressManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressManagerSession struct {
	Contract     *AddressManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressManagerCallerSession struct {
	Contract *AddressManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AddressManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressManagerTransactorSession struct {
	Contract     *AddressManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AddressManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressManagerRaw struct {
	Contract *AddressManager // Generic contract binding to access the raw methods on
}

// AddressManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressManagerCallerRaw struct {
	Contract *AddressManagerCaller // Generic read-only contract binding to access the raw methods on
}

// AddressManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressManagerTransactorRaw struct {
	Contract *AddressManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressManager creates a new instance of AddressManager, bound to a specific deployed contract.
func NewAddressManager(address common.Address, backend bind.ContractBackend) (*AddressManager, error) {
	contract, err := bindAddressManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddressManager{AddressManagerCaller: AddressManagerCaller{contract: contract}, AddressManagerTransactor: AddressManagerTransactor{contract: contract}, AddressManagerFilterer: AddressManagerFilterer{contract: contract}}, nil
}

// NewAddressManagerCaller creates a new read-only instance of AddressManager, bound to a specific deployed contract.
func NewAddressManagerCaller(address common.Address, caller bind.ContractCaller) (*AddressManagerCaller, error) {
	contract, err := bindAddressManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressManagerCaller{contract: contract}, nil
}

// NewAddressManagerTransactor creates a new write-only instance of AddressManager, bound to a specific deployed contract.
func NewAddressManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressManagerTransactor, error) {
	contract, err := bindAddressManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressManagerTransactor{contract: contract}, nil
}

// NewAddressManagerFilterer creates a new log filterer instance of AddressManager, bound to a specific deployed contract.
func NewAddressManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressManagerFilterer, error) {
	contract, err := bindAddressManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressManagerFilterer{contract: contract}, nil
}

// bindAddressManager binds a generic wrapper to an already deployed contract.
func bindAddressManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AddressManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressManager *AddressManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressManager.Contract.AddressManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressManager *AddressManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressManager.Contract.AddressManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressManager *AddressManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressManager.Contract.AddressManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressManager *AddressManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressManager *AddressManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressManager *AddressManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressManager.Contract.contract.Transact(opts, method, params...)
}

// GetAddress is a free data retrieval call binding the contract method 0x656b06a4.
//
// Solidity: function getAddress(uint256 domain, bytes32 name) view returns(address)
func (_AddressManager *AddressManagerCaller) GetAddress(opts *bind.CallOpts, domain *big.Int, name [32]byte) (common.Address, error) {
	var out []interface{}
	err := _AddressManager.contract.Call(opts, &out, "getAddress", domain, name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x656b06a4.
//
// Solidity: function getAddress(uint256 domain, bytes32 name) view returns(address)
func (_AddressManager *AddressManagerSession) GetAddress(domain *big.Int, name [32]byte) (common.Address, error) {
	return _AddressManager.Contract.GetAddress(&_AddressManager.CallOpts, domain, name)
}

// GetAddress is a free data retrieval call binding the contract method 0x656b06a4.
//
// Solidity: function getAddress(uint256 domain, bytes32 name) view returns(address)
func (_AddressManager *AddressManagerCallerSession) GetAddress(domain *big.Int, name [32]byte) (common.Address, error) {
	return _AddressManager.Contract.GetAddress(&_AddressManager.CallOpts, domain, name)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddressManager *AddressManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddressManager *AddressManagerSession) Owner() (common.Address, error) {
	return _AddressManager.Contract.Owner(&_AddressManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddressManager *AddressManagerCallerSession) Owner() (common.Address, error) {
	return _AddressManager.Contract.Owner(&_AddressManager.CallOpts)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_AddressManager *AddressManagerTransactor) Init(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressManager.contract.Transact(opts, "init")
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_AddressManager *AddressManagerSession) Init() (*types.Transaction, error) {
	return _AddressManager.Contract.Init(&_AddressManager.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_AddressManager *AddressManagerTransactorSession) Init() (*types.Transaction, error) {
	return _AddressManager.Contract.Init(&_AddressManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AddressManager *AddressManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AddressManager *AddressManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _AddressManager.Contract.RenounceOwnership(&_AddressManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AddressManager *AddressManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AddressManager.Contract.RenounceOwnership(&_AddressManager.TransactOpts)
}

// SetAddress is a paid mutator transaction binding the contract method 0xdecd8e39.
//
// Solidity: function setAddress(uint256 domain, bytes32 name, address newAddress) returns()
func (_AddressManager *AddressManagerTransactor) SetAddress(opts *bind.TransactOpts, domain *big.Int, name [32]byte, newAddress common.Address) (*types.Transaction, error) {
	return _AddressManager.contract.Transact(opts, "setAddress", domain, name, newAddress)
}

// SetAddress is a paid mutator transaction binding the contract method 0xdecd8e39.
//
// Solidity: function setAddress(uint256 domain, bytes32 name, address newAddress) returns()
func (_AddressManager *AddressManagerSession) SetAddress(domain *big.Int, name [32]byte, newAddress common.Address) (*types.Transaction, error) {
	return _AddressManager.Contract.SetAddress(&_AddressManager.TransactOpts, domain, name, newAddress)
}

// SetAddress is a paid mutator transaction binding the contract method 0xdecd8e39.
//
// Solidity: function setAddress(uint256 domain, bytes32 name, address newAddress) returns()
func (_AddressManager *AddressManagerTransactorSession) SetAddress(domain *big.Int, name [32]byte, newAddress common.Address) (*types.Transaction, error) {
	return _AddressManager.Contract.SetAddress(&_AddressManager.TransactOpts, domain, name, newAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AddressManager *AddressManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AddressManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AddressManager *AddressManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AddressManager.Contract.TransferOwnership(&_AddressManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AddressManager *AddressManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AddressManager.Contract.TransferOwnership(&_AddressManager.TransactOpts, newOwner)
}

// AddressManagerAddressSetIterator is returned from FilterAddressSet and is used to iterate over the raw logs and unpacked data for AddressSet events raised by the AddressManager contract.
type AddressManagerAddressSetIterator struct {
	Event *AddressManagerAddressSet // Event containing the contract specifics and raw log

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
func (it *AddressManagerAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressManagerAddressSet)
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
		it.Event = new(AddressManagerAddressSet)
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
func (it *AddressManagerAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressManagerAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressManagerAddressSet represents a AddressSet event raised by the AddressManager contract.
type AddressManagerAddressSet struct {
	Domain     *big.Int
	Name       [32]byte
	NewAddress common.Address
	OldAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddressSet is a free log retrieval operation binding the contract event 0xe41a6e8584d6e19a0dfc5f9331be4ebe61b5f025d45da164c9ca6ee9b837cea9.
//
// Solidity: event AddressSet(uint256 indexed domain, bytes32 indexed name, address newAddress, address oldAddress)
func (_AddressManager *AddressManagerFilterer) FilterAddressSet(opts *bind.FilterOpts, domain []*big.Int, name [][32]byte) (*AddressManagerAddressSetIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _AddressManager.contract.FilterLogs(opts, "AddressSet", domainRule, nameRule)
	if err != nil {
		return nil, err
	}
	return &AddressManagerAddressSetIterator{contract: _AddressManager.contract, event: "AddressSet", logs: logs, sub: sub}, nil
}

// WatchAddressSet is a free log subscription operation binding the contract event 0xe41a6e8584d6e19a0dfc5f9331be4ebe61b5f025d45da164c9ca6ee9b837cea9.
//
// Solidity: event AddressSet(uint256 indexed domain, bytes32 indexed name, address newAddress, address oldAddress)
func (_AddressManager *AddressManagerFilterer) WatchAddressSet(opts *bind.WatchOpts, sink chan<- *AddressManagerAddressSet, domain []*big.Int, name [][32]byte) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _AddressManager.contract.WatchLogs(opts, "AddressSet", domainRule, nameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressManagerAddressSet)
				if err := _AddressManager.contract.UnpackLog(event, "AddressSet", log); err != nil {
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

// ParseAddressSet is a log parse operation binding the contract event 0xe41a6e8584d6e19a0dfc5f9331be4ebe61b5f025d45da164c9ca6ee9b837cea9.
//
// Solidity: event AddressSet(uint256 indexed domain, bytes32 indexed name, address newAddress, address oldAddress)
func (_AddressManager *AddressManagerFilterer) ParseAddressSet(log types.Log) (*AddressManagerAddressSet, error) {
	event := new(AddressManagerAddressSet)
	if err := _AddressManager.contract.UnpackLog(event, "AddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AddressManager contract.
type AddressManagerInitializedIterator struct {
	Event *AddressManagerInitialized // Event containing the contract specifics and raw log

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
func (it *AddressManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressManagerInitialized)
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
		it.Event = new(AddressManagerInitialized)
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
func (it *AddressManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressManagerInitialized represents a Initialized event raised by the AddressManager contract.
type AddressManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AddressManager *AddressManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*AddressManagerInitializedIterator, error) {

	logs, sub, err := _AddressManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AddressManagerInitializedIterator{contract: _AddressManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AddressManager *AddressManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AddressManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _AddressManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressManagerInitialized)
				if err := _AddressManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AddressManager *AddressManagerFilterer) ParseInitialized(log types.Log) (*AddressManagerInitialized, error) {
	event := new(AddressManagerInitialized)
	if err := _AddressManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AddressManager contract.
type AddressManagerOwnershipTransferredIterator struct {
	Event *AddressManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AddressManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressManagerOwnershipTransferred)
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
		it.Event = new(AddressManagerOwnershipTransferred)
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
func (it *AddressManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressManagerOwnershipTransferred represents a OwnershipTransferred event raised by the AddressManager contract.
type AddressManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AddressManager *AddressManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AddressManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AddressManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AddressManagerOwnershipTransferredIterator{contract: _AddressManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AddressManager *AddressManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AddressManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AddressManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressManagerOwnershipTransferred)
				if err := _AddressManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AddressManager *AddressManagerFilterer) ParseOwnershipTransferred(log types.Log) (*AddressManagerOwnershipTransferred, error) {
	event := new(AddressManagerOwnershipTransferred)
	if err := _AddressManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
