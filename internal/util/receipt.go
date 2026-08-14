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
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
)

// ValidateTransactionReceipt verifies that a mined transaction succeeded.
func ValidateTransactionReceipt(receipt *types.Receipt) error {
	if receipt == nil {
		return fmt.Errorf("missing transaction receipt")
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return fmt.Errorf("transaction reverted with receipt status %d", receipt.Status)
	}
	return nil
}
