package main

import "testing"

func TestSendMail(t *testing.T) {
	sendMail("test1")
	sendMail("test2") // should not send
}
