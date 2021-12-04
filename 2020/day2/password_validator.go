package main

import (
	"fmt"
)

func validatePasswordWithRange(passwordDetails []PasswordDetail) int {
	validaPasswordCounter := 0
	for _, passwordDetail := range passwordDetails {
		if passwordDetail.isRangeValid() {
			validaPasswordCounter++
		}
	}

	return validaPasswordCounter
}


func validatePasswordWithPosition(passwordDetails []PasswordDetail) int {
	validaPasswordCounter := 0
	for _, passwordDetail := range passwordDetails {
		if passwordDetail.isPositionValid() {
			validaPasswordCounter++
		}
	}

	return validaPasswordCounter
}


func main() {
	lines, err := getInput()
	if err != nil {
		fmt.Errorf("Error while parsing input: %+v", err)
	}

	res := validatePasswordWithRange(lines)
	fmt.Println("result day1 part1:", res)

	res = validatePasswordWithPosition(lines)
	fmt.Println("result day1 part2:", res)
}
