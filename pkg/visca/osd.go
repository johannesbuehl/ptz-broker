package visca

func (c *Camera) OpenMenu() error {
	command := []byte{
		0x81, 0x01, 0x04, 0x3F, 0x02, 0x5F, 0xFF,
	}
	if err := c.sendCommand(command); err != nil {
		return err
	} else {
		return nil
	}
}

func (c *Camera) Enter() error {
	command := []byte{
		0x81, 0x01, 0x06, 0x06, 0x05, 0xFF,
	}

	if err := c.sendCommand(command); err != nil {
		return err
	} else {
		return nil
	}
}
