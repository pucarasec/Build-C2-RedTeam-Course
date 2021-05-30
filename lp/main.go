package main

import (
	"fmt"
	"log"
	"net/http"

	"../crypto/dh"
	"./comm/layer0"
	"./comm/layer1"
)

type AppHandler struct{}

func (h *AppHandler) HandleAuthenticatedMsg(clientID []byte, msg []byte) ([]byte, error) {
	response := fmt.Sprintf("Your message was: %s", msg)
	return []byte(response), nil
}

func main() {
	key := []byte("some random key!")
	appHandler := &AppHandler{}
	dhHandler := layer1.NewDHHandler(layer1.NewBasicKeyRespository(), dh.NewKeyExchange(), appHandler)
	encryptedHandler := layer0.NewEncryptedHandler(key, dhHandler)
	handler := layer0.NewHTTPHandler(encryptedHandler)
	http.Handle("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
