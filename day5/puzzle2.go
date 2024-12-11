package day5

import "adventofcode-2025/common"

func checkCorrectlyUpdatesAndFix(rules []Rule, instructions *[]Instruction) int {
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

	checkInterdictionList := func(instruction *Instruction) bool {
		var check func(index int) bool
		check = func(index int) bool {
			// Base case: If we've processed all elements, return false
			if index >= len(instruction.updatesOrder)-1 {
				return false
			}
			// Check the interdiction rule for the current element and the next one
			if findInterditionRule(instruction.updatesOrder[index], instruction.updatesOrder[index+1]) {
				instruction.correct = false
				temp := instruction.updatesOrder[index+1]
				instruction.updatesOrder[index+1] = instruction.updatesOrder[index]
				instruction.updatesOrder[index] = temp
				// start check over
				return check(0)
			}
			// Recursive case: Move to the next element
			return check(index + 1)
		}
		// Start recursion from the first index
		return check(0)
	}

	for _, instruction := range *instructions {
		checkInterdictionList(&instruction)

		if !instruction.correct {
			middleIndex := len(instruction.updatesOrder) / 2
			sum += instruction.updatesOrder[middleIndex]
		}
	}

	return sum
}

func Puzzle2() int {
	content := common.ReadFileByLines("./inputs/day5.txt")

	rules, instructions := parseFile(content)

	sum := checkCorrectlyUpdatesAndFix(rules, &instructions)

	return sum
}
