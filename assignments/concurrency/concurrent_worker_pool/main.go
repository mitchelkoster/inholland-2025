package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
)

const FILE = "websites.txt"
const WORKER_POOL_SIZE = 3

func loadWebsites(filename string) []string {
	var websites []string

	// Open file for reading
	fh, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fh.Close()

	// Read file line-by-line
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		websites = append(websites, scanner.Text())
	}

	// Deal with scanning errors
	if scanner.Err() != nil {
		fmt.Printf("Error reading line: %s\n", scanner.Err().Error())
	}

	return websites
}

func main() {
	// Load websites
	websites := loadWebsites(FILE)
	fmt.Println(len(websites), "websites loaded")

	options := HttpOptions{
		httpClient: &http.Client{Timeout: 5 * time.Second},
	}

	// Print websites as they arive
	results := CheckAllWebsites(websites, &options, WORKER_POOL_SIZE)
	for res := range results {
		if res.err != nil {
			fmt.Printf("Worker: %d, Error: %v", res.workerID, res.err)
			continue
		}
		fmt.Printf("Worker: %d, Response: %s (Code: %d, Bytes: %d)\n", res.workerID, res.site, res.statusCode, len(res.body))
	}
}
