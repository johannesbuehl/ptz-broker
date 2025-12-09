package config

import (
	"encoding/json"
	"os"
)

func (conf Config) Save() error {
	if cont, err := json.MarshalIndent(conf, "", "\t"); err != nil {
		return err
	} else if err := os.WriteFile(conf.path, cont, 0644); err != nil {
		return err
	} else {
		return nil
	}
}
