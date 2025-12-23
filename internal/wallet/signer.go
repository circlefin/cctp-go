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
