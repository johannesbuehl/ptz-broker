package main

import (
	"fmt"
	"net"

	"github.com/hypebeast/go-osc/osc"
	"github.com/johannesbuehl/ptz-broker/pkg/cameraControl"
	"github.com/johannesbuehl/ptz-broker/pkg/positionPreset"
)

func getString(msg *osc.Message) (string, error) {
	if msg.CountArguments() != 1 {
		return "", fmt.Errorf("invalid argument count (must be 1)")
	} else if s, ok := msg.Arguments[0].(string); !ok {
		return "", fmt.Errorf("argument must be string")
	} else {
		return s, nil
	}
}

func getInteger(msg *osc.Message) (int32, error) {
	if msg.CountArguments() != 1 {
		return 0, fmt.Errorf("invalid argument count (must be 1)")
	} else if s, ok := msg.Arguments[0].(int32); !ok {
		return 0, fmt.Errorf("argument must be int")
	} else {
		return s, nil
	}
}

func recallPreset(msg *osc.Message, connection *net.TCPConn) {
	if message, err := getString(msg); err == nil {
		if position, check := configFile.Presets.Positions[message]; check {
			position.RecallCameraPosition(connection)
		}
	}
}

func savePreset(msg *osc.Message, connection *net.TCPConn) {
	if message, err := getString(msg); err == nil {
		if position, err := positionPreset.GetCameraPosition(connection); err != nil {
			fmt.Println(err)
		} else {
			configFile.Presets.Positions[message] = position
			configFile.Save()
		}
	}
}

func moveCamera(msg *osc.Message, connection *net.TCPConn) {
	if message, err := getString(msg); err == nil {
		fmt.Println(message)
		cameraControl.Move(configFile.Camera.Speed, message, connection)
	}
}

func zoomCamera(msg *osc.Message, connection *net.TCPConn) {
	if message, err := getString(msg); err == nil {
		cameraControl.Zoom(message, connection)
	}
}

func setSpeed(msg *osc.Message, connection *net.TCPConn) {
	if message, err := getInteger(msg); err == nil {
		configFile.Camera.Speed = byte(message)
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
	if message, err := getString(msg); err == nil {
		cameraControl.ModeWhiteBalance(message, connection)
	}
}

func manuelColorTemperature(msg *osc.Message, connection *net.TCPConn) {
	if message, err := getString(msg); err == nil {
		cameraControl.ManuelColorTemperature(message, connection)
	}
}

func redGain(msg *osc.Message, connection *net.TCPConn) {
	if message, err := getString(msg); err == nil {
		cameraControl.RedGain(message, connection)
	}
}

func blueGain(msg *osc.Message, connection *net.TCPConn) {
	if message, err := getString(msg); err == nil {
		cameraControl.BlueGain(message, connection)
	}
}

// func saveColorTemperatur(msg *osc.Message, connection *net.TCPConn) {
// 	if message, err := cameraControl.SaveColorTemperatur(connection); err == nil {
// 		configFile.Camera.WhiteBalance = bytes(message)
// 		configFile.Save()
// 	}

// }

// func recallColorTemperatur(msg *osc.Message, connection *net.TCPConn) {
// 	cameraControl.RecallColorTemperatur(configFile.Camera.WhiteBalance, connection)
// }
