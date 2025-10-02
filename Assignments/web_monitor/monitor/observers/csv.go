package observers

import (
	"encoding/csv"
	"fmt"
	"os"

	"webmonitor.com/monitor"
)

type CSVObserver struct {
	writer *csv.Writer
	file   *os.File
}

func NewCSVObserver(filename string) (*CSVObserver, error) {
	f, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	w := csv.NewWriter(f)
	w.Write([]string{"Domain", "Status"})
	return &CSVObserver{writer: w, file: f}, nil
}

func (c *CSVObserver) Update(r monitor.Result) {
	if err := c.writer.Write([]string{r.Domain, fmt.Sprintf("%d", r.Status)}); err != nil {
		fmt.Println("CSV write error:", err)
	}
}

func (c *CSVObserver) Close() {
	c.writer.Flush()
	if err := c.writer.Error(); err != nil {
		fmt.Println("CSV flush error:", err)
	}
	c.file.Close()
}
