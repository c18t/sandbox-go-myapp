package main

import (
	"bytes"
	"strings"
	"testing"
)

var buffer *bytes.Buffer

func init() {
	buffer = &bytes.Buffer{}
	writer = buffer
}

func TestMain(t *testing.T) {
	main()

	result := buffer.String()
	lines := strings.Split(result, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	if len(lines) != 100 {
		t.Errorf("default output line count is not 100")
	}
}
