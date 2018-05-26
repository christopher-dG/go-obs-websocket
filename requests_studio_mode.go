package obsws

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#studio-mode-1

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getstudiomodestatus
type getStudioModeStatusRequest request

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getstudiomodestatus
type getStudioModeStatusResponse struct {
	StudioMode bool `json:"studio-mode"`
	response
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getpreviewscene
type getPreviewSceneRequest request

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getpreviewscene
type getPreviewSceneResponse struct {
	Name    string      `json:"name"`
	Sources interface{} `json:"sources"` // TODO: Source|Array.
	response
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setpreviewscene
type setPreviewSceneRequest struct {
	SceneName string `json:"scene-name"`
	request
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setpreviewscene
type setPreviewSceneResponse response

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#transitiontoprogram
type transitionToProgramRequest struct {
	WithTransition         interface{} `json:"with-transition"` // TODO: Object.
	WithTransitionName     string      `json:"with-transition.name"`
	WithTransitionDuration int         `json:"with-transition.duration"`
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#transitiontoprogram
type transitionToProgramResponse response

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#enablestudiomode
type enableStudioModeRequest request

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#enablestudiomode
type enableStudioModeResponse response

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#disablestudiomode
type disableStudioModeRequest request

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#disablestudiomode
type disableStudioModeResponse response

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#togglestudiomode
type toggleStudioModeRequest request

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#togglestudiomode
type toggleStudioModeResponse response
