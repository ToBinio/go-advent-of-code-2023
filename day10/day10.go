package day10

import (
	"advent-of-code-2023/io"
)

func Run() {
	lines := io.ReadLines("resources/day10/input.txt")

	var grid [][]Tile
	startX := 0
	startY := 0

	for y, line := range lines {

		var tiles []Tile

		for x, char := range line {
			tiles = append(tiles, getTileFromAscii(char, x, y))

			if char == 'S' {
				startX = x
				startY = y
			}
		}

		grid = append(grid, tiles)
	}

	pipes := stepThoughPipes(grid, []Tile{grid[startY][startX]}, startX, startY)

	scaledGrid := generateGrid(len(grid[0]), len(grid), pipes)
	touches, newGrid := mapScaledGrid(scaledGrid)

	println(countInner(touches, newGrid))

	//for _, line := range newGrid {
	//	for _, val := range line {
	//		fmt.Printf(" %d ", val)
	//	}
	//
	//	println()
	//}
}

func countInner(touches map[int]bool, grid [][]int) int {

	count := 0

	for y := 1; y < len(grid); y += 3 {
		for x := 1; x < len(grid[0]); x += 3 {
			if !touches[grid[y][x]] {
				count++
			}
		}
	}

	return count
}

func mapScaledGrid(grid [][]int) (map[int]bool, [][]int) {
	currentID := 2

	touched := make(map[int]bool)

	touched[1] = true

	for y, line := range grid {
		for x, _ := range line {
			if grid[y][x] == 0 {
				newGrid, touchedOuter := fillScaledGrid(grid, currentID, x, y)
				grid = newGrid

				touched[currentID] = touchedOuter
				currentID++
			}
		}
	}

	return touched, grid
}

func fillScaledGrid(grid [][]int, ID int, x int, y int) ([][]int, bool) {
	grid[y][x] = ID

	touchesOuter := false

	if x-1 < 0 {
		touchesOuter = true
	} else {
		if grid[y][x-1] == 0 {
			scaledGrid, touches := fillScaledGrid(grid, ID, x-1, y)

			if touches {
				touchesOuter = true
			}

			grid = scaledGrid
		}
	}

	if x+1 >= len(grid[0]) {
		touchesOuter = true
	} else {
		if grid[y][x+1] == 0 {
			scaledGrid, touches := fillScaledGrid(grid, ID, x+1, y)

			if touches {
				touchesOuter = true
			}

			grid = scaledGrid
		}
	}

	if y-1 < 0 {
		touchesOuter = true
	} else {
		if grid[y-1][x] == 0 {
			scaledGrid, touches := fillScaledGrid(grid, ID, x, y-1)

			if touches {
				touchesOuter = true
			}

			grid = scaledGrid
		}
	}

	if y+1 >= len(grid) {
		touchesOuter = true
	} else {
		if grid[y+1][x] == 0 {
			scaledGrid, touches := fillScaledGrid(grid, ID, x, y+1)

			if touches {
				touchesOuter = true
			}

			grid = scaledGrid
		}
	}

	return grid, touchesOuter
}

func generateGrid(width int, height int, pipes []Tile) [][]int {
	var grid [][]int

	for y := 0; y < height*3; y++ {
		var line []int

		for x := 0; x < width*3; x++ {
			line = append(line, 0)
		}

		grid = append(grid, line)
	}

	for _, pipe := range pipes {
		newX := pipe.x*3 + 1
		newY := pipe.y*3 + 1

		grid[newY][newX] = 1

		if pipe.up {
			grid[newY-1][newX] = 1
		}

		if pipe.down {
			grid[newY+1][newX] = 1
		}

		if pipe.left {
			grid[newY][newX-1] = 1
		}

		if pipe.right {
			grid[newY][newX+1] = 1
		}
	}

	return grid
}

func stepThoughPipes(grid [][]Tile, visitedTiles []Tile, startX int, startY int) []Tile {
	currentTile := visitedTiles[len(visitedTiles)-1]
	var previousTile Tile

	if len(visitedTiles) <= 1 {
		previousTile = getTileFromAscii('S', startX, startY)
	} else {
		previousTile = visitedTiles[len(visitedTiles)-2]
	}

	if currentTile.x == startX && currentTile.y == startY && len(visitedTiles) > 1 {
		return visitedTiles
	}

	x := currentTile.x - 1
	if previousTile.x != x && x >= 0 {
		if grid[currentTile.y][x].right && currentTile.left {
			visitedTiles = append(visitedTiles, grid[currentTile.y][x])
			visitedTiles = stepThoughPipes(grid, visitedTiles, startX, startY)
			return visitedTiles
		}
	}

	x = currentTile.x + 1
	if previousTile.x != x && x < len(grid[0]) {
		if grid[currentTile.y][x].left && currentTile.right {
			visitedTiles = append(visitedTiles, grid[currentTile.y][x])
			visitedTiles = stepThoughPipes(grid, visitedTiles, startX, startY)
			return visitedTiles
		}
	}

	y := currentTile.y - 1
	if previousTile.y != y && y >= 0 {
		if grid[y][currentTile.x].down && currentTile.up {
			visitedTiles = append(visitedTiles, grid[y][currentTile.x])
			visitedTiles = stepThoughPipes(grid, visitedTiles, startX, startY)
			return visitedTiles
		}
	}

	y = currentTile.y + 1
	if previousTile.y != y && y < len(grid) {
		if grid[y][currentTile.x].up && currentTile.down {
			visitedTiles = append(visitedTiles, grid[y][currentTile.x])
			visitedTiles = stepThoughPipes(grid, visitedTiles, startX, startY)
			return visitedTiles
		}
	}

	return visitedTiles
}

type Tile struct {
	up    bool
	down  bool
	right bool
	left  bool
	x     int
	y     int
}

func getTileFromAscii(char rune, x int, y int) Tile {
	switch char {
	case '|':
		return Tile{
			up:    true,
			down:  true,
			right: false,
			left:  false,
			x:     x,
			y:     y,
		}
	case '-':
		return Tile{
			up:    false,
			down:  false,
			right: true,
			left:  true,
			x:     x,
			y:     y,
		}
	case 'L':
		return Tile{
			up:    true,
			down:  false,
			right: true,
			left:  false,
			x:     x,
			y:     y,
		}
	case 'J':
		return Tile{
			up:    true,
			down:  false,
			right: false,
			left:  true,
			x:     x,
			y:     y,
		}
	case '7':
		return Tile{
			up:    false,
			down:  true,
			right: false,
			left:  true,
			x:     x,
			y:     y,
		}
	case 'F':
		return Tile{
			up:    false,
			down:  true,
			right: true,
			left:  false,
			x:     x,
			y:     y,
		}
	case 'S':
		return Tile{
			up:    true,
			down:  true,
			right: true,
			left:  true,
			x:     x,
			y:     y,
		}
	default:
		return Tile{
			up:    false,
			down:  false,
			right: false,
			left:  false,
			x:     x,
			y:     y,
		}
	}
}
