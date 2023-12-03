package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	file, _ := os.ReadFile("input.txt")
	ans := 0
	for _, line := range strings.Split(string(file), "\n") {
		if line == "" {
			continue
		}
		maxColors := GetMaxEach(line)
		ans += maxColors["red"] * maxColors["green"] * maxColors["blue"]
	}
	log.Println(int(ans))
}
