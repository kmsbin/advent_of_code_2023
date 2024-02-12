package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	problem2()
}

func problem1() {
	values := readFileLines()
	result := 1
	for i := 0; i < len(values[0]); i++ {
		time := values[0][i]
		record := values[1][i]
		records := 0
		for j := 1; j < time; j++ {
			if (time-j)*j > record {
				records++
			}
		}
		result *= records
	}
	fmt.Printf("Result: %v\n", result)
}

func readFileLines() [][]int {
	f, err := os.OpenFile("./input.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)
	lines := make([][]int, 0)
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	lines = append(lines, stringToInts(scanner.Text()))
	scanner.Scan()
	lines = append(lines, stringToInts(scanner.Text()))

	return lines
}

func stringToInts(value string) []int {
	c := regexp.MustCompile(`\d+`)
	matchs := c.FindAllString(value, -1)
	numValues := make([]int, len(matchs))
	for i, numValue := range matchs {
		numValues[i], _ = strconv.Atoi(numValue)
	}
	return numValues
}

func problem2() {
	time, record := readFileLinesPart2()
	records := 0
	for j := 1; j < time; j++ {
		if (time-j)*j > record {
			records++
		}
	}

	fmt.Printf("Result: %v\n", records)
}

func readFileLinesPart2() (int, int) {
	f, err := os.OpenFile("./input.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	milliseconds := stringToIntsPart2(scanner.Text(), "Time:")
	scanner.Scan()
	millimeters := stringToIntsPart2(scanner.Text(), "Distance:")

	return milliseconds, millimeters
}

func stringToIntsPart2(value, name string) int {
	c := regexp.MustCompile(`\d+`)
	matchs := c.FindAllString(value, -1)
	strings.Join(matchs, "")
	numValue, _ := strconv.Atoi(strings.Join(matchs, ""))
	return numValue
}
