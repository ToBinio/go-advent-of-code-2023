package day7

import (
	"advent-of-code-2023/io"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Run() {
	lines := io.ReadLines("resources/day7/input.txt.txt")

	var hands []Hand

	for _, line := range lines {
		hands = append(hands, HandFromLine(line))
	}

	slices.SortFunc(hands, CompareHand)

	sum := 0

	for i, hand := range hands {
		sum += hand.bit * (i + 1)
	}

	fmt.Println(sum)
}

func CompareHand(hand1 Hand, hand2 Hand) int {
	compare := hand2.level - hand1.level
	if compare != 0 {
		return int(compare)
	}

	for i := 0; i < len(hand1.cards); i++ {
		compare := determineCardLevel(rune(hand2.cards[i])) - determineCardLevel(rune(hand1.cards[i]))
		if compare != 0 {
			return compare
		}
	}

	return 0
}

func HandFromLine(line string) Hand {
	split := strings.Split(line, " ")

	bit, _ := strconv.Atoi(split[1])

	return Hand{
		cards: split[0],
		level: determineHandLevel(maximiseHand(split[0])),
		bit:   bit,
	}
}

func maximiseHand(cards string) string {

	if !strings.Contains(cards, "J") {
		return cards
	}

	bestCards := strings.Replace(cards, "J", string(CardType[0]), 1)

	for i, value := range CardType {

		if i == len(CardType)-1 {
			break
		}

		newCards := maximiseHand(strings.Replace(cards, "J", string(value), 1))

		if determineHandLevel(bestCards) > determineHandLevel(newCards) {
			bestCards = newCards
		}
	}

	return bestCards
}

func determineHandLevel(cards string) HandLevel {

	countMap := make(map[rune]int)
	for _, char := range cards {
		countMap[char]++
	}

	maxCount := 0
	typeCount := len(countMap)

	for _, value := range countMap {
		maxCount = max(maxCount, value)
	}

	if maxCount == 5 {
		return Five
	}

	if maxCount == 4 {
		return Four
	}

	if maxCount == 3 && typeCount == 2 {
		return FullHouse
	}

	if maxCount == 3 {
		return Three
	}

	if maxCount == 2 && typeCount == 3 {
		return TwoPair
	}

	if maxCount == 2 {
		return Pair
	}

	return High
}

var CardType = []rune{'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J'}

func determineCardLevel(card rune) int {
	for i, r := range CardType {
		if r == card {
			return i
		}
	}

	return -1
}

type Hand struct {
	cards string
	level HandLevel
	bit   int
}

type HandLevel int

const (
	Five HandLevel = iota
	Four
	FullHouse
	Three
	TwoPair
	Pair
	High
)
