package util

import "strconv"

func ParseInt(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}

func ToInts(fromStrings []string) (result []int) {
	for _, n := range fromStrings {
		result = append(result, ParseInt(n))
	}
	return result
}
