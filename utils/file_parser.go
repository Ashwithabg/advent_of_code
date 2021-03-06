package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadNumbersFromFile(filePath string) ([]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error parsing the file %+v", err)
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		element, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("error while converting string to int %+v", err)
		}
		lines = append(lines, element)
	}
	return lines, nil
}

func ReadLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error parsing the file %+v", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}
