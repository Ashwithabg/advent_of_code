package main

import (
	"fmt"
)

func validatePassword(passwordDetails []PasswordDetail) int {
	validaPasswordCounter := 0
	for _, passwordDetail := range passwordDetails {
		if passwordDetail.isRangeValid() {
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

	res := validatePassword(lines)
	fmt.Println("result day1 part1:", res)
}
