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

package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
	tea "github.com/charmbracelet/bubbletea"
)

// updateChainSelection handles updates for chain selection state
func (m Model) updateChainSelection(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			// Go back to previous state
			m.goBack()
			return m, nil

		case "enter":
			// Get selected chain
			selectedItem, ok := m.chainList.SelectedItem().(chainItem)
			if !ok {
				return m, nil
			}

		if m.state == StateSelectSourceChain {
			// Source chain selected, move to destination chain selection
			m.sourceChain = &selectedItem.chain
			
			// Update list for destination chain (exclude source chain)
			items := make([]list.Item, 0)
			for _, chain := range m.chains {
				if chain.Domain != m.sourceChain.Domain {
					items = append(items, chainItem{chain: chain})
				}
			}
			
			m.chainList.SetItems(items)
			m.chainList.ResetFilter() // Clear any active filter from source chain selection
			m.chainList.Title = "Select Destination Chain"
			m.changeState(StateSelectDestChain)
			return m, nil
		} else {
			// Destination chain selected, move to transfer input
			m.destChain = &selectedItem.chain
			m.changeState(StateTransferInput)
			m.amountInput.Focus()
			m.inputFocus = 0
			m.loadingBalance = true
			m.usdcBalance = nil
			return m, tea.Batch(
				m.amountInput.Focus(),
				m.fetchUSDCBalance(),
			)
		}
		}
	}

	// Update list
	var cmd tea.Cmd
	m.chainList, cmd = m.chainList.Update(msg)
	return m, cmd
}

// viewChainSelection renders the chain selection screen
func (m Model) viewChainSelection() string {
	network := RenderNetworkIndicator(m.testnet)
	breadcrumb := RenderBreadcrumb(m.state, m.resumeMode)
	
	elements := []string{network, breadcrumb, ""}
	
	if m.state == StateSelectSourceChain {
		elements = append(elements, RenderTitle("CCTP V2 - Select Source Chain"))
	} else {
		elements = append(elements, RenderTitle("CCTP V2 - Select Destination Chain"))
		if m.sourceChain != nil {
			elements = append(elements, lipgloss.JoinHorizontal(
				lipgloss.Left,
				RenderInfo("From: "),
				RenderChainBadge(m.sourceChain.Name, m.testnet),
			))
		}
	}

	help := RenderHelp("↑/↓: navigate • enter: select • esc: back • q: quit")

	elements = append(elements, "", m.chainList.View(), "", help)
	return lipgloss.JoinVertical(lipgloss.Left, elements...)
}

