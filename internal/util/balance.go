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
