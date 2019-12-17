package logger

import (
	"fmt"
	"github.com/xvrzhao/site-monitor/mailer"
	"github.com/xvrzhao/site-monitor/parser"
	"sync"
	"time"
)

var (
	loggerInstance logger
)

type Record struct {
	time.Time
	IsNormal   bool
	StatusCode int
	IsTimeout  bool
	Error      error
}

type logger struct {
	sync.Mutex
	records []Record
}

func (l *logger) string() string {
	var msg string
	for _, record := range l.records {
		msg += fmt.Sprintf("- %v normal: %v | status code: %d | timeout: %v | error: %v\n",
			record.Time, record.IsNormal, record.StatusCode, record.IsTimeout, record.Error)
	}
	return msg
}

func Push(record Record) {
	loggerInstance.Lock()
	defer loggerInstance.Unlock()

	if len(loggerInstance.records) >= 10 {
		loggerInstance.records = loggerInstance.records[1:]
	}
	loggerInstance.records = append(loggerInstance.records, record)

	check()
}

func check() {
	var unNormal int
	for _, r := range loggerInstance.records {
		if !r.IsNormal {
			unNormal++
		}
	}
	if unNormal >= 5 {
		parser.Debug("* Site status: unhealthy\n")
		go mailer.Send(loggerInstance.string())
	} else if float64(unNormal)/float64(len(loggerInstance.records)) >= 0.5 {
		parser.Debug("* Site status: unhealthy\n")
	} else {
		parser.Debug("* Site status: healthy\n")
	}
}

func init() {
	loggerInstance.records = make([]Record, 0, 10)
}
