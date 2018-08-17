package obsws

import (
	"time"
)

const interval = time.Millisecond * 50

// sendRequest sends a request to the WebSocket server.
func (c *Client) sendRequest(req Request) (chan map[string]interface{}, error) {
	future := make(chan map[string]interface{})
	if err := c.conn.WriteJSON(req); err != nil {
		return nil, err
	}
	go func() { future <- c.waitResponse(req) }()
	return future, nil
}

// waitResponse waits until a response matching the request is found.
func (c *Client) waitResponse(req Request) map[string]interface{} {
	for {
		resp := <-c.respQ
		id := resp["message-id"]
		if resp["message-id"] == req.ID() {
			logger.Debug("received response", id)
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
