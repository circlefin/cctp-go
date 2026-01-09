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
