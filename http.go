package main

import (
	"net/http"
	"time"
)

func newHTTPClient(timeout, responseHeaderTimeout time.Duration) http.Client {
	tr := http.Transport{
		DisableKeepAlives:     true,
		ResponseHeaderTimeout: responseHeaderTimeout,
	}
	client := http.Client{
		Transport: &tr,
		Timeout:   timeout,
	}

	return client
}
