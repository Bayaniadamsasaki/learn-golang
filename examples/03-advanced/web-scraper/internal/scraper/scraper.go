package scraper

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"web-scraper/internal/models"
)

type Scraper struct {
	config models.ScrapingConfig
	client *http.Client
}

func New(config models.ScrapingConfig) *Scraper {
	return &Scraper{
		config: config,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (s *Scraper) ScrapeArticles(urls []string) (*models.ScrapingResult, error) {
	start := time.Now()
	articles := make([]models.Article, 0)
	success, failed := 0, 0
	
	jobs := make(chan string, len(urls))
	results := make(chan models.Article, len(urls))
	errors := make(chan error, len(urls))
	
	var wg sync.WaitGroup
	
	for i := 0; i < s.config.Concurrent; i++ {
		wg.Add(1)
		go s.worker(jobs, results, errors, &wg)
	}
	
	for _, url := range urls {
		jobs <- url
	}
	close(jobs)
	
	go func() {
		wg.Wait()
		close(results)
		close(errors)
	}()
	
	for {
		select {
		case article, ok := <-results:
			if !ok {
				results = nil
			} else {
				articles = append(articles, article)
				success++
			}
		case err, ok := <-errors:
			if !ok {
				errors = nil
			} else {
				fmt.Printf("Error: %v\n", err)
				failed++
			}
		}
		
		if results == nil && errors == nil {
			break
		}
	}
	
	return &models.ScrapingResult{
		Articles: articles,
		Success:  success,
		Failed:   failed,
		Duration: time.Since(start).String(),
	}, nil
}

func (s *Scraper) worker(jobs <-chan string, results chan<- models.Article, errors chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for url := range jobs {
		time.Sleep(s.config.Delay)
		
		article, err := s.scrapeArticle(url)
		if err != nil {
			errors <- err
			continue
		}
		
		results <- article
	}
}

func (s *Scraper) scrapeArticle(url string) (models.Article, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return models.Article{}, err
	}
	
	req.Header.Set("User-Agent", s.config.UserAgent)
	
	resp, err := s.client.Do(req)
	if err != nil {
		return models.Article{}, err
	}
	defer resp.Body.Close()
	
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return models.Article{}, err
	}
	
	article := models.Article{
		URL: url,
	}
	
	if s.config.Selectors.Title != "" {
		article.Title = strings.TrimSpace(doc.Find(s.config.Selectors.Title).First().Text())
	}
	
	if s.config.Selectors.Description != "" {
		article.Description = strings.TrimSpace(doc.Find(s.config.Selectors.Description).First().Text())
	}
	
	if s.config.Selectors.Author != "" {
		article.Author = strings.TrimSpace(doc.Find(s.config.Selectors.Author).First().Text())
	}
	
	if s.config.Selectors.Tags != "" {
		doc.Find(s.config.Selectors.Tags).Each(func(i int, sel *goquery.Selection) {
			tag := strings.TrimSpace(sel.Text())
			if tag != "" {
				article.Tags = append(article.Tags, tag)
			}
		})
	}
	
	article.PublishedAt = time.Now()
	
	return article, nil
}
