package day15

import (
	"advent-of-code-2023/io"
	"strings"
)

func Run() {
	line := io.ReadLines("resources/day15/input.txt")[0]

	elements := strings.Split(line, ",")

	sum := 0

	for _, element := range elements {
		sum += hash(element)
	}

	println(sum)
}

func hash(text string) int {
	value := 0

	for _, char := range text {
		value += int(char)
		value *= 17
		value %= 256
	}

	return value
}
