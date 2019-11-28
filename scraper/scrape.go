package scraper

import (
	"fmt"
	"strings"
	"time"

	// "github.com/new-lms/lms-core/core"
	"github.com/tebeka/selenium"
)

// Scrape scrapes stuff using a core.ChromeWebDriver.
func Scrape(driver *selenium.WebDriver) error {
	fmt.Printf("I AM SCRAPING NOW\n\n\n")

	err := (*driver).Get("http://play.golang.org/?simple=1")
	if err != nil {
		return err
	}

	elem, err := (*driver).FindElement(selenium.ByCSSSelector, "#code")
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
	btn, err := (*driver).FindElement(selenium.ByCSSSelector, "#run")
	if err != nil {
		return err
	}
	if err := btn.Click(); err != nil {
		return err
	}

	// Wait for the program to finish running and get the output.
	outputDiv, err := (*driver).FindElement(selenium.ByCSSSelector, "#output")
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

	return nil
}
