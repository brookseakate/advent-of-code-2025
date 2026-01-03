package day4

import (
	"aoc2025/util"
	"fmt"
	"slices"
	"strings"
)

func Main() {
	//input := util.ReadInputFileToStringSlice("day4/sample")
	input := util.ReadInputFileToStringSlice("day4/input")

	//total := calculateAccessibleRollTotal(input) // part 1
	total := calculateRemovableRollTotal(input)
	fmt.Printf("TOTAL: %v\n", total)
}

// part 2
func calculateRemovableRollTotal(input []string) int {
	total := 0
	justRemoved := 1
	var updatedGrid []string

	for justRemoved != 0 {
		justRemoved, updatedGrid = removeRolls(input)

		total += justRemoved

		fmt.Printf("Debug: justRemoved: %v\n", justRemoved)
		fmt.Printf("Debug: input:\n %v\n\n", updatedGrid)

		input = updatedGrid
	}

	return total
}

// for part 2. Returns: totalRemoved this round, updatedGrid
func removeRolls(input []string) (int, []string) {
	roundTotal := 0
	updatedGrid := slices.Clone(input)

	for i, row := range input {
		for j, _ := range row {
			//fmt.Printf("Debug: Examining %v, %v: %c\n", i, j, input[i][j])

			if input[i][j] == '@' && isAccessibleRoll(input, i, j) {
				//fmt.Printf("Debug: it's accessible!\n\n")
				roundTotal++

				splitRow := strings.Split(updatedGrid[i], "")
				splitRow[j] = "."
				updatedRow := strings.Join(splitRow, "")
				updatedGrid[i] = updatedRow
			}
		}
	}

	return roundTotal, updatedGrid
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
