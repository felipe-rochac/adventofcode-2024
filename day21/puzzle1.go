package day21

import (
	"adventofcode-2024/common"
	"strings"
)

var numpad = map[string][2]int{
	"7": {0, 0}, "8": {0, 1}, "9": {0, 2},
	"4": {1, 0}, "5": {1, 1}, "6": {1, 2},
	"1": {2, 0}, "2": {2, 1}, "3": {2, 2},
	"0": {3, 1}, "A": {3, 2},
}

var dirpad = map[string][2]int{
	"^": {0, 1}, "A": {0, 2},
	"<": {1, 0}, "v": {1, 1}, ">": {1, 2},
}

func repeat(r rune, count int) string {
	if count <= 0 {
		return ""
	}
	return strings.Repeat(string(r), count)
}

// CreateGraph function
func createGraph(keypad map[string][2]int, invalidCoords [2]int) map[[2]string]string {
	graph := make(map[[2]string]string)
	for a, coords1 := range keypad {
		x1, y1 := coords1[0], coords1[1]
		for b, coords2 := range keypad {
			x2, y2 := coords2[0], coords2[1]
			path := ""
			path += repeat('<', y1-y2)
			path += repeat('v', x2-x1)
			path += repeat('^', x1-x2)
			path += repeat('>', y2-y1)
			if invalidCoords == [2]int{x1, y2} || invalidCoords == [2]int{x2, y1} {
				path = reverse(path)
			}
			graph[[2]string{a, b}] = path + "A"
		}
	}

	return graph
}

// Reverse a string
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Convert function
func convert(sequence string, graph map[[2]string]string) string {
	value, exists := conversionsCache[sequence]

	if exists {
		return value
	}

	conversion := ""
	prev := "A"
	for _, char := range sequence {
		conversion += graph[[2]string{prev, string(char)}]
		prev = string(char)
	}
	conversionsCache[sequence] = conversion

	if len(conversion) > cacheMaxLenght {
		cacheMaxLenght = len(conversion)
	}

	return conversion
}

func hasOnCache(sequence string) (string, bool) {
	value, exists := conversionsCache[sequence]

	if exists {
		return value, true
	}

	if cacheMaxLenght == 0 {
		return "", false
	}

	return "", false
}

var conversionsCache map[string]string
var cacheMaxLenght int

func Puzzle1() int {
	// Generate graphs
	codes := common.ReadFileByLines("./day21/input.txt")
	numpadGraph := createGraph(numpad, [2]int{3, 0})
	dirpadGraph := createGraph(dirpad, [2]int{0, 0})
	conversionsCache = make(map[string]string)
	cacheMaxLenght = 0

	complexity := 0
	// Example usage of convert
	for _, sequence := range codes {
		code := common.ParseStrToInt(common.RemoveAlpha(sequence))
		numpadSequence := convert(sequence, numpadGraph)

		robot1Sequence := convert(numpadSequence, dirpadGraph)

		robot2Sequence := convert(robot1Sequence, dirpadGraph)

		complexity += len(robot2Sequence) * code
	}
	return complexity
}
