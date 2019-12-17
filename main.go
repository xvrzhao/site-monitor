package main

import (
	m "github.com/xvrzhao/site-monitor/monitor"
)

func main() {
	monitor := m.NewMonitor()
	monitor.Work()
}
