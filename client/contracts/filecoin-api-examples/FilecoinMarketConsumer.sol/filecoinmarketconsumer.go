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

// BigInt is an auto generated low-level Go binding around an user-defined struct.
type BigInt struct {
	Val []byte
	Neg bool
}

// FilecoinMarketConsumerMetaData contains all meta data concerning the FilecoinMarketConsumer contract.
var FilecoinMarketConsumerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"activationStatus\",\"outputs\":[{\"internalType\":\"int64\",\"name\":\"activated\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"terminated\",\"type\":\"int64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clientCollateral\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"neg\",\"type\":\"bool\"}],\"internalType\":\"structBigInt\",\"name\":\"collateral\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dealClientActorId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dealCommitment\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dealLabel\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dealPricePerEpoch\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"neg\",\"type\":\"bool\"}],\"internalType\":\"structBigInt\",\"name\":\"price_per_epoch\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dealProviderActorId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dealTerm\",\"outputs\":[{\"internalType\":\"int64\",\"name\":\"start\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"end\",\"type\":\"int64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isDealActivated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"providerCollateral\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"neg\",\"type\":\"bool\"}],\"internalType\":\"structBigInt\",\"name\":\"collateral\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dealId\",\"type\":\"uint64\"}],\"name\":\"storeAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dealId\",\"type\":\"uint64\"}],\"name\":\"storeClientCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dealId\",\"type\":\"uint64\"}],\"name\":\"storeDealActivationStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dealId\",\"type\":\"uint64\"}],\"name\":\"storeDealClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dealId\",\"type\":\"uint64\"}],\"name\":\"storeDealClientProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dealId\",\"type\":\"uint64\"}],\"name\":\"storeDealCommitment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dealId\",\"type\":\"uint64\"}],\"name\":\"storeDealLabel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dealId\",\"type\":\"uint64\"}],\"name\":\"storeDealTerm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dealId\",\"type\":\"uint64\"}],\"name\":\"storeDealTotalPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dealId\",\"type\":\"uint64\"}],\"name\":\"storeDealVerificaton\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dealId\",\"type\":\"uint64\"}],\"name\":\"storeProviderCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FilecoinMarketConsumerABI is the input ABI used to generate the binding from.
// Deprecated: Use FilecoinMarketConsumerMetaData.ABI instead.
var FilecoinMarketConsumerABI = FilecoinMarketConsumerMetaData.ABI

// FilecoinMarketConsumer is an auto generated Go binding around an Ethereum contract.
type FilecoinMarketConsumer struct {
	FilecoinMarketConsumerCaller     // Read-only binding to the contract
	FilecoinMarketConsumerTransactor // Write-only binding to the contract
	FilecoinMarketConsumerFilterer   // Log filterer for contract events
}

// FilecoinMarketConsumerCaller is an auto generated read-only Go binding around an Ethereum contract.
type FilecoinMarketConsumerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FilecoinMarketConsumerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FilecoinMarketConsumerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FilecoinMarketConsumerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FilecoinMarketConsumerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FilecoinMarketConsumerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FilecoinMarketConsumerSession struct {
	Contract     *FilecoinMarketConsumer // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// FilecoinMarketConsumerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FilecoinMarketConsumerCallerSession struct {
	Contract *FilecoinMarketConsumerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// FilecoinMarketConsumerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FilecoinMarketConsumerTransactorSession struct {
	Contract     *FilecoinMarketConsumerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// FilecoinMarketConsumerRaw is an auto generated low-level Go binding around an Ethereum contract.
type FilecoinMarketConsumerRaw struct {
	Contract *FilecoinMarketConsumer // Generic contract binding to access the raw methods on
}

// FilecoinMarketConsumerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FilecoinMarketConsumerCallerRaw struct {
	Contract *FilecoinMarketConsumerCaller // Generic read-only contract binding to access the raw methods on
}

// FilecoinMarketConsumerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FilecoinMarketConsumerTransactorRaw struct {
	Contract *FilecoinMarketConsumerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFilecoinMarketConsumer creates a new instance of FilecoinMarketConsumer, bound to a specific deployed contract.
func NewFilecoinMarketConsumer(address common.Address, backend bind.ContractBackend) (*FilecoinMarketConsumer, error) {
	contract, err := bindFilecoinMarketConsumer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FilecoinMarketConsumer{FilecoinMarketConsumerCaller: FilecoinMarketConsumerCaller{contract: contract}, FilecoinMarketConsumerTransactor: FilecoinMarketConsumerTransactor{contract: contract}, FilecoinMarketConsumerFilterer: FilecoinMarketConsumerFilterer{contract: contract}}, nil
}

// NewFilecoinMarketConsumerCaller creates a new read-only instance of FilecoinMarketConsumer, bound to a specific deployed contract.
func NewFilecoinMarketConsumerCaller(address common.Address, caller bind.ContractCaller) (*FilecoinMarketConsumerCaller, error) {
	contract, err := bindFilecoinMarketConsumer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FilecoinMarketConsumerCaller{contract: contract}, nil
}

// NewFilecoinMarketConsumerTransactor creates a new write-only instance of FilecoinMarketConsumer, bound to a specific deployed contract.
func NewFilecoinMarketConsumerTransactor(address common.Address, transactor bind.ContractTransactor) (*FilecoinMarketConsumerTransactor, error) {
	contract, err := bindFilecoinMarketConsumer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FilecoinMarketConsumerTransactor{contract: contract}, nil
}

// NewFilecoinMarketConsumerFilterer creates a new log filterer instance of FilecoinMarketConsumer, bound to a specific deployed contract.
func NewFilecoinMarketConsumerFilterer(address common.Address, filterer bind.ContractFilterer) (*FilecoinMarketConsumerFilterer, error) {
	contract, err := bindFilecoinMarketConsumer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FilecoinMarketConsumerFilterer{contract: contract}, nil
}

// bindFilecoinMarketConsumer binds a generic wrapper to an already deployed contract.
func bindFilecoinMarketConsumer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FilecoinMarketConsumerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FilecoinMarketConsumer *FilecoinMarketConsumerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FilecoinMarketConsumer.Contract.FilecoinMarketConsumerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FilecoinMarketConsumer *FilecoinMarketConsumerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FilecoinMarketConsumer.Contract.FilecoinMarketConsumerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FilecoinMarketConsumer *FilecoinMarketConsumerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FilecoinMarketConsumer.Contract.FilecoinMarketConsumerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FilecoinMarketConsumer *FilecoinMarketConsumerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FilecoinMarketConsumer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FilecoinMarketConsumer *FilecoinMarketConsumerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FilecoinMarketConsumer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FilecoinMarketConsumer *FilecoinMarketConsumerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FilecoinMarketConsumer.Contract.contract.Transact(opts, method, params...)
}

// ActivationStatus is a free data retrieval call binding the contract method 0xafaef4b9.
//
// Solidity: function activationStatus() view returns(int64 activated, int64 terminated)
func (_FilecoinMarketConsumer *FilecoinMarketConsumerCaller) ActivationStatus(opts *bind.CallOpts) (struct {
	Activated  int64
	Terminated int64
}, error) {
	var out []interface{}
	err := _FilecoinMarketConsumer.contract.Call(opts, &out, "activationStatus")

	outstruct := new(struct {
		Activated  int64
		Terminated int64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Activated = *abi.ConvertType(out[0], new(int64)).(*int64)
	outstruct.Terminated = *abi.ConvertType(out[1], new(int64)).(*int64)

	return *outstruct, err

}

// ActivationStatus is a free data retrieval call binding the contract method 0xafaef4b9.
//
// Solidity: function activationStatus() view returns(int64 activated, int64 terminated)
func (_FilecoinMarketConsumer *FilecoinMarketConsumerSession) ActivationStatus() (struct {
	Activated  int64
	Terminated int64
}, error) {
	return _FilecoinMarketConsumer.Contract.ActivationStatus(&_FilecoinMarketConsumer.CallOpts)
}

// ActivationStatus is a free data retrieval call binding the contract method 0xafaef4b9.
//
// Solidity: function activationStatus() view returns(int64 activated, int64 terminated)
func (_FilecoinMarketConsumer *FilecoinMarketConsumerCallerSession) ActivationStatus() (struct {
	Activated  int64
	Terminated int64
}, error) {
	return _FilecoinMarketConsumer.Contract.ActivationStatus(&_FilecoinMarketConsumer.CallOpts)
}

// ClientCollateral is a free data retrieval call binding the contract method 0xa0b19fd5.
//
// Solidity: function clientCollateral() view returns((bytes,bool) collateral)
func (_FilecoinMarketConsumer *FilecoinMarketConsumerCaller) ClientCollateral(opts *bind.CallOpts) (BigInt, error) {
	var out []interface{}
	err := _FilecoinMarketConsumer.contract.Call(opts, &out, "clientCollateral")

	if err != nil {
		return *new(BigInt), err
	}

	out0 := *abi.ConvertType(out[0], new(BigInt)).(*BigInt)

	return out0, err

}

// ClientCollateral is a free data retrieval call binding the contract method 0xa0b19fd5.
//
// Solidity: function clientCollateral() view returns((bytes,bool) collateral)
func (_FilecoinMarketConsumer *FilecoinMarketConsumerSession) ClientCollateral() (BigInt, error) {
	return _FilecoinMarketConsumer.Contract.ClientCollateral(&_FilecoinMarketConsumer.CallOpts)
}

// ClientCollateral is a free data retrieval call binding the contract method 0xa0b19fd5.
//
// Solidity: function clientCollateral() view returns((bytes,bool) collateral)
func (_FilecoinMarketConsumer *FilecoinMarketConsumerCallerSession) ClientCollateral() (BigInt, error) {
	return _FilecoinMarketConsumer.Contract.ClientCollateral(&_FilecoinMarketConsumer.CallOpts)
}

// DealClientActorId is a free data retrieval call binding the contract method 0x95f298cc.
//
// Solidity: function dealClientActorId() view returns(uint64)
func (_FilecoinMarketConsumer *FilecoinMarketConsumerCaller) DealClientActorId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FilecoinMarketConsumer.contract.Call(opts, &out, "dealClientActorId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DealClientActorId is a free data retrieval call binding the contract method 0x95f298cc.
//
// Solidity: function dealClientActorId() view returns(uint64)
func (_FilecoinMarketConsumer *FilecoinMarketConsumerSession) DealClientActorId() (uint64, error) {
	return _FilecoinMarketConsumer.Contract.DealClientActorId(&_FilecoinMarketConsumer.CallOpts)
}

// DealClientActorId is a free data retrieval call binding the contract method 0x95f298cc.
//
// Solidity: function dealClientActorId() view returns(uint64)
func (_FilecoinMarketConsumer *FilecoinMarketConsumerCallerSession) DealClientActorId() (uint64, error) {
	return _FilecoinMarketConsumer.Contract.DealClientActorId(&_FilecoinMarketConsumer.CallOpts)
}

// DealCommitment is a free data retrieval call binding the contract method 0x047a90c5.
//
// Solidity: function dealCommitment() view returns(bytes data, uint64 size)
func (_FilecoinMarketConsumer *FilecoinMarketConsumerCaller) DealCommitment(opts *bind.CallOpts) (struct {
	Data []byte
	Size uint64
}, error) {
	var out []interface{}
	err := _FilecoinMarketConsumer.contract.Call(opts, &out, "dealCommitment")

	outstruct := new(struct {
		Data []byte
		Size uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Data = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Size = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// DealCommitment is a free data retrieval call binding the contract method 0x047a90c5.
//
// Solidity: function dealCommitment() view returns(bytes data, uint64 size)
func (_FilecoinMarketConsumer *FilecoinMarketConsumerSession) DealCommitment() (struct {
	Data []byte
	Size uint64
}, error) {
	return _FilecoinMarketConsumer.Contract.DealCommitment(&_FilecoinMarketConsumer.CallOpts)
}

// DealCommitment is a free data retrieval call binding the contract method 0x047a90c5.
//
// Solidity: function dealCommitment() view returns(bytes data, uint64 size)
func (_FilecoinMarketConsumer *FilecoinMarketConsumerCallerSession) DealCommitment() (struct {
	Data []byte
	Size uint64
}, error) {
	return _FilecoinMarketConsumer.Contract.DealCommitment(&_FilecoinMarketConsumer.CallOpts)
}

// DealLabel is a free data retrieval call binding the contract method 0x6dd64d4c.
//
// Solidity: function dealLabel() view returns(string)
func (_FilecoinMarketConsumer *FilecoinMarketConsumerCaller) DealLabel(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FilecoinMarketConsumer.contract.Call(opts, &out, "dealLabel")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DealLabel is a free data retrieval call binding the contract method 0x6dd64d4c.
//
// Solidity: function dealLabel() view returns(string)
func (_FilecoinMarketConsumer *FilecoinMarketConsumerSession) DealLabel() (string, error) {
	return _FilecoinMarketConsumer.Contract.DealLabel(&_FilecoinMarketConsumer.CallOpts)
}

// DealLabel is a free data retrieval call binding the contract method 0x6dd64d4c.
//
// Solidity: function dealLabel() view returns(string)
func (_FilecoinMarketConsumer *FilecoinMarketConsumerCallerSession) DealLabel() (string, error) {
	return _FilecoinMarketConsumer.Contract.DealLabel(&_FilecoinMarketConsumer.CallOpts)
}

// DealPricePerEpoch is a free data retrieval call binding the contract method 0xe6171e1c.
//
// Solidity: function dealPricePerEpoch() view returns((bytes,bool) price_per_epoch)
func (_FilecoinMarketConsumer *FilecoinMarketConsumerCaller) DealPricePerEpoch(opts *bind.CallOpts) (BigInt, error) {
	var out []interface{}
	err := _FilecoinMarketConsumer.contract.Call(opts, &out, "dealPricePerEpoch")

	if err != nil {
		return *new(BigInt), err
	}

	out0 := *abi.ConvertType(out[0], new(BigInt)).(*BigInt)

	return out0, err

}

// DealPricePerEpoch is a free data retrieval call binding the contract method 0xe6171e1c.
//
// Solidity: function dealPricePerEpoch() view returns((bytes,bool) price_per_epoch)
func (_FilecoinMarketConsumer *FilecoinMarketConsumerSession) DealPricePerEpoch() (BigInt, error) {
	return _FilecoinMarketConsumer.Contract.DealPricePerEpoch(&_FilecoinMarketConsumer.CallOpts)
}

// DealPricePerEpoch is a free data retrieval call binding the contract method 0xe6171e1c.
//
// Solidity: function dealPricePerEpoch() view returns((bytes,bool) price_per_epoch)
func (_FilecoinMarketConsumer *FilecoinMarketConsumerCallerSession) DealPricePerEpoch() (BigInt, error) {
	return _FilecoinMarketConsumer.Contract.DealPricePerEpoch(&_FilecoinMarketConsumer.CallOpts)
}

// DealProviderActorId is a free data retrieval call binding the contract method 0xb693420b.
//
// Solidity: function dealProviderActorId() view returns(uint64)
func (_FilecoinMarketConsumer *FilecoinMarketConsumerCaller) DealProviderActorId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FilecoinMarketConsumer.contract.Call(opts, &out, "dealProviderActorId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DealProviderActorId is a free data retrieval call binding the contract method 0xb693420b.
//
// Solidity: function dealProviderActorId() view returns(uint64)
func (_FilecoinMarketConsumer *FilecoinMarketConsumerSession) DealProviderActorId() (uint64, error) {
	return _FilecoinMarketConsumer.Contract.DealProviderActorId(&_FilecoinMarketConsumer.CallOpts)
}

// DealProviderActorId is a free data retrieval call binding the contract method 0xb693420b.
//
// Solidity: function dealProviderActorId() view returns(uint64)
func (_FilecoinMarketConsumer *FilecoinMarketConsumerCallerSession) DealProviderActorId() (uint64, error) {
	return _FilecoinMarketConsumer.Contract.DealProviderActorId(&_FilecoinMarketConsumer.CallOpts)
}

// DealTerm is a free data retrieval call binding the contract method 0x7f6c8235.
//
// Solidity: function dealTerm() view returns(int64 start, int64 end)
func (_FilecoinMarketConsumer *FilecoinMarketConsumerCaller) DealTerm(opts *bind.CallOpts) (struct {
	Start int64
	End   int64
}, error) {
	var out []interface{}
	err := _FilecoinMarketConsumer.contract.Call(opts, &out, "dealTerm")

	outstruct := new(struct {
		Start int64
		End   int64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Start = *abi.ConvertType(out[0], new(int64)).(*int64)
	outstruct.End = *abi.ConvertType(out[1], new(int64)).(*int64)

	return *outstruct, err

}

// DealTerm is a free data retrieval call binding the contract method 0x7f6c8235.
//
// Solidity: function dealTerm() view returns(int64 start, int64 end)
func (_FilecoinMarketConsumer *FilecoinMarketConsumerSession) DealTerm() (struct {
	Start int64
	End   int64
}, error) {
	return _FilecoinMarketConsumer.Contract.DealTerm(&_FilecoinMarketConsumer.CallOpts)
}

// DealTerm is a free data retrieval call binding the contract method 0x7f6c8235.
//
// Solidity: function dealTerm() view returns(int64 start, int64 end)
func (_FilecoinMarketConsumer *FilecoinMarketConsumerCallerSession) DealTerm() (struct {
	Start int64
	End   int64
}, error) {
	return _FilecoinMarketConsumer.Contract.DealTerm(&_FilecoinMarketConsumer.CallOpts)
}

// IsDealActivated is a free data retrieval call binding the contract method 0xe3a5cee8.
//
// Solidity: function isDealActivated() view returns(bool)
func (_FilecoinMarketConsumer *FilecoinMarketConsumerCaller) IsDealActivated(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FilecoinMarketConsumer.contract.Call(opts, &out, "isDealActivated")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDealActivated is a free data retrieval call binding the contract method 0xe3a5cee8.
//
// Solidity: function isDealActivated() view returns(bool)
func (_FilecoinMarketConsumer *FilecoinMarketConsumerSession) IsDealActivated() (bool, error) {
	return _FilecoinMarketConsumer.Contract.IsDealActivated(&_FilecoinMarketConsumer.CallOpts)
}

// IsDealActivated is a free data retrieval call binding the contract method 0xe3a5cee8.
//
// Solidity: function isDealActivated() view returns(bool)
func (_FilecoinMarketConsumer *FilecoinMarketConsumerCallerSession) IsDealActivated() (bool, error) {
	return _FilecoinMarketConsumer.Contract.IsDealActivated(&_FilecoinMarketConsumer.CallOpts)
}

// ProviderCollateral is a free data retrieval call binding the contract method 0x0958f87b.
//
// Solidity: function providerCollateral() view returns((bytes,bool) collateral)
func (_FilecoinMarketConsumer *FilecoinMarketConsumerCaller) ProviderCollateral(opts *bind.CallOpts) (BigInt, error) {
	var out []interface{}
	err := _FilecoinMarketConsumer.contract.Call(opts, &out, "providerCollateral")

	if err != nil {
		return *new(BigInt), err
	}

	out0 := *abi.ConvertType(out[0], new(BigInt)).(*BigInt)

	return out0, err

}

// ProviderCollateral is a free data retrieval call binding the contract method 0x0958f87b.
//
// Solidity: function providerCollateral() view returns((bytes,bool) collateral)
func (_FilecoinMarketConsumer *FilecoinMarketConsumerSession) ProviderCollateral() (BigInt, error) {
	return _FilecoinMarketConsumer.Contract.ProviderCollateral(&_FilecoinMarketConsumer.CallOpts)
}

// ProviderCollateral is a free data retrieval call binding the contract method 0x0958f87b.
//
// Solidity: function providerCollateral() view returns((bytes,bool) collateral)
func (_FilecoinMarketConsumer *FilecoinMarketConsumerCallerSession) ProviderCollateral() (BigInt, error) {
	return _FilecoinMarketConsumer.Contract.ProviderCollateral(&_FilecoinMarketConsumer.CallOpts)
}

// StoreAll is a paid mutator transaction binding the contract method 0xf3d4ef29.
//
// Solidity: function storeAll(uint64 dealId) returns()
func (_FilecoinMarketConsumer *FilecoinMarketConsumerTransactor) StoreAll(opts *bind.TransactOpts, dealId uint64) (*types.Transaction, error) {
	return _FilecoinMarketConsumer.contract.Transact(opts, "storeAll", dealId)
}

// StoreAll is a paid mutator transaction binding the contract method 0xf3d4ef29.
//
// Solidity: function storeAll(uint64 dealId) returns()
func (_FilecoinMarketConsumer *FilecoinMarketConsumerSession) StoreAll(dealId uint64) (*types.Transaction, error) {
	return _FilecoinMarketConsumer.Contract.StoreAll(&_FilecoinMarketConsumer.TransactOpts, dealId)
}

// StoreAll is a paid mutator transaction binding the contract method 0xf3d4ef29.
//
// Solidity: function storeAll(uint64 dealId) returns()
func (_FilecoinMarketConsumer *FilecoinMarketConsumerTransactorSession) StoreAll(dealId uint64) (*types.Transaction, error) {
	return _FilecoinMarketConsumer.Contract.StoreAll(&_FilecoinMarketConsumer.TransactOpts, dealId)
}

// StoreClientCollateral is a paid mutator transaction binding the contract method 0xd5f268e7.
//
// Solidity: function storeClientCollateral(uint64 dealId) returns()
func (_FilecoinMarketConsumer *FilecoinMarketConsumerTransactor) StoreClientCollateral(opts *bind.TransactOpts, dealId uint64) (*types.Transaction, error) {
	return _FilecoinMarketConsumer.contract.Transact(opts, "storeClientCollateral", dealId)
}

// StoreClientCollateral is a paid mutator transaction binding the contract method 0xd5f268e7.
//
// Solidity: function storeClientCollateral(uint64 dealId) returns()
func (_FilecoinMarketConsumer *FilecoinMarketConsumerSession) StoreClientCollateral(dealId uint64) (*types.Transaction, error) {
	return _FilecoinMarketConsumer.Contract.StoreClientCollateral(&_FilecoinMarketConsumer.TransactOpts, dealId)
}

// StoreClientCollateral is a paid mutator transaction binding the contract method 0xd5f268e7.
//
// Solidity: function storeClientCollateral(uint64 dealId) returns()
func (_FilecoinMarketConsumer *FilecoinMarketConsumerTransactorSession) StoreClientCollateral(dealId uint64) (*types.Transaction, error) {
	return _FilecoinMarketConsumer.Contract.StoreClientCollateral(&_FilecoinMarketConsumer.TransactOpts, dealId)
}

// StoreDealActivationStatus is a paid mutator transaction binding the contract method 0x2b9e649f.
//
// Solidity: function storeDealActivationStatus(uint64 dealId) returns()
func (_FilecoinMarketConsumer *FilecoinMarketConsumerTransactor) StoreDealActivationStatus(opts *bind.TransactOpts, dealId uint64) (*types.Transaction, error) {
	return _FilecoinMarketConsumer.contract.Transact(opts, "storeDealActivationStatus", dealId)
}

// StoreDealActivationStatus is a paid mutator transaction binding the contract method 0x2b9e649f.
//
// Solidity: function storeDealActivationStatus(uint64 dealId) returns()
func (_FilecoinMarketConsumer *FilecoinMarketConsumerSession) StoreDealActivationStatus(dealId uint64) (*types.Transaction, error) {
	return _FilecoinMarketConsumer.Contract.StoreDealActivationStatus(&_FilecoinMarketConsumer.TransactOpts, dealId)
}

// StoreDealActivationStatus is a paid mutator transaction binding the contract method 0x2b9e649f.
//
// Solidity: function storeDealActivationStatus(uint64 dealId) returns()
func (_FilecoinMarketConsumer *FilecoinMarketConsumerTransactorSession) StoreDealActivationStatus(dealId uint64) (*types.Transaction, error) {
	return _FilecoinMarketConsumer.Contract.StoreDealActivationStatus(&_FilecoinMarketConsumer.TransactOpts, dealId)
}

// StoreDealClient is a paid mutator transaction binding the contract method 0x9db0f72c.
//
// Solidity: function storeDealClient(uint64 dealId) returns()
func (_FilecoinMarketConsumer *FilecoinMarketConsumerTransactor) StoreDealClient(opts *bind.TransactOpts, dealId uint64) (*types.Transaction, error) {
	return _FilecoinMarketConsumer.contract.Transact(opts, "storeDealClient", dealId)
}

// StoreDealClient is a paid mutator transaction binding the contract method 0x9db0f72c.
//
// Solidity: function storeDealClient(uint64 dealId) returns()
func (_FilecoinMarketConsumer *FilecoinMarketConsumerSession) StoreDealClient(dealId uint64) (*types.Transaction, error) {
	return _FilecoinMarketConsumer.Contract.StoreDealClient(&_FilecoinMarketConsumer.TransactOpts, dealId)
}

// StoreDealClient is a paid mutator transaction binding the contract method 0x9db0f72c.
//
// Solidity: function storeDealClient(uint64 dealId) returns()
func (_FilecoinMarketConsumer *FilecoinMarketConsumerTransactorSession) StoreDealClient(dealId uint64) (*types.Transaction, error) {
	return _FilecoinMarketConsumer.Contract.StoreDealClient(&_FilecoinMarketConsumer.TransactOpts, dealId)
}

// StoreDealClientProvider is a paid mutator transaction binding the contract method 0x481ed8e0.
//
// Solidity: function storeDealClientProvider(uint64 dealId) returns()
func (_FilecoinMarketConsumer *FilecoinMarketConsumerTransactor) StoreDealClientProvider(opts *bind.TransactOpts, dealId uint64) (*types.Transaction, error) {
	return _FilecoinMarketConsumer.contract.Transact(opts, "storeDealClientProvider", dealId)
}

// StoreDealClientProvider is a paid mutator transaction binding the contract method 0x481ed8e0.
//
// Solidity: function storeDealClientProvider(uint64 dealId) returns()
func (_FilecoinMarketConsumer *FilecoinMarketConsumerSession) StoreDealClientProvider(dealId uint64) (*types.Transaction, error) {
	return _FilecoinMarketConsumer.Contract.StoreDealClientProvider(&_FilecoinMarketConsumer.TransactOpts, dealId)
}

// StoreDealClientProvider is a paid mutator transaction binding the contract method 0x481ed8e0.
//
// Solidity: function storeDealClientProvider(uint64 dealId) returns()
func (_FilecoinMarketConsumer *FilecoinMarketConsumerTransactorSession) StoreDealClientProvider(dealId uint64) (*types.Transaction, error) {
	return _FilecoinMarketConsumer.Contract.StoreDealClientProvider(&_FilecoinMarketConsumer.TransactOpts, dealId)
}

// StoreDealCommitment is a paid mutator transaction binding the contract method 0x4c89deac.
//
// Solidity: function storeDealCommitment(uint64 dealId) returns()
func (_FilecoinMarketConsumer *FilecoinMarketConsumerTransactor) StoreDealCommitment(opts *bind.TransactOpts, dealId uint64) (*types.Transaction, error) {
	return _FilecoinMarketConsumer.contract.Transact(opts, "storeDealCommitment", dealId)
}

// StoreDealCommitment is a paid mutator transaction binding the contract method 0x4c89deac.
//
// Solidity: function storeDealCommitment(uint64 dealId) returns()
func (_FilecoinMarketConsumer *FilecoinMarketConsumerSession) StoreDealCommitment(dealId uint64) (*types.Transaction, error) {
	return _FilecoinMarketConsumer.Contract.StoreDealCommitment(&_FilecoinMarketConsumer.TransactOpts, dealId)
}

// StoreDealCommitment is a paid mutator transaction binding the contract method 0x4c89deac.
//
// Solidity: function storeDealCommitment(uint64 dealId) returns()
func (_FilecoinMarketConsumer *FilecoinMarketConsumerTransactorSession) StoreDealCommitment(dealId uint64) (*types.Transaction, error) {
	return _FilecoinMarketConsumer.Contract.StoreDealCommitment(&_FilecoinMarketConsumer.TransactOpts, dealId)
}

// StoreDealLabel is a paid mutator transaction binding the contract method 0xfca9226d.
//
// Solidity: function storeDealLabel(uint64 dealId) returns()
func (_FilecoinMarketConsumer *FilecoinMarketConsumerTransactor) StoreDealLabel(opts *bind.TransactOpts, dealId uint64) (*types.Transaction, error) {
	return _FilecoinMarketConsumer.contract.Transact(opts, "storeDealLabel", dealId)
}

// StoreDealLabel is a paid mutator transaction binding the contract method 0xfca9226d.
//
// Solidity: function storeDealLabel(uint64 dealId) returns()
func (_FilecoinMarketConsumer *FilecoinMarketConsumerSession) StoreDealLabel(dealId uint64) (*types.Transaction, error) {
	return _FilecoinMarketConsumer.Contract.StoreDealLabel(&_FilecoinMarketConsumer.TransactOpts, dealId)
}

// StoreDealLabel is a paid mutator transaction binding the contract method 0xfca9226d.
//
// Solidity: function storeDealLabel(uint64 dealId) returns()
func (_FilecoinMarketConsumer *FilecoinMarketConsumerTransactorSession) StoreDealLabel(dealId uint64) (*types.Transaction, error) {
	return _FilecoinMarketConsumer.Contract.StoreDealLabel(&_FilecoinMarketConsumer.TransactOpts, dealId)
}

// StoreDealTerm is a paid mutator transaction binding the contract method 0xd7d06d44.
//
// Solidity: function storeDealTerm(uint64 dealId) returns()
func (_FilecoinMarketConsumer *FilecoinMarketConsumerTransactor) StoreDealTerm(opts *bind.TransactOpts, dealId uint64) (*types.Transaction, error) {
	return _FilecoinMarketConsumer.contract.Transact(opts, "storeDealTerm", dealId)
}

// StoreDealTerm is a paid mutator transaction binding the contract method 0xd7d06d44.
//
// Solidity: function storeDealTerm(uint64 dealId) returns()
func (_FilecoinMarketConsumer *FilecoinMarketConsumerSession) StoreDealTerm(dealId uint64) (*types.Transaction, error) {
	return _FilecoinMarketConsumer.Contract.StoreDealTerm(&_FilecoinMarketConsumer.TransactOpts, dealId)
}

// StoreDealTerm is a paid mutator transaction binding the contract method 0xd7d06d44.
//
// Solidity: function storeDealTerm(uint64 dealId) returns()
func (_FilecoinMarketConsumer *FilecoinMarketConsumerTransactorSession) StoreDealTerm(dealId uint64) (*types.Transaction, error) {
	return _FilecoinMarketConsumer.Contract.StoreDealTerm(&_FilecoinMarketConsumer.TransactOpts, dealId)
}

// StoreDealTotalPrice is a paid mutator transaction binding the contract method 0xc5f3b1af.
//
// Solidity: function storeDealTotalPrice(uint64 dealId) returns()
func (_FilecoinMarketConsumer *FilecoinMarketConsumerTransactor) StoreDealTotalPrice(opts *bind.TransactOpts, dealId uint64) (*types.Transaction, error) {
	return _FilecoinMarketConsumer.contract.Transact(opts, "storeDealTotalPrice", dealId)
}

// StoreDealTotalPrice is a paid mutator transaction binding the contract method 0xc5f3b1af.
//
// Solidity: function storeDealTotalPrice(uint64 dealId) returns()
func (_FilecoinMarketConsumer *FilecoinMarketConsumerSession) StoreDealTotalPrice(dealId uint64) (*types.Transaction, error) {
	return _FilecoinMarketConsumer.Contract.StoreDealTotalPrice(&_FilecoinMarketConsumer.TransactOpts, dealId)
}

// StoreDealTotalPrice is a paid mutator transaction binding the contract method 0xc5f3b1af.
//
// Solidity: function storeDealTotalPrice(uint64 dealId) returns()
func (_FilecoinMarketConsumer *FilecoinMarketConsumerTransactorSession) StoreDealTotalPrice(dealId uint64) (*types.Transaction, error) {
	return _FilecoinMarketConsumer.Contract.StoreDealTotalPrice(&_FilecoinMarketConsumer.TransactOpts, dealId)
}

// StoreDealVerificaton is a paid mutator transaction binding the contract method 0x10af7c86.
//
// Solidity: function storeDealVerificaton(uint64 dealId) returns()
func (_FilecoinMarketConsumer *FilecoinMarketConsumerTransactor) StoreDealVerificaton(opts *bind.TransactOpts, dealId uint64) (*types.Transaction, error) {
	return _FilecoinMarketConsumer.contract.Transact(opts, "storeDealVerificaton", dealId)
}

// StoreDealVerificaton is a paid mutator transaction binding the contract method 0x10af7c86.
//
// Solidity: function storeDealVerificaton(uint64 dealId) returns()
func (_FilecoinMarketConsumer *FilecoinMarketConsumerSession) StoreDealVerificaton(dealId uint64) (*types.Transaction, error) {
	return _FilecoinMarketConsumer.Contract.StoreDealVerificaton(&_FilecoinMarketConsumer.TransactOpts, dealId)
}

// StoreDealVerificaton is a paid mutator transaction binding the contract method 0x10af7c86.
//
// Solidity: function storeDealVerificaton(uint64 dealId) returns()
func (_FilecoinMarketConsumer *FilecoinMarketConsumerTransactorSession) StoreDealVerificaton(dealId uint64) (*types.Transaction, error) {
	return _FilecoinMarketConsumer.Contract.StoreDealVerificaton(&_FilecoinMarketConsumer.TransactOpts, dealId)
}

// StoreProviderCollateral is a paid mutator transaction binding the contract method 0x519b6837.
//
// Solidity: function storeProviderCollateral(uint64 dealId) returns()
func (_FilecoinMarketConsumer *FilecoinMarketConsumerTransactor) StoreProviderCollateral(opts *bind.TransactOpts, dealId uint64) (*types.Transaction, error) {
	return _FilecoinMarketConsumer.contract.Transact(opts, "storeProviderCollateral", dealId)
}

// StoreProviderCollateral is a paid mutator transaction binding the contract method 0x519b6837.
//
// Solidity: function storeProviderCollateral(uint64 dealId) returns()
func (_FilecoinMarketConsumer *FilecoinMarketConsumerSession) StoreProviderCollateral(dealId uint64) (*types.Transaction, error) {
	return _FilecoinMarketConsumer.Contract.StoreProviderCollateral(&_FilecoinMarketConsumer.TransactOpts, dealId)
}

// StoreProviderCollateral is a paid mutator transaction binding the contract method 0x519b6837.
//
// Solidity: function storeProviderCollateral(uint64 dealId) returns()
func (_FilecoinMarketConsumer *FilecoinMarketConsumerTransactorSession) StoreProviderCollateral(dealId uint64) (*types.Transaction, error) {
	return _FilecoinMarketConsumer.Contract.StoreProviderCollateral(&_FilecoinMarketConsumer.TransactOpts, dealId)
}
