package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func lineDoer(t string) int {
	start, end := 0, len(t)-1
	digits := "0123456789"
	for i := 0; i < len(t); i++ {
		l := t[i]
		if strings.Contains(digits, string(l)) {
			start, _ = strconv.Atoi(string(l))
			break
		}
	}
	for i := len(t) - 1; i >= 0; i-- {
		l := t[i]
		if strings.Contains(digits, string(l)) {
			end, _ = strconv.Atoi(string(l))
			break
		}
	}
	return (start*10 + end)
}

func main() {
	file, err := os.Open("temp.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ans := 0
	for scanner.Scan() {
		ans += lineDoer(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	log.Println(ans)
}
