package obsws

import (
	"time"

	"github.com/mitchellh/mapstructure"
)

const interval = time.Millisecond * 50

// SendRequest sends a request to the WebSocket server.
// It's not recommended to use this directly, use requests' Send functions instead.
func (c *Client) SendRequest(req Request) (chan map[string]interface{}, error) {
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
		if c.noIDMode || resp["message-id"] == req.ID() {
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

// mapToStruct serializes a map into a struct.
func mapToStruct(data map[string]interface{}, dest interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		ZeroFields: true, // TODO: Is this actually working?
		TagName:    "json",
		Result:     dest,
	})
	if err != nil {
		logger.Warning("initializing decoder:", err)
		return err
	}
	if err = decoder.Decode(data); err != nil {
		logger.Warningf("unmarshalling map -> %T: %v", dest, err)
		logger.Debugf("input: %#v\n", data)
		return err
	}
	return nil
}
