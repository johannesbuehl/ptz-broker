package main

import (
	"net"

	"github.com/hypebeast/go-osc/osc"
	"github.com/johannesbuehl/ptz-broker/pkg/cameraControl"
	"github.com/johannesbuehl/ptz-broker/pkg/positionPreset"
)

func recallPreset(msg *osc.Message, connection *net.TCPConn) {
	if typeString, err := msg.TypeTags(); typeString == "s" && err == nil {
		message := msg.String()

		if position, check := configFile.Presets.Positions[message]; check {
			position.RecallCameraPosition(connection)
		}
	}
}

func savePreset(msg *osc.Message, connection *net.TCPConn) {
	if typeString, err := msg.TypeTags(); typeString == "s" && err == nil {
		message := msg.String()

		if position, err := positionPreset.GetCameraPosition(connection); err == nil {
			configFile.Presets.Positions[message] = position
			configFile.Save()
		}
	}
}

func moveCamera(msg *osc.Message, connection *net.TCPConn) {
	if typeString, err := msg.TypeTags(); typeString == "s" && err == nil {
		message := msg.String()

		cameraControl.Move(configFile.Camera.Speed, message, connection)
	}
}

func zoomCamera(msg *osc.Message, connection *net.TCPConn) {
	if typeString, err := msg.TypeTags(); typeString == "s" && err == nil {
		message := msg.String()

		cameraControl.Zoom(message, connection)
	}
}

func setSpeed(msg *osc.Message, connection *net.TCPConn) {
	if typeString, err := msg.TypeTags(); typeString == "s" && err == nil {
		message := msg.String()

		configFile.Camera.Speed = message
		configFile.Save()
	}
}

func openMenu(msg *osc.Message, connection *net.TCPConn) {
	cameraControl.OpenMenu(connection)
}

func enter(msg *osc.Message, connection *net.TCPConn) {
	cameraControl.Enter(connection)
}
