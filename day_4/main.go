package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	problem2()
}

func problem1() {
	winningNumbersRegex := regexp.MustCompile(`:(?P<winners>.*) (\|) (?P<playeds>.*)`)
	result := 0
	for _, line := range readFileLines() {
		matchs := winningNumbersRegex.FindStringSubmatch(line)
		winningNumbers := stringToNumberList(matchs[1])
		playedNumbers := stringToNumberList(matchs[3])
		gameResult := 0
		for _, winnerNumber := range winningNumbers {
			if slices.Contains(playedNumbers, winnerNumber) {
				if gameResult == 0 {
					gameResult = 1
				} else {
					gameResult *= 2
				}
			}
		}
		result += gameResult
	}
	fmt.Println("Result: ", result)
}

func problem2() {
	winningNumbersRegex := regexp.MustCompile(`:(?P<winners>.*) (\|) (?P<playeds>.*)`)
	lines := readFileLines()
	cardMultiplier := make(map[int]int, len(lines))
	for cardIndex := range lines {
		cardMultiplier[cardIndex] = 1
	}

	for lineIndex, line := range lines {
		matchs := winningNumbersRegex.FindStringSubmatch(line)
		winningNumbers := stringToNumberList(matchs[1])
		playedNumbers := stringToNumberList(matchs[3])
		gameResult := 0
		for _, winnerNumber := range winningNumbers {
			if slices.Contains(playedNumbers, winnerNumber) {
				gameResult++
			}
		}
		for i := lineIndex + 1; i < len(lines) && i <= lineIndex+gameResult; i++ {
			cardMultiplier[i] += cardMultiplier[lineIndex]
		}
	}
	result := 0
	for _, value := range cardMultiplier {
		result += value
	}
	fmt.Println("Result: ", result)

}

func stringToNumberList(value string) []int {
	value = strings.Trim(strings.ReplaceAll(value, "  ", " "), " ")
	splittedString := strings.Split(value, " ")
	nums := make([]int, len(splittedString))
	for i, num := range splittedString {
		result, _ := strconv.Atoi(num)
		nums[i] = result
	}
	return nums
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
