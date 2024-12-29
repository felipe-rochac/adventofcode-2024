package day12

import (
	"adventofcode-2025/common"
	"fmt"
)

func calculateSides(grid [][]rune, plant rune) int {
	rows := len(grid)
	cols := len(grid[0])
	sides := 0

	directions := []struct{ x, y int }{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == plant {
				for _, dir := range directions {
					ni, nj := i+dir.x, j+dir.y
					if ni < 0 || nj < 0 || ni >= rows || nj >= cols || grid[ni][nj] != plant {
						sides++
					}
				}
			}
		}
	}

	return sides
}

func calculatePriceV2(garden []garden) int {
	price := 0

	for _, g := range garden {
		fmt.Println("("+string(g.plant)+") area ", g.area, " perimeter ", g.perimeter)
		price += g.size * g.area
	}

	return price
}

func Puzzle2() int {

	grid := common.ReadFileAsGrid(fileName)

	gardens := parseGarden(grid)

	for i := 0; i < len(gardens); i++ {
		gardens[i].sides = calculateSides(grid, gardens[i].plant)
		fmt.Println("Plant ", string(gardens[i].plant), " has ", gardens[i].sides, " sides")
	}

	return calculatePriceV2(gardens)
}
