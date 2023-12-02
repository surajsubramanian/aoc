package main

import (
	"log"
	"os"
	"strings"
)

func lineDoer(t string) int {
	digits := "0123456789"
	start := t[strings.IndexAny(t, digits)] - '0'
	end := t[strings.LastIndexAny(t, digits)] - '0'
	return int(start)*10 + int(end)
}

func main() {
	file, _ := os.ReadFile("temp.txt")
	result := 0
	for _, line := range strings.Split(string(file), "\n") {
		if line != "" {
			result += lineDoer(line)
		}
	}
	log.Println(result)
}
