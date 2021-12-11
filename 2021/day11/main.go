package main

import (
	"advent_of_code/2021/day11/input"
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

func day11(octopuses [10][10]int) int {
	flashCounter := 0
	for k := 0; k < 100; k++ {
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				octopuses[i][j]++
			}
		}

		var flashedOctopuses [10][10]bool
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if (octopuses)[i][j] > 9 {
					flashOctopus(&octopuses, &flashCounter, i, j, &flashedOctopuses)
				}
			}
		}
	}

	return flashCounter
}

func day11Part2(octopuses [10][10]int) int {
	flashCounter := 0
	for k := 0; k < 600; k++ {
		flashStepCounter := flashCounter
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				octopuses[i][j]++
			}
		}

		var flashedOctopuses [10][10]bool
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if (octopuses)[i][j] > 9 {
					flashOctopus(&octopuses, &flashCounter, i, j, &flashedOctopuses)
				}
			}
		}

		if flashCounter - flashStepCounter == 100 {
			return k + 1
		}
	}

	return 0
}

func flashOctopus(octopuses *[10][10]int, flashCounter *int, i int, j int, flashedOctopuses *[10][10]bool) {
	(*flashCounter)++
	octopuses[i][j] = 0
	flashedOctopuses[i][j] = true

	var moves = [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}, {-1, -1}, {-1, 1}, {1, 1}, {1, -1}}

	for _, move := range moves {
		adjacentRow := i + move[0]
		adjacentCol := j + move[1]

		if isValidBoundary(adjacentRow, adjacentCol) && (*flashedOctopuses)[adjacentRow][adjacentCol] != true {
			(*octopuses)[adjacentRow][adjacentCol]++

			if octopuses[adjacentRow][adjacentCol] > 9 {
				flashOctopus(octopuses, flashCounter, adjacentRow, adjacentCol, flashedOctopuses)
			}
		}
	}
}

func isValidBoundary(row int, column int) bool {
	return row >= 0 && column >= 0 && row < 10 && column < 10
}
