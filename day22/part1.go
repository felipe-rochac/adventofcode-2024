package main

import (
	"adventofcode-2024/common"
	day22 "adventofcode-2024/day22/common"
	"fmt"
)

func part1() {
	// buyers := []int{1, 10, 100, 2024}
	lines := common.ReadFileByLines("./input.txt")
	buyers := day22.ParseBuyers(lines)

	total := 0

	for _, b := range buyers {
		secret := b
		for i := 0; i < 2000; i++ {
			secret = day22.Evolve(secret)
		}
		fmt.Println(fmt.Sprintf("%d: %d", b, secret))
		total += secret
	}

	fmt.Println(fmt.Sprintf("Total is %d", total))
}
