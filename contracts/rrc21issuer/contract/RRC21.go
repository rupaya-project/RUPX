// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	ethereum "github.com/rupaya-project/rupx"
	"github.com/rupaya-project/rupx/accounts/abi"
	"github.com/rupaya-project/rupx/accounts/abi/bind"
	"github.com/rupaya-project/rupx/common"
	"github.com/rupaya-project/rupx/core/types"
	"github.com/rupaya-project/rupx/event"
)

// IRRC21ABI is the input ABI used to generate the binding from.
const IRRC21ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"estimateFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Fee\",\"type\":\"event\"}]"

// IRRC21Bin is the compiled bytecode used for deploying new contracts.
const IRRC21Bin = `0x`

// DeployIRRC21 deploys a new Ethereum contract, binding an instance of IRRC21 to it.
func DeployIRRC21(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IRRC21, error) {
	parsed, err := abi.JSON(strings.NewReader(IRRC21ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IRRC21Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IRRC21{IRRC21Caller: IRRC21Caller{contract: contract}, IRRC21Transactor: IRRC21Transactor{contract: contract}, IRRC21Filterer: IRRC21Filterer{contract: contract}}, nil
}

// IRRC21 is an auto generated Go binding around an Ethereum contract.
type IRRC21 struct {
	IRRC21Caller     // Read-only binding to the contract
	IRRC21Transactor // Write-only binding to the contract
	IRRC21Filterer   // Log filterer for contract events
}

// IRRC21Caller is an auto generated read-only Go binding around an Ethereum contract.
type IRRC21Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRRC21Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IRRC21Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRRC21Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRRC21Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRRC21Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRRC21Session struct {
	Contract     *IRRC21           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRRC21CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRRC21CallerSession struct {
	Contract *IRRC21Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IRRC21TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRRC21TransactorSession struct {
	Contract     *IRRC21Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRRC21Raw is an auto generated low-level Go binding around an Ethereum contract.
type IRRC21Raw struct {
	Contract *IRRC21 // Generic contract binding to access the raw methods on
}

// IRRC21CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRRC21CallerRaw struct {
	Contract *IRRC21Caller // Generic read-only contract binding to access the raw methods on
}

// IRRC21TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRRC21TransactorRaw struct {
	Contract *IRRC21Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIRRC21 creates a new instance of IRRC21, bound to a specific deployed contract.
func NewIRRC21(address common.Address, backend bind.ContractBackend) (*IRRC21, error) {
	contract, err := bindIRRC21(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRRC21{IRRC21Caller: IRRC21Caller{contract: contract}, IRRC21Transactor: IRRC21Transactor{contract: contract}, IRRC21Filterer: IRRC21Filterer{contract: contract}}, nil
}

// NewIRRC21Caller creates a new read-only instance of IRRC21, bound to a specific deployed contract.
func NewIRRC21Caller(address common.Address, caller bind.ContractCaller) (*IRRC21Caller, error) {
	contract, err := bindIRRC21(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRRC21Caller{contract: contract}, nil
}

// NewIRRC21Transactor creates a new write-only instance of IRRC21, bound to a specific deployed contract.
func NewIRRC21Transactor(address common.Address, transactor bind.ContractTransactor) (*IRRC21Transactor, error) {
	contract, err := bindIRRC21(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRRC21Transactor{contract: contract}, nil
}

// NewIRRC21Filterer creates a new log filterer instance of IRRC21, bound to a specific deployed contract.
func NewIRRC21Filterer(address common.Address, filterer bind.ContractFilterer) (*IRRC21Filterer, error) {
	contract, err := bindIRRC21(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRRC21Filterer{contract: contract}, nil
}

// bindIRRC21 binds a generic wrapper to an already deployed contract.
func bindIRRC21(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IRRC21ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRRC21 *IRRC21Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IRRC21.Contract.IRRC21Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRRC21 *IRRC21Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRRC21.Contract.IRRC21Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRRC21 *IRRC21Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRRC21.Contract.IRRC21Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRRC21 *IRRC21CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IRRC21.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRRC21 *IRRC21TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRRC21.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRRC21 *IRRC21TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRRC21.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_IRRC21 *IRRC21Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IRRC21.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_IRRC21 *IRRC21Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IRRC21.Contract.Allowance(&_IRRC21.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_IRRC21 *IRRC21CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IRRC21.Contract.Allowance(&_IRRC21.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_IRRC21 *IRRC21Caller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IRRC21.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_IRRC21 *IRRC21Session) BalanceOf(who common.Address) (*big.Int, error) {
	return _IRRC21.Contract.BalanceOf(&_IRRC21.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_IRRC21 *IRRC21CallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _IRRC21.Contract.BalanceOf(&_IRRC21.CallOpts, who)
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(value uint256) constant returns(uint256)
func (_IRRC21 *IRRC21Caller) EstimateFee(opts *bind.CallOpts, value *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IRRC21.contract.Call(opts, out, "estimateFee", value)
	return *ret0, err
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(value uint256) constant returns(uint256)
func (_IRRC21 *IRRC21Session) EstimateFee(value *big.Int) (*big.Int, error) {
	return _IRRC21.Contract.EstimateFee(&_IRRC21.CallOpts, value)
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(value uint256) constant returns(uint256)
func (_IRRC21 *IRRC21CallerSession) EstimateFee(value *big.Int) (*big.Int, error) {
	return _IRRC21.Contract.EstimateFee(&_IRRC21.CallOpts, value)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IRRC21 *IRRC21Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IRRC21.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IRRC21 *IRRC21Session) TotalSupply() (*big.Int, error) {
	return _IRRC21.Contract.TotalSupply(&_IRRC21.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IRRC21 *IRRC21CallerSession) TotalSupply() (*big.Int, error) {
	return _IRRC21.Contract.TotalSupply(&_IRRC21.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_IRRC21 *IRRC21Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IRRC21.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_IRRC21 *IRRC21Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IRRC21.Contract.Approve(&_IRRC21.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_IRRC21 *IRRC21TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IRRC21.Contract.Approve(&_IRRC21.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_IRRC21 *IRRC21Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IRRC21.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_IRRC21 *IRRC21Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IRRC21.Contract.Transfer(&_IRRC21.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_IRRC21 *IRRC21TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IRRC21.Contract.Transfer(&_IRRC21.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_IRRC21 *IRRC21Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IRRC21.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_IRRC21 *IRRC21Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IRRC21.Contract.TransferFrom(&_IRRC21.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_IRRC21 *IRRC21TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IRRC21.Contract.TransferFrom(&_IRRC21.TransactOpts, from, to, value)
}

// IRRC21ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IRRC21 contract.
type IRRC21ApprovalIterator struct {
	Event *IRRC21Approval // Event containing the contract specifics and raw log

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
func (it *IRRC21ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRRC21Approval)
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
		it.Event = new(IRRC21Approval)
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
func (it *IRRC21ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRRC21ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRRC21Approval represents a Approval event raised by the IRRC21 contract.
type IRRC21Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_IRRC21 *IRRC21Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IRRC21ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IRRC21.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IRRC21ApprovalIterator{contract: _IRRC21.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_IRRC21 *IRRC21Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IRRC21Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IRRC21.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRRC21Approval)
				if err := _IRRC21.contract.UnpackLog(event, "Approval", log); err != nil {
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

// IRRC21FeeIterator is returned from FilterFee and is used to iterate over the raw logs and unpacked data for Fee events raised by the IRRC21 contract.
type IRRC21FeeIterator struct {
	Event *IRRC21Fee // Event containing the contract specifics and raw log

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
func (it *IRRC21FeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRRC21Fee)
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
		it.Event = new(IRRC21Fee)
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
func (it *IRRC21FeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRRC21FeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRRC21Fee represents a Fee event raised by the IRRC21 contract.
type IRRC21Fee struct {
	From   common.Address
	To     common.Address
	Issuer common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFee is a free log retrieval operation binding the contract event 0xfcf5b3276434181e3c48bd3fe569b8939808f11e845d4b963693b9d7dbd6dd99.
//
// Solidity: event Fee(from indexed address, to indexed address, issuer indexed address, value uint256)
func (_IRRC21 *IRRC21Filterer) FilterFee(opts *bind.FilterOpts, from []common.Address, to []common.Address, issuer []common.Address) (*IRRC21FeeIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _IRRC21.contract.FilterLogs(opts, "Fee", fromRule, toRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &IRRC21FeeIterator{contract: _IRRC21.contract, event: "Fee", logs: logs, sub: sub}, nil
}

// WatchFee is a free log subscription operation binding the contract event 0xfcf5b3276434181e3c48bd3fe569b8939808f11e845d4b963693b9d7dbd6dd99.
//
// Solidity: event Fee(from indexed address, to indexed address, issuer indexed address, value uint256)
func (_IRRC21 *IRRC21Filterer) WatchFee(opts *bind.WatchOpts, sink chan<- *IRRC21Fee, from []common.Address, to []common.Address, issuer []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _IRRC21.contract.WatchLogs(opts, "Fee", fromRule, toRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRRC21Fee)
				if err := _IRRC21.contract.UnpackLog(event, "Fee", log); err != nil {
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

// IRRC21TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IRRC21 contract.
type IRRC21TransferIterator struct {
	Event *IRRC21Transfer // Event containing the contract specifics and raw log

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
func (it *IRRC21TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRRC21Transfer)
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
		it.Event = new(IRRC21Transfer)
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
func (it *IRRC21TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRRC21TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRRC21Transfer represents a Transfer event raised by the IRRC21 contract.
type IRRC21Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_IRRC21 *IRRC21Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IRRC21TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IRRC21.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IRRC21TransferIterator{contract: _IRRC21.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_IRRC21 *IRRC21Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IRRC21Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IRRC21.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRRC21Transfer)
				if err := _IRRC21.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// MyRRC21ABI is the input ABI used to generate the binding from.
const MyRRC21ABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"estimateFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\"},{\"name\":\"cap\",\"type\":\"uint256\"},{\"name\":\"minFee\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Fee\",\"type\":\"event\"}]"

// MyRRC21Bin is the compiled bytecode used for deploying new contracts.
const MyRRC21Bin = `0x60806040523480156200001157600080fd5b5060405162000b3138038062000b3183398101604090815281516020808401519284015160608501516080860151938601805190969590950194919390929091620000639160059190880190620001e7565b50835162000079906006906020870190620001e7565b506007805460ff191660ff85161790556200009e3383640100000000620000d1810204565b620000b23364010000000062000190810204565b620000c681640100000000620001c8810204565b50505050506200028c565b600160a060020a0382161515620000e757600080fd5b60045462000104908264010000000062000842620001cd82021704565b600455600160a060020a0382166000908152602081905260409020546200013a908264010000000062000842620001cd82021704565b600160a060020a0383166000818152602081815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b600160a060020a0381161515620001a657600080fd5b60028054600160a060020a031916600160a060020a0392909216919091179055565b600155565b600082820183811015620001e057600080fd5b9392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200022a57805160ff19168380011785556200025a565b828001600101855582156200025a579182015b828111156200025a5782518255916020019190600101906200023d565b50620002689291506200026c565b5090565b6200028991905b8082111562000268576000815560010162000273565b90565b610895806200029c6000396000f3006080604052600436106100b95763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100be578063095ea7b314610148578063127e8e4d1461018057806318160ddd146101aa5780631d143848146101bf57806323b872dd146101f057806324ec75901461021a578063313ce5671461022f57806370a082311461025a57806395d89b411461027b578063a9059cbb14610290578063dd62ed3e146102b4575b600080fd5b3480156100ca57600080fd5b506100d36102db565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561010d5781810151838201526020016100f5565b50505050905090810190601f16801561013a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561015457600080fd5b5061016c600160a060020a0360043516602435610371565b604080519115158252519081900360200190f35b34801561018c57600080fd5b5061019860043561042b565b60408051918252519081900360200190f35b3480156101b657600080fd5b50610198610457565b3480156101cb57600080fd5b506101d461045d565b60408051600160a060020a039092168252519081900360200190f35b3480156101fc57600080fd5b5061016c600160a060020a036004358116906024351660443561046c565b34801561022657600080fd5b506101986105ac565b34801561023b57600080fd5b506102446105b2565b6040805160ff9092168252519081900360200190f35b34801561026657600080fd5b50610198600160a060020a03600435166105bb565b34801561028757600080fd5b506100d36105d6565b34801561029c57600080fd5b5061016c600160a060020a0360043516602435610637565b3480156102c057600080fd5b50610198600160a060020a03600435811690602435166106f0565b60058054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156103675780601f1061033c57610100808354040283529160200191610367565b820191906000526020600020905b81548152906001019060200180831161034a57829003601f168201915b5050505050905090565b6000600160a060020a038316151561038857600080fd5b6001543360009081526020819052604090205410156103a657600080fd5b336000818152600360209081526040808320600160a060020a03888116855292529091208490556002546001546103e29392919091169061071b565b604080518381529051600160a060020a0385169133917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259181900360200190a350600192915050565b60015460009061045190610445848463ffffffff61080d16565b9063ffffffff61084216565b92915050565b60045490565b600254600160a060020a031690565b6000806104846001548461084290919063ffffffff16565b9050600160a060020a038416151561049b57600080fd5b808311156104a857600080fd5b600160a060020a03851660009081526003602090815260408083203384529091529020548111156104d857600080fd5b600160a060020a038516600090815260036020908152604080832033845290915290205461050c908263ffffffff61085416565b600160a060020a038616600090815260036020908152604080832033845290915290205561053b85858561071b565b600254600154610558918791600160a060020a039091169061071b565b6002546001546040805191825251600160a060020a039283169287169133917ffcf5b3276434181e3c48bd3fe569b8939808f11e845d4b963693b9d7dbd6dd999181900360200190a4506001949350505050565b60015490565b60075460ff1690565b600160a060020a031660009081526020819052604090205490565b60068054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156103675780601f1061033c57610100808354040283529160200191610367565b60008061064f6001548461084290919063ffffffff16565b9050600160a060020a038416151561066657600080fd5b8083111561067357600080fd5b61067e33858561071b565b60025460015461069b913391600160a060020a039091169061071b565b6002546001546040805191825251600160a060020a039283169287169133917ffcf5b3276434181e3c48bd3fe569b8939808f11e845d4b963693b9d7dbd6dd999181900360200190a4600191505b5092915050565b600160a060020a03918216600090815260036020908152604080832093909416825291909152205490565b600160a060020a03831660009081526020819052604090205481111561074057600080fd5b600160a060020a038216151561075557600080fd5b600160a060020a03831660009081526020819052604090205461077e908263ffffffff61085416565b600160a060020a0380851660009081526020819052604080822093909355908416815220546107b3908263ffffffff61084216565b600160a060020a038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b60008083151561082057600091506106e9565b5082820282848281151561083057fe5b041461083b57600080fd5b9392505050565b60008282018381101561083b57600080fd5b60008282111561086357600080fd5b509003905600a165627a7a72305820fafa860f05b0db36e3b8067a75d9068ddc0c82bea68b125e228ef1a02a3a62140029`

// DeployMyRRC21 deploys a new Ethereum contract, binding an instance of MyRRC21 to it.
func DeployMyRRC21(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string, decimals uint8, cap *big.Int, minFee *big.Int) (common.Address, *types.Transaction, *MyRRC21, error) {
	parsed, err := abi.JSON(strings.NewReader(MyRRC21ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MyRRC21Bin), backend, name, symbol, decimals, cap, minFee)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MyRRC21{MyRRC21Caller: MyRRC21Caller{contract: contract}, MyRRC21Transactor: MyRRC21Transactor{contract: contract}, MyRRC21Filterer: MyRRC21Filterer{contract: contract}}, nil
}

// MyRRC21 is an auto generated Go binding around an Ethereum contract.
type MyRRC21 struct {
	MyRRC21Caller     // Read-only binding to the contract
	MyRRC21Transactor // Write-only binding to the contract
	MyRRC21Filterer   // Log filterer for contract events
}

// MyRRC21Caller is an auto generated read-only Go binding around an Ethereum contract.
type MyRRC21Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyRRC21Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MyRRC21Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyRRC21Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MyRRC21Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyRRC21Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MyRRC21Session struct {
	Contract     *MyRRC21          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MyRRC21CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MyRRC21CallerSession struct {
	Contract *MyRRC21Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MyRRC21TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MyRRC21TransactorSession struct {
	Contract     *MyRRC21Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MyRRC21Raw is an auto generated low-level Go binding around an Ethereum contract.
type MyRRC21Raw struct {
	Contract *MyRRC21 // Generic contract binding to access the raw methods on
}

// MyRRC21CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MyRRC21CallerRaw struct {
	Contract *MyRRC21Caller // Generic read-only contract binding to access the raw methods on
}

// MyRRC21TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MyRRC21TransactorRaw struct {
	Contract *MyRRC21Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMyRRC21 creates a new instance of MyRRC21, bound to a specific deployed contract.
func NewMyRRC21(address common.Address, backend bind.ContractBackend) (*MyRRC21, error) {
	contract, err := bindMyRRC21(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MyRRC21{MyRRC21Caller: MyRRC21Caller{contract: contract}, MyRRC21Transactor: MyRRC21Transactor{contract: contract}, MyRRC21Filterer: MyRRC21Filterer{contract: contract}}, nil
}

// NewMyRRC21Caller creates a new read-only instance of MyRRC21, bound to a specific deployed contract.
func NewMyRRC21Caller(address common.Address, caller bind.ContractCaller) (*MyRRC21Caller, error) {
	contract, err := bindMyRRC21(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MyRRC21Caller{contract: contract}, nil
}

// NewMyRRC21Transactor creates a new write-only instance of MyRRC21, bound to a specific deployed contract.
func NewMyRRC21Transactor(address common.Address, transactor bind.ContractTransactor) (*MyRRC21Transactor, error) {
	contract, err := bindMyRRC21(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MyRRC21Transactor{contract: contract}, nil
}

// NewMyRRC21Filterer creates a new log filterer instance of MyRRC21, bound to a specific deployed contract.
func NewMyRRC21Filterer(address common.Address, filterer bind.ContractFilterer) (*MyRRC21Filterer, error) {
	contract, err := bindMyRRC21(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MyRRC21Filterer{contract: contract}, nil
}

// bindMyRRC21 binds a generic wrapper to an already deployed contract.
func bindMyRRC21(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MyRRC21ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyRRC21 *MyRRC21Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MyRRC21.Contract.MyRRC21Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyRRC21 *MyRRC21Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyRRC21.Contract.MyRRC21Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyRRC21 *MyRRC21Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyRRC21.Contract.MyRRC21Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyRRC21 *MyRRC21CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MyRRC21.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyRRC21 *MyRRC21TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyRRC21.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyRRC21 *MyRRC21TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyRRC21.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_MyRRC21 *MyRRC21Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyRRC21.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_MyRRC21 *MyRRC21Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MyRRC21.Contract.Allowance(&_MyRRC21.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_MyRRC21 *MyRRC21CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MyRRC21.Contract.Allowance(&_MyRRC21.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_MyRRC21 *MyRRC21Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyRRC21.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_MyRRC21 *MyRRC21Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MyRRC21.Contract.BalanceOf(&_MyRRC21.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_MyRRC21 *MyRRC21CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MyRRC21.Contract.BalanceOf(&_MyRRC21.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MyRRC21 *MyRRC21Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _MyRRC21.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MyRRC21 *MyRRC21Session) Decimals() (uint8, error) {
	return _MyRRC21.Contract.Decimals(&_MyRRC21.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MyRRC21 *MyRRC21CallerSession) Decimals() (uint8, error) {
	return _MyRRC21.Contract.Decimals(&_MyRRC21.CallOpts)
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(value uint256) constant returns(uint256)
func (_MyRRC21 *MyRRC21Caller) EstimateFee(opts *bind.CallOpts, value *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyRRC21.contract.Call(opts, out, "estimateFee", value)
	return *ret0, err
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(value uint256) constant returns(uint256)
func (_MyRRC21 *MyRRC21Session) EstimateFee(value *big.Int) (*big.Int, error) {
	return _MyRRC21.Contract.EstimateFee(&_MyRRC21.CallOpts, value)
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(value uint256) constant returns(uint256)
func (_MyRRC21 *MyRRC21CallerSession) EstimateFee(value *big.Int) (*big.Int, error) {
	return _MyRRC21.Contract.EstimateFee(&_MyRRC21.CallOpts, value)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_MyRRC21 *MyRRC21Caller) Issuer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MyRRC21.contract.Call(opts, out, "issuer")
	return *ret0, err
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_MyRRC21 *MyRRC21Session) Issuer() (common.Address, error) {
	return _MyRRC21.Contract.Issuer(&_MyRRC21.CallOpts)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_MyRRC21 *MyRRC21CallerSession) Issuer() (common.Address, error) {
	return _MyRRC21.Contract.Issuer(&_MyRRC21.CallOpts)
}

// MinFee is a free data retrieval call binding the contract method 0x24ec7590.
//
// Solidity: function minFee() constant returns(uint256)
func (_MyRRC21 *MyRRC21Caller) MinFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyRRC21.contract.Call(opts, out, "minFee")
	return *ret0, err
}

// MinFee is a free data retrieval call binding the contract method 0x24ec7590.
//
// Solidity: function minFee() constant returns(uint256)
func (_MyRRC21 *MyRRC21Session) MinFee() (*big.Int, error) {
	return _MyRRC21.Contract.MinFee(&_MyRRC21.CallOpts)
}

// MinFee is a free data retrieval call binding the contract method 0x24ec7590.
//
// Solidity: function minFee() constant returns(uint256)
func (_MyRRC21 *MyRRC21CallerSession) MinFee() (*big.Int, error) {
	return _MyRRC21.Contract.MinFee(&_MyRRC21.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MyRRC21 *MyRRC21Caller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MyRRC21.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MyRRC21 *MyRRC21Session) Name() (string, error) {
	return _MyRRC21.Contract.Name(&_MyRRC21.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MyRRC21 *MyRRC21CallerSession) Name() (string, error) {
	return _MyRRC21.Contract.Name(&_MyRRC21.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MyRRC21 *MyRRC21Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MyRRC21.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MyRRC21 *MyRRC21Session) Symbol() (string, error) {
	return _MyRRC21.Contract.Symbol(&_MyRRC21.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MyRRC21 *MyRRC21CallerSession) Symbol() (string, error) {
	return _MyRRC21.Contract.Symbol(&_MyRRC21.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MyRRC21 *MyRRC21Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyRRC21.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MyRRC21 *MyRRC21Session) TotalSupply() (*big.Int, error) {
	return _MyRRC21.Contract.TotalSupply(&_MyRRC21.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MyRRC21 *MyRRC21CallerSession) TotalSupply() (*big.Int, error) {
	return _MyRRC21.Contract.TotalSupply(&_MyRRC21.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_MyRRC21 *MyRRC21Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyRRC21.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_MyRRC21 *MyRRC21Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyRRC21.Contract.Approve(&_MyRRC21.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_MyRRC21 *MyRRC21TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyRRC21.Contract.Approve(&_MyRRC21.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_MyRRC21 *MyRRC21Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyRRC21.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_MyRRC21 *MyRRC21Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyRRC21.Contract.Transfer(&_MyRRC21.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_MyRRC21 *MyRRC21TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyRRC21.Contract.Transfer(&_MyRRC21.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_MyRRC21 *MyRRC21Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyRRC21.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_MyRRC21 *MyRRC21Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyRRC21.Contract.TransferFrom(&_MyRRC21.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_MyRRC21 *MyRRC21TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyRRC21.Contract.TransferFrom(&_MyRRC21.TransactOpts, from, to, value)
}

// MyRRC21ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MyRRC21 contract.
type MyRRC21ApprovalIterator struct {
	Event *MyRRC21Approval // Event containing the contract specifics and raw log

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
func (it *MyRRC21ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyRRC21Approval)
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
		it.Event = new(MyRRC21Approval)
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
func (it *MyRRC21ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyRRC21ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyRRC21Approval represents a Approval event raised by the MyRRC21 contract.
type MyRRC21Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_MyRRC21 *MyRRC21Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MyRRC21ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MyRRC21.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MyRRC21ApprovalIterator{contract: _MyRRC21.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_MyRRC21 *MyRRC21Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MyRRC21Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MyRRC21.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyRRC21Approval)
				if err := _MyRRC21.contract.UnpackLog(event, "Approval", log); err != nil {
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

// MyRRC21FeeIterator is returned from FilterFee and is used to iterate over the raw logs and unpacked data for Fee events raised by the MyRRC21 contract.
type MyRRC21FeeIterator struct {
	Event *MyRRC21Fee // Event containing the contract specifics and raw log

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
func (it *MyRRC21FeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyRRC21Fee)
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
		it.Event = new(MyRRC21Fee)
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
func (it *MyRRC21FeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyRRC21FeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyRRC21Fee represents a Fee event raised by the MyRRC21 contract.
type MyRRC21Fee struct {
	From   common.Address
	To     common.Address
	Issuer common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFee is a free log retrieval operation binding the contract event 0xfcf5b3276434181e3c48bd3fe569b8939808f11e845d4b963693b9d7dbd6dd99.
//
// Solidity: event Fee(from indexed address, to indexed address, issuer indexed address, value uint256)
func (_MyRRC21 *MyRRC21Filterer) FilterFee(opts *bind.FilterOpts, from []common.Address, to []common.Address, issuer []common.Address) (*MyRRC21FeeIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _MyRRC21.contract.FilterLogs(opts, "Fee", fromRule, toRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &MyRRC21FeeIterator{contract: _MyRRC21.contract, event: "Fee", logs: logs, sub: sub}, nil
}

// WatchFee is a free log subscription operation binding the contract event 0xfcf5b3276434181e3c48bd3fe569b8939808f11e845d4b963693b9d7dbd6dd99.
//
// Solidity: event Fee(from indexed address, to indexed address, issuer indexed address, value uint256)
func (_MyRRC21 *MyRRC21Filterer) WatchFee(opts *bind.WatchOpts, sink chan<- *MyRRC21Fee, from []common.Address, to []common.Address, issuer []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _MyRRC21.contract.WatchLogs(opts, "Fee", fromRule, toRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyRRC21Fee)
				if err := _MyRRC21.contract.UnpackLog(event, "Fee", log); err != nil {
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

// MyRRC21TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MyRRC21 contract.
type MyRRC21TransferIterator struct {
	Event *MyRRC21Transfer // Event containing the contract specifics and raw log

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
func (it *MyRRC21TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyRRC21Transfer)
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
		it.Event = new(MyRRC21Transfer)
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
func (it *MyRRC21TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyRRC21TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyRRC21Transfer represents a Transfer event raised by the MyRRC21 contract.
type MyRRC21Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_MyRRC21 *MyRRC21Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MyRRC21TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MyRRC21.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MyRRC21TransferIterator{contract: _MyRRC21.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_MyRRC21 *MyRRC21Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MyRRC21Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MyRRC21.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyRRC21Transfer)
				if err := _MyRRC21.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// RRC21ABI is the input ABI used to generate the binding from.
const RRC21ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"estimateFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Fee\",\"type\":\"event\"}]"

// RRC21Bin is the compiled bytecode used for deploying new contracts.
const RRC21Bin = `0x608060405234801561001057600080fd5b506106aa806100206000396000f3006080604052600436106100985763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461009d578063127e8e4d146100d557806318160ddd146100ff5780631d1438481461011457806323b872dd1461014557806324ec75901461016f57806370a0823114610184578063a9059cbb146101a5578063dd62ed3e146101c9575b600080fd5b3480156100a957600080fd5b506100c1600160a060020a03600435166024356101f0565b604080519115158252519081900360200190f35b3480156100e157600080fd5b506100ed6004356102aa565b60408051918252519081900360200190f35b34801561010b57600080fd5b506100ed6102d6565b34801561012057600080fd5b506101296102dc565b60408051600160a060020a039092168252519081900360200190f35b34801561015157600080fd5b506100c1600160a060020a03600435811690602435166044356102eb565b34801561017b57600080fd5b506100ed61042b565b34801561019057600080fd5b506100ed600160a060020a0360043516610431565b3480156101b157600080fd5b506100c1600160a060020a036004351660243561044c565b3480156101d557600080fd5b506100ed600160a060020a0360043581169060243516610505565b6000600160a060020a038316151561020757600080fd5b60015433600090815260208190526040902054101561022557600080fd5b336000818152600360209081526040808320600160a060020a038881168552925290912084905560025460015461026193929190911690610530565b604080518381529051600160a060020a0385169133917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259181900360200190a350600192915050565b6001546000906102d0906102c4848463ffffffff61062216565b9063ffffffff61065716565b92915050565b60045490565b600254600160a060020a031690565b6000806103036001548461065790919063ffffffff16565b9050600160a060020a038416151561031a57600080fd5b8083111561032757600080fd5b600160a060020a038516600090815260036020908152604080832033845290915290205481111561035757600080fd5b600160a060020a038516600090815260036020908152604080832033845290915290205461038b908263ffffffff61066916565b600160a060020a03861660009081526003602090815260408083203384529091529020556103ba858585610530565b6002546001546103d7918791600160a060020a0390911690610530565b6002546001546040805191825251600160a060020a039283169287169133917ffcf5b3276434181e3c48bd3fe569b8939808f11e845d4b963693b9d7dbd6dd999181900360200190a4506001949350505050565b60015490565b600160a060020a031660009081526020819052604090205490565b6000806104646001548461065790919063ffffffff16565b9050600160a060020a038416151561047b57600080fd5b8083111561048857600080fd5b610493338585610530565b6002546001546104b0913391600160a060020a0390911690610530565b6002546001546040805191825251600160a060020a039283169287169133917ffcf5b3276434181e3c48bd3fe569b8939808f11e845d4b963693b9d7dbd6dd999181900360200190a4600191505b5092915050565b600160a060020a03918216600090815260036020908152604080832093909416825291909152205490565b600160a060020a03831660009081526020819052604090205481111561055557600080fd5b600160a060020a038216151561056a57600080fd5b600160a060020a038316600090815260208190526040902054610593908263ffffffff61066916565b600160a060020a0380851660009081526020819052604080822093909355908416815220546105c8908263ffffffff61065716565b600160a060020a038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b60008083151561063557600091506104fe565b5082820282848281151561064557fe5b041461065057600080fd5b9392505050565b60008282018381101561065057600080fd5b60008282111561067857600080fd5b509003905600a165627a7a72305820a40ecfbb2241cf25b0b2e3e120bec997e2d61dc709b4b92c70939c9b604b417f0029`

// DeployRRC21 deploys a new Ethereum contract, binding an instance of RRC21 to it.
func DeployRRC21(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RRC21, error) {
	parsed, err := abi.JSON(strings.NewReader(RRC21ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RRC21Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RRC21{RRC21Caller: RRC21Caller{contract: contract}, RRC21Transactor: RRC21Transactor{contract: contract}, RRC21Filterer: RRC21Filterer{contract: contract}}, nil
}

// RRC21 is an auto generated Go binding around an Ethereum contract.
type RRC21 struct {
	RRC21Caller     // Read-only binding to the contract
	RRC21Transactor // Write-only binding to the contract
	RRC21Filterer   // Log filterer for contract events
}

// RRC21Caller is an auto generated read-only Go binding around an Ethereum contract.
type RRC21Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RRC21Transactor is an auto generated write-only Go binding around an Ethereum contract.
type RRC21Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RRC21Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RRC21Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RRC21Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RRC21Session struct {
	Contract     *RRC21            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RRC21CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RRC21CallerSession struct {
	Contract *RRC21Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RRC21TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RRC21TransactorSession struct {
	Contract     *RRC21Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RRC21Raw is an auto generated low-level Go binding around an Ethereum contract.
type RRC21Raw struct {
	Contract *RRC21 // Generic contract binding to access the raw methods on
}

// RRC21CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RRC21CallerRaw struct {
	Contract *RRC21Caller // Generic read-only contract binding to access the raw methods on
}

// RRC21TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RRC21TransactorRaw struct {
	Contract *RRC21Transactor // Generic write-only contract binding to access the raw methods on
}

// NewRRC21 creates a new instance of RRC21, bound to a specific deployed contract.
func NewRRC21(address common.Address, backend bind.ContractBackend) (*RRC21, error) {
	contract, err := bindRRC21(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RRC21{RRC21Caller: RRC21Caller{contract: contract}, RRC21Transactor: RRC21Transactor{contract: contract}, RRC21Filterer: RRC21Filterer{contract: contract}}, nil
}

// NewRRC21Caller creates a new read-only instance of RRC21, bound to a specific deployed contract.
func NewRRC21Caller(address common.Address, caller bind.ContractCaller) (*RRC21Caller, error) {
	contract, err := bindRRC21(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RRC21Caller{contract: contract}, nil
}

// NewRRC21Transactor creates a new write-only instance of RRC21, bound to a specific deployed contract.
func NewRRC21Transactor(address common.Address, transactor bind.ContractTransactor) (*RRC21Transactor, error) {
	contract, err := bindRRC21(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RRC21Transactor{contract: contract}, nil
}

// NewRRC21Filterer creates a new log filterer instance of RRC21, bound to a specific deployed contract.
func NewRRC21Filterer(address common.Address, filterer bind.ContractFilterer) (*RRC21Filterer, error) {
	contract, err := bindRRC21(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RRC21Filterer{contract: contract}, nil
}

// bindRRC21 binds a generic wrapper to an already deployed contract.
func bindRRC21(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RRC21ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RRC21 *RRC21Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RRC21.Contract.RRC21Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RRC21 *RRC21Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RRC21.Contract.RRC21Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RRC21 *RRC21Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RRC21.Contract.RRC21Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RRC21 *RRC21CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RRC21.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RRC21 *RRC21TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RRC21.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RRC21 *RRC21TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RRC21.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_RRC21 *RRC21Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RRC21.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_RRC21 *RRC21Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _RRC21.Contract.Allowance(&_RRC21.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_RRC21 *RRC21CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _RRC21.Contract.Allowance(&_RRC21.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_RRC21 *RRC21Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RRC21.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_RRC21 *RRC21Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _RRC21.Contract.BalanceOf(&_RRC21.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_RRC21 *RRC21CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _RRC21.Contract.BalanceOf(&_RRC21.CallOpts, owner)
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(value uint256) constant returns(uint256)
func (_RRC21 *RRC21Caller) EstimateFee(opts *bind.CallOpts, value *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RRC21.contract.Call(opts, out, "estimateFee", value)
	return *ret0, err
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(value uint256) constant returns(uint256)
func (_RRC21 *RRC21Session) EstimateFee(value *big.Int) (*big.Int, error) {
	return _RRC21.Contract.EstimateFee(&_RRC21.CallOpts, value)
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(value uint256) constant returns(uint256)
func (_RRC21 *RRC21CallerSession) EstimateFee(value *big.Int) (*big.Int, error) {
	return _RRC21.Contract.EstimateFee(&_RRC21.CallOpts, value)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_RRC21 *RRC21Caller) Issuer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RRC21.contract.Call(opts, out, "issuer")
	return *ret0, err
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_RRC21 *RRC21Session) Issuer() (common.Address, error) {
	return _RRC21.Contract.Issuer(&_RRC21.CallOpts)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_RRC21 *RRC21CallerSession) Issuer() (common.Address, error) {
	return _RRC21.Contract.Issuer(&_RRC21.CallOpts)
}

// MinFee is a free data retrieval call binding the contract method 0x24ec7590.
//
// Solidity: function minFee() constant returns(uint256)
func (_RRC21 *RRC21Caller) MinFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RRC21.contract.Call(opts, out, "minFee")
	return *ret0, err
}

// MinFee is a free data retrieval call binding the contract method 0x24ec7590.
//
// Solidity: function minFee() constant returns(uint256)
func (_RRC21 *RRC21Session) MinFee() (*big.Int, error) {
	return _RRC21.Contract.MinFee(&_RRC21.CallOpts)
}

// MinFee is a free data retrieval call binding the contract method 0x24ec7590.
//
// Solidity: function minFee() constant returns(uint256)
func (_RRC21 *RRC21CallerSession) MinFee() (*big.Int, error) {
	return _RRC21.Contract.MinFee(&_RRC21.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_RRC21 *RRC21Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RRC21.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_RRC21 *RRC21Session) TotalSupply() (*big.Int, error) {
	return _RRC21.Contract.TotalSupply(&_RRC21.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_RRC21 *RRC21CallerSession) TotalSupply() (*big.Int, error) {
	return _RRC21.Contract.TotalSupply(&_RRC21.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_RRC21 *RRC21Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _RRC21.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_RRC21 *RRC21Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _RRC21.Contract.Approve(&_RRC21.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_RRC21 *RRC21TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _RRC21.Contract.Approve(&_RRC21.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_RRC21 *RRC21Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _RRC21.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_RRC21 *RRC21Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _RRC21.Contract.Transfer(&_RRC21.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_RRC21 *RRC21TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _RRC21.Contract.Transfer(&_RRC21.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_RRC21 *RRC21Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _RRC21.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_RRC21 *RRC21Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _RRC21.Contract.TransferFrom(&_RRC21.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_RRC21 *RRC21TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _RRC21.Contract.TransferFrom(&_RRC21.TransactOpts, from, to, value)
}

// RRC21ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the RRC21 contract.
type RRC21ApprovalIterator struct {
	Event *RRC21Approval // Event containing the contract specifics and raw log

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
func (it *RRC21ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RRC21Approval)
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
		it.Event = new(RRC21Approval)
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
func (it *RRC21ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RRC21ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RRC21Approval represents a Approval event raised by the RRC21 contract.
type RRC21Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_RRC21 *RRC21Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*RRC21ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _RRC21.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &RRC21ApprovalIterator{contract: _RRC21.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_RRC21 *RRC21Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *RRC21Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _RRC21.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RRC21Approval)
				if err := _RRC21.contract.UnpackLog(event, "Approval", log); err != nil {
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

// RRC21FeeIterator is returned from FilterFee and is used to iterate over the raw logs and unpacked data for Fee events raised by the RRC21 contract.
type RRC21FeeIterator struct {
	Event *RRC21Fee // Event containing the contract specifics and raw log

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
func (it *RRC21FeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RRC21Fee)
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
		it.Event = new(RRC21Fee)
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
func (it *RRC21FeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RRC21FeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RRC21Fee represents a Fee event raised by the RRC21 contract.
type RRC21Fee struct {
	From   common.Address
	To     common.Address
	Issuer common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFee is a free log retrieval operation binding the contract event 0xfcf5b3276434181e3c48bd3fe569b8939808f11e845d4b963693b9d7dbd6dd99.
//
// Solidity: event Fee(from indexed address, to indexed address, issuer indexed address, value uint256)
func (_RRC21 *RRC21Filterer) FilterFee(opts *bind.FilterOpts, from []common.Address, to []common.Address, issuer []common.Address) (*RRC21FeeIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _RRC21.contract.FilterLogs(opts, "Fee", fromRule, toRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &RRC21FeeIterator{contract: _RRC21.contract, event: "Fee", logs: logs, sub: sub}, nil
}

// WatchFee is a free log subscription operation binding the contract event 0xfcf5b3276434181e3c48bd3fe569b8939808f11e845d4b963693b9d7dbd6dd99.
//
// Solidity: event Fee(from indexed address, to indexed address, issuer indexed address, value uint256)
func (_RRC21 *RRC21Filterer) WatchFee(opts *bind.WatchOpts, sink chan<- *RRC21Fee, from []common.Address, to []common.Address, issuer []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _RRC21.contract.WatchLogs(opts, "Fee", fromRule, toRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RRC21Fee)
				if err := _RRC21.contract.UnpackLog(event, "Fee", log); err != nil {
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

// RRC21TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the RRC21 contract.
type RRC21TransferIterator struct {
	Event *RRC21Transfer // Event containing the contract specifics and raw log

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
func (it *RRC21TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RRC21Transfer)
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
		it.Event = new(RRC21Transfer)
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
func (it *RRC21TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RRC21TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RRC21Transfer represents a Transfer event raised by the RRC21 contract.
type RRC21Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_RRC21 *RRC21Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RRC21TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RRC21.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RRC21TransferIterator{contract: _RRC21.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_RRC21 *RRC21Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *RRC21Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RRC21.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RRC21Transfer)
				if err := _RRC21.contract.UnpackLog(event, "Transfer", log); err != nil {
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
