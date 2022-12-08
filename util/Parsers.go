package util

import "strconv"

func ParseInt(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}
