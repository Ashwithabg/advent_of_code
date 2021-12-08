package main

import (
	"advent_of_code/2021/day8/input"
	"advent_of_code/2021/day8/models"
	"advent_of_code/utils"
	"fmt"
	"strconv"
	"time"
)

const noOfDigits = 10

func main() {
	fmt.Println("Start time: ", time.Now().String())

	occurance := findSumOf4digitOccurance()
	fmt.Println("day 8 part1 ", occurance)

	sum := findSumOfDisplayDigits()
	fmt.Println("day 8 part2 ", sum)

	fmt.Println("end time: ", time.Now().String())
}

func findSumOf4digitOccurance() int {
	sevenSegmentDisplayOutputs := input.GetDisplayData()

	counter := 0
	for _, displayData := range sevenSegmentDisplayOutputs {
		counter += displayData.FindCountOf234Or7InOutputs()
	}

	return counter
}

func findSumOfDisplayDigits() (sum int) {
	sevenSegmentDisplayData := input.GetDisplayData()

	for _, displayData := range sevenSegmentDisplayData {
		signalDigitPattern := &models.SignalPatterns{}

		for _, input := range displayData.SignalInputs {
			signalDigitPattern.FindSignalPatternsBasedOnLength(input)
		}

		for i := 0; i < noOfDigits; i++ {
			for _, input := range displayData.SignalInputs {
				signalDigitPattern.FindSignalPatternsBasedOnOverlaps(input)
			}
		}

		signalDigitPattern.SortValues()
		sortedOutput := displayData.GetSortedOutputs()

		var fourDigitOutputString string
		for _, output := range sortedOutput {
			ok, digit := signalDigitPattern.PatternExistsFor(string(output))
			if ok {
				fourDigitOutputString += strconv.Itoa(digit)
			}
		}

		sum += utils.ToInt(fourDigitOutputString)
	}

	return sum
}
