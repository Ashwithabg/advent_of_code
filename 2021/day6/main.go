package main

import (
	"advent_of_code/2021/day6/input"
	"fmt"
)

const maxTimerForFishes = 8
const timerForFishesThatJustGaveBirth = 6

func calculateNumberOfFishes(fishTimers []int, dayCount int) int {
	fishCountWithTimers := make(map[int]int)

	for _, fishTimer := range fishTimers {
		fishCountWithTimers[fishTimer] += 1
	}

	for dayIndex := 0; dayIndex < dayCount; dayIndex++ {
		fishesThatCanGiveBirth := fishCountWithTimers[0]

		for timer := 0; timer <= maxTimerForFishes; timer++ {
			fishCountWithTimers[timer] = fishCountWithTimers[timer+1]
		}

		fishCountWithTimers[timerForFishesThatJustGaveBirth] += fishesThatCanGiveBirth
		fishCountWithTimers[maxTimerForFishes] += fishesThatCanGiveBirth
	}

	sumOfFishes := 0
	for _, fishCount := range fishCountWithTimers {
		sumOfFishes += fishCount
	}

	return sumOfFishes
}

func main() {
	elements, err := input.GetInput()
	if err != nil {
		fmt.Errorf("Error while parsing input: %+v", err)
	}

	res := calculateNumberOfFishes(elements, 80)
	fmt.Println("result day1 part1:", res)

	res = calculateNumberOfFishes(elements, 256)
	fmt.Println("result day1 part2:", res)
}
