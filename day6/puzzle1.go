package day6

import (
	"adventofcode-2025/common"
)

type Position struct {
	row, col int
}

var rotations = map[rune]rune{
	'^': '>',
	'>': 'v',
	'v': '<',
	'<': '^',
}

var directions = map[rune]Position{
	'^': Position{row: -1, col: 0},
	'>': Position{row: 0, col: 1},
	'v': Position{row: 1, col: 0},
	'<': Position{row: 0, col: -1},
}

var rows, cols int
var _map [][]rune

func findGuard(_map [][]rune) Position {
	for r, row := range _map {
		for c, col := range row {
			_, ok := rotations[col]

			if ok {
				return Position{row: r, col: c}
			}
		}
	}

	panic("Guard not found")
}

func guardBeingSeen(guardPosition Position, rows int, cols int) bool {
	return guardPosition.row < rows && guardPosition.row > -1 && guardPosition.col < cols && guardPosition.col > -1
}

func backStep(guard Position, direction rune) Position {
	p := directions[direction]
	return Position{row: guard.row - p.row, col: guard.col - p.col}
}

func moveForward(guard Position, direction rune) Position {
	p := directions[direction]
	return Position{row: guard.row + p.row, col: guard.col + p.col}
}

func Puzzle1() int {
	_map = common.ReadFileAsGrid("./inputs/day6.txt")
	rows = len(_map)
	cols = len(_map[0])

	guardCoord := findGuard(_map)
	direction := '^'
	//initialPosition := guardCoord

	moves := 0
	seen := make(map[Position]struct{})

	for guardBeingSeen(guardCoord, rows, cols) {
		if _map[guardCoord.row][guardCoord.col] == '#' {
			guardCoord = backStep(guardCoord, direction)
			direction = rotations[direction]
		} else {
			if _, alreadySeen := seen[guardCoord]; !alreadySeen {
				seen[guardCoord] = struct{}{}
				moves++
			}
			guardCoord = moveForward(guardCoord, direction)
		}
	}

	// = 5086
	return moves
}
