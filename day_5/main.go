package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"sync"
)

func main() {
	problem2()
}

// Result: 806029445 == 806029445
func problem1() {
	text := readFileText()
	categories := getAllCategoriesSources(text)
	values := getSeeds(text)
	for i := range values {
		for _, category := range categories {
			values[i] = getSouceFromDestination(values[i], category)
		}
	}

	fmt.Printf("Result: %v", slices.Min(values))
}

// 59370573
// 59370573
func problem2() {
	text := readFileText()
	categories := getAllCategoriesSources(text)
	values := getSeeds(text)
	locationChan := make(chan int)

	wg := new(sync.WaitGroup)
	location := 0
	for i := 0; i < len(values); i += 2 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			getLocation(values[i], values[i+1], categories, locationChan)
		}(i)
	}
	go func() {
		wg.Wait()
		close(locationChan)
	}()
	for message := range locationChan {
		if location == 0 {
			location = message
		} else {
			location = min(location, message)
		}
	}
	fmt.Printf("Result: %v", location)
}

func getLocation(value, valueRange int, categories [][]categorySorce, c chan int) {
	location := 0
	for j := value; j < value+valueRange; j++ {
		result := j
		for _, category := range categories {
			result = getSouceFromDestination(result, category)
		}
		if location == 0 {
			location = result
		} else {
			location = min(location, result)
		}
	}
	c <- location
}

func getSouceFromDestination(destination int, categories []categorySorce) int {
	for _, c := range categories {
		if destination >= c.initialIndex && destination < c.initialIndex+c.rangeSize {
			return c.value + abs(destination, c.initialIndex)
		}
	}
	return destination
}

func getSeeds(source string) []int {
	c := regexp.MustCompile(`seeds: ([\d+ ]*)`)
	seeds := make([]int, 0)
	for _, value := range strings.Split(c.FindStringSubmatch(source)[1], " ") {
		seeds = append(seeds, stringToNumber(value))
	}
	return seeds
}

func getAllCategoriesSources(source string) [][]categorySorce {
	allCategories := make([][]categorySorce, 0)
	for i, seedMap := range regexp.MustCompile(` map:\n((\d+| |\n)*)\n`).FindAllStringSubmatch(source, -1) {
		result := strings.Trim(seedMap[1], "\n")
		result = strings.Trim(result, " ")
		allCategories = append(allCategories, make([]categorySorce, 0))
		for _, value := range strings.Split(result, "\n") {
			values := strings.Split(strings.Trim(value, " "), " ")
			allCategories[i] = append(allCategories[i], categorySorce{
				stringToNumber(values[0]),
				stringToNumber(values[1]),
				stringToNumber(values[2]),
			})
		}
	}

	return allCategories
}

func stringToNumber(value string) int {
	result, _ := strconv.Atoi(strings.Trim(value, "\n"))
	return result
}

type categorySorce struct {
	// value to distribute against the range
	value int
	// where starts the range
	initialIndex int
	// until the range continues
	rangeSize int
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func readFileText() string {
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

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(f)
	if err != nil {
		panic(err)
	}
	contents := buf.String()

	return contents
}
