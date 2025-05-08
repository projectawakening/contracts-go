// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERC2771Forwarder

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

// ERC2771ForwarderForwardRequest is an auto generated low-level Go binding around an user-defined struct.
type ERC2771ForwarderForwardRequest struct {
	From     common.Address
	To       common.Address
	Value    *big.Int
	Gas      *big.Int
	Nonce    *big.Int
	Deadline *big.Int
	Data     []byte
}

// ERC2771ForwarderForwardRequestData is an auto generated low-level Go binding around an user-defined struct.
type ERC2771ForwarderForwardRequestData struct {
	From      common.Address
	To        common.Address
	Value     *big.Int
	Gas       *big.Int
	Nonce     *big.Int
	Deadline  *big.Int
	Data      []byte
	Signature []byte
}

// ERC2771ForwarderMetaData contains all meta data concerning the ERC2771Forwarder contract.
var ERC2771ForwarderMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"request\",\"type\":\"tuple\",\"internalType\":\"structERC2771Forwarder.ForwardRequestData\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"executeBatch\",\"inputs\":[{\"name\":\"requests\",\"type\":\"tuple[]\",\"internalType\":\"structERC2771Forwarder.ForwardRequestData[]\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"refundReceiver\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"isNonceUsed\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"structHash\",\"inputs\":[{\"name\":\"request\",\"type\":\"tuple\",\"internalType\":\"structERC2771Forwarder.ForwardRequest\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verify\",\"inputs\":[{\"name\":\"request\",\"type\":\"tuple\",\"internalType\":\"structERC2771Forwarder.ForwardRequestData\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutedForwardRequest\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC2771ForwarderExpiredRequest\",\"inputs\":[{\"name\":\"deadline\",\"type\":\"uint48\",\"internalType\":\"uint48\"}]},{\"type\":\"error\",\"name\":\"ERC2771ForwarderInvalidSigner\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC2771ForwarderMismatchedValue\",\"inputs\":[{\"name\":\"requestedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"msgValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC2771UntrustfulTarget\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"forwarder\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAccountNonce\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"currentNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
}

// ERC2771ForwarderABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC2771ForwarderMetaData.ABI instead.
var ERC2771ForwarderABI = ERC2771ForwarderMetaData.ABI

// ERC2771Forwarder is an auto generated Go binding around an Ethereum contract.
type ERC2771Forwarder struct {
	ERC2771ForwarderCaller     // Read-only binding to the contract
	ERC2771ForwarderTransactor // Write-only binding to the contract
	ERC2771ForwarderFilterer   // Log filterer for contract events
}

// ERC2771ForwarderCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC2771ForwarderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC2771ForwarderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC2771ForwarderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC2771ForwarderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC2771ForwarderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC2771ForwarderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC2771ForwarderSession struct {
	Contract     *ERC2771Forwarder // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC2771ForwarderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC2771ForwarderCallerSession struct {
	Contract *ERC2771ForwarderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ERC2771ForwarderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC2771ForwarderTransactorSession struct {
	Contract     *ERC2771ForwarderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ERC2771ForwarderRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC2771ForwarderRaw struct {
	Contract *ERC2771Forwarder // Generic contract binding to access the raw methods on
}

// ERC2771ForwarderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC2771ForwarderCallerRaw struct {
	Contract *ERC2771ForwarderCaller // Generic read-only contract binding to access the raw methods on
}

// ERC2771ForwarderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC2771ForwarderTransactorRaw struct {
	Contract *ERC2771ForwarderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC2771Forwarder creates a new instance of ERC2771Forwarder, bound to a specific deployed contract.
func NewERC2771Forwarder(address common.Address, backend bind.ContractBackend) (*ERC2771Forwarder, error) {
	contract, err := bindERC2771Forwarder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC2771Forwarder{ERC2771ForwarderCaller: ERC2771ForwarderCaller{contract: contract}, ERC2771ForwarderTransactor: ERC2771ForwarderTransactor{contract: contract}, ERC2771ForwarderFilterer: ERC2771ForwarderFilterer{contract: contract}}, nil
}

// NewERC2771ForwarderCaller creates a new read-only instance of ERC2771Forwarder, bound to a specific deployed contract.
func NewERC2771ForwarderCaller(address common.Address, caller bind.ContractCaller) (*ERC2771ForwarderCaller, error) {
	contract, err := bindERC2771Forwarder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC2771ForwarderCaller{contract: contract}, nil
}

// NewERC2771ForwarderTransactor creates a new write-only instance of ERC2771Forwarder, bound to a specific deployed contract.
func NewERC2771ForwarderTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC2771ForwarderTransactor, error) {
	contract, err := bindERC2771Forwarder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC2771ForwarderTransactor{contract: contract}, nil
}

// NewERC2771ForwarderFilterer creates a new log filterer instance of ERC2771Forwarder, bound to a specific deployed contract.
func NewERC2771ForwarderFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC2771ForwarderFilterer, error) {
	contract, err := bindERC2771Forwarder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC2771ForwarderFilterer{contract: contract}, nil
}

// bindERC2771Forwarder binds a generic wrapper to an already deployed contract.
func bindERC2771Forwarder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC2771ForwarderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC2771Forwarder *ERC2771ForwarderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC2771Forwarder.Contract.ERC2771ForwarderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC2771Forwarder *ERC2771ForwarderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ERC2771ForwarderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC2771Forwarder *ERC2771ForwarderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ERC2771ForwarderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC2771Forwarder *ERC2771ForwarderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC2771Forwarder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC2771Forwarder *ERC2771ForwarderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC2771Forwarder *ERC2771ForwarderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.contract.Transact(opts, method, params...)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_ERC2771Forwarder *ERC2771ForwarderSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _ERC2771Forwarder.Contract.Eip712Domain(&_ERC2771Forwarder.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _ERC2771Forwarder.Contract.Eip712Domain(&_ERC2771Forwarder.CallOpts)
}

// IsNonceUsed is a free data retrieval call binding the contract method 0xcab7e8eb.
//
// Solidity: function isNonceUsed(address owner, uint256 data) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) IsNonceUsed(opts *bind.CallOpts, owner common.Address, data *big.Int) (bool, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "isNonceUsed", owner, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNonceUsed is a free data retrieval call binding the contract method 0xcab7e8eb.
//
// Solidity: function isNonceUsed(address owner, uint256 data) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderSession) IsNonceUsed(owner common.Address, data *big.Int) (bool, error) {
	return _ERC2771Forwarder.Contract.IsNonceUsed(&_ERC2771Forwarder.CallOpts, owner, data)
}

// IsNonceUsed is a free data retrieval call binding the contract method 0xcab7e8eb.
//
// Solidity: function isNonceUsed(address owner, uint256 data) view returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) IsNonceUsed(owner common.Address, data *big.Int) (bool, error) {
	return _ERC2771Forwarder.Contract.IsNonceUsed(&_ERC2771Forwarder.CallOpts, owner, data)
}

// StructHash is a free data retrieval call binding the contract method 0xf1109062.
//
// Solidity: function structHash((address,address,uint256,uint256,uint256,uint48,bytes) request) view returns(bytes32)
func (_ERC2771Forwarder *ERC2771ForwarderCaller) StructHash(opts *bind.CallOpts, request ERC2771ForwarderForwardRequest) ([32]byte, error) {
	var out []interface{}
	err := _ERC2771Forwarder.contract.Call(opts, &out, "structHash", request)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StructHash is a free data retrieval call binding the contract method 0xf1109062.
//
// Solidity: function structHash((address,address,uint256,uint256,uint256,uint48,bytes) request) view returns(bytes32)
func (_ERC2771Forwarder *ERC2771ForwarderSession) StructHash(request ERC2771ForwarderForwardRequest) ([32]byte, error) {
	return _ERC2771Forwarder.Contract.StructHash(&_ERC2771Forwarder.CallOpts, request)
}

// StructHash is a free data retrieval call binding the contract method 0xf1109062.
//
// Solidity: function structHash((address,address,uint256,uint256,uint256,uint48,bytes) request) view returns(bytes32)
func (_ERC2771Forwarder *ERC2771ForwarderCallerSession) StructHash(request ERC2771ForwarderForwardRequest) ([32]byte, error) {
	return _ERC2771Forwarder.Contract.StructHash(&_ERC2771Forwarder.CallOpts, request)
}

// Execute is a paid mutator transaction binding the contract method 0xccff1ca1.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,uint48,bytes,bytes) request) payable returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) Execute(opts *bind.TransactOpts, request ERC2771ForwarderForwardRequestData) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "execute", request)
}

// Execute is a paid mutator transaction binding the contract method 0xccff1ca1.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,uint48,bytes,bytes) request) payable returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) Execute(request ERC2771ForwarderForwardRequestData) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.Execute(&_ERC2771Forwarder.TransactOpts, request)
}

// Execute is a paid mutator transaction binding the contract method 0xccff1ca1.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,uint48,bytes,bytes) request) payable returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) Execute(request ERC2771ForwarderForwardRequestData) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.Execute(&_ERC2771Forwarder.TransactOpts, request)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0xa6b4806a.
//
// Solidity: function executeBatch((address,address,uint256,uint256,uint256,uint48,bytes,bytes)[] requests, address refundReceiver) payable returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) ExecuteBatch(opts *bind.TransactOpts, requests []ERC2771ForwarderForwardRequestData, refundReceiver common.Address) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "executeBatch", requests, refundReceiver)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0xa6b4806a.
//
// Solidity: function executeBatch((address,address,uint256,uint256,uint256,uint48,bytes,bytes)[] requests, address refundReceiver) payable returns()
func (_ERC2771Forwarder *ERC2771ForwarderSession) ExecuteBatch(requests []ERC2771ForwarderForwardRequestData, refundReceiver common.Address) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ExecuteBatch(&_ERC2771Forwarder.TransactOpts, requests, refundReceiver)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0xa6b4806a.
//
// Solidity: function executeBatch((address,address,uint256,uint256,uint256,uint48,bytes,bytes)[] requests, address refundReceiver) payable returns()
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) ExecuteBatch(requests []ERC2771ForwarderForwardRequestData, refundReceiver common.Address) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.ExecuteBatch(&_ERC2771Forwarder.TransactOpts, requests, refundReceiver)
}

// Verify is a paid mutator transaction binding the contract method 0xafeb5022.
//
// Solidity: function verify((address,address,uint256,uint256,uint256,uint48,bytes,bytes) request) returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderTransactor) Verify(opts *bind.TransactOpts, request ERC2771ForwarderForwardRequestData) (*types.Transaction, error) {
	return _ERC2771Forwarder.contract.Transact(opts, "verify", request)
}

// Verify is a paid mutator transaction binding the contract method 0xafeb5022.
//
// Solidity: function verify((address,address,uint256,uint256,uint256,uint48,bytes,bytes) request) returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderSession) Verify(request ERC2771ForwarderForwardRequestData) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.Verify(&_ERC2771Forwarder.TransactOpts, request)
}

// Verify is a paid mutator transaction binding the contract method 0xafeb5022.
//
// Solidity: function verify((address,address,uint256,uint256,uint256,uint48,bytes,bytes) request) returns(bool)
func (_ERC2771Forwarder *ERC2771ForwarderTransactorSession) Verify(request ERC2771ForwarderForwardRequestData) (*types.Transaction, error) {
	return _ERC2771Forwarder.Contract.Verify(&_ERC2771Forwarder.TransactOpts, request)
}

// ERC2771ForwarderEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the ERC2771Forwarder contract.
type ERC2771ForwarderEIP712DomainChangedIterator struct {
	Event *ERC2771ForwarderEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *ERC2771ForwarderEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC2771ForwarderEIP712DomainChanged)
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
		it.Event = new(ERC2771ForwarderEIP712DomainChanged)
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
func (it *ERC2771ForwarderEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC2771ForwarderEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC2771ForwarderEIP712DomainChanged represents a EIP712DomainChanged event raised by the ERC2771Forwarder contract.
type ERC2771ForwarderEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_ERC2771Forwarder *ERC2771ForwarderFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*ERC2771ForwarderEIP712DomainChangedIterator, error) {

	logs, sub, err := _ERC2771Forwarder.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &ERC2771ForwarderEIP712DomainChangedIterator{contract: _ERC2771Forwarder.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_ERC2771Forwarder *ERC2771ForwarderFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *ERC2771ForwarderEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _ERC2771Forwarder.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC2771ForwarderEIP712DomainChanged)
				if err := _ERC2771Forwarder.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_ERC2771Forwarder *ERC2771ForwarderFilterer) ParseEIP712DomainChanged(log types.Log) (*ERC2771ForwarderEIP712DomainChanged, error) {
	event := new(ERC2771ForwarderEIP712DomainChanged)
	if err := _ERC2771Forwarder.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC2771ForwarderExecutedForwardRequestIterator is returned from FilterExecutedForwardRequest and is used to iterate over the raw logs and unpacked data for ExecutedForwardRequest events raised by the ERC2771Forwarder contract.
type ERC2771ForwarderExecutedForwardRequestIterator struct {
	Event *ERC2771ForwarderExecutedForwardRequest // Event containing the contract specifics and raw log

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
func (it *ERC2771ForwarderExecutedForwardRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC2771ForwarderExecutedForwardRequest)
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
		it.Event = new(ERC2771ForwarderExecutedForwardRequest)
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
func (it *ERC2771ForwarderExecutedForwardRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC2771ForwarderExecutedForwardRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC2771ForwarderExecutedForwardRequest represents a ExecutedForwardRequest event raised by the ERC2771Forwarder contract.
type ERC2771ForwarderExecutedForwardRequest struct {
	Signer  common.Address
	Nonce   *big.Int
	Success bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExecutedForwardRequest is a free log retrieval operation binding the contract event 0x842fb24a83793558587a3dab2be7674da4a51d09c5542d6dd354e5d0ea70813c.
//
// Solidity: event ExecutedForwardRequest(address indexed signer, uint256 nonce, bool success)
func (_ERC2771Forwarder *ERC2771ForwarderFilterer) FilterExecutedForwardRequest(opts *bind.FilterOpts, signer []common.Address) (*ERC2771ForwarderExecutedForwardRequestIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _ERC2771Forwarder.contract.FilterLogs(opts, "ExecutedForwardRequest", signerRule)
	if err != nil {
		return nil, err
	}
	return &ERC2771ForwarderExecutedForwardRequestIterator{contract: _ERC2771Forwarder.contract, event: "ExecutedForwardRequest", logs: logs, sub: sub}, nil
}

// WatchExecutedForwardRequest is a free log subscription operation binding the contract event 0x842fb24a83793558587a3dab2be7674da4a51d09c5542d6dd354e5d0ea70813c.
//
// Solidity: event ExecutedForwardRequest(address indexed signer, uint256 nonce, bool success)
func (_ERC2771Forwarder *ERC2771ForwarderFilterer) WatchExecutedForwardRequest(opts *bind.WatchOpts, sink chan<- *ERC2771ForwarderExecutedForwardRequest, signer []common.Address) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _ERC2771Forwarder.contract.WatchLogs(opts, "ExecutedForwardRequest", signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC2771ForwarderExecutedForwardRequest)
				if err := _ERC2771Forwarder.contract.UnpackLog(event, "ExecutedForwardRequest", log); err != nil {
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

// ParseExecutedForwardRequest is a log parse operation binding the contract event 0x842fb24a83793558587a3dab2be7674da4a51d09c5542d6dd354e5d0ea70813c.
//
// Solidity: event ExecutedForwardRequest(address indexed signer, uint256 nonce, bool success)
func (_ERC2771Forwarder *ERC2771ForwarderFilterer) ParseExecutedForwardRequest(log types.Log) (*ERC2771ForwarderExecutedForwardRequest, error) {
	event := new(ERC2771ForwarderExecutedForwardRequest)
	if err := _ERC2771Forwarder.contract.UnpackLog(event, "ExecutedForwardRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
