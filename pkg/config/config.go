package config

import (
	"fmt"

	"github.com/johannesbuehl/ptz-broker/pkg/visca"
)

type Position struct {
	PanTilt visca.PanTilt `json:"pan_tilt" validate:"required"`
	Zoom    visca.Zoom    `json:"zoom" validate:"required"`
}

type Color struct {
	ColorTemperature visca.ColorTemperature
	RedGain          visca.RedGain
	BlueGain         visca.BlueGain
	Hue              visca.Hue
}

type Presets struct {
	Positions map[string]Position `json:"positions" validate:"required"`
	Color     map[string]Color    `json:"color" validate:"required"`
}

type Adress struct {
	IP   string `json:"ip" validate:"required,ipv4"`
	Port uint16 `json:"port" validate:"required,port"`
}

type Config struct {
	Presets Presets `json:"presets" validate:"required"`
	Camera  struct {
		Adress Adress      `json:"adress" validate:"required"`
		Speed  visca.Speed `json:"speed" validate:"required"`
	} `json:"camera" validate:"required"`
	OSCPort uint `json:"osc_port" validate:"required,port"`

	path string
}

func (a Adress) GetString() string {
	return fmt.Sprintf("%s:%d", a.IP, a.Port)
}
