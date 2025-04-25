package auth

import (
	"errors"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
)

func VerifyMasterPassword(expectedHash string) error {
	fmt.Print("🔐 Enter master password: ")
	pass, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	if err != nil {
		return fmt.Errorf("failed to read password: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(expectedHash), pass); err != nil {
		return errors.New("❌ Invalid master password")
	}
	return nil
}
