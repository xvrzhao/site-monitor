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

func TestConvertRecipient(t *testing.T) {
	var addrs []string
	if convertRecipient("xvrzhao@gmail.com ", &addrs); len(addrs) != 1 || addrs[0] != "xvrzhao@gmail.com" {
		t.Fail()
		return
	}
	var addrs1 []string
	convertRecipient("xvrzhao@gmail.com, tom@example.com", &addrs1)
	if  len(addrs1) != 2 || addrs1[0] != "xvrzhao@gmail.com" || addrs1[1] != "tom@example.com" {
		t.Fail()
	}
}
