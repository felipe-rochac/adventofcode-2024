package day4

func findMatches2(matrix [][]Character) int {
	rows := len(matrix)
	cols := len(matrix[0])
	matches := 0

	// Collect all primary diagonals (top-left to bottom-right)
	for row := 1; row < rows-1; row++ {
		for col := 1; col < cols-1; col++ {
			if matrix[row][col].Char != 'A' {
				continue
			}

			//  1 | 2 | 3 | 4 |
			// S S|S M|M M|M S|
			//  A | A | A | A |
			// M M|S M|S S|M S|
			case1 := matrix[row-1][col-1].Char == 'S' && matrix[row+1][col+1].Char == 'M' && matrix[row-1][col+1].Char == 'S' && matrix[row+1][col-1].Char == 'M'
			case2 := matrix[row-1][col-1].Char == 'S' && matrix[row+1][col+1].Char == 'M' && matrix[row-1][col+1].Char == 'M' && matrix[row+1][col-1].Char == 'S'
			case3 := matrix[row-1][col-1].Char == 'M' && matrix[row+1][col+1].Char == 'S' && matrix[row-1][col+1].Char == 'M' && matrix[row+1][col-1].Char == 'S'
			case4 := matrix[row-1][col-1].Char == 'M' && matrix[row+1][col+1].Char == 'S' && matrix[row-1][col+1].Char == 'S' && matrix[row+1][col-1].Char == 'M'

			if !case1 && !case2 && !case3 && !case4 {
				continue
			}

			matches++
		}
	}

	return matches
}

func Puzzle2() int {
	content := readFileCharactersByLines("./inputs/day4.txt")

	count := findMatches2(content)

	return count
}
