package day20

import (
	"adventofcode-2024/common"
	"fmt"
)

type position struct {
	row, col int
}

type edge struct {
	from, to position
	weight   int
}

type state struct {
	pos               position
	steps, cheatsUsed int
	lastCheated       bool
	path              []edge
}

type Cheat struct {
	path  []edge
	saved int // Savings in picoseconds
}

func calculateSavings(path []edge, cheatDuration int) int {
	savings := 0
	for _, edge := range path {
		if edge.weight > cheatDuration {
			savings += edge.weight - cheatDuration
		}
	}
	return savings
}

func buildEdges(grid [][]rune) map[position][]edge {
	rows, cols := len(grid), len(grid[0])
	directions := []position{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	edges := make(map[position][]edge)

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '#' {
				continue // Walls don't generate edges
			}
			curr := position{r, c}
			for _, d := range directions {
				nr, nc := r+d.row, c+d.col
				if nr >= 0 && nc >= 0 && nr < rows && nc < cols && grid[nr][nc] != '#' {
					next := position{nr, nc}
					edges[curr] = append(edges[curr], edge{from: curr, to: next, weight: 1})
				}
			}
		}
	}

	return edges
}

func findPathsWithCheats(grid [][]rune, start, end position, maxCheats int, cheatDuration int) int {
	edges := buildEdges(grid)
	visited := make(map[position]bool)
	queue := []state{{pos: start, steps: 0, cheatsUsed: 0, lastCheated: false, path: nil}}
	validPaths := 0

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.pos == end && curr.cheatsUsed <= maxCheats {
			totalSavings := calculateSavings(curr.path, cheatDuration)
			if totalSavings >= 100 {
				validPaths++
			}
			continue
		}

		for _, edge := range edges[curr.pos] {
			next := edge.to
			newState := state{
				pos:         next,
				steps:       curr.steps + edge.weight,
				cheatsUsed:  curr.cheatsUsed,
				lastCheated: false,
				path:        append(curr.path, edge),
			}

			if grid[next.row][next.col] == '#' && curr.cheatsUsed < maxCheats {
				// Apply a cheat if possible
				newState.cheatsUsed++
				newState.lastCheated = true
				newState.path[len(newState.path)-1].weight -= cheatDuration
			}

			if !visited[next] || (curr.cheatsUsed < maxCheats) {
				queue = append(queue, newState)
			}
		}

		visited[curr.pos] = true
	}

	return validPaths
}

func Puzzle2() int {
	lines := common.ReadFileByLines("./day20/test.txt")

	grid, startA, endA := parseGrid(lines)

	start := position{startA.Y, startA.X}
	end := position{endA.Y, endA.X}
	maxCheats := 1
	cheatDuration := 20

	paths := findPathsWithCheats(grid, start, end, maxCheats, cheatDuration)
	fmt.Printf("Number of paths saving >= 100 picoseconds: %d\n", paths)

	return 0
}
