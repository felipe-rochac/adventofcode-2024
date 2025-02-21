package day22

import "adventofcode-2024/common"

func mix(number, secret int) int {
	return number ^ secret
}

func prune(secret int) int {
	return secret % 16777216
}

func evolvePart1(secret int) int {
	s := secret * 64
	s = mix(s, secret)
	return prune(s)
}

func evolvePart2(secret int) int {
	s := secret / 32
	s = mix(s, secret)
	return prune(s)
}

func evolvePart3(secret int) int {
	s := secret * 2048
	s = mix(s, secret)
	return prune(s)
}

func Evolve(secret int) int {
	s := evolvePart1(secret)
	s = evolvePart2(s)
	return evolvePart3(s)
}

func ParseBuyers(lines []string) []int {
	buyers := make([]int, len(lines))
	for i, s := range lines {
		buyers[i] = common.ParseStrToInt(s)
	}

	return buyers
}
