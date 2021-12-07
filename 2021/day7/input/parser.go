package input

import (
	"strings"

	"advent_of_code/utils"
)

func GetCrabSubmarineHorizontalPositions() ([]int, error) {
	filePath := "/Users/ashwitha/GolandProjects/advent_of_code/2021/day7/input/files/input1.txt"

	lines, err := utils.ReadLines(filePath)
	if err != nil {
		return nil, err
	}

	internalTimers := strings.Split(lines[0], ",")
	return utils.ToIntSlice(internalTimers), nil
}
