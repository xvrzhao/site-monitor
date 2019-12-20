package logger

import (
	"testing"
	"time"
)

var (
	r = Record{
		Time:       time.Now(),
		IsNormal:   false,
		StatusCode: 502,
		IsTimeout:  false,
		Error:      nil,
	}
	r2 = Record{
		Time:       time.Now(),
		IsNormal:   true,
		StatusCode: 200,
		IsTimeout:  false,
		Error:      nil,
	}
)

func TestPush(t *testing.T) {
	Push(r)

	if lr := loggerInstance.records[len(loggerInstance.records)-1]; lr.StatusCode != 502 {
		t.Fail()
		return
	} else if len(loggerInstance.records) != 1 {
		t.Fail()
		return
	} else {
		t.Logf("time: %v | isNormal: %v | code: %d | timeout: %v | err: %v",
			lr.Time, lr.IsNormal, lr.StatusCode, lr.IsTimeout, lr.Error)
	}
}

// TestPush2 tests max length of loggerInstance.records.
func TestPush2(t *testing.T) {
	for i := 0; i < 15; i++ {
		Push(r2)
	}
	if len(loggerInstance.records) != 10 {
		t.Fail()
	}
}

// TestPush3 tests mail sending.
func TestPush3(t *testing.T) {
	for i := 0; i < 5; i++ {
		Push(r)
	}
	time.Sleep(5 * time.Second)
}
