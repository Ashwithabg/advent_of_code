package main

import (
	"fmt"
	"sort"
	"sync"

	"advent_of_code/2021/day10/input"
	"advent_of_code/2021/day10/models"
)

var wg sync.WaitGroup

func main() {
	brackets, err := input.GetInput()
	if err != nil {
		fmt.Errorf("Error while parsing input: %+v", err)
	}
	wg.Add(2)
	go func() {
		defer wg.Done()
		res := findCorruptedLinesScore(brackets)
		fmt.Println("result day9 part1:", res)

	}()

	go func() {
		defer wg.Done()
		res := findIncompleteLinesScore(brackets)
		fmt.Println("result day9 part2:", res)
	}()

	defer wg.Wait()
}

func findCorruptedLinesScore(lines []models.BracketLine) int {
	sum := 0
	for _, line := range lines {
		sum += line.FindCorruptedScore()
	}
	return sum
}

func findIncompleteLinesScore(lines []models.BracketLine) int {
	var incompleteLines []int
	for _, line := range lines {
		value := line.GetIncompleteLineScore()
		if value != 0 {
			incompleteLines = append(incompleteLines, value)
		}
	}

	sort.Ints(incompleteLines)
	return incompleteLines[len(incompleteLines)/2]
}
