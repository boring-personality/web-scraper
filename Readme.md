
## Configuration

The crawler accepts several parameters:
- `client`: An HTTP client for making requests
- `startURL`: The initial URL to begin crawling from
- `numScrapers`: Number of concurrent scraping workers
- `numValidators`: Number of concurrent external link validators

## Output Indicators

- 🔍 Scraping in progress
- ✅ Successful crawl
- ❌ Failed to process
- 📄 Internal page queued
- 🔗 External link queued
- 🌐 Validating external link
- ⏭️ Skipping already visited URL
- 📊 Final crawl statistics

## Implementation Details

- Uses Go's `sync.WaitGroup` for worker coordination
- Implements thread-safe operations with `sync.RWMutex`
- Employs channels for concurrent communication
- Utilizes `golang.org/x/net/html` for HTML parsing
- Handles relative and absolute URLs using Go's `net/url` package

## Error Handling

The crawler includes robust error handling for:
- Invalid URLs
- Network failures
- HTTP errors
- Malformed HTML
- External link validation


webscraper/
│── cmd/                    # Entry points for CLI/web scraper
│   ├── webscraper/
│   │   ├── main.go         # Main function
│── internal/               # Core business logic
│   ├── crawler/            # Web crawling logic
│   │   ├── crawler.go      # Main recursive crawling logic
│   ├── utils/            # Stores visited URLs
│   │   ├── url_utils.go      # Manages visited links
│   ├── validator/         # Link validation
│   │   ├── link_validator.go    # Checks for valid links
│── go.mod                  # Go module file
│── go.sum                  # Dependencies
│── README.md               # Documentation
