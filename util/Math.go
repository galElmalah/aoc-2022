package util

type NumbersConstraint interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func Abs[T NumbersConstraint](num T) T {
	if num > 0 {
		return num
	}
	return -num
}

// Return the max number from a slice of numbers
func Max[T NumbersConstraint](values ...T) T {
	max := values[0]
	for _, v := range values[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

// Return the min number from a slice of numbers
func Min[T NumbersConstraint](values ...T) T {
	min := values[0]
	for _, v := range values[1:] {
		if v < min {
			min = v
		}
	}
	return min
}
