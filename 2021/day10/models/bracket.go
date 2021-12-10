package models

var bracketMappers = map[Bracket]Bracket{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

type Bracket string

func (bracket Bracket) IsOpening() bool {
	_, ok := bracketMappers[bracket]
	return ok
}

func (bracket Bracket) IsCorruptedClosing(closingBracket Bracket) bool {
	CorrectBracket, _ := bracketMappers[bracket]
	return closingBracket != CorrectBracket
}
