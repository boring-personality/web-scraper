package crawler

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"scraper/internal/utils"
	"scraper/internal/validator"
	"strings"
	"sync"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

var visited = make(map[string]bool)

func Crawl(link string, baseUrl string, wg *sync.WaitGroup) {
	defer wg.Done()
	if visited[link] {
		return
	}
	visited[link] = true
	// fmt.Println("Checking: ", link)
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println("Error getting url: ", link)
		return
	}

	if resp.StatusCode != 200 {
		fmt.Println("Invalid page:", link, "Status:", resp.StatusCode)
		return
	}

	fmt.Println("Valid page:", link)

	body, _ := io.ReadAll(resp.Body)
	node, _ := html.Parse(strings.NewReader(string(body)))

	for n := range node.Descendants() {
		if n.Type == html.ElementNode && n.DataAtom == atom.A {
			for _, a := range n.Attr {
				if a.Key == "href" {
					href := utils.ResolveURL(link, a.Val)
					parsedLink, err := url.Parse(href)
					if err != nil || href == "" {
						continue
					}

					if parsedLink.Host == baseUrl || parsedLink.Host == "" {
						wg.Add(1)
						go Crawl(href, baseUrl, wg)
					} else {
						if visited[href] {
							continue
						}
						validator.ValidateExternalLink(href)
						visited[href] = true
					}

				}
			}
		}
	}
}
