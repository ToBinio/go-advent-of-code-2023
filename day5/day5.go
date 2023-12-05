package day5

import (
	"advent-of-code-2023/io"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func Run() {

	lines := io.ReadLines("resources/day5/example.txt")

	var seeds []int
	var convertedSeeds []int

	for i, line := range lines {
		if line == "" {
			continue
		}

		if i == 0 {
			s := strings.Split(line, ":")[1]

			seeds = io.LineToNumbers(s)
			convertedSeeds = clone(seeds)
			continue
		}

		if !unicode.IsDigit(rune(line[0])) {
			fmt.Println(convertedSeeds)
			seeds = clone(convertedSeeds)
			continue
		}

		split := strings.Split(line, " ")

		destination, _ := strconv.Atoi(split[0])
		source, _ := strconv.Atoi(split[1])
		length, _ := strconv.Atoi(split[2])

		for seedIndex, seed := range seeds {

			if seed < source || seed > source+length {
				continue
			}

			for offset := 0; offset < length; offset++ {
				if seed == source+offset {
					convertedSeeds[seedIndex] = destination + offset
				}
			}
		}
	}

	println(slices.Min(convertedSeeds))
}

func clone(array []int) []int {
	var val []int

	for _, i := range array {
		val = append(val, i)
	}

	return val
}
