package obsws

import (
	"errors"
)

// ErrUnknownEventType is returned when a handler is added for an unknown event.
var ErrUnknownEventType = errors.New("unknown event type")

// AddEventHandler adds a handler function for a given event type.
func (c *Client) AddEventHandler(eventType string, handler func(Event)) error {
	if eventMap[eventType] == nil {
		return ErrUnknownEventType
	}
	c.handlers[eventType] = handler
	return nil
}

// RemoveEventHandler removes the handler for a given event type.
func (c *Client) RemoveEventHandler(eventType string) {
	delete(c.handlers, eventType)
}

// handleEvent runs an event's handler if it exists.
func (c *Client) handleEvent(m map[string]interface{}) {
	t := m["update-type"].(string)

	eventFn, ok := eventMap[t]
	if !ok {
		logger.Warning("unknown event type:", m["update-type"])
		return
	}
	event := eventFn()

	handler, ok := c.handlers[t]
	if !ok {
		return
	}

	if err := mapToStruct(m, event); err != nil {
		logger.Warning("event handler failed:", err)
		return
	}

	go handler(derefEvent(event))
}
