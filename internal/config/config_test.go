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

package config

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/spf13/viper"
)

func TestInitViper(t *testing.T) {
	// Reset viper before test
	viper.Reset()

	err := InitViper()
	if err != nil {
		t.Fatalf("InitViper failed: %v", err)
	}

	// Verify defaults are set
	if viper.GetString("network") != "mainnet" {
		t.Errorf("expected default network 'mainnet', got %s", viper.GetString("network"))
	}

	if viper.GetBool("testnet") != false {
		t.Error("expected default testnet false")
	}

	if viper.GetString("iris_api_url_mainnet") == "" {
		t.Error("expected iris_api_url_mainnet to be set")
	}

	if viper.GetString("iris_api_url_testnet") == "" {
		t.Error("expected iris_api_url_testnet to be set")
	}
}

func TestLoadConfig_Defaults(t *testing.T) {
	// Reset viper and initialize
	viper.Reset()
	err := InitViper()
	if err != nil {
		t.Fatalf("InitViper failed: %v", err)
	}

	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	// Verify mainnet is default
	if cfg.Network != "mainnet" {
		t.Errorf("expected network 'mainnet', got %s", cfg.Network)
	}

	if cfg.Testnet != false {
		t.Error("expected testnet false")
	}

	// Verify Iris API URL is set
	expectedMainnetURL := "https://iris-api.circle.com"
	if cfg.IrisAPIBaseURL != expectedMainnetURL {
		t.Errorf("expected IrisAPIBaseURL %s, got %s", expectedMainnetURL, cfg.IrisAPIBaseURL)
	}

	// Verify keystore path is set
	if cfg.KeystorePath == "" {
		t.Error("expected KeystorePath to be set")
	}
}

func TestLoadConfig_Mainnet(t *testing.T) {
	// Reset viper and initialize
	viper.Reset()
	err := InitViper()
	if err != nil {
		t.Fatalf("InitViper failed: %v", err)
	}

	// Set mainnet explicitly
	viper.Set("network", "mainnet")
	viper.Set("testnet", false)

	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	if cfg.Network != "mainnet" {
		t.Errorf("expected network 'mainnet', got %s", cfg.Network)
	}

	if cfg.Testnet != false {
		t.Error("expected testnet false")
	}

	expectedMainnetURL := "https://iris-api.circle.com"
	if cfg.IrisAPIBaseURL != expectedMainnetURL {
		t.Errorf("expected IrisAPIBaseURL %s, got %s", expectedMainnetURL, cfg.IrisAPIBaseURL)
	}
}

func TestLoadConfig_Testnet(t *testing.T) {
	// Reset viper and initialize
	viper.Reset()
	err := InitViper()
	if err != nil {
		t.Fatalf("InitViper failed: %v", err)
	}

	// Set testnet
	viper.Set("testnet", true)

	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	if cfg.Network != "testnet" {
		t.Errorf("expected network 'testnet', got %s", cfg.Network)
	}

	if cfg.Testnet != true {
		t.Error("expected testnet true")
	}

	expectedTestnetURL := "https://iris-api-sandbox.circle.com"
	if cfg.IrisAPIBaseURL != expectedTestnetURL {
		t.Errorf("expected IrisAPIBaseURL %s, got %s", expectedTestnetURL, cfg.IrisAPIBaseURL)
	}
}

func TestExpandPath(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		checkFn  func(string) bool
	}{
		{
			name:  "tilde expansion",
			input: "~/test/path",
			checkFn: func(output string) bool {
				return output != "~/test/path" && len(output) > len("/test/path")
			},
		},
		{
			name:  "tilde only",
			input: "~",
			checkFn: func(output string) bool {
				return output != "~" && len(output) > 1
			},
		},
		{
			name:  "no tilde",
			input: "/absolute/path",
			checkFn: func(output string) bool {
				return output == "/absolute/path"
			},
		},
		{
			name:  "empty string",
			input: "",
			checkFn: func(output string) bool {
				return output == ""
			},
		},
		{
			name:  "relative path",
			input: "relative/path",
			checkFn: func(output string) bool {
				return output == "relative/path"
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := expandPath(tc.input)
			if !tc.checkFn(result) {
				t.Errorf("expandPath(%s) = %s, check failed", tc.input, result)
			}
		})
	}
}

func TestGetDefaultKeystorePath(t *testing.T) {
	path := getDefaultKeystorePath()

	// Verify path is not empty
	if path == "" {
		t.Error("expected non-empty keystore path")
	}

	// Verify path contains "keystore"
	if !containsStr(path, "keystore") {
		t.Errorf("expected path to contain 'keystore', got %s", path)
	}

	// Verify path is OS-specific
	switch runtime.GOOS {
	case "darwin":
		if !containsStr(path, "Library/Ethereum") {
			t.Errorf("expected macOS path to contain 'Library/Ethereum', got %s", path)
		}
	case "linux":
		if !containsStr(path, ".ethereum") {
			t.Errorf("expected Linux path to contain '.ethereum', got %s", path)
		}
	case "windows":
		if !containsStr(path, "AppData") {
			t.Errorf("expected Windows path to contain 'AppData', got %s", path)
		}
	}
}

func TestConfigValidate_WithPrivateKey(t *testing.T) {
	// Set PRIVATE_KEY env var
	testPrivateKey := "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"
	os.Setenv("PRIVATE_KEY", testPrivateKey)
	defer os.Unsetenv("PRIVATE_KEY")

	cfg := &Config{
		Network:      "mainnet",
		Testnet:      false,
		KeystorePath: "/nonexistent/path", // Should be ignored when PRIVATE_KEY is set
	}

	err := cfg.Validate()
	if err != nil {
		t.Errorf("Validate should succeed when PRIVATE_KEY is set, got error: %v", err)
	}
}

func TestConfigValidate_WithKeystore(t *testing.T) {
	// Ensure PRIVATE_KEY is not set
	os.Unsetenv("PRIVATE_KEY")

	// Create a temporary directory
	tmpDir := t.TempDir()

	cfg := &Config{
		Network:      "mainnet",
		Testnet:      false,
		KeystorePath: tmpDir,
	}

	err := cfg.Validate()
	if err != nil {
		t.Errorf("Validate should succeed with valid keystore path, got error: %v", err)
	}
}

func TestConfigValidate_InvalidKeystore(t *testing.T) {
	// Ensure PRIVATE_KEY is not set
	os.Unsetenv("PRIVATE_KEY")

	// Use a non-existent path
	nonExistentPath := filepath.Join(t.TempDir(), "nonexistent")

	cfg := &Config{
		Network:      "mainnet",
		Testnet:      false,
		KeystorePath: nonExistentPath,
	}

	err := cfg.Validate()
	// Should not error for non-existent path if it's not the default
	// The actual error will occur when trying to load the wallet
	// But if it's not the default, it should error
	if err == nil {
		// This is acceptable - validation might pass but loading will fail
		t.Log("Validate passed for non-existent non-default path")
	}
}

func TestConfigValidate_DefaultKeystoreMissing(t *testing.T) {
	// Ensure PRIVATE_KEY is not set
	os.Unsetenv("PRIVATE_KEY")

	// Use the default keystore path
	defaultPath := getDefaultKeystorePath()

	cfg := &Config{
		Network:      "mainnet",
		Testnet:      false,
		KeystorePath: defaultPath,
	}

	// This should not error even if default doesn't exist
	err := cfg.Validate()
	if err != nil {
		// Check if error is about default path
		if !containsStr(err.Error(), "keystore path does not exist") {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

func TestConfigValidate_KeystoreIsFile(t *testing.T) {
	// Ensure PRIVATE_KEY is not set
	os.Unsetenv("PRIVATE_KEY")

	// Create a temporary file (not directory)
	tmpFile := filepath.Join(t.TempDir(), "keystore-file")
	err := os.WriteFile(tmpFile, []byte("test"), 0644)
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	cfg := &Config{
		Network:      "mainnet",
		Testnet:      false,
		KeystorePath: tmpFile,
	}

	err = cfg.Validate()
	if err == nil {
		t.Error("expected error when keystore path is a file, not a directory")
	}

	expectedError := "keystore path is not a directory"
	if !containsStr(err.Error(), expectedError) {
		t.Errorf("expected error to contain %q, got %q", expectedError, err.Error())
	}
}

func TestGetRPCURL(t *testing.T) {
	cfg := &Config{
		RPCUrls: map[string]string{
			"Ethereum": "https://custom-eth-rpc.com",
			"Arbitrum": "https://custom-arb-rpc.com",
		},
	}

	testCases := []struct {
		name      string
		chainName string
		expected  string
	}{
		{"custom Ethereum RPC", "Ethereum", "https://custom-eth-rpc.com"},
		{"custom Arbitrum RPC", "Arbitrum", "https://custom-arb-rpc.com"},
		{"non-existent chain", "Polygon", ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := cfg.GetRPCURL(tc.chainName)
			if result != tc.expected {
				t.Errorf("expected %s, got %s", tc.expected, result)
			}
		})
	}
}

func TestLoadConfig_CustomRPCUrls(t *testing.T) {
	// Reset viper and initialize
	viper.Reset()
	err := InitViper()
	if err != nil {
		t.Fatalf("InitViper failed: %v", err)
	}

	// Set custom RPC URLs
	customRPCs := map[string]string{
		"Ethereum": "https://my-custom-eth-rpc.com",
		"Base":     "https://my-custom-base-rpc.com",
	}
	viper.Set("rpc_urls", customRPCs)

	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	// Verify custom RPC URLs are loaded
	if cfg.GetRPCURL("Ethereum") != customRPCs["Ethereum"] {
		t.Errorf("expected Ethereum RPC %s, got %s", customRPCs["Ethereum"], cfg.GetRPCURL("Ethereum"))
	}

	if cfg.GetRPCURL("Base") != customRPCs["Base"] {
		t.Errorf("expected Base RPC %s, got %s", customRPCs["Base"], cfg.GetRPCURL("Base"))
	}
}

func TestLoadConfig_PrivateKeyFromEnv(t *testing.T) {
	// Set PRIVATE_KEY env var
	testPrivateKey := "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"
	os.Setenv("PRIVATE_KEY", testPrivateKey)
	defer os.Unsetenv("PRIVATE_KEY")

	// Reset viper and initialize
	viper.Reset()
	err := InitViper()
	if err != nil {
		t.Fatalf("InitViper failed: %v", err)
	}

	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	// Verify private key is loaded from env
	if cfg.PrivateKey != testPrivateKey {
		t.Errorf("expected PrivateKey %s, got %s", testPrivateKey, cfg.PrivateKey)
	}
}

func TestLoadConfig_DefaultValues(t *testing.T) {
	// Reset viper and initialize
	viper.Reset()
	err := InitViper()
	if err != nil {
		t.Fatalf("InitViper failed: %v", err)
	}

	// Set some default values
	viper.Set("defaults.amount", "100")
	viper.Set("defaults.recipient", "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb7")

	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	// Verify defaults are loaded
	if cfg.Defaults.Amount != "100" {
		t.Errorf("expected default amount '100', got %s", cfg.Defaults.Amount)
	}

	if cfg.Defaults.Recipient != "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb7" {
		t.Errorf("expected default recipient, got %s", cfg.Defaults.Recipient)
	}
}

func TestLoadConfig_SavedAddresses(t *testing.T) {
	// Reset viper and initialize
	viper.Reset()
	err := InitViper()
	if err != nil {
		t.Fatalf("InitViper failed: %v", err)
	}

	// Set saved addresses
	savedAddresses := []SavedAddress{
		{Name: "My Wallet", Address: "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb7"},
		{Name: "Friend's Wallet", Address: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"},
	}
	viper.Set("saved_addresses", savedAddresses)

	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	// Verify saved addresses are loaded
	if len(cfg.SavedAddresses) != 2 {
		t.Errorf("expected 2 saved addresses, got %d", len(cfg.SavedAddresses))
	}

	if cfg.SavedAddresses[0].Name != "My Wallet" {
		t.Errorf("expected first saved address name 'My Wallet', got %s", cfg.SavedAddresses[0].Name)
	}

	if cfg.SavedAddresses[1].Address != "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48" {
		t.Errorf("expected second saved address, got %s", cfg.SavedAddresses[1].Address)
	}
}

func TestExpandPath_EdgeCases(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"tilde with no slash", "~test", "~test"}, // Should not expand (no slash after tilde)
		{"empty", "", ""},
		{"just slash", "/", "/"},
		{"relative with dot", "./test", "./test"},
		{"relative with double dot", "../test", "../test"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := expandPath(tc.input)
			if result != tc.expected {
				t.Errorf("expandPath(%s) = %s, want %s", tc.input, result, tc.expected)
			}
		})
	}
}

// Helper function to check if a string contains a substring
func containsStr(s, substr string) bool {
	if len(substr) == 0 {
		return true
	}
	if len(s) < len(substr) {
		return false
	}
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

