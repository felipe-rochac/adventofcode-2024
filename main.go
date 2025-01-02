package main

import (
	day "adventofcode-2024/day20"
	"encoding/base64"
	"fmt"
	"log"
)

func decodeBase64(decimalValues []byte) {
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
