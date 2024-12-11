package main

import (
	day "adventofcode-2025/day2"
	"encoding/base64"
	"fmt"
	"log"
)

func decodeBase64() {
	decimalValues := []byte{
		50,
		57,
		182,
		144,
		98,
		248,
		128,
		71,
		135,
		76,
		41,
		230,
		158,
		122,
		219,
		67,
	}

	// Convert the byte array to a Base64 string
	base64String := base64.URLEncoding.EncodeToString(decimalValues)
	fmt.Println("Base64 String:", base64String)

	// Decode the Base64 string
	decodedBytes, err := base64.URLEncoding.DecodeString(base64String)
	if err != nil {
		fmt.Println("Error decoding Base64:", err)
		return
	}

	// Convert decoded bytes to a string
	decodedString := string(decodedBytes)

	// Print the result
	fmt.Println("Decoded string:", decodedString)
}

func main() {
	//decodeBase64()

	sum := day.Puzzle1()

	log.Println(fmt.Sprintf("Puzzle 1 sum was %d", sum))

	sum = day.Puzzle2()

	log.Println(fmt.Sprintf("Puzzle 2 sum was %d", sum))
}
