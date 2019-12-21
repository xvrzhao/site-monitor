package main

import (
	"fmt"
	"testing"
	"time"
)

var testR1, testR2 logRecord

func TestLoggerString(t *testing.T) {
	logPush(testR2)
	logPush(testR1)
	fmt.Println(loggerInstance.string())
}

func TestLogPush(t *testing.T) {
	for i := 0; i < 12; i++ {
		logPush(testR1)
	}
	if len(loggerInstance.records) != 10 {
		t.Fail()
		return
	}

	// test mail sending
	fmt.Println("---")
	for i := 0; i < 5; i++ {
		logPush(testR2)
	}
	time.Sleep(time.Second * 3) // waiting for sending mail
}

func init() {
	testR1 = logRecord{
		time:       time.Now(),
		isNormal:   true,
		statusCode: 200,
		isTimeout:  false,
		err:        nil,
	}
	testR2 = logRecord{
		time:       time.Now(),
		isNormal:   false,
		statusCode: 500,
		isTimeout:  false,
		err:        nil,
	}
}
