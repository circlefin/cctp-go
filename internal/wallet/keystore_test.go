package wallet

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
)

func TestLoadFromPrivateKey(t *testing.T) {
	// Test private key (DO NOT use in production)
	testPrivateKey := "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"
	// The expected address derived from this private key
	expectedAddress := "0xFCAd0B19bB29D4674531d6f115237E16AfCE377c"

	wallet, err := LoadFromPrivateKey(testPrivateKey)
	if err != nil {
		t.Fatalf("LoadFromPrivateKey failed: %v", err)
	}

	if wallet == nil {
		t.Fatal("expected non-nil wallet")
	}

	if wallet.PrivateKey == nil {
		t.Fatal("expected non-nil private key")
	}

	// Verify address matches expected
	if wallet.Address.Hex() != expectedAddress {
		t.Errorf("expected address %s, got %s", expectedAddress, wallet.Address.Hex())
	}
}

func TestLoadFromPrivateKey_WithPrefix(t *testing.T) {
	testPrivateKey := "0x0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"
	expectedAddress := "0xFCAd0B19bB29D4674531d6f115237E16AfCE377c"

	wallet, err := LoadFromPrivateKey(testPrivateKey)
	if err != nil {
		t.Fatalf("LoadFromPrivateKey with 0x prefix failed: %v", err)
	}

	if wallet.Address.Hex() != expectedAddress {
		t.Errorf("expected address %s, got %s", expectedAddress, wallet.Address.Hex())
	}
}

func TestLoadFromPrivateKey_Invalid(t *testing.T) {
	testCases := []struct {
		name       string
		privateKey string
	}{
		{"too short", "0123456789abcdef"},
		{"too long", "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef00"},
		{"invalid hex", "0123456789abcdefghijklmnopqrstuvwxyz0123456789abcdef0123456789ab"},
		{"empty", ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := LoadFromPrivateKey(tc.privateKey)
			if err == nil {
				t.Error("expected error for invalid private key")
			}
		})
	}
}

func TestLoadFromEnv(t *testing.T) {
	testPrivateKey := "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"
	expectedAddress := "0xFCAd0B19bB29D4674531d6f115237E16AfCE377c"

	// Set environment variable
	os.Setenv("PRIVATE_KEY", testPrivateKey)
	defer os.Unsetenv("PRIVATE_KEY")

	wallet, err := LoadFromEnv()
	if err != nil {
		t.Fatalf("LoadFromEnv failed: %v", err)
	}

	if wallet.Address.Hex() != expectedAddress {
		t.Errorf("expected address %s, got %s", expectedAddress, wallet.Address.Hex())
	}
}

func TestLoadFromEnv_NotSet(t *testing.T) {
	// Ensure PRIVATE_KEY is not set
	os.Unsetenv("PRIVATE_KEY")

	_, err := LoadFromEnv()
	if err == nil {
		t.Fatal("expected error when PRIVATE_KEY not set")
	}

	expectedError := "PRIVATE_KEY environment variable not set"
	if err.Error() != expectedError {
		t.Errorf("expected error %q, got %q", expectedError, err.Error())
	}
}

func TestLoadFromKeystore(t *testing.T) {
	// Create a temporary keystore directory
	tmpDir := t.TempDir()

	// Create a test keystore
	ks := keystore.NewKeyStore(tmpDir, keystore.StandardScryptN, keystore.StandardScryptP)
	password := "test-password"

	// Create a new account
	account, err := ks.NewAccount(password)
	if err != nil {
		t.Fatalf("failed to create test account: %v", err)
	}

	// Load the wallet
	wallet, err := LoadFromKeystore(tmpDir, password)
	if err != nil {
		t.Fatalf("LoadFromKeystore failed: %v", err)
	}

	// Verify the address matches
	if wallet.Address != account.Address {
		t.Errorf("expected address %s, got %s", account.Address.Hex(), wallet.Address.Hex())
	}

	// Verify private key is set
	if wallet.PrivateKey == nil {
		t.Error("expected non-nil private key")
	}
}

func TestLoadFromKeystore_WrongPassword(t *testing.T) {
	// Create a temporary keystore directory
	tmpDir := t.TempDir()

	// Create a test keystore
	ks := keystore.NewKeyStore(tmpDir, keystore.StandardScryptN, keystore.StandardScryptP)
	password := "test-password"

	// Create a new account
	_, err := ks.NewAccount(password)
	if err != nil {
		t.Fatalf("failed to create test account: %v", err)
	}

	// Try to load with wrong password
	_, err = LoadFromKeystore(tmpDir, "wrong-password")
	if err == nil {
		t.Fatal("expected error for wrong password")
	}

	expectedError := "failed to unlock account"
	if !containsString(err.Error(), expectedError) {
		t.Errorf("expected error to contain %q, got %q", expectedError, err.Error())
	}
}

func TestLoadFromKeystore_NotFound(t *testing.T) {
	// Use a non-existent directory
	nonExistentDir := filepath.Join(t.TempDir(), "nonexistent")

	_, err := LoadFromKeystore(nonExistentDir, "password")
	if err == nil {
		t.Fatal("expected error for non-existent keystore")
	}

	expectedError := "keystore path does not exist"
	if !containsString(err.Error(), expectedError) {
		t.Errorf("expected error to contain %q, got %q", expectedError, err.Error())
	}
}

func TestLoadFromKeystore_NoAccounts(t *testing.T) {
	// Create an empty keystore directory
	tmpDir := t.TempDir()

	_, err := LoadFromKeystore(tmpDir, "password")
	if err == nil {
		t.Fatal("expected error for empty keystore")
	}

	expectedError := "no accounts found in keystore"
	if !containsString(err.Error(), expectedError) {
		t.Errorf("expected error to contain %q, got %q", expectedError, err.Error())
	}
}

func TestListKeystoreAccounts(t *testing.T) {
	// Create a temporary keystore directory
	tmpDir := t.TempDir()

	// Create a test keystore with multiple accounts
	ks := keystore.NewKeyStore(tmpDir, keystore.StandardScryptN, keystore.StandardScryptP)
	password := "test-password"

	// Create accounts
	account1, err := ks.NewAccount(password)
	if err != nil {
		t.Fatalf("failed to create account 1: %v", err)
	}

	account2, err := ks.NewAccount(password)
	if err != nil {
		t.Fatalf("failed to create account 2: %v", err)
	}

	// List accounts
	addresses, err := ListKeystoreAccounts(tmpDir)
	if err != nil {
		t.Fatalf("ListKeystoreAccounts failed: %v", err)
	}

	// Verify we got 2 accounts
	if len(addresses) != 2 {
		t.Errorf("expected 2 accounts, got %d", len(addresses))
	}

	// Verify the addresses match
	expectedAddresses := map[common.Address]bool{
		account1.Address: true,
		account2.Address: true,
	}

	for _, addr := range addresses {
		if !expectedAddresses[addr] {
			t.Errorf("unexpected address: %s", addr.Hex())
		}
	}
}

func TestListKeystoreAccounts_NotFound(t *testing.T) {
	nonExistentDir := filepath.Join(t.TempDir(), "nonexistent")

	_, err := ListKeystoreAccounts(nonExistentDir)
	if err == nil {
		t.Fatal("expected error for non-existent keystore")
	}
}

func TestSelectAccountFromKeystore(t *testing.T) {
	// Create a temporary keystore directory
	tmpDir := t.TempDir()

	// Create a test keystore with multiple accounts
	ks := keystore.NewKeyStore(tmpDir, keystore.StandardScryptN, keystore.StandardScryptP)
	password := "test-password"

	// Create accounts
	_, err := ks.NewAccount(password)
	if err != nil {
		t.Fatalf("failed to create account 1: %v", err)
	}

	account2, err := ks.NewAccount(password)
	if err != nil {
		t.Fatalf("failed to create account 2: %v", err)
	}

	// Select the second account
	wallet, err := SelectAccountFromKeystore(tmpDir, account2.Address.Hex(), password)
	if err != nil {
		t.Fatalf("SelectAccountFromKeystore failed: %v", err)
	}

	// Verify the address matches
	if wallet.Address != account2.Address {
		t.Errorf("expected address %s, got %s", account2.Address.Hex(), wallet.Address.Hex())
	}

	// Verify private key is set
	if wallet.PrivateKey == nil {
		t.Error("expected non-nil private key")
	}
}

func TestSelectAccountFromKeystore_NotFound(t *testing.T) {
	// Create a temporary keystore directory
	tmpDir := t.TempDir()

	// Create a test keystore
	ks := keystore.NewKeyStore(tmpDir, keystore.StandardScryptN, keystore.StandardScryptP)
	password := "test-password"

	// Create an account
	_, err := ks.NewAccount(password)
	if err != nil {
		t.Fatalf("failed to create account: %v", err)
	}

	// Try to select a non-existent account
	nonExistentAddr := "0x0000000000000000000000000000000000000001"
	_, err = SelectAccountFromKeystore(tmpDir, nonExistentAddr, password)
	if err == nil {
		t.Fatal("expected error for non-existent account")
	}

	expectedError := "account not found"
	if !containsString(err.Error(), expectedError) {
		t.Errorf("expected error to contain %q, got %q", expectedError, err.Error())
	}
}

func TestFindKeystoreFiles(t *testing.T) {
	// Create a temporary keystore directory
	tmpDir := t.TempDir()

	// Create a test keystore
	ks := keystore.NewKeyStore(tmpDir, keystore.StandardScryptN, keystore.StandardScryptP)
	password := "test-password"

	// Create an account (this creates a keystore file)
	_, err := ks.NewAccount(password)
	if err != nil {
		t.Fatalf("failed to create account: %v", err)
	}

	// Find keystore files
	files, err := FindKeystoreFiles(tmpDir)
	if err != nil {
		t.Fatalf("FindKeystoreFiles failed: %v", err)
	}

	// Should have at least 1 file
	if len(files) == 0 {
		t.Error("expected at least 1 keystore file")
	}

	// Verify files exist
	for _, file := range files {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			t.Errorf("file does not exist: %s", file)
		}
	}
}

func TestFindKeystoreFiles_NotFound(t *testing.T) {
	nonExistentDir := filepath.Join(t.TempDir(), "nonexistent")

	_, err := FindKeystoreFiles(nonExistentDir)
	if err == nil {
		t.Fatal("expected error for non-existent directory")
	}
}

func TestFindKeystoreFiles_EmptyDirectory(t *testing.T) {
	// Create an empty directory
	tmpDir := t.TempDir()

	files, err := FindKeystoreFiles(tmpDir)
	if err != nil {
		t.Fatalf("FindKeystoreFiles failed: %v", err)
	}

	// Should return empty list for empty directory
	if len(files) != 0 {
		t.Errorf("expected 0 files, got %d", len(files))
	}
}

func TestPrivateKeyDerivation(t *testing.T) {
	// Test that the same private key always produces the same address
	testPrivateKey := "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"

	wallet1, err := LoadFromPrivateKey(testPrivateKey)
	if err != nil {
		t.Fatalf("failed to load wallet 1: %v", err)
	}

	wallet2, err := LoadFromPrivateKey(testPrivateKey)
	if err != nil {
		t.Fatalf("failed to load wallet 2: %v", err)
	}

	if wallet1.Address != wallet2.Address {
		t.Errorf("addresses should match: %s != %s", wallet1.Address.Hex(), wallet2.Address.Hex())
	}
}

// Helper function to check if a string contains a substring
func containsString(s, substr string) bool {
	return len(s) >= len(substr) && containsStringHelper(s, substr)
}

func containsStringHelper(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

