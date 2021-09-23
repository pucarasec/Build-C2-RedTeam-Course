package comm

import (
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
	// Se debe utilizar la conexion instanciada en la variable 'conn'
	// para completar un buffer de tamaño NUMERO_TOTALMENTE_ARBITRARIO
	// y devolver el segmento del buffer correspondiente a la cantidad
	// de bytes enviados por el Listener
	return nil, fmt.Errorf("Not Implemented ERROR")

}
