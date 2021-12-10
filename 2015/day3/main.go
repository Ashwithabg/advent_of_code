package main

import (
	"fmt"
	"strings"
	"sync"

	"advent_of_code/utils"
)

var wg sync.WaitGroup

type Point struct {
	x int
	y int
}

func findNumberOfHouses(inputs []string) int {
	row := 0
	col := 0
	visitedHouses := make(map[Point]int)
	visitedHouses[Point{row, col}] = 1

	for i := 0; i < len(inputs); i++ {
		findNextPoint(inputs[i], &row, &col)
		visitedHouses[Point{row, col}] += 1
	}

	return len(visitedHouses)
}

func findNumberOfHousesBySantaAndRobo(inputs []string) int {
	row := 0
	roboRow := 0
	col := 0
	roboCol := 0
	visitedHouses := make(map[Point]int)
	visitedHouses[Point{row, col}] = 1

	for i := 0; i < len(inputs); i += 2 {
		findNextPoint(inputs[i], &row, &col)
		visitedHouses[Point{row, col}] += 1

		findNextPoint(inputs[i+1], &roboRow, &roboCol)
		visitedHouses[Point{roboRow, roboCol}] += 1
	}

	return len(visitedHouses)
}

func findNextPoint(santaInput string, col *int, row *int) {
	if santaInput == ">" {
		*col++
	} else if santaInput == "<" {
		*col--
	} else if santaInput == "^" {
		*row++
	} else {
		*row--
	}
}

func getInput() ([]string, error) {
	filePath := "/Users/ashwitha/GolandProjects/advent_of_code/2015/day3/input1.txt"

	lines, err := utils.ReadLines(filePath)
	if err != nil {
		return nil, err
	}
	directions := strings.Split(lines[0], "")
	return directions, nil
}

func main() {
	parenthesis, err := getInput()
	if err != nil {
		fmt.Errorf("Error while parsing input: %+v", err)
	}

	wg.Add(2)
	go func() {
		defer wg.Done()
		res := findNumberOfHouses(parenthesis)
		fmt.Println("result day3 part1:", res)
	}()

	go func() {
		defer wg.Done()
		res := findNumberOfHousesBySantaAndRobo(parenthesis)
		fmt.Println("result day3 part2:", res)
	}()

	wg.Wait()
}
