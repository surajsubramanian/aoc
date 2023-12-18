package main

import (
	"strings"
	"testing"
)

func TestDoer(t *testing.T) {
	s := "two2seven7"
	x := lineDoer(s)
	if x != 27 {
		t.Errorf("expected 27, got %d", x)
	}
}

func TestReplacer(t *testing.T) {
	s := []struct {
		input  string
		starts string
		ends   string
	}{
		{
			"two2seven7",
			"two2",
			"seven7",
		},
		{
			"eightwothree",
			"eight8",
			"3three",
		},
	}
	for _, tt := range s {
		if got := Replacer(tt.input); !strings.HasPrefix(got, tt.starts) || !strings.HasSuffix(got, tt.ends) {
			t.Errorf("Replacer() = %v should start with %v and end with %v", got, tt.starts, tt.ends)
		}
	}
}

func TestLineCalibrationValue(t *testing.T) {
	s := []struct {
		input  string
		output int
	}{
		{
			"two2seven7",
			27,
		},
		{
			"eightwothree",
			83,
		},
	}
	for _, tt := range s {
		got := LineCalibrationValue(Replacer(tt.input))
		if got != tt.output {
			t.Errorf("LineDoer() = %v, want %v", got, tt.output)
		}
	}
}
