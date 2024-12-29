package day12

import (
	"adventofcode-2025/common"
	"fmt"
)

type position struct {
	row, col int
}

type garden struct {
	plant     rune
	size      int
	area      int
	perimeter int
	sides     int
	positions []position
}

func calculatePerimeter(positions []position) int {
	findMax := func() (row, col int) {
		row, col = 0, 0
		for _, pos := range positions {
			if pos.row > row {
				row = pos.row
			}
			if pos.col > col {
				col = pos.col
			}
		}
		return row + 1, col + 1
	}

	maxRow, maxCol := findMax()

	dirs := [][]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}

	m := make([][]int, common.MaxInt(maxRow, 1))
	for i := 0; i < maxRow; i++ {
		m[i] = make([]int, maxCol)
	}

	// Creating matrix for positions and filling those in use to determine form
	for i := 0; i < len(positions); i++ {
		pos := positions[i]
		m[pos.row][pos.col] = 1
	}

	perimeter := 0
	for i := 0; i < maxRow; i++ {
		for j := 0; j < maxCol; j++ {
			if m[i][j] == 1 {
				for _, dir := range dirs {
					ni, nj := i+dir[0], j+dir[1]
					if ni < 0 || nj < 0 || ni >= maxRow || nj >= maxCol || m[ni][nj] == 0 {
						// Out of bounds or neighbor is empty, add to perimeter
						perimeter++
					}
				}
			}
		}
	}

	return perimeter
}

func parseGarden(grid [][]rune) []garden {
	rows := len(grid)
	cols := len(grid[0])
	visited := make([][]bool, rows)
	gardens := make([]garden, 0)

	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	var dfs func(row, col int, char rune, g *garden)
	dfs = func(row, col int, char rune, g *garden) {
		if row < 0 || col < 0 || row >= rows || col >= cols || visited[row][col] || grid[row][col] != char {
			return
		}
		visited[row][col] = true
		g.size++
		g.positions = append(g.positions, position{row, col})
		dfs(row+1, col, char, g)
		dfs(row-1, col, char, g)
		dfs(row, col+1, char, g)
		dfs(row, col-1, char, g)
	}

	count := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if !visited[i][j] {
				count++
				g := garden{plant: grid[i][j]}
				dfs(i, j, grid[i][j], &g)
				gardens = append(gardens, g)
			}
		}
	}

	for i := 0; i < len(gardens); i++ {
		gardens[i].area = gardens[i].size
		gardens[i].perimeter += calculatePerimeter(gardens[i].positions)
	}

	return gardens
}

func calculatePrice(garden []garden) int {
	price := 0

	for _, g := range garden {
		fmt.Println("("+string(g.plant)+") area ", g.area, " perimeter ", g.perimeter)
		price += g.perimeter * g.area
	}

	return price
}

const fileName = "./inputs/day12.test.txt"

func Puzzle1() int {
	grid := common.ReadFileAsGrid(fileName)

	gardens := parseGarden(grid)

	return calculatePrice(gardens)
}
