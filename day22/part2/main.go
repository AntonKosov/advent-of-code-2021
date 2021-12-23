package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/AntonKosov/advent-of-code-2021/aoc"
)

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

type operation struct {
	isOn bool
	min  aoc.Vector3
	max  aoc.Vector3
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
		min := aoc.NewVector3(aoc.StrToInt(xc[0]), aoc.StrToInt(yc[0]), aoc.StrToInt(zc[0]))
		max := aoc.NewVector3(aoc.StrToInt(xc[1]), aoc.StrToInt(yc[1]), aoc.StrToInt(zc[1]))
		op := operation{isOn: sof[0] == "on", min: min, max: max}
		data = append(data, op)
	}

	return data
}

type collapsedDimension struct {
	coordToIndex map[int]int
	indexToCoord map[int]int
}

func newCollapsedDimension(values map[int]struct{}) *collapsedDimension {
	coordToIndex := make(map[int]int, len(values))
	indexToCoord := make(map[int]int, len(values))

	vs := make([]int, 0, len(values))
	for v := range values {
		vs = append(vs, v)
	}
	sort.Ints(vs)

	for i, v := range vs {
		coordToIndex[v] = i
		indexToCoord[i] = v
	}

	return &collapsedDimension{
		coordToIndex: coordToIndex,
		indexToCoord: indexToCoord,
	}
}

func (cd *collapsedDimension) size(index int) int {
	if index < len(cd.indexToCoord)-1 {
		return cd.indexToCoord[index+1] - cd.indexToCoord[index]
	}

	return 1
}

type collapsedDimensions struct {
	x, y, z *collapsedDimension
}

func newCollapsedDimensions(operations []operation) *collapsedDimensions {
	mx := make(map[int]struct{}, len(operations))
	my := make(map[int]struct{}, len(operations))
	mz := make(map[int]struct{}, len(operations))
	addComponent := func(m map[int]struct{}, c int) {
		// if _, ok := m[c]; ok {
		// m[c-1] = struct{}{}
		// m[c+1] = struct{}{}
		m[c-1] = struct{}{}
		m[c] = struct{}{}
		m[c+1] = struct{}{}
	}
	addCoord := func(v aoc.Vector3) {
		addComponent(mx, v.X)
		addComponent(my, v.Y)
		addComponent(mz, v.Z)
	}
	for _, o := range operations {
		addCoord(o.min)
		addCoord(o.max)
	}

	// addComponent := func(m map[int]struct{}, c int, isSingle bool) {
	// 	if _, ok := m[c]; ok || isSingle {
	// 		m[c-1] = struct{}{}
	// 		m[c+1] = struct{}{}
	// 	}
	// 	m[c] = struct{}{}
	// }
	// addCoord := func(v aoc.Vector3, isXSingle, isYSingle, isZSingle bool) {
	// 	addComponent(mx, v.X, isXSingle)
	// 	addComponent(my, v.Y, isYSingle)
	// 	addComponent(mz, v.Z, isZSingle)
	// }
	// for _, o := range operations {
	// 	isXSingle := o.min.X == o.max.X
	// 	isYSingle := o.min.Y == o.max.Y
	// 	isZSingle := o.min.Z == o.max.Z
	// 	addCoord(o.min, isXSingle, isYSingle, isZSingle)
	// 	addCoord(o.max, isXSingle, isYSingle, isZSingle)
	// }

	return &collapsedDimensions{
		x: newCollapsedDimension(mx),
		y: newCollapsedDimension(my),
		z: newCollapsedDimension(mz),
	}
}

func (cd *collapsedDimensions) getIndexes(pos aoc.Vector3) aoc.Vector3 {
	return aoc.NewVector3(
		cd.x.coordToIndex[pos.X],
		cd.y.coordToIndex[pos.Y],
		cd.z.coordToIndex[pos.Z],
	)
}

func newSpace(x, y, z int) [][][]bool {
	//TODO: optimization: don't create the space in advance. It may be just a map[vector3]struct{}
	// with enabled cells
	space := make([][][]bool, x)
	for x := range space {
		syz := make([][]bool, y)
		for y := range syz {
			syz[y] = make([]bool, z)
		}
		space[x] = syz
	}

	return space
}

func process(data []operation) int {
	cd := newCollapsedDimensions(data)
	space := newSpace(len(cd.x.coordToIndex), len(cd.y.coordToIndex), len(cd.z.coordToIndex))
	// space := map[aoc.Vector3]struct{}{}

	for i, o := range data {
		si := cd.getIndexes(o.min)
		ei := cd.getIndexes(o.max)
		for x := si.X; x <= ei.X; x++ {
			for y := si.Y; y <= ei.Y; y++ {
				for z := si.Z; z <= ei.Z; z++ {
					space[x][y][z] = o.isOn
					// // TODO: optimization: create only once?
					// pos := aoc.NewVector3(x, y, z)
					// if o.isOn {
					// 	space[pos] = struct{}{}
					// } else {
					// 	delete(space, pos)
					// }
				}
			}
		}
		fmt.Printf("%v/%v\n", i+1, len(data))
	}

	sum := 0
	for x, yzv := range space {
		for y, zv := range yzv {
			for z, isOn := range zv {
				if isOn {
					volume := cd.x.size(x) * cd.y.size(y) * cd.z.size(z)
					sum += volume
					// fmt.Printf("%v,%v,%v=%v\n", x, y, z, volume)
				}
			}
		}
	}
	// for c := range space {
	// 	volume := cd.x.size(c.X) * cd.y.size(c.Y) * cd.z.size(c.Z)
	// 	sum += volume
	// }

	return sum
}

// type segment struct {
// 	from, to int
// }

// func (s segment) less(v int) bool {
// 	return v >= s.from
// }

// func (s segment) size() int {
// 	return s.to - s.from + 1
// }
//
// type segments struct {
// 	x, y, z []segment
// }
//
// func newSegments(m map[segment]struct{}) []segment {
// 	s := make([]segment, 0, len(m))
// 	for seg := range m {
// 		s = append(s, seg)
// 	}
// 	return s
// }
//
// func newXYZSegments(data []operation) *segments {
// 	sx := make(map[segment]struct{}, len(data))
// 	sy := make(map[segment]struct{}, len(data))
// 	sz := make(map[segment]struct{}, len(data))
// 	for _, o := range data {
// 		sx[segment{from: o.min.X, to: o.max.X}] = struct{}{}
// 		sy[segment{from: o.min.Y, to: o.max.Y}] = struct{}{}
// 		sz[segment{from: o.min.Z, to: o.max.Z}] = struct{}{}
// 	}
//
// 	return &segments{
// 		x: newSegments(sx),
// 		y: newSegments(sy),
// 		z: newSegments(sz),
// 	}
// }
//
// type cell struct {
// 	x, y, z segment
// }
//
// func (c cell) volume() int {
// 	return c.x.size() * c.y.size() * c.z.size()
// }
//
// func newSpace(s *segments) map[cell]bool {
// 	m := make(map[cell]bool, len(s.x)*len(s.y)*len(s.z))
// 	for _, x := range s.x {
// 		for _, y := range s.y {
// 			for _, z := range s.z {
// 				m[cell{x: x, y: y, z: z}] = false
// 			}
// 		}
// 	}
//
// 	return m
// }
//
// func process(data []operation) int {
// 	s := newXYZSegments(data)
// 	space := newSpace(s)
//
// 	for _, o := range data {
// 		min := o.min
// 		max := o.max
// 		// TODO: optimization: the segments may be sorted
// 		for _, x := range s.x {
// 			for _, y := range s.y {
// 				for _, z := range s.z {
// 					if x.wraps(min.X) && y.wraps(min.Y) && z.wraps(min.Z) {
//
// 					}
// 				}
// 			}
// 		}
// 	}
//
// 	sum := 0
// 	for x, yzv := range space {
// 		for y, zv := range yzv {
// 			for z, isOn := range zv {
// 				if isOn {
// 					volume := cd.x.size(x) * cd.y.size(y) * cd.z.size(z)
// 					sum += volume
// 					fmt.Printf("%v,%v,%v=%v\n", x, y, z, volume)
// 				}
// 			}
// 		}
// 	}
//
// 	return sum
// }
