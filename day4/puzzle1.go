package day4

import (
	"bufio"
	"os"
	"strings"
)

type Character struct {
	Char  rune
	InUse bool
	Count int
}

func readFileCharactersByLines(fileName string) [][]Character {
	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([][]Character, 0)

	for scanner.Scan() {
		// Convert the line to a slice of runes
		line := scanner.Text()
		characters := make([]Character, 0)
		for _, r := range []rune(line) {
			characters = append(characters, Character{Char: r, InUse: false, Count: 0})
		}
		lines = append(lines, characters)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

func rotateMatrix90Clockwise(matrix [][]Character) [][]Character {
	n := len(matrix)
	// Create a new matrix to store the rotated version
	rotated := make([][]Character, n)
	for i := range rotated {
		rotated[i] = make([]Character, n)
	}

	// Perform the rotation
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			rotated[j][n-i-1] = matrix[i][j]
		}
	}

	return rotated
}

func flattenAllDiagonals(matrix [][]Character) ([][]Character, [][]Character) {
	rows := len(matrix)
	cols := len(matrix[0])

	// Store diagonals
	primaryDiagonals := [][]Character{}
	secondaryDiagonals := [][]Character{}

	// Collect all primary diagonals (top-left to bottom-right)
	for startCol := 0; startCol < cols; startCol++ {
		diagonal := []Character{}
		for i, j := 0, startCol; i < rows && j < cols; i, j = i+1, j+1 {
			diagonal = append(diagonal, matrix[i][j])
		}
		primaryDiagonals = append(primaryDiagonals, diagonal)
	}
	for startRow := 1; startRow < rows; startRow++ {
		diagonal := []Character{}
		for i, j := startRow, 0; i < rows && j < cols; i, j = i+1, j+1 {
			diagonal = append(diagonal, matrix[i][j])
		}
		primaryDiagonals = append(primaryDiagonals, diagonal)
	}

	// Collect all secondary diagonals (top-right to bottom-left)
	for startCol := 0; startCol < cols; startCol++ {
		diagonal := []Character{}
		for i, j := 0, startCol; i < rows && j >= 0; i, j = i+1, j-1 {
			diagonal = append(diagonal, matrix[i][j])
		}
		secondaryDiagonals = append(secondaryDiagonals, diagonal)
	}
	for startRow := 1; startRow < rows; startRow++ {
		diagonal := []Character{}
		for i, j := startRow, cols-1; i < rows && j >= 0; i, j = i+1, j-1 {
			diagonal = append(diagonal, matrix[i][j])
		}
		secondaryDiagonals = append(secondaryDiagonals, diagonal)
	}

	return primaryDiagonals, secondaryDiagonals
}

func matrixToString(matrix [][]Character) string {
	var builder strings.Builder

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			builder.WriteString(string(matrix[row][col].Char))
		}
		builder.WriteString("\n")
	}

	return builder.String()
}

func findMatches(matrix [][]Character) int {
	matches := 0
	primaryDiagonal, secondaryDiagonal := flattenAllDiagonals(matrix)
	rotatedMatrix := rotateMatrix90Clockwise(matrix)

	str := matrixToString(matrix)
	xmas := "XMAS"
	reverse := "SAMX"
	matches = strings.Count(str, xmas)
	matches += strings.Count(str, reverse)

	str = matrixToString(rotatedMatrix)
	matches += strings.Count(str, xmas)
	matches += strings.Count(str, reverse)

	str = matrixToString(primaryDiagonal)
	matches += strings.Count(str, xmas)
	matches += strings.Count(str, reverse)

	str = matrixToString(secondaryDiagonal)
	matches += strings.Count(str, xmas)
	matches += strings.Count(str, reverse)

	return matches
}

func Puzzle1() int {
	content := readFileCharactersByLines("./inputs/day4.txt")

	count := findMatches(content)

	return count
}
