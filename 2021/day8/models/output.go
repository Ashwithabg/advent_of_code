package models

type Output string

func (o Output) hasFourDigits() bool {
	return len(o) == 4
}

func (o Output) hasTwoDigits() bool {
	return len(o) == 2
}

func (o Output) hasThreeDigits() bool {
	return len(o) == 3
}

func (o Output) hasSevenDigits() bool {
	return len(o) == 7
}

func (o Output) IsHavingSameValueAs(value string) bool {
	return string(o) == value
}
