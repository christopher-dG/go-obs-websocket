package obsws

import "testing"

func TestGetAuth(t *testing.T) {
	expected := "zTM5ki6L2vVvBQiTG9ckH1Lh64AbnCf6XZ226UmnkIA="
	observed := getAuth("password", "salt", "challenge")
	if observed != expected {
		t.Errorf("expected auth == '%s', got '%s'", expected, observed)
	}
}
