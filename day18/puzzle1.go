package day18

import (
	"adventofcode-2024/common"
	"fmt"
	"math"
	"strings"
)

type coordinate struct {
	x, y int
	path [][2]int
}

func parseCoordinates(lines []string) []coordinate {
	coordinates := make([]coordinate, 0)
	for _, line := range lines {
		parts := strings.Split(line, ",")
		coordinates = append(coordinates, coordinate{common.ParseStrToInt(parts[0]), common.ParseStrToInt(parts[1]), nil})
	}

	return coordinates
}

func generateMemoryMap(coords []coordinate, rows, cols int) [][]rune {
	// Initialize the grid with '.'
	m := make([][]rune, cols)
	for i := range m {
		m[i] = make([]rune, rows)
		for j := range m[i] {
			m[i][j] = '.'
		}
	}

	// Mark the coordinates with '#'
	for _, coord := range coords {
		if coord.y < rows && coord.x < cols { // Ensure within bounds
			m[coord.x][coord.y] = '#'
		}
	}

	return m
}

func coordValid(coord coordinate, rows, cols int) bool {
	if coord.y < 0 || coord.y >= rows {
		return false
	}

	if coord.x < 0 || coord.x >= cols {
		return false
	}

	return true
}

func findShortestPath(grid [][]rune, start, exit coordinate) (paths []coordinate, shortest int) {
	var queue common.Queue[coordinate]
	queue.Push(start)
	paths = make([]coordinate, 0)
	visited := make(map[[2]int]bool)
	shortest = math.MaxInt32

	for !queue.IsEmpty() {

		coord, _ := queue.Pop()
		c := [2]int{coord.x, coord.y}

		if !coordValid(coord, len(grid[0]), len(grid)) {
			continue
		}

		// corrupted memory hit
		if grid[coord.x][coord.y] == '#' {
			continue
		}

		// exit
		if coord.x == exit.x && coord.y == exit.y {
			coord.path = append(coord.path, c)
			paths = append(paths, coord)
			if shortest > len(coord.path) {
				shortest = len(coord.path)
			}
			continue
		}

		// mark as visited
		coord.path = append(coord.path, c)
		visited[c] = true

		directions := [][2]int{
			{coord.x - 1, coord.y}, // left
			{coord.x + 1, coord.y}, // right
			{coord.x, coord.y - 1}, // top
			{coord.x, coord.y + 1}, // bottom
		}

		for _, dir := range directions {
			if !visited[dir] {
				queue.Push(coordinate{dir[0], dir[1], coord.path})
			}
		}
	}

	return paths, shortest
}

func printMap(memory [][]rune, coords [][2]int) {
	for _, c := range coords {
		memory[c[0]][c[1]] = 'O'
	}

	var builder strings.Builder
	for i := 0; i < len(memory); i++ {
		for j := 0; j < len(memory[i]); j++ {
			builder.WriteString(string(memory[j][i]))
		}
		builder.WriteString("\n")
	}

	fmt.Println(builder.String())
}

func Puzzle1() int {
	lines := common.ReadFileByLines("./day18/test.txt")

	path := make([][2]int, 0)
	start := coordinate{0, 0, path}
	end := coordinate{6, 6, nil}

	coordinates := parseCoordinates(lines)

	memory := generateMemoryMap(coordinates[:12], end.y+1, end.x+1)

	paths, shortest := findShortestPath(memory, start, end)
	fmt.Printf("Shortest path is %d\n", shortest)

	for _, p := range paths {
		if len(p.path) > shortest {
			continue
		}

		fmt.Println("")
		printMap(memory, p.path)
		fmt.Println("")
	}

	return 0
}
