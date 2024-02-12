package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
)

// 73201705
// 69786491
func main() {
	problem2()
}

// / 325 734
// / 88 859
// / 921 856
// / 806 137
// / 620 856
func problem2() {
	lines := readFileLines()
	ySize := len(lines)
	var numbersIndexes = make([][][]int, ySize)
	numRegex := regexp.MustCompile(`\d+`)
	for lineIndex, line := range lines {
		for _, value := range numRegex.FindAllStringIndex(line, -1) {
			numbersIndexes[lineIndex] = append(numbersIndexes[lineIndex], value)
		}
	}
	engineRegex := regexp.MustCompile(`\*`)
	resultAcum := 0
	for lineIndex, line := range lines {
		for _, indexes := range engineRegex.FindAllStringIndex(line, -1) {
			xMin, xMax := indexes[0]-1, indexes[0]+2
			numbersAdjecents := make([]string, 0, 2)

			for i := max(lineIndex-1, 0); i < min(ySize, lineIndex+2); i++ {
				for _, numbIndexes := range getIfArrBetween(numbersIndexes[i], xMin, xMax) {
					numbersAdjecents = append(numbersAdjecents, lines[i][numbIndexes[0]:numbIndexes[1]])
				}
			}
			if len(numbersAdjecents) == 2 {
				fmt.Println(stringToInt(numbersAdjecents[0]), stringToInt(numbersAdjecents[1]))
				resultAcum += stringToInt(numbersAdjecents[0]) * stringToInt(numbersAdjecents[1])
			}
		}
	}

	fmt.Printf("Result: %v", resultAcum)
}

func problem1() {
	lines := readFileLines()
	ySize := len(lines)
	var specialCharsIndexes = make([][]int, ySize)
	numRegex := regexp.MustCompile(`\d+`)
	specialCharRegex := regexp.MustCompile(`[^\d\n\s.]`)
	for lineIndex, line := range lines {
		for _, indexes := range specialCharRegex.FindAllStringIndex(line, -1) {
			specialCharsIndexes[lineIndex] = append(specialCharsIndexes[lineIndex], indexes[0])
		}
	}
	resultAcum := 0
	for lineIndex, line := range lines {
		for _, value := range numRegex.FindAllStringIndex(line, -1) {
			xMin, xMax := value[0]-1, value[len(value)-1:][0]
			if (lineIndex > 0 && checkIfBetween(specialCharsIndexes[lineIndex-1], xMin, xMax)) ||
				checkIfBetween(specialCharsIndexes[lineIndex], xMin, xMax) ||
				(lineIndex+1 < ySize && checkIfBetween(specialCharsIndexes[lineIndex+1], xMin, xMax)) {
				result, _ := strconv.Atoi(line[xMin+1 : xMax])
				resultAcum += result
			} else {
				fmt.Println(line[xMin+1 : xMax])

			}
		}
	}
	fmt.Printf("Result: %v", resultAcum)
}

func stringToInt(value string) int {
	result, _ := strconv.Atoi(value)
	return result
}

func checkIfBetween(arr []int, min int, max int) bool {
	return slices.ContainsFunc(arr, func(next int) bool {
		return next >= min && next <= max
	})
}

func getIfArrBetween(matrix [][]int, min int, max int) [][]int {
	matches := make([][]int, 0, 2)
	for _, arr := range matrix {
		if arr[0] >= min && arr[0] < max || arr[1] > min && arr[0] < max {
			matches = append(matches, arr)
		}
		if len(matches) == 2 {
			return matches
		}
	}
	return matches

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
