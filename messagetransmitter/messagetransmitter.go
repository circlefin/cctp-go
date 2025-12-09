// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package messagetransmitter

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

// AttestableMetaData contains all meta data concerning the Attestable contract.
var AttestableMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"}],\"name\":\"AttesterDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"}],\"name\":\"AttesterEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousAttesterManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAttesterManager\",\"type\":\"address\"}],\"name\":\"AttesterManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldSignatureThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSignatureThreshold\",\"type\":\"uint256\"}],\"name\":\"SignatureThresholdUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"attesterManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"}],\"name\":\"disableAttester\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAttester\",\"type\":\"address\"}],\"name\":\"enableAttester\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getEnabledAttester\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumEnabledAttesters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"}],\"name\":\"isEnabledAttester\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newSignatureThreshold\",\"type\":\"uint256\"}],\"name\":\"setSignatureThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signatureThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAttesterManager\",\"type\":\"address\"}],\"name\":\"updateAttesterManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "b867694248a4b71c6f3f96468a6f557e97",
	Bin: "0x60806040523480156200001157600080fd5b50604051620019d5380380620019d5833981810160405260208110156200003757600080fd5b8101908080519060200190929190505050620000686200005c6200009960201b60201c565b620000a160201b60201c565b6200007933620000df60201b60201c565b600160028190555062000092816200024960201b60201c565b5062000632565b600033905090565b600160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055620000dc816200032160201b620009ff1760201c565b50565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141562000183576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f496e76616c6964206174746573746572206d616e61676572206164647265737381525060200191505060405180910390fd5b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f0cee1b7ae04f3c788dd3a46c6fa677eb95b913611ef7ab59524fdc09d346021960405160405180910390a35050565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146200030d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f43616c6c6572206e6f74206174746573746572206d616e61676572000000000081525060200191505060405180910390fd5b6200031e81620003e560201b60201c565b50565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141562000489576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f4e6577206174746573746572206d757374206265206e6f6e7a65726f0000000081525060200191505060405180910390fd5b620004a48160036200055d60201b62000ac31790919060201c565b62000517576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f417474657374657220616c726561647920656e61626c6564000000000000000081525060200191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff167f5b99bab45c72ce67e89466dbc47480b9c1fde1400e7268bbf463b8354ee4653f60405160405180910390a250565b60006200058d836000018373ffffffffffffffffffffffffffffffffffffffff1660001b6200059560201b60201c565b905092915050565b6000620005a983836200060f60201b60201c565b6200060457826000018290806001815401808255809150506001900390600052602060002001600090919091909150558260000180549050836001016000848152602001908152602001600020819055506001905062000609565b600090505b92915050565b600080836001016000848152602001908152602001600020541415905092915050565b61139380620006426000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c8063a82f2e261161008c578063de7769d411610066578063de7769d4146102a6578063e30c3978146102ea578063f2fde38b1461031e578063fae3687914610362576100cf565b8063a82f2e2614610202578063bbde537414610220578063beb673d81461024e576100cf565b80632d025080146100d457806351079a531461011857806379ba5097146101365780637af82f60146101405780638da5cb5b1461019a5780639b0d94b7146101ce575b600080fd5b610116600480360360208110156100ea57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506103a6565b005b61012061062f565b6040518082815260200191505060405180910390f35b61013e610640565b005b6101826004803603602081101561015657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506106e3565b60405180821515815260200191505060405180910390f35b6101a2610700565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101d6610729565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61020a610753565b6040518082815260200191505060405180910390f35b61024c6004803603602081101561023657600080fd5b8101908080359060200190929190505050610759565b005b61027a6004803603602081101561026457600080fd5b8101908080359060200190929190505050610828565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6102e8600480360360208110156102bc57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610845565b005b6102f2610859565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6103606004803603602081101561033457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610883565b005b6103a46004803603602081101561037857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610930565b005b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610469576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f43616c6c6572206e6f74206174746573746572206d616e61676572000000000081525060200191505060405180910390fd5b600061047361062f565b9050600181116104eb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f546f6f2066657720656e61626c6564206174746573746572730000000000000081525060200191505060405180910390fd5b6002548111610562576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f5369676e6174757265207468726573686f6c6420697320746f6f206c6f77000081525060200191505060405180910390fd5b610576826003610af390919063ffffffff16565b6105e8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f417474657374657220616c72656164792064697361626c65640000000000000081525060200191505060405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff167f78e573a18c75957b7cadaab01511aa1c19a659f06ecf53e01de37ed92d3261fc60405160405180910390a25050565b600061063b6003610b23565b905090565b600061064a610b38565b90508073ffffffffffffffffffffffffffffffffffffffff1661066b610859565b73ffffffffffffffffffffffffffffffffffffffff16146106d7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260298152602001806113356029913960400191505060405180910390fd5b6106e081610b40565b50565b60006106f9826003610b7190919063ffffffff16565b9050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60025481565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461081c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f43616c6c6572206e6f74206174746573746572206d616e61676572000000000081525060200191505060405180910390fd5b61082581610ba1565b50565b600061083e826003610d6090919063ffffffff16565b9050919050565b61084d610d7a565b61085681610e2b565b50565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61088b610d7a565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff166108eb610700565b73ffffffffffffffffffffffffffffffffffffffff167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a350565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109f3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f43616c6c6572206e6f74206174746573746572206d616e61676572000000000081525060200191505060405180910390fd5b6109fc81610f94565b50565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000610aeb836000018373ffffffffffffffffffffffffffffffffffffffff1660001b611103565b905092915050565b6000610b1b836000018373ffffffffffffffffffffffffffffffffffffffff1660001b611173565b905092915050565b6000610b318260000161125b565b9050919050565b600033905090565b600160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055610b6e816109ff565b50565b6000610b99836000018373ffffffffffffffffffffffffffffffffffffffff1660001b61126c565b905092915050565b6000811415610c18576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f496e76616c6964207369676e6174757265207468726573686f6c64000000000081525060200191505060405180910390fd5b610c226003610b23565b811115610c97576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4e6577207369676e6174757265207468726573686f6c6420746f6f206869676881525060200191505060405180910390fd5b600254811415610d0f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f5369676e6174757265207468726573686f6c6420616c7265616479207365740081525060200191505060405180910390fd5b60006002549050816002819055507f149153f58b4da003a8cfd4523709a202402182cb5aa335046911277a1be6eede8183604051808381526020018281526020019250505060405180910390a15050565b6000610d6f836000018361128f565b60001c905092915050565b610d82610b38565b73ffffffffffffffffffffffffffffffffffffffff16610da0610700565b73ffffffffffffffffffffffffffffffffffffffff1614610e29576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610ece576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f496e76616c6964206174746573746572206d616e61676572206164647265737381525060200191505060405180910390fd5b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f0cee1b7ae04f3c788dd3a46c6fa677eb95b913611ef7ab59524fdc09d346021960405160405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415611037576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f4e6577206174746573746572206d757374206265206e6f6e7a65726f0000000081525060200191505060405180910390fd5b61104b816003610ac390919063ffffffff16565b6110bd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f417474657374657220616c726561647920656e61626c6564000000000000000081525060200191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff167f5b99bab45c72ce67e89466dbc47480b9c1fde1400e7268bbf463b8354ee4653f60405160405180910390a250565b600061110f838361126c565b61116857826000018290806001815401808255809150506001900390600052602060002001600090919091909150558260000180549050836001016000848152602001908152602001600020819055506001905061116d565b600090505b92915050565b6000808360010160008481526020019081526020016000205490506000811461124f57600060018203905060006001866000018054905003905060008660000182815481106111be57fe5b90600052602060002001549050808760000184815481106111db57fe5b906000526020600020018190555060018301876001016000838152602001908152602001600020819055508660000180548061121357fe5b60019003818190600052602060002001600090559055866001016000878152602001908152602001600020600090556001945050505050611255565b60009150505b92915050565b600081600001805490509050919050565b600080836001016000848152602001908152602001600020541415905092915050565b6000818360000180549050116112f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806113136022913960400191505060405180910390fd5b8260000182815481106112ff57fe5b906000526020600020015490509291505056fe456e756d657261626c655365743a20696e646578206f7574206f6620626f756e64734f776e61626c6532537465703a2063616c6c6572206973206e6f7420746865206e6577206f776e6572a26469706673582212202a1f7f5d405d7746447f56551a9e5f28d58e546083e2710f433f462244af4bfd64736f6c63430007060033",
}

// Attestable is an auto generated Go binding around an Ethereum contract.
type Attestable struct {
	abi abi.ABI
}

// NewAttestable creates a new instance of Attestable.
func NewAttestable() *Attestable {
	parsed, err := AttestableMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Attestable{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Attestable) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address attester) returns()
func (attestable *Attestable) PackConstructor(attester common.Address) []byte {
	enc, err := attestable.abi.Pack("", attester)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackAcceptOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x79ba5097.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function acceptOwnership() returns()
func (attestable *Attestable) PackAcceptOwnership() []byte {
	enc, err := attestable.abi.Pack("acceptOwnership")
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
func (attestable *Attestable) TryPackAcceptOwnership() ([]byte, error) {
	return attestable.abi.Pack("acceptOwnership")
}

// PackAttesterManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9b0d94b7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function attesterManager() view returns(address)
func (attestable *Attestable) PackAttesterManager() []byte {
	enc, err := attestable.abi.Pack("attesterManager")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAttesterManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9b0d94b7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function attesterManager() view returns(address)
func (attestable *Attestable) TryPackAttesterManager() ([]byte, error) {
	return attestable.abi.Pack("attesterManager")
}

// UnpackAttesterManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9b0d94b7.
//
// Solidity: function attesterManager() view returns(address)
func (attestable *Attestable) UnpackAttesterManager(data []byte) (common.Address, error) {
	out, err := attestable.abi.Unpack("attesterManager", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackDisableAttester is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2d025080.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function disableAttester(address attester) returns()
func (attestable *Attestable) PackDisableAttester(attester common.Address) []byte {
	enc, err := attestable.abi.Pack("disableAttester", attester)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDisableAttester is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2d025080.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function disableAttester(address attester) returns()
func (attestable *Attestable) TryPackDisableAttester(attester common.Address) ([]byte, error) {
	return attestable.abi.Pack("disableAttester", attester)
}

// PackEnableAttester is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfae36879.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function enableAttester(address newAttester) returns()
func (attestable *Attestable) PackEnableAttester(newAttester common.Address) []byte {
	enc, err := attestable.abi.Pack("enableAttester", newAttester)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackEnableAttester is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfae36879.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function enableAttester(address newAttester) returns()
func (attestable *Attestable) TryPackEnableAttester(newAttester common.Address) ([]byte, error) {
	return attestable.abi.Pack("enableAttester", newAttester)
}

// PackGetEnabledAttester is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbeb673d8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getEnabledAttester(uint256 index) view returns(address)
func (attestable *Attestable) PackGetEnabledAttester(index *big.Int) []byte {
	enc, err := attestable.abi.Pack("getEnabledAttester", index)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetEnabledAttester is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbeb673d8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getEnabledAttester(uint256 index) view returns(address)
func (attestable *Attestable) TryPackGetEnabledAttester(index *big.Int) ([]byte, error) {
	return attestable.abi.Pack("getEnabledAttester", index)
}

// UnpackGetEnabledAttester is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbeb673d8.
//
// Solidity: function getEnabledAttester(uint256 index) view returns(address)
func (attestable *Attestable) UnpackGetEnabledAttester(data []byte) (common.Address, error) {
	out, err := attestable.abi.Unpack("getEnabledAttester", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetNumEnabledAttesters is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x51079a53.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getNumEnabledAttesters() view returns(uint256)
func (attestable *Attestable) PackGetNumEnabledAttesters() []byte {
	enc, err := attestable.abi.Pack("getNumEnabledAttesters")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetNumEnabledAttesters is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x51079a53.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getNumEnabledAttesters() view returns(uint256)
func (attestable *Attestable) TryPackGetNumEnabledAttesters() ([]byte, error) {
	return attestable.abi.Pack("getNumEnabledAttesters")
}

// UnpackGetNumEnabledAttesters is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x51079a53.
//
// Solidity: function getNumEnabledAttesters() view returns(uint256)
func (attestable *Attestable) UnpackGetNumEnabledAttesters(data []byte) (*big.Int, error) {
	out, err := attestable.abi.Unpack("getNumEnabledAttesters", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackIsEnabledAttester is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7af82f60.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isEnabledAttester(address attester) view returns(bool)
func (attestable *Attestable) PackIsEnabledAttester(attester common.Address) []byte {
	enc, err := attestable.abi.Pack("isEnabledAttester", attester)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsEnabledAttester is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7af82f60.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isEnabledAttester(address attester) view returns(bool)
func (attestable *Attestable) TryPackIsEnabledAttester(attester common.Address) ([]byte, error) {
	return attestable.abi.Pack("isEnabledAttester", attester)
}

// UnpackIsEnabledAttester is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7af82f60.
//
// Solidity: function isEnabledAttester(address attester) view returns(bool)
func (attestable *Attestable) UnpackIsEnabledAttester(data []byte) (bool, error) {
	out, err := attestable.abi.Unpack("isEnabledAttester", data)
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
func (attestable *Attestable) PackOwner() []byte {
	enc, err := attestable.abi.Pack("owner")
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
func (attestable *Attestable) TryPackOwner() ([]byte, error) {
	return attestable.abi.Pack("owner")
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (attestable *Attestable) UnpackOwner(data []byte) (common.Address, error) {
	out, err := attestable.abi.Unpack("owner", data)
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
func (attestable *Attestable) PackPendingOwner() []byte {
	enc, err := attestable.abi.Pack("pendingOwner")
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
func (attestable *Attestable) TryPackPendingOwner() ([]byte, error) {
	return attestable.abi.Pack("pendingOwner")
}

// UnpackPendingOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (attestable *Attestable) UnpackPendingOwner(data []byte) (common.Address, error) {
	out, err := attestable.abi.Unpack("pendingOwner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackSetSignatureThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbbde5374.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setSignatureThreshold(uint256 newSignatureThreshold) returns()
func (attestable *Attestable) PackSetSignatureThreshold(newSignatureThreshold *big.Int) []byte {
	enc, err := attestable.abi.Pack("setSignatureThreshold", newSignatureThreshold)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetSignatureThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbbde5374.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setSignatureThreshold(uint256 newSignatureThreshold) returns()
func (attestable *Attestable) TryPackSetSignatureThreshold(newSignatureThreshold *big.Int) ([]byte, error) {
	return attestable.abi.Pack("setSignatureThreshold", newSignatureThreshold)
}

// PackSignatureThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa82f2e26.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function signatureThreshold() view returns(uint256)
func (attestable *Attestable) PackSignatureThreshold() []byte {
	enc, err := attestable.abi.Pack("signatureThreshold")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSignatureThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa82f2e26.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function signatureThreshold() view returns(uint256)
func (attestable *Attestable) TryPackSignatureThreshold() ([]byte, error) {
	return attestable.abi.Pack("signatureThreshold")
}

// UnpackSignatureThreshold is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa82f2e26.
//
// Solidity: function signatureThreshold() view returns(uint256)
func (attestable *Attestable) UnpackSignatureThreshold(data []byte) (*big.Int, error) {
	out, err := attestable.abi.Unpack("signatureThreshold", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (attestable *Attestable) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := attestable.abi.Pack("transferOwnership", newOwner)
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
func (attestable *Attestable) TryPackTransferOwnership(newOwner common.Address) ([]byte, error) {
	return attestable.abi.Pack("transferOwnership", newOwner)
}

// PackUpdateAttesterManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xde7769d4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updateAttesterManager(address newAttesterManager) returns()
func (attestable *Attestable) PackUpdateAttesterManager(newAttesterManager common.Address) []byte {
	enc, err := attestable.abi.Pack("updateAttesterManager", newAttesterManager)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUpdateAttesterManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xde7769d4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function updateAttesterManager(address newAttesterManager) returns()
func (attestable *Attestable) TryPackUpdateAttesterManager(newAttesterManager common.Address) ([]byte, error) {
	return attestable.abi.Pack("updateAttesterManager", newAttesterManager)
}

// AttestableAttesterDisabled represents a AttesterDisabled event raised by the Attestable contract.
type AttestableAttesterDisabled struct {
	Attester common.Address
	Raw      *types.Log // Blockchain specific contextual infos
}

const AttestableAttesterDisabledEventName = "AttesterDisabled"

// ContractEventName returns the user-defined event name.
func (AttestableAttesterDisabled) ContractEventName() string {
	return AttestableAttesterDisabledEventName
}

// UnpackAttesterDisabledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AttesterDisabled(address indexed attester)
func (attestable *Attestable) UnpackAttesterDisabledEvent(log *types.Log) (*AttestableAttesterDisabled, error) {
	event := "AttesterDisabled"
	if len(log.Topics) == 0 || log.Topics[0] != attestable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AttestableAttesterDisabled)
	if len(log.Data) > 0 {
		if err := attestable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range attestable.abi.Events[event].Inputs {
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

// AttestableAttesterEnabled represents a AttesterEnabled event raised by the Attestable contract.
type AttestableAttesterEnabled struct {
	Attester common.Address
	Raw      *types.Log // Blockchain specific contextual infos
}

const AttestableAttesterEnabledEventName = "AttesterEnabled"

// ContractEventName returns the user-defined event name.
func (AttestableAttesterEnabled) ContractEventName() string {
	return AttestableAttesterEnabledEventName
}

// UnpackAttesterEnabledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AttesterEnabled(address indexed attester)
func (attestable *Attestable) UnpackAttesterEnabledEvent(log *types.Log) (*AttestableAttesterEnabled, error) {
	event := "AttesterEnabled"
	if len(log.Topics) == 0 || log.Topics[0] != attestable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AttestableAttesterEnabled)
	if len(log.Data) > 0 {
		if err := attestable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range attestable.abi.Events[event].Inputs {
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

// AttestableAttesterManagerUpdated represents a AttesterManagerUpdated event raised by the Attestable contract.
type AttestableAttesterManagerUpdated struct {
	PreviousAttesterManager common.Address
	NewAttesterManager      common.Address
	Raw                     *types.Log // Blockchain specific contextual infos
}

const AttestableAttesterManagerUpdatedEventName = "AttesterManagerUpdated"

// ContractEventName returns the user-defined event name.
func (AttestableAttesterManagerUpdated) ContractEventName() string {
	return AttestableAttesterManagerUpdatedEventName
}

// UnpackAttesterManagerUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AttesterManagerUpdated(address indexed previousAttesterManager, address indexed newAttesterManager)
func (attestable *Attestable) UnpackAttesterManagerUpdatedEvent(log *types.Log) (*AttestableAttesterManagerUpdated, error) {
	event := "AttesterManagerUpdated"
	if len(log.Topics) == 0 || log.Topics[0] != attestable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AttestableAttesterManagerUpdated)
	if len(log.Data) > 0 {
		if err := attestable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range attestable.abi.Events[event].Inputs {
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

// AttestableOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the Attestable contract.
type AttestableOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const AttestableOwnershipTransferStartedEventName = "OwnershipTransferStarted"

// ContractEventName returns the user-defined event name.
func (AttestableOwnershipTransferStarted) ContractEventName() string {
	return AttestableOwnershipTransferStartedEventName
}

// UnpackOwnershipTransferStartedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (attestable *Attestable) UnpackOwnershipTransferStartedEvent(log *types.Log) (*AttestableOwnershipTransferStarted, error) {
	event := "OwnershipTransferStarted"
	if len(log.Topics) == 0 || log.Topics[0] != attestable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AttestableOwnershipTransferStarted)
	if len(log.Data) > 0 {
		if err := attestable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range attestable.abi.Events[event].Inputs {
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

// AttestableOwnershipTransferred represents a OwnershipTransferred event raised by the Attestable contract.
type AttestableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const AttestableOwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (AttestableOwnershipTransferred) ContractEventName() string {
	return AttestableOwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (attestable *Attestable) UnpackOwnershipTransferredEvent(log *types.Log) (*AttestableOwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if len(log.Topics) == 0 || log.Topics[0] != attestable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AttestableOwnershipTransferred)
	if len(log.Data) > 0 {
		if err := attestable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range attestable.abi.Events[event].Inputs {
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

// AttestableSignatureThresholdUpdated represents a SignatureThresholdUpdated event raised by the Attestable contract.
type AttestableSignatureThresholdUpdated struct {
	OldSignatureThreshold *big.Int
	NewSignatureThreshold *big.Int
	Raw                   *types.Log // Blockchain specific contextual infos
}

const AttestableSignatureThresholdUpdatedEventName = "SignatureThresholdUpdated"

// ContractEventName returns the user-defined event name.
func (AttestableSignatureThresholdUpdated) ContractEventName() string {
	return AttestableSignatureThresholdUpdatedEventName
}

// UnpackSignatureThresholdUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SignatureThresholdUpdated(uint256 oldSignatureThreshold, uint256 newSignatureThreshold)
func (attestable *Attestable) UnpackSignatureThresholdUpdatedEvent(log *types.Log) (*AttestableSignatureThresholdUpdated, error) {
	event := "SignatureThresholdUpdated"
	if len(log.Topics) == 0 || log.Topics[0] != attestable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AttestableSignatureThresholdUpdated)
	if len(log.Data) > 0 {
		if err := attestable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range attestable.abi.Events[event].Inputs {
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

// AttestableV2MetaData contains all meta data concerning the AttestableV2 contract.
var AttestableV2MetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"}],\"name\":\"AttesterDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"}],\"name\":\"AttesterEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousAttesterManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAttesterManager\",\"type\":\"address\"}],\"name\":\"AttesterManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldSignatureThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSignatureThreshold\",\"type\":\"uint256\"}],\"name\":\"SignatureThresholdUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"attesterManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"}],\"name\":\"disableAttester\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAttester\",\"type\":\"address\"}],\"name\":\"enableAttester\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getEnabledAttester\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumEnabledAttesters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"}],\"name\":\"isEnabledAttester\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newSignatureThreshold\",\"type\":\"uint256\"}],\"name\":\"setSignatureThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signatureThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAttesterManager\",\"type\":\"address\"}],\"name\":\"updateAttesterManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "b75cf973e73db072d5d69a9082668dddfe",
	Bin: "0x60806040523480156200001157600080fd5b503362000033620000276200006460201b60201c565b6200006c60201b60201c565b6200004433620000aa60201b60201c565b60016002819055506200005d816200021460201b60201c565b50620005fd565b600033905090565b600160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055620000a781620002ec60201b620009ff1760201c565b50565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156200014e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f496e76616c6964206174746573746572206d616e61676572206164647265737381525060200191505060405180910390fd5b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f0cee1b7ae04f3c788dd3a46c6fa677eb95b913611ef7ab59524fdc09d346021960405160405180910390a35050565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614620002d8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f43616c6c6572206e6f74206174746573746572206d616e61676572000000000081525060200191505060405180910390fd5b620002e981620003b060201b60201c565b50565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141562000454576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f4e6577206174746573746572206d757374206265206e6f6e7a65726f0000000081525060200191505060405180910390fd5b6200046f8160036200052860201b62000ac31790919060201c565b620004e2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f417474657374657220616c726561647920656e61626c6564000000000000000081525060200191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff167f5b99bab45c72ce67e89466dbc47480b9c1fde1400e7268bbf463b8354ee4653f60405160405180910390a250565b600062000558836000018373ffffffffffffffffffffffffffffffffffffffff1660001b6200056060201b60201c565b905092915050565b6000620005748383620005da60201b60201c565b620005cf578260000182908060018154018082558091505060019003906000526020600020016000909190919091505582600001805490508360010160008481526020019081526020016000208190555060019050620005d4565b600090505b92915050565b600080836001016000848152602001908152602001600020541415905092915050565b611393806200060d6000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c8063a82f2e261161008c578063de7769d411610066578063de7769d4146102a6578063e30c3978146102ea578063f2fde38b1461031e578063fae3687914610362576100cf565b8063a82f2e2614610202578063bbde537414610220578063beb673d81461024e576100cf565b80632d025080146100d457806351079a531461011857806379ba5097146101365780637af82f60146101405780638da5cb5b1461019a5780639b0d94b7146101ce575b600080fd5b610116600480360360208110156100ea57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506103a6565b005b61012061062f565b6040518082815260200191505060405180910390f35b61013e610640565b005b6101826004803603602081101561015657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506106e3565b60405180821515815260200191505060405180910390f35b6101a2610700565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101d6610729565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61020a610753565b6040518082815260200191505060405180910390f35b61024c6004803603602081101561023657600080fd5b8101908080359060200190929190505050610759565b005b61027a6004803603602081101561026457600080fd5b8101908080359060200190929190505050610828565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6102e8600480360360208110156102bc57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610845565b005b6102f2610859565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6103606004803603602081101561033457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610883565b005b6103a46004803603602081101561037857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610930565b005b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610469576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f43616c6c6572206e6f74206174746573746572206d616e61676572000000000081525060200191505060405180910390fd5b600061047361062f565b9050600181116104eb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f546f6f2066657720656e61626c6564206174746573746572730000000000000081525060200191505060405180910390fd5b6002548111610562576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f5369676e6174757265207468726573686f6c6420697320746f6f206c6f77000081525060200191505060405180910390fd5b610576826003610af390919063ffffffff16565b6105e8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f417474657374657220616c72656164792064697361626c65640000000000000081525060200191505060405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff167f78e573a18c75957b7cadaab01511aa1c19a659f06ecf53e01de37ed92d3261fc60405160405180910390a25050565b600061063b6003610b23565b905090565b600061064a610b38565b90508073ffffffffffffffffffffffffffffffffffffffff1661066b610859565b73ffffffffffffffffffffffffffffffffffffffff16146106d7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260298152602001806113356029913960400191505060405180910390fd5b6106e081610b40565b50565b60006106f9826003610b7190919063ffffffff16565b9050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60025481565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461081c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f43616c6c6572206e6f74206174746573746572206d616e61676572000000000081525060200191505060405180910390fd5b61082581610ba1565b50565b600061083e826003610d6090919063ffffffff16565b9050919050565b61084d610d7a565b61085681610e2b565b50565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61088b610d7a565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff166108eb610700565b73ffffffffffffffffffffffffffffffffffffffff167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a350565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109f3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f43616c6c6572206e6f74206174746573746572206d616e61676572000000000081525060200191505060405180910390fd5b6109fc81610f94565b50565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000610aeb836000018373ffffffffffffffffffffffffffffffffffffffff1660001b611103565b905092915050565b6000610b1b836000018373ffffffffffffffffffffffffffffffffffffffff1660001b611173565b905092915050565b6000610b318260000161125b565b9050919050565b600033905090565b600160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055610b6e816109ff565b50565b6000610b99836000018373ffffffffffffffffffffffffffffffffffffffff1660001b61126c565b905092915050565b6000811415610c18576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f496e76616c6964207369676e6174757265207468726573686f6c64000000000081525060200191505060405180910390fd5b610c226003610b23565b811115610c97576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4e6577207369676e6174757265207468726573686f6c6420746f6f206869676881525060200191505060405180910390fd5b600254811415610d0f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f5369676e6174757265207468726573686f6c6420616c7265616479207365740081525060200191505060405180910390fd5b60006002549050816002819055507f149153f58b4da003a8cfd4523709a202402182cb5aa335046911277a1be6eede8183604051808381526020018281526020019250505060405180910390a15050565b6000610d6f836000018361128f565b60001c905092915050565b610d82610b38565b73ffffffffffffffffffffffffffffffffffffffff16610da0610700565b73ffffffffffffffffffffffffffffffffffffffff1614610e29576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610ece576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f496e76616c6964206174746573746572206d616e61676572206164647265737381525060200191505060405180910390fd5b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f0cee1b7ae04f3c788dd3a46c6fa677eb95b913611ef7ab59524fdc09d346021960405160405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415611037576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f4e6577206174746573746572206d757374206265206e6f6e7a65726f0000000081525060200191505060405180910390fd5b61104b816003610ac390919063ffffffff16565b6110bd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f417474657374657220616c726561647920656e61626c6564000000000000000081525060200191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff167f5b99bab45c72ce67e89466dbc47480b9c1fde1400e7268bbf463b8354ee4653f60405160405180910390a250565b600061110f838361126c565b61116857826000018290806001815401808255809150506001900390600052602060002001600090919091909150558260000180549050836001016000848152602001908152602001600020819055506001905061116d565b600090505b92915050565b6000808360010160008481526020019081526020016000205490506000811461124f57600060018203905060006001866000018054905003905060008660000182815481106111be57fe5b90600052602060002001549050808760000184815481106111db57fe5b906000526020600020018190555060018301876001016000838152602001908152602001600020819055508660000180548061121357fe5b60019003818190600052602060002001600090559055866001016000878152602001908152602001600020600090556001945050505050611255565b60009150505b92915050565b600081600001805490509050919050565b600080836001016000848152602001908152602001600020541415905092915050565b6000818360000180549050116112f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806113136022913960400191505060405180910390fd5b8260000182815481106112ff57fe5b906000526020600020015490509291505056fe456e756d657261626c655365743a20696e646578206f7574206f6620626f756e64734f776e61626c6532537465703a2063616c6c6572206973206e6f7420746865206e6577206f776e6572a2646970667358221220e92faa421c7e6d5e309c51fcca5271a52668f109268feda2961668cfa07251b864736f6c63430007060033",
}

// AttestableV2 is an auto generated Go binding around an Ethereum contract.
type AttestableV2 struct {
	abi abi.ABI
}

// NewAttestableV2 creates a new instance of AttestableV2.
func NewAttestableV2() *AttestableV2 {
	parsed, err := AttestableV2MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &AttestableV2{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *AttestableV2) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackAcceptOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x79ba5097.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function acceptOwnership() returns()
func (attestableV2 *AttestableV2) PackAcceptOwnership() []byte {
	enc, err := attestableV2.abi.Pack("acceptOwnership")
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
func (attestableV2 *AttestableV2) TryPackAcceptOwnership() ([]byte, error) {
	return attestableV2.abi.Pack("acceptOwnership")
}

// PackAttesterManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9b0d94b7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function attesterManager() view returns(address)
func (attestableV2 *AttestableV2) PackAttesterManager() []byte {
	enc, err := attestableV2.abi.Pack("attesterManager")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAttesterManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9b0d94b7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function attesterManager() view returns(address)
func (attestableV2 *AttestableV2) TryPackAttesterManager() ([]byte, error) {
	return attestableV2.abi.Pack("attesterManager")
}

// UnpackAttesterManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9b0d94b7.
//
// Solidity: function attesterManager() view returns(address)
func (attestableV2 *AttestableV2) UnpackAttesterManager(data []byte) (common.Address, error) {
	out, err := attestableV2.abi.Unpack("attesterManager", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackDisableAttester is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2d025080.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function disableAttester(address attester) returns()
func (attestableV2 *AttestableV2) PackDisableAttester(attester common.Address) []byte {
	enc, err := attestableV2.abi.Pack("disableAttester", attester)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDisableAttester is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2d025080.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function disableAttester(address attester) returns()
func (attestableV2 *AttestableV2) TryPackDisableAttester(attester common.Address) ([]byte, error) {
	return attestableV2.abi.Pack("disableAttester", attester)
}

// PackEnableAttester is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfae36879.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function enableAttester(address newAttester) returns()
func (attestableV2 *AttestableV2) PackEnableAttester(newAttester common.Address) []byte {
	enc, err := attestableV2.abi.Pack("enableAttester", newAttester)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackEnableAttester is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfae36879.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function enableAttester(address newAttester) returns()
func (attestableV2 *AttestableV2) TryPackEnableAttester(newAttester common.Address) ([]byte, error) {
	return attestableV2.abi.Pack("enableAttester", newAttester)
}

// PackGetEnabledAttester is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbeb673d8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getEnabledAttester(uint256 index) view returns(address)
func (attestableV2 *AttestableV2) PackGetEnabledAttester(index *big.Int) []byte {
	enc, err := attestableV2.abi.Pack("getEnabledAttester", index)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetEnabledAttester is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbeb673d8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getEnabledAttester(uint256 index) view returns(address)
func (attestableV2 *AttestableV2) TryPackGetEnabledAttester(index *big.Int) ([]byte, error) {
	return attestableV2.abi.Pack("getEnabledAttester", index)
}

// UnpackGetEnabledAttester is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbeb673d8.
//
// Solidity: function getEnabledAttester(uint256 index) view returns(address)
func (attestableV2 *AttestableV2) UnpackGetEnabledAttester(data []byte) (common.Address, error) {
	out, err := attestableV2.abi.Unpack("getEnabledAttester", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetNumEnabledAttesters is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x51079a53.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getNumEnabledAttesters() view returns(uint256)
func (attestableV2 *AttestableV2) PackGetNumEnabledAttesters() []byte {
	enc, err := attestableV2.abi.Pack("getNumEnabledAttesters")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetNumEnabledAttesters is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x51079a53.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getNumEnabledAttesters() view returns(uint256)
func (attestableV2 *AttestableV2) TryPackGetNumEnabledAttesters() ([]byte, error) {
	return attestableV2.abi.Pack("getNumEnabledAttesters")
}

// UnpackGetNumEnabledAttesters is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x51079a53.
//
// Solidity: function getNumEnabledAttesters() view returns(uint256)
func (attestableV2 *AttestableV2) UnpackGetNumEnabledAttesters(data []byte) (*big.Int, error) {
	out, err := attestableV2.abi.Unpack("getNumEnabledAttesters", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackIsEnabledAttester is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7af82f60.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isEnabledAttester(address attester) view returns(bool)
func (attestableV2 *AttestableV2) PackIsEnabledAttester(attester common.Address) []byte {
	enc, err := attestableV2.abi.Pack("isEnabledAttester", attester)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsEnabledAttester is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7af82f60.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isEnabledAttester(address attester) view returns(bool)
func (attestableV2 *AttestableV2) TryPackIsEnabledAttester(attester common.Address) ([]byte, error) {
	return attestableV2.abi.Pack("isEnabledAttester", attester)
}

// UnpackIsEnabledAttester is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7af82f60.
//
// Solidity: function isEnabledAttester(address attester) view returns(bool)
func (attestableV2 *AttestableV2) UnpackIsEnabledAttester(data []byte) (bool, error) {
	out, err := attestableV2.abi.Unpack("isEnabledAttester", data)
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
func (attestableV2 *AttestableV2) PackOwner() []byte {
	enc, err := attestableV2.abi.Pack("owner")
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
func (attestableV2 *AttestableV2) TryPackOwner() ([]byte, error) {
	return attestableV2.abi.Pack("owner")
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (attestableV2 *AttestableV2) UnpackOwner(data []byte) (common.Address, error) {
	out, err := attestableV2.abi.Unpack("owner", data)
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
func (attestableV2 *AttestableV2) PackPendingOwner() []byte {
	enc, err := attestableV2.abi.Pack("pendingOwner")
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
func (attestableV2 *AttestableV2) TryPackPendingOwner() ([]byte, error) {
	return attestableV2.abi.Pack("pendingOwner")
}

// UnpackPendingOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (attestableV2 *AttestableV2) UnpackPendingOwner(data []byte) (common.Address, error) {
	out, err := attestableV2.abi.Unpack("pendingOwner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackSetSignatureThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbbde5374.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setSignatureThreshold(uint256 newSignatureThreshold) returns()
func (attestableV2 *AttestableV2) PackSetSignatureThreshold(newSignatureThreshold *big.Int) []byte {
	enc, err := attestableV2.abi.Pack("setSignatureThreshold", newSignatureThreshold)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetSignatureThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbbde5374.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setSignatureThreshold(uint256 newSignatureThreshold) returns()
func (attestableV2 *AttestableV2) TryPackSetSignatureThreshold(newSignatureThreshold *big.Int) ([]byte, error) {
	return attestableV2.abi.Pack("setSignatureThreshold", newSignatureThreshold)
}

// PackSignatureThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa82f2e26.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function signatureThreshold() view returns(uint256)
func (attestableV2 *AttestableV2) PackSignatureThreshold() []byte {
	enc, err := attestableV2.abi.Pack("signatureThreshold")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSignatureThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa82f2e26.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function signatureThreshold() view returns(uint256)
func (attestableV2 *AttestableV2) TryPackSignatureThreshold() ([]byte, error) {
	return attestableV2.abi.Pack("signatureThreshold")
}

// UnpackSignatureThreshold is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa82f2e26.
//
// Solidity: function signatureThreshold() view returns(uint256)
func (attestableV2 *AttestableV2) UnpackSignatureThreshold(data []byte) (*big.Int, error) {
	out, err := attestableV2.abi.Unpack("signatureThreshold", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (attestableV2 *AttestableV2) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := attestableV2.abi.Pack("transferOwnership", newOwner)
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
func (attestableV2 *AttestableV2) TryPackTransferOwnership(newOwner common.Address) ([]byte, error) {
	return attestableV2.abi.Pack("transferOwnership", newOwner)
}

// PackUpdateAttesterManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xde7769d4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updateAttesterManager(address newAttesterManager) returns()
func (attestableV2 *AttestableV2) PackUpdateAttesterManager(newAttesterManager common.Address) []byte {
	enc, err := attestableV2.abi.Pack("updateAttesterManager", newAttesterManager)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUpdateAttesterManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xde7769d4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function updateAttesterManager(address newAttesterManager) returns()
func (attestableV2 *AttestableV2) TryPackUpdateAttesterManager(newAttesterManager common.Address) ([]byte, error) {
	return attestableV2.abi.Pack("updateAttesterManager", newAttesterManager)
}

// AttestableV2AttesterDisabled represents a AttesterDisabled event raised by the AttestableV2 contract.
type AttestableV2AttesterDisabled struct {
	Attester common.Address
	Raw      *types.Log // Blockchain specific contextual infos
}

const AttestableV2AttesterDisabledEventName = "AttesterDisabled"

// ContractEventName returns the user-defined event name.
func (AttestableV2AttesterDisabled) ContractEventName() string {
	return AttestableV2AttesterDisabledEventName
}

// UnpackAttesterDisabledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AttesterDisabled(address indexed attester)
func (attestableV2 *AttestableV2) UnpackAttesterDisabledEvent(log *types.Log) (*AttestableV2AttesterDisabled, error) {
	event := "AttesterDisabled"
	if len(log.Topics) == 0 || log.Topics[0] != attestableV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AttestableV2AttesterDisabled)
	if len(log.Data) > 0 {
		if err := attestableV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range attestableV2.abi.Events[event].Inputs {
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

// AttestableV2AttesterEnabled represents a AttesterEnabled event raised by the AttestableV2 contract.
type AttestableV2AttesterEnabled struct {
	Attester common.Address
	Raw      *types.Log // Blockchain specific contextual infos
}

const AttestableV2AttesterEnabledEventName = "AttesterEnabled"

// ContractEventName returns the user-defined event name.
func (AttestableV2AttesterEnabled) ContractEventName() string {
	return AttestableV2AttesterEnabledEventName
}

// UnpackAttesterEnabledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AttesterEnabled(address indexed attester)
func (attestableV2 *AttestableV2) UnpackAttesterEnabledEvent(log *types.Log) (*AttestableV2AttesterEnabled, error) {
	event := "AttesterEnabled"
	if len(log.Topics) == 0 || log.Topics[0] != attestableV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AttestableV2AttesterEnabled)
	if len(log.Data) > 0 {
		if err := attestableV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range attestableV2.abi.Events[event].Inputs {
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

// AttestableV2AttesterManagerUpdated represents a AttesterManagerUpdated event raised by the AttestableV2 contract.
type AttestableV2AttesterManagerUpdated struct {
	PreviousAttesterManager common.Address
	NewAttesterManager      common.Address
	Raw                     *types.Log // Blockchain specific contextual infos
}

const AttestableV2AttesterManagerUpdatedEventName = "AttesterManagerUpdated"

// ContractEventName returns the user-defined event name.
func (AttestableV2AttesterManagerUpdated) ContractEventName() string {
	return AttestableV2AttesterManagerUpdatedEventName
}

// UnpackAttesterManagerUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AttesterManagerUpdated(address indexed previousAttesterManager, address indexed newAttesterManager)
func (attestableV2 *AttestableV2) UnpackAttesterManagerUpdatedEvent(log *types.Log) (*AttestableV2AttesterManagerUpdated, error) {
	event := "AttesterManagerUpdated"
	if len(log.Topics) == 0 || log.Topics[0] != attestableV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AttestableV2AttesterManagerUpdated)
	if len(log.Data) > 0 {
		if err := attestableV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range attestableV2.abi.Events[event].Inputs {
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

// AttestableV2OwnershipTransferStarted represents a OwnershipTransferStarted event raised by the AttestableV2 contract.
type AttestableV2OwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const AttestableV2OwnershipTransferStartedEventName = "OwnershipTransferStarted"

// ContractEventName returns the user-defined event name.
func (AttestableV2OwnershipTransferStarted) ContractEventName() string {
	return AttestableV2OwnershipTransferStartedEventName
}

// UnpackOwnershipTransferStartedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (attestableV2 *AttestableV2) UnpackOwnershipTransferStartedEvent(log *types.Log) (*AttestableV2OwnershipTransferStarted, error) {
	event := "OwnershipTransferStarted"
	if len(log.Topics) == 0 || log.Topics[0] != attestableV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AttestableV2OwnershipTransferStarted)
	if len(log.Data) > 0 {
		if err := attestableV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range attestableV2.abi.Events[event].Inputs {
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

// AttestableV2OwnershipTransferred represents a OwnershipTransferred event raised by the AttestableV2 contract.
type AttestableV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const AttestableV2OwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (AttestableV2OwnershipTransferred) ContractEventName() string {
	return AttestableV2OwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (attestableV2 *AttestableV2) UnpackOwnershipTransferredEvent(log *types.Log) (*AttestableV2OwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if len(log.Topics) == 0 || log.Topics[0] != attestableV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AttestableV2OwnershipTransferred)
	if len(log.Data) > 0 {
		if err := attestableV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range attestableV2.abi.Events[event].Inputs {
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

// AttestableV2SignatureThresholdUpdated represents a SignatureThresholdUpdated event raised by the AttestableV2 contract.
type AttestableV2SignatureThresholdUpdated struct {
	OldSignatureThreshold *big.Int
	NewSignatureThreshold *big.Int
	Raw                   *types.Log // Blockchain specific contextual infos
}

const AttestableV2SignatureThresholdUpdatedEventName = "SignatureThresholdUpdated"

// ContractEventName returns the user-defined event name.
func (AttestableV2SignatureThresholdUpdated) ContractEventName() string {
	return AttestableV2SignatureThresholdUpdatedEventName
}

// UnpackSignatureThresholdUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SignatureThresholdUpdated(uint256 oldSignatureThreshold, uint256 newSignatureThreshold)
func (attestableV2 *AttestableV2) UnpackSignatureThresholdUpdatedEvent(log *types.Log) (*AttestableV2SignatureThresholdUpdated, error) {
	event := "SignatureThresholdUpdated"
	if len(log.Topics) == 0 || log.Topics[0] != attestableV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AttestableV2SignatureThresholdUpdated)
	if len(log.Data) > 0 {
		if err := attestableV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range attestableV2.abi.Events[event].Inputs {
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

// BaseMessageTransmitterMetaData contains all meta data concerning the BaseMessageTransmitter contract.
var BaseMessageTransmitterMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_localDomain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_version\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"}],\"name\":\"AttesterDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"}],\"name\":\"AttesterEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousAttesterManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAttesterManager\",\"type\":\"address\"}],\"name\":\"AttesterManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxMessageBodySize\",\"type\":\"uint256\"}],\"name\":\"MaxMessageBodySizeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"PauserChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newRescuer\",\"type\":\"address\"}],\"name\":\"RescuerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldSignatureThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSignatureThreshold\",\"type\":\"uint256\"}],\"name\":\"SignatureThresholdUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NONCE_USED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"attesterManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"}],\"name\":\"disableAttester\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAttester\",\"type\":\"address\"}],\"name\":\"enableAttester\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getEnabledAttester\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumEnabledAttesters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedVersion\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"}],\"name\":\"isEnabledAttester\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxMessageBodySize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rescuer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMaxMessageBodySize\",\"type\":\"uint256\"}],\"name\":\"setMaxMessageBodySize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newSignatureThreshold\",\"type\":\"uint256\"}],\"name\":\"setSignatureThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signatureThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAttesterManager\",\"type\":\"address\"}],\"name\":\"updateAttesterManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newPauser\",\"type\":\"address\"}],\"name\":\"updatePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRescuer\",\"type\":\"address\"}],\"name\":\"updateRescuer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"usedNonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "d5613de6bff063d30fe82f88eaa6d98222",
	Bin: "0x60c06040526000600260146101000a81548160ff0219169083151502179055503480156200002c57600080fd5b506040516200299f3803806200299f833981810160405260408110156200005257600080fd5b810190808051906020019092919080519060200190929190505050336200008e62000082620000ef60201b60201c565b620000f760201b60201c565b6200009f336200013560201b60201c565b6001600481905550620000b8816200029f60201b60201c565b508163ffffffff1660808163ffffffff1660e01b815250508063ffffffff1660a08163ffffffff1660e01b81525050505062000688565b600033905090565b600160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905562000132816200037760201b620011731760201c565b50565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415620001d9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f496e76616c6964206174746573746572206d616e61676572206164647265737381525060200191505060405180910390fd5b6000600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f0cee1b7ae04f3c788dd3a46c6fa677eb95b913611ef7ab59524fdc09d346021960405160405180910390a35050565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161462000363576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f43616c6c6572206e6f74206174746573746572206d616e61676572000000000081525060200191505060405180910390fd5b62000374816200043b60201b60201c565b50565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415620004df576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f4e6577206174746573746572206d757374206265206e6f6e7a65726f0000000081525060200191505060405180910390fd5b620004fa816005620005b360201b620012371790919060201c565b6200056d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f417474657374657220616c726561647920656e61626c6564000000000000000081525060200191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff167f5b99bab45c72ce67e89466dbc47480b9c1fde1400e7268bbf463b8354ee4653f60405160405180910390a250565b6000620005e3836000018373ffffffffffffffffffffffffffffffffffffffff1660001b620005eb60201b60201c565b905092915050565b6000620005ff83836200066560201b60201c565b6200065a5782600001829080600181540180825580915050600190039060005260206000200160009091909190915055826000018054905083600101600084815260200190815260200160002081905550600190506200065f565b600090505b92915050565b600080836001016000848152602001908152602001600020541415905092915050565b60805160e01c60a05160e01c6122ed620006b260003980610b21525080610d2052506122ed6000f3fe608060405234801561001057600080fd5b50600436106101c45760003560e01c80638da5cb5b116100f9578063bbde537411610097578063e30c397811610071578063e30c39781461064b578063f2fde38b1461067f578063fae36879146106c3578063feb6172414610707576101c4565b8063bbde537414610581578063beb673d8146105af578063de7769d414610607576101c4565b80639fd0506d116100d35780639fd0506d146104a3578063a82f2e26146104d7578063af47b9bb146104f5578063b2118a8d14610513576101c4565b80638da5cb5b1461040d57806392492c68146104415780639b0d94b71461046f576101c4565b8063554bab3c116101665780637af82f60116101405780637af82f60146103675780637de25ae4146103c15780638456cb59146103df5780638d3638f4146103e9576101c4565b8063554bab3c146102f95780635c975abb1461033d57806379ba50971461035d576101c4565b806338a63183116101a257806338a63183146102795780633f4ba83a146102ad57806351079a53146102b757806354fd4d50146102d5576101c4565b806308828eb7146101c95780632ab60045146101f15780632d02508014610235575b600080fd5b6101d1610749565b604051808267ffffffffffffffff16815260200191505060405180910390f35b6102336004803603602081101561020757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610758565b005b6102776004803603602081101561024b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061076c565b005b6102816109f5565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6102b5610a1f565b005b6102bf610b0e565b6040518082815260200191505060405180910390f35b6102dd610b1f565b604051808263ffffffff16815260200191505060405180910390f35b61033b6004803603602081101561030f57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610b43565b005b610345610b57565b60405180821515815260200191505060405180910390f35b610365610b6a565b005b6103a96004803603602081101561037d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610c0d565b60405180821515815260200191505060405180910390f35b6103c9610c2a565b6040518082815260200191505060405180910390f35b6103e7610c2f565b005b6103f1610d1e565b604051808263ffffffff16815260200191505060405180910390f35b610415610d42565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61046d6004803603602081101561045757600080fd5b8101908080359060200190929190505050610d6b565b005b610477610d7f565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6104ab610da9565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6104df610dd3565b6040518082815260200191505060405180910390f35b6104fd610dd9565b6040518082815260200191505060405180910390f35b61057f6004803603606081101561052957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610ddf565b005b6105ad6004803603602081101561059757600080fd5b8101908080359060200190929190505050610eb5565b005b6105db600480360360208110156105c557600080fd5b8101908080359060200190929190505050610f84565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6106496004803603602081101561061d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610fa1565b005b610653610fb5565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6106c16004803603602081101561069557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610fdf565b005b610705600480360360208110156106d957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061108c565b005b6107336004803603602081101561071d57600080fd5b810190808035906020019092919050505061115b565b6040518082815260200191505060405180910390f35b6000610753611267565b905090565b61076061128e565b6107698161133f565b50565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461082f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f43616c6c6572206e6f74206174746573746572206d616e61676572000000000081525060200191505060405180910390fd5b6000610839610b0e565b9050600181116108b1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f546f6f2066657720656e61626c6564206174746573746572730000000000000081525060200191505060405180910390fd5b6004548111610928576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f5369676e6174757265207468726573686f6c6420697320746f6f206c6f77000081525060200191505060405180910390fd5b61093c82600561144c90919063ffffffff16565b6109ae576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f417474657374657220616c72656164792064697361626c65640000000000000081525060200191505060405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff167f78e573a18c75957b7cadaab01511aa1c19a659f06ecf53e01de37ed92d3261fc60405160405180910390a25050565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ac5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602281526020018061226c6022913960400191505060405180910390fd5b6000600260146101000a81548160ff0219169083151502179055507f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a1565b6000610b1a600561147c565b905090565b7f000000000000000000000000000000000000000000000000000000000000000081565b610b4b61128e565b610b5481611491565b50565b600260149054906101000a900460ff1681565b6000610b746115c0565b90508073ffffffffffffffffffffffffffffffffffffffff16610b95610fb5565b73ffffffffffffffffffffffffffffffffffffffff1614610c01576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260298152602001806121cf6029913960400191505060405180910390fd5b610c0a816115c8565b50565b6000610c238260056115f990919063ffffffff16565b9050919050565b600181565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610cd5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602281526020018061226c6022913960400191505060405180910390fd5b6001600260146101000a81548160ff0219169083151502179055507f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b7f000000000000000000000000000000000000000000000000000000000000000081565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610d7361128e565b610d7c81611629565b50565b6000600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60045481565b601c5481565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610e85576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806122486024913960400191505060405180910390fd5b610eb082828573ffffffffffffffffffffffffffffffffffffffff1661166c9092919063ffffffff16565b505050565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610f78576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f43616c6c6572206e6f74206174746573746572206d616e61676572000000000081525060200191505060405180910390fd5b610f818161170e565b50565b6000610f9a8260056118cd90919063ffffffff16565b9050919050565b610fa961128e565b610fb2816118e7565b50565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610fe761128e565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16611047610d42565b73ffffffffffffffffffffffffffffffffffffffff167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a350565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461114f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f43616c6c6572206e6f74206174746573746572206d616e61676572000000000081525060200191505060405180910390fd5b61115881611a50565b50565b601d6020528060005260406000206000915090505481565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600061125f836000018373ffffffffffffffffffffffffffffffffffffffff1660001b611bbf565b905092915050565b6000611271611c2f565b60000160009054906101000a900467ffffffffffffffff16905090565b6112966115c0565b73ffffffffffffffffffffffffffffffffffffffff166112b4610d42565b73ffffffffffffffffffffffffffffffffffffffff161461133d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156113c5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a8152602001806121f8602a913960400191505060405180910390fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff167fe475e580d85111348e40d8ca33cfdd74c30fe1655c2d8537a13abc10065ffa5a60405160405180910390a250565b6000611474836000018373ffffffffffffffffffffffffffffffffffffffff1660001b611c57565b905092915050565b600061148a82600001611d3f565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415611517576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260288152602001806121a76028913960400191505060405180910390fd5b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fb80482a293ca2e013eda8683c9bd7fc8347cfdaeea5ede58cba46df502c2a60460405160405180910390a250565b600033905090565b600160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556115f681611173565b50565b6000611621836000018373ffffffffffffffffffffffffffffffffffffffff1660001b611d50565b905092915050565b80601c819055507fb13bf6bebed03d1b318e3ea32e4b2a3ad9f5e2312cdf340a2f4bbfaee39f928d601c546040518082815260200191505060405180910390a150565b6117098363a9059cbb60e01b8484604051602401808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611d73565b505050565b6000811415611785576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f496e76616c6964207369676e6174757265207468726573686f6c64000000000081525060200191505060405180910390fd5b61178f600561147c565b811115611804576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4e6577207369676e6174757265207468726573686f6c6420746f6f206869676881525060200191505060405180910390fd5b60045481141561187c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f5369676e6174757265207468726573686f6c6420616c7265616479207365740081525060200191505060405180910390fd5b60006004549050816004819055507f149153f58b4da003a8cfd4523709a202402182cb5aa335046911277a1be6eede8183604051808381526020018281526020019250505060405180910390a15050565b60006118dc8360000183611e62565b60001c905092915050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561198a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f496e76616c6964206174746573746572206d616e61676572206164647265737381525060200191505060405180910390fd5b6000600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f0cee1b7ae04f3c788dd3a46c6fa677eb95b913611ef7ab59524fdc09d346021960405160405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415611af3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f4e6577206174746573746572206d757374206265206e6f6e7a65726f0000000081525060200191505060405180910390fd5b611b0781600561123790919063ffffffff16565b611b79576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f417474657374657220616c726561647920656e61626c6564000000000000000081525060200191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff167f5b99bab45c72ce67e89466dbc47480b9c1fde1400e7268bbf463b8354ee4653f60405160405180910390a250565b6000611bcb8383611d50565b611c24578260000182908060018154018082558091505060019003906000526020600020016000909190919091505582600001805490508360010160008481526020019081526020016000208190555060019050611c29565b600090505b92915050565b60007ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b60008083600101600084815260200190815260200160002054905060008114611d335760006001820390506000600186600001805490500390506000866000018281548110611ca257fe5b9060005260206000200154905080876000018481548110611cbf57fe5b9060005260206000200181905550600183018760010160008381526020019081526020016000208190555086600001805480611cf757fe5b60019003818190600052602060002001600090559055866001016000878152602001908152602001600020600090556001945050505050611d39565b60009150505b92915050565b600081600001805490509050919050565b600080836001016000848152602001908152602001600020541415905092915050565b6000611dd5826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16611ee59092919063ffffffff16565b9050600081511115611e5d57808060200190516020811015611df657600080fd5b8101908080519060200190929190505050611e5c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a81526020018061228e602a913960400191505060405180910390fd5b5b505050565b600081836000018054905011611ec3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806121856022913960400191505060405180910390fd5b826000018281548110611ed257fe5b9060005260206000200154905092915050565b6060611ef48484600085611efd565b90509392505050565b606082471015611f58576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260268152602001806122226026913960400191505060405180910390fd5b611f61856120a5565b611fd3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000081525060200191505060405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040518082805190602001908083835b602083106120225780518252602082019150602081019050602083039250611fff565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114612084576040519150601f19603f3d011682016040523d82523d6000602084013e612089565b606091505b50915091506120998282866120b8565b92505050949350505050565b600080823b905060008111915050919050565b606083156120c85782905061217d565b6000835111156120db5782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015612142578082015181840152602081019050612127565b50505050905090810190601f16801561216f5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b939250505056fe456e756d657261626c655365743a20696e646578206f7574206f6620626f756e64735061757361626c653a206e65772070617573657220697320746865207a65726f20616464726573734f776e61626c6532537465703a2063616c6c6572206973206e6f7420746865206e6577206f776e6572526573637561626c653a206e6577207265736375657220697320746865207a65726f2061646472657373416464726573733a20696e73756666696369656e742062616c616e636520666f722063616c6c526573637561626c653a2063616c6c6572206973206e6f742074686520726573637565725061757361626c653a2063616c6c6572206973206e6f7420746865207061757365725361666545524332303a204552433230206f7065726174696f6e20646964206e6f742073756363656564a2646970667358221220742d0bc974408ab6c36fd1adaeac5a743859995a0673207a73fff25c64904c4964736f6c63430007060033",
}

// BaseMessageTransmitter is an auto generated Go binding around an Ethereum contract.
type BaseMessageTransmitter struct {
	abi abi.ABI
}

// NewBaseMessageTransmitter creates a new instance of BaseMessageTransmitter.
func NewBaseMessageTransmitter() *BaseMessageTransmitter {
	parsed, err := BaseMessageTransmitterMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &BaseMessageTransmitter{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *BaseMessageTransmitter) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(uint32 _localDomain, uint32 _version) returns()
func (baseMessageTransmitter *BaseMessageTransmitter) PackConstructor(_localDomain uint32, _version uint32) []byte {
	enc, err := baseMessageTransmitter.abi.Pack("", _localDomain, _version)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackNONCEUSED is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7de25ae4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function NONCE_USED() view returns(uint256)
func (baseMessageTransmitter *BaseMessageTransmitter) PackNONCEUSED() []byte {
	enc, err := baseMessageTransmitter.abi.Pack("NONCE_USED")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackNONCEUSED is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7de25ae4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function NONCE_USED() view returns(uint256)
func (baseMessageTransmitter *BaseMessageTransmitter) TryPackNONCEUSED() ([]byte, error) {
	return baseMessageTransmitter.abi.Pack("NONCE_USED")
}

// UnpackNONCEUSED is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7de25ae4.
//
// Solidity: function NONCE_USED() view returns(uint256)
func (baseMessageTransmitter *BaseMessageTransmitter) UnpackNONCEUSED(data []byte) (*big.Int, error) {
	out, err := baseMessageTransmitter.abi.Unpack("NONCE_USED", data)
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
func (baseMessageTransmitter *BaseMessageTransmitter) PackAcceptOwnership() []byte {
	enc, err := baseMessageTransmitter.abi.Pack("acceptOwnership")
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
func (baseMessageTransmitter *BaseMessageTransmitter) TryPackAcceptOwnership() ([]byte, error) {
	return baseMessageTransmitter.abi.Pack("acceptOwnership")
}

// PackAttesterManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9b0d94b7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function attesterManager() view returns(address)
func (baseMessageTransmitter *BaseMessageTransmitter) PackAttesterManager() []byte {
	enc, err := baseMessageTransmitter.abi.Pack("attesterManager")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAttesterManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9b0d94b7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function attesterManager() view returns(address)
func (baseMessageTransmitter *BaseMessageTransmitter) TryPackAttesterManager() ([]byte, error) {
	return baseMessageTransmitter.abi.Pack("attesterManager")
}

// UnpackAttesterManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9b0d94b7.
//
// Solidity: function attesterManager() view returns(address)
func (baseMessageTransmitter *BaseMessageTransmitter) UnpackAttesterManager(data []byte) (common.Address, error) {
	out, err := baseMessageTransmitter.abi.Unpack("attesterManager", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackDisableAttester is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2d025080.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function disableAttester(address attester) returns()
func (baseMessageTransmitter *BaseMessageTransmitter) PackDisableAttester(attester common.Address) []byte {
	enc, err := baseMessageTransmitter.abi.Pack("disableAttester", attester)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDisableAttester is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2d025080.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function disableAttester(address attester) returns()
func (baseMessageTransmitter *BaseMessageTransmitter) TryPackDisableAttester(attester common.Address) ([]byte, error) {
	return baseMessageTransmitter.abi.Pack("disableAttester", attester)
}

// PackEnableAttester is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfae36879.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function enableAttester(address newAttester) returns()
func (baseMessageTransmitter *BaseMessageTransmitter) PackEnableAttester(newAttester common.Address) []byte {
	enc, err := baseMessageTransmitter.abi.Pack("enableAttester", newAttester)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackEnableAttester is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfae36879.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function enableAttester(address newAttester) returns()
func (baseMessageTransmitter *BaseMessageTransmitter) TryPackEnableAttester(newAttester common.Address) ([]byte, error) {
	return baseMessageTransmitter.abi.Pack("enableAttester", newAttester)
}

// PackGetEnabledAttester is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbeb673d8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getEnabledAttester(uint256 index) view returns(address)
func (baseMessageTransmitter *BaseMessageTransmitter) PackGetEnabledAttester(index *big.Int) []byte {
	enc, err := baseMessageTransmitter.abi.Pack("getEnabledAttester", index)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetEnabledAttester is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbeb673d8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getEnabledAttester(uint256 index) view returns(address)
func (baseMessageTransmitter *BaseMessageTransmitter) TryPackGetEnabledAttester(index *big.Int) ([]byte, error) {
	return baseMessageTransmitter.abi.Pack("getEnabledAttester", index)
}

// UnpackGetEnabledAttester is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbeb673d8.
//
// Solidity: function getEnabledAttester(uint256 index) view returns(address)
func (baseMessageTransmitter *BaseMessageTransmitter) UnpackGetEnabledAttester(data []byte) (common.Address, error) {
	out, err := baseMessageTransmitter.abi.Unpack("getEnabledAttester", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetNumEnabledAttesters is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x51079a53.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getNumEnabledAttesters() view returns(uint256)
func (baseMessageTransmitter *BaseMessageTransmitter) PackGetNumEnabledAttesters() []byte {
	enc, err := baseMessageTransmitter.abi.Pack("getNumEnabledAttesters")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetNumEnabledAttesters is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x51079a53.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getNumEnabledAttesters() view returns(uint256)
func (baseMessageTransmitter *BaseMessageTransmitter) TryPackGetNumEnabledAttesters() ([]byte, error) {
	return baseMessageTransmitter.abi.Pack("getNumEnabledAttesters")
}

// UnpackGetNumEnabledAttesters is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x51079a53.
//
// Solidity: function getNumEnabledAttesters() view returns(uint256)
func (baseMessageTransmitter *BaseMessageTransmitter) UnpackGetNumEnabledAttesters(data []byte) (*big.Int, error) {
	out, err := baseMessageTransmitter.abi.Unpack("getNumEnabledAttesters", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackInitializedVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x08828eb7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function initializedVersion() view returns(uint64)
func (baseMessageTransmitter *BaseMessageTransmitter) PackInitializedVersion() []byte {
	enc, err := baseMessageTransmitter.abi.Pack("initializedVersion")
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
func (baseMessageTransmitter *BaseMessageTransmitter) TryPackInitializedVersion() ([]byte, error) {
	return baseMessageTransmitter.abi.Pack("initializedVersion")
}

// UnpackInitializedVersion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x08828eb7.
//
// Solidity: function initializedVersion() view returns(uint64)
func (baseMessageTransmitter *BaseMessageTransmitter) UnpackInitializedVersion(data []byte) (uint64, error) {
	out, err := baseMessageTransmitter.abi.Unpack("initializedVersion", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackIsEnabledAttester is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7af82f60.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isEnabledAttester(address attester) view returns(bool)
func (baseMessageTransmitter *BaseMessageTransmitter) PackIsEnabledAttester(attester common.Address) []byte {
	enc, err := baseMessageTransmitter.abi.Pack("isEnabledAttester", attester)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsEnabledAttester is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7af82f60.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isEnabledAttester(address attester) view returns(bool)
func (baseMessageTransmitter *BaseMessageTransmitter) TryPackIsEnabledAttester(attester common.Address) ([]byte, error) {
	return baseMessageTransmitter.abi.Pack("isEnabledAttester", attester)
}

// UnpackIsEnabledAttester is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7af82f60.
//
// Solidity: function isEnabledAttester(address attester) view returns(bool)
func (baseMessageTransmitter *BaseMessageTransmitter) UnpackIsEnabledAttester(data []byte) (bool, error) {
	out, err := baseMessageTransmitter.abi.Unpack("isEnabledAttester", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackLocalDomain is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8d3638f4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function localDomain() view returns(uint32)
func (baseMessageTransmitter *BaseMessageTransmitter) PackLocalDomain() []byte {
	enc, err := baseMessageTransmitter.abi.Pack("localDomain")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackLocalDomain is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8d3638f4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function localDomain() view returns(uint32)
func (baseMessageTransmitter *BaseMessageTransmitter) TryPackLocalDomain() ([]byte, error) {
	return baseMessageTransmitter.abi.Pack("localDomain")
}

// UnpackLocalDomain is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (baseMessageTransmitter *BaseMessageTransmitter) UnpackLocalDomain(data []byte) (uint32, error) {
	out, err := baseMessageTransmitter.abi.Unpack("localDomain", data)
	if err != nil {
		return *new(uint32), err
	}
	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	return out0, nil
}

// PackMaxMessageBodySize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaf47b9bb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function maxMessageBodySize() view returns(uint256)
func (baseMessageTransmitter *BaseMessageTransmitter) PackMaxMessageBodySize() []byte {
	enc, err := baseMessageTransmitter.abi.Pack("maxMessageBodySize")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMaxMessageBodySize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaf47b9bb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function maxMessageBodySize() view returns(uint256)
func (baseMessageTransmitter *BaseMessageTransmitter) TryPackMaxMessageBodySize() ([]byte, error) {
	return baseMessageTransmitter.abi.Pack("maxMessageBodySize")
}

// UnpackMaxMessageBodySize is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xaf47b9bb.
//
// Solidity: function maxMessageBodySize() view returns(uint256)
func (baseMessageTransmitter *BaseMessageTransmitter) UnpackMaxMessageBodySize(data []byte) (*big.Int, error) {
	out, err := baseMessageTransmitter.abi.Unpack("maxMessageBodySize", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function owner() view returns(address)
func (baseMessageTransmitter *BaseMessageTransmitter) PackOwner() []byte {
	enc, err := baseMessageTransmitter.abi.Pack("owner")
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
func (baseMessageTransmitter *BaseMessageTransmitter) TryPackOwner() ([]byte, error) {
	return baseMessageTransmitter.abi.Pack("owner")
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (baseMessageTransmitter *BaseMessageTransmitter) UnpackOwner(data []byte) (common.Address, error) {
	out, err := baseMessageTransmitter.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackPause is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8456cb59.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function pause() returns()
func (baseMessageTransmitter *BaseMessageTransmitter) PackPause() []byte {
	enc, err := baseMessageTransmitter.abi.Pack("pause")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPause is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8456cb59.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function pause() returns()
func (baseMessageTransmitter *BaseMessageTransmitter) TryPackPause() ([]byte, error) {
	return baseMessageTransmitter.abi.Pack("pause")
}

// PackPaused is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c975abb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function paused() view returns(bool)
func (baseMessageTransmitter *BaseMessageTransmitter) PackPaused() []byte {
	enc, err := baseMessageTransmitter.abi.Pack("paused")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPaused is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c975abb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function paused() view returns(bool)
func (baseMessageTransmitter *BaseMessageTransmitter) TryPackPaused() ([]byte, error) {
	return baseMessageTransmitter.abi.Pack("paused")
}

// UnpackPaused is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (baseMessageTransmitter *BaseMessageTransmitter) UnpackPaused(data []byte) (bool, error) {
	out, err := baseMessageTransmitter.abi.Unpack("paused", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackPauser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9fd0506d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function pauser() view returns(address)
func (baseMessageTransmitter *BaseMessageTransmitter) PackPauser() []byte {
	enc, err := baseMessageTransmitter.abi.Pack("pauser")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPauser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9fd0506d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function pauser() view returns(address)
func (baseMessageTransmitter *BaseMessageTransmitter) TryPackPauser() ([]byte, error) {
	return baseMessageTransmitter.abi.Pack("pauser")
}

// UnpackPauser is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (baseMessageTransmitter *BaseMessageTransmitter) UnpackPauser(data []byte) (common.Address, error) {
	out, err := baseMessageTransmitter.abi.Unpack("pauser", data)
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
func (baseMessageTransmitter *BaseMessageTransmitter) PackPendingOwner() []byte {
	enc, err := baseMessageTransmitter.abi.Pack("pendingOwner")
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
func (baseMessageTransmitter *BaseMessageTransmitter) TryPackPendingOwner() ([]byte, error) {
	return baseMessageTransmitter.abi.Pack("pendingOwner")
}

// UnpackPendingOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (baseMessageTransmitter *BaseMessageTransmitter) UnpackPendingOwner(data []byte) (common.Address, error) {
	out, err := baseMessageTransmitter.abi.Unpack("pendingOwner", data)
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
func (baseMessageTransmitter *BaseMessageTransmitter) PackRescueERC20(tokenContract common.Address, to common.Address, amount *big.Int) []byte {
	enc, err := baseMessageTransmitter.abi.Pack("rescueERC20", tokenContract, to, amount)
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
func (baseMessageTransmitter *BaseMessageTransmitter) TryPackRescueERC20(tokenContract common.Address, to common.Address, amount *big.Int) ([]byte, error) {
	return baseMessageTransmitter.abi.Pack("rescueERC20", tokenContract, to, amount)
}

// PackRescuer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x38a63183.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function rescuer() view returns(address)
func (baseMessageTransmitter *BaseMessageTransmitter) PackRescuer() []byte {
	enc, err := baseMessageTransmitter.abi.Pack("rescuer")
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
func (baseMessageTransmitter *BaseMessageTransmitter) TryPackRescuer() ([]byte, error) {
	return baseMessageTransmitter.abi.Pack("rescuer")
}

// UnpackRescuer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x38a63183.
//
// Solidity: function rescuer() view returns(address)
func (baseMessageTransmitter *BaseMessageTransmitter) UnpackRescuer(data []byte) (common.Address, error) {
	out, err := baseMessageTransmitter.abi.Unpack("rescuer", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackSetMaxMessageBodySize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x92492c68.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setMaxMessageBodySize(uint256 newMaxMessageBodySize) returns()
func (baseMessageTransmitter *BaseMessageTransmitter) PackSetMaxMessageBodySize(newMaxMessageBodySize *big.Int) []byte {
	enc, err := baseMessageTransmitter.abi.Pack("setMaxMessageBodySize", newMaxMessageBodySize)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetMaxMessageBodySize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x92492c68.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setMaxMessageBodySize(uint256 newMaxMessageBodySize) returns()
func (baseMessageTransmitter *BaseMessageTransmitter) TryPackSetMaxMessageBodySize(newMaxMessageBodySize *big.Int) ([]byte, error) {
	return baseMessageTransmitter.abi.Pack("setMaxMessageBodySize", newMaxMessageBodySize)
}

// PackSetSignatureThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbbde5374.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setSignatureThreshold(uint256 newSignatureThreshold) returns()
func (baseMessageTransmitter *BaseMessageTransmitter) PackSetSignatureThreshold(newSignatureThreshold *big.Int) []byte {
	enc, err := baseMessageTransmitter.abi.Pack("setSignatureThreshold", newSignatureThreshold)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetSignatureThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbbde5374.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setSignatureThreshold(uint256 newSignatureThreshold) returns()
func (baseMessageTransmitter *BaseMessageTransmitter) TryPackSetSignatureThreshold(newSignatureThreshold *big.Int) ([]byte, error) {
	return baseMessageTransmitter.abi.Pack("setSignatureThreshold", newSignatureThreshold)
}

// PackSignatureThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa82f2e26.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function signatureThreshold() view returns(uint256)
func (baseMessageTransmitter *BaseMessageTransmitter) PackSignatureThreshold() []byte {
	enc, err := baseMessageTransmitter.abi.Pack("signatureThreshold")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSignatureThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa82f2e26.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function signatureThreshold() view returns(uint256)
func (baseMessageTransmitter *BaseMessageTransmitter) TryPackSignatureThreshold() ([]byte, error) {
	return baseMessageTransmitter.abi.Pack("signatureThreshold")
}

// UnpackSignatureThreshold is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa82f2e26.
//
// Solidity: function signatureThreshold() view returns(uint256)
func (baseMessageTransmitter *BaseMessageTransmitter) UnpackSignatureThreshold(data []byte) (*big.Int, error) {
	out, err := baseMessageTransmitter.abi.Unpack("signatureThreshold", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (baseMessageTransmitter *BaseMessageTransmitter) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := baseMessageTransmitter.abi.Pack("transferOwnership", newOwner)
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
func (baseMessageTransmitter *BaseMessageTransmitter) TryPackTransferOwnership(newOwner common.Address) ([]byte, error) {
	return baseMessageTransmitter.abi.Pack("transferOwnership", newOwner)
}

// PackUnpause is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3f4ba83a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function unpause() returns()
func (baseMessageTransmitter *BaseMessageTransmitter) PackUnpause() []byte {
	enc, err := baseMessageTransmitter.abi.Pack("unpause")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUnpause is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3f4ba83a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function unpause() returns()
func (baseMessageTransmitter *BaseMessageTransmitter) TryPackUnpause() ([]byte, error) {
	return baseMessageTransmitter.abi.Pack("unpause")
}

// PackUpdateAttesterManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xde7769d4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updateAttesterManager(address newAttesterManager) returns()
func (baseMessageTransmitter *BaseMessageTransmitter) PackUpdateAttesterManager(newAttesterManager common.Address) []byte {
	enc, err := baseMessageTransmitter.abi.Pack("updateAttesterManager", newAttesterManager)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUpdateAttesterManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xde7769d4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function updateAttesterManager(address newAttesterManager) returns()
func (baseMessageTransmitter *BaseMessageTransmitter) TryPackUpdateAttesterManager(newAttesterManager common.Address) ([]byte, error) {
	return baseMessageTransmitter.abi.Pack("updateAttesterManager", newAttesterManager)
}

// PackUpdatePauser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x554bab3c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updatePauser(address _newPauser) returns()
func (baseMessageTransmitter *BaseMessageTransmitter) PackUpdatePauser(newPauser common.Address) []byte {
	enc, err := baseMessageTransmitter.abi.Pack("updatePauser", newPauser)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUpdatePauser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x554bab3c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function updatePauser(address _newPauser) returns()
func (baseMessageTransmitter *BaseMessageTransmitter) TryPackUpdatePauser(newPauser common.Address) ([]byte, error) {
	return baseMessageTransmitter.abi.Pack("updatePauser", newPauser)
}

// PackUpdateRescuer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2ab60045.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updateRescuer(address newRescuer) returns()
func (baseMessageTransmitter *BaseMessageTransmitter) PackUpdateRescuer(newRescuer common.Address) []byte {
	enc, err := baseMessageTransmitter.abi.Pack("updateRescuer", newRescuer)
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
func (baseMessageTransmitter *BaseMessageTransmitter) TryPackUpdateRescuer(newRescuer common.Address) ([]byte, error) {
	return baseMessageTransmitter.abi.Pack("updateRescuer", newRescuer)
}

// PackUsedNonces is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfeb61724.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function usedNonces(bytes32 ) view returns(uint256)
func (baseMessageTransmitter *BaseMessageTransmitter) PackUsedNonces(arg0 [32]byte) []byte {
	enc, err := baseMessageTransmitter.abi.Pack("usedNonces", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUsedNonces is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfeb61724.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function usedNonces(bytes32 ) view returns(uint256)
func (baseMessageTransmitter *BaseMessageTransmitter) TryPackUsedNonces(arg0 [32]byte) ([]byte, error) {
	return baseMessageTransmitter.abi.Pack("usedNonces", arg0)
}

// UnpackUsedNonces is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfeb61724.
//
// Solidity: function usedNonces(bytes32 ) view returns(uint256)
func (baseMessageTransmitter *BaseMessageTransmitter) UnpackUsedNonces(data []byte) (*big.Int, error) {
	out, err := baseMessageTransmitter.abi.Unpack("usedNonces", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x54fd4d50.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function version() view returns(uint32)
func (baseMessageTransmitter *BaseMessageTransmitter) PackVersion() []byte {
	enc, err := baseMessageTransmitter.abi.Pack("version")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x54fd4d50.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function version() view returns(uint32)
func (baseMessageTransmitter *BaseMessageTransmitter) TryPackVersion() ([]byte, error) {
	return baseMessageTransmitter.abi.Pack("version")
}

// UnpackVersion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x54fd4d50.
//
// Solidity: function version() view returns(uint32)
func (baseMessageTransmitter *BaseMessageTransmitter) UnpackVersion(data []byte) (uint32, error) {
	out, err := baseMessageTransmitter.abi.Unpack("version", data)
	if err != nil {
		return *new(uint32), err
	}
	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	return out0, nil
}

// BaseMessageTransmitterAttesterDisabled represents a AttesterDisabled event raised by the BaseMessageTransmitter contract.
type BaseMessageTransmitterAttesterDisabled struct {
	Attester common.Address
	Raw      *types.Log // Blockchain specific contextual infos
}

const BaseMessageTransmitterAttesterDisabledEventName = "AttesterDisabled"

// ContractEventName returns the user-defined event name.
func (BaseMessageTransmitterAttesterDisabled) ContractEventName() string {
	return BaseMessageTransmitterAttesterDisabledEventName
}

// UnpackAttesterDisabledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AttesterDisabled(address indexed attester)
func (baseMessageTransmitter *BaseMessageTransmitter) UnpackAttesterDisabledEvent(log *types.Log) (*BaseMessageTransmitterAttesterDisabled, error) {
	event := "AttesterDisabled"
	if len(log.Topics) == 0 || log.Topics[0] != baseMessageTransmitter.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BaseMessageTransmitterAttesterDisabled)
	if len(log.Data) > 0 {
		if err := baseMessageTransmitter.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range baseMessageTransmitter.abi.Events[event].Inputs {
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

// BaseMessageTransmitterAttesterEnabled represents a AttesterEnabled event raised by the BaseMessageTransmitter contract.
type BaseMessageTransmitterAttesterEnabled struct {
	Attester common.Address
	Raw      *types.Log // Blockchain specific contextual infos
}

const BaseMessageTransmitterAttesterEnabledEventName = "AttesterEnabled"

// ContractEventName returns the user-defined event name.
func (BaseMessageTransmitterAttesterEnabled) ContractEventName() string {
	return BaseMessageTransmitterAttesterEnabledEventName
}

// UnpackAttesterEnabledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AttesterEnabled(address indexed attester)
func (baseMessageTransmitter *BaseMessageTransmitter) UnpackAttesterEnabledEvent(log *types.Log) (*BaseMessageTransmitterAttesterEnabled, error) {
	event := "AttesterEnabled"
	if len(log.Topics) == 0 || log.Topics[0] != baseMessageTransmitter.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BaseMessageTransmitterAttesterEnabled)
	if len(log.Data) > 0 {
		if err := baseMessageTransmitter.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range baseMessageTransmitter.abi.Events[event].Inputs {
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

// BaseMessageTransmitterAttesterManagerUpdated represents a AttesterManagerUpdated event raised by the BaseMessageTransmitter contract.
type BaseMessageTransmitterAttesterManagerUpdated struct {
	PreviousAttesterManager common.Address
	NewAttesterManager      common.Address
	Raw                     *types.Log // Blockchain specific contextual infos
}

const BaseMessageTransmitterAttesterManagerUpdatedEventName = "AttesterManagerUpdated"

// ContractEventName returns the user-defined event name.
func (BaseMessageTransmitterAttesterManagerUpdated) ContractEventName() string {
	return BaseMessageTransmitterAttesterManagerUpdatedEventName
}

// UnpackAttesterManagerUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AttesterManagerUpdated(address indexed previousAttesterManager, address indexed newAttesterManager)
func (baseMessageTransmitter *BaseMessageTransmitter) UnpackAttesterManagerUpdatedEvent(log *types.Log) (*BaseMessageTransmitterAttesterManagerUpdated, error) {
	event := "AttesterManagerUpdated"
	if len(log.Topics) == 0 || log.Topics[0] != baseMessageTransmitter.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BaseMessageTransmitterAttesterManagerUpdated)
	if len(log.Data) > 0 {
		if err := baseMessageTransmitter.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range baseMessageTransmitter.abi.Events[event].Inputs {
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

// BaseMessageTransmitterInitialized represents a Initialized event raised by the BaseMessageTransmitter contract.
type BaseMessageTransmitterInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const BaseMessageTransmitterInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (BaseMessageTransmitterInitialized) ContractEventName() string {
	return BaseMessageTransmitterInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (baseMessageTransmitter *BaseMessageTransmitter) UnpackInitializedEvent(log *types.Log) (*BaseMessageTransmitterInitialized, error) {
	event := "Initialized"
	if len(log.Topics) == 0 || log.Topics[0] != baseMessageTransmitter.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BaseMessageTransmitterInitialized)
	if len(log.Data) > 0 {
		if err := baseMessageTransmitter.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range baseMessageTransmitter.abi.Events[event].Inputs {
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

// BaseMessageTransmitterMaxMessageBodySizeUpdated represents a MaxMessageBodySizeUpdated event raised by the BaseMessageTransmitter contract.
type BaseMessageTransmitterMaxMessageBodySizeUpdated struct {
	NewMaxMessageBodySize *big.Int
	Raw                   *types.Log // Blockchain specific contextual infos
}

const BaseMessageTransmitterMaxMessageBodySizeUpdatedEventName = "MaxMessageBodySizeUpdated"

// ContractEventName returns the user-defined event name.
func (BaseMessageTransmitterMaxMessageBodySizeUpdated) ContractEventName() string {
	return BaseMessageTransmitterMaxMessageBodySizeUpdatedEventName
}

// UnpackMaxMessageBodySizeUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MaxMessageBodySizeUpdated(uint256 newMaxMessageBodySize)
func (baseMessageTransmitter *BaseMessageTransmitter) UnpackMaxMessageBodySizeUpdatedEvent(log *types.Log) (*BaseMessageTransmitterMaxMessageBodySizeUpdated, error) {
	event := "MaxMessageBodySizeUpdated"
	if len(log.Topics) == 0 || log.Topics[0] != baseMessageTransmitter.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BaseMessageTransmitterMaxMessageBodySizeUpdated)
	if len(log.Data) > 0 {
		if err := baseMessageTransmitter.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range baseMessageTransmitter.abi.Events[event].Inputs {
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

// BaseMessageTransmitterOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the BaseMessageTransmitter contract.
type BaseMessageTransmitterOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const BaseMessageTransmitterOwnershipTransferStartedEventName = "OwnershipTransferStarted"

// ContractEventName returns the user-defined event name.
func (BaseMessageTransmitterOwnershipTransferStarted) ContractEventName() string {
	return BaseMessageTransmitterOwnershipTransferStartedEventName
}

// UnpackOwnershipTransferStartedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (baseMessageTransmitter *BaseMessageTransmitter) UnpackOwnershipTransferStartedEvent(log *types.Log) (*BaseMessageTransmitterOwnershipTransferStarted, error) {
	event := "OwnershipTransferStarted"
	if len(log.Topics) == 0 || log.Topics[0] != baseMessageTransmitter.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BaseMessageTransmitterOwnershipTransferStarted)
	if len(log.Data) > 0 {
		if err := baseMessageTransmitter.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range baseMessageTransmitter.abi.Events[event].Inputs {
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

// BaseMessageTransmitterOwnershipTransferred represents a OwnershipTransferred event raised by the BaseMessageTransmitter contract.
type BaseMessageTransmitterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const BaseMessageTransmitterOwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (BaseMessageTransmitterOwnershipTransferred) ContractEventName() string {
	return BaseMessageTransmitterOwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (baseMessageTransmitter *BaseMessageTransmitter) UnpackOwnershipTransferredEvent(log *types.Log) (*BaseMessageTransmitterOwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if len(log.Topics) == 0 || log.Topics[0] != baseMessageTransmitter.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BaseMessageTransmitterOwnershipTransferred)
	if len(log.Data) > 0 {
		if err := baseMessageTransmitter.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range baseMessageTransmitter.abi.Events[event].Inputs {
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

// BaseMessageTransmitterPause represents a Pause event raised by the BaseMessageTransmitter contract.
type BaseMessageTransmitterPause struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const BaseMessageTransmitterPauseEventName = "Pause"

// ContractEventName returns the user-defined event name.
func (BaseMessageTransmitterPause) ContractEventName() string {
	return BaseMessageTransmitterPauseEventName
}

// UnpackPauseEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Pause()
func (baseMessageTransmitter *BaseMessageTransmitter) UnpackPauseEvent(log *types.Log) (*BaseMessageTransmitterPause, error) {
	event := "Pause"
	if len(log.Topics) == 0 || log.Topics[0] != baseMessageTransmitter.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BaseMessageTransmitterPause)
	if len(log.Data) > 0 {
		if err := baseMessageTransmitter.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range baseMessageTransmitter.abi.Events[event].Inputs {
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

// BaseMessageTransmitterPauserChanged represents a PauserChanged event raised by the BaseMessageTransmitter contract.
type BaseMessageTransmitterPauserChanged struct {
	NewAddress common.Address
	Raw        *types.Log // Blockchain specific contextual infos
}

const BaseMessageTransmitterPauserChangedEventName = "PauserChanged"

// ContractEventName returns the user-defined event name.
func (BaseMessageTransmitterPauserChanged) ContractEventName() string {
	return BaseMessageTransmitterPauserChangedEventName
}

// UnpackPauserChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PauserChanged(address indexed newAddress)
func (baseMessageTransmitter *BaseMessageTransmitter) UnpackPauserChangedEvent(log *types.Log) (*BaseMessageTransmitterPauserChanged, error) {
	event := "PauserChanged"
	if len(log.Topics) == 0 || log.Topics[0] != baseMessageTransmitter.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BaseMessageTransmitterPauserChanged)
	if len(log.Data) > 0 {
		if err := baseMessageTransmitter.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range baseMessageTransmitter.abi.Events[event].Inputs {
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

// BaseMessageTransmitterRescuerChanged represents a RescuerChanged event raised by the BaseMessageTransmitter contract.
type BaseMessageTransmitterRescuerChanged struct {
	NewRescuer common.Address
	Raw        *types.Log // Blockchain specific contextual infos
}

const BaseMessageTransmitterRescuerChangedEventName = "RescuerChanged"

// ContractEventName returns the user-defined event name.
func (BaseMessageTransmitterRescuerChanged) ContractEventName() string {
	return BaseMessageTransmitterRescuerChangedEventName
}

// UnpackRescuerChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RescuerChanged(address indexed newRescuer)
func (baseMessageTransmitter *BaseMessageTransmitter) UnpackRescuerChangedEvent(log *types.Log) (*BaseMessageTransmitterRescuerChanged, error) {
	event := "RescuerChanged"
	if len(log.Topics) == 0 || log.Topics[0] != baseMessageTransmitter.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BaseMessageTransmitterRescuerChanged)
	if len(log.Data) > 0 {
		if err := baseMessageTransmitter.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range baseMessageTransmitter.abi.Events[event].Inputs {
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

// BaseMessageTransmitterSignatureThresholdUpdated represents a SignatureThresholdUpdated event raised by the BaseMessageTransmitter contract.
type BaseMessageTransmitterSignatureThresholdUpdated struct {
	OldSignatureThreshold *big.Int
	NewSignatureThreshold *big.Int
	Raw                   *types.Log // Blockchain specific contextual infos
}

const BaseMessageTransmitterSignatureThresholdUpdatedEventName = "SignatureThresholdUpdated"

// ContractEventName returns the user-defined event name.
func (BaseMessageTransmitterSignatureThresholdUpdated) ContractEventName() string {
	return BaseMessageTransmitterSignatureThresholdUpdatedEventName
}

// UnpackSignatureThresholdUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SignatureThresholdUpdated(uint256 oldSignatureThreshold, uint256 newSignatureThreshold)
func (baseMessageTransmitter *BaseMessageTransmitter) UnpackSignatureThresholdUpdatedEvent(log *types.Log) (*BaseMessageTransmitterSignatureThresholdUpdated, error) {
	event := "SignatureThresholdUpdated"
	if len(log.Topics) == 0 || log.Topics[0] != baseMessageTransmitter.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BaseMessageTransmitterSignatureThresholdUpdated)
	if len(log.Data) > 0 {
		if err := baseMessageTransmitter.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range baseMessageTransmitter.abi.Events[event].Inputs {
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

// BaseMessageTransmitterUnpause represents a Unpause event raised by the BaseMessageTransmitter contract.
type BaseMessageTransmitterUnpause struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const BaseMessageTransmitterUnpauseEventName = "Unpause"

// ContractEventName returns the user-defined event name.
func (BaseMessageTransmitterUnpause) ContractEventName() string {
	return BaseMessageTransmitterUnpauseEventName
}

// UnpackUnpauseEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Unpause()
func (baseMessageTransmitter *BaseMessageTransmitter) UnpackUnpauseEvent(log *types.Log) (*BaseMessageTransmitterUnpause, error) {
	event := "Unpause"
	if len(log.Topics) == 0 || log.Topics[0] != baseMessageTransmitter.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BaseMessageTransmitterUnpause)
	if len(log.Data) > 0 {
		if err := baseMessageTransmitter.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range baseMessageTransmitter.abi.Events[event].Inputs {
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

// ECDSAMetaData contains all meta data concerning the ECDSA contract.
var ECDSAMetaData = bind.MetaData{
	ABI: "[]",
	ID:  "af14cda50dc10ffd723ca623e5dc6e1102",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122011d62dcf47c83b4780ac4c7fc6a975be401ad9880850d3c6ce0d8a240930db4564736f6c63430007060033",
}

// ECDSA is an auto generated Go binding around an Ethereum contract.
type ECDSA struct {
	abi abi.ABI
}

// NewECDSA creates a new instance of ECDSA.
func NewECDSA() *ECDSA {
	parsed, err := ECDSAMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ECDSA{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ECDSA) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// EnumerableSetMetaData contains all meta data concerning the EnumerableSet contract.
var EnumerableSetMetaData = bind.MetaData{
	ABI: "[]",
	ID:  "9f2da77e677d64f91377233b6f7ea225d6",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122075d3abd9deeb9cc2b2478eb7cdbebabf7e66401c459dcf84c72d3ccc38aa686864736f6c63430007060033",
}

// EnumerableSet is an auto generated Go binding around an Ethereum contract.
type EnumerableSet struct {
	abi abi.ABI
}

// NewEnumerableSet creates a new instance of EnumerableSet.
func NewEnumerableSet() *EnumerableSet {
	parsed, err := EnumerableSetMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &EnumerableSet{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *EnumerableSet) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
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

// IMessageTransmitterV2MetaData contains all meta data concerning the IMessageTransmitterV2 contract.
var IMessageTransmitterV2MetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"receiveMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCaller\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"minFinalityThreshold\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "598147a9182fd279a8b6346d117230ae53",
}

// IMessageTransmitterV2 is an auto generated Go binding around an Ethereum contract.
type IMessageTransmitterV2 struct {
	abi abi.ABI
}

// NewIMessageTransmitterV2 creates a new instance of IMessageTransmitterV2.
func NewIMessageTransmitterV2() *IMessageTransmitterV2 {
	parsed, err := IMessageTransmitterV2MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IMessageTransmitterV2{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IMessageTransmitterV2) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackReceiveMessage is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x57ecfd28.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function receiveMessage(bytes message, bytes signature) returns(bool success)
func (iMessageTransmitterV2 *IMessageTransmitterV2) PackReceiveMessage(message []byte, signature []byte) []byte {
	enc, err := iMessageTransmitterV2.abi.Pack("receiveMessage", message, signature)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackReceiveMessage is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x57ecfd28.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function receiveMessage(bytes message, bytes signature) returns(bool success)
func (iMessageTransmitterV2 *IMessageTransmitterV2) TryPackReceiveMessage(message []byte, signature []byte) ([]byte, error) {
	return iMessageTransmitterV2.abi.Pack("receiveMessage", message, signature)
}

// UnpackReceiveMessage is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x57ecfd28.
//
// Solidity: function receiveMessage(bytes message, bytes signature) returns(bool success)
func (iMessageTransmitterV2 *IMessageTransmitterV2) UnpackReceiveMessage(data []byte) (bool, error) {
	out, err := iMessageTransmitterV2.abi.Unpack("receiveMessage", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackSendMessage is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x14b157ab.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function sendMessage(uint32 destinationDomain, bytes32 recipient, bytes32 destinationCaller, uint32 minFinalityThreshold, bytes messageBody) returns()
func (iMessageTransmitterV2 *IMessageTransmitterV2) PackSendMessage(destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, minFinalityThreshold uint32, messageBody []byte) []byte {
	enc, err := iMessageTransmitterV2.abi.Pack("sendMessage", destinationDomain, recipient, destinationCaller, minFinalityThreshold, messageBody)
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
func (iMessageTransmitterV2 *IMessageTransmitterV2) TryPackSendMessage(destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, minFinalityThreshold uint32, messageBody []byte) ([]byte, error) {
	return iMessageTransmitterV2.abi.Pack("sendMessage", destinationDomain, recipient, destinationCaller, minFinalityThreshold, messageBody)
}

// IReceiverMetaData contains all meta data concerning the IReceiver contract.
var IReceiverMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"receiveMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "3e5c09894de133f8c8ea2352dcd5d7cff7",
}

// IReceiver is an auto generated Go binding around an Ethereum contract.
type IReceiver struct {
	abi abi.ABI
}

// NewIReceiver creates a new instance of IReceiver.
func NewIReceiver() *IReceiver {
	parsed, err := IReceiverMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IReceiver{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IReceiver) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackReceiveMessage is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x57ecfd28.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function receiveMessage(bytes message, bytes signature) returns(bool success)
func (iReceiver *IReceiver) PackReceiveMessage(message []byte, signature []byte) []byte {
	enc, err := iReceiver.abi.Pack("receiveMessage", message, signature)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackReceiveMessage is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x57ecfd28.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function receiveMessage(bytes message, bytes signature) returns(bool success)
func (iReceiver *IReceiver) TryPackReceiveMessage(message []byte, signature []byte) ([]byte, error) {
	return iReceiver.abi.Pack("receiveMessage", message, signature)
}

// UnpackReceiveMessage is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x57ecfd28.
//
// Solidity: function receiveMessage(bytes message, bytes signature) returns(bool success)
func (iReceiver *IReceiver) UnpackReceiveMessage(data []byte) (bool, error) {
	out, err := iReceiver.abi.Unpack("receiveMessage", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// IReceiverV2MetaData contains all meta data concerning the IReceiverV2 contract.
var IReceiverV2MetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"receiveMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "6b880bd596b196d29f213d80b398dfe6be",
}

// IReceiverV2 is an auto generated Go binding around an Ethereum contract.
type IReceiverV2 struct {
	abi abi.ABI
}

// NewIReceiverV2 creates a new instance of IReceiverV2.
func NewIReceiverV2() *IReceiverV2 {
	parsed, err := IReceiverV2MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IReceiverV2{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IReceiverV2) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackReceiveMessage is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x57ecfd28.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function receiveMessage(bytes message, bytes signature) returns(bool success)
func (iReceiverV2 *IReceiverV2) PackReceiveMessage(message []byte, signature []byte) []byte {
	enc, err := iReceiverV2.abi.Pack("receiveMessage", message, signature)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackReceiveMessage is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x57ecfd28.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function receiveMessage(bytes message, bytes signature) returns(bool success)
func (iReceiverV2 *IReceiverV2) TryPackReceiveMessage(message []byte, signature []byte) ([]byte, error) {
	return iReceiverV2.abi.Pack("receiveMessage", message, signature)
}

// UnpackReceiveMessage is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x57ecfd28.
//
// Solidity: function receiveMessage(bytes message, bytes signature) returns(bool success)
func (iReceiverV2 *IReceiverV2) UnpackReceiveMessage(data []byte) (bool, error) {
	out, err := iReceiverV2.abi.Unpack("receiveMessage", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
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

// MessageTransmitterV2MetaData contains all meta data concerning the MessageTransmitterV2 contract.
var MessageTransmitterV2MetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_localDomain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_version\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"}],\"name\":\"AttesterDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"}],\"name\":\"AttesterEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousAttesterManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAttesterManager\",\"type\":\"address\"}],\"name\":\"AttesterManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxMessageBodySize\",\"type\":\"uint256\"}],\"name\":\"MaxMessageBodySizeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"sourceDomain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"finalityThresholdExecuted\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"}],\"name\":\"MessageReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"MessageSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"PauserChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newRescuer\",\"type\":\"address\"}],\"name\":\"RescuerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldSignatureThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSignatureThreshold\",\"type\":\"uint256\"}],\"name\":\"SignatureThresholdUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NONCE_USED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"attesterManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"}],\"name\":\"disableAttester\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAttester\",\"type\":\"address\"}],\"name\":\"enableAttester\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getEnabledAttester\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumEnabledAttesters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pauser_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rescuer_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"attesterManager_\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"attesters_\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"signatureThreshold_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxMessageBodySize_\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedVersion\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"}],\"name\":\"isEnabledAttester\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxMessageBodySize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"name\":\"receiveMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rescuer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCaller\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"minFinalityThreshold\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMaxMessageBodySize\",\"type\":\"uint256\"}],\"name\":\"setMaxMessageBodySize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newSignatureThreshold\",\"type\":\"uint256\"}],\"name\":\"setSignatureThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signatureThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAttesterManager\",\"type\":\"address\"}],\"name\":\"updateAttesterManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newPauser\",\"type\":\"address\"}],\"name\":\"updatePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRescuer\",\"type\":\"address\"}],\"name\":\"updateRescuer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"usedNonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "dad69a46ecc99afcf5676880a11b9fd8c2",
	Bin: "0x60c06040526000600260146101000a81548160ff0219169083151502179055503480156200002c57600080fd5b506040516200521238038062005212833981810160405260408110156200005257600080fd5b81019080805190602001909291908051906020019092919050505081813362000090620000846200010360201b60201c565b6200010b60201b60201c565b620000a1336200014960201b60201c565b6001600481905550620000ba81620002b360201b60201c565b508163ffffffff1660808163ffffffff1660e01b815250508063ffffffff1660a08163ffffffff1660e01b815250505050620000fb6200038b60201b60201c565b5050620007f6565b600033905090565b600160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556200014681620004bd60201b620020341760201c565b50565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415620001ed576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f496e76616c6964206174746573746572206d616e61676572206164647265737381525060200191505060405180910390fd5b6000600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f0cee1b7ae04f3c788dd3a46c6fa677eb95b913611ef7ab59524fdc09d346021960405160405180910390a35050565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161462000377576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f43616c6c6572206e6f74206174746573746572206d616e61676572000000000081525060200191505060405180910390fd5b62000388816200058160201b60201c565b50565b60006200039d620006f960201b60201c565b90508060000160089054906101000a900460ff161562000409576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180620051ed6025913960400191505060405180910390fd5b67ffffffffffffffff80168160000160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1614620004ba5767ffffffffffffffff8160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d267ffffffffffffffff604051808267ffffffffffffffff16815260200191505060405180910390a15b50565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141562000625576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f4e6577206174746573746572206d757374206265206e6f6e7a65726f0000000081525060200191505060405180910390fd5b620006408160056200072160201b620020f81790919060201c565b620006b3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f417474657374657220616c726561647920656e61626c6564000000000000000081525060200191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff167f5b99bab45c72ce67e89466dbc47480b9c1fde1400e7268bbf463b8354ee4653f60405160405180910390a250565b60007ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b600062000751836000018373ffffffffffffffffffffffffffffffffffffffff1660001b6200075960201b60201c565b905092915050565b60006200076d8383620007d360201b60201c565b620007c8578260000182908060018154018082558091505060019003906000526020600020016000909190919091505582600001805490508360010160008481526020019081526020016000208190555060019050620007cd565b600090505b92915050565b600080836001016000848152602001908152602001600020541415905092915050565b60805160e01c60a05160e01c6149b46200083960003980610c5a52806111175280612743525080610a9c5280610c7b5280611be152806125da52506149b46000f3fe608060405234801561001057600080fd5b50600436106101e55760003560e01c80638456cb591161010f578063b2118a8d116100a2578063e30c397811610071578063e30c39781461090a578063f2fde38b1461093e578063fae3687914610982578063feb61724146109c6576101e5565b8063b2118a8d146107d2578063bbde537414610840578063beb673d81461086e578063de7769d4146108c6576101e5565b80639b0d94b7116100de5780639b0d94b71461072e5780639fd0506d14610762578063a82f2e2614610796578063af47b9bb146107b4576101e5565b80638456cb591461069e5780638d3638f4146106a85780638da5cb5b146106cc57806392492c6814610700576101e5565b806354fd4d50116101875780635f747e7d116101565780635f747e7d1461050f57806379ba50971461061c5780637af82f60146106265780637de25ae414610680576101e5565b806354fd4d50146103a3578063554bab3c146103c757806357ecfd281461040b5780635c975abb146104ef576101e5565b80632d025080116101c35780632d0250801461030357806338a63183146103475780633f4ba83a1461037b57806351079a5314610385576101e5565b806308828eb7146101ea57806314b157ab146102125780632ab60045146102bf575b600080fd5b6101f2610a08565b604051808267ffffffffffffffff16815260200191505060405180910390f35b6102bd600480360360a081101561022857600080fd5b81019080803563ffffffff1690602001909291908035906020019092919080359060200190929190803563ffffffff1690602001909291908035906020019064010000000081111561027957600080fd5b82018360208201111561028b57600080fd5b803590602001918460018302840111640100000000831117156102ad57600080fd5b9091929391929390505050610a17565b005b610301600480360360208110156102d557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610d4e565b005b6103456004803603602081101561031957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610d62565b005b61034f610feb565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610383611015565b005b61038d611104565b6040518082815260200191505060405180910390f35b6103ab611115565b604051808263ffffffff16815260200191505060405180910390f35b610409600480360360208110156103dd57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611139565b005b6104d76004803603604081101561042157600080fd5b810190808035906020019064010000000081111561043e57600080fd5b82018360208201111561045057600080fd5b8035906020019184600183028401116401000000008311171561047257600080fd5b90919293919293908035906020019064010000000081111561049357600080fd5b8201836020820111156104a557600080fd5b803590602001918460018302840111640100000000831117156104c757600080fd5b909192939192939050505061114d565b60405180821515815260200191505060405180910390f35b6104f76115e3565b60405180821515815260200191505060405180910390f35b61061a600480360360e081101561052557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001906401000000008111156105c257600080fd5b8201836020820111156105d457600080fd5b803590602001918460208302840111640100000000831117156105f657600080fd5b909192939192939080359060200190929190803590602001909291905050506115f6565b005b610624611a2b565b005b6106686004803603602081101561063c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611ace565b60405180821515815260200191505060405180910390f35b610688611aeb565b6040518082815260200191505060405180910390f35b6106a6611af0565b005b6106b0611bdf565b604051808263ffffffff16815260200191505060405180910390f35b6106d4611c03565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61072c6004803603602081101561071657600080fd5b8101908080359060200190929190505050611c2c565b005b610736611c40565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61076a611c6a565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61079e611c94565b6040518082815260200191505060405180910390f35b6107bc611c9a565b6040518082815260200191505060405180910390f35b61083e600480360360608110156107e857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611ca0565b005b61086c6004803603602081101561085657600080fd5b8101908080359060200190929190505050611d76565b005b61089a6004803603602081101561088457600080fd5b8101908080359060200190929190505050611e45565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610908600480360360208110156108dc57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611e62565b005b610912611e76565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6109806004803603602081101561095457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611ea0565b005b6109c46004803603602081101561099857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611f4d565b005b6109f2600480360360208110156109dc57600080fd5b810190808035906020019092919050505061201c565b6040518082815260200191505060405180910390f35b6000610a12612128565b905090565b600260149054906101000a900460ff1615610a9a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f5061757361626c653a207061757365640000000000000000000000000000000081525060200191505060405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000063ffffffff168663ffffffff161415610b3c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f446f6d61696e206973206c6f63616c20646f6d61696e0000000000000000000081525060200191505060405180910390fd5b601c54828290501115610bb7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f4d65737361676520626f64792065786365656473206d61782073697a6500000081525060200191505060405180910390fd5b6000801b851415610c30576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f526563697069656e74206d757374206265206e6f6e7a65726f0000000000000081525060200191505060405180910390fd5b6000610c513373ffffffffffffffffffffffffffffffffffffffff1661214f565b90506000610ca67f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000008a858b8b8b8b8b612172565b90507f8c5261668696ce22758910d05bab8f186d6eb247ceac2af2e82c7dc17669b036816040518080602001828103825283818151815260200191508051906020019080838360005b83811015610d0a578082015181840152602081019050610cef565b50505050905090810190601f168015610d375780820380516001836020036101000a031916815260200191505b509250505060405180910390a15050505050505050565b610d56612225565b610d5f816122d6565b50565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610e25576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f43616c6c6572206e6f74206174746573746572206d616e61676572000000000081525060200191505060405180910390fd5b6000610e2f611104565b905060018111610ea7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f546f6f2066657720656e61626c6564206174746573746572730000000000000081525060200191505060405180910390fd5b6004548111610f1e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f5369676e6174757265207468726573686f6c6420697320746f6f206c6f77000081525060200191505060405180910390fd5b610f328260056123e390919063ffffffff16565b610fa4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f417474657374657220616c72656164792064697361626c65640000000000000081525060200191505060405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff167f78e573a18c75957b7cadaab01511aa1c19a659f06ecf53e01de37ed92d3261fc60405160405180910390a25050565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146110bb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806148836022913960400191505060405180910390fd5b6000600260146101000a81548160ff0219169083151502179055507f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a1565b60006111106005612413565b905090565b7f000000000000000000000000000000000000000000000000000000000000000081565b611141612225565b61114a81612428565b50565b6000600260149054906101000a900460ff16156111d2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f5061757361626c653a207061757365640000000000000000000000000000000081525060200191505060405180910390fd5b6000806000806000806111e78b8b8b8b612557565b9550955095509550955095506001601d6000888152602001908152602001600020819055506107d063ffffffff168263ffffffff161015611393578273ffffffffffffffffffffffffffffffffffffffff16637c92f219868685856040518563ffffffff1660e01b8152600401808563ffffffff1681526020018481526020018363ffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b838110156112af578082015181840152602081019050611294565b50505050905090810190601f1680156112dc5780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b1580156112fe57600080fd5b505af1158015611312573d6000803e3d6000fd5b505050506040513d602081101561132857600080fd5b810190808051906020019092919050505061138e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260288152602001806149576028913960400191505060405180910390fd5b611500565b8273ffffffffffffffffffffffffffffffffffffffff166311cffb67868685856040518563ffffffff1660e01b8152600401808563ffffffff1681526020018481526020018363ffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b83811015611420578082015181840152602081019050611405565b50505050905090810190601f16801561144d5780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b15801561146f57600080fd5b505af1158015611483573d6000803e3d6000fd5b505050506040513d602081101561149957600080fd5b81019080805190602001909291905050506114ff576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260268152602001806147516026913960400191505060405180910390fd5b5b8163ffffffff16863373ffffffffffffffffffffffffffffffffffffffff167fff48c13eda96b1cceacc6b9edeedc9e9db9d6226afbc30146b720c19d3addb1c888886604051808463ffffffff16815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561159557808201518184015260208101905061157a565b50505050905090810190601f1680156115c25780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a460019650505050505050949350505050565b600260149054906101000a900460ff1681565b6000611600612903565b905060008160000160089054906101000a900460ff1615905060008260000160009054906101000a900467ffffffffffffffff1690506000808267ffffffffffffffff1614801561164e5750825b9050600060018367ffffffffffffffff1614801561167257506116703061292b565b155b9050818061167d5750805b6116d2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806146b46025913960400191505060405180910390fd5b60018560000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555083156117225760018560000160086101000a81548160ff0219169083151502179055505b600073ffffffffffffffffffffffffffffffffffffffff168d73ffffffffffffffffffffffffffffffffffffffff1614156117c5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f4f776e657220697320746865207a65726f20616464726573730000000000000081525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168a73ffffffffffffffffffffffffffffffffffffffff16141561184b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260238152602001806148266023913960400191505060405180910390fd5b888890508711156118a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602581526020018061472c6025913960400191505060405180910390fd5b6000861161191d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f4d61784d657373616765426f647953697a65206973207a65726f00000000000081525060200191505060405180910390fd5b6119268d61293e565b61192f8b6122d6565b6119388c612428565b6119418a61296f565b61194a86612ad8565b600089899050905060005b818110156119995761198e8b8b8381811061196c57fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff16612b1b565b806001019050611955565b506119a388612c8a565b6001601d60008060001b815260200190815260200160002081905550508315611a1c5760008560000160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d260016040518082815260200191505060405180910390a15b50505050505050505050505050565b6000611a35612e49565b90508073ffffffffffffffffffffffffffffffffffffffff16611a56611e76565b73ffffffffffffffffffffffffffffffffffffffff1614611ac2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260298152602001806146d96029913960400191505060405180910390fd5b611acb8161293e565b50565b6000611ae4826005612e5190919063ffffffff16565b9050919050565b600181565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611b96576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806148836022913960400191505060405180910390fd5b6001600260146101000a81548160ff0219169083151502179055507f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b7f000000000000000000000000000000000000000000000000000000000000000081565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b611c34612225565b611c3d81612ad8565b50565b6000600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60045481565b601c5481565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611d46576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806148026024913960400191505060405180910390fd5b611d7182828573ffffffffffffffffffffffffffffffffffffffff16612e819092919063ffffffff16565b505050565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611e39576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f43616c6c6572206e6f74206174746573746572206d616e61676572000000000081525060200191505060405180910390fd5b611e4281612c8a565b50565b6000611e5b826005612f2390919063ffffffff16565b9050919050565b611e6a612225565b611e738161296f565b50565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b611ea8612225565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16611f08611c03565b73ffffffffffffffffffffffffffffffffffffffff167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a350565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612010576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f43616c6c6572206e6f74206174746573746572206d616e61676572000000000081525060200191505060405180910390fd5b61201981612b1b565b50565b601d6020528060005260406000206000915090505481565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000612120836000018373ffffffffffffffffffffffffffffffffffffffff1660001b612f3d565b905092915050565b6000612132612903565b60000160009054906101000a900467ffffffffffffffff16905090565b60008173ffffffffffffffffffffffffffffffffffffffff1660001b9050919050565b60608989896000801b8a8a8a8a60008b8b604051602001808c63ffffffff1660e01b81526004018b63ffffffff1660e01b81526004018a63ffffffff1660e01b81526004018981526020018881526020018781526020018681526020018563ffffffff1660e01b81526004018463ffffffff1660e01b8152600401838380828437808301925050509b50505050505050505050505060405160208183030381529060405290509998505050505050505050565b61222d612e49565b73ffffffffffffffffffffffffffffffffffffffff1661224b611c03565b73ffffffffffffffffffffffffffffffffffffffff16146122d4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561235c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a815260200180614702602a913960400191505060405180910390fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff167fe475e580d85111348e40d8ca33cfdd74c30fe1655c2d8537a13abc10065ffa5a60405160405180910390a250565b600061240b836000018373ffffffffffffffffffffffffffffffffffffffff1660001b612fad565b905092915050565b600061242182600001613095565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156124ae576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602881526020018061468c6028913960400191505060405180910390fd5b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fb80482a293ca2e013eda8683c9bd7fc8347cfdaeea5ede58cba46df502c2a60460405160405180910390a250565b6000806000806000606061256d8a8a8a8a6130a6565b60006125c760008c8c8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050506132f890919063ffffffff16565b90506125d88162ffffff1916613323565b7f000000000000000000000000000000000000000000000000000000000000000063ffffffff1661260e8262ffffff191661343d565b63ffffffff1614612687576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f496e76616c69642064657374696e6174696f6e20646f6d61696e00000000000081525060200191505060405180910390fd5b6000801b61269a8262ffffff1916613466565b14612741576126be3373ffffffffffffffffffffffffffffffffffffffff1661214f565b6126cd8262ffffff1916613466565b14612740576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f496e76616c69642063616c6c657220666f72206d65737361676500000000000081525060200191505060405180910390fd5b5b7f000000000000000000000000000000000000000000000000000000000000000063ffffffff166127778262ffffff191661348f565b63ffffffff16146127f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f496e76616c6964206d6573736167652076657273696f6e00000000000000000081525060200191505060405180910390fd5b6127ff8162ffffff19166134b8565b96506000601d6000898152602001908152602001600020541461288a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f4e6f6e636520616c72656164792075736564000000000000000000000000000081525060200191505060405180910390fd5b6128998162ffffff19166134e1565b95506128aa8162ffffff191661350a565b94506128c36128be8262ffffff1916613533565b61355c565b93506128d48162ffffff1916613569565b92506128f36128e88262ffffff1916613592565b62ffffff19166135df565b9150509499939850945094509450565b60007ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b600080823b905060008111915050919050565b600160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905561296c81612034565b50565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415612a12576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f496e76616c6964206174746573746572206d616e61676572206164647265737381525060200191505060405180910390fd5b6000600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f0cee1b7ae04f3c788dd3a46c6fa677eb95b913611ef7ab59524fdc09d346021960405160405180910390a35050565b80601c819055507fb13bf6bebed03d1b318e3ea32e4b2a3ad9f5e2312cdf340a2f4bbfaee39f928d601c546040518082815260200191505060405180910390a150565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415612bbe576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f4e6577206174746573746572206d757374206265206e6f6e7a65726f0000000081525060200191505060405180910390fd5b612bd28160056120f890919063ffffffff16565b612c44576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f417474657374657220616c726561647920656e61626c6564000000000000000081525060200191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff167f5b99bab45c72ce67e89466dbc47480b9c1fde1400e7268bbf463b8354ee4653f60405160405180910390a250565b6000811415612d01576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f496e76616c6964207369676e6174757265207468726573686f6c64000000000081525060200191505060405180910390fd5b612d0b6005612413565b811115612d80576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4e6577207369676e6174757265207468726573686f6c6420746f6f206869676881525060200191505060405180910390fd5b600454811415612df8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f5369676e6174757265207468726573686f6c6420616c7265616479207365740081525060200191505060405180910390fd5b60006004549050816004819055507f149153f58b4da003a8cfd4523709a202402182cb5aa335046911277a1be6eede8183604051808381526020018281526020019250505060405180910390a15050565b600033905090565b6000612e79836000018373ffffffffffffffffffffffffffffffffffffffff1660001b613626565b905092915050565b612f1e8363a9059cbb60e01b8484604051602401808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050613649565b505050565b6000612f328360000183613738565b60001c905092915050565b6000612f498383613626565b612fa2578260000182908060018154018082558091505060019003906000526020600020016000909190919091505582600001805490508360010160008481526020019081526020016000208190555060019050612fa7565b600090505b92915050565b600080836001016000848152602001908152602001600020549050600081146130895760006001820390506000600186600001805490500390506000866000018281548110612ff857fe5b906000526020600020015490508087600001848154811061301557fe5b906000526020600020018190555060018301876001016000838152602001908152602001600020819055508660000180548061304d57fe5b6001900381819060005260206000200160009055905586600101600087815260200190815260200160002060009055600194505050505061308f565b60009150505b92915050565b600081600001805490509050919050565b6004546041028282905014613123576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f496e76616c6964206174746573746174696f6e206c656e67746800000000000081525060200191505060405180910390fd5b60008085856040518083838082843780830192505050925050506040518091039020905060005b6004548110156132ef576000858560418402906041808602019261317093929190614636565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050905060006131c184836137bb565b90508473ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1611613264576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f496e76616c6964207369676e6174757265206f72646572206f7220647570650081525060200191505060405180910390fd5b61326d81611ace565b6132df576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f496e76616c6964207369676e61747572653a206e6f742061747465737465720081525060200191505060405180910390fd5b809450505080600101905061314a565b50505050505050565b6000808351905060006020850190506133198464ffffffffff1682846137cf565b9250505092915050565b6133328162ffffff191661383e565b6133a4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f4d616c666f726d6564206d65737361676500000000000000000000000000000081525060200191505060405180910390fd5b609460ff166133b88262ffffff1916613881565b6bffffffffffffffffffffffff16101561343a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f496e76616c6964206d6573736167653a20746f6f2073686f727400000000000081525060200191505060405180910390fd5b50565b600061345f600860ff1660048462ffffff19166138a79092919063ffffffff16565b9050919050565b6000613488606c60ff1660208462ffffff19166138cc9092919063ffffffff16565b9050919050565b60006134b1600060ff1660048462ffffff19166138a79092919063ffffffff16565b9050919050565b60006134da600c60ff1660208462ffffff19166138cc9092919063ffffffff16565b9050919050565b6000613503600460ff1660048462ffffff19166138a79092919063ffffffff16565b9050919050565b600061352c602c60ff1660208462ffffff19166138cc9092919063ffffffff16565b9050919050565b6000613555604c60ff1660208462ffffff19166138cc9092919063ffffffff16565b9050919050565b60008160001c9050919050565b600061358b609060ff1660048462ffffff19166138a79092919063ffffffff16565b9050919050565b60006135d8609460ff16609460ff166135b08562ffffff1916613881565b036bffffffffffffffffffffffff1660008562ffffff1916613a93909392919063ffffffff16565b9050919050565b60606000806135ed84613881565b6bffffffffffffffffffffffff16905060405191508192506136128460208401613b3f565b506020818301016040528082525050919050565b600080836001016000848152602001908152602001600020541415905092915050565b60006136ab826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16613cec9092919063ffffffff16565b9050600081511115613733578080602001905160208110156136cc57600080fd5b8101908080519060200190929190505050613732576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a8152602001806148a5602a913960400191505060405180910390fd5b5b505050565b600081836000018054905011613799576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602281526020018061466a6022913960400191505060405180910390fd5b8260000182815481106137a857fe5b9060005260206000200154905092915050565b60006137c78383613d04565b905092915050565b6000806137e58385613db190919063ffffffff16565b90506040518111156137f657600090505b6000811415613828577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000915050613837565b613833858585613e34565b9150505b9392505050565b600064ffffffffff61384f83613e5f565b64ffffffffff161415613865576000905061387c565b600061387083613e74565b90506040518111159150505b919050565b6000806bffffffffffffffffffffffff90506000601890508184821c1692505050919050565b60006008826020030260ff166138be8585856138cc565b60001c901c90509392505050565b6000808260ff1614156138e4576000801b9050613a8c565b6138ed84613881565b6bffffffffffffffffffffffff166139118360ff1685613db190919063ffffffff16565b11156139f35761395261392385613e9e565b6bffffffffffffffffffffffff1661393a86613881565b6bffffffffffffffffffffffff16858560ff16613ec7565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b838110156139b857808201518184015260208101905061399d565b50505050905090810190601f1680156139e55780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b60208260ff161115613a50576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603a815260200180614849603a913960400191505060405180910390fd5b60006008830290506000613a6386613e9e565b6bffffffffffffffffffffffff1690506000613a7e83614001565b905080868301511693505050505b9392505050565b600080613a9f86613e9e565b6bffffffffffffffffffffffff169050613ab886613e74565b613add85613acf8885613db190919063ffffffff16565b613db190919063ffffffff16565b1115613b0c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000915050613b37565b613b1f8582613db190919063ffffffff16565b9050613b338364ffffffffff1682866137cf565b9150505b949350505050565b6000613b4a83614030565b613b9f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260288152602001806148cf6028913960400191505060405180910390fd5b613ba88361383e565b613bfd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602b8152602001806148f7602b913960400191505060405180910390fd5b6000613c0884613881565b6bffffffffffffffffffffffff1690506000613c2385613e9e565b6bffffffffffffffffffffffff169050600080604051915085821115613c495760206060fd5b8386858560045afa905080613cc6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f6964656e74697479204f4f47000000000000000000000000000000000000000081525060200191505060405180910390fd5b613ce0613cd288613e5f565b64ffffffffff168786613e34565b94505050505092915050565b6060613cfb8484600085614043565b90509392505050565b60006041825114613d7d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f45434453413a20696e76616c6964207369676e6174757265206c656e6774680081525060200191505060405180910390fd5b60008060006020850151925060408501519150606085015160001a9050613da6868285856141eb565b935050505092915050565b6000818301905082811015613e2e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f4f766572666c6f7720647572696e67206164646974696f6e2e0000000000000081525060200191505060405180910390fd5b92915050565b60008060609050600060189050858317821b9250848317821b9250838317811b925050509392505050565b60008060d860ff16905082811c915050919050565b6000613e7f82613881565b613e8883613e9e565b016bffffffffffffffffffffffff169050919050565b6000806bffffffffffffffffffffffff90506000607860ff1690508184821c1692505050919050565b60606000613ed4866143ea565b9150506000613ee2866143ea565b9150506000613ef0866143ea565b9150506000613efe866143ea565b915050838383836040516020018080614922603591396035018565ffffffffffff1660d01b8152600601807f2077697468206c656e6774682030780000000000000000000000000000000000815250600f018465ffffffffffff1660d01b8152600601806147bf602191396021018365ffffffffffff1660d01b8152600601807f2077697468206c656e6774682030780000000000000000000000000000000000815250600f018265ffffffffffff1660d01b8152600601807f2e00000000000000000000000000000000000000000000000000000000000000815250600101945050505050604051602081830303815290604052945050505050949350505050565b60007f8000000000000000000000000000000000000000000000000000000000000000600183031d9050919050565b600061403b82614496565b159050919050565b60608247101561409e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260268152602001806147996026913960400191505060405180910390fd5b6140a78561292b565b614119576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000081525060200191505060405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040518082805190602001908083835b602083106141685780518252602082019150602081019050602083039250614145565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d80600081146141ca576040519150601f19603f3d011682016040523d82523d6000602084013e6141cf565b606091505b50915091506141df8282866144ce565b92505050949350505050565b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08260001c1115614269576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806147776022913960400191505060405180910390fd5b601b8460ff16148061427e5750601c8460ff16145b6142d3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806147e06022913960400191505060405180910390fd5b600060018686868660405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa15801561432f573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156143de576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f45434453413a20696e76616c6964207369676e6174757265000000000000000081525060200191505060405180910390fd5b80915050949350505050565b6000806000601f90505b600f8160ff16111561443e5760006008820260ff1685901c90506144178161459a565b61ffff168417935060108260ff161461443257601084901b93505b506001810390506143f4565b506000600f90505b60ff8160ff1610156144905760006008820260ff1685901c90506144698161459a565b61ffff168317925060008260ff161461448457601083901b92505b50600181039050614446565b50915091565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000062ffffff19168262ffffff1916149050919050565b606083156144de57829050614593565b6000835111156144f15782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561455857808201518184015260208101905061453d565b50505050905090810190601f1680156145855780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b9392505050565b60006145ac60048360ff16901c6145d3565b60ff168117905060088161ffff16901b90506145c7826145d3565b60ff1681179050919050565b600080600f831690506040518060400160405280601081526020017f30313233343536373839616263646566000000000000000000000000000000008152508160ff168151811061462057fe5b602001015160f81c60f81b60f81c915050919050565b6000808585111561464657600080fd5b8386111561465357600080fd5b600185028301915084860390509450949250505056fe456e756d657261626c655365743a20696e646578206f7574206f6620626f756e64735061757361626c653a206e65772070617573657220697320746865207a65726f2061646472657373496e697469616c697a61626c653a20696e76616c696420696e697469616c697a6174696f6e4f776e61626c6532537465703a2063616c6c6572206973206e6f7420746865206e6577206f776e6572526573637561626c653a206e6577207265736375657220697320746865207a65726f20616464726573735369676e6174757265207468726573686f6c6420657863656564732061747465737465727368616e646c655265636569766546696e616c697a65644d6573736167652829206661696c656445434453413a20696e76616c6964207369676e6174757265202773272076616c7565416464726573733a20696e73756666696369656e742062616c616e636520666f722063616c6c2e20417474656d7074656420746f20696e646578206174206f666673657420307845434453413a20696e76616c6964207369676e6174757265202776272076616c7565526573637561626c653a2063616c6c6572206973206e6f7420746865207265736375657241747465737465724d616e6167657220697320746865207a65726f206164647265737354797065644d656d566965772f696e646578202d20417474656d7074656420746f20696e646578206d6f7265207468616e2033322062797465735061757361626c653a2063616c6c6572206973206e6f7420746865207061757365725361666545524332303a204552433230206f7065726174696f6e20646964206e6f74207375636365656454797065644d656d566965772f636f7079546f202d204e756c6c20706f696e74657220646572656654797065644d656d566965772f636f7079546f202d20496e76616c696420706f696e74657220646572656654797065644d656d566965772f696e646578202d204f76657272616e2074686520766965772e20536c69636520697320617420307868616e646c6552656365697665556e66696e616c697a65644d6573736167652829206661696c6564a264697066735822122090a8f59d6573c85f087ed4737f218ce45c30830634bd70dd594c10b9efc7672164736f6c63430007060033496e697469616c697a61626c653a20696e76616c696420696e697469616c697a6174696f6e",
}

// MessageTransmitterV2 is an auto generated Go binding around an Ethereum contract.
type MessageTransmitterV2 struct {
	abi abi.ABI
}

// NewMessageTransmitterV2 creates a new instance of MessageTransmitterV2.
func NewMessageTransmitterV2() *MessageTransmitterV2 {
	parsed, err := MessageTransmitterV2MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &MessageTransmitterV2{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *MessageTransmitterV2) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(uint32 _localDomain, uint32 _version) returns()
func (messageTransmitterV2 *MessageTransmitterV2) PackConstructor(_localDomain uint32, _version uint32) []byte {
	enc, err := messageTransmitterV2.abi.Pack("", _localDomain, _version)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackNONCEUSED is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7de25ae4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function NONCE_USED() view returns(uint256)
func (messageTransmitterV2 *MessageTransmitterV2) PackNONCEUSED() []byte {
	enc, err := messageTransmitterV2.abi.Pack("NONCE_USED")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackNONCEUSED is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7de25ae4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function NONCE_USED() view returns(uint256)
func (messageTransmitterV2 *MessageTransmitterV2) TryPackNONCEUSED() ([]byte, error) {
	return messageTransmitterV2.abi.Pack("NONCE_USED")
}

// UnpackNONCEUSED is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7de25ae4.
//
// Solidity: function NONCE_USED() view returns(uint256)
func (messageTransmitterV2 *MessageTransmitterV2) UnpackNONCEUSED(data []byte) (*big.Int, error) {
	out, err := messageTransmitterV2.abi.Unpack("NONCE_USED", data)
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
func (messageTransmitterV2 *MessageTransmitterV2) PackAcceptOwnership() []byte {
	enc, err := messageTransmitterV2.abi.Pack("acceptOwnership")
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
func (messageTransmitterV2 *MessageTransmitterV2) TryPackAcceptOwnership() ([]byte, error) {
	return messageTransmitterV2.abi.Pack("acceptOwnership")
}

// PackAttesterManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9b0d94b7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function attesterManager() view returns(address)
func (messageTransmitterV2 *MessageTransmitterV2) PackAttesterManager() []byte {
	enc, err := messageTransmitterV2.abi.Pack("attesterManager")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAttesterManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9b0d94b7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function attesterManager() view returns(address)
func (messageTransmitterV2 *MessageTransmitterV2) TryPackAttesterManager() ([]byte, error) {
	return messageTransmitterV2.abi.Pack("attesterManager")
}

// UnpackAttesterManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9b0d94b7.
//
// Solidity: function attesterManager() view returns(address)
func (messageTransmitterV2 *MessageTransmitterV2) UnpackAttesterManager(data []byte) (common.Address, error) {
	out, err := messageTransmitterV2.abi.Unpack("attesterManager", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackDisableAttester is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2d025080.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function disableAttester(address attester) returns()
func (messageTransmitterV2 *MessageTransmitterV2) PackDisableAttester(attester common.Address) []byte {
	enc, err := messageTransmitterV2.abi.Pack("disableAttester", attester)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDisableAttester is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2d025080.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function disableAttester(address attester) returns()
func (messageTransmitterV2 *MessageTransmitterV2) TryPackDisableAttester(attester common.Address) ([]byte, error) {
	return messageTransmitterV2.abi.Pack("disableAttester", attester)
}

// PackEnableAttester is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfae36879.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function enableAttester(address newAttester) returns()
func (messageTransmitterV2 *MessageTransmitterV2) PackEnableAttester(newAttester common.Address) []byte {
	enc, err := messageTransmitterV2.abi.Pack("enableAttester", newAttester)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackEnableAttester is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfae36879.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function enableAttester(address newAttester) returns()
func (messageTransmitterV2 *MessageTransmitterV2) TryPackEnableAttester(newAttester common.Address) ([]byte, error) {
	return messageTransmitterV2.abi.Pack("enableAttester", newAttester)
}

// PackGetEnabledAttester is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbeb673d8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getEnabledAttester(uint256 index) view returns(address)
func (messageTransmitterV2 *MessageTransmitterV2) PackGetEnabledAttester(index *big.Int) []byte {
	enc, err := messageTransmitterV2.abi.Pack("getEnabledAttester", index)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetEnabledAttester is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbeb673d8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getEnabledAttester(uint256 index) view returns(address)
func (messageTransmitterV2 *MessageTransmitterV2) TryPackGetEnabledAttester(index *big.Int) ([]byte, error) {
	return messageTransmitterV2.abi.Pack("getEnabledAttester", index)
}

// UnpackGetEnabledAttester is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbeb673d8.
//
// Solidity: function getEnabledAttester(uint256 index) view returns(address)
func (messageTransmitterV2 *MessageTransmitterV2) UnpackGetEnabledAttester(data []byte) (common.Address, error) {
	out, err := messageTransmitterV2.abi.Unpack("getEnabledAttester", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetNumEnabledAttesters is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x51079a53.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getNumEnabledAttesters() view returns(uint256)
func (messageTransmitterV2 *MessageTransmitterV2) PackGetNumEnabledAttesters() []byte {
	enc, err := messageTransmitterV2.abi.Pack("getNumEnabledAttesters")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetNumEnabledAttesters is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x51079a53.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getNumEnabledAttesters() view returns(uint256)
func (messageTransmitterV2 *MessageTransmitterV2) TryPackGetNumEnabledAttesters() ([]byte, error) {
	return messageTransmitterV2.abi.Pack("getNumEnabledAttesters")
}

// UnpackGetNumEnabledAttesters is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x51079a53.
//
// Solidity: function getNumEnabledAttesters() view returns(uint256)
func (messageTransmitterV2 *MessageTransmitterV2) UnpackGetNumEnabledAttesters(data []byte) (*big.Int, error) {
	out, err := messageTransmitterV2.abi.Unpack("getNumEnabledAttesters", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5f747e7d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function initialize(address owner_, address pauser_, address rescuer_, address attesterManager_, address[] attesters_, uint256 signatureThreshold_, uint256 maxMessageBodySize_) returns()
func (messageTransmitterV2 *MessageTransmitterV2) PackInitialize(owner common.Address, pauser common.Address, rescuer common.Address, attesterManager common.Address, attesters []common.Address, signatureThreshold *big.Int, maxMessageBodySize *big.Int) []byte {
	enc, err := messageTransmitterV2.abi.Pack("initialize", owner, pauser, rescuer, attesterManager, attesters, signatureThreshold, maxMessageBodySize)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5f747e7d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function initialize(address owner_, address pauser_, address rescuer_, address attesterManager_, address[] attesters_, uint256 signatureThreshold_, uint256 maxMessageBodySize_) returns()
func (messageTransmitterV2 *MessageTransmitterV2) TryPackInitialize(owner common.Address, pauser common.Address, rescuer common.Address, attesterManager common.Address, attesters []common.Address, signatureThreshold *big.Int, maxMessageBodySize *big.Int) ([]byte, error) {
	return messageTransmitterV2.abi.Pack("initialize", owner, pauser, rescuer, attesterManager, attesters, signatureThreshold, maxMessageBodySize)
}

// PackInitializedVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x08828eb7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function initializedVersion() view returns(uint64)
func (messageTransmitterV2 *MessageTransmitterV2) PackInitializedVersion() []byte {
	enc, err := messageTransmitterV2.abi.Pack("initializedVersion")
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
func (messageTransmitterV2 *MessageTransmitterV2) TryPackInitializedVersion() ([]byte, error) {
	return messageTransmitterV2.abi.Pack("initializedVersion")
}

// UnpackInitializedVersion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x08828eb7.
//
// Solidity: function initializedVersion() view returns(uint64)
func (messageTransmitterV2 *MessageTransmitterV2) UnpackInitializedVersion(data []byte) (uint64, error) {
	out, err := messageTransmitterV2.abi.Unpack("initializedVersion", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackIsEnabledAttester is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7af82f60.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isEnabledAttester(address attester) view returns(bool)
func (messageTransmitterV2 *MessageTransmitterV2) PackIsEnabledAttester(attester common.Address) []byte {
	enc, err := messageTransmitterV2.abi.Pack("isEnabledAttester", attester)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsEnabledAttester is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7af82f60.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isEnabledAttester(address attester) view returns(bool)
func (messageTransmitterV2 *MessageTransmitterV2) TryPackIsEnabledAttester(attester common.Address) ([]byte, error) {
	return messageTransmitterV2.abi.Pack("isEnabledAttester", attester)
}

// UnpackIsEnabledAttester is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7af82f60.
//
// Solidity: function isEnabledAttester(address attester) view returns(bool)
func (messageTransmitterV2 *MessageTransmitterV2) UnpackIsEnabledAttester(data []byte) (bool, error) {
	out, err := messageTransmitterV2.abi.Unpack("isEnabledAttester", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackLocalDomain is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8d3638f4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function localDomain() view returns(uint32)
func (messageTransmitterV2 *MessageTransmitterV2) PackLocalDomain() []byte {
	enc, err := messageTransmitterV2.abi.Pack("localDomain")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackLocalDomain is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8d3638f4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function localDomain() view returns(uint32)
func (messageTransmitterV2 *MessageTransmitterV2) TryPackLocalDomain() ([]byte, error) {
	return messageTransmitterV2.abi.Pack("localDomain")
}

// UnpackLocalDomain is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (messageTransmitterV2 *MessageTransmitterV2) UnpackLocalDomain(data []byte) (uint32, error) {
	out, err := messageTransmitterV2.abi.Unpack("localDomain", data)
	if err != nil {
		return *new(uint32), err
	}
	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	return out0, nil
}

// PackMaxMessageBodySize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaf47b9bb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function maxMessageBodySize() view returns(uint256)
func (messageTransmitterV2 *MessageTransmitterV2) PackMaxMessageBodySize() []byte {
	enc, err := messageTransmitterV2.abi.Pack("maxMessageBodySize")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMaxMessageBodySize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaf47b9bb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function maxMessageBodySize() view returns(uint256)
func (messageTransmitterV2 *MessageTransmitterV2) TryPackMaxMessageBodySize() ([]byte, error) {
	return messageTransmitterV2.abi.Pack("maxMessageBodySize")
}

// UnpackMaxMessageBodySize is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xaf47b9bb.
//
// Solidity: function maxMessageBodySize() view returns(uint256)
func (messageTransmitterV2 *MessageTransmitterV2) UnpackMaxMessageBodySize(data []byte) (*big.Int, error) {
	out, err := messageTransmitterV2.abi.Unpack("maxMessageBodySize", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function owner() view returns(address)
func (messageTransmitterV2 *MessageTransmitterV2) PackOwner() []byte {
	enc, err := messageTransmitterV2.abi.Pack("owner")
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
func (messageTransmitterV2 *MessageTransmitterV2) TryPackOwner() ([]byte, error) {
	return messageTransmitterV2.abi.Pack("owner")
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (messageTransmitterV2 *MessageTransmitterV2) UnpackOwner(data []byte) (common.Address, error) {
	out, err := messageTransmitterV2.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackPause is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8456cb59.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function pause() returns()
func (messageTransmitterV2 *MessageTransmitterV2) PackPause() []byte {
	enc, err := messageTransmitterV2.abi.Pack("pause")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPause is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8456cb59.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function pause() returns()
func (messageTransmitterV2 *MessageTransmitterV2) TryPackPause() ([]byte, error) {
	return messageTransmitterV2.abi.Pack("pause")
}

// PackPaused is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c975abb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function paused() view returns(bool)
func (messageTransmitterV2 *MessageTransmitterV2) PackPaused() []byte {
	enc, err := messageTransmitterV2.abi.Pack("paused")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPaused is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c975abb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function paused() view returns(bool)
func (messageTransmitterV2 *MessageTransmitterV2) TryPackPaused() ([]byte, error) {
	return messageTransmitterV2.abi.Pack("paused")
}

// UnpackPaused is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (messageTransmitterV2 *MessageTransmitterV2) UnpackPaused(data []byte) (bool, error) {
	out, err := messageTransmitterV2.abi.Unpack("paused", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackPauser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9fd0506d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function pauser() view returns(address)
func (messageTransmitterV2 *MessageTransmitterV2) PackPauser() []byte {
	enc, err := messageTransmitterV2.abi.Pack("pauser")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPauser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9fd0506d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function pauser() view returns(address)
func (messageTransmitterV2 *MessageTransmitterV2) TryPackPauser() ([]byte, error) {
	return messageTransmitterV2.abi.Pack("pauser")
}

// UnpackPauser is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (messageTransmitterV2 *MessageTransmitterV2) UnpackPauser(data []byte) (common.Address, error) {
	out, err := messageTransmitterV2.abi.Unpack("pauser", data)
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
func (messageTransmitterV2 *MessageTransmitterV2) PackPendingOwner() []byte {
	enc, err := messageTransmitterV2.abi.Pack("pendingOwner")
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
func (messageTransmitterV2 *MessageTransmitterV2) TryPackPendingOwner() ([]byte, error) {
	return messageTransmitterV2.abi.Pack("pendingOwner")
}

// UnpackPendingOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (messageTransmitterV2 *MessageTransmitterV2) UnpackPendingOwner(data []byte) (common.Address, error) {
	out, err := messageTransmitterV2.abi.Unpack("pendingOwner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackReceiveMessage is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x57ecfd28.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function receiveMessage(bytes message, bytes attestation) returns(bool success)
func (messageTransmitterV2 *MessageTransmitterV2) PackReceiveMessage(message []byte, attestation []byte) []byte {
	enc, err := messageTransmitterV2.abi.Pack("receiveMessage", message, attestation)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackReceiveMessage is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x57ecfd28.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function receiveMessage(bytes message, bytes attestation) returns(bool success)
func (messageTransmitterV2 *MessageTransmitterV2) TryPackReceiveMessage(message []byte, attestation []byte) ([]byte, error) {
	return messageTransmitterV2.abi.Pack("receiveMessage", message, attestation)
}

// UnpackReceiveMessage is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x57ecfd28.
//
// Solidity: function receiveMessage(bytes message, bytes attestation) returns(bool success)
func (messageTransmitterV2 *MessageTransmitterV2) UnpackReceiveMessage(data []byte) (bool, error) {
	out, err := messageTransmitterV2.abi.Unpack("receiveMessage", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackRescueERC20 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb2118a8d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function rescueERC20(address tokenContract, address to, uint256 amount) returns()
func (messageTransmitterV2 *MessageTransmitterV2) PackRescueERC20(tokenContract common.Address, to common.Address, amount *big.Int) []byte {
	enc, err := messageTransmitterV2.abi.Pack("rescueERC20", tokenContract, to, amount)
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
func (messageTransmitterV2 *MessageTransmitterV2) TryPackRescueERC20(tokenContract common.Address, to common.Address, amount *big.Int) ([]byte, error) {
	return messageTransmitterV2.abi.Pack("rescueERC20", tokenContract, to, amount)
}

// PackRescuer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x38a63183.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function rescuer() view returns(address)
func (messageTransmitterV2 *MessageTransmitterV2) PackRescuer() []byte {
	enc, err := messageTransmitterV2.abi.Pack("rescuer")
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
func (messageTransmitterV2 *MessageTransmitterV2) TryPackRescuer() ([]byte, error) {
	return messageTransmitterV2.abi.Pack("rescuer")
}

// UnpackRescuer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x38a63183.
//
// Solidity: function rescuer() view returns(address)
func (messageTransmitterV2 *MessageTransmitterV2) UnpackRescuer(data []byte) (common.Address, error) {
	out, err := messageTransmitterV2.abi.Unpack("rescuer", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackSendMessage is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x14b157ab.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function sendMessage(uint32 destinationDomain, bytes32 recipient, bytes32 destinationCaller, uint32 minFinalityThreshold, bytes messageBody) returns()
func (messageTransmitterV2 *MessageTransmitterV2) PackSendMessage(destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, minFinalityThreshold uint32, messageBody []byte) []byte {
	enc, err := messageTransmitterV2.abi.Pack("sendMessage", destinationDomain, recipient, destinationCaller, minFinalityThreshold, messageBody)
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
func (messageTransmitterV2 *MessageTransmitterV2) TryPackSendMessage(destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, minFinalityThreshold uint32, messageBody []byte) ([]byte, error) {
	return messageTransmitterV2.abi.Pack("sendMessage", destinationDomain, recipient, destinationCaller, minFinalityThreshold, messageBody)
}

// PackSetMaxMessageBodySize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x92492c68.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setMaxMessageBodySize(uint256 newMaxMessageBodySize) returns()
func (messageTransmitterV2 *MessageTransmitterV2) PackSetMaxMessageBodySize(newMaxMessageBodySize *big.Int) []byte {
	enc, err := messageTransmitterV2.abi.Pack("setMaxMessageBodySize", newMaxMessageBodySize)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetMaxMessageBodySize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x92492c68.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setMaxMessageBodySize(uint256 newMaxMessageBodySize) returns()
func (messageTransmitterV2 *MessageTransmitterV2) TryPackSetMaxMessageBodySize(newMaxMessageBodySize *big.Int) ([]byte, error) {
	return messageTransmitterV2.abi.Pack("setMaxMessageBodySize", newMaxMessageBodySize)
}

// PackSetSignatureThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbbde5374.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setSignatureThreshold(uint256 newSignatureThreshold) returns()
func (messageTransmitterV2 *MessageTransmitterV2) PackSetSignatureThreshold(newSignatureThreshold *big.Int) []byte {
	enc, err := messageTransmitterV2.abi.Pack("setSignatureThreshold", newSignatureThreshold)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetSignatureThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbbde5374.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setSignatureThreshold(uint256 newSignatureThreshold) returns()
func (messageTransmitterV2 *MessageTransmitterV2) TryPackSetSignatureThreshold(newSignatureThreshold *big.Int) ([]byte, error) {
	return messageTransmitterV2.abi.Pack("setSignatureThreshold", newSignatureThreshold)
}

// PackSignatureThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa82f2e26.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function signatureThreshold() view returns(uint256)
func (messageTransmitterV2 *MessageTransmitterV2) PackSignatureThreshold() []byte {
	enc, err := messageTransmitterV2.abi.Pack("signatureThreshold")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSignatureThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa82f2e26.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function signatureThreshold() view returns(uint256)
func (messageTransmitterV2 *MessageTransmitterV2) TryPackSignatureThreshold() ([]byte, error) {
	return messageTransmitterV2.abi.Pack("signatureThreshold")
}

// UnpackSignatureThreshold is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa82f2e26.
//
// Solidity: function signatureThreshold() view returns(uint256)
func (messageTransmitterV2 *MessageTransmitterV2) UnpackSignatureThreshold(data []byte) (*big.Int, error) {
	out, err := messageTransmitterV2.abi.Unpack("signatureThreshold", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (messageTransmitterV2 *MessageTransmitterV2) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := messageTransmitterV2.abi.Pack("transferOwnership", newOwner)
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
func (messageTransmitterV2 *MessageTransmitterV2) TryPackTransferOwnership(newOwner common.Address) ([]byte, error) {
	return messageTransmitterV2.abi.Pack("transferOwnership", newOwner)
}

// PackUnpause is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3f4ba83a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function unpause() returns()
func (messageTransmitterV2 *MessageTransmitterV2) PackUnpause() []byte {
	enc, err := messageTransmitterV2.abi.Pack("unpause")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUnpause is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3f4ba83a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function unpause() returns()
func (messageTransmitterV2 *MessageTransmitterV2) TryPackUnpause() ([]byte, error) {
	return messageTransmitterV2.abi.Pack("unpause")
}

// PackUpdateAttesterManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xde7769d4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updateAttesterManager(address newAttesterManager) returns()
func (messageTransmitterV2 *MessageTransmitterV2) PackUpdateAttesterManager(newAttesterManager common.Address) []byte {
	enc, err := messageTransmitterV2.abi.Pack("updateAttesterManager", newAttesterManager)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUpdateAttesterManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xde7769d4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function updateAttesterManager(address newAttesterManager) returns()
func (messageTransmitterV2 *MessageTransmitterV2) TryPackUpdateAttesterManager(newAttesterManager common.Address) ([]byte, error) {
	return messageTransmitterV2.abi.Pack("updateAttesterManager", newAttesterManager)
}

// PackUpdatePauser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x554bab3c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updatePauser(address _newPauser) returns()
func (messageTransmitterV2 *MessageTransmitterV2) PackUpdatePauser(newPauser common.Address) []byte {
	enc, err := messageTransmitterV2.abi.Pack("updatePauser", newPauser)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUpdatePauser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x554bab3c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function updatePauser(address _newPauser) returns()
func (messageTransmitterV2 *MessageTransmitterV2) TryPackUpdatePauser(newPauser common.Address) ([]byte, error) {
	return messageTransmitterV2.abi.Pack("updatePauser", newPauser)
}

// PackUpdateRescuer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2ab60045.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updateRescuer(address newRescuer) returns()
func (messageTransmitterV2 *MessageTransmitterV2) PackUpdateRescuer(newRescuer common.Address) []byte {
	enc, err := messageTransmitterV2.abi.Pack("updateRescuer", newRescuer)
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
func (messageTransmitterV2 *MessageTransmitterV2) TryPackUpdateRescuer(newRescuer common.Address) ([]byte, error) {
	return messageTransmitterV2.abi.Pack("updateRescuer", newRescuer)
}

// PackUsedNonces is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfeb61724.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function usedNonces(bytes32 ) view returns(uint256)
func (messageTransmitterV2 *MessageTransmitterV2) PackUsedNonces(arg0 [32]byte) []byte {
	enc, err := messageTransmitterV2.abi.Pack("usedNonces", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUsedNonces is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfeb61724.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function usedNonces(bytes32 ) view returns(uint256)
func (messageTransmitterV2 *MessageTransmitterV2) TryPackUsedNonces(arg0 [32]byte) ([]byte, error) {
	return messageTransmitterV2.abi.Pack("usedNonces", arg0)
}

// UnpackUsedNonces is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfeb61724.
//
// Solidity: function usedNonces(bytes32 ) view returns(uint256)
func (messageTransmitterV2 *MessageTransmitterV2) UnpackUsedNonces(data []byte) (*big.Int, error) {
	out, err := messageTransmitterV2.abi.Unpack("usedNonces", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x54fd4d50.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function version() view returns(uint32)
func (messageTransmitterV2 *MessageTransmitterV2) PackVersion() []byte {
	enc, err := messageTransmitterV2.abi.Pack("version")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x54fd4d50.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function version() view returns(uint32)
func (messageTransmitterV2 *MessageTransmitterV2) TryPackVersion() ([]byte, error) {
	return messageTransmitterV2.abi.Pack("version")
}

// UnpackVersion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x54fd4d50.
//
// Solidity: function version() view returns(uint32)
func (messageTransmitterV2 *MessageTransmitterV2) UnpackVersion(data []byte) (uint32, error) {
	out, err := messageTransmitterV2.abi.Unpack("version", data)
	if err != nil {
		return *new(uint32), err
	}
	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	return out0, nil
}

// MessageTransmitterV2AttesterDisabled represents a AttesterDisabled event raised by the MessageTransmitterV2 contract.
type MessageTransmitterV2AttesterDisabled struct {
	Attester common.Address
	Raw      *types.Log // Blockchain specific contextual infos
}

const MessageTransmitterV2AttesterDisabledEventName = "AttesterDisabled"

// ContractEventName returns the user-defined event name.
func (MessageTransmitterV2AttesterDisabled) ContractEventName() string {
	return MessageTransmitterV2AttesterDisabledEventName
}

// UnpackAttesterDisabledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AttesterDisabled(address indexed attester)
func (messageTransmitterV2 *MessageTransmitterV2) UnpackAttesterDisabledEvent(log *types.Log) (*MessageTransmitterV2AttesterDisabled, error) {
	event := "AttesterDisabled"
	if len(log.Topics) == 0 || log.Topics[0] != messageTransmitterV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MessageTransmitterV2AttesterDisabled)
	if len(log.Data) > 0 {
		if err := messageTransmitterV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range messageTransmitterV2.abi.Events[event].Inputs {
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

// MessageTransmitterV2AttesterEnabled represents a AttesterEnabled event raised by the MessageTransmitterV2 contract.
type MessageTransmitterV2AttesterEnabled struct {
	Attester common.Address
	Raw      *types.Log // Blockchain specific contextual infos
}

const MessageTransmitterV2AttesterEnabledEventName = "AttesterEnabled"

// ContractEventName returns the user-defined event name.
func (MessageTransmitterV2AttesterEnabled) ContractEventName() string {
	return MessageTransmitterV2AttesterEnabledEventName
}

// UnpackAttesterEnabledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AttesterEnabled(address indexed attester)
func (messageTransmitterV2 *MessageTransmitterV2) UnpackAttesterEnabledEvent(log *types.Log) (*MessageTransmitterV2AttesterEnabled, error) {
	event := "AttesterEnabled"
	if len(log.Topics) == 0 || log.Topics[0] != messageTransmitterV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MessageTransmitterV2AttesterEnabled)
	if len(log.Data) > 0 {
		if err := messageTransmitterV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range messageTransmitterV2.abi.Events[event].Inputs {
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

// MessageTransmitterV2AttesterManagerUpdated represents a AttesterManagerUpdated event raised by the MessageTransmitterV2 contract.
type MessageTransmitterV2AttesterManagerUpdated struct {
	PreviousAttesterManager common.Address
	NewAttesterManager      common.Address
	Raw                     *types.Log // Blockchain specific contextual infos
}

const MessageTransmitterV2AttesterManagerUpdatedEventName = "AttesterManagerUpdated"

// ContractEventName returns the user-defined event name.
func (MessageTransmitterV2AttesterManagerUpdated) ContractEventName() string {
	return MessageTransmitterV2AttesterManagerUpdatedEventName
}

// UnpackAttesterManagerUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AttesterManagerUpdated(address indexed previousAttesterManager, address indexed newAttesterManager)
func (messageTransmitterV2 *MessageTransmitterV2) UnpackAttesterManagerUpdatedEvent(log *types.Log) (*MessageTransmitterV2AttesterManagerUpdated, error) {
	event := "AttesterManagerUpdated"
	if len(log.Topics) == 0 || log.Topics[0] != messageTransmitterV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MessageTransmitterV2AttesterManagerUpdated)
	if len(log.Data) > 0 {
		if err := messageTransmitterV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range messageTransmitterV2.abi.Events[event].Inputs {
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

// MessageTransmitterV2Initialized represents a Initialized event raised by the MessageTransmitterV2 contract.
type MessageTransmitterV2Initialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const MessageTransmitterV2InitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (MessageTransmitterV2Initialized) ContractEventName() string {
	return MessageTransmitterV2InitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (messageTransmitterV2 *MessageTransmitterV2) UnpackInitializedEvent(log *types.Log) (*MessageTransmitterV2Initialized, error) {
	event := "Initialized"
	if len(log.Topics) == 0 || log.Topics[0] != messageTransmitterV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MessageTransmitterV2Initialized)
	if len(log.Data) > 0 {
		if err := messageTransmitterV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range messageTransmitterV2.abi.Events[event].Inputs {
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

// MessageTransmitterV2MaxMessageBodySizeUpdated represents a MaxMessageBodySizeUpdated event raised by the MessageTransmitterV2 contract.
type MessageTransmitterV2MaxMessageBodySizeUpdated struct {
	NewMaxMessageBodySize *big.Int
	Raw                   *types.Log // Blockchain specific contextual infos
}

const MessageTransmitterV2MaxMessageBodySizeUpdatedEventName = "MaxMessageBodySizeUpdated"

// ContractEventName returns the user-defined event name.
func (MessageTransmitterV2MaxMessageBodySizeUpdated) ContractEventName() string {
	return MessageTransmitterV2MaxMessageBodySizeUpdatedEventName
}

// UnpackMaxMessageBodySizeUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MaxMessageBodySizeUpdated(uint256 newMaxMessageBodySize)
func (messageTransmitterV2 *MessageTransmitterV2) UnpackMaxMessageBodySizeUpdatedEvent(log *types.Log) (*MessageTransmitterV2MaxMessageBodySizeUpdated, error) {
	event := "MaxMessageBodySizeUpdated"
	if len(log.Topics) == 0 || log.Topics[0] != messageTransmitterV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MessageTransmitterV2MaxMessageBodySizeUpdated)
	if len(log.Data) > 0 {
		if err := messageTransmitterV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range messageTransmitterV2.abi.Events[event].Inputs {
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

// MessageTransmitterV2MessageReceived represents a MessageReceived event raised by the MessageTransmitterV2 contract.
type MessageTransmitterV2MessageReceived struct {
	Caller                    common.Address
	SourceDomain              uint32
	Nonce                     [32]byte
	Sender                    [32]byte
	FinalityThresholdExecuted uint32
	MessageBody               []byte
	Raw                       *types.Log // Blockchain specific contextual infos
}

const MessageTransmitterV2MessageReceivedEventName = "MessageReceived"

// ContractEventName returns the user-defined event name.
func (MessageTransmitterV2MessageReceived) ContractEventName() string {
	return MessageTransmitterV2MessageReceivedEventName
}

// UnpackMessageReceivedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MessageReceived(address indexed caller, uint32 sourceDomain, bytes32 indexed nonce, bytes32 sender, uint32 indexed finalityThresholdExecuted, bytes messageBody)
func (messageTransmitterV2 *MessageTransmitterV2) UnpackMessageReceivedEvent(log *types.Log) (*MessageTransmitterV2MessageReceived, error) {
	event := "MessageReceived"
	if len(log.Topics) == 0 || log.Topics[0] != messageTransmitterV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MessageTransmitterV2MessageReceived)
	if len(log.Data) > 0 {
		if err := messageTransmitterV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range messageTransmitterV2.abi.Events[event].Inputs {
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

// MessageTransmitterV2MessageSent represents a MessageSent event raised by the MessageTransmitterV2 contract.
type MessageTransmitterV2MessageSent struct {
	Message []byte
	Raw     *types.Log // Blockchain specific contextual infos
}

const MessageTransmitterV2MessageSentEventName = "MessageSent"

// ContractEventName returns the user-defined event name.
func (MessageTransmitterV2MessageSent) ContractEventName() string {
	return MessageTransmitterV2MessageSentEventName
}

// UnpackMessageSentEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MessageSent(bytes message)
func (messageTransmitterV2 *MessageTransmitterV2) UnpackMessageSentEvent(log *types.Log) (*MessageTransmitterV2MessageSent, error) {
	event := "MessageSent"
	if len(log.Topics) == 0 || log.Topics[0] != messageTransmitterV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MessageTransmitterV2MessageSent)
	if len(log.Data) > 0 {
		if err := messageTransmitterV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range messageTransmitterV2.abi.Events[event].Inputs {
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

// MessageTransmitterV2OwnershipTransferStarted represents a OwnershipTransferStarted event raised by the MessageTransmitterV2 contract.
type MessageTransmitterV2OwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const MessageTransmitterV2OwnershipTransferStartedEventName = "OwnershipTransferStarted"

// ContractEventName returns the user-defined event name.
func (MessageTransmitterV2OwnershipTransferStarted) ContractEventName() string {
	return MessageTransmitterV2OwnershipTransferStartedEventName
}

// UnpackOwnershipTransferStartedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (messageTransmitterV2 *MessageTransmitterV2) UnpackOwnershipTransferStartedEvent(log *types.Log) (*MessageTransmitterV2OwnershipTransferStarted, error) {
	event := "OwnershipTransferStarted"
	if len(log.Topics) == 0 || log.Topics[0] != messageTransmitterV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MessageTransmitterV2OwnershipTransferStarted)
	if len(log.Data) > 0 {
		if err := messageTransmitterV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range messageTransmitterV2.abi.Events[event].Inputs {
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

// MessageTransmitterV2OwnershipTransferred represents a OwnershipTransferred event raised by the MessageTransmitterV2 contract.
type MessageTransmitterV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const MessageTransmitterV2OwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (MessageTransmitterV2OwnershipTransferred) ContractEventName() string {
	return MessageTransmitterV2OwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (messageTransmitterV2 *MessageTransmitterV2) UnpackOwnershipTransferredEvent(log *types.Log) (*MessageTransmitterV2OwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if len(log.Topics) == 0 || log.Topics[0] != messageTransmitterV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MessageTransmitterV2OwnershipTransferred)
	if len(log.Data) > 0 {
		if err := messageTransmitterV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range messageTransmitterV2.abi.Events[event].Inputs {
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

// MessageTransmitterV2Pause represents a Pause event raised by the MessageTransmitterV2 contract.
type MessageTransmitterV2Pause struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const MessageTransmitterV2PauseEventName = "Pause"

// ContractEventName returns the user-defined event name.
func (MessageTransmitterV2Pause) ContractEventName() string {
	return MessageTransmitterV2PauseEventName
}

// UnpackPauseEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Pause()
func (messageTransmitterV2 *MessageTransmitterV2) UnpackPauseEvent(log *types.Log) (*MessageTransmitterV2Pause, error) {
	event := "Pause"
	if len(log.Topics) == 0 || log.Topics[0] != messageTransmitterV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MessageTransmitterV2Pause)
	if len(log.Data) > 0 {
		if err := messageTransmitterV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range messageTransmitterV2.abi.Events[event].Inputs {
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

// MessageTransmitterV2PauserChanged represents a PauserChanged event raised by the MessageTransmitterV2 contract.
type MessageTransmitterV2PauserChanged struct {
	NewAddress common.Address
	Raw        *types.Log // Blockchain specific contextual infos
}

const MessageTransmitterV2PauserChangedEventName = "PauserChanged"

// ContractEventName returns the user-defined event name.
func (MessageTransmitterV2PauserChanged) ContractEventName() string {
	return MessageTransmitterV2PauserChangedEventName
}

// UnpackPauserChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PauserChanged(address indexed newAddress)
func (messageTransmitterV2 *MessageTransmitterV2) UnpackPauserChangedEvent(log *types.Log) (*MessageTransmitterV2PauserChanged, error) {
	event := "PauserChanged"
	if len(log.Topics) == 0 || log.Topics[0] != messageTransmitterV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MessageTransmitterV2PauserChanged)
	if len(log.Data) > 0 {
		if err := messageTransmitterV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range messageTransmitterV2.abi.Events[event].Inputs {
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

// MessageTransmitterV2RescuerChanged represents a RescuerChanged event raised by the MessageTransmitterV2 contract.
type MessageTransmitterV2RescuerChanged struct {
	NewRescuer common.Address
	Raw        *types.Log // Blockchain specific contextual infos
}

const MessageTransmitterV2RescuerChangedEventName = "RescuerChanged"

// ContractEventName returns the user-defined event name.
func (MessageTransmitterV2RescuerChanged) ContractEventName() string {
	return MessageTransmitterV2RescuerChangedEventName
}

// UnpackRescuerChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RescuerChanged(address indexed newRescuer)
func (messageTransmitterV2 *MessageTransmitterV2) UnpackRescuerChangedEvent(log *types.Log) (*MessageTransmitterV2RescuerChanged, error) {
	event := "RescuerChanged"
	if len(log.Topics) == 0 || log.Topics[0] != messageTransmitterV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MessageTransmitterV2RescuerChanged)
	if len(log.Data) > 0 {
		if err := messageTransmitterV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range messageTransmitterV2.abi.Events[event].Inputs {
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

// MessageTransmitterV2SignatureThresholdUpdated represents a SignatureThresholdUpdated event raised by the MessageTransmitterV2 contract.
type MessageTransmitterV2SignatureThresholdUpdated struct {
	OldSignatureThreshold *big.Int
	NewSignatureThreshold *big.Int
	Raw                   *types.Log // Blockchain specific contextual infos
}

const MessageTransmitterV2SignatureThresholdUpdatedEventName = "SignatureThresholdUpdated"

// ContractEventName returns the user-defined event name.
func (MessageTransmitterV2SignatureThresholdUpdated) ContractEventName() string {
	return MessageTransmitterV2SignatureThresholdUpdatedEventName
}

// UnpackSignatureThresholdUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SignatureThresholdUpdated(uint256 oldSignatureThreshold, uint256 newSignatureThreshold)
func (messageTransmitterV2 *MessageTransmitterV2) UnpackSignatureThresholdUpdatedEvent(log *types.Log) (*MessageTransmitterV2SignatureThresholdUpdated, error) {
	event := "SignatureThresholdUpdated"
	if len(log.Topics) == 0 || log.Topics[0] != messageTransmitterV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MessageTransmitterV2SignatureThresholdUpdated)
	if len(log.Data) > 0 {
		if err := messageTransmitterV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range messageTransmitterV2.abi.Events[event].Inputs {
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

// MessageTransmitterV2Unpause represents a Unpause event raised by the MessageTransmitterV2 contract.
type MessageTransmitterV2Unpause struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const MessageTransmitterV2UnpauseEventName = "Unpause"

// ContractEventName returns the user-defined event name.
func (MessageTransmitterV2Unpause) ContractEventName() string {
	return MessageTransmitterV2UnpauseEventName
}

// UnpackUnpauseEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Unpause()
func (messageTransmitterV2 *MessageTransmitterV2) UnpackUnpauseEvent(log *types.Log) (*MessageTransmitterV2Unpause, error) {
	event := "Unpause"
	if len(log.Topics) == 0 || log.Topics[0] != messageTransmitterV2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MessageTransmitterV2Unpause)
	if len(log.Data) > 0 {
		if err := messageTransmitterV2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range messageTransmitterV2.abi.Events[event].Inputs {
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

// MessageV2MetaData contains all meta data concerning the MessageV2 contract.
var MessageV2MetaData = bind.MetaData{
	ABI: "[]",
	ID:  "074256211ec4313b9462c04e13b87fb77b",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212205faa5375926d8652a981c0d794e5bfbdefd14aa30bb1bc970da4ea34955d27ae64736f6c63430007060033",
}

// MessageV2 is an auto generated Go binding around an Ethereum contract.
type MessageV2 struct {
	abi abi.ABI
}

// NewMessageV2 creates a new instance of MessageV2.
func NewMessageV2() *MessageV2 {
	parsed, err := MessageV2MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &MessageV2{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *MessageV2) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
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

// PausableMetaData contains all meta data concerning the Pausable contract.
var PausableMetaData = bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"PauserChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newPauser\",\"type\":\"address\"}],\"name\":\"updatePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "5d4add9cf6040c733f44d451520b3ebcdd",
	Bin: "0x60806040526000600260146101000a81548160ff02191690831515021790555034801561002b57600080fd5b5061004861003d61004d60201b60201c565b61005560201b60201c565b610154565b600033905090565b600160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905561008d8161009060201b6105cc1760201c565b50565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b610952806101636000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80638456cb59116100665780638456cb59146101105780638da5cb5b1461011a5780639fd0506d1461014e578063e30c397814610182578063f2fde38b146101b657610093565b80633f4ba83a14610098578063554bab3c146100a25780635c975abb146100e657806379ba509714610106575b600080fd5b6100a06101fa565b005b6100e4600480360360208110156100b857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506102e9565b005b6100ee6102fd565b60405180821515815260200191505060405180910390f35b61010e610310565b005b6101186103b3565b005b6101226104a2565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101566104cb565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61018a6104f5565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101f8600480360360208110156101cc57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061051f565b005b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146102a0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806108fb6022913960400191505060405180910390fd5b6000600260146101000a81548160ff0219169083151502179055507f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a1565b6102f1610690565b6102fa81610741565b50565b600260149054906101000a900460ff1681565b600061031a610870565b90508073ffffffffffffffffffffffffffffffffffffffff1661033b6104f5565b73ffffffffffffffffffffffffffffffffffffffff16146103a7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260298152602001806108d26029913960400191505060405180910390fd5b6103b081610878565b50565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610459576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806108fb6022913960400191505060405180910390fd5b6001600260146101000a81548160ff0219169083151502179055507f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610527610690565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff166105876104a2565b73ffffffffffffffffffffffffffffffffffffffff167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a350565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b610698610870565b73ffffffffffffffffffffffffffffffffffffffff166106b66104a2565b73ffffffffffffffffffffffffffffffffffffffff161461073f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156107c7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260288152602001806108aa6028913960400191505060405180910390fd5b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fb80482a293ca2e013eda8683c9bd7fc8347cfdaeea5ede58cba46df502c2a60460405160405180910390a250565b600033905090565b600160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556108a6816105cc565b5056fe5061757361626c653a206e65772070617573657220697320746865207a65726f20616464726573734f776e61626c6532537465703a2063616c6c6572206973206e6f7420746865206e6577206f776e65725061757361626c653a2063616c6c6572206973206e6f742074686520706175736572a264697066735822122031bfd474d57cc3df03942066b0652f9e20d6d9e0ec9aebfd4ad8f609b507806a64736f6c63430007060033",
}

// Pausable is an auto generated Go binding around an Ethereum contract.
type Pausable struct {
	abi abi.ABI
}

// NewPausable creates a new instance of Pausable.
func NewPausable() *Pausable {
	parsed, err := PausableMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Pausable{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Pausable) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackAcceptOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x79ba5097.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function acceptOwnership() returns()
func (pausable *Pausable) PackAcceptOwnership() []byte {
	enc, err := pausable.abi.Pack("acceptOwnership")
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
func (pausable *Pausable) TryPackAcceptOwnership() ([]byte, error) {
	return pausable.abi.Pack("acceptOwnership")
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function owner() view returns(address)
func (pausable *Pausable) PackOwner() []byte {
	enc, err := pausable.abi.Pack("owner")
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
func (pausable *Pausable) TryPackOwner() ([]byte, error) {
	return pausable.abi.Pack("owner")
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (pausable *Pausable) UnpackOwner(data []byte) (common.Address, error) {
	out, err := pausable.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackPause is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8456cb59.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function pause() returns()
func (pausable *Pausable) PackPause() []byte {
	enc, err := pausable.abi.Pack("pause")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPause is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8456cb59.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function pause() returns()
func (pausable *Pausable) TryPackPause() ([]byte, error) {
	return pausable.abi.Pack("pause")
}

// PackPaused is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c975abb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function paused() view returns(bool)
func (pausable *Pausable) PackPaused() []byte {
	enc, err := pausable.abi.Pack("paused")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPaused is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c975abb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function paused() view returns(bool)
func (pausable *Pausable) TryPackPaused() ([]byte, error) {
	return pausable.abi.Pack("paused")
}

// UnpackPaused is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (pausable *Pausable) UnpackPaused(data []byte) (bool, error) {
	out, err := pausable.abi.Unpack("paused", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackPauser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9fd0506d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function pauser() view returns(address)
func (pausable *Pausable) PackPauser() []byte {
	enc, err := pausable.abi.Pack("pauser")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPauser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9fd0506d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function pauser() view returns(address)
func (pausable *Pausable) TryPackPauser() ([]byte, error) {
	return pausable.abi.Pack("pauser")
}

// UnpackPauser is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (pausable *Pausable) UnpackPauser(data []byte) (common.Address, error) {
	out, err := pausable.abi.Unpack("pauser", data)
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
func (pausable *Pausable) PackPendingOwner() []byte {
	enc, err := pausable.abi.Pack("pendingOwner")
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
func (pausable *Pausable) TryPackPendingOwner() ([]byte, error) {
	return pausable.abi.Pack("pendingOwner")
}

// UnpackPendingOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (pausable *Pausable) UnpackPendingOwner(data []byte) (common.Address, error) {
	out, err := pausable.abi.Unpack("pendingOwner", data)
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
func (pausable *Pausable) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := pausable.abi.Pack("transferOwnership", newOwner)
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
func (pausable *Pausable) TryPackTransferOwnership(newOwner common.Address) ([]byte, error) {
	return pausable.abi.Pack("transferOwnership", newOwner)
}

// PackUnpause is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3f4ba83a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function unpause() returns()
func (pausable *Pausable) PackUnpause() []byte {
	enc, err := pausable.abi.Pack("unpause")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUnpause is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3f4ba83a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function unpause() returns()
func (pausable *Pausable) TryPackUnpause() ([]byte, error) {
	return pausable.abi.Pack("unpause")
}

// PackUpdatePauser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x554bab3c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updatePauser(address _newPauser) returns()
func (pausable *Pausable) PackUpdatePauser(newPauser common.Address) []byte {
	enc, err := pausable.abi.Pack("updatePauser", newPauser)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUpdatePauser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x554bab3c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function updatePauser(address _newPauser) returns()
func (pausable *Pausable) TryPackUpdatePauser(newPauser common.Address) ([]byte, error) {
	return pausable.abi.Pack("updatePauser", newPauser)
}

// PausableOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the Pausable contract.
type PausableOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const PausableOwnershipTransferStartedEventName = "OwnershipTransferStarted"

// ContractEventName returns the user-defined event name.
func (PausableOwnershipTransferStarted) ContractEventName() string {
	return PausableOwnershipTransferStartedEventName
}

// UnpackOwnershipTransferStartedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (pausable *Pausable) UnpackOwnershipTransferStartedEvent(log *types.Log) (*PausableOwnershipTransferStarted, error) {
	event := "OwnershipTransferStarted"
	if len(log.Topics) == 0 || log.Topics[0] != pausable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PausableOwnershipTransferStarted)
	if len(log.Data) > 0 {
		if err := pausable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range pausable.abi.Events[event].Inputs {
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

// PausableOwnershipTransferred represents a OwnershipTransferred event raised by the Pausable contract.
type PausableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const PausableOwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (PausableOwnershipTransferred) ContractEventName() string {
	return PausableOwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (pausable *Pausable) UnpackOwnershipTransferredEvent(log *types.Log) (*PausableOwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if len(log.Topics) == 0 || log.Topics[0] != pausable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PausableOwnershipTransferred)
	if len(log.Data) > 0 {
		if err := pausable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range pausable.abi.Events[event].Inputs {
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

// PausablePause represents a Pause event raised by the Pausable contract.
type PausablePause struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const PausablePauseEventName = "Pause"

// ContractEventName returns the user-defined event name.
func (PausablePause) ContractEventName() string {
	return PausablePauseEventName
}

// UnpackPauseEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Pause()
func (pausable *Pausable) UnpackPauseEvent(log *types.Log) (*PausablePause, error) {
	event := "Pause"
	if len(log.Topics) == 0 || log.Topics[0] != pausable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PausablePause)
	if len(log.Data) > 0 {
		if err := pausable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range pausable.abi.Events[event].Inputs {
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

// PausablePauserChanged represents a PauserChanged event raised by the Pausable contract.
type PausablePauserChanged struct {
	NewAddress common.Address
	Raw        *types.Log // Blockchain specific contextual infos
}

const PausablePauserChangedEventName = "PauserChanged"

// ContractEventName returns the user-defined event name.
func (PausablePauserChanged) ContractEventName() string {
	return PausablePauserChangedEventName
}

// UnpackPauserChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PauserChanged(address indexed newAddress)
func (pausable *Pausable) UnpackPauserChangedEvent(log *types.Log) (*PausablePauserChanged, error) {
	event := "PauserChanged"
	if len(log.Topics) == 0 || log.Topics[0] != pausable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PausablePauserChanged)
	if len(log.Data) > 0 {
		if err := pausable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range pausable.abi.Events[event].Inputs {
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

// PausableUnpause represents a Unpause event raised by the Pausable contract.
type PausableUnpause struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const PausableUnpauseEventName = "Unpause"

// ContractEventName returns the user-defined event name.
func (PausableUnpause) ContractEventName() string {
	return PausableUnpauseEventName
}

// UnpackUnpauseEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Unpause()
func (pausable *Pausable) UnpackUnpauseEvent(log *types.Log) (*PausableUnpause, error) {
	event := "Unpause"
	if len(log.Topics) == 0 || log.Topics[0] != pausable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PausableUnpause)
	if len(log.Data) > 0 {
		if err := pausable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range pausable.abi.Events[event].Inputs {
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
