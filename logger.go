package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	loggerInstance logger
)

type logRecord struct {
	time       time.Time
	isNormal   bool
	statusCode int
	isTimeout  bool
	err        error
}

type logger struct {
	sync.Mutex
	records []logRecord
}

func (l *logger) string() string {
	var msg string
	for _, record := range l.records {
		msg += fmt.Sprintf("* %s\tnormal: %v\tstatus code: %d\ttimeout: %v\terror: %v\n",
			record.time.Format("2006-01-02 15:04:05"), record.isNormal, record.statusCode, record.isTimeout, record.err)
	}
	return msg
}

func logPush(record logRecord) {
	loggerInstance.Lock()
	defer loggerInstance.Unlock()

	if len(loggerInstance.records) >= 10 {
		loggerInstance.records = loggerInstance.records[1:]
	}
	loggerInstance.records = append(loggerInstance.records, record)

	logCheck()
}

func logCheck() {
	var unNormal int
	for _, r := range loggerInstance.records {
		if !r.isNormal {
			unNormal++
		}
	}
	if unNormal >= 5 {
		debug("* Site status: unhealthy")
		go sendMail(loggerInstance.string())
	} else if float64(unNormal)/float64(len(loggerInstance.records)) > 0.5 {
		debug("* Site status: unhealthy")
	} else {
		debug("* Site status: healthy")
	}
}

func init() {
	loggerInstance.records = make([]logRecord, 0, 10)
}
