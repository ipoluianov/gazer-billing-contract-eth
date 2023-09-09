// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

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

// PayloadShopRecord is an auto generated low-level Go binding around an user-defined struct.
type PayloadShopRecord struct {
	IsRegistered bool
	Owner        common.Address
	Payload      [32]byte
}

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_payload\",\"type\":\"bytes32\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_payload\",\"type\":\"bytes32\"}],\"name\":\"buy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getRecordsOfUser\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"payload\",\"type\":\"bytes32\"}],\"internalType\":\"structPayloadShop.Record[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"records\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"payload\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recordsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040525f600255348015610013575f80fd5b506104a0806100215f395ff3fe60806040526004361061003e575f3560e01c806334461067146100425780639c9a1061146100b2578063c4f82d7d146100d3578063eb3dbf83146100ff575b5f80fd5b34801561004d575f80fd5b5061008861005c366004610356565b5f602081905290815260409020805460019091015460ff82169161010090046001600160a01b03169083565b6040805193151584526001600160a01b039092166020840152908201526060015b60405180910390f35b6100c56100c0366004610356565b610114565b6040519081526020016100a9565b3480156100de575f80fd5b506100f26100ed36600461036d565b6101fa565b6040516100a9919061039a565b34801561010a575f80fd5b506100c560025481565b5f8060025460016101259190610411565b6040805160608101825260018082523360208084018281528486018a81525f8881528084528781209651875493516001600160a81b0319909416901515610100600160a81b031916176101006001600160a01b0390941693909302929092178655519484019490945590835281815292822080549182018155825291812090910182905560028054929350906101ba8361042a565b909155505060405183815233907fb8142d42f05d95abf0a6570799774d59276e49ea32a04d9a4ec316ea4a6886bc9060200160405180910390a292915050565b6001600160a01b0381165f908152600160205260408120546060919067ffffffffffffffff81111561022e5761022e610442565b60405190808252806020026020018201604052801561027757816020015b604080516060810182525f80825260208083018290529282015282525f1990920191018161024c5790505b5090505f5b6001600160a01b0384165f9081526001602052604090205481101561034f576001600160a01b0384165f908152600160205260408120805482918291859081106102c8576102c8610456565b5f918252602080832090910154835282810193909352604091820190208151606081018352815460ff8116151582526001600160a01b036101009091041693810193909352600181015491830191909152845190925084908490811061033057610330610456565b60200260200101819052505080806103479061042a565b91505061027c565b5092915050565b5f60208284031215610366575f80fd5b5035919050565b5f6020828403121561037d575f80fd5b81356001600160a01b0381168114610393575f80fd5b9392505050565b602080825282518282018190525f919060409081850190868401855b828110156103f0578151805115158552868101516001600160a01b03168786015285015185850152606090930192908501906001016103b6565b5091979650505050505050565b634e487b7160e01b5f52601160045260245ffd5b80820180821115610424576104246103fd565b92915050565b5f6001820161043b5761043b6103fd565b5060010190565b634e487b7160e01b5f52604160045260245ffd5b634e487b7160e01b5f52603260045260245ffdfea264697066735822122009a7691cf32d4eb6f001620ad1b3dd3cffea1b4cd76a3492b5c02b3ca2aef08164736f6c63430008150033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// GetRecordsOfUser is a free data retrieval call binding the contract method 0xc4f82d7d.
//
// Solidity: function getRecordsOfUser(address addr) view returns((bool,address,bytes32)[])
func (_Api *ApiCaller) GetRecordsOfUser(opts *bind.CallOpts, addr common.Address) ([]PayloadShopRecord, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getRecordsOfUser", addr)

	if err != nil {
		return *new([]PayloadShopRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]PayloadShopRecord)).(*[]PayloadShopRecord)

	return out0, err

}

// GetRecordsOfUser is a free data retrieval call binding the contract method 0xc4f82d7d.
//
// Solidity: function getRecordsOfUser(address addr) view returns((bool,address,bytes32)[])
func (_Api *ApiSession) GetRecordsOfUser(addr common.Address) ([]PayloadShopRecord, error) {
	return _Api.Contract.GetRecordsOfUser(&_Api.CallOpts, addr)
}

// GetRecordsOfUser is a free data retrieval call binding the contract method 0xc4f82d7d.
//
// Solidity: function getRecordsOfUser(address addr) view returns((bool,address,bytes32)[])
func (_Api *ApiCallerSession) GetRecordsOfUser(addr common.Address) ([]PayloadShopRecord, error) {
	return _Api.Contract.GetRecordsOfUser(&_Api.CallOpts, addr)
}

// Records is a free data retrieval call binding the contract method 0x34461067.
//
// Solidity: function records(uint256 ) view returns(bool isRegistered, address owner, bytes32 payload)
func (_Api *ApiCaller) Records(opts *bind.CallOpts, arg0 *big.Int) (struct {
	IsRegistered bool
	Owner        common.Address
	Payload      [32]byte
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "records", arg0)

	outstruct := new(struct {
		IsRegistered bool
		Owner        common.Address
		Payload      [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsRegistered = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Payload = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Records is a free data retrieval call binding the contract method 0x34461067.
//
// Solidity: function records(uint256 ) view returns(bool isRegistered, address owner, bytes32 payload)
func (_Api *ApiSession) Records(arg0 *big.Int) (struct {
	IsRegistered bool
	Owner        common.Address
	Payload      [32]byte
}, error) {
	return _Api.Contract.Records(&_Api.CallOpts, arg0)
}

// Records is a free data retrieval call binding the contract method 0x34461067.
//
// Solidity: function records(uint256 ) view returns(bool isRegistered, address owner, bytes32 payload)
func (_Api *ApiCallerSession) Records(arg0 *big.Int) (struct {
	IsRegistered bool
	Owner        common.Address
	Payload      [32]byte
}, error) {
	return _Api.Contract.Records(&_Api.CallOpts, arg0)
}

// RecordsCount is a free data retrieval call binding the contract method 0xeb3dbf83.
//
// Solidity: function recordsCount() view returns(uint256)
func (_Api *ApiCaller) RecordsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "recordsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RecordsCount is a free data retrieval call binding the contract method 0xeb3dbf83.
//
// Solidity: function recordsCount() view returns(uint256)
func (_Api *ApiSession) RecordsCount() (*big.Int, error) {
	return _Api.Contract.RecordsCount(&_Api.CallOpts)
}

// RecordsCount is a free data retrieval call binding the contract method 0xeb3dbf83.
//
// Solidity: function recordsCount() view returns(uint256)
func (_Api *ApiCallerSession) RecordsCount() (*big.Int, error) {
	return _Api.Contract.RecordsCount(&_Api.CallOpts)
}

// Buy is a paid mutator transaction binding the contract method 0x9c9a1061.
//
// Solidity: function buy(bytes32 _payload) payable returns(uint256)
func (_Api *ApiTransactor) Buy(opts *bind.TransactOpts, _payload [32]byte) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "buy", _payload)
}

// Buy is a paid mutator transaction binding the contract method 0x9c9a1061.
//
// Solidity: function buy(bytes32 _payload) payable returns(uint256)
func (_Api *ApiSession) Buy(_payload [32]byte) (*types.Transaction, error) {
	return _Api.Contract.Buy(&_Api.TransactOpts, _payload)
}

// Buy is a paid mutator transaction binding the contract method 0x9c9a1061.
//
// Solidity: function buy(bytes32 _payload) payable returns(uint256)
func (_Api *ApiTransactorSession) Buy(_payload [32]byte) (*types.Transaction, error) {
	return _Api.Contract.Buy(&_Api.TransactOpts, _payload)
}

// ApiRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the Api contract.
type ApiRegisteredIterator struct {
	Event *ApiRegistered // Event containing the contract specifics and raw log

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
func (it *ApiRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiRegistered)
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
		it.Event = new(ApiRegistered)
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
func (it *ApiRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiRegistered represents a Registered event raised by the Api contract.
type ApiRegistered struct {
	Owner   common.Address
	Payload [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0xb8142d42f05d95abf0a6570799774d59276e49ea32a04d9a4ec316ea4a6886bc.
//
// Solidity: event Registered(address indexed _owner, bytes32 _payload)
func (_Api *ApiFilterer) FilterRegistered(opts *bind.FilterOpts, _owner []common.Address) (*ApiRegisteredIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "Registered", _ownerRule)
	if err != nil {
		return nil, err
	}
	return &ApiRegisteredIterator{contract: _Api.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0xb8142d42f05d95abf0a6570799774d59276e49ea32a04d9a4ec316ea4a6886bc.
//
// Solidity: event Registered(address indexed _owner, bytes32 _payload)
func (_Api *ApiFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *ApiRegistered, _owner []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "Registered", _ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiRegistered)
				if err := _Api.contract.UnpackLog(event, "Registered", log); err != nil {
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

// ParseRegistered is a log parse operation binding the contract event 0xb8142d42f05d95abf0a6570799774d59276e49ea32a04d9a4ec316ea4a6886bc.
//
// Solidity: event Registered(address indexed _owner, bytes32 _payload)
func (_Api *ApiFilterer) ParseRegistered(log types.Log) (*ApiRegistered, error) {
	event := new(ApiRegistered)
	if err := _Api.contract.UnpackLog(event, "Registered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
