package main

import (
	"fmt"

	"advent_of_code/utils"
)

func CheckIncrease() (int, error) {
	filePath := "/Users/ashwitha/github/advent-of-code-go/2021/day1/input.txt"

	elements, err := utils.ReadNumbersFromFile(filePath)
	if err != nil {
		return 0, err
	}

	result := 0
	for i := 1; i < len(elements); i++ {
		if elements[i] > elements[i-1] {
			result++
		}
	}

	return result, nil
}

func checkIncreaseWithSlidingWindow() (int, error) {
	filePath := "/Users/ashwitha/github/advent-of-code-go/2021/day1/input.txt"

	elements, err := utils.ReadNumbersFromFile(filePath)
	if err != nil {
		return 0, err
	}

	result := 0
	previousSum := -1
	for i := 0; i < len(elements)-3; i++ {
		currentSum := 0
		for j := i; j < i+3; j++ {
			currentSum += elements[j]
		}

		if currentSum > previousSum {
			result++
		}

		previousSum = currentSum
	}

	return result, nil
}

func main() {
	res, err := CheckIncrease()
	if err != nil {
		fmt.Errorf("error %+v", err)
	}

	fmt.Println("result day1 part1:", res)

	res, err = checkIncreaseWithSlidingWindow()
	if err != nil {
		fmt.Errorf("error %+v", err)
	}

	fmt.Println("result day1 part2:", res)
}
