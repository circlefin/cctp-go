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

