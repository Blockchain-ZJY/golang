package main

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	var (
		input    = 5
		expected = 120
	)
	actual := Factorial(input)
	if actual != expected {
		t.Errorf("Factorial(%d) = %d; expected %d", input, actual, expected)
	}

}
