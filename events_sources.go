package obsws

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#sources

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#sourceorderchanged
type sourceOrderChangedEvent struct {
	Name    string   `json:"name"`
	Sources []string `json:"sources"`
	event
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#sceneitemadded
type sceneItemAddedEvent struct {
	SceneName string `json:"scene-name"`
	ItemName  string `json:"item-name"`
	event
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#sceneitemremoved
type sceneItemRemovedEvent sceneItemAddedEvent

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#sceneitemvisibilitychanged
type sceneItemVisibilityChanged struct {
	ItemVisible bool `json:"item-visible"`
	sceneItemAddedEvent
}
