package main

import (
	"fmt"

	"../crypto/dh"
	"./comm/layer0"
	"./comm/layer1"
)

func main() {
	targetUrl := "http://localhost:8080"
	outgoingMessage := "My message"
	key := []byte("some random key!")

	httpClient := layer0.NewHttpClient(targetUrl)
	encHttpClient := layer0.NewEncryptedClient(httpClient, key)
	client := layer1.NewDHClient(dh.NewKeyExchange(), encHttpClient)

	response, err := client.SendMsg([]byte(outgoingMessage))
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	fmt.Printf("Response: %s\n", response)

}
