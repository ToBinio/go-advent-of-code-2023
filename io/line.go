package io

import (
	"strconv"
	"strings"
)

func LineToNumbers(line string) []int {
	split := strings.Split(strings.Trim(line, " "), " ")

	var numbers []int

	for _, val := range split {
		if val == "" {
			continue
		}

		number, _ := strconv.Atoi(val)

		numbers = append(numbers, number)
	}

	return numbers
}
