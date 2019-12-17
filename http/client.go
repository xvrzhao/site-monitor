package http

import (
	"net/http"
	"time"
)

func NewClient(totalTimeout, respHeaderTimeout time.Duration) http.Client {
	tr := http.Transport{
		DisableKeepAlives:     true,
		ResponseHeaderTimeout: respHeaderTimeout, // see go src
	}
	client := http.Client{
		Transport: &tr,
		Timeout:   totalTimeout, // see go src
	}

	return client
}
