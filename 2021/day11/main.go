package main

import (
	"advent_of_code/2021/day11/input"
	"advent_of_code/2021/day11/input/models"
	"fmt"
)

func main() {
	octopuses, err := input.GetInput()
	if err != nil {
		fmt.Errorf("Error while parsing input: %+v", err)
	}
	res := day11(octopuses)
	fmt.Println("result day11 part1:", res)

	res = day11Part2(octopuses)
	fmt.Println("result day11 part2:", res)
}

func day11(octopuses [10][10]models.Octopus) int {
	flashCounter := 0
	for k := 0; k < 100; k++ {
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				octopuses[i][j].IsFlashing = false
				(&octopuses[i][j]).Energize()
			}
		}

		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if (octopuses)[i][j].EnergyLevel > 9 {
					flashOctopus(&octopuses, &flashCounter, i, j)
				}
			}
		}
	}

	return flashCounter
}

func day11Part2(octopuses [10][10]models.Octopus) int {
	flashCounter := 0
	for k := 0; k < 600; k++ {
		flashStepCounter := flashCounter
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				octopuses[i][j].IsFlashing = false
				octopuses[i][j].EnergyLevel++
			}
		}

		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if (octopuses)[i][j].EnergyLevel > 9 {
					flashOctopus(&octopuses, &flashCounter, i, j)
				}
			}
		}

		if flashCounter-flashStepCounter == 100 {
			return k + 1
		}
	}

	return 0
}

func flashOctopus(octopuses *[10][10]models.Octopus, flashCounter *int, i int, j int) {
	(*flashCounter)++
	octopuses[i][j].Flash()
	octopuses[i][j].Reset()


	var moves = [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}, {-1, -1}, {-1, 1}, {1, 1}, {1, -1}}

	for _, move := range moves {
		adjacentRow := i + move[0]
		adjacentCol := j + move[1]

		if isValidBoundary(adjacentRow, adjacentCol) && !octopuses[adjacentRow][adjacentCol].IsFlashing {
			(*octopuses)[adjacentRow][adjacentCol].EnergyLevel++

			if octopuses[adjacentRow][adjacentCol].EnergyLevel > 9 {
				flashOctopus(octopuses, flashCounter, adjacentRow, adjacentCol)
			}
		}
	}
}

func isValidBoundary(row int, column int) bool {
	return row >= 0 && column >= 0 && row < 10 && column < 10
}
