package utils

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
