package handler

type Handler interface {
	HandleMsg(msg []byte) ([]byte, error)
}

type AuthenticatedHandler interface {
	HandleAuthenticatedMsg(clientID []byte, msg []byte) ([]byte, error)
}
