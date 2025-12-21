package main

import (
	"fmt"

	"github.com/hypebeast/go-osc/osc"
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

func recallPositionPreset(msg *osc.Message, camera *visca.Camera) {
	if message, err := getString(msg); err == nil {
		if position, check := configFile.Presets.Positions[message]; check {
			if camera.PanTiltAbsolute(position.PanTilt) == nil {
				camera.ZoomAbsolute(position.Zoom)
			}
		}
	}
}

func savePositionPreset(msg *osc.Message, camera *visca.Camera) {
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
	camera.OpenMenu()
}

func enter(msg *osc.Message, camera *visca.Camera) {
	camera.Enter()
}

func autoWhiteBalance(msg *osc.Message, camera *visca.Camera) {
	camera.ModeWhiteBalance(visca.WBAuto)
}

func onePushWhiteBalance(msg *osc.Message, camera *visca.Camera) {
	camera.ModeWhiteBalance(visca.WBOnePush)
}

func indoorWhiteBalance(msg *osc.Message, camera *visca.Camera) {
	camera.ModeWhiteBalance(visca.WBIndoor)
}

func outdoorWhiteBalance(msg *osc.Message, camera *visca.Camera) {
	camera.ModeWhiteBalance(visca.WBOutdoor)
}

func manuelWhiteBalance(msg *osc.Message, camera *visca.Camera) {
	camera.ModeWhiteBalance(visca.WBManuel)
}

func triggerWhiteBalance(msg *osc.Message, camera *visca.Camera) {
	camera.ModeWhiteBalance(visca.WBTrigger)
}

func decreaseColorTemperature(msg *osc.Message, camera *visca.Camera) {
	camera.DecreaseColorTemperature()
}

func increaseColorTemperature(msg *osc.Message, camera *visca.Camera) {
	camera.IncreaseColorTemperature()
}

func decreaseRedGain(msg *osc.Message, camera *visca.Camera) {
	camera.DecreaseRedGain()
}

func increaseRedGain(msg *osc.Message, camera *visca.Camera) {
	camera.IncreaseRedGain()
}

func decreaseBlueGain(msg *osc.Message, camera *visca.Camera) {
	camera.DecreaseBlueGain()
}

func increaseBlueGain(msg *osc.Message, camera *visca.Camera) {
	camera.IncreaseBlueGain()
}

func recallColorPreset(msg *osc.Message, camera *visca.Camera) {
	if message, err := getString(msg); err == nil {
		if color, check := configFile.Presets.Color[message]; check {
			if camera.SetColorTemperature(color.ColorTemperature) == nil {
				if camera.SetRedGain(color.RedGain) == nil {
					if camera.SetBlueGain(color.BlueGain) == nil {
						camera.SetHue(color.Hue)
					}
				}
			}
		}
	}
}

func saveColorPreset(msg *osc.Message, camera *visca.Camera) {
	if message, err := getString(msg); err == nil {
		if colorTemperature, err := camera.GetColorTemperature(); err == nil {
			if redGain, err := camera.GetRedGain(); err == nil {
				if blueGain, err := camera.GetBlueGain(); err == nil {
					if hue, err := camera.GetHue(); err == nil {
						configFile.Presets.Color[message] = config.Color{
							ColorTemperature: colorTemperature,
							RedGain:          redGain,
							BlueGain:         blueGain,
							Hue:              hue,
						}
						configFile.Save()
					}
				}
			}
		}
	}
}
