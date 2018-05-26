package obsws

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#sources-1

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getsourceslist
type getSourcesListRequest request

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getsourceslist
type getSourcesListResponse struct {
	Sources       []interface{} `json:"sources"` // TODO: Array of Objects.
	SourcesName   string        `json:"sources.*.name"`
	SourcesTypeID string        `json:"sources.*.typeId"`
	SourcesType   string        `json:"sources.*.type"`
	response
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getsourcestypeslist
type getSourcesTypesListRequest request

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getsourcestypeslist
type getSourcesTypesListResponse struct {
	IDs                     []interface{} `json:"ids"` // TODO: Array of Objects.
	IDsTypeID               string        `json:"ids.*.typeID"`
	IDsDisplayName          string        `json:"ids.*.displayName"`
	IDsType                 string        `json:"ids.*.type"`
	IDsDefaultSettings      interface{}   `json:"ids.*.defaultSettings"` // TODO: Object.
	IDsCaps                 interface{}   `json:"ids.*.caps"`            // TODO: Object.
	IDsCapsIsAsync          bool          `json:"ids.*.caps.isAsync"`
	IDsCapsHasVideo         bool          `json:"ids.*.caps.hasVideo"`
	IDsCapsHasAudio         bool          `json:"ids.*.caps.hasAudio"`
	IDsCapsCanInteract      bool          `json:"ids.*.caps.canInteract"`
	IDsCapsIsComposite      bool          `json:"ids.*.caps.isComposite"`
	IDsCapsDoNotDuplicate   bool          `json:"ids.*.caps.doNotDuplicate"`
	IDsCapsDoNotSelfMonitor bool          `json:"ids.*.caps.doNotSelfMonitor"`
	response
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getvolume
type getVolumeRequest struct {
	Source string `json:"source"`
	request
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getvolume
type getVolumeResponse struct {
	Name   string  `json:"name"`
	Volume float64 `json:"volume"`
	Mute   bool    `json:"mute"`
	response
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setvolume
type setVolumeRequest struct {
	Source string  `json:"source"`
	Volume float64 `json:"volume"`
	request
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setvolume
type setVolumeResponse response

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getmute
type getMuteRequest struct {
	Source string `json:"source"`
	request
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getmute
type getMuteResponse struct {
	Source string `json:"source"`
	Muted  bool   `json:"muted"`
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setmute
type setMuteRequest struct {
	Source string `json:"source"`
	Mute   bool   `json:"mute"`
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setmute
type setMuteResponse response

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#togglemute
type toggleMuteRequest struct {
	Source string `json:"source"`
	request
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#togglemute
type toggleMuteResponse response

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setsyncoffset
type setSyncOffsetRequest struct {
	Source string `json:"source"`
	Offset int    `json:"offset"`
	request
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setsyncoffset
type setSyncOffsetResponse response

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getsyncoffset
type getSyncOffsetRequest struct {
	Source string `json:"source"`
	request
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getsyncoffset
type getSyncOffsetResponse struct {
	Name   string `json:"name"`
	Offset int    `json:"offset"`
	response
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getsourcesettings
type getSourceSettingsRequest struct {
	SourceName string `json:"sourceName"`
	SourceType string `json:"sourceType"`
	request
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getsourcesettings
type getSourceSettingsResponse struct {
	SourceName     string      `json:"sourceName"`
	SourceType     string      `json:"sourceType"`
	SourceSettings interface{} `json:"sourceSettings"` // TODO: Object.
	response
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setsourcesettings
type setSourceSettingsRequest struct {
	SourceName     string      `json:"sourceName"`
	SourceType     string      `json:"sourceType"`
	SourceSettings interface{} `json:"sourceSettings"` // TODO: Object.
	request
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setsourcesettings
type setSourceSettingsResponse struct {
	SourceName     string      `json:"sourceName"`
	SourceType     string      `json:"sourceType"`
	SourceSettings interface{} `json:"sourceSettings"` // TODO: Object.
	response
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#gettextgdiplusproperties
type getTextGDIPlusPropertiesRequest struct {
	SceneName string `json:"scene-name"`
	Source    string `json:"source"`
	request
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#gettextgdiplusproperties
type getTextGDIPlusPropertiesResponse struct {
	Align           string      `json:"align"`
	BKColor         int         `json:"bk-color"`
	BKOpacity       int         `json:"bk-opacity"`
	Chatlog         bool        `json:"chatlog"`
	ChatlogLines    int         `json:"chatlog_lines"`
	Color           int         `json:"color"`
	Extents         bool        `json:"extents"`
	ExtentsCX       int         `json:"extents_cx"`
	ExtentsCY       int         `json:"extents_cy"`
	File            string      `json:"file"`
	ReadFromFile    bool        `json:"read_from_file"`
	Font            interface{} `json:"font"` // TODO: Object.
	FontFace        string      `json:"font.face"`
	FontFlags       int         `json:"font.flags"`
	FontSize        int         `json:"font.size"`
	FontStyle       string      `json:"font.style"`
	Gradient        bool        `json:"gradient"`
	GradientColor   int         `json:"gradient_color"`
	GradientDir     float64     `json:"gradient_dir"`
	GradientOpacity int         `json:"gradient_opacity"`
	Outline         bool        `json:"outline"`
	OutlineColour   int         `json:"outline_color"`
	OutlineSize     int         `json:"outline_size"`
	OutlineOpacity  int         `json:"outline_opacity"`
	Text            string      `json:"text"`
	VAlign          string      `json:"valign"`
	Vertical        bool        `json:"vertical"`
	Render          bool        `json:"render"`
	response
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#settextgdiplusproperties
type setTextGDIPlusPropertiesRequest struct {
	SceneName       string      `json:"scene-name"`
	Align           string      `json:"align"`
	BKColor         int         `json:"bk-color"`
	BKOpacity       int         `json:"bk-opacity"`
	Chatlog         bool        `json:"chatlog"`
	ChatlogLines    int         `json:"chatlog_lines"`
	Color           int         `json:"color"`
	Extents         bool        `json:"extents"`
	ExtentsCX       int         `json:"extents_cx"`
	ExtentsCY       int         `json:"extents_cy"`
	File            string      `json:"file"`
	ReadFromFile    bool        `json:"read_from_file"`
	Font            interface{} `json:"font"` // TODO: Object.
	FontFace        string      `json:"font.face"`
	FontFlags       int         `json:"font.flags"`
	FontSize        int         `json:"font.size"`
	FontStyle       string      `json:"font.style"`
	Gradient        bool        `json:"gradient"`
	GradientColor   int         `json:"gradient_color"`
	GradientDir     float64     `json:"gradient_dir"`
	GradientOpacity int         `json:"gradient_opacity"`
	Outline         bool        `json:"outline"`
	OutlineColour   int         `json:"outline_color"`
	OutlineSize     int         `json:"outline_size"`
	OutlineOpacity  int         `json:"outline_opacity"`
	Text            string      `json:"text"`
	VAlign          string      `json:"valign"`
	Vertical        bool        `json:"vertical"`
	Render          bool        `json:"render"`
	request
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#settextgdiplusproperties
type setTextGDIPlusPropertiesResponse response

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#gettextfreetype2properties
type getTextFreetype2PropertiesRequest struct {
	SceneName string `json:"scene-name"`
	Source    string `json:"source"`
	request
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#gettextfreetype2properties
type getTextFreetype2PropertiesResponse struct {
	Color1      int         `json:"color1"`
	Color2      int         `json:"color2"`
	CustomWidth int         `json:"custom_width"`
	DropShadow  bool        `json:"drop_shadow"`
	Font        interface{} `json:"font"` // TODO: Object.
	FontFace    string      `json:"font.face"`
	FontFlags   int         `json:"font.flags"`
	FontSize    int         `json:"font.size"`
	FontStyle   string      `json:"font.style"`
	FromFile    bool        `json:"from_file"`
	LogMode     bool        `json:"log_mode"`
	Outline     bool        `json:"outline"`
	Text        string      `json:"text"`
	TextFile    string      `json:"text_file"`
	WordWrap    bool        `json:"word_wrap"`
	Render      bool        `json:"render"`
	response
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#settextfreetype2properties
type setTextFreetype2PropertiesRequest struct {
	SceneName   string      `json:"scene-name"`
	Color1      int         `json:"color1"`
	Color2      int         `json:"color2"`
	CustomWidth int         `json:"custom_width"`
	DropShadow  bool        `json:"drop_shadow"`
	Font        interface{} `json:"font"` // TODO: Object.
	FontFace    string      `json:"font.face"`
	FontFlags   int         `json:"font.flags"`
	FontSize    int         `json:"font.size"`
	FontStyle   string      `json:"font.style"`
	FromFile    bool        `json:"from_file"`
	LogMode     bool        `json:"log_mode"`
	Outline     bool        `json:"outline"`
	Text        string      `json:"text"`
	TextFile    string      `json:"text_file"`
	WordWrap    bool        `json:"word_wrap"`
	Render      bool        `json:"render"`
	request
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#settextfreetype2properties
type setTextFreetype2PropertiesResponse response

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getbrowsersourceproperties
type getBrowserSourcePropertiesRequest struct {
	SceneName string `json:"scene-name"`
	Source    string `json:"source"`
	request
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getbrowsersourceproperties
type getBrowserSourcePropertiesResponse struct {
	IsLocalFile bool   `json:"is_local_file"`
	LocalFile   string `json:"local_file"`
	URL         string `json:"url"`
	CSS         string `json:"css"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	FPS         int    `json:"fps"`
	Shutdown    bool   `json:"shutdown"`
	Render      bool   `json:"render"`
	response
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setbrowsersourceproperties
type setBrowserSourcePropertiesRequest struct {
	SceneName   string `json:"scene-name"`
	Source      string `json:"source"`
	IsLocalFile bool   `json:"is_local_file"`
	LocalFile   string `json:"local_file"`
	URL         string `json:"url"`
	CSS         string `json:"css"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	FPS         int    `json:"fps"`
	Shutdown    bool   `json:"shutdown"`
	Render      bool   `json:"render"`
	response
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setbrowsersourceproperties
type setBrowserSourcePropertiesResponse response

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#deletesceneitem
type deleteSceneItemRequest struct { // Unreleased.
	Scene    string      `json:"scene"`
	Item     interface{} `json:"item"` // TODO: Object.
	ItemName string      `json:"item.name"`
	ItemID   int         `json:"item.id"`
	request
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#deletesceneitem
type deleteSceneItemResponse response

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#duplicatesceneitem
type duplicateSceneItemRequest struct { // Unreleased.
	FromScene string      `json:"fromScene"`
	ToScene   string      `json:"toScene"`
	Item      interface{} `json:"item"` // TODO: Object.
	ItemName  string      `json:"item.name"`
	ItemID    int         `json:"item.id"`
	request
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#duplicatesceneitem
type duplicateSceneItemResponse response

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getspecialsources
type getSpecialSourcesRequest request

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getspecialsources
type getSpecialSourcesResponse struct {
	Desktop1 string `json:"desktop-1"`
	Desktop2 string `json:"desktop-2"`
	Mic1     string `json:"mic-1"`
	Mic2     string `json:"mic-2"`
	Mic3     string `json:"mic-3"`
	response
}
