package obsws

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#authentication
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getauthrequired
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#authenticate

// Connect makes a WebSocket connection and authenticates if necessary.
func (c *client) Connect() error {
	conn, err := connect(c.Host, c.Port)
	if err != nil {
		return err
	}
	c.conn = conn

	reqGAR := getAuthRequiredRequest{
		MessageID:   c.getMessageID(),
		RequestType: "GetAuthRequired",
	}

	if err = c.conn.WriteJSON(reqGAR); err != nil {
		return errors.Wrap(err, "write Authenticate")
	}

	respGAR := &getAuthRequiredResponse{}
	if err = c.conn.ReadJSON(respGAR); err != nil {
		return errors.Wrap(err, "read GetAuthRequired")
	}

	if !respGAR.AuthRequired {
		logger.Info("no authentication required")
		return nil
	}

	auth := getAuth(c.Password, respGAR.Salt, respGAR.Challenge)
	logger.Debugf("auth: %s", auth)

	reqA := authenticateRequest{
		Auth: auth,
		request: request{
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

// Close closes the WebSocket connection.
func (c *client) Close() error {
	return c.conn.Close()
}

func connect(host string, port int) (*websocket.Conn, error) {
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
