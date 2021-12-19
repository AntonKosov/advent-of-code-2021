package aoc_test

import (
	"testing"

	"github.com/AntonKosov/advent-of-code-2021/aoc"
)

func TestMatrix3x3MulVector3(t *testing.T) {
	m := aoc.NewMatrix3x3([9]int{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	})
	v := aoc.NewVector3(11, 12, 13)
	actual := m.MulVector3(v)
	expected := aoc.NewVector3(74, 182, 290)
	if expected != actual {
		t.Fatalf("TestMatrix3x3MulVector3: expected %v, actual %v", expected, actual)
	}
}

func testRotation(t *testing.T, f func(angle int) aoc.Matrix3x3, angle int, expected aoc.Vector3) {
	testAngle := func(a int) {
		rm := f(a)
		actual := rm.MulVector3(aoc.NewVector3(1, 1, 1))
		if actual != expected {
			t.Fatalf("expected %v, actual %v", expected, actual)
		}
	}
	testAngle(angle)
	testAngle(angle + 360)
	testAngle(angle - 360)
	testAngle(angle - 720)
}

func TestRotationMatrixXY0(t *testing.T) {
	testRotation(t, aoc.RotationMatrixXY, 0, aoc.NewVector3(1, 1, 1))
}

func TestRotationMatrixXY90(t *testing.T) {
	testRotation(t, aoc.RotationMatrixXY, 90, aoc.NewVector3(-1, 1, 1))
}

func TestRotationMatrixXY180(t *testing.T) {
	testRotation(t, aoc.RotationMatrixXY, 180, aoc.NewVector3(-1, -1, 1))
}

func TestRotationMatrixXY270(t *testing.T) {
	testRotation(t, aoc.RotationMatrixXY, 270, aoc.NewVector3(1, -1, 1))
}

func TestRotationMatrixXZ0(t *testing.T) {
	testRotation(t, aoc.RotationMatrixXZ, 0, aoc.NewVector3(1, 1, 1))
}

func TestRotationMatrixXZ90(t *testing.T) {
	testRotation(t, aoc.RotationMatrixXZ, 90, aoc.NewVector3(1, 1, -1))
}

func TestRotationMatrixXZ180(t *testing.T) {
	testRotation(t, aoc.RotationMatrixXZ, 180, aoc.NewVector3(-1, 1, -1))
}

func TestRotationMatrixXZ270(t *testing.T) {
	testRotation(t, aoc.RotationMatrixXZ, 270, aoc.NewVector3(-1, 1, 1))
}

func TestRotationMatrixYZ0(t *testing.T) {
	testRotation(t, aoc.RotationMatrixYZ, 0, aoc.NewVector3(1, 1, 1))
}

func TestRotationMatrixYZ90(t *testing.T) {
	testRotation(t, aoc.RotationMatrixYZ, 90, aoc.NewVector3(1, -1, 1))
}

func TestRotationMatrixYZ180(t *testing.T) {
	testRotation(t, aoc.RotationMatrixYZ, 180, aoc.NewVector3(1, -1, -1))
}

func TestRotationMatrixYZ270(t *testing.T) {
	testRotation(t, aoc.RotationMatrixYZ, 270, aoc.NewVector3(1, 1, -1))
}

func TestMatrix3x3Mul(t *testing.T) {
	a := aoc.NewMatrix3x3([9]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	b := aoc.NewMatrix3x3([9]int{10, 11, 12, 13, 14, 15, 16, 17, 18})
	actual := a.Mul(b)
	expected := aoc.NewMatrix3x3([9]int{
		84, 90, 96,
		201, 216, 231,
		318, 342, 366,
	})
	if actual != expected {
		t.Fatalf("expected: %v, actual: %v", expected, actual)
	}
}
