package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

func Encrypt(plaintext, base64Key string) (string, error) {
	key, err := base64.StdEncoding.DecodeString(base64Key)
	if err != nil {
		return "", err
	}
	if len(key) != 32 {
		return "", errors.New("invalid AES key length")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := aesGCM.Seal(nil, nonce, []byte(plaintext), nil)

	final := append(nonce, ciphertext...)
	return base64.StdEncoding.EncodeToString(final), nil
}

func Decrypt(base64Input, base64Key string) (string, error) {
	key, err := base64.StdEncoding.DecodeString(base64Key)
	if err != nil {
		return "", err
	}
	if len(key) != 32 {
		return "", errors.New("invalid AES key length")
	}

	data, err := base64.StdEncoding.DecodeString(base64Input)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	if len(data) < aesGCM.NonceSize() {
		return "", errors.New("malformed encrypted data")
	}

	nonce := data[:aesGCM.NonceSize()]
	ciphertext := data[aesGCM.NonceSize():]

	plain, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plain), nil
}
