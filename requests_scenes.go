package obsws

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#scenes-1

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setcurrentscene
type setCurrentSceneRequest struct {
	SceneName string `json:"scene-name"`
	request
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setcurrentscene
type setCurrentSceneResponse response

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getcurrentscene
type getCurrentSceneRequest request

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getcurrentscene
type getCurrentSceneResponse struct {
	Name    string      `json:"name"`
	Sources interface{} `json:"sources"` // TODO: Source|Array.
	response
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getscenelist
type getSceneListRequest request

//https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getscenelist
type getSceneListResponse struct {
	CurrentScene string      `json:"current-scene"`
	Scenes       interface{} `json:"scenes"` // TODO: Scene|Array.
	response
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setsceneitemorder
type setSceneItemOrderRequest struct { // Unreleased.
	Scene     string      `json:"scene"`
	Items     interface{} `json:"items"`
	ItemsID   int         `json:"items[].id"`
	ItemsName string      `json:"items[].name"`
}
