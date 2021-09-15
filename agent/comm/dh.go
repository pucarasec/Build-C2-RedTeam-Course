package comm

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"../cryptoutil/dh"
	"../cryptoutil/sym"
)

type DHClient struct {
	sharedKey   []byte
	keyExchange *dh.KeyExchange
	subclient   Client
}

type HandshakeMsg struct {
	PublicKey []byte `json:"public_key"`
}

type ServerMsg struct {
	Payload []byte `json:"payload"`
}

type ClientMsg struct {
	ClientID string `json:"client_id"`
	Payload  []byte `json:"payload"`
}

type ErrorMsg struct {
	Type string `json:"type"`
}

type BaseMsg struct {
	HandshakeMsg *HandshakeMsg `json:"handshake_msg,omitempty"`
	ClientMsg    *ClientMsg    `json:"client_msg,omitempty"`
	ServerMsg    *ServerMsg    `json:"server_msg,omitempty"`
	ErrorMsg     *ErrorMsg     `json:"error_msg,omitempty"`
}

func NewDHClient(keyExchange *dh.KeyExchange, subclient Client) *DHClient {
	return &DHClient{
		sharedKey:   nil,
		keyExchange: keyExchange,
		subclient:   subclient,
	}
}

func (c *DHClient) sendBaseMsg(baseMsg *BaseMsg) (*BaseMsg, error) {
	var responseBaseMsg BaseMsg
	msg, err := json.Marshal(baseMsg)
	if err != nil {
		return nil, err
	}
	response, err := c.subclient.SendMsg(msg)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(response, &responseBaseMsg)
	if err != nil {
		return nil, err
	}
	return &responseBaseMsg, nil
}

func (c *DHClient) NegotiateKey() error {
	msg := BaseMsg{
		HandshakeMsg: &HandshakeMsg{
			PublicKey: c.keyExchange.GetPublicKey(),
		},
	}
	response, err := c.sendBaseMsg(&msg)
	if err != nil {
		return err
	}

	if handshakeMsg := response.HandshakeMsg; handshakeMsg != nil {
		c.sharedKey, err = c.keyExchange.GetSharedKey(handshakeMsg.PublicKey)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("expected HandshakeMsg")
	}

	return nil
}

func (c *DHClient) GetClientID() []byte {
	return dh.GetClientID(c.keyExchange.GetPublicKey())
}

func (c *DHClient) SendMsg(msg []byte) ([]byte, error) {
	if c.sharedKey == nil {
		err := c.NegotiateKey()
		if err != nil {
			return nil, err
		}
	}

	encryptedMsg, err := sym.EncryptThenSignMessage(msg, c.sharedKey)
	if err != nil {
		return nil, err
	}

	baseMsg := &BaseMsg{
		ClientMsg: &ClientMsg{
			ClientID: hex.EncodeToString(c.GetClientID()),
			Payload:  encryptedMsg,
		},
	}
	responseBaseMsg, err := c.sendBaseMsg(baseMsg)
	if err != nil {
		return nil, err
	}

	if serverMsg := responseBaseMsg.ServerMsg; serverMsg != nil {
		return sym.ValidateThenDecryptMessage(serverMsg.Payload, c.sharedKey)
	} else if errorMsg := responseBaseMsg.ErrorMsg; errorMsg != nil {
		if errorMsg.Type == "HANDSHAKE_EXPIRED" {
			c.NegotiateKey()
			return c.SendMsg(msg)
		} else {
			return nil, fmt.Errorf("server returned unknown error")
		}
	} else {
		return nil, fmt.Errorf("expected ServerMsg")
	}
}
