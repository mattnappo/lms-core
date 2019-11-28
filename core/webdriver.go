package core

import (
	"encoding/json"
	"fmt"
	"os"
	// "strings"
	// "time"

	// "github.com/new-lms/lms-core/scraper"
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

// NewChromeWebDriver returns a new LIVE web driver.
func NewChromeWebDriver(port int) (*ChromeWebDriver, error) {
	fmt.Printf("\n\nSeleniumPath: [%s]\nChromeDriverPath: [%s]\n\n", SeleniumPath, ChromeDriverPath)
	
	options := []selenium.ServiceOption{
		selenium.StartFrameBuffer(),             // Start an X frame buffer for the browser to run in
		selenium.ChromeDriver(ChromeDriverPath), // Specify the path to the chroem driver
		selenium.Output(os.Stderr),              // Output debug information to STDERR
	}

	// Initialize the selenium service
	_, err := selenium.NewSeleniumService(SeleniumPath, port, options...)
	if err != nil {
		return nil, err
	}
	// defer service.Stop()

	// Connect to the webdriver instance running locally.
	caps := selenium.Capabilities{"browser": "chrome"}

	// Declare the capabilities for chrome
	var chromeCaps chrome.Capabilities
	chromeCaps.Path = ChromeBinPath
	caps.AddChrome(chromeCaps)

	// Create the web driver remote itself
	webDriver, err := selenium.NewRemote(
		caps, // The capabilities
		fmt.Sprintf("http://localhost:%d/wd/hub", // The ip to listen on
			port), // The port to listen on
	)
	if err != nil {
		return nil, err
	}

	fmt.Printf("I AM RETURNING THE CWD NOW\n\n\n")

	// defer webDriver.Quit()

	//err = scraper.Scrape(&webDriver)
	//if err != nil {
	// 	 return nil, err
	//}

	// Construct the ChromeWebDriver
	newCWD := &ChromeWebDriver{
		WebDriver: &webDriver, // The live webdriver itself

		Options:      options, // The options declared earlier
		Capabilities: caps,    // The capabilities declared earlier

		Port:    port, // The port of the instance
		Running: true, // The instance is currently running
	}
	return newCWD, nil

}

// String marshals a ChromeWebDriver.
func (cwd *ChromeWebDriver) String() string {
	json, _ := json.MarshalIndent(*cwd, "", "  ")
	return string(json)
}
