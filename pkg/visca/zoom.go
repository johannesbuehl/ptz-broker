package visca

type Zoom [4]byte

// ZoomAbsolute the camera
func (c *Camera) ZoomAbsolute(zoom Zoom) error {
	command := []byte{
		0x81, 0x01, 0x04, 0x47, zoom[0], zoom[1], zoom[2], zoom[3], 0xFF,
	}

	if err := c.sendCommand(command); err != nil {
		return err
	} else {
		return nil
	}
}

type ZoomDirection string

const (
	ZoomIn   ZoomDirection = "in"
	ZoomOut  ZoomDirection = "out"
	ZoomStop ZoomDirection = "stop"
)

func (c *Camera) Zoom(direction ZoomDirection) error {
	command := []byte{
		0x81, 0x01, 0x04, 0x07, 0x00, 0xFF,
	}
	switch direction {
	case ZoomIn:
		command[4] = 0x02
	case ZoomOut:
		command[4] = 0x03
	case ZoomStop:
		command[4] = 0x00
	}

	if err := c.sendCommand(command); err != nil {
		return err
	} else {
		return nil
	}
}

// GetZoom of the camera
func (c *Camera) GetZoom() (Zoom, error) {
	if response, err := c.sendInquiry([]byte{0x81, 0x09, 0x06, 0x12, 0xFF}); err != nil {
		return Zoom{}, err
	} else {

		return Zoom{
			response[2],
			response[3],
			response[4],
			response[5],
		}, err
	}
}
