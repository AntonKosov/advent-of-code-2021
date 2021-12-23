package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2021/aoc"
)

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

type cuboid struct {
	min aoc.Vector3
	max aoc.Vector3
}

func newCuboid(v1 aoc.Vector3, v2 aoc.Vector3) cuboid {
	v1, v2 = aoc.NewVector3(aoc.Min(v1.X, v2.X), aoc.Min(v1.Y, v2.Y), aoc.Min(v1.Z, v2.Z)),
		aoc.NewVector3(aoc.Max(v1.X, v2.X), aoc.Max(v1.Y, v2.Y), aoc.Max(v1.Z, v2.Z))
	return cuboid{
		min: v1,
		max: v2,
	}
}

type operation struct {
	isOn bool
	c    cuboid
}

func read() (data []operation) {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	for _, line := range lines {
		sof := strings.Split(line, " ")
		comps := strings.Split(sof[1], ",")
		xc := strings.Split(comps[0][2:], "..")
		yc := strings.Split(comps[1][2:], "..")
		zc := strings.Split(comps[2][2:], "..")
		v1 := aoc.NewVector3(aoc.StrToInt(xc[0]), aoc.StrToInt(yc[0]), aoc.StrToInt(zc[0]))
		v2 := aoc.NewVector3(aoc.StrToInt(xc[1]), aoc.StrToInt(yc[1]), aoc.StrToInt(zc[1]))
		c := newCuboid(v1, v2)
		op := operation{isOn: sof[0] == "on", c: c}
		if c.min.X >= -50 && c.min.Y >= -50 && c.min.Z >= -50 && c.max.X <= 50 && c.max.Y <= 50 && c.max.Z <= 50 {
			data = append(data, op)
		}
	}

	return data
}

func process(data []operation) int {
	m := [101][101][101]bool{}
	for _, operation := range data {
		min := operation.c.min
		max := operation.c.max
		for x := min.X; x <= max.X; x++ {
			for y := min.Y; y <= max.Y; y++ {
				for z := min.Z; z <= max.Z; z++ {
					m[x+50][y+50][z+50] = operation.isOn
				}
			}
		}
	}

	sum := 0
	for _, i := range m {
		for _, j := range i {
			for _, k := range j {
				if k {
					sum++
				}
			}
		}
	}

	return sum
}
