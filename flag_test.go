package main

import (
	"testing"
	"time"
)

func TestConvertDuration(t *testing.T) {
	var du time.Duration
	if convertDuration(10, &du); du.String() != "10ms" {
		t.Fail()
	}
}
