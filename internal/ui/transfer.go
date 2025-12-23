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

package ui

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pxgray/cctp-go"
	"github.com/pxgray/cctp-go/internal/logger"
	"github.com/pxgray/cctp-go/internal/util"
)

// updateTransferInput handles updates for transfer input state
func (m Model) updateTransferInput(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case BalanceUpdateMsg:
		// Handle balance update
		m.loadingBalance = false
		if msg.Error != nil {
			// Balance fetch failed, but don't block the user
			m.usdcBalance = nil
		} else {
			m.usdcBalance = msg.Balance
		}
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			// Go back to destination chain selection
			m.goBack()
			return m, nil

		case "tab", "shift+tab":
			// Switch focus between inputs (amount, recipient, transfer type)
			m.inputFocus = (m.inputFocus + 1) % 3
			m.amountInput.Blur()
			m.recipientInput.Blur()
			switch m.inputFocus {
			case 0:
				m.amountInput.Focus()
			case 1:
				m.recipientInput.Focus()
				// case 2: transfer type selection (no text input to focus)
			}
			return m, nil

		case "left", "right":
			// Handle transfer type selection when on that field
			if m.inputFocus == 2 {
				m.cycleTransferType(msg.String() == "right")
				return m, nil
			}

		case "enter":
			// Validate and start transfer
			amount := m.amountInput.Value()
			recipient := strings.TrimSpace(m.recipientInput.Value())

			// Validate amount
			if amount == "" {
				m.err = fmt.Errorf("amount is required")
				m.changeState(StateError)
				return m, nil
			}

			// Parse amount (assuming 6 decimals for USDC)
			amountFloat := new(big.Float)
			_, ok := amountFloat.SetString(amount)
			if !ok {
				m.err = fmt.Errorf("invalid amount format")
				m.changeState(StateError)
				return m, nil
			}

			// Convert to wei (multiply by 10^6 for USDC)
			multiplier := new(big.Float).SetInt(big.NewInt(1000000))
			amountWei := new(big.Float).Mul(amountFloat, multiplier)
			amountBigInt, _ := amountWei.Int(nil)

			// Validate recipient address
			if !common.IsHexAddress(recipient) {
				m.err = fmt.Errorf("invalid recipient address format")
				m.changeState(StateError)
				return m, nil
			}

			// Start transfer
			m.changeState(StateProgress)
			m.currentStep = cctp.StepCheckBalance
			return m, m.startTransfer(amountBigInt, common.HexToAddress(recipient))
		}
	}

	// Update inputs
	if m.inputFocus == 0 {
		m.amountInput, cmd = m.amountInput.Update(msg)
	} else {
		m.recipientInput, cmd = m.recipientInput.Update(msg)
	}

	return m, cmd
}

// viewTransferInput renders the transfer input screen
func (m Model) viewTransferInput() string {
	network := RenderNetworkIndicator(m.testnet)
	breadcrumb := RenderBreadcrumb(m.state, m.resumeMode)
	title := RenderTitle("Transfer Details")

	chainInfo := lipgloss.JoinHorizontal(
		lipgloss.Center,
		RenderChainBadge(m.sourceChain.Name, m.testnet),
		" → ",
		RenderChainBadge(m.destChain.Name, m.testnet),
	)

	walletInfo := RenderInfo(fmt.Sprintf("Wallet: %s", m.wallet.Address.Hex()))

	// Format and display balance
	var balanceInfo string
	if m.loadingBalance {
		balanceInfo = RenderInfo(fmt.Sprintf("%s Fetching USDC balance...", m.spinner.View()))
	} else if m.usdcBalance != nil {
		balanceInfo = RenderInfo(fmt.Sprintf("Balance: %s USDC", util.FormatUSDCBalance(m.usdcBalance)))
	} else {
		balanceInfo = RenderInfo("Balance: Unable to fetch")
	}

	amountLabel := "Amount (USDC):"
	amountStyle := BlurredInputStyle
	if m.inputFocus == 0 {
		amountStyle = FocusedInputStyle
	}
	amountView := amountStyle.Render(m.amountInput.View())

	recipientLabel := "Recipient Address:"
	recipientStyle := BlurredInputStyle
	if m.inputFocus == 1 {
		recipientStyle = FocusedInputStyle
	}
	recipientView := recipientStyle.Render(m.recipientInput.View())

	// Transfer type selector
	transferTypeLabel := "Transfer Type:"
	transferType := m.transferType
	if transferType == "" {
		transferType = cctp.TransferTypeAuto
	}
	transferTypeText := getTransferTypeLabel(transferType, m.sourceChain)
	transferTypeStyle := BlurredInputStyle
	if m.inputFocus == 2 {
		transferTypeStyle = FocusedInputStyle
	}
	transferTypeView := transferTypeStyle.Render(fmt.Sprintf("< %s >", transferTypeText))

	help := RenderHelp("tab: switch field • ←→: change type • enter: confirm • esc: back • q: quit")

	return lipgloss.JoinVertical(
		lipgloss.Left,
		network,
		breadcrumb,
		"",
		title,
		"",
		chainInfo,
		walletInfo,
		balanceInfo,
		"",
		amountLabel,
		amountView,
		"",
		recipientLabel,
		recipientView,
		"",
		transferTypeLabel,
		transferTypeView,
		"",
		help,
	)
}

// TransferUpdateMsg is a message containing transfer updates
type TransferUpdateMsg struct {
	Update cctp.TransferUpdate
}

// InitTransferMsg indicates transfer initialization is complete
type InitTransferMsg struct {
	UpdatesChan chan cctp.TransferUpdate
}

// LogUpdateMsg is a message containing a new log entry
type LogUpdateMsg struct {
	Entry logger.LogEntry
}

// BalanceUpdateMsg is a message containing the USDC balance
type BalanceUpdateMsg struct {
	Balance *big.Int
	Error   error
}

// cycleTransferType cycles through transfer type options
func (m *Model) cycleTransferType(forward bool) {
	types := m.getAvailableTransferTypes()
	if len(types) == 0 {
		return
	}

	// Find current type index
	currentIdx := -1
	for i, tt := range types {
		if tt == m.transferType {
			currentIdx = i
			break
		}
	}

	// If not set, default to first option
	if currentIdx == -1 {
		m.transferType = types[0]
		return
	}

	// Cycle to next/previous
	if forward {
		currentIdx = (currentIdx + 1) % len(types)
	} else {
		currentIdx = (currentIdx - 1 + len(types)) % len(types)
	}

	m.transferType = types[currentIdx]
}

// getAvailableTransferTypes returns available transfer types based on source chain
func (m *Model) getAvailableTransferTypes() []cctp.TransferType {
	types := []cctp.TransferType{cctp.TransferTypeAuto}

	if m.sourceChain != nil {
		// Fast Transfer not available for instant finality chains as source
		if !m.sourceChain.InstantFinality {
			types = append(types, cctp.TransferTypeFast)
		}
		// Standard Transfer always available
		types = append(types, cctp.TransferTypeStandard)
	}

	return types
}

// getTransferTypeLabel returns a user-friendly label for a transfer type
func getTransferTypeLabel(tt cctp.TransferType, sourceChain *cctp.Chain) string {
	switch tt {
	case cctp.TransferTypeFast:
		return "Fast Transfer (~8-20 seconds, with fee)"
	case cctp.TransferTypeStandard:
		if sourceChain != nil && sourceChain.InstantFinality {
			return "Standard Transfer (~8 seconds, no fee)"
		}
		return "Standard Transfer (~13-19 minutes, no fee)"
	case cctp.TransferTypeAuto:
		return "Auto (recommended)"
	default:
		return "Auto (recommended)"
	}
}

// startTransfer starts the transfer process
func (m Model) startTransfer(amount *big.Int, recipient common.Address) tea.Cmd {
	return func() tea.Msg {
		// Default to auto if not set
		transferType := m.transferType
		if transferType == "" {
			transferType = cctp.TransferTypeAuto
		}

		// Create transfer params with cached balance to avoid redundant RPC call
		params := &cctp.TransferParams{
			SourceChain:      m.sourceChain,
			DestChain:        m.destChain,
			Amount:           amount,
			RecipientAddress: recipient,
			Testnet:          m.testnet,
			TransferType:     transferType,
			CachedBalance:    m.usdcBalance, // Pass cached balance from UI
		}

		// Determine Iris API base URL
		irisURL := "https://iris-api.circle.com"
		if m.testnet {
			irisURL = "https://iris-api-sandbox.circle.com"
		}

		// Create orchestrator
		orchestrator, err := cctp.NewTransferOrchestrator(m.wallet, params, irisURL)
		if err != nil {
			return TransferUpdateMsg{
				Update: cctp.TransferUpdate{
					Step:  cctp.StepError,
					Error: err,
				},
			}
		}

		// Execute transfer in background
		updates := make(chan cctp.TransferUpdate, 10) // Buffered channel
		ctx := context.Background()

		go func() {
			defer orchestrator.Close()
			_ = orchestrator.Execute(ctx, updates)
		}()

		// Return the channel so the UI can start listening
		return InitTransferMsg{UpdatesChan: updates}
	}
}

// waitForNextUpdate waits for the next transfer update
func (m Model) waitForNextUpdate() tea.Cmd {
	return func() tea.Msg {
		if m.updatesChan == nil {
			return nil
		}
		update, ok := <-m.updatesChan
		if !ok {
			// Channel closed, no more updates
			return nil
		}
		return TransferUpdateMsg{Update: update}
	}
}

// waitForLogUpdate waits for the next log entry
func waitForLogUpdate() tea.Cmd {
	return func() tea.Msg {
		// Get logger and wait for next entry
		log := logger.GetLogger()
		entry, ok := <-log.UpdateChan()
		if !ok {
			return nil
		}
		return LogUpdateMsg{Entry: entry}
	}
}

// fetchUSDCBalance fetches the USDC balance for the wallet on the source chain
func (m Model) fetchUSDCBalance() tea.Cmd {
	return func() tea.Msg {
		if m.sourceChain == nil || m.wallet == nil {
			return BalanceUpdateMsg{Error: fmt.Errorf("source chain or wallet not configured")}
		}

		ctx := context.Background()
		
		// Connect to source chain
		client, err := ethclient.Dial(m.sourceChain.RPC)
		if err != nil {
			return BalanceUpdateMsg{Error: fmt.Errorf("failed to connect to source chain: %w", err)}
		}
		defer client.Close()

		// Use shared balance checking utility
		usdcAddress := common.HexToAddress(m.sourceChain.USDC)
		balance, err := util.GetUSDCBalance(ctx, client, usdcAddress, m.wallet.Address)
		if err != nil {
			return BalanceUpdateMsg{Error: err}
		}

		return BalanceUpdateMsg{Balance: balance}
	}
}
