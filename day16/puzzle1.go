package day16

import (
	"adventofcode-2024/common"
	"container/heap"
)

type Point struct {
	score, direction, x, y int
	path                   map[[2]int]bool
}

// PriorityQueue implements a priority queue for Points
type PriorityQueue []Point

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].score < pq[j].score } // Min-Heap based on score
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Point))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func seekExit(grid [][]rune, start, end [2]int) int {
	directions := [][2]int{
		{0, 1},  // Right
		{1, 0},  // Down
		{0, -1}, // Left
		{-1, 0}, // Up
	}

	grid[end[0]][end[1]] = '.'
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Point{score: 0, direction: 0, x: start[0], y: start[1]})

	visited := make(map[[3]int]bool)

	for pq.Len() > 0 {
		point := heap.Pop(pq).(Point)
		score, d, i, j := point.score, point.direction, point.x, point.y

		if [2]int{i, j} == end {
			return score
		}

		if visited[[3]int{d, i, j}] {
			continue
		}

		visited[[3]int{d, i, j}] = true

		// Move forward
		x, y := i+directions[d][0], j+directions[d][1]
		if grid[x][y] == '.' && !visited[[3]int{d, x, y}] {
			heap.Push(pq, Point{score: score + 1, direction: d, x: x, y: y})
		}

		// Turn left
		left := (d - 1 + 4) % 4
		if !visited[[3]int{left, i, j}] {
			heap.Push(pq, Point{score: score + 1000, direction: left, x: i, y: j})
		}

		// Turn right
		right := (d + 1) % 4
		if !visited[[3]int{right, i, j}] {
			heap.Push(pq, Point{score: score + 1000, direction: right, x: i, y: j})
		}
	}

	return -1 // No path found
}

func findReindeerAndExit(grid [][]rune) (reindeer, exit Point) {
	for r, row := range grid {
		for c, col := range row {
			if col == 'E' {
				exit = Point{0, 0, r, c, nil}
			} else if col == 'S' {
				reindeer = Point{0, 0, r, c, nil}
			}
		}
	}

	return reindeer, exit
}

func Puzzle1() int {
	grid := common.ReadFileByGrid("./day16/day16.txt")

	start, exit := findReindeerAndExit(grid)

	score := seekExit(grid, [2]int{start.x, start.y}, [2]int{exit.x, exit.y})

	// < 500716
	return score
}
