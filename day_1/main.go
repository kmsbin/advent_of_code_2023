package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	problem1()
	problem2()
}

func problem1() {
	lines := readFileLines()
	var acum int
	for _, line := range lines {
		var start, end rune = 0, 0
		for _, char := range line {
			if unicode.IsDigit(char) {
				if start == 0 {
					start = char
				}
				end = char
			}
		}
		digit, err := strconv.Atoi(string(start) + string(end))
		if err != nil {
			panic(err)
		}
		acum += digit
	}
	fmt.Printf("Result: %v\n\n\n", acum)
}

func problem2() {
	lines := readFileLines()
	var acum int
	for _, line := range lines {
		var start, end = findFirstNumber(line), findLastNumber(line)
		//fmt.Printf("%v -> %v %v \n", line, start, end)
		digit, err := strconv.Atoi(start + end)
		if err != nil {
			fmt.Printf("%v -> %v %v \n", line, start, end)
			panic(err)
		}
		acum += digit

	}
	fmt.Printf("Result: %v\n\n\n", acum)
}

var wordToIntMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func containsNumber(line string) (string, bool) {
	for key, value := range wordToIntMap {
		if strings.Contains(line, key) {
			return value, true
		}
	}
	return "", false
}

func findFirstNumber(line string) string {
	for i := 0; i < len(line); i++ {
		if unicode.IsDigit(rune(line[i])) {
			return line[i : i+1]
		}
		if number, contains := containsNumber(line[:i+1]); contains {
			return number
		}
	}
	fmt.Printf("%v", line)

	return line
}

func findLastNumber(line string) string {
	for i := len(line) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(line[i])) {
			return line[i : i+1]
		}
		if number, contains := containsNumber(line[i:]); contains {
			return number
		}
	}
	fmt.Printf("%v", line)

	return line
}

func readFileLines() []string {
	f, err := os.OpenFile("./input.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	lines := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
