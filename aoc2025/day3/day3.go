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
		//maxJoltages = append(maxJoltages, getMaxJoltageTwoBatteries(bankString)) // part 1
		maxJoltages = append(maxJoltages, naiveGetMaxJoltageTwelveBatteries(bankString)) // part 2
	}
	return maxJoltages
}

func getMaxJoltageTwoBatteries(bankString string) int {
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

// expecting this to not work due to time complexity, but will try naive version first
func naiveGetMaxJoltageTwelveBatteries(bankString string) int {
	currentMax := 0
	for i, digitOne := range bankString[:len(bankString)-12] {
		for j, digitTwo := range bankString[i+1 : len(bankString)-11] {
			for k, digitThree := range bankString[j+1 : len(bankString)-10] {
				for l, digitFour := range bankString[k+1 : len(bankString)-9] {
					for m, digitFive := range bankString[l+1 : len(bankString)-8] {
						for n, digitSix := range bankString[m+1 : len(bankString)-7] {
							for o, digitSeven := range bankString[n+1 : len(bankString)-6] {
								for p, digitEight := range bankString[o+1 : len(bankString)-5] {
									for q, digitNine := range bankString[p+1 : len(bankString)-4] {
										for r, digitTen := range bankString[q+1 : len(bankString)-3] {
											for s, digitEleven := range bankString[r+1 : len(bankString)-2] {
												for _, digitTwelve := range bankString[s+1:] {

													joltage, _ := strconv.Atoi(fmt.Sprintf("%c%c%c%c%c%c%c%c%c%c%c%c", digitOne, digitTwo, digitThree, digitFour, digitFive, digitSix, digitSeven, digitEight, digitNine, digitTen, digitEleven, digitTwelve))

													if joltage > currentMax {
														currentMax = joltage
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return currentMax
}
