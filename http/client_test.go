package http

import (
	"net/http"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	c := NewClient(time.Second*5, time.Second*2)
	if tr := c.Transport.(*http.Transport); c.Timeout != time.Second*5 || tr.ResponseHeaderTimeout != time.Second*2 {
		t.Fail()
	}
}