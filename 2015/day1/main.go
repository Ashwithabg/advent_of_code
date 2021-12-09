package main

import (
	"fmt"
	"strings"

	"advent_of_code/utils"
)

func findFloorNumber(parenthesis []string) int {
	openedBraces := 0
	for i := 0; i < len(parenthesis); i++ {
		if parenthesis[i] == "(" {
			openedBraces++
		}
		if parenthesis[i] == ")" {
			openedBraces--
		}
	}

	return openedBraces
}

func findBasementIndex(parenthesis []string) int {
	openedBraces := 0
	for i := 0; i < len(parenthesis); i++ {
		if parenthesis[i] == "(" {
			openedBraces++
		}
		if parenthesis[i] == ")" {
			openedBraces--
		}
		if openedBraces == -1 {
			return i
		}
	}

	return 0
}

func main() {
	parenthesis, err := getInput()
	if err != nil {
		fmt.Errorf("Error while parsing input: %+v", err)
	}

	res := findFloorNumber(parenthesis)
	fmt.Println("result day1 part1:", res)

	res = findBasementIndex(parenthesis)
	fmt.Println("result day1 part2:", res+1)
}

func getInput() ([]string, error) {
	filePath := "/Users/ashwitha/GolandProjects/advent_of_code/2015/day1/input.txt"

	lines, err := utils.ReadLines(filePath)
	if err != nil {
		return nil, err
	}
	parenthesis := strings.Split(lines[0], "")
	return parenthesis, nil
}
