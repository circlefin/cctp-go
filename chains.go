package cctp

import (
	"fmt"
	"math/big"
)

// Chain represents a blockchain configuration for CCTP V2
type Chain struct {
	Name                 string
	ChainID              *big.Int
	Domain               uint32
	RPC                  string
	TokenMessengerV2     string
	MessageTransmitterV2 string
	USDC                 string
	Explorer             string
	IsTestnet            bool
	InstantFinality      bool // Chains with instant finality don't need Fast Transfer distinction
}

// Chain name constants for mainnet chains
const (
	Ethereum   = "Ethereum"
	Avalanche  = "Avalanche"
	OPMainnet  = "OP Mainnet"
	Arbitrum   = "Arbitrum"
	Base       = "Base"
	PolygonPoS = "Polygon PoS"
	Unichain   = "Unichain"
	Linea      = "Linea"
	Codex      = "Codex"
	Sonic      = "Sonic"
	WorldChain = "World Chain"
	Sei        = "Sei"
	XDC        = "XDC"
	HyperEVM   = "HyperEVM"
	Ink        = "Ink"
	Plume      = "Plume"
)

// Chain name constants for testnet chains
const (
	EthereumSepolia   = "Ethereum Sepolia"
	AvalancheFuji     = "Avalanche Fuji"
	OPSepolia         = "OP Sepolia"
	ArbitrumSepolia   = "Arbitrum Sepolia"
	BaseSepolia       = "Base Sepolia"
	PolygonPoSAmoy    = "Polygon PoS Amoy"
	LineaSepolia      = "Linea Sepolia"
	ArcTestnet        = "Arc Testnet"
	UnichainSepolia   = "Unichain Sepolia"
	CodexTestnet      = "Codex Testnet"
	SonicTestnet      = "Sonic Testnet"
	WorldChainSepolia = "World Chain Sepolia"
	SeiTestnet        = "Sei Testnet"
	XDCApothem        = "XDC Apothem"
	HyperEVMTestnet   = "HyperEVM Testnet"
	InkTestnet        = "Ink Testnet"
	PlumeTestnet      = "Plume Testnet"
)

// GetChains returns all configured chains
func GetChains(testnet bool) []Chain {
	if testnet {
		return GetTestnetChains()
	}
	return GetMainnetChains()
}

// GetMainnetChains returns all mainnet chain configurations
func GetMainnetChains() []Chain {
	// All V2 mainnet chains use the same contract addresses
	const (
		tokenMessengerV2Addr     = "0x28b5a0e9C621a5BadaA536219b3a228C8168cf5d"
		messageTransmitterV2Addr = "0x81D40F21F12A8F0E3252Bccb954D722d4c464B64"
	)

	return []Chain{
		{
			Name:                 "Ethereum",
			ChainID:              big.NewInt(1),
			Domain:               0,
			RPC:                  "https://ethereum-rpc.publicnode.com",
			TokenMessengerV2:     tokenMessengerV2Addr,
			MessageTransmitterV2: messageTransmitterV2Addr,
			USDC:                 "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
			Explorer:             "https://etherscan.io",
			IsTestnet:            false,
			InstantFinality:      false,
		},
		{
			Name:                 "Avalanche",
			ChainID:              big.NewInt(43114),
			Domain:               1,
			RPC:                  "https://api.avax.network/ext/bc/C/rpc",
			TokenMessengerV2:     tokenMessengerV2Addr,
			MessageTransmitterV2: messageTransmitterV2Addr,
			USDC:                 "0xB97EF9Ef8734C71904D8002F8b6Bc66Dd9c48a6E",
			Explorer:             "https://snowtrace.io",
			IsTestnet:            false,
			InstantFinality:      true,
		},
		{
			Name:                 "OP Mainnet",
			ChainID:              big.NewInt(10),
			Domain:               2,
			RPC:                  "https://mainnet.optimism.io",
			TokenMessengerV2:     tokenMessengerV2Addr,
			MessageTransmitterV2: messageTransmitterV2Addr,
			USDC:                 "0x0b2C639c533813f4Aa9D7837CAf62653d097Ff85",
			Explorer:             "https://optimistic.etherscan.io",
			IsTestnet:            false,
			InstantFinality:      false,
		},
		{
			Name:                 "Arbitrum",
			ChainID:              big.NewInt(42161),
			Domain:               3,
			RPC:                  "https://arb1.arbitrum.io/rpc",
			TokenMessengerV2:     tokenMessengerV2Addr,
			MessageTransmitterV2: messageTransmitterV2Addr,
			USDC:                 "0xaf88d065e77c8cC2239327C5EDb3A432268e5831",
			Explorer:             "https://arbiscan.io",
			IsTestnet:            false,
			InstantFinality:      false,
		},
		{
			Name:                 "Base",
			ChainID:              big.NewInt(8453),
			Domain:               6,
			RPC:                  "https://mainnet.base.org",
			TokenMessengerV2:     tokenMessengerV2Addr,
			MessageTransmitterV2: messageTransmitterV2Addr,
			USDC:                 "0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913",
			Explorer:             "https://basescan.org",
			IsTestnet:            false,
			InstantFinality:      false,
		},
		{
			Name:                 "Polygon PoS",
			ChainID:              big.NewInt(137),
			Domain:               7,
			RPC:                  "https://polygon-rpc.com",
			TokenMessengerV2:     tokenMessengerV2Addr,
			MessageTransmitterV2: messageTransmitterV2Addr,
			USDC:                 "0x3c499c542cEF5E3811e1192ce70d8cC03d5c3359",
			Explorer:             "https://polygonscan.com",
			IsTestnet:            false,
			InstantFinality:      true, // Polygon has fast finality
		},
		{
			Name:                 "Unichain",
			ChainID:              big.NewInt(1301),
			Domain:               10,
			RPC:                  "https://sepolia.unichain.org",
			TokenMessengerV2:     tokenMessengerV2Addr,
			MessageTransmitterV2: messageTransmitterV2Addr,
			USDC:                 "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
			Explorer:             "https://unichain.org",
			IsTestnet:            false,
			InstantFinality:      false,
		},
		{
			Name:                 "Linea",
			ChainID:              big.NewInt(59144),
			Domain:               11,
			RPC:                  "https://rpc.linea.build",
			TokenMessengerV2:     tokenMessengerV2Addr,
			MessageTransmitterV2: messageTransmitterV2Addr,
			USDC:                 "0x176211869cA2b568f2A7D4EE941E073a821EE1ff",
			Explorer:             "https://lineascan.build",
			IsTestnet:            false,
			InstantFinality:      false,
		},
		{
			Name:                 "Codex",
			ChainID:              big.NewInt(5115),
			Domain:               12,
			RPC:                  "https://rpc.codex.io",
			TokenMessengerV2:     tokenMessengerV2Addr,
			MessageTransmitterV2: messageTransmitterV2Addr,
			USDC:                 "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
			Explorer:             "https://explorer.codex.io",
			IsTestnet:            false,
			InstantFinality:      false,
		},
		{
			Name:                 "Sonic",
			ChainID:              big.NewInt(146),
			Domain:               13,
			RPC:                  "https://rpc.soniclabs.com",
			TokenMessengerV2:     tokenMessengerV2Addr,
			MessageTransmitterV2: messageTransmitterV2Addr,
			USDC:                 "0x29219dd400f2Bf60E5a23d13Be72B486D4038894",
			Explorer:             "https://sonicscan.org",
			IsTestnet:            false,
			InstantFinality:      true, // Sonic has instant finality
		},
		{
			Name:                 "World Chain",
			ChainID:              big.NewInt(480),
			Domain:               14,
			RPC:                  "https://worldchain-mainnet.g.alchemy.com/public",
			TokenMessengerV2:     tokenMessengerV2Addr,
			MessageTransmitterV2: messageTransmitterV2Addr,
			USDC:                 "0x79A02482A880bCE3F13e09Da970dC34db4CD24d1",
			Explorer:             "https://worldscan.org",
			IsTestnet:            false,
			InstantFinality:      false,
		},
		{
			Name:                 "Sei",
			ChainID:              big.NewInt(1329),
			Domain:               16,
			RPC:                  "https://evm-rpc.sei-apis.com",
			TokenMessengerV2:     tokenMessengerV2Addr,
			MessageTransmitterV2: messageTransmitterV2Addr,
			USDC:                 "0x3894085Ef7Ff0f0aeDf52E2A2704928d1Ec074F1",
			Explorer:             "https://seitrace.com",
			IsTestnet:            false,
			InstantFinality:      true, // Sei has instant finality
		},
		{
			Name:                 "XDC",
			ChainID:              big.NewInt(50),
			Domain:               18,
			RPC:                  "https://rpc.xdcrpc.com",
			TokenMessengerV2:     tokenMessengerV2Addr,
			MessageTransmitterV2: messageTransmitterV2Addr,
			USDC:                 "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
			Explorer:             "https://explorer.xdc.org",
			IsTestnet:            false,
			InstantFinality:      true, // XDC has instant finality
		},
		{
			Name:                 "HyperEVM",
			ChainID:              big.NewInt(998),
			Domain:               19,
			RPC:                  "https://api.hyperevm.com",
			TokenMessengerV2:     tokenMessengerV2Addr,
			MessageTransmitterV2: messageTransmitterV2Addr,
			USDC:                 "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
			Explorer:             "https://explorer.hyperevm.com",
			IsTestnet:            false,
			InstantFinality:      true, // HyperEVM has instant finality
		},
		{
			Name:                 "Ink",
			ChainID:              big.NewInt(57073),
			Domain:               21,
			RPC:                  "https://rpc-gel.inkonchain.com",
			TokenMessengerV2:     tokenMessengerV2Addr,
			MessageTransmitterV2: messageTransmitterV2Addr,
			USDC:                 "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
			Explorer:             "https://explorer.inkonchain.com",
			IsTestnet:            false,
			InstantFinality:      false,
		},
		{
			Name:                 "Plume",
			ChainID:              big.NewInt(98865),
			Domain:               22,
			RPC:                  "https://rpc.plumenetwork.xyz",
			TokenMessengerV2:     tokenMessengerV2Addr,
			MessageTransmitterV2: messageTransmitterV2Addr,
			USDC:                 "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
			Explorer:             "https://explorer.plumenetwork.xyz",
			IsTestnet:            false,
			InstantFinality:      false,
		},
	}
}

// GetTestnetChains returns all testnet chain configurations
func GetTestnetChains() []Chain {
	// All V2 testnet chains use the same contract addresses
	const (
		tokenMessengerV2Addr     = "0x8FE6B999Dc680CcFDD5Bf7EB0974218be2542DAA"
		messageTransmitterV2Addr = "0xE737e5cEBEEBa77EFE34D4aa090756590b1CE275"
	)

	return []Chain{
		{
			Name:                 "Ethereum Sepolia",
			ChainID:              big.NewInt(11155111),
			Domain:               0,
			RPC:                  "https://ethereum-sepolia-rpc.publicnode.com",
			TokenMessengerV2:     tokenMessengerV2Addr,
			MessageTransmitterV2: messageTransmitterV2Addr,
			USDC:                 "0x1c7D4B196Cb0C7B01d743Fbc6116a902379C7238",
			Explorer:             "https://sepolia.etherscan.io",
			IsTestnet:            true,
			InstantFinality:      false,
		},
		{
			Name:                 "Avalanche Fuji",
			ChainID:              big.NewInt(43113),
			Domain:               1,
			RPC:                  "https://api.avax-test.network/ext/bc/C/rpc",
			TokenMessengerV2:     tokenMessengerV2Addr,
			MessageTransmitterV2: messageTransmitterV2Addr,
			USDC:                 "0x5425890298aed601595a70AB815c96711a31Bc65",
			Explorer:             "https://testnet.snowtrace.io",
			IsTestnet:            true,
			InstantFinality:      true,
		},
		{
			Name:                 "OP Sepolia",
			ChainID:              big.NewInt(11155420),
			Domain:               2,
			RPC:                  "https://sepolia.optimism.io",
			TokenMessengerV2:     tokenMessengerV2Addr,
			MessageTransmitterV2: messageTransmitterV2Addr,
			USDC:                 "0x5fd84259d66Cd46123540766Be93DFE6D43130D7",
			Explorer:             "https://sepolia-optimism.etherscan.io",
			IsTestnet:            true,
			InstantFinality:      false,
		},
		{
			Name:                 "Arbitrum Sepolia",
			ChainID:              big.NewInt(421614),
			Domain:               3,
			RPC:                  "https://sepolia-rollup.arbitrum.io/rpc",
			TokenMessengerV2:     tokenMessengerV2Addr,
			MessageTransmitterV2: messageTransmitterV2Addr,
			USDC:                 "0x75faf114eafb1BDbe2F0316DF893fd58CE46AA4d",
			Explorer:             "https://sepolia.arbiscan.io",
			IsTestnet:            true,
			InstantFinality:      false,
		},
		{
			Name:                 "Base Sepolia",
			ChainID:              big.NewInt(84532),
			Domain:               6,
			RPC:                  "https://sepolia.base.org",
			TokenMessengerV2:     tokenMessengerV2Addr,
			MessageTransmitterV2: messageTransmitterV2Addr,
			USDC:                 "0x036CbD53842c5426634e7929541eC2318f3dCF7e",
			Explorer:             "https://sepolia.basescan.org",
			IsTestnet:            true,
			InstantFinality:      false,
		},
		{
			Name:                 "Polygon PoS Amoy",
			ChainID:              big.NewInt(80002),
			Domain:               7,
			RPC:                  "https://rpc-amoy.polygon.technology",
			TokenMessengerV2:     tokenMessengerV2Addr,
			MessageTransmitterV2: messageTransmitterV2Addr,
			USDC:                 "0x41E94Eb019C0762f9Bfcf9Fb1E58725BfB0e7582",
			Explorer:             "https://amoy.polygonscan.com",
			IsTestnet:            true,
			InstantFinality:      true, // Polygon has fast finality
		},
		{
			Name:                 "Linea Sepolia",
			ChainID:              big.NewInt(59141),
			Domain:               11,
			RPC:                  "https://rpc.sepolia.linea.build",
			TokenMessengerV2:     tokenMessengerV2Addr,
			MessageTransmitterV2: messageTransmitterV2Addr,
			USDC:                 "0x5fd84259d66Cd46123540766Be93DFE6D43130D7",
			Explorer:             "https://sepolia.lineascan.build",
			IsTestnet:            true,
			InstantFinality:      false,
		},
		{
			Name:                 "Arc Testnet",
			ChainID:              big.NewInt(5042002),
			Domain:               26,
			RPC:                  "https://rpc.testnet.arc.network",
			TokenMessengerV2:     tokenMessengerV2Addr,
			MessageTransmitterV2: messageTransmitterV2Addr,
			USDC:                 "0x3600000000000000000000000000000000000000",
			Explorer:             "https://testnet.arcscan.app",
			IsTestnet:            true,
			InstantFinality:      true,
		},
		{
			Name:                 "Unichain Sepolia",
			ChainID:              big.NewInt(1301),
			Domain:               10,
			RPC:                  "https://sepolia.unichain.org",
			TokenMessengerV2:     tokenMessengerV2Addr,
			MessageTransmitterV2: messageTransmitterV2Addr,
			USDC:                 "0x5fd84259d66Cd46123540766Be93DFE6D43130D7",
			Explorer:             "https://sepolia.uniscan.xyz",
			IsTestnet:            true,
			InstantFinality:      false,
		},
		{
			Name:                 "Codex Testnet",
			ChainID:              big.NewInt(5115),
			Domain:               12,
			RPC:                  "https://rpc.testnet.codex.io",
			TokenMessengerV2:     tokenMessengerV2Addr,
			MessageTransmitterV2: messageTransmitterV2Addr,
			USDC:                 "0x5fd84259d66Cd46123540766Be93DFE6D43130D7",
			Explorer:             "https://testnet.codexscan.io",
			IsTestnet:            true,
			InstantFinality:      false,
		},
		{
			Name:                 "Sonic Testnet",
			ChainID:              big.NewInt(57054),
			Domain:               13,
			RPC:                  "https://rpc.testnet.soniclabs.com",
			TokenMessengerV2:     tokenMessengerV2Addr,
			MessageTransmitterV2: messageTransmitterV2Addr,
			USDC:                 "0x5fd84259d66Cd46123540766Be93DFE6D43130D7",
			Explorer:             "https://testnet.sonicscan.org",
			IsTestnet:            true,
			InstantFinality:      true,
		},
		{
			Name:                 "World Chain Sepolia",
			ChainID:              big.NewInt(4801),
			Domain:               14,
			RPC:                  "https://worldchain-sepolia.g.alchemy.com/public",
			TokenMessengerV2:     tokenMessengerV2Addr,
			MessageTransmitterV2: messageTransmitterV2Addr,
			USDC:                 "0x5fd84259d66Cd46123540766Be93DFE6D43130D7",
			Explorer:             "https://sepolia.worldscan.org",
			IsTestnet:            true,
			InstantFinality:      false,
		},
		{
			Name:                 "Sei Testnet",
			ChainID:              big.NewInt(713715),
			Domain:               16,
			RPC:                  "https://evm-rpc-testnet.sei-apis.com",
			TokenMessengerV2:     tokenMessengerV2Addr,
			MessageTransmitterV2: messageTransmitterV2Addr,
			USDC:                 "0x5fd84259d66Cd46123540766Be93DFE6D43130D7",
			Explorer:             "https://testnet.seitrace.com",
			IsTestnet:            true,
			InstantFinality:      true,
		},
		{
			Name:                 "XDC Apothem",
			ChainID:              big.NewInt(51),
			Domain:               18,
			RPC:                  "https://erpc.apothem.network",
			TokenMessengerV2:     tokenMessengerV2Addr,
			MessageTransmitterV2: messageTransmitterV2Addr,
			USDC:                 "0x5fd84259d66Cd46123540766Be93DFE6D43130D7",
			Explorer:             "https://apothem.xdcscan.io",
			IsTestnet:            true,
			InstantFinality:      true,
		},
		{
			Name:                 "HyperEVM Testnet",
			ChainID:              big.NewInt(999),
			Domain:               19,
			RPC:                  "https://api.testnet.hyperevm.com",
			TokenMessengerV2:     tokenMessengerV2Addr,
			MessageTransmitterV2: messageTransmitterV2Addr,
			USDC:                 "0x5fd84259d66Cd46123540766Be93DFE6D43130D7",
			Explorer:             "https://testnet.explorer.hyperevm.com",
			IsTestnet:            true,
			InstantFinality:      true,
		},
		{
			Name:                 "Ink Testnet",
			ChainID:              big.NewInt(763373),
			Domain:               21,
			RPC:                  "https://rpc-gel-sepolia.inkonchain.com",
			TokenMessengerV2:     tokenMessengerV2Addr,
			MessageTransmitterV2: messageTransmitterV2Addr,
			USDC:                 "0x5fd84259d66Cd46123540766Be93DFE6D43130D7",
			Explorer:             "https://sepolia.explorer.inkonchain.com",
			IsTestnet:            true,
			InstantFinality:      false,
		},
		{
			Name:                 "Plume Testnet",
			ChainID:              big.NewInt(161221135),
			Domain:               22,
			RPC:                  "https://testnet-rpc.plumenetwork.xyz",
			TokenMessengerV2:     tokenMessengerV2Addr,
			MessageTransmitterV2: messageTransmitterV2Addr,
			USDC:                 "0x5fd84259d66Cd46123540766Be93DFE6D43130D7",
			Explorer:             "https://testnet.explorer.plumenetwork.xyz",
			IsTestnet:            true,
			InstantFinality:      false,
		},
	}
}

// GetChainByDomain returns a chain by its domain ID
func GetChainByDomain(domain uint32, testnet bool) (*Chain, error) {
	chains := GetChains(testnet)
	for _, chain := range chains {
		if chain.Domain == domain {
			return &chain, nil
		}
	}
	return nil, fmt.Errorf("chain with domain %d not found", domain)
}

// GetChainByName returns a chain by its name
func GetChainByName(name string, testnet bool) (*Chain, error) {
	chains := GetChains(testnet)
	for _, chain := range chains {
		if chain.Name == name {
			return &chain, nil
		}
	}
	return nil, fmt.Errorf("chain %s not found", name)
}

// ApplyRPCOverrides applies RPC URL overrides from a map to the given chains.
// The map keys should be chain names, and the values should be the override RPC URLs.
// This allows integrating packages to customize RPC endpoints based on their configuration.
func ApplyRPCOverrides(chains []Chain, overrides map[string]string) []Chain {
	if len(overrides) == 0 {
		return chains
	}

	result := make([]Chain, len(chains))
	for i, chain := range chains {
		result[i] = chain
		if rpcURL, ok := overrides[chain.Name]; ok && rpcURL != "" {
			result[i].RPC = rpcURL
		}
	}
	return result
}
