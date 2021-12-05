package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2021/aoc"
)

type line struct {
	start aoc.Vector2
	end   aoc.Vector2
}

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func read() (data []line) {
	lines := aoc.ReadAllInput()

	for _, l := range lines {
		if l == "" {
			continue
		}
		dots := strings.Split(l, " -> ")
		start := aoc.StrToInts(dots[0], ",")
		end := aoc.StrToInts(dots[1], ",")
		if start[0] == end[0] || start[1] == end[1] {
			data = append(data, line{
				start: aoc.NewVector2(start[0], start[1]),
				end:   aoc.NewVector2(end[0], end[1]),
			})
		}
	}

	return data
}

func process(data []line) int {
	m := make(map[aoc.Vector2]int)
	intersections := 0
	for _, l := range data {
		dir := l.end.Sub(l.start).Norm()
		for cp := l.start; ; cp = cp.Add(dir) {
			m[cp]++
			if m[cp] == 2 {
				intersections++
			}
			if cp == l.end {
				break
			}
		}
	}

	return intersections
}
