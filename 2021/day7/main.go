package main

import (
	"fmt"
	"sort"

	"advent_of_code/2021/day7/input"
	"advent_of_code/utils"
)

const veryLargeNumber = 999999999

func findLeastFuel(elements []int) int {
	result := 0
	sort.Ints(elements)
	midIndex := len(elements) / 2
	midElement := elements[midIndex]

	for _, ele := range elements {
		result += utils.Abs(ele - midElement)
	}

	return result
}

func findLeastFuelByIteratingAllElements(elements []int) int {
	sort.Ints(elements)
	maxIndex := len(elements)
	sum := veryLargeNumber

	for i := 0; i < elements[maxIndex-1]; i++ {
		elementSums := 0
		for _, element := range elements {
			elementSums += utils.Abs(element - i)
		}

		if elementSums < sum {
			sum = elementSums
		}
	}

	return sum
}

func findLeastFuelWithCounter(elements []int) int {
	maxIndex := len(elements)
	sum := veryLargeNumber

	for i := 0; i < elements[maxIndex-1]; i++ {
		elementSums := 0
		for _, element := range elements {
			elementSums +=  sumOfNaturalNumbers(utils.Abs(element - i))
		}

		if elementSums < sum {
			sum = elementSums
		}
	}

	return sum
}

func sumOfNaturalNumbers(numberOfElements int) int {
	return numberOfElements * (numberOfElements + 1) / 2
}

func main() {
	elements, err := input.GetInput()
	if err != nil {
		fmt.Errorf("Error while parsing input: %+v", err)
	}

	res := findLeastFuel(elements)
	fmt.Println("result findLeastFuel part1:", res)

	res = findLeastFuelByIteratingAllElements(elements)
	fmt.Println("result findLeastFuel part1:", res)

	res = findLeastFuelWithCounter(elements)
	fmt.Println("result findLeastFuel part2:", res)
}
