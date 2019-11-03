package core

import (
	"fmt"
	"os"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

// WebDriver is an abstraction class for a selenium webdriver.
type WebDriver struct {
	WebDriver *selenium.WebDriver `json:"web_driver"` // The web driver itself

	Options      []selenium.ServiceOption `json:"options"`      // The service configuration/options
	Capabilities selenium.Capabilities    `json:"Capabilities"` // The capabilities (further browser configuration)
}

// NewWebDriver constructs a new web driver.
func NewWebDriver() {
	serviceOptions := []selenium.ServiceOption{
		selenium.StartFrameBuffer(),             // Start an X frame buffer for the browser to run in
		selenium.ChromeDriver(ChromeDriverPath), // Specify the path to the chroem driver
		selenium.Output(os.Stderr),              // Output debug information to STDERR
	}

	selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(SeleniumPath, WebDriverPort, serviceOptions...)
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}
	defer service.Stop()

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browser": "chrome"}

	var chromeCaps chrome.Capabilities
	chromeCaps.Path = ChromeBinPath

	caps.AddChrome(chromeCaps)

	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", WebDriverPort))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

}
