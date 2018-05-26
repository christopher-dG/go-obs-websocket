package obsws

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#profiles-1

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#setcurrentprofile
type setCurrentProfileRequest struct {
	ProfileName string `json:"profile-name"`
	request
}

type setCurrentProfileResponse response

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getcurrentprofile
type getCurrentProfileRequest request

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#getcurrentprofile
type getCurrentProfileResponse struct {
	ProfileName string `json:"profile-name"`
	response
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#listprofiles
type listProfilesRequest request

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#listprofiles
type listProfilesResponse struct {
	Profiles interface{} `json:"profiles"` // TODO: Object|Array.
	response
}
