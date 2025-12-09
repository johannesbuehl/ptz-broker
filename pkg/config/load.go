package config

import (
	"encoding/json"
	"os"
)

func Load(pth string) (Config, error) {
	var conf Config

	if cont, err := os.ReadFile(pth); err != nil {
		return conf, err
	} else if err := json.Unmarshal(cont, &conf); err != nil {
		return conf, err
	} else {
		conf.path = pth
		return conf, nil
	}
}
