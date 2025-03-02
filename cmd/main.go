package main

import (
	"fmt"
	"net/url"
	"scraper/internal/crawler"
	"sync"
)

const URL string = "https://scrape-me.dreamsofcode.io/"

func main() {
	var wg sync.WaitGroup
	baseUrl, _ := url.Parse(URL)
	fmt.Println("🔍 Scraping in progress")

	crawler.Crawl(URL, baseUrl.Host, &wg)

	wg.Wait()
	fmt.Println("✅ Successful crawl")
}
