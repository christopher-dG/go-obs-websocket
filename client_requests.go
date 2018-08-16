package obsws

import (
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

const noID = "noID"

// SendRequest sends a request to the WebSocket server.
// It's not recommended to use this directly, use requests' Send functions instead.
func (c *Client) SendRequest(req Request) (chan Response, error) {
	future := make(chan Response)
	if err := c.conn.WriteJSON(req); err != nil {
		return nil, errors.Wrapf(err, "write %s", req.Type())
	}

	var key string
	if c.noIDMode {
		key = noID
	} else {
		key = req.ID()
	}
	c.requestTypes[key] = req.Type()

	go func() { future <- c.waitResponse(req) }()
	return future, nil
}

// waitResponse waits until a response matching the request is found.
func (c *Client) waitResponse(req Request) Response {
	for {
		resp := <-c.respQ
		if c.noIDMode || resp.ID() == req.ID() {
			logger.Debug("received response", resp.ID())
			return resp
		}

		if c.responseTimeout > 0 && time.Since(c.arrivalTimes[resp.ID()]) > c.responseTimeout {
			delete(c.arrivalTimes, resp.ID())
		} else {
			c.respQ <- resp
		}

		time.Sleep(time.Millisecond * 50)
	}
}

// handleResponse sends a response into the queue.
func (c *Client) handleResponse(m map[string]interface{}) {
	var key string
	if c.noIDMode {
		key = noID
	} else {
		key = m["message-id"].(string)
	}

	respType := c.requestTypes[key]
	if respType == "" {
		logger.Warning("no requestTypes entry for message", key)
		return
	}
	delete(c.requestTypes, key)

	resp := respMap[respType]
	if resp == nil {
		logger.Warning("unknown response type", respType)
		return
	}

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		ZeroFields: true, // TODO: Is this actually working?
		TagName:    "json",
		Result:     resp,
	})
	if err != nil {
		logger.Warning("initializing decoder:", err)
		return
	}

	if err = decoder.Decode(m); err != nil {
		logger.Warningf("unmarshalling map -> %T: %v", resp, err)
		return
	}

	c.arrivalTimes[resp.ID()] = time.Now()
	c.respQ <- derefResponse(resp)
}
