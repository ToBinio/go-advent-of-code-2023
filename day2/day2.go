package day2

import (
	"advent-of-code-2023/io"
)

func Run() {
	lines := io.ReadLines("resources/day2/input.txt")

	var games []Game

	for _, line := range lines {
		games = append(games, GameFromLine(line))
	}

	minRed := 12
	minBlue := 14
	minGreen := 13

	sum := 0

outer:
	for _, game := range games {
		for _, move := range game.Moves {
			if move.Red > minRed || move.Blue > minBlue || move.Green > minGreen {
				continue outer
			}
		}

		sum += game.Id
	}

	println(sum)
}
