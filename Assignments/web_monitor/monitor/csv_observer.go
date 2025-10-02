package monitor

import (
	"encoding/csv"
	"fmt"
	"os"
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

func (c *CSVObserver) Update(r Result) {
	c.writer.Write([]string{r.Domain, fmt.Sprintf("%d", r.Status)})
}

func (c *CSVObserver) Close() {
	c.writer.Flush()
	c.file.Close()
}
