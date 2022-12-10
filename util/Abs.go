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
