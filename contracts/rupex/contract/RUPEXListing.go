// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"strings"

	"github.com/rupaya-project/rupx/accounts/abi"
	"github.com/rupaya-project/rupx/accounts/abi/bind"
	"github.com/rupaya-project/rupx/common"
	"github.com/rupaya-project/rupx/core/types"
)

// RUPEXListingABI is the input ABI used to generate the binding from.
const RUPEXListingABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"tokens\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"apply\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// RUPEXListingBin is the compiled bytecode used for deploying new contracts.
const RUPEXListingBin = `0x608060405234801561001057600080fd5b506102be806100206000396000f3006080604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416639d63848a811461005b578063a3ff31b5146100c0578063c6b32f34146100f5575b600080fd5b34801561006757600080fd5b5061007061010b565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156100ac578181015183820152602001610094565b505050509050019250505060405180910390f35b3480156100cc57600080fd5b506100e1600160a060020a036004351661016d565b604080519115158252519081900360200190f35b610109600160a060020a036004351661018b565b005b6060600080548060200260200160405190810160405280929190818152602001828054801561016357602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610145575b5050505050905090565b600160a060020a031660009081526001602052604090205460ff1690565b80600160a060020a03811615156101a157600080fd5b600160a060020a03811660009081526001602081905260409091205460ff16151514156101cd57600080fd5b683635c9adc5dea0000034146101e257600080fd5b6040516068903480156108fc02916000818181858888f1935050505015801561020f573d6000803e3d6000fd5b505060008054600180820183557f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563909101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039490941693841790556040805160208082018352838252948452919093529190209051815460ff19169015151790555600a165627a7a723058206d2dc0ce827743c25efa82f99e7830ade39d28e17f4d651573f89e0460a6626a0029`

// DeployRUPEXListing deploys a new Ethereum contract, binding an instance of RUPEXListing to it.
func DeployRUPEXListing(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RUPEXListing, error) {
	parsed, err := abi.JSON(strings.NewReader(RUPEXListingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RUPEXListingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RUPEXListing{RUPEXListingCaller: RUPEXListingCaller{contract: contract}, RUPEXListingTransactor: RUPEXListingTransactor{contract: contract}, RUPEXListingFilterer: RUPEXListingFilterer{contract: contract}}, nil
}

// RUPEXListing is an auto generated Go binding around an Ethereum contract.
type RUPEXListing struct {
	RUPEXListingCaller     // Read-only binding to the contract
	RUPEXListingTransactor // Write-only binding to the contract
	RUPEXListingFilterer   // Log filterer for contract events
}

// RUPEXListingCaller is an auto generated read-only Go binding around an Ethereum contract.
type RUPEXListingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RUPEXListingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RUPEXListingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RUPEXListingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RUPEXListingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RUPEXListingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RUPEXListingSession struct {
	Contract     *RUPEXListing     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RUPEXListingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RUPEXListingCallerSession struct {
	Contract *RUPEXListingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RUPEXListingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RUPEXListingTransactorSession struct {
	Contract     *RUPEXListingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RUPEXListingRaw is an auto generated low-level Go binding around an Ethereum contract.
type RUPEXListingRaw struct {
	Contract *RUPEXListing // Generic contract binding to access the raw methods on
}

// RUPEXListingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RUPEXListingCallerRaw struct {
	Contract *RUPEXListingCaller // Generic read-only contract binding to access the raw methods on
}

// RUPEXListingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RUPEXListingTransactorRaw struct {
	Contract *RUPEXListingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRUPEXListing creates a new instance of RUPEXListing, bound to a specific deployed contract.
func NewRUPEXListing(address common.Address, backend bind.ContractBackend) (*RUPEXListing, error) {
	contract, err := bindRUPEXListing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RUPEXListing{RUPEXListingCaller: RUPEXListingCaller{contract: contract}, RUPEXListingTransactor: RUPEXListingTransactor{contract: contract}, RUPEXListingFilterer: RUPEXListingFilterer{contract: contract}}, nil
}

// NewRUPEXListingCaller creates a new read-only instance of RUPEXListing, bound to a specific deployed contract.
func NewRUPEXListingCaller(address common.Address, caller bind.ContractCaller) (*RUPEXListingCaller, error) {
	contract, err := bindRUPEXListing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RUPEXListingCaller{contract: contract}, nil
}

// NewRUPEXListingTransactor creates a new write-only instance of RUPEXListing, bound to a specific deployed contract.
func NewRUPEXListingTransactor(address common.Address, transactor bind.ContractTransactor) (*RUPEXListingTransactor, error) {
	contract, err := bindRUPEXListing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RUPEXListingTransactor{contract: contract}, nil
}

// NewRUPEXListingFilterer creates a new log filterer instance of RUPEXListing, bound to a specific deployed contract.
func NewRUPEXListingFilterer(address common.Address, filterer bind.ContractFilterer) (*RUPEXListingFilterer, error) {
	contract, err := bindRUPEXListing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RUPEXListingFilterer{contract: contract}, nil
}

// bindRUPEXListing binds a generic wrapper to an already deployed contract.
func bindRUPEXListing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RUPEXListingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RUPEXListing *RUPEXListingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RUPEXListing.Contract.RUPEXListingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RUPEXListing *RUPEXListingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RUPEXListing.Contract.RUPEXListingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RUPEXListing *RUPEXListingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RUPEXListing.Contract.RUPEXListingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RUPEXListing *RUPEXListingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RUPEXListing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RUPEXListing *RUPEXListingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RUPEXListing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RUPEXListing *RUPEXListingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RUPEXListing.Contract.contract.Transact(opts, method, params...)
}

// GetTokenStatus is a free data retrieval call binding the contract method 0xa3ff31b5.
//
// Solidity: function getTokenStatus(token address) constant returns(bool)
func (_RUPEXListing *RUPEXListingCaller) GetTokenStatus(opts *bind.CallOpts, token common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RUPEXListing.contract.Call(opts, out, "getTokenStatus", token)
	return *ret0, err
}

// GetTokenStatus is a free data retrieval call binding the contract method 0xa3ff31b5.
//
// Solidity: function getTokenStatus(token address) constant returns(bool)
func (_RUPEXListing *RUPEXListingSession) GetTokenStatus(token common.Address) (bool, error) {
	return _RUPEXListing.Contract.GetTokenStatus(&_RUPEXListing.CallOpts, token)
}

// GetTokenStatus is a free data retrieval call binding the contract method 0xa3ff31b5.
//
// Solidity: function getTokenStatus(token address) constant returns(bool)
func (_RUPEXListing *RUPEXListingCallerSession) GetTokenStatus(token common.Address) (bool, error) {
	return _RUPEXListing.Contract.GetTokenStatus(&_RUPEXListing.CallOpts, token)
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(address[])
func (_RUPEXListing *RUPEXListingCaller) Tokens(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _RUPEXListing.contract.Call(opts, out, "tokens")
	return *ret0, err
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(address[])
func (_RUPEXListing *RUPEXListingSession) Tokens() ([]common.Address, error) {
	return _RUPEXListing.Contract.Tokens(&_RUPEXListing.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(address[])
func (_RUPEXListing *RUPEXListingCallerSession) Tokens() ([]common.Address, error) {
	return _RUPEXListing.Contract.Tokens(&_RUPEXListing.CallOpts)
}

// Apply is a paid mutator transaction binding the contract method 0xc6b32f34.
//
// Solidity: function apply(token address) returns()
func (_RUPEXListing *RUPEXListingTransactor) Apply(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _RUPEXListing.contract.Transact(opts, "apply", token)
}

// Apply is a paid mutator transaction binding the contract method 0xc6b32f34.
//
// Solidity: function apply(token address) returns()
func (_RUPEXListing *RUPEXListingSession) Apply(token common.Address) (*types.Transaction, error) {
	return _RUPEXListing.Contract.Apply(&_RUPEXListing.TransactOpts, token)
}

// Apply is a paid mutator transaction binding the contract method 0xc6b32f34.
//
// Solidity: function apply(token address) returns()
func (_RUPEXListing *RUPEXListingTransactorSession) Apply(token common.Address) (*types.Transaction, error) {
	return _RUPEXListing.Contract.Apply(&_RUPEXListing.TransactOpts, token)
}
