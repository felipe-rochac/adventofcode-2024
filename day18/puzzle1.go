package day18

import (
	"adventofcode-2024/common"
	"fmt"
	"strings"
)

func initializeGrid(size int) [][]rune {
	grid := make([][]rune, size)
	for i := 0; i < size; i++ {
		grid[i] = make([]rune, size)
		for j := 0; j < size; j++ {
			grid[i][j] = '.'
		}
	}

	return grid
}

func parsePoints(content string) (coords []common.Point) {
	lines := strings.Split(content, "\n")
	coords = make([]common.Point, 0)
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		var x, y int
		fmt.Sscanf(lines[i], "%d,%d", &x, &y)
		coords = append(coords, common.Point{X: x, Y: y})
	}
	return coords
}

func fillGridWithInput(grid *[][]rune, points []common.Point) {
	for _, p := range points {
		(*grid)[p.Y][p.X] = '#'
	}
}

func Puzzle1() int {
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
	result := d.MatrixWithObstacles(grid, '#', start, end)

	return result
}
