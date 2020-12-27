package obsws

import (
	"errors"
	"time"
)

const interval = time.Millisecond * 50

var (
	// ErrNotConnected is returned when a request is sent by a client which is not connected.
	ErrNotConnected = errors.New("not connected")
	// ErrReceiveTimeout is returned when a response takes too long to arrive.
	ErrReceiveTimeout = errors.New("receive timed out")
)

// SendRequest sends a request to the WebSocket server.
func (c *Client) SendRequest(req Request) (chan map[string]interface{}, error) {
	if !c.connected {
		return nil, ErrNotConnected
	}
	future := make(chan map[string]interface{})
	if err := c.conn.WriteJSON(req); err != nil {
		return nil, err
	}
	c.logger.Println("sent request", req.ID())
	go func() { future <- c.receive(req.ID()) }()
	return future, nil
}

// receive waits until a response matching the given ID arrives.
func (c *Client) receive(id string) map[string]interface{} {
	for {
		resp := <-c.respQ
		if resp["message-id"] == id {
			c.logger.Println("received response", resp["message-id"])
			return resp
		}
		c.respQ <- resp
		time.Sleep(interval)
	}
}

// handleResponse sends a response into the queue.
func (c *Client) handleResponse(m map[string]interface{}) {
	c.respQ <- m
}
