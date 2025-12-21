// Package visca implements communication with visca-cameras over ip
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
	defaultChannel chan []byte
	socketChannels map[viscaSocket]chan []byte
}

// Connect to the camera
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
		c.socketChannels = make(map[viscaSocket]chan []byte)
		for socket := range viscaSocket(8) {
			c.socketChannels[socket] = make(chan []byte)
		}

		// response reader
		go func() {
			reader := bufio.NewReader(c.connection)

			for {
				if content, err := reader.ReadBytes(0xFF); err == nil {
					// check for a resonse with a socket (that is not an ACK)
					if response, socket := parseViscaResponse(content); response == ResponseCompletion || response == ResponseCommandCancelled || response == ResponseNoSocket || response == ResponseCommandNotExecutable {
						c.socketChannels[socket] <- content
					} else {
						c.defaultChannel <- content
					}
				}
			}
		}()

		return nil
	}
}

// Disconnect from the camera
func (c *Camera) Disconnect() error {
	// TODO: stop receive-goroutine
	return c.connection.Close()
}

type (
	viscaResponse uint8
	viscaSocket   byte
)

const (
	ResponseAcknowledge viscaResponse = iota
	ResponseCompletion
	ResponseSyntaxError
	ResponseCommandBufferFull
	ResponseCommandCancelled
	ResponseNoSocket
	ResponseCommandNotExecutable
	ResponseInquiry
)

// parse a visca response
func parseViscaResponse(msg []byte) (viscaResponse, viscaSocket) {
	var response viscaResponse
	socket := viscaSocket(msg[1] % 16)

	switch msg[1] >> 4 {
	case 4:
		response = ResponseAcknowledge
	case 5:
		if len(msg) == 3 {
			response = ResponseCompletion
		} else {
			socket = math.MaxUint8
			response = ResponseInquiry
		}
	case 6:
		switch msg[2] {
		case 2:
			socket = math.MaxUint8
			response = ResponseSyntaxError
		case 3:
			socket = math.MaxUint8
			response = ResponseCommandBufferFull
		case 4:
			response = ResponseCommandCancelled
		case 5:
			response = ResponseNoSocket
		case 0x41:
			response = ResponseCommandNotExecutable
		}
	}

	return response, socket
}

// wait for a response from the channel and return it, time out after 1 second
func getResponse(ch chan []byte) ([]byte, error) {
	select {
	case response := <-ch:
		return response, nil
	case <-time.After(1 * time.Second):
		return nil, fmt.Errorf("timed out waiting for channel message")
	}
}

// send a command to the camera and get a response
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
		} else {
			responseChannel := make(chan []byte)

			// wait for the responses in the background and send them back through the channel
			go func() {
				defer close(responseChannel)

				// if it was an acknowledge, wait for the completion message
				if resp, socket := parseViscaResponse(response); resp == ResponseAcknowledge {
					if response, err := getResponse(c.socketChannels[socket]); err != nil {
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
