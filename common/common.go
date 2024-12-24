package common

import (
	"bufio"
	"log"
	"os"
)

func ReadFileByLines(filename string) []string {
	file, err := os.Open(filename)
	lines := make([]string, 0)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func ReadFileByGrid(filename string) [][]rune {
	file, err := os.Open(filename)

	grid := make([][]rune, 0)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := []rune(scanner.Text())
		gridLine := make([]rune, 0)
		for _, r := range line {
			gridLine = append(gridLine, rune(r))
		}
		grid = append(grid, gridLine)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return grid
}
