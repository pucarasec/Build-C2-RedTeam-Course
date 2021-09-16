package comm

import (
	"bufio"
	"fmt"
	"net"
)

type UdpClient struct {
	conn net.Conn
}

func NewUdpClient(targetHost string, targetPort int) (*UdpClient, error) {
	conn, err := net.Dial("udp", fmt.Sprintf("%s:%d", targetHost, targetPort))
	if err != nil {
		return nil, err
	}
	return &UdpClient{
		conn: conn,
	}, nil
}

const NUMERO_TOTALMENTE_ARBITRARIO = 2048

func (client *UdpClient) SendMsg(outgoingMsg []byte) ([]byte, error) {
	bytes_written, err := client.conn.Write(outgoingMsg)
	if err != nil {
		return nil, err
	}
	if bytes_written < len(outgoingMsg) {
		return nil, fmt.Errorf("error writing to UDP socket")
	}

	buffer := make([]byte, NUMERO_TOTALMENTE_ARBITRARIO)
	bytes_read, err := bufio.NewReader(client.conn).Read(buffer)
	if err != nil {
		return nil, err
	}
	return buffer[:bytes_read], nil
}
