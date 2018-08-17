package obsws

// Request is a request to obs-websocket.
type Request interface {
	ID() string
	Type() string
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
}

func (r _request) ID() string { return r.ID_ }

func (r _request) Type() string { return r.Type_ }

func (r _request) Send(c Client) (chan _response, error) {
	generic, err := c.SendRequest(r)
	if err != nil {
		return nil, err
	}
	future := make(chan _response)
	go func() { future <- (<-generic).(_response) }()
	return future, nil
}

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
func (r _response) Error() string { return r.Error_ }
