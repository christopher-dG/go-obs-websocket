package obsws

type OBSStats struct {
	FPS                float64 `json:"fps"`
	RenderTotalFrames  int     `json:"render-total-frames"`
	RenderMissedFrames int     `json:"render-missed-frames"`
	// "{int} `output-total-frames` Number of frames outputted",
	// "{int} `output-skipped-frames` Number of frames skipped due to encoding lag",
	// "{double} `average-frame-time` Average frame render time (in milliseconds)",
	// "{double} `cpu-usage` Current CPU usage (percentage)",
	// "{double} `memory-usage` Current RAM usage (in megabytes)",
	// "{double} `free-disk-space` Free recording disk space (in megabytes)"
}

type Scene struct {
	Name    string       `json:"name"`
	Sources []*SceneItem `json:"sources"`
}

type SceneItem struct {
	CY int `json:"cy"`
	// "{Number} `cx`",
	// "{String} `name` The name of this Scene Item.",
	// "{int} `id` Scene item ID",
	// "{Boolean} `render` Whether or not this Scene Item is set to \"visible\".",
	// "{Boolean} `locked` Whether or not this Scene Item is locked and can't be moved around",
	// "{Number} `source_cx`",
	// "{Number} `source_cy`",
	// "{String} `type` Source type. Value is one of the following: \"input\", \"filter\", \"transition\", \"scene\" or \"unknown\"",
	// "{Number} `volume`",
	// "{Number} `x`",
	// "{Number} `y`",
	// "{String (optional)} `parentGroupName` Name of the item's parent (if this item belongs to a group)",
	// "{Array<SceneItem> (optional)} `groupChildren` List of children (if this item is a group)"
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
