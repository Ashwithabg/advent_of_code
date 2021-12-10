package main

import (
	"fmt"
	"sort"
	"sync"

	"advent_of_code/2021/day10/input"
	"advent_of_code/utils"
)

var pointsMap = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

var ClosingPointsMap = map[string]int{
	"(": 1,
	"[": 2,
	"{": 3,
	"<": 4,
}

var brackets = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

var wg sync.WaitGroup

func main() {
	matrix, err := input.GetInput()
	if err != nil {
		fmt.Errorf("Error while parsing input: %+v", err)
	}
	wg.Add(2)
	go func() {
		defer wg.Done()
		res := findCorruptedLinesScore(matrix)
		fmt.Println("result day9 part1:", res)

	}()

	go func() {
		defer wg.Done()
		res := findIncompleteLinesScore(matrix)
		fmt.Println("result day9 part2:", res)
	}()

	defer wg.Wait()
}

func findCorruptedLinesScore(lines [][]string) int {
	sum := 0
	for _, line := range lines {
		sum += findCorruptedScoreFor(line)
	}
	return sum
}

func findCorruptedScoreFor(line []string) int {
	var bracketTracker utils.Stack
	score := 0

	for _, value := range line {
		if isOpeningBraces(value) {
			bracketTracker.Push(value)
		} else {
			poppedValue, _ := bracketTracker.Pop()
			if isCorrupted(poppedValue, value) {
				score += pointsMap[value]
			}
		}
	}

	return score
}

func findIncompleteLinesScore(lines [][]string) int {
	var incompleteLines []int
	for _, line := range lines {
		value := getIncompleteLineScoreFor(line)
		if value != 0 {
			incompleteLines = append(incompleteLines, value)
		}
	}

	sort.Ints(incompleteLines)
	return incompleteLines[len(incompleteLines)/2]
}

func getIncompleteLineScoreFor(line []string) int {
	var stack utils.Stack

	for _, bracket := range line {
		if isOpeningBraces(bracket) {
			stack.Push(bracket)
		} else {
			openingBracket, _ := stack.Pop()
			if isCorrupted(openingBracket, bracket) {
				return 0
			}
		}
	}

	score := 0
	for stack.Size() != 0 {
		poppedValue, _ := stack.Pop()
		score *= 5
		score += ClosingPointsMap[poppedValue]
	}

	return score
}

func isOpeningBraces(key string) bool {
	_, ok := brackets[key]
	return ok
}

func isCorrupted(openingBracket string, value string) bool {
	closingBracket, _ := brackets[openingBracket]
	return value != closingBracket
}
