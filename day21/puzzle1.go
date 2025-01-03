package day21

import (
	"adventofcode-2024/common"
	"math"
)

type position struct {
	row, col int
}

type edge struct {
	to     position
	weight int
}

type actor struct {
	curr, dest position
	pad        [][]rune
}

type path struct {
	current position
	actions []rune
}

func buildGraph(matrix [][]rune) map[position][]edge {
	rows, cols := len(matrix), len(matrix[0])
	directions := []position{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // Right, Down, Left, Up
	graph := make(map[position][]edge)

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			curr := position{r, c}
			for _, d := range directions {
				nr, nc := r+d.row, c+d.col
				if nr >= 0 && nc >= 0 && nr < rows && nc < cols {
					next := position{nr, nc}
					graph[curr] = append(graph[curr], edge{to: next, weight: 1})
				}
			}
		}
	}

	return graph
}

func getKeypadGraph() [][]rune {
	return [][]rune{
		{'7', '8', '9'},
		{'4', '5', '6'},
		{'1', '2', '3'},
		{' ', '0', 'A'},
	}
}

func getDestinationKeypadPosition(code rune) position {
	switch code {
	case '9':
		return position{0, 2}
	case '8':
		return position{0, 1}
	case '7':
		return position{0, 0}
	case '6':
		return position{1, 2}
	case '5':
		return position{1, 1}
	case '4':
		return position{1, 0}
	case '3':
		return position{2, 2}
	case '2':
		return position{2, 1}
	case '1':
		return position{2, 0}
	case '0':
		return position{3, 1}
	default:
		return position{3, 2}
	}
}

func getDirectionalPadGraph() [][]rune {
	return [][]rune{
		{' ', '^', 'A'},
		{'<', 'v', '>'},
	}
}

func getDestinationDirectionalpadPosition(code rune) position {
	switch code {
	case '^':
		return position{0, 1}
	case '<':
		return position{1, 0}
	case 'v':
		return position{1, 1}
	case '>':
		return position{1, 2}
	default:
		return position{0, 2}
	}
}

func getShortestPath(actor actor) []rune {
	rows := len(actor.pad)
	cols := len(actor.pad[0])
	start := actor.curr
	end := actor.dest
	visited := make(map[position]bool)

	// Distance map
	distances := make([][]int, rows)
	for i := 0; i < rows; i++ {
		distances[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			distances[i][j] = math.MaxInt32
		}
	}
	distances[start.row][start.col] = 0

	// Priority queue
	queue := []path{{start, []rune{}}}

	// Directions for moving up, down, left, and right
	directions := []position{
		{row: -1, col: 0}, // Up
		{row: 1, col: 0},  // Down
		{row: 0, col: -1}, // Left
		{row: 0, col: 1},  // Right
	}

	action := map[position]rune{
		{row: -1, col: 0}: '^',
		{row: 1, col: 0}:  'v',
		{row: 0, col: 1}:  '>',
		{row: 0, col: -1}: '<',
	}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		// If we reach the end point
		if curr.current == end {
			return curr.actions
		}

		// Explore neighbors
		for _, d := range directions {
			next := position{curr.current.row + d.row, curr.current.col + d.col}

			if next.row < 0 || next.col < 0 || next.row >= rows || next.col >= cols {
				continue
			}

			if visited[next] {
				continue
			}

			queue = append(queue, path{next, append(curr.actions, action[d])})
		}
	}

	return []rune{}
}

func pressDirectionalPad(actions []rune, actor *actor) []rune {
	requiredActions := make([]rune, 0)
	for _, a := range actions {
		actor.dest = getDestinationDirectionalpadPosition(a)
		directionActions := getShortestPath(*actor)
		actor.curr = actor.dest
		directionActions = append(directionActions, 'A')
		requiredActions = append(requiredActions, directionActions...)
	}
	return requiredActions
}

func Puzzle1() int {
	codes := common.ReadFileByLines("./day21/test.txt")

	keypad := getKeypadGraph()
	directionPad := getDirectionalPadGraph()

	initialPosKeypad := getDestinationKeypadPosition('A')
	initialPosDirectionalPad := getDestinationDirectionalpadPosition('A')

	commander := actor{initialPosDirectionalPad, position{0, 0}, directionPad}
	directionRobot1 := actor{initialPosDirectionalPad, position{0, 0}, directionPad}
	directionRobot2 := actor{initialPosDirectionalPad, position{0, 0}, directionPad}
	keypadRobot := actor{initialPosKeypad, position{0, 0}, keypad}
	pressedKeys := make([]rune, 0)

	for _, code := range codes {
		for _, r := range code {
			keypadRobot.dest = getDestinationKeypadPosition(r)
			actions := getShortestPath(keypadRobot)
			actions = append(actions, 'A')

			actions2 := pressDirectionalPad(actions, &directionRobot2)

			actions3 := pressDirectionalPad(actions2, &directionRobot1)

			actions4 := pressDirectionalPad(actions3, &commander)

			pressedKeys = append(pressedKeys, actions4...)
		}
	}

	return 0
}
