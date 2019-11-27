package core

import "github.com/tebeka/selenium"

// InitSelenium initializes selenium.
func InitSelenium() {
	InitPaths()
	selenium.SetDebug(DebugMode)
}
