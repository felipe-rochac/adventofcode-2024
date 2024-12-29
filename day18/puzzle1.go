package day18

import (
	"adventofcode-2024/common"
	"fmt"
	"math"
	"strings"
)

type coordinate struct {
	row, col int
	visited  map[[2]int]bool
	path     [][2]int
}

func parseCoordinates(lines []string) []coordinate {
	coordinates := make([]coordinate, 0)
	for _, line := range lines {
		parts := strings.Split(line, ",")
		coordinates = append(coordinates, coordinate{common.ParseStrToInt(parts[1]), common.ParseStrToInt(parts[0]), nil, nil})
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
		if coord.col < rows && coord.row < cols { // Ensure within bounds
			m[coord.row][coord.col] = '#'
		}
	}

	return m
}

func coordValid(row, col, rows, cols int) bool {
	if col < 0 || col >= cols {
		return false
	}

	if row < 0 || row >= rows {
		return false
	}

	return true
}

func findShortestPath(memory [][]rune, start, exit coordinate) (paths []coordinate, shortest int) {
	var queue common.Queue[coordinate]
	queue.Push(start)
	paths = make([]coordinate, 0)
	shortest = math.MaxInt32

	for !queue.IsEmpty() {
		coord, _ := queue.Pop()
		c := [2]int{coord.row, coord.col}

		if !coordValid(coord.row, coord.col, len(memory), len(memory[0])) {
			continue
		}

		if memory[coord.row][coord.col] == '#' {
			continue
		}

		if coord.row == exit.row && coord.col == exit.col && len(coord.path) < 25 {
			newPath := append(coord.path, c)
			paths = append(paths, coordinate{coord.row, coord.col, coord.visited, newPath})
			if shortest > len(newPath) {
				shortest = len(newPath)
			}
			continue
		}

		// mark as visited
		newVisited := make(map[[2]int]bool)
		for k, v := range coord.visited {
			newVisited[k] = v
		}
		newVisited[c] = true

		directions := [][2]int{
			{coord.row - 1, coord.col}, // top
			{coord.row + 1, coord.col}, // bottom
			{coord.row, coord.col - 1}, // left
			{coord.row, coord.col + 1}, // right
		}

		for _, dir := range directions {
			if dir[0] >= 0 && dir[0] < len(memory) && dir[1] >= 0 && dir[1] < len(memory[0]) && !newVisited[dir] {
				queue.Push(coordinate{dir[0], dir[1], newVisited, append(coord.path, c)})
			}
		}
	}

	return paths, shortest
}

func printMap(memory [][]rune, coords [][2]int) {
	var last [2]int
	for _, c := range coords {
		memory[c[0]][c[1]] = 'O'
		last = c
	}

	var builder strings.Builder

	// Print column numbers
	builder.WriteString("    ") // Space for row numbers
	for j := 0; j < len(memory[0]); j++ {
		builder.WriteString(fmt.Sprintf("%2d ", j))
	}
	builder.WriteString("\n")

	// Print each row with row numbers
	for i := 0; i < len(memory); i++ {
		builder.WriteString(fmt.Sprintf("%2d |", i)) // Row number
		for j := 0; j < len(memory[i]); j++ {
			builder.WriteString(fmt.Sprintf(" %s ", string(memory[i][j])))
		}
		builder.WriteString("\n")
	}

	fmt.Printf("Last Coordinate: (%d, %d)\n", last[0], last[1])
	fmt.Println(builder.String())
}

func Puzzle1() int {
	lines := common.ReadFileByLines("./day18/test.txt")

	visited := make(map[[2]int]bool)
	path := make([][2]int, 0)
	start := coordinate{0, 0, visited, path}
	end := coordinate{6, 6, nil, nil}

	coordinates := parseCoordinates(lines)

	memory := generateMemoryMap(coordinates[:12], end.col+1, end.row+1)

	paths, shortest := findShortestPath(memory, start, end)
	fmt.Printf("Shortest path is %d\n", shortest)

	for _, p := range paths {
		if len(p.path) != shortest {
			continue
		}

		fmt.Println("")
		printMap(memory, p.path)
		fmt.Println("")
	}

	return 0
}
