package input

import (
	"advent_of_code/2021/day5/models"
	"advent_of_code/utils"
	"strconv"
	"strings"
)

func GetPositions() ([]models.Position, error) {
	filePath := "/Users/ashwitha/GolandProjects/advent_of_code/2021/day5/input/files/input1.txt"

	lines, err := utils.ReadLines(filePath)
	if err != nil {
		return []models.Position{}, err
	}

	positions := []models.Position{}
	for _, line := range lines {
		coordinateStrings := strings.Split(line, " -> ")
		firstCoordinate := strings.Split(coordinateStrings[0], ",")
		secondCoordinate := strings.Split(coordinateStrings[1], ",")

		x1, _ := strconv.Atoi(firstCoordinate[0])
		y1, _ := strconv.Atoi(firstCoordinate[1])

		x2, _ := strconv.Atoi(secondCoordinate[0])
		y2, _ := strconv.Atoi(secondCoordinate[1])

		position := models.NewPosition(x1, y1, x2, y2)
		positions = append(positions, position)
	}

	return positions, nil
}
