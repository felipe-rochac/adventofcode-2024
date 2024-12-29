package day15

import (
	"adventofcode-2025/common"
	"fmt"
)

func parseWideWarehouse(lines []string) warehouse {
	_map := make([][]rune, 0)
	movements := make([]rune, 0)
	robot := position{row: 0, col: 0}
	index := 0
	getWiderObject := func(obj rune) (first rune, second rune) {
		if obj == '#' {
			return '#', '#'
		} else if obj == 'O' {
			return '[', ']'
		} else if obj == '.' {
			return '.', '.'
		} else if obj == '@' {
			return '@', '.'
		}
		return obj, '?'
	}

	for i := 0; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			break
		}

		runes := []rune(lines[i])
		mapLine := make([]rune, 0)
		for j := 0; j < len(runes); j++ {
			first, second := getWiderObject(runes[j])
			mapLine = append(mapLine, first)
			mapLine = append(mapLine, second)
		}
		_map = append(_map, mapLine)
		index = i
	}

	for i := 0; i < len(_map); i++ {
		runes := _map[i]
		found := false
		for j := 0; j < len(runes); j++ {
			if runes[j] == '@' {
				robot = position{row: i, col: j}
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	for i := index + 1; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			movements = append(movements, rune(lines[i][j]))
		}
	}

	return warehouse{grid: _map, movements: movements, robot: robot}
}

func moveRobotV2(warehouse *warehouse) {
	robot := warehouse.robot

	moveObject := func(p position, direction rune) position {
		coords := directions[direction]
		return position{p.row + coords.row, p.col + coords.col}
	}

	var moveBox func(current position, direction rune)
	moveBox = func(current position, direction rune) {
		var chain []position

		// Detect the chain of boxes
		for {
			chain = append(chain, current)
			next := moveObject(current, direction)
			if warehouse.grid[next.row][next.col] == '.' {
				break
			}
			current = next
		}

		// Check if there is enough free space to move the entire chain
		end := moveObject(current, direction)
		if warehouse.grid[end.row][end.col] != '.' {
			return // Not enough space to move
		}

		// Move the entire chain from the back to the front
		for i := len(chain) - 1; i >= 0; i-- {
			newPos := moveObject(chain[i], direction)
			warehouse.grid[newPos.row][newPos.col] = warehouse.grid[chain[i].row][chain[i].col]
			warehouse.grid[chain[i].row][chain[i].col] = '.'
		}
	}

	for i, direction := range warehouse.movements {
		robotNewPosition := moveObject(robot, direction)
		object := warehouse.grid[robotNewPosition.row][robotNewPosition.col]

		fmt.Println("Movement (", i, "): ", string(direction))
		printGrid(*warehouse)

		// Skip if hitting a wall
		if object == '#' {
			continue
		}

		// Handle box movement
		if object == '[' {
			moveBox(robotNewPosition, direction)
			if direction == '^' || direction == 'v' {
				moveBox(position{robotNewPosition.row, robotNewPosition.col + 1}, direction)
			}
		} else if object == ']' {
			moveBox(robotNewPosition, direction)
			if direction == '^' || direction == 'v' {
				moveBox(position{robotNewPosition.row, robotNewPosition.col - 1}, direction)
			}
		}

		// Move the robot if the space is free
		if warehouse.grid[robotNewPosition.row][robotNewPosition.col] == '.' {
			warehouse.grid[robot.row][robot.col] = '.'
			warehouse.grid[robotNewPosition.row][robotNewPosition.col] = '@'
			robot = robotNewPosition
		}
	}
}

func Puzzle2() int {

	lines := common.ReadFileByLines("./inputs/day15.test3.txt")

	_map := parseWideWarehouse(lines)

	moveRobotV2(&_map)

	sum := sumBoxesCoordinates(&_map)

	return sum
}
