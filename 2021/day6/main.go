package main

import (
	"fmt"
	"strconv"
	"strings"

	"advent_of_code/utils"
)

func calculateNumberOfFishes(fishes []int, days int) int {
	for j := 0; j < days; j++ {
		var newFishes []int
		for i := 0; i < len(fishes); i++ {
			if fishes[i] == 0 {
				newFishes = append(newFishes, 8)
				fishes[i]=6
			} else {
				fishes[i]--
			}
		}

		fishes = append(fishes, newFishes...)
	}

	return len(fishes)
}

func main() {
	elements, err := getInput()
	if err != nil {
		fmt.Errorf("Error while parsing input: %+v", err)
	}

	res := calculateNumberOfFishes(elements, 80)
	fmt.Println("result day1 part1:", res)
}

func getInput() ([]int, error) {
	filePath := "/Users/ashwitha/GolandProjects/advent_of_code/2021/day6/input.txt"

	lines, err := utils.ReadLines(filePath)
	if err != nil {
		return nil, err
	}

	elements := strings.Split(lines[0], ",")

	var numbers []int
	for _, ele := range elements{
		n, _ := strconv.Atoi(ele)
		numbers = append(numbers, n)
	}

	return numbers, nil
}
