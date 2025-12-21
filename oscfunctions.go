package main

import (
	"fmt"

	"github.com/hypebeast/go-osc/osc"
	"github.com/johannesbuehl/ptz-broker/pkg/cameraControl"
	"github.com/johannesbuehl/ptz-broker/pkg/config"
	"github.com/johannesbuehl/ptz-broker/pkg/visca"
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

func recallPreset(msg *osc.Message, camera *visca.Camera) {
	if message, err := getString(msg); err == nil {
		if position, check := configFile.Presets.Positions[message]; check {
			if camera.PanTiltAbsolute(position.PanTilt) == nil {
				camera.Zoom(position.Zoom)
			}
		}
	}
}

func savePreset(msg *osc.Message, camera *visca.Camera) {
	if message, err := getString(msg); err == nil {
		if panTilt, err := camera.GetPanTilt(); err == nil {
			if zoom, err := camera.GetZoom(); err == nil {
				configFile.Presets.Positions[message] = config.Position{
					PanTilt: panTilt,
					Zoom:    zoom,
				}
				configFile.Save()
			}
		}
	}
}

func moveCamera(msg *osc.Message, camera *visca.Camera) {
	if message, err := getString(msg); err == nil {
		fmt.Println(message)
		camera.Move(visca.Direction(message), configFile.Camera.Speed)
	}
}

func zoomCamera(msg *osc.Message, camera *visca.Camera) {
	if message, err := getString(msg); err == nil {
		camera.Zoom(visca.ZoomDirection(message))
	}
}

func setSpeed(msg *osc.Message, camera *visca.Camera) {
	if message, err := getInteger(msg); err == nil {
		configFile.Camera.Speed = visca.Speed(message)
		configFile.Save()
	}
}

func openMenu(msg *osc.Message, camera *visca.Camera) {
	cameraControl.OpenMenu(connection)
}

func enter(msg *osc.Message, camera *visca.Camera) {
	cameraControl.Enter(connection)
}

func modeWhiteBalance(msg *osc.Message, camera *visca.Camera) {
	if message, err := getString(msg); err == nil {
		cameraControl.ModeWhiteBalance(message, connection)
	}
}

func manuelColorTemperature(msg *osc.Message, camera *visca.Camera) {
	if message, err := getString(msg); err == nil {
		cameraControl.ManuelColorTemperature(message, connection)
	}
}

func redGain(msg *osc.Message, camera *visca.Camera) {
	if message, err := getString(msg); err == nil {
		cameraControl.RedGain(message, connection)
	}
}

func blueGain(msg *osc.Message, camera *visca.Camera) {
	if message, err := getString(msg); err == nil {
		cameraControl.BlueGain(message, connection)
	}
}

// func saveColorTemperatur(msg *osc.Message, camera *visca.Camera) {
// 	if message, err := cameraControl.SaveColorTemperatur(connection); err == nil {
// 		configFile.Camera.WhiteBalance = bytes(message)
// 		configFile.Save()
// 	}

// }

// func recallColorTemperatur(msg *osc.Message, camera *visca.Camera) {
// 	cameraControl.RecallColorTemperatur(configFile.Camera.WhiteBalance, connection)
// }
