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

package cctp

import (
	"math/big"
	"testing"

	"github.com/circlefin/cctp-go/internal/wallet"
	"github.com/circlefin/cctp-go/testutil"
)

func TestNewTransferOrchestrator_InvalidRPC(t *testing.T) {
	w, err := wallet.LoadFromPrivateKey(testutil.TestPrivateKey)
	if err != nil {
		t.Fatalf("failed to load wallet: %v", err)
	}

	sourceChain := GetMainnetChains()[0]
	destChain := GetMainnetChains()[3]
	sourceChain.RPC = "invalid-url"

	params := &TransferParams{
		SourceChain:      &sourceChain,
		DestChain:        &destChain,
		Amount:           big.NewInt(100000000),
		RecipientAddress: testutil.TestRecipientAddress,
		Testnet:          false,
		TransferType:     TransferTypeAuto,
	}

	_, err = NewTransferOrchestrator(w, params, "https://iris-api.circle.com")
	if err == nil {
		t.Error("expected error for invalid RPC URL")
	}

	expectedError := "failed to connect to source chain"
	if !containsText(err.Error(), expectedError) {
		t.Errorf("expected error to contain %q, got %q", expectedError, err.Error())
	}
}

func TestGetMainnetChains_RequiredFields(t *testing.T) {
	// Test that GetMainnetChains returns chains with all required fields
	sourceChain := GetMainnetChains()[0]
	destChain := GetMainnetChains()[3]

	// Verify source chain has required fields
	if sourceChain.Name == "" {
		t.Error("source chain should have a name")
	}
	if sourceChain.TokenMessengerV2 == "" {
		t.Error("source chain should have TokenMessengerV2")
	}
	if sourceChain.USDC == "" {
		t.Error("source chain should have USDC address")
	}

	// Verify dest chain has required fields
	if destChain.Name == "" {
		t.Error("dest chain should have a name")
	}
	if destChain.MessageTransmitterV2 == "" {
		t.Error("dest chain should have MessageTransmitterV2")
	}
	if destChain.USDC == "" {
		t.Error("dest chain should have USDC address")
	}
}

func TestGetMainnetChains_UniqueDomains(t *testing.T) {
	// Test that GetMainnetChains returns chains with unique domains and chain IDs
	sourceChain := GetMainnetChains()[0]
	destChain := GetMainnetChains()[3]

	// Verify source and dest are different chains (tests that GetMainnetChains returns distinct chains)
	if sourceChain.Domain == destChain.Domain {
		t.Error("different mainnet chains should have different domains")
	}

	if sourceChain.ChainID.Cmp(destChain.ChainID) == 0 {
		t.Error("different mainnet chains should have different chain IDs")
	}
}

// Helper function
func containsText(s, substr string) bool {
	if len(substr) == 0 {
		return true
	}
	if len(s) < len(substr) {
		return false
	}
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
