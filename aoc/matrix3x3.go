package aoc

type Matrix3x3 struct {
	// 0 1 2
	// 3 4 5
	// 6 7 8
	v [9]int
}

func NewMatrix3x3(values [9]int) Matrix3x3 {
	return Matrix3x3{v: values}
}

func NewIdentityMatrix3x3() Matrix3x3 {
	return NewMatrix3x3([9]int{1, 0, 0, 0, 1, 0, 0, 0, 1})
}

func (m Matrix3x3) MulVector3(v Vector3) Vector3 {
	return Vector3{
		X: m.v[0]*v.X + m.v[1]*v.Y + m.v[2]*v.Z,
		Y: m.v[3]*v.X + m.v[4]*v.Y + m.v[5]*v.Z,
		Z: m.v[6]*v.X + m.v[7]*v.Y + m.v[8]*v.Z,
	}
}

func (m Matrix3x3) Mul(am Matrix3x3) Matrix3x3 {
	return NewMatrix3x3([9]int{
		// 0 1 2
		// 3 4 5
		// 6 7 8
		m.v[0]*am.v[0] + m.v[1]*am.v[3] + m.v[2]*am.v[6], // 0
		m.v[0]*am.v[1] + m.v[1]*am.v[4] + m.v[2]*am.v[7], // 1
		m.v[0]*am.v[2] + m.v[1]*am.v[5] + m.v[2]*am.v[8], // 2

		m.v[3]*am.v[0] + m.v[4]*am.v[3] + m.v[5]*am.v[6], // 3
		m.v[3]*am.v[1] + m.v[4]*am.v[4] + m.v[5]*am.v[7], // 4
		m.v[3]*am.v[2] + m.v[4]*am.v[5] + m.v[5]*am.v[8], // 5

		m.v[6]*am.v[0] + m.v[7]*am.v[3] + m.v[8]*am.v[6], // 6
		m.v[6]*am.v[1] + m.v[7]*am.v[4] + m.v[8]*am.v[7], // 7
		m.v[6]*am.v[2] + m.v[7]*am.v[5] + m.v[8]*am.v[8], // 8
	})
}

var rotationMatrixXY map[int]Matrix3x3
var rotationMatrixXZ map[int]Matrix3x3
var rotationMatrixYZ map[int]Matrix3x3

func init() {
	cos90 := 0
	sin90 := 1
	cos180 := -1
	sin180 := 0
	cos270 := 0
	sin270 := -1
	rotMatrixes := func(f func(c, s int) Matrix3x3) map[int]Matrix3x3 {
		return map[int]Matrix3x3{
			0:   NewIdentityMatrix3x3(),
			90:  f(cos90, sin90),
			180: f(cos180, sin180),
			270: f(cos270, sin270),
		}
	}
	rotationMatrixXY = rotMatrixes(func(c, s int) Matrix3x3 {
		return NewMatrix3x3([9]int{c, -s, 0, s, c, 0, 0, 0, 1})
	})
	rotationMatrixXZ = rotMatrixes(func(c, s int) Matrix3x3 {
		return NewMatrix3x3([9]int{c, 0, s, 0, 1, 0, -s, 0, c})
	})
	rotationMatrixYZ = rotMatrixes(func(c, s int) Matrix3x3 {
		return NewMatrix3x3([9]int{1, 0, 0, 0, c, -s, 0, s, c})
	})
}

func RotationMatrixXY(angle int) Matrix3x3 {
	return rotationMatrixXY[normAngle(angle)]
}

func RotationMatrixXZ(angle int) Matrix3x3 {
	return rotationMatrixXZ[normAngle(angle)]
}

func RotationMatrixYZ(angle int) Matrix3x3 {
	return rotationMatrixYZ[normAngle(angle)]
}

func normAngle(angle int) int {
	angle = angle % 360
	if angle < 0 {
		angle += 360
	}
	return angle
}
