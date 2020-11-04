// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package accessControlContract

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AccessControlContractABI is the input ABI used to generate the binding from.
const AccessControlContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"newAddrRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"newAddrRemove\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"PubKeysKeystore\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"producerAddr\",\"type\":\"address\"}],\"name\":\"addAccountToRegister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"pubKey\",\"type\":\"string\"}],\"name\":\"addPubKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"adminPublicKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedAccounts\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"producerAddr\",\"type\":\"address\"}],\"name\":\"removeAccountFromRegister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"returnAllowedAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AccessControlContractFuncSigs maps the 4-byte function signature to its string representation.
var AccessControlContractFuncSigs = map[string]string{
	"489616fd": "PubKeysKeystore(address)",
	"c1a9e99a": "addAccountToRegister(address)",
	"670d65ea": "addPubKey(string)",
	"f851a440": "admin()",
	"be943a59": "adminPublicKey()",
	"e04610ed": "allowedAccounts(address)",
	"089a6e91": "removeAccountFromRegister(address)",
	"664ae384": "returnAllowedAddresses()",
}

// AccessControlContractBin is the compiled bytecode used for deploying new contracts.
var AccessControlContractBin = "0x6080604052600080546001600160a01b03191673647f089f75db1874e574419d20c34b078797c4c517905534801561003657600080fd5b506107fe806100466000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063be943a591161005b578063be943a591461024e578063c1a9e99a14610256578063e04610ed1461027c578063f851a440146102b657610088565b8063089a6e911461008d578063489616fd146100b5578063664ae38414610150578063670d65ea146101a8575b600080fd5b6100b3600480360360208110156100a357600080fd5b50356001600160a01b03166102da565b005b6100db600480360360208110156100cb57600080fd5b50356001600160a01b03166103d4565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101155781810151838201526020016100fd565b50505050905090810190601f1680156101425780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61015861046f565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561019457818101518382015260200161017c565b505050509050019250505060405180910390f35b6100b3600480360360208110156101be57600080fd5b8101906020810181356401000000008111156101d957600080fd5b8201836020820111156101eb57600080fd5b8035906020019184600183028401116401000000008311171561020d57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506104d2945050505050565b6100db610523565b6100b36004803603602081101561026c57600080fd5b50356001600160a01b031661057e565b6102a26004803603602081101561029257600080fd5b50356001600160a01b031661070d565b604080519115158252519081900360200190f35b6102be610722565b604080516001600160a01b039092168252519081900360200190f35b6000546001600160a01b031633146102f157600080fd5b6001600160a01b03811660009081526001602081905260409091205460ff1615151461031c57600080fd5b6001600160a01b0381166000908152600160209081526040808320805460ff191690556003909152902054600280548290811061035557fe5b6000918252602082200180546001600160a01b0319169055600480546001810182559082527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b018290556040516001600160a01b038416917fc80a30448ebca053eacc14b559ce6e66b3e353e7ba16587686d9c2ea98af7abe91a25050565b60056020908152600091825260409182902080548351601f6002600019610100600186161502019093169290920491820184900484028101840190945280845290918301828280156104675780601f1061043c57610100808354040283529160200191610467565b820191906000526020600020905b81548152906001019060200180831161044a57829003601f168201915b505050505081565b606060028054806020026020016040519081016040528092919081815260200182805480156104c757602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116104a9575b505050505090505b90565b6000546001600160a01b03163314156104fe5780516104f8906006906020840190610731565b50610520565b336000908152600560209081526040909120825161051e92840190610731565b505b50565b6006805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156104675780601f1061043c57610100808354040283529160200191610467565b6000546001600160a01b0316331461059557600080fd5b6001600160a01b03811660009081526001602052604090205460ff16156105bb57600080fd5b6001600160a01b0381166000908152600160208190526040909120805460ff191690911790556004548015610673576000600460018303815481106105fc57fe5b90600052602060002001549050826002828154811061061757fe5b600091825260208083209190910180546001600160a01b0319166001600160a01b03948516179055918516815260039091526040902081905560048054600019840190811061066257fe5b6000918252602082200155506106d5565b600280546001810182557f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace0180546001600160a01b0319166001600160a01b038516908117909155905460009182526003602052604090912060001990910190555b6040516001600160a01b038316907fa7cab406c53b3f8f59aa392a2950c6e96c73a82560b7d0da98a17a0929cbfe8490600090a25050565b60016020526000908152604090205460ff1681565b6000546001600160a01b031681565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061077257805160ff191683800117855561079f565b8280016001018555821561079f579182015b8281111561079f578251825591602001919060010190610784565b506107ab9291506107af565b5090565b6104cf91905b808211156107ab57600081556001016107b556fea265627a7a72315820477b38ffeedd8aa996a245fdee158ecd5b5ed6e32e34d66e890950e6b93592fd64736f6c63430005100032"

// DeployAccessControlContract deploys a new Ethereum contract, binding an instance of AccessControlContract to it.
func DeployAccessControlContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AccessControlContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccessControlContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AccessControlContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AccessControlContract{AccessControlContractCaller: AccessControlContractCaller{contract: contract}, AccessControlContractTransactor: AccessControlContractTransactor{contract: contract}, AccessControlContractFilterer: AccessControlContractFilterer{contract: contract}}, nil
}

// AccessControlContract is an auto generated Go binding around an Ethereum contract.
type AccessControlContract struct {
	AccessControlContractCaller     // Read-only binding to the contract
	AccessControlContractTransactor // Write-only binding to the contract
	AccessControlContractFilterer   // Log filterer for contract events
}

// AccessControlContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccessControlContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccessControlContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccessControlContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccessControlContractSession struct {
	Contract     *AccessControlContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AccessControlContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccessControlContractCallerSession struct {
	Contract *AccessControlContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// AccessControlContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccessControlContractTransactorSession struct {
	Contract     *AccessControlContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// AccessControlContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccessControlContractRaw struct {
	Contract *AccessControlContract // Generic contract binding to access the raw methods on
}

// AccessControlContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccessControlContractCallerRaw struct {
	Contract *AccessControlContractCaller // Generic read-only contract binding to access the raw methods on
}

// AccessControlContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccessControlContractTransactorRaw struct {
	Contract *AccessControlContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccessControlContract creates a new instance of AccessControlContract, bound to a specific deployed contract.
func NewAccessControlContract(address common.Address, backend bind.ContractBackend) (*AccessControlContract, error) {
	contract, err := bindAccessControlContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccessControlContract{AccessControlContractCaller: AccessControlContractCaller{contract: contract}, AccessControlContractTransactor: AccessControlContractTransactor{contract: contract}, AccessControlContractFilterer: AccessControlContractFilterer{contract: contract}}, nil
}

// NewAccessControlContractCaller creates a new read-only instance of AccessControlContract, bound to a specific deployed contract.
func NewAccessControlContractCaller(address common.Address, caller bind.ContractCaller) (*AccessControlContractCaller, error) {
	contract, err := bindAccessControlContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlContractCaller{contract: contract}, nil
}

// NewAccessControlContractTransactor creates a new write-only instance of AccessControlContract, bound to a specific deployed contract.
func NewAccessControlContractTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessControlContractTransactor, error) {
	contract, err := bindAccessControlContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlContractTransactor{contract: contract}, nil
}

// NewAccessControlContractFilterer creates a new log filterer instance of AccessControlContract, bound to a specific deployed contract.
func NewAccessControlContractFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessControlContractFilterer, error) {
	contract, err := bindAccessControlContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessControlContractFilterer{contract: contract}, nil
}

// bindAccessControlContract binds a generic wrapper to an already deployed contract.
func bindAccessControlContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccessControlContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControlContract *AccessControlContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControlContract.Contract.AccessControlContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControlContract *AccessControlContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControlContract.Contract.AccessControlContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControlContract *AccessControlContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControlContract.Contract.AccessControlContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControlContract *AccessControlContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControlContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControlContract *AccessControlContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControlContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControlContract *AccessControlContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControlContract.Contract.contract.Transact(opts, method, params...)
}

// PubKeysKeystore is a free data retrieval call binding the contract method 0x489616fd.
//
// Solidity: function PubKeysKeystore(address ) view returns(string)
func (_AccessControlContract *AccessControlContractCaller) PubKeysKeystore(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _AccessControlContract.contract.Call(opts, &out, "PubKeysKeystore", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PubKeysKeystore is a free data retrieval call binding the contract method 0x489616fd.
//
// Solidity: function PubKeysKeystore(address ) view returns(string)
func (_AccessControlContract *AccessControlContractSession) PubKeysKeystore(arg0 common.Address) (string, error) {
	return _AccessControlContract.Contract.PubKeysKeystore(&_AccessControlContract.CallOpts, arg0)
}

// PubKeysKeystore is a free data retrieval call binding the contract method 0x489616fd.
//
// Solidity: function PubKeysKeystore(address ) view returns(string)
func (_AccessControlContract *AccessControlContractCallerSession) PubKeysKeystore(arg0 common.Address) (string, error) {
	return _AccessControlContract.Contract.PubKeysKeystore(&_AccessControlContract.CallOpts, arg0)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_AccessControlContract *AccessControlContractCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccessControlContract.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_AccessControlContract *AccessControlContractSession) Admin() (common.Address, error) {
	return _AccessControlContract.Contract.Admin(&_AccessControlContract.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_AccessControlContract *AccessControlContractCallerSession) Admin() (common.Address, error) {
	return _AccessControlContract.Contract.Admin(&_AccessControlContract.CallOpts)
}

// AdminPublicKey is a free data retrieval call binding the contract method 0xbe943a59.
//
// Solidity: function adminPublicKey() view returns(string)
func (_AccessControlContract *AccessControlContractCaller) AdminPublicKey(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AccessControlContract.contract.Call(opts, &out, "adminPublicKey")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AdminPublicKey is a free data retrieval call binding the contract method 0xbe943a59.
//
// Solidity: function adminPublicKey() view returns(string)
func (_AccessControlContract *AccessControlContractSession) AdminPublicKey() (string, error) {
	return _AccessControlContract.Contract.AdminPublicKey(&_AccessControlContract.CallOpts)
}

// AdminPublicKey is a free data retrieval call binding the contract method 0xbe943a59.
//
// Solidity: function adminPublicKey() view returns(string)
func (_AccessControlContract *AccessControlContractCallerSession) AdminPublicKey() (string, error) {
	return _AccessControlContract.Contract.AdminPublicKey(&_AccessControlContract.CallOpts)
}

// AllowedAccounts is a free data retrieval call binding the contract method 0xe04610ed.
//
// Solidity: function allowedAccounts(address ) view returns(bool)
func (_AccessControlContract *AccessControlContractCaller) AllowedAccounts(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _AccessControlContract.contract.Call(opts, &out, "allowedAccounts", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedAccounts is a free data retrieval call binding the contract method 0xe04610ed.
//
// Solidity: function allowedAccounts(address ) view returns(bool)
func (_AccessControlContract *AccessControlContractSession) AllowedAccounts(arg0 common.Address) (bool, error) {
	return _AccessControlContract.Contract.AllowedAccounts(&_AccessControlContract.CallOpts, arg0)
}

// AllowedAccounts is a free data retrieval call binding the contract method 0xe04610ed.
//
// Solidity: function allowedAccounts(address ) view returns(bool)
func (_AccessControlContract *AccessControlContractCallerSession) AllowedAccounts(arg0 common.Address) (bool, error) {
	return _AccessControlContract.Contract.AllowedAccounts(&_AccessControlContract.CallOpts, arg0)
}

// ReturnAllowedAddresses is a free data retrieval call binding the contract method 0x664ae384.
//
// Solidity: function returnAllowedAddresses() view returns(address[])
func (_AccessControlContract *AccessControlContractCaller) ReturnAllowedAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _AccessControlContract.contract.Call(opts, &out, "returnAllowedAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ReturnAllowedAddresses is a free data retrieval call binding the contract method 0x664ae384.
//
// Solidity: function returnAllowedAddresses() view returns(address[])
func (_AccessControlContract *AccessControlContractSession) ReturnAllowedAddresses() ([]common.Address, error) {
	return _AccessControlContract.Contract.ReturnAllowedAddresses(&_AccessControlContract.CallOpts)
}

// ReturnAllowedAddresses is a free data retrieval call binding the contract method 0x664ae384.
//
// Solidity: function returnAllowedAddresses() view returns(address[])
func (_AccessControlContract *AccessControlContractCallerSession) ReturnAllowedAddresses() ([]common.Address, error) {
	return _AccessControlContract.Contract.ReturnAllowedAddresses(&_AccessControlContract.CallOpts)
}

// AddAccountToRegister is a paid mutator transaction binding the contract method 0xc1a9e99a.
//
// Solidity: function addAccountToRegister(address producerAddr) returns()
func (_AccessControlContract *AccessControlContractTransactor) AddAccountToRegister(opts *bind.TransactOpts, producerAddr common.Address) (*types.Transaction, error) {
	return _AccessControlContract.contract.Transact(opts, "addAccountToRegister", producerAddr)
}

// AddAccountToRegister is a paid mutator transaction binding the contract method 0xc1a9e99a.
//
// Solidity: function addAccountToRegister(address producerAddr) returns()
func (_AccessControlContract *AccessControlContractSession) AddAccountToRegister(producerAddr common.Address) (*types.Transaction, error) {
	return _AccessControlContract.Contract.AddAccountToRegister(&_AccessControlContract.TransactOpts, producerAddr)
}

// AddAccountToRegister is a paid mutator transaction binding the contract method 0xc1a9e99a.
//
// Solidity: function addAccountToRegister(address producerAddr) returns()
func (_AccessControlContract *AccessControlContractTransactorSession) AddAccountToRegister(producerAddr common.Address) (*types.Transaction, error) {
	return _AccessControlContract.Contract.AddAccountToRegister(&_AccessControlContract.TransactOpts, producerAddr)
}

// AddPubKey is a paid mutator transaction binding the contract method 0x670d65ea.
//
// Solidity: function addPubKey(string pubKey) returns()
func (_AccessControlContract *AccessControlContractTransactor) AddPubKey(opts *bind.TransactOpts, pubKey string) (*types.Transaction, error) {
	return _AccessControlContract.contract.Transact(opts, "addPubKey", pubKey)
}

// AddPubKey is a paid mutator transaction binding the contract method 0x670d65ea.
//
// Solidity: function addPubKey(string pubKey) returns()
func (_AccessControlContract *AccessControlContractSession) AddPubKey(pubKey string) (*types.Transaction, error) {
	return _AccessControlContract.Contract.AddPubKey(&_AccessControlContract.TransactOpts, pubKey)
}

// AddPubKey is a paid mutator transaction binding the contract method 0x670d65ea.
//
// Solidity: function addPubKey(string pubKey) returns()
func (_AccessControlContract *AccessControlContractTransactorSession) AddPubKey(pubKey string) (*types.Transaction, error) {
	return _AccessControlContract.Contract.AddPubKey(&_AccessControlContract.TransactOpts, pubKey)
}

// RemoveAccountFromRegister is a paid mutator transaction binding the contract method 0x089a6e91.
//
// Solidity: function removeAccountFromRegister(address producerAddr) returns()
func (_AccessControlContract *AccessControlContractTransactor) RemoveAccountFromRegister(opts *bind.TransactOpts, producerAddr common.Address) (*types.Transaction, error) {
	return _AccessControlContract.contract.Transact(opts, "removeAccountFromRegister", producerAddr)
}

// RemoveAccountFromRegister is a paid mutator transaction binding the contract method 0x089a6e91.
//
// Solidity: function removeAccountFromRegister(address producerAddr) returns()
func (_AccessControlContract *AccessControlContractSession) RemoveAccountFromRegister(producerAddr common.Address) (*types.Transaction, error) {
	return _AccessControlContract.Contract.RemoveAccountFromRegister(&_AccessControlContract.TransactOpts, producerAddr)
}

// RemoveAccountFromRegister is a paid mutator transaction binding the contract method 0x089a6e91.
//
// Solidity: function removeAccountFromRegister(address producerAddr) returns()
func (_AccessControlContract *AccessControlContractTransactorSession) RemoveAccountFromRegister(producerAddr common.Address) (*types.Transaction, error) {
	return _AccessControlContract.Contract.RemoveAccountFromRegister(&_AccessControlContract.TransactOpts, producerAddr)
}

// AccessControlContractNewAddrRegisteredIterator is returned from FilterNewAddrRegistered and is used to iterate over the raw logs and unpacked data for NewAddrRegistered events raised by the AccessControlContract contract.
type AccessControlContractNewAddrRegisteredIterator struct {
	Event *AccessControlContractNewAddrRegistered // Event containing the contract specifics and raw log

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
func (it *AccessControlContractNewAddrRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlContractNewAddrRegistered)
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
		it.Event = new(AccessControlContractNewAddrRegistered)
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
func (it *AccessControlContractNewAddrRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlContractNewAddrRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlContractNewAddrRegistered represents a NewAddrRegistered event raised by the AccessControlContract contract.
type AccessControlContractNewAddrRegistered struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNewAddrRegistered is a free log retrieval operation binding the contract event 0xa7cab406c53b3f8f59aa392a2950c6e96c73a82560b7d0da98a17a0929cbfe84.
//
// Solidity: event newAddrRegistered(address indexed _addr)
func (_AccessControlContract *AccessControlContractFilterer) FilterNewAddrRegistered(opts *bind.FilterOpts, _addr []common.Address) (*AccessControlContractNewAddrRegisteredIterator, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}

	logs, sub, err := _AccessControlContract.contract.FilterLogs(opts, "newAddrRegistered", _addrRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlContractNewAddrRegisteredIterator{contract: _AccessControlContract.contract, event: "newAddrRegistered", logs: logs, sub: sub}, nil
}

// WatchNewAddrRegistered is a free log subscription operation binding the contract event 0xa7cab406c53b3f8f59aa392a2950c6e96c73a82560b7d0da98a17a0929cbfe84.
//
// Solidity: event newAddrRegistered(address indexed _addr)
func (_AccessControlContract *AccessControlContractFilterer) WatchNewAddrRegistered(opts *bind.WatchOpts, sink chan<- *AccessControlContractNewAddrRegistered, _addr []common.Address) (event.Subscription, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}

	logs, sub, err := _AccessControlContract.contract.WatchLogs(opts, "newAddrRegistered", _addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlContractNewAddrRegistered)
				if err := _AccessControlContract.contract.UnpackLog(event, "newAddrRegistered", log); err != nil {
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

// ParseNewAddrRegistered is a log parse operation binding the contract event 0xa7cab406c53b3f8f59aa392a2950c6e96c73a82560b7d0da98a17a0929cbfe84.
//
// Solidity: event newAddrRegistered(address indexed _addr)
func (_AccessControlContract *AccessControlContractFilterer) ParseNewAddrRegistered(log types.Log) (*AccessControlContractNewAddrRegistered, error) {
	event := new(AccessControlContractNewAddrRegistered)
	if err := _AccessControlContract.contract.UnpackLog(event, "newAddrRegistered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccessControlContractNewAddrRemoveIterator is returned from FilterNewAddrRemove and is used to iterate over the raw logs and unpacked data for NewAddrRemove events raised by the AccessControlContract contract.
type AccessControlContractNewAddrRemoveIterator struct {
	Event *AccessControlContractNewAddrRemove // Event containing the contract specifics and raw log

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
func (it *AccessControlContractNewAddrRemoveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlContractNewAddrRemove)
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
		it.Event = new(AccessControlContractNewAddrRemove)
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
func (it *AccessControlContractNewAddrRemoveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlContractNewAddrRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlContractNewAddrRemove represents a NewAddrRemove event raised by the AccessControlContract contract.
type AccessControlContractNewAddrRemove struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNewAddrRemove is a free log retrieval operation binding the contract event 0xc80a30448ebca053eacc14b559ce6e66b3e353e7ba16587686d9c2ea98af7abe.
//
// Solidity: event newAddrRemove(address indexed _addr)
func (_AccessControlContract *AccessControlContractFilterer) FilterNewAddrRemove(opts *bind.FilterOpts, _addr []common.Address) (*AccessControlContractNewAddrRemoveIterator, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}

	logs, sub, err := _AccessControlContract.contract.FilterLogs(opts, "newAddrRemove", _addrRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlContractNewAddrRemoveIterator{contract: _AccessControlContract.contract, event: "newAddrRemove", logs: logs, sub: sub}, nil
}

// WatchNewAddrRemove is a free log subscription operation binding the contract event 0xc80a30448ebca053eacc14b559ce6e66b3e353e7ba16587686d9c2ea98af7abe.
//
// Solidity: event newAddrRemove(address indexed _addr)
func (_AccessControlContract *AccessControlContractFilterer) WatchNewAddrRemove(opts *bind.WatchOpts, sink chan<- *AccessControlContractNewAddrRemove, _addr []common.Address) (event.Subscription, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}

	logs, sub, err := _AccessControlContract.contract.WatchLogs(opts, "newAddrRemove", _addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlContractNewAddrRemove)
				if err := _AccessControlContract.contract.UnpackLog(event, "newAddrRemove", log); err != nil {
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

// ParseNewAddrRemove is a log parse operation binding the contract event 0xc80a30448ebca053eacc14b559ce6e66b3e353e7ba16587686d9c2ea98af7abe.
//
// Solidity: event newAddrRemove(address indexed _addr)
func (_AccessControlContract *AccessControlContractFilterer) ParseNewAddrRemove(log types.Log) (*AccessControlContractNewAddrRemove, error) {
	event := new(AccessControlContractNewAddrRemove)
	if err := _AccessControlContract.contract.UnpackLog(event, "newAddrRemove", log); err != nil {
		return nil, err
	}
	return event, nil
}
