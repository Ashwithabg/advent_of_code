package main

import (
	"fmt"
	"time"

	"advent_of_code/2021/day8/input"
)

func day8fn(crabPositions []int) int {

	return 0
}

func main() {
	fmt.Println("Start time: ", time.Now().String())
	_, err := input.Day8Part2Parse()
	if err != nil {
		fmt.Errorf("Error while parsing input: %+v\n", err)
	}

	//res := day8fn(elements)
	//fmt.Println("day8 part1:", res)

	fmt.Println("end time: ", time.Now().String())
}
