package day8

import (
	"advent-of-code-2023/io"
	"fmt"
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

	startLocations := getStartingLocations(locations)

	fmt.Println(locations)

	var steps []int

	for _, baseCurrent := range startLocations {
		stepCount := 0
		current := baseCurrent
	outer:
		for {
			for _, move := range moves {
				switch move {
				case 'L':
					current = locations[current].left
				case 'R':
					current = locations[current].right
				}

				stepCount++

				if strings.HasSuffix(current, "Z") {
					break outer
				}
			}
		}

		steps = append(steps, stepCount)
	}

	fmt.Println(LCM(steps[0], steps[1], steps...))
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

func getStartingLocations(locations map[string]Location) []string {

	var startLocations []string

	for location := range locations {
		if strings.HasSuffix(location, "A") {
			startLocations = append(startLocations, location)
		}
	}

	return startLocations
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
