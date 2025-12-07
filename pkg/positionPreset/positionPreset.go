package positionpreset

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

func (p Position) Recall(connection *net.TCPConn) error {
	command := "81 01 06 02 18 14 " + strings.Join([]string{p.Pan, p.Tilt}, " ") + " FF"

	if _, err := connection.Write([]byte(command)); err != nil {
		return err
	} else {
		return nil
	}
}

func GetCameraPosition(connection *net.TCPConn) (Position, error) {
	mu := getMutex{
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

type getMutex struct {
	mu       sync.Mutex
	ch       chan error
	Position Position
}

func (c *getMutex) getPanTilt(connection *net.TCPConn) {
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

func (c *getMutex) getZoom(connection *net.TCPConn) {
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
