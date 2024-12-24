package day3

import (
	"adventofcode-2024/common"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func parseFileV2(content string) [][]int {
	pattern := `mul\(\d+,\d+\)|do\(\)|don\'t\(\)`
	muls := make([][]int, 0)
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(content, -1)
	ignoreNext := false
	extractData := func(content string) (int, int) {
		data := strings.Replace(content, "mul(", "", 1)
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

		return num1, num2
	}

	for _, m := range matches {

		if m == "don't()" {
			ignoreNext = true
			continue
		}

		if m == "do()" {
			ignoreNext = false
			continue
		}

		if ignoreNext {
			continue
		}

		num1, num2 := extractData(m)
		muls = append(muls, []int{num1, num2})
	}
	return muls
}

func Puzzle2() int {
	content := common.ReadFileText("./inputs/day3.txt")

	muls := parseFileV2(content)

	sum := sumMuls(muls)

	return sum
}
