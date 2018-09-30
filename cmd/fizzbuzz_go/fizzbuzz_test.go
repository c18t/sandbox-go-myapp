package main

import (
	"strings"
	"testing"
)

func TestFizz(t *testing.T) {
	if !strings.EqualFold("FizzBuzz", FizzBuzz(15)) {
		t.Errorf("15 != FizzBuzz")
	}
	if !strings.EqualFold("Fizz", FizzBuzz(3)) {
		t.Errorf("3 != Fizz")
	}
	if !strings.EqualFold("Buzz", FizzBuzz(5)) {
		t.Errorf("5 != Buzz")
	}
	if !strings.EqualFold("2", FizzBuzz(2)) {
		t.Errorf("2 != 2")
	}
}
