package day8

import (
	"advent-of-code-2023/io"
	"strings"
)

func Run() {
	lines := io.ReadLines("resources/day8/input.txt")
	var moves []rune

	locations := make(map[string]Location)

	for index, line := range lines {
		if index == 0 {
			moves = parseMoves(line, moves)
		}

		if index > 1 {
			parseLocation(line, locations)
		}
	}

	current := "AAA"
	steps := 0

outer:
	for {
		for _, move := range moves {
			switch move {
			case 'L':
				current = locations[current].left
			case 'R':
				current = locations[current].right
			}

			steps++

			if current == "ZZZ" {
				break outer
			}
		}
	}

	println(steps)
}

type Location struct {
	left  string
	right string
}

func parseLocation(line string, locations map[string]Location) {
	split := strings.Split(line, "=")

	current := strings.Trim(split[0], " ")

	left := split[1][2:5]
	right := split[1][7:10]

	locations[current] = Location{
		left:  left,
		right: right,
	}
}

func parseMoves(line string, moves []rune) []rune {
	for _, char := range line {
		moves = append(moves, char)
	}

	return moves
}
