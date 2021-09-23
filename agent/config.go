package main

import (
	_ "embed"
	"encoding/json"
	"os"
)

type Config struct {
	PrivateKey []byte `json:"PrivateKey"`
	SymKey     []byte `json:"SymKey"`
	Type       string `json:"Type"`
	Host       string `json:"Host"`
	Port       int    `json:"Port"`
	IntervalMs int    `json:"IntervalMs"`
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
