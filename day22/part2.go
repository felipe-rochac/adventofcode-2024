package main

import (
	"adventofcode-2024/common"
	day22 "adventofcode-2024/day22/common"
	"fmt"
)

// Gets the sequence and max from an array using start post
func getdMaxAndSequenceFromArrayStartPost(entries []int, startPos int) ([]int, int) {
	seq := make([]int, 0)
	max := 0
	for j := 0; j < 4; j++ {
		cur := entries[startPos-3+j] % 10
		prev := entries[startPos-4+j] % 10
		r := cur - prev
		if cur > max {
			max = cur
		}
		seq = append(seq, r)
	}
	return seq, max
}

// Find max and sequence of differences in an aray
func findSequenceAndMaxFromArray(entries []int) (int, []int) {
	max := 0
	index := 0

	for i, v := range entries {
		if i < 4 {
			continue
		}

		cur := v % 10

		if cur > max {
			max = cur
			index = i
		}
	}

	seq := make([]int, 4)
	seq[0] = entries[index-3]%10 - entries[index-4]%10
	seq[1] = entries[index-2]%10 - entries[index-3]%10
	seq[2] = entries[index-1]%10 - entries[index-2]%10
	seq[3] = entries[index]%10 - entries[index-1]%10

	return max, seq
}

// Determines if array contains sequence and returns it max
func findMaxForArrayFromSequence(entries, seq []int) (bool, int) {
	for i := 3; i < len(entries); i++ {
		seq2, max := getdMaxAndSequenceFromArrayStartPost(entries, i)

		if common.CompareIntSlices(seq, seq2) {
			return true, max
		}
	}

	return false, 0
}

func main() {
	buyers := []int{1, 2, 3, 2024}
	// lines := common.ReadFileByLines("./input.txt")
	// buyers := parseBuyers(lines)
	secrets := make([]int, 0)
	secret := buyers[0]
	// generate 2000 secrets for buyers 1 and retrieve max and sequence
	secrets = append(secrets, buyers[0])
	generations := 1999

	for range generations {
		secret = day22.Evolve(secret)
		secrets = append(secrets, secret)
	}
	max, seq := findSequenceAndMaxFromArray(secrets)
	total := max

	// once sequence is stablished, tries to find it for other buyers
	for i := 1; i < len(buyers); i++ {
		secrets = make([]int, 0)
		secrets = append(secrets, buyers[i])
		for range generations {
			secret := day22.Evolve(secret)
			secrets = append(secrets, secret)
		}

		found, max := findMaxForArrayFromSequence(secrets, seq)

		if found {
			total += max
		}
	}

	fmt.Print(len(seq))
	fmt.Print(max)
}
