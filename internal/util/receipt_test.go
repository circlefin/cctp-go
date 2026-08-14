// Copyright (C) 2026 Circle Internet Group, Inc.
// This file is part of the cctp-go project.

// The cctp-go project is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// The cctp-go project is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.

// You should have received a copy of the GNU Lesser General Public License
// along with the cctp-go project. If not, see <http://www.gnu.org/licenses/>.

package util

import (
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
)

func TestValidateTransactionReceipt(t *testing.T) {
	tests := []struct {
		name    string
		receipt *types.Receipt
		wantErr string
	}{
		{
			name:    "successful transaction",
			receipt: &types.Receipt{Status: types.ReceiptStatusSuccessful},
		},
		{
			name:    "reverted transaction",
			receipt: &types.Receipt{Status: types.ReceiptStatusFailed},
			wantErr: "transaction reverted with receipt status 0",
		},
		{
			name:    "missing receipt",
			receipt: nil,
			wantErr: "missing transaction receipt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateTransactionReceipt(tt.receipt)
			if tt.wantErr == "" && err != nil {
				t.Fatalf("ValidateTransactionReceipt() unexpected error = %v", err)
			}
			if tt.wantErr != "" && (err == nil || !strings.Contains(err.Error(), tt.wantErr)) {
				t.Fatalf("ValidateTransactionReceipt() error = %v, want error containing %q", err, tt.wantErr)
			}
		})
	}
}
