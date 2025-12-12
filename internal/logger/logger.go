package logger

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/charmbracelet/log"
)

// LogLevel represents the severity level of a log entry
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

// String returns the string representation of a log level
func (l LogLevel) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

// LogEntry represents a single log entry
type LogEntry struct {
	Timestamp time.Time
	Level     LogLevel
	Message   string
}

// LogCapture is a custom io.Writer that captures log output
// and optionally writes to a file on error
type LogCapture struct {
	mu            sync.RWMutex
	entries       []LogEntry
	maxEntries    int
	updateChan    chan LogEntry
	errorOccurred bool
	logFilePath   string
}

var (
	defaultCapture *LogCapture
	defaultLogger  *log.Logger
	once           sync.Once
)

// Init initializes the logger and sets up charmbracelet/log
func Init() {
	once.Do(func() {
		defaultCapture = &LogCapture{
			entries:    make([]LogEntry, 0, 100),
			maxEntries: 100, // Keep last 100 entries in memory
			updateChan: make(chan LogEntry, 10),
		}

		// Create charmbracelet logger with our custom writer
		defaultLogger = log.NewWithOptions(defaultCapture, log.Options{
			ReportTimestamp: false, // We'll add timestamps in the capture
			Level:           log.DebugLevel,
		})

		// Set as default logger for the package
		log.SetDefault(defaultLogger)
	})
}

// Write implements io.Writer interface for capturing logs
func (lc *LogCapture) Write(p []byte) (n int, err error) {
	if lc == nil {
		return 0, fmt.Errorf("logger not initialized")
	}

	msg := string(p)
	// Remove trailing newline if present
	if len(msg) > 0 && msg[len(msg)-1] == '\n' {
		msg = msg[:len(msg)-1]
	}

	// Parse log level from charmbracelet/log output
	// Format is typically: "LEVEL message" or with colors/styles
	level := INFO // default level
	cleanMsg := msg

	// Remove ANSI color codes for parsing
	cleanForParsing := stripANSI(msg)

	// Parse level from the log message
	if strings.Contains(strings.ToUpper(cleanForParsing), "DEBU") {
		level = DEBUG
		// Extract message after level indicator
		if idx := strings.Index(cleanForParsing, "DEBU"); idx >= 0 {
			cleanMsg = strings.TrimSpace(cleanForParsing[idx+4:])
		}
	} else if strings.Contains(strings.ToUpper(cleanForParsing), "INFO") {
		level = INFO
		if idx := strings.Index(cleanForParsing, "INFO"); idx >= 0 {
			cleanMsg = strings.TrimSpace(cleanForParsing[idx+4:])
		}
	} else if strings.Contains(strings.ToUpper(cleanForParsing), "WARN") {
		level = WARN
		if idx := strings.Index(cleanForParsing, "WARN"); idx >= 0 {
			cleanMsg = strings.TrimSpace(cleanForParsing[idx+4:])
		}
	} else if strings.Contains(strings.ToUpper(cleanForParsing), "ERRO") {
		level = ERROR
		if idx := strings.Index(cleanForParsing, "ERRO"); idx >= 0 {
			cleanMsg = strings.TrimSpace(cleanForParsing[idx+4:])
		}
	}

	entry := LogEntry{
		Timestamp: time.Now(),
		Level:     level,
		Message:   cleanMsg,
	}

	lc.mu.Lock()
	lc.entries = append(lc.entries, entry)
	// Keep only the last maxEntries
	if len(lc.entries) > lc.maxEntries {
		lc.entries = lc.entries[len(lc.entries)-lc.maxEntries:]
	}
	lc.mu.Unlock()

	// Non-blocking send to update channel
	select {
	case lc.updateChan <- entry:
	default:
		// Channel full, skip update
	}

	// Don't write to stdout/stderr - the TUI will display logs through its own UI
	// Writing here would interfere with the TUI rendering

	return len(p), nil
}

// stripANSI removes ANSI color codes from a string
func stripANSI(s string) string {
	var result strings.Builder
	inEscape := false
	for _, r := range s {
		if r == '\x1b' {
			inEscape = true
			continue
		}
		if inEscape {
			if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
				inEscape = false
			}
			continue
		}
		result.WriteRune(r)
	}
	return result.String()
}

// GetEntries returns all current log entries
func (lc *LogCapture) GetEntries() []LogEntry {
	if lc == nil {
		return []LogEntry{}
	}

	lc.mu.RLock()
	defer lc.mu.RUnlock()

	// Return a copy
	entries := make([]LogEntry, len(lc.entries))
	copy(entries, lc.entries)
	return entries
}

// GetRecentEntries returns the last n log entries
func (lc *LogCapture) GetRecentEntries(n int) []LogEntry {
	if lc == nil || n <= 0 {
		return []LogEntry{}
	}

	lc.mu.RLock()
	defer lc.mu.RUnlock()

	start := len(lc.entries) - n
	if start < 0 {
		start = 0
	}

	// Return a copy
	entries := make([]LogEntry, len(lc.entries)-start)
	copy(entries, lc.entries[start:])
	return entries
}

// UpdateChan returns the channel that receives log updates
func (lc *LogCapture) UpdateChan() <-chan LogEntry {
	if lc == nil {
		return nil
	}
	return lc.updateChan
}

// WriteToFile writes all buffered log entries to a file
// This is called when an error occurs
func (lc *LogCapture) WriteToFile() error {
	if lc == nil {
		return fmt.Errorf("logger not initialized")
	}

	lc.mu.Lock()
	defer lc.mu.Unlock()

	if lc.errorOccurred {
		// Already wrote to file
		return nil
	}

	// Generate filename with timestamp
	timestamp := time.Now().Format("2006-01-02-150405")
	filename := fmt.Sprintf("cctp-error-%s.log", timestamp)

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create log file: %w", err)
	}
	defer file.Close()

	// Write header
	fmt.Fprintf(file, "CCTP CLI Error Log\n")
	fmt.Fprintf(file, "Generated: %s\n", time.Now().Format(time.RFC3339))
	fmt.Fprintf(file, "===========================================\n\n")

	// Write all entries with levels
	for _, entry := range lc.entries {
		fmt.Fprintf(file, "[%s] [%s] %s\n",
			entry.Timestamp.Format("15:04:05"),
			entry.Level.String(),
			entry.Message)
	}

	lc.errorOccurred = true
	lc.logFilePath = filename

	return nil
}

// GetLogFilePath returns the path to the log file if it was created
func (lc *LogCapture) GetLogFilePath() string {
	if lc == nil {
		return ""
	}

	lc.mu.RLock()
	defer lc.mu.RUnlock()
	return lc.logFilePath
}

// Helper functions for convenience

// GetLogger returns the default log capture instance
func GetLogger() *LogCapture {
	return defaultCapture
}

// WriteToFile writes logs to a file (called on error)
func WriteToFile() error {
	if defaultCapture == nil {
		return fmt.Errorf("logger not initialized")
	}
	return defaultCapture.WriteToFile()
}

// GetRecentEntries returns recent entries from the default logger
func GetRecentEntries(n int) []LogEntry {
	if defaultCapture == nil {
		return []LogEntry{}
	}
	return defaultCapture.GetRecentEntries(n)
}
