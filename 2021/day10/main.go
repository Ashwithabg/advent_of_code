package main

import (
	"fmt"

	"advent_of_code/2021/day10/input"
)

var pointsMap = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

func main() {
	matrix, err := input.GetInput()
	if err != nil {
		fmt.Errorf("Error while parsing input: %+v", err)
	}

	res := findCorruptedLinePoints(matrix)
	fmt.Println("result day9 part1:", res)
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
		if value == "(" || value == "{" || value == "[" || value == "<" {
			stack.Push(value)
		} else {
			poppedValue, _ := stack.Pop()
			if poppedValue == "(" && value != ")" ||
				poppedValue == "[" && value != "]" ||
				poppedValue == "{" && value != "}" ||
				poppedValue == "<" && value != ">" {
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
