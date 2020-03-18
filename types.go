package obsws

type OBSStats struct {
	FPS                float64 `json:"fps"`
	RenderTotalFrames  int     `json:"render-total-frames"`
	RenderMissedFrames int     `json:"render-missed-frames"`
	OutputTotalFrames  int     `json:"output-total-frames"`
	OutputMissedFrames int     `json:"output-missed-frames"`
	AverageFrameTime   float64 `json:"average-frame-time"`
	CPUUsage           float64 `json:"cpu-usage"`
	MemoryUsage        float64 `json:"memory-usage"`
	FreeDiskSpace      float64 `json:"free-disk-space`
}

type Scene struct {
	Name    string       `json:"name"`
	Sources []*SceneItem `json:"sources"`
}

type SceneItem struct {
	CY              int          `json:"cy"`
	CX              int          `json:"cx"`
	Name            string       `json:"name"`
	ID              int          `json:"id"`
	Render          bool         `json:"render"` // Visible or not
	Locked          bool         `json:"locked"`
	SourceCX        int          `json:"source_cx"`
	SourceCY        int          `json:"source_cy"`
	Type            string       `json:"type"` // One of: "input", "filter", "transition", "scene" or "unknown"
	Volume          int          `json:"volume"`
	X               int          `json:"x"`
	Y               int          `json:"y"`
	ParentGroupName string       `json:"parentGroupName,omitempty"` // Name of the item's parent (if this item belongs to a group)
	GroupChildren   []*SceneItem `json:"groupChildren"`             // List of children (if this item is a group)
}

type SceneItemTransform struct {
	// "{int} `position.x` The x position of the scene item from the left.",
	// "{int} `position.y` The y position of the scene item from the top.",
	// "{int} `position.alignment` The point on the scene item that the item is manipulated from.",
	// "{double} `rotation` The clockwise rotation of the scene item in degrees around the point of alignment.",
	// "{double} `scale.x` The x-scale factor of the scene item.",
	// "{double} `scale.y` The y-scale factor of the scene item.",
	// "{int} `crop.top` The number of pixels cropped off the top of the scene item before scaling.",
	// "{int} `crop.right` The number of pixels cropped off the right of the scene item before scaling.",
	// "{int} `crop.bottom` The number of pixels cropped off the bottom of the scene item before scaling.",
	// "{int} `crop.left` The number of pixels cropped off the left of the scene item before scaling.",
	// "{bool} `visible` If the scene item is visible.",
	// "{bool} `locked` If the scene item is locked in position.",
	// "{String} `bounds.type` Type of bounding box. Can be \"OBS_BOUNDS_STRETCH\", \"OBS_BOUNDS_SCALE_INNER\", \"OBS_BOUNDS_SCALE_OUTER\", \"OBS_BOUNDS_SCALE_TO_WIDTH\", \"OBS_BOUNDS_SCALE_TO_HEIGHT\", \"OBS_BOUNDS_MAX_ONLY\" or \"OBS_BOUNDS_NONE\".",
	// "{int} `bounds.alignment` Alignment of the bounding box.",
	// "{double} `bounds.x` Width of the bounding box.",
	// "{double} `bounds.y` Height of the bounding box.",
	// "{int} `sourceWidth` Base width (without scaling) of the source",
	// "{int} `sourceHeight` Base source (without scaling) of the source",
	// "{double} `width` Scene item width (base source width multiplied by the horizontal scaling factor)",
	// "{double} `height` Scene item height (base source height multiplied by the vertical scaling factor)",
	// "{String (optional)} `parentGroupName` Name of the item's parent (if this item belongs to a group)",
	// "{Array<SceneItemTransform> (optional)} `groupChildren` List of children (if this item is a group)"
}
