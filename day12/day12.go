package day12

import (
	"advent-of-code-2023/io"
	"strconv"
	"strings"
)

func Run() {
	//1 4 1 1 4 10

	lines := io.ReadLines("resources/day12/input.txt")

	var readings []Reading

	for _, line := range lines {
		readings = append(readings, readingFromLine(line))
	}

	sum := 0

	for index, reading := range readings {
		i := solutionsFromData(reading.line, reading.sizes)
		println(index, i)
		sum += i
	}

	println(sum)
}

type Reading struct {
	line  string
	sizes []int
}

func readingFromLine(line string) Reading {
	split := strings.Split(line, " ")

	var numbers []int

	for _, number := range strings.Split(split[1], ",") {
		value, _ := strconv.Atoi(number)

		numbers = append(numbers, value)
	}

	var scaledLines []string
	var scaledNumbers []int

	for i := 0; i < 5; i++ {
		scaledLines = append(scaledLines, split[0])

		for _, val := range numbers {
			scaledNumbers = append(scaledNumbers, val)
		}
	}

	return Reading{line: strings.Join(scaledLines, "?"), sizes: scaledNumbers}
}

func solutionsFromData(line string, sizes []int) int {
	index := strings.Index(line, "?")

	if index <= -1 {

		realSizes := sizesFromLine(line)

		if len(realSizes) != len(sizes) {
			return 0
		}

		for i := 0; i < len(realSizes); i++ {
			if realSizes[i] != sizes[i] {
				return 0
			}
		}

		return 1
	}

	fixedSizes := sizesFromLine(line[0:index])

	if len(fixedSizes) > len(sizes) {
		return 0
	}

	for i := 0; i < len(fixedSizes)-1; i++ {
		if fixedSizes[i] != sizes[i] {
			return 0
		}
	}

	if len(fixedSizes) > 0 {
		if fixedSizes[len(fixedSizes)-1] > sizes[len(fixedSizes)-1] {
			return 0
		}
	}

	var solutions int

	newLine := strings.Replace(line, "?", "#", 1)
	solutions += solutionsFromData(newLine, sizes)

	newLine = strings.Replace(line, "?", ".", 1)
	solutions += solutionsFromData(newLine, sizes)

	return solutions
}

func sizesFromLine(line string) []int {
	var sizes []int

	lastSize := 0

	for _, char := range line {
		if char != '.' {
			lastSize++
			continue
		}

		if lastSize <= 0 {
			continue
		}

		sizes = append(sizes, lastSize)
		lastSize = 0
	}

	if lastSize > 0 {
		sizes = append(sizes, lastSize)
	}

	return sizes
}
