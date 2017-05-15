package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppend(t *testing.T) {

	l := &LinkedList{}

	l.Append("data")
	assert.Equal(t, l.Length(), 1)
}
