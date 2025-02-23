package common

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func ParseInt(r rune) int {
	i, err := strconv.Atoi(string(r))

	if err != nil {
		panic(fmt.Sprintf("Invalid int %c\n", r))
	}

	return i
}

func SplitStr(str string, sep string) []string {
	return strings.FieldsFunc(str, func(r rune) bool {
		return strings.ContainsRune(sep, r)
	})
}

func ReplaceStr(str string, old []string, new string) string {
	for _, s := range old {
		str = strings.Replace(str, s, new, 1)
	}
	return str
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func ParseStrToInt(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		panic(fmt.Sprintf("Invalid int %s\n", s))
	}

	return i
}

func ReadFileAsGrid(fileName string) [][]rune {
	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	defer file.Close()
	grid := make([][]rune, 0)

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		gridLine := make([]rune, 0)
		for _, char := range line {
			gridLine = append(gridLine, char)
		}
		grid = append(grid, gridLine)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return grid
}

func ReplaceCharAtIndex(str string, index int, newChar rune) string {
	runes := []rune(str)
	// Replace the character at the specified index
	runes[index] = newChar
	// Convert the rune slice back to a string
	return string(runes)
}

func ReadFileByLines(fileName string) []string {
	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	defer file.Close()
	var lines []string

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

func ReadFileText(fileName string) string {
	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	defer file.Close()
	var content string

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return content
}

func SliceContains[T comparable](slice []T, element T) bool {
	for _, value := range slice {
		if value == element {
			return true
		}
	}
	return false
}

func CompareIntSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func WriteToFile(fileName string, content string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(content)

	if err != nil {
		panic(err)
	}
}

func AppendToFile(filename, content string) error {
	// Open the file in append mode, create it if it doesn't exist
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the content to the file
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

func CopyMap[K comparable, V any](original map[K]V) map[K]V {
	copy := make(map[K]V, len(original))
	for key, value := range original {
		copy[key] = value
	}
	return copy
}

func CopyGrid[V any](original [][]V) [][]V {
	temp := make([][]V, len(original))
	for i, row := range original {
		// Create a new slice for each inner slice
		temp[i] = make([]V, len(row))
		copy(temp[i], row) // Copy elements from the original inner slice
	}
	return temp
}

func CopyArray[V any](original []V) []V {
	temp := make([]V, len(original))
	for i, row := range original {
		temp[i] = row
	}
	return temp
}

func CountOccurrences[T comparable](arr []T, value T) int {
	count := 0
	for _, v := range arr {
		if v == value {
			count++
		}
	}
	return count
}

func Combine[T comparable](array []T, x int) [][]T {
	if x == 0 {
		return [][]T{{}} // Return an empty combination
	}
	if len(array) < x {
		return [][]T{} // Not enough elements to create a combination
	}

	// Exclude the first element and combine the rest
	exclude := Combine(array[1:], x)

	// Include the first element
	include := Combine(array[1:], x-1)
	for i := range include {
		include[i] = append([]T{array[0]}, include[i]...)
	}

	// Return both included and excluded combinations
	return append(exclude, include...)
}

func RemoveAlpha(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) {
			return -1 // Removes the character
		}
		return r
	}, s)
}
