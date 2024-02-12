package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const (
	blueMax  = 14
	greenMax = 13
	redMax   = 12
)

func main() {
	problem1()
	problem2()
}

func problem2() {
	gameAcum := 0
	for _, line := range readFileLines() {
		gameAcum += getSumOfMinimumGameSet(line)
	}
	fmt.Printf("Result 2: %v", gameAcum)
}

func problem1() {
	gameAcum := 0
	for _, line := range readFileLines() {
		if validateGameSet(line) {
			gameAcum += getGameId(line)
		}
	}
	fmt.Printf("Result 1: %v\n", gameAcum)
}

func getGameId(game string) int {
	e := regexp.MustCompile("Game (\\d+)")
	result := e.FindStringSubmatch(game)
	id, _ := strconv.Atoi(result[1])
	return id
}

func validateGameSet(game string) bool {
	e := regexp.MustCompile("(\\d+) blue|(\\d+) red|(\\d+) green")
	blues := e.FindAllStringSubmatch(game, -1)
	for _, result := range blues {
		if stringToInt(result[1]) > blueMax || stringToInt(result[2]) > redMax || stringToInt(result[3]) > greenMax {
			return false
		}
	}
	return true
}

func stringToInt(value string) int {
	result, _ := strconv.Atoi(value)
	return result
}

func getSumOfMinimumGameSet(game string) int {
	e := regexp.MustCompile("(\\d+) blue|(\\d+) red|(\\d+) green")
	blues := e.FindAllStringSubmatch(game, -1)
	blue := 0
	red := 0
	green := 0
	for _, result := range blues {
		blue = max(stringToInt(result[1]), blue)
		red = max(stringToInt(result[2]), red)
		green = max(stringToInt(result[3]), green)
	}
	return blue * red * green
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
