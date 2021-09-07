package comm

import (
	"../../crypto/sym"
)

type EncryptedHandler struct {
	key        []byte
	subhandler Handler
}

func NewEncryptedHandler(key []byte, subhandler Handler) *EncryptedHandler {
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
