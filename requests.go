package obsws

import "errors"

const (
	// StatusOK indicates that the request was successful.
	StatusOK = "ok"
	// StatusError indicates that the request was unsuccessful.
	StatusError = "error"
)

var (
	// ErrNotSent is returned when you call Receive on a request that has not been sent.
	ErrNotSent = errors.New("request not yet sent")
	// ErrAlreadySent is returned when a request has already been sent.
	ErrAlreadySent = errors.New("request already sent")
)

// Request is a request to obs-websocket.
type Request interface {
	ID() string
	Type() string
	Send(Client) error
}

// Response is a response from obs-websocket.
type Response interface {
	ID() string
	Status() string
	Error() string
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#requests
type _request struct {
	ID_   string `json:"message-id"`
	Type_ string `json:"request-type"`
	sent  bool
	err   chan error
}

func (r _request) Send(c Client) error { return nil }

// ID returns the request's message ID.
func (r _request) ID() string { return r.ID_ }

// Type returns the request's message type.
func (r _request) Type() string { return r.Type_ }

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#requests
type _response struct {
	ID_     string `json:"message-id"`
	Status_ string `json:"status"`
	Error_  string `json:"error"`
}

// ID returns the response's message ID.
func (r _response) ID() string { return r.ID_ }

// Status returns the response's status.
func (r _response) Status() string { return r.Status_ }

// Error returns the response's error.
// When using Receive or SendReceive, this should always return an empty string,
// because the error will have been returned explictly instead of stored here.
func (r _response) Error() string { return r.Error_ }
