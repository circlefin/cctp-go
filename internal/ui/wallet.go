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
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/circlefin/cctp-go/internal/wallet"
	"github.com/ethereum/go-ethereum/common"
)

// accountItem implements list.Item for account selection
type accountItem struct {
	address common.Address
	index   int
}

func (i accountItem) Title() string {
	return i.address.Hex()
}

func (i accountItem) Description() string {
	return fmt.Sprintf("Account #%d", i.index+1)
}

func (i accountItem) FilterValue() string {
	return i.address.Hex()
}

// AccountsLoadedMsg is sent when accounts are loaded from keystore
type AccountsLoadedMsg struct {
	Accounts []common.Address
	Error    error
}

// WalletLoadedMsg is sent when wallet is successfully loaded
type WalletLoadedMsg struct {
	Wallet *wallet.Wallet
	Error  error
}

// loadAccounts loads accounts from keystore
func (m Model) loadAccounts() tea.Cmd {
	return func() tea.Msg {
		if m.config == nil || m.config.KeystorePath == "" {
			return AccountsLoadedMsg{
				Error: fmt.Errorf("no keystore path configured"),
			}
		}

		accounts, err := wallet.ListKeystoreAccounts(m.config.KeystorePath)
		if err != nil {
			return AccountsLoadedMsg{
				Error: fmt.Errorf("failed to list accounts: %w", err),
			}
		}

		if len(accounts) == 0 {
			return AccountsLoadedMsg{
				Error: fmt.Errorf("no accounts found in keystore"),
			}
		}

		return AccountsLoadedMsg{
			Accounts: accounts,
		}
	}
}

// loadWallet loads the wallet from keystore with the given password
func (m Model) loadWallet(address common.Address, password string) tea.Cmd {
	return func() tea.Msg {
		if m.config == nil || m.config.KeystorePath == "" {
			return WalletLoadedMsg{
				Error: fmt.Errorf("no keystore path configured"),
			}
		}

		w, err := wallet.SelectAccountFromKeystore(
			m.config.KeystorePath,
			address.Hex(),
			password,
		)
		if err != nil {
			return WalletLoadedMsg{
				Error: fmt.Errorf("failed to load wallet: %w", err),
			}
		}

		return WalletLoadedMsg{
			Wallet: w,
		}
	}
}

// updateAccountSelection handles updates for account selection state
func (m Model) updateAccountSelection(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case AccountsLoadedMsg:
		if msg.Error != nil {
			m.err = msg.Error
			m.changeState(StateError)
			return m, nil
		}

		// Store accounts and populate list
		m.availableAccounts = msg.Accounts
		items := make([]list.Item, len(msg.Accounts))
		for i, addr := range msg.Accounts {
			items[i] = accountItem{address: addr, index: i}
		}
		m.accountList.SetItems(items)

		// Set size if we have window dimensions
		if m.width > 0 && m.height > 0 {
			m.accountList.SetSize(m.width-4, m.height-12)
		}

		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			// Get selected account
			selected := m.accountList.SelectedItem()
			if selected == nil {
				return m, nil
			}

			item, ok := selected.(accountItem)
			if !ok {
				return m, nil
			}

			m.selectedAccount = item.index
			// Reset password attempts and error for new account selection
			m.passwordAttempts = 0
			m.err = nil
			m.changeState(StateEnterPassword)
			m.passwordInput.Focus()
			return m, m.passwordInput.Focus()
		}
	}

	// Update the account list
	var cmd tea.Cmd
	m.accountList, cmd = m.accountList.Update(msg)
	return m, cmd
}

// updatePasswordInput handles updates for password input state
func (m Model) updatePasswordInput(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case WalletLoadedMsg:
		m.loadingWallet = false

		// SECURITY: Always clear password immediately, even on error
		m.passwordInput.SetValue("")

		if msg.Error != nil {
			// Password was incorrect or other error
			m.passwordAttempts++

			if m.passwordAttempts >= m.maxPasswordAttempts {
				// Max attempts reached, exit
				m.err = fmt.Errorf("maximum password attempts (%d) exceeded: %w", m.maxPasswordAttempts, msg.Error)
				m.changeState(StateError)
				return m, nil
			}

			// Allow retry - stay on password input screen with error message
			m.err = msg.Error
			m.passwordInput.Focus()
			return m, m.passwordInput.Focus()
		}

		// Wallet loaded successfully - reset attempt counter
		m.wallet = msg.Wallet
		m.passwordAttempts = 0

		// Set default recipient to wallet address
		m.recipientInput.SetValue(msg.Wallet.Address.Hex())

		// Check if we should go to resume mode
		if m.resumeMode {
			m.changeState(StateResumeInput)
			m.txHashInput.Focus()
			return m, m.txHashInput.Focus()
		}

		// Go to network selection
		m.changeState(StateSelectNetwork)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			// Go back to account selection
			if !m.loadingWallet {
				m.changeState(StateSelectAccount)
				m.passwordInput.SetValue("")
				m.passwordInput.Blur()
				// Reset password attempts when going back
				m.passwordAttempts = 0
				m.err = nil
			}
			return m, nil

		case "enter":
			if m.loadingWallet {
				return m, nil
			}

			password := m.passwordInput.Value()
			if password == "" {
				// Don't count empty password as an attempt, just show error
				m.err = fmt.Errorf("password cannot be empty")
				return m, nil
			}

			// Clear any previous error
			m.err = nil

			// Start loading wallet
			m.loadingWallet = true
			selectedAddress := m.availableAccounts[m.selectedAccount]
			return m, m.loadWallet(selectedAddress, password)
		}
	}

	// Update the password input if not loading
	if !m.loadingWallet {
		var cmd tea.Cmd
		m.passwordInput, cmd = m.passwordInput.Update(msg)
		return m, cmd
	}

	return m, nil
}

// viewAccountSelection renders the account selection screen
func (m Model) viewAccountSelection() string {
	title := RenderTitle("Select Keystore Account")

	var content string
	if len(m.availableAccounts) == 0 {
		content = lipgloss.JoinHorizontal(lipgloss.Left, m.spinner.View(), " Loading accounts...")
	} else {
		content = m.accountList.View()
	}

	help := RenderHelp("↑/↓: navigate • enter: select • q: quit")

	return lipgloss.JoinVertical(lipgloss.Left, title, "", content, "", help)
}

// viewPasswordInput renders the password input screen
func (m Model) viewPasswordInput() string {
	title := RenderTitle("Enter Keystore Password")

	selectedAddr := ""
	if m.selectedAccount < len(m.availableAccounts) {
		selectedAddr = m.availableAccounts[m.selectedAccount].Hex()
	}

	accountInfo := RenderInfo(fmt.Sprintf("Account: %s", selectedAddr))

	elements := []string{title, "", accountInfo, ""}

	// Show error message if there was a failed attempt
	if m.err != nil && m.passwordAttempts > 0 {
		attemptsLeft := m.maxPasswordAttempts - m.passwordAttempts
		elements = append(elements, RenderError(fmt.Sprintf("Incorrect password (attempt %d/%d). %d attempt(s) remaining.",
			m.passwordAttempts, m.maxPasswordAttempts, attemptsLeft)))
		elements = append(elements, "")
	} else if m.err != nil {
		// Show error for empty password or other validation errors (doesn't count as attempt)
		elements = append(elements, RenderWarning(m.err.Error()))
		elements = append(elements, "")
	}

	var content string
	if m.loadingWallet {
		content = lipgloss.JoinHorizontal(lipgloss.Left, m.spinner.View(), " Loading wallet...")
	} else {
		passwordView := FocusedInputStyle.Render(m.passwordInput.View())
		content = lipgloss.JoinVertical(lipgloss.Left, "Password:", passwordView)
	}

	elements = append(elements, content, "", RenderHelp("enter: unlock • esc: back • q: quit"))

	return lipgloss.JoinVertical(lipgloss.Left, elements...)
}
