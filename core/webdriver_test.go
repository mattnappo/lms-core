package core

import "testing"

func TestNewChromeWebDriver(t *testing.T) {
	InitSelenium()

	port := 8081
	cwd, err := NewChromeWebDriver(port)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(cwd.String())
}
