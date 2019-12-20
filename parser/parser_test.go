package parser

import "testing"

func TestDebug(t *testing.T) {
	Debug("test: %s, %s", "Hello", "world")
}
