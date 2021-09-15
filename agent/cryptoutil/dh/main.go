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

func NewKeyExchange(privateKeyBytes []byte) (*KeyExchange, error) {
	var privateKey *ecdsa.PrivateKey
	var err error
	if privateKeyBytes == nil {
		privateKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	} else {
		privateKey, err = x509.ParseECPrivateKey(privateKeyBytes)
	}

	if err != nil {
		return nil, err
	}

	return &KeyExchange{
		PrivateKey: privateKey,
	}, nil
}

func (ke *KeyExchange) GetSharedKey(publicKeyBytes []byte) ([]byte, error) {
	publicKey, err := x509.ParsePKIXPublicKey(publicKeyBytes)
	if err != nil {
		return nil, err
	}
	ecdsaPublicKey := publicKey.(*ecdsa.PublicKey)
	x, _ := ecdsaPublicKey.Curve.ScalarMult(ecdsaPublicKey.X, ecdsaPublicKey.Y, ke.PrivateKey.D.Bytes())
	return x.Bytes(), nil
}

func (ke *KeyExchange) GetPublicKey() []byte {
	bytes, err := x509.MarshalPKIXPublicKey(&ke.PrivateKey.PublicKey)
	if err != nil {
		fmt.Println(err)
	}
	return bytes
}

func (ke *KeyExchange) GetPrivateKey() []byte {
	keyBytes, _ := x509.MarshalECPrivateKey(ke.PrivateKey)
	return keyBytes
}

func GetClientID(publicKey []byte) []byte {
	h := sha256.New()
	h.Write(publicKey)
	return h.Sum(nil)[:16]
}
