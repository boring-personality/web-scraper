package main

import (
	"net/url"
	"scraper/internal/crawler"
)

const URL string = "https://scrape-me.dreamsofcode.io/"

func main() {
	baseUrl, _ := url.Parse(URL)
	crawler.Crawl(URL, baseUrl.Host)
}
