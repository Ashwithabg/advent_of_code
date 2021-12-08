package input

import (
	"fmt"
	"strings"

	"advent_of_code/utils"
)

func Day8Parse() ([]string, error) {
	filePath := "/Users/ashwitha/GolandProjects/advent_of_code/2021/day8/input/files/input1.txt"

	lines, err := utils.ReadLines(filePath)
	if err != nil {
		return nil, err
	}

	var values []string
	for i := 0; i < len(lines); i++ {
		data := strings.Split(lines[i], " | ")
		values = append(values, data[1])
	}

	counter := 0
	for i := 0; i < len(values); i++ {
		data := strings.Split(values[i], " ")
		for j := 0; j < len(data); j++ {
			length := len(data[j])
			if length == 2 || length == 4 || length == 3 || length == 7 {
				counter++
			}

		}
	}

	return lines, nil
}

func Day8Part2Parse() ([]string, error) {
	filePath := "/Users/ashwitha/GolandProjects/advent_of_code/2021/day8/input/files/input1.txt"

	lines, err := utils.ReadLines(filePath)
	if err != nil {
		return nil, err
	}

	var part2 int
	for _, line := range lines {
		data := strings.Split(line, " | ")
		inputs := strings.Split(data[0], " ")
		outputs := strings.Split(data[1], " ")

		result := [10]string{}
		for _, input := range inputs {
			length := len(input)
			if length == 2 {
				result[1] = input
			} else if length == 3 {
				result[7] = input
			} else if length == 4 {
				result[4] = input
			} else if length == 7 {
				result[8] = input
			}
		}

		for i := 0; i < 10; i++ {
			for _, input := range inputs {
				if isTwo(result, input) {
					result[2] = input
				}
				if isThree(result, input) {
					result[3] = input
				}
				if isNine(result, input) {
					result[9] = input
				}
				if isFive(result, input) {
					result[5] = input
				}
				if isSix(result, input) {
					result[6] = input
				}
				if isZero(result, input) {
					result[0] = input
				}
			}
		}

		sortedInput := utils.SortValues(result)

		var outputSum string
		for _, output := range outputs {
			output = utils.SortString(output)
			for k, input := range sortedInput {
				if input == output {
					outputSum += fmt.Sprintf("%d", k)
				}
			}
		}

		part2 += utils.ToInt(outputSum)
	}

	fmt.Println(part2)

	return lines, nil
}

func isZero(wireMaps [10]string, input string) bool {
	return wireMaps[0] == "" && len(input) == 6 &&
		utils.Overlaps(input, wireMaps[1]) == 2 &&
		utils.Overlaps(input, wireMaps[2]) == 4 &&
		utils.Overlaps(input, wireMaps[3]) == 4
}

func isSix(wireMaps [10]string, input string) bool {
	return wireMaps[6] == "" && len(input) == 6 &&
		utils.Overlaps(input, wireMaps[1]) == 1 &&
		utils.Overlaps(input, wireMaps[2]) == 4
}

func isFive(wireMaps [10]string, input string) bool {
	return wireMaps[5] == "" && len(input) == 5 &&
		utils.Overlaps(input, wireMaps[2]) == 3 &&
		utils.Overlaps(input, wireMaps[4]) == 3
}

func isNine(wireMaps [10]string, input string) bool {
	return wireMaps[9] == "" && len(input) == 6 &&
		utils.Overlaps(input, wireMaps[7]) == 3 &&
		utils.Overlaps(input, wireMaps[0]) == 5
}

func isThree(wireMaps [10]string, input string) bool {
	return wireMaps[3] == "" && len(input) == 5 &&
		utils.Overlaps(input, wireMaps[1]) == 2
}

func isTwo(misplacedWireMaps [10]string, input string) bool {
	return misplacedWireMaps[2] == "" && len(input) == 5 &&
		utils.Overlaps(input, misplacedWireMaps[7]) == 2 &&
		utils.Overlaps(input, misplacedWireMaps[1]) == 1 &&
		utils.Overlaps(input, misplacedWireMaps[3]) == 4 &&
		utils.Overlaps(input, misplacedWireMaps[4]) == 2
}
