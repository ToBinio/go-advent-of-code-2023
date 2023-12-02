package day2

import (
	"strconv"
	"strings"
)

type Move struct {
	Red   int `json:"red,omitempty"`
	Blue  int `json:"blue,omitempty"`
	Green int `json:"green,omitempty"`
}

func MoveFromString(text string) Move {
	text = strings.Trim(text, " ")

	move := Move{
		Red:   0,
		Blue:  0,
		Green: 0,
	}

	for _, data := range strings.Split(text, ",") {
		data = strings.Trim(data, " ")

		split := strings.Split(data, " ")

		val, _ := strconv.Atoi(split[0])

		switch split[1] {
		case "red":
			move.Red = val
		case "blue":
			move.Blue = val
		case "green":
			move.Green = val
		}
	}

	return move
}
