package obsws

import "strconv"

func (c *client) getMessageID() string {
	c.id++
	return strconv.Itoa(c.id)
}

func (c *client) validateMessageID(req, resp reqOrResp) bool {
	return req.ID() == resp.ID()
}
