// Copyright (C) 2026 Circle Internet Group, Inc.
// This file is part of the cctp-go project.

// The cctp-go project is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The cctp-go project is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the cctp-go project. If not, see <http://www.gnu.org/licenses/>.

package testutil

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Test constants
const (
	TestPrivateKey = "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"
	TestAddress    = "0xFCAd0B19bB29D4674531d6f115237E16AfCE377c" // Derived from TestPrivateKey
)

// Common test addresses
var (
	TestUSDCAddress           = common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
	TestTokenMessengerAddress = common.HexToAddress("0x28b5a0e9C621a5BadaA536219b3a228C8168cf5d")
	TestRecipientAddress      = common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb7")
)

// EncodeBigInt encodes a big.Int as 32 bytes for ABI encoding
func EncodeBigInt(value *big.Int) []byte {
	result := make([]byte, 32)
	valueBytes := value.Bytes()
	copy(result[32-len(valueBytes):], valueBytes)
	return result
}

// EncodeAddress encodes an address as 32 bytes for ABI encoding
func EncodeAddress(addr common.Address) []byte {
	result := make([]byte, 32)
	copy(result[12:], addr.Bytes())
	return result
}

// StandardUSDCAmount returns standard test amounts in USDC (6 decimals)
func StandardUSDCAmount() *big.Int {
	return big.NewInt(100000000) // 100 USDC
}

// SmallUSDCAmount returns a small test amount
func SmallUSDCAmount() *big.Int {
	return big.NewInt(1000000) // 1 USDC
}

// LargeUSDCAmount returns a large test amount
func LargeUSDCAmount() *big.Int {
	return big.NewInt(1000000000000) // 1 million USDC
}
