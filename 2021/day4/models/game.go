package models

import "fmt"

type Game struct {
	Boards []Board
}

func (game *Game) PlayFor(input int) (bool, int) {
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
