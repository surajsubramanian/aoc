package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var lines []string

func lineDoer(t string) int {
	digits := "0123456789"
	start := t[strings.IndexAny(t, digits)] - '0'
	end := t[strings.LastIndexAny(t, digits)] - '0'
	return int(start)*10 + int(end)
}

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

func first() int {
	result := 0
	for _, line := range lines {
		if line != "" {
			result += lineDoer(line)
		}
	}
	return result
}

func second() int {
	result := 0
	for _, line := range lines {
		line = Replacer(line)
		if line != "" {
			result += LineCalibrationValue(line)
		}
	}
	return result
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()

	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines = strings.Split(string(file), "\n")

	if part == 1 {
		ans := first()
		fmt.Println("Output: ", ans)
	} else {
		ans := second()
		fmt.Println("Output: ", ans)
	}
}
