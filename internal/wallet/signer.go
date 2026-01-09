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

package wallet

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// CreateTransactOpts creates transaction options for V2 contract bindings
// This uses bind.NewKeyedTransactor which handles transaction signing automatically
func (w *Wallet) CreateTransactOpts(ctx context.Context, client *ethclient.Client, chainID *big.Int) (*bind.TransactOpts, error) {
	// Create TransactOpts with built-in signing using the wallet's private key
	auth := bind.NewKeyedTransactor(w.PrivateKey, chainID)

	// Set context for the transaction
	auth.Context = ctx

	// Note: Nonce and gas parameters are automatically managed by bind.Transact
	// unless explicitly set here. The V2 system will:
	// - Automatically fetch the nonce from the network
	// - Automatically estimate gas limits
	// - Automatically determine EIP-1559 vs legacy transaction type

	return auth, nil
}

// EstimateGas estimates gas for a transaction
func (w *Wallet) EstimateGas(ctx context.Context, client *ethclient.Client, to common.Address, data []byte, value *big.Int) (uint64, error) {
	// Create call message
	callMsg := ethereum.CallMsg{
		From:  w.Address,
		To:    &to,
		Data:  data,
		Value: value,
	}

	gasLimit, err := client.EstimateGas(ctx, callMsg)
	if err != nil {
		// Try to call the contract to get more detailed error
		_, callErr := client.CallContract(ctx, callMsg, nil)
		if callErr != nil {
			return 0, fmt.Errorf("failed to estimate gas (contract call error: %v): %w", callErr, err)
		}
		return 0, fmt.Errorf("failed to estimate gas: %w", err)
	}

	// Add 20% buffer to gas limit
	gasLimit = gasLimit * 120 / 100
	return gasLimit, nil
}
