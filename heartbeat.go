package obsws

import (
	"github.com/pkg/errors"
)

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setheartbeat
type setHeartbeatRequest struct {
	Enable bool `json:"enable"`
	request
}

// SetHeartbeat enables or disables sending of the Heartbeat event.
func (c *client) SetHeartbeat(enable bool) error {
	reqH := setHeartbeatRequest{
		Enable: enable,
		request: request{
			MessageID:   c.getMessageID(),
			RequestType: "SetHeartbeat",
		},
	}
	if err := c.conn.WriteJSON(reqH); err != nil {
		return errors.Wrap(err, "write SetHeartbeat")
	}

	logger.Infof("set heartbeat to %t", enable)
	return nil
}
