package main

import (
	"fmt"
	"net"
	"strings"

	"github.com/hypebeast/go-osc/osc"
	"github.com/johannesbuehl/ptz-broker/pkg/cameraControl"
	"github.com/johannesbuehl/ptz-broker/pkg/positionPreset"
)

func getData(msg *osc.Message) string {
	message := getData(msg)
	lastSpaceIndex := strings.LastIndex(message, " ")
	return message[lastSpaceIndex+1:]
}

func recallPreset(msg *osc.Message, connection *net.TCPConn) {
	if typeString, err := msg.TypeTags(); typeString == ",s" && err == nil {
		message := getData(msg)

		if position, check := configFile.Presets.Positions[message]; check {
			position.RecallCameraPosition(connection)
		}
	}
}

func savePreset(msg *osc.Message, connection *net.TCPConn) {
	if typeString, err := msg.TypeTags(); typeString == ",s" && err == nil {
		message := getData(msg)

		if position, err := positionPreset.GetCameraPosition(connection); err == nil {
			configFile.Presets.Positions[message] = position
			configFile.Save()
		}
	}
}

func moveCamera(msg *osc.Message, connection *net.TCPConn) {
	if typeString, err := msg.TypeTags(); typeString == ",s" && err == nil {
		data := getData(msg)
		fmt.Println(data)
		cameraControl.Move(configFile.Camera.Speed, data, connection)
	}
}

func zoomCamera(msg *osc.Message, connection *net.TCPConn) {
	if typeString, err := msg.TypeTags(); typeString == ",s" && err == nil {
		message := getData(msg)

		cameraControl.Zoom(message, connection)
	}
}

func setSpeed(msg *osc.Message, connection *net.TCPConn) {
	if typeString, err := msg.TypeTags(); typeString == ",s" && err == nil {
		message := getData(msg)

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

func modeWhiteBalance(msg *osc.Message, connection *net.TCPConn) {
	if typeString, err := msg.TypeTags(); typeString == ",s" && err == nil {
		message := getData(msg)

		cameraControl.ModeWhiteBalance(message, connection)
	}
}

func manuelColorTemperature(msg *osc.Message, connection *net.TCPConn) {
	if typeString, err := msg.TypeTags(); typeString == ",s" && err == nil {
		message := getData(msg)

		cameraControl.ManuelColorTemperature(message, connection)
	}
}

func redGain(msg *osc.Message, connection *net.TCPConn) {
	if typeString, err := msg.TypeTags(); typeString == ",s" && err == nil {
		message := getData(msg)

		cameraControl.RedGain(message, connection)
	}
}

func blueGain(msg *osc.Message, connection *net.TCPConn) {
	if typeString, err := msg.TypeTags(); typeString == ",s" && err == nil {
		message := getData(msg)

		cameraControl.BlueGain(message, connection)
	}
}

func saveColorTemperatur(msg *osc.Message, connection *net.TCPConn) {
	if message, err := cameraControl.SaveColorTemperatur(connection); message == ",s" && err == nil {
		configFile.Camera.WhiteBalance = message
		configFile.Save()
	}

}

func recallColorTemperatur(msg *osc.Message, connection *net.TCPConn) {
	cameraControl.RecallColorTemperatur(configFile.Camera.WhiteBalance, connection)
}
