package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/pxgray/cctp-go"
)

// updateNetworkSelection handles updates for network selection state
func (m Model) updateNetworkSelection(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			// Go back to account/password selection (if using keystore)
			if m.config != nil && m.config.KeystorePath != "" {
				m.changeState(StateSelectAccount)
				return m, nil
			}
			// If using private key, can't go back - network selection is first screen
			return m, nil
		
		case "enter":
			// Get selected network
			selectedItem, ok := m.networkList.SelectedItem().(networkItem)
			if !ok {
				return m, nil
			}
			
			// Apply network selection
			m.testnet = selectedItem.isTestnet
			m.networkSetByUser = true
			
			// Reload chains for selected network with RPC overrides from config
			m.chains = cctp.GetChains(m.testnet)
			m.chains = cctp.ApplyRPCOverrides(m.chains, m.config.RPCUrls)
			
			// Update chain list items
			m.updateChainListItems()
			
			m.changeState(StateMainMenu)
			return m, nil
		}
	}
	
	// Update network list
	var cmd tea.Cmd
	m.networkList, cmd = m.networkList.Update(msg)
	return m, cmd
}

// updateChainListItems refreshes the chain list items based on current chains
func (m *Model) updateChainListItems() {
	items := make([]list.Item, len(m.chains))
	for i, chain := range m.chains {
		items[i] = chainItem{chain: chain}
	}
	m.chainList.SetItems(items)
	m.chainList.ResetFilter() // Clear any active filter from previous network
	m.chainList.Title = "Select Source Chain"
}

// viewNetworkSelection renders the network selection screen
func (m Model) viewNetworkSelection() string {
	title := RenderTitle("CCTP V2 CLI - Select Network")
	
	elements := []string{title}
	
	if m.wallet != nil {
		elements = append(elements, RenderInfo(fmt.Sprintf("Wallet: %s", m.wallet.Address.Hex())), "")
	}
	
	elements = append(elements, m.networkList.View())
	
	// Show escape option if using keystore (can go back to account selection)
	helpText := "↑/↓: navigate • enter: select • q: quit"
	if m.config != nil && m.config.KeystorePath != "" {
		helpText = "↑/↓: navigate • enter: select • esc: back • q: quit"
	}
	elements = append(elements, "", RenderHelp(helpText))

	return lipgloss.JoinVertical(lipgloss.Left, elements...)
}

