package main

import (
	"bufio"
	"fmt"
	"os"

	"parser.com/http"
)

const VERSION = "1.0.0"

func parseLogEntries(file string) ([]http.LogEntry, error) {
	var logEntries []http.LogEntry

	// Open file (read only)
	fh, err := os.Open("logs/access.log")
	if err != nil {
		fmt.Println("Error while opening file:", err)
		return []http.LogEntry{}, err
	}
	defer fh.Close()

	// Read file line by line (could be large)
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		entry, err := http.Parse(scanner.Text())

		// Return parial when error
		if err != nil {
			fmt.Println("Error while parsing:", err)
			return logEntries, err
		}

		logEntries = append(logEntries, entry)
	}

	return logEntries, nil
}

func main() {
	fmt.Printf("Log parser v%s\n", VERSION)

	logEntries, err := parseLogEntries("logs/access.log")
	if err != nil {
		fmt.Println("Error parsing log enrties:", err)
	}

	for _, entry := range logEntries {
		fmt.Printf("Accessed from: %s\n", entry.IP)
	}
}
