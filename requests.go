package obsws

type request interface {
	ID() string
	Type() string
}

type response interface {
	ID() string
	Stat() string
	Err() string
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#requests
type _request struct {
	MessageID   string `json:"message-id"`
	RequestType string `json:"request-type"`
}

func (r _request) ID() string { return r.MessageID }

func (r _request) Type() string { return r.RequestType }

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#requests
type _response struct {
	MessageID string `mapstructure:"message-id"`
	Status    string `mapstructure:"status"`
	Error     string `mapstructure:"error"`
}

func (r _response) ID() string { return r.MessageID }

func (r _response) Stat() string { return r.Status }

func (r _response) Err() string { return r.Error }
