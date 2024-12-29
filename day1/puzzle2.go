package day1

import (
	"adventofcode-2025/common"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func parseFiles2(lines []string) (*[]int, *map[int]int) {
	var list []int
	rightDict := make(map[int]int)

	for _, value := range lines {
		elements := strings.Split(value, "   ")

		key, err := strconv.Atoi(elements[0])

		list = append(list, key)

		key, err = strconv.Atoi(elements[1])

		if err != nil {
			panic(fmt.Sprintf("Invalid integer %s", elements[0]))
		}

		rightDict[key]++
	}

	return &list, &rightDict
}

func calculateAppearances(list *[]int, rightDict *map[int]int) *[]int {
	var appearances []int
	for i := 0; i < len(*list); i++ {
		element := (*list)[i]
		value, exists := (*rightDict)[element]

		if exists {
			appearances = append(appearances, value*element)
		}
	}
	return &appearances
}

func sumAppearances(list *[]int) int {
	var sum = 0
	for i := 0; i < len(*list); i++ {
		sum += (*list)[i]
	}
	return sum
}

func Puzzle2() {
	file := "./inputs/day1.txt"
	lines := common.ReadFileByLines(file)

	if len(lines) == 0 {
		panic(fmt.Sprintf("No line found on file %s", file))
	}

	list, rightDict := parseFiles2(lines)

	appearances := calculateAppearances(list, rightDict)

	sum := sumAppearances(appearances)

	log.Println(fmt.Sprintf("Your total distance is %d", sum))
}
