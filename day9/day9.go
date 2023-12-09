package day9

import (
	"advent-of-code-2023/io"
)

func Run() {
	lines := io.ReadLines("resources/day9/input.txt")

	sum := 0

	for _, line := range lines {
		prediction := getPrediction(line)
		sum += prediction
	}

	println(sum)
}

func getPrediction(line string) int {
	numbers := [][]int{io.LineToNumbers(line)}

	numbers = generatePredictionValues(numbers)

	prediction := 0

	for i := len(numbers) - 2; i >= 0; i-- {
		values := numbers[i]
		prediction = values[0] - prediction
	}

	return prediction
}

func generatePredictionValues(numbers [][]int) [][]int {
outer:
	for {
		var nextNumbers []int

		lastIndex := len(numbers) - 1

		for i := 0; i < len(numbers[lastIndex])-1; i++ {
			nextNumbers = append(nextNumbers, numbers[lastIndex][i+1]-numbers[lastIndex][i])
		}

		numbers = append(numbers, nextNumbers)

		for _, val := range nextNumbers {
			if val != 0 {
				continue outer
			}
		}

		break
	}

	return numbers
}
