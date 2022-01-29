// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abibindings

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
)

// ContractabiMetaData contains all meta data concerning the Contractabi contract.
var ContractabiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_baseURI\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_PER_TX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_PER_WALLET\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_SUPPLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_WHITELIST\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESERVES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addressToMinted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data_\",\"type\":\"bytes\"}],\"name\":\"batchSafeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"batchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collectReserves\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extension\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"isOwnerOf\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_state\",\"type\":\"bool\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceInWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceWhitelist\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"revealedURI\",\"type\":\"string\"}],\"name\":\"reveal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revealed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_baseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_extension\",\"type\":\"string\"}],\"name\":\"setExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_max_per_tx\",\"type\":\"uint256\"}],\"name\":\"setMaxPerTx\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_max_per_wallet\",\"type\":\"uint256\"}],\"name\":\"setMaxPerWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_max_whitelist\",\"type\":\"uint256\"}],\"name\":\"setMaxWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setPriceInWei\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setPriceWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_whitelistActive\",\"type\":\"bool\"}],\"name\":\"setWhitelistActive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_whitelistMerkleRoot\",\"type\":\"bytes32\"}],\"name\":\"setWhitelistMerkleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"walletOfOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dummy\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"whitelistMint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// ContractabiABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractabiMetaData.ABI instead.
var ContractabiABI = ContractabiMetaData.ABI

// Contractabi is an auto generated Go binding around an Ethereum contract.
type Contractabi struct {
	ContractabiCaller     // Read-only binding to the contract
	ContractabiTransactor // Write-only binding to the contract
	ContractabiFilterer   // Log filterer for contract events
}

// ContractabiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractabiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractabiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractabiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractabiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractabiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractabiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractabiSession struct {
	Contract     *Contractabi      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractabiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractabiCallerSession struct {
	Contract *ContractabiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ContractabiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractabiTransactorSession struct {
	Contract     *ContractabiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ContractabiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractabiRaw struct {
	Contract *Contractabi // Generic contract binding to access the raw methods on
}

// ContractabiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractabiCallerRaw struct {
	Contract *ContractabiCaller // Generic read-only contract binding to access the raw methods on
}

// ContractabiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractabiTransactorRaw struct {
	Contract *ContractabiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractabi creates a new instance of Contractabi, bound to a specific deployed contract.
func NewContractabi(address common.Address, backend bind.ContractBackend) (*Contractabi, error) {
	contract, err := bindContractabi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contractabi{ContractabiCaller: ContractabiCaller{contract: contract}, ContractabiTransactor: ContractabiTransactor{contract: contract}, ContractabiFilterer: ContractabiFilterer{contract: contract}}, nil
}

// NewContractabiCaller creates a new read-only instance of Contractabi, bound to a specific deployed contract.
func NewContractabiCaller(address common.Address, caller bind.ContractCaller) (*ContractabiCaller, error) {
	contract, err := bindContractabi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractabiCaller{contract: contract}, nil
}

// NewContractabiTransactor creates a new write-only instance of Contractabi, bound to a specific deployed contract.
func NewContractabiTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractabiTransactor, error) {
	contract, err := bindContractabi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractabiTransactor{contract: contract}, nil
}

// NewContractabiFilterer creates a new log filterer instance of Contractabi, bound to a specific deployed contract.
func NewContractabiFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractabiFilterer, error) {
	contract, err := bindContractabi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractabiFilterer{contract: contract}, nil
}

// bindContractabi binds a generic wrapper to an already deployed contract.
func bindContractabi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractabiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contractabi *ContractabiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contractabi.Contract.ContractabiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contractabi *ContractabiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contractabi.Contract.ContractabiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contractabi *ContractabiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contractabi.Contract.ContractabiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contractabi *ContractabiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contractabi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contractabi *ContractabiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contractabi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contractabi *ContractabiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contractabi.Contract.contract.Transact(opts, method, params...)
}

// MAXPERTX is a free data retrieval call binding the contract method 0xf43a22dc.
//
// Solidity: function MAX_PER_TX() view returns(uint256)
func (_Contractabi *ContractabiCaller) MAXPERTX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contractabi.contract.Call(opts, &out, "MAX_PER_TX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXPERTX is a free data retrieval call binding the contract method 0xf43a22dc.
//
// Solidity: function MAX_PER_TX() view returns(uint256)
func (_Contractabi *ContractabiSession) MAXPERTX() (*big.Int, error) {
	return _Contractabi.Contract.MAXPERTX(&_Contractabi.CallOpts)
}

// MAXPERTX is a free data retrieval call binding the contract method 0xf43a22dc.
//
// Solidity: function MAX_PER_TX() view returns(uint256)
func (_Contractabi *ContractabiCallerSession) MAXPERTX() (*big.Int, error) {
	return _Contractabi.Contract.MAXPERTX(&_Contractabi.CallOpts)
}

// MAXPERWALLET is a free data retrieval call binding the contract method 0x0f2cdd6c.
//
// Solidity: function MAX_PER_WALLET() view returns(uint256)
func (_Contractabi *ContractabiCaller) MAXPERWALLET(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contractabi.contract.Call(opts, &out, "MAX_PER_WALLET")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXPERWALLET is a free data retrieval call binding the contract method 0x0f2cdd6c.
//
// Solidity: function MAX_PER_WALLET() view returns(uint256)
func (_Contractabi *ContractabiSession) MAXPERWALLET() (*big.Int, error) {
	return _Contractabi.Contract.MAXPERWALLET(&_Contractabi.CallOpts)
}

// MAXPERWALLET is a free data retrieval call binding the contract method 0x0f2cdd6c.
//
// Solidity: function MAX_PER_WALLET() view returns(uint256)
func (_Contractabi *ContractabiCallerSession) MAXPERWALLET() (*big.Int, error) {
	return _Contractabi.Contract.MAXPERWALLET(&_Contractabi.CallOpts)
}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint256)
func (_Contractabi *ContractabiCaller) MAXSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contractabi.contract.Call(opts, &out, "MAX_SUPPLY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint256)
func (_Contractabi *ContractabiSession) MAXSUPPLY() (*big.Int, error) {
	return _Contractabi.Contract.MAXSUPPLY(&_Contractabi.CallOpts)
}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint256)
func (_Contractabi *ContractabiCallerSession) MAXSUPPLY() (*big.Int, error) {
	return _Contractabi.Contract.MAXSUPPLY(&_Contractabi.CallOpts)
}

// MAXWHITELIST is a free data retrieval call binding the contract method 0x95f61c48.
//
// Solidity: function MAX_WHITELIST() view returns(uint256)
func (_Contractabi *ContractabiCaller) MAXWHITELIST(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contractabi.contract.Call(opts, &out, "MAX_WHITELIST")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXWHITELIST is a free data retrieval call binding the contract method 0x95f61c48.
//
// Solidity: function MAX_WHITELIST() view returns(uint256)
func (_Contractabi *ContractabiSession) MAXWHITELIST() (*big.Int, error) {
	return _Contractabi.Contract.MAXWHITELIST(&_Contractabi.CallOpts)
}

// MAXWHITELIST is a free data retrieval call binding the contract method 0x95f61c48.
//
// Solidity: function MAX_WHITELIST() view returns(uint256)
func (_Contractabi *ContractabiCallerSession) MAXWHITELIST() (*big.Int, error) {
	return _Contractabi.Contract.MAXWHITELIST(&_Contractabi.CallOpts)
}

// RESERVES is a free data retrieval call binding the contract method 0x0922f9c5.
//
// Solidity: function RESERVES() view returns(uint256)
func (_Contractabi *ContractabiCaller) RESERVES(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contractabi.contract.Call(opts, &out, "RESERVES")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RESERVES is a free data retrieval call binding the contract method 0x0922f9c5.
//
// Solidity: function RESERVES() view returns(uint256)
func (_Contractabi *ContractabiSession) RESERVES() (*big.Int, error) {
	return _Contractabi.Contract.RESERVES(&_Contractabi.CallOpts)
}

// RESERVES is a free data retrieval call binding the contract method 0x0922f9c5.
//
// Solidity: function RESERVES() view returns(uint256)
func (_Contractabi *ContractabiCallerSession) RESERVES() (*big.Int, error) {
	return _Contractabi.Contract.RESERVES(&_Contractabi.CallOpts)
}

// AddressToMinted is a free data retrieval call binding the contract method 0x9ec00c95.
//
// Solidity: function addressToMinted(address ) view returns(uint256)
func (_Contractabi *ContractabiCaller) AddressToMinted(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contractabi.contract.Call(opts, &out, "addressToMinted", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressToMinted is a free data retrieval call binding the contract method 0x9ec00c95.
//
// Solidity: function addressToMinted(address ) view returns(uint256)
func (_Contractabi *ContractabiSession) AddressToMinted(arg0 common.Address) (*big.Int, error) {
	return _Contractabi.Contract.AddressToMinted(&_Contractabi.CallOpts, arg0)
}

// AddressToMinted is a free data retrieval call binding the contract method 0x9ec00c95.
//
// Solidity: function addressToMinted(address ) view returns(uint256)
func (_Contractabi *ContractabiCallerSession) AddressToMinted(arg0 common.Address) (*big.Int, error) {
	return _Contractabi.Contract.AddressToMinted(&_Contractabi.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Contractabi *ContractabiCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contractabi.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Contractabi *ContractabiSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Contractabi.Contract.BalanceOf(&_Contractabi.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Contractabi *ContractabiCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Contractabi.Contract.BalanceOf(&_Contractabi.CallOpts, owner)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Contractabi *ContractabiCaller) BaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contractabi.contract.Call(opts, &out, "baseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Contractabi *ContractabiSession) BaseURI() (string, error) {
	return _Contractabi.Contract.BaseURI(&_Contractabi.CallOpts)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Contractabi *ContractabiCallerSession) BaseURI() (string, error) {
	return _Contractabi.Contract.BaseURI(&_Contractabi.CallOpts)
}

// Extension is a free data retrieval call binding the contract method 0x2d5537b0.
//
// Solidity: function extension() view returns(string)
func (_Contractabi *ContractabiCaller) Extension(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contractabi.contract.Call(opts, &out, "extension")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Extension is a free data retrieval call binding the contract method 0x2d5537b0.
//
// Solidity: function extension() view returns(string)
func (_Contractabi *ContractabiSession) Extension() (string, error) {
	return _Contractabi.Contract.Extension(&_Contractabi.CallOpts)
}

// Extension is a free data retrieval call binding the contract method 0x2d5537b0.
//
// Solidity: function extension() view returns(string)
func (_Contractabi *ContractabiCallerSession) Extension() (string, error) {
	return _Contractabi.Contract.Extension(&_Contractabi.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Contractabi *ContractabiCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contractabi.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Contractabi *ContractabiSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Contractabi.Contract.GetApproved(&_Contractabi.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Contractabi *ContractabiCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Contractabi.Contract.GetApproved(&_Contractabi.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Contractabi *ContractabiCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Contractabi.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Contractabi *ContractabiSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Contractabi.Contract.IsApprovedForAll(&_Contractabi.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Contractabi *ContractabiCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Contractabi.Contract.IsApprovedForAll(&_Contractabi.CallOpts, owner, operator)
}

// IsOwnerOf is a free data retrieval call binding the contract method 0x4d44660c.
//
// Solidity: function isOwnerOf(address account, uint256[] _tokenIds) view returns(bool)
func (_Contractabi *ContractabiCaller) IsOwnerOf(opts *bind.CallOpts, account common.Address, _tokenIds []*big.Int) (bool, error) {
	var out []interface{}
	err := _Contractabi.contract.Call(opts, &out, "isOwnerOf", account, _tokenIds)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwnerOf is a free data retrieval call binding the contract method 0x4d44660c.
//
// Solidity: function isOwnerOf(address account, uint256[] _tokenIds) view returns(bool)
func (_Contractabi *ContractabiSession) IsOwnerOf(account common.Address, _tokenIds []*big.Int) (bool, error) {
	return _Contractabi.Contract.IsOwnerOf(&_Contractabi.CallOpts, account, _tokenIds)
}

// IsOwnerOf is a free data retrieval call binding the contract method 0x4d44660c.
//
// Solidity: function isOwnerOf(address account, uint256[] _tokenIds) view returns(bool)
func (_Contractabi *ContractabiCallerSession) IsOwnerOf(account common.Address, _tokenIds []*big.Int) (bool, error) {
	return _Contractabi.Contract.IsOwnerOf(&_Contractabi.CallOpts, account, _tokenIds)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_Contractabi *ContractabiCaller) MerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contractabi.contract.Call(opts, &out, "merkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_Contractabi *ContractabiSession) MerkleRoot() ([32]byte, error) {
	return _Contractabi.Contract.MerkleRoot(&_Contractabi.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_Contractabi *ContractabiCallerSession) MerkleRoot() ([32]byte, error) {
	return _Contractabi.Contract.MerkleRoot(&_Contractabi.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contractabi *ContractabiCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contractabi.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contractabi *ContractabiSession) Name() (string, error) {
	return _Contractabi.Contract.Name(&_Contractabi.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contractabi *ContractabiCallerSession) Name() (string, error) {
	return _Contractabi.Contract.Name(&_Contractabi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contractabi *ContractabiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contractabi.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contractabi *ContractabiSession) Owner() (common.Address, error) {
	return _Contractabi.Contract.Owner(&_Contractabi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contractabi *ContractabiCallerSession) Owner() (common.Address, error) {
	return _Contractabi.Contract.Owner(&_Contractabi.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Contractabi *ContractabiCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contractabi.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Contractabi *ContractabiSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Contractabi.Contract.OwnerOf(&_Contractabi.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Contractabi *ContractabiCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Contractabi.Contract.OwnerOf(&_Contractabi.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Contractabi *ContractabiCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contractabi.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Contractabi *ContractabiSession) Paused() (bool, error) {
	return _Contractabi.Contract.Paused(&_Contractabi.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Contractabi *ContractabiCallerSession) Paused() (bool, error) {
	return _Contractabi.Contract.Paused(&_Contractabi.CallOpts)
}

// PriceInWei is a free data retrieval call binding the contract method 0x3c8da588.
//
// Solidity: function priceInWei() view returns(uint256)
func (_Contractabi *ContractabiCaller) PriceInWei(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contractabi.contract.Call(opts, &out, "priceInWei")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceInWei is a free data retrieval call binding the contract method 0x3c8da588.
//
// Solidity: function priceInWei() view returns(uint256)
func (_Contractabi *ContractabiSession) PriceInWei() (*big.Int, error) {
	return _Contractabi.Contract.PriceInWei(&_Contractabi.CallOpts)
}

// PriceInWei is a free data retrieval call binding the contract method 0x3c8da588.
//
// Solidity: function priceInWei() view returns(uint256)
func (_Contractabi *ContractabiCallerSession) PriceInWei() (*big.Int, error) {
	return _Contractabi.Contract.PriceInWei(&_Contractabi.CallOpts)
}

// PriceWhitelist is a free data retrieval call binding the contract method 0x2fff1796.
//
// Solidity: function priceWhitelist() view returns(uint256)
func (_Contractabi *ContractabiCaller) PriceWhitelist(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contractabi.contract.Call(opts, &out, "priceWhitelist")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceWhitelist is a free data retrieval call binding the contract method 0x2fff1796.
//
// Solidity: function priceWhitelist() view returns(uint256)
func (_Contractabi *ContractabiSession) PriceWhitelist() (*big.Int, error) {
	return _Contractabi.Contract.PriceWhitelist(&_Contractabi.CallOpts)
}

// PriceWhitelist is a free data retrieval call binding the contract method 0x2fff1796.
//
// Solidity: function priceWhitelist() view returns(uint256)
func (_Contractabi *ContractabiCallerSession) PriceWhitelist() (*big.Int, error) {
	return _Contractabi.Contract.PriceWhitelist(&_Contractabi.CallOpts)
}

// Revealed is a free data retrieval call binding the contract method 0x51830227.
//
// Solidity: function revealed() view returns(bool)
func (_Contractabi *ContractabiCaller) Revealed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contractabi.contract.Call(opts, &out, "revealed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Revealed is a free data retrieval call binding the contract method 0x51830227.
//
// Solidity: function revealed() view returns(bool)
func (_Contractabi *ContractabiSession) Revealed() (bool, error) {
	return _Contractabi.Contract.Revealed(&_Contractabi.CallOpts)
}

// Revealed is a free data retrieval call binding the contract method 0x51830227.
//
// Solidity: function revealed() view returns(bool)
func (_Contractabi *ContractabiCallerSession) Revealed() (bool, error) {
	return _Contractabi.Contract.Revealed(&_Contractabi.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contractabi *ContractabiCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Contractabi.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contractabi *ContractabiSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Contractabi.Contract.SupportsInterface(&_Contractabi.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contractabi *ContractabiCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Contractabi.Contract.SupportsInterface(&_Contractabi.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contractabi *ContractabiCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contractabi.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contractabi *ContractabiSession) Symbol() (string, error) {
	return _Contractabi.Contract.Symbol(&_Contractabi.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contractabi *ContractabiCallerSession) Symbol() (string, error) {
	return _Contractabi.Contract.Symbol(&_Contractabi.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Contractabi *ContractabiCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contractabi.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Contractabi *ContractabiSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Contractabi.Contract.TokenByIndex(&_Contractabi.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Contractabi *ContractabiCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Contractabi.Contract.TokenByIndex(&_Contractabi.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256 tokenId)
func (_Contractabi *ContractabiCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contractabi.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256 tokenId)
func (_Contractabi *ContractabiSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Contractabi.Contract.TokenOfOwnerByIndex(&_Contractabi.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256 tokenId)
func (_Contractabi *ContractabiCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Contractabi.Contract.TokenOfOwnerByIndex(&_Contractabi.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Contractabi *ContractabiCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Contractabi.contract.Call(opts, &out, "tokenURI", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Contractabi *ContractabiSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Contractabi.Contract.TokenURI(&_Contractabi.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Contractabi *ContractabiCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Contractabi.Contract.TokenURI(&_Contractabi.CallOpts, _tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contractabi *ContractabiCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contractabi.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contractabi *ContractabiSession) TotalSupply() (*big.Int, error) {
	return _Contractabi.Contract.TotalSupply(&_Contractabi.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contractabi *ContractabiCallerSession) TotalSupply() (*big.Int, error) {
	return _Contractabi.Contract.TotalSupply(&_Contractabi.CallOpts)
}

// WalletOfOwner is a free data retrieval call binding the contract method 0x438b6300.
//
// Solidity: function walletOfOwner(address _owner) view returns(uint256[])
func (_Contractabi *ContractabiCaller) WalletOfOwner(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Contractabi.contract.Call(opts, &out, "walletOfOwner", _owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// WalletOfOwner is a free data retrieval call binding the contract method 0x438b6300.
//
// Solidity: function walletOfOwner(address _owner) view returns(uint256[])
func (_Contractabi *ContractabiSession) WalletOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _Contractabi.Contract.WalletOfOwner(&_Contractabi.CallOpts, _owner)
}

// WalletOfOwner is a free data retrieval call binding the contract method 0x438b6300.
//
// Solidity: function walletOfOwner(address _owner) view returns(uint256[])
func (_Contractabi *ContractabiCallerSession) WalletOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _Contractabi.Contract.WalletOfOwner(&_Contractabi.CallOpts, _owner)
}

// WhitelistActive is a free data retrieval call binding the contract method 0x02ce5813.
//
// Solidity: function whitelistActive() view returns(bool)
func (_Contractabi *ContractabiCaller) WhitelistActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contractabi.contract.Call(opts, &out, "whitelistActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistActive is a free data retrieval call binding the contract method 0x02ce5813.
//
// Solidity: function whitelistActive() view returns(bool)
func (_Contractabi *ContractabiSession) WhitelistActive() (bool, error) {
	return _Contractabi.Contract.WhitelistActive(&_Contractabi.CallOpts)
}

// WhitelistActive is a free data retrieval call binding the contract method 0x02ce5813.
//
// Solidity: function whitelistActive() view returns(bool)
func (_Contractabi *ContractabiCallerSession) WhitelistActive() (bool, error) {
	return _Contractabi.Contract.WhitelistActive(&_Contractabi.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Contractabi *ContractabiTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contractabi.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Contractabi *ContractabiSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contractabi.Contract.Approve(&_Contractabi.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Contractabi *ContractabiTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contractabi.Contract.Approve(&_Contractabi.TransactOpts, to, tokenId)
}

// BatchSafeTransferFrom is a paid mutator transaction binding the contract method 0x5a4fee30.
//
// Solidity: function batchSafeTransferFrom(address _from, address _to, uint256[] _tokenIds, bytes data_) returns()
func (_Contractabi *ContractabiTransactor) BatchSafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenIds []*big.Int, data_ []byte) (*types.Transaction, error) {
	return _Contractabi.contract.Transact(opts, "batchSafeTransferFrom", _from, _to, _tokenIds, data_)
}

// BatchSafeTransferFrom is a paid mutator transaction binding the contract method 0x5a4fee30.
//
// Solidity: function batchSafeTransferFrom(address _from, address _to, uint256[] _tokenIds, bytes data_) returns()
func (_Contractabi *ContractabiSession) BatchSafeTransferFrom(_from common.Address, _to common.Address, _tokenIds []*big.Int, data_ []byte) (*types.Transaction, error) {
	return _Contractabi.Contract.BatchSafeTransferFrom(&_Contractabi.TransactOpts, _from, _to, _tokenIds, data_)
}

// BatchSafeTransferFrom is a paid mutator transaction binding the contract method 0x5a4fee30.
//
// Solidity: function batchSafeTransferFrom(address _from, address _to, uint256[] _tokenIds, bytes data_) returns()
func (_Contractabi *ContractabiTransactorSession) BatchSafeTransferFrom(_from common.Address, _to common.Address, _tokenIds []*big.Int, data_ []byte) (*types.Transaction, error) {
	return _Contractabi.Contract.BatchSafeTransferFrom(&_Contractabi.TransactOpts, _from, _to, _tokenIds, data_)
}

// BatchTransferFrom is a paid mutator transaction binding the contract method 0xf3993d11.
//
// Solidity: function batchTransferFrom(address _from, address _to, uint256[] _tokenIds) returns()
func (_Contractabi *ContractabiTransactor) BatchTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _Contractabi.contract.Transact(opts, "batchTransferFrom", _from, _to, _tokenIds)
}

// BatchTransferFrom is a paid mutator transaction binding the contract method 0xf3993d11.
//
// Solidity: function batchTransferFrom(address _from, address _to, uint256[] _tokenIds) returns()
func (_Contractabi *ContractabiSession) BatchTransferFrom(_from common.Address, _to common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _Contractabi.Contract.BatchTransferFrom(&_Contractabi.TransactOpts, _from, _to, _tokenIds)
}

// BatchTransferFrom is a paid mutator transaction binding the contract method 0xf3993d11.
//
// Solidity: function batchTransferFrom(address _from, address _to, uint256[] _tokenIds) returns()
func (_Contractabi *ContractabiTransactorSession) BatchTransferFrom(_from common.Address, _to common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _Contractabi.Contract.BatchTransferFrom(&_Contractabi.TransactOpts, _from, _to, _tokenIds)
}

// CollectReserves is a paid mutator transaction binding the contract method 0x029877b6.
//
// Solidity: function collectReserves() returns()
func (_Contractabi *ContractabiTransactor) CollectReserves(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contractabi.contract.Transact(opts, "collectReserves")
}

// CollectReserves is a paid mutator transaction binding the contract method 0x029877b6.
//
// Solidity: function collectReserves() returns()
func (_Contractabi *ContractabiSession) CollectReserves() (*types.Transaction, error) {
	return _Contractabi.Contract.CollectReserves(&_Contractabi.TransactOpts)
}

// CollectReserves is a paid mutator transaction binding the contract method 0x029877b6.
//
// Solidity: function collectReserves() returns()
func (_Contractabi *ContractabiTransactorSession) CollectReserves() (*types.Transaction, error) {
	return _Contractabi.Contract.CollectReserves(&_Contractabi.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 count) payable returns()
func (_Contractabi *ContractabiTransactor) Mint(opts *bind.TransactOpts, count *big.Int) (*types.Transaction, error) {
	return _Contractabi.contract.Transact(opts, "mint", count)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 count) payable returns()
func (_Contractabi *ContractabiSession) Mint(count *big.Int) (*types.Transaction, error) {
	return _Contractabi.Contract.Mint(&_Contractabi.TransactOpts, count)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 count) payable returns()
func (_Contractabi *ContractabiTransactorSession) Mint(count *big.Int) (*types.Transaction, error) {
	return _Contractabi.Contract.Mint(&_Contractabi.TransactOpts, count)
}

// Pause is a paid mutator transaction binding the contract method 0x02329a29.
//
// Solidity: function pause(bool _state) returns()
func (_Contractabi *ContractabiTransactor) Pause(opts *bind.TransactOpts, _state bool) (*types.Transaction, error) {
	return _Contractabi.contract.Transact(opts, "pause", _state)
}

// Pause is a paid mutator transaction binding the contract method 0x02329a29.
//
// Solidity: function pause(bool _state) returns()
func (_Contractabi *ContractabiSession) Pause(_state bool) (*types.Transaction, error) {
	return _Contractabi.Contract.Pause(&_Contractabi.TransactOpts, _state)
}

// Pause is a paid mutator transaction binding the contract method 0x02329a29.
//
// Solidity: function pause(bool _state) returns()
func (_Contractabi *ContractabiTransactorSession) Pause(_state bool) (*types.Transaction, error) {
	return _Contractabi.Contract.Pause(&_Contractabi.TransactOpts, _state)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contractabi *ContractabiTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contractabi.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contractabi *ContractabiSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contractabi.Contract.RenounceOwnership(&_Contractabi.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contractabi *ContractabiTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contractabi.Contract.RenounceOwnership(&_Contractabi.TransactOpts)
}

// Reveal is a paid mutator transaction binding the contract method 0x4c261247.
//
// Solidity: function reveal(string revealedURI) returns()
func (_Contractabi *ContractabiTransactor) Reveal(opts *bind.TransactOpts, revealedURI string) (*types.Transaction, error) {
	return _Contractabi.contract.Transact(opts, "reveal", revealedURI)
}

// Reveal is a paid mutator transaction binding the contract method 0x4c261247.
//
// Solidity: function reveal(string revealedURI) returns()
func (_Contractabi *ContractabiSession) Reveal(revealedURI string) (*types.Transaction, error) {
	return _Contractabi.Contract.Reveal(&_Contractabi.TransactOpts, revealedURI)
}

// Reveal is a paid mutator transaction binding the contract method 0x4c261247.
//
// Solidity: function reveal(string revealedURI) returns()
func (_Contractabi *ContractabiTransactorSession) Reveal(revealedURI string) (*types.Transaction, error) {
	return _Contractabi.Contract.Reveal(&_Contractabi.TransactOpts, revealedURI)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Contractabi *ContractabiTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contractabi.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Contractabi *ContractabiSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contractabi.Contract.SafeTransferFrom(&_Contractabi.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Contractabi *ContractabiTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contractabi.Contract.SafeTransferFrom(&_Contractabi.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Contractabi *ContractabiTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Contractabi.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Contractabi *ContractabiSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Contractabi.Contract.SafeTransferFrom0(&_Contractabi.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Contractabi *ContractabiTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Contractabi.Contract.SafeTransferFrom0(&_Contractabi.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Contractabi *ContractabiTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Contractabi.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Contractabi *ContractabiSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Contractabi.Contract.SetApprovalForAll(&_Contractabi.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Contractabi *ContractabiTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Contractabi.Contract.SetApprovalForAll(&_Contractabi.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _baseURI) returns()
func (_Contractabi *ContractabiTransactor) SetBaseURI(opts *bind.TransactOpts, _baseURI string) (*types.Transaction, error) {
	return _Contractabi.contract.Transact(opts, "setBaseURI", _baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _baseURI) returns()
func (_Contractabi *ContractabiSession) SetBaseURI(_baseURI string) (*types.Transaction, error) {
	return _Contractabi.Contract.SetBaseURI(&_Contractabi.TransactOpts, _baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _baseURI) returns()
func (_Contractabi *ContractabiTransactorSession) SetBaseURI(_baseURI string) (*types.Transaction, error) {
	return _Contractabi.Contract.SetBaseURI(&_Contractabi.TransactOpts, _baseURI)
}

// SetExtension is a paid mutator transaction binding the contract method 0x7e2285aa.
//
// Solidity: function setExtension(string _extension) returns()
func (_Contractabi *ContractabiTransactor) SetExtension(opts *bind.TransactOpts, _extension string) (*types.Transaction, error) {
	return _Contractabi.contract.Transact(opts, "setExtension", _extension)
}

// SetExtension is a paid mutator transaction binding the contract method 0x7e2285aa.
//
// Solidity: function setExtension(string _extension) returns()
func (_Contractabi *ContractabiSession) SetExtension(_extension string) (*types.Transaction, error) {
	return _Contractabi.Contract.SetExtension(&_Contractabi.TransactOpts, _extension)
}

// SetExtension is a paid mutator transaction binding the contract method 0x7e2285aa.
//
// Solidity: function setExtension(string _extension) returns()
func (_Contractabi *ContractabiTransactorSession) SetExtension(_extension string) (*types.Transaction, error) {
	return _Contractabi.Contract.SetExtension(&_Contractabi.TransactOpts, _extension)
}

// SetMaxPerTx is a paid mutator transaction binding the contract method 0xc6f6f216.
//
// Solidity: function setMaxPerTx(uint256 _max_per_tx) returns()
func (_Contractabi *ContractabiTransactor) SetMaxPerTx(opts *bind.TransactOpts, _max_per_tx *big.Int) (*types.Transaction, error) {
	return _Contractabi.contract.Transact(opts, "setMaxPerTx", _max_per_tx)
}

// SetMaxPerTx is a paid mutator transaction binding the contract method 0xc6f6f216.
//
// Solidity: function setMaxPerTx(uint256 _max_per_tx) returns()
func (_Contractabi *ContractabiSession) SetMaxPerTx(_max_per_tx *big.Int) (*types.Transaction, error) {
	return _Contractabi.Contract.SetMaxPerTx(&_Contractabi.TransactOpts, _max_per_tx)
}

// SetMaxPerTx is a paid mutator transaction binding the contract method 0xc6f6f216.
//
// Solidity: function setMaxPerTx(uint256 _max_per_tx) returns()
func (_Contractabi *ContractabiTransactorSession) SetMaxPerTx(_max_per_tx *big.Int) (*types.Transaction, error) {
	return _Contractabi.Contract.SetMaxPerTx(&_Contractabi.TransactOpts, _max_per_tx)
}

// SetMaxPerWallet is a paid mutator transaction binding the contract method 0xe268e4d3.
//
// Solidity: function setMaxPerWallet(uint256 _max_per_wallet) returns()
func (_Contractabi *ContractabiTransactor) SetMaxPerWallet(opts *bind.TransactOpts, _max_per_wallet *big.Int) (*types.Transaction, error) {
	return _Contractabi.contract.Transact(opts, "setMaxPerWallet", _max_per_wallet)
}

// SetMaxPerWallet is a paid mutator transaction binding the contract method 0xe268e4d3.
//
// Solidity: function setMaxPerWallet(uint256 _max_per_wallet) returns()
func (_Contractabi *ContractabiSession) SetMaxPerWallet(_max_per_wallet *big.Int) (*types.Transaction, error) {
	return _Contractabi.Contract.SetMaxPerWallet(&_Contractabi.TransactOpts, _max_per_wallet)
}

// SetMaxPerWallet is a paid mutator transaction binding the contract method 0xe268e4d3.
//
// Solidity: function setMaxPerWallet(uint256 _max_per_wallet) returns()
func (_Contractabi *ContractabiTransactorSession) SetMaxPerWallet(_max_per_wallet *big.Int) (*types.Transaction, error) {
	return _Contractabi.Contract.SetMaxPerWallet(&_Contractabi.TransactOpts, _max_per_wallet)
}

// SetMaxWhitelist is a paid mutator transaction binding the contract method 0x72d9eb1e.
//
// Solidity: function setMaxWhitelist(uint256 _max_whitelist) returns()
func (_Contractabi *ContractabiTransactor) SetMaxWhitelist(opts *bind.TransactOpts, _max_whitelist *big.Int) (*types.Transaction, error) {
	return _Contractabi.contract.Transact(opts, "setMaxWhitelist", _max_whitelist)
}

// SetMaxWhitelist is a paid mutator transaction binding the contract method 0x72d9eb1e.
//
// Solidity: function setMaxWhitelist(uint256 _max_whitelist) returns()
func (_Contractabi *ContractabiSession) SetMaxWhitelist(_max_whitelist *big.Int) (*types.Transaction, error) {
	return _Contractabi.Contract.SetMaxWhitelist(&_Contractabi.TransactOpts, _max_whitelist)
}

// SetMaxWhitelist is a paid mutator transaction binding the contract method 0x72d9eb1e.
//
// Solidity: function setMaxWhitelist(uint256 _max_whitelist) returns()
func (_Contractabi *ContractabiTransactorSession) SetMaxWhitelist(_max_whitelist *big.Int) (*types.Transaction, error) {
	return _Contractabi.Contract.SetMaxWhitelist(&_Contractabi.TransactOpts, _max_whitelist)
}

// SetPriceInWei is a paid mutator transaction binding the contract method 0x8774e5d0.
//
// Solidity: function setPriceInWei(uint256 _price) returns()
func (_Contractabi *ContractabiTransactor) SetPriceInWei(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _Contractabi.contract.Transact(opts, "setPriceInWei", _price)
}

// SetPriceInWei is a paid mutator transaction binding the contract method 0x8774e5d0.
//
// Solidity: function setPriceInWei(uint256 _price) returns()
func (_Contractabi *ContractabiSession) SetPriceInWei(_price *big.Int) (*types.Transaction, error) {
	return _Contractabi.Contract.SetPriceInWei(&_Contractabi.TransactOpts, _price)
}

// SetPriceInWei is a paid mutator transaction binding the contract method 0x8774e5d0.
//
// Solidity: function setPriceInWei(uint256 _price) returns()
func (_Contractabi *ContractabiTransactorSession) SetPriceInWei(_price *big.Int) (*types.Transaction, error) {
	return _Contractabi.Contract.SetPriceInWei(&_Contractabi.TransactOpts, _price)
}

// SetPriceWhitelist is a paid mutator transaction binding the contract method 0xb95f8b5f.
//
// Solidity: function setPriceWhitelist(uint256 _price) returns()
func (_Contractabi *ContractabiTransactor) SetPriceWhitelist(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _Contractabi.contract.Transact(opts, "setPriceWhitelist", _price)
}

// SetPriceWhitelist is a paid mutator transaction binding the contract method 0xb95f8b5f.
//
// Solidity: function setPriceWhitelist(uint256 _price) returns()
func (_Contractabi *ContractabiSession) SetPriceWhitelist(_price *big.Int) (*types.Transaction, error) {
	return _Contractabi.Contract.SetPriceWhitelist(&_Contractabi.TransactOpts, _price)
}

// SetPriceWhitelist is a paid mutator transaction binding the contract method 0xb95f8b5f.
//
// Solidity: function setPriceWhitelist(uint256 _price) returns()
func (_Contractabi *ContractabiTransactorSession) SetPriceWhitelist(_price *big.Int) (*types.Transaction, error) {
	return _Contractabi.Contract.SetPriceWhitelist(&_Contractabi.TransactOpts, _price)
}

// SetWhitelistActive is a paid mutator transaction binding the contract method 0xc3b754dc.
//
// Solidity: function setWhitelistActive(bool _whitelistActive) returns()
func (_Contractabi *ContractabiTransactor) SetWhitelistActive(opts *bind.TransactOpts, _whitelistActive bool) (*types.Transaction, error) {
	return _Contractabi.contract.Transact(opts, "setWhitelistActive", _whitelistActive)
}

// SetWhitelistActive is a paid mutator transaction binding the contract method 0xc3b754dc.
//
// Solidity: function setWhitelistActive(bool _whitelistActive) returns()
func (_Contractabi *ContractabiSession) SetWhitelistActive(_whitelistActive bool) (*types.Transaction, error) {
	return _Contractabi.Contract.SetWhitelistActive(&_Contractabi.TransactOpts, _whitelistActive)
}

// SetWhitelistActive is a paid mutator transaction binding the contract method 0xc3b754dc.
//
// Solidity: function setWhitelistActive(bool _whitelistActive) returns()
func (_Contractabi *ContractabiTransactorSession) SetWhitelistActive(_whitelistActive bool) (*types.Transaction, error) {
	return _Contractabi.Contract.SetWhitelistActive(&_Contractabi.TransactOpts, _whitelistActive)
}

// SetWhitelistMerkleRoot is a paid mutator transaction binding the contract method 0xbd32fb66.
//
// Solidity: function setWhitelistMerkleRoot(bytes32 _whitelistMerkleRoot) returns()
func (_Contractabi *ContractabiTransactor) SetWhitelistMerkleRoot(opts *bind.TransactOpts, _whitelistMerkleRoot [32]byte) (*types.Transaction, error) {
	return _Contractabi.contract.Transact(opts, "setWhitelistMerkleRoot", _whitelistMerkleRoot)
}

// SetWhitelistMerkleRoot is a paid mutator transaction binding the contract method 0xbd32fb66.
//
// Solidity: function setWhitelistMerkleRoot(bytes32 _whitelistMerkleRoot) returns()
func (_Contractabi *ContractabiSession) SetWhitelistMerkleRoot(_whitelistMerkleRoot [32]byte) (*types.Transaction, error) {
	return _Contractabi.Contract.SetWhitelistMerkleRoot(&_Contractabi.TransactOpts, _whitelistMerkleRoot)
}

// SetWhitelistMerkleRoot is a paid mutator transaction binding the contract method 0xbd32fb66.
//
// Solidity: function setWhitelistMerkleRoot(bytes32 _whitelistMerkleRoot) returns()
func (_Contractabi *ContractabiTransactorSession) SetWhitelistMerkleRoot(_whitelistMerkleRoot [32]byte) (*types.Transaction, error) {
	return _Contractabi.Contract.SetWhitelistMerkleRoot(&_Contractabi.TransactOpts, _whitelistMerkleRoot)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Contractabi *ContractabiTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contractabi.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Contractabi *ContractabiSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contractabi.Contract.TransferFrom(&_Contractabi.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Contractabi *ContractabiTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contractabi.Contract.TransferFrom(&_Contractabi.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contractabi *ContractabiTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contractabi.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contractabi *ContractabiSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contractabi.Contract.TransferOwnership(&_Contractabi.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contractabi *ContractabiTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contractabi.Contract.TransferOwnership(&_Contractabi.TransactOpts, newOwner)
}

// WhitelistMint is a paid mutator transaction binding the contract method 0xc4be5b59.
//
// Solidity: function whitelistMint(uint256 count, uint256 dummy, bytes32[] proof) payable returns()
func (_Contractabi *ContractabiTransactor) WhitelistMint(opts *bind.TransactOpts, count *big.Int, dummy *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Contractabi.contract.Transact(opts, "whitelistMint", count, dummy, proof)
}

// WhitelistMint is a paid mutator transaction binding the contract method 0xc4be5b59.
//
// Solidity: function whitelistMint(uint256 count, uint256 dummy, bytes32[] proof) payable returns()
func (_Contractabi *ContractabiSession) WhitelistMint(count *big.Int, dummy *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Contractabi.Contract.WhitelistMint(&_Contractabi.TransactOpts, count, dummy, proof)
}

// WhitelistMint is a paid mutator transaction binding the contract method 0xc4be5b59.
//
// Solidity: function whitelistMint(uint256 count, uint256 dummy, bytes32[] proof) payable returns()
func (_Contractabi *ContractabiTransactorSession) WhitelistMint(count *big.Int, dummy *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Contractabi.Contract.WhitelistMint(&_Contractabi.TransactOpts, count, dummy, proof)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_Contractabi *ContractabiTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contractabi.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_Contractabi *ContractabiSession) Withdraw() (*types.Transaction, error) {
	return _Contractabi.Contract.Withdraw(&_Contractabi.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_Contractabi *ContractabiTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Contractabi.Contract.Withdraw(&_Contractabi.TransactOpts)
}

// ContractabiApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Contractabi contract.
type ContractabiApprovalIterator struct {
	Event *ContractabiApproval // Event containing the contract specifics and raw log

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
func (it *ContractabiApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractabiApproval)
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
		it.Event = new(ContractabiApproval)
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
func (it *ContractabiApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractabiApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractabiApproval represents a Approval event raised by the Contractabi contract.
type ContractabiApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Contractabi *ContractabiFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ContractabiApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contractabi.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractabiApprovalIterator{contract: _Contractabi.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Contractabi *ContractabiFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ContractabiApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contractabi.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractabiApproval)
				if err := _Contractabi.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Contractabi *ContractabiFilterer) ParseApproval(log types.Log) (*ContractabiApproval, error) {
	event := new(ContractabiApproval)
	if err := _Contractabi.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractabiApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Contractabi contract.
type ContractabiApprovalForAllIterator struct {
	Event *ContractabiApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ContractabiApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractabiApprovalForAll)
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
		it.Event = new(ContractabiApprovalForAll)
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
func (it *ContractabiApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractabiApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractabiApprovalForAll represents a ApprovalForAll event raised by the Contractabi contract.
type ContractabiApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Contractabi *ContractabiFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ContractabiApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Contractabi.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractabiApprovalForAllIterator{contract: _Contractabi.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Contractabi *ContractabiFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ContractabiApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Contractabi.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractabiApprovalForAll)
				if err := _Contractabi.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Contractabi *ContractabiFilterer) ParseApprovalForAll(log types.Log) (*ContractabiApprovalForAll, error) {
	event := new(ContractabiApprovalForAll)
	if err := _Contractabi.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractabiOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contractabi contract.
type ContractabiOwnershipTransferredIterator struct {
	Event *ContractabiOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractabiOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractabiOwnershipTransferred)
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
		it.Event = new(ContractabiOwnershipTransferred)
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
func (it *ContractabiOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractabiOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractabiOwnershipTransferred represents a OwnershipTransferred event raised by the Contractabi contract.
type ContractabiOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contractabi *ContractabiFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractabiOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contractabi.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractabiOwnershipTransferredIterator{contract: _Contractabi.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contractabi *ContractabiFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractabiOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contractabi.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractabiOwnershipTransferred)
				if err := _Contractabi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contractabi *ContractabiFilterer) ParseOwnershipTransferred(log types.Log) (*ContractabiOwnershipTransferred, error) {
	event := new(ContractabiOwnershipTransferred)
	if err := _Contractabi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractabiTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Contractabi contract.
type ContractabiTransferIterator struct {
	Event *ContractabiTransfer // Event containing the contract specifics and raw log

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
func (it *ContractabiTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractabiTransfer)
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
		it.Event = new(ContractabiTransfer)
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
func (it *ContractabiTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractabiTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractabiTransfer represents a Transfer event raised by the Contractabi contract.
type ContractabiTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Contractabi *ContractabiFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ContractabiTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contractabi.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractabiTransferIterator{contract: _Contractabi.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Contractabi *ContractabiFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ContractabiTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contractabi.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractabiTransfer)
				if err := _Contractabi.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Contractabi *ContractabiFilterer) ParseTransfer(log types.Log) (*ContractabiTransfer, error) {
	event := new(ContractabiTransfer)
	if err := _Contractabi.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
