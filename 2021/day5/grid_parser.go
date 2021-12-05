package main

import (
	"advent_of_code/2021/day5/models"
	"fmt"
)

func main() {
	positions, err := getPositions()
	if err != nil {
		fmt.Errorf("Error while parsing positions: %+v", err)
	}

	grid := models.NewGridWithRowAndColumnwise(positions)
	res := grid.CheckOverlaps()
	fmt.Println("result day5 part1:", res)
}
