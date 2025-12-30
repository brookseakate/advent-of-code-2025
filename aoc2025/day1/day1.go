package day1

import (
	"aoc2025/util"
	"fmt"
	"strconv"
)

func Main() {
	inputLines := util.ReadInputFileToStringSlice("day1/sample")
	//inputLines := util.ReadInputFileToStringSlice("day1/input")
	totalZeros := part2(inputLines)

	fmt.Printf(">>>>>Total Zeros: %v\n", totalZeros)
}

// modified part1 -> part2 in place
func part2(instructionLines []string) int {
	currentPosition := 50
	totalZeros := 0
	dialMax := 99

	for _, line := range instructionLines {
		fmt.Printf("Debug: PROCESSING instruction: %v\n", line)
		var zeroPasses int
		currentPosition, zeroPasses = rotate(dialMax, currentPosition, line)

		totalZeros += zeroPasses

		fmt.Printf("Debug: CURRENT POSITION after instruction: %v == %v. zeroPasses: %v\n\n", line, currentPosition, zeroPasses)
		if currentPosition == 0 {
			totalZeros++
		}
	}

	return totalZeros
}

// returns: position after rotation, how many times we rotated past 0
func rotate(dialMax int, currentPosition int, instruction string) (int, int) {
	rotationInt := parseInstruction(instruction)

	mightBeOOB := currentPosition + rotationInt

	endPosition, zeroPasses := normalizeToDialBounds(mightBeOOB, dialMax, 0) // reset zeroPasses to 0 for every new instruction
	return endPosition, zeroPasses
}

// returns: in-bounds position after rotation, how many times we rotated past 0
func normalizeToDialBounds(potentialEndPosition int, dialMax int, zeroPasses int) (int, int) {
	fmt.Printf("Debug: MIGHT BE OOB: %v\n", potentialEndPosition)

	// break condition
	if potentialEndPosition >= 0 && potentialEndPosition <= dialMax {
		return potentialEndPosition, zeroPasses
	}

	// recurse
	var endPosition int
	if potentialEndPosition < 0 {
		endPosition = dialMax + 1 + potentialEndPosition

		if endPosition != 0 { // don't double-count if ends on zero
			zeroPasses++
		}
	} else if potentialEndPosition > dialMax {
		endPosition = potentialEndPosition - (dialMax + 1)

		if endPosition != 0 { // don't double-count if ends on zero
			zeroPasses++
		}
	} else {
		//println("THIS SHOULDN'T HAVE HAPPENED, bad normalizeToDialBounds condition")
		panic("THIS SHOULDN'T HAVE HAPPENED, bad normalizeToDialBounds condition")
	}
	return normalizeToDialBounds(endPosition, dialMax, zeroPasses)
}

func parseInstruction(instructionString string) int {
	direction := instructionString[0]
	clicks := instructionString[1:]

	var rotationInt int
	if direction == 'L' {
		rotationInt, _ = strconv.Atoi("-" + clicks)
	} else if direction == 'R' {
		rotationInt, _ = strconv.Atoi(clicks)
	} else {
		//println("OOOOPS that's not L or R")
		panic("OOOOPS that's not L or R")
	}

	return rotationInt
}
