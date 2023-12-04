package day4

import (
	"advent-of-code-2023/io"
	"strconv"
	"strings"
)

func Run() {
	lines := io.ReadLines("resources/day4/input.txt")

	var games []Game

	for _, line := range lines {
		game := strings.Split(line, ":")[1]
		values := strings.Split(game, "|")

		games = append(games, Game{
			winning: lineToNumbers(values[0]),
			having:  lineToNumbers(values[1]),
		})
	}

	sum := 0

	for _, game := range games {
		sum += getGameValue(game)
	}

	println(sum)
}

type Game struct {
	winning []int
	having  []int
}

func getGameValue(game Game) int {

	value := 0

	for _, have := range game.having {
		for _, win := range game.winning {
			if have == win {
				if value == 0 {
					value = 1
				} else {
					value *= 2
				}
			}
		}
	}

	return value
}

func lineToNumbers(line string) []int {
	split := strings.Split(strings.Trim(line, " "), " ")

	var numbers []int

	for _, val := range split {
		if val == "" {
			continue
		}

		number, _ := strconv.Atoi(val)

		numbers = append(numbers, number)
	}

	return numbers
}
