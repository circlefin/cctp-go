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

package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/circlefin/cctp-go"
	"github.com/circlefin/cctp-go/internal/config"
	"github.com/circlefin/cctp-go/internal/ui"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

var (
	transferSource    string
	transferDest      string
	transferAmount    string
	transferRecipient string
	transferType      string
)

var transferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "Start a new USDC transfer",
	Long: `Start a new cross-chain USDC transfer using the interactive TUI.

You can optionally specify transfer parameters as flags to pre-populate the TUI:
  --source: Source chain name
  --dest: Destination chain name
  --amount: Amount of USDC to transfer
  --recipient: Recipient address
  --type: Transfer type (fast, standard, or auto)

If any parameters are not provided, you will be prompted in the TUI.

Transfer Types:
  fast:     Fast Transfer (~8-20 seconds, with fee)
  standard: Standard Transfer (~13-19 minutes, no fee)
  auto:     Auto-select based on chain capabilities (default)`,
	Example: `  # Start interactive transfer
  cctp transfer

  # Pre-populate source and destination
  cctp transfer --source ethereum --dest arbitrum

  # Use Standard Transfer for lower fees
  cctp transfer --source ethereum --dest base --type standard

  # Pre-populate all parameters
  cctp transfer --source ethereum --dest polygon --amount 100 --recipient 0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb --type fast`,
	RunE: runTransfer,
}

func init() {
	transferCmd.Flags().StringVar(&transferSource, "source", "", "source chain name")
	transferCmd.Flags().StringVar(&transferDest, "dest", "", "destination chain name")
	transferCmd.Flags().StringVar(&transferAmount, "amount", "", "amount of USDC to transfer")
	transferCmd.Flags().StringVar(&transferRecipient, "recipient", "", "recipient address")
	transferCmd.Flags().StringVar(&transferType, "type", "auto", "transfer type: fast, standard, or auto (default: auto)")

	rootCmd.AddCommand(transferCmd)
}

func runTransfer(cmd *cobra.Command, args []string) error {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	// Validate configuration
	if err := cfg.Validate(); err != nil {
		return fmt.Errorf("configuration error: %w", err)
	}

	// Validate transfer type if provided
	if transferType != "" && transferType != "auto" && transferType != "fast" && transferType != "standard" {
		return fmt.Errorf("invalid transfer type: %s (must be 'fast', 'standard', or 'auto')", transferType)
	}

	// Create the TUI model (wallet loading happens in TUI)
	model := ui.NewModel(cfg)

	// Pre-populate values from flags or config defaults
	if transferAmount == "" && cfg.Defaults.Amount != "" {
		transferAmount = cfg.Defaults.Amount
	}
	if transferRecipient == "" && cfg.Defaults.Recipient != "" {
		transferRecipient = cfg.Defaults.Recipient
	}

	// Apply pre-populated values
	if err := applyTransferPresets(&model, cfg, transferSource, transferDest, transferAmount, transferRecipient, transferType); err != nil {
		return fmt.Errorf("failed to apply transfer presets: %w", err)
	}

	// Run the TUI
	p := tea.NewProgram(
		model,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		return fmt.Errorf("error running program: %w", err)
	}

	return nil
}

func applyTransferPresets(model *ui.Model, cfg *config.Config, source, dest, amount, recipient, transferType string) error {
	chains := cctp.GetChains(cfg.Testnet)
	chains = cctp.ApplyRPCOverrides(chains, cfg.RPCUrls)

	// Find and set source chain
	if source != "" {
		for i := range chains {
			if chains[i].Name == source {
				model.SetSourceChain(&chains[i])
				break
			}
		}
	}

	// Find and set destination chain
	if dest != "" {
		for i := range chains {
			if chains[i].Name == dest {
				model.SetDestChain(&chains[i])
				break
			}
		}
	}

	// If both chains are set, advance to transfer input
	if model.GetSourceChain() != nil && model.GetDestChain() != nil {
		model.AdvanceToTransferInput()
	}

	// Set transfer type if provided
	if transferType != "" && transferType != "auto" {
		var tt cctp.TransferType
		switch transferType {
		case "fast":
			tt = cctp.TransferTypeFast
		case "standard":
			tt = cctp.TransferTypeStandard
		default:
			tt = cctp.TransferTypeAuto
		}
		model.SetTransferType(tt)
	}

	// Set amount if provided
	if amount != "" {
		model.SetTransferAmount(amount)
	}

	// Set recipient if provided and valid
	if recipient != "" {
		if common.IsHexAddress(recipient) {
			model.SetTransferRecipient(recipient)
		}
	}

	return nil
}
