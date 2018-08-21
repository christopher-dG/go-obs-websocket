package obsws

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/gorilla/websocket"
)

// Connect opens a WebSocket connection and authenticates if necessary.
func (c *Client) Connect() error {
	c.handlers = make(map[string]func(Event))
	c.respQ = make(chan map[string]interface{}, bufferSize)

	conn, err := connectWS(c.Host, c.Port)
	if err != nil {
		return err
	}
	c.conn = conn

	// We can't use SendReceive yet because we haven't started polling.

	reqGAR := NewGetAuthRequiredRequest()
	if err = c.conn.WriteJSON(reqGAR); err != nil {
		return err
	}

	respGAR := &GetAuthRequiredResponse{}
	if err = c.conn.ReadJSON(respGAR); err != nil {
		return err
	}

	if !respGAR.AuthRequired {
		logger.Info("logged in (no authentication required)")
		go c.poll()
		return nil
	}

	auth := getAuth(c.Password, respGAR.Salt, respGAR.Challenge)
	logger.Debug("auth:", auth)

	reqA := NewAuthenticateRequest(auth)
	if err = c.conn.WriteJSON(reqA); err != nil {
		return err
	}

	respA := &AuthenticateResponse{}
	if err = c.conn.ReadJSON(respA); err != nil {
		return err
	}
	if respA.Status() != "ok" {
		return errors.New(respA.Error())
	}

	logger.Info("logged in (authentication successful)")
	go c.poll()
	return nil
}

// Disconnect closes the WebSocket connection.
func (c *Client) Disconnect() error {
	c.active = false
	if err := c.conn.Close(); err != nil {
		return err
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
