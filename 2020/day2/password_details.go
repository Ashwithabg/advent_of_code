package main

type PasswordDetail struct {
	minCharacter int
	maxCharacter int
	letter       string
	password     string
}

func (passwordDetail PasswordDetail) isRangeValid() bool {
	letterCounter := 0
	for _, letter := range passwordDetail.password {
		if string(letter) == passwordDetail.letter {
			letterCounter++
		}
	}

	return letterCounter <= passwordDetail.maxCharacter &&
		letterCounter >= passwordDetail.minCharacter
}

func (passwordDetail PasswordDetail) isPositionValid() bool {
	letters := []rune(passwordDetail.password)
	firstLetter := string(letters[passwordDetail.minCharacter-1])
	secondLetter := string(letters[passwordDetail.maxCharacter-1])

	return (firstLetter == passwordDetail.letter && secondLetter != passwordDetail.letter) ||
		(firstLetter != passwordDetail.letter && secondLetter == passwordDetail.letter)
}
