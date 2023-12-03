package day3

import (
	"advent-of-code-2023/io"
	"strconv"
	"unicode"
)

func Run() {
	lines := io.ReadLines("resources/day3/input.txt")

	var symbols []Symbol
	var serialNumbers []SerialNumber

	for y, line := range lines {

		number := ""

		tryAddSerial := func(x int) {
			if number != "" {
				val, _ := strconv.Atoi(number)

				serialNumbers = append(serialNumbers, SerialNumber{
					startX: x - len(number),
					startY: y,
					value:  val,
					length: len(number),
				})

				number = ""
			}
		}

		for x, char := range line {
			if unicode.IsDigit(char) {
				number += string(char)
				continue
			}

			tryAddSerial(x)

			if char == '.' {
				continue
			}

			symbols = append(symbols, Symbol{
				x: x,
				y: y,
			})
		}

		tryAddSerial(len(line) - 1)
	}

	sum := 0

	for _, number := range serialNumbers {
		if isSerialNumber(number, symbols) {
			sum += number.value
		}
	}

	println(sum)
}

type SerialNumber struct {
	startX int
	startY int
	value  int
	length int
}

type Symbol struct {
	x int
	y int
}

func isSerialNumber(number SerialNumber, symbols []Symbol) bool {
	startX := number.startX - 1
	endX := number.startX + number.length

	startY := number.startY - 1
	endY := number.startY + 1

	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			for _, symbol := range symbols {
				if symbol.x == x && symbol.y == y {
					return true
				}
			}
		}
	}

	return false
}
