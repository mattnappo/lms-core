package main

import "github.com/new-lms/lms-core/scraper"

func main() {
	err := scraper.Scrape()
	panic(err)
}
