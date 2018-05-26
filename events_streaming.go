package obsws

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#streaming

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#streamstarting
type streamStartingEvent struct {
	PreviewOnly bool `json:"preview-only"`
	event
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#streamstarted
type streamStartedEvent event

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#streamstopping
type streamStoppingEvent streamStartingEvent

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#streamstopped
type streamStoppedEvent event

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#streamstatus
type streamStatusEvent struct {
	Streaming        bool    `json:"streaming"`
	Recording        bool    `json:"recording"`
	PreviewOnly      bool    `json:"preview-only"`
	BytesPerSec      int     `json:"bytes-per-sec"`
	KBitsPerSec      int     `json:"kbits-per-sec"`
	Strain           float64 `json:"strain"`
	TotalStreamTime  int     `json:"total-stream-time"`
	NumTotalFrames   int     `json:"num-total-frames"`
	NumDroppedFrames int     `json:"num-dropped-frames"`
	FPS              float64 `json:"fps"`
	event
}
