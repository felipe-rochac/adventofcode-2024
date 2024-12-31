package day14

import (
	"adventofcode-2024/common"
	"sort"
	"strings"
)

func parseInput(lines []string) (towels, patterns []string) {
	parts := strings.Split(lines[0], ",")
	towels = make([]string, len(parts))

	for i, p := range parts {
		towels[i] = strings.TrimSpace(p)
	}

	patterns = make([]string, len(lines[2:]))
	for i, p := range lines[2:] {
		patterns[i] = strings.TrimSpace(p)
	}

	return towels, patterns
}

func tryBuildPattern(pattern string, towels []string) bool {
	for _, t := range towels {
		// If does not match prefix, continue
		if !strings.HasPrefix(pattern, t) {
			continue
		}

		l := len(t)
		remaining := len(pattern[l:])

		if remaining == 0 || tryBuildPattern(pattern[l:], towels) {
			return true
		}
	}
	return false
}

func seekPossibleDesigns(towels, patterns []string) int {

	// Sorting towels for longest combination to the shortest
	sort.Slice(towels, func(i, j int) bool {
		return len(towels[i]) > len(towels[j])
	})

	count := 0
	for _, p := range patterns {
		if tryBuildPattern(p, towels) {
			count++
		}
	}

	return count
}

func Puzzle1() int {
	lines := common.ReadFileByLines("./day19/input.txt")

	towels, patterns := parseInput(lines)

	count := seekPossibleDesigns(towels, patterns)

	return count
}
