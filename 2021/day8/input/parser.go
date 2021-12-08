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

	m := make(map[int]int)
	m[1] = 2
	m[4] = 5
	m[7] = 3
	m[8] = 7

	var values []string
	for i:=0; i< len(lines); i++ {
		data := strings.Split(lines[i], " | ")
		values = append(values, data[1])
	}

	counter := 0
	for i:=0; i< len(values); i++ {
		data := strings.Split(values[i], " ")
		for j:=0; j<len(data); j ++ {
			length := len(data[j])
			if length == 2 || length ==4 || length ==3 || length==7 {
				counter++
			}

		}
	}

	fmt.Println("counter-------->", counter)


	return lines, nil
}
