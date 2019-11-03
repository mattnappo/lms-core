package core

import "testing"

func TestNewChromeWebDriver(t *testing.T) {
	port := 8080

	cwd, err := NewChromeWebDriver(port)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(cwd.String())
}
