package main

import (
	"fmt"
	"strconv"
	"strings"

	"advent_of_code/utils"
)

func calculatePosition() (int, error) {
	filePath := "/Users/ashwitha/GolandProjects/advent_of_code/2021/day2/input.txt"

	elements, err := utils.ReadLines(filePath)
	if err != nil {
		return 0, err
	}

	horizontalPosition := 0
	depth := 0

	for i := 0; i < len(elements); i++ {
		values := strings.Split(elements[i], " ")
		command := values[0]
		value, err := strconv.Atoi(values[1])
		if err != nil {
			fmt.Errorf("%+v", err)
			return 0, err
		}

		switch command {
		case "forward":
			horizontalPosition += value
		case "down":
			depth += value
		case "up":
			depth -= value
		}
	}

	if horizontalPosition != 0 && depth !=0 {
		return horizontalPosition * depth, nil
	}

	if horizontalPosition != 0 {
		return horizontalPosition, nil
	}

	return depth, nil
}

func calculatePositionWithDepthAimAndHorizontalPosition() (int, error) {
	filePath := "/Users/ashwitha/GolandProjects/advent_of_code/2021/day2/input.txt"

	elements, err := utils.ReadLines(filePath)
	if err != nil {
		return 0, err
	}

	horizontalPosition := 0
	depth := 0
	aim := 0

	for i := 0; i < len(elements); i++ {
		values := strings.Split(elements[i], " ")
		command := values[0]
		value, err := strconv.Atoi(values[1])
		if err != nil {
			fmt.Errorf("%+v", err)
			return 0, err
		}

		switch command {
		case "forward":
			horizontalPosition += value
			depth += aim * value
		case "down":
			aim += value
		case "up":
			aim -= value
		}
	}

	if horizontalPosition != 0 && depth !=0 {
		return horizontalPosition * depth, nil
	}

	if horizontalPosition != 0 {
		return horizontalPosition, nil
	}

	return depth, nil
}



func main() {
	res, err := calculatePosition()
	if err != nil {
		fmt.Errorf("error %+v", err)
	}

	fmt.Println("result day2 part1:", res)

	res, err = calculatePositionWithDepthAimAndHorizontalPosition()
	if err != nil {
		fmt.Errorf("error %+v", err)
	}

	fmt.Println("result day2 part2:", res)
}
