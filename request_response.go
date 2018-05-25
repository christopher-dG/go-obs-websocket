package obsws

type reqOrResp interface {
	ID() string
}

type genericRequest struct {
	MessageID   string `json:"message-id"`
	RequestType string `json:"request-type"`
}

func (r *genericRequest) ID() string {
	return r.MessageID
}

type genericResponse struct {
	MessageID string `json:"message-id"`
	Status    string `json:"status"`
	Error     string `json:"error"`
}

func (r *genericResponse) ID() string {
	return r.MessageID
}

// ReceiveGeneric receives a minimal response with only the required fields.
func (c *client) ReceiveGeneric() (*genericResponse, error) {
	resp := &genericResponse{}
	if err := c.conn.ReadJSON(resp); err != nil {
		return nil, err
	}
	return resp, nil
}
