package crypto

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateKey() (string, error) {
	key := make([]byte, 32) // 256-bit key
	if _, err := rand.Read(key); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(key), nil
}
