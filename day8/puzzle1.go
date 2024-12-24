package day8

import (
	"adventofcode-2024/common"
)

type Coordinate struct {
	Row, Col int
}

var antennas []struct {
	Freq  rune
	Coord Coordinate
}

var antennaToCoordinates map[rune][]Coordinate
var rows, cols int

func (c Coordinate) Add(other Coordinate) Coordinate {
	return Coordinate{c.Row + other.Row, c.Col + other.Col}
}

func (c Coordinate) Sub(other Coordinate) Coordinate {
	return Coordinate{c.Row - other.Row, c.Col - other.Col}
}

func IsInGrid(coord Coordinate, rowLen, colLen int) bool {
	return coord.Row >= 0 && coord.Row < rowLen && coord.Col >= 0 && coord.Col < colLen
}

func Puzzle1() int {
	input := common.ReadFileAsGrid("./inputs/day8.txt")

	rows = len(input)
	cols = len(input[0])
	uniqueAntinodes := make(map[Coordinate]struct{})
	antennaToCoordinates = make(map[rune][]Coordinate)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if input[i][j] != '.' {
				antennas = append(antennas, struct {
					Freq  rune
					Coord Coordinate
				}{Freq: input[i][j], Coord: Coordinate{i, j}})
			}
		}
	}

	for _, antenna := range antennas {
		antennaToCoordinates[antenna.Freq] = append(antennaToCoordinates[antenna.Freq], antenna.Coord)
	}

	for _, antenna := range antennas {
		sameFrequencyAntennas := antennaToCoordinates[antenna.Freq]
		for _, coord := range sameFrequencyAntennas {
			if coord.Row == antenna.Coord.Row && coord.Col == antenna.Coord.Col {
				continue
			}

			delta := antenna.Coord.Sub(coord)
			potentialAntinode := antenna.Coord.Add(delta)

			if IsInGrid(potentialAntinode, rows, cols) {
				uniqueAntinodes[potentialAntinode] = struct{}{}
			}
		}
	}

	return len(uniqueAntinodes)
}
