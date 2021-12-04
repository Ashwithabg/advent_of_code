package main

import (
	"fmt"
	"strconv"
	"strings"

	"advent_of_code/utils"
)

type SubmarineReading struct {
	Name  string
	Value int
}

func calculatePosition(readings []SubmarineReading) int {
	horizontalPosition := 0
	depth := 0

	for _, reading := range readings {
		switch reading.Name {
		case "forward":
			horizontalPosition += reading.Value
		case "down":
			depth += reading.Value
		case "up":
			depth -= reading.Value
		}
	}

	return horizontalPosition * depth
}

func calculatePositionWithDepthAimAndHorizontalPosition(readings []SubmarineReading) int {
	horizontalPosition := 0
	depth := 0
	aim := 0

	for _, reading := range readings {
		switch reading.Name {
		case "forward":
			horizontalPosition += reading.Value
			depth += aim * reading.Value
		case "down":
			aim += reading.Value
		case "up":
			aim -= reading.Value
		}
	}
	return horizontalPosition * depth
}

func getSubmarineReadings() ([]SubmarineReading, error) {
	filePath := "/Users/ashwitha/GolandProjects/advent_of_code/2021/day2/input.txt"

	elements, err := utils.ReadLines(filePath)
	if err != nil {
		return nil, err
	}

	var submarineReadings []SubmarineReading
	for i := 0; i < len(elements); i++ {
		values := strings.Split(elements[i], " ")
		convertedValue, err := strconv.Atoi(values[1])
		if err != nil {
			fmt.Errorf("%+v", err)
			return nil, err
		}
		reading := SubmarineReading{
			Name:  values[0],
			Value: convertedValue,
		}
		submarineReadings = append(submarineReadings, reading)
	}
	return submarineReadings, nil
}

func main() {
	readings, err := getSubmarineReadings()
	if err != nil {
		fmt.Errorf("Error while parsing the input: %+v", err)
	}

	res := calculatePosition(readings)
	if err != nil {
		fmt.Errorf("error %+v", err)
	}

	fmt.Println("result day2 part1:", res)

	res = calculatePositionWithDepthAimAndHorizontalPosition(readings)

	fmt.Println("result day2 part2:", res)
}
