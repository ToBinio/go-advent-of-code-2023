package day6

import (
	"advent-of-code-2023/io"
	"strings"
)

func Run() {
	lines := io.ReadLines("resources/day6/input.txt.txt")

	times := io.LineToNumbers(strings.Split(lines[0], ":")[1])
	distances := io.LineToNumbers(strings.Split(lines[1], ":")[1])

	var games []Game

	for i := 0; i < len(times); i++ {
		games = append(games, Game{
			time:           times[i],
			recordDistance: distances[i],
		})
	}

	sum := 1

	for _, game := range games {
		sum *= len(getWinTimes(game))
	}

	println(sum)
}

type Game struct {
	time           int
	recordDistance int
}

func getWinTimes(game Game) []int {
	var times []int

	for time := 1; time < game.time; time++ {
		if getDistance(game, time) > game.recordDistance {
			times = append(times, time)
		}
	}

	return times
}

func getDistance(game Game, holdTime int) int {
	return holdTime * (game.time - holdTime)
}
