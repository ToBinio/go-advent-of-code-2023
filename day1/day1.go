package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func Run() {
	println("day1")

	file, err := os.Open("resources/day1/example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var values []string

		for i := range line {
			char := string(line[i])

			val, err := strconv.Atoi(char)

			if err != nil {
				continue
			}

			values = append(values, strconv.Itoa(val))
		}

		val, _ := strconv.Atoi(values[0] + values[len(values)-1])

		sum += val
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	print(sum)
}
