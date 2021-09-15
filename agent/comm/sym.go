package comm

import (
	"../cryptoutil/sym"
)

type EncryptedClient struct {
	key    []byte
	client Client
}

func NewEncryptedClient(client Client, key []byte) *EncryptedClient {
	return &EncryptedClient{
		key:    key,
		client: client,
	}
}

func (client *EncryptedClient) SendMsg(outgoingMsg []byte) ([]byte, error) {
	outgoingMsg, err := sym.EncryptThenSignMessage([]byte(outgoingMsg), client.key)
	if err != nil {
		return nil, err
	}
	incomingMsg, err := client.client.SendMsg(outgoingMsg)
	if err != nil {
		return nil, err
	}
	incomingMsg, err = sym.ValidateThenDecryptMessage(incomingMsg, client.key)
	if err != nil {
		return nil, err
	}
	return incomingMsg, nil
}
