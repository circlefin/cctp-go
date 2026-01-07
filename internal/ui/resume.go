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
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pxgray/cctp-go"
	"github.com/pxgray/cctp-go/internal/logger"
	"github.com/pxgray/cctp-go/messagetransmitter"
)

// updateResumeInput handles updates for resume input state
func (m Model) updateResumeInput(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			// Go back to main menu
			m.goBack()
			return m, nil

		case "enter":
			// Validate and start resume
			txHash := strings.TrimSpace(m.txHashInput.Value())

			// Validate tx hash
			if !strings.HasPrefix(txHash, "0x") || len(txHash) != 66 {
				m.err = fmt.Errorf("invalid transaction hash format (must be 0x... with 66 characters)")
				m.changeState(StateError)
				return m, nil
			}

			// Start resume - source and destination will be auto-detected
			m.changeState(StateProgress)
			m.currentStep = cctp.StepPollingAttestation
			return m, m.resumeTransfer(txHash)
		}
	}

	// Update input
	m.txHashInput, cmd = m.txHashInput.Update(msg)

	return m, cmd
}

// viewResumeInput renders the resume input screen
func (m Model) viewResumeInput() string {
	breadcrumb := RenderBreadcrumb(m.state, m.resumeMode)
	title := RenderTitle("Resume Transfer")

	walletInfo := RenderInfo(fmt.Sprintf("Wallet: %s", m.wallet.Address.Hex()))

	network := RenderNetworkIndicator(m.testnet)
	info := RenderWarning("Enter the burn transaction hash from the source chain\nSource and destination chains will be auto-detected.")

	txHashLabel := "Burn Transaction Hash:"
	txHashStyle := FocusedInputStyle
	txHashView := txHashStyle.Render(m.txHashInput.View())

	help := RenderHelp("enter: resume transfer • esc: back • q: quit")

	return lipgloss.JoinVertical(
		lipgloss.Left,
		network,
		breadcrumb,
		"",
		title,
		walletInfo,
		"",
		info,
		"",
		txHashLabel,
		txHashView,
		"",
		help,
	)
}

// resumeTransfer resumes a transfer from a burn transaction hash
func (m Model) resumeTransfer(txHash string) tea.Cmd {
	return func() tea.Msg {
		// Determine Iris API base URL
		irisURL := "https://iris-api.circle.com"
		if m.testnet {
			irisURL = "https://iris-api-sandbox.circle.com"
		}

		// Create Iris client
		irisClient := cctp.NewIrisClient(irisURL)

		updates := make(chan cctp.TransferUpdate, 10)
		ctx := context.Background()

		go func() {
			defer close(updates)

			// Try to detect source and destination domains by querying all possible chains
			var sourceDomain uint32
			var destDomain uint32
			var detectedSourceChain *cctp.Chain
			var detectedDestChain *cctp.Chain
			var messageResp *cctp.Message

			chains := cctp.GetChains(m.testnet)
			chains = cctp.ApplyRPCOverrides(chains, m.config.RPCUrls)

			updates <- cctp.TransferUpdate{
				Step:    cctp.StepPollingAttestation,
				Message: "Auto-detecting source and destination chains...",
			}

			for _, chain := range chains {
				// Try to get messages from this domain
				log.Debug("Checking for burn transaction", "tx_hash", txHash, "chain", chain.Name, "domain", chain.Domain)
				resp, err := irisClient.GetMessages(ctx, chain.Domain, txHash)
				if err != nil {
					log.Debug("API error", "chain", chain.Name, "error", err)
					continue
				}

				log.Debug("Got messages", "count", len(resp.Messages), "chain", chain.Name)

				if len(resp.Messages) > 0 {
					msg := &resp.Messages[0]

					// Check if decodedMessage is available
					if msg.DecodedMessage == nil {
						log.Debug("Message found but DecodedMessage is nil, skipping")
						continue
					}

					log.Debug("DecodedMessage present", "source_domain", msg.DecodedMessage.SourceDomain, "destination_domain", msg.DecodedMessage.DestinationDomain)

					// Parse destination domain from decodedMessage
					destDomainParsed, parseErr := strconv.ParseUint(msg.DecodedMessage.DestinationDomain, 10, 32)
					if parseErr != nil {
						log.Debug("Failed to parse destination domain", "destination_domain", msg.DecodedMessage.DestinationDomain, "error", parseErr)
						continue
					}

					sourceDomain = chain.Domain
					destDomain = uint32(destDomainParsed)
					detectedSourceChain = &chain
					messageResp = msg
					log.Info("Found burn transaction on source chain", "chain", chain.Name, "source_domain", sourceDomain, "dest_domain", destDomain)
					break
				}
			}

			if detectedSourceChain == nil {
				updates <- cctp.TransferUpdate{
					Step:  cctp.StepError,
					Error: fmt.Errorf("could not find burn transaction on any chain. Please verify the tx hash and network (mainnet/testnet)"),
				}
				return
			}

			// Log the detected source and destination domains for debugging
			updates <- cctp.TransferUpdate{
				Step:    cctp.StepPollingAttestation,
				Message: fmt.Sprintf("Found burn on %s (domain %d) → destination domain %d", detectedSourceChain.Name, sourceDomain, destDomain),
			}

			// Get destination chain by domain
			detectedDestChain, err := cctp.GetChainByDomain(destDomain, m.testnet)
			if err != nil {
				// Provide helpful error message with available domains
				networkType := "mainnet"
				if m.testnet {
					networkType = "testnet"
				}

				// List available domains for the current network
				availableDomains := ""
				for i, chain := range chains {
					if i > 0 {
						availableDomains += ", "
					}
					availableDomains += fmt.Sprintf("%d (%s)", chain.Domain, chain.Name)
				}

				log.Error("Destination domain not found in configuration", "dest_domain", destDomain, "network", networkType)
				log.Error("Available domains", "network", networkType, "domains", availableDomains)
				
				// Write log to file on error
				_ = logger.WriteToFile()

				updates <- cctp.TransferUpdate{
					Step: cctp.StepError,
					Error: fmt.Errorf(
						"destination domain %d not found in %s configuration\n\n"+
							"This usually means:\n"+
							"1. You're in %s mode but the burn was for a %s chain\n"+
							"2. The destination domain is not supported\n\n"+
							"Available %s domains: %s\n\n"+
							"Try running the CLI with %s flag",
						destDomain,
						networkType,
						networkType,
						map[bool]string{true: "mainnet", false: "testnet"}[m.testnet],
						networkType,
						availableDomains,
						map[bool]string{true: "", false: "--testnet"}[m.testnet],
					),
				}
				return
			}

			m.sourceChain = detectedSourceChain
			m.destChain = detectedDestChain

			// Send the burn transaction hash first so it gets stored
			updates <- cctp.TransferUpdate{
				Step:        cctp.StepBurning,
				Message:     fmt.Sprintf("Burn transaction detected on %s", detectedSourceChain.Name),
				TxHash:      txHash,
				SourceChain: detectedSourceChain,
				DestChain:   detectedDestChain,
			}

			updates <- cctp.TransferUpdate{
				Step:        cctp.StepPollingAttestation,
				Message:     fmt.Sprintf("Detected: %s → %s", detectedSourceChain.Name, detectedDestChain.Name),
				SourceChain: detectedSourceChain,
				DestChain:   detectedDestChain,
			}

			// Poll for attestation if needed
			var msg *cctp.Message

			if messageResp != nil && messageResp.Status == cctp.AttestationStatusComplete {
				msg = messageResp
				updates <- cctp.TransferUpdate{
					Step:    cctp.StepPollingAttestation,
					Message: "Attestation already available!",
				}
			} else {
				// Poll for attestation
				updates <- cctp.TransferUpdate{
					Step:    cctp.StepPollingAttestation,
					Message: fmt.Sprintf("Polling for attestation (tx: %s)...", txHash[:10]+"..."),
				}

				var pollErr error
				msg, pollErr = irisClient.PollForAttestation(
					ctx,
					sourceDomain,
					txHash,
					func(attempt int, elapsed time.Duration) {
						updates <- cctp.TransferUpdate{
							Step:    cctp.StepPollingAttestation,
							Message: fmt.Sprintf("Polling for attestation (attempt %d, %s elapsed)...", attempt, elapsed.Round(time.Second)),
						}
					},
				)
				if pollErr != nil {
					updates <- cctp.TransferUpdate{
						Step:  cctp.StepError,
						Error: fmt.Errorf("failed to get attestation: %w", pollErr),
					}
					return
				}
			}

			updates <- cctp.TransferUpdate{
				Step:    cctp.StepPollingAttestation,
				Message: "Attestation received!",
			}

			// Now mint on destination chain
			updates <- cctp.TransferUpdate{
				Step:    cctp.StepMinting,
				Message: "Minting USDC on destination chain...",
			}

			// Connect to destination chain
			destClient, err := ethclient.Dial(m.destChain.RPC)
			if err != nil {
				updates <- cctp.TransferUpdate{
					Step:  cctp.StepError,
					Error: fmt.Errorf("failed to connect to destination chain: %w", err),
				}
				return
			}
			defer destClient.Close()

			// Decode attestation and message
			messageBytes := common.FromHex(msg.Message)
			attestationBytes := common.FromHex(msg.Attestation)

			updates <- cctp.TransferUpdate{
				Step:    cctp.StepMinting,
				Message: fmt.Sprintf("Preparing mint (msg: %d bytes, att: %d bytes)", len(messageBytes), len(attestationBytes)),
			}

		// Use V2 bindings to mint
		messageTransmitterV2 := messagetransmitter.NewMessageTransmitterV2()
		messageTransmitterAddr := common.HexToAddress(m.destChain.MessageTransmitterV2)
		messageTransmitterInstance := messageTransmitterV2.Instance(destClient, messageTransmitterAddr)
		
		// Pack the receiveMessage call
		mintData := messageTransmitterV2.PackReceiveMessage(messageBytes, attestationBytes)

		// Create transaction options for mint
		authDest, err := m.wallet.CreateTransactOpts(ctx, destClient, m.destChain.ChainID)
		if err != nil {
			updates <- cctp.TransferUpdate{
				Step:  cctp.StepError,
				Error: fmt.Errorf("failed to create auth: %w", err),
			}
			return
		}

		// Use V2 bindings to send receiveMessage transaction
		mintTx, err := bind.Transact(messageTransmitterInstance, authDest, mintData)
		if err != nil {
			updates <- cctp.TransferUpdate{
				Step:  cctp.StepError,
				Error: fmt.Errorf("failed to send mint tx: %w", err),
			}
			return
		}

		updates <- cctp.TransferUpdate{
			Step:    cctp.StepMintingWait,
			Message: "Waiting for mint confirmation...",
			TxHash:  mintTx.Hash().Hex(),
		}

		// Wait for mint confirmation
		_, err = bind.WaitMined(ctx, destClient, mintTx.Hash())
			if err != nil {
				updates <- cctp.TransferUpdate{
					Step:  cctp.StepError,
					Error: fmt.Errorf("mint transaction failed: %w", err),
				}
				return
			}

			// Complete
			updates <- cctp.TransferUpdate{
				Step:    cctp.StepComplete,
				Message: fmt.Sprintf("Transfer resumed and completed! USDC minted on %s", m.destChain.Name),
				TxHash:  mintTx.Hash().Hex(),
			}
		}()

		return InitTransferMsg{UpdatesChan: updates}
	}
}
