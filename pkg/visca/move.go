package visca

type Direction string

const (
	DirectionUp    Direction = "up"
	DirectionDown  Direction = "down"
	DirectionRight Direction = "right"
	DirectionLeft  Direction = "left"
	DirectionStop  Direction = "stop"
)

type Speed byte

func (c *Camera) Move(direction Direction, speed Speed) error {
	command := []byte{
		0x81, 0x01, 0x06, 0x01, byte(speed), byte(speed), 0x00, 0x00, 0xFF,
	}

	switch direction {
	case DirectionUp:
		command[6] = 0x03
		command[7] = 0x01
	case DirectionDown:
		command[6] = 0x03
		command[7] = 0x02
	case DirectionRight:
		command[6] = 0x02
		command[7] = 0x03
	case DirectionLeft:
		command[6] = 0x01
		command[7] = 0x03
	case DirectionStop:
		command[6] = 0x03
		command[7] = 0x03
	}

	if err := c.sendCommand(command); err != nil {
		return err
	} else {
		return nil
	}
}
