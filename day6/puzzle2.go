package day6

//func isCycle(coord Coordinate, direction Direction, rows int, cols int, input [][]rune) bool {
//	obstacleSeen := make(map[Coordinate]Direction)
//	for isInGrid(coord, rows, cols) {
//		if input[coord.Row][coord.Col] == '#' {
//			_, ok := obstacleSeen[coord]
//			if ok {
//				return true
//			}
//			obstacleSeen[coord] = direction
//			coord = getPrevCoordinate(coord, direction)
//			direction = getNextDirection(direction)
//		}
//		coord = getNextCoordinate(coord, direction)
//	}
//	return false
//}

func stuckInLoop(guard Position, direction rune, _map [][]rune) bool {
	obstaclesFound := make(map[struct {
		Coord     Position
		Direction rune
	}]bool)

	for guardBeingSeen(guard, rows, cols) {
		if _map[guard.row][guard.col] == '#' {
			// Create a composite key of (Coordinate, Direction)
			key := struct {
				Coord     Position
				Direction rune
			}{Coord: guard, Direction: direction}

			// Check if this (Coordinate, Direction) pair was already seen
			if obstaclesFound[key] {
				return true
			}

			// Mark this obstacle and direction as seen
			obstaclesFound[key] = true

			// Move to the previous coordinate and change direction
			guard = backStep(guard, direction)
			direction = rotations[direction]
		}

		// Move to the next coordinate in the current direction
		guard = moveForward(guard, direction)
	}

	// No cycle detected
	return false
}

func Puzzle2() int {

	guardCoord := findGuard(_map)
	moves := 0

	_map2 := make([][]rune, len(_map))
	for i, row := range _map {
		_map2[i] = make([]rune, len(row))
		copy(_map2[i], row)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			temp := _map2[i][j]
			if temp == '#' {
				continue
			}

			_map2[i][j] = '#'
			tempGuardCoord := guardCoord
			direction := '^'

			if stuckInLoop(tempGuardCoord, direction, _map2) {
				moves++
			}

			_map2[i][j] = temp
		}
	}

	// < 11021
	return moves
}
