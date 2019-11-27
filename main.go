package main

import (
	"fmt"
	"os"

	"github.com/new-lms/lms-core/scraper"
	"github.com/new-lms/lms-core/core"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

const (
	SeleniumPath     = "vendor/selenium-server.jar"
	ChromeDriverPath = "vendor/chromedriver"
	ChromeBinPath    = "vendor/chrome-linux/chrome"
)

func mainExample(port int) error {
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
	// core.InitSelenium()

	// cwd, err := core.NewChromeWebDriver(8081)
	// if err != nil {
	// 	panic(err)
	// }

	// err = scraper.Scrape(*cwd)
	// if err != nil {
	// 	panic(err)
	// }

	// mainExample(8001)
	cwd, err := core.NewChromeWebDriver(8081)
	if err != nil {
		panic(err)
	}

	err = scraper.Scrape(cwd.WebDriver)
	if err != nil {
		panic(err)
	}

}
