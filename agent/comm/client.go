package comm

type Client interface {
	SendMsg(msg []byte) ([]byte, error)
}
