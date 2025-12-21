package visca

type PanTilt struct {
	Pan  [4]byte
	Tilt [4]byte
}

func (c *Camera) PanTiltAbsolute(pt PanTilt) error {
	command := []byte{
		0x81, 0x01, 0x06, 0x02, 0x18, 0x14, pt.Pan[0], pt.Pan[1], pt.Pan[2], pt.Pan[3], pt.Tilt[0], pt.Tilt[1], pt.Tilt[2], pt.Tilt[3], 0xFF,
	}

	if ch, err := c.sendCommand(command); err != nil {
		return err
	} else {
		// wait for message
		if msg, err := getResponse(ch); err != nil {
			return err
		} else {
			// if the message is an acknowledge, wait for the completion
			if resp, _ := parseViscaResponse(msg); resp == RESP_ACK {
				if _, err := getResponse(ch); err != nil {
					return err
				}
			}

			return nil
		}
	}
}

func (c *Camera) GetPanTilt() (PanTilt, error) {
	if ch, err := c.sendCommand([]byte{0x81, 0x09, 0x06, 0x12, 0xFF}); err != nil {
		return PanTilt{}, err
	} else {
		response := <-ch

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
		}, err
	}
}

type Zoom [4]byte

func (c *Camera) Zoom(zoom Zoom) error {
	command := []byte{
		0x81, 0x01, 0x04, 0x47, zoom[0], zoom[1], zoom[2], zoom[3], 0xFF,
	}

	if ch, err := c.sendCommand(command); err != nil {
		return err
	} else {
		// wait for message
		if msg, err := getResponse(ch); err != nil {
			return err
		} else {
			// if the message is an acknowledge, wait for the completion
			if resp, _ := parseViscaResponse(msg); resp == RESP_ACK {
				if _, err := getResponse(ch); err != nil {
					return err
				}
			}

			return nil
		}
	}
}

func (c *Camera) GetZoom() (Zoom, error) {
	if ch, err := c.sendCommand([]byte{0x81, 0x09, 0x06, 0x12, 0xFF}); err != nil {
		return Zoom{}, err
	} else {
		response := <-ch

		return Zoom{
			response[2],
			response[3],
			response[4],
			response[5],
		}, err
	}
}
