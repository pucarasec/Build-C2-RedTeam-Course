package dh

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"fmt"
)

type KeyExchange struct {
	PrivateKey *ecdsa.PrivateKey
}

func NewKeyExchange() *KeyExchange {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	return &KeyExchange{
		PrivateKey: privateKey,
	}
}

func (ke *KeyExchange) GetSharedKey(publicKeyBytes []byte) ([]byte, error) {
	publicKey, err := x509.ParsePKIXPublicKey(publicKeyBytes)
	if err != nil {
		return nil, err
	}
	ecdsaPublicKey := publicKey.(*ecdsa.PublicKey)
	x, y := ecdsaPublicKey.Curve.ScalarMult(ecdsaPublicKey.X, ecdsaPublicKey.Y, ke.PrivateKey.D.Bytes())
	h := sha256.New()
	h.Write(x.Bytes())
	h.Write(y.Bytes())
	return h.Sum(nil), nil
}

func (ke *KeyExchange) GetPublicKey() []byte {
	bytes, err := x509.MarshalPKIXPublicKey(&ke.PrivateKey.PublicKey)
	if err != nil {
		fmt.Println(err)
	}
	return bytes
}

func GetClientID(publicKey []byte) []byte {
	h := sha256.New()
	h.Write(publicKey)
	return h.Sum(nil)[:16]
}
