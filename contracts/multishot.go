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

// MultishotABI is the input ABI used to generate the binding from.
const MultishotABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"stake\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"decisions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"decision\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weight_received\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentMax\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"majorityStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"txOrigin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txHash\",\"type\":\"uint256\"}],\"name\":\"propose\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"txOrigin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"read\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"txHash\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validatorStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MultishotFuncSigs maps the 4-byte function signature to its string representation.
var MultishotFuncSigs = map[string]string{
	"d6a315a3": "decisions(address,uint256)",
	"f96e8001": "majorityStake()",
	"0d34b2ad": "propose(address,uint256,uint256)",
	"014c2add": "read(address,uint256)",
	"39b7fcc6": "validatorStake(address)",
}

// MultishotBin is the compiled bytecode used for deploying new contracts.
var MultishotBin = "0x608060405234801561001057600080fd5b5060405161079938038061079983398101604081905261002f91610166565b6000805b83518110156100d05782818151811061004e5761004e61031a565b602002602001015160008086848151811061006b5761006b61031a565b60200260200101516001600160a01b03166001600160a01b03168152602001908152602001600020819055508281815181106100a9576100a961031a565b6020026020010151826100bc9190610290565b9150806100c8816102e9565b915050610033565b5060056100de8260046102ca565b6100e891906102a8565b60015550610346915050565b600082601f83011261010557600080fd5b8151602061011a6101158361026d565b61023d565b80838252828201915082860187848660051b890101111561013a57600080fd5b60005b858110156101595781518452928401929084019060010161013d565b5090979650505050505050565b6000806040838503121561017957600080fd5b82516001600160401b038082111561019057600080fd5b818501915085601f8301126101a457600080fd5b815160206101b46101158361026d565b8083825282820191508286018a848660051b89010111156101d457600080fd5b600096505b8487101561020c5780516001600160a01b03811681146101f857600080fd5b8352600196909601959183019183016101d9565b509188015191965090935050508082111561022657600080fd5b50610233858286016100f4565b9150509250929050565b604051601f8201601f191681016001600160401b038111828210171561026557610265610330565b604052919050565b60006001600160401b0382111561028657610286610330565b5060051b60200190565b600082198211156102a3576102a3610304565b500190565b6000826102c557634e487b7160e01b600052601260045260246000fd5b500490565b60008160001904831182151516156102e4576102e4610304565b500290565b60006000198214156102fd576102fd610304565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b610444806103556000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063014c2add1461005c5780630d34b2ad1461008257806339b7fcc614610097578063d6a315a3146100b7578063f96e800114610108575b600080fd5b61006f61006a36600461038b565b610111565b6040519081526020015b60405180910390f35b6100956100903660046103b5565b61014f565b005b61006f6100a5366004610369565b60006020819052908152604090205481565b6100ed6100c536600461038b565b6002602081815260009384526040808520909152918352912080546001820154919092015483565b60408051938452602084019290925290820152606001610079565b61006f60015481565b6001600160a01b0382166000908152600260209081526040808320848452909152812080541561014357549050610149565b60009150505b92915050565b6001600160a01b03831660009081526002602090815260408083208584528252808320338452600381019092529091205460ff16156101cd5760405162461bcd60e51b81526020600482015260156024820152742cb7ba9030b63932b0b23c90383937b837b9b2b21760591b60448201526064015b60405180910390fd5b80541561021c5760405162461bcd60e51b815260206004820152601c60248201527f436f6e73656e73757320616c72656164792064656c6976657265642e0000000060448201526064016101c4565b3360009081526020819052604090205461028e5760405162461bcd60e51b815260206004820152602d60248201527f596f75206e656564207374616b6520746f2062652061626c6520746f2070726f60448201526c3837b9b29030903b30b63ab29760991b60648201526084016101c4565b33600090815260208181526040808320548584526004850190925282208054919290916102bc9084906103e8565b909155505033600090815260208190526040812054600183018054919290916102e69084906103e8565b909155505060008281526004820160205260408082205460028401548352912054101561031557600281018290555b6001548160010154111561032b57600281015481555b336000908152600390910160205260409020805460ff19166001179055505050565b80356001600160a01b038116811461036457600080fd5b919050565b60006020828403121561037b57600080fd5b6103848261034d565b9392505050565b6000806040838503121561039e57600080fd5b6103a78361034d565b946020939093013593505050565b6000806000606084860312156103ca57600080fd5b6103d38461034d565b95602085013595506040909401359392505050565b6000821982111561040957634e487b7160e01b600052601160045260246000fd5b50019056fea2646970667358221220f09c6d1da56e4d970d790747a5b509a8f661a69c0a7656d8804976388e174b5664736f6c63430008060033"

// DeployMultishot deploys a new Ethereum contract, binding an instance of Multishot to it.
func DeployMultishot(auth *bind.TransactOpts, backend bind.ContractBackend, validators []common.Address, stake []*big.Int) (common.Address, *types.Transaction, *Multishot, error) {
	parsed, err := abi.JSON(strings.NewReader(MultishotABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MultishotBin), backend, validators, stake)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Multishot{MultishotCaller: MultishotCaller{contract: contract}, MultishotTransactor: MultishotTransactor{contract: contract}, MultishotFilterer: MultishotFilterer{contract: contract}}, nil
}

// Multishot is an auto generated Go binding around an Ethereum contract.
type Multishot struct {
	MultishotCaller     // Read-only binding to the contract
	MultishotTransactor // Write-only binding to the contract
	MultishotFilterer   // Log filterer for contract events
}

// MultishotCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultishotCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultishotTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultishotTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultishotFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultishotFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultishotSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultishotSession struct {
	Contract     *Multishot        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MultishotCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultishotCallerSession struct {
	Contract *MultishotCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MultishotTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultishotTransactorSession struct {
	Contract     *MultishotTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MultishotRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultishotRaw struct {
	Contract *Multishot // Generic contract binding to access the raw methods on
}

// MultishotCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultishotCallerRaw struct {
	Contract *MultishotCaller // Generic read-only contract binding to access the raw methods on
}

// MultishotTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultishotTransactorRaw struct {
	Contract *MultishotTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultishot creates a new instance of Multishot, bound to a specific deployed contract.
func NewMultishot(address common.Address, backend bind.ContractBackend) (*Multishot, error) {
	contract, err := bindMultishot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Multishot{MultishotCaller: MultishotCaller{contract: contract}, MultishotTransactor: MultishotTransactor{contract: contract}, MultishotFilterer: MultishotFilterer{contract: contract}}, nil
}

// NewMultishotCaller creates a new read-only instance of Multishot, bound to a specific deployed contract.
func NewMultishotCaller(address common.Address, caller bind.ContractCaller) (*MultishotCaller, error) {
	contract, err := bindMultishot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultishotCaller{contract: contract}, nil
}

// NewMultishotTransactor creates a new write-only instance of Multishot, bound to a specific deployed contract.
func NewMultishotTransactor(address common.Address, transactor bind.ContractTransactor) (*MultishotTransactor, error) {
	contract, err := bindMultishot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultishotTransactor{contract: contract}, nil
}

// NewMultishotFilterer creates a new log filterer instance of Multishot, bound to a specific deployed contract.
func NewMultishotFilterer(address common.Address, filterer bind.ContractFilterer) (*MultishotFilterer, error) {
	contract, err := bindMultishot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultishotFilterer{contract: contract}, nil
}

// bindMultishot binds a generic wrapper to an already deployed contract.
func bindMultishot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MultishotABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multishot *MultishotRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multishot.Contract.MultishotCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multishot *MultishotRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multishot.Contract.MultishotTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multishot *MultishotRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multishot.Contract.MultishotTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multishot *MultishotCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multishot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multishot *MultishotTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multishot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multishot *MultishotTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multishot.Contract.contract.Transact(opts, method, params...)
}

// Decisions is a free data retrieval call binding the contract method 0xd6a315a3.
//
// Solidity: function decisions(address , uint256 ) view returns(uint256 decision, uint256 weight_received, uint256 currentMax)
func (_Multishot *MultishotCaller) Decisions(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Decision       *big.Int
	WeightReceived *big.Int
	CurrentMax     *big.Int
}, error) {
	var out []interface{}
	err := _Multishot.contract.Call(opts, &out, "decisions", arg0, arg1)

	outstruct := new(struct {
		Decision       *big.Int
		WeightReceived *big.Int
		CurrentMax     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Decision = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.WeightReceived = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CurrentMax = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Decisions is a free data retrieval call binding the contract method 0xd6a315a3.
//
// Solidity: function decisions(address , uint256 ) view returns(uint256 decision, uint256 weight_received, uint256 currentMax)
func (_Multishot *MultishotSession) Decisions(arg0 common.Address, arg1 *big.Int) (struct {
	Decision       *big.Int
	WeightReceived *big.Int
	CurrentMax     *big.Int
}, error) {
	return _Multishot.Contract.Decisions(&_Multishot.CallOpts, arg0, arg1)
}

// Decisions is a free data retrieval call binding the contract method 0xd6a315a3.
//
// Solidity: function decisions(address , uint256 ) view returns(uint256 decision, uint256 weight_received, uint256 currentMax)
func (_Multishot *MultishotCallerSession) Decisions(arg0 common.Address, arg1 *big.Int) (struct {
	Decision       *big.Int
	WeightReceived *big.Int
	CurrentMax     *big.Int
}, error) {
	return _Multishot.Contract.Decisions(&_Multishot.CallOpts, arg0, arg1)
}

// MajorityStake is a free data retrieval call binding the contract method 0xf96e8001.
//
// Solidity: function majorityStake() view returns(uint256)
func (_Multishot *MultishotCaller) MajorityStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Multishot.contract.Call(opts, &out, "majorityStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MajorityStake is a free data retrieval call binding the contract method 0xf96e8001.
//
// Solidity: function majorityStake() view returns(uint256)
func (_Multishot *MultishotSession) MajorityStake() (*big.Int, error) {
	return _Multishot.Contract.MajorityStake(&_Multishot.CallOpts)
}

// MajorityStake is a free data retrieval call binding the contract method 0xf96e8001.
//
// Solidity: function majorityStake() view returns(uint256)
func (_Multishot *MultishotCallerSession) MajorityStake() (*big.Int, error) {
	return _Multishot.Contract.MajorityStake(&_Multishot.CallOpts)
}

// Read is a free data retrieval call binding the contract method 0x014c2add.
//
// Solidity: function read(address txOrigin, uint256 nonce) view returns(uint256 txHash)
func (_Multishot *MultishotCaller) Read(opts *bind.CallOpts, txOrigin common.Address, nonce *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Multishot.contract.Call(opts, &out, "read", txOrigin, nonce)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Read is a free data retrieval call binding the contract method 0x014c2add.
//
// Solidity: function read(address txOrigin, uint256 nonce) view returns(uint256 txHash)
func (_Multishot *MultishotSession) Read(txOrigin common.Address, nonce *big.Int) (*big.Int, error) {
	return _Multishot.Contract.Read(&_Multishot.CallOpts, txOrigin, nonce)
}

// Read is a free data retrieval call binding the contract method 0x014c2add.
//
// Solidity: function read(address txOrigin, uint256 nonce) view returns(uint256 txHash)
func (_Multishot *MultishotCallerSession) Read(txOrigin common.Address, nonce *big.Int) (*big.Int, error) {
	return _Multishot.Contract.Read(&_Multishot.CallOpts, txOrigin, nonce)
}

// ValidatorStake is a free data retrieval call binding the contract method 0x39b7fcc6.
//
// Solidity: function validatorStake(address ) view returns(uint256)
func (_Multishot *MultishotCaller) ValidatorStake(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Multishot.contract.Call(opts, &out, "validatorStake", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorStake is a free data retrieval call binding the contract method 0x39b7fcc6.
//
// Solidity: function validatorStake(address ) view returns(uint256)
func (_Multishot *MultishotSession) ValidatorStake(arg0 common.Address) (*big.Int, error) {
	return _Multishot.Contract.ValidatorStake(&_Multishot.CallOpts, arg0)
}

// ValidatorStake is a free data retrieval call binding the contract method 0x39b7fcc6.
//
// Solidity: function validatorStake(address ) view returns(uint256)
func (_Multishot *MultishotCallerSession) ValidatorStake(arg0 common.Address) (*big.Int, error) {
	return _Multishot.Contract.ValidatorStake(&_Multishot.CallOpts, arg0)
}

// Propose is a paid mutator transaction binding the contract method 0x0d34b2ad.
//
// Solidity: function propose(address txOrigin, uint256 nonce, uint256 txHash) returns()
func (_Multishot *MultishotTransactor) Propose(opts *bind.TransactOpts, txOrigin common.Address, nonce *big.Int, txHash *big.Int) (*types.Transaction, error) {
	return _Multishot.contract.Transact(opts, "propose", txOrigin, nonce, txHash)
}

// Propose is a paid mutator transaction binding the contract method 0x0d34b2ad.
//
// Solidity: function propose(address txOrigin, uint256 nonce, uint256 txHash) returns()
func (_Multishot *MultishotSession) Propose(txOrigin common.Address, nonce *big.Int, txHash *big.Int) (*types.Transaction, error) {
	return _Multishot.Contract.Propose(&_Multishot.TransactOpts, txOrigin, nonce, txHash)
}

// Propose is a paid mutator transaction binding the contract method 0x0d34b2ad.
//
// Solidity: function propose(address txOrigin, uint256 nonce, uint256 txHash) returns()
func (_Multishot *MultishotTransactorSession) Propose(txOrigin common.Address, nonce *big.Int, txHash *big.Int) (*types.Transaction, error) {
	return _Multishot.Contract.Propose(&_Multishot.TransactOpts, txOrigin, nonce, txHash)
}
