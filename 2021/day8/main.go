package main

import (
	"advent_of_code/2021/day8/input"
	"advent_of_code/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const noOfDigits = 10

func main() {
	fmt.Println("Start time: ", time.Now().String())

	parseDisplayFor4digits()
	parseDisplayForAllDigits()

	fmt.Println("end time: ", time.Now().String())
}

func parseDisplayFor4digits() {
	lines := input.GetInput()

	var values []string
	for i := 0; i < len(lines); i++ {
		data := strings.Split(lines[i], " | ")
		values = append(values, data[1])
	}

	counter := 0
	for i := 0; i < len(values); i++ {
		data := strings.Split(values[i], " ")
		for j := 0; j < len(data); j++ {
			if doesLengthMatch(len(data[j])) {
				counter++
			}

		}
	}
	fmt.Println("day8 part1", counter)
}

func doesLengthMatch(length int) bool {
	return length == 2 || length == 4 || length == 3 || length == 7
}

func parseDisplayForAllDigits() {
	lines := input.GetInput()
	var totalSum int
	for _, line := range lines {
		data := strings.Split(line, " | ")
		inputs := strings.Split(data[0], " ")
		outputs := strings.Split(data[1], " ")

		result := findBasedOnLength(inputs)
		result = findBasedOnCharacterOverlaps(inputs, result)
		sortedInput := utils.SortValues(result)

		var outputString string
		for _, output := range outputs {
			output = utils.SortString(output)
			for k, input := range sortedInput {
				if input == output {
					outputString += strconv.Itoa(k)
				}
			}
		}

		totalSum += utils.ToInt(outputString)
	}

	fmt.Println(totalSum)
}

func findBasedOnCharacterOverlaps(inputs []string, result [10]string) [10]string {
	for i := 0; i < noOfDigits; i++ {
		for _, input := range inputs {
			if isTwo(result, input) {
				result[2] = input
			} else if isZero(result, input) {
				result[0] = input
			} else if isThree(result, input) {
				result[3] = input
			} else if isFive(result, input) {
				result[5] = input
			} else if isSix(result, input) {
				result[6] = input
			} else if isNine(result, input) {
				result[9] = input
			}
		}
	}
	return result
}

func findBasedOnLength(inputs []string) [10]string {
	result := [10]string{}
	for _, input := range inputs {
		length := len(input)
		if length == 2 {
			result[1] = input
		} else if length == 3 {
			result[7] = input
		} else if length == 4 {
			result[4] = input
		} else if length == 7 {
			result[8] = input
		}
	}
	return result
}

func isZero(wireMaps [10]string, input string) bool {
	return wireMaps[0] == "" &&
		len(input) == 6 &&
		utils.Overlaps(input, wireMaps[1]) == 2 &&
		utils.Overlaps(input, wireMaps[2]) == 4 &&
		utils.Overlaps(input, wireMaps[3]) == 4
}

func isSix(wireMaps [10]string, input string) bool {
	return wireMaps[6] == "" &&
		len(input) == 6 &&
		utils.Overlaps(input, wireMaps[1]) == 1 &&
		utils.Overlaps(input, wireMaps[2]) == 4
}

func isFive(wireMaps [10]string, input string) bool {
	return wireMaps[5] == "" &&
		len(input) == 5 &&
		utils.Overlaps(input, wireMaps[2]) == 3 &&
		utils.Overlaps(input, wireMaps[4]) == 3
}

func isNine(wireMaps [10]string, input string) bool {
	return wireMaps[9] == "" &&
		len(input) == 6 &&
		utils.Overlaps(input, wireMaps[0]) == 5 &&
		utils.Overlaps(input, wireMaps[7]) == 3

}

func isThree(wireMaps [10]string, input string) bool {
	return wireMaps[3] == "" &&
		len(input) == 5 &&
		utils.Overlaps(input, wireMaps[1]) == 2
}

func isTwo(misplacedWireMaps [10]string, input string) bool {
	return misplacedWireMaps[2] == "" &&
		len(input) == 5 &&
		utils.Overlaps(input, misplacedWireMaps[1]) == 1 &&
		utils.Overlaps(input, misplacedWireMaps[4]) == 2 &&
		utils.Overlaps(input, misplacedWireMaps[3]) == 4 &&
		utils.Overlaps(input, misplacedWireMaps[7]) == 2
}
