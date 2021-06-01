package layer1

import (
	"encoding/hex"

	"../../../crypto/dh"
	"../../../crypto/sym"
	"../../../protocol"
	"../handler"
	"google.golang.org/protobuf/proto"
)

type DHHandler struct {
	keyRepository KeyRepository
	keyExchange   *dh.KeyExchange
	subhandler    handler.AuthenticatedHandler
}

func NewDHHandler(keyRepository KeyRepository, keyExchange *dh.KeyExchange, subhandler handler.AuthenticatedHandler) *DHHandler {
	return &DHHandler{
		keyRepository: keyRepository,
		keyExchange:   keyExchange,
		subhandler:    subhandler,
	}
}

func (h *DHHandler) HandleMsg(msg []byte) ([]byte, error) {
	var baseMsg protocol.BaseMsg
	var responseBaseMsg *protocol.BaseMsg
	err := proto.Unmarshal(msg, &baseMsg)
	if err != nil {
		return nil, err
	}

	if handshakeMsg := baseMsg.GetHandshakeMsg(); handshakeMsg != nil {
		responseBaseMsg, err = h.handleHandshakeMsg(handshakeMsg)
	} else if clientMsg := baseMsg.GetClientMsg(); clientMsg != nil {
		responseBaseMsg, err = h.handleClientMsg(clientMsg)
	}

	if err != nil {
		return nil, err
	}

	response, err := proto.Marshal(responseBaseMsg)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (h *DHHandler) handleHandshakeMsg(msg *protocol.HandshakeMsg) (*protocol.BaseMsg, error) {
	clientID := hex.EncodeToString(dh.GetClientID(msg.PublicKey))
	err := h.keyRepository.Set(clientID, msg.PublicKey)
	if err != nil {
		return nil, err
	}

	return &protocol.BaseMsg{
		MsgType: &protocol.BaseMsg_HandshakeMsg{
			HandshakeMsg: &protocol.HandshakeMsg{
				PublicKey: h.keyExchange.GetPublicKey(),
			},
		},
	}, nil
}

func (h *DHHandler) handleClientMsg(msg *protocol.ClientMsg) (*protocol.BaseMsg, error) {
	clientID := hex.EncodeToString(msg.ClientID)
	publicKey, err := h.keyRepository.Get(clientID)
	if err != nil {
		return nil, err
	}

	if publicKey == nil {
		return &protocol.BaseMsg{
			MsgType: &protocol.BaseMsg_ErrorMsg{
				ErrorMsg: &protocol.ErrorMsg{
					Type: protocol.ErrorType_HANDSHAKE_EXPIRED,
				},
			},
		}, nil
	}

	sharedKey, err := h.keyExchange.GetSharedKey(publicKey)
	if err != nil {
		return nil, err
	}

	decryptedPayload, err := sym.ValidateThenDecryptMessage(msg.Payload, sharedKey)
	if err != nil {
		return nil, err
	}

	response, err := h.subhandler.HandleAuthenticatedMsg(clientID, decryptedPayload)
	if err != nil {
		return nil, err
	}

	encryptedResponse, err := sym.EncryptThenSignMessage(response, sharedKey)
	if err != nil {
		return nil, err
	}

	return &protocol.BaseMsg{
		MsgType: &protocol.BaseMsg_ServerMsg{
			ServerMsg: &protocol.ServerMsg{
				Payload: encryptedResponse,
			},
		},
	}, nil
}
