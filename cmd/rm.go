package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/VerTrillion/passpop/internal/auth"
	"github.com/VerTrillion/passpop/internal/config"
	"github.com/VerTrillion/passpop/internal/store"
)

var rmCmd = &cobra.Command{
	Use:   "rm [key]",
	Short: "Remove a stored credential",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		key := args[0]

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

		if _, found := credMap[key]; !found {
			fmt.Printf("❌ Key '%s' not found.\n", key)
			return nil
		}

		delete(credMap, key)

		if err := store.SaveCredentials(credMap); err != nil {
			return fmt.Errorf("failed to save credentials: %w", err)
		}

		fmt.Printf("🗑️  Key '%s' has been removed.\n", key)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
