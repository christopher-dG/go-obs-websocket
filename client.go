package obsws

import "github.com/gorilla/websocket"

// client is the interface to obs-websocket.
type client struct {
	Host     string // Host (probably "localhost").
	Port     int    // Port (OBS default is 4444).
	Password string // Password (OBS default is "").
	conn     *websocket.Conn
	id       int
}

// NewClient creates a new client. If you haven't configured obs-websocket at
// all, then host should be "localhost", port should be 4444, and password
// should be "".
func NewClient(host string, port int, password string) *client {
	return &client{Host: host, Port: port, Password: password}
}
