package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigPath() (string, error) {

	hom, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Home directory not found: %v", err)
	}
	return filepath.Join(hom, configFileName), nil

}

func Read() (Config, error) {

	pth, err := getConfigPath()
	if err != nil {
		return Config{}, err
	}

	dat, err := os.ReadFile(pth)
	if err != nil {
		return Config{}, fmt.Errorf("Error reading file: %v", err)
	}

	cfg := Config{}
	err = json.Unmarshal(dat, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("Error decoding data: %v", err)
	}

	return cfg, nil
}

func write(cfg Config) error {
	pth, err := getConfigPath()
	if err != nil {
		return err
	}

	dat, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("Error encoding config file: %v", err)
	}

	err = os.WriteFile(pth, dat, 0644)
	if err != nil {
		return fmt.Errorf("Error writing config file: %v", err)
	}

	return nil
}

func (c *Config) SetUser(current_user_name string) error {
	c.CurrentUserName = current_user_name
	err := write(*c)
	return err
}
