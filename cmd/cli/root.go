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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/pxgray/cctp-go/internal/config"
	"github.com/pxgray/cctp-go/internal/logger"
)

const version = "0.2.0"

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "cctp",
		Short: "CCTP V2 CLI Tool - Transfer USDC across chains with Fast or Standard Transfer",
		Long: `CCTP (Cross-Chain Transfer Protocol) CLI enables you to transfer
USDC seamlessly across multiple blockchain networks using Circle's CCTP V2.

Choose between Fast Transfer (~8-20 seconds, with fee) or Standard Transfer
(~13-19 minutes, no fee) based on your needs.

The CLI provides an interactive terminal UI for managing cross-chain transfers
with support for multiple wallet sources including private keys and keystores.`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Initialize logger (captures standard log output)
			logger.Init()

			// Initialize Viper configuration
			if err := config.InitViper(); err != nil {
				return fmt.Errorf("failed to initialize configuration: %w", err)
			}

			// Bind flags to viper
			if err := viper.BindPFlag("testnet", cmd.Flags().Lookup("testnet")); err != nil {
				return err
			}
			if err := viper.BindPFlag("keystore_path", cmd.Flags().Lookup("keystore")); err != nil {
				return err
			}

			return nil
		},
	}
)

func init() {
	// Global flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cctp/config.yaml)")
	rootCmd.PersistentFlags().Bool("testnet", false, "use testnet instead of mainnet")
	rootCmd.PersistentFlags().String("keystore", "", "path to Ethereum keystore directory")
}

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}

