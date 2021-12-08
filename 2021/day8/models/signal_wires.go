package models

import (
	"advent_of_code/utils"
)

type SignalPatterns [10]string

var SignalLengthAndDigits = map[int]int{
	2: 1,
	3: 7,
	4: 4,
	7: 8,
}

func (signalPatterns *SignalPatterns) FindSignalPatternsBasedOnOverlaps(input string) {
	if signalPatterns.isTwoFor(input) {
		signalPatterns[2] = input
	}
	if signalPatterns.isZeroFor(input) {
		signalPatterns[0] = input
	}
	if signalPatterns.isThreeFor(input) {
		signalPatterns[3] = input
	}
	if signalPatterns.isFiveFor(input) {
		signalPatterns[5] = input
	}
	if signalPatterns.isSixFor(input) {
		signalPatterns[6] = input
	}
	if signalPatterns.isNineFOr(input) {
		signalPatterns[9] = input
	}
}

func (signalPatterns *SignalPatterns) FindSignalPatternsBasedOnLength(input string) {
	length := len(input)
	val, ok := SignalLengthAndDigits[length]
	if ok {
		signalPatterns[val] = input
	}
}

func (signalPatterns *SignalPatterns) SortValues() {
	for key, value := range signalPatterns {
		signalPatterns[key] = utils.SortString(value)
	}
}

func (signalPatterns SignalPatterns) isTwoFor(input string) bool {
	return signalPatterns[2] == "" && len(input) == 5 &&
		utils.GetOverlapCount(input, signalPatterns[1]) == 1 &&
		utils.GetOverlapCount(input, signalPatterns[4]) == 2 &&
		utils.GetOverlapCount(input, signalPatterns[3]) == 4 &&
		utils.GetOverlapCount(input, signalPatterns[7]) == 2
}

func (signalPatterns SignalPatterns) isZeroFor(input string) bool {
	return signalPatterns[0] == "" &&
		len(input) == 6 &&
		utils.GetOverlapCount(input, signalPatterns[1]) == 2 &&
		utils.GetOverlapCount(input, signalPatterns[2]) == 4 &&
		utils.GetOverlapCount(input, signalPatterns[3]) == 4
}

func (signalPatterns SignalPatterns) isSixFor(input string) bool {
	return signalPatterns[6] == "" &&
		len(input) == 6 &&
		utils.GetOverlapCount(input, signalPatterns[1]) == 1 &&
		utils.GetOverlapCount(input, signalPatterns[2]) == 4
}

func (signalPatterns SignalPatterns) isFiveFor(input string) bool {
	return signalPatterns[5] == "" &&
		len(input) == 5 &&
		utils.GetOverlapCount(input, signalPatterns[2]) == 3 &&
		utils.GetOverlapCount(input, signalPatterns[4]) == 3
}

func (signalPatterns SignalPatterns) isNineFOr(input string) bool {
	return signalPatterns[9] == "" &&
		len(input) == 6 &&
		utils.GetOverlapCount(input, signalPatterns[0]) == 5 &&
		utils.GetOverlapCount(input, signalPatterns[7]) == 3

}

func (signalPatterns SignalPatterns) isThreeFor(input string) bool {
	return signalPatterns[3] == "" &&
		len(input) == 5 &&
		utils.GetOverlapCount(input, signalPatterns[1]) == 2
}

func (signalPatterns *SignalPatterns) PatternExistsFor(output string) (bool, int) {
	for digit, pattern := range signalPatterns {
		if pattern == output {
			return true, digit
		}
	}
	return false, -1
}
