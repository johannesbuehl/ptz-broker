package config

import positionpreset "github.com/johannesbuehl/ptz-broker/pkg/positionPreset"

type presets struct {
	Positions map[string]positionpreset.Position `json:"positions"`
}

type Config struct {
	Presets presets `json:"presets"`
	path    string
}
