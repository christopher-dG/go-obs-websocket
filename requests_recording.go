package obsws

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#recording-1

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#startstoprecording
type startStopRecordingRequest request

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#startstoprecording
type startStopRecordingResponse response

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#startrecording
type startRecordingRequest request

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#startrecording
type startRecordingResponse response

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#stoprecording
type stopRecordingRequest request

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#stoprecording
type stopRecordingResponse response

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setrecordingfolder
type setRecordingFolderRequest struct {
	RecFolder string `json:"rec-folder"`
	request
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setrecordingfolder
type setRecordingFolderResponse response

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setrecordingfolder
type getRecordingFolderRequest request

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setrecordingfolder
type getRecordingFolderResponse struct {
	RecFolder string `json:"rec-folder"`
	response
}
