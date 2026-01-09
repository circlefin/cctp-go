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
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/charmbracelet/log"
)

// AttestationStatus represents the status of an attestation
type AttestationStatus string

const (
	AttestationStatusPending  AttestationStatus = "pending"
	AttestationStatusComplete AttestationStatus = "complete"
)

// FeeInfo represents fee information for a specific finality threshold
type FeeInfo struct {
	FinalityThreshold uint32 `json:"finalityThreshold"`
	MinimumFee        uint32 `json:"minimumFee"` // in basis points (bps)
}

// FeesResponse represents the response from /v2/burn/USDC/fees
type FeesResponse struct {
	Data []FeeInfo `json:"data"`
}

// DecodedMessage represents the decoded CCTP message fields
type DecodedMessage struct {
	SourceDomain              string `json:"sourceDomain"`
	DestinationDomain         string `json:"destinationDomain"`
	Nonce                     string `json:"nonce"`
	Sender                    string `json:"sender"`
	Recipient                 string `json:"recipient"`
	DestinationCaller         string `json:"destinationCaller"`
	MinFinalityThreshold      string `json:"minFinalityThreshold"`
	FinalityThresholdExecuted string `json:"finalityThresholdExecuted"`
	MessageBody               string `json:"messageBody"`
}

// Message represents a CCTP message
type Message struct {
	Message        string            `json:"message"`
	EventNonce     string            `json:"eventNonce"`
	Attestation    string            `json:"attestation"`
	DecodedMessage *DecodedMessage   `json:"decodedMessage"`
	CctpVersion    int               `json:"cctpVersion"`
	Status         AttestationStatus `json:"status"`
	DelayReason    string            `json:"delayReason"`
}

// MessagesResponse represents the response from /v2/messages
type MessagesResponse struct {
	Messages []Message `json:"messages"`
}

// IrisClient is a client for Circle's Iris attestation service
type IrisClient struct {
	baseURL    string
	httpClient *http.Client
}

// NewIrisClient creates a new Iris API client
func NewIrisClient(baseURL string) *IrisClient {
	return &IrisClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// GetMessages fetches messages and attestations for a transaction
func (c *IrisClient) GetMessages(ctx context.Context, sourceDomain uint32, txHash string) (*MessagesResponse, error) {
	url := fmt.Sprintf("%s/v2/messages/%d?transactionHash=%s", c.baseURL, sourceDomain, txHash)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusTooManyRequests {
		return nil, fmt.Errorf("rate limit exceeded (429): please wait 5 minutes before retrying")
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(body))
	}

	var result MessagesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// checkAttestationStatus checks if an attestation is ready for a given transaction
// Returns the message if attestation is ready, or an error indicating the status
func (c *IrisClient) checkAttestationStatus(ctx context.Context, sourceDomain uint32, txHash string, attempt int) (*Message, error) {
	resp, err := c.GetMessages(ctx, sourceDomain, txHash)
	if err != nil {
		return nil, err
	}

	// Check if we have messages
	if len(resp.Messages) == 0 {
		return nil, fmt.Errorf("no messages found")
	}

	// Get the first message
	msg := resp.Messages[0]

	// Log status for debugging
	if attempt == 1 || attempt%5 == 0 {
		log.Debug("Attestation status", "status", msg.Status, "has_attestation", msg.Attestation != "")
	}

	// Check if attestation is complete
	if msg.Status == AttestationStatusComplete && msg.Attestation != "" {
		return &msg, nil
	}

	// If status is complete but no attestation, there's an issue
	if msg.Status == AttestationStatusComplete && msg.Attestation == "" {
		return nil, fmt.Errorf("attestation status is complete but attestation is empty")
	}

	// Attestation not ready yet
	return nil, fmt.Errorf("attestation not ready, status: %s", msg.Status)
}

// PollForAttestation polls for an attestation until it's ready or context is cancelled
func (c *IrisClient) PollForAttestation(ctx context.Context, sourceDomain uint32, txHash string, progressCallback func(attempt int, elapsed time.Duration)) (*Message, error) {
	startTime := time.Now()
	attempt := 0
	backoff := 2 * time.Second
	maxBackoff := 10 * time.Second
	consecutive404Count := 0
	max404Retries := 3

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		attempt++
		elapsed := time.Since(startTime)

		if progressCallback != nil {
			progressCallback(attempt, elapsed)
		}

		msg, err := c.checkAttestationStatus(ctx, sourceDomain, txHash, attempt)
		if err == nil {
			// Attestation is ready
			return msg, nil
		}

		// Check if this is a 404 error
		errMsg := err.Error()
		is404 := false

		// Parse "API returned status 404: ..." format
		var statusCode int
		if n, scanErr := fmt.Sscanf(errMsg, "API returned status %d:", &statusCode); scanErr == nil && n == 1 && statusCode == 404 {
			is404 = true
		}

		if is404 {
			consecutive404Count++
			log.Debug("Received 404 error - API may not have picked up transaction yet", "attempt", consecutive404Count, "max_retries", max404Retries)

			if consecutive404Count >= max404Retries {
				return nil, fmt.Errorf("received 404 error %d times: %w", max404Retries, err)
			}
		} else {
			// Reset 404 counter on successful response (even if attestation not ready)
			consecutive404Count = 0
		}

		// Check if it's a fatal error (not 404, not "no messages", not "attestation not ready")
		if !is404 && errMsg != "no messages found" && !strings.Contains(errMsg, "attestation not ready") {
			return nil, err
		}

		// Wait before retrying
		select {
		case <-time.After(backoff):
		case <-ctx.Done():
			return nil, ctx.Err()
		}

		// Exponential backoff with cap
		backoff = backoff * 2
		if backoff > maxBackoff {
			backoff = maxBackoff
		}
	}
}

// GetTransferFees fetches the fee information for a transfer between two domains
func (c *IrisClient) GetTransferFees(ctx context.Context, sourceDomain, destDomain uint32) (*FeesResponse, error) {
	url := fmt.Sprintf("%s/v2/burn/USDC/fees/%d/%d", c.baseURL, sourceDomain, destDomain)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusTooManyRequests {
		return nil, fmt.Errorf("rate limit exceeded (429): please wait 5 minutes before retrying")
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(body))
	}

	var result FeesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// EstimatedAttestationTime returns the estimated time for attestation based on the chain
func EstimatedAttestationTime(instantFinality bool) time.Duration {
	if instantFinality {
		return 10 * time.Second // ~8-10 seconds for instant finality chains
	}
	return 15 * time.Second // ~8-20 seconds for Fast Transfer on most EVM chains
}
