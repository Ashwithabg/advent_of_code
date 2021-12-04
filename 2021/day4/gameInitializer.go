package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getInput(line string) ([]int, error) {
	values := strings.Split(line, ",")

	var inputs []int
	for _, v := range values {
		val, err := strconv.Atoi(v)
		if err != nil {
			return inputs, err
		}
		inputs = append(inputs, val)
	}
	fmt.Printf("inputs  %+v \n", inputs)
	return inputs, nil
}

func getBoards(rawData []string) (Game, error) {
	var game Game
	var board Board
	boardRowIndex := 0
	boardIndex := 0

	for index := 2; index < len(rawData); index++ {
		if rawData[index] != "" {
			var cells []Cell

			rowValues := strings.Fields(rawData[index])
			for _, rowValue := range rowValues {
				convertedValue, err := strconv.Atoi(rowValue)
				if err != nil {
					return game, err
				}

				cells = append(cells, Cell{
					Value: convertedValue,
				})
			}

			board.Cells = append(board.Cells, cells)
			boardRowIndex++
		} else {
			game.boards = append(game.boards, board)
			boardIndex++
			boardRowIndex = 0
			board = Board{}
		}
	}

	game.boards = append(game.boards, board)
	return game, nil
}
