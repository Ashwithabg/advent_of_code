package models

type OctopusGrid [][]Octopus

var neibourOctopusPoints = [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}, {-1, -1}, {-1, 1}, {1, 1}, {1, -1}}

func (og *OctopusGrid) EnergizeOctopuses() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			(*og)[i][j].turnOffFlash()
			(*og)[i][j].Energize()
		}
	}
}

func (og *OctopusGrid) FlashEnergizedOctopuses(flashCounter *int) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if (*og)[i][j].isEnergized() {
				(*og).flash(flashCounter, i, j)
			}
		}
	}
}

func (og *OctopusGrid) flash(flashCounter *int, row int, column int) {
	*flashCounter++
	(*og)[row][column].Flash()
	(*og)[row][column].Reset()

	for _, point := range neibourOctopusPoints {
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
