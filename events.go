package obsws

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#events
type event struct {
	UpdateType     string `json:"update-type"`
	StreamTimecode string `json:"stream-timecode"`
	RecTimecode    string `json:"rec-timecode"`
}

type sceneNameSourcesEvent struct {
	SceneName string   `json:"scene-name"`
	Sources   []string `json:"sources"`
	event
}
