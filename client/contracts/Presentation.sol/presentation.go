// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package client

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
)

// PresentationMetaData contains all meta data concerning the Presentation contract.
var PresentationMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"VOTE_MAX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VOTE_MIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getVote\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_metadata\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_ranking\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pieceCID\",\"type\":\"bytes\"}],\"name\":\"updateCID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"grade\",\"type\":\"uint8\"}],\"name\":\"updateRanking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"votes\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PresentationABI is the input ABI used to generate the binding from.
// Deprecated: Use PresentationMetaData.ABI instead.
var PresentationABI = PresentationMetaData.ABI

// Presentation is an auto generated Go binding around an Ethereum contract.
type Presentation struct {
	PresentationCaller     // Read-only binding to the contract
	PresentationTransactor // Write-only binding to the contract
	PresentationFilterer   // Log filterer for contract events
}

// PresentationCaller is an auto generated read-only Go binding around an Ethereum contract.
type PresentationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PresentationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PresentationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PresentationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PresentationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PresentationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PresentationSession struct {
	Contract     *Presentation     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PresentationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PresentationCallerSession struct {
	Contract *PresentationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PresentationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PresentationTransactorSession struct {
	Contract     *PresentationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PresentationRaw is an auto generated low-level Go binding around an Ethereum contract.
type PresentationRaw struct {
	Contract *Presentation // Generic contract binding to access the raw methods on
}

// PresentationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PresentationCallerRaw struct {
	Contract *PresentationCaller // Generic read-only contract binding to access the raw methods on
}

// PresentationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PresentationTransactorRaw struct {
	Contract *PresentationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPresentation creates a new instance of Presentation, bound to a specific deployed contract.
func NewPresentation(address common.Address, backend bind.ContractBackend) (*Presentation, error) {
	contract, err := bindPresentation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Presentation{PresentationCaller: PresentationCaller{contract: contract}, PresentationTransactor: PresentationTransactor{contract: contract}, PresentationFilterer: PresentationFilterer{contract: contract}}, nil
}

// NewPresentationCaller creates a new read-only instance of Presentation, bound to a specific deployed contract.
func NewPresentationCaller(address common.Address, caller bind.ContractCaller) (*PresentationCaller, error) {
	contract, err := bindPresentation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PresentationCaller{contract: contract}, nil
}

// NewPresentationTransactor creates a new write-only instance of Presentation, bound to a specific deployed contract.
func NewPresentationTransactor(address common.Address, transactor bind.ContractTransactor) (*PresentationTransactor, error) {
	contract, err := bindPresentation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PresentationTransactor{contract: contract}, nil
}

// NewPresentationFilterer creates a new log filterer instance of Presentation, bound to a specific deployed contract.
func NewPresentationFilterer(address common.Address, filterer bind.ContractFilterer) (*PresentationFilterer, error) {
	contract, err := bindPresentation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PresentationFilterer{contract: contract}, nil
}

// bindPresentation binds a generic wrapper to an already deployed contract.
func bindPresentation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PresentationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Presentation *PresentationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Presentation.Contract.PresentationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Presentation *PresentationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Presentation.Contract.PresentationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Presentation *PresentationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Presentation.Contract.PresentationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Presentation *PresentationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Presentation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Presentation *PresentationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Presentation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Presentation *PresentationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Presentation.Contract.contract.Transact(opts, method, params...)
}

// VOTEMAX is a free data retrieval call binding the contract method 0xb41cea53.
//
// Solidity: function VOTE_MAX() view returns(uint256)
func (_Presentation *PresentationCaller) VOTEMAX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Presentation.contract.Call(opts, &out, "VOTE_MAX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VOTEMAX is a free data retrieval call binding the contract method 0xb41cea53.
//
// Solidity: function VOTE_MAX() view returns(uint256)
func (_Presentation *PresentationSession) VOTEMAX() (*big.Int, error) {
	return _Presentation.Contract.VOTEMAX(&_Presentation.CallOpts)
}

// VOTEMAX is a free data retrieval call binding the contract method 0xb41cea53.
//
// Solidity: function VOTE_MAX() view returns(uint256)
func (_Presentation *PresentationCallerSession) VOTEMAX() (*big.Int, error) {
	return _Presentation.Contract.VOTEMAX(&_Presentation.CallOpts)
}

// VOTEMIN is a free data retrieval call binding the contract method 0x8fb5557f.
//
// Solidity: function VOTE_MIN() view returns(uint256)
func (_Presentation *PresentationCaller) VOTEMIN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Presentation.contract.Call(opts, &out, "VOTE_MIN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VOTEMIN is a free data retrieval call binding the contract method 0x8fb5557f.
//
// Solidity: function VOTE_MIN() view returns(uint256)
func (_Presentation *PresentationSession) VOTEMIN() (*big.Int, error) {
	return _Presentation.Contract.VOTEMIN(&_Presentation.CallOpts)
}

// VOTEMIN is a free data retrieval call binding the contract method 0x8fb5557f.
//
// Solidity: function VOTE_MIN() view returns(uint256)
func (_Presentation *PresentationCallerSession) VOTEMIN() (*big.Int, error) {
	return _Presentation.Contract.VOTEMIN(&_Presentation.CallOpts)
}

// GetVote is a free data retrieval call binding the contract method 0x8d337b81.
//
// Solidity: function getVote(address _voter) view returns(uint8)
func (_Presentation *PresentationCaller) GetVote(opts *bind.CallOpts, _voter common.Address) (uint8, error) {
	var out []interface{}
	err := _Presentation.contract.Call(opts, &out, "getVote", _voter)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetVote is a free data retrieval call binding the contract method 0x8d337b81.
//
// Solidity: function getVote(address _voter) view returns(uint8)
func (_Presentation *PresentationSession) GetVote(_voter common.Address) (uint8, error) {
	return _Presentation.Contract.GetVote(&_Presentation.CallOpts, _voter)
}

// GetVote is a free data retrieval call binding the contract method 0x8d337b81.
//
// Solidity: function getVote(address _voter) view returns(uint8)
func (_Presentation *PresentationCallerSession) GetVote(_voter common.Address) (uint8, error) {
	return _Presentation.Contract.GetVote(&_Presentation.CallOpts, _voter)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Presentation *PresentationCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Presentation.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Presentation *PresentationSession) Owner() (common.Address, error) {
	return _Presentation.Contract.Owner(&_Presentation.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Presentation *PresentationCallerSession) Owner() (common.Address, error) {
	return _Presentation.Contract.Owner(&_Presentation.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0x3c130d90.
//
// Solidity: function tokenURI() view returns(string)
func (_Presentation *PresentationCaller) TokenURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Presentation.contract.Call(opts, &out, "tokenURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0x3c130d90.
//
// Solidity: function tokenURI() view returns(string)
func (_Presentation *PresentationSession) TokenURI() (string, error) {
	return _Presentation.Contract.TokenURI(&_Presentation.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0x3c130d90.
//
// Solidity: function tokenURI() view returns(string)
func (_Presentation *PresentationCallerSession) TokenURI() (string, error) {
	return _Presentation.Contract.TokenURI(&_Presentation.CallOpts)
}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes(address ) view returns(uint8)
func (_Presentation *PresentationCaller) Votes(opts *bind.CallOpts, arg0 common.Address) (uint8, error) {
	var out []interface{}
	err := _Presentation.contract.Call(opts, &out, "votes", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes(address ) view returns(uint8)
func (_Presentation *PresentationSession) Votes(arg0 common.Address) (uint8, error) {
	return _Presentation.Contract.Votes(&_Presentation.CallOpts, arg0)
}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes(address ) view returns(uint8)
func (_Presentation *PresentationCallerSession) Votes(arg0 common.Address) (uint8, error) {
	return _Presentation.Contract.Votes(&_Presentation.CallOpts, arg0)
}

// Initialize is a paid mutator transaction binding the contract method 0x8beaf7d7.
//
// Solidity: function initialize(string _metadata, uint256 _ranking) returns()
func (_Presentation *PresentationTransactor) Initialize(opts *bind.TransactOpts, _metadata string, _ranking *big.Int) (*types.Transaction, error) {
	return _Presentation.contract.Transact(opts, "initialize", _metadata, _ranking)
}

// Initialize is a paid mutator transaction binding the contract method 0x8beaf7d7.
//
// Solidity: function initialize(string _metadata, uint256 _ranking) returns()
func (_Presentation *PresentationSession) Initialize(_metadata string, _ranking *big.Int) (*types.Transaction, error) {
	return _Presentation.Contract.Initialize(&_Presentation.TransactOpts, _metadata, _ranking)
}

// Initialize is a paid mutator transaction binding the contract method 0x8beaf7d7.
//
// Solidity: function initialize(string _metadata, uint256 _ranking) returns()
func (_Presentation *PresentationTransactorSession) Initialize(_metadata string, _ranking *big.Int) (*types.Transaction, error) {
	return _Presentation.Contract.Initialize(&_Presentation.TransactOpts, _metadata, _ranking)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Presentation *PresentationTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Presentation.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Presentation *PresentationSession) RenounceOwnership() (*types.Transaction, error) {
	return _Presentation.Contract.RenounceOwnership(&_Presentation.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Presentation *PresentationTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Presentation.Contract.RenounceOwnership(&_Presentation.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Presentation *PresentationTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Presentation.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Presentation *PresentationSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Presentation.Contract.TransferOwnership(&_Presentation.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Presentation *PresentationTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Presentation.Contract.TransferOwnership(&_Presentation.TransactOpts, newOwner)
}

// UpdateCID is a paid mutator transaction binding the contract method 0x264981ca.
//
// Solidity: function updateCID(bytes _pieceCID) returns()
func (_Presentation *PresentationTransactor) UpdateCID(opts *bind.TransactOpts, _pieceCID []byte) (*types.Transaction, error) {
	return _Presentation.contract.Transact(opts, "updateCID", _pieceCID)
}

// UpdateCID is a paid mutator transaction binding the contract method 0x264981ca.
//
// Solidity: function updateCID(bytes _pieceCID) returns()
func (_Presentation *PresentationSession) UpdateCID(_pieceCID []byte) (*types.Transaction, error) {
	return _Presentation.Contract.UpdateCID(&_Presentation.TransactOpts, _pieceCID)
}

// UpdateCID is a paid mutator transaction binding the contract method 0x264981ca.
//
// Solidity: function updateCID(bytes _pieceCID) returns()
func (_Presentation *PresentationTransactorSession) UpdateCID(_pieceCID []byte) (*types.Transaction, error) {
	return _Presentation.Contract.UpdateCID(&_Presentation.TransactOpts, _pieceCID)
}

// UpdateRanking is a paid mutator transaction binding the contract method 0x3cb6ee75.
//
// Solidity: function updateRanking(uint8 grade) returns()
func (_Presentation *PresentationTransactor) UpdateRanking(opts *bind.TransactOpts, grade uint8) (*types.Transaction, error) {
	return _Presentation.contract.Transact(opts, "updateRanking", grade)
}

// UpdateRanking is a paid mutator transaction binding the contract method 0x3cb6ee75.
//
// Solidity: function updateRanking(uint8 grade) returns()
func (_Presentation *PresentationSession) UpdateRanking(grade uint8) (*types.Transaction, error) {
	return _Presentation.Contract.UpdateRanking(&_Presentation.TransactOpts, grade)
}

// UpdateRanking is a paid mutator transaction binding the contract method 0x3cb6ee75.
//
// Solidity: function updateRanking(uint8 grade) returns()
func (_Presentation *PresentationTransactorSession) UpdateRanking(grade uint8) (*types.Transaction, error) {
	return _Presentation.Contract.UpdateRanking(&_Presentation.TransactOpts, grade)
}

// PresentationInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Presentation contract.
type PresentationInitializedIterator struct {
	Event *PresentationInitialized // Event containing the contract specifics and raw log

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
func (it *PresentationInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PresentationInitialized)
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
		it.Event = new(PresentationInitialized)
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
func (it *PresentationInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PresentationInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PresentationInitialized represents a Initialized event raised by the Presentation contract.
type PresentationInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Presentation *PresentationFilterer) FilterInitialized(opts *bind.FilterOpts) (*PresentationInitializedIterator, error) {

	logs, sub, err := _Presentation.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PresentationInitializedIterator{contract: _Presentation.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Presentation *PresentationFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PresentationInitialized) (event.Subscription, error) {

	logs, sub, err := _Presentation.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PresentationInitialized)
				if err := _Presentation.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Presentation *PresentationFilterer) ParseInitialized(log types.Log) (*PresentationInitialized, error) {
	event := new(PresentationInitialized)
	if err := _Presentation.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PresentationOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Presentation contract.
type PresentationOwnershipTransferredIterator struct {
	Event *PresentationOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PresentationOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PresentationOwnershipTransferred)
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
		it.Event = new(PresentationOwnershipTransferred)
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
func (it *PresentationOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PresentationOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PresentationOwnershipTransferred represents a OwnershipTransferred event raised by the Presentation contract.
type PresentationOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Presentation *PresentationFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PresentationOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Presentation.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PresentationOwnershipTransferredIterator{contract: _Presentation.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Presentation *PresentationFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PresentationOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Presentation.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PresentationOwnershipTransferred)
				if err := _Presentation.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Presentation *PresentationFilterer) ParseOwnershipTransferred(log types.Log) (*PresentationOwnershipTransferred, error) {
	event := new(PresentationOwnershipTransferred)
	if err := _Presentation.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
