package main

import (
	"advent_of_code/utils"
	"strconv"
	"strings"
)

func getInput() ([]PasswordDetail, error) {
	filePath := "/Users/ashwitha/GolandProjects/advent_of_code/2020/day2/input.txt"

	lines, err := utils.ReadLines(filePath)
	if err != nil {
		return nil, err
	}

	var passwordDetails []PasswordDetail
	for _, line := range lines {
		words := strings.Split(line, " ")
		numbers := strings.Split(words[0], "-")
		letters := strings.Split(words[1], ":")

		minValue, err := strconv.Atoi(numbers[0])
		if err != nil {
			return passwordDetails, err
		}

		maxValue, err := strconv.Atoi(numbers[1])
		if err != nil {
			return passwordDetails, err
		}

		passwordDetails = append(passwordDetails, PasswordDetail{
			minCharacter: minValue,
			maxCharacter: maxValue,
			letter:       letters[0],
			password:     words[2],
		})
	}

	return passwordDetails, nil
}
