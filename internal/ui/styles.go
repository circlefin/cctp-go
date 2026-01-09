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

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/list"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/pxgray/cctp-go/internal/logger"
)

// Color palette with adaptive colors for light/dark terminal support
var (
	ColorPrimary = lipgloss.AdaptiveColor{
		Light: "#2563EB", // Darker blue for light terminals
		Dark:  "#60A5FA", // Lighter blue for dark terminals
	}
	ColorSuccess = lipgloss.AdaptiveColor{
		Light: "#059669", // Darker green for light terminals
		Dark:  "#34D399", // Lighter green for dark terminals
	}
	ColorError = lipgloss.AdaptiveColor{
		Light: "#DC2626", // Darker red for light terminals
		Dark:  "#F87171", // Lighter red for dark terminals
	}
	ColorWarning = lipgloss.AdaptiveColor{
		Light: "#D97706", // Darker yellow for light terminals
		Dark:  "#FBBF24", // Lighter yellow for dark terminals
	}
	ColorSecondary = lipgloss.AdaptiveColor{
		Light: "#4B5563", // Darker gray for light terminals
		Dark:  "#9CA3AF", // Lighter gray for dark terminals
	}
	ColorAccent = lipgloss.AdaptiveColor{
		Light: "#7C3AED", // Darker purple for light terminals
		Dark:  "#A78BFA", // Lighter purple for dark terminals
	}
	ColorWhite = lipgloss.AdaptiveColor{
		Light: "#FFFFFF",
		Dark:  "#FFFFFF",
	}
	ColorBlack = lipgloss.AdaptiveColor{
		Light: "#000000",
		Dark:  "#000000",
	}
)

// Global styles
var (
	// Header styles
	TitleStyle = lipgloss.NewStyle().
			Foreground(ColorPrimary).
			Bold(true).
			MarginBottom(1)

	SubtitleStyle = lipgloss.NewStyle().
			Foreground(ColorSecondary).
			Italic(true)

	// Message styles
	SuccessStyle = lipgloss.NewStyle().
			Foreground(ColorSuccess).
			Bold(true)

	ErrorStyle = lipgloss.NewStyle().
			Foreground(ColorError).
			Bold(true)

	WarningStyle = lipgloss.NewStyle().
			Foreground(ColorWarning)

	InfoStyle = lipgloss.NewStyle().
			Foreground(ColorPrimary)

	// Chain badge style
	ChainBadgeStyle = lipgloss.NewStyle().
			Foreground(ColorWhite).
			Background(ColorPrimary).
			Padding(0, 1).
			Bold(true)

	TestnetBadgeStyle = lipgloss.NewStyle().
				Foreground(ColorWhite).
				Background(ColorWarning).
				Padding(0, 1).
				Bold(true)

	// Network badge styles
	MainnetBadgeStyle = lipgloss.NewStyle().
				Foreground(ColorWhite).
				Background(ColorSuccess).
				Padding(0, 1).
				Bold(true)

	// List styles
	SelectedItemStyle = lipgloss.NewStyle().
				Foreground(ColorPrimary).
				Bold(true).
				PaddingLeft(2)

	UnselectedItemStyle = lipgloss.NewStyle().
				Foreground(ColorSecondary).
				PaddingLeft(2)

	// Progress styles
	ProgressBarStyle = lipgloss.NewStyle().
				Foreground(ColorPrimary)

	ProgressCompleteStyle = lipgloss.NewStyle().
				Foreground(ColorSuccess)

	// Input styles
	FocusedInputStyle = lipgloss.NewStyle().
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(ColorPrimary).
				Padding(0, 1)

	BlurredInputStyle = lipgloss.NewStyle().
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(ColorSecondary).
				Padding(0, 1)

	// Transaction link style
	TxLinkStyle = lipgloss.NewStyle().
			Foreground(ColorAccent).
			Underline(true)

	// Step indicator styles
	StepActiveStyle = lipgloss.NewStyle().
			Foreground(ColorPrimary).
			Bold(true)

	StepCompleteStyle = lipgloss.NewStyle().
				Foreground(ColorSuccess)

	StepPendingStyle = lipgloss.NewStyle().
				Foreground(ColorSecondary)

	// Breadcrumb styles
	BreadcrumbActiveStyle = lipgloss.NewStyle().
				Foreground(ColorPrimary).
				Bold(true)

	BreadcrumbCompleteStyle = lipgloss.NewStyle().
				Foreground(ColorSuccess)

	BreadcrumbPendingStyle = lipgloss.NewStyle().
				Foreground(ColorSecondary)

	BreadcrumbSeparatorStyle = lipgloss.NewStyle().
					Foreground(ColorSecondary)

	// Container styles
	BoxStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(ColorPrimary).
			Padding(1, 2).
			MarginBottom(1)

	// Help text style
	HelpStyle = lipgloss.NewStyle().
			Foreground(ColorSecondary).
			Italic(true).
			MarginTop(1)

	// Spinner style
	SpinnerStyle = lipgloss.NewStyle().
			Foreground(ColorPrimary)

	// Log styles
	LogStyle = lipgloss.NewStyle().
			Foreground(ColorSecondary).
			Faint(true)

	LogDebugStyle = lipgloss.NewStyle().
			Foreground(ColorSecondary).
			Faint(true)

	LogInfoStyle = lipgloss.NewStyle().
			Foreground(ColorPrimary)

	LogWarnStyle = lipgloss.NewStyle().
			Foreground(ColorWarning)

	LogErrorStyle = lipgloss.NewStyle().
			Foreground(ColorError).
			Bold(true)

	// Table styles
	TableHeaderStyle = lipgloss.NewStyle().
				Foreground(ColorPrimary).
				Bold(true).
				Align(lipgloss.Center)

	TableCellStyle = lipgloss.NewStyle().
			Padding(0, 1)

	TableBorderStyle = lipgloss.NewStyle().
				Foreground(ColorPrimary)
)

// Render functions
func RenderTitle(text string) string {
	return TitleStyle.Render(text)
}

func RenderSubtitle(text string) string {
	return SubtitleStyle.Render(text)
}

func RenderSuccess(text string) string {
	return SuccessStyle.Render("✓ " + text)
}

func RenderError(text string) string {
	return ErrorStyle.Render("✗ " + text)
}

func RenderWarning(text string) string {
	return WarningStyle.Render("⚠ " + text)
}

func RenderInfo(text string) string {
	return InfoStyle.Render("ℹ " + text)
}

func RenderChainBadge(chainName string, isTestnet bool) string {
	if isTestnet {
		return TestnetBadgeStyle.Render(chainName + " (TESTNET)")
	}
	return ChainBadgeStyle.Render(chainName)
}

func RenderNetworkBadge(network string) string {
	if network == "Testnet" {
		return TestnetBadgeStyle.Render("🧪 TESTNET")
	}
	return MainnetBadgeStyle.Render("🌐 MAINNET")
}

func RenderNetworkIndicator(testnet bool) string {
	if testnet {
		return RenderNetworkBadge("Testnet")
	}
	return RenderNetworkBadge("Mainnet")
}

func RenderTxLink(txHash, explorerURL string) string {
	return TxLinkStyle.Render(explorerURL + "/tx/" + txHash)
}

func RenderBox(content string) string {
	return BoxStyle.Render(content)
}

func RenderHelp(text string) string {
	return HelpStyle.Render(text)
}

func RenderLogEntry(entry logger.LogEntry) string {
	timestamp := entry.Timestamp.Format("15:04:05")
	levelStr := entry.Level.String()

	var style lipgloss.Style
	switch entry.Level {
	case logger.DEBUG:
		style = LogDebugStyle
	case logger.INFO:
		style = LogInfoStyle
	case logger.WARN:
		style = LogWarnStyle
	case logger.ERROR:
		style = LogErrorStyle
	default:
		style = LogStyle
	}

	return style.Render(fmt.Sprintf("[%s] [%s] %s", timestamp, levelStr, entry.Message))
}

// Layout helpers
func CenterHorizontal(width int, text string) string {
	return lipgloss.Place(
		width,
		1,
		lipgloss.Center,
		lipgloss.Center,
		text,
	)
}

func CenterVertical(height int, text string) string {
	return lipgloss.Place(
		1,
		height,
		lipgloss.Center,
		lipgloss.Center,
		text,
	)
}

func Center(width, height int, text string) string {
	return lipgloss.Place(
		width,
		height,
		lipgloss.Center,
		lipgloss.Center,
		text,
	)
}

// RenderBreadcrumb renders a breadcrumb navigation indicator
func RenderBreadcrumb(currentState AppState, resumeMode bool) string {
	separator := BreadcrumbSeparatorStyle.Render(" → ")

	if resumeMode {
		// Different breadcrumb for resume flow
		steps := []struct {
			name  string
			state AppState
		}{
			{"Menu", StateMainMenu},
			{"Resume", StateResumeInput},
			{"Progress", StateProgress},
		}

		var parts []string
		for i, step := range steps {
			var rendered string
			if step.state == currentState {
				rendered = BreadcrumbActiveStyle.Render(step.name)
			} else if i < getStateOrder(currentState, resumeMode) {
				rendered = BreadcrumbCompleteStyle.Render(step.name)
			} else {
				rendered = BreadcrumbPendingStyle.Render(step.name)
			}
			parts = append(parts, rendered)
			if i < len(steps)-1 {
				parts = append(parts, separator)
			}
		}

		return lipgloss.JoinHorizontal(lipgloss.Left, parts...)
	}

	// Normal transfer flow
	steps := []struct {
		name  string
		state AppState
	}{
		{"Menu", StateMainMenu},
		{"Source", StateSelectSourceChain},
		{"Destination", StateSelectDestChain},
		{"Details", StateTransferInput},
		{"Progress", StateProgress},
	}

	var parts []string
	for i, step := range steps {
		var rendered string
		if step.state == currentState {
			rendered = BreadcrumbActiveStyle.Render(step.name)
		} else if i < getStateOrder(currentState, resumeMode) {
			rendered = BreadcrumbCompleteStyle.Render(step.name)
		} else {
			rendered = BreadcrumbPendingStyle.Render(step.name)
		}
		parts = append(parts, rendered)
		if i < len(steps)-1 {
			parts = append(parts, separator)
		}
	}

	return lipgloss.JoinHorizontal(lipgloss.Left, parts...)
}

// getStateOrder returns the order of a state in the flow
func getStateOrder(state AppState, resumeMode bool) int {
	if resumeMode {
		switch state {
		case StateMainMenu:
			return 0
		case StateResumeInput:
			return 1
		case StateProgress, StateComplete, StateError:
			return 2
		default:
			return 0
		}
	}

	switch state {
	case StateMainMenu:
		return 0
	case StateSelectSourceChain:
		return 1
	case StateSelectDestChain:
		return 2
	case StateTransferInput:
		return 3
	case StateProgress, StateComplete, StateError:
		return 4
	default:
		return 0
	}
}

// Table rendering utilities

// RenderTransactionTable renders a table of transaction hashes
func RenderTransactionTable(transactions []TransactionRow) string {
	if len(transactions) == 0 {
		return ""
	}

	t := table.New().
		Border(lipgloss.RoundedBorder()).
		BorderStyle(TableBorderStyle).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == table.HeaderRow {
				return TableHeaderStyle
			}
			return TableCellStyle
		}).
		Headers("Step", "Transaction Hash", "Explorer Link")

	for _, tx := range transactions {
		t.Row(tx.Step, tx.TxHash, tx.ExplorerLink)
	}

	return t.Render()
}

// TransactionRow represents a row in the transaction table
type TransactionRow struct {
	Step         string
	TxHash       string
	ExplorerLink string
}

// RenderChainInfoTable renders chain information in a table format
func RenderChainInfoTable(sourceChain, destChain string) string {
	t := table.New().
		Border(lipgloss.RoundedBorder()).
		BorderStyle(TableBorderStyle).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == table.HeaderRow {
				return TableHeaderStyle
			}
			return TableCellStyle
		}).
		Headers("", "Chain").
		Row("Source", sourceChain).
		Row("Destination", destChain)

	return t.Render()
}

// Layout utility functions

// RenderTwoColumn creates a two-column layout
func RenderTwoColumn(left, right string) string {
	return lipgloss.JoinHorizontal(lipgloss.Top, left, "  ", right)
}

// RenderThreeColumn creates a three-column layout
func RenderThreeColumn(left, middle, right string) string {
	return lipgloss.JoinHorizontal(lipgloss.Top, left, "  ", middle, "  ", right)
}

// RenderSection creates a section with title and content
func RenderSection(title, content string) string {
	return lipgloss.JoinVertical(lipgloss.Left, RenderTitle(title), "", content)
}

// RenderCard creates a bordered card with content
func RenderCard(content string, width int) string {
	style := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(ColorPrimary).
		Padding(1, 2)

	if width > 0 {
		style = style.Width(width)
	}

	return style.Render(content)
}

// RenderInfoBox creates an information box with a title and content
func RenderInfoBox(title, content string) string {
	titleStyle := lipgloss.NewStyle().
		Foreground(ColorPrimary).
		Bold(true).
		Padding(0, 1)

	contentStyle := lipgloss.NewStyle().
		Padding(0, 1)

	boxStyle := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(ColorPrimary)

	return boxStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			titleStyle.Render(title),
			contentStyle.Render(content),
		),
	)
}

// WithMaxWidth applies a maximum width to text, wrapping as needed
func WithMaxWidth(text string, maxWidth int) string {
	if maxWidth <= 0 {
		return text
	}
	return lipgloss.NewStyle().MaxWidth(maxWidth).Render(text)
}

// WithMaxHeight applies a maximum height to text, truncating as needed
func WithMaxHeight(text string, maxHeight int) string {
	if maxHeight <= 0 {
		return text
	}
	return lipgloss.NewStyle().MaxHeight(maxHeight).Render(text)
}

// RenderStepList renders a list of steps with custom styling based on status
func RenderStepList(steps []StepItem) string {
	items := make([]any, len(steps))
	for i, step := range steps {
		items[i] = step.Text
	}

	l := list.New(items...)

	// Use custom enumerator function for step indicators
	l.Enumerator(func(items list.Items, i int) string {
		if i >= len(steps) {
			return ""
		}
		step := steps[i]
		switch step.Status {
		case "complete":
			return "✓ "
		case "active":
			return "◉ "
		case "pending":
			return "○ "
		default:
			return "○ "
		}
	})

	// Apply styling based on step status
	l.ItemStyleFunc(func(items list.Items, i int) lipgloss.Style {
		if i >= len(steps) {
			return lipgloss.NewStyle()
		}
		step := steps[i]
		switch step.Status {
		case "complete":
			return StepCompleteStyle
		case "active":
			return StepActiveStyle
		case "pending":
			return StepPendingStyle
		default:
			return StepPendingStyle
		}
	})

	return l.String()
}

// StepItem represents a step in a process
type StepItem struct {
	Text   string
	Status string // "complete", "active", "pending"
}
