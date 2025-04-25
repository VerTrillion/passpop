package cmd

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"

	"github.com/VerTrillion/passpop/internal/auth"
	"github.com/VerTrillion/passpop/internal/config"
	"github.com/VerTrillion/passpop/internal/store"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all stored keys",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadConfig()
		if err != nil {
			return fmt.Errorf("cannot load config: %w", err)
		}

		if cfg.Mode == "secure" {
			if err := auth.VerifyMasterPassword(cfg.Hash); err != nil {
				return err
			}
		}

		credMap, err := store.LoadCredentials()
		if err != nil {
			return fmt.Errorf("cannot load credentials: %w", err)
		}

		if len(credMap) == 0 {
			fmt.Println("📭 No credentials found.")
			return nil
		}

		fmt.Println("🔑 Stored keys:")
		keys := make([]string, 0, len(credMap))
		for k := range credMap {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			fmt.Printf("- %s\n", k)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
