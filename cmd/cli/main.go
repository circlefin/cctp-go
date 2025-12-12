package main

import (
	"os"

	"github.com/charmbracelet/log"
)

func main() {
	if err := Execute(); err != nil {
		log.Error("Command execution failed", "error", err)
		os.Exit(1)
	}
}
