package models

import (
	"fmt"
)

type Grid [1000][1000]int

func NewGridWithRowAndColumnwise(positions []Position) Grid {
	var grid Grid
	for _, position := range positions {
		startPoint := position.startPoint
		endPoint := position.endPoint

		if startPoint.hasSameXCoordinateAs(endPoint) {
			if startPoint.isYCoordinateLesserOrEqualTo(endPoint) {
				for i := startPoint.y; i <= endPoint.y; i++ {
					grid[endPoint.x][i]++
				}
			} else {
				for i := endPoint.y; i <= startPoint.y; i++ {
					grid[endPoint.x][i]++
				}
			}

		} else if startPoint.hasSameYCoordinateAs(endPoint) {
			if startPoint.isXCoordinateLesserOrEqualTo(endPoint) {
				for i := startPoint.x; i <= endPoint.x; i++ {
					grid[i][startPoint.y]++
				}
			} else {
				for i := endPoint.x; i <= startPoint.x; i++ {
					grid[i][startPoint.y]++
				}
			}
		}
	}

	return grid
}

func NewGridWithAllThreeDirections(positions []Position) Grid {
	var grid Grid
	for _, position := range positions {
		startPoint := position.startPoint
		endPoint := position.endPoint

		if startPoint.hasSameXCoordinateAs(endPoint) {
			if startPoint.isYCoordinateLesserOrEqualTo(endPoint) {
				for i := startPoint.y; i <= endPoint.y; i++ {
					grid[endPoint.x][i]++
				}
			} else {
				for i := endPoint.y; i <= startPoint.y; i++ {
					grid[endPoint.x][i]++
				}
			}

		} else if startPoint.hasSameYCoordinateAs(endPoint) {
			if startPoint.isXCoordinateLesserOrEqualTo(endPoint) {
				for i := startPoint.x; i <= endPoint.x; i++ {
					grid[i][startPoint.y]++
				}
			} else {
				for i := endPoint.x; i <= startPoint.x; i++ {
					grid[i][startPoint.y]++
				}
			}
		} else {
			if endPoint.isXCoordinateLesserOrEqualTo(startPoint) &&
				endPoint.isYCoordinateLesserOrEqualTo(startPoint) {

				for column, row := endPoint.y, endPoint.x; column <= startPoint.y; column, row = column+1, row+1 {
					grid[row][column]++
				}
			} else if endPoint.isXCoordinateLesserOrEqualTo(startPoint) &&
				endPoint.isYCoordinateGreaterThan(startPoint) {
				for i, row := endPoint.y, endPoint.x; i >= startPoint.y; i, row = i-1, row+1 {
					grid[row][i]++
				}
			} else if endPoint.isXCoordinateGreaterThan(startPoint) &&
				endPoint.isYCoordinateGreaterThan(startPoint) {
				for i, row := endPoint.y, endPoint.x; i >= startPoint.y; i, row = i-1, row-1 {
					grid[row][i]++
				}
			} else {
				for i, row := endPoint.y, endPoint.x; i <= startPoint.y; i, row = i+1, row-1 {
					grid[row][i]++
				}
			}
		}
	}

	return grid
}

func (grid Grid) display() {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			fmt.Printf("%+v", grid[i][j])
		}
		fmt.Println()
	}
	fmt.Println("-------------")
}

func (grid Grid) CheckOverlaps() int {
	counter := 0

	for xCoord := 0; xCoord < len(grid); xCoord++ {
		for yCoord := 0; yCoord < len(grid); yCoord++ {
			if grid[xCoord][yCoord] > 1 {
				counter++
			}
		}

	}

	return counter
}
