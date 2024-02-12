package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	problem2()
}

// answer: 250058342
func problem1() {
	cardsPlays := readFileLines()
	cardsPlayeds := make([]CalculedCard, len(cardsPlays))
	for i := 0; i < len(cardsPlays); i++ {
		cardsPlayeds[i] = getCardWeight(cardsPlays[i])
	}
	slices.SortFunc(cardsPlayeds, func(firstHand, secondHand CalculedCard) int {
		if compare := cmp.Compare(secondHand.cardType, firstHand.cardType); compare != 0 {
			return compare
		}
		for i := range firstHand.cards {
			firstIndex, secondIndex := slices.Index(cards, firstHand.cards[i]), slices.Index(cards, secondHand.cards[i])
			if compare := cmp.Compare(secondIndex, firstIndex); compare != 0 {
				return compare
			}
		}
		panic("PARES IGUAIS")
	})
	result := 0
	for i, cardPlayed := range cardsPlayeds {
		fmt.Printf("%v: %v\n", strings.Join(cardPlayed.cards, ""), cardPlayed.bid)
		result += cardPlayed.bid * (i + 1)
	}
	fmt.Printf("Result: %v\n", result)
}

func getCardWeight(cardPlay CardPlay) CalculedCard {
	cardsWeight := map[string]int{}
	for _, card := range cardPlay.cards {
		cardsWeight[card]++
	}
	keys := Keys(cardsWeight)
	return CalculedCard{
		bid:      cardPlay.bind,
		cards:    cardPlay.cards,
		cardType: getCardType(cardsWeight, keys, 0),
	}
}

func getCardType(cardsWeight map[string]int, keys []string, sum int) CardType {
	switch cardsWeight[keys[0]] + sum {
	case 5:
		return five
	case 4:
		return four
	case 3:
		if cardsWeight[keys[1]] == 2 {
			return fullHouse
		}
		return three
	case 2:
		if cardsWeight[keys[1]] == 2 {
			return twoPair
		}
		return pair
	default:
		return cardHigh
	}
}

var cards = []string{"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"}

func Keys(m map[string]int) []string {
	r := make([]string, 0, 0)
	for k := range m {
		r = append(r, k)
	}
	slices.SortFunc(r, func(a, b string) int {
		if m[a] > m[b] {
			return -1
		} else if m[a] < m[b] {
			return +1
		}
		return cmp.Compare(slices.Index(cards, a), slices.Index(cards, b))
	})
	return r
}

func readFileLines() []CardPlay {
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
	lines := make([]CardPlay, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		lineParts := strings.Split(text, " ")
		bind, _ := strconv.Atoi(lineParts[1])
		play := CardPlay{
			cards: strings.Split(lineParts[0], ""),
			bind:  bind,
		}
		lines = append(lines, play)
	}
	return lines
}

type CardType uint32

const (
	five CardType = iota
	four
	fullHouse
	three
	twoPair
	pair
	cardHigh
)

type CardPlay struct {
	cards []string
	bind  int
}

type CalculedCard struct {
	cardType    CardType
	cards       []string
	bid         int
	jokersCount int
}
