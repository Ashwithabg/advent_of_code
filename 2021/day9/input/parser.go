package input

import (
	"strings"

	"advent_of_code/2021/day9/input/models"
	"advent_of_code/utils"
)

func GetInput() (models.HeightMap, error) {
	filePath := "/Users/ashwitha/GolandProjects/advent_of_code/2021/day9/input/files/input1.txt"

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
