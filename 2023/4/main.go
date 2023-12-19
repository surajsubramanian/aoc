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

func lineParse(line string) ([]int, []int) {
	use := strings.Split(strings.Split(line, ": ")[1], "|")
	wAll, hAll := strings.TrimSpace(use[0]), strings.TrimSpace(use[1])

	win, have := []int{}, []int{}
	for _, w := range strings.Split(wAll, " ") {
		temp := strings.TrimSpace(w)
		if temp != "" {
			t, _ := strconv.Atoi(temp)
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
	return win, have
}

func getChosen(line string) int {
	chosen := []int{}
	if strings.TrimSpace(line) != "" {
		win, have := lineParse(line)
		for _, h := range have {
			if slices.Contains(win, h) {
				chosen = append(chosen, h)
			}
		}
		return len(chosen)
	}
	return 0
}

func first() int {
	result := 0
	for _, line := range lines {
		chosen := getChosen(line)
		if chosen == 0 {
			result += 0
		}
		result += int(math.Pow(float64(2), float64(chosen-1)))
	}
	return result
}

func second() int {
	rList := []int{}
	for i := 0; i < len(lines); i++ {
		rList = append(rList, 1)
	}
	for i := 0; i < len(lines); i++ {
		chosen := getChosen(lines[i])
		for j := i + 1; j < i+1+chosen; j++ {
			if j >= len(lines) {
				continue
			}
			rList[j] += rList[i]
		}
	}
	result := 0
	for _, r := range rList {
		result += r
	}
	return result
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	tempLines := strings.Split(string(file), "\n")
	for _, line := range tempLines {
		if strings.TrimSpace(line) != "" {
			lines = append(lines, line)
		}
	}

	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()

	if part == 1 {
		result := first()
		fmt.Println("Output: ", result)
	} else {
		result := second()
		fmt.Println("Output: ", result)
	}
}
