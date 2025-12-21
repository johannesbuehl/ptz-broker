package visca

type ColorTemperature [2]byte

type BlueGain [2]byte

type RedGain [2]byte

type Hue [1]byte

type WhiteBalanceMode uint8

const (
	WBAuto WhiteBalanceMode = iota
	WBOnePush
	WBIndoor
	WBOutdoor
	WBManuel
	WBTrigger
)

func (c *Camera) ModeWhiteBalance(whiteBalanceMode WhiteBalanceMode) error {
	command := []byte{
		0x81, 0x01, 0x04, 0x00, 0x00, 0xFF,
	}

	switch whiteBalanceMode {
	case WBAuto:
		command[3] = 0x35
		command[4] = 0x00
	case WBOnePush:
		command[3] = 0x35
		command[4] = 0x03
	case WBIndoor:
		command[3] = 0x35
		command[4] = 0x01
	case WBOutdoor:
		command[3] = 0x35
		command[4] = 0x02
	case WBManuel:
		command[3] = 0x35
		command[4] = 0x05
	case WBTrigger:
		command[3] = 0x10
		command[4] = 0x05
	}

	if err := c.sendCommand(command); err != nil {
		return err
	} else {
		return nil
	}
}

func (c *Camera) DecreaseColorTemperature() error {
	command := []byte{
		0x81, 0x01, 0x04, 0x20, 0x03, 0xFF,
	}

	if err := c.sendCommand(command); err != nil {
		return err
	} else {
		return nil
	}
}

func (c *Camera) IncreaseColorTemperature() error {
	command := []byte{
		0x81, 0x01, 0x04, 0x20, 0x02, 0xFF,
	}

	if err := c.sendCommand(command); err != nil {
		return err
	} else {
		return nil
	}
}

func (c *Camera) DecreaseRedGain() error {
	command := []byte{
		0x81, 0x01, 0x04, 0x03, 0x03, 0xFF,
	}

	if err := c.sendCommand(command); err != nil {
		return err
	} else {
		return nil
	}
}

func (c *Camera) IncreaseRedGain() error {
	command := []byte{
		0x81, 0x01, 0x04, 0x03, 0x02, 0xFF,
	}

	if err := c.sendCommand(command); err != nil {
		return err
	} else {
		return nil
	}
}

func (c *Camera) DecreaseBlueGain() error {
	command := []byte{
		0x81, 0x01, 0x04, 0x04, 0x03, 0xFF,
	}

	if err := c.sendCommand(command); err != nil {
		return err
	} else {
		return nil
	}
}

func (c *Camera) IncreaseBlueGain() error {
	command := []byte{
		0x81, 0x01, 0x04, 0x04, 0x02, 0xFF,
	}

	if err := c.sendCommand(command); err != nil {
		return err
	} else {
		return nil
	}
}

func (c *Camera) GetColorTemperature() (ColorTemperature, error) {
	command := []byte{
		0x81, 0x09, 0x04, 0x20, 0xFF,
	}
	if response, err := c.sendInquiry(command); err != nil {
		return ColorTemperature{}, err
	} else {
		return ColorTemperature{
			(response[2] & 0b11110000) >> 4,
			response[2] & 0b00001111,
		}, nil
	}
}

func (c *Camera) SetColorTemperature(ColorTemperature ColorTemperature) error {
	command := []byte{
		0x81, 0x01, 0x04, 0x20, ColorTemperature[0], ColorTemperature[1], 0xFF,
	}

	if err := c.sendCommand(command); err != nil {
		return err
	} else {
		return nil
	}
}

func (c *Camera) GetRedGain() (RedGain, error) {
	command := []byte{
		0x81, 0x09, 0x04, 0x43, 0xFF,
	}
	if response, err := c.sendInquiry(command); err != nil {
		return RedGain{}, err
	} else {
		return RedGain{
			response[4],
			response[5],
		}, nil
	}
}

func (c *Camera) SetRedGain(RedGain RedGain) error {
	command := []byte{
		0x81, 0x01, 0x04, 0x43, 0x00, 0x00, RedGain[0], RedGain[1], 0xFF,
	}

	if err := c.sendCommand(command); err != nil {
		return err
	} else {
		return nil
	}
}

func (c *Camera) GetBlueGain() (BlueGain, error) {
	command := []byte{
		0x81, 0x09, 0x04, 0x44, 0xFF,
	}
	if response, err := c.sendInquiry(command); err != nil {
		return BlueGain{}, err
	} else {
		return BlueGain{
			response[4],
			response[5],
		}, nil
	}
}

func (c *Camera) SetBlueGain(BlueGain BlueGain) error {
	command := []byte{
		0x81, 0x01, 0x04, 0x44, 0x00, 0x00, BlueGain[0], BlueGain[1], 0xFF,
	}

	if err := c.sendCommand(command); err != nil {
		return err
	} else {
		return nil
	}
}

func (c *Camera) GetHue() (Hue, error) {
	command := []byte{
		0x81, 0x09, 0x04, 0x4F, 0xFF,
	}
	if response, err := c.sendInquiry(command); err != nil {
		return Hue{}, err
	} else {
		return Hue{
			response[5],
		}, nil
	}
}

func (c *Camera) SetHue(Hue Hue) error {
	command := []byte{
		0x81, 0x01, 0x04, 0x4F, 0x00, 0x00, 0x00, Hue[0], 0xFF,
	}

	if err := c.sendCommand(command); err != nil {
		return err
	} else {
		return nil
	}
}
