package main

import (
	"fmt"
	"strconv"
	"strings"

	"advent_of_code/utils"
)

func calculatePower() (int64, error) {
	readings, err := getReadings()
	if err != nil {
		return 0, err
	}

	readingLength := len(readings[0])
	countOfOnesInAllReadings := make([]int, readingLength)
	for i := 0; i < len(readings); i++ {
		bitValuesForReading := strings.Split(readings[i], "")
		for position, bitValue := range bitValuesForReading {
			if bitValue == "1" {
				countOfOnesInAllReadings[position]++
			}
		}
	}

	gammaRate := findGammaRate(countOfOnesInAllReadings, readings)
	epsilonRate := getEpsilonRate(countOfOnesInAllReadings, readings)
	return gammaRate * epsilonRate, nil
}

func getEpsilonRate(countOfOnesInAllReadings []int, readings []string) int64 {
	epsilonRateBinaryValue := ""
	for i := 0; i < len(countOfOnesInAllReadings); i++ {
		if countOfOnesInAllReadings[i] > (len(readings) / 2) {
			epsilonRateBinaryValue += "0"
		} else {
			epsilonRateBinaryValue += "1"
		}
	}

	return convertBinaryToDecimal(epsilonRateBinaryValue)
}

func findGammaRate(countOfOnesInAllReadings []int, readings []string) int64 {
	gammaRateInBinary := ""

	for _, count := range countOfOnesInAllReadings {
		if count > len(readings)/2 {
			gammaRateInBinary += "1"
		} else {
			gammaRateInBinary += "0"
		}
	}

	return convertBinaryToDecimal(gammaRateInBinary)
}

func getReadings() ([]string, error) {
	filePath := "/Users/ashwitha/GolandProjects/advent_of_code/2021/day3/input.txt"
	elements, err := utils.ReadLines(filePath)
	if err != nil {
		return nil, err
	}
	return elements, nil
}

func calculateLifeSupportRatingOfSubmarine() (int64, error) {
	elements, err := getReadings()
	if err != nil {
		return 0, err
	}

	oxygenRate := getOxygenGenerationRate(elements)
	carbonDioxideRate := getCO2ScrubberRating(elements)
	return oxygenRate * carbonDioxideRate, nil
}

func getCO2ScrubberRating(ratings []string) int64 {
	filteredRatings := ratings
	ratingSize := len(ratings[0])

	for ratingIndex := 0; ratingIndex < ratingSize && len(filteredRatings) > 1; ratingIndex++ {
		countOfOnes := findNumberOfOnes(filteredRatings, ratingIndex)
		countOfZeros := len(filteredRatings) - countOfOnes
		if countOfOnes < countOfZeros {
			filteredRatings = getRatingsWith(1, filteredRatings, ratingIndex)
		} else {
			filteredRatings = getRatingsWith(0, filteredRatings, ratingIndex)
		}
	}

	return convertBinaryToDecimal(filteredRatings[0])
}

func getOxygenGenerationRate(ratings []string) int64 {
	filteredRatings := ratings
	ratingSize := len(ratings[0])

	for ratingIndex := 0; ratingIndex < ratingSize && len(filteredRatings) > 1; ratingIndex++ {
		countOfOnes := findNumberOfOnes(filteredRatings, ratingIndex)
		countOfZeros := len(filteredRatings) - countOfOnes
		if countOfOnes >= countOfZeros {
			filteredRatings = getRatingsWith(1, filteredRatings, ratingIndex)
		} else {
			filteredRatings = getRatingsWith(0, filteredRatings, ratingIndex)
		}
	}

	return convertBinaryToDecimal(filteredRatings[0])
}

func getRatingsWith(value int, ratings []string, position int) []string {
	var filteredValues []string
	for _, rating := range ratings {
		stringifiedRating := strings.Split(rating, "")
		if stringifiedRating[position] == strconv.Itoa(value) {
			filteredValues = append(filteredValues, rating)
		}
	}

	return filteredValues
}

func findNumberOfOnes(filteredRatings []string, ratingIndex int) int {
	countOfOnes := 0

	for _, rating := range filteredRatings {
		convertedValue := strings.Split(rating, "")
		if convertedValue[ratingIndex] == "1" {
			countOfOnes++
		}
	}

	return countOfOnes
}

func convertBinaryToDecimal(binaryReading string) int64 {
	reading, err := strconv.ParseInt(binaryReading, 2, 64)
	if err != nil {
		fmt.Errorf("Error on converting binary To decimal %+v\n", err)
		return 0
	}

	return reading
}

func main() {
	res, err := calculatePower()
	if err != nil {
		fmt.Errorf("error %+v", err)
	}

	fmt.Println("result day3 part1:", res)

	res, err = calculateLifeSupportRatingOfSubmarine()
	if err != nil {
		fmt.Errorf("error %+v", err)
	}

	fmt.Println("result day3 part2:", res)
}
