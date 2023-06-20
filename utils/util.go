package utils

type Vec2 struct {
	X int
	Y int
}

func Abs(X int) int {
	if X < 0 {
		return -X
	}
	return X
}

func Clamp(X int, min int, max int) int {
	if X < min {
		return min
	}
	if X > max {
		return max
	}
	return X
}