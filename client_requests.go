package obsws

import (
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

// SendRequest sends a request to the WebSocket server.
// The return value is a channel to which a response will be written when it
// is received. Note that to access any fields that are not defined in the base
// response, a type assertion is required.
// The following pattern is recommended:
//
//     future, err := client.SendRequest(NewGetStreamingStatusRequest())
//     if err != nil { /* ... */  }
//     status := (<-future).(GetStreamingStatusResponse)
func (c *Client) SendRequest(req Request) (chan Response, error) {
	future := make(chan Response)
	if err := c.conn.WriteJSON(req); err != nil {
		return nil, errors.Wrapf(err, "write %s", req.Type())
	}
	c.requestTypes[req.ID()] = req.Type()
	go func() { future <- c.waitResponse(req, true) }()
	return future, nil
}

// SendRequestNoID is the same as SendRequest, except that it does not guarantee
// that the returned reponse's ID is the same as that of the request.
func (c *Client) SendRequestNoID(req Request) (chan Response, error) {
	future := make(chan Response)
	if err := c.conn.WriteJSON(req); err != nil {
		return nil, errors.Wrapf(err, "write %s", req.Type())
	}
	c.requestTypes[req.ID()] = req.Type()
	go func() { future <- c.waitResponse(req, false) }()
	return future, nil
}

// waitResponse waits until a response matching the request is found.
func (c *Client) waitResponse(req Request, checkID bool) Response {
	for {
		resp := <-c.respQ
		if !checkID || resp.ID() == req.ID() {
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
	mID := m["message-id"].(string)
	respType := c.requestTypes[mID]
	if respType == "" {
		logger.Warning("no requestTypes entry for message", mID)
		return
	}
	delete(c.requestTypes, mID)

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
