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

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestAddressToBytes32(t *testing.T) {
	testCases := []struct {
		name     string
		address  string
		expected [32]byte
	}{
		{
			name:    "Zero address",
			address: "0x0000000000000000000000000000000000000000",
			expected: [32]byte{
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
		},
		{
			name:    "Non-zero address",
			address: "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb7",
			expected: [32]byte{
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x74, 0x2d, 0x35, 0xcc,
				0x66, 0x34, 0xc0, 0x53, 0x29, 0x25, 0xa3, 0xb8,
				0x44, 0xbc, 0x9e, 0x75, 0x95, 0xf0, 0xbe, 0xb7,
			},
		},
		{
			name:    "Max address (all 0xFF)",
			address: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF",
			expected: [32]byte{
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF, 0xFF, 0xFF,
				0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			addr := common.HexToAddress(tc.address)
			result := AddressToBytes32(addr)

			if result != tc.expected {
				t.Errorf("AddressToBytes32(%s) = %x, want %x", tc.address, result, tc.expected)
			}

			// Verify first 12 bytes are zero (padding)
			for i := 0; i < 12; i++ {
				if result[i] != 0x00 {
					t.Errorf("byte %d should be 0x00, got 0x%02x", i, result[i])
				}
			}
		})
	}
}

func TestBytes32ToAddress(t *testing.T) {
	testCases := []struct {
		name     string
		bytes32  [32]byte
		expected string
	}{
		{
			name: "Zero address",
			bytes32: [32]byte{
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
			expected: "0x0000000000000000000000000000000000000000",
		},
		{
			name: "Non-zero address",
			bytes32: [32]byte{
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x74, 0x2d, 0x35, 0xcc,
				0x66, 0x34, 0xc0, 0x53, 0x29, 0x25, 0xa3, 0xb8,
				0x44, 0xbc, 0x9e, 0x75, 0x95, 0xf0, 0xbe, 0xb7,
			},
			expected: "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb7",
		},
		{
			name: "Max address (all 0xFF)",
			bytes32: [32]byte{
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF, 0xFF, 0xFF,
				0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			},
			expected: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Bytes32ToAddress(tc.bytes32)
			expected := common.HexToAddress(tc.expected)

			if result != expected {
				t.Errorf("Bytes32ToAddress() = %s, want %s", result.Hex(), expected.Hex())
			}
		})
	}
}

func TestAddressToBytes32_RoundTrip(t *testing.T) {
	testAddresses := []string{
		"0x0000000000000000000000000000000000000000",
		"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb7",
		"0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
		"0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF",
		"0x1234567890123456789012345678901234567890",
	}

	for _, addrStr := range testAddresses {
		t.Run(addrStr, func(t *testing.T) {
			addr := common.HexToAddress(addrStr)
			bytes32 := AddressToBytes32(addr)
			recovered := Bytes32ToAddress(bytes32)

			if recovered != addr {
				t.Errorf("round trip failed: original %s, recovered %s", addr.Hex(), recovered.Hex())
			}
		})
	}
}

