package layer0

import (
	"../../../crypto/sym"
	"../handler"
)

type EncryptedHandler struct {
	key        []byte
	subhandler handler.Handler
}

func NewEncryptedHandler(key []byte, subhandler handler.Handler) *EncryptedHandler {
	return &EncryptedHandler{
		key:        key,
		subhandler: subhandler,
	}
}

func (h *EncryptedHandler) HandleMsg(msg []byte) ([]byte, error) {
	decryptedMsg, err := sym.ValidateThenDecryptMessage(msg, h.key)
	if err != nil {
		return nil, err
	}
	response, err := h.subhandler.HandleMsg(decryptedMsg)
	if err != nil {
		return nil, err
	}
	encryptedResponse, err := sym.EncryptThenSignMessage(response, h.key)
	if err != nil {
		return nil, err
	}
	return encryptedResponse, nil
}
