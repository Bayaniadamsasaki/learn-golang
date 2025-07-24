package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"web-scraper/internal/models"
	"web-scraper/internal/scraper"
)

func main() {
	config := models.ScrapingConfig{
		BaseURL: "https://example.com",
		Selectors: models.Selectors{
			Title:       "h1",
			Description: "meta[name='description']",
			Author:      ".author",
			Tags:        ".tags a",
		},
		Delay:      time.Second,
		Concurrent: 3,
		UserAgent:  "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	}

	urls := []string{
		"https://golang.org/doc/",
		"https://pkg.go.dev/",
		"https://go.dev/blog/",
		"https://github.com/golang/go",
	}

	s := scraper.New(config)
	
	fmt.Println("Starting web scraping...")
	result, err := s.ScrapeArticles(urls)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Scraping completed!\n")
	fmt.Printf("Success: %d, Failed: %d\n", result.Success, result.Failed)
	fmt.Printf("Duration: %s\n", result.Duration)

	output, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(output))

	err = saveToFile(result, "scraped_data.json")
	if err != nil {
		fmt.Printf("Error saving to file: %v\n", err)
	} else {
		fmt.Println("Data saved to scraped_data.json")
	}
}

func saveToFile(result *models.ScrapingResult, filename string) error {
	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
