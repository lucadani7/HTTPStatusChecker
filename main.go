package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func checkStatus(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", url, err)
		return
	}

	defer resp.Body.Close()

	fmt.Printf("%s -> %d %s\n", url, resp.StatusCode, http.StatusText(resp.StatusCode))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter URLs (one per line). Type 'done' to start checking:")

	var urls []string

	for scanner.Scan() {
		url := strings.TrimSpace(scanner.Text())
		if url == "done" {
			break
		}
		if url != "" {
			urls = append(urls, url)
		}
	}

	if len(urls) == 0 {
		fmt.Println("No URLs entered. Exiting.")
		return
	}

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go checkStatus(url, &wg)
	}

	wg.Wait()
}
