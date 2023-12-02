package main

import (
	"log"
	"os"
	"strings"
)

func LineCalibrationValue(t string) int {
	digits := "0123456789"
	start := t[strings.IndexAny(t, digits)] - '0'
	end := t[strings.LastIndexAny(t, digits)] - '0'
	return int(start)*10 + int(end)
}

func Replacer(line string) string {
	dict := map[string]string{
		"zero":  "zero0zero",
		"one":   "one1one",
		"two":   "two2two",
		"three": "three3three",
		"four":  "four4four",
		"five":  "five5five",
		"six":   "six6six",
		"seven": "seven7seven",
		"eight": "eight8eight",
		"nine":  "nine9nine",
	}

	for k, v := range dict {
		line = strings.ReplaceAll(line, k, v)
	}
	return line
}

func main() {
	file, _ := os.ReadFile("temp.txt")
	result := 0

	for _, line := range strings.Split(string(file), "\n") {
		line = Replacer(line)
		if line != "" {
			result += LineCalibrationValue(line)
		}
	}
	log.Println(result)
}
