package main

import (
	"reflect"
	"testing"
)

func TestGetMaxEach(t *testing.T) {
	s := []struct {
		input string
		want  map[string]int
	}{
		{
			"Game 43: 4 green, 6 red, 9 blue; 4 green, 3 red, 18 blue; 6 green, 7 blue; 4 red, 7 blue; 8 blue, 7 green, 1 red; 5 red, 14 blue",
			map[string]int{
				"green": 7,
				"red":   6,
				"blue":  18,
			},
		},
	}

	for _, tt := range s {
		if got := GetMaxEach(tt.input); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("GetMaxEach() = %v, want %v", got, tt.want)
		}
	}

}
