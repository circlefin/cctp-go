package cctp

import (
	"math/big"
	"testing"
)

func TestGetChains(t *testing.T) {
	t.Run("mainnet", func(t *testing.T) {
		chains := GetChains(false)
		if len(chains) == 0 {
			t.Fatal("expected mainnet chains, got none")
		}
		// Verify we get mainnet chains
		for _, chain := range chains {
			if chain.IsTestnet {
				t.Errorf("expected mainnet chain, got testnet chain: %s", chain.Name)
			}
		}
	})

	t.Run("testnet", func(t *testing.T) {
		chains := GetChains(true)
		if len(chains) == 0 {
			t.Fatal("expected testnet chains, got none")
		}
		// Verify we get testnet chains
		for _, chain := range chains {
			if !chain.IsTestnet {
				t.Errorf("expected testnet chain, got mainnet chain: %s", chain.Name)
			}
		}
	})
}

func TestGetMainnetChains(t *testing.T) {
	chains := GetMainnetChains()

	if len(chains) != 16 {
		t.Errorf("expected 16 mainnet chains, got %d", len(chains))
	}

	// Expected contract addresses for V2 mainnet
	expectedTokenMessenger := "0x28b5a0e9C621a5BadaA536219b3a228C8168cf5d"
	expectedMessageTransmitter := "0x81D40F21F12A8F0E3252Bccb954D722d4c464B64"

	// Track domains to ensure uniqueness
	seenDomains := make(map[uint32]bool)

	for _, chain := range chains {
		// Check required fields are populated
		if chain.Name == "" {
			t.Error("chain has empty name")
		}
		if chain.ChainID == nil {
			t.Errorf("chain %s has nil ChainID", chain.Name)
		}
		if chain.RPC == "" {
			t.Errorf("chain %s has empty RPC", chain.Name)
		}
		if chain.USDC == "" {
			t.Errorf("chain %s has empty USDC address", chain.Name)
		}
		if chain.Explorer == "" {
			t.Errorf("chain %s has empty Explorer", chain.Name)
		}
		if chain.IsTestnet {
			t.Errorf("chain %s should not be marked as testnet", chain.Name)
		}

		// Check V2 contract addresses match
		if chain.TokenMessengerV2 != expectedTokenMessenger {
			t.Errorf("chain %s has incorrect TokenMessengerV2: got %s, want %s",
				chain.Name, chain.TokenMessengerV2, expectedTokenMessenger)
		}
		if chain.MessageTransmitterV2 != expectedMessageTransmitter {
			t.Errorf("chain %s has incorrect MessageTransmitterV2: got %s, want %s",
				chain.Name, chain.MessageTransmitterV2, expectedMessageTransmitter)
		}

		// Check domain uniqueness
		if seenDomains[chain.Domain] {
			t.Errorf("duplicate domain %d found for chain %s", chain.Domain, chain.Name)
		}
		seenDomains[chain.Domain] = true
	}

	// Verify specific chains exist with correct configuration
	testCases := []struct {
		name            string
		domain          uint32
		chainID         int64
		instantFinality bool
	}{
		{"Ethereum", 0, 1, false},
		{"Avalanche", 1, 43114, true},
		{"OP Mainnet", 2, 10, false},
		{"Arbitrum", 3, 42161, false},
		{"Base", 6, 8453, false},
		{"Polygon PoS", 7, 137, true},
		{"Sonic", 13, 146, true},
		{"Sei", 16, 1329, true},
		{"XDC", 18, 50, true},
		{"HyperEVM", 19, 998, true},
	}

	chainMap := make(map[string]Chain)
	for _, chain := range chains {
		chainMap[chain.Name] = chain
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			chain, ok := chainMap[tc.name]
			if !ok {
				t.Fatalf("chain %s not found in mainnet chains", tc.name)
			}
			if chain.Domain != tc.domain {
				t.Errorf("expected domain %d, got %d", tc.domain, chain.Domain)
			}
			if chain.ChainID.Cmp(big.NewInt(tc.chainID)) != 0 {
				t.Errorf("expected chainID %d, got %s", tc.chainID, chain.ChainID.String())
			}
			if chain.InstantFinality != tc.instantFinality {
				t.Errorf("expected instantFinality %v, got %v", tc.instantFinality, chain.InstantFinality)
			}
		})
	}
}

func TestGetTestnetChains(t *testing.T) {
	chains := GetTestnetChains()

	if len(chains) != 17 {
		t.Errorf("expected 17 testnet chains, got %d", len(chains))
	}

	// Expected contract addresses for V2 testnet
	expectedTokenMessenger := "0x8FE6B999Dc680CcFDD5Bf7EB0974218be2542DAA"
	expectedMessageTransmitter := "0xE737e5cEBEEBa77EFE34D4aa090756590b1CE275"

	// Track domains to ensure uniqueness
	seenDomains := make(map[uint32]bool)

	for _, chain := range chains {
		// Check required fields are populated
		if chain.Name == "" {
			t.Error("chain has empty name")
		}
		if chain.ChainID == nil {
			t.Errorf("chain %s has nil ChainID", chain.Name)
		}
		if chain.RPC == "" {
			t.Errorf("chain %s has empty RPC", chain.Name)
		}
		if chain.USDC == "" {
			t.Errorf("chain %s has empty USDC address", chain.Name)
		}
		if chain.Explorer == "" {
			t.Errorf("chain %s has empty Explorer", chain.Name)
		}
		if !chain.IsTestnet {
			t.Errorf("chain %s should be marked as testnet", chain.Name)
		}

		// Check V2 contract addresses match
		if chain.TokenMessengerV2 != expectedTokenMessenger {
			t.Errorf("chain %s has incorrect TokenMessengerV2: got %s, want %s",
				chain.Name, chain.TokenMessengerV2, expectedTokenMessenger)
		}
		if chain.MessageTransmitterV2 != expectedMessageTransmitter {
			t.Errorf("chain %s has incorrect MessageTransmitterV2: got %s, want %s",
				chain.Name, chain.MessageTransmitterV2, expectedMessageTransmitter)
		}

		// Check domain uniqueness
		if seenDomains[chain.Domain] {
			t.Errorf("duplicate domain %d found for chain %s", chain.Domain, chain.Name)
		}
		seenDomains[chain.Domain] = true
	}

	// Verify specific chains exist with correct configuration
	testCases := []struct {
		name            string
		domain          uint32
		chainID         int64
		instantFinality bool
	}{
		{"Ethereum Sepolia", 0, 11155111, false},
		{"Avalanche Fuji", 1, 43113, true},
		{"OP Sepolia", 2, 11155420, false},
		{"Arbitrum Sepolia", 3, 421614, false},
		{"Base Sepolia", 6, 84532, false},
		{"Polygon PoS Amoy", 7, 80002, true},
		{"Arc Testnet", 26, 5042002, true},
		{"Sonic Testnet", 13, 57054, true},
		{"Sei Testnet", 16, 713715, true},
	}

	chainMap := make(map[string]Chain)
	for _, chain := range chains {
		chainMap[chain.Name] = chain
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			chain, ok := chainMap[tc.name]
			if !ok {
				t.Fatalf("chain %s not found in testnet chains", tc.name)
			}
			if chain.Domain != tc.domain {
				t.Errorf("expected domain %d, got %d", tc.domain, chain.Domain)
			}
			if chain.ChainID.Cmp(big.NewInt(tc.chainID)) != 0 {
				t.Errorf("expected chainID %d, got %s", tc.chainID, chain.ChainID.String())
			}
			if chain.InstantFinality != tc.instantFinality {
				t.Errorf("expected instantFinality %v, got %v", tc.instantFinality, chain.InstantFinality)
			}
		})
	}
}

func TestGetChainByDomain(t *testing.T) {
	testCases := []struct {
		name        string
		domain      uint32
		testnet     bool
		expectError bool
		expectName  string
	}{
		{"mainnet ethereum", 0, false, false, "Ethereum"},
		{"mainnet avalanche", 1, false, false, "Avalanche"},
		{"mainnet arbitrum", 3, false, false, "Arbitrum"},
		{"testnet ethereum", 0, true, false, "Ethereum Sepolia"},
		{"testnet avalanche", 1, true, false, "Avalanche Fuji"},
		{"invalid mainnet domain", 999, false, true, ""},
		{"invalid testnet domain", 999, true, true, ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			chain, err := GetChainByDomain(tc.domain, tc.testnet)
			if tc.expectError {
				if err == nil {
					t.Error("expected error but got none")
				}
				if chain != nil {
					t.Error("expected nil chain on error")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if chain == nil {
					t.Fatal("expected chain but got nil")
				}
				if chain.Name != tc.expectName {
					t.Errorf("expected chain name %s, got %s", tc.expectName, chain.Name)
				}
				if chain.Domain != tc.domain {
					t.Errorf("expected domain %d, got %d", tc.domain, chain.Domain)
				}
			}
		})
	}
}

func TestGetChainByName(t *testing.T) {
	testCases := []struct {
		name         string
		chainName    string
		testnet      bool
		expectError  bool
		expectDomain uint32
	}{
		{"mainnet ethereum", "Ethereum", false, false, 0},
		{"mainnet avalanche", "Avalanche", false, false, 1},
		{"mainnet arbitrum", "Arbitrum", false, false, 3},
		{"mainnet base", "Base", false, false, 6},
		{"testnet ethereum", "Ethereum Sepolia", true, false, 0},
		{"testnet avalanche", "Avalanche Fuji", true, false, 1},
		{"invalid chain name", "NonExistent Chain", false, true, 0},
		{"wrong network", "Ethereum Sepolia", false, true, 0}, // Looking for testnet chain in mainnet
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			chain, err := GetChainByName(tc.chainName, tc.testnet)
			if tc.expectError {
				if err == nil {
					t.Error("expected error but got none")
				}
				if chain != nil {
					t.Error("expected nil chain on error")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if chain == nil {
					t.Fatal("expected chain but got nil")
				}
				if chain.Name != tc.chainName {
					t.Errorf("expected chain name %s, got %s", tc.chainName, chain.Name)
				}
				if chain.Domain != tc.expectDomain {
					t.Errorf("expected domain %d, got %d", tc.expectDomain, chain.Domain)
				}
			}
		})
	}
}

func TestChainStructValidation(t *testing.T) {
	// Test both mainnet and testnet chains
	allChains := append(GetMainnetChains(), GetTestnetChains()...)

	for _, chain := range allChains {
		t.Run(chain.Name, func(t *testing.T) {
			// Name should not be empty
			if chain.Name == "" {
				t.Error("chain Name is empty")
			}

			// ChainID should not be nil
			if chain.ChainID == nil {
				t.Error("chain ChainID is nil")
			}

			// RPC should not be empty and should be a valid URL format
			if chain.RPC == "" {
				t.Error("chain RPC is empty")
			}
			if len(chain.RPC) > 0 && chain.RPC[:4] != "http" {
				t.Errorf("chain RPC does not start with http: %s", chain.RPC)
			}

			// Contract addresses should be valid Ethereum addresses (42 chars with 0x prefix)
			if len(chain.TokenMessengerV2) != 42 || chain.TokenMessengerV2[:2] != "0x" {
				t.Errorf("invalid TokenMessengerV2 address: %s", chain.TokenMessengerV2)
			}
			if len(chain.MessageTransmitterV2) != 42 || chain.MessageTransmitterV2[:2] != "0x" {
				t.Errorf("invalid MessageTransmitterV2 address: %s", chain.MessageTransmitterV2)
			}
			if len(chain.USDC) != 42 || chain.USDC[:2] != "0x" {
				t.Errorf("invalid USDC address: %s", chain.USDC)
			}

			// Explorer should not be empty and should be a valid URL
			if chain.Explorer == "" {
				t.Error("chain Explorer is empty")
			}
			if len(chain.Explorer) > 0 && chain.Explorer[:4] != "http" {
				t.Errorf("chain Explorer does not start with http: %s", chain.Explorer)
			}
		})
	}
}

func TestApplyRPCOverrides(t *testing.T) {
	t.Run("empty overrides returns original chains", func(t *testing.T) {
		chains := GetMainnetChains()
		originalRPC := chains[0].RPC

		result := ApplyRPCOverrides(chains, map[string]string{})

		if result[0].RPC != originalRPC {
			t.Errorf("expected RPC to remain %s, got %s", originalRPC, result[0].RPC)
		}
	})

	t.Run("nil overrides returns original chains", func(t *testing.T) {
		chains := GetMainnetChains()
		originalRPC := chains[0].RPC

		result := ApplyRPCOverrides(chains, nil)

		if result[0].RPC != originalRPC {
			t.Errorf("expected RPC to remain %s, got %s", originalRPC, result[0].RPC)
		}
	})

	t.Run("applies single override", func(t *testing.T) {
		chains := GetMainnetChains()
		customRPC := "https://custom-ethereum-rpc.example.com"

		overrides := map[string]string{
			"Ethereum": customRPC,
		}

		result := ApplyRPCOverrides(chains, overrides)

		// Find Ethereum chain
		var ethereumChain *Chain
		for i := range result {
			if result[i].Name == "Ethereum" {
				ethereumChain = &result[i]
				break
			}
		}

		if ethereumChain == nil {
			t.Fatal("Ethereum chain not found in result")
		}

		if ethereumChain.RPC != customRPC {
			t.Errorf("expected Ethereum RPC to be %s, got %s", customRPC, ethereumChain.RPC)
		}
	})

	t.Run("applies multiple overrides", func(t *testing.T) {
		chains := GetMainnetChains()
		customEthRPC := "https://custom-ethereum-rpc.example.com"
		customArbRPC := "https://custom-arbitrum-rpc.example.com"

		overrides := map[string]string{
			"Ethereum": customEthRPC,
			"Arbitrum": customArbRPC,
		}

		result := ApplyRPCOverrides(chains, overrides)

		// Verify Ethereum override
		var ethereumChain *Chain
		for i := range result {
			if result[i].Name == "Ethereum" {
				ethereumChain = &result[i]
				break
			}
		}

		if ethereumChain == nil {
			t.Fatal("Ethereum chain not found in result")
		}

		if ethereumChain.RPC != customEthRPC {
			t.Errorf("expected Ethereum RPC to be %s, got %s", customEthRPC, ethereumChain.RPC)
		}

		// Verify Arbitrum override
		var arbitrumChain *Chain
		for i := range result {
			if result[i].Name == "Arbitrum" {
				arbitrumChain = &result[i]
				break
			}
		}

		if arbitrumChain == nil {
			t.Fatal("Arbitrum chain not found in result")
		}

		if arbitrumChain.RPC != customArbRPC {
			t.Errorf("expected Arbitrum RPC to be %s, got %s", customArbRPC, arbitrumChain.RPC)
		}
	})

	t.Run("ignores empty string overrides", func(t *testing.T) {
		chains := GetMainnetChains()
		originalRPC := chains[0].RPC

		overrides := map[string]string{
			chains[0].Name: "",
		}

		result := ApplyRPCOverrides(chains, overrides)

		if result[0].RPC != originalRPC {
			t.Errorf("expected RPC to remain %s when override is empty, got %s", originalRPC, result[0].RPC)
		}
	})

	t.Run("ignores non-matching chain names", func(t *testing.T) {
		chains := GetMainnetChains()
		originalRPC := chains[0].RPC

		overrides := map[string]string{
			"NonExistentChain": "https://custom-rpc.example.com",
		}

		result := ApplyRPCOverrides(chains, overrides)

		if result[0].RPC != originalRPC {
			t.Errorf("expected RPC to remain %s for unaffected chains, got %s", originalRPC, result[0].RPC)
		}
	})

	t.Run("does not mutate original chains slice", func(t *testing.T) {
		chains := GetMainnetChains()
		originalRPC := chains[0].RPC

		overrides := map[string]string{
			chains[0].Name: "https://custom-rpc.example.com",
		}

		result := ApplyRPCOverrides(chains, overrides)

		// Original should be unchanged
		if chains[0].RPC != originalRPC {
			t.Errorf("original chains slice was mutated, expected %s, got %s", originalRPC, chains[0].RPC)
		}

		// Result should have override
		if result[0].RPC == originalRPC {
			t.Errorf("result chains should have override applied")
		}
	})
}

func TestChainDomainUniqueness(t *testing.T) {
	t.Run("mainnet domains unique", func(t *testing.T) {
		chains := GetMainnetChains()
		domains := make(map[uint32]string)

		for _, chain := range chains {
			if existingChain, exists := domains[chain.Domain]; exists {
				t.Errorf("duplicate domain %d found: %s and %s", chain.Domain, existingChain, chain.Name)
			}
			domains[chain.Domain] = chain.Name
		}
	})

	t.Run("testnet domains unique", func(t *testing.T) {
		chains := GetTestnetChains()
		domains := make(map[uint32]string)

		for _, chain := range chains {
			if existingChain, exists := domains[chain.Domain]; exists {
				t.Errorf("duplicate domain %d found: %s and %s", chain.Domain, existingChain, chain.Name)
			}
			domains[chain.Domain] = chain.Name
		}
	})
}

func TestChainIDUniqueness(t *testing.T) {
	t.Run("mainnet chain IDs unique", func(t *testing.T) {
		chains := GetMainnetChains()
		chainIDs := make(map[string]string)

		for _, chain := range chains {
			chainIDStr := chain.ChainID.String()
			if existingChain, exists := chainIDs[chainIDStr]; exists {
				t.Errorf("duplicate chain ID %s found: %s and %s", chainIDStr, existingChain, chain.Name)
			}
			chainIDs[chainIDStr] = chain.Name
		}
	})

	t.Run("testnet chain IDs unique", func(t *testing.T) {
		chains := GetTestnetChains()
		chainIDs := make(map[string]string)

		for _, chain := range chains {
			chainIDStr := chain.ChainID.String()
			if existingChain, exists := chainIDs[chainIDStr]; exists {
				t.Errorf("duplicate chain ID %s found: %s and %s", chainIDStr, existingChain, chain.Name)
			}
			chainIDs[chainIDStr] = chain.Name
		}
	})
}

func TestChainNameConstants(t *testing.T) {
	t.Run("mainnet constants match chain names", func(t *testing.T) {
		chains := GetMainnetChains()

		// Create a map of expected constants to verify they all exist
		expectedConstants := map[string]string{
			Ethereum:   Ethereum,
			Avalanche:  Avalanche,
			OPMainnet:  OPMainnet,
			Arbitrum:   Arbitrum,
			Base:       Base,
			PolygonPoS: PolygonPoS,
			Unichain:   Unichain,
			Linea:      Linea,
			Codex:      Codex,
			Sonic:      Sonic,
			WorldChain: WorldChain,
			Sei:        Sei,
			XDC:        XDC,
			HyperEVM:   HyperEVM,
			Ink:        Ink,
			Plume:      Plume,
		}

		// Verify each chain name exists in the constants
		for _, chain := range chains {
			if _, exists := expectedConstants[chain.Name]; !exists {
				t.Errorf("chain %q does not have a corresponding constant", chain.Name)
			}
		}

		// Verify each constant matches an actual chain
		for constant := range expectedConstants {
			found := false
			for _, chain := range chains {
				if chain.Name == constant {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("constant %q does not match any mainnet chain name", constant)
			}
		}
	})

	t.Run("testnet constants match chain names", func(t *testing.T) {
		chains := GetTestnetChains()

		// Create a map of expected constants to verify they all exist
		expectedConstants := map[string]string{
			EthereumSepolia:   EthereumSepolia,
			AvalancheFuji:     AvalancheFuji,
			OPSepolia:         OPSepolia,
			ArbitrumSepolia:   ArbitrumSepolia,
			BaseSepolia:       BaseSepolia,
			PolygonPoSAmoy:    PolygonPoSAmoy,
			LineaSepolia:      LineaSepolia,
			ArcTestnet:        ArcTestnet,
			UnichainSepolia:   UnichainSepolia,
			CodexTestnet:      CodexTestnet,
			SonicTestnet:      SonicTestnet,
			WorldChainSepolia: WorldChainSepolia,
			SeiTestnet:        SeiTestnet,
			XDCApothem:        XDCApothem,
			HyperEVMTestnet:   HyperEVMTestnet,
			InkTestnet:        InkTestnet,
			PlumeTestnet:      PlumeTestnet,
		}

		// Verify each chain name exists in the constants
		for _, chain := range chains {
			if _, exists := expectedConstants[chain.Name]; !exists {
				t.Errorf("chain %q does not have a corresponding constant", chain.Name)
			}
		}

		// Verify each constant matches an actual chain
		for constant := range expectedConstants {
			found := false
			for _, chain := range chains {
				if chain.Name == constant {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("constant %q does not match any testnet chain name", constant)
			}
		}
	})

	t.Run("constants can be used for RPC overrides", func(t *testing.T) {
		chains := GetMainnetChains()

		// Test using constants in RPC overrides map
		customRPCs := map[string]string{
			Ethereum: "https://custom-ethereum-rpc.example.com",
			Arbitrum: "https://custom-arbitrum-rpc.example.com",
			Base:     "https://custom-base-rpc.example.com",
		}

		updatedChains := ApplyRPCOverrides(chains, customRPCs)

		// Verify the overrides were applied
		for _, chain := range updatedChains {
			switch chain.Name {
			case Ethereum:
				if chain.RPC != customRPCs[Ethereum] {
					t.Errorf("expected Ethereum RPC to be overridden to %s, got %s", customRPCs[Ethereum], chain.RPC)
				}
			case Arbitrum:
				if chain.RPC != customRPCs[Arbitrum] {
					t.Errorf("expected Arbitrum RPC to be overridden to %s, got %s", customRPCs[Arbitrum], chain.RPC)
				}
			case Base:
				if chain.RPC != customRPCs[Base] {
					t.Errorf("expected Base RPC to be overridden to %s, got %s", customRPCs[Base], chain.RPC)
				}
			}
		}
	})
}
