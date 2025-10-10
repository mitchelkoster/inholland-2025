package observers

import (
	"fmt"

	"webmonitor.com/monitor"
)

// ConsoleObserver
type ConsoleObserver struct{}

func (c ConsoleObserver) Update(r monitor.Result) {
	fmt.Printf("%s (%d)\n", r.Domain, r.Status)
}
