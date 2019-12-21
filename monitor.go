package main

import (
	"net/http"
	"net/url"
	"time"
)

type monitor struct {
	url    string
	cycle  time.Duration
	client http.Client
}

func newMonitor() monitor {
	m := monitor{
		url:   urlFlag,
		cycle: cycleFlag,
	}
	m.client = newHTTPClient(timeoutFlag, headerTimeoutFlag)

	debug("* Monitor instance has been created: <url: %s, cycle: %v, timeout: %v, header timeout: %v>",
		urlFlag, cycleFlag, timeoutFlag, headerTimeoutFlag)

	return m
}

func (m *monitor) work() {
	ticker := time.NewTicker(m.cycle)
	for taskTime := range ticker.C {
		go m.detect(taskTime)
	}
}

func (m *monitor) detect(taskTime time.Time) {
	debug("* Issuing a detection signal to %s ...", urlFlag)
	resp, err := m.client.Get(m.url)
	if err != nil {
		urlErr := err.(*url.Error)
		if urlErr.Timeout() {
			// response timeout error
			debug("* Result (timeout): %v", urlErr)
			r := logRecord{
				time:      taskTime,
				isTimeout: true,
				err:       urlErr,
			}
			logPush(r)
		} else {
			// other errors
			debug("* Result (error): %v", urlErr)
			r := logRecord{
				time: taskTime,
				err:  urlErr,
			}
			logPush(r)
		}
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		// non-2xx
		debug("* Result (non-2xx): status code %d", resp.StatusCode)
		r := logRecord{
			time:       taskTime,
			statusCode: resp.StatusCode,
		}
		logPush(r)
	} else {
		// 200
		debug("* Result (normal): status code %d", resp.StatusCode)
		r := logRecord{
			time:       taskTime,
			isNormal:   true,
			statusCode: 200,
		}
		logPush(r)
	}
}
