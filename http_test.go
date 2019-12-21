package main

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestNewHTTPClient(t *testing.T) {
	c := newHTTPClient(time.Second*5, time.Second*2)
	if c.Timeout.String() != "5s" {
		t.Fail()
		return
	}
	if tr := c.Transport.(*http.Transport); tr.ResponseHeaderTimeout.String() != "2s" {
		t.Fail()
		return
	}
	resp, err := c.Get("https://jsonplaceholder.typicode.com/users/1")
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
}
