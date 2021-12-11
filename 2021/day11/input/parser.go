package input

import (
	"strings"

	"advent_of_code/2021/day11/input/models"
	"advent_of_code/utils"
)

func GetOctopusGrid() (models.OctopusGrid, error) {
	filePath := "/Users/ashwitha/GolandProjects/advent_of_code/2021/day11/input/files/input1.txt"
	var octopuses models.OctopusGrid
	lines, err := utils.ReadLines(filePath)
	if err != nil {
		return octopuses, err
	}

	for _, line := range lines {
		values := strings.Split(line, "")
		var OctopusRow []models.Octopus

		for _, value := range values {
			OctopusRow = append(OctopusRow, models.Octopus{
				EnergyLevel: utils.ToInt(value),
			})
		}
		octopuses = append(octopuses, OctopusRow)
	}

	return octopuses, nil
}
