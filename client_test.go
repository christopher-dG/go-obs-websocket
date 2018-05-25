package obsws

import "testing"

func TestNewClient(t *testing.T) {
	c := NewClient("localhost", 4444, "")
	if c.Host != "localhost" {
		t.Errorf("expected c.Host == 'localhost', got '%s'", c.Host)
	}
	if c.Port != 4444 {
		t.Errorf("expected c.Port == 4444, got '%d'", c.Port)
	}
	if c.Password != "" {
		t.Errorf("expected c.Password == '', got '%s'", c.Password)
	}
}
