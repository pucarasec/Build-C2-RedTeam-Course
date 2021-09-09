package main

import (
	"encoding/json"
	"os"
)

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
