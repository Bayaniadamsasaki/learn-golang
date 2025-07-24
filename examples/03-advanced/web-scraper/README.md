# Web Scraper

Concurrent web scraper untuk mengekstrak data dari website dengan goroutines dan channels.

## Fitur

- Concurrent scraping dengan goroutines
- Rate limiting untuk menghindari server overload
- JSON output untuk hasil scraping
- Support untuk multiple URLs
- Error handling dan retry mechanism
- Configurable timeout dan delays

## Setup

```bash
go mod tidy
go run main.go
```

## Konfigurasi

Edit file `config.json` untuk mengatur:

- URLs yang akan di-scrape
- Jumlah worker goroutines
- Delay antar request
- Timeout settings

## Contoh config.json

```json
{
  "urls": ["https://example.com", "https://httpbin.org/json"],
  "workers": 5,
  "delay_ms": 1000,
  "timeout_seconds": 30
}
```

## Output

Hasil scraping disimpan dalam file `results.json` dengan format:

```json
{
  "url": "https://example.com",
  "title": "Example Domain",
  "status_code": 200,
  "content_length": 1256,
  "scraped_at": "2023-01-01T00:00:00Z"
}
```

## Pembelajaran

Project ini mengajarkan:

- Goroutines untuk concurrency
- Channels untuk komunikasi
- HTTP client dengan timeout
- HTML parsing dengan goquery
- Rate limiting patterns
- Error handling dalam concurrent code
- JSON configuration dan output
