package input

import (
	"advent_of_code/2021/day4/models"

	"fmt"
	"strconv"
	"strings"
)

func GetInput(line string) ([]int, error) {
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

func GetBoards(rawData []string) (models.Game, error) {
	var game models.Game
	var board models.Board
	boardRowIndex := 0
	boardIndex := 0

	for index := 2; index < len(rawData); index++ {
		if rawData[index] != "" {
			var cells []models.Cell

			rowValues := strings.Fields(rawData[index])
			for _, rowValue := range rowValues {
				convertedValue, err := strconv.Atoi(rowValue)
				if err != nil {
					return game, err
				}

				cells = append(cells, models.Cell{
					Value: convertedValue,
				})
			}

			board.Cells = append(board.Cells, cells)
			boardRowIndex++
		} else {
			game.Boards = append(game.Boards, board)
			boardIndex++
			boardRowIndex = 0
			board = models.Board{}
		}
	}

	game.Boards = append(game.Boards, board)
	return game, nil
}
