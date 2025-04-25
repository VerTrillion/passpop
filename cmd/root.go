package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "1.0.0"

var rootCmd = &cobra.Command{
	Use:   "passpop",
	Short: "Passpop is a simple CLI password manager for macOS",
	Long: `Passpop is a simple, secure, and lightweight CLI password manager
for macOS that uses AES encryption and YAML storage.`,
	Version: version,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
