package obsws

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

// Connect opens a WebSocket connection and authenticates if necessary.
func (c *Client) Connect() error {
	c.init()
	conn, err := connectWS(c.Host, c.Port)
	if err != nil {
		return err
	}
	c.conn = conn

	// We can't use SendRequest yet because we haven't started polling.

	reqGAR := GetAuthRequiredRequest{
		ID_:   getMessageID(),
		Type_: "GetAuthRequired",
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
		go c.poll()
		return nil
	}

	auth := getAuth(c.Password, respGAR.Salt, respGAR.Challenge)
	logger.Debug("auth:", auth)

	reqA := AuthenticateRequest{
		Auth: auth,
		_request: _request{
			ID_:   getMessageID(),
			Type_: "Authenticate",
		},
	}
	if err = c.conn.WriteJSON(reqA); err != nil {
		return errors.Wrap(err, "write Authenticate")
	}

	respA := &AuthenticateResponse{}
	if err = c.conn.ReadJSON(respA); err != nil {
		return errors.Wrap(err, "read Authenticate")
	}
	if respA.Status() != "ok" {
		return errors.Errorf("login failed: %s", respA.Error())
	}

	logger.Info("logged in (authentication successful)")
	go c.poll()
	return nil
}

// Disconnect closes the WebSocket connection.
func (c *Client) Disconnect() error {
	c.active = false
	if err := c.conn.Close(); err != nil {
		return errors.Wrap(err, "logout failed")
	}
	return nil
}

// connectWS opens the WebSocket connection.
func connectWS(host string, port int) (*websocket.Conn, error) {
	url := fmt.Sprintf("ws://%s:%d", host, port)
	logger.Debug("connecting to", url)
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// getAuth computes the auth challenge response.
func getAuth(password, salt, challenge string) string {
	sha := sha256.Sum256([]byte(password + salt))
	b64 := base64.StdEncoding.EncodeToString([]byte(sha[:]))

	sha = sha256.Sum256([]byte(b64 + challenge))
	b64 = base64.StdEncoding.EncodeToString([]byte(sha[:]))

	return b64
}
