package visca

import (
	"bufio"
	"fmt"
	"math"
	"net"
	"sync"
	"time"
)

type Camera struct {
	IP             string
	Port           uint16
	connection     *net.TCPConn
	mu             sync.Mutex
	channels       map[viscaSocket]chan []byte
	defaultChannel chan []byte
}

func (c *Camera) Connect() error {
	if tcpAddress, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", c.IP, c.Port)); err != nil {
		return err
	} else if c.connection, err = net.DialTCP("tcp", nil, tcpAddress); err != nil {
		return err
	} else if err := c.connection.SetKeepAlive(true); err != nil {
		return err
	} else {
		// initialize sockets
		c.defaultChannel = make(chan []byte)
		c.channels = make(map[viscaSocket]chan []byte)

		for socket := range viscaSocket(8) {
			c.channels[socket] = make(chan []byte)
		}

		go func() {
			reader := bufio.NewReader(c.connection)

			for {
				if content, err := reader.ReadBytes(0xFF); err == nil {
					// check for a resonse with a socket (that is not an ACK)
					if response, socket := parseViscaResponse(content); response == RESP_COMPLETION || response == RESP_COMMAND_CANCELLED || response == RESP_NO_SOCKET || response == RESP_COMMAND_NOT_EXECUTABLE {
						c.channels[socket] <- content
					} else {
						c.defaultChannel <- content
					}
				}
			}
		}()

		return nil
	}
}

func (c *Camera) Disconnect() error {
	return c.connection.Close()
}

const BUFFER_LEN = 32

type (
	viscaResponse uint8
	viscaSocket   byte
)

const (
	RESP_ACK viscaResponse = iota
	RESP_COMPLETION
	RESP_SYNTAX_ERROR
	RESP_COMMAND_BUFFER_FULL
	RESP_COMMAND_CANCELLED
	RESP_NO_SOCKET
	RESP_COMMAND_NOT_EXECUTABLE
	RESP_INQUIRY
)

func parseViscaResponse(msg []byte) (viscaResponse, viscaSocket) {
	var response viscaResponse
	socket := viscaSocket(msg[1] % 16)

	switch msg[1] >> 4 {
	case 4:
		response = RESP_ACK
	case 5:
		if len(msg) == 3 {
			response = RESP_COMPLETION
		} else {
			response = RESP_INQUIRY
		}
	case 6:
		switch msg[2] {
		case 2:
			socket = math.MaxUint8
			response = RESP_SYNTAX_ERROR
		case 3:
			socket = math.MaxUint8
			response = RESP_COMMAND_BUFFER_FULL
		case 4:
			response = RESP_COMMAND_CANCELLED
		case 5:
			response = RESP_NO_SOCKET
		case 0x41:
			response = RESP_COMMAND_NOT_EXECUTABLE
		}
	}

	return response, socket
}

func getResponse(ch chan []byte) ([]byte, error) {
	select {
	case response := <-ch:
		return response, nil
	case <-time.After(10 * time.Second):
		return nil, fmt.Errorf("timed out waiting for channel message")
	}
}

func (c *Camera) sendCommand(b []byte) (chan []byte, error) {
	c.mu.Lock()

	if _, err := c.connection.Write(b); err != nil {
		c.mu.Unlock()
		return nil, err
	} else {
		response, err := getResponse(c.defaultChannel)
		c.mu.Unlock()

		if err != nil {
			return nil, err

			// if it was an acknowledge, wait for the completion message
		} else {
			responseChannel := make(chan []byte)

			go func() {
				defer close(responseChannel)

				if resp, socket := parseViscaResponse(response); resp == RESP_ACK {
					if response, err := getResponse(c.channels[socket]); err != nil {
					} else {
						responseChannel <- response
					}
				} else {
					responseChannel <- response
				}
			}()

			return responseChannel, nil
		}
	}
}
