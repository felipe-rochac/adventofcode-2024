package day20

import (
	"adventofcode-2024/common"
)

func parseGrid(input []string) ([][]rune, common.Point, common.Point) {
	var start, end common.Point
	grid := make([][]rune, len(input))
	for i, line := range input {
		grid[i] = []rune(line)
		for j, cell := range grid[i] {
			if cell == 'S' {
				start = common.Point{X: i, Y: j}
			} else if cell == 'E' {
				end = common.Point{X: i, Y: j}
			}
		}
	}
	return grid, start, end
}

func countWalls(grid [][]rune) (int, []common.Point) {
	count := 0
	boundary := make([]common.Point, 0)
	for i, row := range grid {
		for j, col := range row {

			// not counting boundary
			if i == 0 || j == 0 {
				boundary = append(boundary, common.Point{X: j, Y: i})
				continue
			}

			if i == len(grid)-1 || j == len(row)-1 {
				boundary = append(boundary, common.Point{X: j, Y: i})
				continue
			}

			if col != '#' {
				continue
			}

			count++
		}
	}

	return count, boundary
}

func Puzzle1() int {
	lines := common.ReadFileByLines("./day20/test.txt")

	grid, start, end := parseGrid(lines)

	_, boundary := countWalls(grid)

	cheatsPosition := make(map[common.Point]bool)
	distances := make([]int, 0)

	// Ignoring boundary
	for _, b := range boundary {
		// obstaclesCount++
		cheatsPosition[b] = true
	}

	var d common.Dijkstra
	noCheatDistance := d.MatrixWithObstaclesAndCheat(grid, '#', start, end, 0, cheatsPosition)

	for {
		distance := d.MatrixWithObstaclesAndCheat(grid, '#', start, end, 1, cheatsPosition)

		if distance == -1 {
			break
		}

		distances = append(distances, distance)
	}

	summary := make(map[int]int)
	save100picosec := 0

	for _, d := range distances {
		if d == noCheatDistance {
			continue
		}
		if (noCheatDistance - d) >= 100 {
			save100picosec++
		}
		summary[d]++
	}

	// for key, value := range summary {
	// 	fmt.Println(fmt.Sprintf("There are %d cheats that save %d picoseconds", value, noCheatDistance-key))
	// }

	return save100picosec
}
