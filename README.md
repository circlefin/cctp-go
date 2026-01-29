# CCTP Go SDK and CLI

A Go SDK and Terminal UI-based CLI tool for executing Circle's CCTP V2 cross-chain USDC transfers across all supported EVM chains.

## Features

- 📦 **Importable Go SDK** - Use CCTP functionality in your Go applications
- 🚀 **Fast Transfer** support (~8-20 seconds for EVM chains, with fee)
- ⏱️ **Standard Transfer** support (~13-19 minutes, no fee)
- 🔗 Support for all CCTP V2 EVM chains (18 mainnet + 10 testnet chains)
- 💼 Multiple wallet support (keystore files, private keys)
- 📊 Real-time progress tracking with detailed fee information
- 🎨 Beautiful terminal user interface

## Usage

### CLI Usage

Install the CLI with the following command:

```shell
go install github.com/circlefin/cctp-go/cli
```

To use the CLI, you must have your wallet's private key stored in the [Geth keystore](https://geth.ethereum.org/docs/fundamentals/account-management#importing-a-keyfile).

Ensure that the path that Go builds executables is in your `PATH` variable.

The following are example CLI commands:

```bash
# Run with interactive TUI
./bin/cctp transfer

# Or run directly without building
go run ./cmd/cli transfer

# Use testnet
./bin/cctp transfer --testnet

# Specify transfer type
./bin/cctp transfer --type fast      # Fast Transfer (~8-20 seconds, with fee)
./bin/cctp transfer --type standard  # Standard Transfer (~13-19 minutes, no fee)
./bin/cctp transfer --type auto      # Auto-select based on chain (default)

# Pre-populate transfer details
./bin/cctp transfer --source ethereum --dest arbitrum --amount 100 --recipient 0x... --type fast

# Resume existing transfer
./bin/cctp resume 0x...  # Provide burn transaction hash

# Specify keystore directory
./bin/cctp --keystore ~/.ethereum/keystore transfer
```

For detailed CLI usage, see [CLI-USAGE.md](CLI-USAGE.md).

### Library Usage

The CCTP SDK can be imported and used in your Go applications.

```
go get github.com/circlefin/cctp-go
```

Example usage in your Go application:

```go
package main

import (
    "context"
    "fmt"
    "math/big"

    "github.com/ethereum/go-ethereum/common"
    "github.com/circlefin/cctp-go"
)

func main() {
    // Get available chains
    chains := cctp.GetChains(false) // false = mainnet

    // Create an Iris client for attestations
    irisClient := cctp.NewIrisClient("https://iris-api.circle.com")

    // Poll for an attestation
    ctx := context.Background()
    msg, err := irisClient.PollForAttestation(
        ctx,
        0, // source domain (Ethereum)
        "0x...", // burn transaction hash
        nil, // optional progress callback
    )
    if err != nil {
        panic(err)
    }

    fmt.Printf("Attestation received: %s\n", msg.Attestation)
}
```

### Overriding RPC Endpoints

The library provides default RPC endpoints for all supported chains, but integrating packages can override them with custom endpoints (e.g., from their own configuration or API keys):

```go
package main

import (
    "github.com/circlefin/cctp-go"
)

func main() {
    // Get chains with default RPC endpoints
    chains := cctp.GetChains(false) // false = mainnet

    // Define custom RPC endpoints using the chain name constants
    // This provides IDE autocomplete and prevents typos
    customRPCs := map[string]string{
        cctp.Ethereum: "https://eth-mainnet.g.alchemy.com/v2/YOUR_API_KEY",
        cctp.Arbitrum: "https://arb-mainnet.g.alchemy.com/v2/YOUR_API_KEY",
        cctp.Base:     "https://base-mainnet.g.alchemy.com/v2/YOUR_API_KEY",
    }

    // Apply RPC overrides
    chains = cctp.ApplyRPCOverrides(chains, customRPCs)

    // Now use chains with your custom RPC endpoints
    for _, chain := range chains {
        if chain.Name == cctp.Ethereum {
            fmt.Printf("Ethereum RPC: %s\n", chain.RPC)
        }
    }
}
```

**Available Chain Constants:**

Mainnet: `cctp.Ethereum`, `cctp.Avalanche`, `cctp.OPMainnet`, `cctp.Arbitrum`, `cctp.Base`, `cctp.PolygonPoS`, `cctp.Unichain`, `cctp.Linea`, `cctp.Codex`, `cctp.Sonic`, `cctp.WorldChain`, `cctp.Sei`, `cctp.XDC`, `cctp.HyperEVM`, `cctp.Ink`, `cctp.Plume`

Testnet: `cctp.EthereumSepolia`, `cctp.AvalancheFuji`, `cctp.OPSepolia`, `cctp.ArbitrumSepolia`, `cctp.BaseSepolia`, `cctp.PolygonPoSAmoy`, `cctp.LineaSepolia`, `cctp.ArcTestnet`, `cctp.UnichainSepolia`, `cctp.CodexTestnet`, `cctp.SonicTestnet`, `cctp.WorldChainSepolia`, `cctp.SeiTestnet`, `cctp.XDCApothem`, `cctp.HyperEVMTestnet`, `cctp.InkTestnet`, `cctp.PlumeTestnet`

**Notes:**

- Use the chain name constants (e.g., `cctp.Ethereum`) for IDE autocomplete support and to avoid typos
- You can still use string literals if needed, but they must match exactly (case-sensitive)
- Empty string overrides are ignored (original RPC is preserved)
- Non-matching chain names are ignored
- The original chains slice is not mutated (returns a new slice)

### Contract Interactions

```go
package main

import (
    "context"
    "fmt"
    "math/big"

    "github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/circlefin/cctp-go/tokenmessenger"
)

func main() {
    // Connect to a chain
    client, err := ethclient.Dial("https://ethereum-rpc.publicnode.com")
    if err != nil {
        panic(err)
    }

    // Create V2 contract binding
    ctx := context.Background()
    ierc20 := tokenmessenger.NewIERC20()
    usdcAddress := common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
    usdcInstance := ierc20.Instance(client, usdcAddress)

    // Check USDC balance using V2 bindings
    walletAddress := common.HexToAddress("0x...")
    balance, err := bind.Call(usdcInstance, &bind.CallOpts{Context: ctx},
        ierc20.PackBalanceOf(walletAddress), ierc20.UnpackBalanceOf)
    if err != nil {
        panic(err)
    }

    fmt.Printf("USDC Balance: %s\n", balance.String())
}
```

### Full Transfer Example

See the [CLI implementation](cmd/cli/main.go) for a complete example of orchestrating a full CCTP transfer.

## Supported Chains

### Mainnet (18 chains)

- **Ethereum** (Domain 0)
- **Avalanche** (Domain 1) - Instant finality
- **OP Mainnet** (Domain 2)
- **Arbitrum** (Domain 3)
- **Base** (Domain 6)
- **Polygon PoS** (Domain 7) - Instant finality
- **Unichain** (Domain 10)
- **Linea** (Domain 11)
- **Codex** (Domain 12)
- **Sonic** (Domain 13) - Instant finality
- **World Chain** (Domain 14)
- **Sei** (Domain 16) - Instant finality
- **XDC** (Domain 18) - Instant finality
- **HyperEVM** (Domain 19) - Instant finality
- **Ink** (Domain 21)
- **Plume** (Domain 22)

### Testnet (10 chains)

- Ethereum Sepolia, Arbitrum Sepolia, Avalanche Fuji, Base Sepolia, Polygon PoS Amoy, Linea Sepolia, Arc Testnet, Unichain Sepolia, Codex Testnet, Sonic Testnet, World Chain Sepolia, Sei Testnet, XDC Apothem, HyperEVM Testnet, Ink Testnet, Plume Testnet

### Transfer Types

**Fast Transfer** (~8-20 seconds)

- Available on most chains when used as source
- Requires fee (1-14 bps depending on source chain)
- Best for time-sensitive transfers
- Not available when instant finality chains are the source

**Standard Transfer** (~13-19 minutes or ~8 seconds for instant finality chains)

- Available on all chains
- No fee
- Best for cost-sensitive transfers
- Instant finality chains: Avalanche, Polygon PoS, Sei, Sonic, XDC, HyperEVM, Arc Testnet

## Setup and Building

### Prerequisites

- Go 1.25.2 or later
- [Task](https://taskfile.dev/) (optional, for using Taskfile commands)

### Getting Started

```bash
# Clone the repository
git clone https://github.com/circlefin/cctp-go
cd cctp-go

# Install dependencies
go mod download

# Build the CLI binary (using Task)
task build

# Or build directly with Go
go build -o bin/cctp ./cmd/cli

# Run the CLI
./bin/cctp transfer

# Or run directly without building
go run ./cmd/cli transfer
```

### Available Build Tasks

```bash
# Build the CLI binary
task build

# Run tests
task test

# Run tests with verbose output
task test-verbose

# Run tests with coverage report
task test-coverage

# Generate HTML coverage report
task test-coverage-html

# Clean build artifacts
task clean
```

## Package Structure

```
cctp-go/
├── client.go          # Iris attestation client
├── transfer.go        # Transfer orchestration
├── chains.go          # Chain configurations
├── contracts.go       # Contract interactions
├── cmd/
│   └── cli/           # CLI application
│       └── main.go
└── internal/          # CLI-specific code
    ├── config/
    ├── ui/
    └── wallet/
```

## API Documentation

### Main Types

- `Chain` - Blockchain configuration with RPC, contracts, and metadata
- `IrisClient` - Client for Circle's attestation service with fee API support
- `TransferOrchestrator` - Complete transfer workflow orchestration
- `TransferType` - Transfer type enum (Fast, Standard, Auto)
- `TransferParams` - Parameters for a cross-chain transfer

### Key Functions

- `GetChains(testnet bool)` - Get all configured chains
- `GetChainByDomain(domain uint32, testnet bool)` - Find chain by domain ID
- `GetChainByName(name string, testnet bool)` - Find chain by name
- `NewIrisClient(baseURL string)` - Create attestation client
- `IrisClient.GetTransferFees(ctx, sourceDomain, destDomain)` - Query transfer fees from API
- `NewTransferOrchestrator(...)` - Create transfer orchestrator

### V2 Contract Bindings

This library uses [Go Ethereum V2 Contract Bindings](https://geth.ethereum.org/docs/developers/dapp-developer/native-bindings-v2) for type-safe contract interactions:

- `NewIERC20()` - Create ERC20 contract binding
- `NewTokenMessengerV2()` - Create TokenMessengerV2 contract binding
- `NewMessageTransmitterV2()` - Create MessageTransmitterV2 contract binding
- `BindCall()` - Make read-only contract calls
- `BindTransact()` - Send transactions to contracts
- `BindWaitMined()` - Wait for transaction confirmation
