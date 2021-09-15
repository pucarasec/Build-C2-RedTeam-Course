package sym

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"io"
)

func DecryptMessage(message []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	iv := message[:aes.BlockSize]
	data := message[aes.BlockSize:]
	stream := cipher.NewOFB(block, iv)
	stream.XORKeyStream(data, data)
	return data, nil
}

func EncryptMessage(message []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	encrypted := make([]byte, aes.BlockSize+len(message))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	data := encrypted[aes.BlockSize:]
	stream := cipher.NewOFB(block, iv)
	stream.XORKeyStream(data, message)
	return encrypted, nil
}

func SignMessage(message []byte, key []byte) []byte {
	mac := hmac.New(sha1.New, key)
	mac.Write(message)
	signed := make([]byte, len(message)+mac.Size())
	copy(signed[:mac.Size()], mac.Sum((nil)))
	copy(signed[mac.Size():], message)
	return signed
}

func ValidateMessage(message []byte, key []byte) ([]byte, bool) {
	mac := hmac.New(sha1.New, key)
	mac.Write(message[mac.Size():])
	expectedMAC := mac.Sum(nil)
	return message[mac.Size():], hmac.Equal(message[:mac.Size()], expectedMAC)
}

func EncryptThenSignMessage(message []byte, key []byte) ([]byte, error) {
	encrypted, err := EncryptMessage(message, key)
	if err != nil {
		return nil, err
	}

	return SignMessage(encrypted, key), nil
}

func ValidateThenDecryptMessage(signedMessage []byte, key []byte) ([]byte, error) {
	message, valid := ValidateMessage(signedMessage, key)
	if !valid {
		return nil, fmt.Errorf("invalid message")
	}
	decrypted, err := DecryptMessage(message, key)
	if err != nil {
		return nil, err
	}
	return decrypted, nil
}
