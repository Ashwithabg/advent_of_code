package main

import (
	"fmt"

	"advent_of_code/utils"
)

func checkIncrease(elements []int) int {
	result := 0
	for i := 1; i < len(elements); i++ {
		if elements[i] > elements[i-1] {
			result++
		}
	}

	return result
}

func checkIncreaseWithSlidingWindow(elements []int) int {
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

	return result
}

func main() {
	elements, err := getInput()
	if err != nil {
		fmt.Errorf("Error while parsing input: %+v", err)
	}

	res := checkIncrease(elements)
	fmt.Println("result day1 part1:", res)

	res = checkIncreaseWithSlidingWindow(elements)
	fmt.Println("result day1 part2:", res)
}

func getInput() ([]int, error) {
	filePath := "/Users/ashwitha/GolandProjects/advent_of_code/2021/day1/input.txt"

	elements, err := utils.ReadNumbersFromFile(filePath)
	if err != nil {
		return nil, err
	}
	return elements, nil
}
