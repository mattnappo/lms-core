package core

import "path/filepath"

// DebugMode determines whether to run selenium in debug mode
const DebugMode = true

var (
	// SeleniumPath is the path to the standalone Selenium server jar.
	SeleniumPath = "vendor/selenium-server.jar"

	// ChromeDriverPath is the path to the chrome driver.
	ChromeDriverPath = "vendor/chromedriver"

	// ChromeBinPath is the path to the chrome binary.
	ChromeBinPath = "vendor/chrome-linux/chrome"
)

// InitPaths initializes the absolute paths of the above paths.
func InitPaths() {
	SeleniumPath, _ = filepath.Abs(SeleniumPath)
	ChromeDriverPath, _ = filepath.Abs(ChromeDriverPath)
	ChromeBinPath, _ = filepath.Abs(ChromeBinPath)

}
