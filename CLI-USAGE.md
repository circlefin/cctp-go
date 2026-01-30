# CCTP CLI Usage Guide

This guide explains how to use the CCTP CLI with Cobra commands and Viper configuration management.

## Installation

Use the following command to install the CLI:

```shell
go install github.com/circlefin/cctp-go/cmd/cctp
```

## Basic Usage

### Commands

The CLI provides the following commands:

- `cctp transfer` - Start a new USDC transfer
- `cctp resume [tx-hash]` - Resume an existing transfer
- `cctp version` - Print version information
- `cctp help` - Show help information
- `cctp completion` - Generate shell completion scripts

### Global Flags

These flags are available for all commands:

- `--config string` - Path to config file (default: `$HOME/.cctp/config.yaml`)
- `--testnet` - Use testnet instead of mainnet
- `--keystore string` - Path to Ethereum keystore directory

## Configuration

The CLI supports multiple configuration sources with the following precedence (highest to lowest):

1. Command-line flags
2. Environment variables (with `CCTP_` prefix)
3. Configuration file
4. Default values

### Configuration File

Create a configuration file at `~/.cctp/config.yaml` or `./config.yaml`:

```yaml
# Network preference (mainnet or testnet)
network: mainnet

# Keystore path (optional - uses OS-specific defaults)
# macOS: ~/Library/Ethereum/keystore
# Linux: ~/.ethereum/keystore  
# Windows: %APPDATA%\Ethereum\keystore
# keystore_path: /custom/path/to/keystore

# Custom RPC endpoints (optional)
rpc_urls:
  ethereum: https://eth-mainnet.g.alchemy.com/v2/YOUR_KEY
  arbitrum: https://arb-mainnet.g.alchemy.com/v2/YOUR_KEY
  polygon: https://polygon-mainnet.g.alchemy.com/v2/YOUR_KEY

# Default transfer values (optional)
defaults:
  amount: "100"
  recipient: "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"

# Saved addresses for quick access (optional)
saved_addresses:
  - name: "My Wallet"
    address: "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"
  - name: "Treasury"
    address: "0x1234567890123456789012345678901234567890"
  - name: "Partner"
    address: "0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"
```

### Environment Variables

Set environment variables with the `CCTP_` prefix:

```bash
export CCTP_TESTNET=true
export CCTP_KEYSTORE_PATH=/path/to/keystore
export PRIVATE_KEY=0x...  # For private key authentication
```

## Wallet Authentication

The CLI supports two methods for wallet authentication:

### 1. Private Key (Environment Variable)

```bash
export PRIVATE_KEY=0x1234567890abcdef...
cctp transfer
```

### 2. Keystore File

```bash
cctp transfer --keystore ~/.ethereum/keystore
# You'll be prompted to select an account and enter password
```

## Transfer Command

Start a new cross-chain USDC transfer.

### Usage

```bash
cctp transfer [flags]
```

### Flags

- `--source string` - Source chain name (e.g., "Ethereum", "Arbitrum")
- `--dest string` - Destination chain name
- `--amount string` - Amount of USDC to transfer
- `--recipient string` - Recipient address
- `--type string` - Transfer type: `fast`, `standard`, or `auto` (default: `auto`)

### Transfer Types

**Fast Transfer** (~8-20 seconds)
- Faster-than-finality transfers using Circle's Fast Transfer Allowance
- Requires a fee (1-14 bps depending on source chain)
- Best for time-sensitive transfers
- Not available when instant finality chains (Avalanche, Polygon, Sei, Sonic, XDC, HyperEVM, Arc) are used as source

**Standard Transfer** (~13-19 minutes or ~8 seconds for instant finality chains)
- Waits for hard finality on the source chain before attestation
- No fee (Standard Transfer fees are currently 0 bps)
- Available on all chains
- Instant finality chains complete in ~8 seconds with no fee

**Auto Mode** (default)
- Automatically selects Fast Transfer when available
- Falls back to Standard Transfer for instant finality source chains
- Recommended for most users

### Examples

```bash
# Interactive mode (prompts for all parameters)
cctp transfer

# Pre-populate source and destination
cctp transfer --source ethereum --dest arbitrum

# Use Fast Transfer explicitly (with fee)
cctp transfer --source ethereum --dest base --type fast

# Use Standard Transfer (no fee, takes longer)
cctp transfer --source ethereum --dest arbitrum --type standard --amount 1000

# Auto mode (default - uses Fast when available)
cctp transfer --source ethereum --dest polygon --type auto

# Pre-populate all parameters with Fast Transfer
cctp transfer --source ethereum --dest polygon --amount 100 --recipient 0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb --type fast

# Use testnet
cctp transfer --testnet --source "Ethereum Sepolia" --dest "Arbitrum Sepolia"

# Use custom config file
cctp transfer --config ./my-config.yaml

# Transfer from instant finality chain (auto-selects Standard, ~8 seconds)
cctp transfer --source avalanche --dest ethereum --amount 50
```

## Resume Command

Resume an existing transfer using the burn transaction hash.

### Usage

```bash
cctp resume [tx-hash] [flags]
```

### Flags

- `--tx-hash string` - Burn transaction hash

### Examples

```bash
# Interactive mode (prompts for tx hash)
cctp resume

# With tx hash as argument
cctp resume 0x1234567890abcdef...

# With tx hash as flag
cctp resume --tx-hash 0x1234567890abcdef...

# On testnet
cctp resume --testnet 0x1234567890abcdef...
```

## Shell Completion

Generate shell completion scripts for faster command-line usage.

### Bash

```bash
# Generate completion script
cctp completion bash > /etc/bash_completion.d/cctp

# Or for current session only
source <(cctp completion bash)
```

### Zsh

```bash
# Add to ~/.zshrc
source <(cctp completion zsh)

# Or generate to completion directory
cctp completion zsh > "${fpath[1]}/_cctp"
```

### Fish

```bash
cctp completion fish > ~/.config/fish/completions/cctp.fish
```

### PowerShell

```powershell
cctp completion powershell | Out-String | Invoke-Expression
```

## Supported Chains

### Mainnet (18 chains)

| Chain | Domain | Fast Transfer | Standard Transfer | Notes |
|-------|--------|---------------|-------------------|-------|
| Ethereum | 0 | ✅ (~8-20s, 1 bps) | ✅ (~13-19min, 0 bps) | |
| Avalanche | 1 | ❌ | ✅ (~8s, 0 bps) | Instant finality |
| OP Mainnet | 2 | ✅ (~8-20s, 1 bps) | ✅ (~13-19min, 0 bps) | |
| Arbitrum | 3 | ✅ (~8-20s, 1 bps) | ✅ (~13-19min, 0 bps) | |
| Base | 6 | ✅ (~8-20s, 1 bps) | ✅ (~13-19min, 0 bps) | |
| Polygon PoS | 7 | ❌ | ✅ (~8s, 0 bps) | Instant finality |
| Unichain | 10 | ✅ (~8-20s, 1 bps) | ✅ (~13-19min, 0 bps) | |
| Linea | 11 | ✅ (~8-20s, 14 bps) | ✅ (~13-19min, 0 bps) | Highest Fast fee |
| Codex | 12 | ✅ (~8-20s, 1 bps) | ✅ (~13-19min, 0 bps) | |
| Sonic | 13 | ❌ | ✅ (~8s, 0 bps) | Instant finality |
| World Chain | 14 | ✅ (~8-20s, 1 bps) | ✅ (~13-19min, 0 bps) | |
| Sei | 16 | ❌ | ✅ (~8s, 0 bps) | Instant finality |
| XDC | 18 | ❌ | ✅ (~8s, 0 bps) | Instant finality |
| HyperEVM | 19 | ❌ | ✅ (~8s, 0 bps) | Instant finality |
| Ink | 21 | ✅ (~8-20s, 2 bps) | ✅ (~13-19min, 0 bps) | |
| Plume | 22 | ✅ (~8-20s, 2 bps) | ✅ (~13-19min, 0 bps) | |

### Testnet (16 chains)

| Chain | Domain | Notes |
|-------|--------|-------|
| Ethereum Sepolia | 0 | |
| Avalanche Fuji | 1 | Instant finality |
| OP Sepolia | 2 | |
| Arbitrum Sepolia | 3 | |
| Base Sepolia | 6 | |
| Polygon PoS Amoy | 7 | Instant finality |
| Unichain Sepolia | 10 | |
| Linea Sepolia | 11 | |
| Codex Testnet | 12 | |
| Sonic Testnet | 13 | Instant finality |
| World Chain Sepolia | 14 | |
| Sei Testnet | 16 | Instant finality |
| XDC Apothem | 18 | Instant finality |
| HyperEVM Testnet | 19 | Instant finality |
| Ink Testnet | 21 | |
| Plume Testnet | 22 | |
| Arc Testnet | 26 | Instant finality |

**Fee Notes:**
- Fast Transfer fees range from 1-14 bps (0.01%-0.14%)
- Standard Transfer fees are currently 0 bps (no fee)
- Instant finality chains only support Standard Transfer when used as source
- Fees are subject to change with advance notice

## Configuration Precedence Example

With the following setup:

**Config file** (`~/.cctp/config.yaml`):
```yaml
network: mainnet
keystore_path: ~/.ethereum/keystore
```

**Environment variable**:
```bash
export CCTP_TESTNET=true
```

**Command**:
```bash
cctp transfer --keystore /custom/path
```

The final configuration will be:
- `testnet`: `true` (from environment variable, overrides config file)
- `keystore_path`: `/custom/path` (from flag, overrides environment and config)

## Security Best Practices

1. **Never share your private keys** - Keep `PRIVATE_KEY` environment variable secure
2. **Use testnet for testing** - Always test with testnet before mainnet transfers
3. **Verify recipient addresses** - Double-check addresses before confirming transfers
4. **Protect your keystore** - Keep keystore files secure and use strong passwords
5. **Use config files carefully** - Don't commit config files with sensitive data to version control

## Troubleshooting

### No wallet source available

**Error**: `no wallet source available: set PRIVATE_KEY env variable or provide --keystore path`

**Solution**: Either set the `PRIVATE_KEY` environment variable or provide the `--keystore` flag with a valid keystore path.

### Config file not found

This is not an error - the CLI works perfectly fine without a config file. Config files are optional and provide convenience for setting defaults.

### Keystore path does not exist

**Error**: `keystore path does not exist: /path/to/keystore`

**Solution**: Ensure the keystore directory exists and contains valid keystore files. You can also omit the path to use the default location.

## Advanced Usage

### Custom RPC Endpoints

Add custom RPC endpoints to your config file to use your own infrastructure:

```yaml
rpc_urls:
  ethereum: https://eth.llamarpc.com
  arbitrum: https://arb1.arbitrum.io/rpc
  polygon: https://polygon-rpc.com
```

### Saved Addresses

Save frequently used addresses in your config:

```yaml
saved_addresses:
  - name: "Treasury"
    address: "0x1234567890123456789012345678901234567890"
  - name: "Hot Wallet"
    address: "0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"
```

While the TUI doesn't automatically present these yet, they're available for future enhancements.

## Getting Help

- Use `cctp --help` for general help
- Use `cctp [command] --help` for command-specific help
- Visit: https://www.circle.com/en/cross-chain-transfer-protocol

