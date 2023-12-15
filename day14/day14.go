package day14

import (
	"advent-of-code-2023/io"
)

func Run() {
	lines := io.ReadLines("resources/day14/example.txt")

	var grid [][]rune

	for _, line := range lines {
		var gridLine []rune

		for _, char := range line {
			gridLine = append(gridLine, char)
		}

		grid = append(grid, gridLine)
	}

	for i := 0; i < 1000_000_000; i++ {
		grid = moveRocksUp(grid)
		grid = moveRocksLeft(grid)
		grid = moveRocksDown(grid)
		grid = moveRocksRight(grid)

		if i%1000_000 == 0 {
			println(i)
		}
	}

	println(getWeight(grid))
}

func moveRocksUp(grid [][]rune) [][]rune {
	hasChanges := true
	for hasChanges {
		hasChanges = false
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

				hasChanges = true
			}
		}
	}

	return grid
}

func moveRocksDown(grid [][]rune) [][]rune {
	hasChanges := true
	for hasChanges {
		hasChanges = false
		for y := len(grid) - 1; y >= 0; y-- {
			for x := range grid[y] {
				if grid[y][x] != 'O' {
					continue
				}

				if y+1 >= len(grid) || grid[y+1][x] != '.' {
					continue
				}

				grid[y][x] = '.'
				grid[y+1][x] = 'O'

				hasChanges = true
			}
		}
	}

	return grid
}

func moveRocksLeft(grid [][]rune) [][]rune {
	hasChanges := true
	for hasChanges {
		hasChanges = false
		for x := range grid[0] {
			for y := range grid {
				if grid[y][x] != 'O' {
					continue
				}

				if x-1 < 0 || grid[y][x-1] != '.' {
					continue
				}

				grid[y][x] = '.'
				grid[y][x-1] = 'O'

				hasChanges = true
			}
		}
	}

	return grid
}

func moveRocksRight(grid [][]rune) [][]rune {
	hasChanges := true
	for hasChanges {
		hasChanges = false
		for x := len(grid[0]) - 1; x >= 0; x-- {
			for y := range grid {
				if grid[y][x] != 'O' {
					continue
				}

				if x+1 >= len(grid[0]) || grid[y][x+1] != '.' {
					continue
				}

				grid[y][x] = '.'
				grid[y][x+1] = 'O'

				hasChanges = true
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
