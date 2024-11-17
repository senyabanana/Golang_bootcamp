package main

import (
	"testing"
	"time"
)

func TestCrawlWeb(t *testing.T) {
	urls := make(chan string, 3)
	done := make(chan struct{})
	urls <- "https://example.com"
	urls <- "https://golang.org"
	close(urls)

	results := crawlWeb(urls, done)

	count := 0
	for range results {
		count++
	}

	if count < 2 {
		t.Errorf("Expected at least 2 pages to be crawled, got %d", count)
	}
}

func TestCrawlWebCancel(t *testing.T) {
	urls := make(chan string, 5)
	done := make(chan struct{})

	go func() {
		urls <- "https://example.com"
		time.Sleep(500 * time.Millisecond)
		close(urls)
	}()

	results := crawlWeb(urls, done)
	select {
	case <-results:
		// Ожидаем, что результат придет до того, как тест завершится
	case <-time.After(3 * time.Second):
		t.Errorf("Test timed out, cancellation may not be working correctly")
	}
}
