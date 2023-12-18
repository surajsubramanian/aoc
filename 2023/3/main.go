package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var input []string

func init() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input = strings.Split(string(f), "\n")
}

func isSymbol(s string) bool {
	_, err := strconv.Atoi(s)
	return err != nil && s != "."
}

func isNumber(f string) bool {
	_, err := strconv.Atoi(f)
	return err == nil
}

func getNumbers() [][]int {
	pattern := regexp.MustCompile(`\d+`)
	numbers := [][]int{}
	for i, line := range input {
		matches := pattern.FindAllStringIndex(line, -1)
		for _, match := range matches {
			numbers = append(numbers, []int{i, match[0], match[1]})
		}
	}
	return numbers
}

func isPartNumber(lines []string, i, start int, end int) bool {
	for line := i - 1; line <= i+1; line++ {
		if line < 0 || line >= len(lines) {
			continue
		}
		for index := start - 1; index <= end+1; index++ {
			if index < 0 || index >= len(lines[line]) {
				continue
			}
			if isSymbol(string(lines[line][index])) {
				return true
			}
		}
	}
	return false
}

func first() int {
	count := 0
	numbers := getNumbers()
	for _, num := range numbers {
		i, start, end := num[0], num[1], num[2]-1
		if isPartNumber(input, i, start, end) {
			c, err := strconv.Atoi(input[i][start : end+1])
			if err != nil {
				log.Fatal(err)
			}
			count += c
		}
	}
	return count
}

func getGears(f string) [][]int {
	pattern := regexp.MustCompile(`\d+`)
	numbers := [][]int{}
	for i, line := range strings.Split(f, "\n") {
		matches := pattern.FindAllStringIndex(line, -1)
		for _, match := range matches {
			numbers = append(numbers, []int{i, match[0], match[1]})
		}
	}
	return numbers
}

func getNumberFromCoordinates(row, col int) ([]int, error) {
	numbers := getNumbers()
	for _, num := range numbers {
		if num[0] == row && num[1] <= col && num[2] >= col {
			return num, nil
		}
	}
	return []int{}, fmt.Errorf("No number found at row %d and column %d", row, col)
}

func isGear(lines []string, i, start int) int {
	adjNums := [][]int{}
	for line := i - 1; line <= i+1; line++ {
		if line < 0 || line >= len(lines) {
			continue
		}
		for index := start - 1; index <= start+1; index++ {
			if index < 0 || index >= len(lines[line]) {
				continue
			}
			if isNumber(string(lines[line][index])) {
				adjNums = append(adjNums, []int{line, index})
			}
		}
	}
	nums := [][]int{}
	for _, adjNum := range adjNums {
		num, err := getNumberFromCoordinates(adjNum[0], adjNum[1])
		if err != nil {
			log.Fatal(err)
		}
		unique := true
		for i = 0; i < len(nums); i++ {
			if nums[i][0] == num[0] && nums[i][1] == num[1] && nums[i][2] == num[2] {
				unique = false
			}
		}
		if unique {
			nums = append(nums, num)
		}
	}

	if len(nums) == 2 {
		x := []int{}
		for _, num := range nums {
			line, start, end := num[0], num[1], num[2]
			t, _ := strconv.Atoi(string(lines[line][start:end]))
			x = append(x, t)
		}
		return x[0] * x[1]
	}
	return 0
}

func second() int {
	count := 0
	for i := range input {
		for j := range input[i] {
			if string(input[i][j]) == "*" {
				count += isGear(input, i, j)
			}
		}
	}
	return count
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part ", part)

	if part == 1 {
		ans := first()
		fmt.Println("Output: ", ans)
	} else {
		ans := second()
		fmt.Println("Output: ", ans)
	}
}
