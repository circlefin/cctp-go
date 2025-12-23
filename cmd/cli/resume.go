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

package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pxgray/cctp-go/internal/config"
	"github.com/pxgray/cctp-go/internal/ui"
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

