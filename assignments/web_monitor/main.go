package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"webmonitor.com/monitor"
	"webmonitor.com/monitor/observers"
)

func top1000() ([]string, error) {
	// Read top 100 website file
	fh, err := os.Open("top_100.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return []string{}, err
	}
	defer fh.Close()

	// Read domains one-by-one
	var domains []string
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		domains = append(domains, strings.TrimSpace(scanner.Text()))
	}

	return domains, nil
}

func main() {
	// Read file
	domains, err := top1000()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Setup monitor
	m := monitor.NewMonitor()

	// Create CSV observer
	csvObs, err := observers.NewCSVObserver("results.csv")
	if err != nil {
		fmt.Println("Error creating CSV observer:", err)
		return
	}
	defer csvObs.Close()

	// Register observers
	m.Register(observers.ConsoleObserver{})
	m.Register(csvObs)

	// Run the checks
	m.Check(domains)
}
