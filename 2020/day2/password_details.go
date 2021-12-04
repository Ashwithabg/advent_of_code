package main

type PasswordDetail struct {
	minCharacter int
	maxCharacter int
	letter       string
	password     string
}


func (passwordDetail PasswordDetail) isRangeValid() bool{
	letterCounter := 0
	for _, letter := range passwordDetail.password {
		if string(letter) == passwordDetail.letter {
			letterCounter++
		}
	}

	return letterCounter <= passwordDetail.maxCharacter &&
		letterCounter >= passwordDetail.minCharacter
}
