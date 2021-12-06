package main

import (
	"fmt"
)

func calculateNumberOfFishes(fishTimers []int, dayCount int) int {
	const maxTimerForFishes = 8
	fishCountWithTimers := make(map[int]int)

	for _, fishTimer := range fishTimers {
		fishCountWithTimers[fishTimer] += 1
	}

	for dayIndex := 0; dayIndex < dayCount; dayIndex++ {
		newlyCreatedFishCountWithTimers := make(map[int]int)

		for timer := 0; timer <= maxTimerForFishes; timer++ {
			noOfFishes := fishCountWithTimers[timer]

			if timer == 0 {
				delete(fishCountWithTimers, 0)
				newlyCreatedFishCountWithTimers[6] += noOfFishes
				newlyCreatedFishCountWithTimers[8] += noOfFishes
			} else {
				fishCountWithTimers[timer] -= noOfFishes
				fishCountWithTimers[timer-1] += noOfFishes
			}
		}

		for timer, numberOfFishes := range newlyCreatedFishCountWithTimers {
			fishCountWithTimers[timer] += numberOfFishes
		}
	}

	sumOfFishes := 0
	for _, fishCount := range fishCountWithTimers {
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
