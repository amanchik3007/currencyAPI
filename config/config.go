package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Database struct {
		ConnectionString string `json:"connectionString"`
	} `json:"database"`
	Port string `json:"port"`
}

func LoadConfiguration(file string) (Config, error) {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		return Config{}, err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
