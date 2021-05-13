package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
)

func encryptMessage(message []byte, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	encrypted := make([]byte, aes.BlockSize+len(message))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	data := encrypted[aes.BlockSize:]
	stream := cipher.NewOFB(block, iv)
	stream.XORKeyStream(data, message)
	return encrypted
}

func signMessage(message []byte, key []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	signed := make([]byte, len(message)+mac.Size())
	copy(signed[:mac.Size()], mac.Sum((nil)))
	copy(signed[mac.Size():], message)
	return signed
}

func handler(w http.ResponseWriter, r *http.Request) {
	key := []byte("some random key!")

	message := []byte("Hi Agent, I'm a Listening Post")
	encoded := base64.StdEncoding.EncodeToString(signMessage(encryptMessage(message, key), key))

	fmt.Fprintf(w, "<html><body>Hi there, this is a totally innocent web page. <!--%s--></body></html>", encoded)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
