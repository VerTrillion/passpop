package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"

	"github.com/VerTrillion/passpop/internal/auth"
	"github.com/VerTrillion/passpop/internal/config"
	"github.com/VerTrillion/passpop/internal/crypto"
)

var secureMode bool

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolVarP(&secureMode, "secure", "s", false, "Enable secure mode with master password")
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Passpop with a new encryption key",
	RunE: func(cmd *cobra.Command, args []string) error {
		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("cannot get user home directory: %w", err)
		}

		passpopDir := filepath.Join(home, ".passpop")
		credPath := filepath.Join(passpopDir, "credentials.yml")
		configPath := filepath.Join(passpopDir, "config.yml")

		if _, err := os.Stat(credPath); err == nil {
			fmt.Println("⚠️  Passpop has already been initialized.")
			fmt.Print("Reinitializing will delete all stored credentials. Proceed? (y/N): ")
			var confirm string
			fmt.Scanln(&confirm)
			if confirm != "y" && confirm != "Y" {
				fmt.Println("🚫 Operation cancelled.")
				return nil
			}

			os.Remove(credPath)
			os.Remove(configPath)

			if err := config.RemoveKeyFromZshrc(); err != nil {
				return fmt.Errorf("failed to clean up .zshrc: %w", err)
			}
			fmt.Println("🧹 Cleaned old configuration.")
		}

		if err := os.MkdirAll(passpopDir, 0700); err != nil {
			return fmt.Errorf("cannot create passpop directory: %w", err)
		}

		key, err := crypto.GenerateKey()
		if err != nil {
			return fmt.Errorf("failed to generate key: %w", err)
		}

		isPasspopKeyExist, err := config.AppendKeyToZshrc(key)
		if err != nil {
			return fmt.Errorf("failed to write key to .zshrc: %w", err)
		}
		if !isPasspopKeyExist {
			fmt.Println("✅ Passpop key has been and saved to ~/.zshrc")
			fmt.Println("👉 Please run `source ~/.zshrc` or restart your terminal to use the key")
		}

		var appCfg *config.AppConfig
		if secureMode {
			password, err := auth.PromptMasterPassword()
			if err != nil {
				return err
			}

			hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err != nil {
				return fmt.Errorf("failed to hash password: %w", err)
			}

			appCfg = &config.AppConfig{
				Version: 1,
				Mode:    "secure",
				Hash:    string(hash),
			}
			fmt.Println("🔐 Master password set and stored securely")
		} else {
			appCfg = &config.AppConfig{
				Version: 1,
				Mode:    "basic",
			}
		}

		if err := config.SaveConfig(appCfg); err != nil {
			return fmt.Errorf("failed to write config.yml: %w", err)
		}

		fmt.Println("🚀 Passpop initialized successfully!")
		return nil
	},
}
