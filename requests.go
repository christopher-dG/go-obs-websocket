package obsws

type message interface {
	ID() string
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#requests
type request struct {
	MessageID   string `json:"message-id"`
	RequestType string `json:"request-type"`
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#requests
type response struct {
	MessageID string `json:"message-id"`
	Status    string `json:"status"`
	Error     string `json:"error"`
}

func (r *request) ID() string {
	return r.MessageID
}

func (r *response) ID() string {
	return r.MessageID
}
