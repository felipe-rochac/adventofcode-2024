package day14

import (
	"adventofcode-2024/common"
	"fmt"
	"strings"
)

var cache map[string][]string

func findMatches(pattern string, towels []string) []string {
	matches := make([]string, 0)
	for _, t := range towels {
		if strings.HasPrefix(pattern, t) {
			matches = append(matches, t)
		}
	}

	return matches
}

func findArrangements(pattern string, towels []string) []string {
	value, exists := cache[pattern]

	if exists {
		return value
	}

	if len(pattern) == 0 {
		return []string{}
	}

	matches := findMatches(pattern, towels)

	var arrangements []string
	for _, m := range matches {
		arr := findArrangements(pattern[len(m):], towels)
		if len(arr) == 0 {
			arrangements = append(arrangements, m)
			continue
		}

		cache[m] = arr

		for _, a := range arr {
			arrangements = append(arrangements, m+a)
		}
	}

	return arrangements
}

func seekPossibleDesignsV2(towels, patterns []string) int {
	count := 0
	for _, p := range patterns {
		arrangement := findArrangements(p, towels)
		total := common.CountOccurrences(arrangement, p)
		fmt.Printf("For %s it was found %d combinations\n", p, total)
		count += total
	}

	return count
}

func Puzzle2() int {
	lines := common.ReadFileByLines("./day19/input.txt")
	cache = make(map[string][]string)

	towels, patterns := parseInput(lines)

	count := seekPossibleDesignsV2(towels, patterns)

	return count
}
