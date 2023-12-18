package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

var lines []string

func firstLine(line string) int {
	use := strings.Split(strings.Split(line, ": ")[1], "|")
	wAll, hAll := strings.TrimSpace(use[0]), strings.TrimSpace(use[1])

	win, have := []int{}, []int{}
	chosen := []int{}
	for _, w := range strings.Split(wAll, " ") {
		temp := strings.TrimSpace(w)
		if temp != "" {
			t, _ := strconv.Atoi(strings.TrimSpace(w))
			win = append(win, t)
		}
	}
	for _, h := range strings.Split(hAll, " ") {
		temp := strings.TrimSpace(h)
		if temp != "" {
			t, _ := strconv.Atoi(temp)
			have = append(have, t)
		}
	}
	for _, h := range have {
		if slices.Contains(win, h) {
			chosen = append(chosen, h)
		}
	}
	if len(chosen) == 0 {
		return 0
	}
	return int(math.Pow(float64(2), float64(len(chosen)-1)))
}

func first() int {
	result := 0
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			result += firstLine(line)
		}
	}
	return result
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines = strings.Split(string(file), "\n")

	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()

	result := first()
	fmt.Println("Output: ", result)
}
