package models

import (
	"fmt"
)

type Game struct {
	Boards []Board
}

func (game *Game) CheckBoardThatWinsFirstFor(input int) (bool, int) {
	for _, board := range game.Boards {
		if board.doesPlayerWinsFor(input) {
			return true, board.calculateScore()
		}
	}

	return false, 0
}

func (game Game) Display() {
	for _, board := range game.Boards {
		board.display()
		fmt.Println("")
	}
}

func (game *Game) CheckBoardThatWinsLastFor(input int) (bool, int) {
	var boardIndexesToDelete []int

	for index, board := range game.Boards {
		board.mark(input)

		if game.hasOnlyOneBoard() && board.isDone() {
			return true, board.calculateScore()
		}

		if board.isDone() {
			boardIndexesToDelete = append(boardIndexesToDelete, index)
		}
	}

	for index, value := range boardIndexesToDelete {
		game.RemoveBoard(value - index)
	}

	return false, 0
}

func (game *Game) hasOnlyOneBoard() bool {
	return len(game.Boards) == 1
}

func (game *Game) RemoveBoard(index int) {
	game.Boards = append(game.Boards[:index], game.Boards[index+1:]...)
}
