package main

import (
	"slices"
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test(t *testing.T) {
	array := []string{"a", "b", "d"}

	assert.True(t, slices.Contains(array, "b"))
	assert.False(t, slices.Contains(array, "c"))
}