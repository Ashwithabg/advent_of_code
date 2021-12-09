package main

import (
	"advent_of_code/2015/day2/models"
	utils "advent_of_code/utils"

	"fmt"
	"sort"
	"strings"
)

func findSurfaceLength(rectangularPrisms []models.RectangularPrism) int {
	sum := 0

	for _, rectangularPrism := range rectangularPrisms {
		sum += rectangularPrism.FindSurfaceArea()
	}

	return sum
}

func findRibbonLength(rectangularPrisms []models.RectangularPrism) int {
	sum := 0

	for _, rectangularPrism := range rectangularPrisms {
		sum += rectangularPrism.FindTotalLengthOfRibbonRequired()
	}

	return sum
}

func main() {
	prisms, err := getInput()
	if err != nil {
		fmt.Errorf("Error while parsing input: %+v", err)
	}

	res := findSurfaceLength(prisms)
	fmt.Println("result day2 part1:", res)


	res = findRibbonLength(prisms)
	fmt.Println("result day2 part2:", res)
}

func getInput() ([]models.RectangularPrism, error) {
	filePath := "/Users/ashwitha/GolandProjects/advent_of_code/2015/day2/input.txt"

	lines, err := utils.ReadLines(filePath)
	if err != nil {
		return nil, err
	}

	var prisms []models.RectangularPrism
	for _, line := range lines {
		values := strings.Split(line, "x")
		numbers := utils.ToIntSlice(values)
		sort.Ints(numbers)
		prisms = append(prisms, models.RectangularPrism{
			Length: numbers[0],
			Width:  numbers[1],
			Height: numbers[2],
		})
	}

	return prisms, nil
}
