package main

import (
	"fmt"
	"sync"

	"advent_of_code/utils"
)

var wg sync.WaitGroup

func findTwoNumbersThatSumTo2020(elements []int) int {
	result := 0
	for i := 0; i < len(elements); i++ {
		for j := 0; j < len(elements); j++ {
			if i != j && (elements[i]+elements[j]) == 2020 {
				return elements[i] * elements[j]
			}
		}
	}

	return result
}

func findThreeNumbersThatSumTo2020(elements []int) int {
	result := 0
	for i := 0; i < len(elements); i++ {
		for j := 0; j < len(elements); j++ {
			for k := 0; k < len(elements); k++ {
				if i != j && j != k && (elements[i]+elements[j]+elements[k]) == 2020 {
					return elements[i] * elements[j] * elements[k]
				}
			}
		}
	}

	return result
}

func main() {
	elements, err := getInput()
	if err != nil {
		fmt.Errorf("Error while parsing input: %+v", err)
	}

	wg.Add(2)

	go func() {
		defer wg.Done()
		res := findTwoNumbersThatSumTo2020(elements)
		fmt.Println("result day1 part1:", res)
	}()

	go func() {
		defer wg.Done()
		res := findThreeNumbersThatSumTo2020(elements)
		fmt.Println("result day1 part2:", res)
	}()

	wg.Wait()

}

func getInput() ([]int, error) {
	filePath := "/Users/ashwitha/GolandProjects/advent_of_code/2020/day1/input.txt"

	elements, err := utils.ReadNumbersFromFile(filePath)
	if err != nil {
		return nil, err
	}
	return elements, nil
}
