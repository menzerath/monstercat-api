package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "monstercat",
	Short: "monstercat-api cli",
	Long:  "a commandline client for the Monstercat api",
}

// Execute is our main entry point and is called by main.go
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
