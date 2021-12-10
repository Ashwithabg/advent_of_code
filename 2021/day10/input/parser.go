package input

import (
	"strings"

	"advent_of_code/2021/day10/models"
	"advent_of_code/utils"
)

func GetInput() ([]models.BracketLine, error) {
	filePath := "/Users/ashwitha/GolandProjects/advent_of_code/2021/day10/input/files/input1.txt"

	lines, err := utils.ReadLines(filePath)
	if err != nil {
		return nil, err
	}

	var brackets []models.BracketLine
	for _, line := range lines {
		values := strings.Split(line, "")

		var bracketLine models.BracketLine
		for _, value := range values {
			bracketLine = append(bracketLine, models.Bracket(value))
		}
		brackets = append(brackets, bracketLine)
	}

	return brackets, nil
}
