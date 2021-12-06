package main

import (
	"fmt"
	"strconv"
	"strings"

	"advent_of_code/utils"
)

func calculateNumberOfFishes(fishAges []int, dayCount int) int {
	const maxAgeForFishes = 8

	fishCountWithAges := make(map[int]int)

	for _, fishAge := range fishAges {
		fishCountWithAges[fishAge] += 1
	}

	for dayIndex := 0; dayIndex < dayCount; dayIndex++ {
		newlyCreatedFishCountWithAge := make(map[int]int)

		for age := 0; age <= maxAgeForFishes; age++ {
			noOfFishes := fishCountWithAges[age]

			if age == 0 {
				delete(fishCountWithAges, 0)
				newlyCreatedFishCountWithAge[6] += noOfFishes
				newlyCreatedFishCountWithAge[8] += noOfFishes
			} else {
				fishCountWithAges[age] -= noOfFishes
				fishCountWithAges[age-1] += noOfFishes
			}
		}

		for age, numberOfFishes := range newlyCreatedFishCountWithAge {
			fishCountWithAges[age] += numberOfFishes
		}
	}

	sumOfFishes := 0
	for _, fishCount := range fishCountWithAges {
		sumOfFishes += fishCount
	}

	return sumOfFishes
}

func main() {
	elements, err := getInput()
	if err != nil {
		fmt.Errorf("Error while parsing input: %+v", err)
	}

	res := calculateNumberOfFishes(elements, 256)
	fmt.Println("result day1 part1:", res)
}

func getInput() ([]int, error) {
	filePath := "/Users/ashwitha/GolandProjects/advent_of_code/2021/day6/input1.txt"

	lines, err := utils.ReadLines(filePath)
	if err != nil {
		return nil, err
	}

	elements := strings.Split(lines[0], ",")

	var numbers []int
	for _, ele := range elements {
		n, _ := strconv.Atoi(ele)
		numbers = append(numbers, n)
	}

	return numbers, nil
}
