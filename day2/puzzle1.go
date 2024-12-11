package day2

import (
	"adventofcode-2025/common"
	"math"
	"strconv"
	"strings"
)

const increasing string = "increasing"
const decreasing string = "decreasing"
const safe string = "safe"
const unsafe string = "unsafe"

func parseFile(lines []string) [][]int {
	reports := make([][]int, len(lines))
	for index, line := range lines {
		split := strings.Split(line, " ")
		cols := make([]int, len(split))

		for i, s := range split {
			cols[i], _ = strconv.Atoi(s)
		}

		reports[index] = cols
	}

	return reports
}

func countSafeReports(reports [][]int) int {
	safeReports := 0

	isFaulty := func(direction string, element int, nextElement int) bool {
		difference := math.Abs(float64(element - nextElement))
		if difference > 3 || difference == 0 {
			return true
		} else if direction == decreasing && nextElement > element {
			return true
		} else if direction == increasing && nextElement < element {
			return true
		}
		return false
	}

	for _, report := range reports {
		firstElement := report[0]
		element := report[1]
		direction := increasing
		reportState := safe

		if element < firstElement {
			direction = decreasing
		}

		diff := math.Abs(float64(element - firstElement))
		if diff == 0 || diff > 3 {
			continue
		}

		for i := 2; i < len(report); i++ {
			nextElement := report[i]

			if isFaulty(direction, element, nextElement) {
				reportState = unsafe
				break
			}

			element = nextElement
		}

		if reportState == safe {
			safeReports++
		}
	}

	return safeReports
}

func Puzzle1() int {
	lines := common.ReadFileByLines("./inputs/day2.txt")

	reports := parseFile(lines)

	count := countSafeReports(reports)

	return count
}
