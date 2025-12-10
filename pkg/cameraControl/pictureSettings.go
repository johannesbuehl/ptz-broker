package cameraControl

import (
	"bufio"
	"net"
)

func ModeWhiteBalance(whiteBalanceMode string, connection *net.TCPConn) error {
	var mode string

	switch whiteBalanceMode {
	case "auto":
		mode = "35 00"
	case "onepush":
		mode = "35 03"
	case "indoor":
		mode = "35 01"
	case "outdoor":
		mode = "35 02"
	case "manuel":
		mode = "35 05"
	case "trigger":
		mode = "10 05"
	}
	command := "81 01 04 " + mode + " FF"

	if _, err := connection.Write([]byte(command)); err != nil {
		return err
	} else {
		return nil
	}
}

func ManuelColorTemperature(valueManuel string, connection *net.TCPConn) error {
	var value string
	switch valueManuel {
	case "up":
		value = "02"
	case "down":
		value = "03"
	}
	command := "81 01 04 20 " + value + " FF"

	if _, err := connection.Write([]byte(command)); err != nil {
		return err
	} else {
		return nil
	}
}

func RedGain(valueManuel string, connection *net.TCPConn) error {
	var value string
	switch valueManuel {
	case "up":
		value = "02"
	case "down":
		value = "03"
	}
	command := "81 01 04 03 " + value + " FF"

	if _, err := connection.Write([]byte(command)); err != nil {
		return err
	} else {
		return nil
	}
}

func BlueGain(valueManuel string, connection *net.TCPConn) error {
	var value string
	switch valueManuel {
	case "up":
		value = "02"
	case "down":
		value = "03"
	}
	command := "81 01 04 04 " + value + " FF"

	if _, err := connection.Write([]byte(command)); err != nil {
		return err
	} else {
		return nil
	}
}

func SaveColorTemperatur(connection *net.TCPConn) (string, error) {
	command := "81 09 04 20 FF"

	if _, err := connection.Write([]byte(command)); err != nil {
		return "", err
	} else {
		reader := bufio.NewReader(connection)
		if line, err := reader.ReadString('\n'); err != nil {
			return "", err
		} else {
			whiteBalance := line[6:8]
			return whiteBalance, nil
		}
	}
}

func RecallColorTemperatur(whiteBalance string, connection *net.TCPConn) error {
	command := "81 01 04 20 " + whiteBalance + " FF"

	if _, err := connection.Write([]byte(command)); err != nil {
		return err
	} else {
		return nil
	}
}
