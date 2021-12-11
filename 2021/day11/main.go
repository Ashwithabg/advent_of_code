package main

import (
	"advent_of_code/2021/day11/input"
	"advent_of_code/2021/day11/input/models"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go func() {
		defer wg.Done()
		octopuses, err := input.GetOctopusGrid()
		if err != nil {
			fmt.Errorf("Error while parsing input: %+v", err)
		}
		res := find100thStepFlashCount(octopuses)
		fmt.Println("result part1:", res)
	}()

	go func() {
		defer wg.Done()
		octopuses, err := input.GetOctopusGrid()
		if err != nil {
			fmt.Errorf("Error while parsing input: %+v", err)
		}
		res := findStepAtWhichAllOctopusEnergize(octopuses)
		fmt.Println("result part2:", res)
	}()

	wg.Wait()
}

func find100thStepFlashCount(octopuses models.OctopusGrid) int {
	flashCounter := 0
	for step := 0; step < 100; step++ {
		(&octopuses).EnergizeOctopuses()
		(&octopuses).FlashEnergizedOctopuses(&flashCounter)
	}

	return flashCounter
}

func findStepAtWhichAllOctopusEnergize(octopuses models.OctopusGrid) int {
	flashCounter := 0
	for step := 1; step <= 600; step++ {
		stepFlashCounter := 0
		(&octopuses).EnergizeOctopuses()
		(&octopuses).FlashEnergizedOctopuses(&stepFlashCounter)
		flashCounter += stepFlashCounter

		if stepFlashCounter == 100 {
			return step
		}
	}

	return 0
}
