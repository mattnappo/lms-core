package main

import (
	"github.com/new-lms/lms-core/core"
	"github.com/new-lms/lms-core/scraper"
)

func main() {
	core.InitSelenium()

	cwd, err := core.NewChromeWebDriver(8081)
	if err != nil {
		panic(err)
	}

	err = scraper.Scrape(*cwd)
	if err != nil {
		panic(err)
	}
}
