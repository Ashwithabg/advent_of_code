package models

import (
	"fmt"
	"sort"
)

type Game struct {
	Boards []Board
}

func (game *Game) CheckBoardThatWinsFirstFor(input int) (bool, int) {
	for _, board := range game.Boards {
		if board.playerWinsFor(input) {
			return true, board.calculateScore()
		}
	}

	return false, 0
}

func (game *Game) markCell(boardIndex int, cellRowIndex int, cellIndex int) {
	game.Boards[boardIndex].markCell(cellRowIndex, cellIndex)
}

func (game Game) Display() {
	for _, board := range game.Boards {
		board.display()
		fmt.Println("")
	}
}

func (game *Game) findBoard(index int) Board {
	return game.Boards[index]
}

func (game *Game) CheckBoardThatWinsLastFor(input int) (bool, int) {
	var boardIndexesToDelete []int

	for index, board := range game.Boards {

		for cellRowIndex, cellRow := range board.Cells {
			for cellIndex, cell := range cellRow {
				if cell.Value == input {
					game.Boards[index].Cells[cellRowIndex][cellIndex].markDone()
				}
			}
		}

		if len(game.Boards) == 1 && (board.isRowMarked() || board.isColumnMarked()){
			return true, board.calculateScore()
		}

		if board.isRowMarked() || board.isColumnMarked() {
			boardIndexesToDelete = append(boardIndexesToDelete, index)
		}
	}

	sort.Ints(boardIndexesToDelete)
	for index, value := range boardIndexesToDelete {
		game.RemoveBoard(value-index)
	}

	return false, 0
}

func (game *Game) RemoveBoard(index int) {
	game.Boards = append(game.Boards[:index], game.Boards[index+1:]...)
}
