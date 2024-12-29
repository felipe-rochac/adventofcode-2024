package day11

import (
	"adventofcode-2025/common"
	"fmt"
	"strconv"
	"sync"
)

type intTuple struct {
	a, b int
}

var cache sync.Map
var splceCache sync.Map

func cachedMultiply(value, multiplier, module int) int {
	if result, exists := cache.Load(value); exists {
		return result.(int)
	}
	result := (value * multiplier) % module
	cache.Store(value, result)
	return result
}

func cachedSplit(value, module int) intTuple {
	if result, exists := splceCache.Load(value); exists {
		return result.(intTuple)
	}

	valueStr := strconv.Itoa(value)
	half := len(valueStr) / 2
	part1 := common.ParseStrToInt(valueStr[:half]) % module
	part2 := common.ParseStrToInt(valueStr[half:]) % module

	tuple := intTuple{part1, part2}
	splceCache.Store(value, tuple)
	return tuple
}

func stonesToMap(stones []stone) map[int]int {
	_map := make(map[int]int)
	for _, stone := range stones {
		_map[stone.value]++
	}
	return _map
}

func blinkV2(stonesMap map[int]int, module int) map[int]int {
	newStonesMap := make(map[int]int)
	for key, s := range stonesMap {
		value := key
		valueStr := strconv.Itoa(value)

		if value == 0 {
			newStonesMap[1] += s
		} else if len(valueStr)%2 == 0 {
			tuple := cachedSplit(value, module)
			newStonesMap[tuple.a] += s
			newStonesMap[tuple.b] += s
		} else {
			newValue := cachedMultiply(value, 2024, module)
			newStonesMap[newValue] += s
		}
	}
	return newStonesMap
}

func Puzzle2() int {
	lines := common.ReadFileByLines(fileName)
	stones := parseStones(lines)
	values := make([]int, len(stones))
	for i, s := range stones {
		values[i] = s.value
	}

	module := lcmForList(stones)
	generations := 75

	_map := stonesToMap(stones)

	for i := 0; i < generations; i++ {
		_map = blinkV2(_map, module)
		if (i+1)%10 == 0 {
			fmt.Println("Blink (", i+1, ") has ", len(_map), " stones")
		}
	}

	sum := 0
	for _, value := range _map {
		sum += value
	}

	return sum
}
