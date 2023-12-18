package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var lines []string

func GetMaxEach(line string) map[string]int {
	draws := strings.Split(strings.Split(line, ":")[1], ";")
	cubes := map[string]int{
		"green": 0,
		"red":   0,
		"blue":  0,
	}
	for i := 0; i < len(draws); i++ {
		for _, colorAndVal := range strings.Split(draws[i], ",") {
			for k, v := range cubes {
				if strings.Contains(colorAndVal, k) {
					curV, _ := strconv.Atoi(strings.TrimSpace(strings.Trim(colorAndVal, k)))
					if curV > v {
						cubes[k] = curV
					}
				}
			}
		}
	}
	return cubes
}

func first() int {
	red, green, blue := 12, 13, 14
	result := 0
	for i, line := range lines {
		if line == "" {
			continue
		}
		maxColors := GetMaxEach(line)
		if maxColors["red"] <= red && maxColors["green"] <= green && maxColors["blue"] <= blue {
			result += i + 1
		}
	}
	return result
}

func second() int {
	result := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		maxColors := GetMaxEach(line)
		result += maxColors["red"] * maxColors["green"] * maxColors["blue"]
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
		result := first()
		fmt.Println("Output: ", result)
	} else {
		result := second()
		fmt.Println("Output: ", result)
	}
}
