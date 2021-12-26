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
	charge int
	min    aoc.Vector3
	max    aoc.Vector3
}

func (c cuboid) volume() int {
	return (c.max.X - c.min.X + 1) * (c.max.Y - c.min.Y + 1) * (c.max.Z - c.min.Z + 1)
}

func (c cuboid) intersection(ac cuboid) (inter cuboid, ok bool) {
	min := aoc.NewVector3(
		aoc.Max(c.min.X, ac.min.X),
		aoc.Max(c.min.Y, ac.min.Y),
		aoc.Max(c.min.Z, ac.min.Z),
	)
	max := aoc.NewVector3(
		aoc.Min(c.max.X, ac.max.X),
		aoc.Min(c.max.Y, ac.max.Y),
		aoc.Min(c.max.Z, ac.max.Z),
	)

	if min.X > max.X || min.Y > max.Y || min.Z > max.Z {
		return cuboid{}, false
	}

	return cuboid{min: min, max: max}, true
}

func read() (data []cuboid) {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	for _, line := range lines {
		sof := strings.Split(line, " ")
		comps := strings.Split(sof[1], ",")
		xc := strings.Split(comps[0][2:], "..")
		yc := strings.Split(comps[1][2:], "..")
		zc := strings.Split(comps[2][2:], "..")
		min := aoc.NewVector3(aoc.StrToInt(xc[0]), aoc.StrToInt(yc[0]), aoc.StrToInt(zc[0]))
		max := aoc.NewVector3(aoc.StrToInt(xc[1]), aoc.StrToInt(yc[1]), aoc.StrToInt(zc[1]))
		charge := 1
		if sof[0] == "off" {
			charge = -1
		}
		op := cuboid{charge: charge, min: min, max: max}
		data = append(data, op)
	}

	return data
}

func process(input []cuboid) int {
	cuboids := make([]cuboid, 0, len(input))
	for _, nc := range input {
		count := len(cuboids)
		if nc.charge > 0 {
			cuboids = append(cuboids, nc)
		}
		for i := 0; i < count; i++ {
			c := cuboids[i]
			intersecton, ok := nc.intersection(c)
			if !ok {
				continue
			}
			charge := 1
			if nc.charge > 0 && c.charge > 0 || nc.charge < 0 && c.charge > 0 {
				charge = -1
			}
			intersecton.charge = charge
			cuboids = append(cuboids, intersecton)
		}
	}

	volume := 0
	for _, c := range cuboids {
		volume += c.charge * c.volume()
	}

	return volume
}
