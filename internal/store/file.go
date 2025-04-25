package store

import (
	"errors"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const (
	credentialsFileName = "credentials.yml"
)

type CredentialsFile struct {
	Version     int               `yaml:"version"`
	Credentials map[string]string `yaml:"credentials"`
}

func GetCredentialFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".passpop", credentialsFileName), nil
}

func LoadCredentials() (map[string]string, error) {
	path, err := GetCredentialFilePath()
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return make(map[string]string), nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var f CredentialsFile
	if err := yaml.Unmarshal(data, &f); err != nil {
		return nil, err
	}

	return f.Credentials, nil
}

func SaveCredentials(creds map[string]string) error {
	path, err := GetCredentialFilePath()
	if err != nil {
		return err
	}

	f := CredentialsFile{
		Version:     1,
		Credentials: creds,
	}

	data, err := yaml.Marshal(f)
	if err != nil {
		return err
	}

	if err := os.WriteFile(path, data, 0600); err != nil {
		return err
	}

	return nil
}
