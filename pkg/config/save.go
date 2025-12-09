package config

import (
	"encoding/json"
	"os"
)

func (conf Config) Save() error {
	if cont, err := json.Marshal(conf); err != nil {
		return err
	} else {
		os.WriteFile(conf.path, cont, 0644)
		return nil
	}
}
