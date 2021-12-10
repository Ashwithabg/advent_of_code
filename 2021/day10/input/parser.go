package input

import (
	"strings"

	"advent_of_code/utils"
)

func GetInput() ([][]string, error) {
	filePath := "/Users/ashwitha/GolandProjects/advent_of_code/2021/day10/input/files/input1.txt"

	lines, err := utils.ReadLines(filePath)
	if err != nil {
		return nil, err
	}

	var values [][]string
	for _, line := range lines {
		val := strings.Split(line, "")
		values = append(values, val)
	}

	return values, nil
}
