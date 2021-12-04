package main

import "fmt"

type Game struct {
	boards []Board
}

func (game *Game) playFor(input int) (bool, int) {
	for _, board := range game.boards {
		if board.playerWinsFor(input) {
			return true, board.calculateScore()
		}
	}

	return false, 0
}

func (game *Game) markCell(boardIndex int, cellRowIndex int, cellIndex int) {
	game.boards[boardIndex].markCell(cellRowIndex, cellIndex)
}

func (game Game) display() {
	for _, board := range game.boards {
		board.display()
		fmt.Println("")
	}
}

func (game *Game) findBoard(index int) Board {
	return game.boards[index]
}
