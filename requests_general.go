package obsws

import "github.com/pkg/errors"

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#general-1

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getversion
type getVersionRequest request

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getversion
type getVersionResponse struct {
	Version             float64 `json:"version"`
	OBSWebsocketVersion string  `json:"obs-websocket-version"`
	OBSStudioVersion    string  `json:"obs-studio-version"`
	AvailableRequests   string  `json:"available-requests"`
	response
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getauthrequired
type getAuthRequiredRequest request

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getauthrequired
type getAuthRequiredResponse struct {
	AuthRequired bool   `json:"authRequired"`
	Challenge    string `json:"challenge"`
	Salt         string `json:"salt"`
	response
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#authenticate
type authenticateRequest struct {
	Auth string `json:"auth"`
	request
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#authenticate
type authenticateResponse response

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setheartbeat
type setHeartbeatRequest struct {
	Enable bool `json:"enable"`
	request
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setheartbeat
type setHeartbeatResponse response

// SetHeartbeat enables/disables sending of the Heartbeat event.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setheartbeat
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

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setfilenameformatting
type setFileNameFormattingRequest struct {
	FilenameFormatting string `json:"filename-formatting"`
	request
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setfilenameformatting
type setFileNameFormattingResponse response

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getfilenameformatting
type getFileNameFormattingRequest request

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getfilenameformatting
type getFileNameFormattingResponse struct {
	FilenameFormatting string `json:"filename-formatting"`
	response
}
