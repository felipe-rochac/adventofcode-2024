package day10

import "adventofcode-2025/common"

func doHikingV2(topography [][]int, trailHead *trailHead, curPos position, nextPos position) {
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

	if topography[nextPos.row][nextPos.col] == 9 {
		(*trailHead).score++
		(*trailHead).visited = append((*trailHead).visited, nextPos)
		return
	}

	doHikingV2(topography, trailHead, nextPos, position{row: nextPos.row - 1, col: nextPos.col}) // up
	doHikingV2(topography, trailHead, nextPos, position{row: nextPos.row + 1, col: nextPos.col}) // down
	doHikingV2(topography, trailHead, nextPos, position{row: nextPos.row, col: nextPos.col - 1}) // left
	doHikingV2(topography, trailHead, nextPos, position{row: nextPos.row, col: nextPos.col + 1}) // right
}

func Puzzle2() int {

	grid := common.ReadFileAsGrid(fileName)

	topography := gridToTopography(grid)

	trailHeads := findTrailHeads(topography)

	for i := 0; i < len(trailHeads); i++ {
		head := &trailHeads[i]
		doHikingV2(topography, head, head.initialPosition, head.initialPosition)
	}

	return sumScores(trailHeads)
}
