package obsws

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#transitions

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#switchtransition
type switchTransitionsEvent struct {
	TransitionName string `json:"transition-name"`
	event
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#transitionlistchanged
type transitionListChangedEvent event

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#transitiondurationchanged
type transitionDurationChangedEvent struct {
	NewDuration int `json:"new-duration"`
	event
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#transitionbegin
type transitionBeginEvent struct {
	Name     string `json:"name"`
	Duration int    `json:"duration"`
	event
}
