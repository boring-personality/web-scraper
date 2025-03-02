## Output Indicators

- 🔍 Scraping in progress
- ✅ Successful crawl
- ❌ Failed to process
- 📄 Internal page queued
- 🔗 External link queued
- 🌐 Validating external link
- ⏭️ Skipping already visited URL


## Implementation Details

- Uses Go's `sync.WaitGroup` for worker coordination
- Utilizes `golang.org/x/net/html` for HTML parsing
- Handles relative and absolute URLs using Go's `net/url` package

## Error Handling

The crawler includes robust error handling for:
- Invalid URLs
- Network failures
- HTTP errors
- Malformed HTML
- External link validation


## Project Directory Structure

```plaintext
webscraper/
├── cmd/                    # Entry points for CLI/web scraper
│   ├── webscraper/
│   │   ├── main.go         # Main function
├── internal/               # Core business logic
│   ├── crawler/            # Web crawling logic
│   │   ├── crawler.go      # Main recursive crawling logic
│   ├── utils/              # Stores visited URLs
│   │   ├── url_utils.go    # Manages visited links
│   ├── validator/          # Link validation
│   │   ├── link_validator.go    # Checks for valid links
├── go.mod                  # Go module file
├── go.sum                  # Dependencies
├── README.md               # Documentation
