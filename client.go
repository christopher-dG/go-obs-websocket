package obsws

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

// Client is the interface to obs-websocket.
// Client{Host: "localhost", Port: 4444} will probably work if you haven't configured OBS.
type Client struct {
	Host       string // Host (probably "localhost").
	Port       int    // Port (OBS default is 4444).
	Password   string // Password (OBS default is "").
	conn       *websocket.Conn
	id         int
	respQ      []response
	timeout    int
	timeoutSet bool
}

// Connect opens a WebSocket connection and authenticates if necessary.
func (c *Client) Connect() error {
	conn, err := connectWS(c.Host, c.Port)
	if err != nil {
		return err
	}
	c.conn = conn

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
		logger.Info("no authentication required")
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

	logger.Info("logged in")
	return nil
}

// Disconnect closes the WebSocket connection.
func (c *Client) Disconnect() error {
	return c.conn.Close()
}

// SetTimeout sets a timeout in seconds for receiving responses.
// A value of 0 indicates that responses must be received instantly.
func (c *Client) SetTimeout(seconds int) {
	c.timeoutSet = true
	c.timeout = seconds
}

// DisableTimeout disables the response timeout.
func (c *Client) DisableTimeout() {
	c.timeoutSet = false
}

// SendRequest sends a request to the WebSocket server.
func (c *Client) SendRequest(req request) (chan response, error) {
	future := make(chan response)
	if err := c.conn.WriteJSON(req); err != nil {
		return nil, errors.Wrapf(err, "write %s", req.Type())
	}
	go func() { future <- c.waitResponse(req) }()
	return future, nil
}

func (c *Client) waitResponse(req request) response {
	start := time.Now()
	for {
		for i, resp := range c.respQ {
			if resp.(message).ID() == req.(message).ID() {
				c.respQ = append(c.respQ[:i], c.respQ[i+1:]...)
				return resp
			}
		}

		resp, err := c.receiveResponse(req)
		if err != nil {
			logger.Warning(err)
		}

		if resp.(message).ID() == req.(message).ID() {
			return resp
		}
		c.respQ = append(c.respQ, resp)

		if c.timeoutSet && time.Since(start).Seconds() > float64(c.timeout) {
			logger.Infof("request %s (%s) timed out", req.(message).ID(), req.Type())
			return nil
		}
	}
}

func (c *Client) receiveResponse(req request) (response, error) {
	m := make(map[string]interface{})

	_, bytes, err := c.conn.ReadMessage()
	if err != nil {
		return nil, errors.Wrap(err, "read from WS")
	}

	if err = json.Unmarshal(bytes, &m); err != nil {
		logger.Warningf("unmarshalling JSON -> map failed: %v", err)
		return nil, errors.Wrap(err, "unmarshalling JSON -> map")
	}

	resp, ok := respMap[req.Type()]
	if !ok {
		return nil, errors.Errorf("unknown request type '%s'", req.Type())
	}

	if err = mapstructure.Decode(m, &resp); err != nil {
		return nil, errors.Wrapf(err, "unmarshalling map -> %T", resp)
	}

	return resp, nil
}

func connectWS(host string, port int) (*websocket.Conn, error) {
	url := fmt.Sprintf("ws://%s:%d", host, port)
	logger.Infof("connecting to %s", url)
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func getAuth(password, salt, challenge string) string {
	sha := sha256.Sum256([]byte(password + salt))
	b64 := base64.StdEncoding.EncodeToString([]byte(sha[:]))

	sha = sha256.Sum256([]byte(b64 + challenge))
	b64 = base64.StdEncoding.EncodeToString([]byte(sha[:]))

	return b64
}

func (c *Client) getMessageID() string {
	c.id++
	return strconv.Itoa(c.id)
}
