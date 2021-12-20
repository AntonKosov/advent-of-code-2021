package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2021/aoc"
)

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

type state struct {
	enhancement []bool
	litPixels   map[aoc.Vector2]bool
	litRest     bool

	minX, minY, maxX, maxY int
}

func (s *state) isLit(x, y int) bool {
	if x >= s.minX && x <= s.maxX && y >= s.minY && y <= s.maxY {
		return s.litPixels[aoc.NewVector2(x, y)]
	}

	return s.litRest
}

func (s *state) index(x, y int) int {
	index := 0
	for iy := y - 1; iy <= y+1; iy++ {
		for ix := x - 1; ix <= x+1; ix++ {
			index <<= 1
			if s.isLit(ix, iy) {
				index |= 1
			}
		}
	}

	return index
}

func read() (data *state) {
	lines := aoc.ReadAllInput()

	data = &state{
		litPixels:   map[aoc.Vector2]bool{},
		enhancement: make([]bool, 0, len(lines[0])),
		minX:        0,
		minY:        0,
		maxX:        len(lines[2]) - 1,
		maxY:        len(lines) - 4,
	}

	for _, r := range lines[0] {
		data.enhancement = append(data.enhancement, r == '#')
	}

	for y := 0; y <= data.maxY; y++ {
		line := lines[y+2]
		for x, r := range line {
			if r == '#' {
				data.litPixels[aoc.NewVector2(x, y)] = true
			}
		}
	}

	return data
}

func process(data *state) int {
	for i := 0; i < 50; i++ {
		nextFrame := make(map[aoc.Vector2]bool, len(data.litPixels))

		for x := data.minX - 1; x <= data.maxX+1; x++ {
			for y := data.minY - 1; y <= data.maxY+1; y++ {
				if data.enhancement[data.index(x, y)] {
					nextFrame[aoc.NewVector2(x, y)] = true
				}
			}
		}

		data.litPixels = nextFrame
		data.minX--
		data.minY--
		data.maxX++
		data.maxY++
		if data.litRest {
			data.litRest = data.enhancement[511]
		} else {
			data.litRest = data.enhancement[0]
		}
	}

	return len(data.litPixels)
}
