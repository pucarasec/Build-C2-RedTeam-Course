package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"../crypto/dh"
	"./comm/client"
	"./comm/layer0"
	"./comm/layer1"
)

func readFromTerminal(client client.Client) {
	reader := bufio.NewReader(os.Stdin)

	for {
		text, _ := reader.ReadString('\n')
		msg := fmt.Sprintf("!%s", text)
		response, err := client.SendMsg([]byte(msg))

		if err != nil {
			fmt.Printf("Error: %s\n", err)
		} else {
			fmt.Printf("%s\n", response)
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	targetUrl := "http://localhost:8080"
	key := []byte("some random key!")

	httpClient := layer0.NewHttpClient(targetUrl)
	encHttpClient := layer0.NewEncryptedClient(httpClient, key)
	client := layer1.NewDHClient(dh.NewKeyExchange(), encHttpClient)

	go readFromTerminal(client)

	for {
		response, err := client.SendMsg([]byte("PING!"))
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		} else if len(response) > 0 {
			fmt.Printf("%s\n", response)
		}
		time.Sleep(1 * time.Second)
	}

}
