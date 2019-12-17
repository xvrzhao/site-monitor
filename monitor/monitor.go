package monitor

import (
	monitorHTTP "github.com/xvrzhao/site-monitor/http"
	"github.com/xvrzhao/site-monitor/logger"
	"github.com/xvrzhao/site-monitor/parser"
	"net/http"
	"net/url"
	"time"
)

type Monitor struct {
	url      string
	interval time.Duration
	client   http.Client
}

func NewMonitor() Monitor {
	m := Monitor{
		url:      parser.MonitorURL,
		interval: parser.MonitorInterval,
	}
	m.client = monitorHTTP.NewClient(parser.MonitorTotalTimeout, parser.MonitorRespHeaderTimeout)

	parser.Debug("* Monitor instance has been created: <url: %s, interval: %v, total-timeout: %v, header-timeout: %v>\n",
		parser.MonitorURL, parser.MonitorInterval, parser.MonitorTotalTimeout, parser.MonitorRespHeaderTimeout)

	return m
}

func (m *Monitor) Work() {
	ticker := time.NewTicker(m.interval)
	for taskTime := range ticker.C {
		go m.issue(taskTime)
	}
}

func (m *Monitor) issue(taskTime time.Time) {
	parser.Debug("* Issuing a detection signal to %s ...\n", parser.MonitorURL)
	resp, err := m.client.Get(m.url)
	if err != nil {
		urlErr, _ := err.(*url.Error)
		if urlErr.Timeout() {
			// response timeout error
			parser.Debug("* Result (timeout): %s\n", urlErr.Error())
			r := logger.Record{
				Time:      taskTime,
				IsTimeout: true,
				Error:     urlErr,
			}
			logger.Push(r)
		} else {
			// other errors
			parser.Debug("* Result (error): %s\n", urlErr.Error())
			r := logger.Record{
				Time:  taskTime,
				Error: urlErr,
			}
			logger.Push(r)
		}
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		// non-2xx
		parser.Debug("* Result (non-2xx): status code %d\n", resp.StatusCode)
		r := logger.Record{
			Time:       taskTime,
			StatusCode: resp.StatusCode,
		}
		logger.Push(r)
	} else {
		// 200
		parser.Debug("* Result (normal): status code %d\n", resp.StatusCode)
		r := logger.Record{
			Time:       taskTime,
			IsNormal:   true,
			StatusCode: 200,
		}
		logger.Push(r)
	}
}
