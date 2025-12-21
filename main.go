package main

import (
	"fmt"

	"github.com/hypebeast/go-osc/osc"
	"github.com/johannesbuehl/ptz-broker/pkg/config"
	"github.com/johannesbuehl/ptz-broker/pkg/visca"
)

var configFile config.Config

func main() {
	var err error
	if configFile, err = config.Load("config.json"); err != nil {
		panic(err)
	} else {
		camera := visca.Camera{
			IP:   configFile.Camera.Adress.IP,
			Port: configFile.Camera.Adress.Port,
		}

		if err := camera.Connect(); err != nil {
			panic(err)
		} else {
			defer camera.Disconnect()

			addr := fmt.Sprintf("0.0.0.0:%d", configFile.OSCPort)
			d := osc.NewStandardDispatcher()

			endpoints := map[string]func(*osc.Message, *visca.Camera){
				"/preset/position/recall":           recallPreset,
				"/preset/position/save":             savePreset,
				"/control/move":                     moveCamera,
				"/control/zoom":                     zoomCamera,
				"/control/speed/set":                setSpeed,
				"/control/menu/open":                openMenu,
				"/control/menu/enter":               enter,
				"/picture/color/whitebalance":       modeWhiteBalance,
				"/picture/color/temperature/manuel": manuelColorTemperature,
				"/picture/color/redgain":            redGain,
				"/picture/color/bluegain":           blueGain,
				// "/picture/color/temperature/save":   saveColorTemperatur,
				// "/picture/color/temperature/recall": recallColorTemperatur,
			}

			for endpoint, ff := range endpoints {
				d.AddMsgHandler(endpoint, func(msg *osc.Message) {
					ff(msg, &camera)
				})
			}

			server := &osc.Server{
				Addr:       addr,
				Dispatcher: d,
			}
			server.ListenAndServe()

			// Testcode Get Position from camera
			// if Position, err := positionPreset.GetCameraPosition(connection); err != nil {
			// 	fmt.Println(err)
			// } else {
			// 	configFile.Presets.Positions["Altar"] = Position
			// }
		}
	}
}
