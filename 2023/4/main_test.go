package main

import (
	"testing"
)

func TestFirstLine(t *testing.T) {
	s := []struct {
		line     string
		winWant  []int
		haveWant []int
	}{
		{
			line:     "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			winWant:  []int{41, 48, 83, 86, 17},
			haveWant: []int{83, 86, 6, 31, 17, 9, 48, 53},
		},
	}

	for _, tt := range s {
		win, have := lineParse(tt.line)
		for i := 0; i < len(tt.winWant); i++ {
			if win[i] != tt.winWant[i] {
				t.Errorf("win = %v, winWant %v", win, tt.winWant)
			}
		}
		for i := 0; i < len(tt.haveWant); i++ {
			if have[i] != tt.haveWant[i] {
				t.Errorf("have = %v, haveWant %v", have, tt.haveWant)
			}
		}
	}
}
