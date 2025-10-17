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
	workerID   int
	site       string
	statusCode int
	body       string
	err        error
}

func websiteWorker(workerID int, jobs <-chan string, results chan<- HttpResult, options *HttpOptions, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done() // Decrease worker count by one once finished

	for website := range jobs {
		// Make request
		resp, err := options.httpClient.Get("http://" + website)
		if err != nil {
			results <- HttpResult{
				workerID:   workerID,
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
			workerID:   workerID,
			site:       website,
			statusCode: resp.StatusCode,
			body:       string(body),
			err:        nil,
		}
	}
}

func CheckAllWebsites(websites []string, options *HttpOptions, workerPoolSize int) <-chan HttpResult {
	var waitGroup sync.WaitGroup

	jobs := make(chan string, workerPoolSize)       // Items to process
	results := make(chan HttpResult, len(websites)) // Results delivered

	// Start workers
	fmt.Println("Worker pool count: ", workerPoolSize)
	for workerID := 0; workerID < WORKER_POOL_SIZE; workerID++ {
		waitGroup.Add(1) // Add active worker count
		go websiteWorker(workerID, jobs, results, options, &waitGroup)
	}

	// Supply workers
	go func() {
		for _, website := range websites {
			jobs <- website
		}
	}()

	// Wait for workers to finish
	go func() {
		waitGroup.Wait() // Wait till active worker count is 0
		close(results)
	}()

	return results
}
