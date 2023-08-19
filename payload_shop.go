Fatal: Failed to generate ABI binding: 48:8: string literal not terminated (and 2 more errors)

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

import (
	"math/big"
	"strings"
	"errors"

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
	Owner common.Address
	Payload [32]byte
	}



	// StoreMetaData contains all meta data concerning the Store contract.
	var StoreMetaData = &bind.MetaData{
		ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_payload\",\"type\":\"bytes32\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_payload\",\"type\":\"bytes32\"}],\"name\":\"buy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getRecordsOfUser\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"payload\",\"type\":\"bytes32\"}],\"internalType\":\"structPayloadShop.Record[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"records\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"payload\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recordsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
		Bin: "0x======= contracts/PayloadShop.sol:PayloadShop =======
Binary:
60806040525f600255348015610013575f80fd5b506109af806100215f395ff3fe60806040526004361061003e575f3560e01c806334461067146100425780639c9a106114610080578063c4f82d7d146100b0578063eb3dbf83146100ec575b5f80fd5b34801561004d575f80fd5b5061006860048036038101906100639190610587565b610116565b60405161007793929190610623565b60405180910390f35b61009a60048036038101906100959190610682565b610166565b6040516100a791906106bc565b60405180910390f35b3480156100bb575f80fd5b506100d660048036038101906100d191906106ff565b6102fe565b6040516100e3919061083f565b60405180910390f35b3480156100f7575f80fd5b50610100610511565b60405161010d91906106bc565b60405180910390f35b5f602052805f5260405f205f91509050805f015f9054906101000a900460ff1690805f0160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154905083565b5f806001600254610177919061088c565b905060405180606001604052806001151581526020013373ffffffffffffffffffffffffffffffffffffffff168152602001848152505f808381526020019081526020015f205f820151815f015f6101000a81548160ff0219169083151502179055506020820151815f0160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816001015590505060015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f0181908060018154018082558091505060019003905f5260205f20015f909190919091505560025f8154809291906102a2906108bf565b91905055503373ffffffffffffffffffffffffffffffffffffffff167fb8142d42f05d95abf0a6570799774d59276e49ea32a04d9a4ec316ea4a6886bc846040516102ed9190610906565b60405180910390a280915050919050565b60605f60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f018054905067ffffffffffffffff81111561035d5761035c61091f565b5b60405190808252806020026020018201604052801561039657816020015b610383610517565b81526020019060019003908161037b5790505b5090505f5b60015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f0180549050811015610507575f805f60015f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f0184815481106104395761043861094c565b5b905f5260205f20015481526020019081526020015f209050806040518060600160405290815f82015f9054906101000a900460ff161515151581526020015f820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820154815250508383815181106104e8576104e761094c565b5b60200260200101819052505080806104ff906108bf565b91505061039b565b5080915050919050565b60025481565b60405180606001604052805f151581526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f80191681525090565b5f80fd5b5f819050919050565b61056681610554565b8114610570575f80fd5b50565b5f813590506105818161055d565b92915050565b5f6020828403121561059c5761059b610550565b5b5f6105a984828501610573565b91505092915050565b5f8115159050919050565b6105c6816105b2565b82525050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6105f5826105cc565b9050919050565b610605816105eb565b82525050565b5f819050919050565b61061d8161060b565b82525050565b5f6060820190506106365f8301866105bd565b61064360208301856105fc565b6106506040830184610614565b949350505050565b6106618161060b565b811461066b575f80fd5b50565b5f8135905061067c81610658565b92915050565b5f6020828403121561069757610696610550565b5b5f6106a48482850161066e565b91505092915050565b6106b681610554565b82525050565b5f6020820190506106cf5f8301846106ad565b92915050565b6106de816105eb565b81146106e8575f80fd5b50565b5f813590506106f9816106d5565b92915050565b5f6020828403121561071457610713610550565b5b5f610721848285016106eb565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b61075c816105b2565b82525050565b61076b816105eb565b82525050565b61077a8161060b565b82525050565b606082015f8201516107945f850182610753565b5060208201516107a76020850182610762565b5060408201516107ba6040850182610771565b50505050565b5f6107cb8383610780565b60608301905092915050565b5f602082019050919050565b5f6107ed8261072a565b6107f78185610734565b935061080283610744565b805f5b8381101561083257815161081988826107c0565b9750610824836107d7565b925050600181019050610805565b5085935050505092915050565b5f6020820190508181035f83015261085781846107e3565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61089682610554565b91506108a183610554565b92508282019050808211156108b9576108b861085f565b5b92915050565b5f6108c982610554565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036108fb576108fa61085f565b5b600182019050919050565b5f6020820190506109195f830184610614565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffdfea2646970667358221220074847d1b729f1ba6cb54fb83c7d81a88bf3c87d89588d011cb68199a231aded64736f6c63430008150033",
		
	}
	// StoreABI is the input ABI used to generate the binding from.
	// Deprecated: Use StoreMetaData.ABI instead.
	var StoreABI = StoreMetaData.ABI

	

	
		// StoreBin is the compiled bytecode used for deploying new contracts.
		// Deprecated: Use StoreMetaData.Bin instead.
		var StoreBin = StoreMetaData.Bin

		// DeployStore deploys a new Ethereum contract, binding an instance of Store to it.
		func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend ) (common.Address, *types.Transaction, *Store, error) {
		  parsed, err := StoreMetaData.GetAbi()
		  if err != nil {
		    return common.Address{}, nil, nil, err
		  }
		  if parsed == nil {
			return common.Address{}, nil, nil, errors.New("GetABI returned nil")
		  }
		  
		  address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StoreBin), backend )
		  if err != nil {
		    return common.Address{}, nil, nil, err
		  }
		  return address, tx, &Store{ StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract} }, nil
		}
	

	// Store is an auto generated Go binding around an Ethereum contract.
	type Store struct {
	  StoreCaller     // Read-only binding to the contract
	  StoreTransactor // Write-only binding to the contract
	  StoreFilterer   // Log filterer for contract events
	}

	// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
	type StoreCaller struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
	type StoreTransactor struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
	type StoreFilterer struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// StoreSession is an auto generated Go binding around an Ethereum contract,
	// with pre-set call and transact options.
	type StoreSession struct {
	  Contract     *Store        // Generic contract binding to set the session for
	  CallOpts     bind.CallOpts     // Call options to use throughout this session
	  TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
	}

	// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
	// with pre-set call options.
	type StoreCallerSession struct {
	  Contract *StoreCaller // Generic contract caller binding to set the session for
	  CallOpts bind.CallOpts    // Call options to use throughout this session
	}

	// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
	// with pre-set transact options.
	type StoreTransactorSession struct {
	  Contract     *StoreTransactor // Generic contract transactor binding to set the session for
	  TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
	}

	// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
	type StoreRaw struct {
	  Contract *Store // Generic contract binding to access the raw methods on
	}

	// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
	type StoreCallerRaw struct {
		Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
	}

	// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
	type StoreTransactorRaw struct {
		Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
	}

	// NewStore creates a new instance of Store, bound to a specific deployed contract.
	func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	  contract, err := bindStore(address, backend, backend, backend)
	  if err != nil {
	    return nil, err
	  }
	  return &Store{ StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract} }, nil
	}

	// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
	func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	  contract, err := bindStore(address, caller, nil, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &StoreCaller{contract: contract}, nil
	}

	// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
	func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	  contract, err := bindStore(address, nil, transactor, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &StoreTransactor{contract: contract}, nil
	}

	// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
 	func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
 	  contract, err := bindStore(address, nil, nil, filterer)
 	  if err != nil {
 	    return nil, err
 	  }
 	  return &StoreFilterer{contract: contract}, nil
 	}

	// bindStore binds a generic wrapper to an already deployed contract.
	func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	  parsed, err := StoreMetaData.GetAbi()
	  if err != nil {
	    return nil, err
	  }
	  return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_Store *StoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
		return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _Store.Contract.StoreTransactor.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
		return _Store.Contract.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _Store.Contract.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _Store.Contract.contract.Transact(opts, method, params...)
	}

	
		// GetRecordsOfUser is a free data retrieval call binding the contract method 0xc4f82d7d.
		//
		// Solidity: function getRecordsOfUser(address addr) view returns((bool,address,bytes32)[])
		func (_Store *StoreCaller) GetRecordsOfUser(opts *bind.CallOpts , addr common.Address ) ([]PayloadShopRecord, error) {
			var out []interface{}
			err := _Store.contract.Call(opts, &out, "getRecordsOfUser" , addr)
			
			if err != nil {
				return *new([]PayloadShopRecord),  err
			}
			
			out0 := *abi.ConvertType(out[0], new([]PayloadShopRecord)).(*[]PayloadShopRecord)
			
			return out0,  err
			
		}

		// GetRecordsOfUser is a free data retrieval call binding the contract method 0xc4f82d7d.
		//
		// Solidity: function getRecordsOfUser(address addr) view returns((bool,address,bytes32)[])
		func (_Store *StoreSession) GetRecordsOfUser( addr common.Address ) ( []PayloadShopRecord,  error) {
		  return _Store.Contract.GetRecordsOfUser(&_Store.CallOpts , addr)
		}

		// GetRecordsOfUser is a free data retrieval call binding the contract method 0xc4f82d7d.
		//
		// Solidity: function getRecordsOfUser(address addr) view returns((bool,address,bytes32)[])
		func (_Store *StoreCallerSession) GetRecordsOfUser( addr common.Address ) ( []PayloadShopRecord,  error) {
		  return _Store.Contract.GetRecordsOfUser(&_Store.CallOpts , addr)
		}
	
		// Records is a free data retrieval call binding the contract method 0x34461067.
		//
		// Solidity: function records(uint256 ) view returns(bool isRegistered, address owner, bytes32 payload)
		func (_Store *StoreCaller) Records(opts *bind.CallOpts , arg0 *big.Int ) (struct{ IsRegistered bool;Owner common.Address;Payload [32]byte; }, error) {
			var out []interface{}
			err := _Store.contract.Call(opts, &out, "records" , arg0)
			
			outstruct := new(struct{  IsRegistered bool;  Owner common.Address;  Payload [32]byte;  })
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
		func (_Store *StoreSession) Records( arg0 *big.Int ) (struct{ IsRegistered bool;Owner common.Address;Payload [32]byte; },  error) {
		  return _Store.Contract.Records(&_Store.CallOpts , arg0)
		}

		// Records is a free data retrieval call binding the contract method 0x34461067.
		//
		// Solidity: function records(uint256 ) view returns(bool isRegistered, address owner, bytes32 payload)
		func (_Store *StoreCallerSession) Records( arg0 *big.Int ) (struct{ IsRegistered bool;Owner common.Address;Payload [32]byte; },  error) {
		  return _Store.Contract.Records(&_Store.CallOpts , arg0)
		}
	
		// RecordsCount is a free data retrieval call binding the contract method 0xeb3dbf83.
		//
		// Solidity: function recordsCount() view returns(uint256)
		func (_Store *StoreCaller) RecordsCount(opts *bind.CallOpts ) (*big.Int, error) {
			var out []interface{}
			err := _Store.contract.Call(opts, &out, "recordsCount" )
			
			if err != nil {
				return *new(*big.Int),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
			
			return out0,  err
			
		}

		// RecordsCount is a free data retrieval call binding the contract method 0xeb3dbf83.
		//
		// Solidity: function recordsCount() view returns(uint256)
		func (_Store *StoreSession) RecordsCount() ( *big.Int,  error) {
		  return _Store.Contract.RecordsCount(&_Store.CallOpts )
		}

		// RecordsCount is a free data retrieval call binding the contract method 0xeb3dbf83.
		//
		// Solidity: function recordsCount() view returns(uint256)
		func (_Store *StoreCallerSession) RecordsCount() ( *big.Int,  error) {
		  return _Store.Contract.RecordsCount(&_Store.CallOpts )
		}
	

	
		// Buy is a paid mutator transaction binding the contract method 0x9c9a1061.
		//
		// Solidity: function buy(bytes32 _payload) payable returns(uint256)
		func (_Store *StoreTransactor) Buy(opts *bind.TransactOpts , _payload [32]byte ) (*types.Transaction, error) {
			return _Store.contract.Transact(opts, "buy" , _payload)
		}

		// Buy is a paid mutator transaction binding the contract method 0x9c9a1061.
		//
		// Solidity: function buy(bytes32 _payload) payable returns(uint256)
		func (_Store *StoreSession) Buy( _payload [32]byte ) (*types.Transaction, error) {
		  return _Store.Contract.Buy(&_Store.TransactOpts , _payload)
		}

		// Buy is a paid mutator transaction binding the contract method 0x9c9a1061.
		//
		// Solidity: function buy(bytes32 _payload) payable returns(uint256)
		func (_Store *StoreTransactorSession) Buy( _payload [32]byte ) (*types.Transaction, error) {
		  return _Store.Contract.Buy(&_Store.TransactOpts , _payload)
		}
	

	

	

	
		// StoreRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the Store contract.
		type StoreRegisteredIterator struct {
			Event *StoreRegistered // Event containing the contract specifics and raw log

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
		func (it *StoreRegisteredIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(StoreRegistered)
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
				it.Event = new(StoreRegistered)
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
		func (it *StoreRegisteredIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *StoreRegisteredIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// StoreRegistered represents a Registered event raised by the Store contract.
		type StoreRegistered struct { 
			Owner common.Address; 
			Payload [32]byte; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterRegistered is a free log retrieval operation binding the contract event 0xb8142d42f05d95abf0a6570799774d59276e49ea32a04d9a4ec316ea4a6886bc.
		//
		// Solidity: event Registered(address indexed _owner, bytes32 _payload)
 		func (_Store *StoreFilterer) FilterRegistered(opts *bind.FilterOpts, _owner []common.Address) (*StoreRegisteredIterator, error) {
			
			var _ownerRule []interface{}
			for _, _ownerItem := range _owner {
				_ownerRule = append(_ownerRule, _ownerItem)
			}
			

			logs, sub, err := _Store.contract.FilterLogs(opts, "Registered", _ownerRule)
			if err != nil {
				return nil, err
			}
			return &StoreRegisteredIterator{contract: _Store.contract, event: "Registered", logs: logs, sub: sub}, nil
 		}

		// WatchRegistered is a free log subscription operation binding the contract event 0xb8142d42f05d95abf0a6570799774d59276e49ea32a04d9a4ec316ea4a6886bc.
		//
		// Solidity: event Registered(address indexed _owner, bytes32 _payload)
		func (_Store *StoreFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *StoreRegistered, _owner []common.Address) (event.Subscription, error) {
			
			var _ownerRule []interface{}
			for _, _ownerItem := range _owner {
				_ownerRule = append(_ownerRule, _ownerItem)
			}
			

			logs, sub, err := _Store.contract.WatchLogs(opts, "Registered", _ownerRule)
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(StoreRegistered)
						if err := _Store.contract.UnpackLog(event, "Registered", log); err != nil {
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
		func (_Store *StoreFilterer) ParseRegistered(log types.Log) (*StoreRegistered, error) {
			event := new(StoreRegistered)
			if err := _Store.contract.UnpackLog(event, "Registered", log); err != nil {
				return nil, err
			}
			event.Raw = log
			return event, nil
		}

 	


