package utils

import "sort"

func GetMidElement(elements []int) int {
	midIndex := len(elements) / 2
	return elements[midIndex]
}

func ToIntSlice(elements []string) []int {
	var numbers []int
	for _, element := range elements {
		numbers = append(numbers, ToInt(element))
	}

	return numbers
}

func SortValues(words [10]string) [10]string {
	var sortedWords [10]string

	for key, value := range words {
		sortedWords[key] = SortString(value)
	}

	return sortedWords
}

func SortDesc(numbers []int) []int {
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[j] < numbers[i]
	})

	return numbers
}

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}