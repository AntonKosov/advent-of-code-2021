package main

import (
	"testing"

	"github.com/AntonKosov/advent-of-code-2021/aoc"
)

func Test2D(t *testing.T) {
	actual := process([][]aoc.Vector3{
		{
			aoc.NewVector3(0, 2, 1),
			aoc.NewVector3(4, 1, 1),
			aoc.NewVector3(3, 3, 1),
		},
		{
			aoc.NewVector3(-1, -1, 1),
			aoc.NewVector3(-5, 0, 1),
			aoc.NewVector3(-2, 1, 1),
		},
	}, 3)
	expected := 3
	if actual != expected {
		t.Fatalf("Expected: %v, actual: %v", expected, actual)
	}
}
