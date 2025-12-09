package config

import (
	"fmt"

	"github.com/johannesbuehl/ptz-broker/pkg/positionPreset"
)

type presets struct {
	Positions map[string]positionPreset.Position `json:"positions" validate:"required"`
}

type Adress struct {
	Ip   string `json:"ip" validate:"required,ipv4"`
	Port uint   `json:"port" validate:"required,port"`
}

type Config struct {
	Presets presets `json:"presets" validate:"required"`
	Camera  struct {
		Adress Adress `json:"adress" validate:"required"`
		Speed  string `json:"speed" validate:"required"`
	} `json:"camera" validate:"required"`
	OSCPort uint `json:"osc_port" validate:"required,port"`

	path string
}

func (a Adress) GetString() string {
	return fmt.Sprintf("%s:%d", a.Ip, a.Port)
}
