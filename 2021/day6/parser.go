package main

import (
	"advent_of_code/utils"
	"strconv"
	"strings"
)

func getInput() ([]int, error) {
	filePath := "/Users/ashwitha/GolandProjects/advent_of_code/2021/day6/input1.txt"

	lines, err := utils.ReadLines(filePath)
	if err != nil {
		return nil, err
	}

	internalTimers := strings.Split(lines[0], ",")
	var numbers []int
	for _, element := range internalTimers {
		n, _ := strconv.Atoi(element)
		numbers = append(numbers, n)
	}

	return numbers, nil
}
