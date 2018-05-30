package obsws

import (
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

var messageID = 0

// Client is the interface to obs-websocket.
// Client{Host: "localhost", Port: 4444} will probably work if you haven't configured OBS.
type Client struct {
	Host            string                   // Host (probably "localhost").
	Port            int                      // Port (OBS default is 4444).
	Password        string                   // Password (OBS default is "").
	conn            *websocket.Conn          // Underlying connection to OBS.
	id              int                      // Counter for creating message IDs.
	responseTimeout time.Duration            // Time to keep unhandled responses.
	arrivalTimes    map[string]time.Time     // Arrival time of each response.
	requestTypes    map[string]string        // Mapping of sent requests to their types.
	handlers        map[string]func(e Event) // Event handlers.
	respQ           chan Response            // Queue of received responses.
}

// SetResponseTimeout sets the number of seconds before a response is discarded.
// A value of 0 indicates no timeout.
func (c *Client) SetResponseTimeout(seconds int) {
	c.responseTimeout = time.Duration(time.Duration(seconds) * time.Second)
}

// init prepares the client's internal fields.
func (c *Client) init() {
	c.respQ = make(chan Response)
	c.requestTypes = make(map[string]string)
	c.handlers = make(map[string]func(Event))
	c.arrivalTimes = make(map[string]time.Time)
}

// poll listens for responses/events. This function blocks forever.
func (c *Client) poll() {
	logger.Debug("started polling")

	for {
		m := make(map[string]interface{})

		if err := c.conn.ReadJSON(&m); err != nil {
			logger.Warning("read from WS:", err)
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
	messageID++
	return strconv.Itoa(messageID)
}
