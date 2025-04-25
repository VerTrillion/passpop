package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"

	"github.com/VerTrillion/passpop/internal/auth"
	"github.com/VerTrillion/passpop/internal/config"
	"github.com/VerTrillion/passpop/internal/crypto"
	"github.com/VerTrillion/passpop/internal/store"
)

var getCmd = &cobra.Command{
	Use:   "get [key]",
	Short: "Retrieve a password and copy it to the clipboard",
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

		passKey := os.Getenv("PASSPOP_KEY")
		if passKey == "" {
			return errors.New("PASSPOP_KEY not found in environment")
		}

		credMap, err := store.LoadCredentials()
		if err != nil {
			return err
		}

		enc, ok := credMap[key]
		if !ok {
			return fmt.Errorf("❌ key '%s' not found", key)
		}

		plain, err := crypto.Decrypt(enc, passKey)
		if err != nil {
			return fmt.Errorf("decryption failed: %w", err)
		}

		if err := clipboard.WriteAll(plain); err != nil {
			return fmt.Errorf("failed to copy to clipboard: %w", err)
		}

		fmt.Printf("🔑 Password for '%s' copied to clipboard!\n", key)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
