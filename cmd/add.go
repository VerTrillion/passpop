package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/VerTrillion/passpop/internal/auth"
	"github.com/VerTrillion/passpop/internal/config"
	"github.com/VerTrillion/passpop/internal/crypto"
	"github.com/VerTrillion/passpop/internal/store"
	"github.com/spf13/cobra"
)

var (
	addKey      string
	addPassword string
)

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&addKey, "key", "k", "", "Key for the password (e.g., gmail)")
	addCmd.Flags().StringVarP(&addPassword, "password", "p", "", "Password to encrypt and store")
	addCmd.MarkFlagRequired("key")
	addCmd.MarkFlagRequired("password")
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add or update a credential",
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

		secret := os.Getenv("PASSPOP_KEY")
		if secret == "" {
			return errors.New("PASSPOP_KEY not set in environment (please source your .zshrc)")
		}

		enc, err := crypto.Encrypt(addPassword, secret)
		if err != nil {
			return fmt.Errorf("encryption failed: %w", err)
		}

		credMap, err := store.LoadCredentials()
		if err != nil {
			return fmt.Errorf("cannot load credentials: %w", err)
		}

		credMap[addKey] = enc

		if err := store.SaveCredentials(credMap); err != nil {
			return fmt.Errorf("cannot save credentials: %w", err)
		}

		fmt.Printf("✅ Credential for '%s' saved successfully.\n", addKey)
		return nil
	},
}
