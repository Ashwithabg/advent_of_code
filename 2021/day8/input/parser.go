package input

import (
	"advent_of_code/utils"
)

func GetInput() []string {
	filePath := "/Users/ashwitha/GolandProjects/advent_of_code/2021/day8/input/files/input1.txt"

	lines, err := utils.ReadLines(filePath)
	if err != nil {
		panic("can not read input")
	}
	return lines
}
