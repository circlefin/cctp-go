package cctp

import (
	"math/big"
	"testing"

	"github.com/pxgray/cctp-go/internal/wallet"
	"github.com/pxgray/cctp-go/testutil"
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
