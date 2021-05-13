package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

func getMessages(body []byte) [][]byte {
	re := regexp.MustCompile(`(<!--)([A-Za-z0-9/+=]*|=[^=]|={3,})(-->)`)
	matches := re.FindAll(body, -1)
	var messages [][]byte
	for _, match := range matches {
		encoded := string(match[4 : len(match)-3])
		data, err := base64.StdEncoding.DecodeString(encoded)
		if err == nil {
			messages = append(messages, data)
		}
	}
	return messages
}

func decryptMessage(message []byte, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	iv := message[:aes.BlockSize]
	data := message[aes.BlockSize:]
	stream := cipher.NewOFB(block, iv)
	stream.XORKeyStream(data, data)
	return data
}

func validateMessage(message []byte, key []byte) (bool, []byte) {
	mac := hmac.New(sha256.New, key)
	mac.Write(message[mac.Size():])
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(message[:mac.Size()], expectedMAC), message[mac.Size():]
}

func main() {

	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		fmt.Println("Error: %w", err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body: %w", err)
		return
	}

	key := []byte("some random key!")

	messages := getMessages(body)
	fmt.Printf("Got %d messages.\n", len(messages))
	for _, signedMessage := range messages {
		valid, message := validateMessage(signedMessage, key)
		if valid {
			fmt.Printf("Message: %s\n", decryptMessage(message, key))
		} else {
			fmt.Printf("Invalid message.\n")
		}
	}
	fmt.Printf("%s", body)
}
