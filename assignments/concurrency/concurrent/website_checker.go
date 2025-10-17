package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

type HttpOptions struct {
	httpClient *http.Client
}

type HttpResult struct {
	site       string
	statusCode int
	body       string
	err        error
}

func checkWebsite(website string, results chan HttpResult, options HttpOptions, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	// Make response
	resp, err := options.httpClient.Get("http://" + website)
	if err != nil {
		results <- HttpResult{
			site:       website,
			statusCode: -1,
			body:       "",
			err:        err,
		}
		return
	}
	defer resp.Body.Close()

	// Parse body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read body: ", err)
	}

	// Send result to channel
	results <- HttpResult{
		site:       website,
		statusCode: resp.StatusCode,
		body:       string(body),
		err:        nil,
	}
}

func CheckAllWebsites(websites []string, options *HttpOptions) <-chan HttpResult {
	var waitGroup sync.WaitGroup
	results := make(chan HttpResult, len(websites))

	// Visit websites
	for _, web := range websites {
		waitGroup.Add(1)
		go checkWebsite(web, results, *options, &waitGroup)
	}

	// Wait for all thread to finish
	go func() {
		waitGroup.Wait()
		close(results)
	}()

	return results
}
