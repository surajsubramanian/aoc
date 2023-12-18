package main

import "testing"

func TestFirstLine(t *testing.T) {
	s := []struct {
		line string
		want int
	}{
		{
			line: "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			want: 8,
		},
		{
			line: "Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
			want: 2,
		},
		{
			line: "Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
			want: 2,
		},
		{
			line: "Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
			want: 0,
		},
	}

	for _, tt := range s {
		if got := firstLine(tt.line); got != tt.want {
			t.Errorf("firstLine() = %v, want %v", got, tt.want)
		}
	}
}
