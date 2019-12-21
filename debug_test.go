package main

import "testing"

func TestDebug(t *testing.T) {
	debug("%s %d", "test", 123)
}
