package main

import (
	"fmt"
	"os"

	"github.com/new-lms/lms-core/core"
	"github.com/new-lms/lms-core/scraper"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

const (
	SeleniumPath     = "vendor/selenium-server.jar"
	ChromeDriverPath = "vendor/chromedriver"
	ChromeBinPath    = "vendor/chrome-linux/chrome"
)

func mainExample(port int) (*selenium.WebDriver, error) {
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

	return &webDriver, nil
}

func otherExample(port int) error {
	options := []selenium.ServiceOption{
		selenium.StartFrameBuffer(),             // Start an X frame buffer for the browser to run in
		selenium.ChromeDriver(ChromeDriverPath), // Specify the path to the chroem driver
		selenium.Output(os.Stderr),              // Output debug information to STDERR
	}

	// Initialize the selenium service
	service, err := selenium.NewSeleniumService(SeleniumPath, port, options...)
	if err != nil {
		return err
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
		return err
	}

	err = scraper.Scrape(&webDriver)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	core.InitSelenium()
	core.InitPaths()

	// scripts.InstallVendor()

	// webDriver, err := mainExample(8081)
	// if err != nil {
	// 	panic(err)
	// }

	// err = scraper.Scrape(webDriver)
	// if err != nil {
	// 	panic(err)
	// }

	// err := otherExample(8081)
	// if err != nil {
	// 	panic(err)
	// }

	_, err := core.NewChromeWebDriver(8081)
	if err != nil {
		panic(err)
	}
}
