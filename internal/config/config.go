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

package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

// Config holds the application configuration
type Config struct {
	Network        string
	Testnet        bool
	PrivateKey     string // From PRIVATE_KEY env var
	KeystorePath   string
	IrisAPIBaseURL string
	RPCUrls        map[string]string
	Defaults       DefaultValues
	SavedAddresses []SavedAddress
}

// DefaultValues holds default values for transfer parameters
type DefaultValues struct {
	Amount    string `mapstructure:"amount"`
	Recipient string `mapstructure:"recipient"`
}

// SavedAddress represents a saved recipient address
type SavedAddress struct {
	Name    string `mapstructure:"name"`
	Address string `mapstructure:"address"`
}

// InitViper initializes Viper with default configuration
func InitViper() error {
	// Set config file name and paths
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// Add config paths
	home, err := os.UserHomeDir()
	if err == nil {
		viper.AddConfigPath(filepath.Join(home, ".cctp"))
	}
	viper.AddConfigPath(".")

	// Set environment variable prefix
	viper.SetEnvPrefix("CCTP")
	viper.AutomaticEnv()

	// Set defaults
	viper.SetDefault("network", "mainnet")
	viper.SetDefault("testnet", false)
	viper.SetDefault("keystore_path", getDefaultKeystorePath())
	viper.SetDefault("iris_api_url_mainnet", "https://iris-api.circle.com")
	viper.SetDefault("iris_api_url_testnet", "https://iris-api-sandbox.circle.com")
	viper.SetDefault("rpc_urls", map[string]string{})
	viper.SetDefault("defaults.amount", "")
	viper.SetDefault("defaults.recipient", "")
	viper.SetDefault("saved_addresses", []SavedAddress{})

	// Try to read config file (it's okay if it doesn't exist)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			// Config file was found but another error occurred
			return fmt.Errorf("error reading config file: %w", err)
		}
		// Config file not found; ignore error
	}

	return nil
}

// LoadConfig loads configuration from Viper
func LoadConfig() (*Config, error) {
	cfg := &Config{
		Network:      viper.GetString("network"),
		Testnet:      viper.GetBool("testnet"),
		PrivateKey:   os.Getenv("PRIVATE_KEY"), // Load from env
		KeystorePath: expandPath(viper.GetString("keystore_path")),
		RPCUrls:      viper.GetStringMapString("rpc_urls"),
	}

	// Determine network and Iris API URL
	if cfg.Testnet || cfg.Network == "testnet" {
		cfg.Testnet = true
		cfg.Network = "testnet"
		cfg.IrisAPIBaseURL = viper.GetString("iris_api_url_testnet")
	} else {
		cfg.Testnet = false
		cfg.Network = "mainnet"
		cfg.IrisAPIBaseURL = viper.GetString("iris_api_url_mainnet")
	}

	// Load default values
	cfg.Defaults = DefaultValues{
		Amount:    viper.GetString("defaults.amount"),
		Recipient: viper.GetString("defaults.recipient"),
	}

	// Load saved addresses
	if err := viper.UnmarshalKey("saved_addresses", &cfg.SavedAddresses); err != nil {
		// It's okay if this fails, just use empty list
		cfg.SavedAddresses = []SavedAddress{}
	}

	return cfg, nil
}

// expandPath expands ~ in file paths to the user's home directory
func expandPath(path string) string {
	if path == "" {
		return path
	}

	if len(path) > 0 && path[0] == '~' {
		home, err := os.UserHomeDir()
		if err != nil {
			return path
		}

		if len(path) == 1 {
			return home
		}

		if path[1] == '/' || path[1] == filepath.Separator {
			return filepath.Join(home, path[2:])
		}
	}

	return path
}

// getDefaultKeystorePath returns the default Ethereum keystore path
// based on the operating system (matches geth defaults)
func getDefaultKeystorePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}

	var keystorePath string
	switch runtime.GOOS {
	case "darwin": // macOS
		keystorePath = filepath.Join(home, "Library", "Ethereum", "keystore")
	case "linux":
		keystorePath = filepath.Join(home, ".ethereum", "keystore")
	case "windows":
		keystorePath = filepath.Join(home, "AppData", "Roaming", "Ethereum", "keystore")
	default:
		// Default to Linux path for other Unix-like systems
		keystorePath = filepath.Join(home, ".ethereum", "keystore")
	}

	return keystorePath
}

// Validate validates the configuration
// If PRIVATE_KEY env var is set, keystore validation is skipped
func (c *Config) Validate() error {
	// If PRIVATE_KEY is set, we don't need to validate keystore
	if os.Getenv("PRIVATE_KEY") != "" {
		return nil
	}

	// Only validate keystore if we'll actually use it
	if c.KeystorePath != "" {
		info, err := os.Stat(c.KeystorePath)
		if err != nil {
			if os.IsNotExist(err) {
				// If it's the default path and doesn't exist, that's okay
				// The user will get a better error message when trying to load the wallet
				defaultPath := getDefaultKeystorePath()
				if c.KeystorePath == defaultPath {
					return nil
				}
				return fmt.Errorf("keystore path does not exist: %s", c.KeystorePath)
			}
			return fmt.Errorf("error accessing keystore path: %w", err)
		}
		if !info.IsDir() {
			return fmt.Errorf("keystore path is not a directory: %s", c.KeystorePath)
		}
	}
	return nil
}

// GetRPCURL returns the RPC URL for a given chain name
func (c *Config) GetRPCURL(chainName string) string {
	if url, ok := c.RPCUrls[chainName]; ok {
		return url
	}
	return ""
}
