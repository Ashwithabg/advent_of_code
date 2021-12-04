package models

import "fmt"

type Board struct {
	Cells [][]Cell
}

const CellCount = 5

func (board *Board) markCell(cellRowIndex int, cellIndex int) {
	board.Cells[cellRowIndex][cellIndex].markDone()
}

func (board Board) display() {
	for _, cellRows := range board.Cells {
		for _, cell := range cellRows {
			cell.display()
		}

		fmt.Println("")
	}
}

func (board Board) calculateScore() int {
	sum := 0
	for _, rowCells := range board.Cells {
		for _, cell := range rowCells {
			if cell.Value != -1 {
				sum += cell.Value
			}
		}
	}

	return sum
}

func (board *Board) doesPlayerWinsFor(input int) bool {
	for cellRowIndex, cellRow := range board.Cells {
		for cellIndex, cell := range cellRow {
			if cell.isValue(input) {
				board.markCell(cellRowIndex, cellIndex)
			}
		}

		if board.isDone() {
			return true
		}
	}

	return false
}

func (board *Board) isRowMarked() bool {
	for _, cellRow := range board.Cells {
		markedRowCount := 0
		for _, cell := range cellRow {
			if cell.isMarked() {
				markedRowCount++
			}

			if board.doesAllCellsTraversed(markedRowCount) {
				return true
			}
		}
	}

	return false
}

func (board *Board) isColumnMarked() bool {
	for rowIndex := 0; rowIndex < CellCount; rowIndex++ {
		markedColumnCount := 0

		for columnIndex := 0; columnIndex < CellCount; columnIndex++ {
			if board.Cells[columnIndex][rowIndex].isMarked() {
				markedColumnCount++
			}

			if board.doesAllCellsTraversed(markedColumnCount) {
				return true
			}
		}
	}

	return false
}

func (board *Board) doesAllCellsTraversed(columnWinnerCount int) bool {

	return columnWinnerCount == CellCount
}

func (board Board) isDone() bool {
	return board.isRowMarked() || board.isColumnMarked()
}

func (board *Board) mark(input int) {
	for cellRowIndex, cellRow := range board.Cells {
		for cellIndex, cell := range cellRow {
			if cell.Value == input {
				board.Cells[cellRowIndex][cellIndex].markDone()
			}
		}
	}
}
