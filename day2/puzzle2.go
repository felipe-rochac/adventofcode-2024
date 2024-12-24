package day2

import (
	"adventofcode-2024/common"
)

func sign(n int) int {
	if n > 0 {
		return 1
	} else if n < 0 {
		return -1
	}
	return 0
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func isSafeReport(values []int) bool {
	direction := sign(values[0] - values[1])

	left, right := 0, 1
	for left < len(values)-1 && right < len(values) {
		rawDiff := values[left] - values[right]
		diff := abs(rawDiff)

		// Check conditions for safety
		if direction != sign(rawDiff) || diff == 0 || diff > 3 {
			return false
		}
		left++
		right++
	}
	return true
}

func countSafeReportsV2(reports [][]int) int {
	safeCount := 0
	for _, line := range reports {

		if len(line) <= 1 {
			continue
		}

		// Check if the line is safe
		if isSafeReport(line) {
			safeCount++
		} else {
			// Try removing each element and recheck
			for i := 0; i < len(line); i++ {
				// Create a copy of data without the i-th element
				temp := append([]int{}, line[:i]...)
				temp = append(temp, line[i+1:]...)

				if isSafeReport(temp) {
					safeCount++
					break
				}
			}
		}
	}
	return safeCount
}

func Puzzle2() int {
	lines := common.ReadFileByLines("./inputs/day2.txt")

	reports := parseFile(lines)

	count := countSafeReportsV2(reports)

	// > 594 && < 726 && !673
	return count
}
