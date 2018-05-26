package obsws

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#transitions-1

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#gettransitionlist
type getTransitionListRequest request

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#gettransitionlist
type getTransitionListResponse struct {
	CurrentTransition string      `json:"current-transition"`
	Transitions       interface{} `json:"transitions"` // TODO: Object|Array.
	TransitionsName   string      `json:"transitions[].name"`
	response
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getcurrenttransition
type getCurrentTransitionRequest request

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getcurrenttransition
type getCurrentTransitionResponse struct {
	Name     string `json:"name"`
	Duration int    `json:"duration"`
	response
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setcurrenttransition
type setCurrentTransitionRequest struct {
	TransitionName string `json:"transition-name"`
	request
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setcurrenttransition
type setCurrentTransitionResponse response

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#settransitionduration
type setTransitionDurationRequest struct {
	Duration int `json:"duration"`
	request
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#settransitionduration
type setTransitionDurationResponse response

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#gettransitionduration
type getTransitionDurationRequest request

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#gettransitionduration
type getTransitionDurationResponse struct {
	TransitionDuration int `json:"transition-duration"`
	response
}
