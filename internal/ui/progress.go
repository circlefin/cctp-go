package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/pxgray/cctp-go"
	"github.com/pxgray/cctp-go/internal/logger"
)

// updateProgress handles updates for progress state
func (m Model) updateProgress(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case InitTransferMsg:
		// Store the updates channel and start listening
		m.updatesChan = msg.UpdatesChan
		// Start both update listeners
		return m, tea.Batch(m.waitForNextUpdate(), waitForLogUpdate())

	case LogUpdateMsg:
		// Add log entry to the list
		m.logEntries = append(m.logEntries, msg.Entry)
		// Continue listening for log updates
		return m, waitForLogUpdate()

	case TransferUpdateMsg:
		update := msg.Update

		// Update current step and message
		m.currentStep = update.Step
		m.stepMessage = update.Message

		// Store transaction hash if provided
		if update.TxHash != "" {
			m.txHashes[update.Step] = update.TxHash
		}

		// Store chain information if provided (for resume flow)
		if update.SourceChain != nil {
			m.sourceChain = update.SourceChain
		}
		if update.DestChain != nil {
			m.destChain = update.DestChain
		}

		// Handle errors
		if update.Error != nil {
			m.err = update.Error
			m.changeState(StateError)
			// Write logs to file on error
			_ = logger.GetLogger().WriteToFile()
			return m, nil
		}

		// Check if complete
		if update.Step == cctp.StepComplete {
			m.changeState(StateComplete)
			return m, nil
		}

		// Update progress value based on step
		m.progressValue = getProgressValue(update.Step)

		// Wait for next update
		return m, m.waitForNextUpdate()
	}

	return m, nil
}

// viewProgress renders the progress screen
func (m Model) viewProgress() string {
	network := RenderNetworkIndicator(m.testnet)
	breadcrumb := RenderBreadcrumb(m.state, m.resumeMode)
	title := RenderTitle("Transfer in Progress")

	chainInfo := ""
	if m.sourceChain != nil && m.destChain != nil {
		chainInfo = lipgloss.JoinHorizontal(
			lipgloss.Center,
			RenderChainBadge(m.sourceChain.Name, m.testnet),
			" → ",
			RenderChainBadge(m.destChain.Name, m.testnet),
		)
	} else if m.destChain != nil {
		chainInfo = lipgloss.JoinHorizontal(
			lipgloss.Center,
			"?",
			" → ",
			RenderChainBadge(m.destChain.Name, m.testnet),
		)
	}

	// Progress bar
	progressBar := m.progress.ViewAs(m.progressValue)

	// Current step
	currentStepView := lipgloss.JoinHorizontal(
		lipgloss.Left,
		m.spinner.View(),
		" ",
		StepActiveStyle.Render(m.stepMessage),
	)

	// Build step list using lipgloss/list
	steps := []StepItem{
		{Text: "Check Balance", Status: getStepStatus(cctp.StepCheckBalance, m.currentStep)},
		{Text: "Check Allowance", Status: getStepStatus(cctp.StepCheckAllowance, m.currentStep)},
		{Text: "Approve USDC", Status: getStepStatus(cctp.StepApproving, m.currentStep)},
		{Text: "Burn USDC", Status: getStepStatus(cctp.StepBurning, m.currentStep)},
		{Text: "Poll Attestation", Status: getStepStatus(cctp.StepPollingAttestation, m.currentStep)},
		{Text: "Mint USDC", Status: getStepStatus(cctp.StepMinting, m.currentStep)},
	}

	stepsList := RenderStepList(steps)

	// Transaction links
	txLinks := ""
	if len(m.txHashes) > 0 {
		var txRows []TransactionRow
		displayed := make(map[string]bool)

		// Check for approve tx
		if hash, ok := m.txHashes[cctp.StepApproving]; ok && !displayed[hash] {
			if m.sourceChain != nil {
				txRows = append(txRows, TransactionRow{
					Step:         "Approve",
					TxHash:       formatTxHash(hash),
					ExplorerLink: m.sourceChain.Explorer + "/tx/" + hash,
				})
				displayed[hash] = true
			}
		} else if hash, ok := m.txHashes[cctp.StepApprovingWait]; ok && !displayed[hash] {
			if m.sourceChain != nil {
				txRows = append(txRows, TransactionRow{
					Step:         "Approve",
					TxHash:       formatTxHash(hash),
					ExplorerLink: m.sourceChain.Explorer + "/tx/" + hash,
				})
				displayed[hash] = true
			}
		}

		// Check for burn tx
		if hash, ok := m.txHashes[cctp.StepBurning]; ok && !displayed[hash] {
			if m.sourceChain != nil {
				txRows = append(txRows, TransactionRow{
					Step:         "Burn",
					TxHash:       formatTxHash(hash),
					ExplorerLink: m.sourceChain.Explorer + "/tx/" + hash,
				})
				displayed[hash] = true
			}
		} else if hash, ok := m.txHashes[cctp.StepBurningWait]; ok && !displayed[hash] {
			if m.sourceChain != nil {
				txRows = append(txRows, TransactionRow{
					Step:         "Burn",
					TxHash:       formatTxHash(hash),
					ExplorerLink: m.sourceChain.Explorer + "/tx/" + hash,
				})
				displayed[hash] = true
			}
		}

		// Check for mint tx
		if hash, ok := m.txHashes[cctp.StepMinting]; ok && !displayed[hash] {
			if m.destChain != nil {
				txRows = append(txRows, TransactionRow{
					Step:         "Mint",
					TxHash:       formatTxHash(hash),
					ExplorerLink: m.destChain.Explorer + "/tx/" + hash,
				})
				displayed[hash] = true
			}
		} else if hash, ok := m.txHashes[cctp.StepMintingWait]; ok && !displayed[hash] {
			if m.destChain != nil {
				txRows = append(txRows, TransactionRow{
					Step:         "Mint",
					TxHash:       formatTxHash(hash),
					ExplorerLink: m.destChain.Explorer + "/tx/" + hash,
				})
				displayed[hash] = true
			}
		}

		if len(txRows) > 0 {
			txLinks = "\n\n" + RenderTitle("Transaction Hashes:") + "\n" + RenderTransactionTable(txRows)
		}
	}

	// Log pane - show last 5 log entries with max width constraint
	var logPaneElements []string
	if len(m.logEntries) > 0 {
		logPaneElements = append(logPaneElements, "", "") // Two empty lines for spacing
		logPaneElements = append(logPaneElements, RenderTitle("Recent Logs:"))

		// Get last 5 entries
		startIdx := len(m.logEntries) - 5
		if startIdx < 0 {
			startIdx = 0
		}

		// Apply max width to prevent overflow
		maxWidth := m.width
		if maxWidth > 100 {
			maxWidth = 100 // Cap at 100 characters for readability
		}

		for _, entry := range m.logEntries[startIdx:] {
			logEntry := RenderLogEntry(entry)
			if maxWidth > 0 {
				logEntry = WithMaxWidth(logEntry, maxWidth)
			}
			logPaneElements = append(logPaneElements, logEntry)
		}
	}

	help := RenderHelp("q: quit")

	// Build the final view using JoinVertical
	elements := []string{
		network,
		breadcrumb,
		"",
		title,
		"",
		chainInfo,
		"",
		progressBar,
		"",
		currentStepView,
		"",
		stepsList,
	}

	if txLinks != "" {
		elements = append(elements, txLinks)
	}

	elements = append(elements, logPaneElements...)
	elements = append(elements, "", help)

	return lipgloss.JoinVertical(lipgloss.Left, elements...)
}

// updateComplete handles updates for complete state
func (m Model) updateComplete(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "m":
			// Go back to main menu for another transfer
			m.resetForNewTransfer()
			m.changeState(StateMainMenu)
			return m, nil
		case "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

// viewComplete renders the completion screen
func (m Model) viewComplete() string {
	network := RenderNetworkIndicator(m.testnet)
	title := RenderSuccess("Transfer Complete!")

	chainInfo := ""
	if m.sourceChain != nil && m.destChain != nil {
		chainInfo = lipgloss.JoinHorizontal(
			lipgloss.Center,
			RenderChainBadge(m.sourceChain.Name, m.testnet),
			" → ",
			RenderChainBadge(m.destChain.Name, m.testnet),
		)
	} else if m.destChain != nil {
		chainInfo = lipgloss.JoinHorizontal(
			lipgloss.Left,
			"Destination: ",
			RenderChainBadge(m.destChain.Name, m.testnet),
		)
	}

	message := RenderInfo("Your USDC has been successfully transferred!")

	// Transaction links
	var txRows []TransactionRow
	displayed := make(map[string]bool)

	// Check for approve tx
	if hash, ok := m.txHashes[cctp.StepApproving]; ok && !displayed[hash] {
		if m.sourceChain != nil {
			txRows = append(txRows, TransactionRow{
				Step:         "Approve",
				TxHash:       formatTxHash(hash),
				ExplorerLink: m.sourceChain.Explorer + "/tx/" + hash,
			})
			displayed[hash] = true
		}
	} else if hash, ok := m.txHashes[cctp.StepApprovingWait]; ok && !displayed[hash] {
		if m.sourceChain != nil {
			txRows = append(txRows, TransactionRow{
				Step:         "Approve",
				TxHash:       formatTxHash(hash),
				ExplorerLink: m.sourceChain.Explorer + "/tx/" + hash,
			})
			displayed[hash] = true
		}
	}

	// Check for burn tx
	if hash, ok := m.txHashes[cctp.StepBurning]; ok && !displayed[hash] {
		if m.sourceChain != nil {
			txRows = append(txRows, TransactionRow{
				Step:         "Burn",
				TxHash:       formatTxHash(hash),
				ExplorerLink: m.sourceChain.Explorer + "/tx/" + hash,
			})
			displayed[hash] = true
		}
	} else if hash, ok := m.txHashes[cctp.StepBurningWait]; ok && !displayed[hash] {
		if m.sourceChain != nil {
			txRows = append(txRows, TransactionRow{
				Step:         "Burn",
				TxHash:       formatTxHash(hash),
				ExplorerLink: m.sourceChain.Explorer + "/tx/" + hash,
			})
			displayed[hash] = true
		}
	}

	// Check for mint tx
	if hash, ok := m.txHashes[cctp.StepMinting]; ok && !displayed[hash] {
		if m.destChain != nil {
			txRows = append(txRows, TransactionRow{
				Step:         "Mint",
				TxHash:       formatTxHash(hash),
				ExplorerLink: m.destChain.Explorer + "/tx/" + hash,
			})
			displayed[hash] = true
		}
	} else if hash, ok := m.txHashes[cctp.StepMintingWait]; ok && !displayed[hash] {
		if m.destChain != nil {
			txRows = append(txRows, TransactionRow{
				Step:         "Mint",
				TxHash:       formatTxHash(hash),
				ExplorerLink: m.destChain.Explorer + "/tx/" + hash,
			})
			displayed[hash] = true
		}
	}

	txLinks := lipgloss.JoinVertical(
		lipgloss.Left,
		RenderTitle("Transaction Hashes:"),
		RenderTransactionTable(txRows),
	)

	help := RenderHelp("m: main menu • q: quit")

	return lipgloss.JoinVertical(
		lipgloss.Left,
		network,
		"",
		title,
		"",
		chainInfo,
		"",
		message,
		"",
		txLinks,
		"",
		help,
	)
}

// updateError handles updates for error state
func (m Model) updateError(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "m":
			// Go back to main menu to try again
			m.resetForNewTransfer()
			m.changeState(StateMainMenu)
			return m, nil
		case "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

// viewError renders the error screen
func (m Model) viewError() string {
	network := RenderNetworkIndicator(m.testnet)
	title := RenderError("Transfer Failed")

	// Apply max width to error message to prevent overflow
	maxWidth := m.width
	if maxWidth > 80 {
		maxWidth = 80 // Cap at 80 characters for readability
	}

	errText := m.err.Error()
	if maxWidth > 0 {
		errText = WithMaxWidth(errText, maxWidth-8) // Account for box padding
	}
	errorMsg := RenderBox(errText)

	elements := []string{
		network,
		"",
		title,
		"",
		errorMsg,
	}

	// Add info about log file if it was created
	if logPath := logger.GetLogger().GetLogFilePath(); logPath != "" {
		elements = append(elements, RenderInfo(fmt.Sprintf("Detailed logs saved to: %s", logPath)))
	}

	help := RenderHelp("m: main menu • q: quit")
	elements = append(elements, "", help)

	return lipgloss.JoinVertical(lipgloss.Left, elements...)
}

// Helper functions

func getStepStatus(step cctp.TransferStep, currentStep cctp.TransferStep) string {
	stepOrder := map[cctp.TransferStep]int{
		cctp.StepCheckBalance:       0,
		cctp.StepCheckAllowance:     1,
		cctp.StepApproving:          2,
		cctp.StepApprovingWait:      2,
		cctp.StepBurning:            3,
		cctp.StepBurningWait:        3,
		cctp.StepPollingAttestation: 4,
		cctp.StepMinting:            5,
		cctp.StepMintingWait:        5,
		cctp.StepComplete:           6,
	}

	currentOrder := stepOrder[currentStep]
	stepOrderValue := stepOrder[step]

	if stepOrderValue < currentOrder {
		return "complete"
	} else if stepOrderValue == currentOrder {
		return "active"
	} else {
		return "pending"
	}
}

func getProgressValue(step cctp.TransferStep) float64 {
	switch step {
	case cctp.StepCheckBalance:
		return 0.1
	case cctp.StepCheckAllowance:
		return 0.2
	case cctp.StepApproving, cctp.StepApprovingWait:
		return 0.3
	case cctp.StepBurning, cctp.StepBurningWait:
		return 0.5
	case cctp.StepPollingAttestation:
		return 0.7
	case cctp.StepMinting, cctp.StepMintingWait:
		return 0.9
	case cctp.StepComplete:
		return 1.0
	default:
		return 0.0
	}
}

// formatTxHash shortens a transaction hash for display
func formatTxHash(hash string) string {
	if len(hash) <= 16 {
		return hash
	}
	return hash[:10] + "..." + hash[len(hash)-6:]
}
