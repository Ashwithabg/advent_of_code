package main

import (
	"fmt"

	"advent_of_code/utils"
)

func main() {
	filePath := "/Users/ashwitha/GolandProjects/advent_of_code/2021/day4/input.txt"
	lines, err := utils.ReadLines(filePath)
	if err != nil {
		fmt.Errorf("Error while reading file: %+v\n", err)
		return
	}

	inputs, err := getInput(lines[0])
	if err != nil {
		fmt.Errorf("Error while parsing input: %+v\n", err)
		return
	}

	game, err := getBoards(lines)
	if err != nil {
		fmt.Errorf("Error while parsing game: %+v\n", err)
		return
	}

	game.display()

	for _, input := range inputs {
		isGameOver, score := game.playFor(input)
		if isGameOver {
			fmt.Println("day4 part1 score: ", score * input)
			break
		}
	}

	fmt.Println("Game over")
}
