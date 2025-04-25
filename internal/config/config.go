package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	Version int    `yaml:"version"`
	Mode    string `yaml:"mode"`           // "basic" or "secure"
	Hash    string `yaml:"hash,omitempty"` // only if secure mode
}

func ConfigFilePath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".passpop", "config.yml")
}

func LoadConfig() (*AppConfig, error) {
	path := ConfigFilePath()
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg AppConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func SaveConfig(cfg *AppConfig) error {
	path := ConfigFilePath()
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0600)
}

func AppendKeyToZshrc(encodedKey string) (bool, error) {
	home, err := os.UserHomeDir()
	var isPasspopKeyExist bool

	if err != nil {
		return isPasspopKeyExist, err
	}

	zshrcPath := filepath.Join(home, ".zshrc")
	exportLine := fmt.Sprintf("export PASSPOP_KEY=\"%s\"", encodedKey)

	if content, err := os.ReadFile(zshrcPath); err == nil {
		if strings.Contains(string(content), "PASSPOP_KEY=") {
			isPasspopKeyExist = true
			fmt.Println("⚠️  PASSPOP_KEY already set in .zshrc, skipping append")
			return isPasspopKeyExist, nil
		}
	}

	f, err := os.OpenFile(zshrcPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return isPasspopKeyExist, err
	}
	defer f.Close()

	if _, err := f.WriteString("\n" + exportLine + "\n"); err != nil {
		return isPasspopKeyExist, err
	}

	return isPasspopKeyExist, nil
}

func RemoveKeyFromZshrc() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	zshrcPath := filepath.Join(home, ".zshrc")

	data, err := os.ReadFile(zshrcPath)
	if err != nil {
		return err
	}

	lines := []string{}
	for _, line := range strings.Split(string(data), "\n") {
		if !strings.Contains(line, "PASSPOP_KEY=") {
			lines = append(lines, line)
		}
	}

	newContent := strings.Join(lines, "\n")
	return os.WriteFile(zshrcPath, []byte(newContent), 0644)
}
