package config

import (
	"encoding/json"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/johannesbuehl/ptz-broker/pkg/positionPreset"
)

func Load(pth string) (Config, error) {
	conf := Config{
		path: pth,
	}

	if cont, err := os.ReadFile(pth); err != nil {
		if os.IsNotExist(err) {
			conf.Presets.Positions = map[string]positionPreset.Position{}
			if err := conf.Save(); err != nil {
				return conf, err
			} else {
				os.Exit(0)
			}
		}

		return conf, err
	} else if err := json.Unmarshal(cont, &conf); err != nil {
		return conf, err
	} else {
		validate := validator.New(validator.WithRequiredStructEnabled())
		if err := validate.Struct(conf); err != nil {
			return conf, err
		} else {
			return conf, nil
		}
	}
}
