package layer1

import (
	"fmt"

	"../../../crypto/dh"
	"../../../crypto/sym"
	protocol "../../../protocol/base"
	"../client"
	"google.golang.org/protobuf/proto"
)

type DHClient struct {
	sharedKey   []byte
	keyExchange *dh.KeyExchange
	subclient   client.Client
}

func NewDHClient(keyExchange *dh.KeyExchange, subclient client.Client) *DHClient {
	return &DHClient{
		sharedKey:   nil,
		keyExchange: keyExchange,
		subclient:   subclient,
	}
}

func (c *DHClient) sendBaseMsg(baseMsg *protocol.BaseMsg) (*protocol.BaseMsg, error) {
	var responseBaseMsg protocol.BaseMsg
	msg, err := proto.Marshal(baseMsg)
	if err != nil {
		return nil, err
	}
	response, err := c.subclient.SendMsg(msg)
	if err != nil {
		return nil, err
	}
	err = proto.Unmarshal(response, &responseBaseMsg)
	if err != nil {
		return nil, err
	}
	return &responseBaseMsg, nil
}

func (c *DHClient) NegotiateKey() error {
	msg := &protocol.BaseMsg{
		MsgType: &protocol.BaseMsg_HandshakeMsg{
			HandshakeMsg: &protocol.HandshakeMsg{
				PublicKey: c.keyExchange.GetPublicKey(),
			},
		},
	}
	response, err := c.sendBaseMsg(msg)
	if err != nil {
		return err
	}

	if handshakeMsg := response.GetHandshakeMsg(); handshakeMsg != nil {
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

	baseMsg := &protocol.BaseMsg{
		MsgType: &protocol.BaseMsg_ClientMsg{
			ClientMsg: &protocol.ClientMsg{
				ClientID: c.GetClientID(),
				Payload:  encryptedMsg,
			},
		},
	}
	responseBaseMsg, err := c.sendBaseMsg(baseMsg)
	if err != nil {
		return nil, err
	}

	if serverMsg := responseBaseMsg.GetServerMsg(); serverMsg != nil {
		return sym.ValidateThenDecryptMessage(serverMsg.Payload, c.sharedKey)
	} else if errorMsg := responseBaseMsg.GetErrorMsg(); errorMsg != nil {
		if errorMsg.Type == protocol.ErrorType_HANDSHAKE_EXPIRED {
			c.NegotiateKey()
			return c.SendMsg(msg)
		} else {
			return nil, fmt.Errorf("server returned unknown error")
		}
	} else {
		return nil, fmt.Errorf("expected ServerMsg")
	}
}
