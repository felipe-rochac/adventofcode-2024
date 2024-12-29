package day14

import (
	"adventofcode-2025/common"
	"fmt"
	"strconv"
	"strings"
)

type position struct {
	x, y int
}

type velocity struct {
	x, y int
}

type robot struct {
	id              int
	initialPosition position
	velocity        velocity
	direction       string
	lastPosition    position
}

type tile struct {
	x, y   int
	robots []robot
}

type space struct {
	id                     int
	tiles                  [][]tile
	haveTwoRobotsSamePlace bool
}

type quadrant struct {
	id         int
	robotCount int
	tiles      [][]tile
}

func removeRobot(tile *tile, robot robot) {
	for i := 0; i < len(tile.robots); i++ {
		if tile.robots[i].id == robot.id {
			tile.robots = append(tile.robots[:i], tile.robots[i+1:]...)
		}
	}
}

func addRobot(tile *tile, robot robot) {

	for i := 0; i < len(tile.robots); i++ {
		// to avoid adding the same robot twice
		if tile.robots[i].id == robot.id {
			return
		}
	}
	tile.robots = append(tile.robots, robot)
}

func parseRobots(lines []string) []robot {
	robots := make([]robot, len(lines))
	determineDirection := func(velocity velocity) string {
		switch {
		case velocity.x > 0 && velocity.y > 0:
			return ">v"
		case velocity.x > 0 && velocity.y <= 0:
			return ">^"
		case velocity.x <= 0 && velocity.y > 0:
			return "<v"
		default:
			return "<^"
		}
	}

	for i := 0; i < len(lines); i++ {
		line := common.ReplaceStr(lines[i], []string{"p=", "v="}, "")
		line = common.ReplaceStr(line, []string{" "}, ",")
		parts := strings.Split(line, ",")
		p := position{common.ParseStrToInt(parts[0]), common.ParseStrToInt(parts[1])}
		v := velocity{common.ParseStrToInt(parts[2]), common.ParseStrToInt(parts[3])}
		d := determineDirection(v)
		robots[i] = robot{i, p, v, d, p}
	}

	return robots
}

func createSpace(tall, wide int) space {
	s := space{id: 0}

	s.tiles = make([][]tile, tall)

	for i := range s.tiles {
		s.tiles[i] = make([]tile, wide)
	}

	return s
}

func splitQuadrants(s space, tall, wide int) []quadrant {
	vertMid := tall / 2
	horMid := wide / 2
	quadrants := make([]quadrant, 4)

	// Allocate tiles for each quadrant
	extractSubslice := func(tiles [][]tile, startRow, endRow, startCol, endCol int) [][]tile {
		sub := make([][]tile, endRow-startRow)
		for i := startRow; i < endRow; i++ {
			sub[i-startRow] = tiles[i][startCol:endCol]
		}
		return sub
	}

	countRobots := func(q *quadrant) {
		rows := len(q.tiles)
		cols := len(q.tiles[0])

		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				tile := q.tiles[i][j]
				if len(tile.robots) > 0 {
					q.robotCount += len(tile.robots)
				}
			}
		}
	}

	// Extract tiles for each quadrant
	quadrants[0] = quadrant{id: 1, tiles: extractSubslice(s.tiles, 0, vertMid, 0, horMid)}           // Top-left
	quadrants[1] = quadrant{id: 2, tiles: extractSubslice(s.tiles, 0, vertMid, horMid+1, wide)}      // Top-right
	quadrants[2] = quadrant{id: 3, tiles: extractSubslice(s.tiles, vertMid+1, tall, 0, horMid)}      // Bottom-left
	quadrants[3] = quadrant{id: 4, tiles: extractSubslice(s.tiles, vertMid+1, tall, horMid+1, wide)} // Bottom-right

	for i := 0; i < len(quadrants); i++ {
		countRobots(&quadrants[i])
	}

	return quadrants
}

func moveRobot(robot robot, tall, wide int) position {
	x := robot.lastPosition.x + robot.velocity.x
	y := robot.lastPosition.y + robot.velocity.y

	if x < 0 {
		x = wide + x
	}
	if x >= wide {
		x = x - wide
	}

	if y < 0 {
		y = tall + y
	}
	if y >= tall {
		y = y - tall
	}

	return position{x, y}
}

func teleportRobots(space *space, robots []robot, seconds int, printSteps bool) {
	tall := len((*space).tiles)
	wide := len((*space).tiles[0])
	checkRobotsSamePlace := func(tiles [][]tile) bool {
		for i := 0; i < tall; i++ {
			for j := 0; j < wide; j++ {
				tile := tiles[i][j]
				if len(tile.robots) > 1 {
					return true
				}
			}
		}
		return false
	}

	for i := 0; i < seconds; i++ {
		for r := 0; r < len(robots); r++ {
			robot := &robots[r]
			previousTile := robot.lastPosition
			nextPosition := moveRobot(*robot, tall, wide)
			robot.lastPosition = nextPosition
			removeRobot(&(*space).tiles[previousTile.y][previousTile.x], *robot)
			addRobot(&(*space).tiles[nextPosition.y][nextPosition.x], *robot)

		}

		if !checkRobotsSamePlace((*space).tiles) {
			fmt.Println("Easter egg at", i+1, " seconds")
		}
	}
}

func printTiles(tiles [][]tile) string {
	var builder strings.Builder
	for i := 0; i < len(tiles); i++ {
		for j := 0; j < len(tiles[i]); j++ {
			robots := len(tiles[i][j].robots)
			if robots > 0 {
				builder.WriteString(strconv.Itoa(robots))
			} else {
				builder.WriteString(".")
			}
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func Puzzle1() int {
	lines := common.ReadFileByLines("./inputs/day14.txt")
	tall := 103
	wide := 101

	robots := parseRobots(lines)

	space := createSpace(tall, wide)

	//fmt.Println("Initial state")
	//printSpaces(space)

	teleportRobots(&space, robots, 10000, true)

	fmt.Println("Last state")
	printTiles(space.tiles)

	quadrants := splitQuadrants(space, tall, wide)

	factor := 1
	for i := 0; i < len(quadrants); i++ {
		quadrant := quadrants[i]
		//fmt.Println("quadrant ", quadrant.id)
		//printTiles(quadrant.tiles)
		if quadrant.robotCount == 0 {
			continue
		}

		factor *= quadrant.robotCount
	}

	return factor
}
