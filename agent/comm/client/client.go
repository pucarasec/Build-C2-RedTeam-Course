package client

type Client interface {
	SendMsg(msg []byte) ([]byte, error)
}
