package wallet

import (
	"crypto/ecdsa"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Wallet represents an Ethereum wallet
type Wallet struct {
	Address    common.Address
	PrivateKey *ecdsa.PrivateKey
	keystore   *keystore.KeyStore
	account    *accounts.Account
}

// LoadFromKeystore loads a wallet from a keystore directory
func LoadFromKeystore(keystorePath, password string) (*Wallet, error) {
	// Check if keystore path exists
	if _, err := os.Stat(keystorePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("keystore path does not exist: %s", keystorePath)
	}

	// Create keystore instance
	ks := keystore.NewKeyStore(keystorePath, keystore.StandardScryptN, keystore.StandardScryptP)
	
	// Get all accounts
	accounts := ks.Accounts()
	if len(accounts) == 0 {
		return nil, fmt.Errorf("no accounts found in keystore: %s", keystorePath)
	}

	// Use the first account
	account := accounts[0]
	
	// Unlock the account
	err := ks.Unlock(account, password)
	if err != nil {
		return nil, fmt.Errorf("failed to unlock account: %w", err)
	}

	// Get the private key by exporting and re-importing
	// Note: This is a workaround since KeyStore doesn't expose private keys directly
	jsonKey, err := ks.Export(account, password, password)
	if err != nil {
		return nil, fmt.Errorf("failed to export key: %w", err)
	}

	key, err := keystore.DecryptKey(jsonKey, password)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt key: %w", err)
	}

	return &Wallet{
		Address:    account.Address,
		PrivateKey: key.PrivateKey,
		keystore:   ks,
		account:    &account,
	}, nil
}

// LoadFromPrivateKey loads a wallet from a private key hex string
func LoadFromPrivateKey(privateKeyHex string) (*Wallet, error) {
	// Remove 0x prefix if present
	if len(privateKeyHex) >= 2 && privateKeyHex[:2] == "0x" {
		privateKeyHex = privateKeyHex[2:]
	}

	// Parse private key
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	// Derive public key and address
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("failed to cast public key to ECDSA")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	return &Wallet{
		Address:    address,
		PrivateKey: privateKey,
	}, nil
}

// LoadFromEnv loads a wallet from the PRIVATE_KEY environment variable
func LoadFromEnv() (*Wallet, error) {
	privateKey := os.Getenv("PRIVATE_KEY")
	if privateKey == "" {
		return nil, fmt.Errorf("PRIVATE_KEY environment variable not set")
	}
	return LoadFromPrivateKey(privateKey)
}

// ListKeystoreAccounts lists all accounts in a keystore directory
func ListKeystoreAccounts(keystorePath string) ([]common.Address, error) {
	if _, err := os.Stat(keystorePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("keystore path does not exist: %s", keystorePath)
	}

	ks := keystore.NewKeyStore(keystorePath, keystore.StandardScryptN, keystore.StandardScryptP)
	accounts := ks.Accounts()
	
	addresses := make([]common.Address, len(accounts))
	for i, account := range accounts {
		addresses[i] = account.Address
	}
	
	return addresses, nil
}

// SelectAccountFromKeystore loads a specific account from keystore by address
func SelectAccountFromKeystore(keystorePath, addressHex, password string) (*Wallet, error) {
	if _, err := os.Stat(keystorePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("keystore path does not exist: %s", keystorePath)
	}

	ks := keystore.NewKeyStore(keystorePath, keystore.StandardScryptN, keystore.StandardScryptP)
	
	// Parse target address
	targetAddress := common.HexToAddress(addressHex)
	
	// Find the account
	account, err := ks.Find(accounts.Account{Address: targetAddress})
	if err != nil {
		return nil, fmt.Errorf("account not found: %w", err)
	}

	// Unlock the account
	err = ks.Unlock(account, password)
	if err != nil {
		return nil, fmt.Errorf("failed to unlock account: %w", err)
	}

	// Export and decrypt to get private key
	jsonKey, err := ks.Export(account, password, password)
	if err != nil {
		return nil, fmt.Errorf("failed to export key: %w", err)
	}

	key, err := keystore.DecryptKey(jsonKey, password)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt key: %w", err)
	}

	return &Wallet{
		Address:    account.Address,
		PrivateKey: key.PrivateKey,
		keystore:   ks,
		account:    &account,
	}, nil
}

// FindKeystoreFiles finds all keystore files in a directory
func FindKeystoreFiles(keystorePath string) ([]string, error) {
	if _, err := os.Stat(keystorePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("keystore path does not exist: %s", keystorePath)
	}

	var keystoreFiles []string
	err := filepath.Walk(keystorePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Mode().IsRegular() {
			// Keystore files typically start with "UTC--"
			if len(info.Name()) > 4 {
				keystoreFiles = append(keystoreFiles, path)
			}
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to scan keystore directory: %w", err)
	}

	return keystoreFiles, nil
}

