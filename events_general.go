package obsws

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#general

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#heartbeat
type heartbeatEvent struct {
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
	event
}
