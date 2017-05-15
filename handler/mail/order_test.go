package mail

import "testing"

func TestNotify(t *testing.T) {
	_, err := OrderNotify("ABC123")
	if err != nil {
		t.Error("Expected order notification to pass, got ", err)
	}
}
