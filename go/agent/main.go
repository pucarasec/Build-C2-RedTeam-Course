package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"time"

	"../crypto/dh"
	"./comm"
)

func main() {
	config, err := LoadConfig("config.json")
	if err != nil {
		fmt.Printf("Error reading config file: %s\n", err)
		return
	}

	httpClient := comm.NewHttpClient(config.TargetUrl)
	encHttpClient := comm.NewEncryptedClient(httpClient, config.SymKey)

	keyExchange, err := dh.NewKeyExchange(config.PrivateKey)
	if err != nil {
		fmt.Printf("Error creating key exchange: %s\n", err)
		return
	}

	client := comm.NewDHClient(keyExchange, encHttpClient)

	agent := NewAgent(client)

	encodedPrivateKey := base64.StdEncoding.EncodeToString(keyExchange.GetPrivateKey())
	fmt.Printf("Private Key: %s\n", encodedPrivateKey)

	encodedPublicKey := base64.StdEncoding.EncodeToString(keyExchange.GetPublicKey())
	fmt.Printf("Public Key: %s\n", encodedPublicKey)

	selfSharedKey, _ := keyExchange.GetSharedKey(keyExchange.GetPublicKey())
	encodedSelfSharedKey := base64.StdEncoding.EncodeToString(selfSharedKey)
	fmt.Printf("Self Shared Key: %s\n", encodedSelfSharedKey)
	fmt.Printf("Client ID: %s\n", hex.EncodeToString(client.GetClientID()))

	for {
		err := agent.heartbeat()
		if err != nil {
			fmt.Printf("Error: %e", err)
		}
		time.Sleep(time.Duration(config.ConnectionIntervalMs) * time.Millisecond)
	}

}
