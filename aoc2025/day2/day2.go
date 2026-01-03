package day2

import (
	"aoc2025/util"
	"fmt"
	"strconv"
	"strings"
)

func Main() {
	//inputLine := util.ReadSingleLineInputFileToString("day2/sample")
	inputLine := util.ReadSingleLineInputFileToString("day2/input")

	idRanges := parseIdRanges(inputLine)
	//invalidIds := collectTwiceRepeatingSubstringIds(idRanges) // part 1
	invalidIds := collectArbitraryRepeatingSubstringIds(idRanges) // part 2

	var invalidIdTotal int64
	for _, invalidId := range invalidIds {
		invalidIdTotal += invalidId
	}

	fmt.Printf(">>>> INVALID ID TOTAL: %v", invalidIdTotal)
}

func parseIdRanges(inputLine string) [][]int64 {
	idRanges := [][]int64{}

	rangeStrings := strings.Split(inputLine, ",")
	for _, rangeString := range rangeStrings {
		rangeStartAndEnd := strings.Split(rangeString, "-")
		rangeStart, _ := strconv.ParseInt(rangeStartAndEnd[0], 10, 64)
		rangeEnd, _ := strconv.ParseInt(rangeStartAndEnd[1], 10, 64)
		rangeSlice := make([]int64, rangeEnd+1-rangeStart)

		for i := range rangeSlice {
			rangeSlice[i] = rangeStart + int64(i)
		}

		if rangeSlice[0] != rangeStart {
			panic("oops, rangeStart isn't right")
		}

		if rangeSlice[len(rangeSlice)-1] != rangeEnd {
			panic("oops, rangeEnd isn't right")
		}

		idRanges = append(idRanges, rangeSlice)
	}

	return idRanges
}

// part 1 method; returns ids composed of a substring repeated exactly twice
func collectTwiceRepeatingSubstringIds(idRanges [][]int64) []int64 {
	invalidIds := []int64{}

	for _, idRange := range idRanges {
		for _, id := range idRange {
			idString := strconv.FormatInt(id, 10)
			firstHalf := idString[0 : len(idString)/2]
			lastHalf := idString[len(idString)/2:]

			if firstHalf == lastHalf {
				invalidIds = append(invalidIds, id)
			}
		}
	}
	return invalidIds
}

// part 2 method: returns ids composed of a substring repeated any number of times
func collectArbitraryRepeatingSubstringIds(idRanges [][]int64) []int64 {
	invalidIds := []int64{}

	for _, idRange := range idRanges {
		for _, id := range idRange {
			// iterate substrings (half-string down to 1 char), check that string contains only substring matches
			idString := strconv.FormatInt(id, 10)
			for i := len(idString) / 2; i > 0; i-- {
				pattern := idString[0:i]
				strippedString := strings.ReplaceAll(idString, pattern, "")

				if len(strippedString) == 0 {
					fmt.Printf("Debug: invalidId: %v\n", id)
					invalidIds = append(invalidIds, id)
					break
				}
			}
		}
	}
	return invalidIds
}
