package input

import (
	"advent_of_code/2021/day8/models"
	"advent_of_code/utils"
	"strings"
)

const delimiter = " | "

func GetDisplayData() []models.SevenSegmentDisplayEntry {
	filePath := "/Users/ashwitha/GolandProjects/advent_of_code/2021/day8/input/files/input1.txt"

	lines, err := utils.ReadLines(filePath)
	if err != nil {
		panic("can not read input")
	}

	return getSevenSegmentDisplayData(lines)
}

func getSevenSegmentDisplayData(lines []string) (sevenSegmentdisplayOutputs []models.SevenSegmentDisplayEntry) {
	for i := 0; i < len(lines); i++ {
		displayData := strings.Split(lines[i], delimiter)
		inputData := strings.Split(displayData[0], " ")
		outputData := strings.Split(displayData[1], " ")

		var outputs []models.Output
		for _, val := range outputData {
			outputs = append(outputs, models.Output(val))
		}

		var inputs []string
		for _, val := range inputData {
			inputs = append(inputs, val)
		}

		sevenSegmentdisplayOutputs = append(sevenSegmentdisplayOutputs, models.SevenSegmentDisplayEntry{
			SignalInputs: inputs,
			Outputs:      outputs,
		})
	}
	return sevenSegmentdisplayOutputs
}
