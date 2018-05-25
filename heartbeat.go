package obsws

import (
	"github.com/pkg/errors"
)

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setheartbeat
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#heartbeat

type setHeartbeatRequest struct {
	MessageID   string `json:"message-id"`
	RequestType string `json:"request-type"`
	Enable      bool   `json:"enable"`
}

func (r *setHeartbeatRequest) ID() string {
	return r.MessageID
}

type heartbeatResponse struct {
	MessageID         string `json:"message-id"`
	Status            string `json:"status"`
	Error             string `json:"error"`
	Pulse             bool   `json:"pulse"`
	CurrentProfile    string `json:"current-profile"`
	CurrentScene      string `json:"current-scene"`
	Streaming         bool   `json:"streaming"`
	TotalStreamTime   int    `json:"total-stream-time"`
	TotalStreamBytes  int    `json:"total-stream-bytes"`
	TotalStreamFrames int    `json:"total-stream-frames"`
	Recording         int    `json:"recording"`
	TotalRecordTime   int    `json:"total-record-time"`
	TotalRecordBytes  int    `json:"total-record-bytes"`
	TotalRecordFrames int    `json:"total-record-frames"`
}

func (r *heartbeatResponse) ID() string {
	return r.MessageID
}

// SetHeartbeat enables or disables sending of the Heartbeat event.
func (c *client) SetHeartbeat(enable bool) error {
	reqH := setHeartbeatRequest{
		MessageID:   c.getMessageID(),
		RequestType: "SetHeartbeat",
		Enable:      enable,
	}
	if err := c.conn.WriteJSON(reqH); err != nil {
		return errors.Wrap(err, "write SetHeartbeat")
	}

	logger.Infof("set heartbeat to %t", enable)
	return nil
}
