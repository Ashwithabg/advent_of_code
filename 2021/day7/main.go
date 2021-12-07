package main

import (
	"fmt"
	"sort"
	"time"

	"advent_of_code/2021/day7/input"
	"advent_of_code/utils"
)

const veryLargeNumber = 999999999

func findLeastFuelWhenFuelConsumptionIsAtConstantRate(crabPositions []int) int {
	sort.Ints(crabPositions)
	midPosition := utils.GetMidElement(crabPositions)

	result := 0
	for _, currentPosition := range crabPositions {
		result += utils.Abs(currentPosition - midPosition)
	}

	return result
}

func findLeastFuelConsumption(crabPositions []int, getFuelConsumption func(int,int)(int)) int {
	sort.Ints(crabPositions)
	maxIndex := len(crabPositions)
	leastFuelConsumption := veryLargeNumber

	for i := 0; i < crabPositions[maxIndex-1]; i++ {
		fuelConsumption := 0
		for _, position := range crabPositions {
			fuelConsumption += getFuelConsumption(position, i)
		}

		if fuelConsumption < leastFuelConsumption {
			leastFuelConsumption = fuelConsumption
		}
	}

	return leastFuelConsumption
}

func findLeastFuelWhenFuelConsumptionIsAtConstantRateByIteratingAllPositions(crabPositions []int) int {
	return findLeastFuelConsumption(crabPositions, func(rate int, position int) int {
		return utils.Abs(position - rate)
	})
}

func findLeastFuelWhenFuelConsumptionIsNotInConstantRate(crabPositions []int) int {
	return findLeastFuelConsumption(crabPositions, func(rate int, position int) int {
		return utils.SummationOfNaturalNumbers(utils.Abs(position - rate))
	})
}

func main() {
	fmt.Println("Start time: ", time.Now().String())
	elements, err := input.GetCrabSubmarineHorizontalPositions()
	if err != nil {
		fmt.Errorf("Error while parsing input: %+v\n", err)
	}

	res := findLeastFuelWhenFuelConsumptionIsAtConstantRate(elements)
	fmt.Println("lead fuel consumption part1:", res)

	res = findLeastFuelWhenFuelConsumptionIsAtConstantRateByIteratingAllPositions(elements)
	fmt.Println("lead fuel consumption part1:", res)

	res = findLeastFuelWhenFuelConsumptionIsNotInConstantRate(elements)
	fmt.Println("lead fuel consumption part2:", res)

	fmt.Println("end time: ", time.Now().String())
}
