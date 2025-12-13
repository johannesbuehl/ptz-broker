package cameraControl

import (
	"net"
)

func Move(speed byte, direction string, connection *net.TCPConn) error {
	direc := []byte{
		0x81, 0x01, 0x06, 0x01, speed, speed, 0x00, 0x00, 0xFF,
	}

	switch direction {
	case "up":
		direc[6] = 0x03
		direc[7] = 0x01
	case "down":
		direc[6] = 0x03
		direc[7] = 0x02
	case "right":
		direc[6] = 0x02
		direc[7] = 0x03
	case "left":
		direc[6] = 0x01
		direc[7] = 0x03
	case "stop":
		direc[6] = 0x03
		direc[7] = 0x03
	}

	if _, err := connection.Write(direc); err != nil {
		return err
	} else {
		return nil
	}
}

func Zoom(zoomType string, connection *net.TCPConn) error {
	zoomCommand := []byte{
		0x81, 0x01, 0x04, 0x07, 0x00, 0xFF,
	}
	switch zoomType {
	case "in":
		zoomCommand[4] = 0x02
	case "out":
		zoomCommand[4] = 0x03
	case "stop":
		zoomCommand[4] = 0x00
	}

	if _, err := connection.Write(zoomCommand); err != nil {
		return err
	} else {
		return nil
	}
}
