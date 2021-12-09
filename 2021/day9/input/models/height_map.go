package models

import "sort"

type HeightMap [][]int

var moves = [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

const border = 9

func (heightMap HeightMap) FindSizeOfThreeLargestBasins() int {
	basins := heightMap.findBasins()
	sort.Ints(basins)

	size := 1
	for i := len(basins) - 1; i > len(basins)-4; i-- {
		size *= basins[i]
	}

	return size
}

func (heightMap HeightMap) FindRiskLevel() int {
	riskLevel := 0
	lowPoints := heightMap.findLowPoints()

	for _, lowPoint := range lowPoints {
		riskLevel += lowPoint + 1
	}

	return riskLevel
}

func (heightMap HeightMap) findBasins() []int {
	var locationCounter int
	var basins []int

	for i := 0; i < len(heightMap); i++ {
		for j := 0; j < len(heightMap[0]); j++ {
			(&heightMap).findBasinCount(i, j, &locationCounter)
			basins = append(basins, locationCounter)
			locationCounter = 0
		}
	}

	return basins
}

func (heightMap HeightMap) findLowPoints() []int {
	var lowPoints []int

	for row := 0; row < len(heightMap); row++ {
		for column := 0; column < len(heightMap[0]); column++ {
			if heightMap.isLow(row, column) {
				lowPoints = append(lowPoints, heightMap[row][column])
			}
		}
	}
	return lowPoints
}

func (heightMap HeightMap) isLow(row int, col int) bool {
	for _, move := range moves {
		adjacentRow := row + move[0]
		adjacentCol := col + move[1]

		if heightMap.isValidBoundary(adjacentRow, adjacentCol) {
			if heightMap[row][col] >= heightMap[adjacentRow][adjacentCol] {
				return false
			}
		}
	}

	return true
}

func (heightMap HeightMap) isValidBasinBoundary(row int, column int) bool {
	return heightMap.isValidBoundary(row, column) && heightMap[row][column] != border
}

func (heightMap HeightMap) isValidBoundary(row int, column int) bool {
	return row >= 0 && column >= 0 && row < len(heightMap) && column < len(heightMap[0])
}

func (heightMap *HeightMap) findBasinCount(row int, column int, locationCounter *int) {
	if heightMap.isValidBasinBoundary(row, column) {
		(*heightMap)[row][column] = border
		(*locationCounter) += 1

		for i := 0; i < len(moves); i++ {
			adjacentRow := row + moves[i][0]
			adjacentColumn := column + moves[i][1]
			heightMap.findBasinCount(adjacentRow, adjacentColumn, locationCounter)
		}
	}
}
