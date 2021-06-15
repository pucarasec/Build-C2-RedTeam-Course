package main

import (
	"log"
	"net/http"

	"../crypto/dh"
	"./comm/layer0"
	"./comm/layer1"
	"./models"

	"./admin"
)

func main() {
	key := []byte("some random key!")
	db := models.GetDB()
	appHandler := NewAppHandler(db)
	keyExchange, _ := dh.NewKeyExchange(nil)
	dhHandler := layer1.NewDHHandler(layer1.NewBasicKeyRespository(), keyExchange, appHandler)
	encryptedHandler := layer0.NewEncryptedHandler(key, dhHandler)
	handler := layer0.NewHTTPHandler(encryptedHandler)
	http.Handle("/", handler)

	adminHandler := admin.NewAdminHandler("/admin", db)
	http.Handle("/admin/", adminHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
