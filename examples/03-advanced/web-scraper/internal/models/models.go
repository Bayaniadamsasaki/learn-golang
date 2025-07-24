package models

import "time"

type Article struct {
	Title       string    `json:"title"`
	URL         string    `json:"url"`
	Description string    `json:"description"`
	Author      string    `json:"author"`
	PublishedAt time.Time `json:"published_at"`
	Tags        []string  `json:"tags"`
}

type ScrapingConfig struct {
	BaseURL     string        `json:"base_url"`
	Selectors   Selectors     `json:"selectors"`
	Delay       time.Duration `json:"delay"`
	Concurrent  int           `json:"concurrent"`
	UserAgent   string        `json:"user_agent"`
}

type Selectors struct {
	Title       string `json:"title"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Date        string `json:"date"`
	Tags        string `json:"tags"`
}

type ScrapingResult struct {
	Articles []Article `json:"articles"`
	Success  int       `json:"success"`
	Failed   int       `json:"failed"`
	Duration string    `json:"duration"`
}
