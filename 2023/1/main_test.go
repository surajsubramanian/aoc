package main

import (
	"testing"
)

func TestDoer(t *testing.T) {
	s := "two2seven7"
	x := lineDoer(s)
	if x != 27 {
		t.Errorf("expected 27, got %d", x)
	}
}
