package obsws

import (
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
)

const bufferSize = 100

var (
	messageID = 0
	lock      = sync.Mutex{}
)

// Client is the interface to obs-websocket.
// Client{Host: "localhost", Port: 4444} will probably work if you haven't configured OBS.
type Client struct {
	Host     string                      // Host (probably "localhost").
	Port     int                         // Port (OBS default is 4444).
	Password string                      // Password (OBS default is "").
	conn     *websocket.Conn             // Underlying connection to OBS.
	active   bool                        // True until Disconnect is called.
	noIDMode bool                        // When true, don't verify response IDs.
	handlers map[string]func(e Event)    // Event handlers.
	respQ    chan map[string]interface{} // Queue of received responses.
}

// NoIDMode disables response ID checking when set to true.
// This means that there is no guarantee that a response obtained from SendRequest
// corresponds to the request that was sent.
func (c *Client) NoIDMode(enable bool) {
	c.noIDMode = enable
}

// init prepares the client's internal fields.
func (c *Client) init() {
	c.handlers = make(map[string]func(Event))
	c.respQ = make(chan map[string]interface{}, bufferSize)
}

// poll listens for responses/events.
// This function blocks until Disconnect is called.
func (c *Client) poll() {
	c.active = true
	logger.Debug("started polling")

	for c.active {
		m := make(map[string]interface{})
		if err := c.conn.ReadJSON(&m); err != nil {
			if !c.active {
				return
			}
			logger.Warning("read from WS:", err)
			continue
		}

		if _, ok := m["message-id"]; ok {
			c.handleResponse(m)
		} else {
			c.handleEvent(m)
		}
	}
}

// getMessageID generates a string that the client has not yet used.
func getMessageID() string {
	lock.Lock()
	messageID++
	id := strconv.Itoa(messageID)
	lock.Unlock()
	return id
}
