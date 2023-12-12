package day11

import (
	"advent-of-code-2023/io"
)

func Run() {
	lines := io.ReadLines("resources/day11/input.txt")

	var grid [][]bool
	var stars []PreStar

	for y, line := range lines {
		var chars []bool

		for x, char := range line {
			chars = append(chars, char == '#')
			if char == '#' {
				stars = append(stars, PreStar{x: x, y: y, xShifts: 0, yShifts: 0})
			}
		}

		grid = append(grid, chars)
	}

	rows, columns := expandGrid(grid)

	for _, column := range columns {
		for i, star := range stars {
			if star.x > column {
				stars[i].xShifts++
			}
		}
	}

	for _, row := range rows {
		for i, star := range stars {
			if star.y > row {
				stars[i].yShifts++
			}
		}
	}

	for i, star := range stars {
		stars[i].x = star.x + star.xShifts*(1000000-1)
		stars[i].y = star.y + star.yShifts*(1000000-1)
	}

	sum := 0

	for i := 0; i < len(stars); i++ {
		for j := i; j < len(stars); j++ {

			x := stars[i].x - stars[j].x
			y := stars[i].y - stars[j].y

			if x < 0 {
				x = x * -1
			}

			if y < 0 {
				y = y * -1
			}

			sum += x + y
		}
	}

	println(sum)
}

func expandGrid(grid [][]bool) ([]int, []int) {
	var emptyRows []int
	var emptyColumns []int

	for y, row := range grid {

		isEmpty := true

		for _, isStart := range row {
			if isStart {
				isEmpty = false
				break
			}
		}

		if isEmpty {
			emptyRows = append(emptyRows, y)
		}
	}

	var isEmptyColumns []bool

	for i := 0; i < len(grid[0]); i++ {
		isEmptyColumns = append(isEmptyColumns, true)
	}

	for _, row := range grid {
		for x, isStart := range row {
			if isStart {
				isEmptyColumns[x] = false
			}
		}
	}

	for i, column := range isEmptyColumns {
		if column {
			emptyColumns = append(emptyColumns, i)
		}
	}

	return emptyRows, emptyColumns
}

type PreStar struct {
	x       int
	y       int
	xShifts int
	yShifts int
}
