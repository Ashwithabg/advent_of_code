package models

import "advent_of_code/utils"

type BracketLine []Bracket

var pointsMap = map[Bracket]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

var ClosingPointsMap = map[string]int{
	"(": 1,
	"[": 2,
	"{": 3,
	"<": 4,
}

func (bracketLine BracketLine) FindCorruptedScore() int {
	var bracketTracker utils.Stack
	score := 0

	for _, bracket := range bracketLine {
		if bracket.IsOpening() {
			bracketTracker.Push(string(bracket))
		} else {
			poppedValue, _ := bracketTracker.Pop()
			if Bracket(poppedValue).IsCorruptedClosing(bracket) {
				score += pointsMap[bracket]
			}
		}
	}

	return score
}

func (bracketLine BracketLine) GetIncompleteLineScore() int {
	var openingBrackets utils.Stack

	for _, bracket := range bracketLine {
		if bracket.IsOpening() {
			openingBrackets.Push(string(bracket))
		} else {
			openingBracket, _ := openingBrackets.Pop()
			if Bracket(openingBracket).IsCorruptedClosing(bracket) {
				return 0
			}
		}
	}

	return bracketLine.getIncompleteBracketScore(openingBrackets)
}

func (bracketLine BracketLine) getIncompleteBracketScore(openingBrackets utils.Stack) int {
	score := 0
	for openingBrackets.Size() != 0 {
		poppedValue, _ := openingBrackets.Pop()
		score *= 5
		score += ClosingPointsMap[poppedValue]
	}
	return score
}
