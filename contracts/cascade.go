// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// CascadeABI is the input ABI used to generate the binding from.
const CascadeABI = "[{\"inputs\":[],\"name\":\"message\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"initialMessage\",\"type\":\"string\"}],\"name\":\"readMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newMessage\",\"type\":\"string\"}],\"name\":\"setMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CascadeFuncSigs maps the 4-byte function signature to its string representation.
var CascadeFuncSigs = map[string]string{
	"e21f37ce": "message()",
	"5771d6e4": "readMessage(string)",
	"368b8772": "setMessage(string)",
}

// CascadeBin is the compiled bytecode used for deploying new contracts.
var CascadeBin = "0x608060405234801561001057600080fd5b50610344806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063368b8772146100465780635771d6e414610046578063e21f37ce1461005b575b600080fd5b6100596100543660046101b7565b610079565b005b610063610090565b6040516100709190610268565b60405180910390f35b805161008c90600090602084019061011e565b5050565b6000805461009d906102bd565b80601f01602080910402602001604051908101604052809291908181526020018280546100c9906102bd565b80156101165780601f106100eb57610100808354040283529160200191610116565b820191906000526020600020905b8154815290600101906020018083116100f957829003601f168201915b505050505081565b82805461012a906102bd565b90600052602060002090601f01602090048101928261014c5760008555610192565b82601f1061016557805160ff1916838001178555610192565b82800160010185558215610192579182015b82811115610192578251825591602001919060010190610177565b5061019e9291506101a2565b5090565b5b8082111561019e57600081556001016101a3565b6000602082840312156101c957600080fd5b813567ffffffffffffffff808211156101e157600080fd5b818401915084601f8301126101f557600080fd5b813581811115610207576102076102f8565b604051601f8201601f19908116603f0116810190838211818310171561022f5761022f6102f8565b8160405282815287602084870101111561024857600080fd5b826020860160208301376000928101602001929092525095945050505050565b600060208083528351808285015260005b8181101561029557858101830151858201604001528201610279565b818111156102a7576000604083870101525b50601f01601f1916929092016040019392505050565b600181811c908216806102d157607f821691505b602082108114156102f257634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fdfea2646970667358221220bc9300070ab313522b77153767a367eef38e8aa994e33c4aeefca22f52f601b264736f6c63430008060033"

// DeployCascade deploys a new Ethereum contract, binding an instance of Cascade to it.
func DeployCascade(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Cascade, error) {
	parsed, err := abi.JSON(strings.NewReader(CascadeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CascadeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Cascade{CascadeCaller: CascadeCaller{contract: contract}, CascadeTransactor: CascadeTransactor{contract: contract}, CascadeFilterer: CascadeFilterer{contract: contract}}, nil
}

// Cascade is an auto generated Go binding around an Ethereum contract.
type Cascade struct {
	CascadeCaller     // Read-only binding to the contract
	CascadeTransactor // Write-only binding to the contract
	CascadeFilterer   // Log filterer for contract events
}

// CascadeCaller is an auto generated read-only Go binding around an Ethereum contract.
type CascadeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CascadeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CascadeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CascadeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CascadeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CascadeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CascadeSession struct {
	Contract     *Cascade          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CascadeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CascadeCallerSession struct {
	Contract *CascadeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CascadeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CascadeTransactorSession struct {
	Contract     *CascadeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CascadeRaw is an auto generated low-level Go binding around an Ethereum contract.
type CascadeRaw struct {
	Contract *Cascade // Generic contract binding to access the raw methods on
}

// CascadeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CascadeCallerRaw struct {
	Contract *CascadeCaller // Generic read-only contract binding to access the raw methods on
}

// CascadeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CascadeTransactorRaw struct {
	Contract *CascadeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCascade creates a new instance of Cascade, bound to a specific deployed contract.
func NewCascade(address common.Address, backend bind.ContractBackend) (*Cascade, error) {
	contract, err := bindCascade(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cascade{CascadeCaller: CascadeCaller{contract: contract}, CascadeTransactor: CascadeTransactor{contract: contract}, CascadeFilterer: CascadeFilterer{contract: contract}}, nil
}

// NewCascadeCaller creates a new read-only instance of Cascade, bound to a specific deployed contract.
func NewCascadeCaller(address common.Address, caller bind.ContractCaller) (*CascadeCaller, error) {
	contract, err := bindCascade(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CascadeCaller{contract: contract}, nil
}

// NewCascadeTransactor creates a new write-only instance of Cascade, bound to a specific deployed contract.
func NewCascadeTransactor(address common.Address, transactor bind.ContractTransactor) (*CascadeTransactor, error) {
	contract, err := bindCascade(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CascadeTransactor{contract: contract}, nil
}

// NewCascadeFilterer creates a new log filterer instance of Cascade, bound to a specific deployed contract.
func NewCascadeFilterer(address common.Address, filterer bind.ContractFilterer) (*CascadeFilterer, error) {
	contract, err := bindCascade(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CascadeFilterer{contract: contract}, nil
}

// bindCascade binds a generic wrapper to an already deployed contract.
func bindCascade(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CascadeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cascade *CascadeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cascade.Contract.CascadeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cascade *CascadeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cascade.Contract.CascadeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cascade *CascadeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cascade.Contract.CascadeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cascade *CascadeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cascade.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cascade *CascadeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cascade.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cascade *CascadeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cascade.Contract.contract.Transact(opts, method, params...)
}

// Message is a free data retrieval call binding the contract method 0xe21f37ce.
//
// Solidity: function message() view returns(string)
func (_Cascade *CascadeCaller) Message(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Cascade.contract.Call(opts, &out, "message")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Message is a free data retrieval call binding the contract method 0xe21f37ce.
//
// Solidity: function message() view returns(string)
func (_Cascade *CascadeSession) Message() (string, error) {
	return _Cascade.Contract.Message(&_Cascade.CallOpts)
}

// Message is a free data retrieval call binding the contract method 0xe21f37ce.
//
// Solidity: function message() view returns(string)
func (_Cascade *CascadeCallerSession) Message() (string, error) {
	return _Cascade.Contract.Message(&_Cascade.CallOpts)
}

// ReadMessage is a paid mutator transaction binding the contract method 0x5771d6e4.
//
// Solidity: function readMessage(string initialMessage) returns()
func (_Cascade *CascadeTransactor) ReadMessage(opts *bind.TransactOpts, initialMessage string) (*types.Transaction, error) {
	return _Cascade.contract.Transact(opts, "readMessage", initialMessage)
}

// ReadMessage is a paid mutator transaction binding the contract method 0x5771d6e4.
//
// Solidity: function readMessage(string initialMessage) returns()
func (_Cascade *CascadeSession) ReadMessage(initialMessage string) (*types.Transaction, error) {
	return _Cascade.Contract.ReadMessage(&_Cascade.TransactOpts, initialMessage)
}

// ReadMessage is a paid mutator transaction binding the contract method 0x5771d6e4.
//
// Solidity: function readMessage(string initialMessage) returns()
func (_Cascade *CascadeTransactorSession) ReadMessage(initialMessage string) (*types.Transaction, error) {
	return _Cascade.Contract.ReadMessage(&_Cascade.TransactOpts, initialMessage)
}

// SetMessage is a paid mutator transaction binding the contract method 0x368b8772.
//
// Solidity: function setMessage(string newMessage) returns()
func (_Cascade *CascadeTransactor) SetMessage(opts *bind.TransactOpts, newMessage string) (*types.Transaction, error) {
	return _Cascade.contract.Transact(opts, "setMessage", newMessage)
}

// SetMessage is a paid mutator transaction binding the contract method 0x368b8772.
//
// Solidity: function setMessage(string newMessage) returns()
func (_Cascade *CascadeSession) SetMessage(newMessage string) (*types.Transaction, error) {
	return _Cascade.Contract.SetMessage(&_Cascade.TransactOpts, newMessage)
}

// SetMessage is a paid mutator transaction binding the contract method 0x368b8772.
//
// Solidity: function setMessage(string newMessage) returns()
func (_Cascade *CascadeTransactorSession) SetMessage(newMessage string) (*types.Transaction, error) {
	return _Cascade.Contract.SetMessage(&_Cascade.TransactOpts, newMessage)
}
