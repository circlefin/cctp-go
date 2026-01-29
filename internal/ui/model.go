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
	"math/big"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/circlefin/cctp-go"
	"github.com/circlefin/cctp-go/internal/config"
	"github.com/circlefin/cctp-go/internal/logger"
	"github.com/circlefin/cctp-go/internal/wallet"
)

// AppState represents the current state of the application
type AppState int

const (
	StateSelectAccount AppState = iota
	StateEnterPassword
	StateSelectNetwork
	StateMainMenu
	StateSelectSourceChain
	StateSelectDestChain
	StateTransferInput
	StateResumeInput
	StateProgress
	StateComplete
	StateError
)

// Model is the main Bubbletea model
type Model struct {
	// State
	state            AppState
	previousState    AppState
	err              error
	testnet          bool
	networkSetByUser bool // true if user selected network in TUI
	config           *config.Config
	wallet           *wallet.Wallet

	// Wallet loading
	availableAccounts   []common.Address
	selectedAccount     int
	accountList         list.Model
	passwordInput       textinput.Model
	loadingWallet       bool
	passwordAttempts    int
	maxPasswordAttempts int

	// Chain selection
	chains      []cctp.Chain
	sourceChain *cctp.Chain
	destChain   *cctp.Chain
	chainList   list.Model

	// Network selection
	networkList list.Model

	// Main menu
	mainMenuList list.Model

	// Transfer input
	amountInput    textinput.Model
	recipientInput textinput.Model
	txHashInput    textinput.Model
	inputFocus     int
	resumeMode     bool
	transferType   cctp.TransferType
	usdcBalance    *big.Int
	loadingBalance bool

	// Progress
	spinner       spinner.Model
	progress      progress.Model
	currentStep   cctp.TransferStep
	stepMessage   string
	txHashes      map[cctp.TransferStep]string
	progressValue float64
	updatesChan   chan cctp.TransferUpdate

	// Logging
	logEntries []logger.LogEntry

	// Window size
	width  int
	height int
}

// NewModel creates a new UI model
func NewModel(cfg *config.Config) Model {
	testnet := cfg.Testnet

	// Create chain list with RPC overrides from config
	chains := cctp.GetChains(testnet)
	chains = cctp.ApplyRPCOverrides(chains, cfg.RPCUrls)
	items := make([]list.Item, len(chains))
	for i, chain := range chains {
		items[i] = chainItem{chain: chain}
	}

	chainList := list.New(items, list.NewDefaultDelegate(), 0, 0)
	chainList.Title = "Select Source Chain"
	chainList.SetShowHelp(true)

	// Create network list
	networkItems := []list.Item{
		networkItem{name: "Mainnet", description: "Production - Real USDC", isTestnet: false},
		networkItem{name: "Testnet", description: "Testing - Test USDC", isTestnet: true},
	}
	networkList := list.New(networkItems, list.NewDefaultDelegate(), 0, 0)
	networkList.Title = "Select Network"
	networkList.SetShowHelp(false)
	networkList.SetShowStatusBar(false)
	networkList.SetFilteringEnabled(false)

	// Create main menu list
	mainMenuItems := []list.Item{
		mainMenuItem{action: "transfer", name: "Start New Transfer", description: "Initiate a new cross-chain USDC transfer"},
		mainMenuItem{action: "resume", name: "Resume Transfer", description: "Continue an existing transfer using burn transaction hash"},
	}
	mainMenuList := list.New(mainMenuItems, list.NewDefaultDelegate(), 0, 0)
	mainMenuList.Title = "Main Menu"
	mainMenuList.SetShowHelp(false)
	mainMenuList.SetShowStatusBar(false)
	mainMenuList.SetFilteringEnabled(false)

	// Create empty account list (will be populated when accounts are loaded)
	accountList := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	accountList.Title = "Available Accounts"
	accountList.SetShowHelp(false)
	accountList.SetShowStatusBar(false)
	accountList.SetFilteringEnabled(false)

	// Create spinner
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = SpinnerStyle

	// Create progress bar
	prog := progress.New(progress.WithDefaultGradient())

	// Create password input
	passwordInput := textinput.New()
	passwordInput.Placeholder = "Enter keystore password"
	passwordInput.CharLimit = 256
	passwordInput.Width = 50
	passwordInput.EchoMode = textinput.EchoPassword
	passwordInput.EchoCharacter = '•'

	// Create amount input
	amountInput := textinput.New()
	amountInput.Placeholder = "0.00"
	amountInput.CharLimit = 20
	amountInput.Width = 30

	// Create recipient input
	recipientInput := textinput.New()
	recipientInput.Placeholder = "0x..."
	recipientInput.CharLimit = 42
	recipientInput.Width = 50

	// Create transaction hash input
	txHashInput := textinput.New()
	txHashInput.Placeholder = "0x..."
	txHashInput.CharLimit = 66
	txHashInput.Width = 70

	// Determine initial state based on wallet source
	initialState := StateSelectNetwork // Start with network selection
	var w *wallet.Wallet
	networkSetByUser := false

	// Check if testnet was set via CLI flag (not from config file default)
	// If testnet flag was explicitly set, skip network selection
	if cfg.Testnet {
		// Check if it was set via flag by seeing if it differs from default
		// For now, if testnet is true, we'll still show the selector
		// User can always skip by pressing the default
		networkSetByUser = false
	}

	// Try loading from PRIVATE_KEY environment variable first
	privateKey := ""
	if pk := cfg.PrivateKey; pk != "" {
		privateKey = pk
	}

	if privateKey != "" {
		// Load wallet directly from private key
		var err error
		w, err = wallet.LoadFromPrivateKey(privateKey)
		if err != nil {
			// Handle error in the UI
			initialState = StateError
		}
		// If wallet loaded, start with network selection
		initialState = StateSelectNetwork
	} else if cfg.KeystorePath != "" {
		// Need to load from keystore - start with account selection
		// Network selection will come after wallet is loaded
		initialState = StateSelectAccount
	}

	// Set default recipient to wallet address if available
	if w != nil {
		recipientInput.SetValue(w.Address.Hex())
	}

	return Model{
		state:               initialState,
		testnet:             testnet,
		networkSetByUser:    networkSetByUser,
		config:              cfg,
		wallet:              w,
		chains:              chains,
		chainList:           chainList,
		networkList:         networkList,
		mainMenuList:        mainMenuList,
		accountList:         accountList,
		passwordInput:       passwordInput,
		spinner:             s,
		progress:            prog,
		amountInput:         amountInput,
		recipientInput:      recipientInput,
		txHashInput:         txHashInput,
		txHashes:            make(map[cctp.TransferStep]string),
		logEntries:          make([]logger.LogEntry, 0),
		passwordAttempts:    0,
		maxPasswordAttempts: 3,
	}
}

// Init initializes the model
func (m Model) Init() tea.Cmd {
	cmds := []tea.Cmd{m.spinner.Tick}

	// If starting with account selection, load accounts
	if m.state == StateSelectAccount && m.config != nil && m.config.KeystorePath != "" {
		cmds = append(cmds, m.loadAccounts())
	}

	return tea.Batch(cmds...)
}

// Update handles messages and updates the model
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.chainList.SetSize(msg.Width, msg.Height-10)
		m.networkList.SetSize(msg.Width, msg.Height-10)
		m.mainMenuList.SetSize(msg.Width, msg.Height-12)
		m.accountList.SetSize(msg.Width-4, msg.Height-12)
		return m, nil
	case spinner.TickMsg:
		// Handle spinner tick at the top level so it animates in all states
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	// Delegate to state-specific update functions
	switch m.state {
	case StateSelectAccount:
		return m.updateAccountSelection(msg)
	case StateEnterPassword:
		return m.updatePasswordInput(msg)
	case StateSelectNetwork:
		return m.updateNetworkSelection(msg)
	case StateMainMenu:
		return m.updateMainMenu(msg)
	case StateSelectSourceChain, StateSelectDestChain:
		return m.updateChainSelection(msg)
	case StateTransferInput:
		return m.updateTransferInput(msg)
	case StateResumeInput:
		return m.updateResumeInput(msg)
	case StateProgress:
		return m.updateProgress(msg)
	case StateComplete:
		return m.updateComplete(msg)
	case StateError:
		return m.updateError(msg)
	}

	return m, nil
}

// View renders the UI
func (m Model) View() string {
	switch m.state {
	case StateSelectAccount:
		return m.viewAccountSelection()
	case StateEnterPassword:
		return m.viewPasswordInput()
	case StateSelectNetwork:
		return m.viewNetworkSelection()
	case StateMainMenu:
		return m.viewMainMenu()
	case StateSelectSourceChain, StateSelectDestChain:
		return m.viewChainSelection()
	case StateTransferInput:
		return m.viewTransferInput()
	case StateResumeInput:
		return m.viewResumeInput()
	case StateProgress:
		return m.viewProgress()
	case StateComplete:
		return m.viewComplete()
	case StateError:
		return m.viewError()
	}
	return ""
}

// changeState changes the state and tracks the previous state
func (m *Model) changeState(newState AppState) {
	m.previousState = m.state
	m.state = newState
}

// goBack navigates to the previous state with appropriate cleanup
func (m *Model) goBack() {
	switch m.state {
	case StateSelectNetwork:
		// Go back to account selection (if using keystore)
		if m.config != nil && m.config.KeystorePath != "" {
			m.changeState(StateSelectAccount)
		}
		return

	case StateSelectSourceChain:
		// Go back to main menu
		m.changeState(StateMainMenu)
		m.sourceChain = nil

	case StateSelectDestChain:
		// Go back to source chain selection
		m.changeState(StateSelectSourceChain)
		m.destChain = nil
		// Restore full chain list
		items := make([]list.Item, len(m.chains))
		for i, chain := range m.chains {
			items[i] = chainItem{chain: chain}
		}
		m.chainList.SetItems(items)
		m.chainList.ResetFilter() // Clear any active filter
		m.chainList.Title = "Select Source Chain"

	case StateTransferInput:
		// Go back to destination chain selection
		m.changeState(StateSelectDestChain)
		// Restore destination chain list (exclude source chain)
		items := make([]list.Item, 0)
		for _, chain := range m.chains {
			if m.sourceChain != nil && chain.Domain != m.sourceChain.Domain {
				items = append(items, chainItem{chain: chain})
			}
		}
		m.chainList.SetItems(items)
		m.chainList.ResetFilter() // Clear any active filter
		m.chainList.Title = "Select Destination Chain"
		m.destChain = nil
		// Reset balance state
		m.usdcBalance = nil
		m.loadingBalance = false

	case StateResumeInput:
		// Go back to main menu
		m.changeState(StateMainMenu)
		m.txHashInput.SetValue("")
		m.txHashInput.Blur()
	}
}

// updateMainMenu handles the main menu state
func (m Model) updateMainMenu(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			// Go back to network selection
			m.changeState(StateSelectNetwork)
			return m, nil
		case "w":
			// Go back to wallet selection (if using keystore)
			if m.config != nil && m.config.KeystorePath != "" {
				m.changeState(StateSelectAccount)
				// SECURITY: Clear password when changing wallet
				m.passwordInput.SetValue("")
				m.passwordInput.Blur()
				return m, nil
			}
			return m, nil
		case "enter":
			// Get selected menu item
			selectedItem, ok := m.mainMenuList.SelectedItem().(mainMenuItem)
			if !ok {
				return m, nil
			}

			switch selectedItem.action {
			case "transfer":
				// Start new transfer
				m.changeState(StateSelectSourceChain)
				m.resumeMode = false
				// Refresh chain list with current network's chains
				items := make([]list.Item, len(m.chains))
				for i, chain := range m.chains {
					items[i] = chainItem{chain: chain}
				}
				m.chainList.SetItems(items)
				m.chainList.ResetFilter() // Clear any active filter from previous transfer
				m.chainList.Title = "Select Source Chain"
				return m, nil
			case "resume":
				// Resume transfer
				m.changeState(StateResumeInput)
				m.resumeMode = true
				m.txHashInput.Focus()
				return m, m.txHashInput.Focus()
			}
			return m, nil
		}
	}

	// Update main menu list
	var cmd tea.Cmd
	m.mainMenuList, cmd = m.mainMenuList.Update(msg)
	return m, cmd
}

// viewMainMenu renders the main menu
func (m Model) viewMainMenu() string {
	network := RenderNetworkIndicator(m.testnet)
	title := RenderTitle("CCTP V2 CLI - Cross-Chain USDC Transfers")
	wallet := RenderInfo(fmt.Sprintf("Wallet: %s", m.wallet.Address.Hex()))

	// Build help text based on wallet source
	helpText := "↑/↓: navigate • enter: select • esc: change network • q: quit"
	if m.config != nil && m.config.KeystorePath != "" {
		helpText = "↑/↓: navigate • enter: select • w: change wallet • esc: change network • q: quit"
	}
	help := RenderHelp(helpText)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		network,
		"",
		title,
		wallet,
		"",
		m.mainMenuList.View(),
		"",
		help,
	)
}

// chainItem implements list.Item for chain selection
type chainItem struct {
	chain cctp.Chain
}

func (i chainItem) Title() string {
	return i.chain.Name
}

func (i chainItem) Description() string {
	return fmt.Sprintf("Domain: %d | Chain ID: %s", i.chain.Domain, i.chain.ChainID.String())
}

func (i chainItem) FilterValue() string {
	return i.chain.Name
}

// networkItem implements list.Item for network selection
type networkItem struct {
	name        string
	description string
	isTestnet   bool
}

func (i networkItem) Title() string {
	return i.name
}

func (i networkItem) Description() string {
	return i.description
}

func (i networkItem) FilterValue() string {
	return i.name
}

// mainMenuItem implements list.Item for main menu selection
type mainMenuItem struct {
	action      string
	name        string
	description string
}

func (i mainMenuItem) Title() string {
	return i.name
}

func (i mainMenuItem) Description() string {
	return i.description
}

func (i mainMenuItem) FilterValue() string {
	return i.name
}

// SetSourceChain sets the source chain for the transfer
func (m *Model) SetSourceChain(chain *cctp.Chain) {
	m.sourceChain = chain
}

// SetDestChain sets the destination chain for the transfer
func (m *Model) SetDestChain(chain *cctp.Chain) {
	m.destChain = chain
}

// GetSourceChain returns the currently selected source chain
func (m *Model) GetSourceChain() *cctp.Chain {
	return m.sourceChain
}

// GetDestChain returns the currently selected destination chain
func (m *Model) GetDestChain() *cctp.Chain {
	return m.destChain
}

// AdvanceToTransferInput advances the model to the transfer input state
func (m *Model) AdvanceToTransferInput() {
	m.state = StateTransferInput
	m.amountInput.Focus()
}

// SetTransferAmount sets the transfer amount value
func (m *Model) SetTransferAmount(amount string) {
	m.amountInput.SetValue(amount)
}

// SetTransferRecipient sets the recipient address value
func (m *Model) SetTransferRecipient(recipient string) {
	m.recipientInput.SetValue(recipient)
}

// SetTransferType sets the transfer type
func (m *Model) SetTransferType(tt cctp.TransferType) {
	m.transferType = tt
}

// resetForNewTransfer resets transfer-specific state for a new transfer
func (m *Model) resetForNewTransfer() {
	m.sourceChain = nil
	m.destChain = nil
	m.amountInput.SetValue("")
	m.recipientInput.SetValue(m.wallet.Address.Hex()) // Reset to wallet address
	m.transferType = cctp.TransferTypeAuto
	m.txHashInput.SetValue("")
	m.currentStep = ""
	m.stepMessage = ""
	m.txHashes = make(map[cctp.TransferStep]string)
	m.progressValue = 0
	m.err = nil
	m.resumeMode = false
	m.inputFocus = 0
	m.usdcBalance = nil
	m.loadingBalance = false
}

// SetResumeMode sets the model to resume mode with optional tx hash
// Only applies if wallet is already loaded, otherwise it's deferred
func (m *Model) SetResumeMode(txHash string) {
	m.resumeMode = true
	if txHash != "" {
		m.txHashInput.SetValue(txHash)
	}

	// Only switch to resume input if wallet is loaded
	// Otherwise, wallet loading will happen first
	if m.wallet != nil {
		m.state = StateResumeInput
		m.txHashInput.Focus()
	}
}
