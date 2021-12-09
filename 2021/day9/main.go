package main

import (
	"fmt"

	"advent_of_code/2021/day9/input"
)

var moves = [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

func main() {
	matrix, err := input.GetInput()
	if err != nil {
		fmt.Errorf("Error while parsing input: %+v", err)
	}

	res := matrix.FindRiskLevel()
	fmt.Println("result day9 part1:", res)

	finalResult := matrix.FindSizeOfThreeLargestBasins()
	fmt.Println("result day9 part2:", finalResult)
}
