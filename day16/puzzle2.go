package day16

import (
	"adventofcode-2024/common"
	"container/heap"
)

func canVisit(visited map[[3]int]int, d, x, y, score int) bool {
	prevScore, exists := visited[[3]int{d, x, y}]
	if exists && prevScore < score {
		return false
	}
	visited[[3]int{d, x, y}] = score
	return true
}

func seekExitV2(grid [][]rune, start, end [2]int) int {
	// Replace 'E' with '.' in the grid
	grid[end[0]][end[1]] = '.'

	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	heapQueue := &PriorityQueue{}
	heap.Init(heapQueue)
	heap.Push(heapQueue, Point{score: 0, direction: 0, x: start[0], y: start[1], path: map[[2]int]bool{start: true}})
	visited := make(map[[3]int]int)

	lowestScore := -1
	winningPaths := make(map[[2]int]bool)

	for heapQueue.Len() > 0 {
		PointV2 := heap.Pop(heapQueue).(Point)
		score, d, x, y, path := PointV2.score, PointV2.direction, PointV2.x, PointV2.y, PointV2.path

		if lowestScore != -1 && lowestScore < score {
			break
		}

		if [2]int{x, y} == end {
			if lowestScore == -1 || score < lowestScore {
				lowestScore = score
			}
			for key := range path {
				winningPaths[key] = true
			}
			continue
		}

		if !canVisit(visited, d, x, y, score) {
			continue
		}

		// Move forward
		nextX, nextY := x+directions[d][0], y+directions[d][1]
		if grid[nextX][nextY] == '.' && canVisit(visited, d, nextX, nextY, score+1) {
			newPath := make(map[[2]int]bool)
			for key := range path {
				newPath[key] = true
			}
			newPath[[2]int{nextX, nextY}] = true
			heap.Push(heapQueue, Point{score: score + 1, direction: d, x: nextX, y: nextY, path: newPath})
		}

		// Turn left
		left := (d - 1 + 4) % 4
		if canVisit(visited, left, x, y, score+1000) {
			heap.Push(heapQueue, Point{score: score + 1000, direction: left, x: x, y: y, path: path})
		}

		// Turn right
		right := (d + 1) % 4
		if canVisit(visited, right, x, y, score+1000) {
			heap.Push(heapQueue, Point{score: score + 1000, direction: right, x: x, y: y, path: path})
		}
	}

	return len(winningPaths)
}

func Puzzle2() int {
	grid := common.ReadFileByGrid("./day16/day16.txt")

	start, exit := findReindeerAndExit(grid)

	score := seekExitV2(grid, [2]int{start.x, start.y}, [2]int{exit.x, exit.y})

	// < 500716
	return score
}
