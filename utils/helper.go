package utils

import (
	"sort"
	"strconv"
	"strings"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ToInt(value string) int {
	number, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return number
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func SummationOfNaturalNumbers(numberOfElements int) int {
	return numberOfElements * (numberOfElements + 1) / 2
}

func SortString(value string) string {
	splittedValue := strings.Split(value, "")
	sort.Strings(splittedValue)
	return strings.Join(splittedValue, "")
}

func GetOverlapCount(letters, anotherSetOfLetters string) int {
	count := 0
	for _, letter := range letters {
		if strings.Contains(anotherSetOfLetters, string(letter)) {
			count++
		}
	}
	return count
}
