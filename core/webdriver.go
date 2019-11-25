package core

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

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

	// Create the web driver remote itself
	webDriver, err := selenium.NewRemote(
		caps, // The capabilities
		fmt.Sprintf("http://localhost:%d/wd/hub", // The ip to listen on
			port), // The port to listen on
	)
	if err != nil {
		return nil, err
	}

	// defer webDriver.Quit()

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

// Start starts the web driver service for a given ChromeWebDriver.
func (cwd *ChromeWebDriver) Start() error {
	// Create the web driver remote itself
	webDriver, err := selenium.NewRemote(
		(*cwd).Capabilities, // The capabilities
		fmt.Sprintf("http://localhost:%d/wd/hub", // The ip to listen on
			(*cwd).Port), // The port to listen on
	)
	if err != nil {
		return err
	}
	err = webDriver.Get("http://play.golang.org/?simple=1")
	if err != nil {
		return err
	}

	elem, err := webDriver.FindElement(selenium.ByCSSSelector, "#code")
	if err != nil {
		return err
	}

	err = elem.Clear()
	if err != nil {
		return err
	}

	// Enter some new code in text box.
	err = elem.SendKeys(`
		package main
		import "fmt"
		func main() {
			fmt.Println("Hello WebDriver thing!\n")
		}
	`)
	if err != nil {
		return err
	}

	// Click the run button.
	btn, err := webDriver.FindElement(selenium.ByCSSSelector, "#run")
	if err != nil {
		return err
	}
	if err := btn.Click(); err != nil {
		return err
	}

	// Wait for the program to finish running and get the output.
	outputDiv, err := webDriver.FindElement(selenium.ByCSSSelector, "#output")
	if err != nil {
		return err
	}

	var output string
	for {
		output, err = outputDiv.Text()
		if err != nil {
			return err
		}
		if output != "Waiting for remote server..." {
			break
		}
		time.Sleep(time.Millisecond * 100)
	}

	fmt.Printf("%s", strings.Replace(output, "\n\n", "\n", -1))

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

// String marshals a ChromeWebDriver.
func (cwd *ChromeWebDriver) String() string {
	json, _ := json.MarshalIndent(*cwd, "", "  ")
	return string(json)
}
