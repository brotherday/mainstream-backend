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

// DealRewarderMetaData contains all meta data concerning the DealRewarder contract.
var DealRewarderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"cidraw\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"addCID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"cidraw\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"provider\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"authorizeData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"method\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"flags\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"codec\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"name\":\"call_actor_id\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"cidProviders\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"cidSet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"cidSizes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"deal_id\",\"type\":\"uint64\"}],\"name\":\"claim_bounty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"unused\",\"type\":\"uint64\"}],\"name\":\"fund\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DealRewarderABI is the input ABI used to generate the binding from.
// Deprecated: Use DealRewarderMetaData.ABI instead.
var DealRewarderABI = DealRewarderMetaData.ABI

// DealRewarder is an auto generated Go binding around an Ethereum contract.
type DealRewarder struct {
	DealRewarderCaller     // Read-only binding to the contract
	DealRewarderTransactor // Write-only binding to the contract
	DealRewarderFilterer   // Log filterer for contract events
}

// DealRewarderCaller is an auto generated read-only Go binding around an Ethereum contract.
type DealRewarderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DealRewarderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DealRewarderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DealRewarderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DealRewarderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DealRewarderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DealRewarderSession struct {
	Contract     *DealRewarder     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DealRewarderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DealRewarderCallerSession struct {
	Contract *DealRewarderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// DealRewarderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DealRewarderTransactorSession struct {
	Contract     *DealRewarderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DealRewarderRaw is an auto generated low-level Go binding around an Ethereum contract.
type DealRewarderRaw struct {
	Contract *DealRewarder // Generic contract binding to access the raw methods on
}

// DealRewarderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DealRewarderCallerRaw struct {
	Contract *DealRewarderCaller // Generic read-only contract binding to access the raw methods on
}

// DealRewarderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DealRewarderTransactorRaw struct {
	Contract *DealRewarderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDealRewarder creates a new instance of DealRewarder, bound to a specific deployed contract.
func NewDealRewarder(address common.Address, backend bind.ContractBackend) (*DealRewarder, error) {
	contract, err := bindDealRewarder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DealRewarder{DealRewarderCaller: DealRewarderCaller{contract: contract}, DealRewarderTransactor: DealRewarderTransactor{contract: contract}, DealRewarderFilterer: DealRewarderFilterer{contract: contract}}, nil
}

// NewDealRewarderCaller creates a new read-only instance of DealRewarder, bound to a specific deployed contract.
func NewDealRewarderCaller(address common.Address, caller bind.ContractCaller) (*DealRewarderCaller, error) {
	contract, err := bindDealRewarder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DealRewarderCaller{contract: contract}, nil
}

// NewDealRewarderTransactor creates a new write-only instance of DealRewarder, bound to a specific deployed contract.
func NewDealRewarderTransactor(address common.Address, transactor bind.ContractTransactor) (*DealRewarderTransactor, error) {
	contract, err := bindDealRewarder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DealRewarderTransactor{contract: contract}, nil
}

// NewDealRewarderFilterer creates a new log filterer instance of DealRewarder, bound to a specific deployed contract.
func NewDealRewarderFilterer(address common.Address, filterer bind.ContractFilterer) (*DealRewarderFilterer, error) {
	contract, err := bindDealRewarder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DealRewarderFilterer{contract: contract}, nil
}

// bindDealRewarder binds a generic wrapper to an already deployed contract.
func bindDealRewarder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DealRewarderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DealRewarder *DealRewarderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DealRewarder.Contract.DealRewarderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DealRewarder *DealRewarderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DealRewarder.Contract.DealRewarderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DealRewarder *DealRewarderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DealRewarder.Contract.DealRewarderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DealRewarder *DealRewarderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DealRewarder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DealRewarder *DealRewarderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DealRewarder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DealRewarder *DealRewarderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DealRewarder.Contract.contract.Transact(opts, method, params...)
}

// CidProviders is a free data retrieval call binding the contract method 0x1a7fd6c6.
//
// Solidity: function cidProviders(bytes , uint64 ) view returns(bool)
func (_DealRewarder *DealRewarderCaller) CidProviders(opts *bind.CallOpts, arg0 []byte, arg1 uint64) (bool, error) {
	var out []interface{}
	err := _DealRewarder.contract.Call(opts, &out, "cidProviders", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CidProviders is a free data retrieval call binding the contract method 0x1a7fd6c6.
//
// Solidity: function cidProviders(bytes , uint64 ) view returns(bool)
func (_DealRewarder *DealRewarderSession) CidProviders(arg0 []byte, arg1 uint64) (bool, error) {
	return _DealRewarder.Contract.CidProviders(&_DealRewarder.CallOpts, arg0, arg1)
}

// CidProviders is a free data retrieval call binding the contract method 0x1a7fd6c6.
//
// Solidity: function cidProviders(bytes , uint64 ) view returns(bool)
func (_DealRewarder *DealRewarderCallerSession) CidProviders(arg0 []byte, arg1 uint64) (bool, error) {
	return _DealRewarder.Contract.CidProviders(&_DealRewarder.CallOpts, arg0, arg1)
}

// CidSet is a free data retrieval call binding the contract method 0x3d7650b2.
//
// Solidity: function cidSet(bytes ) view returns(bool)
func (_DealRewarder *DealRewarderCaller) CidSet(opts *bind.CallOpts, arg0 []byte) (bool, error) {
	var out []interface{}
	err := _DealRewarder.contract.Call(opts, &out, "cidSet", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CidSet is a free data retrieval call binding the contract method 0x3d7650b2.
//
// Solidity: function cidSet(bytes ) view returns(bool)
func (_DealRewarder *DealRewarderSession) CidSet(arg0 []byte) (bool, error) {
	return _DealRewarder.Contract.CidSet(&_DealRewarder.CallOpts, arg0)
}

// CidSet is a free data retrieval call binding the contract method 0x3d7650b2.
//
// Solidity: function cidSet(bytes ) view returns(bool)
func (_DealRewarder *DealRewarderCallerSession) CidSet(arg0 []byte) (bool, error) {
	return _DealRewarder.Contract.CidSet(&_DealRewarder.CallOpts, arg0)
}

// CidSizes is a free data retrieval call binding the contract method 0x8f8a6d8f.
//
// Solidity: function cidSizes(bytes ) view returns(uint256)
func (_DealRewarder *DealRewarderCaller) CidSizes(opts *bind.CallOpts, arg0 []byte) (*big.Int, error) {
	var out []interface{}
	err := _DealRewarder.contract.Call(opts, &out, "cidSizes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CidSizes is a free data retrieval call binding the contract method 0x8f8a6d8f.
//
// Solidity: function cidSizes(bytes ) view returns(uint256)
func (_DealRewarder *DealRewarderSession) CidSizes(arg0 []byte) (*big.Int, error) {
	return _DealRewarder.Contract.CidSizes(&_DealRewarder.CallOpts, arg0)
}

// CidSizes is a free data retrieval call binding the contract method 0x8f8a6d8f.
//
// Solidity: function cidSizes(bytes ) view returns(uint256)
func (_DealRewarder *DealRewarderCallerSession) CidSizes(arg0 []byte) (*big.Int, error) {
	return _DealRewarder.Contract.CidSizes(&_DealRewarder.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DealRewarder *DealRewarderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DealRewarder.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DealRewarder *DealRewarderSession) Owner() (common.Address, error) {
	return _DealRewarder.Contract.Owner(&_DealRewarder.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DealRewarder *DealRewarderCallerSession) Owner() (common.Address, error) {
	return _DealRewarder.Contract.Owner(&_DealRewarder.CallOpts)
}

// AddCID is a paid mutator transaction binding the contract method 0xd4a0cd0a.
//
// Solidity: function addCID(bytes cidraw, uint256 size) returns()
func (_DealRewarder *DealRewarderTransactor) AddCID(opts *bind.TransactOpts, cidraw []byte, size *big.Int) (*types.Transaction, error) {
	return _DealRewarder.contract.Transact(opts, "addCID", cidraw, size)
}

// AddCID is a paid mutator transaction binding the contract method 0xd4a0cd0a.
//
// Solidity: function addCID(bytes cidraw, uint256 size) returns()
func (_DealRewarder *DealRewarderSession) AddCID(cidraw []byte, size *big.Int) (*types.Transaction, error) {
	return _DealRewarder.Contract.AddCID(&_DealRewarder.TransactOpts, cidraw, size)
}

// AddCID is a paid mutator transaction binding the contract method 0xd4a0cd0a.
//
// Solidity: function addCID(bytes cidraw, uint256 size) returns()
func (_DealRewarder *DealRewarderTransactorSession) AddCID(cidraw []byte, size *big.Int) (*types.Transaction, error) {
	return _DealRewarder.Contract.AddCID(&_DealRewarder.TransactOpts, cidraw, size)
}

// AuthorizeData is a paid mutator transaction binding the contract method 0xb97c601a.
//
// Solidity: function authorizeData(bytes cidraw, uint64 provider, uint256 size) returns()
func (_DealRewarder *DealRewarderTransactor) AuthorizeData(opts *bind.TransactOpts, cidraw []byte, provider uint64, size *big.Int) (*types.Transaction, error) {
	return _DealRewarder.contract.Transact(opts, "authorizeData", cidraw, provider, size)
}

// AuthorizeData is a paid mutator transaction binding the contract method 0xb97c601a.
//
// Solidity: function authorizeData(bytes cidraw, uint64 provider, uint256 size) returns()
func (_DealRewarder *DealRewarderSession) AuthorizeData(cidraw []byte, provider uint64, size *big.Int) (*types.Transaction, error) {
	return _DealRewarder.Contract.AuthorizeData(&_DealRewarder.TransactOpts, cidraw, provider, size)
}

// AuthorizeData is a paid mutator transaction binding the contract method 0xb97c601a.
//
// Solidity: function authorizeData(bytes cidraw, uint64 provider, uint256 size) returns()
func (_DealRewarder *DealRewarderTransactorSession) AuthorizeData(cidraw []byte, provider uint64, size *big.Int) (*types.Transaction, error) {
	return _DealRewarder.Contract.AuthorizeData(&_DealRewarder.TransactOpts, cidraw, provider, size)
}

// CallActorId is a paid mutator transaction binding the contract method 0x6aba510b.
//
// Solidity: function call_actor_id(uint64 method, uint256 value, uint64 flags, uint64 codec, bytes params, uint64 id) returns(bool, int256, uint64, bytes)
func (_DealRewarder *DealRewarderTransactor) CallActorId(opts *bind.TransactOpts, method uint64, value *big.Int, flags uint64, codec uint64, params []byte, id uint64) (*types.Transaction, error) {
	return _DealRewarder.contract.Transact(opts, "call_actor_id", method, value, flags, codec, params, id)
}

// CallActorId is a paid mutator transaction binding the contract method 0x6aba510b.
//
// Solidity: function call_actor_id(uint64 method, uint256 value, uint64 flags, uint64 codec, bytes params, uint64 id) returns(bool, int256, uint64, bytes)
func (_DealRewarder *DealRewarderSession) CallActorId(method uint64, value *big.Int, flags uint64, codec uint64, params []byte, id uint64) (*types.Transaction, error) {
	return _DealRewarder.Contract.CallActorId(&_DealRewarder.TransactOpts, method, value, flags, codec, params, id)
}

// CallActorId is a paid mutator transaction binding the contract method 0x6aba510b.
//
// Solidity: function call_actor_id(uint64 method, uint256 value, uint64 flags, uint64 codec, bytes params, uint64 id) returns(bool, int256, uint64, bytes)
func (_DealRewarder *DealRewarderTransactorSession) CallActorId(method uint64, value *big.Int, flags uint64, codec uint64, params []byte, id uint64) (*types.Transaction, error) {
	return _DealRewarder.Contract.CallActorId(&_DealRewarder.TransactOpts, method, value, flags, codec, params, id)
}

// ClaimBounty is a paid mutator transaction binding the contract method 0xe1fb0bba.
//
// Solidity: function claim_bounty(uint64 deal_id) returns()
func (_DealRewarder *DealRewarderTransactor) ClaimBounty(opts *bind.TransactOpts, deal_id uint64) (*types.Transaction, error) {
	return _DealRewarder.contract.Transact(opts, "claim_bounty", deal_id)
}

// ClaimBounty is a paid mutator transaction binding the contract method 0xe1fb0bba.
//
// Solidity: function claim_bounty(uint64 deal_id) returns()
func (_DealRewarder *DealRewarderSession) ClaimBounty(deal_id uint64) (*types.Transaction, error) {
	return _DealRewarder.Contract.ClaimBounty(&_DealRewarder.TransactOpts, deal_id)
}

// ClaimBounty is a paid mutator transaction binding the contract method 0xe1fb0bba.
//
// Solidity: function claim_bounty(uint64 deal_id) returns()
func (_DealRewarder *DealRewarderTransactorSession) ClaimBounty(deal_id uint64) (*types.Transaction, error) {
	return _DealRewarder.Contract.ClaimBounty(&_DealRewarder.TransactOpts, deal_id)
}

// Fund is a paid mutator transaction binding the contract method 0x4f9c09cc.
//
// Solidity: function fund(uint64 unused) payable returns()
func (_DealRewarder *DealRewarderTransactor) Fund(opts *bind.TransactOpts, unused uint64) (*types.Transaction, error) {
	return _DealRewarder.contract.Transact(opts, "fund", unused)
}

// Fund is a paid mutator transaction binding the contract method 0x4f9c09cc.
//
// Solidity: function fund(uint64 unused) payable returns()
func (_DealRewarder *DealRewarderSession) Fund(unused uint64) (*types.Transaction, error) {
	return _DealRewarder.Contract.Fund(&_DealRewarder.TransactOpts, unused)
}

// Fund is a paid mutator transaction binding the contract method 0x4f9c09cc.
//
// Solidity: function fund(uint64 unused) payable returns()
func (_DealRewarder *DealRewarderTransactorSession) Fund(unused uint64) (*types.Transaction, error) {
	return _DealRewarder.Contract.Fund(&_DealRewarder.TransactOpts, unused)
}
