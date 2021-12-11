package models

type OctopusGrid [][]Octopus

var neighbourOctopusPoints = [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}, {-1, -1}, {-1, 1}, {1, 1}, {1, -1}}

func (og *OctopusGrid) EnergizeOctopuses() {
	for row := 0; row < 10; row++ {
		for column := 0; column < 10; column++ {
			(*og)[row][column].turnOffFlash()
			(*og)[row][column].Energize()
		}
	}
}

func (og *OctopusGrid) FlashEnergizedOctopuses(flashCounter *int) {
	for row := 0; row < 10; row++ {
		for column := 0; column < 10; column++ {
			if (*og)[row][column].isEnergized() {
				(*og).flash(flashCounter, row, column)
			}
		}
	}
}

func (og *OctopusGrid) flash(flashCounter *int, row int, column int) {
	*flashCounter++
	(*og)[row][column].Flash()
	(*og)[row][column].Reset()

	for _, point := range neighbourOctopusPoints {
		adjacentRow := row + point[0]
		adjacentCol := column + point[1]

		if og.isValidGridBoundary(adjacentRow, adjacentCol) &&
			!(*og)[adjacentRow][adjacentCol].IsFlashing {
			(*og)[adjacentRow][adjacentCol].EnergyLevel++

			if (*og)[adjacentRow][adjacentCol].isEnergized() {
				(*og).flash(flashCounter, adjacentRow, adjacentCol)
			}
		}
	}
}

func (og OctopusGrid) isValidGridBoundary(row int, column int) bool {
	return row >= 0 &&
		column >= 0 &&
		row < len(og) &&
		column < len(og[0])
}
