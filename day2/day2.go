package day2

import (
	"advent-of-code-2023/io"
	"fmt"
)

func Run() {
	lines := io.ReadLines("resources/day2/input.txt.txt")

	var games []Game

	for _, line := range lines {
		games = append(games, GameFromLine(line))
	}

	sum := 0

	for _, game := range games {
		minRed := 0
		minBlue := 0
		minGreen := 0

		for _, move := range game.Moves {
			minRed = max(move.Red, minRed)
			minBlue = max(move.Blue, minBlue)
			minGreen = max(move.Green, minGreen)
		}

		sum += minRed * minBlue * minGreen
	}

	fmt.Printf("%+v\n", games)
	println(sum)
}
