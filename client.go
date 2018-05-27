package obsws

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

// Client is the interface to obs-websocket.
// Client{Host: "localhost", Port: 4444} will probably work if you haven't configured OBS.
type Client struct {
	Host     string // Host (probably "localhost").
	Port     int    // Port (OBS default is 4444).
	Password string // Password (OBS default is "").
	conn     *websocket.Conn
	id       int
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
