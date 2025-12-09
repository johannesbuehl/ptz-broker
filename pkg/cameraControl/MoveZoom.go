package cameraControl

import (
	"net"
	"strings"
)

func Move(speed, direction string, connection *net.TCPConn) error {
	var direc string
	switch direction {
	case "up":
		direc = "03 01"
	case "down":
		direc = "03 02"
	case "right":
		direc = "02 03"
	case "left":
		direc = "01 03"
	case "stop":
		direc = "03 03"
	}
	command := "81 01 06 01 " + strings.Join([]string{speed, speed, direc}, " ") + " FF"

	if _, err := connection.Write([]byte(command)); err != nil {
		return err
	} else {
		return nil
	}
}

func Zoom(zoomType string, connection *net.TCPConn) error {
	var zoomCommand string
	switch zoomType {
	case "in":
		zoomCommand = "02"
	case "out":
		zoomCommand = "03"
	case "stop":
		zoomCommand = "00"
	}

	command := "81 01 04 07 " + zoomCommand + " FF"

	if _, err := connection.Write([]byte(command)); err != nil {
		return err
	} else {
		return nil
	}
}
