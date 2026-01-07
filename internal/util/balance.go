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

package util

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pxgray/cctp-go/tokenmessenger"
)

// GetUSDCBalance fetches the USDC balance for an address on a specific chain
// This is a shared utility used by both the UI and transfer orchestrator
func GetUSDCBalance(ctx context.Context, client *ethclient.Client, usdcAddress, walletAddress common.Address) (*big.Int, error) {
	ierc20 := tokenmessenger.NewIERC20()
	usdcInstance := ierc20.Instance(client, usdcAddress)

	balance, err := bind.Call(usdcInstance, &bind.CallOpts{Context: ctx},
		ierc20.PackBalanceOf(walletAddress), ierc20.UnpackBalanceOf)
	if err != nil {
		return nil, fmt.Errorf("failed to check balance: %w", err)
	}

	return balance, nil
}

// FormatUSDCBalance converts a USDC balance (6 decimals) to a human-readable string
func FormatUSDCBalance(balance *big.Int) string {
	if balance == nil {
		return "0.000000"
	}
	balanceFloat := new(big.Float).SetInt(balance)
	divisor := new(big.Float).SetInt(big.NewInt(1000000)) // 10^6
	balanceFloat.Quo(balanceFloat, divisor)
	return balanceFloat.Text('f', 6)
}
