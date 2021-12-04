package main

import (
	"fmt"
)

func validatePassword(passwordDetails []PasswordDetail) int {
	validaPasswordCounter := 0
	for _, passwordDetail := range passwordDetails {

		letterCounter := 0
		for _, letter := range passwordDetail.password {
			if string(letter) == passwordDetail.letter {
				letterCounter++
			}
		}

		if letterCounter <= passwordDetail.maxCharacter &&
			letterCounter >= passwordDetail.minCharacter {
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
