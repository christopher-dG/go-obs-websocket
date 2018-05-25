package obsws

// Client is the interface to OBS's websockets.
type Client struct {
	host     string
	port     int
	password string
}

// NewClient creates a new Client. If you haven't configured obs-websocket at
// all, then host should be "localhost", port should be 4444, and password
// should be "".
func NewClient(host string, port int, password string) *Client {
	return &Client{host: host, port: port, password: password}
}

// Host returns the Client's host.
func (c *Client) Host() string {
	return c.host
}

// Port returns the Client's port.
func (c *Client) Port() int {
	return c.port

}

// Password returns the Client's password.
func (c *Client) Password() string {
	return c.password
}
