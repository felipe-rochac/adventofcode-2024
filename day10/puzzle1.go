package day10

import (
	"adventofcode-2025/common"
)

const fileName = "./inputs/day10.txt"

type position struct {
	row, col int
}

type trailHead struct {
	initialPosition position
	score           int
	visited         []position
}

func gridToTopography(grid [][]rune) (topography [][]int) {
	topography = make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		topography[i] = make([]int, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '.' {
				topography[i][j] = -1
			} else {
				topography[i][j] = common.ParseInt(grid[i][j])
			}
		}
	}
	return topography
}

func findTrailHeads(topography [][]int) (trailHeads []trailHead) {
	for i := 0; i < len(topography); i++ {
		for j := 0; j < len(topography[i]); j++ {
			if topography[i][j] == 0 {
				trailHeads = append(trailHeads, trailHead{initialPosition: position{row: i, col: j}, score: 0})
			}
		}
	}
	return trailHeads
}

func doHiking(topography [][]int, trailHead *trailHead, curPos position, nextPos position) {
	rows := len(topography)
	cols := len(topography[0])

	//fmt.Println("Next position (%d, %d)", nextPos.row, nextPos.col)

	// out of bounds
	if nextPos.row < 0 || nextPos.row >= rows || nextPos.col < 0 || nextPos.col >= cols {
		return
	}

	// Hiking not feasible when next position is bigger than 1 level
	if curPos != nextPos && topography[nextPos.row][nextPos.col]-topography[curPos.row][curPos.col] != 1 {
		return
	}

	if topography[nextPos.row][nextPos.col] == 9 && !common.SliceContains((*trailHead).visited, nextPos) {
		(*trailHead).score++
		(*trailHead).visited = append((*trailHead).visited, nextPos)
		return
	}

	doHiking(topography, trailHead, nextPos, position{row: nextPos.row - 1, col: nextPos.col}) // up
	doHiking(topography, trailHead, nextPos, position{row: nextPos.row + 1, col: nextPos.col}) // down
	doHiking(topography, trailHead, nextPos, position{row: nextPos.row, col: nextPos.col - 1}) // left
	doHiking(topography, trailHead, nextPos, position{row: nextPos.row, col: nextPos.col + 1}) // right
}

func sumScores(trailHeads []trailHead) (sum int) {
	for i := 0; i < len(trailHeads); i++ {
		sum += trailHeads[i].score
	}
	return sum
}

func Puzzle1() int {
	grid := common.ReadFileAsGrid(fileName)

	topography := gridToTopography(grid)

	trailHeads := findTrailHeads(topography)

	for i := 0; i < len(trailHeads); i++ {
		head := &trailHeads[i]
		doHiking(topography, head, head.initialPosition, head.initialPosition)
	}

	return sumScores(trailHeads)
}
