package obsws

// Event is broadcast by the server to each connected client when a recognized action occurs within OBS.
type Event interface {
	Type() string
	StreamTC() string
	RecTC() string
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#events
type _event struct {
	UpdateType     string `json:"update-type"`
	StreamTimecode string `json:"stream-timecode"`
	RecTimecode    string `json:"rec-timecode"`
}

func (e _event) Type() string { return e.UpdateType }

func (e _event) StreamTC() string { return e.StreamTimecode }

func (e _event) RecTC() string { return e.RecTimecode }
