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

package cctp

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/charmbracelet/log"
	"github.com/circlefin/cctp-go/internal/util"
	"github.com/circlefin/cctp-go/internal/wallet"
	"github.com/circlefin/cctp-go/messagetransmitter"
	"github.com/circlefin/cctp-go/tokenmessenger"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// TransferStep represents a step in the transfer process
type TransferStep string

const (
	StepCheckBalance       TransferStep = "checking_balance"
	StepCheckAllowance     TransferStep = "checking_allowance"
	StepApproving          TransferStep = "approving"
	StepApprovingWait      TransferStep = "approving_wait"
	StepBurning            TransferStep = "burning"
	StepBurningWait        TransferStep = "burning_wait"
	StepPollingAttestation TransferStep = "polling_attestation"
	StepMinting            TransferStep = "minting"
	StepMintingWait        TransferStep = "minting_wait"
	StepComplete           TransferStep = "complete"
	StepError              TransferStep = "error"
)

// TransferType represents the type of CCTP transfer
type TransferType string

const (
	TransferTypeFast     TransferType = "fast"     // Fast Transfer (~8-20 seconds, with fee)
	TransferTypeStandard TransferType = "standard" // Standard Transfer (~13-19 minutes, no fee)
	TransferTypeAuto     TransferType = "auto"     // Auto-select based on chain capabilities
)

// TransferUpdate represents an update in the transfer process
type TransferUpdate struct {
	Step        TransferStep
	Message     string
	TxHash      string
	Error       error
	SourceChain *Chain
	DestChain   *Chain
}

// TransferParams holds parameters for a transfer
type TransferParams struct {
	SourceChain      *Chain
	DestChain        *Chain
	Amount           *big.Int
	RecipientAddress common.Address
	Testnet          bool
	TransferType     TransferType // Fast, Standard, or Auto (default: Auto)
	CachedBalance    *big.Int     // Optional: pre-fetched balance to skip redundant check
}

// TransferOrchestrator orchestrates the entire transfer workflow
type TransferOrchestrator struct {
	wallet       *wallet.Wallet
	irisClient   *IrisClient
	sourceClient *ethclient.Client
	destClient   *ethclient.Client
	params       *TransferParams
}

// NewTransferOrchestrator creates a new transfer orchestrator
func NewTransferOrchestrator(
	w *wallet.Wallet,
	params *TransferParams,
	irisBaseURL string,
) (*TransferOrchestrator, error) {
	// Create Iris client
	irisClient := NewIrisClient(irisBaseURL)

	// Create RPC clients
	sourceClient, err := ethclient.Dial(params.SourceChain.RPC)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to source chain: %w", err)
	}

	destClient, err := ethclient.Dial(params.DestChain.RPC)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to destination chain: %w", err)
	}

	return &TransferOrchestrator{
		wallet:       w,
		irisClient:   irisClient,
		sourceClient: sourceClient,
		destClient:   destClient,
		params:       params,
	}, nil
}

// Execute executes the complete transfer workflow
func (t *TransferOrchestrator) Execute(ctx context.Context, updates chan<- TransferUpdate) error {
	defer close(updates)

	// Step 1: Check USDC balance (use cached if available)
	var balance *big.Int
	if t.params.CachedBalance != nil {
		// Use cached balance from UI to avoid redundant RPC call
		balance = t.params.CachedBalance
		updates <- TransferUpdate{Step: StepCheckBalance, Message: fmt.Sprintf("Balance: %s USDC (cached)", util.FormatUSDCBalance(balance))}
	} else {
		// Fetch balance if not cached
		updates <- TransferUpdate{Step: StepCheckBalance, Message: "Checking USDC balance..."}

		usdcAddress := common.HexToAddress(t.params.SourceChain.USDC)
		var err error
		balance, err = util.GetUSDCBalance(ctx, t.sourceClient, usdcAddress, t.wallet.Address)
		if err != nil {
			updates <- TransferUpdate{Step: StepError, Error: err}
			return err
		}

		updates <- TransferUpdate{Step: StepCheckBalance, Message: fmt.Sprintf("Balance: %s USDC", util.FormatUSDCBalance(balance))}
	}

	// Validate balance regardless of whether it was cached or fetched
	if balance.Cmp(t.params.Amount) < 0 {
		err := fmt.Errorf("insufficient USDC balance: have %s, need %s", util.FormatUSDCBalance(balance), util.FormatUSDCBalance(t.params.Amount))
		updates <- TransferUpdate{Step: StepError, Error: err}
		return err
	}

	// Step 2: Check allowance
	updates <- TransferUpdate{Step: StepCheckAllowance, Message: "Checking USDC allowance..."}
	tokenMessengerAddr := common.HexToAddress(t.params.SourceChain.TokenMessengerV2)

	// Create USDC contract instance for allowance and approval operations
	ierc20 := tokenmessenger.NewIERC20()
	usdcAddress := common.HexToAddress(t.params.SourceChain.USDC)
	usdcInstance := ierc20.Instance(t.sourceClient, usdcAddress)

	// Use V2 bindings to check allowance
	allowance, err := bind.Call(usdcInstance, &bind.CallOpts{Context: ctx},
		ierc20.PackAllowance(t.wallet.Address, tokenMessengerAddr), ierc20.UnpackAllowance)
	if err != nil {
		updates <- TransferUpdate{Step: StepError, Error: fmt.Errorf("failed to check allowance: %w", err)}
		return err
	}

	// Step 3: Approve if needed
	if allowance.Cmp(t.params.Amount) < 0 {
		updates <- TransferUpdate{Step: StepApproving, Message: "Approving USDC..."}

		// Create transaction options
		auth, err := t.wallet.CreateTransactOpts(ctx, t.sourceClient, t.params.SourceChain.ChainID)
		if err != nil {
			updates <- TransferUpdate{Step: StepError, Error: fmt.Errorf("failed to create auth: %w", err)}
			return err
		}

		// Use V2 bindings to send approve transaction
		approveTx, err := bind.Transact(usdcInstance, auth, ierc20.PackApprove(tokenMessengerAddr, t.params.Amount))
		if err != nil {
			updates <- TransferUpdate{Step: StepError, Error: fmt.Errorf("failed to send approve tx: %w", err)}
			return err
		}

		updates <- TransferUpdate{
			Step:    StepApprovingWait,
			Message: "Waiting for approval confirmation...",
			TxHash:  approveTx.Hash().Hex(),
		}

		// Wait for approval confirmation using V2 helper
		_, err = bind.WaitMined(ctx, t.sourceClient, approveTx.Hash())
		if err != nil {
			updates <- TransferUpdate{Step: StepError, Error: fmt.Errorf("approval transaction failed: %w", err)}
			return err
		}

		updates <- TransferUpdate{
			Step:    StepApproving,
			Message: "USDC approved successfully",
			TxHash:  approveTx.Hash().Hex(),
		}

		// Verify the allowance was updated
		newAllowance, err := bind.Call(usdcInstance, &bind.CallOpts{Context: ctx},
			ierc20.PackAllowance(t.wallet.Address, tokenMessengerAddr), ierc20.UnpackAllowance)
		if err != nil {
			updates <- TransferUpdate{Step: StepError, Error: fmt.Errorf("failed to verify allowance: %w", err)}
			return err
		}
		if newAllowance.Cmp(t.params.Amount) < 0 {
			updates <- TransferUpdate{Step: StepError, Error: fmt.Errorf("allowance not updated correctly: have %s, need %s", newAllowance.String(), t.params.Amount.String())}
			return fmt.Errorf("allowance verification failed")
		}
	}

	// Step 4: Determine transfer type and fetch fees
	transferType := t.params.TransferType
	if transferType == "" || transferType == TransferTypeAuto {
		// Auto-select: use Fast if available, otherwise Standard
		if t.params.SourceChain.InstantFinality {
			transferType = TransferTypeStandard
			log.Info("Auto-selected Standard Transfer (instant finality chain)")
		} else {
			transferType = TransferTypeFast
			log.Info("Auto-selected Fast Transfer")
		}
	}

	// Determine finality threshold based on transfer type
	var minFinalityThreshold uint32
	var transferTypeLabel string
	if transferType == TransferTypeFast {
		minFinalityThreshold = 1000
		transferTypeLabel = "Fast Transfer"
	} else {
		minFinalityThreshold = 2000
		transferTypeLabel = "Standard Transfer"
	}

	updates <- TransferUpdate{
		Step:    StepBurning,
		Message: fmt.Sprintf("Querying fees for %s...", transferTypeLabel),
	}

	// Fetch fee information from API
	feesResp, err := t.irisClient.GetTransferFees(ctx, t.params.SourceChain.Domain, t.params.DestChain.Domain)
	if err != nil {
		log.Warn("Failed to fetch fees from API, using fallback", "error", err)
		// Fallback to conservative maxFee if API fails
		feesResp = &FeesResponse{
			Data: []FeeInfo{
				{FinalityThreshold: 1000, MinimumFee: 14}, // Conservative: highest known fee (Linea)
				{FinalityThreshold: 2000, MinimumFee: 0},
			},
		}
	}

	// Find the appropriate fee for the selected finality threshold
	var feeBps uint32 = 0
	for _, feeInfo := range feesResp.Data {
		if feeInfo.FinalityThreshold == minFinalityThreshold {
			feeBps = feeInfo.MinimumFee
			break
		}
	}

	// Calculate maxFee based on fee rate (with 1 bps safety buffer)
	// maxFee = (amount * (feeBps + 1)) / 10000
	var maxFee *big.Int
	if feeBps > 0 {
		// Add 1 bps buffer for fee fluctuations
		feeWithBuffer := big.NewInt(int64(feeBps + 1))
		maxFee = new(big.Int).Mul(t.params.Amount, feeWithBuffer)
		maxFee = maxFee.Div(maxFee, big.NewInt(10000))
		// Ensure minimum of 1 unit
		if maxFee.Cmp(big.NewInt(0)) == 0 {
			maxFee = big.NewInt(1)
		}
	} else {
		// Standard Transfer with no fee - still need non-zero maxFee for contract
		maxFee = big.NewInt(1)
	}

	// Step 5: Burn USDC
	updates <- TransferUpdate{
		Step:    StepBurning,
		Message: fmt.Sprintf("Burning USDC on source chain (%s)...", transferTypeLabel),
	}

	log.Info("Using transfer mode", "mode", transferTypeLabel, "threshold", minFinalityThreshold, "fee_bps", feeBps)

	// Convert recipient address to bytes32
	recipientBytes32 := util.AddressToBytes32(t.params.RecipientAddress)

	log.Debug("=== BURN PARAMETERS ===")
	log.Debug("Transfer parameters", "type", transferTypeLabel, "amount", t.params.Amount.String(), "dest_domain", t.params.DestChain.Domain)
	log.Debug("Recipient info", "address", t.params.RecipientAddress.Hex(), "bytes32", fmt.Sprintf("%x", recipientBytes32))
	log.Debug("Token and fees", "burn_token", t.params.SourceChain.USDC, "max_fee", maxFee.String(), "fee_bps", feeBps)
	log.Debug("Finality", "min_finality_threshold", minFinalityThreshold)
	log.Debug("======================")

	// Create V2 token messenger instance
	tokenMessengerV2 := tokenmessenger.NewTokenMessengerV2()
	tokenMessengerInstance := tokenMessengerV2.Instance(t.sourceClient, tokenMessengerAddr)

	// Pack the depositForBurn call
	burnData := tokenMessengerV2.PackDepositForBurn(
		t.params.Amount,
		t.params.DestChain.Domain,
		recipientBytes32,
		common.HexToAddress(t.params.SourceChain.USDC),
		[32]byte{},           // empty destinationCaller (bytes32(0) = any address can call)
		maxFee,               // maxFee calculated based on fee API response
		minFinalityThreshold, // 1000 = Fast Transfer, 2000 = Standard Transfer
	)

	// Try to call the contract first to get detailed error
	callMsg := ethereum.CallMsg{
		From:  t.wallet.Address,
		To:    &tokenMessengerAddr,
		Data:  burnData,
		Value: big.NewInt(0),
	}

	_, callErr := t.sourceClient.CallContract(ctx, callMsg, nil)
	if callErr != nil {
		// Log detailed debug info
		log.Error("=== BURN TRANSACTION FAILED ===")
		log.Error("Burn transaction would fail", "error", callErr)
		log.Error("Chain info", "source", t.params.SourceChain.Name, "source_domain", t.params.SourceChain.Domain, "dest", t.params.DestChain.Name, "dest_domain", t.params.DestChain.Domain)
		log.Error("Transaction details", "amount", t.params.Amount.String(), "recipient", t.params.RecipientAddress.Hex())
		log.Error("Contract addresses", "usdc", t.params.SourceChain.USDC, "token_messenger", tokenMessengerAddr.Hex(), "wallet", t.wallet.Address.Hex())
		log.Error("Fee parameters", "max_fee", maxFee.String(), "min_finality_threshold", minFinalityThreshold)
		log.Error("Call data", "data", fmt.Sprintf("%x", burnData))
		log.Error("================================")

		updates <- TransferUpdate{
			Step:  StepError,
			Error: fmt.Errorf("burn call would fail: %w", callErr),
		}
		return callErr
	}

	// Create transaction options for burn
	auth, err := t.wallet.CreateTransactOpts(ctx, t.sourceClient, t.params.SourceChain.ChainID)
	if err != nil {
		updates <- TransferUpdate{Step: StepError, Error: fmt.Errorf("failed to create auth: %w", err)}
		return err
	}

	// Use V2 bindings to send depositForBurn transaction
	burnTx, err := bind.Transact(tokenMessengerInstance, auth, burnData)
	if err != nil {
		updates <- TransferUpdate{Step: StepError, Error: fmt.Errorf("failed to send burn tx: %w", err)}
		return err
	}

	updates <- TransferUpdate{
		Step:    StepBurningWait,
		Message: "Waiting for burn confirmation...",
		TxHash:  burnTx.Hash().Hex(),
	}

	// Wait for burn confirmation using V2 helper
	_, err = bind.WaitMined(ctx, t.sourceClient, burnTx.Hash())
	if err != nil {
		updates <- TransferUpdate{Step: StepError, Error: fmt.Errorf("burn transaction failed: %w", err)}
		return err
	}

	updates <- TransferUpdate{
		Step:    StepBurning,
		Message: fmt.Sprintf("USDC burned successfully (%s)", transferTypeLabel),
		TxHash:  burnTx.Hash().Hex(),
	}

	// Step 6: Poll for attestation
	var estimatedTime string
	if transferType == TransferTypeFast {
		estimatedTime = "~8-20 seconds"
	} else {
		estimatedTime = "~13-19 minutes"
	}
	updates <- TransferUpdate{
		Step:    StepPollingAttestation,
		Message: fmt.Sprintf("Waiting for attestation from Circle (%s expected)...", estimatedTime),
	}

	msg, err := t.irisClient.PollForAttestation(
		ctx,
		t.params.SourceChain.Domain,
		burnTx.Hash().Hex(),
		func(attempt int, elapsed time.Duration) {
			updates <- TransferUpdate{
				Step:    StepPollingAttestation,
				Message: fmt.Sprintf("Polling for attestation (attempt %d, %s elapsed)...", attempt, elapsed.Round(time.Second)),
			}
		},
	)
	if err != nil {
		updates <- TransferUpdate{Step: StepError, Error: fmt.Errorf("failed to get attestation: %w", err)}
		return err
	}

	updates <- TransferUpdate{
		Step:    StepPollingAttestation,
		Message: "Attestation received!",
	}

	// Step 7: Mint on destination chain
	updates <- TransferUpdate{Step: StepMinting, Message: "Minting USDC on destination chain..."}

	// Decode hex strings (remove 0x prefix if present)
	messageBytes := common.FromHex(msg.Message)
	attestationBytes := common.FromHex(msg.Attestation)

	log.Debug("=== MINT TRANSACTION ===")
	log.Debug("Destination chain info", "chain", t.params.DestChain.Name, "domain", t.params.DestChain.Domain, "message_transmitter", t.params.DestChain.MessageTransmitterV2)
	log.Debug("Message sizes", "message_len", len(messageBytes), "attestation_len", len(attestationBytes))
	log.Debug("Message hex", "message", fmt.Sprintf("%x", messageBytes), "attestation", fmt.Sprintf("%x", attestationBytes))
	log.Debug("Message raw", "message", msg.Message, "attestation", msg.Attestation)
	log.Debug("=======================")

	// Create V2 message transmitter instance
	messageTransmitterV2 := messagetransmitter.NewMessageTransmitterV2()
	messageTransmitterAddr := common.HexToAddress(t.params.DestChain.MessageTransmitterV2)
	messageTransmitterInstance := messageTransmitterV2.Instance(t.destClient, messageTransmitterAddr)

	// Pack the receiveMessage call
	mintData := messageTransmitterV2.PackReceiveMessage(messageBytes, attestationBytes)

	// Create transaction options for mint (on destination chain)
	authDest, err := t.wallet.CreateTransactOpts(ctx, t.destClient, t.params.DestChain.ChainID)
	if err != nil {
		updates <- TransferUpdate{Step: StepError, Error: fmt.Errorf("failed to create auth: %w", err)}
		return err
	}

	// Use V2 bindings to send receiveMessage transaction
	mintTx, err := bind.Transact(messageTransmitterInstance, authDest, mintData)
	if err != nil {
		updates <- TransferUpdate{Step: StepError, Error: fmt.Errorf("failed to send mint tx: %w", err)}
		return err
	}

	updates <- TransferUpdate{
		Step:    StepMintingWait,
		Message: "Waiting for mint confirmation...",
		TxHash:  mintTx.Hash().Hex(),
	}

	// Wait for mint confirmation using V2 helper
	receipt, err := bind.WaitMined(ctx, t.destClient, mintTx.Hash())
	if err != nil {
		updates <- TransferUpdate{Step: StepError, Error: fmt.Errorf("mint transaction failed: %w", err)}
		return err
	}

	log.Info("Mint transaction confirmed", "tx_hash", mintTx.Hash().Hex())
	log.Debug("Receipt details", "status", receipt.Status, "gas_used", receipt.GasUsed, "num_logs", len(receipt.Logs))
	log.Info("USDC minted to recipient", "recipient", t.params.RecipientAddress.Hex(), "chain", t.params.DestChain.Name)

	if len(receipt.Logs) == 0 {
		log.Warn("No events emitted in mint transaction!")
		updates <- TransferUpdate{
			Step:    StepMinting,
			Message: "WARNING: Mint tx succeeded but no events emitted. Message may have been already received.",
		}
	}

	// Step 8: Complete
	updates <- TransferUpdate{
		Step:    StepComplete,
		Message: fmt.Sprintf("Transfer complete! USDC minted to %s on %s (%s)", t.params.RecipientAddress.Hex(), t.params.DestChain.Name, transferTypeLabel),
		TxHash:  mintTx.Hash().Hex(),
	}

	return nil
}

// Close closes the RPC clients
func (t *TransferOrchestrator) Close() {
	if t.sourceClient != nil {
		t.sourceClient.Close()
	}
	if t.destClient != nil {
		t.destClient.Close()
	}
}
