package aoc

// Vector2 represents a two-dimensional vector in left-handed Cartesian
// coordinate system
type Vector2 struct {
	X, Y int
}

func NewVector2(x, y int) Vector2 {
	return Vector2{X: x, Y: y}
}

func (v Vector2) RotateLeft() Vector2 {
	rm := [2][2]int{
		{0, -1},
		{1, 0},
	}
	return v.Rotate(rm)
}

func (v Vector2) RotateRight() Vector2 {
	rm := [2][2]int{
		{0, 1},
		{-1, 0},
	}
	return v.Rotate(rm)
}

func (v Vector2) Rotate(rm [2][2]int) Vector2 {
	newX := rm[1][0]*v.Y + rm[1][1]*v.X
	newY := rm[0][0]*v.Y + rm[0][1]*v.X
	return Vector2{X: newX, Y: newY}
}

func (v Vector2) Add(s Vector2) Vector2 {
	return Vector2{X: v.X + s.X, Y: v.Y + s.Y}
}

func (v Vector2) Mul(scalar int) Vector2 {
	return Vector2{X: v.X * scalar, Y: v.Y * scalar}
}

func (v Vector2) Sub(s Vector2) Vector2 {
	return Vector2{X: v.X - s.X, Y: v.Y - s.Y}
}

func (v Vector2) Norm() Vector2 {
	if v.X != 0 {
		v = NewVector2(v.X/Abs(v.X), v.Y)
	}
	if v.Y != 0 {
		v = NewVector2(v.X, v.Y/Abs(v.Y))
	}
	return v
}

func (v Vector2) Adjacent() []Vector2 {
	adj := make([]Vector2, 0, 8)
	for x := v.X - 1; x <= v.X+1; x++ {
		for y := v.Y - 1; y <= v.Y+1; y++ {
			if x == v.X && y == v.Y {
				continue
			}
			adj = append(adj, NewVector2(x, y))
		}
	}
	return adj
}
