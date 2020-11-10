// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dataContract

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
const AccessControlContractABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedAccounts\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AccessControlContractFuncSigs maps the 4-byte function signature to its string representation.
var AccessControlContractFuncSigs = map[string]string{
	"e04610ed": "allowedAccounts(address)",
}

// AccessControlContractBin is the compiled bytecode used for deploying new contracts.
var AccessControlContractBin = "0x6080604052348015600f57600080fd5b5060ae8061001e6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063e04610ed14602d575b600080fd5b605060048036036020811015604157600080fd5b50356001600160a01b03166064565b604080519115158252519081900360200190f35b60006020819052908152604090205460ff168156fea265627a7a7231582024235e4ac188c2b8fd32973e19d03c220741b499c6f50d1cff5879511074828864736f6c63430005100032"

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

// DataLedgerContractABI is the input ABI used to generate the binding from.
const DataLedgerContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"deleteInfo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"}],\"name\":\"evtStoreInfo\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"deleteMeasurement\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getIoTAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ledger\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"storeInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DataLedgerContractFuncSigs maps the 4-byte function signature to its string representation.
var DataLedgerContractFuncSigs = map[string]string{
	"77ad95ca": "deleteMeasurement(bytes32)",
	"6ade0219": "getIoTAddress(bytes32)",
	"15977d45": "ledger(bytes32)",
	"e30081a0": "setAddress(address)",
	"b7e2a1b8": "storeInfo(bytes32,string,string)",
}

// DataLedgerContractBin is the compiled bytecode used for deploying new contracts.
var DataLedgerContractBin = "0x6080604052600180546001600160a01b03191673647f089f75db1874e574419d20c34b078797c4c517905534801561003657600080fd5b506109c4806100466000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c806315977d451461005c5780636ade02191461017057806377ad95ca146101a9578063b7e2a1b8146101c8578063e30081a0146102fc575b600080fd5b6100796004803603602081101561007257600080fd5b5035610322565b604051808060200180602001846001600160a01b03166001600160a01b03168152602001838103835286818151815260200191508051906020019080838360005b838110156100d25781810151838201526020016100ba565b50505050905090810190601f1680156100ff5780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b8381101561013257818101518382015260200161011a565b50505050905090810190601f16801561015f5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b61018d6004803603602081101561018657600080fd5b503561046a565b604080516001600160a01b039092168252519081900360200190f35b6101c6600480360360208110156101bf57600080fd5b5035610489565b005b6101c6600480360360608110156101de57600080fd5b8135919081019060408101602082013564010000000081111561020057600080fd5b82018360208201111561021257600080fd5b8035906020019184600183028401116401000000008311171561023457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561028757600080fd5b82018360208201111561029957600080fd5b803590602001918460018302840111640100000000831117156102bb57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610539945050505050565b6101c66004803603602081101561031257600080fd5b50356001600160a01b031661070e565b600260208181526000928352604092839020805484516001821615610100026000190190911693909304601f81018390048302840183019094528383529283918301828280156103b35780601f10610388576101008083540402835291602001916103b3565b820191906000526020600020905b81548152906001019060200180831161039657829003601f168201915b505050505090806001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104515780601f1061042657610100808354040283529160200191610451565b820191906000526020600020905b81548152906001019060200180831161043457829003601f168201915b505050600290930154919250506001600160a01b031683565b600090815260026020819052604090912001546001600160a01b031690565b6001546001600160a01b031633146104d25760405162461bcd60e51b81526004018080602001828103825260338152602001806109066033913960400191505060405180910390fd5b6000818152600260205260408120906104eb82826107f9565b6104f96001830160006107f9565b5060020180546001600160a01b031916905560405181907f072007d551e16de6c1b8938fdd0559f70033d87037e5dffa28631256df69f9fe90600090a250565b61054233610779565b15156001146105825760405162461bcd60e51b815260040180806020018281038252602b815260200180610965602b913960400191505060405180910390fd5b61058a610840565b828152602080820183905233604080840191909152600086815260028352208251805184936105bd92849291019061086a565b5060208281015180516105d6926001850192019061086a565b5060408201518160020160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550905050837fafeab48c124d8423d588c0406f25d8386c21751c23cdf11491137a519343ec228484604051808060200180602001838103835285818151815260200191508051906020019080838360005b8381101561066c578181015183820152602001610654565b50505050905090810190601f1680156106995780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b838110156106cc5781810151838201526020016106b4565b50505050905090810190601f1680156106f95780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a250505050565b6001546001600160a01b031633146107575760405162461bcd60e51b815260040180806020018281038252602c815260200180610939602c913960400191505060405180910390fd5b600080546001600160a01b0319166001600160a01b0392909216919091179055565b600080546040805163e04610ed60e01b81526001600160a01b0385811660048301529151919092169163e04610ed916024808301926020929190829003018186803b1580156107c757600080fd5b505afa1580156107db573d6000803e3d6000fd5b505050506040513d60208110156107f157600080fd5b505192915050565b50805460018160011615610100020316600290046000825580601f1061081f575061083d565b601f01602090049060005260206000209081019061083d91906108e8565b50565b6040518060600160405280606081526020016060815260200160006001600160a01b031681525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106108ab57805160ff19168380011785556108d8565b828001600101855582156108d8579182015b828111156108d85782518255916020019190600101906108bd565b506108e49291506108e8565b5090565b61090291905b808211156108e457600081556001016108ee565b9056fe596f7520646f206e6f74206861766520656e6f7567682070726976696c6567657320746f20646f207468697320616374696f6e596f7520646f206e6f7420686176652070726976696c6567657320746f20646f207468697320616374696f6e546865204944207468617420796f7520617265207573696e67206973206e6f742072656769737465726564a265627a7a7231582077ef5bdb589b78dcda372229e743bb61782b2ac839e2f2b82bb59b3fa14c7aab64736f6c63430005100032"

// DeployDataLedgerContract deploys a new Ethereum contract, binding an instance of DataLedgerContract to it.
func DeployDataLedgerContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DataLedgerContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DataLedgerContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DataLedgerContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DataLedgerContract{DataLedgerContractCaller: DataLedgerContractCaller{contract: contract}, DataLedgerContractTransactor: DataLedgerContractTransactor{contract: contract}, DataLedgerContractFilterer: DataLedgerContractFilterer{contract: contract}}, nil
}

// DataLedgerContract is an auto generated Go binding around an Ethereum contract.
type DataLedgerContract struct {
	DataLedgerContractCaller     // Read-only binding to the contract
	DataLedgerContractTransactor // Write-only binding to the contract
	DataLedgerContractFilterer   // Log filterer for contract events
}

// DataLedgerContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type DataLedgerContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataLedgerContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DataLedgerContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataLedgerContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DataLedgerContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataLedgerContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DataLedgerContractSession struct {
	Contract     *DataLedgerContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DataLedgerContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DataLedgerContractCallerSession struct {
	Contract *DataLedgerContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// DataLedgerContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DataLedgerContractTransactorSession struct {
	Contract     *DataLedgerContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// DataLedgerContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type DataLedgerContractRaw struct {
	Contract *DataLedgerContract // Generic contract binding to access the raw methods on
}

// DataLedgerContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DataLedgerContractCallerRaw struct {
	Contract *DataLedgerContractCaller // Generic read-only contract binding to access the raw methods on
}

// DataLedgerContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DataLedgerContractTransactorRaw struct {
	Contract *DataLedgerContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDataLedgerContract creates a new instance of DataLedgerContract, bound to a specific deployed contract.
func NewDataLedgerContract(address common.Address, backend bind.ContractBackend) (*DataLedgerContract, error) {
	contract, err := bindDataLedgerContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DataLedgerContract{DataLedgerContractCaller: DataLedgerContractCaller{contract: contract}, DataLedgerContractTransactor: DataLedgerContractTransactor{contract: contract}, DataLedgerContractFilterer: DataLedgerContractFilterer{contract: contract}}, nil
}

// NewDataLedgerContractCaller creates a new read-only instance of DataLedgerContract, bound to a specific deployed contract.
func NewDataLedgerContractCaller(address common.Address, caller bind.ContractCaller) (*DataLedgerContractCaller, error) {
	contract, err := bindDataLedgerContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DataLedgerContractCaller{contract: contract}, nil
}

// NewDataLedgerContractTransactor creates a new write-only instance of DataLedgerContract, bound to a specific deployed contract.
func NewDataLedgerContractTransactor(address common.Address, transactor bind.ContractTransactor) (*DataLedgerContractTransactor, error) {
	contract, err := bindDataLedgerContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DataLedgerContractTransactor{contract: contract}, nil
}

// NewDataLedgerContractFilterer creates a new log filterer instance of DataLedgerContract, bound to a specific deployed contract.
func NewDataLedgerContractFilterer(address common.Address, filterer bind.ContractFilterer) (*DataLedgerContractFilterer, error) {
	contract, err := bindDataLedgerContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DataLedgerContractFilterer{contract: contract}, nil
}

// bindDataLedgerContract binds a generic wrapper to an already deployed contract.
func bindDataLedgerContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DataLedgerContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataLedgerContract *DataLedgerContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataLedgerContract.Contract.DataLedgerContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataLedgerContract *DataLedgerContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.DataLedgerContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataLedgerContract *DataLedgerContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.DataLedgerContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataLedgerContract *DataLedgerContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataLedgerContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataLedgerContract *DataLedgerContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataLedgerContract *DataLedgerContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.contract.Transact(opts, method, params...)
}

// GetIoTAddress is a free data retrieval call binding the contract method 0x6ade0219.
//
// Solidity: function getIoTAddress(bytes32 hash) view returns(address)
func (_DataLedgerContract *DataLedgerContractCaller) GetIoTAddress(opts *bind.CallOpts, hash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _DataLedgerContract.contract.Call(opts, &out, "getIoTAddress", hash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetIoTAddress is a free data retrieval call binding the contract method 0x6ade0219.
//
// Solidity: function getIoTAddress(bytes32 hash) view returns(address)
func (_DataLedgerContract *DataLedgerContractSession) GetIoTAddress(hash [32]byte) (common.Address, error) {
	return _DataLedgerContract.Contract.GetIoTAddress(&_DataLedgerContract.CallOpts, hash)
}

// GetIoTAddress is a free data retrieval call binding the contract method 0x6ade0219.
//
// Solidity: function getIoTAddress(bytes32 hash) view returns(address)
func (_DataLedgerContract *DataLedgerContractCallerSession) GetIoTAddress(hash [32]byte) (common.Address, error) {
	return _DataLedgerContract.Contract.GetIoTAddress(&_DataLedgerContract.CallOpts, hash)
}

// Ledger is a free data retrieval call binding the contract method 0x15977d45.
//
// Solidity: function ledger(bytes32 ) view returns(string uri, string description, address addr)
func (_DataLedgerContract *DataLedgerContractCaller) Ledger(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Uri         string
	Description string
	Addr        common.Address
}, error) {
	var out []interface{}
	err := _DataLedgerContract.contract.Call(opts, &out, "ledger", arg0)

	outstruct := new(struct {
		Uri         string
		Description string
		Addr        common.Address
	})

	outstruct.Uri = out[0].(string)
	outstruct.Description = out[1].(string)
	outstruct.Addr = out[2].(common.Address)

	return *outstruct, err

}

// Ledger is a free data retrieval call binding the contract method 0x15977d45.
//
// Solidity: function ledger(bytes32 ) view returns(string uri, string description, address addr)
func (_DataLedgerContract *DataLedgerContractSession) Ledger(arg0 [32]byte) (struct {
	Uri         string
	Description string
	Addr        common.Address
}, error) {
	return _DataLedgerContract.Contract.Ledger(&_DataLedgerContract.CallOpts, arg0)
}

// Ledger is a free data retrieval call binding the contract method 0x15977d45.
//
// Solidity: function ledger(bytes32 ) view returns(string uri, string description, address addr)
func (_DataLedgerContract *DataLedgerContractCallerSession) Ledger(arg0 [32]byte) (struct {
	Uri         string
	Description string
	Addr        common.Address
}, error) {
	return _DataLedgerContract.Contract.Ledger(&_DataLedgerContract.CallOpts, arg0)
}

// DeleteMeasurement is a paid mutator transaction binding the contract method 0x77ad95ca.
//
// Solidity: function deleteMeasurement(bytes32 hash) returns()
func (_DataLedgerContract *DataLedgerContractTransactor) DeleteMeasurement(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _DataLedgerContract.contract.Transact(opts, "deleteMeasurement", hash)
}

// DeleteMeasurement is a paid mutator transaction binding the contract method 0x77ad95ca.
//
// Solidity: function deleteMeasurement(bytes32 hash) returns()
func (_DataLedgerContract *DataLedgerContractSession) DeleteMeasurement(hash [32]byte) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.DeleteMeasurement(&_DataLedgerContract.TransactOpts, hash)
}

// DeleteMeasurement is a paid mutator transaction binding the contract method 0x77ad95ca.
//
// Solidity: function deleteMeasurement(bytes32 hash) returns()
func (_DataLedgerContract *DataLedgerContractTransactorSession) DeleteMeasurement(hash [32]byte) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.DeleteMeasurement(&_DataLedgerContract.TransactOpts, hash)
}

// SetAddress is a paid mutator transaction binding the contract method 0xe30081a0.
//
// Solidity: function setAddress(address _address) returns()
func (_DataLedgerContract *DataLedgerContractTransactor) SetAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _DataLedgerContract.contract.Transact(opts, "setAddress", _address)
}

// SetAddress is a paid mutator transaction binding the contract method 0xe30081a0.
//
// Solidity: function setAddress(address _address) returns()
func (_DataLedgerContract *DataLedgerContractSession) SetAddress(_address common.Address) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.SetAddress(&_DataLedgerContract.TransactOpts, _address)
}

// SetAddress is a paid mutator transaction binding the contract method 0xe30081a0.
//
// Solidity: function setAddress(address _address) returns()
func (_DataLedgerContract *DataLedgerContractTransactorSession) SetAddress(_address common.Address) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.SetAddress(&_DataLedgerContract.TransactOpts, _address)
}

// StoreInfo is a paid mutator transaction binding the contract method 0xb7e2a1b8.
//
// Solidity: function storeInfo(bytes32 hash, string uri, string description) returns()
func (_DataLedgerContract *DataLedgerContractTransactor) StoreInfo(opts *bind.TransactOpts, hash [32]byte, uri string, description string) (*types.Transaction, error) {
	return _DataLedgerContract.contract.Transact(opts, "storeInfo", hash, uri, description)
}

// StoreInfo is a paid mutator transaction binding the contract method 0xb7e2a1b8.
//
// Solidity: function storeInfo(bytes32 hash, string uri, string description) returns()
func (_DataLedgerContract *DataLedgerContractSession) StoreInfo(hash [32]byte, uri string, description string) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.StoreInfo(&_DataLedgerContract.TransactOpts, hash, uri, description)
}

// StoreInfo is a paid mutator transaction binding the contract method 0xb7e2a1b8.
//
// Solidity: function storeInfo(bytes32 hash, string uri, string description) returns()
func (_DataLedgerContract *DataLedgerContractTransactorSession) StoreInfo(hash [32]byte, uri string, description string) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.StoreInfo(&_DataLedgerContract.TransactOpts, hash, uri, description)
}

// DataLedgerContractDeleteInfoIterator is returned from FilterDeleteInfo and is used to iterate over the raw logs and unpacked data for DeleteInfo events raised by the DataLedgerContract contract.
type DataLedgerContractDeleteInfoIterator struct {
	Event *DataLedgerContractDeleteInfo // Event containing the contract specifics and raw log

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
func (it *DataLedgerContractDeleteInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataLedgerContractDeleteInfo)
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
		it.Event = new(DataLedgerContractDeleteInfo)
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
func (it *DataLedgerContractDeleteInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataLedgerContractDeleteInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataLedgerContractDeleteInfo represents a DeleteInfo event raised by the DataLedgerContract contract.
type DataLedgerContractDeleteInfo struct {
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDeleteInfo is a free log retrieval operation binding the contract event 0x072007d551e16de6c1b8938fdd0559f70033d87037e5dffa28631256df69f9fe.
//
// Solidity: event deleteInfo(bytes32 indexed _hash)
func (_DataLedgerContract *DataLedgerContractFilterer) FilterDeleteInfo(opts *bind.FilterOpts, _hash [][32]byte) (*DataLedgerContractDeleteInfoIterator, error) {

	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _DataLedgerContract.contract.FilterLogs(opts, "deleteInfo", _hashRule)
	if err != nil {
		return nil, err
	}
	return &DataLedgerContractDeleteInfoIterator{contract: _DataLedgerContract.contract, event: "deleteInfo", logs: logs, sub: sub}, nil
}

// WatchDeleteInfo is a free log subscription operation binding the contract event 0x072007d551e16de6c1b8938fdd0559f70033d87037e5dffa28631256df69f9fe.
//
// Solidity: event deleteInfo(bytes32 indexed _hash)
func (_DataLedgerContract *DataLedgerContractFilterer) WatchDeleteInfo(opts *bind.WatchOpts, sink chan<- *DataLedgerContractDeleteInfo, _hash [][32]byte) (event.Subscription, error) {

	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _DataLedgerContract.contract.WatchLogs(opts, "deleteInfo", _hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataLedgerContractDeleteInfo)
				if err := _DataLedgerContract.contract.UnpackLog(event, "deleteInfo", log); err != nil {
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

// ParseDeleteInfo is a log parse operation binding the contract event 0x072007d551e16de6c1b8938fdd0559f70033d87037e5dffa28631256df69f9fe.
//
// Solidity: event deleteInfo(bytes32 indexed _hash)
func (_DataLedgerContract *DataLedgerContractFilterer) ParseDeleteInfo(log types.Log) (*DataLedgerContractDeleteInfo, error) {
	event := new(DataLedgerContractDeleteInfo)
	if err := _DataLedgerContract.contract.UnpackLog(event, "deleteInfo", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DataLedgerContractEvtStoreInfoIterator is returned from FilterEvtStoreInfo and is used to iterate over the raw logs and unpacked data for EvtStoreInfo events raised by the DataLedgerContract contract.
type DataLedgerContractEvtStoreInfoIterator struct {
	Event *DataLedgerContractEvtStoreInfo // Event containing the contract specifics and raw log

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
func (it *DataLedgerContractEvtStoreInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataLedgerContractEvtStoreInfo)
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
		it.Event = new(DataLedgerContractEvtStoreInfo)
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
func (it *DataLedgerContractEvtStoreInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataLedgerContractEvtStoreInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataLedgerContractEvtStoreInfo represents a EvtStoreInfo event raised by the DataLedgerContract contract.
type DataLedgerContractEvtStoreInfo struct {
	Hash        [32]byte
	Uri         string
	Description string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEvtStoreInfo is a free log retrieval operation binding the contract event 0xafeab48c124d8423d588c0406f25d8386c21751c23cdf11491137a519343ec22.
//
// Solidity: event evtStoreInfo(bytes32 indexed _hash, string _uri, string _description)
func (_DataLedgerContract *DataLedgerContractFilterer) FilterEvtStoreInfo(opts *bind.FilterOpts, _hash [][32]byte) (*DataLedgerContractEvtStoreInfoIterator, error) {

	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _DataLedgerContract.contract.FilterLogs(opts, "evtStoreInfo", _hashRule)
	if err != nil {
		return nil, err
	}
	return &DataLedgerContractEvtStoreInfoIterator{contract: _DataLedgerContract.contract, event: "evtStoreInfo", logs: logs, sub: sub}, nil
}

// WatchEvtStoreInfo is a free log subscription operation binding the contract event 0xafeab48c124d8423d588c0406f25d8386c21751c23cdf11491137a519343ec22.
//
// Solidity: event evtStoreInfo(bytes32 indexed _hash, string _uri, string _description)
func (_DataLedgerContract *DataLedgerContractFilterer) WatchEvtStoreInfo(opts *bind.WatchOpts, sink chan<- *DataLedgerContractEvtStoreInfo, _hash [][32]byte) (event.Subscription, error) {

	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _DataLedgerContract.contract.WatchLogs(opts, "evtStoreInfo", _hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataLedgerContractEvtStoreInfo)
				if err := _DataLedgerContract.contract.UnpackLog(event, "evtStoreInfo", log); err != nil {
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

// ParseEvtStoreInfo is a log parse operation binding the contract event 0xafeab48c124d8423d588c0406f25d8386c21751c23cdf11491137a519343ec22.
//
// Solidity: event evtStoreInfo(bytes32 indexed _hash, string _uri, string _description)
func (_DataLedgerContract *DataLedgerContractFilterer) ParseEvtStoreInfo(log types.Log) (*DataLedgerContractEvtStoreInfo, error) {
	event := new(DataLedgerContractEvtStoreInfo)
	if err := _DataLedgerContract.contract.UnpackLog(event, "evtStoreInfo", log); err != nil {
		return nil, err
	}
	return event, nil
}
