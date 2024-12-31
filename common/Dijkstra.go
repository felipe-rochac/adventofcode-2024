package common

import (
	"container/heap"
	"math"
)

type Point struct {
	X, Y int
}

type PriorityQueueItem struct {
	Point    Point
	Distance int
}

type PriorityQueue []PriorityQueueItem

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Distance < pq[j].Distance }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(PriorityQueueItem))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

type Dijkstra struct {
}

func (dijkstra *Dijkstra) Matrix(matrix [][]int, start, end Point) int {
	rows := len(matrix)
	cols := len(matrix[0])

	// Distance map
	distances := make([][]int, rows)
	for i := 0; i < rows; i++ {
		distances[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			distances[i][j] = math.MaxInt32
		}
	}
	distances[start.X][start.Y] = 0

	// Priority queue
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, PriorityQueueItem{Point: start, Distance: 0})

	// Directions for moving up, down, left, and right
	directions := []Point{
		{X: -1, Y: 0}, // Up
		{X: 1, Y: 0},  // Down
		{X: 0, Y: -1}, // Left
		{X: 0, Y: 1},  // Right
	}

	for pq.Len() > 0 {
		current := heap.Pop(pq).(PriorityQueueItem)
		currPoint := current.Point
		currDistance := current.Distance

		// If we reach the end point
		if currPoint == end {
			return currDistance
		}

		// Explore neighbors
		for _, d := range directions {
			neighbor := Point{X: currPoint.X + d.X, Y: currPoint.Y + d.Y}
			if neighbor.X >= 0 && neighbor.X < rows && neighbor.Y >= 0 && neighbor.Y < cols {
				newDist := currDistance + matrix[neighbor.X][neighbor.Y]
				if newDist < distances[neighbor.X][neighbor.Y] {
					distances[neighbor.X][neighbor.Y] = newDist
					heap.Push(pq, PriorityQueueItem{Point: neighbor, Distance: newDist})
				}
			}
		}
	}

	return -1 // Return -1 if no path exists
}

// Find shortest path between start and end on the grid, using obstacle as stoppers
func (dijkstra *Dijkstra) MatrixWithObstacles(grid [][]rune, obstacle rune, start, end Point) int {
	rows := len(grid)
	cols := len(grid[0])

	// Distance map
	distances := make([][]int, rows)
	for i := range distances {
		distances[i] = make([]int, cols)
		for j := range distances[i] {
			distances[i][j] = math.MaxInt32
		}
	}
	distances[start.X][start.Y] = 0

	// Priority queue
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, PriorityQueueItem{Point: start, Distance: 0})

	// Directions for moving up, down, left, and right
	directions := []Point{
		{X: -1, Y: 0}, // Up
		{X: 1, Y: 0},  // Down
		{X: 0, Y: -1}, // Left
		{X: 0, Y: 1},  // Right
	}

	for pq.Len() > 0 {
		current := heap.Pop(pq).(PriorityQueueItem)
		currPoint := current.Point
		currDistance := current.Distance

		// If we reach the end point
		if currPoint == end {
			return currDistance
		}

		// Explore neighbors
		for _, d := range directions {
			neighbor := Point{X: currPoint.X + d.X, Y: currPoint.Y + d.Y}
			if neighbor.X >= 0 && neighbor.X < rows && neighbor.Y >= 0 && neighbor.Y < cols {
				// Skip obstacles
				if grid[neighbor.X][neighbor.Y] == obstacle {
					continue
				}

				// Calculate cost based on the rune value
				cost := 1 // Default cost for traversable cells
				newDist := currDistance + cost
				if newDist < distances[neighbor.X][neighbor.Y] {
					distances[neighbor.X][neighbor.Y] = newDist
					heap.Push(pq, PriorityQueueItem{Point: neighbor, Distance: newDist})
				}
			}
		}
	}

	return -1 // Return -1 if no path exists
}
