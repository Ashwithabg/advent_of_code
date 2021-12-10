package main

import (
	"fmt"
	"sort"

	"advent_of_code/2021/day10/input"
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

var ClosingBraces = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

func main() {
	matrix, err := input.GetInput()
	if err != nil {
		fmt.Errorf("Error while parsing input: %+v", err)
	}

	res := findCorruptedLinePoints(matrix)
	fmt.Println("result day9 part1:", res)

	res = findIncomepleteLinePoints(matrix)
	fmt.Println("result day9 part2:", res)
}

func findCorruptedLinePoints(paranethesisLines [][]string) int {
	sum := 0
	for _, paranethesisLine := range paranethesisLines {
		sum += findValue(paranethesisLine)
	}
	return sum
}

func findValue(line []string) int {
	var stack Stack
	var corruptedValue []string

	for _, value := range line {
		if isOpeningBraces(value) {
			stack.Push(value)
		} else {
			poppedValue, _ := stack.Pop()
			if isCorrupted(poppedValue, value) {
				corruptedValue = append(corruptedValue, value)
			}
		}
	}

	sum := 0
	for _, cv := range corruptedValue {
		sum += pointsMap[cv]
	}

	return sum
}

func findIncomepleteLinePoints(paranethesisLines [][]string) int {
	var values []int
	for _, paranethesisLine := range paranethesisLines {
		value := findIncompleteLineValue(paranethesisLine)
		values = append(values, value)
	}
	fmt.Printf("%+v\n", values)
	sort.Ints(values)
	var data []int
	for _, val := range values {
		if val != 0{
			data = append(data, val)
		}
	}

	return data[len(data)/2]
}

func findIncompleteLineValue(line []string) int {
	var stack Stack
	for _, value := range line {
		if isOpeningBraces(value) {
			stack.Push(value)
		} else {
			poppedValue, ok := stack.Pop()
			if ok {
				if isCorrupted(poppedValue, value) {
					return 0
				}
			}
		}
	}

	sum := 0
	for stack.Size() != 0 {
		poppedValue, _ := stack.Pop()
		sum *= 5
		sum += ClosingPointsMap[poppedValue]
	}
	return sum
}

func isOpeningBraces(value string) bool {
	return value == "(" || value == "{" || value == "[" || value == "<"
}

func isCorrupted(poppedValue string, value string) bool {
	return poppedValue == "(" && value != ")" ||
		poppedValue == "[" && value != "]" ||
		poppedValue == "{" && value != "}" ||
		poppedValue == "<" && value != ">"
}
