package config

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/RevoltSecurities/Crlfix/crlfix/modules/logger"
	"gopkg.in/yaml.v3"
)

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func ExistDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}
	return nil
}

func ExistFile(file string) error {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		content := map[string][]string{
			"slack": {},
		}
		data, err := yaml.Marshal(content)
		if err != nil {
			return fmt.Errorf("failed to marshal default config: %w", err)
		}
		if err := os.WriteFile(file, data, 0644); err != nil {
			return fmt.Errorf("failed to create config file %s: %w", file, err)
		}
	}
	return nil
}

func GetConfig() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		logger.Logger("Unable to get the Config Directory!", "warn")
		return "", err
	}

	crlfiXDir := filepath.Join(configDir, "crlfix")
	if err := ExistDir(crlfiXDir); err != nil {
		return "", err
	}

	configFile := filepath.Join(crlfiXDir, "crlfiX.yaml")
	if err := ExistFile(configFile); err != nil {
		return "", err
	}

	return configFile, nil
}

type Config struct {
	Slack []string `yaml:"slack"`
}

func SetConfig(configPath string) (*Config, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config file: %w", err)
	}

	return &config, nil
}

func (c *Config) GetRandomKey() (string, error) {
	if len(c.Slack) == 0 {
		return "", fmt.Errorf("no url available in Slack")
	}
	randomIndex := seededRand.Intn(len(c.Slack))
	return c.Slack[randomIndex], nil
}
