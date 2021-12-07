package utils

import "strconv"

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ToInt(value string) int {
	number, err := strconv.Atoi(value)
	if err!=nil {
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
