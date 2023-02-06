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

// LibraryMetaData contains all meta data concerning the Library contract.
var LibraryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_provider\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updatedCount\",\"type\":\"uint256\"}],\"name\":\"UnsynchronizedCount\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"Count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pieceCid\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"dealId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"add\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"contributors\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_num\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAll\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getAllByAccount\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEpochs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"}],\"name\":\"getNumberOfPresentationsInEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"numPresentationsByEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"presentations\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"presentationsByAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endEpoch\",\"type\":\"uint256\"}],\"name\":\"range\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_num\",\"type\":\"uint256\"}],\"name\":\"remove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_presentation\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_contributors\",\"type\":\"address[]\"}],\"name\":\"setContributors\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"setCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_presentation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_pieceCID\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"dealId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LibraryABI is the input ABI used to generate the binding from.
// Deprecated: Use LibraryMetaData.ABI instead.
var LibraryABI = LibraryMetaData.ABI

// Library is an auto generated Go binding around an Ethereum contract.
type Library struct {
	LibraryCaller     // Read-only binding to the contract
	LibraryTransactor // Write-only binding to the contract
	LibraryFilterer   // Log filterer for contract events
}

// LibraryCaller is an auto generated read-only Go binding around an Ethereum contract.
type LibraryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibraryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LibraryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibraryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LibraryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibrarySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LibrarySession struct {
	Contract     *Library          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LibraryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LibraryCallerSession struct {
	Contract *LibraryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// LibraryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LibraryTransactorSession struct {
	Contract     *LibraryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LibraryRaw is an auto generated low-level Go binding around an Ethereum contract.
type LibraryRaw struct {
	Contract *Library // Generic contract binding to access the raw methods on
}

// LibraryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LibraryCallerRaw struct {
	Contract *LibraryCaller // Generic read-only contract binding to access the raw methods on
}

// LibraryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LibraryTransactorRaw struct {
	Contract *LibraryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLibrary creates a new instance of Library, bound to a specific deployed contract.
func NewLibrary(address common.Address, backend bind.ContractBackend) (*Library, error) {
	contract, err := bindLibrary(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Library{LibraryCaller: LibraryCaller{contract: contract}, LibraryTransactor: LibraryTransactor{contract: contract}, LibraryFilterer: LibraryFilterer{contract: contract}}, nil
}

// NewLibraryCaller creates a new read-only instance of Library, bound to a specific deployed contract.
func NewLibraryCaller(address common.Address, caller bind.ContractCaller) (*LibraryCaller, error) {
	contract, err := bindLibrary(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LibraryCaller{contract: contract}, nil
}

// NewLibraryTransactor creates a new write-only instance of Library, bound to a specific deployed contract.
func NewLibraryTransactor(address common.Address, transactor bind.ContractTransactor) (*LibraryTransactor, error) {
	contract, err := bindLibrary(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LibraryTransactor{contract: contract}, nil
}

// NewLibraryFilterer creates a new log filterer instance of Library, bound to a specific deployed contract.
func NewLibraryFilterer(address common.Address, filterer bind.ContractFilterer) (*LibraryFilterer, error) {
	contract, err := bindLibrary(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LibraryFilterer{contract: contract}, nil
}

// bindLibrary binds a generic wrapper to an already deployed contract.
func bindLibrary(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LibraryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Library *LibraryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Library.Contract.LibraryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Library *LibraryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Library.Contract.LibraryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Library *LibraryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Library.Contract.LibraryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Library *LibraryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Library.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Library *LibraryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Library.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Library *LibraryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Library.Contract.contract.Transact(opts, method, params...)
}

// Count is a free data retrieval call binding the contract method 0x93a8333e.
//
// Solidity: function Count() view returns(uint256)
func (_Library *LibraryCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Library.contract.Call(opts, &out, "Count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x93a8333e.
//
// Solidity: function Count() view returns(uint256)
func (_Library *LibrarySession) Count() (*big.Int, error) {
	return _Library.Contract.Count(&_Library.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x93a8333e.
//
// Solidity: function Count() view returns(uint256)
func (_Library *LibraryCallerSession) Count() (*big.Int, error) {
	return _Library.Contract.Count(&_Library.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Library *LibraryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Library.contract.Call(opts, &out, "Owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Library *LibrarySession) Owner() (common.Address, error) {
	return _Library.Contract.Owner(&_Library.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Library *LibraryCallerSession) Owner() (common.Address, error) {
	return _Library.Contract.Owner(&_Library.CallOpts)
}

// Contributors is a free data retrieval call binding the contract method 0x73d659b3.
//
// Solidity: function contributors(address , uint256 ) view returns(address)
func (_Library *LibraryCaller) Contributors(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Library.contract.Call(opts, &out, "contributors", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Contributors is a free data retrieval call binding the contract method 0x73d659b3.
//
// Solidity: function contributors(address , uint256 ) view returns(address)
func (_Library *LibrarySession) Contributors(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Library.Contract.Contributors(&_Library.CallOpts, arg0, arg1)
}

// Contributors is a free data retrieval call binding the contract method 0x73d659b3.
//
// Solidity: function contributors(address , uint256 ) view returns(address)
func (_Library *LibraryCallerSession) Contributors(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Library.Contract.Contributors(&_Library.CallOpts, arg0, arg1)
}

// Get is a free data retrieval call binding the contract method 0x669e48aa.
//
// Solidity: function get(uint256 _epoch, uint256 _num) view returns(address)
func (_Library *LibraryCaller) Get(opts *bind.CallOpts, _epoch *big.Int, _num *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Library.contract.Call(opts, &out, "get", _epoch, _num)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x669e48aa.
//
// Solidity: function get(uint256 _epoch, uint256 _num) view returns(address)
func (_Library *LibrarySession) Get(_epoch *big.Int, _num *big.Int) (common.Address, error) {
	return _Library.Contract.Get(&_Library.CallOpts, _epoch, _num)
}

// Get is a free data retrieval call binding the contract method 0x669e48aa.
//
// Solidity: function get(uint256 _epoch, uint256 _num) view returns(address)
func (_Library *LibraryCallerSession) Get(_epoch *big.Int, _num *big.Int) (common.Address, error) {
	return _Library.Contract.Get(&_Library.CallOpts, _epoch, _num)
}

// GetEpochs is a free data retrieval call binding the contract method 0xd6b1aa5d.
//
// Solidity: function getEpochs() view returns(uint256[])
func (_Library *LibraryCaller) GetEpochs(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Library.contract.Call(opts, &out, "getEpochs")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetEpochs is a free data retrieval call binding the contract method 0xd6b1aa5d.
//
// Solidity: function getEpochs() view returns(uint256[])
func (_Library *LibrarySession) GetEpochs() ([]*big.Int, error) {
	return _Library.Contract.GetEpochs(&_Library.CallOpts)
}

// GetEpochs is a free data retrieval call binding the contract method 0xd6b1aa5d.
//
// Solidity: function getEpochs() view returns(uint256[])
func (_Library *LibraryCallerSession) GetEpochs() ([]*big.Int, error) {
	return _Library.Contract.GetEpochs(&_Library.CallOpts)
}

// GetNumberOfPresentationsInEpoch is a free data retrieval call binding the contract method 0x8f7041f3.
//
// Solidity: function getNumberOfPresentationsInEpoch(uint256 _epoch) view returns(uint256)
func (_Library *LibraryCaller) GetNumberOfPresentationsInEpoch(opts *bind.CallOpts, _epoch *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Library.contract.Call(opts, &out, "getNumberOfPresentationsInEpoch", _epoch)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumberOfPresentationsInEpoch is a free data retrieval call binding the contract method 0x8f7041f3.
//
// Solidity: function getNumberOfPresentationsInEpoch(uint256 _epoch) view returns(uint256)
func (_Library *LibrarySession) GetNumberOfPresentationsInEpoch(_epoch *big.Int) (*big.Int, error) {
	return _Library.Contract.GetNumberOfPresentationsInEpoch(&_Library.CallOpts, _epoch)
}

// GetNumberOfPresentationsInEpoch is a free data retrieval call binding the contract method 0x8f7041f3.
//
// Solidity: function getNumberOfPresentationsInEpoch(uint256 _epoch) view returns(uint256)
func (_Library *LibraryCallerSession) GetNumberOfPresentationsInEpoch(_epoch *big.Int) (*big.Int, error) {
	return _Library.Contract.GetNumberOfPresentationsInEpoch(&_Library.CallOpts, _epoch)
}

// NumPresentationsByEpoch is a free data retrieval call binding the contract method 0x1d196bc2.
//
// Solidity: function numPresentationsByEpoch(uint256 ) view returns(uint256)
func (_Library *LibraryCaller) NumPresentationsByEpoch(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Library.contract.Call(opts, &out, "numPresentationsByEpoch", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumPresentationsByEpoch is a free data retrieval call binding the contract method 0x1d196bc2.
//
// Solidity: function numPresentationsByEpoch(uint256 ) view returns(uint256)
func (_Library *LibrarySession) NumPresentationsByEpoch(arg0 *big.Int) (*big.Int, error) {
	return _Library.Contract.NumPresentationsByEpoch(&_Library.CallOpts, arg0)
}

// NumPresentationsByEpoch is a free data retrieval call binding the contract method 0x1d196bc2.
//
// Solidity: function numPresentationsByEpoch(uint256 ) view returns(uint256)
func (_Library *LibraryCallerSession) NumPresentationsByEpoch(arg0 *big.Int) (*big.Int, error) {
	return _Library.Contract.NumPresentationsByEpoch(&_Library.CallOpts, arg0)
}

// Presentations is a free data retrieval call binding the contract method 0x2d193300.
//
// Solidity: function presentations(uint256 , uint256 ) view returns(address)
func (_Library *LibraryCaller) Presentations(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Library.contract.Call(opts, &out, "presentations", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Presentations is a free data retrieval call binding the contract method 0x2d193300.
//
// Solidity: function presentations(uint256 , uint256 ) view returns(address)
func (_Library *LibrarySession) Presentations(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Library.Contract.Presentations(&_Library.CallOpts, arg0, arg1)
}

// Presentations is a free data retrieval call binding the contract method 0x2d193300.
//
// Solidity: function presentations(uint256 , uint256 ) view returns(address)
func (_Library *LibraryCallerSession) Presentations(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Library.Contract.Presentations(&_Library.CallOpts, arg0, arg1)
}

// PresentationsByAccount is a free data retrieval call binding the contract method 0x2ddf7c0b.
//
// Solidity: function presentationsByAccount(address , uint256 ) view returns(address)
func (_Library *LibraryCaller) PresentationsByAccount(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Library.contract.Call(opts, &out, "presentationsByAccount", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PresentationsByAccount is a free data retrieval call binding the contract method 0x2ddf7c0b.
//
// Solidity: function presentationsByAccount(address , uint256 ) view returns(address)
func (_Library *LibrarySession) PresentationsByAccount(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Library.Contract.PresentationsByAccount(&_Library.CallOpts, arg0, arg1)
}

// PresentationsByAccount is a free data retrieval call binding the contract method 0x2ddf7c0b.
//
// Solidity: function presentationsByAccount(address , uint256 ) view returns(address)
func (_Library *LibraryCallerSession) PresentationsByAccount(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Library.Contract.PresentationsByAccount(&_Library.CallOpts, arg0, arg1)
}

// Add is a paid mutator transaction binding the contract method 0x185425d6.
//
// Solidity: function add(bytes _pieceCid, uint64 dealId, uint64 size, string metadata) returns(address)
func (_Library *LibraryTransactor) Add(opts *bind.TransactOpts, _pieceCid []byte, dealId uint64, size uint64, metadata string) (*types.Transaction, error) {
	return _Library.contract.Transact(opts, "add", _pieceCid, dealId, size, metadata)
}

// Add is a paid mutator transaction binding the contract method 0x185425d6.
//
// Solidity: function add(bytes _pieceCid, uint64 dealId, uint64 size, string metadata) returns(address)
func (_Library *LibrarySession) Add(_pieceCid []byte, dealId uint64, size uint64, metadata string) (*types.Transaction, error) {
	return _Library.Contract.Add(&_Library.TransactOpts, _pieceCid, dealId, size, metadata)
}

// Add is a paid mutator transaction binding the contract method 0x185425d6.
//
// Solidity: function add(bytes _pieceCid, uint64 dealId, uint64 size, string metadata) returns(address)
func (_Library *LibraryTransactorSession) Add(_pieceCid []byte, dealId uint64, size uint64, metadata string) (*types.Transaction, error) {
	return _Library.Contract.Add(&_Library.TransactOpts, _pieceCid, dealId, size, metadata)
}

// GetAll is a paid mutator transaction binding the contract method 0x53ed5143.
//
// Solidity: function getAll() returns(address[])
func (_Library *LibraryTransactor) GetAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Library.contract.Transact(opts, "getAll")
}

// GetAll is a paid mutator transaction binding the contract method 0x53ed5143.
//
// Solidity: function getAll() returns(address[])
func (_Library *LibrarySession) GetAll() (*types.Transaction, error) {
	return _Library.Contract.GetAll(&_Library.TransactOpts)
}

// GetAll is a paid mutator transaction binding the contract method 0x53ed5143.
//
// Solidity: function getAll() returns(address[])
func (_Library *LibraryTransactorSession) GetAll() (*types.Transaction, error) {
	return _Library.Contract.GetAll(&_Library.TransactOpts)
}

// GetAllByAccount is a paid mutator transaction binding the contract method 0x5415c3e1.
//
// Solidity: function getAllByAccount(address _account) returns(address[])
func (_Library *LibraryTransactor) GetAllByAccount(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Library.contract.Transact(opts, "getAllByAccount", _account)
}

// GetAllByAccount is a paid mutator transaction binding the contract method 0x5415c3e1.
//
// Solidity: function getAllByAccount(address _account) returns(address[])
func (_Library *LibrarySession) GetAllByAccount(_account common.Address) (*types.Transaction, error) {
	return _Library.Contract.GetAllByAccount(&_Library.TransactOpts, _account)
}

// GetAllByAccount is a paid mutator transaction binding the contract method 0x5415c3e1.
//
// Solidity: function getAllByAccount(address _account) returns(address[])
func (_Library *LibraryTransactorSession) GetAllByAccount(_account common.Address) (*types.Transaction, error) {
	return _Library.Contract.GetAllByAccount(&_Library.TransactOpts, _account)
}

// Range is a paid mutator transaction binding the contract method 0x55680991.
//
// Solidity: function range(uint256 startEpoch, uint256 endEpoch) returns(address[])
func (_Library *LibraryTransactor) Range(opts *bind.TransactOpts, startEpoch *big.Int, endEpoch *big.Int) (*types.Transaction, error) {
	return _Library.contract.Transact(opts, "range", startEpoch, endEpoch)
}

// Range is a paid mutator transaction binding the contract method 0x55680991.
//
// Solidity: function range(uint256 startEpoch, uint256 endEpoch) returns(address[])
func (_Library *LibrarySession) Range(startEpoch *big.Int, endEpoch *big.Int) (*types.Transaction, error) {
	return _Library.Contract.Range(&_Library.TransactOpts, startEpoch, endEpoch)
}

// Range is a paid mutator transaction binding the contract method 0x55680991.
//
// Solidity: function range(uint256 startEpoch, uint256 endEpoch) returns(address[])
func (_Library *LibraryTransactorSession) Range(startEpoch *big.Int, endEpoch *big.Int) (*types.Transaction, error) {
	return _Library.Contract.Range(&_Library.TransactOpts, startEpoch, endEpoch)
}

// Remove is a paid mutator transaction binding the contract method 0x6526db7a.
//
// Solidity: function remove(uint256 _epoch, uint256 _num) returns()
func (_Library *LibraryTransactor) Remove(opts *bind.TransactOpts, _epoch *big.Int, _num *big.Int) (*types.Transaction, error) {
	return _Library.contract.Transact(opts, "remove", _epoch, _num)
}

// Remove is a paid mutator transaction binding the contract method 0x6526db7a.
//
// Solidity: function remove(uint256 _epoch, uint256 _num) returns()
func (_Library *LibrarySession) Remove(_epoch *big.Int, _num *big.Int) (*types.Transaction, error) {
	return _Library.Contract.Remove(&_Library.TransactOpts, _epoch, _num)
}

// Remove is a paid mutator transaction binding the contract method 0x6526db7a.
//
// Solidity: function remove(uint256 _epoch, uint256 _num) returns()
func (_Library *LibraryTransactorSession) Remove(_epoch *big.Int, _num *big.Int) (*types.Transaction, error) {
	return _Library.Contract.Remove(&_Library.TransactOpts, _epoch, _num)
}

// SetContributors is a paid mutator transaction binding the contract method 0x79afebca.
//
// Solidity: function setContributors(address _presentation, address[] _contributors) returns()
func (_Library *LibraryTransactor) SetContributors(opts *bind.TransactOpts, _presentation common.Address, _contributors []common.Address) (*types.Transaction, error) {
	return _Library.contract.Transact(opts, "setContributors", _presentation, _contributors)
}

// SetContributors is a paid mutator transaction binding the contract method 0x79afebca.
//
// Solidity: function setContributors(address _presentation, address[] _contributors) returns()
func (_Library *LibrarySession) SetContributors(_presentation common.Address, _contributors []common.Address) (*types.Transaction, error) {
	return _Library.Contract.SetContributors(&_Library.TransactOpts, _presentation, _contributors)
}

// SetContributors is a paid mutator transaction binding the contract method 0x79afebca.
//
// Solidity: function setContributors(address _presentation, address[] _contributors) returns()
func (_Library *LibraryTransactorSession) SetContributors(_presentation common.Address, _contributors []common.Address) (*types.Transaction, error) {
	return _Library.Contract.SetContributors(&_Library.TransactOpts, _presentation, _contributors)
}

// SetCount is a paid mutator transaction binding the contract method 0xd14e62b8.
//
// Solidity: function setCount(uint256 _count) returns()
func (_Library *LibraryTransactor) SetCount(opts *bind.TransactOpts, _count *big.Int) (*types.Transaction, error) {
	return _Library.contract.Transact(opts, "setCount", _count)
}

// SetCount is a paid mutator transaction binding the contract method 0xd14e62b8.
//
// Solidity: function setCount(uint256 _count) returns()
func (_Library *LibrarySession) SetCount(_count *big.Int) (*types.Transaction, error) {
	return _Library.Contract.SetCount(&_Library.TransactOpts, _count)
}

// SetCount is a paid mutator transaction binding the contract method 0xd14e62b8.
//
// Solidity: function setCount(uint256 _count) returns()
func (_Library *LibraryTransactorSession) SetCount(_count *big.Int) (*types.Transaction, error) {
	return _Library.Contract.SetCount(&_Library.TransactOpts, _count)
}

// Update is a paid mutator transaction binding the contract method 0xb4e2008b.
//
// Solidity: function update(address _presentation, bytes _pieceCID, uint64 dealId, uint64 size) returns()
func (_Library *LibraryTransactor) Update(opts *bind.TransactOpts, _presentation common.Address, _pieceCID []byte, dealId uint64, size uint64) (*types.Transaction, error) {
	return _Library.contract.Transact(opts, "update", _presentation, _pieceCID, dealId, size)
}

// Update is a paid mutator transaction binding the contract method 0xb4e2008b.
//
// Solidity: function update(address _presentation, bytes _pieceCID, uint64 dealId, uint64 size) returns()
func (_Library *LibrarySession) Update(_presentation common.Address, _pieceCID []byte, dealId uint64, size uint64) (*types.Transaction, error) {
	return _Library.Contract.Update(&_Library.TransactOpts, _presentation, _pieceCID, dealId, size)
}

// Update is a paid mutator transaction binding the contract method 0xb4e2008b.
//
// Solidity: function update(address _presentation, bytes _pieceCID, uint64 dealId, uint64 size) returns()
func (_Library *LibraryTransactorSession) Update(_presentation common.Address, _pieceCID []byte, dealId uint64, size uint64) (*types.Transaction, error) {
	return _Library.Contract.Update(&_Library.TransactOpts, _presentation, _pieceCID, dealId, size)
}

// LibraryUnsynchronizedCountIterator is returned from FilterUnsynchronizedCount and is used to iterate over the raw logs and unpacked data for UnsynchronizedCount events raised by the Library contract.
type LibraryUnsynchronizedCountIterator struct {
	Event *LibraryUnsynchronizedCount // Event containing the contract specifics and raw log

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
func (it *LibraryUnsynchronizedCountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LibraryUnsynchronizedCount)
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
		it.Event = new(LibraryUnsynchronizedCount)
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
func (it *LibraryUnsynchronizedCountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LibraryUnsynchronizedCountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LibraryUnsynchronizedCount represents a UnsynchronizedCount event raised by the Library contract.
type LibraryUnsynchronizedCount struct {
	UpdatedCount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUnsynchronizedCount is a free log retrieval operation binding the contract event 0x05aa2541666252702c740092e7016de5077e6bd8946a479951ad76b955c96528.
//
// Solidity: event UnsynchronizedCount(uint256 updatedCount)
func (_Library *LibraryFilterer) FilterUnsynchronizedCount(opts *bind.FilterOpts) (*LibraryUnsynchronizedCountIterator, error) {

	logs, sub, err := _Library.contract.FilterLogs(opts, "UnsynchronizedCount")
	if err != nil {
		return nil, err
	}
	return &LibraryUnsynchronizedCountIterator{contract: _Library.contract, event: "UnsynchronizedCount", logs: logs, sub: sub}, nil
}

// WatchUnsynchronizedCount is a free log subscription operation binding the contract event 0x05aa2541666252702c740092e7016de5077e6bd8946a479951ad76b955c96528.
//
// Solidity: event UnsynchronizedCount(uint256 updatedCount)
func (_Library *LibraryFilterer) WatchUnsynchronizedCount(opts *bind.WatchOpts, sink chan<- *LibraryUnsynchronizedCount) (event.Subscription, error) {

	logs, sub, err := _Library.contract.WatchLogs(opts, "UnsynchronizedCount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LibraryUnsynchronizedCount)
				if err := _Library.contract.UnpackLog(event, "UnsynchronizedCount", log); err != nil {
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

// ParseUnsynchronizedCount is a log parse operation binding the contract event 0x05aa2541666252702c740092e7016de5077e6bd8946a479951ad76b955c96528.
//
// Solidity: event UnsynchronizedCount(uint256 updatedCount)
func (_Library *LibraryFilterer) ParseUnsynchronizedCount(log types.Log) (*LibraryUnsynchronizedCount, error) {
	event := new(LibraryUnsynchronizedCount)
	if err := _Library.contract.UnpackLog(event, "UnsynchronizedCount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
