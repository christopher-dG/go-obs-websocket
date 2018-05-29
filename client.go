package obsws

import (
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

// Client is the interface to obs-websocket.
// Client{Host: "localhost", Port: 4444} will probably work if you haven't configured OBS.
type Client struct {
	Host         string // Host (probably "localhost").
	Port         int    // Port (OBS default is 4444).
	Password     string // Password (OBS default is "").
	conn         *websocket.Conn
	id           int
	respQ        []response
	requestTypes map[string]string
}

// Connect opens a WebSocket connection and authenticates if necessary.
func (c *Client) Connect() error {
	c.requestTypes = make(map[string]string)
	conn, err := connectWS(c.Host, c.Port)
	if err != nil {
		return err
	}
	c.conn = conn

	// We can't use SendRequest yet because we aren't polling.

	reqGAR := GetAuthRequiredRequest{
		MessageID:   c.getMessageID(),
		RequestType: "GetAuthRequired",
	}
	if err = c.conn.WriteJSON(reqGAR); err != nil {
		return errors.Wrap(err, "write Authenticate")
	}

	respGAR := &GetAuthRequiredResponse{}
	if err = c.conn.ReadJSON(respGAR); err != nil {
		return errors.Wrap(err, "read GetAuthRequired")
	}

	if !respGAR.AuthRequired {
		logger.Info("logged in (no authentication required)")
		return nil
	}

	auth := getAuth(c.Password, respGAR.Salt, respGAR.Challenge)
	logger.Debugf("auth: %s", auth)

	reqA := AuthenticateRequest{
		Auth: auth,
		_request: _request{
			MessageID:   c.getMessageID(),
			RequestType: "Authenticate",
		},
	}
	if err = c.conn.WriteJSON(reqA); err != nil {
		return errors.Wrap(err, "write Authenticate")
	}

	respA := &AuthenticateResponse{}
	if err = c.conn.ReadJSON(respA); err != nil {
		return errors.Wrap(err, "read Authenticate")
	}
	if respA.Stat() != "ok" {
		return errors.Errorf("login failed: %s", respA.Err())
	}

	logger.Info("logged in (authentication successful)")
	go c.poll()
	return nil
}

// Disconnect closes the WebSocket connection.
func (c *Client) Disconnect() error {
	return c.conn.Close()
}

// SendRequest sends a request to the WebSocket server.
func (c *Client) SendRequest(req request) (chan response, error) {
	future := make(chan response)
	if err := c.conn.WriteJSON(req); err != nil {
		return nil, errors.Wrapf(err, "write %s", req.Type())
	}
	c.requestTypes[req.ID()] = req.Type()
	go func() { future <- c.waitResponse(req) }()
	return future, nil
}

func (c *Client) waitResponse(req request) response {
	for {
		for i, resp := range c.respQ {
			if resp.ID() == req.ID() {
				logger.Debug("found message", resp.ID())
				c.respQ = removeResponses(c.respQ, i)
				return resp
			}
		}
	}
}

func (c *Client) receiveResponse(req request) (response, error) {
	m := make(map[string]interface{})

	if err := c.conn.ReadJSON(&m); err != nil {
		return nil, errors.Wrap(err, "read from WS")
	}

	resp, ok := respMap[req.Type()]
	if !ok {
		return nil, errors.Errorf("unknown request type '%s'", req.Type())
	}

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		ZeroFields: true, // TODO: I don't think this is working.
		Result:     resp,
	})
	if err != nil {
		return nil, errors.Wrap(err, "initializing decoder")
	}

	if err := decoder.Decode(m); err != nil {
		return nil, errors.Wrapf(err, "unmarshalling map -> %T", resp)
	}

	return resp, nil
}

// Listen for responses/events and send them into queues.
// This function blocks forever.
func (c *Client) poll() {
	logger.Debug("started polling")
	for {
		m := make(map[string]interface{})

		if err := c.conn.ReadJSON(&m); err != nil {
			logger.Warning("read from WS:", err)
		}

		if mID, ok := m["message-id"]; ok { // Response.
			respType := c.requestTypes[mID.(string)]
			if respType == "" {
				logger.Warning("no requestTypes entry for message", mID)
				continue
			}
			delete(c.requestTypes, mID.(string))

			resp := respMap[respType]
			if resp == nil {
				logger.Warning("unknown response type", respType)
				continue
			}

			decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				ZeroFields: true, // TODO: I don't think this is working.
				Result:     resp,
			})
			if err != nil {
				logger.Warning("initializing decoder:", err)
				continue
			}

			if err = decoder.Decode(m); err != nil {
				logger.Warningf("unmarshalling map -> %T: %v", resp, err)
				continue
			}

			c.respQ = append(c.respQ, deref(resp))
		} else { // Event.
			logger.Debug("event")
		}
	}
}

func (c *Client) getMessageID() string {
	c.id++
	return strconv.Itoa(c.id)
}
