package main

import (
	"advent_of_code/utils"
	"fmt"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type position struct {
	startPoint point
	endPoint   point
}

type Grid [1000][1000]int

func checkOverlapsInRowAndColumnOfGrid(grid Grid) int {
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

func main() {
	positions, err := getPositions()
	if err != nil {
		fmt.Errorf("Error while parsing positions: %+v", err)
	}

	grid := constructGridRowAndColumnwise(positions)
	res := checkOverlapsInRowAndColumnOfGrid(grid)
	fmt.Println("result day5 part1:", res)
}

func constructGridRowAndColumnwise(positions []position) Grid {
	var grid Grid
	for _, position := range positions {
		startPoint := position.startPoint
		endPoint := position.endPoint

		if startPoint.x == endPoint.x {
			if startPoint.y <= endPoint.y {
				for i := startPoint.y; i <= endPoint.y; i++ {
					grid[endPoint.x][i]++
				}
			} else {
				for i := endPoint.y; i <= startPoint.y; i++ {
					grid[endPoint.x][i]++
				}
			}

		} else if startPoint.y == endPoint.y {
			if startPoint.x <= endPoint.x {
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

func getPositions() ([]position, error) {
	filePath := "/Users/ashwitha/GolandProjects/advent_of_code/2021/day5/input.txt"

	lines, err := utils.ReadLines(filePath)
	if err != nil {
		return []position{}, err
	}

	positions := []position{}
	for _, line := range lines {
		coordinateStrings := strings.Split(line, " -> ")
		firstCoordinate := strings.Split(coordinateStrings[0], ",")
		secondCoordinate := strings.Split(coordinateStrings[1], ",")

		x1, _ := strconv.Atoi(firstCoordinate[0])
		y1, _ := strconv.Atoi(firstCoordinate[1])

		x2, _ := strconv.Atoi(secondCoordinate[0])
		y2, _ := strconv.Atoi(secondCoordinate[1])

		position := position{
			startPoint: point{x: x1, y: y1},
			endPoint:   point{x: x2, y: y2},
		}
		positions = append(positions, position)
	}

	return positions, nil
}
