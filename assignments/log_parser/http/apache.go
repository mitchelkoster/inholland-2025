package http

import (
	"fmt"
	"strings"
)

type LogEntry struct {
	IP     string
	Method string
	Path   string
	Status string
}

func Parse(line string) (LogEntry, error) {
	/*
		go doc strings.fields
		package strings // import "strings"

		func Fields(s string) []string
			Fields splits the string s around each instance of one or more consecutive
			white space characters
	*/
	parts := strings.Fields(line)

	if len(parts) < 9 {
		return LogEntry{}, fmt.Errorf("invalid log line")
	}

	entry := LogEntry{
		IP:     parts[0],
		Method: parts[5],
		Path:   parts[6],
		Status: parts[8],
	}

	return entry, nil
}
