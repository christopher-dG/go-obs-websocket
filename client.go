package obsws

import (
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
)

const bufferSize = 100

var (
	receiveTimeout = time.Duration(0)
	messageID      = 0
	lock           = sync.Mutex{}
)

// Client is the interface to obs-websocket.
// Client{Host: "localhost", Port: 4444} will probably work if you haven't configured OBS.
type Client struct {
	Host           string                      // Host (probably "localhost").
	Port           int                         // Port (OBS default is 4444).
	Password       string                      // Password (OBS default is "").
	conn           *websocket.Conn             // Underlying connection to OBS.
	Logger         *log.Logger                 // Logger to use for most messages
	receiveTimeout time.Duration               // Maximum blocking time for receiving request responses
	connected      bool                        // True until Disconnect is called.
	handlers       map[string]func(e Event)    // Event handlers.
	respQ          chan map[string]interface{} // Queue of received responses.
}

// poll listens for responses/events.
// This function blocks until Disconnect is called.
func (c *Client) poll() {
	c.Logger.Println("started polling")

	for c.connected {
		m := make(map[string]interface{})
		if err := c.conn.ReadJSON(&m); err != nil {
			if !c.connected {
				return
			} else if websocket.IsUnexpectedCloseError(err) {
				c.Disconnect()
			}
			Logger.Println("read from WS:", err)
			continue
		}

		if _, ok := m["message-id"]; ok {
			c.handleResponse(m)
		} else {
			c.handleEvent(m)
		}
	}
}

// Connected returns wheter or not the client is connected.
func (c Client) Connected() bool {
	return c.connected
}

// SetReceiveTimeout sets the maximum blocking time for receiving request responses.
// If set to 0 (the default), there is no timeout.
func SetReceiveTimeout(timeout time.Duration) {
	receiveTimeout = timeout
}

// GetMessageID generates a string that the client has not yet used.
func GetMessageID() string {
	lock.Lock()
	messageID++
	id := strconv.Itoa(messageID)
	lock.Unlock()
	return id
}

// mapToStruct serializes a map into a struct.
func mapToStruct(data map[string]interface{}, dest interface{}) error {
	// func mapToStruct(resp Response, dest interface{}) error {
	// data := resp.(map[string])
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName: "json",
		Result:  dest,
	})
	if err != nil {
		Logger.Println("initializing decoder:", err)
		return err
	}
	if err = decoder.Decode(data); err != nil {
		Logger.Printf("unmarshalling map -> %T: %v", dest, err)
		Logger.Printf("input: %#v\n", data)
		return err
	}
	return nil
}
