package handler

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintStdOut(t *testing.T) {
	var buffer bytes.Buffer

	expectedOutput := "line1\nline2\nline3"
	PrintStdOut(&buffer, expectedOutput)
	actualOutput := buffer.String()
	assert.Equal(t, expectedOutput+"\n", actualOutput)
}
