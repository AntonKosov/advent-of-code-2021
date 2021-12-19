package aoc

type Vector3 struct {
	X, Y, Z int
}

func NewVector3(x, y, z int) Vector3 {
	return Vector3{X: x, Y: y, Z: z}
}

func (v Vector3) Add(av Vector3) Vector3 {
	return NewVector3(v.X+av.X, v.Y+av.Y, v.Z+av.Z)
}

func (v Vector3) Sub(av Vector3) Vector3 {
	return NewVector3(v.X-av.X, v.Y-av.Y, v.Z-av.Z)
}

func (v Vector3) ManhattanLength() int {
	return Abs(v.X) + Abs(v.Y) + Abs(v.Z)
}
