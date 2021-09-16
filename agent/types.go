package main

import (
	"encoding/hex"
	"fmt"

	"./comm"
	"./cryptoutil/dh"
)

func createKeyExchange(config *Config) (*dh.KeyExchange, error) {
	keyExchange, err := dh.NewKeyExchange(config.PrivateKey)
	if err != nil {
		return nil, err
	}

	clientID := dh.GetClientID(keyExchange.GetPrivateKey())
	fmt.Printf("Client ID: %s\n", hex.EncodeToString(clientID))
	return keyExchange, nil
}

func CreateHttpClient(config *Config) (comm.Client, error) {
	keyExchange, err := createKeyExchange(config)
	if err != nil {
		return nil, err
	}

	httpClient := comm.NewHttpClient(fmt.Sprintf("http://%s:%d", config.Host, config.Port))
	encHttpClient := comm.NewEncryptedClient(httpClient, config.SymKey)
	dhClient := comm.NewDHClient(keyExchange, encHttpClient)

	return dhClient, nil
}

func CreateUnencHttpClient(config *Config) (comm.Client, error) {
	keyExchange, err := createKeyExchange(config)
	if err != nil {
		return nil, err
	}

	httpClient := comm.NewHttpClient(fmt.Sprintf("http://%s:%d", config.Host, config.Port))
	clientID := dh.GetClientID(keyExchange.GetPublicKey())
	dummyAuthClient := comm.NewDummyAuthClient(clientID, httpClient)

	return dummyAuthClient, nil
}

func CreateClient(config *Config) (comm.Client, error) {
	switch config.Type {
	case "http":
		return CreateHttpClient(config)
	case "unenc-http":
		return CreateUnencHttpClient(config)
	default:
		return nil, fmt.Errorf("unknown client type")
	}
}
