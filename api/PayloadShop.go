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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_payload\",\"type\":\"bytes32\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_payload\",\"type\":\"bytes32\"}],\"name\":\"buy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getRecordsOfUser\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"payload\",\"type\":\"bytes32\"}],\"internalType\":\"structPayloadShop.Record[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"records\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"payload\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recordsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"withdrawAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040525f6002555f600455348015610017575f80fd5b50600380546001600160a01b031916331790556106d2806100375f395ff3fe608060405260043610610084575f3560e01c80639c9a1061116100575780639c9a10611461016c5780639d1b464a1461017f578063c4f82d7d14610194578063eb3dbf83146101c0578063fa09e630146101d5575f80fd5b806312065fe01461008857806334461067146100a95780638da5cb5b1461011457806391b7f5ed1461014b575b5f80fd5b348015610093575f80fd5b50475b6040519081526020015b60405180910390f35b3480156100b4575f80fd5b506100ef6100c336600461057c565b5f602081905290815260409020805460019091015460ff82169161010090046001600160a01b03169083565b6040805193151584526001600160a01b039092166020840152908201526060016100a0565b34801561011f575f80fd5b50600354610133906001600160a01b031681565b6040516001600160a01b0390911681526020016100a0565b348015610156575f80fd5b5061016a61016536600461057c565b6101f4565b005b61009661017a36600461057c565b610258565b34801561018a575f80fd5b5061009660045481565b34801561019f575f80fd5b506101b36101ae3660046105aa565b610390565b6040516100a091906105cc565b3480156101cb575f80fd5b5061009660025481565b3480156101e0575f80fd5b5061016a6101ef3660046105aa565b6104ec565b6003546001600160a01b031633146102535760405162461bcd60e51b815260206004820152601a60248201527f2d2d2d3a434552523a4143434553535f44454e4945443a2d2d2d00000000000060448201526064015b60405180910390fd5b600455565b5f6004543410156102ab5760405162461bcd60e51b815260206004820152601860248201527f2d2d2d3a434552523a57524f4e475f50524943453a2d2d2d0000000000000000604482015260640161024a565b5f60025460016102bb9190610643565b6040805160608101825260018082523360208084018281528486018a81525f8881528084528781209651875493516001600160a81b0319909416901515610100600160a81b031916176101006001600160a01b0390941693909302929092178655519484019490945590835281815292822080549182018155825291812090910182905560028054929350906103508361065c565b909155505060405183815233907fb8142d42f05d95abf0a6570799774d59276e49ea32a04d9a4ec316ea4a6886bc9060200160405180910390a292915050565b6001600160a01b0381165f908152600160205260408120546060919067ffffffffffffffff8111156103c4576103c4610674565b60405190808252806020026020018201604052801561040d57816020015b604080516060810182525f80825260208083018290529282015282525f199092019101816103e25790505b5090505f5b6001600160a01b0384165f908152600160205260409020548110156104e5576001600160a01b0384165f9081526001602052604081208054829182918590811061045e5761045e610688565b5f918252602080832090910154835282810193909352604091820190208151606081018352815460ff8116151582526001600160a01b03610100909104169381019390935260018101549183019190915284519092508490849081106104c6576104c6610688565b60200260200101819052505080806104dd9061065c565b915050610412565b5092915050565b6003546001600160a01b031633146105465760405162461bcd60e51b815260206004820152601a60248201527f2d2d2d3a434552523a4143434553535f44454e4945443a2d2d2d000000000000604482015260640161024a565b6040516001600160a01b038216904780156108fc02915f818181858888f19350505050158015610578573d5f803e3d5ffd5b5050565b5f6020828403121561058c575f80fd5b5035919050565b6001600160a01b03811681146105a7575f80fd5b50565b5f602082840312156105ba575f80fd5b81356105c581610593565b9392505050565b602080825282518282018190525f919060409081850190868401855b82811015610622578151805115158552868101516001600160a01b03168786015285015185850152606090930192908501906001016105e8565b5091979650505050505050565b634e487b7160e01b5f52601160045260245ffd5b808201808211156106565761065661062f565b92915050565b5f6001820161066d5761066d61062f565b5060010190565b634e487b7160e01b5f52604160045260245ffd5b634e487b7160e01b5f52603260045260245ffdfea26469706673582212201e72ec265515b439af53966259b4505687d4a13572ba4ad7e2e7cb3819315c1464736f6c63430008150033",
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

// CurrentPrice is a free data retrieval call binding the contract method 0x9d1b464a.
//
// Solidity: function currentPrice() view returns(uint256)
func (_Api *ApiCaller) CurrentPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "currentPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentPrice is a free data retrieval call binding the contract method 0x9d1b464a.
//
// Solidity: function currentPrice() view returns(uint256)
func (_Api *ApiSession) CurrentPrice() (*big.Int, error) {
	return _Api.Contract.CurrentPrice(&_Api.CallOpts)
}

// CurrentPrice is a free data retrieval call binding the contract method 0x9d1b464a.
//
// Solidity: function currentPrice() view returns(uint256)
func (_Api *ApiCallerSession) CurrentPrice() (*big.Int, error) {
	return _Api.Contract.CurrentPrice(&_Api.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Api *ApiCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Api *ApiSession) GetBalance() (*big.Int, error) {
	return _Api.Contract.GetBalance(&_Api.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Api *ApiCallerSession) GetBalance() (*big.Int, error) {
	return _Api.Contract.GetBalance(&_Api.CallOpts)
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

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Api *ApiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Api *ApiSession) Owner() (common.Address, error) {
	return _Api.Contract.Owner(&_Api.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Api *ApiCallerSession) Owner() (common.Address, error) {
	return _Api.Contract.Owner(&_Api.CallOpts)
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

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 price) returns()
func (_Api *ApiTransactor) SetPrice(opts *bind.TransactOpts, price *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setPrice", price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 price) returns()
func (_Api *ApiSession) SetPrice(price *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SetPrice(&_Api.TransactOpts, price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 price) returns()
func (_Api *ApiTransactorSession) SetPrice(price *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SetPrice(&_Api.TransactOpts, price)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0xfa09e630.
//
// Solidity: function withdrawAll(address _to) returns()
func (_Api *ApiTransactor) WithdrawAll(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "withdrawAll", _to)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0xfa09e630.
//
// Solidity: function withdrawAll(address _to) returns()
func (_Api *ApiSession) WithdrawAll(_to common.Address) (*types.Transaction, error) {
	return _Api.Contract.WithdrawAll(&_Api.TransactOpts, _to)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0xfa09e630.
//
// Solidity: function withdrawAll(address _to) returns()
func (_Api *ApiTransactorSession) WithdrawAll(_to common.Address) (*types.Transaction, error) {
	return _Api.Contract.WithdrawAll(&_Api.TransactOpts, _to)
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
