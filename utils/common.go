package utils

type Vec2 struct {
	X int
	Y int
}

func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{v.X + other.X, v.Y + other.Y}
}

func GetDelta(start Vec2, end Vec2) Vec2 {
	return Vec2{end.X - start.X, end.Y - start.Y}
}

type numeric interface {
	uint8 | uint16 | uint32 | uint64 | int8 | int16 | int32 | int64 | float32 | float64 | int
}

func Abs[N numeric](X N) N {
	if X < 0 {
		return -X
	}
	return X
}

func Clamp[N numeric](X N, min N, max N) N {
	if X < min {
		return min
	}
	if X > max {
		return max
	}
	return X
}

func Min[N numeric](X N, Y N) N {
	if X < Y {
		return X
	}
	return Y
}

func Max[N numeric](X N, Y N) N {
	if X > Y {
		return X
	}
	return Y
}

