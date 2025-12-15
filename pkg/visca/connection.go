package visca

import (
	"errors"
	"fmt"
	"io"
	"net"
)

type Camera struct {
	IP         string
	Port       uint16
	connection *net.TCPConn
}

func (c *Camera) Connect() error {
	if tcpAddress, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", c.IP, c.Port)); err != nil {
		return err
	} else if c.connection, err = net.DialTCP("tcp", nil, tcpAddress); err != nil {
		return err
	} else if err := c.connection.SetKeepAlive(true); err != nil {
		return err
	} else {
		return nil
	}
}

func (c *Camera) Disconnect() error {
	return c.connection.Close()
}

const BUFFER_LEN = 32

func (c *Camera) sendViscaBytes(b []byte) ([]byte, error) {
	if _, err := c.connection.Write(b); err != nil {
		return nil, err
	} else {
		// read all data
		responseBuffer := make([]byte, BUFFER_LEN)
		response := []byte{}

		for {
			if readBytes, err := c.connection.Read(responseBuffer); err != nil {
				if errors.Is(err, io.EOF) {
					break
				} else {
					return nil, err
				}
			} else if readBytes < BUFFER_LEN {
				response = append(response, responseBuffer[:readBytes]...)
				break
			} else {
				response = append(response, responseBuffer...)
			}
		}

		// TODO: parse returned message for response message

		return response, nil
	}
}
