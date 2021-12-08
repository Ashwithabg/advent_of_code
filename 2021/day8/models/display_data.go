package models

import "advent_of_code/utils"

type SevenSegmentDisplayEntry struct {
	SignalInputs []string
	Outputs      []Output
}

func (dd SevenSegmentDisplayEntry) FindCountOf234Or7InOutputs() int {
	counter := 0

	for _, output := range dd.Outputs {
		if output.hasThreeDigits() || output.hasFourDigits() ||
			output.hasSevenDigits() || output.hasTwoDigits() {
			counter++
		}
	}

	return counter
}

func (dd SevenSegmentDisplayEntry) GetSortedOutputs() (sortedOutput [10]Output) {
	for key, value := range dd.Outputs {
		sortedOutput[key] = Output(utils.SortString(string(value)))
	}
	return
}
