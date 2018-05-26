package obsws

type reqOrResp interface {
	ID() string
}

type request struct {
	MessageID   string `json:"message-id"`
	RequestType string `json:"request-type"`
}

func (r *request) ID() string {
	return r.MessageID
}

type response struct {
	MessageID string `json:"message-id"`
	Status    string `json:"status"`
	Error     string `json:"error"`
}

func (r *response) ID() string {
	return r.MessageID
}

// ReceiveMinimal receives a minimal response with only the required fields.
func (c *client) ReceiveMinimal() (*response, error) {
	resp := &response{}
	if err := c.conn.ReadJSON(resp); err != nil {
		return nil, err
	}
	return resp, nil
}
