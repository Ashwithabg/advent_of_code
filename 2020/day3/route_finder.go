package main

import (
	"fmt"

	"advent_of_code/utils"
)

func findNumberOfTrees(elements []int) int {
	result := 0
	for i := 0; i < len(elements); i++ {




	}

	return result
}

func main() {
	elements, err := getInput()
	if err != nil {
		fmt.Errorf("Error while parsing input: %+v", err)
	}

	res := findNumberOfTrees(elements)
	fmt.Println("result day1 part1:", res)
}

func getInput() ([]int, error) {
	filePath := "/Users/ashwitha/GolandProjects/advent_of_code/2020/day3/input.txt"

	elements, err := utils.ReadNumbersFromFile(filePath)
	if err != nil {
		return nil, err
	}
	return elements, nil
}
