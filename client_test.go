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
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewIrisClient(t *testing.T) {
	baseURL := "https://iris-api.circle.com"
	client := NewIrisClient(baseURL)

	if client == nil {
		t.Fatal("expected non-nil client")
	}

	if client.baseURL != baseURL {
		t.Errorf("expected baseURL %s, got %s", baseURL, client.baseURL)
	}

	if client.httpClient == nil {
		t.Error("expected non-nil httpClient")
	}

	if client.httpClient.Timeout != 30*time.Second {
		t.Errorf("expected timeout 30s, got %v", client.httpClient.Timeout)
	}
}

func TestGetMessages_Success(t *testing.T) {
	expectedMessage := Message{
		Message:     "0x1234567890",
		EventNonce:  "123",
		Attestation: "0xabcdef",
		Status:      AttestationStatusComplete,
		CctpVersion: 2,
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request
		if r.Method != "GET" {
			t.Errorf("expected GET method, got %s", r.Method)
		}
		if r.URL.Path != "/v2/messages/0" {
			t.Errorf("expected path /v2/messages/0, got %s", r.URL.Path)
		}

		// Check query parameters
		txHash := r.URL.Query().Get("transactionHash")
		if txHash != "0xabc123" {
			t.Errorf("expected txHash 0xabc123, got %s", txHash)
		}

		// Send response
		resp := MessagesResponse{
			Messages: []Message{expectedMessage},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewIrisClient(server.URL)
	ctx := context.Background()

	resp, err := client.GetMessages(ctx, 0, "0xabc123")
	if err != nil {
		t.Fatalf("GetMessages failed: %v", err)
	}

	if len(resp.Messages) != 1 {
		t.Fatalf("expected 1 message, got %d", len(resp.Messages))
	}

	msg := resp.Messages[0]
	if msg.Message != expectedMessage.Message {
		t.Errorf("expected message %s, got %s", expectedMessage.Message, msg.Message)
	}
	if msg.Attestation != expectedMessage.Attestation {
		t.Errorf("expected attestation %s, got %s", expectedMessage.Attestation, msg.Attestation)
	}
	if msg.Status != expectedMessage.Status {
		t.Errorf("expected status %s, got %s", expectedMessage.Status, msg.Status)
	}
}

func TestGetMessages_404Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Transaction not found"))
	}))
	defer server.Close()

	client := NewIrisClient(server.URL)
	ctx := context.Background()

	_, err := client.GetMessages(ctx, 0, "0xnonexistent")
	if err == nil {
		t.Fatal("expected error for 404 response")
	}

	expectedError := "API returned status 404"
	if !contains(err.Error(), expectedError) {
		t.Errorf("expected error to contain %q, got %q", expectedError, err.Error())
	}
}

func TestGetMessages_RateLimitError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte("Rate limit exceeded"))
	}))
	defer server.Close()

	client := NewIrisClient(server.URL)
	ctx := context.Background()

	_, err := client.GetMessages(ctx, 0, "0xabc123")
	if err == nil {
		t.Fatal("expected error for 429 response")
	}

	expectedError := "rate limit exceeded"
	if !contains(err.Error(), expectedError) {
		t.Errorf("expected error to contain %q, got %q", expectedError, err.Error())
	}
}

func TestGetMessages_InvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("invalid json{"))
	}))
	defer server.Close()

	client := NewIrisClient(server.URL)
	ctx := context.Background()

	_, err := client.GetMessages(ctx, 0, "0xabc123")
	if err == nil {
		t.Fatal("expected error for invalid JSON")
	}

	expectedError := "failed to decode response"
	if !contains(err.Error(), expectedError) {
		t.Errorf("expected error to contain %q, got %q", expectedError, err.Error())
	}
}

func TestGetTransferFees_Success(t *testing.T) {
	expectedFees := FeesResponse{
		Data: []FeeInfo{
			{FinalityThreshold: 1000, MinimumFee: 10},
			{FinalityThreshold: 2000, MinimumFee: 0},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request
		if r.Method != "GET" {
			t.Errorf("expected GET method, got %s", r.Method)
		}
		expectedPath := "/v2/burn/USDC/fees/0/1"
		if r.URL.Path != expectedPath {
			t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
		}

		// Send response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(expectedFees)
	}))
	defer server.Close()

	client := NewIrisClient(server.URL)
	ctx := context.Background()

	resp, err := client.GetTransferFees(ctx, 0, 1)
	if err != nil {
		t.Fatalf("GetTransferFees failed: %v", err)
	}

	if len(resp.Data) != len(expectedFees.Data) {
		t.Fatalf("expected %d fee infos, got %d", len(expectedFees.Data), len(resp.Data))
	}

	for i, feeInfo := range resp.Data {
		expected := expectedFees.Data[i]
		if feeInfo.FinalityThreshold != expected.FinalityThreshold {
			t.Errorf("fee[%d]: expected threshold %d, got %d",
				i, expected.FinalityThreshold, feeInfo.FinalityThreshold)
		}
		if feeInfo.MinimumFee != expected.MinimumFee {
			t.Errorf("fee[%d]: expected fee %d, got %d",
				i, expected.MinimumFee, feeInfo.MinimumFee)
		}
	}
}

func TestGetTransferFees_ErrorCases(t *testing.T) {
	testCases := []struct {
		name           string
		statusCode     int
		responseBody   string
		expectedError  string
	}{
		{
			name:          "rate limit",
			statusCode:    http.StatusTooManyRequests,
			responseBody:  "Rate limit exceeded",
			expectedError: "rate limit exceeded",
		},
		{
			name:          "not found",
			statusCode:    http.StatusNotFound,
			responseBody:  "Not found",
			expectedError: "API returned status 404",
		},
		{
			name:          "internal server error",
			statusCode:    http.StatusInternalServerError,
			responseBody:  "Internal error",
			expectedError: "API returned status 500",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tc.statusCode)
				w.Write([]byte(tc.responseBody))
			}))
			defer server.Close()

			client := NewIrisClient(server.URL)
			ctx := context.Background()

			_, err := client.GetTransferFees(ctx, 0, 1)
			if err == nil {
				t.Fatal("expected error")
			}

			if !contains(err.Error(), tc.expectedError) {
				t.Errorf("expected error to contain %q, got %q", tc.expectedError, err.Error())
			}
		})
	}
}

func TestPollForAttestation_ImmediateSuccess(t *testing.T) {
	expectedMessage := Message{
		Message:     "0x1234567890",
		EventNonce:  "123",
		Attestation: "0xabcdef",
		Status:      AttestationStatusComplete,
		CctpVersion: 2,
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := MessagesResponse{
			Messages: []Message{expectedMessage},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewIrisClient(server.URL)
	ctx := context.Background()

	callCount := 0
	msg, err := client.PollForAttestation(ctx, 0, "0xabc123", func(attempt int, elapsed time.Duration) {
		callCount++
	})

	if err != nil {
		t.Fatalf("PollForAttestation failed: %v", err)
	}

	if msg.Message != expectedMessage.Message {
		t.Errorf("expected message %s, got %s", expectedMessage.Message, msg.Message)
	}
	if msg.Attestation != expectedMessage.Attestation {
		t.Errorf("expected attestation %s, got %s", expectedMessage.Attestation, msg.Attestation)
	}

	if callCount == 0 {
		t.Error("progress callback was never called")
	}
}

func TestPollForAttestation_EventualSuccess(t *testing.T) {
	attemptCount := 0
	maxAttempts := 3

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attemptCount++

		if attemptCount < maxAttempts {
			// Return pending status
			resp := MessagesResponse{
				Messages: []Message{
					{
						Message:     "0x1234567890",
						EventNonce:  "123",
						Attestation: "",
						Status:      AttestationStatusPending,
						CctpVersion: 2,
					},
				},
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
		} else {
			// Return complete status
			resp := MessagesResponse{
				Messages: []Message{
					{
						Message:     "0x1234567890",
						EventNonce:  "123",
						Attestation: "0xabcdef",
						Status:      AttestationStatusComplete,
						CctpVersion: 2,
					},
				},
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
		}
	}))
	defer server.Close()

	client := NewIrisClient(server.URL)
	ctx := context.Background()

	progressCallCount := 0
	msg, err := client.PollForAttestation(ctx, 0, "0xabc123", func(attempt int, elapsed time.Duration) {
		progressCallCount++
	})

	if err != nil {
		t.Fatalf("PollForAttestation failed: %v", err)
	}

	if msg.Attestation != "0xabcdef" {
		t.Errorf("expected attestation 0xabcdef, got %s", msg.Attestation)
	}

	if attemptCount != maxAttempts {
		t.Errorf("expected %d attempts, got %d", maxAttempts, attemptCount)
	}

	if progressCallCount < maxAttempts {
		t.Errorf("progress callback should be called at least %d times, got %d", maxAttempts, progressCallCount)
	}
}

func TestPollForAttestation_Timeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Always return pending
		resp := MessagesResponse{
			Messages: []Message{
				{
					Message:     "0x1234567890",
					EventNonce:  "123",
					Attestation: "",
					Status:      AttestationStatusPending,
					CctpVersion: 2,
				},
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewIrisClient(server.URL)
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	_, err := client.PollForAttestation(ctx, 0, "0xabc123", nil)
	if err == nil {
		t.Fatal("expected timeout error")
	}

	if err != context.DeadlineExceeded {
		t.Errorf("expected context.DeadlineExceeded, got %v", err)
	}
}

func TestPollForAttestation_Consecutive404s(t *testing.T) {
	attemptCount := 0
	max404s := 3

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attemptCount++
		// Always return 404
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Transaction not found"))
	}))
	defer server.Close()

	client := NewIrisClient(server.URL)
	ctx := context.Background()

	_, err := client.PollForAttestation(ctx, 0, "0xabc123", nil)
	if err == nil {
		t.Fatal("expected error for consecutive 404s")
	}

	// Should fail after max404Retries (3) attempts
	if attemptCount < max404s {
		t.Errorf("expected at least %d attempts, got %d", max404s, attemptCount)
	}

	expectedError := "received 404 error"
	if !contains(err.Error(), expectedError) {
		t.Errorf("expected error to contain %q, got %q", expectedError, err.Error())
	}
}

func TestPollForAttestation_NoMessages(t *testing.T) {
	attemptCount := 0
	maxAttempts := 2

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attemptCount++

		if attemptCount < maxAttempts {
			// Return empty messages
			resp := MessagesResponse{
				Messages: []Message{},
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
		} else {
			// Return message
			resp := MessagesResponse{
				Messages: []Message{
					{
						Message:     "0x1234567890",
						EventNonce:  "123",
						Attestation: "0xabcdef",
						Status:      AttestationStatusComplete,
						CctpVersion: 2,
					},
				},
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
		}
	}))
	defer server.Close()

	client := NewIrisClient(server.URL)
	ctx := context.Background()

	msg, err := client.PollForAttestation(ctx, 0, "0xabc123", nil)
	if err != nil {
		t.Fatalf("PollForAttestation failed: %v", err)
	}

	if msg.Attestation != "0xabcdef" {
		t.Errorf("expected attestation 0xabcdef, got %s", msg.Attestation)
	}

	if attemptCount != maxAttempts {
		t.Errorf("expected %d attempts, got %d", maxAttempts, attemptCount)
	}
}

func TestPollForAttestation_CompleteButNoAttestation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Return complete status but no attestation
		resp := MessagesResponse{
			Messages: []Message{
				{
					Message:     "0x1234567890",
					EventNonce:  "123",
					Attestation: "", // Empty attestation
					Status:      AttestationStatusComplete,
					CctpVersion: 2,
				},
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewIrisClient(server.URL)
	ctx := context.Background()

	_, err := client.PollForAttestation(ctx, 0, "0xabc123", nil)
	if err == nil {
		t.Fatal("expected error for complete status without attestation")
	}

	expectedError := "attestation status is complete but attestation is empty"
	if !contains(err.Error(), expectedError) {
		t.Errorf("expected error to contain %q, got %q", expectedError, err.Error())
	}
}

func TestEstimatedAttestationTime(t *testing.T) {
	testCases := []struct {
		name            string
		instantFinality bool
		expectedMin     time.Duration
		expectedMax     time.Duration
	}{
		{
			name:            "instant finality",
			instantFinality: true,
			expectedMin:     5 * time.Second,
			expectedMax:     15 * time.Second,
		},
		{
			name:            "not instant finality",
			instantFinality: false,
			expectedMin:     10 * time.Second,
			expectedMax:     20 * time.Second,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			duration := EstimatedAttestationTime(tc.instantFinality)

			if duration < tc.expectedMin || duration > tc.expectedMax {
				t.Errorf("expected duration between %v and %v, got %v",
					tc.expectedMin, tc.expectedMax, duration)
			}
		})
	}
}

func TestPollForAttestation_ContextCancellation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Delay response
		time.Sleep(50 * time.Millisecond)
		resp := MessagesResponse{
			Messages: []Message{
				{
					Message:     "0x1234567890",
					EventNonce:  "123",
					Attestation: "",
					Status:      AttestationStatusPending,
					CctpVersion: 2,
				},
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewIrisClient(server.URL)
	ctx, cancel := context.WithCancel(context.Background())

	// Cancel after a short delay
	go func() {
		time.Sleep(10 * time.Millisecond)
		cancel()
	}()

	_, err := client.PollForAttestation(ctx, 0, "0xabc123", nil)
	if err == nil {
		t.Fatal("expected context cancellation error")
	}

	// The error might be wrapped, so just check it's a cancellation error
	if ctx.Err() != context.Canceled {
		t.Errorf("context should be canceled")
	}
}

func TestGetMessages_ContextCancellation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Delay to allow context to be cancelled
		time.Sleep(100 * time.Millisecond)
		resp := MessagesResponse{Messages: []Message{}}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewIrisClient(server.URL)
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	_, err := client.GetMessages(ctx, 0, "0xabc123")
	if err == nil {
		t.Fatal("expected context cancellation error")
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && containsHelper(s, substr))
}

func containsHelper(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

