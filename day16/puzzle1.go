package day16

import (
	"adventofcode-2024/common"
	"fmt"
	"math"
)

type position struct {
	row, col int
}

type location struct {
	location  position
	direction rune
}

type game struct {
	path  [][]location
	score []int
}

func findReindeerAndExit(grid [][]rune) (reindeer, exit position) {
	for r, row := range grid {
		for c, col := range row {
			if col == 'E' {
				exit = position{r, c}
			} else if col == 'S' {
				reindeer = position{r, c}
			}
		}
	}

	return reindeer, exit
}

func seekExit(grid [][]rune, path []location, visited map[position]bool, pos position, direction rune, exitPos position, g *game) {
	if visited[pos] {
		return
	}

	// Mark the position as visited
	visited[pos] = true
	p := location{pos, direction}
	path = append(path, p)

	// If we reach the exit, save the path and backtrack
	if pos == exitPos {
		score := calculateScore(path)
		g.score = append(g.score, score)
		// Store a copy of the path to the game
		g.path = append(g.path, append([]location{}, path...))
		visited[pos] = false // Unmark on backtrack
		return
	}

	directions := []rune{'U', 'D', 'L', 'R'}
	move := func(p position, direction rune) position {
		switch direction {
		case 'U':
			return position{p.row - 1, p.col}
		case 'D':
			return position{p.row + 1, p.col}
		case 'L':
			return position{p.row, p.col - 1}
		case 'R':
			return position{p.row, p.col + 1}
		}
		return p
	}

	// Explore all possible directions
	for _, dir := range directions {
		nextPos := move(pos, dir)
		// Check if the next position is valid (not a wall and within bounds)
		if nextPos.row >= 0 && nextPos.row < len(grid) &&
			nextPos.col >= 0 && nextPos.col < len(grid[0]) &&
			grid[nextPos.row][nextPos.col] != '#' {
			seekExit(grid, path, visited, nextPos, dir, exitPos, g)
		}
	}

	// Backtrack
	visited[pos] = false
	path = path[:len(path)-1]
}

func isClockwise(prev, after rune) bool {
	return ((prev == 'L' || prev == 'R') && (after == 'U' || after == 'D')) || ((prev == 'U' || prev == 'D') && (after == 'L' || after == 'R'))
}

func calculateScore(path []location) int {
	dir := path[0].direction
	score := 0
	for i := 1; i < len(path); i++ {
		if isClockwise(dir, path[i].direction) {
			score += 1001
		} else {
			score += 1
		}
		dir = path[i].direction
	}

	return score
}

func copyMatrix(src [][]rune) [][]rune {
	// Create a new matrix with the same dimensions as the source
	dst := make([][]rune, len(src))
	for i := range src {
		// Allocate a new slice for each row and copy the elements
		dst[i] = make([]rune, len(src[i]))
		copy(dst[i], src[i])
	}
	return dst
}

func printPath(grid [][]rune, path []location, score, index int) {
	g := copyMatrix(grid)

	for _, p := range path {
		g[p.location.row][p.location.col] = p.direction
	}

	fmt.Println(fmt.Sprintf(`%d: Movements %d Score %d`, index, len(path), score))
	for _, row := range g {
		for _, col := range row {
			fmt.Print(string(col))
		}
		fmt.Println()
	}
	fmt.Println()
}

func Puzzle1() int {
	grid := common.ReadFileByGrid("./day16/day16.txt")

	reindeer, exit := findReindeerAndExit(grid)
	visited := make(map[position]bool)
	path := []location{}
	g := &game{path: [][]location{}}

	seekExit(grid, path, visited, reindeer, 'L', exit, g)

	lowestStepsIndex := -1
	lowestScore := math.MaxInt32

	for i, _ := range g.path {
		if g.score[i] < lowestScore {
			lowestScore = g.score[i]
			lowestStepsIndex = i
		}
	}

	printPath(grid, g.path[lowestStepsIndex], g.score[lowestStepsIndex], lowestStepsIndex)

	return lowestScore
}
