package main

import (
	"testing"
	"time"
)

var testM monitor

func TestDetect(t *testing.T) {
	testM.detect(time.Now())
	if len(loggerInstance.records) != 1 {
		t.Fail()
	}
}

func init() {
	testM = newMonitor()
}