package day7

import (
	"adventofcode-2024/common"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Equation struct {
	result    int
	valuesStr string
	values    []int
}

func parseContent(lines []string) []Equation {
	equations := make([]Equation, 0)
	for _, line := range lines {
		elements := strings.Split(line, ":")

		result, err := strconv.Atoi(elements[0])

		if err != nil {
			panic(err)
		}

		valueStr := elements[1]
		elements = strings.Split(elements[1], " ")
		values := make([]int, 0)
		for _, element := range elements {
			if element == "" {
				continue
			}
			value, err := strconv.Atoi(element)

			if err != nil {
				panic(err)
			}

			values = append(values, value)
		}

		equations = append(equations, Equation{result: result, values: values, valuesStr: valueStr})
	}

	return equations
}

func calculateCalibrationResults(equations []Equation, operators []rune) int {
	calibration := 0

	var generatePermutations func(chars []rune, size int, current []rune, result *[]string)
	generatePermutations = func(chars []rune, size int, current []rune, result *[]string) {
		// Base case: if the current permutation's length equals the specified size, add to the result
		if len(current) == size {
			// Convert current permutation to string and add it to the result
			*result = append(*result, string(current))
			return
		}

		// Recursively generate permutations by adding each character from the array
		for i := 0; i < len(chars); i++ {
			// Append the current character to the current permutation
			current = append(current, chars[i])
			// Recurse with the next characters
			generatePermutations(chars, size, current, result)
			// Backtrack by removing the last character added
			current = current[:len(current)-1]
		}
	}

	for _, equation := range equations {
		var permutations []string
		var current []rune
		generatePermutations(operators, len(equation.values), current, &permutations)

		validCalibration := false
		for _, permutation := range permutations {
			validCalibration = false
			sum := 0
			for index, operator := range permutation {
				if operator == '+' {
					sum += equation.values[index]
				} else if operator == '|' {
					str := strconv.Itoa(sum) + strconv.Itoa(equation.values[index])
					value, err := strconv.Atoi(str)
					if err != nil {
						panic(err)
					}
					sum = value
				} else {
					if sum == 0 {
						sum = 1
					}
					sum *= equation.values[index]
				}

				if sum == equation.result {
					log.Println(fmt.Sprintf("Permutation %s matches test result %d with values (%s)", permutation, equation.result, equation.valuesStr))
					calibration += equation.result
					validCalibration = true
					break
				}
			}

			if validCalibration {
				break
			}
		}
	}

	return calibration
}

func Puzzle1() int {
	lines := common.ReadFileByLines("./inputs/day7.txt")

	equations := parseContent(lines)

	sum := calculateCalibrationResults(equations, []rune{'+', '*'})

	return sum
}
