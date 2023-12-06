package day5

import (
	"advent-of-code-2023/io"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Run() {

	lines := io.ReadLines("resources/day5/input.txt.txt")

	var seeds []Range
	var convertedSeeds []Range

	for i, line := range lines {
		if line == "" {
			continue
		}

		if i == 0 {
			s := strings.Split(line, ":")[1]
			numbers := io.LineToNumbers(s)

			for i := 0; i < len(numbers); i += 2 {
				seeds = append(seeds, Range{
					min: numbers[i],
					max: numbers[i] + numbers[i+1] - 1,
				})
			}
			continue
		}

		if !unicode.IsDigit(rune(line[0])) {
			for _, seed := range seeds {
				convertedSeeds = append(convertedSeeds, seed)
			}

			var temp []Range

			for _, seed := range convertedSeeds {
				if seed.min == -1 {
					continue
				}

				temp = append(temp, seed)
			}

			seeds = clone(temp)
			fmt.Println(seeds)
			convertedSeeds = []Range{}

			continue
		}

		split := strings.Split(line, " ")

		destination, _ := strconv.Atoi(split[0])
		source, _ := strconv.Atoi(split[1])
		length, _ := strconv.Atoi(split[2])

		for i, seed := range seeds {
			if source > seed.max || (source+length-1) < seed.min {
				continue
			}

			rangeMin := max(seed.min, source)
			rangeMax := min(seed.max, source+length-1)

			realLength := rangeMax - rangeMin
			realOffset := rangeMin - source

			convertedSeeds = append(convertedSeeds, Range{
				min: destination + realOffset,
				max: destination + realOffset + realLength - 1,
			})

			if rangeMax != seed.max {
				seeds = append(seeds, Range{min: rangeMax + 1, max: seed.max})
			}

			if rangeMin != seed.min {
				seeds = append(seeds, Range{min: seed.min, max: rangeMin - 1})
			}

			seeds[i].max = -1
			seeds[i].min = -1
		}
	}

	for _, seed := range seeds {
		convertedSeeds = append(convertedSeeds, seed)
	}

	var temp []Range

	for _, seed := range convertedSeeds {
		if seed.min == -1 {
			continue
		}

		temp = append(temp, seed)
	}

	seeds = clone(temp)
	fmt.Println("current", getSmallest(seeds))
}

func getSmallest(array []Range) int {

	smallest := array[0].min

	for _, val := range array {
		smallest = min(smallest, val.min)
	}

	return smallest
}

func clone(array []Range) []Range {
	var val []Range

	for _, i := range array {
		val = append(val, i)
	}

	return val
}

type Range struct {
	min int
	max int
}
