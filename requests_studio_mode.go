package obsws

import (
	"errors"
	"time"
)

// This file is automatically generated.
// https://github.com/christopher-dG/go-obs-websocket/blob/master/codegen/protocol.py

// GetStudioModeStatusRequest : Indicates if Studio Mode is currently enabled.
//
// Since obs-websocket version: 4.1.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#getstudiomodestatus
type GetStudioModeStatusRequest struct {
	_request `json:",squash"`
	response chan GetStudioModeStatusResponse
}

// NewGetStudioModeStatusRequest returns a new GetStudioModeStatusRequest.
func NewGetStudioModeStatusRequest() GetStudioModeStatusRequest {
	return GetStudioModeStatusRequest{
		_request{
			ID_:   GetMessageID(),
			Type_: "GetStudioModeStatus",
			err:   make(chan error, 1),
		},
		make(chan GetStudioModeStatusResponse, 1),
	}
}

// Send sends the request.
func (r *GetStudioModeStatusRequest) Send(c Client) error {
	if r.sent {
		return ErrAlreadySent
	}
	future, err := c.SendRequest(r)
	if err != nil {
		return err
	}
	r.sent = true
	go func() {
		m := <-future
		var resp GetStudioModeStatusResponse
		if err = mapToStruct(m, &resp); err != nil {
			r.err <- err
		} else if resp.Status() != StatusOK {
			r.err <- errors.New(resp.Error())
		} else {
			r.response <- resp
		}
	}()
	return nil
}

// Receive waits for the response.
func (r GetStudioModeStatusRequest) Receive() (GetStudioModeStatusResponse, error) {
	if !r.sent {
		return GetStudioModeStatusResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetStudioModeStatusResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetStudioModeStatusResponse{}, err
		case <-time.After(receiveTimeout):
			return GetStudioModeStatusResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r GetStudioModeStatusRequest) SendReceive(c Client) (GetStudioModeStatusResponse, error) {
	if err := r.Send(c); err != nil {
		return GetStudioModeStatusResponse{}, err
	}
	return r.Receive()
}

// GetStudioModeStatusResponse : Response for GetStudioModeStatusRequest.
//
// Since obs-websocket version: 4.1.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#getstudiomodestatus
type GetStudioModeStatusResponse struct {
	// Indicates if Studio Mode is enabled.
	// Required: Yes.
	StudioMode bool `json:"studio-mode"`
	_response  `json:",squash"`
}

// GetPreviewSceneRequest : Get the name of the currently previewed scene and its list of sources.
// Will return an `error` if Studio Mode is not enabled.
//
// Since obs-websocket version: 4.1.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#getpreviewscene
type GetPreviewSceneRequest struct {
	_request `json:",squash"`
	response chan GetPreviewSceneResponse
}

// NewGetPreviewSceneRequest returns a new GetPreviewSceneRequest.
func NewGetPreviewSceneRequest() GetPreviewSceneRequest {
	return GetPreviewSceneRequest{
		_request{
			ID_:   GetMessageID(),
			Type_: "GetPreviewScene",
			err:   make(chan error, 1),
		},
		make(chan GetPreviewSceneResponse, 1),
	}
}

// Send sends the request.
func (r *GetPreviewSceneRequest) Send(c Client) error {
	if r.sent {
		return ErrAlreadySent
	}
	future, err := c.SendRequest(r)
	if err != nil {
		return err
	}
	r.sent = true
	go func() {
		m := <-future
		var resp GetPreviewSceneResponse
		if err = mapToStruct(m, &resp); err != nil {
			r.err <- err
		} else if resp.Status() != StatusOK {
			r.err <- errors.New(resp.Error())
		} else {
			r.response <- resp
		}
	}()
	return nil
}

// Receive waits for the response.
func (r GetPreviewSceneRequest) Receive() (GetPreviewSceneResponse, error) {
	if !r.sent {
		return GetPreviewSceneResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetPreviewSceneResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetPreviewSceneResponse{}, err
		case <-time.After(receiveTimeout):
			return GetPreviewSceneResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r GetPreviewSceneRequest) SendReceive(c Client) (GetPreviewSceneResponse, error) {
	if err := r.Send(c); err != nil {
		return GetPreviewSceneResponse{}, err
	}
	return r.Receive()
}

// GetPreviewSceneResponse : Response for GetPreviewSceneRequest.
//
// Since obs-websocket version: 4.1.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#getpreviewscene
type GetPreviewSceneResponse struct {
	// The name of the active preview scene.
	// Required: Yes.
	Name string `json:"name"`
	// Required: Yes.
	Sources   []*SceneItem `json:"sources"`
	_response `json:",squash"`
}

// SetPreviewSceneRequest : Set the active preview scene.
// Will return an `error` if Studio Mode is not enabled.
//
// Since obs-websocket version: 4.1.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#setpreviewscene
type SetPreviewSceneRequest struct {
	// The name of the scene to preview.
	// Required: Yes.
	SceneName string `json:"scene-name"`
	_request  `json:",squash"`
	response  chan SetPreviewSceneResponse
}

// NewSetPreviewSceneRequest returns a new SetPreviewSceneRequest.
func NewSetPreviewSceneRequest(sceneName string) SetPreviewSceneRequest {
	return SetPreviewSceneRequest{
		sceneName,
		_request{
			ID_:   GetMessageID(),
			Type_: "SetPreviewScene",
			err:   make(chan error, 1),
		},
		make(chan SetPreviewSceneResponse, 1),
	}
}

// Send sends the request.
func (r *SetPreviewSceneRequest) Send(c Client) error {
	if r.sent {
		return ErrAlreadySent
	}
	future, err := c.SendRequest(r)
	if err != nil {
		return err
	}
	r.sent = true
	go func() {
		m := <-future
		var resp SetPreviewSceneResponse
		if err = mapToStruct(m, &resp); err != nil {
			r.err <- err
		} else if resp.Status() != StatusOK {
			r.err <- errors.New(resp.Error())
		} else {
			r.response <- resp
		}
	}()
	return nil
}

// Receive waits for the response.
func (r SetPreviewSceneRequest) Receive() (SetPreviewSceneResponse, error) {
	if !r.sent {
		return SetPreviewSceneResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return SetPreviewSceneResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return SetPreviewSceneResponse{}, err
		case <-time.After(receiveTimeout):
			return SetPreviewSceneResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r SetPreviewSceneRequest) SendReceive(c Client) (SetPreviewSceneResponse, error) {
	if err := r.Send(c); err != nil {
		return SetPreviewSceneResponse{}, err
	}
	return r.Receive()
}

// SetPreviewSceneResponse : Response for SetPreviewSceneRequest.
//
// Since obs-websocket version: 4.1.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#setpreviewscene
type SetPreviewSceneResponse struct {
	_response `json:",squash"`
}

// TransitionToProgramRequest : Transitions the currently previewed scene to the main output.
// Will return an `error` if Studio Mode is not enabled.
//
// Since obs-websocket version: 4.1.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#transitiontoprogram
type TransitionToProgramRequest struct {
	// Change the active transition before switching scenes.
	// Defaults to the active transition.
	// Required: No.
	WithTransition map[string]interface{} `json:"with-transition"`
	// Name of the transition.
	// Required: Yes.
	WithTransitionName string `json:"with-transition.name"`
	// Transition duration (in milliseconds).
	// Required: No.
	WithTransitionDuration int `json:"with-transition.duration"`
	_request               `json:",squash"`
	response               chan TransitionToProgramResponse
}

// NewTransitionToProgramRequest returns a new TransitionToProgramRequest.
func NewTransitionToProgramRequest(
	withTransition map[string]interface{},
	withTransitionName string,
	withTransitionDuration int,
) TransitionToProgramRequest {
	return TransitionToProgramRequest{
		withTransition,
		withTransitionName,
		withTransitionDuration,
		_request{
			ID_:   GetMessageID(),
			Type_: "TransitionToProgram",
			err:   make(chan error, 1),
		},
		make(chan TransitionToProgramResponse, 1),
	}
}

// Send sends the request.
func (r *TransitionToProgramRequest) Send(c Client) error {
	if r.sent {
		return ErrAlreadySent
	}
	future, err := c.SendRequest(r)
	if err != nil {
		return err
	}
	r.sent = true
	go func() {
		m := <-future
		var resp TransitionToProgramResponse
		if err = mapToStruct(m, &resp); err != nil {
			r.err <- err
		} else if resp.Status() != StatusOK {
			r.err <- errors.New(resp.Error())
		} else {
			r.response <- resp
		}
	}()
	return nil
}

// Receive waits for the response.
func (r TransitionToProgramRequest) Receive() (TransitionToProgramResponse, error) {
	if !r.sent {
		return TransitionToProgramResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return TransitionToProgramResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return TransitionToProgramResponse{}, err
		case <-time.After(receiveTimeout):
			return TransitionToProgramResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r TransitionToProgramRequest) SendReceive(c Client) (TransitionToProgramResponse, error) {
	if err := r.Send(c); err != nil {
		return TransitionToProgramResponse{}, err
	}
	return r.Receive()
}

// TransitionToProgramResponse : Response for TransitionToProgramRequest.
//
// Since obs-websocket version: 4.1.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#transitiontoprogram
type TransitionToProgramResponse struct {
	_response `json:",squash"`
}

// EnableStudioModeRequest : Enables Studio Mode.
//
// Since obs-websocket version: 4.1.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#enablestudiomode
type EnableStudioModeRequest struct {
	_request `json:",squash"`
	response chan EnableStudioModeResponse
}

// NewEnableStudioModeRequest returns a new EnableStudioModeRequest.
func NewEnableStudioModeRequest() EnableStudioModeRequest {
	return EnableStudioModeRequest{
		_request{
			ID_:   GetMessageID(),
			Type_: "EnableStudioMode",
			err:   make(chan error, 1),
		},
		make(chan EnableStudioModeResponse, 1),
	}
}

// Send sends the request.
func (r *EnableStudioModeRequest) Send(c Client) error {
	if r.sent {
		return ErrAlreadySent
	}
	future, err := c.SendRequest(r)
	if err != nil {
		return err
	}
	r.sent = true
	go func() {
		m := <-future
		var resp EnableStudioModeResponse
		if err = mapToStruct(m, &resp); err != nil {
			r.err <- err
		} else if resp.Status() != StatusOK {
			r.err <- errors.New(resp.Error())
		} else {
			r.response <- resp
		}
	}()
	return nil
}

// Receive waits for the response.
func (r EnableStudioModeRequest) Receive() (EnableStudioModeResponse, error) {
	if !r.sent {
		return EnableStudioModeResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return EnableStudioModeResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return EnableStudioModeResponse{}, err
		case <-time.After(receiveTimeout):
			return EnableStudioModeResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r EnableStudioModeRequest) SendReceive(c Client) (EnableStudioModeResponse, error) {
	if err := r.Send(c); err != nil {
		return EnableStudioModeResponse{}, err
	}
	return r.Receive()
}

// EnableStudioModeResponse : Response for EnableStudioModeRequest.
//
// Since obs-websocket version: 4.1.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#enablestudiomode
type EnableStudioModeResponse struct {
	_response `json:",squash"`
}

// DisableStudioModeRequest : Disables Studio Mode.
//
// Since obs-websocket version: 4.1.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#disablestudiomode
type DisableStudioModeRequest struct {
	_request `json:",squash"`
	response chan DisableStudioModeResponse
}

// NewDisableStudioModeRequest returns a new DisableStudioModeRequest.
func NewDisableStudioModeRequest() DisableStudioModeRequest {
	return DisableStudioModeRequest{
		_request{
			ID_:   GetMessageID(),
			Type_: "DisableStudioMode",
			err:   make(chan error, 1),
		},
		make(chan DisableStudioModeResponse, 1),
	}
}

// Send sends the request.
func (r *DisableStudioModeRequest) Send(c Client) error {
	if r.sent {
		return ErrAlreadySent
	}
	future, err := c.SendRequest(r)
	if err != nil {
		return err
	}
	r.sent = true
	go func() {
		m := <-future
		var resp DisableStudioModeResponse
		if err = mapToStruct(m, &resp); err != nil {
			r.err <- err
		} else if resp.Status() != StatusOK {
			r.err <- errors.New(resp.Error())
		} else {
			r.response <- resp
		}
	}()
	return nil
}

// Receive waits for the response.
func (r DisableStudioModeRequest) Receive() (DisableStudioModeResponse, error) {
	if !r.sent {
		return DisableStudioModeResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return DisableStudioModeResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return DisableStudioModeResponse{}, err
		case <-time.After(receiveTimeout):
			return DisableStudioModeResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r DisableStudioModeRequest) SendReceive(c Client) (DisableStudioModeResponse, error) {
	if err := r.Send(c); err != nil {
		return DisableStudioModeResponse{}, err
	}
	return r.Receive()
}

// DisableStudioModeResponse : Response for DisableStudioModeRequest.
//
// Since obs-websocket version: 4.1.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#disablestudiomode
type DisableStudioModeResponse struct {
	_response `json:",squash"`
}

// ToggleStudioModeRequest : Toggles Studio Mode (depending on the current state of studio mode).
//
// Since obs-websocket version: 4.1.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#togglestudiomode
type ToggleStudioModeRequest struct {
	_request `json:",squash"`
	response chan ToggleStudioModeResponse
}

// NewToggleStudioModeRequest returns a new ToggleStudioModeRequest.
func NewToggleStudioModeRequest() ToggleStudioModeRequest {
	return ToggleStudioModeRequest{
		_request{
			ID_:   GetMessageID(),
			Type_: "ToggleStudioMode",
			err:   make(chan error, 1),
		},
		make(chan ToggleStudioModeResponse, 1),
	}
}

// Send sends the request.
func (r *ToggleStudioModeRequest) Send(c Client) error {
	if r.sent {
		return ErrAlreadySent
	}
	future, err := c.SendRequest(r)
	if err != nil {
		return err
	}
	r.sent = true
	go func() {
		m := <-future
		var resp ToggleStudioModeResponse
		if err = mapToStruct(m, &resp); err != nil {
			r.err <- err
		} else if resp.Status() != StatusOK {
			r.err <- errors.New(resp.Error())
		} else {
			r.response <- resp
		}
	}()
	return nil
}

// Receive waits for the response.
func (r ToggleStudioModeRequest) Receive() (ToggleStudioModeResponse, error) {
	if !r.sent {
		return ToggleStudioModeResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return ToggleStudioModeResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return ToggleStudioModeResponse{}, err
		case <-time.After(receiveTimeout):
			return ToggleStudioModeResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r ToggleStudioModeRequest) SendReceive(c Client) (ToggleStudioModeResponse, error) {
	if err := r.Send(c); err != nil {
		return ToggleStudioModeResponse{}, err
	}
	return r.Receive()
}

// ToggleStudioModeResponse : Response for ToggleStudioModeRequest.
//
// Since obs-websocket version: 4.1.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#togglestudiomode
type ToggleStudioModeResponse struct {
	_response `json:",squash"`
}
