package comm

type Handler interface {
	HandleMsg(msg []byte) ([]byte, error)
}

type AuthenticatedHandler interface {
	HandleAuthenticatedMsg(clientID string, msg []byte) ([]byte, error)
}
