package day3

import (
	"aoc2025/util"
	"fmt"
	"strconv"
)

func Main() {
	//inputlines := util.ReadInputFileToStringSlice("day3/sample")
	inputlines := util.ReadInputFileToStringSlice("day3/input")

	bankMaxJoltages := getMaxJoltages(inputlines)
	var totalJoltage int
	for _, joltage := range bankMaxJoltages {
		totalJoltage += joltage
	}

	fmt.Printf(">>>>>MAX JOLTAGE TOTAL: %v\n", totalJoltage)
}

func getMaxJoltages(bankStrings []string) []int {
	var maxJoltages []int

	for _, bankString := range bankStrings {
		maxJoltages = append(maxJoltages, getMaxJoltage(bankString))
	}
	return maxJoltages
}

func getMaxJoltage(bankString string) int {
	currentMax := 0
	for i, digitOne := range bankString[:len(bankString)-1] {
		for _, digitTwo := range bankString[i+1:] {
			joltage, _ := strconv.Atoi(fmt.Sprintf("%c%c", digitOne, digitTwo))

			if joltage > currentMax {
				currentMax = joltage
			}
		}
	}
	return currentMax
}
