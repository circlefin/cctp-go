package util

import "github.com/ethereum/go-ethereum/common"

// AddressToBytes32 converts an Ethereum address to bytes32 format for CCTP
// The address is right-padded with 12 zero bytes to create a 32-byte value
func AddressToBytes32(addr common.Address) [32]byte {
	var bytes32 [32]byte
	copy(bytes32[12:], addr.Bytes())
	return bytes32
}

// Bytes32ToAddress converts bytes32 back to an Ethereum address
// Extracts the last 20 bytes from the 32-byte value
func Bytes32ToAddress(bytes32 [32]byte) common.Address {
	return common.BytesToAddress(bytes32[12:])
}

