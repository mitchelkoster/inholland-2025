package monitor

import (
	"net/http"
	"time"
)

// Monitor implements Subject
type Monitor struct {
	observers []Observer
	client    *http.Client
}

func NewMonitor() *Monitor {
	return &Monitor{
		client: &http.Client{Timeout: 5 * time.Second},
	}
}

func (m *Monitor) Register(o Observer) {
	m.observers = append(m.observers, o)
}

func (m *Monitor) Deregister(o Observer) {
	for i, obs := range m.observers {
		if obs == o {
			m.observers = append(m.observers[:i], m.observers[i+1:]...)
			break
		}
	}
}

func (m *Monitor) NotifyAll(r Result) {
	for _, o := range m.observers {
		o.Update(r)
	}
}

func (m *Monitor) Check(domains []string) {
	for _, d := range domains {
		status := 0
		resp, err := m.client.Get("http://" + d)
		if err == nil {
			status = resp.StatusCode
			resp.Body.Close()
		}
		m.NotifyAll(Result{Domain: d, Status: status})
	}
}
