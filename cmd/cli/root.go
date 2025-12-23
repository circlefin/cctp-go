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

