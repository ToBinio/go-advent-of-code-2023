package day2

import (
	"strconv"
	"strings"
)

type Game struct {
	Id    int    `json:"id"`
	Moves []Move `json:"moves"`
}

func GameFromLine(line string) Game {
	split := strings.Split(line, ":")

	id, _ := strconv.Atoi(strings.Split(split[0], " ")[1])
	var moves []Move

	for _, move := range strings.Split(split[1], ";") {
		moves = append(moves, MoveFromString(move))
	}

	return Game{
		Id:    id,
		Moves: moves,
	}
}
