package obsws

import (
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

// AddEventHandler adds a handler function for a given event type.
func (c *Client) AddEventHandler(eventType string, handler func(Event)) error {
	if eventMap[eventType] == nil {
		return errors.Errorf("unknown event type %s", eventType)
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
	event := eventMap[m["update-type"].(string)]
	if event == nil {
		logger.Warning("unknown event type", m["update-type"])
		return
	}

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		ZeroFields: true, // TODO: Is this actually working?
		TagName:    "json",
		Result:     event,
	})
	if err != nil {
		logger.Warning("initializing decoder", err)
		return
	}

	if err = decoder.Decode(m); err != nil {
		logger.Warningf("unmarshalling map -> %T: %v", event, err)
		return
	}

	handler := c.handlers[event.Type()]
	if handler != nil {
		go handler(derefEvent(event))
	}
}
