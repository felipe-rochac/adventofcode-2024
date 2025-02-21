package day21

import (
	"adventofcode-2024/common"
)

func Puzzle2() int {
	// Generate graphs
	codes := common.ReadFileByLines("./day21/input.txt")
	numpadGraph := createGraph(numpad, [2]int{3, 0})
	dirpadGraph := createGraph(dirpad, [2]int{0, 0})
	conversionsCache = make(map[string]string)

	complexity := 0
	// Example usage of convert
	for _, sequence := range codes {
		code := common.ParseStrToInt(common.RemoveAlpha(sequence))
		numpadSequence := convert(sequence, numpadGraph)
		seq := numpadSequence

		for i := 0; i < 25; i++ {
			seq = convert(seq, dirpadGraph)
		}

		complexity += len(seq) * code
	}
	return complexity
}
