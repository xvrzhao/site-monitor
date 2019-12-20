package monitor

import (
	"testing"
	"time"
)

func TestNewMonitor(t *testing.T) {
	m := NewMonitor()
	t.Logf("url: %s | interval: %v | client: %v", m.url, m.interval, m.client)
}

func TestMonitorIssue(t *testing.T) {
	m := NewMonitor()
	m.issue(time.Now())
}
