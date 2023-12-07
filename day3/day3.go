package day3

import (
	"advent-of-code-2023/io"
	"strconv"
	"unicode"
)

func Run() {
	lines := io.ReadLines("resources/day3/input.txt")

	var gears []GearSymbol
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

			if char == '*' {
				gears = append(gears, GearSymbol{
					x: x,
					y: y,
				})
			}
		}

		tryAddSerial(len(line) - 1)
	}

	sum := 0

	for _, gear := range gears {
		sum += getRatio(gear, serialNumbers)
	}

	println(sum)
}

type SerialNumber struct {
	startX int
	startY int
	value  int
	length int
}

type GearSymbol struct {
	x int
	y int
}

func getRatio(gear GearSymbol, numbers []SerialNumber) int {

	var adjacent []SerialNumber

number:
	for _, number := range numbers {
		for x := number.startX; x <= number.startX+number.length-1; x++ {
			for gearX := gear.x - 1; gearX <= gear.x+1; gearX++ {
				for gearY := gear.y - 1; gearY <= gear.y+1; gearY++ {
					if x == gearX && number.startY == gearY {
						adjacent = append(adjacent, number)
						continue number
					}
				}
			}
		}
	}

	if len(adjacent) != 2 {
		return 0
	}

	return adjacent[0].value * adjacent[1].value
}
