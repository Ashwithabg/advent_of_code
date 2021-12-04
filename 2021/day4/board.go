package main

import "fmt"

type Board struct {
	Cells [][]Cell
}

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

func (board *Board) playerWinsFor(input int) bool {
	for cellRowIndex, cellRow := range board.Cells {
		for cellIndex, cell := range cellRow {
			if cell.Value == input {
				board.markCell(cellRowIndex, cellIndex)
			}
		}

		if board.isRowMarked() || board.isColumnMarked() {
			return true
		}
	}

	return false
}

func (board *Board) isRowMarked() bool {
	for _, cellRow := range board.Cells {
		rowWinnerCount := 0
		for _, cell := range cellRow {
			if cell.isMarked() {
				rowWinnerCount++
			}

			if rowWinnerCount == 5 {
				return true
			}
		}
	}

	return false
}

func (board *Board) isColumnMarked() bool {
	for rowIndex := 0; rowIndex < 5; rowIndex++ {
		columnWinnerCount := 0
		for columnIndex := 0; columnIndex < 5; columnIndex++ {
			if board.Cells[columnIndex][rowIndex].isMarked() {
				columnWinnerCount++
			}

			if columnWinnerCount == 5 {
				return true
			}
		}
	}

	return false
}
