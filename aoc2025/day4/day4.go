package day4

import (
	"aoc2025/util"
	"fmt"
)

func Main() {
	//input := util.ReadInputFileToStringSlice("day4/sample")
	input := util.ReadInputFileToStringSlice("day4/input")

	accessiblePaperRollTotal := calculateAccessibleRollTotal(input)
	fmt.Printf("TOTAL: %v\n", accessiblePaperRollTotal)
}

// part 1
func calculateAccessibleRollTotal(input []string) int {
	total := 0

	for i, row := range input {
		for j, _ := range row {
			//fmt.Printf("Debug: Examining %v, %v: %c\n", i, j, input[i][j])

			if input[i][j] == '@' && isAccessibleRoll(input, i, j) {
				//fmt.Printf("Debug: it's accessible!\n\n")
				total++
			}
		}
	}

	return total
}

func isAccessibleRoll(input []string, i int, j int) bool {
	adjacentRolls := 0

	if i > 0 && j > 0 && input[i-1][j-1] == '@' {
		adjacentRolls++
	}

	if i > 0 && input[i-1][j] == '@' {
		adjacentRolls++
	}

	if i > 0 && j < len(input[i])-1 && input[i-1][j+1] == '@' {
		adjacentRolls++
	}

	if j < len(input[i])-1 && input[i][j+1] == '@' {
		adjacentRolls++
	}

	if i < len(input)-1 && j < len(input[i])-1 && input[i+1][j+1] == '@' {
		adjacentRolls++
	}

	if i < len(input)-1 && input[i+1][j] == '@' {
		adjacentRolls++
	}

	if i < len(input)-1 && j > 0 && input[i+1][j-1] == '@' {
		adjacentRolls++
	}

	if j > 0 && input[i][j-1] == '@' {
		adjacentRolls++
	}

	if adjacentRolls < 4 {
		return true
	}

	return false
}
