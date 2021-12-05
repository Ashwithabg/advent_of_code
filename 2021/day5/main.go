package main

import (
	"advent_of_code/2021/day5/input"
	"advent_of_code/2021/day5/models"
	"fmt"
)

func main() {
	positions, err := input.GetPositions()
	if err != nil {
		fmt.Errorf("Error while parsing positions: %+v", err)
	}

	grid := models.NewGridWithRowAndColumnwise(positions)
	res := grid.CheckOverlaps()
	fmt.Println("result day5 part1:", res)

	secondGrid := models.NewGridWithAllThreeDirections(positions)
	res = secondGrid.CheckOverlaps()
	fmt.Println("result day5 part2:", res)
}
