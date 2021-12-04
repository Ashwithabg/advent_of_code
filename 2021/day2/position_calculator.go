package main

import (
	"fmt"
)

func calculatePosition(readings []SubmarineReading) int {
	horizontalPosition := 0
	depth := 0

	for _, reading := range readings {
		switch reading.Name {
		case FORWARD:
			horizontalPosition += reading.Value
		case DOWN:
			depth += reading.Value
		case UP:
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
		case FORWARD:
			horizontalPosition += reading.Value
			depth += aim * reading.Value
		case DOWN:
			aim += reading.Value
		case UP:
			aim -= reading.Value
		}
	}
	return horizontalPosition * depth
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
