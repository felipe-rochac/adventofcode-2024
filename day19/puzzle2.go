package day14

import (
	"adventofcode-2024/common"
	"fmt"
	"strings"
)

var cache map[string]int

// Finds towels that match the prefix of the pattern
func findMatches(pattern string, towels []string) []string {
	matches := make([]string, 0)
	for _, t := range towels {
		if strings.HasPrefix(pattern, t) {
			matches = append(matches, t)
		}
	}
	return matches
}

// Finds the count of arrangements for a given pattern
func countArrangements(pattern string, towels []string) int {
	// Check cache
	if value, exists := cache[pattern]; exists {
		return value
	}

	// Base case: Empty pattern
	if len(pattern) == 0 {
		return 1 // There's exactly one way to arrange an empty pattern (do nothing)
	}

	// Find matching towels
	matches := findMatches(pattern, towels)
	count := 0

	// Count arrangements recursively for each match
	for _, m := range matches {
		count += countArrangements(pattern[len(m):], towels)
	}

	// Store in cache
	cache[pattern] = count
	return count
}

// Seeks possible designs for multiple patterns
func seekPossibleDesignsV2(towels, patterns []string) int {
	totalCount := 0
	for _, p := range patterns {
		count := countArrangements(p, towels)
		fmt.Printf("For pattern '%s', found %d combinations\n", p, count)
		totalCount += count
	}
	return totalCount
}

func Puzzle2() int {
	lines := common.ReadFileByLines("./day19/input.txt")
	cache = make(map[string]int)

	towels, patterns := parseInput(lines)

	count := seekPossibleDesignsV2(towels, patterns)

	return count
}
