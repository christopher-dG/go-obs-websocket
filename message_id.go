package obsws

import "strconv"

func (c *client) getMessageID() string {
	c.id++
	return strconv.Itoa(c.id)
}

func (c *client) validateMessageID(x, y message) bool {
	return x.ID() == y.ID()
}
