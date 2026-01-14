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
	"github.com/circlefin/cctp-go/internal/config"
	"github.com/circlefin/cctp-go/internal/ui"
	"github.com/spf13/cobra"
)

var resumeTxHash string

var resumeCmd = &cobra.Command{
	Use:   "resume [tx-hash]",
	Short: "Resume an existing transfer",
	Long: `Resume an existing cross-chain USDC transfer using the burn transaction hash.

You can provide the transaction hash as a positional argument or use the --tx-hash flag.
If not provided, you will be prompted in the TUI.`,
	Example: `  # Resume with tx hash as argument
  cctp resume 0x1234567890abcdef...

  # Resume with tx hash as flag
  cctp resume --tx-hash 0x1234567890abcdef...

  # Start resume flow (will prompt for tx hash)
  cctp resume`,
	Args: cobra.MaximumNArgs(1),
	RunE: runResume,
}

func init() {
	resumeCmd.Flags().StringVar(&resumeTxHash, "tx-hash", "", "burn transaction hash to resume")

	rootCmd.AddCommand(resumeCmd)
}

func runResume(cmd *cobra.Command, args []string) error {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	// Validate configuration
	if err := cfg.Validate(); err != nil {
		return fmt.Errorf("configuration error: %w", err)
	}

	// Get tx hash from args if provided
	if len(args) > 0 {
		resumeTxHash = args[0]
	}

	// Create the TUI model in resume mode (wallet loading happens in TUI)
	model := ui.NewModel(cfg)
	model.SetResumeMode(resumeTxHash)

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
