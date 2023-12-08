package day4

import (
	"advent-of-code-2023/io"
	"strings"
)

func Run() {
	lines := io.ReadLines("resources/day4/input.txt.txt.txt")

	var games []Game

	for _, line := range lines {
		game := strings.Split(line, ":")[1]
		values := strings.Split(game, "|")

		games = append(games, Game{
			winning: io.LineToNumbers(values[0]),
			having:  io.LineToNumbers(values[1]),
			count:   1,
		})
	}

	for index, game := range games {
		if game.count <= 0 {
			break
		}

		for i := 0; i < getGameWinCount(game); i++ {
			games[index+1+i].count += game.count
		}
	}

	sum := 0

	for _, game := range games {
		sum += game.count
	}

	println(sum)
}

type Game struct {
	winning []int
	having  []int
	count   int
}

func getGameWinCount(game Game) int {

	count := 0

	for _, have := range game.having {
		for _, win := range game.winning {
			if have == win {
				count++
			}
		}
	}

	return count
}
