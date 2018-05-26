package obsws

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#studio-mode

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#previewscenechanged
type previewSceneChangedEvent sceneNameSourcesEvent

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#studiomodeswitched
type studioModeSwitchedEvent struct {
	NewState bool `json:"new-state"`
	event
}
