package input

import (
	"advent_of_code/utils"
	"strings"
)

func GetInput() ([10][10]int, error) {
	filePath := "/Users/ashwitha/GolandProjects/advent_of_code/2021/day11/input/files/input1.txt"
	var numbers [10][10]int
	lines, err := utils.ReadLines(filePath)
	if err != nil {
		return numbers, err
	}

	for j, line := range lines {
		values := strings.Split(line, "")
		var row [10]int
		for i, value := range values {
			row[i] = utils.ToInt(value)
		}
		numbers[j] = row
	}

	return numbers, nil
}
