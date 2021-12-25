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

type sea struct {
	width  int
	height int
	east   map[aoc.Vector2]bool
	south  map[aoc.Vector2]bool
}

func read() (data sea) {
	lines := aoc.ReadAllInput()

	data.width = len(lines[0])
	data.height = len(lines) - 1
	data.east = map[aoc.Vector2]bool{}
	data.south = map[aoc.Vector2]bool{}

	for y, line := range lines {
		for x, v := range line {
			switch rune(v) {
			case '>':
				data.east[aoc.NewVector2(x, y)] = true
			case 'v':
				data.south[aoc.NewVector2(x, y)] = true
			}
		}
	}

	return data
}

func process(data sea) int {
	for i := 1; ; i++ {
		movedEast := moveEast(data)
		movedSouth := moveSouth(data)
		if !(movedEast || movedSouth) {
			return i
		}
	}
}

func moveEast(data sea) bool {
	move := map[aoc.Vector2]aoc.Vector2{}
	for f := range data.east {
		np := aoc.NewVector2((f.X+1)%data.width, f.Y)
		if !data.east[np] && !data.south[np] {
			move[f] = np
		}
	}

	for from, to := range move {
		delete(data.east, from)
		data.east[to] = true
	}

	return len(move) != 0
}

func moveSouth(data sea) bool {
	move := map[aoc.Vector2]aoc.Vector2{}
	for f := range data.south {
		np := aoc.NewVector2(f.X, (f.Y+1)%data.height)
		if !data.east[np] && !data.south[np] {
			move[f] = np
		}
	}

	for from, to := range move {
		delete(data.south, from)
		data.south[to] = true
	}

	return len(move) != 0
}
