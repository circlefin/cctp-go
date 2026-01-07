// Copyright 2025 Circle Internet Group, Inc.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

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

