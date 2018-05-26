package obsws

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#scene-collections

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setcurrentscenecollection
type setCurrentSceneCollectionRequest struct {
	SCName string `json:"sc-name"`
	request
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setcurrentscenecollection
type setCurrentSceneCollectionResponse response

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getcurrentscenecollection
type getCurrentSceneCollectionRequest request

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getcurrentscenecollection
type getCurrentSceneCollectionResponse struct {
	SCName string `json:"sc-name"`
	response
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#listscenecollections
type listSceneCollectionsRequest request

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#listscenecollections
type listSceneCollectionsResponse struct {
	SceneCollections     interface{} `json:"scene-collections"` // TODO: Object|Array.
	SceneCollectionsStar string      `json:"scene-collections.*."`
	response
}
