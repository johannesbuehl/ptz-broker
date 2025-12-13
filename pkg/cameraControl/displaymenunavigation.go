package cameraControl

import "net"

func OpenMenu(connection *net.TCPConn) error {
	command := []byte{
		0x81, 0x01, 0x04, 0x3F, 0x02, 0x5F, 0xFF,
	}
	if _, err := connection.Write([]byte(command)); err != nil {
		return err
	} else {
		return nil
	}
}

func Enter(connection *net.TCPConn) error {
	command := []byte{
		0x81, 0x01, 0x06, 0x06, 0x05, 0xFF,
	}

	if _, err := connection.Write([]byte(command)); err != nil {
		return err
	} else {
		return nil
	}
}
