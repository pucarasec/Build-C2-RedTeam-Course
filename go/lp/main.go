package main

import (
	"log"
	"net/http"

	"../crypto/dh"
	"./comm"
	"./models"

	"./admin"
)

func main() {
	key := []byte("some random key!")
	db := models.GetDB()
	appHandler := NewAppHandler(db)
	keyExchange, _ := dh.NewKeyExchange(nil)
	dhHandler := comm.NewDHHandler(comm.NewBasicKeyRespository(), keyExchange, appHandler)
	encryptedHandler := comm.NewEncryptedHandler(key, dhHandler)
	handler := comm.NewHTTPHandler(encryptedHandler)
	http.Handle("/", handler)

	adminHandler := admin.NewAdminHandler("/admin", db)
	http.Handle("/admin/", adminHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
