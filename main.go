package main

import (
	"fmt"
	"net"

	positionpreset "github.com/johannesbuehl/ptz-broker/pkg/positionPreset"
)

type presets struct {
	Positions map[string]positionpreset.Position `json:"positions"`
}

type config struct {
	Presets presets `json:"presets"`
}

func main() {
	fmt.Println("Hello World")

	if tcpAddress, err := net.ResolveTCPAddr("tcp", "tcpbin.com:4242"); err != nil {
		panic(err)
	} else if connection, err := net.DialTCP("tcp", nil, tcpAddress); err != nil {
		panic(err)

	} else {
		defer connection.Close()
		connection.SetKeepAlive(true)
		positionpreset.GetCameraPosition(connection)
	}
}
