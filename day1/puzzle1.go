package day1

import (
	"adventofcode-2025/common"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func findMin(dict *map[int]int) int {
	currentMin := math.MaxInt32
	for key, _ := range *dict {
		if key < currentMin {
			currentMin = key
		}
	}

	// subtract 1 from current key
	(*dict)[currentMin]--

	// if no more values, then remove key
	if (*dict)[currentMin] == 0 {
		delete(*dict, currentMin)
	}

	return currentMin
}

func parseFiles(lines []string) (*map[int]int, *map[int]int) {
	leftDict := make(map[int]int)
	rightDict := make(map[int]int)

	for _, value := range lines {
		elements := strings.Split(value, "   ")

		key, err := strconv.Atoi(elements[0])

		if err != nil {
			panic(fmt.Sprintf("Invalid integer %s", elements[0]))
		}

		leftDict[key]++

		key, err = strconv.Atoi(elements[1])

		if err != nil {
			panic(fmt.Sprintf("Invalid integer %s", elements[0]))
		}

		rightDict[key]++
	}

	return &leftDict, &rightDict
}

func calculateDistances(leftDict *map[int]int, rightDict *map[int]int) []int {
	var distances []int
	for len(*leftDict) > 0 {
		leftMin := findMin(leftDict)
		rightMin := findMin(rightDict)

		distances = append(distances, int(math.Abs(float64(rightMin-leftMin))))
	}

	return distances
}

func sumDistances(distances []int) int {
	sum := 0
	for i := 0; i < len(distances); i++ {
		sum += distances[i]
	}

	return sum
}

func Puzzle1() {
	file := "./inputs/day1.txt"
	lines := common.ReadFileByLines(file)

	if len(lines) == 0 {
		panic(fmt.Sprintf("No line found on file %s", file))
	}

	leftDict, rightDict := parseFiles(lines)

	distances := calculateDistances(leftDict, rightDict)

	sum := sumDistances(distances)

	log.Println(fmt.Sprintf("Your total distance is %d", sum))
}
