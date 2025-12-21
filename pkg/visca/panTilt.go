package visca

type PanTilt struct {
	Pan  [4]byte `json:"pan"`
	Tilt [4]byte `json:"tilt"`
}

// PanTiltAbsolute sends absolute pan and tilt positions to the camera
func (c *Camera) PanTiltAbsolute(pt PanTilt) error {
	command := []byte{
		0x81, 0x01, 0x06, 0x02, 0x18, 0x14, pt.Pan[0], pt.Pan[1], pt.Pan[2], pt.Pan[3], pt.Tilt[0], pt.Tilt[1], pt.Tilt[2], pt.Tilt[3], 0xFF,
	}

	if err := c.sendCommand(command); err != nil {
		return err
	} else {
		return nil
	}
}

// GetPanTilt position from the camera
func (c *Camera) GetPanTilt() (PanTilt, error) {
	if response, err := c.sendInquiry([]byte{0x81, 0x09, 0x06, 0x12, 0xFF}); err != nil {
		return PanTilt{}, err
	} else {
		return PanTilt{
			Pan: [4]byte{
				response[2],
				response[3],
				response[4],
				response[5],
			},
			Tilt: [4]byte{
				response[6],
				response[7],
				response[8],
				response[9],
			},
		}, nil
	}
}
