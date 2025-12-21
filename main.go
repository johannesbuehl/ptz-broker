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

			// endpoints := map[string]func(*osc.Message, *visca.Camera){
			// 	"/preset/position/recall": recallPreset,
			// 	"/preset/position/save":   savePreset,
			// 	"/control/move":           moveCamera,
			// 	"/control/zoom":           zoomCamera,
			// 	"/control/speed/set":      setSpeed,
			// 	"/control/menu/open":                openMenu,
			// 	"/control/menu/enter":               enter,
			// 	"/picture/color/whitebalance/auto":       autoWhiteBalance,
			// 	"/picture/color/whitebalance/onepush":    onePushWhiteBalance,
			// 	"/picture/color/whitebalance/indoor":     indoorWhiteBalance,
			// 	"/picture/color/whitebalance/outdoor":    outdoorWhiteBalance,
			// 	"/picture/color/whitebalance/manuel":     manuelWhiteBalance,
			// 	"/picture/color/whitebalance/trigger":    triggerWhiteBalance,
			// 	"/picture/color/temperature/manuel/up":   increaseColorTemperature,
			// 	"/picture/color/temperature/manuel/down": decreaseColorTemperature,
			// 	"/picture/color/redgain":                 redGain,
			// 	"/picture/color/bluegain":                blueGain,
			//  "/picture/color/temperature/save":   saveColorTemperatur,
			//  "/picture/color/temperature/recall": recallColorTemperatur,
			// }
			type oscmap map[string]any

			newendpoints := oscmap{
				"preset": oscmap{
					"position": oscmap{
						"recall": recallPositionPreset,
						"save":   savePositionPreset,
					},
					"color": oscmap{
						"recall": recallColorPreset,
						"save":   saveColorPreset,
					},
				},
				"control": oscmap{
					"move": moveCamera,
					"zoom": zoomCamera,
					"speed": oscmap{
						"set": setSpeed,
					},
					"menu": oscmap{
						"open":  openMenu,
						"enter": enter,
					},
				},
				"picture": oscmap{
					"color": oscmap{
						"whitebalance": oscmap{
							"auto":    autoWhiteBalance,
							"onepush": onePushWhiteBalance,
							"indoor":  indoorWhiteBalance,
							"outdoor": outdoorWhiteBalance,
							"manuel":  manuelWhiteBalance,
							"trigger": triggerWhiteBalance,
						},
						"temperature": oscmap{
							"manuel": oscmap{
								"up":   increaseColorTemperature,
								"down": decreaseColorTemperature,
							},
						},
						"redgain": oscmap{
							"up":   increaseRedGain,
							"down": decreaseRedGain,
						},
						"bluegain": oscmap{
							"up":   increaseBlueGain,
							"down": decreaseBlueGain,
						},
					},
				},
			}

			var addendpoints func(oscmap, string)
			addendpoints = func(endpoints oscmap, path string) {
				for endpoint, ff := range endpoints {
					completeendpoint := path + "/" + endpoint
					if v, ok := ff.(func(*osc.Message, *visca.Camera)); ok {
						d.AddMsgHandler(completeendpoint, func(msg *osc.Message) {
							v(msg, &camera)
						})
					} else {
						addendpoints(ff.(oscmap), completeendpoint)
					}
				}
			}

			addendpoints(newendpoints, "")

			// for endpoint, ff := range endpoints {
			// 	d.AddMsgHandler(endpoint, func(msg *osc.Message) {
			// 		ff(msg, &camera)
			// 	})
			// }

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
