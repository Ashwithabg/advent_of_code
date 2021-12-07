package main

import (
	"fmt"
	"sort"

	"advent_of_code/2021/day7/input"
	"advent_of_code/utils"
)

const veryLargeNumber = 999999999

func findLeastFuelWhenFuelConsumptionIsAtConstantRate(crabPositions []int) int {
	result := 0
	sort.Ints(crabPositions)
	midIndex := len(crabPositions) / 2
	midPosition := crabPositions[midIndex]

	for _, currentPosition := range crabPositions {
		result += utils.Abs(currentPosition - midPosition)
	}

	return result
}

func findLeastFuelWhenFuelConsumptionIsAtConstantRateByIteratingAllPositions(crabPositions []int) int {
	sort.Ints(crabPositions)
	maxIndex := len(crabPositions)
	leastFuelConsumption := veryLargeNumber

	for i := 0; i < crabPositions[maxIndex-1]; i++ {
		fuelConsumption := 0
		for _, position := range crabPositions {
			fuelConsumption += utils.Abs(position - i)
		}

		if fuelConsumption < leastFuelConsumption {
			leastFuelConsumption = fuelConsumption
		}
	}

	return leastFuelConsumption
}

func findLeastFuelWhenFuelConsumptionIsNotInConstantRate(crabPositions []int) int {
	maxIndex := len(crabPositions)
	leastFuelConsumption := veryLargeNumber

	for i := 0; i < crabPositions[maxIndex-1]; i++ {
		fuelConsumption := 0
		for _, position := range crabPositions {
			fuelConsumption +=  sumOfNaturalNumbers(utils.Abs(position - i))
		}

		if fuelConsumption < leastFuelConsumption {
			leastFuelConsumption = fuelConsumption
		}
	}

	return leastFuelConsumption
}

func sumOfNaturalNumbers(numberOfElements int) int {
	return numberOfElements * (numberOfElements + 1) / 2
}

func main() {
	elements, err := input.GetCrabSubmarineHorizontalPositions()
	if err != nil {
		fmt.Errorf("Error while parsing input: %+v", err)
	}

	res := findLeastFuelWhenFuelConsumptionIsAtConstantRate(elements)
	fmt.Println("lead fuel consumption part1:", res)

	res = findLeastFuelWhenFuelConsumptionIsAtConstantRateByIteratingAllPositions(elements)
	fmt.Println("lead fuel consumption part1:", res)

	res = findLeastFuelWhenFuelConsumptionIsNotInConstantRate(elements)
	fmt.Println("lead fuel consumption part2:", res)
}
