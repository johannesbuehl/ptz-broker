package cameraControl

import (
	"bufio"
	"net"
)

func ModeWhiteBalance(whiteBalanceMode string, connection *net.TCPConn) error {
	command := []byte{
		0x81, 0x01, 0x04, 0x00, 0x00, 0xFF,
	}

	switch whiteBalanceMode {
	case "auto":
		command[3] = 0x35
		command[4] = 0x00
	case "onepush":
		command[3] = 0x35
		command[4] = 0x03
	case "indoor":
		command[3] = 0x35
		command[4] = 0x01
	case "outdoor":
		command[3] = 0x35
		command[4] = 0x02
	case "manuel":
		command[3] = 0x35
		command[4] = 0x05
	case "trigger":
		command[3] = 0x10
		command[4] = 0x05
	}

	if _, err := connection.Write([]byte(command)); err != nil {
		return err
	} else {
		return nil
	}
}

func ManuelColorTemperature(valueManuel string, connection *net.TCPConn) error {
	command := []byte{
		0x81, 0x01, 0x04, 0x20, 0x00, 0xFF,
	}
	switch valueManuel {
	case "up":
		command[4] = 0x02
	case "down":
		command[4] = 0x03
	}

	if _, err := connection.Write([]byte(command)); err != nil {
		return err
	} else {
		return nil
	}
}

func RedGain(valueManuel string, connection *net.TCPConn) error {
	command := []byte{
		0x81, 0x01, 0x04, 0x03, 0x00, 0xFF,
	}
	switch valueManuel {
	case "up":
		command[4] = 0x02
	case "down":
		command[4] = 0x03
	}

	if _, err := connection.Write([]byte(command)); err != nil {
		return err
	} else {
		return nil
	}
}

func BlueGain(valueManuel string, connection *net.TCPConn) error {
	command := []byte{
		0x81, 0x01, 0x04, 0x04, 0x00, 0xFF,
	}
	switch valueManuel {
	case "up":
		command[4] = 0x02
	case "down":
		command[4] = 0x03
	}

	if _, err := connection.Write([]byte(command)); err != nil {
		return err
	} else {
		return nil
	}
}

func SaveColorTemperatur(connection *net.TCPConn) (int, error) {
	command := []byte{
		0x81, 0x09, 0x04, 0x20, 0xFF,
	}
	if _, err := connection.Write([]byte(command)); err != nil {
		return 0, err
	} else {
		reader := bufio.NewReader(connection)
		if _, err := reader.ReadString('\n'); err != nil {
			return 0, err
		} else {
			// whiteBalance := (line[6:8])
			return 0, nil
		}
	}
}

func RecallColorTemperatur(whiteBalance byte, connection *net.TCPConn) error {
	command := []byte{
		0x81, 0x01, 0x04, 0x20, whiteBalance, 0xFF,
	}

	if _, err := connection.Write([]byte(command)); err != nil {
		return err
	} else {
		return nil
	}
}
