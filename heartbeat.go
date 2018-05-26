package obsws

import (
	"github.com/pkg/errors"
)

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setheartbeat
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#heartbeat

type setHeartbeatRequest struct {
	Enable bool `json:"enable"`
	request
}

type heartbeatResponse struct {
	Pulse             bool   `json:"pulse"`
	CurrentProfile    string `json:"current-profile"`
	CurrentScene      string `json:"current-scene"`
	Streaming         bool   `json:"streaming"`
	TotalStreamTime   int    `json:"total-stream-time"`
	TotalStreamBytes  int    `json:"total-stream-bytes"`
	TotalStreamFrames int    `json:"total-stream-frames"`
	Recording         bool   `json:"recording"`
	TotalRecordTime   int    `json:"total-record-time"`
	TotalRecordBytes  int    `json:"total-record-bytes"`
	TotalRecordFrames int    `json:"total-record-frames"`
	response
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
