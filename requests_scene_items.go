package obsws

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#scene-items

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getsceneitemproperties
type getSceneItemPropertiesRequest struct {
	Scene    string `json:"scene"`
	ItemID   string `json:"item.id"`
	ItemName string `json:"item.name"`
	request
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getsceneitemproperties
type getSceneItemPropertiesResponse struct {
	Scene                 string  `json:"scene"`
	ItemName              string  `json:"item.name"`
	ItemID                string  `json:"item.id"`
	ItemPositionX         int     `json:"item.position.x"`
	ItemPositionY         int     `json:"item.position.y"`
	ItemPositionAlignment int     `json:"item.position.alignment"`
	ItemRotation          float64 `json:"item.rotation"`
	ItemScaleX            float64 `json:"item.scale.x"`
	ItemScaleY            float64 `json:"item.scale.y"`
	ItemCropTop           int     `json:"item.crop.top"`
	ItemCropRigt          int     `json:"item.crop.right"`
	ItemCropBottom        int     `json:"item.crop.bottom"`
	ItemCropLeft          int     `json:"item.crop.left"`
	ItemVisible           bool    `json:"item.visible"`
	ItemLocked            bool    `json:"item.locked"`
	ItemBoundsType        string  `json:"item.bounds.type"`
	ItemBoundsAlignment   int     `json:"item.bounds.alignment"`
	ItemBoundsX           float64 `json:"item.bounds.x"`
	ItemBoundsY           float64 `json:"item.bounds.y"`
	response
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setsceneitemproperties
type setSceneItemPropertiesRequest struct {
	Scene                 string  `json:"scene"`
	ItemName              string  `json:"item.name"`
	ItemID                string  `json:"item.id"`
	ItemPositionX         int     `json:"item.position.x"`
	ItemPositionY         int     `json:"item.position.y"`
	ItemPositionAlignment int     `json:"item.position.alignment"`
	ItemRotation          float64 `json:"item.rotation"`
	ItemScaleX            float64 `json:"item.scale.x"`
	ItemScaleY            float64 `json:"item.scale.y"`
	ItemCropTop           int     `json:"item.crop.top"`
	ItemCropRigt          int     `json:"item.crop.right"`
	ItemCropBottom        int     `json:"item.crop.bottom"`
	ItemCropLeft          int     `json:"item.crop.left"`
	ItemVisible           bool    `json:"item.visible"`
	ItemLocked            bool    `json:"item.locked"`
	ItemBoundsType        string  `json:"item.bounds.type"`
	ItemBoundsAlignment   int     `json:"item.bounds.alignment"`
	ItemBoundsX           float64 `json:"item.bounds.x"`
	ItemBoundsY           float64 `json:"item.bounds.y"`
	request
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setsceneitemproperties
type setSceneItemPropertiesResponse response

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#resetsceneitem
type resetSceneItemRequest struct {
	SceneName string `json:"scene-name"`
	Item      string `json:"item"`
	request
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#resetsceneitem
type resetSceneItemResponse response

/*
* Deprecated:
*   SetSceneItemRender
*   SetSceneItemPosition
*   SetSceneItemTransform
*   SetSceneItemCrop
 */
