package inmemory

type numbers interface {
	int | int8 | int16 | uint32 | int64 | float32 | float64
}

func max[T numbers](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func min[T numbers](a, b T) T {
	if a < b {
		return a
	}
	return b
}
