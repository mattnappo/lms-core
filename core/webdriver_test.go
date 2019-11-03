package core

import "testing"

func TestNewWebDriver(t *testing.T) {
	port := 8080

	cwd, err := NewWebDriver(port)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(cwd.String())
}
