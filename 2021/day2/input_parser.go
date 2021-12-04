package main

import (
	"advent_of_code/utils"
	"fmt"
	"strconv"
	"strings"
)

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
			Name:  CommandName(values[0]),
			Value: convertedValue,
		}
		submarineReadings = append(submarineReadings, reading)
	}
	return submarineReadings, nil
}
