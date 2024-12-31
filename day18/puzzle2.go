package day18

import (
	"adventofcode-2024/common"
	"fmt"
)

func Puzzle2() int {
	content := common.ReadFileText("./day18/input.txt")

	// size := 7
	// bytes := 12
	size := 71
	bytes := 1024

	// Initialize the grid
	grid := initializeGrid(size)

	// Parse input into the grid
	coords := parsePoints(content)

	fillGridWithInput(&grid, coords[:bytes])

	start := common.Point{X: 0, Y: 0}
	end := common.Point{X: size - 1, Y: size - 1}
	var d common.Dijkstra

	for i := bytes; i < len(coords); i++ {
		c := coords[i]
		grid[c.Y][c.X] = '#'
		if d.MatrixWithObstacles(grid, '#', start, end) == -1 {
			fmt.Printf("Coordinate that will prevent exit is (%d, %d)\n", c.X, c.Y)
			break
		}
	}

	return 0
}
