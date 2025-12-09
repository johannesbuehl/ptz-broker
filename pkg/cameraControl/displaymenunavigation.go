package cameraControl

import "net"

func OpenMenu(connection *net.TCPConn) error {
	command := "81 01 04 3F 02 5F FF"

	if _, err := connection.Write([]byte(command)); err != nil {
		return err
	} else {
		return nil
	}
}

func Enter(connection *net.TCPConn) error {
	command := "81 01 06 06 05 FF"

	if _, err := connection.Write([]byte(command)); err != nil {
		return err
	} else {
		return nil
	}
}
