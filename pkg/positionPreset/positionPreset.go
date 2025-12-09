package positionPreset

import (
	"bufio"
	"net"
	"strings"
	"sync"
)

type Position struct {
	Pan  string `json:"pan"`
	Tilt string `json:"tilt"`
	Zoom string `json:"zoom"`
}

type positionMutex struct {
	mu       sync.Mutex
	ch       chan error
	Position Position
}

func (p Position) RecallCameraPosition(connection *net.TCPConn) error {
	mu := positionMutex{
		ch: make(chan error),
	}

	go mu.Position.recallPanTilt(connection)
	go mu.Position.recallZoom(connection)

	err1, err2 := <-mu.ch, <-mu.ch

	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	return nil

}

func (p Position) recallPanTilt(connection *net.TCPConn) error {
	command := "81 01 06 02 18 14 " + strings.Join([]string{p.Pan, p.Tilt}, " ") + " FF"

	if _, err := connection.Write([]byte(command)); err != nil {
		return err
	} else {
		return nil
	}
}
func (z Position) recallZoom(connection *net.TCPConn) error {
	command := "81 01 04 47 " + z.Zoom + " FF"

	if _, err := connection.Write([]byte(command)); err != nil {
		return err
	} else {
		return nil
	}
}

func GetCameraPosition(connection *net.TCPConn) (Position, error) {
	mu := positionMutex{
		ch: make(chan error),
	}

	go mu.getPanTilt(connection)
	go mu.getZoom(connection)

	err1, err2 := <-mu.ch, <-mu.ch

	if err1 != nil {
		return Position{}, err1
	}
	if err2 != nil {
		return Position{}, err2
	}
	return mu.Position, nil
}

func (c *positionMutex) getPanTilt(connection *net.TCPConn) {
	if _, err := connection.Write([]byte("81 09 06 12 FF\n")); err != nil {
		c.ch <- err
	} else {
		reader := bufio.NewReader(connection)

		if line, err := reader.ReadString('\n'); err != nil {
			c.ch <- err
		} else {
			c.mu.Lock()
			c.Position.Pan = line[6:17]
			c.Position.Tilt = line[18:29]
			c.mu.Unlock()
			c.ch <- nil
		}
	}
}

func (c *positionMutex) getZoom(connection *net.TCPConn) {
	if _, err := connection.Write([]byte("81 09 04 47 FF\n")); err != nil {
		c.ch <- err
	} else {
		reader := bufio.NewReader(connection)

		if line, err := reader.ReadString('\n'); err != nil {
			c.ch <- err
		} else {
			c.mu.Lock()
			c.Position.Zoom = line[6:17]
			c.mu.Unlock()
			c.ch <- nil
		}
	}
}
