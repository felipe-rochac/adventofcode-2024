package day3

import (
	"adventofcode-2024/common"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func parseFile(content string) [][]int {
	pattern := `mul\(\d+,\d+\)`
	muls := make([][]int, 0)
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(content, -1)

	for _, m := range matches {

		data := strings.Replace(m, "mul(", "", 1)
		data = strings.Replace(data, ")", "", 1)
		elements := strings.Split(data, ",")

		num1, err := strconv.Atoi(elements[0])
		if err != nil {
			panic(fmt.Sprintf("Invalid num %s", num1))
		}

		num2, err := strconv.Atoi(elements[1])
		if err != nil {
			panic(fmt.Sprintf("Invalid num %s", num2))
		}

		muls = append(muls, []int{num1, num2})
	}
	return muls
}

func sumMuls(muls [][]int) int {
	sum := 0
	for i := 0; i < len(muls); i++ {
		sum += muls[i][0] * muls[i][1]
	}
	return sum
}

func Puzzle1() int {
	content := common.ReadFileText("./inputs/day3.txt")

	muls := parseFile(content)

	sum := sumMuls(muls)

	return sum
}
