package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Config struct {
	URLs           []string `json:"urls"`
	Workers        int      `json:"workers"`
	DelayMs        int      `json:"delay_ms"`
	TimeoutSeconds int      `json:"timeout_seconds"`
}

type ScrapedData struct {
	URL           string    `json:"url"`
	Title         string    `json:"title"`
	StatusCode    int       `json:"status_code"`
	ContentLength int64     `json:"content_length"`
	ScrapedAt     time.Time `json:"scraped_at"`
	Error         string    `json:"error,omitempty"`
}

type Scraper struct {
	config Config
	client *http.Client
}

func main() {
	config := loadConfig()
	scraper := NewScraper(config)
	
	results := scraper.ScrapeURLs(config.URLs)
	
	err := saveResults(results)
	if err != nil {
		log.Printf("Error saving results: %v", err)
	}
	
	fmt.Printf("Scraping completed. Found %d results.\n", len(results))
}

func loadConfig() Config {
	defaultConfig := Config{
		URLs: []string{
			"https://httpbin.org/json",
			"https://jsonplaceholder.typicode.com/posts/1",
		},
		Workers:        3,
		DelayMs:        1000,
		TimeoutSeconds: 30,
	}
	
	file, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Println("Using default config, creating config.json")
		data, _ := json.MarshalIndent(defaultConfig, "", "  ")
		os.WriteFile("config.json", data, 0644)
		return defaultConfig
	}
	
	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Printf("Error parsing config: %v", err)
		return defaultConfig
	}
	
	return config
}

func NewScraper(config Config) *Scraper {
	return &Scraper{
		config: config,
		client: &http.Client{
			Timeout: time.Duration(config.TimeoutSeconds) * time.Second,
		},
	}
}

func (s *Scraper) ScrapeURLs(urls []string) []ScrapedData {
	urlChannel := make(chan string, len(urls))
	resultChannel := make(chan ScrapedData, len(urls))
	
	for _, url := range urls {
		urlChannel <- url
	}
	close(urlChannel)
	
	var wg sync.WaitGroup
	
	for i := 0; i < s.config.Workers; i++ {
		wg.Add(1)
		go s.worker(urlChannel, resultChannel, &wg)
	}
	
	go func() {
		wg.Wait()
		close(resultChannel)
	}()
	
	var results []ScrapedData
	for result := range resultChannel {
		results = append(results, result)
	}
	
	return results
}

func (s *Scraper) worker(urlChannel <-chan string, resultChannel chan<- ScrapedData, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for url := range urlChannel {
		result := s.scrapeURL(url)
		resultChannel <- result
		
		if s.config.DelayMs > 0 {
			time.Sleep(time.Duration(s.config.DelayMs) * time.Millisecond)
		}
	}
}

func (s *Scraper) scrapeURL(url string) ScrapedData {
	result := ScrapedData{
		URL:       url,
		ScrapedAt: time.Now(),
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.config.TimeoutSeconds)*time.Second)
	defer cancel()
	
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		result.Error = fmt.Sprintf("Error creating request: %v", err)
		return result
	}
	
	req.Header.Set("User-Agent", "GoScraper/1.0")
	
	resp, err := s.client.Do(req)
	if err != nil {
		result.Error = fmt.Sprintf("Error making request: %v", err)
		return result
	}
	defer resp.Body.Close()
	
	result.StatusCode = resp.StatusCode
	result.ContentLength = resp.ContentLength
	
	if resp.StatusCode != http.StatusOK {
		result.Error = fmt.Sprintf("HTTP %d", resp.StatusCode)
		return result
	}
	
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		result.Error = fmt.Sprintf("Error parsing HTML: %v", err)
		return result
	}
	
	title := doc.Find("title").First().Text()
	if title == "" {
		title = "No title found"
	}
	result.Title = title
	
	fmt.Printf("Scraped: %s - %s\n", url, title)
	
	return result
}

func saveResults(results []ScrapedData) error {
	data, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return err
	}
	
	return os.WriteFile("results.json", data, 0644)
}
