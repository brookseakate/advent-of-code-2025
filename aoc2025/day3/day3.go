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
		maxJoltages = append(maxJoltages, getMaxJoltageTwelveBatteries(bankString)) // part 2
	}
	return maxJoltages
}

// part 1
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

// part 2 – expecting this to not work due to time complexity, but will try naive version first
func naiveGetMaxJoltageTwelveBatteries(bankString string) int {
	currentMax := 0
	for i, digitOne := range bankString[:len(bankString)-11] {
		for j, digitTwo := range bankString[i+1 : len(bankString)-10] {
			for k, digitThree := range bankString[j+1 : len(bankString)-9] {
				for l, digitFour := range bankString[k+1 : len(bankString)-8] {
					for m, digitFive := range bankString[l+1 : len(bankString)-7] {
						for n, digitSix := range bankString[m+1 : len(bankString)-6] {
							for o, digitSeven := range bankString[n+1 : len(bankString)-5] {
								for p, digitEight := range bankString[o+1 : len(bankString)-4] {
									for q, digitNine := range bankString[p+1 : len(bankString)-3] {
										for r, digitTen := range bankString[q+1 : len(bankString)-2] {
											for s, digitEleven := range bankString[r+1 : len(bankString)-1] {
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

// part 2 optimized
func getMaxJoltageTwelveBatteries(bankString string) int {
	digitOneInd := getRangeMaxIndex(bankString[:len(bankString)-11])
	digitTwoInd := digitOneInd + 1 + getRangeMaxIndex(bankString[digitOneInd+1:len(bankString)-10])
	digitThreeInd := digitTwoInd + 1 + getRangeMaxIndex(bankString[digitTwoInd+1:len(bankString)-9])
	digitFourInd := digitThreeInd + 1 + getRangeMaxIndex(bankString[digitThreeInd+1:len(bankString)-8])
	digitFiveInd := digitFourInd + 1 + getRangeMaxIndex(bankString[digitFourInd+1:len(bankString)-7])
	digitSixInd := digitFiveInd + 1 + getRangeMaxIndex(bankString[digitFiveInd+1:len(bankString)-6])
	digitSevenInd := digitSixInd + 1 + getRangeMaxIndex(bankString[digitSixInd+1:len(bankString)-5])
	digitEightInd := digitSevenInd + 1 + getRangeMaxIndex(bankString[digitSevenInd+1:len(bankString)-4])
	digitNineInd := digitEightInd + 1 + getRangeMaxIndex(bankString[digitEightInd+1:len(bankString)-3])
	digitTenInd := digitNineInd + 1 + getRangeMaxIndex(bankString[digitNineInd+1:len(bankString)-2])
	digitElevenInd := digitTenInd + 1 + getRangeMaxIndex(bankString[digitTenInd+1:len(bankString)-1])
	digitTwelveInd := digitElevenInd + 1 + getRangeMaxIndex(bankString[digitElevenInd+1:])

	joltage, _ := strconv.Atoi(fmt.Sprintf("%c%c%c%c%c%c%c%c%c%c%c%c", bankString[digitOneInd], bankString[digitTwoInd], bankString[digitThreeInd], bankString[digitFourInd], bankString[digitFiveInd], bankString[digitSixInd], bankString[digitSevenInd], bankString[digitEightInd], bankString[digitNineInd], bankString[digitTenInd], bankString[digitElevenInd], bankString[digitTwelveInd]))

	fmt.Printf("Debug: Joltage: %v\n", joltage)

	return joltage
}

func getRangeMaxIndex(bankString string) int {
	currentMaxIndex := 0

	for i, _ := range bankString {
		if bankString[i] > bankString[currentMaxIndex] {
			currentMaxIndex = i
		}
	}
	return currentMaxIndex
}
