package config // this is the config page

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	HomeDir, err := os.UserHomeDir()
	if err != nil {
		return Config{}, fmt.Errorf("Error no HomeDir: %w", err)
	}

	path := filepath.Join(HomeDir, ".gatorconfig.json")
	file, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
