package day5

import (
	"adventofcode-2024/common"
	"strconv"
	"strings"
)

type Rule struct {
	left  int
	right int
}

type Instruction struct {
	updatesOrder []int
	correct      bool
}

func parseFile(lines []string) ([]Rule, []Instruction) {
	rules := make([]Rule, 0)
	count := len(lines)
	index := 0

	for i := 0; i < count; i++ {
		line := lines[i]
		if line == "" {
			index = i + 1
			break
		}

		elements := strings.Split(line, "|")
		e1, err := strconv.Atoi(elements[0])

		if err != nil {
			panic(err)
		}

		e2, err := strconv.Atoi(elements[1])

		if err != nil {
			panic(err)
		}

		rule := Rule{left: e1, right: e2}
		rules = append(rules, rule)
	}

	instructions := make([]Instruction, 0)
	for i := index; i < count; i++ {
		line := lines[i]

		elements := strings.Split(line, ",")
		order := make([]int, 0)

		for _, element := range elements {
			e, err := strconv.Atoi(element)

			if err != nil {
				panic(err)
			}

			order = append(order, e)
		}
		instructions = append(instructions, Instruction{updatesOrder: order, correct: true})
	}

	return rules, instructions
}

func checkCorrectlyUpdates(rules []Rule, instructions *[]Instruction) int {
	sum := 0

	findInterditionRule := func(input int, nextInput int) bool {
		for _, rule := range rules {
			if rule.right != input {
				continue
			}

			if rule.left == nextInput {
				return true
			}
		}
		return false
	}

	checkInterdictionList := func(instructions []int) bool {
		var check func(index int) bool
		check = func(index int) bool {
			// Base case: If we've processed all elements, return false
			if index >= len(instructions)-1 {
				return false
			}
			// Check the interdiction rule for the current element and the next one
			if findInterditionRule(instructions[index], instructions[index+1]) {
				return true
			}
			// Recursive case: Move to the next element
			return check(index + 1)
		}
		// Start recursion from the first index
		return check(0)
	}

	for _, instruction := range *instructions {
		invalidInstruction := checkInterdictionList(instruction.updatesOrder)

		if !invalidInstruction {
			middleIndex := len(instruction.updatesOrder) / 2
			sum += instruction.updatesOrder[middleIndex]
		} else {
			instruction.correct = false
		}
	}

	return sum
}

func Puzzle1() int {
	content := common.ReadFileByLines("./inputs/day5.txt")

	rules, instructions := parseFile(content)

	sum := checkCorrectlyUpdates(rules, &instructions)

	return sum
}
