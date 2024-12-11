package day7

import "adventofcode-2025/common"

func Puzzle2() int {
	lines := common.ReadFileByLines("./inputs/day7.txt")

	equations := parseContent(lines)

	sum := calculateCalibrationResults(equations, []rune{'+', '*', '|'})

	return sum

}
