package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"webmonitor.com/monitor"
)

func main() {
	// Read top 100 website file
	f, err := os.Open("top_100.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer f.Close()

	// Read domains one-by-one
	var domains []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		domains = append(domains, strings.TrimSpace(scanner.Text()))
	}

	// Setup monitor
	m := monitor.NewMonitor()

	// Add console observer
	m.Register(monitor.ConsoleObserver{})

	// Add CSV observer
	csvObs, err := monitor.NewCSVObserver("results.csv")
	if err != nil {
		fmt.Println("Error creating CSV observer:", err)
		return
	}
	defer csvObs.Close()
	m.Register(csvObs)

	// Run the checks
	m.Check(domains)
}
