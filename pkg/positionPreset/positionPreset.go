package positionPreset

import (
	"net"
	"sync"
)

type Position struct {
	Pan  [4]byte `json:"pan"`
	Tilt [4]byte `json:"tilt"`
	Zoom [4]byte `json:"zoom"`
}

type positionMutex struct {
	mu       sync.Mutex
	ch       chan error
	Position Position
}

func (p Position) RecallCameraPosition(connection *net.TCPConn) error {
	mu := positionMutex{
		ch:       make(chan error),
		Position: p,
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
	command := []byte{
		0x81, 0x01, 0x06, 0x02, 0x18, 0x14, p.Pan[0], p.Pan[1], p.Pan[2], p.Pan[3], p.Tilt[0], p.Tilt[1], p.Tilt[2], p.Tilt[3], 0xFF,
	}

	if _, err := connection.Write(command); err != nil {
		return err
	} else {
		return nil
	}
}
func (z Position) recallZoom(connection *net.TCPConn) error {
	command := []byte{
		0x81, 0x01, 0x04, 0x47, z.Zoom[0], z.Zoom[1], z.Zoom[2], z.Zoom[3], 0xFF,
	}

	if _, err := connection.Write(command); err != nil {
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
	if _, err := connection.Write([]byte{0x81, 0x09, 0x06, 0x12, 0xFF}); err != nil {
		c.ch <- err
	} else {
		response := make([]byte, 32)

		if _, err := connection.Read(response); err != nil {
			c.ch <- err
		} else {
			c.mu.Lock()
			// c.Position.Pan = line[6:17]
			// c.Position.Tilt = line[18:29]
			c.Position.Pan[0] = response[2]
			c.Position.Pan[1] = response[3]
			c.Position.Pan[2] = response[4]
			c.Position.Pan[3] = response[5]
			c.Position.Tilt[0] = response[6]
			c.Position.Tilt[1] = response[7]
			c.Position.Tilt[2] = response[8]
			c.Position.Tilt[3] = response[9]
			c.mu.Unlock()
			c.ch <- nil
		}
	}
}

func (c *positionMutex) getZoom(connection *net.TCPConn) {
	if _, err := connection.Write([]byte{0x81, 0x09, 0x04, 0x47, 0xFF}); err != nil {
		c.ch <- err
	} else {
		response := make([]byte, 7)

		if _, err := connection.Read(response); err != nil {
			c.ch <- err
		} else {
			c.mu.Lock()
			c.Position.Zoom[0] = response[2]
			c.Position.Zoom[1] = response[3]
			c.Position.Zoom[2] = response[4]
			c.Position.Zoom[3] = response[5]
			c.mu.Unlock()

			c.ch <- nil
		}
	}
}
