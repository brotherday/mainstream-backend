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

// ProfileMetaData contains all meta data concerning the Profile contract.
var ProfileMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_role\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Host\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Public\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Speaker\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"emails\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"names\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pictures\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"roles\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_email\",\"type\":\"bytes32\"}],\"name\":\"setEmail\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_picture\",\"type\":\"bytes32\"}],\"name\":\"setPicture\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ProfileABI is the input ABI used to generate the binding from.
// Deprecated: Use ProfileMetaData.ABI instead.
var ProfileABI = ProfileMetaData.ABI

// Profile is an auto generated Go binding around an Ethereum contract.
type Profile struct {
	ProfileCaller     // Read-only binding to the contract
	ProfileTransactor // Write-only binding to the contract
	ProfileFilterer   // Log filterer for contract events
}

// ProfileCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProfileCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProfileTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProfileTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProfileFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProfileFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProfileSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProfileSession struct {
	Contract     *Profile          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProfileCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProfileCallerSession struct {
	Contract *ProfileCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ProfileTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProfileTransactorSession struct {
	Contract     *ProfileTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ProfileRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProfileRaw struct {
	Contract *Profile // Generic contract binding to access the raw methods on
}

// ProfileCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProfileCallerRaw struct {
	Contract *ProfileCaller // Generic read-only contract binding to access the raw methods on
}

// ProfileTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProfileTransactorRaw struct {
	Contract *ProfileTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProfile creates a new instance of Profile, bound to a specific deployed contract.
func NewProfile(address common.Address, backend bind.ContractBackend) (*Profile, error) {
	contract, err := bindProfile(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Profile{ProfileCaller: ProfileCaller{contract: contract}, ProfileTransactor: ProfileTransactor{contract: contract}, ProfileFilterer: ProfileFilterer{contract: contract}}, nil
}

// NewProfileCaller creates a new read-only instance of Profile, bound to a specific deployed contract.
func NewProfileCaller(address common.Address, caller bind.ContractCaller) (*ProfileCaller, error) {
	contract, err := bindProfile(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProfileCaller{contract: contract}, nil
}

// NewProfileTransactor creates a new write-only instance of Profile, bound to a specific deployed contract.
func NewProfileTransactor(address common.Address, transactor bind.ContractTransactor) (*ProfileTransactor, error) {
	contract, err := bindProfile(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProfileTransactor{contract: contract}, nil
}

// NewProfileFilterer creates a new log filterer instance of Profile, bound to a specific deployed contract.
func NewProfileFilterer(address common.Address, filterer bind.ContractFilterer) (*ProfileFilterer, error) {
	contract, err := bindProfile(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProfileFilterer{contract: contract}, nil
}

// bindProfile binds a generic wrapper to an already deployed contract.
func bindProfile(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProfileABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Profile *ProfileRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Profile.Contract.ProfileCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Profile *ProfileRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Profile.Contract.ProfileTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Profile *ProfileRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Profile.Contract.ProfileTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Profile *ProfileCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Profile.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Profile *ProfileTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Profile.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Profile *ProfileTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Profile.Contract.contract.Transact(opts, method, params...)
}

// Host is a free data retrieval call binding the contract method 0xdc392c32.
//
// Solidity: function Host() view returns(bytes32)
func (_Profile *ProfileCaller) Host(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "Host")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Host is a free data retrieval call binding the contract method 0xdc392c32.
//
// Solidity: function Host() view returns(bytes32)
func (_Profile *ProfileSession) Host() ([32]byte, error) {
	return _Profile.Contract.Host(&_Profile.CallOpts)
}

// Host is a free data retrieval call binding the contract method 0xdc392c32.
//
// Solidity: function Host() view returns(bytes32)
func (_Profile *ProfileCallerSession) Host() ([32]byte, error) {
	return _Profile.Contract.Host(&_Profile.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Profile *ProfileCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "Owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Profile *ProfileSession) Owner() (common.Address, error) {
	return _Profile.Contract.Owner(&_Profile.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Profile *ProfileCallerSession) Owner() (common.Address, error) {
	return _Profile.Contract.Owner(&_Profile.CallOpts)
}

// Public is a free data retrieval call binding the contract method 0x5097e51f.
//
// Solidity: function Public() view returns(bytes32)
func (_Profile *ProfileCaller) Public(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "Public")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Public is a free data retrieval call binding the contract method 0x5097e51f.
//
// Solidity: function Public() view returns(bytes32)
func (_Profile *ProfileSession) Public() ([32]byte, error) {
	return _Profile.Contract.Public(&_Profile.CallOpts)
}

// Public is a free data retrieval call binding the contract method 0x5097e51f.
//
// Solidity: function Public() view returns(bytes32)
func (_Profile *ProfileCallerSession) Public() ([32]byte, error) {
	return _Profile.Contract.Public(&_Profile.CallOpts)
}

// Speaker is a free data retrieval call binding the contract method 0xc85759c8.
//
// Solidity: function Speaker() view returns(bytes32)
func (_Profile *ProfileCaller) Speaker(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "Speaker")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Speaker is a free data retrieval call binding the contract method 0xc85759c8.
//
// Solidity: function Speaker() view returns(bytes32)
func (_Profile *ProfileSession) Speaker() ([32]byte, error) {
	return _Profile.Contract.Speaker(&_Profile.CallOpts)
}

// Speaker is a free data retrieval call binding the contract method 0xc85759c8.
//
// Solidity: function Speaker() view returns(bytes32)
func (_Profile *ProfileCallerSession) Speaker() ([32]byte, error) {
	return _Profile.Contract.Speaker(&_Profile.CallOpts)
}

// Emails is a free data retrieval call binding the contract method 0xcb2da07c.
//
// Solidity: function emails(address ) view returns(bytes32)
func (_Profile *ProfileCaller) Emails(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "emails", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Emails is a free data retrieval call binding the contract method 0xcb2da07c.
//
// Solidity: function emails(address ) view returns(bytes32)
func (_Profile *ProfileSession) Emails(arg0 common.Address) ([32]byte, error) {
	return _Profile.Contract.Emails(&_Profile.CallOpts, arg0)
}

// Emails is a free data retrieval call binding the contract method 0xcb2da07c.
//
// Solidity: function emails(address ) view returns(bytes32)
func (_Profile *ProfileCallerSession) Emails(arg0 common.Address) ([32]byte, error) {
	return _Profile.Contract.Emails(&_Profile.CallOpts, arg0)
}

// Names is a free data retrieval call binding the contract method 0x5cf3d346.
//
// Solidity: function names(address ) view returns(string)
func (_Profile *ProfileCaller) Names(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "names", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Names is a free data retrieval call binding the contract method 0x5cf3d346.
//
// Solidity: function names(address ) view returns(string)
func (_Profile *ProfileSession) Names(arg0 common.Address) (string, error) {
	return _Profile.Contract.Names(&_Profile.CallOpts, arg0)
}

// Names is a free data retrieval call binding the contract method 0x5cf3d346.
//
// Solidity: function names(address ) view returns(string)
func (_Profile *ProfileCallerSession) Names(arg0 common.Address) (string, error) {
	return _Profile.Contract.Names(&_Profile.CallOpts, arg0)
}

// Pictures is a free data retrieval call binding the contract method 0xc6fd8be3.
//
// Solidity: function pictures(address ) view returns(bytes32)
func (_Profile *ProfileCaller) Pictures(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "pictures", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Pictures is a free data retrieval call binding the contract method 0xc6fd8be3.
//
// Solidity: function pictures(address ) view returns(bytes32)
func (_Profile *ProfileSession) Pictures(arg0 common.Address) ([32]byte, error) {
	return _Profile.Contract.Pictures(&_Profile.CallOpts, arg0)
}

// Pictures is a free data retrieval call binding the contract method 0xc6fd8be3.
//
// Solidity: function pictures(address ) view returns(bytes32)
func (_Profile *ProfileCallerSession) Pictures(arg0 common.Address) ([32]byte, error) {
	return _Profile.Contract.Pictures(&_Profile.CallOpts, arg0)
}

// Roles is a free data retrieval call binding the contract method 0x99374642.
//
// Solidity: function roles(address ) view returns(bytes32)
func (_Profile *ProfileCaller) Roles(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "roles", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Roles is a free data retrieval call binding the contract method 0x99374642.
//
// Solidity: function roles(address ) view returns(bytes32)
func (_Profile *ProfileSession) Roles(arg0 common.Address) ([32]byte, error) {
	return _Profile.Contract.Roles(&_Profile.CallOpts, arg0)
}

// Roles is a free data retrieval call binding the contract method 0x99374642.
//
// Solidity: function roles(address ) view returns(bytes32)
func (_Profile *ProfileCallerSession) Roles(arg0 common.Address) ([32]byte, error) {
	return _Profile.Contract.Roles(&_Profile.CallOpts, arg0)
}

// SetEmail is a paid mutator transaction binding the contract method 0xc8dd612e.
//
// Solidity: function setEmail(bytes32 _email) returns()
func (_Profile *ProfileTransactor) SetEmail(opts *bind.TransactOpts, _email [32]byte) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "setEmail", _email)
}

// SetEmail is a paid mutator transaction binding the contract method 0xc8dd612e.
//
// Solidity: function setEmail(bytes32 _email) returns()
func (_Profile *ProfileSession) SetEmail(_email [32]byte) (*types.Transaction, error) {
	return _Profile.Contract.SetEmail(&_Profile.TransactOpts, _email)
}

// SetEmail is a paid mutator transaction binding the contract method 0xc8dd612e.
//
// Solidity: function setEmail(bytes32 _email) returns()
func (_Profile *ProfileTransactorSession) SetEmail(_email [32]byte) (*types.Transaction, error) {
	return _Profile.Contract.SetEmail(&_Profile.TransactOpts, _email)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string _name) returns()
func (_Profile *ProfileTransactor) SetName(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "setName", _name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string _name) returns()
func (_Profile *ProfileSession) SetName(_name string) (*types.Transaction, error) {
	return _Profile.Contract.SetName(&_Profile.TransactOpts, _name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string _name) returns()
func (_Profile *ProfileTransactorSession) SetName(_name string) (*types.Transaction, error) {
	return _Profile.Contract.SetName(&_Profile.TransactOpts, _name)
}

// SetPicture is a paid mutator transaction binding the contract method 0xc9a3b592.
//
// Solidity: function setPicture(bytes32 _picture) returns()
func (_Profile *ProfileTransactor) SetPicture(opts *bind.TransactOpts, _picture [32]byte) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "setPicture", _picture)
}

// SetPicture is a paid mutator transaction binding the contract method 0xc9a3b592.
//
// Solidity: function setPicture(bytes32 _picture) returns()
func (_Profile *ProfileSession) SetPicture(_picture [32]byte) (*types.Transaction, error) {
	return _Profile.Contract.SetPicture(&_Profile.TransactOpts, _picture)
}

// SetPicture is a paid mutator transaction binding the contract method 0xc9a3b592.
//
// Solidity: function setPicture(bytes32 _picture) returns()
func (_Profile *ProfileTransactorSession) SetPicture(_picture [32]byte) (*types.Transaction, error) {
	return _Profile.Contract.SetPicture(&_Profile.TransactOpts, _picture)
}
