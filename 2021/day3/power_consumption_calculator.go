package main

import (
	"fmt"
	"strconv"
	"strings"

	"advent_of_code/utils"
)

func calculatePower() (int64, error) {
	filePath := "/Users/ashwitha/GolandProjects/advent_of_code/2021/day3/input.txt"
	elements, err := utils.ReadLines(filePath)
	if err != nil {
		return 0, err
	}

	onesCount := make([]int, len(elements[0]))
	for i := 0; i < len(elements); i++ {
		convertedValue := strings.Split(elements[i], "")
		for j := 0; j < len(elements[i]); j++ {
			if convertedValue[j] == "1" {
				onesCount[j]++
			}
		}
	}

	value := ""
	for i := 0; i < len(onesCount); i++ {
		if onesCount[i] > (len(elements) / 2) {
			value += "1"
		} else {
			value += "0"
		}
	}
	firstValue:= convertBinaryToDecimal(value)

	secondBinaryValue := ""
	for i := 0; i < len(onesCount); i++ {
		if onesCount[i] > (len(elements) / 2) {
			secondBinaryValue += "0"
		} else {
			secondBinaryValue += "1"
		}
	}
	secondValue := convertBinaryToDecimal(secondBinaryValue)

	return firstValue * secondValue, nil
}

func calculateLifeSupportRatingOfSubmarine() (int64, error) {
	filePath := "/Users/ashwitha/GolandProjects/advent_of_code/2021/day3/input.txt"

	elements, err := utils.ReadLines(filePath)
	if err != nil {
		return 0, err
	}

	var onebestValues []string
	var zerobestValues []string
	var bestValue = elements

	for k := 0; k < len(elements[0]); k++ {
		onebestValues = []string{}
		zerobestValues = []string{}
		for j := 0; j < len(bestValue); j++ {
			convertedValue := strings.Split(bestValue[j], "")
			if convertedValue[k] == "1" {
				onebestValues = append(onebestValues, bestValue[j])
			} else {
				zerobestValues = append(zerobestValues, bestValue[j])
			}
		}

		if len(onebestValues) >= len(zerobestValues) {
			bestValue = onebestValues
		} else {
			bestValue = zerobestValues
		}

		if len(bestValue) == 1 {
			break
		}
	}

	firstValue := convertBinaryToDecimal(bestValue[0])

	bestValue = elements
	for k := 0; k < len(elements[0]); k++ {
		onebestValues = []string{}
		zerobestValues = []string{}
		for j := 0; j < len(bestValue); j++ {
			convertedValue := strings.Split(bestValue[j], "")
			if convertedValue[k] == "1" {
				onebestValues = append(onebestValues, bestValue[j])
			} else {
				zerobestValues = append(zerobestValues, bestValue[j])
			}
		}

		if len(onebestValues) < len(zerobestValues) {
			bestValue = onebestValues
		} else {
			bestValue = zerobestValues
		}
		if len(bestValue) == 1 {
			break
		}
	}
	secondValue := convertBinaryToDecimal(bestValue[0])
	return firstValue * secondValue, nil
}

func convertBinaryToDecimal(binaryReading string) int64 {
	reading, err := strconv.ParseInt(binaryReading, 2, 64)
	if err != nil {
		fmt.Errorf("Error on converting binary To decimal %+v\n", err)
		return 0
	}

	return reading
}

func main() {
	res, err := calculatePower()
	if err != nil {
		fmt.Errorf("error %+v", err)
	}

	fmt.Println("result day1 part1:", res)

	res, err = calculateLifeSupportRatingOfSubmarine()
	if err != nil {
		fmt.Errorf("error %+v", err)
	}

	fmt.Println("result day1 part2:", res)
}
