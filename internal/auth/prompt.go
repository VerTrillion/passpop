package auth

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func PromptMasterPassword() (string, error) {
	fmt.Println("🔐 Secure mode enabled. Please set a master password.")
	for {
		fmt.Print("Enter master password: ")
		pass1, err := term.ReadPassword(int(os.Stdin.Fd()))
		fmt.Println()
		if err != nil {
			return "", fmt.Errorf("failed to read password: %w", err)
		}

		fmt.Print("Confirm master password: ")
		pass2, err := term.ReadPassword(int(os.Stdin.Fd()))
		fmt.Println()
		if err != nil {
			return "", fmt.Errorf("failed to read password: %w", err)
		}

		if strings.TrimSpace(string(pass1)) != strings.TrimSpace(string(pass2)) {
			fmt.Println("❌ Passwords do not match. Please try again.")
			continue
		}

		if len(pass1) < 6 {
			fmt.Println("⚠️  Password must be at least 6 characters.")
			continue
		}

		return string(pass1), nil
	}
}
