package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {
	var (
		input    = 5
		expected = 120
	)
	actual := Factorial(input)
	assert.Equal(t, expected, actual)
}
