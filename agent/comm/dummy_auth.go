package comm

import (
	"encoding/hex"
	"encoding/json"
)

type DummyAuthClient struct {
	clientID  []byte
	subclient Client
}

func NewDummyAuthClient(clientID []byte, subclient Client) *DummyAuthClient {
	return &DummyAuthClient{
		clientID:  clientID,
		subclient: subclient,
	}
}

func (c *DummyAuthClient) GetClientID() []byte {
	return c.clientID
}

func (c *DummyAuthClient) SendMsg(msg []byte) ([]byte, error) {
	clientMsg := BaseMsg{
		ClientMsg: &ClientMsg{
			ClientID: hex.EncodeToString(c.clientID),
			Payload:  msg,
		},
	}

	jsonMsg, err := json.Marshal(clientMsg)
	if err != nil {
		return nil, err
	}

	response, err := c.subclient.SendMsg(jsonMsg)
	if err != nil {
		return nil, err
	}

	return response, nil
}
