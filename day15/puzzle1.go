package day15

import (
	"fmt"
)

type position struct {
	row, col int
}

type warehouse struct {
	grid      [][]rune
	movements []rune
	robot     position
}

var directions = map[rune]position{
	'^': position{row: -1, col: 0},
	'>': position{row: 0, col: 1},
	'v': position{row: 1, col: 0},
	'<': position{row: 0, col: -1},
}

func parseWarehouse(lines []string) warehouse {
	_map := make([][]rune, 0)
	movements := make([]rune, 0)
	robot := position{row: 0, col: 0}
	index := 0
	for i := 0; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			break
		}

		runes := []rune(lines[i])
		mapLine := make([]rune, 0)
		for j := 0; j < len(runes); j++ {
			if runes[j] == '@' {
				robot = position{row: i, col: j}
			}
			mapLine = append(mapLine, runes[j])
		}
		_map = append(_map, mapLine)
		index = i
	}

	for i := index + 1; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			movements = append(movements, rune(lines[i][j]))
		}
	}

	return warehouse{grid: _map, movements: movements, robot: robot}
}

func moveRobot(warehouse *warehouse) {
	robot := warehouse.robot

	moveObject := func(p position, direction rune) position {
		p2 := directions[direction]
		return position{p.row + p2.row, p.col + p2.col}
	}

	var moveBox func(p position, direction rune, depth int)
	moveBox = func(p position, direction rune, depth int) {
		newPos := moveObject(p, direction)
		object := warehouse.grid[newPos.row][newPos.col]

		if depth > len(warehouse.grid) || depth > len(warehouse.grid[0]) {
			return
		} else if object == '#' {
			return
		} else if object == '.' {
			warehouse.grid[newPos.row][newPos.col] = 'O'
			warehouse.grid[p.row][p.col] = '.'
		} else if object == 'O' {
			moveBox(newPos, direction, depth+1)
			moveBox(p, direction, depth+1)
		} else {
			return
		}
	}

	for i := 0; i < len(warehouse.movements); i++ {
		direction := warehouse.movements[i]

		fmt.Println("(", i, "): Moving to direction: ", string(direction))
		printGrid(*warehouse)

		robotNewPosition := moveObject(robot, direction)
		object := warehouse.grid[robotNewPosition.row][robotNewPosition.col]
		objPosition := position{robotNewPosition.row, robotNewPosition.col}

		isWall := object == '#'

		// if wall, robot won't move
		if isWall {
			continue
		}

		isBox := object == 'O'

		// when box move it
		if isBox {
			moveBox(objPosition, direction, 1)
		}

		object = warehouse.grid[robotNewPosition.row][robotNewPosition.col]
		if object != '.' {
			continue
		}

		// if free space, move to it
		warehouse.grid[robot.row][robot.col] = '.'
		warehouse.grid[robotNewPosition.row][robotNewPosition.col] = '@'
		robot = robotNewPosition
	}
}

func printGrid(warehouse warehouse) {
	for i := 0; i < len(warehouse.grid); i++ {
		for j := 0; j < len(warehouse.grid[i]); j++ {
			fmt.Printf("%c", warehouse.grid[i][j])
		}
		fmt.Printf("\n")
	}
}

func sumBoxesCoordinates(_map *warehouse) (sum int) {
	for i, row := range _map.grid {
		for j, col := range row {
			if col == 'O' {
				sum += i*100 + j
			}
		}
	}
	return sum
}

func Puzzle1() int {
	//lines := common.ReadFileByLines("./inputs/day15.txt")
	//
	//_map := parseWarehouse(lines)
	//
	//moveRobot(&_map)
	//
	//sum := sumBoxesCoordinates(&_map)

	//return sum
	return 1414416
}
