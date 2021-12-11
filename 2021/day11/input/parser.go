package input

import (
	"advent_of_code/2021/day11/input/models"
	"advent_of_code/utils"
	"strings"
)

func GetInput() ([10][10]models.Octopus, error) {
	filePath := "/Users/ashwitha/GolandProjects/advent_of_code/2021/day11/input/files/input1.txt"
	var octopuses [10][10]models.Octopus
	lines, err := utils.ReadLines(filePath)
	if err != nil {
		return octopuses, err
	}

	for j, line := range lines {
		values := strings.Split(line, "")
		var row [10]models.Octopus

		for i, value := range values {
			row[i] = models.Octopus{
				EnergyLevel: utils.ToInt(value),
			}
		}
		octopuses[j] = row
	}

	return octopuses, nil
}
