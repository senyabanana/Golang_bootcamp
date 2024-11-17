package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

func crawlWeb(urls <-chan string, done <-chan struct{}) <-chan string {
	out := make(chan string)
	sem := make(chan struct{}, 8)
	wg := sync.WaitGroup{}

	go func() {
		for url := range urls {
			select {
			case <-done:
				close(out)
				return
			default:
				sem <- struct{}{}
				wg.Add(1)
				go func(url string) {
					defer wg.Done()
					defer func() { <-sem }()

					resp, err := http.Get(url)
					if err != nil {
						log.Println("error fetching url", err)
						out <- fmt.Sprintf("Error: %s", err)
						return
					}
					defer resp.Body.Close()

					body, err := io.ReadAll(resp.Body)
					if err != nil {
						log.Println("error reading response body", err)
						out <- fmt.Sprintf("Error reading body: %s", url)
						return
					}

					out <- string(body)
				}(url)
			}
		}
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	urls := make(chan string)
	done := make(chan struct{})

	go func() {
		defer close(urls)
		urls <- "https://example.com"
		urls <- "https://example.org"
		//urls <- "https://example.net"
		//urls <- "https://google.com"
		//urls <- "https://youtube.com"
	}()

	results := crawlWeb(urls, done)

	for result := range results {
		fmt.Println(result)
	}

	close(done)
}
