package main

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Long:  `Print the version number of CCTP CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("CCTP CLI version", "version", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

