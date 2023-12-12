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

	println((len(pipes) - 1) / 2)
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
	default:
		return Tile{
			up:    true,
			down:  true,
			right: true,
			left:  true,
			x:     x,
			y:     y,
		}
	}
}
