package monitor

import (
	"fmt"
)

// ConsoleObserver
type ConsoleObserver struct{}

func (c ConsoleObserver) Update(r Result) {
	fmt.Printf("%s → %d\n", r.Domain, r.Status)
}
