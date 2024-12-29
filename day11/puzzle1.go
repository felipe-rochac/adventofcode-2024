package day11

import (
	"adventofcode-2025/common"
	"strconv"
	"strings"
)

const fileName = "./inputs/day11.txt"

type stone struct {
	value int
}

func parseStones(lines []string) (stones []stone) {
	if len(lines) > 1 {
		panic("too many lines")
	}

	elements := strings.Split(lines[0], " ")
	stones = make([]stone, len(elements))

	for i, element := range elements {
		stones[i] = stone{common.ParseStrToInt(element)}
	}

	return stones
}

func blink(stones *[]stone, module int) {
	newStones := make([]stone, 0, len(*stones)*2) // Preallocate to minimize reallocations

	for _, s := range *stones {
		value := s.value
		valueStr := strconv.Itoa(value)

		if value == 0 {
			// Case 1: If the value is 0, replace it with 1
			newStones = append(newStones, stone{value: 1})
		} else if len(valueStr)%2 == 0 {
			// Case 2: If the length of the number is even, split it into two parts
			half := len(valueStr) / 2
			part1 := common.ParseStrToInt(valueStr[:half]) % module
			part2 := common.ParseStrToInt(valueStr[half:]) % module
			newStones = append(newStones, stone{value: part1}, stone{value: part2})
		} else {
			// Case 3: If the length of the number is odd, apply the transformation
			newValue := value * 2024 % module
			newStones = append(newStones, stone{value: newValue})
		}
	}

	*stones = newStones
}

func lcmForList(stones []stone) int {
	lcm := 1
	for _, stone := range stones {
		if stone.value == 0 {
			continue
		}

		lcm *= stone.value
	}

	return lcm
}

func Puzzle1() int {
	lines := common.ReadFileByLines(fileName)

	stones := parseStones(lines)
	modulo := lcmForList(stones)

	for i := 0; i < 25; i++ {
		blink(&stones, modulo)
		//fmt.Println("Blink (", i+1, ") has ", len(stones), " stones")
	}

	return len(stones)
}
