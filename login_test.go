package ravello

import "testing"

func TestLogin(t *testing.T) {
	u, err := Login()
	if err != nil {
		t.Fatal("Login failed:", err)
	}
	if u.Email == "" {
		t.Fail()
		t.Logf("Expected email to not be empty, got: %v", u.Email)
	}
}
