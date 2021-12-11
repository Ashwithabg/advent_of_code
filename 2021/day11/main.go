package main

import (
	"advent_of_code/2021/day11/input"
	"advent_of_code/2021/day11/input/models"
	"fmt"
)

func main() {
	octopuses, err := input.GetOctopusGrid()
	if err != nil {
		fmt.Errorf("Error while parsing input: %+v", err)
	}
	res := day11(octopuses)
	fmt.Println("result day11 part1:", res)

	octopuses, err = input.GetOctopusGrid()
	if err != nil {
		fmt.Errorf("Error while parsing input: %+v", err)
	}
	res = findStepAtWhichAllOctopusEnergize(octopuses)
	fmt.Println("result day11 part2:", res)
}

func day11(octopuses models.OctopusGrid) int {
	flashCounter := 0
	for step := 0; step < 100; step++ {
		(&octopuses).EnergizeOctopuses()
		(&octopuses).FlashEnergizedOctopuses(&flashCounter)
	}

	return flashCounter
}

func findStepAtWhichAllOctopusEnergize(octopuses models.OctopusGrid) int {
	flashCounter := 0
	for step := 0; step < 600; step++ {
		stepFlashCounter := 0
		(&octopuses).EnergizeOctopuses()
		(&octopuses).FlashEnergizedOctopuses(&stepFlashCounter)
		flashCounter += stepFlashCounter

		if stepFlashCounter == 100 {
			return step + 1
		}
	}

	return 0
}
