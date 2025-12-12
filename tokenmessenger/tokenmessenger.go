// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokenmessenger

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = bytes.Equal
	_ = errors.New
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = abi.ConvertType
)

// TokenMessengerV2TokenMessengerV2Roles is an auto generated low-level Go binding around an user-defined struct.
type TokenMessengerV2TokenMessengerV2Roles struct {
	Owner            common.Address
	Rescuer          common.Address
	FeeRecipient     common.Address
	Denylister       common.Address
	TokenMinter      common.Address
	MinFeeController common.Address
}

// AddressMetaData contains all meta data concerning the Address contract.
var AddressMetaData = bind.MetaData{
	ABI: "[]",
	ID:  "864c39d6bc14f3cdf318269836e6f76017",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122039be6c42587447fefccc15e3b845c926751e3ab8ff8dd627d8b07309ea08837864736f6c63430007060033",
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	abi abi.ABI
}

// NewAddress creates a new instance of Address.
func NewAddress() *Address {
	parsed, err := AddressMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Address{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Address) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// AddressUtilsMetaData contains all meta data concerning the AddressUtils contract.
var AddressUtilsMetaData = bind.MetaData{
	ABI: "[]",
	ID:  "d9e77df3dd34cb371c86d0d3a808bb9d7d",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122036a3ae6640d4c10bd8b8ba8a537c928a7491e64144dea0fa914319cdf2b471fb64736f6c63430007060033",
}

// AddressUtils is an auto generated Go binding around an Ethereum contract.
type AddressUtils struct {
	abi abi.ABI
}

// NewAddressUtils creates a new instance of AddressUtils.
func NewAddressUtils() *AddressUtils {
	parsed, err := AddressUtilsMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &AddressUtils{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *AddressUtils) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// BaseTokenMessengerMetaData contains all meta data concerning the BaseTokenMessenger contract.
var BaseTokenMessengerMetaData = bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Denylisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldDenylister\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newDenylister\",\"type\":\"address\"}],\"name\":\"DenylisterChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"FeeRecipientSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"localMinter\",\"type\":\"address\"}],\"name\":\"LocalMinterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"localMinter\",\"type\":\"address\"}],\"name\":\"LocalMinterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minFeeController\",\"type\":\"address\"}],\"name\":\"MinFeeControllerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minFee\",\"type\":\"uint256\"}],\"name\":\"MinFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mintRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mintToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeCollected\",\"type\":\"uint256\"}],\"name\":\"MintAndWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"tokenMessenger\",\"type\":\"bytes32\"}],\"name\":\"RemoteTokenMessengerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"tokenMessenger\",\"type\":\"bytes32\"}],\"name\":\"RemoteTokenMessengerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newRescuer\",\"type\":\"address\"}],\"name\":\"RescuerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"UnDenylisted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MIN_FEE_MULTIPLIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newLocalMinter\",\"type\":\"address\"}],\"name\":\"addLocalMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"tokenMessenger\",\"type\":\"bytes32\"}],\"name\":\"addRemoteTokenMessenger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"denylist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denylister\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedVersion\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isDenylisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localMessageTransmitter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localMinter\",\"outputs\":[{\"internalType\":\"contractITokenMinterV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageBodyVersion\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minFeeController\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"remoteTokenMessengers\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeLocalMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"removeRemoteTokenMessenger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rescuer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeRecipient\",\"type\":\"address\"}],\"name\":\"setFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minFee\",\"type\":\"uint256\"}],\"name\":\"setMinFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minFeeController\",\"type\":\"address\"}],\"name\":\"setMinFeeController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"unDenylist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDenylister\",\"type\":\"address\"}],\"name\":\"updateDenylister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRescuer\",\"type\":\"address\"}],\"name\":\"updateRescuer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "f61bb739c8e8a54b3089bcdb98ceea1556",
}

// BaseTokenMessenger is an auto generated Go binding around an Ethereum contract.
type BaseTokenMessenger struct {
	abi abi.ABI
}

// NewBaseTokenMessenger creates a new instance of BaseTokenMessenger.
func NewBaseTokenMessenger() *BaseTokenMessenger {
	parsed, err := BaseTokenMessengerMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &BaseTokenMessenger{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *BaseTokenMessenger) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackMINFEEMULTIPLIER is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x966dfbd5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function MIN_FEE_MULTIPLIER() view returns(uint256)
func (baseTokenMessenger *BaseTokenMessenger) PackMINFEEMULTIPLIER() []byte {
	enc, err := baseTokenMessenger.abi.Pack("MIN_FEE_MULTIPLIER")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMINFEEMULTIPLIER is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x966dfbd5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function MIN_FEE_MULTIPLIER() view returns(uint256)
func (baseTokenMessenger *BaseTokenMessenger) TryPackMINFEEMULTIPLIER() ([]byte, error) {
	return baseTokenMessenger.abi.Pack("MIN_FEE_MULTIPLIER")
}

// UnpackMINFEEMULTIPLIER is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x966dfbd5.
//
// Solidity: function MIN_FEE_MULTIPLIER() view returns(uint256)
func (baseTokenMessenger *BaseTokenMessenger) UnpackMINFEEMULTIPLIER(data []byte) (*big.Int, error) {
	out, err := baseTokenMessenger.abi.Unpack("MIN_FEE_MULTIPLIER", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackAcceptOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x79ba5097.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function acceptOwnership() returns()
func (baseTokenMessenger *BaseTokenMessenger) PackAcceptOwnership() []byte {
	enc, err := baseTokenMessenger.abi.Pack("acceptOwnership")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAcceptOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x79ba5097.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function acceptOwnership() returns()
func (baseTokenMessenger *BaseTokenMessenger) TryPackAcceptOwnership() ([]byte, error) {
	return baseTokenMessenger.abi.Pack("acceptOwnership")
}

// PackAddLocalMinter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8197beb9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addLocalMinter(address newLocalMinter) returns()
func (baseTokenMessenger *BaseTokenMessenger) PackAddLocalMinter(newLocalMinter common.Address) []byte {
	enc, err := baseTokenMessenger.abi.Pack("addLocalMinter", newLocalMinter)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddLocalMinter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8197beb9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addLocalMinter(address newLocalMinter) returns()
func (baseTokenMessenger *BaseTokenMessenger) TryPackAddLocalMinter(newLocalMinter common.Address) ([]byte, error) {
	return baseTokenMessenger.abi.Pack("addLocalMinter", newLocalMinter)
}

// PackAddRemoteTokenMessenger is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xda87e448.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addRemoteTokenMessenger(uint32 domain, bytes32 tokenMessenger) returns()
func (baseTokenMessenger *BaseTokenMessenger) PackAddRemoteTokenMessenger(domain uint32, tokenMessenger [32]byte) []byte {
	enc, err := baseTokenMessenger.abi.Pack("addRemoteTokenMessenger", domain, tokenMessenger)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddRemoteTokenMessenger is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xda87e448.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addRemoteTokenMessenger(uint32 domain, bytes32 tokenMessenger) returns()
func (baseTokenMessenger *BaseTokenMessenger) TryPackAddRemoteTokenMessenger(domain uint32, tokenMessenger [32]byte) ([]byte, error) {
	return baseTokenMessenger.abi.Pack("addRemoteTokenMessenger", domain, tokenMessenger)
}

// PackDenylist is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3371bfff.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function denylist(address account) returns()
func (baseTokenMessenger *BaseTokenMessenger) PackDenylist(account common.Address) []byte {
	enc, err := baseTokenMessenger.abi.Pack("denylist", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDenylist is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3371bfff.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function denylist(address account) returns()
func (baseTokenMessenger *BaseTokenMessenger) TryPackDenylist(account common.Address) ([]byte, error) {
	return baseTokenMessenger.abi.Pack("denylist", account)
}

// PackDenylister is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbcc76c60.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function denylister() view returns(address)
func (baseTokenMessenger *BaseTokenMessenger) PackDenylister() []byte {
	enc, err := baseTokenMessenger.abi.Pack("denylister")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDenylister is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbcc76c60.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function denylister() view returns(address)
func (baseTokenMessenger *BaseTokenMessenger) TryPackDenylister() ([]byte, error) {
	return baseTokenMessenger.abi.Pack("denylister")
}

// UnpackDenylister is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbcc76c60.
//
// Solidity: function denylister() view returns(address)
func (baseTokenMessenger *BaseTokenMessenger) UnpackDenylister(data []byte) (common.Address, error) {
	out, err := baseTokenMessenger.abi.Unpack("denylister", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackFeeRecipient is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x46904840.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function feeRecipient() view returns(address)
func (baseTokenMessenger *BaseTokenMessenger) PackFeeRecipient() []byte {
	enc, err := baseTokenMessenger.abi.Pack("feeRecipient")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFeeRecipient is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x46904840.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function feeRecipient() view returns(address)
func (baseTokenMessenger *BaseTokenMessenger) TryPackFeeRecipient() ([]byte, error) {
	return baseTokenMessenger.abi.Pack("feeRecipient")
}

// UnpackFeeRecipient is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (baseTokenMessenger *BaseTokenMessenger) UnpackFeeRecipient(data []byte) (common.Address, error) {
	out, err := baseTokenMessenger.abi.Unpack("feeRecipient", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackInitializedVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x08828eb7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function initializedVersion() view returns(uint64)
func (baseTokenMessenger *BaseTokenMessenger) PackInitializedVersion() []byte {
	enc, err := baseTokenMessenger.abi.Pack("initializedVersion")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackInitializedVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x08828eb7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function initializedVersion() view returns(uint64)
func (baseTokenMessenger *BaseTokenMessenger) TryPackInitializedVersion() ([]byte, error) {
	return baseTokenMessenger.abi.Pack("initializedVersion")
}

// UnpackInitializedVersion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x08828eb7.
//
// Solidity: function initializedVersion() view returns(uint64)
func (baseTokenMessenger *BaseTokenMessenger) UnpackInitializedVersion(data []byte) (uint64, error) {
	out, err := baseTokenMessenger.abi.Unpack("initializedVersion", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackIsDenylisted is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe877a526.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isDenylisted(address account) view returns(bool)
func (baseTokenMessenger *BaseTokenMessenger) PackIsDenylisted(account common.Address) []byte {
	enc, err := baseTokenMessenger.abi.Pack("isDenylisted", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsDenylisted is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe877a526.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isDenylisted(address account) view returns(bool)
func (baseTokenMessenger *BaseTokenMessenger) TryPackIsDenylisted(account common.Address) ([]byte, error) {
	return baseTokenMessenger.abi.Pack("isDenylisted", account)
}

// UnpackIsDenylisted is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe877a526.
//
// Solidity: function isDenylisted(address account) view returns(bool)
func (baseTokenMessenger *BaseTokenMessenger) UnpackIsDenylisted(data []byte) (bool, error) {
	out, err := baseTokenMessenger.abi.Unpack("isDenylisted", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackLocalMessageTransmitter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2c121921.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function localMessageTransmitter() view returns(address)
func (baseTokenMessenger *BaseTokenMessenger) PackLocalMessageTransmitter() []byte {
	enc, err := baseTokenMessenger.abi.Pack("localMessageTransmitter")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackLocalMessageTransmitter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2c121921.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function localMessageTransmitter() view returns(address)
func (baseTokenMessenger *BaseTokenMessenger) TryPackLocalMessageTransmitter() ([]byte, error) {
	return baseTokenMessenger.abi.Pack("localMessageTransmitter")
}

// UnpackLocalMessageTransmitter is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2c121921.
//
// Solidity: function localMessageTransmitter() view returns(address)
func (baseTokenMessenger *BaseTokenMessenger) UnpackLocalMessageTransmitter(data []byte) (common.Address, error) {
	out, err := baseTokenMessenger.abi.Unpack("localMessageTransmitter", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackLocalMinter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcb75c11c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function localMinter() view returns(address)
func (baseTokenMessenger *BaseTokenMessenger) PackLocalMinter() []byte {
	enc, err := baseTokenMessenger.abi.Pack("localMinter")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackLocalMinter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcb75c11c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function localMinter() view returns(address)
func (baseTokenMessenger *BaseTokenMessenger) TryPackLocalMinter() ([]byte, error) {
	return baseTokenMessenger.abi.Pack("localMinter")
}

// UnpackLocalMinter is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcb75c11c.
//
// Solidity: function localMinter() view returns(address)
func (baseTokenMessenger *BaseTokenMessenger) UnpackLocalMinter(data []byte) (common.Address, error) {
	out, err := baseTokenMessenger.abi.Unpack("localMinter", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackMessageBodyVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9cdbb181.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function messageBodyVersion() view returns(uint32)
func (baseTokenMessenger *BaseTokenMessenger) PackMessageBodyVersion() []byte {
	enc, err := baseTokenMessenger.abi.Pack("messageBodyVersion")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMessageBodyVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9cdbb181.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function messageBodyVersion() view returns(uint32)
func (baseTokenMessenger *BaseTokenMessenger) TryPackMessageBodyVersion() ([]byte, error) {
	return baseTokenMessenger.abi.Pack("messageBodyVersion")
}

// UnpackMessageBodyVersion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9cdbb181.
//
// Solidity: function messageBodyVersion() view returns(uint32)
func (baseTokenMessenger *BaseTokenMessenger) UnpackMessageBodyVersion(data []byte) (uint32, error) {
	out, err := baseTokenMessenger.abi.Unpack("messageBodyVersion", data)
	if err != nil {
		return *new(uint32), err
	}
	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	return out0, nil
}

// PackMinFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x24ec7590.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function minFee() view returns(uint256)
func (baseTokenMessenger *BaseTokenMessenger) PackMinFee() []byte {
	enc, err := baseTokenMessenger.abi.Pack("minFee")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMinFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x24ec7590.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function minFee() view returns(uint256)
func (baseTokenMessenger *BaseTokenMessenger) TryPackMinFee() ([]byte, error) {
	return baseTokenMessenger.abi.Pack("minFee")
}

// UnpackMinFee is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x24ec7590.
//
// Solidity: function minFee() view returns(uint256)
func (baseTokenMessenger *BaseTokenMessenger) UnpackMinFee(data []byte) (*big.Int, error) {
	out, err := baseTokenMessenger.abi.Unpack("minFee", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackMinFeeController is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x369c4f1b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function minFeeController() view returns(address)
func (baseTokenMessenger *BaseTokenMessenger) PackMinFeeController() []byte {
	enc, err := baseTokenMessenger.abi.Pack("minFeeController")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMinFeeController is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x369c4f1b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function minFeeController() view returns(address)
func (baseTokenMessenger *BaseTokenMessenger) TryPackMinFeeController() ([]byte, error) {
	return baseTokenMessenger.abi.Pack("minFeeController")
}

// UnpackMinFeeController is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x369c4f1b.
//
// Solidity: function minFeeController() view returns(address)
func (baseTokenMessenger *BaseTokenMessenger) UnpackMinFeeController(data []byte) (common.Address, error) {
	out, err := baseTokenMessenger.abi.Unpack("minFeeController", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function owner() view returns(address)
func (baseTokenMessenger *BaseTokenMessenger) PackOwner() []byte {
	enc, err := baseTokenMessenger.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function owner() view returns(address)
func (baseTokenMessenger *BaseTokenMessenger) TryPackOwner() ([]byte, error) {
	return baseTokenMessenger.abi.Pack("owner")
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (baseTokenMessenger *BaseTokenMessenger) UnpackOwner(data []byte) (common.Address, error) {
	out, err := baseTokenMessenger.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackPendingOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe30c3978.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function pendingOwner() view returns(address)
func (baseTokenMessenger *BaseTokenMessenger) PackPendingOwner() []byte {
	enc, err := baseTokenMessenger.abi.Pack("pendingOwner")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPendingOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe30c3978.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function pendingOwner() view returns(address)
func (baseTokenMessenger *BaseTokenMessenger) TryPackPendingOwner() ([]byte, error) {
	return baseTokenMessenger.abi.Pack("pendingOwner")
}

// UnpackPendingOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (baseTokenMessenger *BaseTokenMessenger) UnpackPendingOwner(data []byte) (common.Address, error) {
	out, err := baseTokenMessenger.abi.Unpack("pendingOwner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackRemoteTokenMessengers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x82a5e665.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function remoteTokenMessengers(uint32 ) view returns(bytes32)
func (baseTokenMessenger *BaseTokenMessenger) PackRemoteTokenMessengers(arg0 uint32) []byte {
	enc, err := baseTokenMessenger.abi.Pack("remoteTokenMessengers", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoteTokenMessengers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x82a5e665.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function remoteTokenMessengers(uint32 ) view returns(bytes32)
func (baseTokenMessenger *BaseTokenMessenger) TryPackRemoteTokenMessengers(arg0 uint32) ([]byte, error) {
	return baseTokenMessenger.abi.Pack("remoteTokenMessengers", arg0)
}

// UnpackRemoteTokenMessengers is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x82a5e665.
//
// Solidity: function remoteTokenMessengers(uint32 ) view returns(bytes32)
func (baseTokenMessenger *BaseTokenMessenger) UnpackRemoteTokenMessengers(data []byte) ([32]byte, error) {
	out, err := baseTokenMessenger.abi.Unpack("remoteTokenMessengers", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackRemoveLocalMinter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91f17888.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeLocalMinter() returns()
func (baseTokenMessenger *BaseTokenMessenger) PackRemoveLocalMinter() []byte {
	enc, err := baseTokenMessenger.abi.Pack("removeLocalMinter")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveLocalMinter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91f17888.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeLocalMinter() returns()
func (baseTokenMessenger *BaseTokenMessenger) TryPackRemoveLocalMinter() ([]byte, error) {
	return baseTokenMessenger.abi.Pack("removeLocalMinter")
}

// PackRemoveRemoteTokenMessenger is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf79fd08e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeRemoteTokenMessenger(uint32 domain) returns()
func (baseTokenMessenger *BaseTokenMessenger) PackRemoveRemoteTokenMessenger(domain uint32) []byte {
	enc, err := baseTokenMessenger.abi.Pack("removeRemoteTokenMessenger", domain)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveRemoteTokenMessenger is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf79fd08e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeRemoteTokenMessenger(uint32 domain) returns()
func (baseTokenMessenger *BaseTokenMessenger) TryPackRemoveRemoteTokenMessenger(domain uint32) ([]byte, error) {
	return baseTokenMessenger.abi.Pack("removeRemoteTokenMessenger", domain)
}

// PackRescueERC20 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb2118a8d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function rescueERC20(address tokenContract, address to, uint256 amount) returns()
func (baseTokenMessenger *BaseTokenMessenger) PackRescueERC20(tokenContract common.Address, to common.Address, amount *big.Int) []byte {
	enc, err := baseTokenMessenger.abi.Pack("rescueERC20", tokenContract, to, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRescueERC20 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb2118a8d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function rescueERC20(address tokenContract, address to, uint256 amount) returns()
func (baseTokenMessenger *BaseTokenMessenger) TryPackRescueERC20(tokenContract common.Address, to common.Address, amount *big.Int) ([]byte, error) {
	return baseTokenMessenger.abi.Pack("rescueERC20", tokenContract, to, amount)
}

// PackRescuer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x38a63183.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function rescuer() view returns(address)
func (baseTokenMessenger *BaseTokenMessenger) PackRescuer() []byte {
	enc, err := baseTokenMessenger.abi.Pack("rescuer")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRescuer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x38a63183.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function rescuer() view returns(address)
func (baseTokenMessenger *BaseTokenMessenger) TryPackRescuer() ([]byte, error) {
	return baseTokenMessenger.abi.Pack("rescuer")
}

// UnpackRescuer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x38a63183.
//
// Solidity: function rescuer() view returns(address)
func (baseTokenMessenger *BaseTokenMessenger) UnpackRescuer(data []byte) (common.Address, error) {
	out, err := baseTokenMessenger.abi.Unpack("rescuer", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackSetFeeRecipient is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe74b981b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setFeeRecipient(address _feeRecipient) returns()
func (baseTokenMessenger *BaseTokenMessenger) PackSetFeeRecipient(feeRecipient common.Address) []byte {
	enc, err := baseTokenMessenger.abi.Pack("setFeeRecipient", feeRecipient)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetFeeRecipient is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe74b981b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setFeeRecipient(address _feeRecipient) returns()
func (baseTokenMessenger *BaseTokenMessenger) TryPackSetFeeRecipient(feeRecipient common.Address) ([]byte, error) {
	return baseTokenMessenger.abi.Pack("setFeeRecipient", feeRecipient)
}

// PackSetMinFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x31ac9920.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setMinFee(uint256 _minFee) returns()
func (baseTokenMessenger *BaseTokenMessenger) PackSetMinFee(minFee *big.Int) []byte {
	enc, err := baseTokenMessenger.abi.Pack("setMinFee", minFee)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetMinFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x31ac9920.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setMinFee(uint256 _minFee) returns()
func (baseTokenMessenger *BaseTokenMessenger) TryPackSetMinFee(minFee *big.Int) ([]byte, error) {
	return baseTokenMessenger.abi.Pack("setMinFee", minFee)
}

// PackSetMinFeeController is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa5b8d04e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setMinFeeController(address _minFeeController) returns()
func (baseTokenMessenger *BaseTokenMessenger) PackSetMinFeeController(minFeeController common.Address) []byte {
	enc, err := baseTokenMessenger.abi.Pack("setMinFeeController", minFeeController)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetMinFeeController is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa5b8d04e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setMinFeeController(address _minFeeController) returns()
func (baseTokenMessenger *BaseTokenMessenger) TryPackSetMinFeeController(minFeeController common.Address) ([]byte, error) {
	return baseTokenMessenger.abi.Pack("setMinFeeController", minFeeController)
}

// PackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (baseTokenMessenger *BaseTokenMessenger) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := baseTokenMessenger.abi.Pack("transferOwnership", newOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (baseTokenMessenger *BaseTokenMessenger) TryPackTransferOwnership(newOwner common.Address) ([]byte, error) {
	return baseTokenMessenger.abi.Pack("transferOwnership", newOwner)
}

// PackUnDenylist is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9cab0c1c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function unDenylist(address account) returns()
func (baseTokenMessenger *BaseTokenMessenger) PackUnDenylist(account common.Address) []byte {
	enc, err := baseTokenMessenger.abi.Pack("unDenylist", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUnDenylist is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9cab0c1c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function unDenylist(address account) returns()
func (baseTokenMessenger *BaseTokenMessenger) TryPackUnDenylist(account common.Address) ([]byte, error) {
	return baseTokenMessenger.abi.Pack("unDenylist", account)
}

// PackUpdateDenylister is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa946de04.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updateDenylister(address newDenylister) returns()
func (baseTokenMessenger *BaseTokenMessenger) PackUpdateDenylister(newDenylister common.Address) []byte {
	enc, err := baseTokenMessenger.abi.Pack("updateDenylister", newDenylister)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUpdateDenylister is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa946de04.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function updateDenylister(address newDenylister) returns()
func (baseTokenMessenger *BaseTokenMessenger) TryPackUpdateDenylister(newDenylister common.Address) ([]byte, error) {
	return baseTokenMessenger.abi.Pack("updateDenylister", newDenylister)
}

// PackUpdateRescuer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2ab60045.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updateRescuer(address newRescuer) returns()
func (baseTokenMessenger *BaseTokenMessenger) PackUpdateRescuer(newRescuer common.Address) []byte {
	enc, err := baseTokenMessenger.abi.Pack("updateRescuer", newRescuer)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUpdateRescuer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2ab60045.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function updateRescuer(address newRescuer) returns()
func (baseTokenMessenger *BaseTokenMessenger) TryPackUpdateRescuer(newRescuer common.Address) ([]byte, error) {
	return baseTokenMessenger.abi.Pack("updateRescuer", newRescuer)
}

// BaseTokenMessengerDenylisted represents a Denylisted event raised by the BaseTokenMessenger contract.
type BaseTokenMessengerDenylisted struct {
	Account common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const BaseTokenMessengerDenylistedEventName = "Denylisted"

// ContractEventName returns the user-defined event name.
func (BaseTokenMessengerDenylisted) ContractEventName() string {
	return BaseTokenMessengerDenylistedEventName
}

// UnpackDenylistedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Denylisted(address indexed account)
func (baseTokenMessenger *BaseTokenMessenger) UnpackDenylistedEvent(log *types.Log) (*BaseTokenMessengerDenylisted, error) {
	event := "Denylisted"
	if len(log.Topics) == 0 || log.Topics[0] != baseTokenMessenger.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BaseTokenMessengerDenylisted)
	if len(log.Data) > 0 {
		if err := baseTokenMessenger.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range baseTokenMessenger.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// BaseTokenMessengerDenylisterChanged represents a DenylisterChanged event raised by the BaseTokenMessenger contract.
type BaseTokenMessengerDenylisterChanged struct {
	OldDenylister common.Address
	NewDenylister common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const BaseTokenMessengerDenylisterChangedEventName = "DenylisterChanged"

// ContractEventName returns the user-defined event name.
func (BaseTokenMessengerDenylisterChanged) ContractEventName() string {
	return BaseTokenMessengerDenylisterChangedEventName
}

// UnpackDenylisterChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DenylisterChanged(address indexed oldDenylister, address indexed newDenylister)
func (baseTokenMessenger *BaseTokenMessenger) UnpackDenylisterChangedEvent(log *types.Log) (*BaseTokenMessengerDenylisterChanged, error) {
	event := "DenylisterChanged"
	if len(log.Topics) == 0 || log.Topics[0] != baseTokenMessenger.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BaseTokenMessengerDenylisterChanged)
	if len(log.Data) > 0 {
		if err := baseTokenMessenger.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range baseTokenMessenger.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// BaseTokenMessengerFeeRecipientSet represents a FeeRecipientSet event raised by the BaseTokenMessenger contract.
type BaseTokenMessengerFeeRecipientSet struct {
	FeeRecipient common.Address
	Raw          *types.Log // Blockchain specific contextual infos
}

const BaseTokenMessengerFeeRecipientSetEventName = "FeeRecipientSet"

// ContractEventName returns the user-defined event name.
func (BaseTokenMessengerFeeRecipientSet) ContractEventName() string {
	return BaseTokenMessengerFeeRecipientSetEventName
}

// UnpackFeeRecipientSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event FeeRecipientSet(address feeRecipient)
func (baseTokenMessenger *BaseTokenMessenger) UnpackFeeRecipientSetEvent(log *types.Log) (*BaseTokenMessengerFeeRecipientSet, error) {
	event := "FeeRecipientSet"
	if len(log.Topics) == 0 || log.Topics[0] != baseTokenMessenger.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BaseTokenMessengerFeeRecipientSet)
	if len(log.Data) > 0 {
		if err := baseTokenMessenger.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range baseTokenMessenger.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// BaseTokenMessengerInitialized represents a Initialized event raised by the BaseTokenMessenger contract.
type BaseTokenMessengerInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const BaseTokenMessengerInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (BaseTokenMessengerInitialized) ContractEventName() string {
	return BaseTokenMessengerInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (baseTokenMessenger *BaseTokenMessenger) UnpackInitializedEvent(log *types.Log) (*BaseTokenMessengerInitialized, error) {
	event := "Initialized"
	if len(log.Topics) == 0 || log.Topics[0] != baseTokenMessenger.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BaseTokenMessengerInitialized)
	if len(log.Data) > 0 {
		if err := baseTokenMessenger.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range baseTokenMessenger.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// BaseTokenMessengerLocalMinterAdded represents a LocalMinterAdded event raised by the BaseTokenMessenger contract.
type BaseTokenMessengerLocalMinterAdded struct {
	LocalMinter common.Address
	Raw         *types.Log // Blockchain specific contextual infos
}

const BaseTokenMessengerLocalMinterAddedEventName = "LocalMinterAdded"

// ContractEventName returns the user-defined event name.
func (BaseTokenMessengerLocalMinterAdded) ContractEventName() string {
	return BaseTokenMessengerLocalMinterAddedEventName
}

// UnpackLocalMinterAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event LocalMinterAdded(address localMinter)
func (baseTokenMessenger *BaseTokenMessenger) UnpackLocalMinterAddedEvent(log *types.Log) (*BaseTokenMessengerLocalMinterAdded, error) {
	event := "LocalMinterAdded"
	if len(log.Topics) == 0 || log.Topics[0] != baseTokenMessenger.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BaseTokenMessengerLocalMinterAdded)
	if len(log.Data) > 0 {
		if err := baseTokenMessenger.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range baseTokenMessenger.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// BaseTokenMessengerLocalMinterRemoved represents a LocalMinterRemoved event raised by the BaseTokenMessenger contract.
type BaseTokenMessengerLocalMinterRemoved struct {
	LocalMinter common.Address
	Raw         *types.Log // Blockchain specific contextual infos
}

const BaseTokenMessengerLocalMinterRemovedEventName = "LocalMinterRemoved"

// ContractEventName returns the user-defined event name.
func (BaseTokenMessengerLocalMinterRemoved) ContractEventName() string {
	return BaseTokenMessengerLocalMinterRemovedEventName
}

// UnpackLocalMinterRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event LocalMinterRemoved(address localMinter)
func (baseTokenMessenger *BaseTokenMessenger) UnpackLocalMinterRemovedEvent(log *types.Log) (*BaseTokenMessengerLocalMinterRemoved, error) {
	event := "LocalMinterRemoved"
	if len(log.Topics) == 0 || log.Topics[0] != baseTokenMessenger.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BaseTokenMessengerLocalMinterRemoved)
	if len(log.Data) > 0 {
		if err := baseTokenMessenger.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range baseTokenMessenger.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// BaseTokenMessengerMinFeeControllerSet represents a MinFeeControllerSet event raised by the BaseTokenMessenger contract.
type BaseTokenMessengerMinFeeControllerSet struct {
	MinFeeController common.Address
	Raw              *types.Log // Blockchain specific contextual infos
}

const BaseTokenMessengerMinFeeControllerSetEventName = "MinFeeControllerSet"

// ContractEventName returns the user-defined event name.
func (BaseTokenMessengerMinFeeControllerSet) ContractEventName() string {
	return BaseTokenMessengerMinFeeControllerSetEventName
}

// UnpackMinFeeControllerSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MinFeeControllerSet(address minFeeController)
func (baseTokenMessenger *BaseTokenMessenger) UnpackMinFeeControllerSetEvent(log *types.Log) (*BaseTokenMessengerMinFeeControllerSet, error) {
	event := "MinFeeControllerSet"
	if len(log.Topics) == 0 || log.Topics[0] != baseTokenMessenger.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BaseTokenMessengerMinFeeControllerSet)
	if len(log.Data) > 0 {
		if err := baseTokenMessenger.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range baseTokenMessenger.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// BaseTokenMessengerMinFeeSet represents a MinFeeSet event raised by the BaseTokenMessenger contract.
type BaseTokenMessengerMinFeeSet struct {
	MinFee *big.Int
	Raw    *types.Log // Blockchain specific contextual infos
}

const BaseTokenMessengerMinFeeSetEventName = "MinFeeSet"

// ContractEventName returns the user-defined event name.
func (BaseTokenMessengerMinFeeSet) ContractEventName() string {
	return BaseTokenMessengerMinFeeSetEventName
}

// UnpackMinFeeSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MinFeeSet(uint256 minFee)
func (baseTokenMessenger *BaseTokenMessenger) UnpackMinFeeSetEvent(log *types.Log) (*BaseTokenMessengerMinFeeSet, error) {
	event := "MinFeeSet"
	if len(log.Topics) == 0 || log.Topics[0] != baseTokenMessenger.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BaseTokenMessengerMinFeeSet)
	if len(log.Data) > 0 {
		if err := baseTokenMessenger.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range baseTokenMessenger.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// BaseTokenMessengerMintAndWithdraw represents a MintAndWithdraw event raised by the BaseTokenMessenger contract.
type BaseTokenMessengerMintAndWithdraw struct {
	MintRecipient common.Address
	Amount        *big.Int
	MintToken     common.Address
	FeeCollected  *big.Int
	Raw           *types.Log // Blockchain specific contextual infos
}

const BaseTokenMessengerMintAndWithdrawEventName = "MintAndWithdraw"

// ContractEventName returns the user-defined event name.
func (BaseTokenMessengerMintAndWithdraw) ContractEventName() string {
	return BaseTokenMessengerMintAndWithdrawEventName
}

// UnpackMintAndWithdrawEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MintAndWithdraw(address indexed mintRecipient, uint256 amount, address indexed mintToken, uint256 feeCollected)
func (baseTokenMessenger *BaseTokenMessenger) UnpackMintAndWithdrawEvent(log *types.Log) (*BaseTokenMessengerMintAndWithdraw, error) {
	event := "MintAndWithdraw"
	if len(log.Topics) == 0 || log.Topics[0] != baseTokenMessenger.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BaseTokenMessengerMintAndWithdraw)
	if len(log.Data) > 0 {
		if err := baseTokenMessenger.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range baseTokenMessenger.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// BaseTokenMessengerOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the BaseTokenMessenger contract.
type BaseTokenMessengerOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const BaseTokenMessengerOwnershipTransferStartedEventName = "OwnershipTransferStarted"

// ContractEventName returns the user-defined event name.
func (BaseTokenMessengerOwnershipTransferStarted) ContractEventName() string {
	return BaseTokenMessengerOwnershipTransferStartedEventName
}

// UnpackOwnershipTransferStartedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (baseTokenMessenger *BaseTokenMessenger) UnpackOwnershipTransferStartedEvent(log *types.Log) (*BaseTokenMessengerOwnershipTransferStarted, error) {
	event := "OwnershipTransferStarted"
	if len(log.Topics) == 0 || log.Topics[0] != baseTokenMessenger.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BaseTokenMessengerOwnershipTransferStarted)
	if len(log.Data) > 0 {
		if err := baseTokenMessenger.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range baseTokenMessenger.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// BaseTokenMessengerOwnershipTransferred represents a OwnershipTransferred event raised by the BaseTokenMessenger contract.
type BaseTokenMessengerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const BaseTokenMessengerOwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (BaseTokenMessengerOwnershipTransferred) ContractEventName() string {
	return BaseTokenMessengerOwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (baseTokenMessenger *BaseTokenMessenger) UnpackOwnershipTransferredEvent(log *types.Log) (*BaseTokenMessengerOwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if len(log.Topics) == 0 || log.Topics[0] != baseTokenMessenger.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BaseTokenMessengerOwnershipTransferred)
	if len(log.Data) > 0 {
		if err := baseTokenMessenger.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range baseTokenMessenger.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// BaseTokenMessengerRemoteTokenMessengerAdded represents a RemoteTokenMessengerAdded event raised by the BaseTokenMessenger contract.
type BaseTokenMessengerRemoteTokenMessengerAdded struct {
	Domain         uint32
	TokenMessenger [32]byte
	Raw            *types.Log // Blockchain specific contextual infos
}

const BaseTokenMessengerRemoteTokenMessengerAddedEventName = "RemoteTokenMessengerAdded"

// ContractEventName returns the user-defined event name.
func (BaseTokenMessengerRemoteTokenMessengerAdded) ContractEventName() string {
	return BaseTokenMessengerRemoteTokenMessengerAddedEventName
}

// UnpackRemoteTokenMessengerAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RemoteTokenMessengerAdded(uint32 domain, bytes32 tokenMessenger)
func (baseTokenMessenger *BaseTokenMessenger) UnpackRemoteTokenMessengerAddedEvent(log *types.Log) (*BaseTokenMessengerRemoteTokenMessengerAdded, error) {
	event := "RemoteTokenMessengerAdded"
	if len(log.Topics) == 0 || log.Topics[0] != baseTokenMessenger.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BaseTokenMessengerRemoteTokenMessengerAdded)
	if len(log.Data) > 0 {
		if err := baseTokenMessenger.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range baseTokenMessenger.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// BaseTokenMessengerRemoteTokenMessengerRemoved represents a RemoteTokenMessengerRemoved event raised by the BaseTokenMessenger contract.
type BaseTokenMessengerRemoteTokenMessengerRemoved struct {
	Domain         uint32
	TokenMessenger [32]byte
	Raw            *types.Log // Blockchain specific contextual infos
}

const BaseTokenMessengerRemoteTokenMessengerRemovedEventName = "RemoteTokenMessengerRemoved"

// ContractEventName returns the user-defined event name.
func (BaseTokenMessengerRemoteTokenMessengerRemoved) ContractEventName() string {
	return BaseTokenMessengerRemoteTokenMessengerRemovedEventName
}

// UnpackRemoteTokenMessengerRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RemoteTokenMessengerRemoved(uint32 domain, bytes32 tokenMessenger)
func (baseTokenMessenger *BaseTokenMessenger) UnpackRemoteTokenMessengerRemovedEvent(log *types.Log) (*BaseTokenMessengerRemoteTokenMessengerRemoved, error) {
	event := "RemoteTokenMessengerRemoved"
	if len(log.Topics) == 0 || log.Topics[0] != baseTokenMessenger.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BaseTokenMessengerRemoteTokenMessengerRemoved)
	if len(log.Data) > 0 {
		if err := baseTokenMessenger.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range baseTokenMessenger.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// BaseTokenMessengerRescuerChanged represents a RescuerChanged event raised by the BaseTokenMessenger contract.
type BaseTokenMessengerRescuerChanged struct {
	NewRescuer common.Address
	Raw        *types.Log // Blockchain specific contextual infos
}

const BaseTokenMessengerRescuerChangedEventName = "RescuerChanged"

// ContractEventName returns the user-defined event name.
func (BaseTokenMessengerRescuerChanged) ContractEventName() string {
	return BaseTokenMessengerRescuerChangedEventName
}

// UnpackRescuerChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RescuerChanged(address indexed newRescuer)
func (baseTokenMessenger *BaseTokenMessenger) UnpackRescuerChangedEvent(log *types.Log) (*BaseTokenMessengerRescuerChanged, error) {
	event := "RescuerChanged"
	if len(log.Topics) == 0 || log.Topics[0] != baseTokenMessenger.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BaseTokenMessengerRescuerChanged)
	if len(log.Data) > 0 {
		if err := baseTokenMessenger.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range baseTokenMessenger.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// BaseTokenMessengerUnDenylisted represents a UnDenylisted event raised by the BaseTokenMessenger contract.
type BaseTokenMessengerUnDenylisted struct {
	Account common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const BaseTokenMessengerUnDenylistedEventName = "UnDenylisted"

// ContractEventName returns the user-defined event name.
func (BaseTokenMessengerUnDenylisted) ContractEventName() string {
	return BaseTokenMessengerUnDenylistedEventName
}

// UnpackUnDenylistedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event UnDenylisted(address indexed account)
func (baseTokenMessenger *BaseTokenMessenger) UnpackUnDenylistedEvent(log *types.Log) (*BaseTokenMessengerUnDenylisted, error) {
	event := "UnDenylisted"
	if len(log.Topics) == 0 || log.Topics[0] != baseTokenMessenger.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BaseTokenMessengerUnDenylisted)
	if len(log.Data) > 0 {
		if err := baseTokenMessenger.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range baseTokenMessenger.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// BurnMessageMetaData contains all meta data concerning the BurnMessage contract.
var BurnMessageMetaData = bind.MetaData{
	ABI: "[]",
	ID:  "4d2b92996d0b2cb87698e3b8b8e79defb9",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220d0623bec34baadf35164b357dd631a8ce860dfc9fe42a735fb2fbc012a1ac8df64736f6c63430007060033",
}

// BurnMessage is an auto generated Go binding around an Ethereum contract.
type BurnMessage struct {
	abi abi.ABI
}

// NewBurnMessage creates a new instance of BurnMessage.
func NewBurnMessage() *BurnMessage {
	parsed, err := BurnMessageMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &BurnMessage{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *BurnMessage) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// BurnMessageV2MetaData contains all meta data concerning the BurnMessageV2 contract.
var BurnMessageV2MetaData = bind.MetaData{
	ABI: "[]",
	ID:  "621755e1c0c621d756c1d1eeb5ff564093",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220f4a3bd2533b42a8780092a38ca233a08e174864110dec6c839608615f521a27764736f6c63430007060033",
}

// BurnMessageV2 is an auto generated Go binding around an Ethereum contract.
type BurnMessageV2 struct {
	abi abi.ABI
}

// NewBurnMessageV2 creates a new instance of BurnMessageV2.
func NewBurnMessageV2() *BurnMessageV2 {
	parsed, err := BurnMessageV2MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &BurnMessageV2{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *BurnMessageV2) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// ContextMetaData contains all meta data concerning the Context contract.
var ContextMetaData = bind.MetaData{
	ABI: "[]",
	ID:  "5dc51e25e355e604896e8034ea2ca694f7",
}

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	abi abi.ABI
}

// NewContext creates a new instance of Context.
func NewContext() *Context {
	parsed, err := ContextMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Context{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Context) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// DenylistableMetaData contains all meta data concerning the Denylistable contract.
var DenylistableMetaData = bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Denylisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldDenylister\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newDenylister\",\"type\":\"address\"}],\"name\":\"DenylisterChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"UnDenylisted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"denylist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denylister\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isDenylisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"unDenylist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDenylister\",\"type\":\"address\"}],\"name\":\"updateDenylister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "67fbeb2998d4dc1595d236f4b9542aadca",
}

// Denylistable is an auto generated Go binding around an Ethereum contract.
type Denylistable struct {
	abi abi.ABI
}

// NewDenylistable creates a new instance of Denylistable.
func NewDenylistable() *Denylistable {
	parsed, err := DenylistableMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Denylistable{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Denylistable) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackAcceptOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x79ba5097.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function acceptOwnership() returns()
func (denylistable *Denylistable) PackAcceptOwnership() []byte {
	enc, err := denylistable.abi.Pack("acceptOwnership")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAcceptOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x79ba5097.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function acceptOwnership() returns()
func (denylistable *Denylistable) TryPackAcceptOwnership() ([]byte, error) {
	return denylistable.abi.Pack("acceptOwnership")
}

// PackDenylist is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3371bfff.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function denylist(address account) returns()
func (denylistable *Denylistable) PackDenylist(account common.Address) []byte {
	enc, err := denylistable.abi.Pack("denylist", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDenylist is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3371bfff.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function denylist(address account) returns()
func (denylistable *Denylistable) TryPackDenylist(account common.Address) ([]byte, error) {
	return denylistable.abi.Pack("denylist", account)
}

// PackDenylister is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbcc76c60.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function denylister() view returns(address)
func (denylistable *Denylistable) PackDenylister() []byte {
	enc, err := denylistable.abi.Pack("denylister")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDenylister is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbcc76c60.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function denylister() view returns(address)
func (denylistable *Denylistable) TryPackDenylister() ([]byte, error) {
	return denylistable.abi.Pack("denylister")
}

// UnpackDenylister is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbcc76c60.
//
// Solidity: function denylister() view returns(address)
func (denylistable *Denylistable) UnpackDenylister(data []byte) (common.Address, error) {
	out, err := denylistable.abi.Unpack("denylister", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackIsDenylisted is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe877a526.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isDenylisted(address account) view returns(bool)
func (denylistable *Denylistable) PackIsDenylisted(account common.Address) []byte {
	enc, err := denylistable.abi.Pack("isDenylisted", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsDenylisted is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe877a526.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isDenylisted(address account) view returns(bool)
func (denylistable *Denylistable) TryPackIsDenylisted(account common.Address) ([]byte, error) {
	return denylistable.abi.Pack("isDenylisted", account)
}

// UnpackIsDenylisted is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe877a526.
//
// Solidity: function isDenylisted(address account) view returns(bool)
func (denylistable *Denylistable) UnpackIsDenylisted(data []byte) (bool, error) {
	out, err := denylistable.abi.Unpack("isDenylisted", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function owner() view returns(address)
func (denylistable *Denylistable) PackOwner() []byte {
	enc, err := denylistable.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function owner() view returns(address)
func (denylistable *Denylistable) TryPackOwner() ([]byte, error) {
	return denylistable.abi.Pack("owner")
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (denylistable *Denylistable) UnpackOwner(data []byte) (common.Address, error) {
	out, err := denylistable.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackPendingOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe30c3978.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function pendingOwner() view returns(address)
func (denylistable *Denylistable) PackPendingOwner() []byte {
	enc, err := denylistable.abi.Pack("pendingOwner")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPendingOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe30c3978.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function pendingOwner() view returns(address)
func (denylistable *Denylistable) TryPackPendingOwner() ([]byte, error) {
	return denylistable.abi.Pack("pendingOwner")
}

// UnpackPendingOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (denylistable *Denylistable) UnpackPendingOwner(data []byte) (common.Address, error) {
	out, err := denylistable.abi.Unpack("pendingOwner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (denylistable *Denylistable) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := denylistable.abi.Pack("transferOwnership", newOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (denylistable *Denylistable) TryPackTransferOwnership(newOwner common.Address) ([]byte, error) {
	return denylistable.abi.Pack("transferOwnership", newOwner)
}

// PackUnDenylist is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9cab0c1c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function unDenylist(address account) returns()
func (denylistable *Denylistable) PackUnDenylist(account common.Address) []byte {
	enc, err := denylistable.abi.Pack("unDenylist", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUnDenylist is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9cab0c1c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function unDenylist(address account) returns()
func (denylistable *Denylistable) TryPackUnDenylist(account common.Address) ([]byte, error) {
	return denylistable.abi.Pack("unDenylist", account)
}

// PackUpdateDenylister is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa946de04.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updateDenylister(address newDenylister) returns()
func (denylistable *Denylistable) PackUpdateDenylister(newDenylister common.Address) []byte {
	enc, err := denylistable.abi.Pack("updateDenylister", newDenylister)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUpdateDenylister is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa946de04.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function updateDenylister(address newDenylister) returns()
func (denylistable *Denylistable) TryPackUpdateDenylister(newDenylister common.Address) ([]byte, error) {
	return denylistable.abi.Pack("updateDenylister", newDenylister)
}

// DenylistableDenylisted represents a Denylisted event raised by the Denylistable contract.
type DenylistableDenylisted struct {
	Account common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const DenylistableDenylistedEventName = "Denylisted"

// ContractEventName returns the user-defined event name.
func (DenylistableDenylisted) ContractEventName() string {
	return DenylistableDenylistedEventName
}

// UnpackDenylistedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Denylisted(address indexed account)
func (denylistable *Denylistable) UnpackDenylistedEvent(log *types.Log) (*DenylistableDenylisted, error) {
	event := "Denylisted"
	if len(log.Topics) == 0 || log.Topics[0] != denylistable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(DenylistableDenylisted)
	if len(log.Data) > 0 {
		if err := denylistable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range denylistable.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// DenylistableDenylisterChanged represents a DenylisterChanged event raised by the Denylistable contract.
type DenylistableDenylisterChanged struct {
	OldDenylister common.Address
	NewDenylister common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const DenylistableDenylisterChangedEventName = "DenylisterChanged"

// ContractEventName returns the user-defined event name.
func (DenylistableDenylisterChanged) ContractEventName() string {
	return DenylistableDenylisterChangedEventName
}

// UnpackDenylisterChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DenylisterChanged(address indexed oldDenylister, address indexed newDenylister)
func (denylistable *Denylistable) UnpackDenylisterChangedEvent(log *types.Log) (*DenylistableDenylisterChanged, error) {
	event := "DenylisterChanged"
	if len(log.Topics) == 0 || log.Topics[0] != denylistable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(DenylistableDenylisterChanged)
	if len(log.Data) > 0 {
		if err := denylistable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range denylistable.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// DenylistableOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the Denylistable contract.
type DenylistableOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const DenylistableOwnershipTransferStartedEventName = "OwnershipTransferStarted"

// ContractEventName returns the user-defined event name.
func (DenylistableOwnershipTransferStarted) ContractEventName() string {
	return DenylistableOwnershipTransferStartedEventName
}

// UnpackOwnershipTransferStartedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (denylistable *Denylistable) UnpackOwnershipTransferStartedEvent(log *types.Log) (*DenylistableOwnershipTransferStarted, error) {
	event := "OwnershipTransferStarted"
	if len(log.Topics) == 0 || log.Topics[0] != denylistable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(DenylistableOwnershipTransferStarted)
	if len(log.Data) > 0 {
		if err := denylistable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range denylistable.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// DenylistableOwnershipTransferred represents a OwnershipTransferred event raised by the Denylistable contract.
type DenylistableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const DenylistableOwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (DenylistableOwnershipTransferred) ContractEventName() string {
	return DenylistableOwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (denylistable *Denylistable) UnpackOwnershipTransferredEvent(log *types.Log) (*DenylistableOwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if len(log.Topics) == 0 || log.Topics[0] != denylistable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(DenylistableOwnershipTransferred)
	if len(log.Data) > 0 {
		if err := denylistable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range denylistable.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// DenylistableUnDenylisted represents a UnDenylisted event raised by the Denylistable contract.
type DenylistableUnDenylisted struct {
	Account common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const DenylistableUnDenylistedEventName = "UnDenylisted"

// ContractEventName returns the user-defined event name.
func (DenylistableUnDenylisted) ContractEventName() string {
	return DenylistableUnDenylistedEventName
}

// UnpackUnDenylistedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event UnDenylisted(address indexed account)
func (denylistable *Denylistable) UnpackUnDenylistedEvent(log *types.Log) (*DenylistableUnDenylisted, error) {
	event := "UnDenylisted"
	if len(log.Topics) == 0 || log.Topics[0] != denylistable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(DenylistableUnDenylisted)
	if len(log.Data) > 0 {
		if err := denylistable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range denylistable.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// IERC20MetaData contains all meta data concerning the IERC20 contract.
var IERC20MetaData = bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "9947c78930534b1030055e875b67d16e98",
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	abi abi.ABI
}

// NewIERC20 creates a new instance of IERC20.
func NewIERC20() *IERC20 {
	parsed, err := IERC20MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IERC20{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IERC20) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackAllowance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdd62ed3e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (iERC20 *IERC20) PackAllowance(owner common.Address, spender common.Address) []byte {
	enc, err := iERC20.abi.Pack("allowance", owner, spender)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAllowance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdd62ed3e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (iERC20 *IERC20) TryPackAllowance(owner common.Address, spender common.Address) ([]byte, error) {
	return iERC20.abi.Pack("allowance", owner, spender)
}

// UnpackAllowance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (iERC20 *IERC20) UnpackAllowance(data []byte) (*big.Int, error) {
	out, err := iERC20.abi.Unpack("allowance", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackApprove is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x095ea7b3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (iERC20 *IERC20) PackApprove(spender common.Address, amount *big.Int) []byte {
	enc, err := iERC20.abi.Pack("approve", spender, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackApprove is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x095ea7b3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (iERC20 *IERC20) TryPackApprove(spender common.Address, amount *big.Int) ([]byte, error) {
	return iERC20.abi.Pack("approve", spender, amount)
}

// UnpackApprove is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (iERC20 *IERC20) UnpackApprove(data []byte) (bool, error) {
	out, err := iERC20.abi.Unpack("approve", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (iERC20 *IERC20) PackBalanceOf(account common.Address) []byte {
	enc, err := iERC20.abi.Pack("balanceOf", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (iERC20 *IERC20) TryPackBalanceOf(account common.Address) ([]byte, error) {
	return iERC20.abi.Pack("balanceOf", account)
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (iERC20 *IERC20) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := iERC20.abi.Unpack("balanceOf", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackTotalSupply is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x18160ddd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function totalSupply() view returns(uint256)
func (iERC20 *IERC20) PackTotalSupply() []byte {
	enc, err := iERC20.abi.Pack("totalSupply")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTotalSupply is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x18160ddd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function totalSupply() view returns(uint256)
func (iERC20 *IERC20) TryPackTotalSupply() ([]byte, error) {
	return iERC20.abi.Pack("totalSupply")
}

// UnpackTotalSupply is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (iERC20 *IERC20) UnpackTotalSupply(data []byte) (*big.Int, error) {
	out, err := iERC20.abi.Unpack("totalSupply", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa9059cbb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (iERC20 *IERC20) PackTransfer(recipient common.Address, amount *big.Int) []byte {
	enc, err := iERC20.abi.Pack("transfer", recipient, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa9059cbb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (iERC20 *IERC20) TryPackTransfer(recipient common.Address, amount *big.Int) ([]byte, error) {
	return iERC20.abi.Pack("transfer", recipient, amount)
}

// UnpackTransfer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (iERC20 *IERC20) UnpackTransfer(data []byte) (bool, error) {
	out, err := iERC20.abi.Unpack("transfer", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x23b872dd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (iERC20 *IERC20) PackTransferFrom(sender common.Address, recipient common.Address, amount *big.Int) []byte {
	enc, err := iERC20.abi.Pack("transferFrom", sender, recipient, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x23b872dd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (iERC20 *IERC20) TryPackTransferFrom(sender common.Address, recipient common.Address, amount *big.Int) ([]byte, error) {
	return iERC20.abi.Pack("transferFrom", sender, recipient, amount)
}

// UnpackTransferFrom is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (iERC20 *IERC20) UnpackTransferFrom(data []byte) (bool, error) {
	out, err := iERC20.abi.Unpack("transferFrom", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const IERC20ApprovalEventName = "Approval"

// ContractEventName returns the user-defined event name.
func (IERC20Approval) ContractEventName() string {
	return IERC20ApprovalEventName
}

// UnpackApprovalEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (iERC20 *IERC20) UnpackApprovalEvent(log *types.Log) (*IERC20Approval, error) {
	event := "Approval"
	if len(log.Topics) == 0 || log.Topics[0] != iERC20.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IERC20Approval)
	if len(log.Data) > 0 {
		if err := iERC20.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iERC20.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   *types.Log // Blockchain specific contextual infos
}

const IERC20TransferEventName = "Transfer"

// ContractEventName returns the user-defined event name.
func (IERC20Transfer) ContractEventName() string {
	return IERC20TransferEventName
}

// UnpackTransferEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (iERC20 *IERC20) UnpackTransferEvent(log *types.Log) (*IERC20Transfer, error) {
	event := "Transfer"
	if len(log.Topics) == 0 || log.Topics[0] != iERC20.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IERC20Transfer)
	if len(log.Data) > 0 {
		if err := iERC20.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iERC20.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// IMessageHandlerV2MetaData contains all meta data concerning the IMessageHandlerV2 contract.
var IMessageHandlerV2MetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"sourceDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"finalityThresholdExecuted\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"}],\"name\":\"handleReceiveFinalizedMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"sourceDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"finalityThresholdExecuted\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"}],\"name\":\"handleReceiveUnfinalizedMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "56adac06142e507eae8fa000370543e7b8",
}

// IMessageHandlerV2 is an auto generated Go binding around an Ethereum contract.
type IMessageHandlerV2 struct {
	abi abi.ABI
}

// NewIMessageHandlerV2 creates a new instance of IMessageHandlerV2.
func NewIMessageHandlerV2() *IMessageHandlerV2 {
	parsed, err := IMessageHandlerV2MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IMessageHandlerV2{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IMessageHandlerV2) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackHandleReceiveFinalizedMessage is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x11cffb67.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function handleReceiveFinalizedMessage(uint32 sourceDomain, bytes32 sender, uint32 finalityThresholdExecuted, bytes messageBody) returns(bool)
func (iMessageHandlerV2 *IMessageHandlerV2) PackHandleReceiveFinalizedMessage(sourceDomain uint32, sender [32]byte, finalityThresholdExecuted uint32, messageBody []byte) []byte {
	enc, err := iMessageHandlerV2.abi.Pack("handleReceiveFinalizedMessage", sourceDomain, sender, finalityThresholdExecuted, messageBody)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackHandleReceiveFinalizedMessage is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x11cffb67.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function handleReceiveFinalizedMessage(uint32 sourceDomain, bytes32 sender, uint32 finalityThresholdExecuted, bytes messageBody) returns(bool)
func (iMessageHandlerV2 *IMessageHandlerV2) TryPackHandleReceiveFinalizedMessage(sourceDomain uint32, sender [32]byte, finalityThresholdExecuted uint32, messageBody []byte) ([]byte, error) {
	return iMessageHandlerV2.abi.Pack("handleReceiveFinalizedMessage", sourceDomain, sender, finalityThresholdExecuted, messageBody)
}

// UnpackHandleReceiveFinalizedMessage is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x11cffb67.
//
// Solidity: function handleReceiveFinalizedMessage(uint32 sourceDomain, bytes32 sender, uint32 finalityThresholdExecuted, bytes messageBody) returns(bool)
func (iMessageHandlerV2 *IMessageHandlerV2) UnpackHandleReceiveFinalizedMessage(data []byte) (bool, error) {
	out, err := iMessageHandlerV2.abi.Unpack("handleReceiveFinalizedMessage", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackHandleReceiveUnfinalizedMessage is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7c92f219.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function handleReceiveUnfinalizedMessage(uint32 sourceDomain, bytes32 sender, uint32 finalityThresholdExecuted, bytes messageBody) returns(bool)
func (iMessageHandlerV2 *IMessageHandlerV2) PackHandleReceiveUnfinalizedMessage(sourceDomain uint32, sender [32]byte, finalityThresholdExecuted uint32, messageBody []byte) []byte {
	enc, err := iMessageHandlerV2.abi.Pack("handleReceiveUnfinalizedMessage", sourceDomain, sender, finalityThresholdExecuted, messageBody)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackHandleReceiveUnfinalizedMessage is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7c92f219.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function handleReceiveUnfinalizedMessage(uint32 sourceDomain, bytes32 sender, uint32 finalityThresholdExecuted, bytes messageBody) returns(bool)
func (iMessageHandlerV2 *IMessageHandlerV2) TryPackHandleReceiveUnfinalizedMessage(sourceDomain uint32, sender [32]byte, finalityThresholdExecuted uint32, messageBody []byte) ([]byte, error) {
	return iMessageHandlerV2.abi.Pack("handleReceiveUnfinalizedMessage", sourceDomain, sender, finalityThresholdExecuted, messageBody)
}

// UnpackHandleReceiveUnfinalizedMessage is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7c92f219.
//
// Solidity: function handleReceiveUnfinalizedMessage(uint32 sourceDomain, bytes32 sender, uint32 finalityThresholdExecuted, bytes messageBody) returns(bool)
func (iMessageHandlerV2 *IMessageHandlerV2) UnpackHandleReceiveUnfinalizedMessage(data []byte) (bool, error) {
	out, err := iMessageHandlerV2.abi.Unpack("handleReceiveUnfinalizedMessage", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// IMintBurnTokenMetaData contains all meta data concerning the IMintBurnToken contract.
var IMintBurnTokenMetaData = bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "fd492ea973a3eb81ca3f3494f57f59e82b",
}

// IMintBurnToken is an auto generated Go binding around an Ethereum contract.
type IMintBurnToken struct {
	abi abi.ABI
}

// NewIMintBurnToken creates a new instance of IMintBurnToken.
func NewIMintBurnToken() *IMintBurnToken {
	parsed, err := IMintBurnTokenMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IMintBurnToken{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IMintBurnToken) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackAllowance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdd62ed3e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (iMintBurnToken *IMintBurnToken) PackAllowance(owner common.Address, spender common.Address) []byte {
	enc, err := iMintBurnToken.abi.Pack("allowance", owner, spender)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAllowance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdd62ed3e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (iMintBurnToken *IMintBurnToken) TryPackAllowance(owner common.Address, spender common.Address) ([]byte, error) {
	return iMintBurnToken.abi.Pack("allowance", owner, spender)
}

// UnpackAllowance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (iMintBurnToken *IMintBurnToken) UnpackAllowance(data []byte) (*big.Int, error) {
	out, err := iMintBurnToken.abi.Unpack("allowance", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackApprove is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x095ea7b3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (iMintBurnToken *IMintBurnToken) PackApprove(spender common.Address, amount *big.Int) []byte {
	enc, err := iMintBurnToken.abi.Pack("approve", spender, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackApprove is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x095ea7b3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (iMintBurnToken *IMintBurnToken) TryPackApprove(spender common.Address, amount *big.Int) ([]byte, error) {
	return iMintBurnToken.abi.Pack("approve", spender, amount)
}

// UnpackApprove is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (iMintBurnToken *IMintBurnToken) UnpackApprove(data []byte) (bool, error) {
	out, err := iMintBurnToken.abi.Unpack("approve", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (iMintBurnToken *IMintBurnToken) PackBalanceOf(account common.Address) []byte {
	enc, err := iMintBurnToken.abi.Pack("balanceOf", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (iMintBurnToken *IMintBurnToken) TryPackBalanceOf(account common.Address) ([]byte, error) {
	return iMintBurnToken.abi.Pack("balanceOf", account)
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (iMintBurnToken *IMintBurnToken) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := iMintBurnToken.abi.Unpack("balanceOf", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackBurn is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42966c68.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function burn(uint256 amount) returns()
func (iMintBurnToken *IMintBurnToken) PackBurn(amount *big.Int) []byte {
	enc, err := iMintBurnToken.abi.Pack("burn", amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBurn is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42966c68.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function burn(uint256 amount) returns()
func (iMintBurnToken *IMintBurnToken) TryPackBurn(amount *big.Int) ([]byte, error) {
	return iMintBurnToken.abi.Pack("burn", amount)
}

// PackMint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x40c10f19.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function mint(address to, uint256 amount) returns(bool)
func (iMintBurnToken *IMintBurnToken) PackMint(to common.Address, amount *big.Int) []byte {
	enc, err := iMintBurnToken.abi.Pack("mint", to, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x40c10f19.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function mint(address to, uint256 amount) returns(bool)
func (iMintBurnToken *IMintBurnToken) TryPackMint(to common.Address, amount *big.Int) ([]byte, error) {
	return iMintBurnToken.abi.Pack("mint", to, amount)
}

// UnpackMint is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns(bool)
func (iMintBurnToken *IMintBurnToken) UnpackMint(data []byte) (bool, error) {
	out, err := iMintBurnToken.abi.Unpack("mint", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackTotalSupply is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x18160ddd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function totalSupply() view returns(uint256)
func (iMintBurnToken *IMintBurnToken) PackTotalSupply() []byte {
	enc, err := iMintBurnToken.abi.Pack("totalSupply")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTotalSupply is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x18160ddd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function totalSupply() view returns(uint256)
func (iMintBurnToken *IMintBurnToken) TryPackTotalSupply() ([]byte, error) {
	return iMintBurnToken.abi.Pack("totalSupply")
}

// UnpackTotalSupply is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (iMintBurnToken *IMintBurnToken) UnpackTotalSupply(data []byte) (*big.Int, error) {
	out, err := iMintBurnToken.abi.Unpack("totalSupply", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa9059cbb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (iMintBurnToken *IMintBurnToken) PackTransfer(recipient common.Address, amount *big.Int) []byte {
	enc, err := iMintBurnToken.abi.Pack("transfer", recipient, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa9059cbb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (iMintBurnToken *IMintBurnToken) TryPackTransfer(recipient common.Address, amount *big.Int) ([]byte, error) {
	return iMintBurnToken.abi.Pack("transfer", recipient, amount)
}

// UnpackTransfer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (iMintBurnToken *IMintBurnToken) UnpackTransfer(data []byte) (bool, error) {
	out, err := iMintBurnToken.abi.Unpack("transfer", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x23b872dd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (iMintBurnToken *IMintBurnToken) PackTransferFrom(sender common.Address, recipient common.Address, amount *big.Int) []byte {
	enc, err := iMintBurnToken.abi.Pack("transferFrom", sender, recipient, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x23b872dd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (iMintBurnToken *IMintBurnToken) TryPackTransferFrom(sender common.Address, recipient common.Address, amount *big.Int) ([]byte, error) {
	return iMintBurnToken.abi.Pack("transferFrom", sender, recipient, amount)
}

// UnpackTransferFrom is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (iMintBurnToken *IMintBurnToken) UnpackTransferFrom(data []byte) (bool, error) {
	out, err := iMintBurnToken.abi.Unpack("transferFrom", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// IMintBurnTokenApproval represents a Approval event raised by the IMintBurnToken contract.
type IMintBurnTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const IMintBurnTokenApprovalEventName = "Approval"

// ContractEventName returns the user-defined event name.
func (IMintBurnTokenApproval) ContractEventName() string {
	return IMintBurnTokenApprovalEventName
}

// UnpackApprovalEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (iMintBurnToken *IMintBurnToken) UnpackApprovalEvent(log *types.Log) (*IMintBurnTokenApproval, error) {
	event := "Approval"
	if len(log.Topics) == 0 || log.Topics[0] != iMintBurnToken.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IMintBurnTokenApproval)
	if len(log.Data) > 0 {
		if err := iMintBurnToken.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iMintBurnToken.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// IMintBurnTokenTransfer represents a Transfer event raised by the IMintBurnToken contract.
type IMintBurnTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   *types.Log // Blockchain specific contextual infos
}

const IMintBurnTokenTransferEventName = "Transfer"

// ContractEventName returns the user-defined event name.
func (IMintBurnTokenTransfer) ContractEventName() string {
	return IMintBurnTokenTransferEventName
}

// UnpackTransferEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (iMintBurnToken *IMintBurnToken) UnpackTransferEvent(log *types.Log) (*IMintBurnTokenTransfer, error) {
	event := "Transfer"
	if len(log.Topics) == 0 || log.Topics[0] != iMintBurnToken.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IMintBurnTokenTransfer)
	if len(log.Data) > 0 {
		if err := iMintBurnToken.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iMintBurnToken.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// IRelayerV2MetaData contains all meta data concerning the IRelayerV2 contract.
var IRelayerV2MetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCaller\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"minFinalityThreshold\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "80bb51bbf05a5f10b8061006b35952bd1b",
}

// IRelayerV2 is an auto generated Go binding around an Ethereum contract.
type IRelayerV2 struct {
	abi abi.ABI
}

// NewIRelayerV2 creates a new instance of IRelayerV2.
func NewIRelayerV2() *IRelayerV2 {
	parsed, err := IRelayerV2MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IRelayerV2{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IRelayerV2) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackSendMessage is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x14b157ab.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function sendMessage(uint32 destinationDomain, bytes32 recipient, bytes32 destinationCaller, uint32 minFinalityThreshold, bytes messageBody) returns()
func (iRelayerV2 *IRelayerV2) PackSendMessage(destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, minFinalityThreshold uint32, messageBody []byte) []byte {
	enc, err := iRelayerV2.abi.Pack("sendMessage", destinationDomain, recipient, destinationCaller, minFinalityThreshold, messageBody)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSendMessage is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x14b157ab.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function sendMessage(uint32 destinationDomain, bytes32 recipient, bytes32 destinationCaller, uint32 minFinalityThreshold, bytes messageBody) returns()
func (iRelayerV2 *IRelayerV2) TryPackSendMessage(destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, minFinalityThreshold uint32, messageBody []byte) ([]byte, error) {
	return iRelayerV2.abi.Pack("sendMessage", destinationDomain, recipient, destinationCaller, minFinalityThreshold, messageBody)
}

// ITokenMinterMetaData contains all meta data concerning the ITokenMinter contract.
var ITokenMinterMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"burnToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"remoteToken\",\"type\":\"bytes32\"}],\"name\":\"getLocalToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"sourceDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"burnToken\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"mintToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTokenController\",\"type\":\"address\"}],\"name\":\"setTokenController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "fe0a26e0c102c76b6f7c1c915601fc2eeb",
}

// ITokenMinter is an auto generated Go binding around an Ethereum contract.
type ITokenMinter struct {
	abi abi.ABI
}

// NewITokenMinter creates a new instance of ITokenMinter.
func NewITokenMinter() *ITokenMinter {
	parsed, err := ITokenMinterMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ITokenMinter{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ITokenMinter) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackBurn is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9dc29fac.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function burn(address burnToken, uint256 amount) returns()
func (iTokenMinter *ITokenMinter) PackBurn(burnToken common.Address, amount *big.Int) []byte {
	enc, err := iTokenMinter.abi.Pack("burn", burnToken, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBurn is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9dc29fac.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function burn(address burnToken, uint256 amount) returns()
func (iTokenMinter *ITokenMinter) TryPackBurn(burnToken common.Address, amount *big.Int) ([]byte, error) {
	return iTokenMinter.abi.Pack("burn", burnToken, amount)
}

// PackGetLocalToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x78a0565e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getLocalToken(uint32 remoteDomain, bytes32 remoteToken) view returns(address)
func (iTokenMinter *ITokenMinter) PackGetLocalToken(remoteDomain uint32, remoteToken [32]byte) []byte {
	enc, err := iTokenMinter.abi.Pack("getLocalToken", remoteDomain, remoteToken)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetLocalToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x78a0565e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getLocalToken(uint32 remoteDomain, bytes32 remoteToken) view returns(address)
func (iTokenMinter *ITokenMinter) TryPackGetLocalToken(remoteDomain uint32, remoteToken [32]byte) ([]byte, error) {
	return iTokenMinter.abi.Pack("getLocalToken", remoteDomain, remoteToken)
}

// UnpackGetLocalToken is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x78a0565e.
//
// Solidity: function getLocalToken(uint32 remoteDomain, bytes32 remoteToken) view returns(address)
func (iTokenMinter *ITokenMinter) UnpackGetLocalToken(data []byte) (common.Address, error) {
	out, err := iTokenMinter.abi.Unpack("getLocalToken", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackMint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd54de06f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function mint(uint32 sourceDomain, bytes32 burnToken, address to, uint256 amount) returns(address mintToken)
func (iTokenMinter *ITokenMinter) PackMint(sourceDomain uint32, burnToken [32]byte, to common.Address, amount *big.Int) []byte {
	enc, err := iTokenMinter.abi.Pack("mint", sourceDomain, burnToken, to, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd54de06f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function mint(uint32 sourceDomain, bytes32 burnToken, address to, uint256 amount) returns(address mintToken)
func (iTokenMinter *ITokenMinter) TryPackMint(sourceDomain uint32, burnToken [32]byte, to common.Address, amount *big.Int) ([]byte, error) {
	return iTokenMinter.abi.Pack("mint", sourceDomain, burnToken, to, amount)
}

// UnpackMint is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd54de06f.
//
// Solidity: function mint(uint32 sourceDomain, bytes32 burnToken, address to, uint256 amount) returns(address mintToken)
func (iTokenMinter *ITokenMinter) UnpackMint(data []byte) (common.Address, error) {
	out, err := iTokenMinter.abi.Unpack("mint", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackSetTokenController is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe102baab.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setTokenController(address newTokenController) returns()
func (iTokenMinter *ITokenMinter) PackSetTokenController(newTokenController common.Address) []byte {
	enc, err := iTokenMinter.abi.Pack("setTokenController", newTokenController)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetTokenController is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe102baab.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setTokenController(address newTokenController) returns()
func (iTokenMinter *ITokenMinter) TryPackSetTokenController(newTokenController common.Address) ([]byte, error) {
	return iTokenMinter.abi.Pack("setTokenController", newTokenController)
}

// ITokenMinterV2MetaData contains all meta data concerning the ITokenMinterV2 contract.
var ITokenMinterV2MetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"burnToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"remoteToken\",\"type\":\"bytes32\"}],\"name\":\"getLocalToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"sourceDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"burnToken\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"recipientOne\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientTwo\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOne\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTwo\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"sourceDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"burnToken\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"mintToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTokenController\",\"type\":\"address\"}],\"name\":\"setTokenController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "719940ed2968a32aec303e576bf80a5bde",
}

// ITokenMinterV2 is an auto generated Go binding around an Ethereum contract.
type ITokenMinterV2 struct {
	abi abi.ABI
}

// NewITokenMinterV2 creates a new instance of ITokenMinterV2.
func NewITokenMinterV2() *ITokenMinterV2 {
	parsed, err := ITokenMinterV2MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ITokenMinterV2{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ITokenMinterV2) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackBurn is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9dc29fac.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function burn(address burnToken, uint256 amount) returns()
func (iTokenMinterV2 *ITokenMinterV2) PackBurn(burnToken common.Address, amount *big.Int) []byte {
	enc, err := iTokenMinterV2.abi.Pack("burn", burnToken, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBurn is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9dc29fac.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function burn(address burnToken, uint256 amount) returns()
func (iTokenMinterV2 *ITokenMinterV2) TryPackBurn(burnToken common.Address, amount *big.Int) ([]byte, error) {
	return iTokenMinterV2.abi.Pack("burn", burnToken, amount)
}

// PackGetLocalToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x78a0565e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getLocalToken(uint32 remoteDomain, bytes32 remoteToken) view returns(address)
func (iTokenMinterV2 *ITokenMinterV2) PackGetLocalToken(remoteDomain uint32, remoteToken [32]byte) []byte {
	enc, err := iTokenMinterV2.abi.Pack("getLocalToken", remoteDomain, remoteToken)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetLocalToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x78a0565e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getLocalToken(uint32 remoteDomain, bytes32 remoteToken) view returns(address)
func (iTokenMinterV2 *ITokenMinterV2) TryPackGetLocalToken(remoteDomain uint32, remoteToken [32]byte) ([]byte, error) {
	return iTokenMinterV2.abi.Pack("getLocalToken", remoteDomain, remoteToken)
}

// UnpackGetLocalToken is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x78a0565e.
//
// Solidity: function getLocalToken(uint32 remoteDomain, bytes32 remoteToken) view returns(address)
func (iTokenMinterV2 *ITokenMinterV2) UnpackGetLocalToken(data []byte) (common.Address, error) {
	out, err := iTokenMinterV2.abi.Unpack("getLocalToken", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackMint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8dfcfa90.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function mint(uint32 sourceDomain, bytes32 burnToken, address recipientOne, address recipientTwo, uint256 amountOne, uint256 amountTwo) returns(address)
func (iTokenMinterV2 *ITokenMinterV2) PackMint(sourceDomain uint32, burnToken [32]byte, recipientOne common.Address, recipientTwo common.Address, amountOne *big.Int, amountTwo *big.Int) []byte {
	enc, err := iTokenMinterV2.abi.Pack("mint", sourceDomain, burnToken, recipientOne, recipientTwo, amountOne, amountTwo)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8dfcfa90.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function mint(uint32 sourceDomain, bytes32 burnToken, address recipientOne, address recipientTwo, uint256 amountOne, uint256 amountTwo) returns(address)
func (iTokenMinterV2 *ITokenMinterV2) TryPackMint(sourceDomain uint32, burnToken [32]byte, recipientOne common.Address, recipientTwo common.Address, amountOne *big.Int, amountTwo *big.Int) ([]byte, error) {
	return iTokenMinterV2.abi.Pack("mint", sourceDomain, burnToken, recipientOne, recipientTwo, amountOne, amountTwo)
}

// UnpackMint is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8dfcfa90.
//
// Solidity: function mint(uint32 sourceDomain, bytes32 burnToken, address recipientOne, address recipientTwo, uint256 amountOne, uint256 amountTwo) returns(address)
func (iTokenMinterV2 *ITokenMinterV2) UnpackMint(data []byte) (common.Address, error) {
	out, err := iTokenMinterV2.abi.Unpack("mint", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackMint0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd54de06f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function mint(uint32 sourceDomain, bytes32 burnToken, address to, uint256 amount) returns(address mintToken)
func (iTokenMinterV2 *ITokenMinterV2) PackMint0(sourceDomain uint32, burnToken [32]byte, to common.Address, amount *big.Int) []byte {
	enc, err := iTokenMinterV2.abi.Pack("mint0", sourceDomain, burnToken, to, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMint0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd54de06f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function mint(uint32 sourceDomain, bytes32 burnToken, address to, uint256 amount) returns(address mintToken)
func (iTokenMinterV2 *ITokenMinterV2) TryPackMint0(sourceDomain uint32, burnToken [32]byte, to common.Address, amount *big.Int) ([]byte, error) {
	return iTokenMinterV2.abi.Pack("mint0", sourceDomain, burnToken, to, amount)
}

// UnpackMint0 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd54de06f.
//
// Solidity: function mint(uint32 sourceDomain, bytes32 burnToken, address to, uint256 amount) returns(address mintToken)
func (iTokenMinterV2 *ITokenMinterV2) UnpackMint0(data []byte) (common.Address, error) {
	out, err := iTokenMinterV2.abi.Unpack("mint0", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackSetTokenController is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe102baab.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setTokenController(address newTokenController) returns()
func (iTokenMinterV2 *ITokenMinterV2) PackSetTokenController(newTokenController common.Address) []byte {
	enc, err := iTokenMinterV2.abi.Pack("setTokenController", newTokenController)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetTokenController is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe102baab.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setTokenController(address newTokenController) returns()
func (iTokenMinterV2 *ITokenMinterV2) TryPackSetTokenController(newTokenController common.Address) ([]byte, error) {
	return iTokenMinterV2.abi.Pack("setTokenController", newTokenController)
}

// InitializableMetaData contains all meta data concerning the Initializable contract.
var InitializableMetaData = bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"}]",
	ID:  "8a814103c3b4eabfa4e9bfe471c975eee3",
}

// Initializable is an auto generated Go binding around an Ethereum contract.
type Initializable struct {
	abi abi.ABI
}

// NewInitializable creates a new instance of Initializable.
func NewInitializable() *Initializable {
	parsed, err := InitializableMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Initializable{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Initializable) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// InitializableInitialized represents a Initialized event raised by the Initializable contract.
type InitializableInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const InitializableInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (InitializableInitialized) ContractEventName() string {
	return InitializableInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (initializable *Initializable) UnpackInitializedEvent(log *types.Log) (*InitializableInitialized, error) {
	event := "Initialized"
	if len(log.Topics) == 0 || log.Topics[0] != initializable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(InitializableInitialized)
	if len(log.Data) > 0 {
		if err := initializable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range initializable.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// OwnableMetaData contains all meta data concerning the Ownable contract.
var OwnableMetaData = bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "fb41fcf4b096041dd27ea053966ea80b71",
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	abi abi.ABI
}

// NewOwnable creates a new instance of Ownable.
func NewOwnable() *Ownable {
	parsed, err := OwnableMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Ownable{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Ownable) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function owner() view returns(address)
func (ownable *Ownable) PackOwner() []byte {
	enc, err := ownable.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function owner() view returns(address)
func (ownable *Ownable) TryPackOwner() ([]byte, error) {
	return ownable.abi.Pack("owner")
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (ownable *Ownable) UnpackOwner(data []byte) (common.Address, error) {
	out, err := ownable.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (ownable *Ownable) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := ownable.abi.Pack("transferOwnership", newOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (ownable *Ownable) TryPackTransferOwnership(newOwner common.Address) ([]byte, error) {
	return ownable.abi.Pack("transferOwnership", newOwner)
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const OwnableOwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (OwnableOwnershipTransferred) ContractEventName() string {
	return OwnableOwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (ownable *Ownable) UnpackOwnershipTransferredEvent(log *types.Log) (*OwnableOwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if len(log.Topics) == 0 || log.Topics[0] != ownable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(OwnableOwnershipTransferred)
	if len(log.Data) > 0 {
		if err := ownable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range ownable.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// Ownable2StepMetaData contains all meta data concerning the Ownable2Step contract.
var Ownable2StepMetaData = bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "502091fe46efe2c19a7fc465e240c7b440",
}

// Ownable2Step is an auto generated Go binding around an Ethereum contract.
type Ownable2Step struct {
	abi abi.ABI
}

// NewOwnable2Step creates a new instance of Ownable2Step.
func NewOwnable2Step() *Ownable2Step {
	parsed, err := Ownable2StepMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Ownable2Step{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Ownable2Step) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackAcceptOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x79ba5097.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function acceptOwnership() returns()
func (ownable2Step *Ownable2Step) PackAcceptOwnership() []byte {
	enc, err := ownable2Step.abi.Pack("acceptOwnership")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAcceptOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x79ba5097.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function acceptOwnership() returns()
func (ownable2Step *Ownable2Step) TryPackAcceptOwnership() ([]byte, error) {
	return ownable2Step.abi.Pack("acceptOwnership")
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function owner() view returns(address)
func (ownable2Step *Ownable2Step) PackOwner() []byte {
	enc, err := ownable2Step.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function owner() view returns(address)
func (ownable2Step *Ownable2Step) TryPackOwner() ([]byte, error) {
	return ownable2Step.abi.Pack("owner")
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (ownable2Step *Ownable2Step) UnpackOwner(data []byte) (common.Address, error) {
	out, err := ownable2Step.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackPendingOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe30c3978.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function pendingOwner() view returns(address)
func (ownable2Step *Ownable2Step) PackPendingOwner() []byte {
	enc, err := ownable2Step.abi.Pack("pendingOwner")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPendingOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe30c3978.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function pendingOwner() view returns(address)
func (ownable2Step *Ownable2Step) TryPackPendingOwner() ([]byte, error) {
	return ownable2Step.abi.Pack("pendingOwner")
}

// UnpackPendingOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (ownable2Step *Ownable2Step) UnpackPendingOwner(data []byte) (common.Address, error) {
	out, err := ownable2Step.abi.Unpack("pendingOwner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (ownable2Step *Ownable2Step) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := ownable2Step.abi.Pack("transferOwnership", newOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (ownable2Step *Ownable2Step) TryPackTransferOwnership(newOwner common.Address) ([]byte, error) {
	return ownable2Step.abi.Pack("transferOwnership", newOwner)
}

// Ownable2StepOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the Ownable2Step contract.
type Ownable2StepOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const Ownable2StepOwnershipTransferStartedEventName = "OwnershipTransferStarted"

// ContractEventName returns the user-defined event name.
func (Ownable2StepOwnershipTransferStarted) ContractEventName() string {
	return Ownable2StepOwnershipTransferStartedEventName
}

// UnpackOwnershipTransferStartedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (ownable2Step *Ownable2Step) UnpackOwnershipTransferStartedEvent(log *types.Log) (*Ownable2StepOwnershipTransferStarted, error) {
	event := "OwnershipTransferStarted"
	if len(log.Topics) == 0 || log.Topics[0] != ownable2Step.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Ownable2StepOwnershipTransferStarted)
	if len(log.Data) > 0 {
		if err := ownable2Step.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range ownable2Step.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// Ownable2StepOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable2Step contract.
type Ownable2StepOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const Ownable2StepOwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (Ownable2StepOwnershipTransferred) ContractEventName() string {
	return Ownable2StepOwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (ownable2Step *Ownable2Step) UnpackOwnershipTransferredEvent(log *types.Log) (*Ownable2StepOwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if len(log.Topics) == 0 || log.Topics[0] != ownable2Step.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Ownable2StepOwnershipTransferred)
	if len(log.Data) > 0 {
		if err := ownable2Step.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range ownable2Step.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// RescuableMetaData contains all meta data concerning the Rescuable contract.
var RescuableMetaData = bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newRescuer\",\"type\":\"address\"}],\"name\":\"RescuerChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rescuer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRescuer\",\"type\":\"address\"}],\"name\":\"updateRescuer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "ce5ac842fd8e3ca6b5c5e1a61f6cd5e2c4",
	Bin: "0x608060405234801561001057600080fd5b5061002d61002261003260201b60201c565b61003a60201b60201c565b610139565b600033905090565b600160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556100728161007560201b6104d51760201c565b50565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b610cbd806101486000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80638da5cb5b1161005b5780638da5cb5b14610104578063b2118a8d14610138578063e30c3978146101a6578063f2fde38b146101da5761007d565b80632ab600451461008257806338a63183146100c657806379ba5097146100fa575b600080fd5b6100c46004803603602081101561009857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061021e565b005b6100ce610232565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61010261025c565b005b61010c6102ff565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101a46004803603606081101561014e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610328565b005b6101ae6103fe565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61021c600480360360208110156101f057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610428565b005b610226610599565b61022f8161064a565b50565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000610266610757565b90508073ffffffffffffffffffffffffffffffffffffffff166102876103fe565b73ffffffffffffffffffffffffffffffffffffffff16146102f3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526029815260200180610bc16029913960400191505060405180910390fd5b6102fc8161075f565b50565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103ce576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526024815260200180610c3a6024913960400191505060405180910390fd5b6103f982828573ffffffffffffffffffffffffffffffffffffffff166107909092919063ffffffff16565b505050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610430610599565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff166104906102ff565b73ffffffffffffffffffffffffffffffffffffffff167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a350565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6105a1610757565b73ffffffffffffffffffffffffffffffffffffffff166105bf6102ff565b73ffffffffffffffffffffffffffffffffffffffff1614610648576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156106d0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a815260200180610bea602a913960400191505060405180910390fd5b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff167fe475e580d85111348e40d8ca33cfdd74c30fe1655c2d8537a13abc10065ffa5a60405160405180910390a250565b600033905090565b600160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905561078d816104d5565b50565b61082d8363a9059cbb60e01b8484604051602401808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610832565b505050565b6000610894826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166109219092919063ffffffff16565b905060008151111561091c578080602001905160208110156108b557600080fd5b810190808051906020019092919050505061091b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a815260200180610c5e602a913960400191505060405180910390fd5b5b505050565b60606109308484600085610939565b90509392505050565b606082471015610994576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180610c146026913960400191505060405180910390fd5b61099d85610ae1565b610a0f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000081525060200191505060405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040518082805190602001908083835b60208310610a5e5780518252602082019150602081019050602083039250610a3b565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114610ac0576040519150601f19603f3d011682016040523d82523d6000602084013e610ac5565b606091505b5091509150610ad5828286610af4565b92505050949350505050565b600080823b905060008111915050919050565b60608315610b0457829050610bb9565b600083511115610b175782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610b7e578082015181840152602081019050610b63565b50505050905090810190601f168015610bab5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b939250505056fe4f776e61626c6532537465703a2063616c6c6572206973206e6f7420746865206e6577206f776e6572526573637561626c653a206e6577207265736375657220697320746865207a65726f2061646472657373416464726573733a20696e73756666696369656e742062616c616e636520666f722063616c6c526573637561626c653a2063616c6c6572206973206e6f742074686520726573637565725361666545524332303a204552433230206f7065726174696f6e20646964206e6f742073756363656564a264697066735822122091e03e77a38a106000348bc038000cd66d39f988448e6c61c299359d83ca6b2064736f6c63430007060033",
}

// Rescuable is an auto generated Go binding around an Ethereum contract.
type Rescuable struct {
	abi abi.ABI
}

// NewRescuable creates a new instance of Rescuable.
func NewRescuable() *Rescuable {
	parsed, err := RescuableMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Rescuable{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Rescuable) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackAcceptOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x79ba5097.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function acceptOwnership() returns()
func (rescuable *Rescuable) PackAcceptOwnership() []byte {
	enc, err := rescuable.abi.Pack("acceptOwnership")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAcceptOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x79ba5097.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function acceptOwnership() returns()
func (rescuable *Rescuable) TryPackAcceptOwnership() ([]byte, error) {
	return rescuable.abi.Pack("acceptOwnership")
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function owner() view returns(address)
func (rescuable *Rescuable) PackOwner() []byte {
	enc, err := rescuable.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function owner() view returns(address)
func (rescuable *Rescuable) TryPackOwner() ([]byte, error) {
	return rescuable.abi.Pack("owner")
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (rescuable *Rescuable) UnpackOwner(data []byte) (common.Address, error) {
	out, err := rescuable.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackPendingOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe30c3978.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function pendingOwner() view returns(address)
func (rescuable *Rescuable) PackPendingOwner() []byte {
	enc, err := rescuable.abi.Pack("pendingOwner")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPendingOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe30c3978.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function pendingOwner() view returns(address)
func (rescuable *Rescuable) TryPackPendingOwner() ([]byte, error) {
	return rescuable.abi.Pack("pendingOwner")
}

// UnpackPendingOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (rescuable *Rescuable) UnpackPendingOwner(data []byte) (common.Address, error) {
	out, err := rescuable.abi.Unpack("pendingOwner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackRescueERC20 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb2118a8d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function rescueERC20(address tokenContract, address to, uint256 amount) returns()
func (rescuable *Rescuable) PackRescueERC20(tokenContract common.Address, to common.Address, amount *big.Int) []byte {
	enc, err := rescuable.abi.Pack("rescueERC20", tokenContract, to, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRescueERC20 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb2118a8d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function rescueERC20(address tokenContract, address to, uint256 amount) returns()
func (rescuable *Rescuable) TryPackRescueERC20(tokenContract common.Address, to common.Address, amount *big.Int) ([]byte, error) {
	return rescuable.abi.Pack("rescueERC20", tokenContract, to, amount)
}

// PackRescuer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x38a63183.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function rescuer() view returns(address)
func (rescuable *Rescuable) PackRescuer() []byte {
	enc, err := rescuable.abi.Pack("rescuer")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRescuer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x38a63183.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function rescuer() view returns(address)
func (rescuable *Rescuable) TryPackRescuer() ([]byte, error) {
	return rescuable.abi.Pack("rescuer")
}

// UnpackRescuer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x38a63183.
//
// Solidity: function rescuer() view returns(address)
func (rescuable *Rescuable) UnpackRescuer(data []byte) (common.Address, error) {
	out, err := rescuable.abi.Unpack("rescuer", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (rescuable *Rescuable) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := rescuable.abi.Pack("transferOwnership", newOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (rescuable *Rescuable) TryPackTransferOwnership(newOwner common.Address) ([]byte, error) {
	return rescuable.abi.Pack("transferOwnership", newOwner)
}

// PackUpdateRescuer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2ab60045.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updateRescuer(address newRescuer) returns()
func (rescuable *Rescuable) PackUpdateRescuer(newRescuer common.Address) []byte {
	enc, err := rescuable.abi.Pack("updateRescuer", newRescuer)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUpdateRescuer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2ab60045.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function updateRescuer(address newRescuer) returns()
func (rescuable *Rescuable) TryPackUpdateRescuer(newRescuer common.Address) ([]byte, error) {
	return rescuable.abi.Pack("updateRescuer", newRescuer)
}

// RescuableOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the Rescuable contract.
type RescuableOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const RescuableOwnershipTransferStartedEventName = "OwnershipTransferStarted"

// ContractEventName returns the user-defined event name.
func (RescuableOwnershipTransferStarted) ContractEventName() string {
	return RescuableOwnershipTransferStartedEventName
}

// UnpackOwnershipTransferStartedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (rescuable *Rescuable) UnpackOwnershipTransferStartedEvent(log *types.Log) (*RescuableOwnershipTransferStarted, error) {
	event := "OwnershipTransferStarted"
	if len(log.Topics) == 0 || log.Topics[0] != rescuable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RescuableOwnershipTransferStarted)
	if len(log.Data) > 0 {
		if err := rescuable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rescuable.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// RescuableOwnershipTransferred represents a OwnershipTransferred event raised by the Rescuable contract.
type RescuableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const RescuableOwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (RescuableOwnershipTransferred) ContractEventName() string {
	return RescuableOwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (rescuable *Rescuable) UnpackOwnershipTransferredEvent(log *types.Log) (*RescuableOwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if len(log.Topics) == 0 || log.Topics[0] != rescuable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RescuableOwnershipTransferred)
	if len(log.Data) > 0 {
		if err := rescuable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rescuable.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// RescuableRescuerChanged represents a RescuerChanged event raised by the Rescuable contract.
type RescuableRescuerChanged struct {
	NewRescuer common.Address
	Raw        *types.Log // Blockchain specific contextual infos
}

const RescuableRescuerChangedEventName = "RescuerChanged"

// ContractEventName returns the user-defined event name.
func (RescuableRescuerChanged) ContractEventName() string {
	return RescuableRescuerChangedEventName
}

// UnpackRescuerChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RescuerChanged(address indexed newRescuer)
func (rescuable *Rescuable) UnpackRescuerChangedEvent(log *types.Log) (*RescuableRescuerChanged, error) {
	event := "RescuerChanged"
	if len(log.Topics) == 0 || log.Topics[0] != rescuable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RescuableRescuerChanged)
	if len(log.Data) > 0 {
		if err := rescuable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rescuable.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// SafeERC20MetaData contains all meta data concerning the SafeERC20 contract.
var SafeERC20MetaData = bind.MetaData{
	ABI: "[]",
	ID:  "b1b345c69d5781d026cff4956083592b48",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122010377c05667b32ae0d943b9f0fb7ed710906ea71bf224f4b419df0213030089464736f6c63430007060033",
}

// SafeERC20 is an auto generated Go binding around an Ethereum contract.
type SafeERC20 struct {
	abi abi.ABI
}

// NewSafeERC20 creates a new instance of SafeERC20.
func NewSafeERC20() *SafeERC20 {
	parsed, err := SafeERC20MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &SafeERC20{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *SafeERC20) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// SafeMathMetaData contains all meta data concerning the SafeMath contract.
var SafeMathMetaData = bind.MetaData{
	ABI: "[]",
	ID:  "2a563a746d6470214bcc8039e89cdbaae7",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212201690a6717ad67f308153285c5009bd923951998d48e50e33dafc720c69f47d2864736f6c63430007060033",
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	abi abi.ABI
}

// NewSafeMath creates a new instance of SafeMath.
func NewSafeMath() *SafeMath {
	parsed, err := SafeMathMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &SafeMath{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *SafeMath) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// TokenMessengerV2MetaData contains all meta data concerning the TokenMessengerV2 contract.
var TokenMessengerV2MetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messageTransmitter\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_messageBodyVersion\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Denylisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldDenylister\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newDenylister\",\"type\":\"address\"}],\"name\":\"DenylisterChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"burnToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"mintRecipient\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"destinationTokenMessenger\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"destinationCaller\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"minFinalityThreshold\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"hookData\",\"type\":\"bytes\"}],\"name\":\"DepositForBurn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"FeeRecipientSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"localMinter\",\"type\":\"address\"}],\"name\":\"LocalMinterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"localMinter\",\"type\":\"address\"}],\"name\":\"LocalMinterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minFeeController\",\"type\":\"address\"}],\"name\":\"MinFeeControllerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minFee\",\"type\":\"uint256\"}],\"name\":\"MinFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mintRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mintToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeCollected\",\"type\":\"uint256\"}],\"name\":\"MintAndWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"tokenMessenger\",\"type\":\"bytes32\"}],\"name\":\"RemoteTokenMessengerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"tokenMessenger\",\"type\":\"bytes32\"}],\"name\":\"RemoteTokenMessengerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newRescuer\",\"type\":\"address\"}],\"name\":\"RescuerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"UnDenylisted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MIN_FEE_MULTIPLIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newLocalMinter\",\"type\":\"address\"}],\"name\":\"addLocalMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"tokenMessenger\",\"type\":\"bytes32\"}],\"name\":\"addRemoteTokenMessenger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"denylist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denylister\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"mintRecipient\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"burnToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCaller\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"minFinalityThreshold\",\"type\":\"uint32\"}],\"name\":\"depositForBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"mintRecipient\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"burnToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCaller\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"minFinalityThreshold\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"hookData\",\"type\":\"bytes\"}],\"name\":\"depositForBurnWithHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getMinFeeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"}],\"name\":\"handleReceiveFinalizedMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"finalityThresholdExecuted\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"}],\"name\":\"handleReceiveUnfinalizedMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rescuer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"denylister\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenMinter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minFeeController\",\"type\":\"address\"}],\"internalType\":\"structTokenMessengerV2.TokenMessengerV2Roles\",\"name\":\"roles\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"minFee_\",\"type\":\"uint256\"},{\"internalType\":\"uint32[]\",\"name\":\"remoteDomains_\",\"type\":\"uint32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"remoteTokenMessengers_\",\"type\":\"bytes32[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedVersion\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isDenylisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localMessageTransmitter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localMinter\",\"outputs\":[{\"internalType\":\"contractITokenMinterV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageBodyVersion\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minFeeController\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"remoteTokenMessengers\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeLocalMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"removeRemoteTokenMessenger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rescuer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeRecipient\",\"type\":\"address\"}],\"name\":\"setFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minFee\",\"type\":\"uint256\"}],\"name\":\"setMinFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minFeeController\",\"type\":\"address\"}],\"name\":\"setMinFeeController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"unDenylist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDenylister\",\"type\":\"address\"}],\"name\":\"updateDenylister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRescuer\",\"type\":\"address\"}],\"name\":\"updateRescuer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "ce165425b06c6a9259fe2b5f62cdc5058f",
	Bin: "0x60c06040523480156200001157600080fd5b50604051620053b4380380620053b48339818101604052810190620000379190620003f7565b8181620000596200004d6200016560201b60201c565b6200016d60201b60201c565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415620000fd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f4d6573736167655472616e736d6974746572206e6f742073657400000000000081525060200191505060405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1660601b815250508063ffffffff1660a08163ffffffff1660e01b8152505050506200015d620001ab60201b60201c565b5050620004b0565b600033905090565b600160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055620001a881620002dd60201b620018021760201c565b50565b6000620001bd620003a160201b60201c565b90508060000160089054906101000a900460ff161562000229576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806200538f6025913960400191505060405180910390fd5b67ffffffffffffffff80168160000160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1614620002da5767ffffffffffffffff8160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d267ffffffffffffffff604051808267ffffffffffffffff16815260200191505060405180910390a15b50565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b60007ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b600081519050620003da816200047c565b92915050565b600081519050620003f18162000496565b92915050565b600080604083850312156200040b57600080fd5b60006200041b85828601620003c9565b92505060206200042e85828601620003e0565b9150509250929050565b600062000445826200044c565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600063ffffffff82169050919050565b620004878162000438565b81146200049357600080fd5b50565b620004a1816200046c565b8114620004ad57600080fd5b50565b60805160601c60a05160e01c614ea1620004ee60003980611415528061260652806128b9525080610af7528061224c52806126715250614ea16000f3fe608060405234801561001057600080fd5b50600436106102065760003560e01c80638da5cb5b1161011a578063b2118a8d116100ad578063e30c39781161007c578063e30c397814610579578063e74b981b14610597578063e877a526146105b3578063f2fde38b146105e3578063f79fd08e146105ff57610206565b8063b2118a8d14610505578063bcc76c6014610521578063cb75c11c1461053f578063da87e4481461055d57610206565b80639cab0c1c116100e95780639cab0c1c146104935780639cdbb181146104af578063a5b8d04e146104cd578063a946de04146104e957610206565b80638da5cb5b146104315780638e0250ee1461044f57806391f178881461046b578063966dfbd51461047557610206565b8063369c4f1b1161019d578063779b432d1161016c578063779b432d1461038f57806379ba5097146103ab5780637c92f219146103b55780638197beb9146103e557806382a5e6651461040157610206565b8063369c4f1b1461030557806338a63183146103235780634690484014610341578063516990e31461035f57610206565b80632ab60045116101d95780632ab60045146102935780632c121921146102af57806331ac9920146102cd5780633371bfff146102e957610206565b806302db402e1461020b57806308828eb71461022757806311cffb671461024557806324ec759014610275575b600080fd5b61022560048036038101906102209190613f9c565b61061b565b005b61022f610981565b60405161023c9190614a22565b60405180910390f35b61025f600480360381019061025a9190614233565b610990565b60405161026c9190614729565b60405180910390f35b61027d610adb565b60405161028a919061491a565b60405180910390f35b6102ad60048036038101906102a89190613f24565b610ae1565b005b6102b7610af5565b6040516102c4919061470e565b60405180910390f35b6102e760048036038101906102e29190614039565b610b19565b005b61030360048036038101906102fe9190613f24565b610bcb565b005b61030d610cfc565b60405161031a919061470e565b60405180910390f35b61032b610d22565b604051610338919061470e565b60405180910390f35b610349610d4c565b604051610356919061470e565b60405180910390f35b61037960048036038101906103749190614039565b610d72565b604051610386919061491a565b60405180910390f35b6103a960048036038101906103a49190614100565b610ddb565b005b6103b3610e83565b005b6103cf60048036038101906103ca9190614233565b610f26565b6040516103dc9190614729565b60405180910390f35b6103ff60048036038101906103fa9190613f24565b6110c2565b005b61041b600480360381019061041691906141ce565b6110d6565b6040516104289190614744565b60405180910390f35b6104396110ee565b604051610446919061470e565b60405180910390f35b61046960048036038101906104649190614062565b611117565b005b610473611194565b005b61047d6112db565b60405161048a919061491a565b60405180910390f35b6104ad60048036038101906104a89190613f24565b6112e2565b005b6104b7611413565b6040516104c491906149ad565b60405180910390f35b6104e760048036038101906104e29190613f24565b611437565b005b61050360048036038101906104fe9190613f24565b61144b565b005b61051f600480360381019061051a9190613f4d565b61145f565b005b610529611535565b604051610536919061470e565b60405180910390f35b61054761155f565b604051610554919061475f565b60405180910390f35b610577600480360381019061057291906141f7565b611585565b005b61058161159b565b60405161058e919061470e565b60405180910390f35b6105b160048036038101906105ac9190613f24565b6115c5565b005b6105cd60048036038101906105c89190613f24565b6115d9565b6040516105da9190614729565b60405180910390f35b6105fd60048036038101906105f89190613f24565b611625565b005b610619600480360381019061061491906141ce565b6116d2565b005b60006106256118c6565b905060008160000160089054906101000a900460ff1615905060008260000160009054906101000a900467ffffffffffffffff1690506000808267ffffffffffffffff161480156106735750825b9050600060018367ffffffffffffffff161480156106975750610695306118ee565b155b905081806106a25750805b6106f7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180614c106025913960400191505060405180910390fd5b60018560000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555083156107475760018560000160086101000a81548160ff0219169083151502179055505b600073ffffffffffffffffffffffffffffffffffffffff168b60000160208101906107729190613f24565b73ffffffffffffffffffffffffffffffffffffffff1614156107c9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107c09061481a565b60405180910390fd5b868690508989905014610811576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610808906147ba565b60405180910390fd5b61082c8b60000160208101906108279190613f24565b611901565b6108478b60200160208101906108429190613f24565b611932565b6108628b606001602081019061085d9190613f24565b611a3f565b61087d8b60400160208101906108789190613f24565b611b8b565b6108988b60800160208101906108939190613f24565b611cbf565b6108b38b60a00160208101906108ae9190613f24565b611eb7565b6108bc8a611feb565b600089899050905060005b818110156109165761090b8b8b838181106108de57fe5b90506020020160208101906108f391906141ce565b8a8a848181106108ff57fe5b905060200201356120a4565b8060010190506108c7565b505083156109745760008560000160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d260016040518082815260200191505060405180910390a15b5050505050505050505050565b600061098b612221565b905090565b600061099a612248565b610a0c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f496e76616c6964206d657373616765207472616e736d6974746572000000000081525060200191505060405180910390fd5b8585610a18828261229e565b610a6d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526021815260200180614cb86021913960400191505060405180910390fd5b610ace610ac8600087878080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050506122d990919063ffffffff16565b89612304565b9250505095945050505050565b601d5481565b610ae961233a565b610af281611932565b50565b7f000000000000000000000000000000000000000000000000000000000000000081565b601c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610bbf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526024815260200180614dc36024913960400191505060405180910390fd5b610bc881611feb565b50565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610c71576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180614e116026913960400191505060405180910390fd5b6001600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508073ffffffffffffffffffffffffffffffffffffffff167ffa4507bc1f9c730e6e95897024f1fe7d576cf2deb53579d55c14f1ac3439e11460405160405180910390a250565b601c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b601b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600080601d541415610d875760009050610dd6565b60018211610dca576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dc1906148fa565b60405180910390fd5b610dd3826123eb565b90505b919050565b610de43361242a565b3273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610e2157610e203261242a565b5b60008282905011610e67576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e5e9061485a565b60405180910390fd5b610e788989898989898989896124c5565b505050505050505050565b6000610e8d612789565b90508073ffffffffffffffffffffffffffffffffffffffff16610eae61159b565b73ffffffffffffffffffffffffffffffffffffffff1614610f1a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526029815260200180614c356029913960400191505060405180910390fd5b610f2381611901565b50565b6000610f30612248565b610fa2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f496e76616c6964206d657373616765207472616e736d6974746572000000000081525060200191505060405180910390fd5b8585610fae828261229e565b611003576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526021815260200180614cb86021913960400191505060405180910390fd5b6101f463ffffffff168663ffffffff161015611054576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161104b9061487a565b60405180910390fd5b6110b56110af600087878080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050506122d990919063ffffffff16565b89612304565b9250505095945050505050565b6110ca61233a565b6110d381611cbf565b50565b601a6020528060005260406000206000915090505481565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6111203361242a565b3273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461115d5761115c3261242a565b5b366000803660009060009261117493929190614a6a565b915091506111898989898989898989896124c5565b505050505050505050565b61119c61233a565b6000601960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415611266576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f4e6f206c6f63616c206d696e746572206973207365742e00000000000000000081525060200191505060405180910390fd5b601960006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690557f2db49fbf671271826a27b02ebc496209c85fffffb4bccc67430d2a0f22b4d1ac81604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b6298968081565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611388576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180614e116026913960400191505060405180910390fd5b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508073ffffffffffffffffffffffffffffffffffffffff167fc904e1b03de0c20d7fcf9dbd056daf1bd3815e93f251199de815fd0f0b96e16660405160405180910390a250565b7f000000000000000000000000000000000000000000000000000000000000000081565b61143f61233a565b61144881611eb7565b50565b61145361233a565b61145c81611a3f565b50565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611505576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526024815260200180614d206024913960400191505060405180910390fd5b61153082828573ffffffffffffffffffffffffffffffffffffffff166127919092919063ffffffff16565b505050565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b601960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b61158d61233a565b61159782826120a4565b5050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6115cd61233a565b6115d681611b8b565b50565b60006001600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054149050919050565b61162d61233a565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff1661168d6110ee565b73ffffffffffffffffffffffffffffffffffffffff167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a350565b6116da61233a565b6000801b601a60008363ffffffff1663ffffffff168152602001908152602001600020541415611772576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f4e6f20546f6b656e4d657373656e67657220736574000000000000000000000081525060200191505060405180910390fd5b6000601a60008363ffffffff1663ffffffff168152602001908152602001600020549050601a60008363ffffffff1663ffffffff168152602001908152602001600020600090557f3dcea012093dbca2bb8ed7fd2b2ff90305ab70bddda8bbb94d4152735a98f0b18282604051808363ffffffff1681526020018281526020019250505060405180910390a15050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b60007ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b600080823b905060008111915050919050565b600160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905561192f81611802565b50565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156119b8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a815260200180614c5e602a913960400191505060405180910390fd5b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff167fe475e580d85111348e40d8ca33cfdd74c30fe1655c2d8537a13abc10065ffa5a60405160405180910390a250565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415611ac5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526030815260200180614c886030913960400191505060405180910390fd5b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167fe144e84038182cefebda68c192c222085b2c12a85d135d3c938498c0165c01d360405160405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415611c2e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f5a65726f2061646472657373206e6f7420616c6c6f776564000000000000000081525060200191505060405180910390fd5b80601b60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fbf9a9534339a9d6b81696e05dcfb614b7dc518a31d48be3cfb757988381fb32381604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415611d62576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f5a65726f2061646472657373206e6f7420616c6c6f776564000000000000000081525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16601960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611e26576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f4c6f63616c206d696e74657220697320616c7265616479207365742e0000000081525060200191505060405180910390fd5b80601960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f109bb3e70cbf1931e295b49e75c67013b85ff80d64e6f1d321f37157b90c383081604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415611f5a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f5a65726f2061646472657373206e6f7420616c6c6f776564000000000000000081525060200191505060405180910390fd5b80601c60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f4a381820aeb373f402117f6280a310f11cf06f1ef064b9550f5b6f94f1aaa2a881604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b629896808110612063576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f4d696e2066656520746f6f20686967680000000000000000000000000000000081525060200191505060405180910390fd5b80601d819055507fd4c19c993aeeb50b76da8b158f84806c9d235eff5b86f3a36f12a2e1dd4e6eac816040518082815260200191505060405180910390a150565b6000801b81141561211d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f62797465733332283029206e6f7420616c6c6f7765640000000000000000000081525060200191505060405180910390fd5b6000801b601a60008463ffffffff1663ffffffff16815260200190815260200160002054146121b4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f546f6b656e4d657373656e67657220616c72656164792073657400000000000081525060200191505060405180910390fd5b80601a60008463ffffffff1663ffffffff168152602001908152602001600020819055507f4bba2b08298cf59661b4895e384cc2ac3962ce2d71f1b7c11bca52e1169f95998282604051808363ffffffff1681526020018281526020019250505060405180910390a15050565b600061222b6118c6565b60000160009054906101000a900467ffffffffffffffff16905090565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b60008060001b82141580156122d1575081601a60008563ffffffff1663ffffffff16815260200190815260200160002054145b905092915050565b6000808351905060006020850190506122fa8464ffffffffff168284612833565b9250505092915050565b6000806000806000612315876128a2565b935093509350935061232c86848684860385612a85565b600194505050505092915050565b612342612789565b73ffffffffffffffffffffffffffffffffffffffff166123606110ee565b73ffffffffffffffffffffffffffffffffffffffff16146123e9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b565b60008062989680612407601d5485612ce790919063ffffffff16565b8161240e57fe5b0490506000811461241f5780612422565b60015b915050919050565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054146124c2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526024815260200180614d9f6024913960400191505060405180910390fd5b50565b60008911612508576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016124ff906147fa565b60405180910390fd5b6000801b87141561254e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612545906148ba565b60405180910390fd5b888410612590576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016125879061477a565b60405180910390fd5b6000601d5411156125e7576125a4896123eb565b8410156125e6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016125dd9061489a565b60405180910390fd5b5b60006125f289612d6d565b90506125ff87338c612e14565b600061266d7f00000000000000000000000000000000000000000000000000000000000000006126448a73ffffffffffffffffffffffffffffffffffffffff16612ff1565b8b8e6126653373ffffffffffffffffffffffffffffffffffffffff16612ff1565b8b8a8a613014565b90507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166314b157ab8b848a89866040518663ffffffff1660e01b81526004016126d09594939291906149c8565b600060405180830381600087803b1580156126ea57600080fd5b505af11580156126fe573d6000803e3d6000fd5b505050508463ffffffff163373ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff167f0c8c1cbdc5190613ebd485511d4e2812cfa45eecb79d845893331fedad5130a58e8d8f888e8e8d8d604051612774989796959493929190614935565b60405180910390a45050505050505050505050565b600033905090565b61282e8363a9059cbb60e01b8484604051602401808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050613097565b505050565b600080612849838561318690919063ffffffff16565b905060405181111561285a57600090505b600081141561288c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000091505061289b565b612897858585613209565b9150505b9392505050565b6000806000806128b78562ffffff1916613234565b7f000000000000000000000000000000000000000000000000000000000000000063ffffffff166128ed8662ffffff191661334e565b63ffffffff1614612933576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161292a9061479a565b60405180910390fd5b60006129448662ffffff1916613366565b9050600081148061295457504381115b612993576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161298a906147da565b60405180910390fd5b6129a28662ffffff191661338f565b92506129b38662ffffff19166133a7565b915060008214806129c357508282105b612a02576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016129f99061483a565b60405180910390fd5b612a118662ffffff19166133d0565b821115612a53576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612a4a906148da565b60405180910390fd5b612a6a612a658762ffffff19166133f9565b613411565b9450612a7b8662ffffff191661341e565b9350509193509193565b6000612a8f613436565b9050600080831115612bac578173ffffffffffffffffffffffffffffffffffffffff16638dfcfa90888888601b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1689896040518763ffffffff1660e01b8152600401808763ffffffff1681526020018681526020018573ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018281526020019650505050505050602060405180830381600087803b158015612b6a57600080fd5b505af1158015612b7e573d6000803e3d6000fd5b505050506040513d6020811015612b9457600080fd5b81019080805190602001909291905050509050612c71565b8173ffffffffffffffffffffffffffffffffffffffff1663d54de06f888888886040518563ffffffff1660e01b8152600401808563ffffffff1681526020018481526020018373ffffffffffffffffffffffffffffffffffffffff168152602001828152602001945050505050602060405180830381600087803b158015612c3357600080fd5b505af1158015612c47573d6000803e3d6000fd5b505050506040513d6020811015612c5d57600080fd5b810190808051906020019092919050505090505b8073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f50c55e915134d457debfa58eb6f4342956f8b0616d51a89a3659360178e1ab638686604051808381526020018281526020019250505060405180910390a350505050505050565b600080831415612cfa5760009050612d67565b6000828402905082848281612d0b57fe5b0414612d62576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526021815260200180614d446021913960400191505060405180910390fd5b809150505b92915050565b600080601a60008463ffffffff1663ffffffff1681526020019081526020016000205490506000801b811415612e0b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f4e6f20546f6b656e4d657373656e67657220666f7220646f6d61696e0000000081525060200191505060405180910390fd5b80915050919050565b6000612e1e613436565b905060008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8584866040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b158015612eb457600080fd5b505af1158015612ec8573d6000803e3d6000fd5b505050506040513d6020811015612ede57600080fd5b8101908080519060200190929190505050612f61576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f5472616e73666572206f7065726174696f6e206661696c65640000000000000081525060200191505060405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff16639dc29fac86856040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b158015612fd257600080fd5b505af1158015612fe6573d6000803e3d6000fd5b505050505050505050565b60008173ffffffffffffffffffffffffffffffffffffffff1660001b9050919050565b60608888888888886000808a8a604051602001808b63ffffffff1660e01b81526004018a8152602001898152602001888152602001878152602001868152602001858152602001848152602001838380828437808301925050509a5050505050505050505050604051602081830303815290604052905098975050505050505050565b60006130f9826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166135249092919063ffffffff16565b90506000815111156131815780806020019051602081101561311a57600080fd5b8101908080519060200190929190505050613180576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a815260200180614de7602a913960400191505060405180910390fd5b5b505050565b6000818301905082811015613203576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f4f766572666c6f7720647572696e67206164646974696f6e2e0000000000000081525060200191505060405180910390fd5b92915050565b60008060609050600060189050858317821b9250848317821b9250838317811b925050509392505050565b6132438162ffffff191661353c565b6132b5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f4d616c666f726d6564206d65737361676500000000000000000000000000000081525060200191505060405180910390fd5b60e460ff166132c98262ffffff191661357f565b6bffffffffffffffffffffffff16101561334b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f496e76616c6964206275726e206d6573736167653a20746f6f2073686f72740081525060200191505060405180910390fd5b50565b600061335f8262ffffff19166135a5565b9050919050565b600061338860c460ff1660208462ffffff19166135ce9092919063ffffffff16565b9050919050565b60006133a08262ffffff19166135f3565b9050919050565b60006133c960a460ff1660208462ffffff19166135ce9092919063ffffffff16565b9050919050565b60006133f2608460ff1660208462ffffff19166135ce9092919063ffffffff16565b9050919050565b600061340a8262ffffff191661361c565b9050919050565b60008160001c9050919050565b600061342f8262ffffff1916613645565b9050919050565b60008073ffffffffffffffffffffffffffffffffffffffff16601960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156134fc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f4c6f63616c206d696e746572206973206e6f742073657400000000000000000081525060200191505060405180910390fd5b601960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6060613533848460008561366e565b90509392505050565b600064ffffffffff61354d83613816565b64ffffffffff161415613563576000905061357a565b600061356e8361382b565b90506040518111159150505b919050565b6000806bffffffffffffffffffffffff90506000601890508184821c1692505050919050565b60006135c7600060ff1660048462ffffff19166135ce9092919063ffffffff16565b9050919050565b60006008826020030260ff166135e5858585613855565b60001c901c90509392505050565b6000613615604460ff1660208462ffffff19166135ce9092919063ffffffff16565b9050919050565b600061363e602460ff1660208462ffffff19166138559092919063ffffffff16565b9050919050565b6000613667600460ff1660208462ffffff19166138559092919063ffffffff16565b9050919050565b6060824710156136c9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180614cd96026913960400191505060405180910390fd5b6136d2856118ee565b613744576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000081525060200191505060405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040518082805190602001908083835b602083106137935780518252602082019150602081019050602083039250613770565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d80600081146137f5576040519150601f19603f3d011682016040523d82523d6000602084013e6137fa565b606091505b509150915061380a828286613a1c565b92505050949350505050565b60008060d860ff16905082811c915050919050565b60006138368261357f565b61383f83613ae8565b016bffffffffffffffffffffffff169050919050565b6000808260ff16141561386d576000801b9050613a15565b6138768461357f565b6bffffffffffffffffffffffff1661389a8360ff168561318690919063ffffffff16565b111561397c576138db6138ac85613ae8565b6bffffffffffffffffffffffff166138c38661357f565b6bffffffffffffffffffffffff16858560ff16613b11565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015613941578082015181840152602081019050613926565b50505050905090810190601f16801561396e5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b60208260ff1611156139d9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603a815260200180614d65603a913960400191505060405180910390fd5b600060088302905060006139ec86613ae8565b6bffffffffffffffffffffffff1690506000613a0783613c4b565b905080868301511693505050505b9392505050565b60608315613a2c57829050613ae1565b600083511115613a3f5782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015613aa6578082015181840152602081019050613a8b565b50505050905090810190601f168015613ad35780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b9392505050565b6000806bffffffffffffffffffffffff90506000607860ff1690508184821c1692505050919050565b60606000613b1e86613c7a565b9150506000613b2c86613c7a565b9150506000613b3a86613c7a565b9150506000613b4886613c7a565b915050838383836040516020018080614e37603591396035018565ffffffffffff1660d01b8152600601807f2077697468206c656e6774682030780000000000000000000000000000000000815250600f018465ffffffffffff1660d01b815260060180614cff602191396021018365ffffffffffff1660d01b8152600601807f2077697468206c656e6774682030780000000000000000000000000000000000815250600f018265ffffffffffff1660d01b8152600601807f2e00000000000000000000000000000000000000000000000000000000000000815250600101945050505050604051602081830303815290604052945050505050949350505050565b60007f8000000000000000000000000000000000000000000000000000000000000000600183031d9050919050565b6000806000601f90505b600f8160ff161115613cce5760006008820260ff1685901c9050613ca781613d26565b61ffff168417935060108260ff1614613cc257601084901b93505b50600181039050613c84565b506000600f90505b60ff8160ff161015613d205760006008820260ff1685901c9050613cf981613d26565b61ffff168317925060008260ff1614613d1457601083901b92505b50600181039050613cd6565b50915091565b6000613d3860048360ff16901c613d5f565b60ff168117905060088161ffff16901b9050613d5382613d5f565b60ff1681179050919050565b600080600f831690506040518060400160405280601081526020017f30313233343536373839616263646566000000000000000000000000000000008152508160ff1681518110613dac57fe5b602001015160f81c60f81b60f81c915050919050565b600081359050613dd181614b9c565b92915050565b60008083601f840112613de957600080fd5b8235905067ffffffffffffffff811115613e0257600080fd5b602083019150836020820283011115613e1a57600080fd5b9250929050565b60008083601f840112613e3357600080fd5b8235905067ffffffffffffffff811115613e4c57600080fd5b602083019150836020820283011115613e6457600080fd5b9250929050565b600081359050613e7a81614bb3565b92915050565b60008083601f840112613e9257600080fd5b8235905067ffffffffffffffff811115613eab57600080fd5b602083019150836001820283011115613ec357600080fd5b9250929050565b600081359050613ed981614bca565b92915050565b600060c08284031215613ef157600080fd5b81905092915050565b600081359050613f0981614be1565b92915050565b600081359050613f1e81614bf8565b92915050565b600060208284031215613f3657600080fd5b6000613f4484828501613dc2565b91505092915050565b600080600060608486031215613f6257600080fd5b6000613f7086828701613eca565b9350506020613f8186828701613dc2565b9250506040613f9286828701613efa565b9150509250925092565b6000806000806000806101208789031215613fb657600080fd5b6000613fc489828a01613edf565b96505060c0613fd589828a01613efa565b95505060e087013567ffffffffffffffff811115613ff257600080fd5b613ffe89828a01613e21565b945094505061010087013567ffffffffffffffff81111561401e57600080fd5b61402a89828a01613dd7565b92509250509295509295509295565b60006020828403121561404b57600080fd5b600061405984828501613efa565b91505092915050565b600080600080600080600060e0888a03121561407d57600080fd5b600061408b8a828b01613efa565b975050602061409c8a828b01613f0f565b96505060406140ad8a828b01613e6b565b95505060606140be8a828b01613dc2565b94505060806140cf8a828b01613e6b565b93505060a06140e08a828b01613efa565b92505060c06140f18a828b01613f0f565b91505092959891949750929550565b60008060008060008060008060006101008a8c03121561411f57600080fd5b600061412d8c828d01613efa565b995050602061413e8c828d01613f0f565b985050604061414f8c828d01613e6b565b97505060606141608c828d01613dc2565b96505060806141718c828d01613e6b565b95505060a06141828c828d01613efa565b94505060c06141938c828d01613f0f565b93505060e08a013567ffffffffffffffff8111156141b057600080fd5b6141bc8c828d01613e80565b92509250509295985092959850929598565b6000602082840312156141e057600080fd5b60006141ee84828501613f0f565b91505092915050565b6000806040838503121561420a57600080fd5b600061421885828601613f0f565b925050602061422985828601613e6b565b9150509250929050565b60008060008060006080868803121561424b57600080fd5b600061425988828901613f0f565b955050602061426a88828901613e6b565b945050604061427b88828901613f0f565b935050606086013567ffffffffffffffff81111561429857600080fd5b6142a488828901613e80565b92509250509295509295909350565b6142bc81614a9d565b82525050565b6142cb81614aaf565b82525050565b6142da81614abb565b82525050565b60006142ec8385614a48565b93506142f9838584614b49565b61430283614b8b565b840190509392505050565b600061431882614a3d565b6143228185614a48565b9350614332818560208601614b58565b61433b81614b8b565b840191505092915050565b61434f81614b25565b82525050565b6000614362602083614a59565b91507f4d617820666565206d757374206265206c657373207468616e20616d6f756e746000830152602082019050919050565b60006143a2601c83614a59565b91507f496e76616c6964206d65737361676520626f64792076657273696f6e000000006000830152602082019050919050565b60006143e2602383614a59565b91507f496e76616c69642072656d6f746520646f6d61696e20636f6e6669677572617460008301527f696f6e00000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000614448602583614a59565b91507f4d657373616765206578706972656420616e64206d7573742062652072652d7360008301527f69676e65640000000000000000000000000000000000000000000000000000006020830152604082019050919050565b60006144ae601683614a59565b91507f416d6f756e74206d757374206265206e6f6e7a65726f000000000000000000006000830152602082019050919050565b60006144ee601983614a59565b91507f4f776e657220697320746865207a65726f2061646472657373000000000000006000830152602082019050919050565b600061452e601c83614a59565b91507f46656520657175616c73206f72206578636565647320616d6f756e74000000006000830152602082019050919050565b600061456e601283614a59565b91507f486f6f6b206461746120697320656d70747900000000000000000000000000006000830152602082019050919050565b60006145ae601e83614a59565b91507f556e737570706f727465642066696e616c697479207468726573686f6c6400006000830152602082019050919050565b60006145ee601483614a59565b91507f496e73756666696369656e74206d6178206665650000000000000000000000006000830152602082019050919050565b600061462e601e83614a59565b91507f4d696e7420726563697069656e74206d757374206265206e6f6e7a65726f00006000830152602082019050919050565b600061466e601383614a59565b91507f4665652065786365656473206d617820666565000000000000000000000000006000830152602082019050919050565b60006146ae600e83614a59565b91507f416d6f756e7420746f6f206c6f770000000000000000000000000000000000006000830152602082019050919050565b6146ea81614af7565b82525050565b6146f981614b01565b82525050565b61470881614b11565b82525050565b600060208201905061472360008301846142b3565b92915050565b600060208201905061473e60008301846142c2565b92915050565b600060208201905061475960008301846142d1565b92915050565b60006020820190506147746000830184614346565b92915050565b6000602082019050818103600083015261479381614355565b9050919050565b600060208201905081810360008301526147b381614395565b9050919050565b600060208201905081810360008301526147d3816143d5565b9050919050565b600060208201905081810360008301526147f38161443b565b9050919050565b60006020820190508181036000830152614813816144a1565b9050919050565b60006020820190508181036000830152614833816144e1565b9050919050565b6000602082019050818103600083015261485381614521565b9050919050565b6000602082019050818103600083015261487381614561565b9050919050565b60006020820190508181036000830152614893816145a1565b9050919050565b600060208201905081810360008301526148b3816145e1565b9050919050565b600060208201905081810360008301526148d381614621565b9050919050565b600060208201905081810360008301526148f381614661565b9050919050565b60006020820190508181036000830152614913816146a1565b9050919050565b600060208201905061492f60008301846146e1565b92915050565b600060e08201905061494a600083018b6146e1565b614957602083018a6142d1565b61496460408301896146f0565b61497160608301886142d1565b61497e60808301876142d1565b61498b60a08301866146e1565b81810360c083015261499e8184866142e0565b90509998505050505050505050565b60006020820190506149c260008301846146f0565b92915050565b600060a0820190506149dd60008301886146f0565b6149ea60208301876142d1565b6149f760408301866142d1565b614a0460608301856146f0565b8181036080830152614a16818461430d565b90509695505050505050565b6000602082019050614a3760008301846146ff565b92915050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b60008085851115614a7a57600080fd5b83861115614a8757600080fd5b6001850283019150848603905094509492505050565b6000614aa882614ad7565b9050919050565b60008115159050919050565b6000819050919050565b6000614ad082614a9d565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600063ffffffff82169050919050565b600067ffffffffffffffff82169050919050565b6000614b3082614b37565b9050919050565b6000614b4282614ad7565b9050919050565b82818337600083830152505050565b60005b83811015614b76578082015181840152602081019050614b5b565b83811115614b85576000848401525b50505050565b6000601f19601f8301169050919050565b614ba581614a9d565b8114614bb057600080fd5b50565b614bbc81614abb565b8114614bc757600080fd5b50565b614bd381614ac5565b8114614bde57600080fd5b50565b614bea81614af7565b8114614bf557600080fd5b50565b614c0181614b01565b8114614c0c57600080fd5b5056fe496e697469616c697a61626c653a20696e76616c696420696e697469616c697a6174696f6e4f776e61626c6532537465703a2063616c6c6572206973206e6f7420746865206e6577206f776e6572526573637561626c653a206e6577207265736375657220697320746865207a65726f206164647265737344656e796c69737461626c653a206e65772064656e796c697374657220697320746865207a65726f206164647265737352656d6f746520546f6b656e4d657373656e67657220756e737570706f72746564416464726573733a20696e73756666696369656e742062616c616e636520666f722063616c6c2e20417474656d7074656420746f20696e646578206174206f6666736574203078526573637561626c653a2063616c6c6572206973206e6f74207468652072657363756572536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f7754797065644d656d566965772f696e646578202d20417474656d7074656420746f20696e646578206d6f7265207468616e20333220627974657344656e796c69737461626c653a206163636f756e74206973206f6e2064656e796c69737443616c6c6572206973206e6f7420746865206d696e2066656520636f6e74726f6c6c65725361666545524332303a204552433230206f7065726174696f6e20646964206e6f74207375636365656444656e796c69737461626c653a2063616c6c6572206973206e6f742064656e796c697374657254797065644d656d566965772f696e646578202d204f76657272616e2074686520766965772e20536c696365206973206174203078a2646970667358221220dc366bd0ba0a49b6a564e130c991cf76103f98fe799d991c488e8b7fa63521d864736f6c63430007060033496e697469616c697a61626c653a20696e76616c696420696e697469616c697a6174696f6e",
}

// TokenMessengerV2 is an auto generated Go binding around an Ethereum contract.
type TokenMessengerV2 struct {
	abi abi.ABI
}

// NewTokenMessengerV2 creates a new instance of TokenMessengerV2.
func NewTokenMessengerV2() *TokenMessengerV2 {
	parsed, err := TokenMessengerV2MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &TokenMessengerV2{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *TokenMessengerV2) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _messageTransmitter, uint32 _messageBodyVersion) returns()
func (tokenMessengerV2 *TokenMessengerV2) PackConstructor(_messageTransmitter common.Address, _messageBodyVersion uint32) []byte {
	enc, err := tokenMessengerV2.abi.Pack("", _messageTransmitter, _messageBodyVersion)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackMINFEEMULTIPLIER is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x966dfbd5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function MIN_FEE_MULTIPLIER() view returns(uint256)
func (tokenMessengerV2 *TokenMessengerV2) PackMINFEEMULTIPLIER() []byte {
	enc, err := tokenMessengerV2.abi.Pack("MIN_FEE_MULTIPLIER")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMINFEEMULTIPLIER is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x966dfbd5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function MIN_FEE_MULTIPLIER() view returns(uint256)
func (tokenMessengerV2 *TokenMessengerV2) TryPackMINFEEMULTIPLIER() ([]byte, error) {
	return tokenMessengerV2.abi.Pack("MIN_FEE_MULTIPLIER")
}

// UnpackMINFEEMULTIPLIER is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x966dfbd5.
//
// Solidity: function MIN_FEE_MULTIPLIER() view returns(uint256)
func (tokenMessengerV2 *TokenMessengerV2) UnpackMINFEEMULTIPLIER(data []byte) (*big.Int, error) {
	out, err := tokenMessengerV2.abi.Unpack("MIN_FEE_MULTIPLIER", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackAcceptOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x79ba5097.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function acceptOwnership() returns()
func (tokenMessengerV2 *TokenMessengerV2) PackAcceptOwnership() []byte {
	enc, err := tokenMessengerV2.abi.Pack("acceptOwnership")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAcceptOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x79ba5097.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function acceptOwnership() returns()
func (tokenMessengerV2 *TokenMessengerV2) TryPackAcceptOwnership() ([]byte, error) {
	return tokenMessengerV2.abi.Pack("acceptOwnership")
}

// PackAddLocalMinter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8197beb9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addLocalMinter(address newLocalMinter) returns()
func (tokenMessengerV2 *TokenMessengerV2) PackAddLocalMinter(newLocalMinter common.Address) []byte {
	enc, err := tokenMessengerV2.abi.Pack("addLocalMinter", newLocalMinter)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddLocalMinter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8197beb9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addLocalMinter(address newLocalMinter) returns()
func (tokenMessengerV2 *TokenMessengerV2) TryPackAddLocalMinter(newLocalMinter common.Address) ([]byte, error) {
	return tokenMessengerV2.abi.Pack("addLocalMinter", newLocalMinter)
}

// PackAddRemoteTokenMessenger is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xda87e448.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addRemoteTokenMessenger(uint32 domain, bytes32 tokenMessenger) returns()
func (tokenMessengerV2 *TokenMessengerV2) PackAddRemoteTokenMessenger(domain uint32, tokenMessenger [32]byte) []byte {
	enc, err := tokenMessengerV2.abi.Pack("addRemoteTokenMessenger", domain, tokenMessenger)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddRemoteTokenMessenger is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xda87e448.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addRemoteTokenMessenger(uint32 domain, bytes32 tokenMessenger) returns()
func (tokenMessengerV2 *TokenMessengerV2) TryPackAddRemoteTokenMessenger(domain uint32, tokenMessenger [32]byte) ([]byte, error) {
	return tokenMessengerV2.abi.Pack("addRemoteTokenMessenger", domain, tokenMessenger)
}

// PackDenylist is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3371bfff.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function denylist(address account) returns()
func (tokenMessengerV2 *TokenMessengerV2) PackDenylist(account common.Address) []byte {
	enc, err := tokenMessengerV2.abi.Pack("denylist", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDenylist is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3371bfff.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function denylist(address account) returns()
func (tokenMessengerV2 *TokenMessengerV2) TryPackDenylist(account common.Address) ([]byte, error) {
	return tokenMessengerV2.abi.Pack("denylist", account)
}

// PackDenylister is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbcc76c60.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function denylister() view returns(address)
func (tokenMessengerV2 *TokenMessengerV2) PackDenylister() []byte {
	enc, err := tokenMessengerV2.abi.Pack("denylister")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDenylister is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbcc76c60.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function denylister() view returns(address)
func (tokenMessengerV2 *TokenMessengerV2) TryPackDenylister() ([]byte, error) {
	return tokenMessengerV2.abi.Pack("denylister")
}

// UnpackDenylister is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbcc76c60.
//
// Solidity: function denylister() view returns(address)
func (tokenMessengerV2 *TokenMessengerV2) UnpackDenylister(data []byte) (common.Address, error) {
	out, err := tokenMessengerV2.abi.Unpack("denylister", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackDepositForBurn is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8e0250ee.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function depositForBurn(uint256 amount, uint32 destinationDomain, bytes32 mintRecipient, address burnToken, bytes32 destinationCaller, uint256 maxFee, uint32 minFinalityThreshold) returns()
func (tokenMessengerV2 *TokenMessengerV2) PackDepositForBurn(amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte, maxFee *big.Int, minFinalityThreshold uint32) []byte {
	enc, err := tokenMessengerV2.abi.Pack("depositForBurn", amount, destinationDomain, mintRecipient, burnToken, destinationCaller, maxFee, minFinalityThreshold)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDepositForBurn is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8e0250ee.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function depositForBurn(uint256 amount, uint32 destinationDomain, bytes32 mintRecipient, address burnToken, bytes32 destinationCaller, uint256 maxFee, uint32 minFinalityThreshold) returns()
func (tokenMessengerV2 *TokenMessengerV2) TryPackDepositForBurn(amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte, maxFee *big.Int, minFinalityThreshold uint32) ([]byte, error) {
	return tokenMessengerV2.abi.Pack("depositForBurn", amount, destinationDomain, mintRecipient, burnToken, destinationCaller, maxFee, minFinalityThreshold)
}

// PackDepositForBurnWithHook is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x779b432d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function depositForBurnWithHook(uint256 amount, uint32 destinationDomain, bytes32 mintRecipient, address burnToken, bytes32 destinationCaller, uint256 maxFee, uint32 minFinalityThreshold, bytes hookData) returns()
func (tokenMessengerV2 *TokenMessengerV2) PackDepositForBurnWithHook(amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte, maxFee *big.Int, minFinalityThreshold uint32, hookData []byte) []byte {
	enc, err := tokenMessengerV2.abi.Pack("depositForBurnWithHook", amount, destinationDomain, mintRecipient, burnToken, destinationCaller, maxFee, minFinalityThreshold, hookData)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDepositForBurnWithHook is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x779b432d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function depositForBurnWithHook(uint256 amount, uint32 destinationDomain, bytes32 mintRecipient, address burnToken, bytes32 destinationCaller, uint256 maxFee, uint32 minFinalityThreshold, bytes hookData) returns()
func (tokenMessengerV2 *TokenMessengerV2) TryPackDepositForBurnWithHook(amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte, maxFee *big.Int, minFinalityThreshold uint32, hookData []byte) ([]byte, error) {
	return tokenMessengerV2.abi.Pack("depositForBurnWithHook", amount, destinationDomain, mintRecipient, burnToken, destinationCaller, maxFee, minFinalityThreshold, hookData)
}

// PackFeeRecipient is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x46904840.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function feeRecipient() view returns(address)
func (tokenMessengerV2 *TokenMessengerV2) PackFeeRecipient() []byte {
	enc, err := tokenMessengerV2.abi.Pack("feeRecipient")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFeeRecipient is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x46904840.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function feeRecipient() view returns(address)
func (tokenMessengerV2 *TokenMessengerV2) TryPackFeeRecipient() ([]byte, error) {
	return tokenMessengerV2.abi.Pack("feeRecipient")
}

// UnpackFeeRecipient is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (tokenMessengerV2 *TokenMessengerV2) UnpackFeeRecipient(data []byte) (common.Address, error) {
	out, err := tokenMessengerV2.abi.Unpack("feeRecipient", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetMinFeeAmount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x516990e3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getMinFeeAmount(uint256 amount) view returns(uint256)
func (tokenMessengerV2 *TokenMessengerV2) PackGetMinFeeAmount(amount *big.Int) []byte {
	enc, err := tokenMessengerV2.abi.Pack("getMinFeeAmount", amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetMinFeeAmount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x516990e3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getMinFeeAmount(uint256 amount) view returns(uint256)
func (tokenMessengerV2 *TokenMessengerV2) TryPackGetMinFeeAmount(amount *big.Int) ([]byte, error) {
	return tokenMessengerV2.abi.Pack("getMinFeeAmount", amount)
}

// UnpackGetMinFeeAmount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x516990e3.
//
// Solidity: function getMinFeeAmount(uint256 amount) view returns(uint256)
func (tokenMessengerV2 *TokenMessengerV2) UnpackGetMinFeeAmount(data []byte) (*big.Int, error) {
	out, err := tokenMessengerV2.abi.Unpack("getMinFeeAmount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackHandleReceiveFinalizedMessage is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x11cffb67.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function handleReceiveFinalizedMessage(uint32 remoteDomain, bytes32 sender, uint32 , bytes messageBody) returns(bool)
func (tokenMessengerV2 *TokenMessengerV2) PackHandleReceiveFinalizedMessage(remoteDomain uint32, sender [32]byte, arg2 uint32, messageBody []byte) []byte {
	enc, err := tokenMessengerV2.abi.Pack("handleReceiveFinalizedMessage", remoteDomain, sender, arg2, messageBody)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackHandleReceiveFinalizedMessage is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x11cffb67.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function handleReceiveFinalizedMessage(uint32 remoteDomain, bytes32 sender, uint32 , bytes messageBody) returns(bool)
func (tokenMessengerV2 *TokenMessengerV2) TryPackHandleReceiveFinalizedMessage(remoteDomain uint32, sender [32]byte, arg2 uint32, messageBody []byte) ([]byte, error) {
	return tokenMessengerV2.abi.Pack("handleReceiveFinalizedMessage", remoteDomain, sender, arg2, messageBody)
}

// UnpackHandleReceiveFinalizedMessage is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x11cffb67.
//
// Solidity: function handleReceiveFinalizedMessage(uint32 remoteDomain, bytes32 sender, uint32 , bytes messageBody) returns(bool)
func (tokenMessengerV2 *TokenMessengerV2) UnpackHandleReceiveFinalizedMessage(data []byte) (bool, error) {
	out, err := tokenMessengerV2.abi.Unpack("handleReceiveFinalizedMessage", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackHandleReceiveUnfinalizedMessage is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7c92f219.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function handleReceiveUnfinalizedMessage(uint32 remoteDomain, bytes32 sender, uint32 finalityThresholdExecuted, bytes messageBody) returns(bool)
func (tokenMessengerV2 *TokenMessengerV2) PackHandleReceiveUnfinalizedMessage(remoteDomain uint32, sender [32]byte, finalityThresholdExecuted uint32, messageBody []byte) []byte {
	enc, err := tokenMessengerV2.abi.Pack("handleReceiveUnfinalizedMessage", remoteDomain, sender, finalityThresholdExecuted, messageBody)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackHandleReceiveUnfinalizedMessage is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7c92f219.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function handleReceiveUnfinalizedMessage(uint32 remoteDomain, bytes32 sender, uint32 finalityThresholdExecuted, bytes messageBody) returns(bool)
func (tokenMessengerV2 *TokenMessengerV2) TryPackHandleReceiveUnfinalizedMessage(remoteDomain uint32, sender [32]byte, finalityThresholdExecuted uint32, messageBody []byte) ([]byte, error) {
	return tokenMessengerV2.abi.Pack("handleReceiveUnfinalizedMessage", remoteDomain, sender, finalityThresholdExecuted, messageBody)
}

// UnpackHandleReceiveUnfinalizedMessage is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7c92f219.
//
// Solidity: function handleReceiveUnfinalizedMessage(uint32 remoteDomain, bytes32 sender, uint32 finalityThresholdExecuted, bytes messageBody) returns(bool)
func (tokenMessengerV2 *TokenMessengerV2) UnpackHandleReceiveUnfinalizedMessage(data []byte) (bool, error) {
	out, err := tokenMessengerV2.abi.Unpack("handleReceiveUnfinalizedMessage", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x02db402e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function initialize((address,address,address,address,address,address) roles, uint256 minFee_, uint32[] remoteDomains_, bytes32[] remoteTokenMessengers_) returns()
func (tokenMessengerV2 *TokenMessengerV2) PackInitialize(roles TokenMessengerV2TokenMessengerV2Roles, minFee *big.Int, remoteDomains []uint32, remoteTokenMessengers [][32]byte) []byte {
	enc, err := tokenMessengerV2.abi.Pack("initialize", roles, minFee, remoteDomains, remoteTokenMessengers)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x02db402e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function initialize((address,address,address,address,address,address) roles, uint256 minFee_, uint32[] remoteDomains_, bytes32[] remoteTokenMessengers_) returns()
func (tokenMessengerV2 *TokenMessengerV2) TryPackInitialize(roles TokenMessengerV2TokenMessengerV2Roles, minFee *big.Int, remoteDomains []uint32, remoteTokenMessengers [][32]byte) ([]byte, error) {
	return tokenMessengerV2.abi.Pack("initialize", roles, minFee, remoteDomains, remoteTokenMessengers)
}

// PackInitializedVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x08828eb7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function initializedVersion() view returns(uint64)
func (tokenMessengerV2 *TokenMessengerV2) PackInitializedVersion() []byte {
	enc, err := tokenMessengerV2.abi.Pack("initializedVersion")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackInitializedVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x08828eb7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function initializedVersion() view returns(uint64)
func (tokenMessengerV2 *TokenMessengerV2) TryPackInitializedVersion() ([]byte, error) {
	return tokenMessengerV2.abi.Pack("initializedVersion")
}

// UnpackInitializedVersion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x08828eb7.
//
// Solidity: function initializedVersion() view returns(uint64)
func (tokenMessengerV2 *TokenMessengerV2) UnpackInitializedVersion(data []byte) (uint64, error) {
	out, err := tokenMessengerV2.abi.Unpack("initializedVersion", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackIsDenylisted is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe877a526.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isDenylisted(address account) view returns(bool)
func (tokenMessengerV2 *TokenMessengerV2) PackIsDenylisted(account common.Address) []byte {
	enc, err := tokenMessengerV2.abi.Pack("isDenylisted", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsDenylisted is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe877a526.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isDenylisted(address account) view returns(bool)
func (tokenMessengerV2 *TokenMessengerV2) TryPackIsDenylisted(account common.Address) ([]byte, error) {
	return tokenMessengerV2.abi.Pack("isDenylisted", account)
}

// UnpackIsDenylisted is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe877a526.
//
// Solidity: function isDenylisted(address account) view returns(bool)
func (tokenMessengerV2 *TokenMessengerV2) UnpackIsDenylisted(data []byte) (bool, error) {
	out, err := tokenMessengerV2.abi.Unpack("isDenylisted", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackLocalMessageTransmitter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2c121921.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function localMessageTransmitter() view returns(address)
func (tokenMessengerV2 *TokenMessengerV2) PackLocalMessageTransmitter() []byte {
	enc, err := tokenMessengerV2.abi.Pack("localMessageTransmitter")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackLocalMessageTransmitter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2c121921.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function localMessageTransmitter() view returns(address)
func (tokenMessengerV2 *TokenMessengerV2) TryPackLocalMessageTransmitter() ([]byte, error) {
	return tokenMessengerV2.abi.Pack("localMessageTransmitter")
}

// UnpackLocalMessageTransmitter is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2c121921.
//
// Solidity: function localMessageTransmitter() view returns(address)
func (tokenMessengerV2 *TokenMessengerV2) UnpackLocalMessageTransmitter(data []byte) (common.Address, error) {
	out, err := tokenMessengerV2.abi.Unpack("localMessageTransmitter", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackLocalMinter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcb75c11c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function localMinter() view returns(address)
func (tokenMessengerV2 *TokenMessengerV2) PackLocalMinter() []byte {
	enc, err := tokenMessengerV2.abi.Pack("localMinter")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackLocalMinter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcb75c11c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function localMinter() view returns(address)
func (tokenMessengerV2 *TokenMessengerV2) TryPackLocalMinter() ([]byte, error) {
	return tokenMessengerV2.abi.Pack("localMinter")
}

// UnpackLocalMinter is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcb75c11c.
//
// Solidity: function localMinter() view returns(address)
func (tokenMessengerV2 *TokenMessengerV2) UnpackLocalMinter(data []byte) (common.Address, error) {
	out, err := tokenMessengerV2.abi.Unpack("localMinter", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackMessageBodyVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9cdbb181.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function messageBodyVersion() view returns(uint32)
func (tokenMessengerV2 *TokenMessengerV2) PackMessageBodyVersion() []byte {
	enc, err := tokenMessengerV2.abi.Pack("messageBodyVersion")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMessageBodyVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9cdbb181.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function messageBodyVersion() view returns(uint32)
func (tokenMessengerV2 *TokenMessengerV2) TryPackMessageBodyVersion() ([]byte, error) {
	return tokenMessengerV2.abi.Pack("messageBodyVersion")
}

// UnpackMessageBodyVersion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9cdbb181.
//
// Solidity: function messageBodyVersion() view returns(uint32)
func (tokenMessengerV2 *TokenMessengerV2) UnpackMessageBodyVersion(data []byte) (uint32, error) {
	out, err := tokenMessengerV2.abi.Unpack("messageBodyVersion", data)
	if err != nil {
		return *new(uint32), err
	}
	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	return out0, nil
}

// PackMinFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x24ec7590.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function minFee() view returns(uint256)
func (tokenMessengerV2 *TokenMessengerV2) PackMinFee() []byte {
	enc, err := tokenMessengerV2.abi.Pack("minFee")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMinFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x24ec7590.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function minFee() view returns(uint256)
func (tokenMessengerV2 *TokenMessengerV2) TryPackMinFee() ([]byte, error) {
	return tokenMessengerV2.abi.Pack("minFee")
}

// UnpackMinFee is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x24ec7590.
//
// Solidity: function minFee() view returns(uint256)
func (tokenMessengerV2 *TokenMessengerV2) UnpackMinFee(data []byte) (*big.Int, error) {
	out, err := tokenMessengerV2.abi.Unpack("minFee", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackMinFeeController is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x369c4f1b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function minFeeController() view returns(address)
func (tokenMessengerV2 *TokenMessengerV2) PackMinFeeController() []byte {
	enc, err := tokenMessengerV2.abi.Pack("minFeeController")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMinFeeController is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x369c4f1b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function minFeeController() view returns(address)
func (tokenMessengerV2 *TokenMessengerV2) TryPackMinFeeController() ([]byte, error) {
	return tokenMessengerV2.abi.Pack("minFeeController")
}

// UnpackMinFeeController is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x369c4f1b.
//
// Solidity: function minFeeController() view returns(address)
func (tokenMessengerV2 *TokenMessengerV2) UnpackMinFeeController(data []byte) (common.Address, error) {
	out, err := tokenMessengerV2.abi.Unpack("minFeeController", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function owner() view returns(address)
func (tokenMessengerV2 *TokenMessengerV2) PackOwner() []byte {
	enc, err := tokenMessengerV2.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function owner() view returns(address)
func (tokenMessengerV2 *TokenMessengerV2) TryPackOwner() ([]byte, error) {
	return tokenMessengerV2.abi.Pack("owner")
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (tokenMessengerV2 *TokenMessengerV2) UnpackOwner(data []byte) (common.Address, error) {
	out, err := tokenMessengerV2.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackPendingOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe30c3978.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function pendingOwner() view returns(address)
func (tokenMessengerV2 *TokenMessengerV2) PackPendingOwner() []byte {
	enc, err := tokenMessengerV2.abi.Pack("pendingOwner")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPendingOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe30c3978.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function pendingOwner() view returns(address)
func (tokenMessengerV2 *TokenMessengerV2) TryPackPendingOwner() ([]byte, error) {
	return tokenMessengerV2.abi.Pack("pendingOwner")
}

// UnpackPendingOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (tokenMessengerV2 *TokenMessengerV2) UnpackPendingOwner(data []byte) (common.Address, error) {
	out, err := tokenMessengerV2.abi.Unpack("pendingOwner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackRemoteTokenMessengers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x82a5e665.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function remoteTokenMessengers(uint32 ) view returns(bytes32)
func (tokenMessengerV2 *TokenMessengerV2) PackRemoteTokenMessengers(arg0 uint32) []byte {
	enc, err := tokenMessengerV2.abi.Pack("remoteTokenMessengers", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoteTokenMessengers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x82a5e665.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function remoteTokenMessengers(uint32 ) view returns(bytes32)
func (tokenMessengerV2 *TokenMessengerV2) TryPackRemoteTokenMessengers(arg0 uint32) ([]byte, error) {
	return tokenMessengerV2.abi.Pack("remoteTokenMessengers", arg0)
}

// UnpackRemoteTokenMessengers is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x82a5e665.
//
// Solidity: function remoteTokenMessengers(uint32 ) view returns(bytes32)
func (tokenMessengerV2 *TokenMessengerV2) UnpackRemoteTokenMessengers(data []byte) ([32]byte, error) {
	out, err := tokenMessengerV2.abi.Unpack("remoteTokenMessengers", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackRemoveLocalMinter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91f17888.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeLocalMinter() returns()
func (tokenMessengerV2 *TokenMessengerV2) PackRemoveLocalMinter() []byte {
	enc, err := tokenMessengerV2.abi.Pack("removeLocalMinter")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveLocalMinter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91f17888.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeLocalMinter() returns()
func (tokenMessengerV2 *TokenMessengerV2) TryPackRemoveLocalMinter() ([]byte, error) {
	return tokenMessengerV2.abi.Pack("removeLocalMinter")
}

// PackRemoveRemoteTokenMessenger is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf79fd08e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeRemoteTokenMessenger(uint32 domain) returns()
func (tokenMessengerV2 *TokenMessengerV2) PackRemoveRemoteTokenMessenger(domain uint32) []byte {
	enc, err := tokenMessengerV2.abi.Pack("removeRemoteTokenMessenger", domain)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveRemoteTokenMessenger is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf79fd08e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeRemoteTokenMessenger(uint32 domain) returns()
func (tokenMessengerV2 *TokenMessengerV2) TryPackRemoveRemoteTokenMessenger(domain uint32) ([]byte, error) {
	return tokenMessengerV2.abi.Pack("removeRemoteTokenMessenger", domain)
}

// PackRescueERC20 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb2118a8d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function rescueERC20(address tokenContract, address to, uint256 amount) returns()
func (tokenMessengerV2 *TokenMessengerV2) PackRescueERC20(tokenContract common.Address, to common.Address, amount *big.Int) []byte {
	enc, err := tokenMessengerV2.abi.Pack("rescueERC20", tokenContract, to, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRescueERC20 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb2118a8d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function rescueERC20(address tokenContract, address to, uint256 amount) returns()
func (tokenMessengerV2 *TokenMessengerV2) TryPackRescueERC20(tokenContract common.Address, to common.Address, amount *big.Int) ([]byte, error) {
	return tokenMessengerV2.abi.Pack("rescueERC20", tokenContract, to, amount)
}

// PackRescuer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x38a63183.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function rescuer() view returns(address)
func (tokenMessengerV2 *TokenMessengerV2) PackRescuer() []byte {
	enc, err := tokenMessengerV2.abi.Pack("rescuer")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRescuer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x38a63183.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function rescuer() view returns(address)
func (tokenMessengerV2 *TokenMessengerV2) TryPackRescuer() ([]byte, error) {
	return tokenMessengerV2.abi.Pack("rescuer")
}

// UnpackRescuer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x38a63183.
//
// Solidity: function rescuer() view returns(address)
func (tokenMessengerV2 *TokenMessengerV2) UnpackRescuer(data []byte) (common.Address, error) {
	out, err := tokenMessengerV2.abi.Unpack("rescuer", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackSetFeeRecipient is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe74b981b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setFeeRecipient(address _feeRecipient) returns()
func (tokenMessengerV2 *TokenMessengerV2) PackSetFeeRecipient(feeRecipient common.Address) []byte {
	enc, err := tokenMessengerV2.abi.Pack("setFeeRecipient", feeRecipient)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetFeeRecipient is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe74b981b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setFeeRecipient(address _feeRecipient) returns()
func (tokenMessengerV2 *TokenMessengerV2) TryPackSetFeeRecipient(feeRecipient common.Address) ([]byte, error) {
	return tokenMessengerV2.abi.Pack("setFeeRecipient", feeRecipient)
}

// PackSetMinFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x31ac9920.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setMinFee(uint256 _minFee) returns()
func (tokenMessengerV2 *TokenMessengerV2) PackSetMinFee(minFee *big.Int) []byte {
	enc, err := tokenMessengerV2.abi.Pack("setMinFee", minFee)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetMinFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x31ac9920.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setMinFee(uint256 _minFee) returns()
func (tokenMessengerV2 *TokenMessengerV2) TryPackSetMinFee(minFee *big.Int) ([]byte, error) {
	return tokenMessengerV2.abi.Pack("setMinFee", minFee)
}

// PackSetMinFeeController is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa5b8d04e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setMinFeeController(address _minFeeController) returns()
func (tokenMessengerV2 *TokenMessengerV2) PackSetMinFeeController(minFeeController common.Address) []byte {
	enc, err := tokenMessengerV2.abi.Pack("setMinFeeController", minFeeController)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetMinFeeController is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa5b8d04e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setMinFeeController(address _minFeeController) returns()
func (tokenMessengerV2 *TokenMessengerV2) TryPackSetMinFeeController(minFeeController common.Address) ([]byte, error) {
	return tokenMessengerV2.abi.Pack("setMinFeeController", minFeeController)
}

// PackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (tokenMessengerV2 *TokenMessengerV2) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := tokenMessengerV2.abi.Pack("transferOwnership", newOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (tokenMessengerV2 *TokenMessengerV2) TryPackTransferOwnership(newOwner common.Address) ([]byte, error) {
	return tokenMessengerV2.abi.Pack("transferOwnership", newOwner)
}

// PackUnDenylist is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9cab0c1c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function unDenylist(address account) returns()
func (tokenMessengerV2 *TokenMessengerV2) PackUnDenylist(account common.Address) []byte {
	enc, err := tokenMessengerV2.abi.Pack("unDenylist", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUnDenylist is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9cab0c1c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function unDenylist(address account) returns()
func (tokenMessengerV2 *TokenMessengerV2) TryPackUnDenylist(account common.Address) ([]byte, error) {
	return tokenMessengerV2.abi.Pack("unDenylist", account)
}

// PackUpdateDenylister is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa946de04.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updateDenylister(address newDenylister) returns()
func (tokenMessengerV2 *TokenMessengerV2) PackUpdateDenylister(newDenylister common.Address) []byte {
	enc, err := tokenMessengerV2.abi.Pack("updateDenylister", newDenylister)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUpdateDenylister is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa946de04.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function updateDenylister(address newDenylister) returns()
func (tokenMessengerV2 *TokenMessengerV2) TryPackUpdateDenylister(newDenylister common.Address) ([]byte, error) {
	return tokenMessengerV2.abi.Pack("updateDenylister", newDenylister)
}

// PackUpdateRescuer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2ab60045.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updateRescuer(address newRescuer) returns()
func (tokenMessengerV2 *TokenMessengerV2) PackUpdateRescuer(newRescuer common.Address) []byte {
	enc, err := tokenMessengerV2.abi.Pack("updateRescuer", newRescuer)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUpdateRescuer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2ab60045.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function updateRescuer(address newRescuer) returns()
func (tokenMessengerV2 *TokenMessengerV2) TryPackUpdateRescuer(newRescuer common.Address) ([]byte, error) {
	return tokenMessengerV2.abi.Pack("updateRescuer", newRescuer)
}

// TokenMessengerV2Denylisted represents a Denylisted event raised by the TokenMessengerV2 contract.
type TokenMessengerV2Denylisted struct {
	Account common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const TokenMessengerV2DenylistedEventName = "Denylisted"

// ContractEventName returns the user-defined event name.
func (TokenMessengerV2Denylisted) ContractEventName() string {
	return TokenMessengerV2DenylistedEventName
}

// UnpackDenylistedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Denylisted(address indexed account)
func (tokenMessengerV2 *TokenMessengerV2) UnpackDenylistedEvent(log *types.Log) (*TokenMessengerV2Denylisted, error) {
	event := "Denylisted"
	if len(log.Topics) == 0 || log.Topics[0] != tokenMessengerV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenMessengerV2Denylisted)
	if len(log.Data) > 0 {
		if err := tokenMessengerV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenMessengerV2.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// TokenMessengerV2DenylisterChanged represents a DenylisterChanged event raised by the TokenMessengerV2 contract.
type TokenMessengerV2DenylisterChanged struct {
	OldDenylister common.Address
	NewDenylister common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const TokenMessengerV2DenylisterChangedEventName = "DenylisterChanged"

// ContractEventName returns the user-defined event name.
func (TokenMessengerV2DenylisterChanged) ContractEventName() string {
	return TokenMessengerV2DenylisterChangedEventName
}

// UnpackDenylisterChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DenylisterChanged(address indexed oldDenylister, address indexed newDenylister)
func (tokenMessengerV2 *TokenMessengerV2) UnpackDenylisterChangedEvent(log *types.Log) (*TokenMessengerV2DenylisterChanged, error) {
	event := "DenylisterChanged"
	if len(log.Topics) == 0 || log.Topics[0] != tokenMessengerV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenMessengerV2DenylisterChanged)
	if len(log.Data) > 0 {
		if err := tokenMessengerV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenMessengerV2.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// TokenMessengerV2DepositForBurn represents a DepositForBurn event raised by the TokenMessengerV2 contract.
type TokenMessengerV2DepositForBurn struct {
	BurnToken                 common.Address
	Amount                    *big.Int
	Depositor                 common.Address
	MintRecipient             [32]byte
	DestinationDomain         uint32
	DestinationTokenMessenger [32]byte
	DestinationCaller         [32]byte
	MaxFee                    *big.Int
	MinFinalityThreshold      uint32
	HookData                  []byte
	Raw                       *types.Log // Blockchain specific contextual infos
}

const TokenMessengerV2DepositForBurnEventName = "DepositForBurn"

// ContractEventName returns the user-defined event name.
func (TokenMessengerV2DepositForBurn) ContractEventName() string {
	return TokenMessengerV2DepositForBurnEventName
}

// UnpackDepositForBurnEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DepositForBurn(address indexed burnToken, uint256 amount, address indexed depositor, bytes32 mintRecipient, uint32 destinationDomain, bytes32 destinationTokenMessenger, bytes32 destinationCaller, uint256 maxFee, uint32 indexed minFinalityThreshold, bytes hookData)
func (tokenMessengerV2 *TokenMessengerV2) UnpackDepositForBurnEvent(log *types.Log) (*TokenMessengerV2DepositForBurn, error) {
	event := "DepositForBurn"
	if len(log.Topics) == 0 || log.Topics[0] != tokenMessengerV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenMessengerV2DepositForBurn)
	if len(log.Data) > 0 {
		if err := tokenMessengerV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenMessengerV2.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// TokenMessengerV2FeeRecipientSet represents a FeeRecipientSet event raised by the TokenMessengerV2 contract.
type TokenMessengerV2FeeRecipientSet struct {
	FeeRecipient common.Address
	Raw          *types.Log // Blockchain specific contextual infos
}

const TokenMessengerV2FeeRecipientSetEventName = "FeeRecipientSet"

// ContractEventName returns the user-defined event name.
func (TokenMessengerV2FeeRecipientSet) ContractEventName() string {
	return TokenMessengerV2FeeRecipientSetEventName
}

// UnpackFeeRecipientSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event FeeRecipientSet(address feeRecipient)
func (tokenMessengerV2 *TokenMessengerV2) UnpackFeeRecipientSetEvent(log *types.Log) (*TokenMessengerV2FeeRecipientSet, error) {
	event := "FeeRecipientSet"
	if len(log.Topics) == 0 || log.Topics[0] != tokenMessengerV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenMessengerV2FeeRecipientSet)
	if len(log.Data) > 0 {
		if err := tokenMessengerV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenMessengerV2.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// TokenMessengerV2Initialized represents a Initialized event raised by the TokenMessengerV2 contract.
type TokenMessengerV2Initialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const TokenMessengerV2InitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (TokenMessengerV2Initialized) ContractEventName() string {
	return TokenMessengerV2InitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (tokenMessengerV2 *TokenMessengerV2) UnpackInitializedEvent(log *types.Log) (*TokenMessengerV2Initialized, error) {
	event := "Initialized"
	if len(log.Topics) == 0 || log.Topics[0] != tokenMessengerV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenMessengerV2Initialized)
	if len(log.Data) > 0 {
		if err := tokenMessengerV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenMessengerV2.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// TokenMessengerV2LocalMinterAdded represents a LocalMinterAdded event raised by the TokenMessengerV2 contract.
type TokenMessengerV2LocalMinterAdded struct {
	LocalMinter common.Address
	Raw         *types.Log // Blockchain specific contextual infos
}

const TokenMessengerV2LocalMinterAddedEventName = "LocalMinterAdded"

// ContractEventName returns the user-defined event name.
func (TokenMessengerV2LocalMinterAdded) ContractEventName() string {
	return TokenMessengerV2LocalMinterAddedEventName
}

// UnpackLocalMinterAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event LocalMinterAdded(address localMinter)
func (tokenMessengerV2 *TokenMessengerV2) UnpackLocalMinterAddedEvent(log *types.Log) (*TokenMessengerV2LocalMinterAdded, error) {
	event := "LocalMinterAdded"
	if len(log.Topics) == 0 || log.Topics[0] != tokenMessengerV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenMessengerV2LocalMinterAdded)
	if len(log.Data) > 0 {
		if err := tokenMessengerV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenMessengerV2.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// TokenMessengerV2LocalMinterRemoved represents a LocalMinterRemoved event raised by the TokenMessengerV2 contract.
type TokenMessengerV2LocalMinterRemoved struct {
	LocalMinter common.Address
	Raw         *types.Log // Blockchain specific contextual infos
}

const TokenMessengerV2LocalMinterRemovedEventName = "LocalMinterRemoved"

// ContractEventName returns the user-defined event name.
func (TokenMessengerV2LocalMinterRemoved) ContractEventName() string {
	return TokenMessengerV2LocalMinterRemovedEventName
}

// UnpackLocalMinterRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event LocalMinterRemoved(address localMinter)
func (tokenMessengerV2 *TokenMessengerV2) UnpackLocalMinterRemovedEvent(log *types.Log) (*TokenMessengerV2LocalMinterRemoved, error) {
	event := "LocalMinterRemoved"
	if len(log.Topics) == 0 || log.Topics[0] != tokenMessengerV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenMessengerV2LocalMinterRemoved)
	if len(log.Data) > 0 {
		if err := tokenMessengerV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenMessengerV2.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// TokenMessengerV2MinFeeControllerSet represents a MinFeeControllerSet event raised by the TokenMessengerV2 contract.
type TokenMessengerV2MinFeeControllerSet struct {
	MinFeeController common.Address
	Raw              *types.Log // Blockchain specific contextual infos
}

const TokenMessengerV2MinFeeControllerSetEventName = "MinFeeControllerSet"

// ContractEventName returns the user-defined event name.
func (TokenMessengerV2MinFeeControllerSet) ContractEventName() string {
	return TokenMessengerV2MinFeeControllerSetEventName
}

// UnpackMinFeeControllerSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MinFeeControllerSet(address minFeeController)
func (tokenMessengerV2 *TokenMessengerV2) UnpackMinFeeControllerSetEvent(log *types.Log) (*TokenMessengerV2MinFeeControllerSet, error) {
	event := "MinFeeControllerSet"
	if len(log.Topics) == 0 || log.Topics[0] != tokenMessengerV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenMessengerV2MinFeeControllerSet)
	if len(log.Data) > 0 {
		if err := tokenMessengerV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenMessengerV2.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// TokenMessengerV2MinFeeSet represents a MinFeeSet event raised by the TokenMessengerV2 contract.
type TokenMessengerV2MinFeeSet struct {
	MinFee *big.Int
	Raw    *types.Log // Blockchain specific contextual infos
}

const TokenMessengerV2MinFeeSetEventName = "MinFeeSet"

// ContractEventName returns the user-defined event name.
func (TokenMessengerV2MinFeeSet) ContractEventName() string {
	return TokenMessengerV2MinFeeSetEventName
}

// UnpackMinFeeSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MinFeeSet(uint256 minFee)
func (tokenMessengerV2 *TokenMessengerV2) UnpackMinFeeSetEvent(log *types.Log) (*TokenMessengerV2MinFeeSet, error) {
	event := "MinFeeSet"
	if len(log.Topics) == 0 || log.Topics[0] != tokenMessengerV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenMessengerV2MinFeeSet)
	if len(log.Data) > 0 {
		if err := tokenMessengerV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenMessengerV2.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// TokenMessengerV2MintAndWithdraw represents a MintAndWithdraw event raised by the TokenMessengerV2 contract.
type TokenMessengerV2MintAndWithdraw struct {
	MintRecipient common.Address
	Amount        *big.Int
	MintToken     common.Address
	FeeCollected  *big.Int
	Raw           *types.Log // Blockchain specific contextual infos
}

const TokenMessengerV2MintAndWithdrawEventName = "MintAndWithdraw"

// ContractEventName returns the user-defined event name.
func (TokenMessengerV2MintAndWithdraw) ContractEventName() string {
	return TokenMessengerV2MintAndWithdrawEventName
}

// UnpackMintAndWithdrawEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MintAndWithdraw(address indexed mintRecipient, uint256 amount, address indexed mintToken, uint256 feeCollected)
func (tokenMessengerV2 *TokenMessengerV2) UnpackMintAndWithdrawEvent(log *types.Log) (*TokenMessengerV2MintAndWithdraw, error) {
	event := "MintAndWithdraw"
	if len(log.Topics) == 0 || log.Topics[0] != tokenMessengerV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenMessengerV2MintAndWithdraw)
	if len(log.Data) > 0 {
		if err := tokenMessengerV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenMessengerV2.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// TokenMessengerV2OwnershipTransferStarted represents a OwnershipTransferStarted event raised by the TokenMessengerV2 contract.
type TokenMessengerV2OwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const TokenMessengerV2OwnershipTransferStartedEventName = "OwnershipTransferStarted"

// ContractEventName returns the user-defined event name.
func (TokenMessengerV2OwnershipTransferStarted) ContractEventName() string {
	return TokenMessengerV2OwnershipTransferStartedEventName
}

// UnpackOwnershipTransferStartedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (tokenMessengerV2 *TokenMessengerV2) UnpackOwnershipTransferStartedEvent(log *types.Log) (*TokenMessengerV2OwnershipTransferStarted, error) {
	event := "OwnershipTransferStarted"
	if len(log.Topics) == 0 || log.Topics[0] != tokenMessengerV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenMessengerV2OwnershipTransferStarted)
	if len(log.Data) > 0 {
		if err := tokenMessengerV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenMessengerV2.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// TokenMessengerV2OwnershipTransferred represents a OwnershipTransferred event raised by the TokenMessengerV2 contract.
type TokenMessengerV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const TokenMessengerV2OwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (TokenMessengerV2OwnershipTransferred) ContractEventName() string {
	return TokenMessengerV2OwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (tokenMessengerV2 *TokenMessengerV2) UnpackOwnershipTransferredEvent(log *types.Log) (*TokenMessengerV2OwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if len(log.Topics) == 0 || log.Topics[0] != tokenMessengerV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenMessengerV2OwnershipTransferred)
	if len(log.Data) > 0 {
		if err := tokenMessengerV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenMessengerV2.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// TokenMessengerV2RemoteTokenMessengerAdded represents a RemoteTokenMessengerAdded event raised by the TokenMessengerV2 contract.
type TokenMessengerV2RemoteTokenMessengerAdded struct {
	Domain         uint32
	TokenMessenger [32]byte
	Raw            *types.Log // Blockchain specific contextual infos
}

const TokenMessengerV2RemoteTokenMessengerAddedEventName = "RemoteTokenMessengerAdded"

// ContractEventName returns the user-defined event name.
func (TokenMessengerV2RemoteTokenMessengerAdded) ContractEventName() string {
	return TokenMessengerV2RemoteTokenMessengerAddedEventName
}

// UnpackRemoteTokenMessengerAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RemoteTokenMessengerAdded(uint32 domain, bytes32 tokenMessenger)
func (tokenMessengerV2 *TokenMessengerV2) UnpackRemoteTokenMessengerAddedEvent(log *types.Log) (*TokenMessengerV2RemoteTokenMessengerAdded, error) {
	event := "RemoteTokenMessengerAdded"
	if len(log.Topics) == 0 || log.Topics[0] != tokenMessengerV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenMessengerV2RemoteTokenMessengerAdded)
	if len(log.Data) > 0 {
		if err := tokenMessengerV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenMessengerV2.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// TokenMessengerV2RemoteTokenMessengerRemoved represents a RemoteTokenMessengerRemoved event raised by the TokenMessengerV2 contract.
type TokenMessengerV2RemoteTokenMessengerRemoved struct {
	Domain         uint32
	TokenMessenger [32]byte
	Raw            *types.Log // Blockchain specific contextual infos
}

const TokenMessengerV2RemoteTokenMessengerRemovedEventName = "RemoteTokenMessengerRemoved"

// ContractEventName returns the user-defined event name.
func (TokenMessengerV2RemoteTokenMessengerRemoved) ContractEventName() string {
	return TokenMessengerV2RemoteTokenMessengerRemovedEventName
}

// UnpackRemoteTokenMessengerRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RemoteTokenMessengerRemoved(uint32 domain, bytes32 tokenMessenger)
func (tokenMessengerV2 *TokenMessengerV2) UnpackRemoteTokenMessengerRemovedEvent(log *types.Log) (*TokenMessengerV2RemoteTokenMessengerRemoved, error) {
	event := "RemoteTokenMessengerRemoved"
	if len(log.Topics) == 0 || log.Topics[0] != tokenMessengerV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenMessengerV2RemoteTokenMessengerRemoved)
	if len(log.Data) > 0 {
		if err := tokenMessengerV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenMessengerV2.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// TokenMessengerV2RescuerChanged represents a RescuerChanged event raised by the TokenMessengerV2 contract.
type TokenMessengerV2RescuerChanged struct {
	NewRescuer common.Address
	Raw        *types.Log // Blockchain specific contextual infos
}

const TokenMessengerV2RescuerChangedEventName = "RescuerChanged"

// ContractEventName returns the user-defined event name.
func (TokenMessengerV2RescuerChanged) ContractEventName() string {
	return TokenMessengerV2RescuerChangedEventName
}

// UnpackRescuerChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RescuerChanged(address indexed newRescuer)
func (tokenMessengerV2 *TokenMessengerV2) UnpackRescuerChangedEvent(log *types.Log) (*TokenMessengerV2RescuerChanged, error) {
	event := "RescuerChanged"
	if len(log.Topics) == 0 || log.Topics[0] != tokenMessengerV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenMessengerV2RescuerChanged)
	if len(log.Data) > 0 {
		if err := tokenMessengerV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenMessengerV2.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// TokenMessengerV2UnDenylisted represents a UnDenylisted event raised by the TokenMessengerV2 contract.
type TokenMessengerV2UnDenylisted struct {
	Account common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const TokenMessengerV2UnDenylistedEventName = "UnDenylisted"

// ContractEventName returns the user-defined event name.
func (TokenMessengerV2UnDenylisted) ContractEventName() string {
	return TokenMessengerV2UnDenylistedEventName
}

// UnpackUnDenylistedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event UnDenylisted(address indexed account)
func (tokenMessengerV2 *TokenMessengerV2) UnpackUnDenylistedEvent(log *types.Log) (*TokenMessengerV2UnDenylisted, error) {
	event := "UnDenylisted"
	if len(log.Topics) == 0 || log.Topics[0] != tokenMessengerV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TokenMessengerV2UnDenylisted)
	if len(log.Data) > 0 {
		if err := tokenMessengerV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range tokenMessengerV2.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// TypedMemViewMetaData contains all meta data concerning the TypedMemView contract.
var TypedMemViewMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"NULL\",\"outputs\":[{\"internalType\":\"bytes29\",\"name\":\"\",\"type\":\"bytes29\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "cb2086366b9cd31468d55b24e8a10d873e",
	Bin: "0x60b4610024600b82828239805160001a607314601757fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361060335760003560e01c8063f26be3fc146038575b600080fd5b603e605a565b604051808262ffffff1916815260200191505060405180910390f35b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000008156fea2646970667358221220fddb4c6b8b1f022c0bf4450f4ed1996879dda9fef5fd337d87742923949f247a64736f6c63430007060033",
}

// TypedMemView is an auto generated Go binding around an Ethereum contract.
type TypedMemView struct {
	abi abi.ABI
}

// NewTypedMemView creates a new instance of TypedMemView.
func NewTypedMemView() *TypedMemView {
	parsed, err := TypedMemViewMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &TypedMemView{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *TypedMemView) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackNULL is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf26be3fc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function NULL() view returns(bytes29)
func (typedMemView *TypedMemView) PackNULL() []byte {
	enc, err := typedMemView.abi.Pack("NULL")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackNULL is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf26be3fc.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function NULL() view returns(bytes29)
func (typedMemView *TypedMemView) TryPackNULL() ([]byte, error) {
	return typedMemView.abi.Pack("NULL")
}

// UnpackNULL is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf26be3fc.
//
// Solidity: function NULL() view returns(bytes29)
func (typedMemView *TypedMemView) UnpackNULL(data []byte) ([29]byte, error) {
	out, err := typedMemView.abi.Unpack("NULL", data)
	if err != nil {
		return *new([29]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([29]byte)).(*[29]byte)
	return out0, nil
}
