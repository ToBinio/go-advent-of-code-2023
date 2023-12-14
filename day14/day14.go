package day14

import (
	"advent-of-code-2023/io"
)

func Run() {
	lines := io.ReadLines("resources/day14/input.txt")

	var grid [][]rune

	for _, line := range lines {
		var gridLine []rune

		for _, char := range line {
			gridLine = append(gridLine, char)
		}

		grid = append(grid, gridLine)
	}

	grid = moveRocksUp(grid)
	println(getWeight(grid))
}

func moveRocksUp(grid [][]rune) [][]rune {
	for i := 0; i < len(grid); i++ {
		for y := range grid {
			for x := range grid[y] {
				if grid[y][x] != 'O' {
					continue
				}

				if y-1 < 0 || grid[y-1][x] != '.' {
					continue
				}

				grid[y][x] = '.'
				grid[y-1][x] = 'O'
			}
		}
	}

	return grid
}

func getWeight(grid [][]rune) int {
	sum := 0

	for y, stones := range grid {
		for _, stone := range stones {
			if stone == 'O' {
				sum += len(grid) - y
			}
		}
	}

	return sum
}
