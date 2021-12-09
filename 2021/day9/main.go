package main

import (
	"fmt"
	"strings"

	"advent_of_code/utils"
)

func findLowPoint(elements [][]int) int {
	var result []int
	sum := 0

	rowLength := len(elements)
	colLength := len(elements[0])

	for i := 0; i < rowLength; i++ {
		for j := 0; j < colLength; j++ {
			if i != 0 &&
				j != 0 &&
				i+1 < rowLength &&
				j+1 < colLength &&
				elements[i][j] < elements[i][j-1] &&
				elements[i][j] < elements[i][j+1] &&
				elements[i][j] < elements[i-1][j] &&
				elements[i][j] < elements[i+1][j] {

				result = append(result, elements[i][j])
			} else if i != 0 &&
				j == 0 &&
				i+1 < rowLength &&
				j+1 < colLength &&
				elements[i][j] < elements[i][j+1] &&
				elements[i][j] < elements[i-1][j] &&
				elements[i][j] < elements[i+1][j] {
				result = append(result, elements[i][j])
			} else if i == 0 &&
				j != 0 &&
				i+1 < rowLength &&
				j+1 < colLength &&
				elements[i][j] < elements[i][j-1] &&
				elements[i][j] < elements[i][j+1] &&
				elements[i][j] < elements[i+1][j] {
				result = append(result, elements[i][j])
			} else if i != 0 &&
				j != 0 &&
				i == rowLength-1 &&
				j+1 < colLength &&
				elements[i][j] < elements[i][j+1] &&
				elements[i][j] < elements[i-1][j] &&
				elements[i][j] < elements[i][j-1] {
				result = append(result, elements[i][j])
			} else if i != 0 &&
				j != 0 &&
				i+1 < rowLength &&
				j+1 == colLength &&
				elements[i][j] < elements[i][j-1] &&
				elements[i][j] < elements[i+1][j] &&
				elements[i][j] < elements[i-1][j] {
				result = append(result, elements[i][j])
			} else if i == 0 &&
				j == 0 &&
				i+1 < rowLength &&
				j+1 < colLength &&
				elements[i][j] < elements[i+1][j] &&
				elements[i][j] < elements[i][j+1] {
				result = append(result, elements[i][j])
			} else if i != 0 && j != 0 &&
				i+1 >= rowLength && j+1 >= colLength &&
				elements[i][j] < elements[i-1][j] &&
				elements[i][j] < elements[i][j-1] {
				result = append(result, elements[i][j])
			} else if i == 0 && j != 0 &&
				i+1 < rowLength && j+1 >= colLength &&
				elements[i][j] < elements[i+1][j] &&
				elements[i][j] < elements[i][j-1] {
				result = append(result, elements[i][j])
			} else if i != 0 && j == 0 &&
				i+1 >= rowLength && j+1 < colLength &&
				elements[i][j] < elements[i-1][j] &&
				elements[i][j] < elements[i][j+1] {
				result = append(result, elements[i][j])
			}

		}
	}

	for _, val := range result {
		sum += val + 1
	}

	return sum
}

func main() {
	matrix, err := getInput()
	if err != nil {
		fmt.Errorf("Error while parsing input: %+v", err)
	}

	res := findLowPoint(matrix)
	fmt.Println("result day1 part1:", res)
}

func getInput() ([][]int, error) {
	filePath := "/Users/ashwitha/GolandProjects/advent_of_code/2021/day9/input1.txt"

	lines, err := utils.ReadLines(filePath)
	if err != nil {
		return nil, err
	}

	var numbers [][]int
	for _, line := range lines {
		values := strings.Split(line, "")
		var row []int
		for _, value := range values {
			rowValue := utils.ToInt(value)
			row = append(row, rowValue)
		}
		numbers = append(numbers, row)
	}
	return numbers, nil
}
