package day1

import (
	"aoc2025/util"
	"fmt"
	"strconv"
)

func Main() {
	//inputLines := util.ReadInputFileToStringSlice("day1/sample")
	inputLines := util.ReadInputFileToStringSlice("day1/input")
	totalZeros := part1(inputLines)

	fmt.Printf(">>>>>Total Zeros: %v\n", totalZeros)
}

func part1(instructionLines []string) int {
	// start at 50
	// track totalZeros

	// for each line:
	// - rotate func:
	//   - rotate per instruction
	//   - return current position (optional print)
	// - if current position == 0: totalZeros++

	// return totalZeros

	currentPosition := 50
	totalZeros := 0
	dialMax := 99

	for _, line := range instructionLines {
		currentPosition = rotate(dialMax, currentPosition, line)

		fmt.Printf("Debug: CURRENT POSITION after instruction: %v == %v\n", line, currentPosition)
		if currentPosition == 0 {
			totalZeros++
		}
	}

	return totalZeros
}

// returns position after rotation
func rotate(dialMax int, currentPosition int, instruction string) int {
	rotationInt := parseInstruction(instruction)

	mightBeOOB := currentPosition + rotationInt

	fmt.Printf("Debug: MIGHT BE OOB: %v\n", mightBeOOB)

	var endPosition int

	// Assumption: only supports clicks <= 100
	if mightBeOOB < 0 {
		endPosition = dialMax + 1 + mightBeOOB
	} else if mightBeOOB > dialMax {
		endPosition = mightBeOOB - (dialMax + 1)
	} else {
		endPosition = mightBeOOB
	}

	return endPosition
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
		println("OOOOPS that's not L or R")
	}

	return rotationInt
}
