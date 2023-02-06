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

// DataCapTypesBurnFromParams is an auto generated low-level Go binding around an user-defined struct.
type DataCapTypesBurnFromParams struct {
	Owner  []byte
	Amount BigInt
}

// DataCapTypesBurnFromReturn is an auto generated low-level Go binding around an user-defined struct.
type DataCapTypesBurnFromReturn struct {
	Balance   BigInt
	Allowance BigInt
}

// DataCapTypesBurnParams is an auto generated low-level Go binding around an user-defined struct.
type DataCapTypesBurnParams struct {
	Amount BigInt
}

// DataCapTypesBurnReturn is an auto generated low-level Go binding around an user-defined struct.
type DataCapTypesBurnReturn struct {
	Balance BigInt
}

// DataCapTypesDecreaseAllowanceParams is an auto generated low-level Go binding around an user-defined struct.
type DataCapTypesDecreaseAllowanceParams struct {
	Operator []byte
	Decrease BigInt
}

// DataCapTypesGetAllowanceParams is an auto generated low-level Go binding around an user-defined struct.
type DataCapTypesGetAllowanceParams struct {
	Owner    []byte
	Operator []byte
}

// DataCapTypesIncreaseAllowanceParams is an auto generated low-level Go binding around an user-defined struct.
type DataCapTypesIncreaseAllowanceParams struct {
	Operator []byte
	Increase BigInt
}

// DataCapTypesRevokeAllowanceParams is an auto generated low-level Go binding around an user-defined struct.
type DataCapTypesRevokeAllowanceParams struct {
	Operator []byte
}

// DataCapTypesTransferFromParams is an auto generated low-level Go binding around an user-defined struct.
type DataCapTypesTransferFromParams struct {
	From         []byte
	To           []byte
	Amount       BigInt
	OperatorData []byte
}

// DataCapTypesTransferFromReturn is an auto generated low-level Go binding around an user-defined struct.
type DataCapTypesTransferFromReturn struct {
	FromBalance   BigInt
	ToBalance     BigInt
	Allowance     BigInt
	RecipientData []byte
}

// DataCapTypesTransferParams is an auto generated low-level Go binding around an user-defined struct.
type DataCapTypesTransferParams struct {
	To           []byte
	Amount       BigInt
	OperatorData []byte
}

// DataCapTypesTransferReturn is an auto generated low-level Go binding around an user-defined struct.
type DataCapTypesTransferReturn struct {
	FromBalance   BigInt
	ToBalance     BigInt
	RecipientData []byte
}

// IDataCapMetaData contains all meta data concerning the IDataCap contract.
var IDataCapMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"owner\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"operator\",\"type\":\"bytes\"}],\"internalType\":\"structDataCapTypes.GetAllowanceParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"allowance\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"neg\",\"type\":\"bool\"}],\"internalType\":\"structBigInt\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"addr\",\"type\":\"bytes\"}],\"name\":\"balance\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"neg\",\"type\":\"bool\"}],\"internalType\":\"structBigInt\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"neg\",\"type\":\"bool\"}],\"internalType\":\"structBigInt\",\"name\":\"amount\",\"type\":\"tuple\"}],\"internalType\":\"structDataCapTypes.BurnParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"burn\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"neg\",\"type\":\"bool\"}],\"internalType\":\"structBigInt\",\"name\":\"balance\",\"type\":\"tuple\"}],\"internalType\":\"structDataCapTypes.BurnReturn\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"owner\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"neg\",\"type\":\"bool\"}],\"internalType\":\"structBigInt\",\"name\":\"amount\",\"type\":\"tuple\"}],\"internalType\":\"structDataCapTypes.BurnFromParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"burnFrom\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"neg\",\"type\":\"bool\"}],\"internalType\":\"structBigInt\",\"name\":\"balance\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"neg\",\"type\":\"bool\"}],\"internalType\":\"structBigInt\",\"name\":\"allowance\",\"type\":\"tuple\"}],\"internalType\":\"structDataCapTypes.BurnFromReturn\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"operator\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"neg\",\"type\":\"bool\"}],\"internalType\":\"structBigInt\",\"name\":\"decrease\",\"type\":\"tuple\"}],\"internalType\":\"structDataCapTypes.DecreaseAllowanceParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"neg\",\"type\":\"bool\"}],\"internalType\":\"structBigInt\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"operator\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"neg\",\"type\":\"bool\"}],\"internalType\":\"structBigInt\",\"name\":\"increase\",\"type\":\"tuple\"}],\"internalType\":\"structDataCapTypes.IncreaseAllowanceParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"neg\",\"type\":\"bool\"}],\"internalType\":\"structBigInt\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"operator\",\"type\":\"bytes\"}],\"internalType\":\"structDataCapTypes.RevokeAllowanceParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"revokeAllowance\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"neg\",\"type\":\"bool\"}],\"internalType\":\"structBigInt\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"neg\",\"type\":\"bool\"}],\"internalType\":\"structBigInt\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"to\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"neg\",\"type\":\"bool\"}],\"internalType\":\"structBigInt\",\"name\":\"amount\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"operator_data\",\"type\":\"bytes\"}],\"internalType\":\"structDataCapTypes.TransferParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"transfer\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"neg\",\"type\":\"bool\"}],\"internalType\":\"structBigInt\",\"name\":\"from_balance\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"neg\",\"type\":\"bool\"}],\"internalType\":\"structBigInt\",\"name\":\"to_balance\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"recipient_data\",\"type\":\"bytes\"}],\"internalType\":\"structDataCapTypes.TransferReturn\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"from\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"to\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"neg\",\"type\":\"bool\"}],\"internalType\":\"structBigInt\",\"name\":\"amount\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"operator_data\",\"type\":\"bytes\"}],\"internalType\":\"structDataCapTypes.TransferFromParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"transferFrom\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"neg\",\"type\":\"bool\"}],\"internalType\":\"structBigInt\",\"name\":\"from_balance\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"neg\",\"type\":\"bool\"}],\"internalType\":\"structBigInt\",\"name\":\"to_balance\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"neg\",\"type\":\"bool\"}],\"internalType\":\"structBigInt\",\"name\":\"allowance\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"recipient_data\",\"type\":\"bytes\"}],\"internalType\":\"structDataCapTypes.TransferFromReturn\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IDataCapABI is the input ABI used to generate the binding from.
// Deprecated: Use IDataCapMetaData.ABI instead.
var IDataCapABI = IDataCapMetaData.ABI

// IDataCap is an auto generated Go binding around an Ethereum contract.
type IDataCap struct {
	IDataCapCaller     // Read-only binding to the contract
	IDataCapTransactor // Write-only binding to the contract
	IDataCapFilterer   // Log filterer for contract events
}

// IDataCapCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDataCapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDataCapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDataCapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDataCapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDataCapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDataCapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDataCapSession struct {
	Contract     *IDataCap         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDataCapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDataCapCallerSession struct {
	Contract *IDataCapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IDataCapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDataCapTransactorSession struct {
	Contract     *IDataCapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IDataCapRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDataCapRaw struct {
	Contract *IDataCap // Generic contract binding to access the raw methods on
}

// IDataCapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDataCapCallerRaw struct {
	Contract *IDataCapCaller // Generic read-only contract binding to access the raw methods on
}

// IDataCapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDataCapTransactorRaw struct {
	Contract *IDataCapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDataCap creates a new instance of IDataCap, bound to a specific deployed contract.
func NewIDataCap(address common.Address, backend bind.ContractBackend) (*IDataCap, error) {
	contract, err := bindIDataCap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDataCap{IDataCapCaller: IDataCapCaller{contract: contract}, IDataCapTransactor: IDataCapTransactor{contract: contract}, IDataCapFilterer: IDataCapFilterer{contract: contract}}, nil
}

// NewIDataCapCaller creates a new read-only instance of IDataCap, bound to a specific deployed contract.
func NewIDataCapCaller(address common.Address, caller bind.ContractCaller) (*IDataCapCaller, error) {
	contract, err := bindIDataCap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDataCapCaller{contract: contract}, nil
}

// NewIDataCapTransactor creates a new write-only instance of IDataCap, bound to a specific deployed contract.
func NewIDataCapTransactor(address common.Address, transactor bind.ContractTransactor) (*IDataCapTransactor, error) {
	contract, err := bindIDataCap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDataCapTransactor{contract: contract}, nil
}

// NewIDataCapFilterer creates a new log filterer instance of IDataCap, bound to a specific deployed contract.
func NewIDataCapFilterer(address common.Address, filterer bind.ContractFilterer) (*IDataCapFilterer, error) {
	contract, err := bindIDataCap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDataCapFilterer{contract: contract}, nil
}

// bindIDataCap binds a generic wrapper to an already deployed contract.
func bindIDataCap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IDataCapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDataCap *IDataCapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDataCap.Contract.IDataCapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDataCap *IDataCapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDataCap.Contract.IDataCapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDataCap *IDataCapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDataCap.Contract.IDataCapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDataCap *IDataCapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDataCap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDataCap *IDataCapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDataCap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDataCap *IDataCapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDataCap.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a paid mutator transaction binding the contract method 0xce0a0b35.
//
// Solidity: function allowance((bytes,bytes) params) returns((bytes,bool))
func (_IDataCap *IDataCapTransactor) Allowance(opts *bind.TransactOpts, params DataCapTypesGetAllowanceParams) (*types.Transaction, error) {
	return _IDataCap.contract.Transact(opts, "allowance", params)
}

// Allowance is a paid mutator transaction binding the contract method 0xce0a0b35.
//
// Solidity: function allowance((bytes,bytes) params) returns((bytes,bool))
func (_IDataCap *IDataCapSession) Allowance(params DataCapTypesGetAllowanceParams) (*types.Transaction, error) {
	return _IDataCap.Contract.Allowance(&_IDataCap.TransactOpts, params)
}

// Allowance is a paid mutator transaction binding the contract method 0xce0a0b35.
//
// Solidity: function allowance((bytes,bytes) params) returns((bytes,bool))
func (_IDataCap *IDataCapTransactorSession) Allowance(params DataCapTypesGetAllowanceParams) (*types.Transaction, error) {
	return _IDataCap.Contract.Allowance(&_IDataCap.TransactOpts, params)
}

// Balance is a paid mutator transaction binding the contract method 0x5363301d.
//
// Solidity: function balance(bytes addr) returns((bytes,bool))
func (_IDataCap *IDataCapTransactor) Balance(opts *bind.TransactOpts, addr []byte) (*types.Transaction, error) {
	return _IDataCap.contract.Transact(opts, "balance", addr)
}

// Balance is a paid mutator transaction binding the contract method 0x5363301d.
//
// Solidity: function balance(bytes addr) returns((bytes,bool))
func (_IDataCap *IDataCapSession) Balance(addr []byte) (*types.Transaction, error) {
	return _IDataCap.Contract.Balance(&_IDataCap.TransactOpts, addr)
}

// Balance is a paid mutator transaction binding the contract method 0x5363301d.
//
// Solidity: function balance(bytes addr) returns((bytes,bool))
func (_IDataCap *IDataCapTransactorSession) Balance(addr []byte) (*types.Transaction, error) {
	return _IDataCap.Contract.Balance(&_IDataCap.TransactOpts, addr)
}

// Burn is a paid mutator transaction binding the contract method 0x5d69fefc.
//
// Solidity: function burn(((bytes,bool)) params) returns(((bytes,bool)))
func (_IDataCap *IDataCapTransactor) Burn(opts *bind.TransactOpts, params DataCapTypesBurnParams) (*types.Transaction, error) {
	return _IDataCap.contract.Transact(opts, "burn", params)
}

// Burn is a paid mutator transaction binding the contract method 0x5d69fefc.
//
// Solidity: function burn(((bytes,bool)) params) returns(((bytes,bool)))
func (_IDataCap *IDataCapSession) Burn(params DataCapTypesBurnParams) (*types.Transaction, error) {
	return _IDataCap.Contract.Burn(&_IDataCap.TransactOpts, params)
}

// Burn is a paid mutator transaction binding the contract method 0x5d69fefc.
//
// Solidity: function burn(((bytes,bool)) params) returns(((bytes,bool)))
func (_IDataCap *IDataCapTransactorSession) Burn(params DataCapTypesBurnParams) (*types.Transaction, error) {
	return _IDataCap.Contract.Burn(&_IDataCap.TransactOpts, params)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x47f0a805.
//
// Solidity: function burnFrom((bytes,(bytes,bool)) params) returns(((bytes,bool),(bytes,bool)))
func (_IDataCap *IDataCapTransactor) BurnFrom(opts *bind.TransactOpts, params DataCapTypesBurnFromParams) (*types.Transaction, error) {
	return _IDataCap.contract.Transact(opts, "burnFrom", params)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x47f0a805.
//
// Solidity: function burnFrom((bytes,(bytes,bool)) params) returns(((bytes,bool),(bytes,bool)))
func (_IDataCap *IDataCapSession) BurnFrom(params DataCapTypesBurnFromParams) (*types.Transaction, error) {
	return _IDataCap.Contract.BurnFrom(&_IDataCap.TransactOpts, params)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x47f0a805.
//
// Solidity: function burnFrom((bytes,(bytes,bool)) params) returns(((bytes,bool),(bytes,bool)))
func (_IDataCap *IDataCapTransactorSession) BurnFrom(params DataCapTypesBurnFromParams) (*types.Transaction, error) {
	return _IDataCap.Contract.BurnFrom(&_IDataCap.TransactOpts, params)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0x024325c4.
//
// Solidity: function decreaseAllowance((bytes,(bytes,bool)) params) returns((bytes,bool))
func (_IDataCap *IDataCapTransactor) DecreaseAllowance(opts *bind.TransactOpts, params DataCapTypesDecreaseAllowanceParams) (*types.Transaction, error) {
	return _IDataCap.contract.Transact(opts, "decreaseAllowance", params)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0x024325c4.
//
// Solidity: function decreaseAllowance((bytes,(bytes,bool)) params) returns((bytes,bool))
func (_IDataCap *IDataCapSession) DecreaseAllowance(params DataCapTypesDecreaseAllowanceParams) (*types.Transaction, error) {
	return _IDataCap.Contract.DecreaseAllowance(&_IDataCap.TransactOpts, params)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0x024325c4.
//
// Solidity: function decreaseAllowance((bytes,(bytes,bool)) params) returns((bytes,bool))
func (_IDataCap *IDataCapTransactorSession) DecreaseAllowance(params DataCapTypesDecreaseAllowanceParams) (*types.Transaction, error) {
	return _IDataCap.Contract.DecreaseAllowance(&_IDataCap.TransactOpts, params)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0xab10649e.
//
// Solidity: function increaseAllowance((bytes,(bytes,bool)) params) returns((bytes,bool))
func (_IDataCap *IDataCapTransactor) IncreaseAllowance(opts *bind.TransactOpts, params DataCapTypesIncreaseAllowanceParams) (*types.Transaction, error) {
	return _IDataCap.contract.Transact(opts, "increaseAllowance", params)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0xab10649e.
//
// Solidity: function increaseAllowance((bytes,(bytes,bool)) params) returns((bytes,bool))
func (_IDataCap *IDataCapSession) IncreaseAllowance(params DataCapTypesIncreaseAllowanceParams) (*types.Transaction, error) {
	return _IDataCap.Contract.IncreaseAllowance(&_IDataCap.TransactOpts, params)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0xab10649e.
//
// Solidity: function increaseAllowance((bytes,(bytes,bool)) params) returns((bytes,bool))
func (_IDataCap *IDataCapTransactorSession) IncreaseAllowance(params DataCapTypesIncreaseAllowanceParams) (*types.Transaction, error) {
	return _IDataCap.Contract.IncreaseAllowance(&_IDataCap.TransactOpts, params)
}

// Name is a paid mutator transaction binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string)
func (_IDataCap *IDataCapTransactor) Name(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDataCap.contract.Transact(opts, "name")
}

// Name is a paid mutator transaction binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string)
func (_IDataCap *IDataCapSession) Name() (*types.Transaction, error) {
	return _IDataCap.Contract.Name(&_IDataCap.TransactOpts)
}

// Name is a paid mutator transaction binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string)
func (_IDataCap *IDataCapTransactorSession) Name() (*types.Transaction, error) {
	return _IDataCap.Contract.Name(&_IDataCap.TransactOpts)
}

// RevokeAllowance is a paid mutator transaction binding the contract method 0xe5cb0214.
//
// Solidity: function revokeAllowance((bytes) params) returns((bytes,bool))
func (_IDataCap *IDataCapTransactor) RevokeAllowance(opts *bind.TransactOpts, params DataCapTypesRevokeAllowanceParams) (*types.Transaction, error) {
	return _IDataCap.contract.Transact(opts, "revokeAllowance", params)
}

// RevokeAllowance is a paid mutator transaction binding the contract method 0xe5cb0214.
//
// Solidity: function revokeAllowance((bytes) params) returns((bytes,bool))
func (_IDataCap *IDataCapSession) RevokeAllowance(params DataCapTypesRevokeAllowanceParams) (*types.Transaction, error) {
	return _IDataCap.Contract.RevokeAllowance(&_IDataCap.TransactOpts, params)
}

// RevokeAllowance is a paid mutator transaction binding the contract method 0xe5cb0214.
//
// Solidity: function revokeAllowance((bytes) params) returns((bytes,bool))
func (_IDataCap *IDataCapTransactorSession) RevokeAllowance(params DataCapTypesRevokeAllowanceParams) (*types.Transaction, error) {
	return _IDataCap.Contract.RevokeAllowance(&_IDataCap.TransactOpts, params)
}

// Symbol is a paid mutator transaction binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(string)
func (_IDataCap *IDataCapTransactor) Symbol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDataCap.contract.Transact(opts, "symbol")
}

// Symbol is a paid mutator transaction binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(string)
func (_IDataCap *IDataCapSession) Symbol() (*types.Transaction, error) {
	return _IDataCap.Contract.Symbol(&_IDataCap.TransactOpts)
}

// Symbol is a paid mutator transaction binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(string)
func (_IDataCap *IDataCapTransactorSession) Symbol() (*types.Transaction, error) {
	return _IDataCap.Contract.Symbol(&_IDataCap.TransactOpts)
}

// TotalSupply is a paid mutator transaction binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns((bytes,bool))
func (_IDataCap *IDataCapTransactor) TotalSupply(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDataCap.contract.Transact(opts, "totalSupply")
}

// TotalSupply is a paid mutator transaction binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns((bytes,bool))
func (_IDataCap *IDataCapSession) TotalSupply() (*types.Transaction, error) {
	return _IDataCap.Contract.TotalSupply(&_IDataCap.TransactOpts)
}

// TotalSupply is a paid mutator transaction binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns((bytes,bool))
func (_IDataCap *IDataCapTransactorSession) TotalSupply() (*types.Transaction, error) {
	return _IDataCap.Contract.TotalSupply(&_IDataCap.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0x38902cbe.
//
// Solidity: function transfer((bytes,(bytes,bool),bytes) params) returns(((bytes,bool),(bytes,bool),bytes))
func (_IDataCap *IDataCapTransactor) Transfer(opts *bind.TransactOpts, params DataCapTypesTransferParams) (*types.Transaction, error) {
	return _IDataCap.contract.Transact(opts, "transfer", params)
}

// Transfer is a paid mutator transaction binding the contract method 0x38902cbe.
//
// Solidity: function transfer((bytes,(bytes,bool),bytes) params) returns(((bytes,bool),(bytes,bool),bytes))
func (_IDataCap *IDataCapSession) Transfer(params DataCapTypesTransferParams) (*types.Transaction, error) {
	return _IDataCap.Contract.Transfer(&_IDataCap.TransactOpts, params)
}

// Transfer is a paid mutator transaction binding the contract method 0x38902cbe.
//
// Solidity: function transfer((bytes,(bytes,bool),bytes) params) returns(((bytes,bool),(bytes,bool),bytes))
func (_IDataCap *IDataCapTransactorSession) Transfer(params DataCapTypesTransferParams) (*types.Transaction, error) {
	return _IDataCap.Contract.Transfer(&_IDataCap.TransactOpts, params)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x24872061.
//
// Solidity: function transferFrom((bytes,bytes,(bytes,bool),bytes) params) returns(((bytes,bool),(bytes,bool),(bytes,bool),bytes))
func (_IDataCap *IDataCapTransactor) TransferFrom(opts *bind.TransactOpts, params DataCapTypesTransferFromParams) (*types.Transaction, error) {
	return _IDataCap.contract.Transact(opts, "transferFrom", params)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x24872061.
//
// Solidity: function transferFrom((bytes,bytes,(bytes,bool),bytes) params) returns(((bytes,bool),(bytes,bool),(bytes,bool),bytes))
func (_IDataCap *IDataCapSession) TransferFrom(params DataCapTypesTransferFromParams) (*types.Transaction, error) {
	return _IDataCap.Contract.TransferFrom(&_IDataCap.TransactOpts, params)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x24872061.
//
// Solidity: function transferFrom((bytes,bytes,(bytes,bool),bytes) params) returns(((bytes,bool),(bytes,bool),(bytes,bool),bytes))
func (_IDataCap *IDataCapTransactorSession) TransferFrom(params DataCapTypesTransferFromParams) (*types.Transaction, error) {
	return _IDataCap.Contract.TransferFrom(&_IDataCap.TransactOpts, params)
}
