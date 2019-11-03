package core

import (
	"fmt"
	"os"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

// ChromeWebDriver is an abstraction class for a selenium chrome webdriver.
type ChromeWebDriver struct {
	WebDriver *selenium.WebDriver `json:"web_driver"` // The web driver itself

	Options      []selenium.ServiceOption `json:"options"`      // The service configuration/options
	Capabilities selenium.Capabilities    `json:"Capabilities"` // The capabilities (further browser configuration)

	Running bool `json:"running"` // The status of the instance
	Port    int  `json:"port"`    // The port that the instance will run on
}

// NewWebDriver constructs a new web driver.
func NewWebDriver(port int) (*ChromeWebDriver, error) {
	options := []selenium.ServiceOption{
		selenium.StartFrameBuffer(),             // Start an X frame buffer for the browser to run in
		selenium.ChromeDriver(ChromeDriverPath), // Specify the path to the chroem driver
		selenium.Output(os.Stderr),              // Output debug information to STDERR
	}

	// Initialize the selenium service
	service, err := selenium.NewSeleniumService(SeleniumPath, port, options...)
	if err != nil {
		return nil, err
	}
	defer service.Stop()

	// Connect to the webdriver instance running locally.
	caps := selenium.Capabilities{"browser": "chrome"}

	// Declare the capabilities for chrome
	var chromeCaps chrome.Capabilities
	chromeCaps.Path = ChromeBinPath
	caps.AddChrome(chromeCaps)

	// Construct the ChromeWebDriver
	newCWD := &ChromeWebDriver{
		WebDriver: nil, // Init as nil b/c webdriver service has not been started yet

		Options:      options, // The options declared earlier
		Capabilities: caps,    // The capabilities declared earlier

		Port:    port,  // The port of the instance
		Running: false, // The instance is not currently running
	}
	return newCWD, nil

}

// Start starts the web driver service for a given ChromeWebDriver.
func (cwd *ChromeWebDriver) Start() error {
	// Create the web driver remote itself
	webDriver, err := selenium.NewRemote((*cwd).Capabilities, fmt.Sprintf("http://localhost:%d/wd/hub", (*cwd).Port))
	if err != nil {
		return err
	}

	// defer webDriver.Quit()

	cwd.WebDriver = &webDriver
	cwd.Running = true

	return nil
}

// Stop stops the web driver service for a given ChromeWebDriver.
func (cwd *ChromeWebDriver) Stop() error {
	// Stop the webdriver
	webDriver := *cwd.WebDriver
	webDriver.Quit()

	// Delete the webdriver
	cwd.WebDriver = nil
	cwd.Running = true

	return nil
}
