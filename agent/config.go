package main

import (
	_ "embed"
	"encoding/json"
	"os"
)

//go:embed config_placeholder.bin
var configBytes []byte

type Config struct {
	PrivateKey           []byte `json:"PrivateKey"`
	SymKey               []byte `json:"SymKey"`
	TargetUrl            string `json:"TargetUrl"`
	ConnectionIntervalMs int    `json:"ConnectionIntervalMs"`
}

func LoadConfig(file string) (*Config, error) {
	var config Config
	jsonFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(jsonFile).Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func LoadConfigBytes(configBytes []byte) (*Config, error) {
	var config Config
	err := json.Unmarshal(configBytes, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func LoadEmbeddedConfig() (*Config, error) {
	return LoadConfigBytes(configBytes)
}
