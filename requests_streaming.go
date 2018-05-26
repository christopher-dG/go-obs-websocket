package obsws

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#streaming-1

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getstreamingstatus
type getStreamingStatusRequest request

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getstreamingstatus
type getStreamingStatusResponse struct {
	Streaming      bool   `json:"streaming"`
	Recording      bool   `json:"recording"`
	StreamTimecode string `json:"stream-timecode"`
	RecTimecode    string `json:"rec-timecode"`
	PreviewOnly    bool   `json:"preview-only"`
	response
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#startstopstreaming
type startStopStreamingRequest request

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#startstopstreaming
type startStopStreamingResponse response

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#startstreaming
type startStreamingRequest struct {
	Stream                 interface{} `json:"stream"` // TODO: Object.
	StreamType             string      `json:"stream.type"`
	StreamMetadata         interface{} `json:"stream.metadata"` // TODO: Object.
	StreamSettings         interface{} `json:"stream.settings"` // TODO: Object.
	StreamSettingsServer   string      `json:"stream.settings.server"`
	StreamSettingsKey      string      `json:"stream.settings.key"`
	StreamSettingsUseAuth  bool        `json:"stream.settings.use-auth"`
	StreamSettingsUsername string      `json:"stream.settings.username"`
	StreamSettingsPassword string      `json:"stream.settings.password"`
	request
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#startstreaming
type startStreamingResponse response

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#stopstreaming
type stopStreamingRequest request

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#stopstreaming
type stopStreamingResponse response

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setstreamsettings
type setStreamSettingsRequest struct {
	Type             string      `json:"type"`
	Settings         interface{} `json:"settings"` // TODO: Object.
	SettingsServer   string      `json:"settings.server"`
	SettingsKey      string      `json:"settings.key"`
	SettingsUseAuth  bool        `json:"settings.use-auth"`
	SettingsUsername string      `json:"settings.username"`
	SettingsPassword string      `json:"settings.password"`
	Save             bool        `json:"settings.save"`
	request
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setstreamsettings
type setStreamSettingsResponse response

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getstreamsettings
type getStreamSettingsRequest request

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getstreamsettings
type getStreamSettingsResponse struct {
	Type             string      `json:"type"`
	Settings         interface{} `json:"settings"` // TODO: Object.
	SettingsServer   string      `json:"settings.server"`
	SettingsKey      string      `json:"settings.key"`
	SettingsUseAuth  bool        `json:"settings.use-auth"`
	SettingsUsername string      `json:"settings.username"`
	SettingsPassword string      `json:"settings.password"`
	response
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#savestreamsettings
type saveStreamSettingsRequest request

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#savestreamsettings
type saveStreamSettingsResponse response
