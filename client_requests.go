package obsws

import (
	"errors"
	"time"
)

const interval = time.Millisecond * 50

var (
	ErrNotConnected   = errors.New("not connected")
	ErrReceiveTimeout = errors.New("receive timed out")
)

// sendRequest sends a request to the WebSocket server.
func (c *Client) sendRequest(req Request) (chan map[string]interface{}, error) {
	if !c.active {
		return nil, ErrNotConnected
	}
	future := make(chan map[string]interface{})
	if err := c.conn.WriteJSON(req); err != nil {
		return nil, err
	}
	logger.Debug("sent request", req.ID())
	go func() { future <- c.receive(req.ID()) }()
	return future, nil
}

// receive waits until a response matching the given ID arrives.
func (c *Client) receive(id string) map[string]interface{} {
	for {
		resp := <-c.respQ
		if resp["message-id"] == id {
			logger.Debug("received response", resp["message-id"])
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
